/*
Licensed under the BSD 3-Clause License (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://opensource.org/licenses/BSD-3-Clause

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package core

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:webhook:path=/dragnet-webhook,mutating=true,failurePolicy=ignore,groups="*",versions="*",resources="*/*",verbs=create;update,name=dragnet.kubemod.io

// DragnetWebhookHandler is the main entrypoint to KubeMod's mutating admission webhook.
type DragnetWebhookHandler struct {
	client       client.Client
	decoder      *admission.Decoder
	log          logr.Logger
	modRuleStore *ModRuleStore
}

// NewDragnetWebhookHandler constructs a new core webhook handler.
func NewDragnetWebhookHandler(manager manager.Manager, modRuleStore *ModRuleStore, log logr.Logger) *DragnetWebhookHandler {
	return &DragnetWebhookHandler{
		client:       manager.GetClient(),
		log:          log.WithName("modrule-webhook"),
		modRuleStore: modRuleStore,
	}
}

// Handle trigers the main mutating logic.
func (h *DragnetWebhookHandler) Handle(ctx context.Context, req admission.Request) admission.Response {
	log := h.log.WithValues("request uid", req.UID, "namespace", req.Namespace, "resource", fmt.Sprintf("%v/%v", req.Resource.Resource, req.Name), "operation", req.Operation)

	patch, err := h.modRuleStore.CalculatePatch(req.Namespace, req.Object.Raw, log)

	if err != nil {
		log.Error(err, "Failed to calculate patch")
		// We don't want to fail the admission just because we cannot decode the JSON.
		return admission.Allowed("failed to calculate patch")
	}

	if patch != nil && len(patch) > 0 {
		log.Info("Applying ModRule patch", "patch", patch)
		return admission.Patched("patched ok", patch...)
	}

	return admission.Allowed("non-patched ok")
}

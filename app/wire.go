//+build wireinject

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

package app

import (
	"github.com/google/wire"
	"github.com/kubemod/kubemod/controllers"
	"github.com/kubemod/kubemod/core"
	"github.com/kubemod/kubemod/expressions"
	"k8s.io/apimachinery/pkg/runtime"
)

type EnableLeaderElection bool
type EnableDevModeLog bool

func InitializeKubeModApp(scheme *runtime.Scheme, metricsAddr string, enableLeaderElection EnableLeaderElection, enableDevModeLog EnableDevModeLog) (*KubeModApp, error) {
	wire.Build(
		expressions.NewJSONPathLanguage,
		core.NewModRuleStoreItemFactory,
		core.NewModRuleStore,
		core.NewDragnetWebhookHandler,
		controllers.NewModRuleReconciler,
		NewControllerManager,
		NewLogger,
		NewKubeModApp,
	)
	return nil, nil
}

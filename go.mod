module github.com/kubemod/kubemod

go 1.13

replace github.com/evanphx/json-patch/v5 v5.1.0 => github.com/vassilvk/json-patch/v5 v5.2.0-beta.1

require (
	github.com/PaesslerAG/gval v1.0.1
	github.com/PaesslerAG/jsonpath v0.1.1
	github.com/evanphx/json-patch/v5 v5.1.0
	github.com/go-logr/logr v0.1.0
	github.com/google/wire v0.4.0
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	github.com/segmentio/ksuid v1.0.3
	go.uber.org/zap v1.10.0
	gomodules.xyz/jsonpatch/v2 v2.0.1
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v0.18.6
	sigs.k8s.io/controller-runtime v0.6.2
	sigs.k8s.io/yaml v1.2.0
)

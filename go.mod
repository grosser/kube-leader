module github.com/grosser/kube-leader

go 1.14

require (
	github.com/go-logr/logr v0.1.0
	github.com/go-logr/zapr v0.1.1
	github.com/operator-framework/operator-sdk v0.17.1
	go.uber.org/zap v1.14.1
	k8s.io/api v0.17.4
	k8s.io/apimachinery v0.17.4
	sigs.k8s.io/controller-runtime v0.5.2
)

// https://github.com/golang/go/issues/33558
replace (
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190718183610-8e956561bbf5
)

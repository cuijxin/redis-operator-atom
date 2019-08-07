package k8s_test

import "k8s.io/apimachinery/pkg/runtime/schema"

var (
	rbGroup = schema.GroupVersionResource{Group: "rbac.authorization.k8s.io", Version: "v1", Resource: "rolebindings"}
)

package k8s

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"redis-operator-atom/log"
)

// RBAC is the service that knows how to interact with k8s to manage RBAC related resources.
type RBAC interface {
	GetClusterRole(name string) (*rbacv1.ClusterRole, error)
}

// NamespaceService is the Namespace service implementation using API calls to kubernetes.
type RBACService struct {
	kubeClient kubernetes.Interface
	logger     log.Logger
}

// NewRBACService returns a new RBAC Kubernetes.
func NewRBACService(kubeClient kubernetes.Interface, logger log.Logger) *RBACService {
	logger = logger.With("service", "k8s.rbac")
	return &RBACService{
		kubeClient: kubeClient,
		logger:     logger,
	}
}

// GetClusterRole 获取ClusterRole对象.
func (r *RBACService) GetClusterRole(name string) (*rbacv1.ClusterRole, error) {
	return r.kubeClient.RbacV1().ClusterRoles().Get(name, metav1.GetOptions{})
}

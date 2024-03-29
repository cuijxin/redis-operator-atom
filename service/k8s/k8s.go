package k8s

import (
	redisfailoverclientset "github.com/cuijxin/redis-operator-atom/client/k8s/clientset/versioned"
	"github.com/cuijxin/redis-operator-atom/log"
	apiextensionscli "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
)

// Service is the K8s service entrypoint.
type Services interface {
	CRD
	ConfigMap
	Pod
	PodDisruptionBudget
	RedisFailover
	Service
	RBAC
	Deployment
	StatefulSet
}

type services struct {
	CRD
	ConfigMap
	Pod
	PodDisruptionBudget
	RedisFailover
	Service
	RBAC
	Deployment
	StatefulSet
}

// New returns a new Kubernetes service.
func New(kubecli kubernetes.Interface, crdcli redisfailoverclientset.Interface, apiextcli apiextensionscli.Interface, logger log.Logger) Services {
	return &services{
		CRD:                 NewCRDService(apiextcli, logger),
		ConfigMap:           NewConfigMapService(kubecli, logger),
		Pod:                 NewPodService(kubecli, logger),
		PodDisruptionBudget: NewPodDisruptionBudgetService(kubecli, logger),
		RedisFailover:       NewRedisFailoverService(crdcli, logger),
		Service:             NewServiceService(kubecli, logger),
		RBAC:                NewRBACService(kubecli, logger),
		Deployment:          NewDeploymentService(kubecli, logger),
		StatefulSet:         NewStatefulSetService(kubecli, logger),
	}
}

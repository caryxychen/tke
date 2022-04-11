package deletion

import (
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiauthzv1 "tkestack.io/tke/api/authz/v1"
	clientset "tkestack.io/tke/api/client/clientset/versioned"
	"tkestack.io/tke/pkg/authz/provider"
	"tkestack.io/tke/pkg/util/log"
)

type ClusterPolicyBindingDeleter interface {
	Delete(ctx context.Context, cpb *apiauthzv1.ClusterPolicyBinding, provider provider.Provider) error
}

func New(client clientset.Interface) ClusterPolicyBindingDeleter {
	return &clusterPolicyBindingResourcesDeleter{
		client: client,
	}
}

type clusterPolicyBindingResourcesDeleter struct {
	client clientset.Interface
}

func (c *clusterPolicyBindingResourcesDeleter) Delete(ctx context.Context, cpb *apiauthzv1.ClusterPolicyBinding, provider provider.Provider) error {
	// 删除集群中对应的资源
	if err := provider.DeleteClusterPolicyBindingResources(cpb); err != nil {
		log.Warnf("Unable to finalize clusterpolicybinding '%s/%s', err: %v", cpb.Namespace, cpb.Name, err)
		return err
	}
	policyFinalize := apiauthzv1.ClusterPolicyBinding{}
	policyFinalize.ObjectMeta = cpb.ObjectMeta
	policyFinalize.Finalizers = []string{}
	if err := c.client.AuthzV1().RESTClient().Put().Resource("clusterpolicybindings").
		Namespace(cpb.Namespace).
		Name(cpb.Name).
		SubResource("finalize").
		Body(&policyFinalize).
		Do(context.Background()).
		Into(&policyFinalize); err != nil {
		log.Warnf("Unable to finalize clusterpolicybinding '%s/%s', err: %v", cpb.Namespace, cpb.Name, err)
		return err
	}
	return c.client.AuthzV1().ClusterPolicyBindings(cpb.Namespace).Delete(ctx, cpb.Name, metav1.DeleteOptions{})
}

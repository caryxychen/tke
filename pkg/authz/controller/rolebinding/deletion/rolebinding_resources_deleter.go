package deletion

import (
	"fmt"
	"golang.org/x/net/context"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiauthzv1 "tkestack.io/tke/api/authz/v1"
	clientset "tkestack.io/tke/api/client/clientset/versioned"
	"tkestack.io/tke/pkg/util/log"
)

type RoleBindingDeleter interface {
	Delete(ctx context.Context, rb *apiauthzv1.RoleBinding) error
}

func New(client clientset.Interface) RoleBindingDeleter {
	return &roleBindingResourcesDeleter{
		client: client,
	}
}

type roleBindingResourcesDeleter struct {
	client clientset.Interface
}

func (c *roleBindingResourcesDeleter) Delete(ctx context.Context, rb *apiauthzv1.RoleBinding) error {
	generatedName := fmt.Sprintf("role-%s", rb.Name)
	err := c.client.AuthzV1().Policies(rb.Namespace).Delete(ctx, generatedName, metav1.DeleteOptions{})
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Unable to delete generated policy '%s/%s'", rb.Namespace, generatedName)
			return err
		}
	}
	roleBindingFinalize := apiauthzv1.RoleBinding{}
	roleBindingFinalize.ObjectMeta = rb.ObjectMeta
	roleBindingFinalize.Finalizers = []string{}
	if err := c.client.AuthzV1().RESTClient().Put().Resource("rolebindings").
		Namespace(rb.Namespace).
		Name(rb.Name).
		SubResource("finalize").
		Body(&roleBindingFinalize).
		Do(context.Background()).
		Into(&roleBindingFinalize); err != nil {
		log.Warnf("Unable to finalize rolebinding '%s/%s', err: %v", rb.Namespace, rb.Name, err)
		return err
	}
	return c.client.AuthzV1().RoleBindings(rb.Namespace).Delete(ctx, rb.Name, metav1.DeleteOptions{})
}

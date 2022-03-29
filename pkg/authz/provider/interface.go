package provider

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
	authzv1 "tkestack.io/tke/api/authz/v1"
	platformversionedclient "tkestack.io/tke/api/client/clientset/versioned/typed/platform/v1"
	platformv1 "tkestack.io/tke/pkg/platform/types/v1"
)

type Provider interface {
	Name() string
	ReconcileRoleTemplate(rt *authzv1.RoleTemplate, platformClient platformversionedclient.PlatformV1Interface) error
	GetClusterRoleBindingSubject(platformUser string, platformGroup string, cluster *platformv1.Cluster) (*rbacv1.Subject, error)
	DispatchClusterRoleBindings(platformClient platformversionedclient.PlatformV1Interface, rt *authzv1.RoleTemplate, crtb *authzv1.ClusterRoleTemplateBinding, clusterSubjects map[string]*rbacv1.Subject) error
}

var _ Provider = &DelegateProvider{}

type DelegateProvider struct {
	ProviderName string
}

func (p *DelegateProvider) DispatchClusterRoleBindings(platformClient platformversionedclient.PlatformV1Interface, rt *authzv1.RoleTemplate, crtb *authzv1.ClusterRoleTemplateBinding, clusterSubjects map[string]*rbacv1.Subject) error {
	return nil
}

func (p *DelegateProvider) GetClusterRoleBindingSubject(platformUser string, platformGroup string, cluster *platformv1.Cluster) (*rbacv1.Subject, error) {
	_, err := cluster.RESTConfig()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (p *DelegateProvider) ReconcileRoleTemplate(rt *authzv1.RoleTemplate, platformClient platformversionedclient.PlatformV1Interface) error {
	rt.Status.Phase = authzv1.Succeeded
	rt.Status.LastTransitionTime = metav1.Time{Time: time.Now()}
	return nil
}

func (p *DelegateProvider) Name() string {
	if p.ProviderName == "" {
		return "unknown"
	}
	return p.ProviderName
}

package provider

import (
	"context"
	rbacv1 "k8s.io/api/rbac/v1"
	authzv1 "tkestack.io/tke/api/authz/v1"
	platformversionedclient "tkestack.io/tke/api/client/clientset/versioned/typed/platform/v1"
	platformv1 "tkestack.io/tke/pkg/platform/types/v1"
)

type Provider interface {
	Name() string
	InitContext(param interface{}) context.Context
	GetSubject(ctx context.Context, userName string, cluster *platformv1.Cluster) (*rbacv1.Subject, error)
	DispatchMultiClusterRoleBinding(ctx context.Context, platformClient platformversionedclient.PlatformV1Interface, mcrb *authzv1.MultiClusterRoleBinding, rules []rbacv1.PolicyRule, clusterSubjects map[string]*rbacv1.Subject) error
	RenderClusterPolicyBinding(ctx context.Context, mcrb *authzv1.MultiClusterRoleBinding) error
	DeleteMultiClusterRoleBindingResources(ctx context.Context, platformClient platformversionedclient.PlatformV1Interface, mcrb *authzv1.MultiClusterRoleBinding) error
}

var _ Provider = &DelegateProvider{}

type DelegateProvider struct {
	ProviderName string
}

func (p *DelegateProvider) Name() string {
	if p.ProviderName == "" {
		return "unknown"
	}
	return p.ProviderName
}

func (p *DelegateProvider) InitContext(param interface{}) context.Context {
	return context.Background()
}

func (p *DelegateProvider) GetSubject(ctx context.Context, platformUser string, cluster *platformv1.Cluster) (*rbacv1.Subject, error) {
	_, err := cluster.RESTConfig()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (p *DelegateProvider) DispatchMultiClusterRoleBinding(ctx context.Context, platformClient platformversionedclient.PlatformV1Interface, mcrb *authzv1.MultiClusterRoleBinding, rules []rbacv1.PolicyRule, clusterSubjects map[string]*rbacv1.Subject) error {
	return nil
}

func (p *DelegateProvider) RenderClusterPolicyBinding(ctx context.Context, mcrb *authzv1.MultiClusterRoleBinding) error {
	return nil
}

func (p *DelegateProvider) DeleteMultiClusterRoleBindingResources(ctx context.Context, platformClient platformversionedclient.PlatformV1Interface, mcrb *authzv1.MultiClusterRoleBinding) error {
	return nil
}

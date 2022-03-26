package provider

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
	authzv1 "tkestack.io/tke/api/authz/v1"
	platformversionedclient "tkestack.io/tke/api/client/clientset/versioned/typed/platform/v1"
)

type Provider interface {
	Name() string
	ReconcileRoleTemplate(*authzv1.RoleTemplate, platformversionedclient.PlatformV1Interface) error
}

var _ Provider = &DelegateProvider{}

type DelegateProvider struct {
	ProviderName string
}

func (p *DelegateProvider) ReconcileRoleTemplate(template *authzv1.RoleTemplate, v1Interface platformversionedclient.PlatformV1Interface) error {
	template.Status.Phase = authzv1.Succeeded
	template.Status.LastTransitionTime = metav1.Time{Time: time.Now()}
	return nil
}

func (p *DelegateProvider) Name() string {
	if p.ProviderName == "" {
		return "unknown"
	}
	return p.ProviderName
}

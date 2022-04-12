package policyrolecache

import (
	"fmt"
	"sync"
	apiauthzv1 "tkestack.io/tke/api/authz/v1"
)

type PolicyRoleCache interface {
	GetRolesByPolicy(policyName string) Set
	PutByRole(role *apiauthzv1.Role)
	DeleteRole(role *apiauthzv1.Role)
}

type Set map[string]interface{}

type policyRoleCache struct {
	rw sync.RWMutex
	// key: policyName
	// value: roleName set
	store map[string]Set
}

var Cache = &policyRoleCache{store: map[string]Set{}}

func (c *policyRoleCache) GetRolesByPolicy(policy *apiauthzv1.Policy) Set {
	c.rw.RLocker().Lock()
	defer c.rw.RLocker().Unlock()
	policyName := fmt.Sprintf("%s/%s", policy.Namespace, policy.Name)
	return c.store[policyName]
}

func (c *policyRoleCache) UpdateByRole(role *apiauthzv1.Role) {
	c.rw.Lock()
	defer c.rw.Unlock()
	roleName := fmt.Sprintf("%s/%s", role.Namespace, role.Name)
	policies := role.Policies
	for _, roleSet := range c.store {
		delete(roleSet, roleName)
	}
	for _, policy := range policies {
		set := c.store[policy]
		if set == nil {
			set = map[string]interface{}{}
		}
		set[roleName] = nil
		c.store[policy] = set
	}
}

func (c *policyRoleCache) PutByRole(role *apiauthzv1.Role) {
	c.rw.Lock()
	defer c.rw.Unlock()
	roleName := fmt.Sprintf("%s/%s", role.Namespace, role.Name)
	policies := role.Policies
	for _, policy := range policies {
		set := c.store[policy]
		if set == nil {
			set = map[string]interface{}{}
		}
		set[roleName] = nil
		c.store[policy] = set
	}
}

func (c *policyRoleCache) DeleteRole(role *apiauthzv1.Role) {
	c.rw.Lock()
	defer c.rw.Unlock()
	roleName := fmt.Sprintf("%s/%s", role.Namespace, role.Name)
	for _, roleSet := range c.store {
		delete(roleSet, roleName)
	}
}

func (c *policyRoleCache) DeletePolicy(policy *apiauthzv1.Policy) {
	c.rw.Lock()
	defer c.rw.Unlock()
	policyName := fmt.Sprintf("%s/%s", policy.Namespace, policy.Name)
	delete(c.store, policyName)
}

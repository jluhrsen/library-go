// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
)

// ClusterResourceQuotaSpecApplyConfiguration represents a declarative configuration of the ClusterResourceQuotaSpec type for use
// with apply.
type ClusterResourceQuotaSpecApplyConfiguration struct {
	Selector *ClusterResourceQuotaSelectorApplyConfiguration `json:"selector,omitempty"`
	Quota    *corev1.ResourceQuotaSpec                       `json:"quota,omitempty"`
}

// ClusterResourceQuotaSpecApplyConfiguration constructs a declarative configuration of the ClusterResourceQuotaSpec type for use with
// apply.
func ClusterResourceQuotaSpec() *ClusterResourceQuotaSpecApplyConfiguration {
	return &ClusterResourceQuotaSpecApplyConfiguration{}
}

// WithSelector sets the Selector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Selector field is set to the value of the last call.
func (b *ClusterResourceQuotaSpecApplyConfiguration) WithSelector(value *ClusterResourceQuotaSelectorApplyConfiguration) *ClusterResourceQuotaSpecApplyConfiguration {
	b.Selector = value
	return b
}

// WithQuota sets the Quota field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Quota field is set to the value of the last call.
func (b *ClusterResourceQuotaSpecApplyConfiguration) WithQuota(value corev1.ResourceQuotaSpec) *ClusterResourceQuotaSpecApplyConfiguration {
	b.Quota = &value
	return b
}

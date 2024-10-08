// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// PowerVSServiceEndpointApplyConfiguration represents a declarative configuration of the PowerVSServiceEndpoint type for use
// with apply.
type PowerVSServiceEndpointApplyConfiguration struct {
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
}

// PowerVSServiceEndpointApplyConfiguration constructs a declarative configuration of the PowerVSServiceEndpoint type for use with
// apply.
func PowerVSServiceEndpoint() *PowerVSServiceEndpointApplyConfiguration {
	return &PowerVSServiceEndpointApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *PowerVSServiceEndpointApplyConfiguration) WithName(value string) *PowerVSServiceEndpointApplyConfiguration {
	b.Name = &value
	return b
}

// WithURL sets the URL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the URL field is set to the value of the last call.
func (b *PowerVSServiceEndpointApplyConfiguration) WithURL(value string) *PowerVSServiceEndpointApplyConfiguration {
	b.URL = &value
	return b
}

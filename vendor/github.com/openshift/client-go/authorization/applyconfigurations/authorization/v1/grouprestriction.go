// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// GroupRestrictionApplyConfiguration represents a declarative configuration of the GroupRestriction type for use
// with apply.
type GroupRestrictionApplyConfiguration struct {
	Groups    []string                                 `json:"groups,omitempty"`
	Selectors []metav1.LabelSelectorApplyConfiguration `json:"labels,omitempty"`
}

// GroupRestrictionApplyConfiguration constructs a declarative configuration of the GroupRestriction type for use with
// apply.
func GroupRestriction() *GroupRestrictionApplyConfiguration {
	return &GroupRestrictionApplyConfiguration{}
}

// WithGroups adds the given value to the Groups field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Groups field.
func (b *GroupRestrictionApplyConfiguration) WithGroups(values ...string) *GroupRestrictionApplyConfiguration {
	for i := range values {
		b.Groups = append(b.Groups, values[i])
	}
	return b
}

// WithSelectors adds the given value to the Selectors field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Selectors field.
func (b *GroupRestrictionApplyConfiguration) WithSelectors(values ...*metav1.LabelSelectorApplyConfiguration) *GroupRestrictionApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSelectors")
		}
		b.Selectors = append(b.Selectors, *values[i])
	}
	return b
}

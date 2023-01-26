//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionIncidentObservation) DeepCopyInto(out *ActionIncidentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionIncidentObservation.
func (in *ActionIncidentObservation) DeepCopy() *ActionIncidentObservation {
	if in == nil {
		return nil
	}
	out := new(ActionIncidentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionIncidentParameters) DeepCopyInto(out *ActionIncidentParameters) {
	*out = *in
	if in.Classification != nil {
		in, out := &in.Classification, &out.Classification
		*out = new(string)
		**out = **in
	}
	if in.ClassificationComment != nil {
		in, out := &in.ClassificationComment, &out.ClassificationComment
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
	if in.OwnerID != nil {
		in, out := &in.OwnerID, &out.OwnerID
		*out = new(string)
		**out = **in
	}
	if in.Severity != nil {
		in, out := &in.Severity, &out.Severity
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionIncidentParameters.
func (in *ActionIncidentParameters) DeepCopy() *ActionIncidentParameters {
	if in == nil {
		return nil
	}
	out := new(ActionIncidentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionPlaybookObservation) DeepCopyInto(out *ActionPlaybookObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionPlaybookObservation.
func (in *ActionPlaybookObservation) DeepCopy() *ActionPlaybookObservation {
	if in == nil {
		return nil
	}
	out := new(ActionPlaybookObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionPlaybookParameters) DeepCopyInto(out *ActionPlaybookParameters) {
	*out = *in
	if in.LogicAppID != nil {
		in, out := &in.LogicAppID, &out.LogicAppID
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionPlaybookParameters.
func (in *ActionPlaybookParameters) DeepCopy() *ActionPlaybookParameters {
	if in == nil {
		return nil
	}
	out := new(ActionPlaybookParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionObservation) DeepCopyInto(out *ConditionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionObservation.
func (in *ConditionObservation) DeepCopy() *ConditionObservation {
	if in == nil {
		return nil
	}
	out := new(ConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionParameters) DeepCopyInto(out *ConditionParameters) {
	*out = *in
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Property != nil {
		in, out := &in.Property, &out.Property
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionParameters.
func (in *ConditionParameters) DeepCopy() *ConditionParameters {
	if in == nil {
		return nil
	}
	out := new(ConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SentinelAutomationRule) DeepCopyInto(out *SentinelAutomationRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SentinelAutomationRule.
func (in *SentinelAutomationRule) DeepCopy() *SentinelAutomationRule {
	if in == nil {
		return nil
	}
	out := new(SentinelAutomationRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SentinelAutomationRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SentinelAutomationRuleList) DeepCopyInto(out *SentinelAutomationRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SentinelAutomationRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SentinelAutomationRuleList.
func (in *SentinelAutomationRuleList) DeepCopy() *SentinelAutomationRuleList {
	if in == nil {
		return nil
	}
	out := new(SentinelAutomationRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SentinelAutomationRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SentinelAutomationRuleObservation) DeepCopyInto(out *SentinelAutomationRuleObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SentinelAutomationRuleObservation.
func (in *SentinelAutomationRuleObservation) DeepCopy() *SentinelAutomationRuleObservation {
	if in == nil {
		return nil
	}
	out := new(SentinelAutomationRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SentinelAutomationRuleParameters) DeepCopyInto(out *SentinelAutomationRuleParameters) {
	*out = *in
	if in.ActionIncident != nil {
		in, out := &in.ActionIncident, &out.ActionIncident
		*out = make([]ActionIncidentParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ActionPlaybook != nil {
		in, out := &in.ActionPlaybook, &out.ActionPlaybook
		*out = make([]ActionPlaybookParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = make([]ConditionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Expiration != nil {
		in, out := &in.Expiration, &out.Expiration
		*out = new(string)
		**out = **in
	}
	if in.LogAnalyticsWorkspaceID != nil {
		in, out := &in.LogAnalyticsWorkspaceID, &out.LogAnalyticsWorkspaceID
		*out = new(string)
		**out = **in
	}
	if in.LogAnalyticsWorkspaceIDRef != nil {
		in, out := &in.LogAnalyticsWorkspaceIDRef, &out.LogAnalyticsWorkspaceIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.LogAnalyticsWorkspaceIDSelector != nil {
		in, out := &in.LogAnalyticsWorkspaceIDSelector, &out.LogAnalyticsWorkspaceIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SentinelAutomationRuleParameters.
func (in *SentinelAutomationRuleParameters) DeepCopy() *SentinelAutomationRuleParameters {
	if in == nil {
		return nil
	}
	out := new(SentinelAutomationRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SentinelAutomationRuleSpec) DeepCopyInto(out *SentinelAutomationRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SentinelAutomationRuleSpec.
func (in *SentinelAutomationRuleSpec) DeepCopy() *SentinelAutomationRuleSpec {
	if in == nil {
		return nil
	}
	out := new(SentinelAutomationRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SentinelAutomationRuleStatus) DeepCopyInto(out *SentinelAutomationRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SentinelAutomationRuleStatus.
func (in *SentinelAutomationRuleStatus) DeepCopy() *SentinelAutomationRuleStatus {
	if in == nil {
		return nil
	}
	out := new(SentinelAutomationRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SentinelWatchlist) DeepCopyInto(out *SentinelWatchlist) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SentinelWatchlist.
func (in *SentinelWatchlist) DeepCopy() *SentinelWatchlist {
	if in == nil {
		return nil
	}
	out := new(SentinelWatchlist)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SentinelWatchlist) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SentinelWatchlistList) DeepCopyInto(out *SentinelWatchlistList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SentinelWatchlist, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SentinelWatchlistList.
func (in *SentinelWatchlistList) DeepCopy() *SentinelWatchlistList {
	if in == nil {
		return nil
	}
	out := new(SentinelWatchlistList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SentinelWatchlistList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SentinelWatchlistObservation) DeepCopyInto(out *SentinelWatchlistObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SentinelWatchlistObservation.
func (in *SentinelWatchlistObservation) DeepCopy() *SentinelWatchlistObservation {
	if in == nil {
		return nil
	}
	out := new(SentinelWatchlistObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SentinelWatchlistParameters) DeepCopyInto(out *SentinelWatchlistParameters) {
	*out = *in
	if in.DefaultDuration != nil {
		in, out := &in.DefaultDuration, &out.DefaultDuration
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ItemSearchKey != nil {
		in, out := &in.ItemSearchKey, &out.ItemSearchKey
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.LogAnalyticsWorkspaceID != nil {
		in, out := &in.LogAnalyticsWorkspaceID, &out.LogAnalyticsWorkspaceID
		*out = new(string)
		**out = **in
	}
	if in.LogAnalyticsWorkspaceIDRef != nil {
		in, out := &in.LogAnalyticsWorkspaceIDRef, &out.LogAnalyticsWorkspaceIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.LogAnalyticsWorkspaceIDSelector != nil {
		in, out := &in.LogAnalyticsWorkspaceIDSelector, &out.LogAnalyticsWorkspaceIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SentinelWatchlistParameters.
func (in *SentinelWatchlistParameters) DeepCopy() *SentinelWatchlistParameters {
	if in == nil {
		return nil
	}
	out := new(SentinelWatchlistParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SentinelWatchlistSpec) DeepCopyInto(out *SentinelWatchlistSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SentinelWatchlistSpec.
func (in *SentinelWatchlistSpec) DeepCopy() *SentinelWatchlistSpec {
	if in == nil {
		return nil
	}
	out := new(SentinelWatchlistSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SentinelWatchlistStatus) DeepCopyInto(out *SentinelWatchlistStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SentinelWatchlistStatus.
func (in *SentinelWatchlistStatus) DeepCopy() *SentinelWatchlistStatus {
	if in == nil {
		return nil
	}
	out := new(SentinelWatchlistStatus)
	in.DeepCopyInto(out)
	return out
}

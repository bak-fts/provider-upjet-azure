// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SentinelDataConnectorIOTInitParameters struct {

	// The ID of the subscription that this Iot Data Connector connects to. Changing this forces a new Iot Data Connector to be created.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

type SentinelDataConnectorIOTObservation struct {

	// The ID of the Iot Data Connector.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Log Analytics Workspace that this Iot Data Connector resides in. Changing this forces a new Iot Data Connector to be created.
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// The ID of the subscription that this Iot Data Connector connects to. Changing this forces a new Iot Data Connector to be created.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

type SentinelDataConnectorIOTParameters struct {

	// The ID of the Log Analytics Workspace that this Iot Data Connector resides in. Changing this forces a new Iot Data Connector to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/securityinsights/v1beta1.SentinelLogAnalyticsWorkspaceOnboarding
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("workspace_id",false)
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// Reference to a SentinelLogAnalyticsWorkspaceOnboarding in securityinsights to populate logAnalyticsWorkspaceId.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceIDRef *v1.Reference `json:"logAnalyticsWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a SentinelLogAnalyticsWorkspaceOnboarding in securityinsights to populate logAnalyticsWorkspaceId.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceIDSelector *v1.Selector `json:"logAnalyticsWorkspaceIdSelector,omitempty" tf:"-"`

	// The ID of the subscription that this Iot Data Connector connects to. Changing this forces a new Iot Data Connector to be created.
	// +kubebuilder:validation:Optional
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

// SentinelDataConnectorIOTSpec defines the desired state of SentinelDataConnectorIOT
type SentinelDataConnectorIOTSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SentinelDataConnectorIOTParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SentinelDataConnectorIOTInitParameters `json:"initProvider,omitempty"`
}

// SentinelDataConnectorIOTStatus defines the observed state of SentinelDataConnectorIOT.
type SentinelDataConnectorIOTStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SentinelDataConnectorIOTObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelDataConnectorIOT is the Schema for the SentinelDataConnectorIOTs API. Manages an Iot Data Connector.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SentinelDataConnectorIOT struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SentinelDataConnectorIOTSpec   `json:"spec"`
	Status            SentinelDataConnectorIOTStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelDataConnectorIOTList contains a list of SentinelDataConnectorIOTs
type SentinelDataConnectorIOTList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SentinelDataConnectorIOT `json:"items"`
}

// Repository type metadata.
var (
	SentinelDataConnectorIOT_Kind             = "SentinelDataConnectorIOT"
	SentinelDataConnectorIOT_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SentinelDataConnectorIOT_Kind}.String()
	SentinelDataConnectorIOT_KindAPIVersion   = SentinelDataConnectorIOT_Kind + "." + CRDGroupVersion.String()
	SentinelDataConnectorIOT_GroupVersionKind = CRDGroupVersion.WithKind(SentinelDataConnectorIOT_Kind)
)

func init() {
	SchemeBuilder.Register(&SentinelDataConnectorIOT{}, &SentinelDataConnectorIOTList{})
}

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

type SpacecraftInitParameters struct {

	// A links block as defined below. Changing this forces a new resource to be created.
	Links []SpacecraftLinksInitParameters `json:"links,omitempty" tf:"links,omitempty"`

	// The location where the Spacecraft exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// NORAD ID of the Spacecraft.
	NoradID *string `json:"noradId,omitempty" tf:"norad_id,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Title of the two line elements (TLE).
	TitleLine *string `json:"titleLine,omitempty" tf:"title_line,omitempty"`

	// A list of the two line elements (TLE), the first string being the first of the TLE, the second string being the second line of the TLE. Changing this forces a new resource to be created.
	TwoLineElements []*string `json:"twoLineElements,omitempty" tf:"two_line_elements,omitempty"`
}

type SpacecraftLinksInitParameters struct {

	// Bandwidth in Mhz.
	BandwidthMhz *float64 `json:"bandwidthMhz,omitempty" tf:"bandwidth_mhz,omitempty"`

	// Center frequency in Mhz.
	CenterFrequencyMhz *float64 `json:"centerFrequencyMhz,omitempty" tf:"center_frequency_mhz,omitempty"`

	// Direction if the communication. Possible values are Uplink and Downlink.
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Name of the link.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Polarization. Possible values are RHCP, LHCP, linearVertical and linearHorizontal.
	Polarization *string `json:"polarization,omitempty" tf:"polarization,omitempty"`
}

type SpacecraftLinksObservation struct {

	// Bandwidth in Mhz.
	BandwidthMhz *float64 `json:"bandwidthMhz,omitempty" tf:"bandwidth_mhz,omitempty"`

	// Center frequency in Mhz.
	CenterFrequencyMhz *float64 `json:"centerFrequencyMhz,omitempty" tf:"center_frequency_mhz,omitempty"`

	// Direction if the communication. Possible values are Uplink and Downlink.
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Name of the link.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Polarization. Possible values are RHCP, LHCP, linearVertical and linearHorizontal.
	Polarization *string `json:"polarization,omitempty" tf:"polarization,omitempty"`
}

type SpacecraftLinksParameters struct {

	// Bandwidth in Mhz.
	// +kubebuilder:validation:Optional
	BandwidthMhz *float64 `json:"bandwidthMhz" tf:"bandwidth_mhz,omitempty"`

	// Center frequency in Mhz.
	// +kubebuilder:validation:Optional
	CenterFrequencyMhz *float64 `json:"centerFrequencyMhz" tf:"center_frequency_mhz,omitempty"`

	// Direction if the communication. Possible values are Uplink and Downlink.
	// +kubebuilder:validation:Optional
	Direction *string `json:"direction" tf:"direction,omitempty"`

	// Name of the link.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Polarization. Possible values are RHCP, LHCP, linearVertical and linearHorizontal.
	// +kubebuilder:validation:Optional
	Polarization *string `json:"polarization" tf:"polarization,omitempty"`
}

type SpacecraftObservation struct {

	// The ID of the Spacecraft.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A links block as defined below. Changing this forces a new resource to be created.
	Links []SpacecraftLinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	// The location where the Spacecraft exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// NORAD ID of the Spacecraft.
	NoradID *string `json:"noradId,omitempty" tf:"norad_id,omitempty"`

	// The name of the Resource Group where the Spacecraft exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Title of the two line elements (TLE).
	TitleLine *string `json:"titleLine,omitempty" tf:"title_line,omitempty"`

	// A list of the two line elements (TLE), the first string being the first of the TLE, the second string being the second line of the TLE. Changing this forces a new resource to be created.
	TwoLineElements []*string `json:"twoLineElements,omitempty" tf:"two_line_elements,omitempty"`
}

type SpacecraftParameters struct {

	// A links block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Links []SpacecraftLinksParameters `json:"links,omitempty" tf:"links,omitempty"`

	// The location where the Spacecraft exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// NORAD ID of the Spacecraft.
	// +kubebuilder:validation:Optional
	NoradID *string `json:"noradId,omitempty" tf:"norad_id,omitempty"`

	// The name of the Resource Group where the Spacecraft exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Title of the two line elements (TLE).
	// +kubebuilder:validation:Optional
	TitleLine *string `json:"titleLine,omitempty" tf:"title_line,omitempty"`

	// A list of the two line elements (TLE), the first string being the first of the TLE, the second string being the second line of the TLE. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	TwoLineElements []*string `json:"twoLineElements,omitempty" tf:"two_line_elements,omitempty"`
}

// SpacecraftSpec defines the desired state of Spacecraft
type SpacecraftSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpacecraftParameters `json:"forProvider"`
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
	InitProvider SpacecraftInitParameters `json:"initProvider,omitempty"`
}

// SpacecraftStatus defines the observed state of Spacecraft.
type SpacecraftStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpacecraftObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Spacecraft is the Schema for the Spacecrafts API. Manages a Spacecraft resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Spacecraft struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.links) || (has(self.initProvider) && has(self.initProvider.links))",message="spec.forProvider.links is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.noradId) || (has(self.initProvider) && has(self.initProvider.noradId))",message="spec.forProvider.noradId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.titleLine) || (has(self.initProvider) && has(self.initProvider.titleLine))",message="spec.forProvider.titleLine is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.twoLineElements) || (has(self.initProvider) && has(self.initProvider.twoLineElements))",message="spec.forProvider.twoLineElements is a required parameter"
	Spec   SpacecraftSpec   `json:"spec"`
	Status SpacecraftStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpacecraftList contains a list of Spacecrafts
type SpacecraftList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Spacecraft `json:"items"`
}

// Repository type metadata.
var (
	Spacecraft_Kind             = "Spacecraft"
	Spacecraft_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Spacecraft_Kind}.String()
	Spacecraft_KindAPIVersion   = Spacecraft_Kind + "." + CRDGroupVersion.String()
	Spacecraft_GroupVersionKind = CRDGroupVersion.WithKind(Spacecraft_Kind)
)

func init() {
	SchemeBuilder.Register(&Spacecraft{}, &SpacecraftList{})
}

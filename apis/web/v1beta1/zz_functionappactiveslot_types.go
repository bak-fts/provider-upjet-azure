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

type FunctionAppActiveSlotInitParameters struct {

	// The swap action should overwrite the Production slot's network configuration with the configuration from this slot. Defaults to true. Changing this forces a new resource to be created.
	// The swap action should overwrite the Production slot's network configuration with the configuration from this slot. Defaults to `true`.
	OverwriteNetworkConfig *bool `json:"overwriteNetworkConfig,omitempty" tf:"overwrite_network_config,omitempty"`
}

type FunctionAppActiveSlotObservation struct {

	// The ID of the Function App Active Slot
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The timestamp of the last successful swap with Production
	// The timestamp of the last successful swap with `Production`
	LastSuccessfulSwap *string `json:"lastSuccessfulSwap,omitempty" tf:"last_successful_swap,omitempty"`

	// The swap action should overwrite the Production slot's network configuration with the configuration from this slot. Defaults to true. Changing this forces a new resource to be created.
	// The swap action should overwrite the Production slot's network configuration with the configuration from this slot. Defaults to `true`.
	OverwriteNetworkConfig *bool `json:"overwriteNetworkConfig,omitempty" tf:"overwrite_network_config,omitempty"`

	// The ID of the Slot to swap with Production.
	// The ID of the Slot to swap with `Production`.
	SlotID *string `json:"slotId,omitempty" tf:"slot_id,omitempty"`
}

type FunctionAppActiveSlotParameters struct {

	// The swap action should overwrite the Production slot's network configuration with the configuration from this slot. Defaults to true. Changing this forces a new resource to be created.
	// The swap action should overwrite the Production slot's network configuration with the configuration from this slot. Defaults to `true`.
	// +kubebuilder:validation:Optional
	OverwriteNetworkConfig *bool `json:"overwriteNetworkConfig,omitempty" tf:"overwrite_network_config,omitempty"`

	// The ID of the Slot to swap with Production.
	// The ID of the Slot to swap with `Production`.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/web/v1beta1.WindowsFunctionAppSlot
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SlotID *string `json:"slotId,omitempty" tf:"slot_id,omitempty"`

	// Reference to a WindowsFunctionAppSlot in web to populate slotId.
	// +kubebuilder:validation:Optional
	SlotIDRef *v1.Reference `json:"slotIdRef,omitempty" tf:"-"`

	// Selector for a WindowsFunctionAppSlot in web to populate slotId.
	// +kubebuilder:validation:Optional
	SlotIDSelector *v1.Selector `json:"slotIdSelector,omitempty" tf:"-"`
}

// FunctionAppActiveSlotSpec defines the desired state of FunctionAppActiveSlot
type FunctionAppActiveSlotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionAppActiveSlotParameters `json:"forProvider"`
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
	InitProvider FunctionAppActiveSlotInitParameters `json:"initProvider,omitempty"`
}

// FunctionAppActiveSlotStatus defines the observed state of FunctionAppActiveSlot.
type FunctionAppActiveSlotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionAppActiveSlotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionAppActiveSlot is the Schema for the FunctionAppActiveSlots API. Manages a Function App Active Slot.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FunctionAppActiveSlot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionAppActiveSlotSpec   `json:"spec"`
	Status            FunctionAppActiveSlotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionAppActiveSlotList contains a list of FunctionAppActiveSlots
type FunctionAppActiveSlotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FunctionAppActiveSlot `json:"items"`
}

// Repository type metadata.
var (
	FunctionAppActiveSlot_Kind             = "FunctionAppActiveSlot"
	FunctionAppActiveSlot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FunctionAppActiveSlot_Kind}.String()
	FunctionAppActiveSlot_KindAPIVersion   = FunctionAppActiveSlot_Kind + "." + CRDGroupVersion.String()
	FunctionAppActiveSlot_GroupVersionKind = CRDGroupVersion.WithKind(FunctionAppActiveSlot_Kind)
)

func init() {
	SchemeBuilder.Register(&FunctionAppActiveSlot{}, &FunctionAppActiveSlotList{})
}

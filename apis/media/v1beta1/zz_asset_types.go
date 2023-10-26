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

type AssetInitParameters struct {

	// The alternate ID of the Asset.
	AlternateID *string `json:"alternateId,omitempty" tf:"alternate_id,omitempty"`

	// The name of the asset blob container. Changing this forces a new Media Asset to be created.
	Container *string `json:"container,omitempty" tf:"container,omitempty"`

	// The Asset description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the storage account where to store the media asset. Changing this forces a new Media Asset to be created.
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`
}

type AssetObservation struct {

	// The alternate ID of the Asset.
	AlternateID *string `json:"alternateId,omitempty" tf:"alternate_id,omitempty"`

	// The name of the asset blob container. Changing this forces a new Media Asset to be created.
	Container *string `json:"container,omitempty" tf:"container,omitempty"`

	// The Asset description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Media Asset.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the Media Services Account. Changing this forces a new Media Asset to be created.
	MediaServicesAccountName *string `json:"mediaServicesAccountName,omitempty" tf:"media_services_account_name,omitempty"`

	// The name of the Resource Group where the Media Asset should exist. Changing this forces a new Media Asset to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The name of the storage account where to store the media asset. Changing this forces a new Media Asset to be created.
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`
}

type AssetParameters struct {

	// The alternate ID of the Asset.
	// +kubebuilder:validation:Optional
	AlternateID *string `json:"alternateId,omitempty" tf:"alternate_id,omitempty"`

	// The name of the asset blob container. Changing this forces a new Media Asset to be created.
	// +kubebuilder:validation:Optional
	Container *string `json:"container,omitempty" tf:"container,omitempty"`

	// The Asset description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the name of the Media Services Account. Changing this forces a new Media Asset to be created.
	// +crossplane:generate:reference:type=ServicesAccount
	// +kubebuilder:validation:Optional
	MediaServicesAccountName *string `json:"mediaServicesAccountName,omitempty" tf:"media_services_account_name,omitempty"`

	// Reference to a ServicesAccount to populate mediaServicesAccountName.
	// +kubebuilder:validation:Optional
	MediaServicesAccountNameRef *v1.Reference `json:"mediaServicesAccountNameRef,omitempty" tf:"-"`

	// Selector for a ServicesAccount to populate mediaServicesAccountName.
	// +kubebuilder:validation:Optional
	MediaServicesAccountNameSelector *v1.Selector `json:"mediaServicesAccountNameSelector,omitempty" tf:"-"`

	// The name of the Resource Group where the Media Asset should exist. Changing this forces a new Media Asset to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The name of the storage account where to store the media asset. Changing this forces a new Media Asset to be created.
	// +kubebuilder:validation:Optional
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`
}

// AssetSpec defines the desired state of Asset
type AssetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AssetParameters `json:"forProvider"`
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
	InitProvider AssetInitParameters `json:"initProvider,omitempty"`
}

// AssetStatus defines the observed state of Asset.
type AssetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AssetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Asset is the Schema for the Assets API. Manages a Media Asset.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Asset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AssetSpec   `json:"spec"`
	Status            AssetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AssetList contains a list of Assets
type AssetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Asset `json:"items"`
}

// Repository type metadata.
var (
	Asset_Kind             = "Asset"
	Asset_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Asset_Kind}.String()
	Asset_KindAPIVersion   = Asset_Kind + "." + CRDGroupVersion.String()
	Asset_GroupVersionKind = CRDGroupVersion.WithKind(Asset_Kind)
)

func init() {
	SchemeBuilder.Register(&Asset{}, &AssetList{})
}

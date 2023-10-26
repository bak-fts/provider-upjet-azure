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

type ProductTagInitParameters struct {
}

type ProductTagObservation struct {

	// The name of the API Management Service. Changing this forces a new resource to be created.
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// The name of the API Management product. Changing this forces a new resource to be created.
	APIManagementProductID *string `json:"apiManagementProductId,omitempty" tf:"api_management_product_id,omitempty"`

	// The ID of the API Management Product.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name which should be used for this API Management Tag. Changing this forces a new API Management Tag to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type ProductTagParameters struct {

	// The name of the API Management Service. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// The name of the API Management product. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Product
	// +kubebuilder:validation:Optional
	APIManagementProductID *string `json:"apiManagementProductId,omitempty" tf:"api_management_product_id,omitempty"`

	// Reference to a Product in apimanagement to populate apiManagementProductId.
	// +kubebuilder:validation:Optional
	APIManagementProductIDRef *v1.Reference `json:"apiManagementProductIdRef,omitempty" tf:"-"`

	// Selector for a Product in apimanagement to populate apiManagementProductId.
	// +kubebuilder:validation:Optional
	APIManagementProductIDSelector *v1.Selector `json:"apiManagementProductIdSelector,omitempty" tf:"-"`

	// The name which should be used for this API Management Tag. Changing this forces a new API Management Tag to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Tag
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Tag in apimanagement to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Tag in apimanagement to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// ProductTagSpec defines the desired state of ProductTag
type ProductTagSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProductTagParameters `json:"forProvider"`
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
	InitProvider ProductTagInitParameters `json:"initProvider,omitempty"`
}

// ProductTagStatus defines the observed state of ProductTag.
type ProductTagStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProductTagObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProductTag is the Schema for the ProductTags API. Manages an API Management Product tag
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ProductTag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProductTagSpec   `json:"spec"`
	Status            ProductTagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProductTagList contains a list of ProductTags
type ProductTagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProductTag `json:"items"`
}

// Repository type metadata.
var (
	ProductTag_Kind             = "ProductTag"
	ProductTag_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProductTag_Kind}.String()
	ProductTag_KindAPIVersion   = ProductTag_Kind + "." + CRDGroupVersion.String()
	ProductTag_GroupVersionKind = CRDGroupVersion.WithKind(ProductTag_Kind)
)

func init() {
	SchemeBuilder.Register(&ProductTag{}, &ProductTagList{})
}

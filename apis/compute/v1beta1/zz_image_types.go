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

type DataDiskInitParameters struct {

	// Specifies the URI in Azure storage of the blob that you want to use to create the image.
	BlobURI *string `json:"blobUri,omitempty" tf:"blob_uri,omitempty"`

	// Specifies the caching mode as ReadWrite, ReadOnly, or None. Defaults to None.
	Caching *string `json:"caching,omitempty" tf:"caching,omitempty"`

	// Specifies the logical unit number of the data disk.
	Lun *float64 `json:"lun,omitempty" tf:"lun,omitempty"`

	// Specifies the ID of the managed disk resource that you want to use to create the image. Changing this forces a new resource to be created.
	ManagedDiskID *string `json:"managedDiskId,omitempty" tf:"managed_disk_id,omitempty"`

	// Specifies the size of the image to be created. The target size can't be smaller than the source size.
	SizeGb *float64 `json:"sizeGb,omitempty" tf:"size_gb,omitempty"`
}

type DataDiskObservation struct {

	// Specifies the URI in Azure storage of the blob that you want to use to create the image.
	BlobURI *string `json:"blobUri,omitempty" tf:"blob_uri,omitempty"`

	// Specifies the caching mode as ReadWrite, ReadOnly, or None. Defaults to None.
	Caching *string `json:"caching,omitempty" tf:"caching,omitempty"`

	// Specifies the logical unit number of the data disk.
	Lun *float64 `json:"lun,omitempty" tf:"lun,omitempty"`

	// Specifies the ID of the managed disk resource that you want to use to create the image. Changing this forces a new resource to be created.
	ManagedDiskID *string `json:"managedDiskId,omitempty" tf:"managed_disk_id,omitempty"`

	// Specifies the size of the image to be created. The target size can't be smaller than the source size.
	SizeGb *float64 `json:"sizeGb,omitempty" tf:"size_gb,omitempty"`
}

type DataDiskParameters struct {

	// Specifies the URI in Azure storage of the blob that you want to use to create the image.
	// +kubebuilder:validation:Optional
	BlobURI *string `json:"blobUri,omitempty" tf:"blob_uri,omitempty"`

	// Specifies the caching mode as ReadWrite, ReadOnly, or None. Defaults to None.
	// +kubebuilder:validation:Optional
	Caching *string `json:"caching,omitempty" tf:"caching,omitempty"`

	// Specifies the logical unit number of the data disk.
	// +kubebuilder:validation:Optional
	Lun *float64 `json:"lun,omitempty" tf:"lun,omitempty"`

	// Specifies the ID of the managed disk resource that you want to use to create the image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ManagedDiskID *string `json:"managedDiskId,omitempty" tf:"managed_disk_id,omitempty"`

	// Specifies the size of the image to be created. The target size can't be smaller than the source size.
	// +kubebuilder:validation:Optional
	SizeGb *float64 `json:"sizeGb,omitempty" tf:"size_gb,omitempty"`
}

type ImageInitParameters struct {

	// One or more data_disk blocks as defined below.
	DataDisk []DataDiskInitParameters `json:"dataDisk,omitempty" tf:"data_disk,omitempty"`

	// The HyperVGenerationType of the VirtualMachine created from the image as V1, V2. Defaults to V1. Changing this forces a new resource to be created.
	HyperVGeneration *string `json:"hyperVGeneration,omitempty" tf:"hyper_v_generation,omitempty"`

	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// One or more os_disk blocks as defined below. Changing this forces a new resource to be created.
	OsDisk []OsDiskInitParameters `json:"osDisk,omitempty" tf:"os_disk,omitempty"`

	// The Virtual Machine ID from which to create the image.
	SourceVirtualMachineID *string `json:"sourceVirtualMachineId,omitempty" tf:"source_virtual_machine_id,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Is zone resiliency enabled? Defaults to false. Changing this forces a new resource to be created.
	ZoneResilient *bool `json:"zoneResilient,omitempty" tf:"zone_resilient,omitempty"`
}

type ImageObservation struct {

	// One or more data_disk blocks as defined below.
	DataDisk []DataDiskObservation `json:"dataDisk,omitempty" tf:"data_disk,omitempty"`

	// The HyperVGenerationType of the VirtualMachine created from the image as V1, V2. Defaults to V1. Changing this forces a new resource to be created.
	HyperVGeneration *string `json:"hyperVGeneration,omitempty" tf:"hyper_v_generation,omitempty"`

	// The ID of the Image.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// One or more os_disk blocks as defined below. Changing this forces a new resource to be created.
	OsDisk []OsDiskObservation `json:"osDisk,omitempty" tf:"os_disk,omitempty"`

	// The name of the resource group in which to create. Changing this forces a new resource to be created.
	// the image. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The Virtual Machine ID from which to create the image.
	SourceVirtualMachineID *string `json:"sourceVirtualMachineId,omitempty" tf:"source_virtual_machine_id,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Is zone resiliency enabled? Defaults to false. Changing this forces a new resource to be created.
	ZoneResilient *bool `json:"zoneResilient,omitempty" tf:"zone_resilient,omitempty"`
}

type ImageParameters struct {

	// One or more data_disk blocks as defined below.
	// +kubebuilder:validation:Optional
	DataDisk []DataDiskParameters `json:"dataDisk,omitempty" tf:"data_disk,omitempty"`

	// The HyperVGenerationType of the VirtualMachine created from the image as V1, V2. Defaults to V1. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	HyperVGeneration *string `json:"hyperVGeneration,omitempty" tf:"hyper_v_generation,omitempty"`

	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// One or more os_disk blocks as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	OsDisk []OsDiskParameters `json:"osDisk,omitempty" tf:"os_disk,omitempty"`

	// The name of the resource group in which to create. Changing this forces a new resource to be created.
	// the image. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Virtual Machine ID from which to create the image.
	// +kubebuilder:validation:Optional
	SourceVirtualMachineID *string `json:"sourceVirtualMachineId,omitempty" tf:"source_virtual_machine_id,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Is zone resiliency enabled? Defaults to false. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ZoneResilient *bool `json:"zoneResilient,omitempty" tf:"zone_resilient,omitempty"`
}

type OsDiskInitParameters struct {

	// Specifies the URI in Azure storage of the blob that you want to use to create the image. Changing this forces a new resource to be created.
	BlobURI *string `json:"blobUri,omitempty" tf:"blob_uri,omitempty"`

	// Specifies the caching mode as ReadWrite, ReadOnly, or None. The default is None.
	Caching *string `json:"caching,omitempty" tf:"caching,omitempty"`

	// The ID of the Disk Encryption Set which should be used to encrypt this image. Changing this forces a new resource to be created.
	DiskEncryptionSetID *string `json:"diskEncryptionSetId,omitempty" tf:"disk_encryption_set_id,omitempty"`

	// Specifies the ID of the managed disk resource that you want to use to create the image.
	ManagedDiskID *string `json:"managedDiskId,omitempty" tf:"managed_disk_id,omitempty"`

	// Specifies the state of the operating system contained in the blob. Currently, the only value is Generalized. Possible values are Generalized and Specialized.
	OsState *string `json:"osState,omitempty" tf:"os_state,omitempty"`

	// Specifies the type of operating system contained in the virtual machine image. Possible values are: Windows or Linux.
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// Specifies the size of the image to be created. Changing this forces a new resource to be created.
	SizeGb *float64 `json:"sizeGb,omitempty" tf:"size_gb,omitempty"`
}

type OsDiskObservation struct {

	// Specifies the URI in Azure storage of the blob that you want to use to create the image. Changing this forces a new resource to be created.
	BlobURI *string `json:"blobUri,omitempty" tf:"blob_uri,omitempty"`

	// Specifies the caching mode as ReadWrite, ReadOnly, or None. The default is None.
	Caching *string `json:"caching,omitempty" tf:"caching,omitempty"`

	// The ID of the Disk Encryption Set which should be used to encrypt this image. Changing this forces a new resource to be created.
	DiskEncryptionSetID *string `json:"diskEncryptionSetId,omitempty" tf:"disk_encryption_set_id,omitempty"`

	// Specifies the ID of the managed disk resource that you want to use to create the image.
	ManagedDiskID *string `json:"managedDiskId,omitempty" tf:"managed_disk_id,omitempty"`

	// Specifies the state of the operating system contained in the blob. Currently, the only value is Generalized. Possible values are Generalized and Specialized.
	OsState *string `json:"osState,omitempty" tf:"os_state,omitempty"`

	// Specifies the type of operating system contained in the virtual machine image. Possible values are: Windows or Linux.
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// Specifies the size of the image to be created. Changing this forces a new resource to be created.
	SizeGb *float64 `json:"sizeGb,omitempty" tf:"size_gb,omitempty"`
}

type OsDiskParameters struct {

	// Specifies the URI in Azure storage of the blob that you want to use to create the image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	BlobURI *string `json:"blobUri,omitempty" tf:"blob_uri,omitempty"`

	// Specifies the caching mode as ReadWrite, ReadOnly, or None. The default is None.
	// +kubebuilder:validation:Optional
	Caching *string `json:"caching,omitempty" tf:"caching,omitempty"`

	// The ID of the Disk Encryption Set which should be used to encrypt this image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DiskEncryptionSetID *string `json:"diskEncryptionSetId,omitempty" tf:"disk_encryption_set_id,omitempty"`

	// Specifies the ID of the managed disk resource that you want to use to create the image.
	// +kubebuilder:validation:Optional
	ManagedDiskID *string `json:"managedDiskId,omitempty" tf:"managed_disk_id,omitempty"`

	// Specifies the state of the operating system contained in the blob. Currently, the only value is Generalized. Possible values are Generalized and Specialized.
	// +kubebuilder:validation:Optional
	OsState *string `json:"osState,omitempty" tf:"os_state,omitempty"`

	// Specifies the type of operating system contained in the virtual machine image. Possible values are: Windows or Linux.
	// +kubebuilder:validation:Optional
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// Specifies the size of the image to be created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SizeGb *float64 `json:"sizeGb,omitempty" tf:"size_gb,omitempty"`
}

// ImageSpec defines the desired state of Image
type ImageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ImageInitParameters `json:"initProvider,omitempty"`
}

// ImageStatus defines the observed state of Image.
type ImageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Image is the Schema for the Images API. Manages a custom virtual machine image that can be used to create virtual machines.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Image struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   ImageSpec   `json:"spec"`
	Status ImageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageList contains a list of Images
type ImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Image `json:"items"`
}

// Repository type metadata.
var (
	Image_Kind             = "Image"
	Image_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Image_Kind}.String()
	Image_KindAPIVersion   = Image_Kind + "." + CRDGroupVersion.String()
	Image_GroupVersionKind = CRDGroupVersion.WithKind(Image_Kind)
)

func init() {
	SchemeBuilder.Register(&Image{}, &ImageList{})
}

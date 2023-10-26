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

type KeyVaultSASTokenInitParameters struct {

	// Specifies the secret name in Azure Key Vault that stores the SAS token.
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`
}

type KeyVaultSASTokenObservation struct {

	// Specifies the name of an existing Key Vault Data Factory Linked Service.
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Specifies the secret name in Azure Key Vault that stores the SAS token.
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`
}

type KeyVaultSASTokenParameters struct {

	// Specifies the name of an existing Key Vault Data Factory Linked Service.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.LinkedServiceKeyVault
	// +kubebuilder:validation:Optional
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Reference to a LinkedServiceKeyVault in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameRef *v1.Reference `json:"linkedServiceNameRef,omitempty" tf:"-"`

	// Selector for a LinkedServiceKeyVault in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameSelector *v1.Selector `json:"linkedServiceNameSelector,omitempty" tf:"-"`

	// Specifies the secret name in Azure Key Vault that stores the SAS token.
	// +kubebuilder:validation:Optional
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type LinkedServiceAzureBlobStorageInitParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The connection string sent insecurely. Conflicts with connection_string, sas_uri and service_endpoint.
	ConnectionStringInsecure *string `json:"connectionStringInsecure,omitempty" tf:"connection_string_insecure,omitempty"`

	// The description for the Data Factory Linked Service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A key_vault_sas_token block as defined below. Use this argument to store SAS Token in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service. A sas_uri is required.
	KeyVaultSASToken []KeyVaultSASTokenInitParameters `json:"keyVaultSasToken,omitempty" tf:"key_vault_sas_token,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The service principal id in which to authenticate against the Azure Blob Storage account.
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// The service principal key in which to authenticate against the AAzure Blob Storage account.
	ServicePrincipalKey *string `json:"servicePrincipalKey,omitempty" tf:"service_principal_key,omitempty"`

	// A service_principal_linked_key_vault_key block as defined below. Use this argument to store Service Principal key in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	ServicePrincipalLinkedKeyVaultKey []ServicePrincipalLinkedKeyVaultKeyInitParameters `json:"servicePrincipalLinkedKeyVaultKey,omitempty" tf:"service_principal_linked_key_vault_key,omitempty"`

	// Specify the kind of the storage account. Allowed values are Storage, StorageV2, BlobStorage and BlockBlobStorage.
	StorageKind *string `json:"storageKind,omitempty" tf:"storage_kind,omitempty"`

	// The tenant id or name in which to authenticate against the Azure Blob Storage account.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Whether to use the Data Factory's managed identity to authenticate against the Azure Blob Storage account. Incompatible with service_principal_id and service_principal_key.
	UseManagedIdentity *bool `json:"useManagedIdentity,omitempty" tf:"use_managed_identity,omitempty"`
}

type LinkedServiceAzureBlobStorageObservation struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The connection string sent insecurely. Conflicts with connection_string, sas_uri and service_endpoint.
	ConnectionStringInsecure *string `json:"connectionStringInsecure,omitempty" tf:"connection_string_insecure,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The description for the Data Factory Linked Service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Data Factory Linked Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A key_vault_sas_token block as defined below. Use this argument to store SAS Token in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service. A sas_uri is required.
	KeyVaultSASToken []KeyVaultSASTokenObservation `json:"keyVaultSasToken,omitempty" tf:"key_vault_sas_token,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The service principal id in which to authenticate against the Azure Blob Storage account.
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// The service principal key in which to authenticate against the AAzure Blob Storage account.
	ServicePrincipalKey *string `json:"servicePrincipalKey,omitempty" tf:"service_principal_key,omitempty"`

	// A service_principal_linked_key_vault_key block as defined below. Use this argument to store Service Principal key in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	ServicePrincipalLinkedKeyVaultKey []ServicePrincipalLinkedKeyVaultKeyObservation `json:"servicePrincipalLinkedKeyVaultKey,omitempty" tf:"service_principal_linked_key_vault_key,omitempty"`

	// Specify the kind of the storage account. Allowed values are Storage, StorageV2, BlobStorage and BlockBlobStorage.
	StorageKind *string `json:"storageKind,omitempty" tf:"storage_kind,omitempty"`

	// The tenant id or name in which to authenticate against the Azure Blob Storage account.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Whether to use the Data Factory's managed identity to authenticate against the Azure Blob Storage account. Incompatible with service_principal_id and service_principal_key.
	UseManagedIdentity *bool `json:"useManagedIdentity,omitempty" tf:"use_managed_identity,omitempty"`
}

type LinkedServiceAzureBlobStorageParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The connection string sent insecurely. Conflicts with connection_string, sas_uri and service_endpoint.
	// +kubebuilder:validation:Optional
	ConnectionStringInsecure *string `json:"connectionStringInsecure,omitempty" tf:"connection_string_insecure,omitempty"`

	// The connection string. Conflicts with connection_string_insecure, sas_uri and service_endpoint.
	// +kubebuilder:validation:Optional
	ConnectionStringSecretRef *v1.SecretKeySelector `json:"connectionStringSecretRef,omitempty" tf:"-"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The description for the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A key_vault_sas_token block as defined below. Use this argument to store SAS Token in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service. A sas_uri is required.
	// +kubebuilder:validation:Optional
	KeyVaultSASToken []KeyVaultSASTokenParameters `json:"keyVaultSasToken,omitempty" tf:"key_vault_sas_token,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The SAS URI. Conflicts with connection_string_insecure, connection_string and service_endpoint.
	// +kubebuilder:validation:Optional
	SASURISecretRef *v1.SecretKeySelector `json:"sasuriSecretRef,omitempty" tf:"-"`

	// The Service Endpoint. Conflicts with connection_string, connection_string_insecure and sas_uri.
	// +kubebuilder:validation:Optional
	ServiceEndpointSecretRef *v1.SecretKeySelector `json:"serviceEndpointSecretRef,omitempty" tf:"-"`

	// The service principal id in which to authenticate against the Azure Blob Storage account.
	// +kubebuilder:validation:Optional
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// The service principal key in which to authenticate against the AAzure Blob Storage account.
	// +kubebuilder:validation:Optional
	ServicePrincipalKey *string `json:"servicePrincipalKey,omitempty" tf:"service_principal_key,omitempty"`

	// A service_principal_linked_key_vault_key block as defined below. Use this argument to store Service Principal key in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	ServicePrincipalLinkedKeyVaultKey []ServicePrincipalLinkedKeyVaultKeyParameters `json:"servicePrincipalLinkedKeyVaultKey,omitempty" tf:"service_principal_linked_key_vault_key,omitempty"`

	// Specify the kind of the storage account. Allowed values are Storage, StorageV2, BlobStorage and BlockBlobStorage.
	// +kubebuilder:validation:Optional
	StorageKind *string `json:"storageKind,omitempty" tf:"storage_kind,omitempty"`

	// The tenant id or name in which to authenticate against the Azure Blob Storage account.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Whether to use the Data Factory's managed identity to authenticate against the Azure Blob Storage account. Incompatible with service_principal_id and service_principal_key.
	// +kubebuilder:validation:Optional
	UseManagedIdentity *bool `json:"useManagedIdentity,omitempty" tf:"use_managed_identity,omitempty"`
}

type ServicePrincipalLinkedKeyVaultKeyInitParameters struct {

	// Specifies the secret name in Azure Key Vault that stores the Service Principal key.
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`
}

type ServicePrincipalLinkedKeyVaultKeyObservation struct {

	// Specifies the name of an existing Key Vault Data Factory Linked Service.
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Specifies the secret name in Azure Key Vault that stores the Service Principal key.
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`
}

type ServicePrincipalLinkedKeyVaultKeyParameters struct {

	// Specifies the name of an existing Key Vault Data Factory Linked Service.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.LinkedServiceKeyVault
	// +kubebuilder:validation:Optional
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Reference to a LinkedServiceKeyVault in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameRef *v1.Reference `json:"linkedServiceNameRef,omitempty" tf:"-"`

	// Selector for a LinkedServiceKeyVault in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameSelector *v1.Selector `json:"linkedServiceNameSelector,omitempty" tf:"-"`

	// Specifies the secret name in Azure Key Vault that stores the Service Principal key.
	// +kubebuilder:validation:Optional
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

// LinkedServiceAzureBlobStorageSpec defines the desired state of LinkedServiceAzureBlobStorage
type LinkedServiceAzureBlobStorageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceAzureBlobStorageParameters `json:"forProvider"`
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
	InitProvider LinkedServiceAzureBlobStorageInitParameters `json:"initProvider,omitempty"`
}

// LinkedServiceAzureBlobStorageStatus defines the observed state of LinkedServiceAzureBlobStorage.
type LinkedServiceAzureBlobStorageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceAzureBlobStorageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceAzureBlobStorage is the Schema for the LinkedServiceAzureBlobStorages API. Manages a Linked Service (connection) between an Azure Blob Storage Account and Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinkedServiceAzureBlobStorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinkedServiceAzureBlobStorageSpec   `json:"spec"`
	Status            LinkedServiceAzureBlobStorageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceAzureBlobStorageList contains a list of LinkedServiceAzureBlobStorages
type LinkedServiceAzureBlobStorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServiceAzureBlobStorage `json:"items"`
}

// Repository type metadata.
var (
	LinkedServiceAzureBlobStorage_Kind             = "LinkedServiceAzureBlobStorage"
	LinkedServiceAzureBlobStorage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServiceAzureBlobStorage_Kind}.String()
	LinkedServiceAzureBlobStorage_KindAPIVersion   = LinkedServiceAzureBlobStorage_Kind + "." + CRDGroupVersion.String()
	LinkedServiceAzureBlobStorage_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServiceAzureBlobStorage_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServiceAzureBlobStorage{}, &LinkedServiceAzureBlobStorageList{})
}

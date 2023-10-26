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

type AdminInitParameters struct {

	// E-mail address of the admin.
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address,omitempty"`

	// First name of the admin.
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// Last name of the admin.
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`

	// Phone number of the admin.
	Phone *string `json:"phone,omitempty" tf:"phone,omitempty"`
}

type AdminObservation struct {

	// E-mail address of the admin.
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address,omitempty"`

	// First name of the admin.
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// Last name of the admin.
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`

	// Phone number of the admin.
	Phone *string `json:"phone,omitempty" tf:"phone,omitempty"`
}

type AdminParameters struct {

	// E-mail address of the admin.
	// +kubebuilder:validation:Optional
	EmailAddress *string `json:"emailAddress" tf:"email_address,omitempty"`

	// First name of the admin.
	// +kubebuilder:validation:Optional
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// Last name of the admin.
	// +kubebuilder:validation:Optional
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`

	// Phone number of the admin.
	// +kubebuilder:validation:Optional
	Phone *string `json:"phone,omitempty" tf:"phone,omitempty"`
}

type CertificateIssuerInitParameters struct {

	// The account number with the third-party Certificate Issuer.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// One or more admin blocks as defined below.
	Admin []AdminInitParameters `json:"admin,omitempty" tf:"admin,omitempty"`

	// The ID of the organization as provided to the issuer.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// The name of the third-party Certificate Issuer. Possible values are: DigiCert, GlobalSign, OneCertV2-PrivateCA, OneCertV2-PublicCA and SslAdminV2.
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`
}

type CertificateIssuerObservation struct {

	// The account number with the third-party Certificate Issuer.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// One or more admin blocks as defined below.
	Admin []AdminObservation `json:"admin,omitempty" tf:"admin,omitempty"`

	// The ID of the Key Vault Certificate Issuer.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Key Vault in which to create the Certificate Issuer. Changing this forces a new resource to be created.
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// The ID of the organization as provided to the issuer.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// The name of the third-party Certificate Issuer. Possible values are: DigiCert, GlobalSign, OneCertV2-PrivateCA, OneCertV2-PublicCA and SslAdminV2.
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`
}

type CertificateIssuerParameters struct {

	// The account number with the third-party Certificate Issuer.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// One or more admin blocks as defined below.
	// +kubebuilder:validation:Optional
	Admin []AdminParameters `json:"admin,omitempty" tf:"admin,omitempty"`

	// The ID of the Key Vault in which to create the Certificate Issuer. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Vault
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Reference to a Vault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDRef *v1.Reference `json:"keyVaultIdRef,omitempty" tf:"-"`

	// Selector for a Vault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDSelector *v1.Selector `json:"keyVaultIdSelector,omitempty" tf:"-"`

	// The ID of the organization as provided to the issuer.
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// The password associated with the account and organization ID at the third-party Certificate Issuer. If not specified, will not overwrite any previous value.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The name of the third-party Certificate Issuer. Possible values are: DigiCert, GlobalSign, OneCertV2-PrivateCA, OneCertV2-PublicCA and SslAdminV2.
	// +kubebuilder:validation:Optional
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`
}

// CertificateIssuerSpec defines the desired state of CertificateIssuer
type CertificateIssuerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateIssuerParameters `json:"forProvider"`
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
	InitProvider CertificateIssuerInitParameters `json:"initProvider,omitempty"`
}

// CertificateIssuerStatus defines the observed state of CertificateIssuer.
type CertificateIssuerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateIssuerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateIssuer is the Schema for the CertificateIssuers API. Manages a Key Vault Certificate Issuer.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type CertificateIssuer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.providerName) || (has(self.initProvider) && has(self.initProvider.providerName))",message="spec.forProvider.providerName is a required parameter"
	Spec   CertificateIssuerSpec   `json:"spec"`
	Status CertificateIssuerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateIssuerList contains a list of CertificateIssuers
type CertificateIssuerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateIssuer `json:"items"`
}

// Repository type metadata.
var (
	CertificateIssuer_Kind             = "CertificateIssuer"
	CertificateIssuer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CertificateIssuer_Kind}.String()
	CertificateIssuer_KindAPIVersion   = CertificateIssuer_Kind + "." + CRDGroupVersion.String()
	CertificateIssuer_GroupVersionKind = CRDGroupVersion.WithKind(CertificateIssuer_Kind)
)

func init() {
	SchemeBuilder.Register(&CertificateIssuer{}, &CertificateIssuerList{})
}

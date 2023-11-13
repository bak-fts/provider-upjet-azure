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

type AzureadAdministratorInitParameters struct {

	// Specifies whether only AD Users and administrators (e.g. azuread_administrator.0.login_username) can be used to login, or also local database users (e.g. administrator_login). When true, the administrator_login and administrator_login_password properties can be omitted.
	AzureadAuthenticationOnly *bool `json:"azureadAuthenticationOnly,omitempty" tf:"azuread_authentication_only,omitempty"`

	// The tenant id of the Azure AD Administrator of this SQL Server.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AzureadAdministratorObservation struct {

	// Specifies whether only AD Users and administrators (e.g. azuread_administrator.0.login_username) can be used to login, or also local database users (e.g. administrator_login). When true, the administrator_login and administrator_login_password properties can be omitted.
	AzureadAuthenticationOnly *bool `json:"azureadAuthenticationOnly,omitempty" tf:"azuread_authentication_only,omitempty"`

	// The login username of the Azure AD Administrator of this SQL Server.
	LoginUsername *string `json:"loginUsername,omitempty" tf:"login_username,omitempty"`

	// The object id of the Azure AD Administrator of this SQL Server.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The tenant id of the Azure AD Administrator of this SQL Server.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AzureadAdministratorParameters struct {

	// Specifies whether only AD Users and administrators (e.g. azuread_administrator.0.login_username) can be used to login, or also local database users (e.g. administrator_login). When true, the administrator_login and administrator_login_password properties can be omitted.
	// +kubebuilder:validation:Optional
	AzureadAuthenticationOnly *bool `json:"azureadAuthenticationOnly,omitempty" tf:"azuread_authentication_only,omitempty"`

	// The login username of the Azure AD Administrator of this SQL Server.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/managedidentity/v1beta1.UserAssignedIdentity
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	LoginUsername *string `json:"loginUsername,omitempty" tf:"login_username,omitempty"`

	// Reference to a UserAssignedIdentity in managedidentity to populate loginUsername.
	// +kubebuilder:validation:Optional
	LoginUsernameRef *v1.Reference `json:"loginUsernameRef,omitempty" tf:"-"`

	// Selector for a UserAssignedIdentity in managedidentity to populate loginUsername.
	// +kubebuilder:validation:Optional
	LoginUsernameSelector *v1.Selector `json:"loginUsernameSelector,omitempty" tf:"-"`

	// The object id of the Azure AD Administrator of this SQL Server.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/managedidentity/v1beta1.UserAssignedIdentity
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("principal_id",true)
	// +kubebuilder:validation:Optional
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// Reference to a UserAssignedIdentity in managedidentity to populate objectId.
	// +kubebuilder:validation:Optional
	ObjectIDRef *v1.Reference `json:"objectIdRef,omitempty" tf:"-"`

	// Selector for a UserAssignedIdentity in managedidentity to populate objectId.
	// +kubebuilder:validation:Optional
	ObjectIDSelector *v1.Selector `json:"objectIdSelector,omitempty" tf:"-"`

	// The tenant id of the Azure AD Administrator of this SQL Server.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type MSSQLServerIdentityInitParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this SQL Server.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this SQL Server. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type MSSQLServerIdentityObservation struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this SQL Server.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID for the Service Principal associated with the Identity of this SQL Server.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID for the Service Principal associated with the Identity of this SQL Server.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this SQL Server. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type MSSQLServerIdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this SQL Server.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this SQL Server. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type MSSQLServerInitParameters struct {

	// The administrator login name for the new server. Required unless azuread_authentication_only in the azuread_administrator block is true. When omitted, Azure will generate a default username which cannot be subsequently changed. Changing this forces a new resource to be created.
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login,omitempty"`

	// An azuread_administrator block as defined below.
	AzureadAdministrator []AzureadAdministratorInitParameters `json:"azureadAdministrator,omitempty" tf:"azuread_administrator,omitempty"`

	// The connection policy the server will use. Possible values are Default, Proxy, and Redirect. Defaults to Default.
	ConnectionPolicy *string `json:"connectionPolicy,omitempty" tf:"connection_policy,omitempty"`

	// An identity block as defined below.
	Identity []MSSQLServerIdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Minimum TLS Version for all SQL Database and SQL Data Warehouse databases associated with the server. Valid values are: 1.0, 1.1 , 1.2 and Disabled. Defaults to 1.2.
	MinimumTLSVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version,omitempty"`

	// Whether outbound network traffic is restricted for this server. Defaults to false.
	OutboundNetworkRestrictionEnabled *bool `json:"outboundNetworkRestrictionEnabled,omitempty" tf:"outbound_network_restriction_enabled,omitempty"`

	// Whether public network access is allowed for this server. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The version for the new server. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server). Changing this forces a new resource to be created.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type MSSQLServerObservation struct {

	// The administrator login name for the new server. Required unless azuread_authentication_only in the azuread_administrator block is true. When omitted, Azure will generate a default username which cannot be subsequently changed. Changing this forces a new resource to be created.
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login,omitempty"`

	// An azuread_administrator block as defined below.
	AzureadAdministrator []AzureadAdministratorObservation `json:"azureadAdministrator,omitempty" tf:"azuread_administrator,omitempty"`

	// The connection policy the server will use. Possible values are Default, Proxy, and Redirect. Defaults to Default.
	ConnectionPolicy *string `json:"connectionPolicy,omitempty" tf:"connection_policy,omitempty"`

	// The fully qualified domain name of the Azure SQL Server (e.g. myServerName.database.windows.net)
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty" tf:"fully_qualified_domain_name,omitempty"`

	// the Microsoft SQL Server ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []MSSQLServerIdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Minimum TLS Version for all SQL Database and SQL Data Warehouse databases associated with the server. Valid values are: 1.0, 1.1 , 1.2 and Disabled. Defaults to 1.2.
	MinimumTLSVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version,omitempty"`

	// Whether outbound network traffic is restricted for this server. Defaults to false.
	OutboundNetworkRestrictionEnabled *bool `json:"outboundNetworkRestrictionEnabled,omitempty" tf:"outbound_network_restriction_enabled,omitempty"`

	// Specifies the primary user managed identity id. Required if type is UserAssigned and should be combined with identity_ids.
	PrimaryUserAssignedIdentityID *string `json:"primaryUserAssignedIdentityId,omitempty" tf:"primary_user_assigned_identity_id,omitempty"`

	// Whether public network access is allowed for this server. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group in which to create the Microsoft SQL Server. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A list of dropped restorable database IDs on the server.
	RestorableDroppedDatabaseIds []*string `json:"restorableDroppedDatabaseIds,omitempty" tf:"restorable_dropped_database_ids,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The fully versioned Key Vault Key URL (e.g. 'https://<YourVaultName>.vault.azure.net/keys/<YourKeyName>/<YourKeyVersion>) to be used as the Customer Managed Key(CMK/BYOK) for the Transparent Data Encryption(TDE) layer.
	TransparentDataEncryptionKeyVaultKeyID *string `json:"transparentDataEncryptionKeyVaultKeyId,omitempty" tf:"transparent_data_encryption_key_vault_key_id,omitempty"`

	// The version for the new server. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server). Changing this forces a new resource to be created.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type MSSQLServerParameters struct {

	// The administrator login name for the new server. Required unless azuread_authentication_only in the azuread_administrator block is true. When omitted, Azure will generate a default username which cannot be subsequently changed. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login,omitempty"`

	// The password associated with the administrator_login user. Needs to comply with Azure's Password Policy. Required unless azuread_authentication_only in the azuread_administrator block is true.
	// +kubebuilder:validation:Optional
	AdministratorLoginPasswordSecretRef *v1.SecretKeySelector `json:"administratorLoginPasswordSecretRef,omitempty" tf:"-"`

	// An azuread_administrator block as defined below.
	// +kubebuilder:validation:Optional
	AzureadAdministrator []AzureadAdministratorParameters `json:"azureadAdministrator,omitempty" tf:"azuread_administrator,omitempty"`

	// The connection policy the server will use. Possible values are Default, Proxy, and Redirect. Defaults to Default.
	// +kubebuilder:validation:Optional
	ConnectionPolicy *string `json:"connectionPolicy,omitempty" tf:"connection_policy,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []MSSQLServerIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Minimum TLS Version for all SQL Database and SQL Data Warehouse databases associated with the server. Valid values are: 1.0, 1.1 , 1.2 and Disabled. Defaults to 1.2.
	// +kubebuilder:validation:Optional
	MinimumTLSVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version,omitempty"`

	// Whether outbound network traffic is restricted for this server. Defaults to false.
	// +kubebuilder:validation:Optional
	OutboundNetworkRestrictionEnabled *bool `json:"outboundNetworkRestrictionEnabled,omitempty" tf:"outbound_network_restriction_enabled,omitempty"`

	// Specifies the primary user managed identity id. Required if type is UserAssigned and should be combined with identity_ids.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/managedidentity/v1beta1.UserAssignedIdentity
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PrimaryUserAssignedIdentityID *string `json:"primaryUserAssignedIdentityId,omitempty" tf:"primary_user_assigned_identity_id,omitempty"`

	// Reference to a UserAssignedIdentity in managedidentity to populate primaryUserAssignedIdentityId.
	// +kubebuilder:validation:Optional
	PrimaryUserAssignedIdentityIDRef *v1.Reference `json:"primaryUserAssignedIdentityIdRef,omitempty" tf:"-"`

	// Selector for a UserAssignedIdentity in managedidentity to populate primaryUserAssignedIdentityId.
	// +kubebuilder:validation:Optional
	PrimaryUserAssignedIdentityIDSelector *v1.Selector `json:"primaryUserAssignedIdentityIdSelector,omitempty" tf:"-"`

	// Whether public network access is allowed for this server. Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group in which to create the Microsoft SQL Server. Changing this forces a new resource to be created.
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

	// The fully versioned Key Vault Key URL (e.g. 'https://<YourVaultName>.vault.azure.net/keys/<YourKeyName>/<YourKeyVersion>) to be used as the Customer Managed Key(CMK/BYOK) for the Transparent Data Encryption(TDE) layer.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TransparentDataEncryptionKeyVaultKeyID *string `json:"transparentDataEncryptionKeyVaultKeyId,omitempty" tf:"transparent_data_encryption_key_vault_key_id,omitempty"`

	// Reference to a Key in keyvault to populate transparentDataEncryptionKeyVaultKeyId.
	// +kubebuilder:validation:Optional
	TransparentDataEncryptionKeyVaultKeyIDRef *v1.Reference `json:"transparentDataEncryptionKeyVaultKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in keyvault to populate transparentDataEncryptionKeyVaultKeyId.
	// +kubebuilder:validation:Optional
	TransparentDataEncryptionKeyVaultKeyIDSelector *v1.Selector `json:"transparentDataEncryptionKeyVaultKeyIdSelector,omitempty" tf:"-"`

	// The version for the new server. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server). Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

// MSSQLServerSpec defines the desired state of MSSQLServer
type MSSQLServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLServerParameters `json:"forProvider"`
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
	InitProvider MSSQLServerInitParameters `json:"initProvider,omitempty"`
}

// MSSQLServerStatus defines the observed state of MSSQLServer.
type MSSQLServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLServer is the Schema for the MSSQLServers API. Manages a Microsoft SQL Azure Database Server.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MSSQLServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.version) || (has(self.initProvider) && has(self.initProvider.version))",message="spec.forProvider.version is a required parameter"
	Spec   MSSQLServerSpec   `json:"spec"`
	Status MSSQLServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLServerList contains a list of MSSQLServers
type MSSQLServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLServer `json:"items"`
}

// Repository type metadata.
var (
	MSSQLServer_Kind             = "MSSQLServer"
	MSSQLServer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLServer_Kind}.String()
	MSSQLServer_KindAPIVersion   = MSSQLServer_Kind + "." + CRDGroupVersion.String()
	MSSQLServer_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLServer_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLServer{}, &MSSQLServerList{})
}

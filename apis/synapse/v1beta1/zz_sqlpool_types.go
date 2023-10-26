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

type RestoreInitParameters struct {

	// Specifies the Snapshot time to restore formatted as an RFC3339 date string. Changing this forces a new Synapse SQL Pool to be created.
	PointInTime *string `json:"pointInTime,omitempty" tf:"point_in_time,omitempty"`

	// The ID of the Synapse SQL Pool or SQL Database which is to restore. Changing this forces a new Synapse SQL Pool to be created.
	SourceDatabaseID *string `json:"sourceDatabaseId,omitempty" tf:"source_database_id,omitempty"`
}

type RestoreObservation struct {

	// Specifies the Snapshot time to restore formatted as an RFC3339 date string. Changing this forces a new Synapse SQL Pool to be created.
	PointInTime *string `json:"pointInTime,omitempty" tf:"point_in_time,omitempty"`

	// The ID of the Synapse SQL Pool or SQL Database which is to restore. Changing this forces a new Synapse SQL Pool to be created.
	SourceDatabaseID *string `json:"sourceDatabaseId,omitempty" tf:"source_database_id,omitempty"`
}

type RestoreParameters struct {

	// Specifies the Snapshot time to restore formatted as an RFC3339 date string. Changing this forces a new Synapse SQL Pool to be created.
	// +kubebuilder:validation:Optional
	PointInTime *string `json:"pointInTime" tf:"point_in_time,omitempty"`

	// The ID of the Synapse SQL Pool or SQL Database which is to restore. Changing this forces a new Synapse SQL Pool to be created.
	// +kubebuilder:validation:Optional
	SourceDatabaseID *string `json:"sourceDatabaseId" tf:"source_database_id,omitempty"`
}

type SQLPoolInitParameters struct {

	// The name of the collation to use with this pool, only applicable when create_mode is set to Default. Azure default is SQL_LATIN1_GENERAL_CP1_CI_AS. Changing this forces a new resource to be created.
	Collation *string `json:"collation,omitempty" tf:"collation,omitempty"`

	// Specifies how to create the SQL Pool. Valid values are: Default, Recovery or PointInTimeRestore. Must be Default to create a new database. Defaults to Default. Changing this forces a new resource to be created.
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode,omitempty"`

	// Is transparent data encryption enabled?
	DataEncrypted *bool `json:"dataEncrypted,omitempty" tf:"data_encrypted,omitempty"`

	// Is geo-backup policy enabled? Defaults to true.
	GeoBackupPolicyEnabled *bool `json:"geoBackupPolicyEnabled,omitempty" tf:"geo_backup_policy_enabled,omitempty"`

	// The ID of the Synapse SQL Pool or SQL Database which is to back up, only applicable when create_mode is set to Recovery. Changing this forces a new Synapse SQL Pool to be created.
	RecoveryDatabaseID *string `json:"recoveryDatabaseId,omitempty" tf:"recovery_database_id,omitempty"`

	// A restore block as defined below. only applicable when create_mode is set to PointInTimeRestore. Changing this forces a new resource to be created.
	Restore []RestoreInitParameters `json:"restore,omitempty" tf:"restore,omitempty"`

	// Specifies the SKU Name for this Synapse SQL Pool. Possible values are DW100c, DW200c, DW300c, DW400c, DW500c, DW1000c, DW1500c, DW2000c, DW2500c, DW3000c, DW5000c, DW6000c, DW7500c, DW10000c, DW15000c or DW30000c.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags which should be assigned to the Synapse SQL Pool.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SQLPoolObservation struct {

	// The name of the collation to use with this pool, only applicable when create_mode is set to Default. Azure default is SQL_LATIN1_GENERAL_CP1_CI_AS. Changing this forces a new resource to be created.
	Collation *string `json:"collation,omitempty" tf:"collation,omitempty"`

	// Specifies how to create the SQL Pool. Valid values are: Default, Recovery or PointInTimeRestore. Must be Default to create a new database. Defaults to Default. Changing this forces a new resource to be created.
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode,omitempty"`

	// Is transparent data encryption enabled?
	DataEncrypted *bool `json:"dataEncrypted,omitempty" tf:"data_encrypted,omitempty"`

	// Is geo-backup policy enabled? Defaults to true.
	GeoBackupPolicyEnabled *bool `json:"geoBackupPolicyEnabled,omitempty" tf:"geo_backup_policy_enabled,omitempty"`

	// The ID of the Synapse SQL Pool.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Synapse SQL Pool or SQL Database which is to back up, only applicable when create_mode is set to Recovery. Changing this forces a new Synapse SQL Pool to be created.
	RecoveryDatabaseID *string `json:"recoveryDatabaseId,omitempty" tf:"recovery_database_id,omitempty"`

	// A restore block as defined below. only applicable when create_mode is set to PointInTimeRestore. Changing this forces a new resource to be created.
	Restore []RestoreObservation `json:"restore,omitempty" tf:"restore,omitempty"`

	// Specifies the SKU Name for this Synapse SQL Pool. Possible values are DW100c, DW200c, DW300c, DW400c, DW500c, DW1000c, DW1500c, DW2000c, DW2500c, DW3000c, DW5000c, DW6000c, DW7500c, DW10000c, DW15000c or DW30000c.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// The ID of Synapse Workspace within which this SQL Pool should be created. Changing this forces a new Synapse SQL Pool to be created.
	SynapseWorkspaceID *string `json:"synapseWorkspaceId,omitempty" tf:"synapse_workspace_id,omitempty"`

	// A mapping of tags which should be assigned to the Synapse SQL Pool.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SQLPoolParameters struct {

	// The name of the collation to use with this pool, only applicable when create_mode is set to Default. Azure default is SQL_LATIN1_GENERAL_CP1_CI_AS. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Collation *string `json:"collation,omitempty" tf:"collation,omitempty"`

	// Specifies how to create the SQL Pool. Valid values are: Default, Recovery or PointInTimeRestore. Must be Default to create a new database. Defaults to Default. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode,omitempty"`

	// Is transparent data encryption enabled?
	// +kubebuilder:validation:Optional
	DataEncrypted *bool `json:"dataEncrypted,omitempty" tf:"data_encrypted,omitempty"`

	// Is geo-backup policy enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	GeoBackupPolicyEnabled *bool `json:"geoBackupPolicyEnabled,omitempty" tf:"geo_backup_policy_enabled,omitempty"`

	// The ID of the Synapse SQL Pool or SQL Database which is to back up, only applicable when create_mode is set to Recovery. Changing this forces a new Synapse SQL Pool to be created.
	// +kubebuilder:validation:Optional
	RecoveryDatabaseID *string `json:"recoveryDatabaseId,omitempty" tf:"recovery_database_id,omitempty"`

	// A restore block as defined below. only applicable when create_mode is set to PointInTimeRestore. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Restore []RestoreParameters `json:"restore,omitempty" tf:"restore,omitempty"`

	// Specifies the SKU Name for this Synapse SQL Pool. Possible values are DW100c, DW200c, DW300c, DW400c, DW500c, DW1000c, DW1500c, DW2000c, DW2500c, DW3000c, DW5000c, DW6000c, DW7500c, DW10000c, DW15000c or DW30000c.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// The ID of Synapse Workspace within which this SQL Pool should be created. Changing this forces a new Synapse SQL Pool to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/synapse/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SynapseWorkspaceID *string `json:"synapseWorkspaceId,omitempty" tf:"synapse_workspace_id,omitempty"`

	// Reference to a Workspace in synapse to populate synapseWorkspaceId.
	// +kubebuilder:validation:Optional
	SynapseWorkspaceIDRef *v1.Reference `json:"synapseWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in synapse to populate synapseWorkspaceId.
	// +kubebuilder:validation:Optional
	SynapseWorkspaceIDSelector *v1.Selector `json:"synapseWorkspaceIdSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Synapse SQL Pool.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SQLPoolSpec defines the desired state of SQLPool
type SQLPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLPoolParameters `json:"forProvider"`
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
	InitProvider SQLPoolInitParameters `json:"initProvider,omitempty"`
}

// SQLPoolStatus defines the observed state of SQLPool.
type SQLPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SQLPool is the Schema for the SQLPools API. Manages a Synapse SQL Pool.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SQLPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.skuName) || (has(self.initProvider) && has(self.initProvider.skuName))",message="spec.forProvider.skuName is a required parameter"
	Spec   SQLPoolSpec   `json:"spec"`
	Status SQLPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLPoolList contains a list of SQLPools
type SQLPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLPool `json:"items"`
}

// Repository type metadata.
var (
	SQLPool_Kind             = "SQLPool"
	SQLPool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLPool_Kind}.String()
	SQLPool_KindAPIVersion   = SQLPool_Kind + "." + CRDGroupVersion.String()
	SQLPool_GroupVersionKind = CRDGroupVersion.WithKind(SQLPool_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLPool{}, &SQLPoolList{})
}

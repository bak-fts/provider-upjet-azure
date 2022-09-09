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

type RoleAssignmentObservation struct {

	// The Role Assignment ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The type of the principal_id, e.g. User, Group, Service Principal, Application, etc.
	PrincipalType *string `json:"principalType,omitempty" tf:"principal_type,omitempty"`
}

type RoleAssignmentParameters struct {

	// The condition that limits the resources that the role can be assigned to. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// The version of the condition. Possible values are 1.0 or 2.0. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ConditionVersion *string `json:"conditionVersion,omitempty" tf:"condition_version,omitempty"`

	// The delegated Azure Resource Id which contains a Managed Identity. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DelegatedManagedIdentityResourceID *string `json:"delegatedManagedIdentityResourceId,omitempty" tf:"delegated_managed_identity_resource_id,omitempty"`

	// The description for this Role Assignment. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A unique UUID/GUID for this Role Assignment - one will be generated if not specified. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the Principal  to assign the Role Definition to. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	PrincipalID *string `json:"principalId" tf:"principal_id,omitempty"`

	// The Scoped-ID of the Role Definition. Changing this forces a new resource to be created. Conflicts with role_definition_name.
	// +kubebuilder:validation:Optional
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty" tf:"role_definition_id,omitempty"`

	// The name of a built-in Role. Changing this forces a new resource to be created. Conflicts with role_definition_id.
	// +kubebuilder:validation:Optional
	RoleDefinitionName *string `json:"roleDefinitionName,omitempty" tf:"role_definition_name,omitempty"`

	// The scope at which the Role Assignment applies to, such as /subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333, /subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup, or /subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM, or /providers/Microsoft.Management/managementGroups/myMG. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Scope *string `json:"scope" tf:"scope,omitempty"`

	// If the principal_id is a newly provisioned Service Principal set this value to true to skip the Azure Active Directory check which may fail due to replication lag. This argument is only valid if the principal_id is a Service Principal identity. If it is not a Service Principal identity it will cause the role assignment to fail. Defaults to false.
	// +kubebuilder:validation:Optional
	SkipServicePrincipalAADCheck *bool `json:"skipServicePrincipalAadCheck,omitempty" tf:"skip_service_principal_aad_check,omitempty"`
}

// RoleAssignmentSpec defines the desired state of RoleAssignment
type RoleAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleAssignmentParameters `json:"forProvider"`
}

// RoleAssignmentStatus defines the observed state of RoleAssignment.
type RoleAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RoleAssignment is the Schema for the RoleAssignments API. Assigns a given Principal (User or Group) to a given Role.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RoleAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleAssignmentSpec   `json:"spec"`
	Status            RoleAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleAssignmentList contains a list of RoleAssignments
type RoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleAssignment `json:"items"`
}

// Repository type metadata.
var (
	RoleAssignment_Kind             = "RoleAssignment"
	RoleAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoleAssignment_Kind}.String()
	RoleAssignment_KindAPIVersion   = RoleAssignment_Kind + "." + CRDGroupVersion.String()
	RoleAssignment_GroupVersionKind = CRDGroupVersion.WithKind(RoleAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&RoleAssignment{}, &RoleAssignmentList{})
}
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

type ExpressRouteCircuitConnectionInitParameters struct {

	// The IPv4 address space from which to allocate customer address for global reach. Changing this forces a new Express Route Circuit Connection to be created.
	AddressPrefixIPv4 *string `json:"addressPrefixIpv4,omitempty" tf:"address_prefix_ipv4,omitempty"`

	// The IPv6 address space from which to allocate customer addresses for global reach.
	AddressPrefixIPv6 *string `json:"addressPrefixIpv6,omitempty" tf:"address_prefix_ipv6,omitempty"`
}

type ExpressRouteCircuitConnectionObservation struct {

	// The IPv4 address space from which to allocate customer address for global reach. Changing this forces a new Express Route Circuit Connection to be created.
	AddressPrefixIPv4 *string `json:"addressPrefixIpv4,omitempty" tf:"address_prefix_ipv4,omitempty"`

	// The IPv6 address space from which to allocate customer addresses for global reach.
	AddressPrefixIPv6 *string `json:"addressPrefixIpv6,omitempty" tf:"address_prefix_ipv6,omitempty"`

	// The ID of the Express Route Circuit Connection.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the peered Express Route Circuit Private Peering. Changing this forces a new Express Route Circuit Connection to be created.
	PeerPeeringID *string `json:"peerPeeringId,omitempty" tf:"peer_peering_id,omitempty"`

	// The ID of the Express Route Circuit Private Peering that this Express Route Circuit Connection connects with. Changing this forces a new Express Route Circuit Connection to be created.
	PeeringID *string `json:"peeringId,omitempty" tf:"peering_id,omitempty"`
}

type ExpressRouteCircuitConnectionParameters struct {

	// The IPv4 address space from which to allocate customer address for global reach. Changing this forces a new Express Route Circuit Connection to be created.
	// +kubebuilder:validation:Optional
	AddressPrefixIPv4 *string `json:"addressPrefixIpv4,omitempty" tf:"address_prefix_ipv4,omitempty"`

	// The IPv6 address space from which to allocate customer addresses for global reach.
	// +kubebuilder:validation:Optional
	AddressPrefixIPv6 *string `json:"addressPrefixIpv6,omitempty" tf:"address_prefix_ipv6,omitempty"`

	// The authorization key which is associated with the Express Route Circuit Connection.
	// +kubebuilder:validation:Optional
	AuthorizationKeySecretRef *v1.SecretKeySelector `json:"authorizationKeySecretRef,omitempty" tf:"-"`

	// The ID of the peered Express Route Circuit Private Peering. Changing this forces a new Express Route Circuit Connection to be created.
	// +crossplane:generate:reference:type=ExpressRouteCircuitPeering
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PeerPeeringID *string `json:"peerPeeringId,omitempty" tf:"peer_peering_id,omitempty"`

	// Reference to a ExpressRouteCircuitPeering to populate peerPeeringId.
	// +kubebuilder:validation:Optional
	PeerPeeringIDRef *v1.Reference `json:"peerPeeringIdRef,omitempty" tf:"-"`

	// Selector for a ExpressRouteCircuitPeering to populate peerPeeringId.
	// +kubebuilder:validation:Optional
	PeerPeeringIDSelector *v1.Selector `json:"peerPeeringIdSelector,omitempty" tf:"-"`

	// The ID of the Express Route Circuit Private Peering that this Express Route Circuit Connection connects with. Changing this forces a new Express Route Circuit Connection to be created.
	// +crossplane:generate:reference:type=ExpressRouteCircuitPeering
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PeeringID *string `json:"peeringId,omitempty" tf:"peering_id,omitempty"`

	// Reference to a ExpressRouteCircuitPeering to populate peeringId.
	// +kubebuilder:validation:Optional
	PeeringIDRef *v1.Reference `json:"peeringIdRef,omitempty" tf:"-"`

	// Selector for a ExpressRouteCircuitPeering to populate peeringId.
	// +kubebuilder:validation:Optional
	PeeringIDSelector *v1.Selector `json:"peeringIdSelector,omitempty" tf:"-"`
}

// ExpressRouteCircuitConnectionSpec defines the desired state of ExpressRouteCircuitConnection
type ExpressRouteCircuitConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExpressRouteCircuitConnectionParameters `json:"forProvider"`
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
	InitProvider ExpressRouteCircuitConnectionInitParameters `json:"initProvider,omitempty"`
}

// ExpressRouteCircuitConnectionStatus defines the observed state of ExpressRouteCircuitConnection.
type ExpressRouteCircuitConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExpressRouteCircuitConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteCircuitConnection is the Schema for the ExpressRouteCircuitConnections API. Manages an Express Route Circuit Connection.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ExpressRouteCircuitConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.addressPrefixIpv4) || (has(self.initProvider) && has(self.initProvider.addressPrefixIpv4))",message="spec.forProvider.addressPrefixIpv4 is a required parameter"
	Spec   ExpressRouteCircuitConnectionSpec   `json:"spec"`
	Status ExpressRouteCircuitConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteCircuitConnectionList contains a list of ExpressRouteCircuitConnections
type ExpressRouteCircuitConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExpressRouteCircuitConnection `json:"items"`
}

// Repository type metadata.
var (
	ExpressRouteCircuitConnection_Kind             = "ExpressRouteCircuitConnection"
	ExpressRouteCircuitConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ExpressRouteCircuitConnection_Kind}.String()
	ExpressRouteCircuitConnection_KindAPIVersion   = ExpressRouteCircuitConnection_Kind + "." + CRDGroupVersion.String()
	ExpressRouteCircuitConnection_GroupVersionKind = CRDGroupVersion.WithKind(ExpressRouteCircuitConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&ExpressRouteCircuitConnection{}, &ExpressRouteCircuitConnectionList{})
}

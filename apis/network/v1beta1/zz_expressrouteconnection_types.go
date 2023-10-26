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

type ExpressRouteConnectionInitParameters struct {

	// The authorization key to establish the Express Route Connection.
	AuthorizationKey *string `json:"authorizationKey,omitempty" tf:"authorization_key,omitempty"`

	// Is Internet security enabled for this Express Route Connection?
	EnableInternetSecurity *bool `json:"enableInternetSecurity,omitempty" tf:"enable_internet_security,omitempty"`

	// Specified whether Fast Path is enabled for Virtual Wan Firewall Hub. Defaults to false.
	ExpressRouteGatewayBypassEnabled *bool `json:"expressRouteGatewayBypassEnabled,omitempty" tf:"express_route_gateway_bypass_enabled,omitempty"`

	// A routing block as defined below.
	Routing []RoutingInitParameters `json:"routing,omitempty" tf:"routing,omitempty"`

	// The routing weight associated to the Express Route Connection. Possible value is between 0 and 32000. Defaults to 0.
	RoutingWeight *float64 `json:"routingWeight,omitempty" tf:"routing_weight,omitempty"`
}

type ExpressRouteConnectionObservation struct {

	// The authorization key to establish the Express Route Connection.
	AuthorizationKey *string `json:"authorizationKey,omitempty" tf:"authorization_key,omitempty"`

	// Is Internet security enabled for this Express Route Connection?
	EnableInternetSecurity *bool `json:"enableInternetSecurity,omitempty" tf:"enable_internet_security,omitempty"`

	// The ID of the Express Route Circuit Peering that this Express Route Connection connects with. Changing this forces a new resource to be created.
	ExpressRouteCircuitPeeringID *string `json:"expressRouteCircuitPeeringId,omitempty" tf:"express_route_circuit_peering_id,omitempty"`

	// Specified whether Fast Path is enabled for Virtual Wan Firewall Hub. Defaults to false.
	ExpressRouteGatewayBypassEnabled *bool `json:"expressRouteGatewayBypassEnabled,omitempty" tf:"express_route_gateway_bypass_enabled,omitempty"`

	// The ID of the Express Route Gateway that this Express Route Connection connects with. Changing this forces a new resource to be created.
	ExpressRouteGatewayID *string `json:"expressRouteGatewayId,omitempty" tf:"express_route_gateway_id,omitempty"`

	// The ID of the Express Route Connection.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A routing block as defined below.
	Routing []RoutingObservation `json:"routing,omitempty" tf:"routing,omitempty"`

	// The routing weight associated to the Express Route Connection. Possible value is between 0 and 32000. Defaults to 0.
	RoutingWeight *float64 `json:"routingWeight,omitempty" tf:"routing_weight,omitempty"`
}

type ExpressRouteConnectionParameters struct {

	// The authorization key to establish the Express Route Connection.
	// +kubebuilder:validation:Optional
	AuthorizationKey *string `json:"authorizationKey,omitempty" tf:"authorization_key,omitempty"`

	// Is Internet security enabled for this Express Route Connection?
	// +kubebuilder:validation:Optional
	EnableInternetSecurity *bool `json:"enableInternetSecurity,omitempty" tf:"enable_internet_security,omitempty"`

	// The ID of the Express Route Circuit Peering that this Express Route Connection connects with. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=ExpressRouteCircuitPeering
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ExpressRouteCircuitPeeringID *string `json:"expressRouteCircuitPeeringId,omitempty" tf:"express_route_circuit_peering_id,omitempty"`

	// Reference to a ExpressRouteCircuitPeering to populate expressRouteCircuitPeeringId.
	// +kubebuilder:validation:Optional
	ExpressRouteCircuitPeeringIDRef *v1.Reference `json:"expressRouteCircuitPeeringIdRef,omitempty" tf:"-"`

	// Selector for a ExpressRouteCircuitPeering to populate expressRouteCircuitPeeringId.
	// +kubebuilder:validation:Optional
	ExpressRouteCircuitPeeringIDSelector *v1.Selector `json:"expressRouteCircuitPeeringIdSelector,omitempty" tf:"-"`

	// Specified whether Fast Path is enabled for Virtual Wan Firewall Hub. Defaults to false.
	// +kubebuilder:validation:Optional
	ExpressRouteGatewayBypassEnabled *bool `json:"expressRouteGatewayBypassEnabled,omitempty" tf:"express_route_gateway_bypass_enabled,omitempty"`

	// The ID of the Express Route Gateway that this Express Route Connection connects with. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=ExpressRouteGateway
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ExpressRouteGatewayID *string `json:"expressRouteGatewayId,omitempty" tf:"express_route_gateway_id,omitempty"`

	// Reference to a ExpressRouteGateway to populate expressRouteGatewayId.
	// +kubebuilder:validation:Optional
	ExpressRouteGatewayIDRef *v1.Reference `json:"expressRouteGatewayIdRef,omitempty" tf:"-"`

	// Selector for a ExpressRouteGateway to populate expressRouteGatewayId.
	// +kubebuilder:validation:Optional
	ExpressRouteGatewayIDSelector *v1.Selector `json:"expressRouteGatewayIdSelector,omitempty" tf:"-"`

	// A routing block as defined below.
	// +kubebuilder:validation:Optional
	Routing []RoutingParameters `json:"routing,omitempty" tf:"routing,omitempty"`

	// The routing weight associated to the Express Route Connection. Possible value is between 0 and 32000. Defaults to 0.
	// +kubebuilder:validation:Optional
	RoutingWeight *float64 `json:"routingWeight,omitempty" tf:"routing_weight,omitempty"`
}

type PropagatedRouteTableInitParameters struct {

	// The list of labels to logically group route tables.
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A list of IDs of the Virtual Hub Route Table to propagate routes from Express Route Connection to the route table.
	RouteTableIds []*string `json:"routeTableIds,omitempty" tf:"route_table_ids,omitempty"`
}

type PropagatedRouteTableObservation struct {

	// The list of labels to logically group route tables.
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A list of IDs of the Virtual Hub Route Table to propagate routes from Express Route Connection to the route table.
	RouteTableIds []*string `json:"routeTableIds,omitempty" tf:"route_table_ids,omitempty"`
}

type PropagatedRouteTableParameters struct {

	// The list of labels to logically group route tables.
	// +kubebuilder:validation:Optional
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A list of IDs of the Virtual Hub Route Table to propagate routes from Express Route Connection to the route table.
	// +kubebuilder:validation:Optional
	RouteTableIds []*string `json:"routeTableIds,omitempty" tf:"route_table_ids,omitempty"`
}

type RoutingInitParameters struct {

	// The ID of the Virtual Hub Route Table associated with this Express Route Connection.
	AssociatedRouteTableID *string `json:"associatedRouteTableId,omitempty" tf:"associated_route_table_id,omitempty"`

	// The ID of the Route Map associated with this Express Route Connection for inbound routes.
	InboundRouteMapID *string `json:"inboundRouteMapId,omitempty" tf:"inbound_route_map_id,omitempty"`

	// The ID of the Route Map associated with this Express Route Connection for outbound routes.
	OutboundRouteMapID *string `json:"outboundRouteMapId,omitempty" tf:"outbound_route_map_id,omitempty"`

	// A propagated_route_table block as defined below.
	PropagatedRouteTable []PropagatedRouteTableInitParameters `json:"propagatedRouteTable,omitempty" tf:"propagated_route_table,omitempty"`
}

type RoutingObservation struct {

	// The ID of the Virtual Hub Route Table associated with this Express Route Connection.
	AssociatedRouteTableID *string `json:"associatedRouteTableId,omitempty" tf:"associated_route_table_id,omitempty"`

	// The ID of the Route Map associated with this Express Route Connection for inbound routes.
	InboundRouteMapID *string `json:"inboundRouteMapId,omitempty" tf:"inbound_route_map_id,omitempty"`

	// The ID of the Route Map associated with this Express Route Connection for outbound routes.
	OutboundRouteMapID *string `json:"outboundRouteMapId,omitempty" tf:"outbound_route_map_id,omitempty"`

	// A propagated_route_table block as defined below.
	PropagatedRouteTable []PropagatedRouteTableObservation `json:"propagatedRouteTable,omitempty" tf:"propagated_route_table,omitempty"`
}

type RoutingParameters struct {

	// The ID of the Virtual Hub Route Table associated with this Express Route Connection.
	// +kubebuilder:validation:Optional
	AssociatedRouteTableID *string `json:"associatedRouteTableId,omitempty" tf:"associated_route_table_id,omitempty"`

	// The ID of the Route Map associated with this Express Route Connection for inbound routes.
	// +kubebuilder:validation:Optional
	InboundRouteMapID *string `json:"inboundRouteMapId,omitempty" tf:"inbound_route_map_id,omitempty"`

	// The ID of the Route Map associated with this Express Route Connection for outbound routes.
	// +kubebuilder:validation:Optional
	OutboundRouteMapID *string `json:"outboundRouteMapId,omitempty" tf:"outbound_route_map_id,omitempty"`

	// A propagated_route_table block as defined below.
	// +kubebuilder:validation:Optional
	PropagatedRouteTable []PropagatedRouteTableParameters `json:"propagatedRouteTable,omitempty" tf:"propagated_route_table,omitempty"`
}

// ExpressRouteConnectionSpec defines the desired state of ExpressRouteConnection
type ExpressRouteConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExpressRouteConnectionParameters `json:"forProvider"`
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
	InitProvider ExpressRouteConnectionInitParameters `json:"initProvider,omitempty"`
}

// ExpressRouteConnectionStatus defines the observed state of ExpressRouteConnection.
type ExpressRouteConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExpressRouteConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteConnection is the Schema for the ExpressRouteConnections API. Manages an Express Route Connection.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ExpressRouteConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExpressRouteConnectionSpec   `json:"spec"`
	Status            ExpressRouteConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteConnectionList contains a list of ExpressRouteConnections
type ExpressRouteConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExpressRouteConnection `json:"items"`
}

// Repository type metadata.
var (
	ExpressRouteConnection_Kind             = "ExpressRouteConnection"
	ExpressRouteConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ExpressRouteConnection_Kind}.String()
	ExpressRouteConnection_KindAPIVersion   = ExpressRouteConnection_Kind + "." + CRDGroupVersion.String()
	ExpressRouteConnection_GroupVersionKind = CRDGroupVersion.WithKind(ExpressRouteConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&ExpressRouteConnection{}, &ExpressRouteConnectionList{})
}

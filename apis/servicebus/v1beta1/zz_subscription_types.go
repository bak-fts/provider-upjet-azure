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

type ClientScopedSubscriptionInitParameters struct {

	// Specifies the Client ID of the application that created the client-scoped subscription. Changing this forces a new resource to be created.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Whether the client scoped subscription is shareable. Defaults to true Changing this forces a new resource to be created.
	IsClientScopedSubscriptionShareable *bool `json:"isClientScopedSubscriptionShareable,omitempty" tf:"is_client_scoped_subscription_shareable,omitempty"`
}

type ClientScopedSubscriptionObservation struct {

	// Specifies the Client ID of the application that created the client-scoped subscription. Changing this forces a new resource to be created.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Whether the client scoped subscription is durable. This property can only be controlled from the application side.
	IsClientScopedSubscriptionDurable *bool `json:"isClientScopedSubscriptionDurable,omitempty" tf:"is_client_scoped_subscription_durable,omitempty"`

	// Whether the client scoped subscription is shareable. Defaults to true Changing this forces a new resource to be created.
	IsClientScopedSubscriptionShareable *bool `json:"isClientScopedSubscriptionShareable,omitempty" tf:"is_client_scoped_subscription_shareable,omitempty"`
}

type ClientScopedSubscriptionParameters struct {

	// Specifies the Client ID of the application that created the client-scoped subscription. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Whether the client scoped subscription is shareable. Defaults to true Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	IsClientScopedSubscriptionShareable *bool `json:"isClientScopedSubscriptionShareable,omitempty" tf:"is_client_scoped_subscription_shareable,omitempty"`
}

type SubscriptionInitParameters struct {

	// The idle interval after which the topic is automatically deleted as an ISO 8601 duration. The minimum duration is 5 minutes or PT5M.
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty" tf:"auto_delete_on_idle,omitempty"`

	// A client_scoped_subscription block as defined below.
	ClientScopedSubscription []ClientScopedSubscriptionInitParameters `json:"clientScopedSubscription,omitempty" tf:"client_scoped_subscription,omitempty"`

	// whether the subscription is scoped to a client id. Defaults to false.
	ClientScopedSubscriptionEnabled *bool `json:"clientScopedSubscriptionEnabled,omitempty" tf:"client_scoped_subscription_enabled,omitempty"`

	// Boolean flag which controls whether the Subscription has dead letter support on filter evaluation exceptions. Defaults to true.
	DeadLetteringOnFilterEvaluationError *bool `json:"deadLetteringOnFilterEvaluationError,omitempty" tf:"dead_lettering_on_filter_evaluation_error,omitempty"`

	// Boolean flag which controls whether the Subscription has dead letter support when a message expires.
	DeadLetteringOnMessageExpiration *bool `json:"deadLetteringOnMessageExpiration,omitempty" tf:"dead_lettering_on_message_expiration,omitempty"`

	// The Default message timespan to live as an ISO 8601 duration. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTTL *string `json:"defaultMessageTtl,omitempty" tf:"default_message_ttl,omitempty"`

	// Boolean flag which controls whether the Subscription supports batched operations.
	EnableBatchedOperations *bool `json:"enableBatchedOperations,omitempty" tf:"enable_batched_operations,omitempty"`

	// The name of a Queue or Topic to automatically forward Dead Letter messages to.
	ForwardDeadLetteredMessagesTo *string `json:"forwardDeadLetteredMessagesTo,omitempty" tf:"forward_dead_lettered_messages_to,omitempty"`

	// The name of a Queue or Topic to automatically forward messages to.
	ForwardTo *string `json:"forwardTo,omitempty" tf:"forward_to,omitempty"`

	// The lock duration for the subscription as an ISO 8601 duration. The default value is 1 minute or P0DT0H1M0S . The maximum value is 5 minutes or P0DT0H5M0S .
	LockDuration *string `json:"lockDuration,omitempty" tf:"lock_duration,omitempty"`

	// The maximum number of deliveries.
	MaxDeliveryCount *float64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`

	// Boolean flag which controls whether this Subscription supports the concept of a session. Changing this forces a new resource to be created.
	RequiresSession *bool `json:"requiresSession,omitempty" tf:"requires_session,omitempty"`

	// The status of the Subscription. Possible values are Active,ReceiveDisabled, or Disabled. Defaults to Active.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type SubscriptionObservation struct {

	// The idle interval after which the topic is automatically deleted as an ISO 8601 duration. The minimum duration is 5 minutes or PT5M.
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty" tf:"auto_delete_on_idle,omitempty"`

	// A client_scoped_subscription block as defined below.
	ClientScopedSubscription []ClientScopedSubscriptionObservation `json:"clientScopedSubscription,omitempty" tf:"client_scoped_subscription,omitempty"`

	// whether the subscription is scoped to a client id. Defaults to false.
	ClientScopedSubscriptionEnabled *bool `json:"clientScopedSubscriptionEnabled,omitempty" tf:"client_scoped_subscription_enabled,omitempty"`

	// Boolean flag which controls whether the Subscription has dead letter support on filter evaluation exceptions. Defaults to true.
	DeadLetteringOnFilterEvaluationError *bool `json:"deadLetteringOnFilterEvaluationError,omitempty" tf:"dead_lettering_on_filter_evaluation_error,omitempty"`

	// Boolean flag which controls whether the Subscription has dead letter support when a message expires.
	DeadLetteringOnMessageExpiration *bool `json:"deadLetteringOnMessageExpiration,omitempty" tf:"dead_lettering_on_message_expiration,omitempty"`

	// The Default message timespan to live as an ISO 8601 duration. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTTL *string `json:"defaultMessageTtl,omitempty" tf:"default_message_ttl,omitempty"`

	// Boolean flag which controls whether the Subscription supports batched operations.
	EnableBatchedOperations *bool `json:"enableBatchedOperations,omitempty" tf:"enable_batched_operations,omitempty"`

	// The name of a Queue or Topic to automatically forward Dead Letter messages to.
	ForwardDeadLetteredMessagesTo *string `json:"forwardDeadLetteredMessagesTo,omitempty" tf:"forward_dead_lettered_messages_to,omitempty"`

	// The name of a Queue or Topic to automatically forward messages to.
	ForwardTo *string `json:"forwardTo,omitempty" tf:"forward_to,omitempty"`

	// The ServiceBus Subscription ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The lock duration for the subscription as an ISO 8601 duration. The default value is 1 minute or P0DT0H1M0S . The maximum value is 5 minutes or P0DT0H5M0S .
	LockDuration *string `json:"lockDuration,omitempty" tf:"lock_duration,omitempty"`

	// The maximum number of deliveries.
	MaxDeliveryCount *float64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`

	// Boolean flag which controls whether this Subscription supports the concept of a session. Changing this forces a new resource to be created.
	RequiresSession *bool `json:"requiresSession,omitempty" tf:"requires_session,omitempty"`

	// The status of the Subscription. Possible values are Active,ReceiveDisabled, or Disabled. Defaults to Active.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The ID of the ServiceBus Topic to create this Subscription in. Changing this forces a new resource to be created.
	TopicID *string `json:"topicId,omitempty" tf:"topic_id,omitempty"`
}

type SubscriptionParameters struct {

	// The idle interval after which the topic is automatically deleted as an ISO 8601 duration. The minimum duration is 5 minutes or PT5M.
	// +kubebuilder:validation:Optional
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty" tf:"auto_delete_on_idle,omitempty"`

	// A client_scoped_subscription block as defined below.
	// +kubebuilder:validation:Optional
	ClientScopedSubscription []ClientScopedSubscriptionParameters `json:"clientScopedSubscription,omitempty" tf:"client_scoped_subscription,omitempty"`

	// whether the subscription is scoped to a client id. Defaults to false.
	// +kubebuilder:validation:Optional
	ClientScopedSubscriptionEnabled *bool `json:"clientScopedSubscriptionEnabled,omitempty" tf:"client_scoped_subscription_enabled,omitempty"`

	// Boolean flag which controls whether the Subscription has dead letter support on filter evaluation exceptions. Defaults to true.
	// +kubebuilder:validation:Optional
	DeadLetteringOnFilterEvaluationError *bool `json:"deadLetteringOnFilterEvaluationError,omitempty" tf:"dead_lettering_on_filter_evaluation_error,omitempty"`

	// Boolean flag which controls whether the Subscription has dead letter support when a message expires.
	// +kubebuilder:validation:Optional
	DeadLetteringOnMessageExpiration *bool `json:"deadLetteringOnMessageExpiration,omitempty" tf:"dead_lettering_on_message_expiration,omitempty"`

	// The Default message timespan to live as an ISO 8601 duration. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	// +kubebuilder:validation:Optional
	DefaultMessageTTL *string `json:"defaultMessageTtl,omitempty" tf:"default_message_ttl,omitempty"`

	// Boolean flag which controls whether the Subscription supports batched operations.
	// +kubebuilder:validation:Optional
	EnableBatchedOperations *bool `json:"enableBatchedOperations,omitempty" tf:"enable_batched_operations,omitempty"`

	// The name of a Queue or Topic to automatically forward Dead Letter messages to.
	// +kubebuilder:validation:Optional
	ForwardDeadLetteredMessagesTo *string `json:"forwardDeadLetteredMessagesTo,omitempty" tf:"forward_dead_lettered_messages_to,omitempty"`

	// The name of a Queue or Topic to automatically forward messages to.
	// +kubebuilder:validation:Optional
	ForwardTo *string `json:"forwardTo,omitempty" tf:"forward_to,omitempty"`

	// The lock duration for the subscription as an ISO 8601 duration. The default value is 1 minute or P0DT0H1M0S . The maximum value is 5 minutes or P0DT0H5M0S .
	// +kubebuilder:validation:Optional
	LockDuration *string `json:"lockDuration,omitempty" tf:"lock_duration,omitempty"`

	// The maximum number of deliveries.
	// +kubebuilder:validation:Optional
	MaxDeliveryCount *float64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`

	// Boolean flag which controls whether this Subscription supports the concept of a session. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	RequiresSession *bool `json:"requiresSession,omitempty" tf:"requires_session,omitempty"`

	// The status of the Subscription. Possible values are Active,ReceiveDisabled, or Disabled. Defaults to Active.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The ID of the ServiceBus Topic to create this Subscription in. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/servicebus/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TopicID *string `json:"topicId,omitempty" tf:"topic_id,omitempty"`

	// Reference to a Topic in servicebus to populate topicId.
	// +kubebuilder:validation:Optional
	TopicIDRef *v1.Reference `json:"topicIdRef,omitempty" tf:"-"`

	// Selector for a Topic in servicebus to populate topicId.
	// +kubebuilder:validation:Optional
	TopicIDSelector *v1.Selector `json:"topicIdSelector,omitempty" tf:"-"`
}

// SubscriptionSpec defines the desired state of Subscription
type SubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubscriptionParameters `json:"forProvider"`
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
	InitProvider SubscriptionInitParameters `json:"initProvider,omitempty"`
}

// SubscriptionStatus defines the observed state of Subscription.
type SubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Subscription is the Schema for the Subscriptions API. Manages a ServiceBus Subscription.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Subscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.maxDeliveryCount) || (has(self.initProvider) && has(self.initProvider.maxDeliveryCount))",message="spec.forProvider.maxDeliveryCount is a required parameter"
	Spec   SubscriptionSpec   `json:"spec"`
	Status SubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionList contains a list of Subscriptions
type SubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subscription `json:"items"`
}

// Repository type metadata.
var (
	Subscription_Kind             = "Subscription"
	Subscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Subscription_Kind}.String()
	Subscription_KindAPIVersion   = Subscription_Kind + "." + CRDGroupVersion.String()
	Subscription_GroupVersionKind = CRDGroupVersion.WithKind(Subscription_Kind)
)

func init() {
	SchemeBuilder.Register(&Subscription{}, &SubscriptionList{})
}

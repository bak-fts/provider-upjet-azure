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

type ApplicationInsightsWorkbookTemplateInitParameters struct {

	// Information about the author of the workbook template.
	Author *string `json:"author,omitempty" tf:"author,omitempty"`

	// A galleries block as defined below.
	Galleries []GalleriesInitParameters `json:"galleries,omitempty" tf:"galleries,omitempty"`

	// Key value pairs of localized gallery. Each key is the locale code of languages supported by the Azure portal.
	Localized *string `json:"localized,omitempty" tf:"localized,omitempty"`

	// Specifies the Azure Region where the Application Insights Workbook Template should exist. Changing this forces a new Application Insights Workbook Template to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name which should be used for this Application Insights Workbook Template. Changing this forces a new Application Insights Workbook Template to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Priority of the template. Determines which template to open when a workbook gallery is opened in viewer mode. Defaults to 0.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// A mapping of tags which should be assigned to the Application Insights Workbook Template.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Valid JSON object containing workbook template payload.
	TemplateData *string `json:"templateData,omitempty" tf:"template_data,omitempty"`
}

type ApplicationInsightsWorkbookTemplateObservation struct {

	// Information about the author of the workbook template.
	Author *string `json:"author,omitempty" tf:"author,omitempty"`

	// A galleries block as defined below.
	Galleries []GalleriesObservation `json:"galleries,omitempty" tf:"galleries,omitempty"`

	// The ID of the Application Insights Workbook Template.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key value pairs of localized gallery. Each key is the locale code of languages supported by the Azure portal.
	Localized *string `json:"localized,omitempty" tf:"localized,omitempty"`

	// Specifies the Azure Region where the Application Insights Workbook Template should exist. Changing this forces a new Application Insights Workbook Template to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name which should be used for this Application Insights Workbook Template. Changing this forces a new Application Insights Workbook Template to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Priority of the template. Determines which template to open when a workbook gallery is opened in viewer mode. Defaults to 0.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Specifies the name of the Resource Group where the Application Insights Workbook Template should exist. Changing this forces a new Application Insights Workbook Template to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags which should be assigned to the Application Insights Workbook Template.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Valid JSON object containing workbook template payload.
	TemplateData *string `json:"templateData,omitempty" tf:"template_data,omitempty"`
}

type ApplicationInsightsWorkbookTemplateParameters struct {

	// Information about the author of the workbook template.
	// +kubebuilder:validation:Optional
	Author *string `json:"author,omitempty" tf:"author,omitempty"`

	// A galleries block as defined below.
	// +kubebuilder:validation:Optional
	Galleries []GalleriesParameters `json:"galleries,omitempty" tf:"galleries,omitempty"`

	// Key value pairs of localized gallery. Each key is the locale code of languages supported by the Azure portal.
	// +kubebuilder:validation:Optional
	Localized *string `json:"localized,omitempty" tf:"localized,omitempty"`

	// Specifies the Azure Region where the Application Insights Workbook Template should exist. Changing this forces a new Application Insights Workbook Template to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name which should be used for this Application Insights Workbook Template. Changing this forces a new Application Insights Workbook Template to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Priority of the template. Determines which template to open when a workbook gallery is opened in viewer mode. Defaults to 0.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Specifies the name of the Resource Group where the Application Insights Workbook Template should exist. Changing this forces a new Application Insights Workbook Template to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Application Insights Workbook Template.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Valid JSON object containing workbook template payload.
	// +kubebuilder:validation:Optional
	TemplateData *string `json:"templateData,omitempty" tf:"template_data,omitempty"`
}

type GalleriesInitParameters struct {

	// Category for the gallery.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// Name of the workbook template in the gallery.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Order of the template within the gallery. Defaults to 0.
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// Azure resource type supported by the gallery. Defaults to Azure Monitor.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// Type of workbook supported by the workbook template. Defaults to workbook.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type GalleriesObservation struct {

	// Category for the gallery.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// Name of the workbook template in the gallery.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Order of the template within the gallery. Defaults to 0.
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// Azure resource type supported by the gallery. Defaults to Azure Monitor.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// Type of workbook supported by the workbook template. Defaults to workbook.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type GalleriesParameters struct {

	// Category for the gallery.
	// +kubebuilder:validation:Optional
	Category *string `json:"category" tf:"category,omitempty"`

	// Name of the workbook template in the gallery.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Order of the template within the gallery. Defaults to 0.
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// Azure resource type supported by the gallery. Defaults to Azure Monitor.
	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// Type of workbook supported by the workbook template. Defaults to workbook.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// ApplicationInsightsWorkbookTemplateSpec defines the desired state of ApplicationInsightsWorkbookTemplate
type ApplicationInsightsWorkbookTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationInsightsWorkbookTemplateParameters `json:"forProvider"`
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
	InitProvider ApplicationInsightsWorkbookTemplateInitParameters `json:"initProvider,omitempty"`
}

// ApplicationInsightsWorkbookTemplateStatus defines the observed state of ApplicationInsightsWorkbookTemplate.
type ApplicationInsightsWorkbookTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationInsightsWorkbookTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationInsightsWorkbookTemplate is the Schema for the ApplicationInsightsWorkbookTemplates API. Manages an Application Insights Workbook Template.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApplicationInsightsWorkbookTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.galleries) || (has(self.initProvider) && has(self.initProvider.galleries))",message="spec.forProvider.galleries is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.templateData) || (has(self.initProvider) && has(self.initProvider.templateData))",message="spec.forProvider.templateData is a required parameter"
	Spec   ApplicationInsightsWorkbookTemplateSpec   `json:"spec"`
	Status ApplicationInsightsWorkbookTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationInsightsWorkbookTemplateList contains a list of ApplicationInsightsWorkbookTemplates
type ApplicationInsightsWorkbookTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationInsightsWorkbookTemplate `json:"items"`
}

// Repository type metadata.
var (
	ApplicationInsightsWorkbookTemplate_Kind             = "ApplicationInsightsWorkbookTemplate"
	ApplicationInsightsWorkbookTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationInsightsWorkbookTemplate_Kind}.String()
	ApplicationInsightsWorkbookTemplate_KindAPIVersion   = ApplicationInsightsWorkbookTemplate_Kind + "." + CRDGroupVersion.String()
	ApplicationInsightsWorkbookTemplate_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationInsightsWorkbookTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationInsightsWorkbookTemplate{}, &ApplicationInsightsWorkbookTemplateList{})
}

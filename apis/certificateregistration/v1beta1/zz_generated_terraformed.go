// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this AppServiceCertificateOrder
func (mg *AppServiceCertificateOrder) GetTerraformResourceType() string {
	return "azurerm_app_service_certificate_order"
}

// GetConnectionDetailsMapping for this AppServiceCertificateOrder
func (tr *AppServiceCertificateOrder) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AppServiceCertificateOrder
func (tr *AppServiceCertificateOrder) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AppServiceCertificateOrder
func (tr *AppServiceCertificateOrder) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AppServiceCertificateOrder
func (tr *AppServiceCertificateOrder) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AppServiceCertificateOrder
func (tr *AppServiceCertificateOrder) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AppServiceCertificateOrder
func (tr *AppServiceCertificateOrder) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this AppServiceCertificateOrder
func (tr *AppServiceCertificateOrder) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this AppServiceCertificateOrder using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AppServiceCertificateOrder) LateInitialize(attrs []byte) (bool, error) {
	params := &AppServiceCertificateOrderParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}
	opts = append(opts, resource.WithNameFilter("Csr"))

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AppServiceCertificateOrder) GetTerraformSchemaVersion() int {
	return 0
}

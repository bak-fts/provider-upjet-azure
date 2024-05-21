// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	rconfig "github.com/upbound/provider-azure/apis/rconfig"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this API.
	apisresolver "github.com/upbound/provider-azure/internal/apis"
)

func (mg *API) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apimanagement.azure.upbound.io", "v1beta2", "Management", "ManagementList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIManagementName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.APIManagementNameRef,
			Selector:     mg.Spec.ForProvider.APIManagementNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIManagementName")
	}
	mg.Spec.ForProvider.APIManagementName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIManagementNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this APIDiagnostic.
func (mg *APIDiagnostic) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apimanagement.azure.upbound.io", "v1beta2", "Logger", "LoggerList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIManagementLoggerID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.APIManagementLoggerIDRef,
			Selector:     mg.Spec.ForProvider.APIManagementLoggerIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIManagementLoggerID")
	}
	mg.Spec.ForProvider.APIManagementLoggerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIManagementLoggerIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apimanagement.azure.upbound.io", "v1beta2", "Management", "ManagementList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIManagementName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.APIManagementNameRef,
			Selector:     mg.Spec.ForProvider.APIManagementNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIManagementName")
	}
	mg.Spec.ForProvider.APIManagementName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIManagementNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apimanagement.azure.upbound.io", "v1beta2", "API", "APIList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.APINameRef,
			Selector:     mg.Spec.ForProvider.APINameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIName")
	}
	mg.Spec.ForProvider.APIName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APINameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apimanagement.azure.upbound.io", "v1beta2", "Logger", "LoggerList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.APIManagementLoggerID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.APIManagementLoggerIDRef,
			Selector:     mg.Spec.InitProvider.APIManagementLoggerIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.APIManagementLoggerID")
	}
	mg.Spec.InitProvider.APIManagementLoggerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.APIManagementLoggerIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this APIOperation.
func (mg *APIOperation) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apimanagement.azure.upbound.io", "v1beta2", "Management", "ManagementList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIManagementName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.APIManagementNameRef,
			Selector:     mg.Spec.ForProvider.APIManagementNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIManagementName")
	}
	mg.Spec.ForProvider.APIManagementName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIManagementNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apimanagement.azure.upbound.io", "v1beta2", "API", "APIList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.APINameRef,
			Selector:     mg.Spec.ForProvider.APINameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIName")
	}
	mg.Spec.ForProvider.APIName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APINameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Backend.
func (mg *Backend) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apimanagement.azure.upbound.io", "v1beta2", "Management", "ManagementList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIManagementName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.APIManagementNameRef,
			Selector:     mg.Spec.ForProvider.APIManagementNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIManagementName")
	}
	mg.Spec.ForProvider.APIManagementName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIManagementNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Diagnostic.
func (mg *Diagnostic) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apimanagement.azure.upbound.io", "v1beta2", "Logger", "LoggerList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIManagementLoggerID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.APIManagementLoggerIDRef,
			Selector:     mg.Spec.ForProvider.APIManagementLoggerIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIManagementLoggerID")
	}
	mg.Spec.ForProvider.APIManagementLoggerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIManagementLoggerIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apimanagement.azure.upbound.io", "v1beta2", "Management", "ManagementList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIManagementName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.APIManagementNameRef,
			Selector:     mg.Spec.ForProvider.APIManagementNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIManagementName")
	}
	mg.Spec.ForProvider.APIManagementName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIManagementNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apimanagement.azure.upbound.io", "v1beta2", "Logger", "LoggerList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.APIManagementLoggerID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.APIManagementLoggerIDRef,
			Selector:     mg.Spec.InitProvider.APIManagementLoggerIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.APIManagementLoggerID")
	}
	mg.Spec.InitProvider.APIManagementLoggerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.APIManagementLoggerIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Gateway.
func (mg *Gateway) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apimanagement.azure.upbound.io", "v1beta2", "Management", "ManagementList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIManagementID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.APIManagementIDRef,
			Selector:     mg.Spec.ForProvider.APIManagementIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIManagementID")
	}
	mg.Spec.ForProvider.APIManagementID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIManagementIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apimanagement.azure.upbound.io", "v1beta2", "Management", "ManagementList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.APIManagementID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.APIManagementIDRef,
			Selector:     mg.Spec.InitProvider.APIManagementIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.APIManagementID")
	}
	mg.Spec.InitProvider.APIManagementID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.APIManagementIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Logger.
func (mg *Logger) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apimanagement.azure.upbound.io", "v1beta2", "Management", "ManagementList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIManagementName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.APIManagementNameRef,
			Selector:     mg.Spec.ForProvider.APIManagementNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIManagementName")
	}
	mg.Spec.ForProvider.APIManagementName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIManagementNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("insights.azure.upbound.io", "v1beta1", "ApplicationInsights", "ApplicationInsightsList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ResourceIDRef,
			Selector:     mg.Spec.ForProvider.ResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceID")
	}
	mg.Spec.ForProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("insights.azure.upbound.io", "v1beta1", "ApplicationInsights", "ApplicationInsightsList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ResourceIDRef,
			Selector:     mg.Spec.InitProvider.ResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceID")
	}
	mg.Spec.InitProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Management.
func (mg *Management) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.AdditionalLocation); i3++ {
		if mg.Spec.ForProvider.AdditionalLocation[i3].VirtualNetworkConfiguration != nil {
			{
				m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AdditionalLocation[i3].VirtualNetworkConfiguration.SubnetID),
					Extract:      rconfig.ExtractResourceID(),
					Reference:    mg.Spec.ForProvider.AdditionalLocation[i3].VirtualNetworkConfiguration.SubnetIDRef,
					Selector:     mg.Spec.ForProvider.AdditionalLocation[i3].VirtualNetworkConfiguration.SubnetIDSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.AdditionalLocation[i3].VirtualNetworkConfiguration.SubnetID")
			}
			mg.Spec.ForProvider.AdditionalLocation[i3].VirtualNetworkConfiguration.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.AdditionalLocation[i3].VirtualNetworkConfiguration.SubnetIDRef = rsp.ResolvedReference

		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	if mg.Spec.ForProvider.VirtualNetworkConfiguration != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualNetworkConfiguration.SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.VirtualNetworkConfiguration.SubnetIDRef,
				Selector:     mg.Spec.ForProvider.VirtualNetworkConfiguration.SubnetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VirtualNetworkConfiguration.SubnetID")
		}
		mg.Spec.ForProvider.VirtualNetworkConfiguration.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.VirtualNetworkConfiguration.SubnetIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.AdditionalLocation); i3++ {
		if mg.Spec.InitProvider.AdditionalLocation[i3].VirtualNetworkConfiguration != nil {
			{
				m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AdditionalLocation[i3].VirtualNetworkConfiguration.SubnetID),
					Extract:      rconfig.ExtractResourceID(),
					Reference:    mg.Spec.InitProvider.AdditionalLocation[i3].VirtualNetworkConfiguration.SubnetIDRef,
					Selector:     mg.Spec.InitProvider.AdditionalLocation[i3].VirtualNetworkConfiguration.SubnetIDSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.AdditionalLocation[i3].VirtualNetworkConfiguration.SubnetID")
			}
			mg.Spec.InitProvider.AdditionalLocation[i3].VirtualNetworkConfiguration.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.AdditionalLocation[i3].VirtualNetworkConfiguration.SubnetIDRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.VirtualNetworkConfiguration != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VirtualNetworkConfiguration.SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.VirtualNetworkConfiguration.SubnetIDRef,
				Selector:     mg.Spec.InitProvider.VirtualNetworkConfiguration.SubnetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.VirtualNetworkConfiguration.SubnetID")
		}
		mg.Spec.InitProvider.VirtualNetworkConfiguration.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.VirtualNetworkConfiguration.SubnetIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this NamedValue.
func (mg *NamedValue) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apimanagement.azure.upbound.io", "v1beta2", "Management", "ManagementList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIManagementName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.APIManagementNameRef,
			Selector:     mg.Spec.ForProvider.APIManagementNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIManagementName")
	}
	mg.Spec.ForProvider.APIManagementName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIManagementNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}
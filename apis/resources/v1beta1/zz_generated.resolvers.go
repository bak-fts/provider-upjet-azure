// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	rconfig "github.com/upbound/provider-azure/apis/rconfig"
	apisresolver "github.com/upbound/provider-azure/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *ResourceDeploymentScriptAzureCli) ResolveReferences( // ResolveReferences of this ResourceDeploymentScriptAzureCli.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Identity); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("managedidentity.azure.upbound.io", "v1beta1", "UserAssignedIdentity", "UserAssignedIdentityList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Identity[i3].IdentityIds),
				Extract:       rconfig.ExtractResourceID(),
				References:    mg.Spec.ForProvider.Identity[i3].IdentityIdsRefs,
				Selector:      mg.Spec.ForProvider.Identity[i3].IdentityIdsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Identity[i3].IdentityIds")
		}
		mg.Spec.ForProvider.Identity[i3].IdentityIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.Identity[i3].IdentityIdsRefs = mrsp.ResolvedReferences

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

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Identity); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("managedidentity.azure.upbound.io", "v1beta1", "UserAssignedIdentity", "UserAssignedIdentityList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Identity[i3].IdentityIds),
				Extract:       rconfig.ExtractResourceID(),
				References:    mg.Spec.InitProvider.Identity[i3].IdentityIdsRefs,
				Selector:      mg.Spec.InitProvider.Identity[i3].IdentityIdsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Identity[i3].IdentityIds")
		}
		mg.Spec.InitProvider.Identity[i3].IdentityIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.Identity[i3].IdentityIdsRefs = mrsp.ResolvedReferences

	}
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.InitProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceGroupName")
	}
	mg.Spec.InitProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ResourceDeploymentScriptAzurePowerShell.
func (mg *ResourceDeploymentScriptAzurePowerShell) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Identity); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("managedidentity.azure.upbound.io", "v1beta1", "UserAssignedIdentity", "UserAssignedIdentityList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Identity[i3].IdentityIds),
				Extract:       rconfig.ExtractResourceID(),
				References:    mg.Spec.ForProvider.Identity[i3].IdentityIdsRefs,
				Selector:      mg.Spec.ForProvider.Identity[i3].IdentityIdsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Identity[i3].IdentityIds")
		}
		mg.Spec.ForProvider.Identity[i3].IdentityIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.Identity[i3].IdentityIdsRefs = mrsp.ResolvedReferences

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

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Identity); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("managedidentity.azure.upbound.io", "v1beta1", "UserAssignedIdentity", "UserAssignedIdentityList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Identity[i3].IdentityIds),
				Extract:       rconfig.ExtractResourceID(),
				References:    mg.Spec.InitProvider.Identity[i3].IdentityIdsRefs,
				Selector:      mg.Spec.InitProvider.Identity[i3].IdentityIdsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Identity[i3].IdentityIds")
		}
		mg.Spec.InitProvider.Identity[i3].IdentityIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.Identity[i3].IdentityIdsRefs = mrsp.ResolvedReferences

	}
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.InitProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceGroupName")
	}
	mg.Spec.InitProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
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

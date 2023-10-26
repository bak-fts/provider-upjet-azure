/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1beta12 "github.com/upbound/provider-azure/apis/azure/v1beta1"
	v1beta11 "github.com/upbound/provider-azure/apis/keyvault/v1beta1"
	v1beta1 "github.com/upbound/provider-azure/apis/managedidentity/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Configuration.
func (mg *Configuration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Encryption); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Encryption[i3].IdentityClientID),
			Extract:      resource.ExtractParamPath("client_id", true),
			Reference:    mg.Spec.ForProvider.Encryption[i3].IdentityClientIDRef,
			Selector:     mg.Spec.ForProvider.Encryption[i3].IdentityClientIDSelector,
			To: reference.To{
				List:    &v1beta1.UserAssignedIdentityList{},
				Managed: &v1beta1.UserAssignedIdentity{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Encryption[i3].IdentityClientID")
		}
		mg.Spec.ForProvider.Encryption[i3].IdentityClientID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Encryption[i3].IdentityClientIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Encryption); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Encryption[i3].KeyVaultKeyIdentifier),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.Encryption[i3].KeyVaultKeyIdentifierRef,
			Selector:     mg.Spec.ForProvider.Encryption[i3].KeyVaultKeyIdentifierSelector,
			To: reference.To{
				List:    &v1beta11.KeyList{},
				Managed: &v1beta11.Key{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Encryption[i3].KeyVaultKeyIdentifier")
		}
		mg.Spec.ForProvider.Encryption[i3].KeyVaultKeyIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Encryption[i3].KeyVaultKeyIdentifierRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta12.ResourceGroupList{},
			Managed: &v1beta12.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

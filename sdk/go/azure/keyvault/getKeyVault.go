// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keyvault

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Key Vault.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := keyvault.LookupKeyVault(ctx, &keyvault.LookupKeyVaultArgs{
// 			Name:              "mykeyvault",
// 			ResourceGroupName: "some-resource-group",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("vaultUri", example.VaultUri)
// 		return nil
// 	})
// }
// ```
func LookupKeyVault(ctx *pulumi.Context, args *LookupKeyVaultArgs, opts ...pulumi.InvokeOption) (*LookupKeyVaultResult, error) {
	var rv LookupKeyVaultResult
	err := ctx.Invoke("azure:keyvault/getKeyVault:getKeyVault", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKeyVault.
type LookupKeyVaultArgs struct {
	// Specifies the name of the Key Vault.
	Name string `pulumi:"name"`
	// The name of the Resource Group in which the Key Vault exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getKeyVault.
type LookupKeyVaultResult struct {
	// One or more `accessPolicy` blocks as defined below.
	AccessPolicies []GetKeyVaultAccessPolicy `pulumi:"accessPolicies"`
	// Can Azure Virtual Machines retrieve certificates stored as secrets from the Key Vault?
	EnabledForDeployment bool `pulumi:"enabledForDeployment"`
	// Can Azure Disk Encryption retrieve secrets from the Key Vault?
	EnabledForDiskEncryption bool `pulumi:"enabledForDiskEncryption"`
	// Can Azure Resource Manager retrieve secrets from the Key Vault?
	EnabledForTemplateDeployment bool `pulumi:"enabledForTemplateDeployment"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Azure Region in which the Key Vault exists.
	Location    string                  `pulumi:"location"`
	Name        string                  `pulumi:"name"`
	NetworkAcls []GetKeyVaultNetworkAcl `pulumi:"networkAcls"`
	// Is purge protection enabled on this Key Vault?
	PurgeProtectionEnabled bool   `pulumi:"purgeProtectionEnabled"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	// The Name of the SKU used for this Key Vault.
	SkuName string `pulumi:"skuName"`
	// Is soft delete enabled on this Key Vault?
	SoftDeleteEnabled bool `pulumi:"softDeleteEnabled"`
	// A mapping of tags assigned to the Key Vault.
	Tags map[string]string `pulumi:"tags"`
	// The Azure Active Directory Tenant ID used to authenticate requests for this Key Vault.
	TenantId string `pulumi:"tenantId"`
	// The URI of the vault for performing operations on keys and secrets.
	VaultUri string `pulumi:"vaultUri"`
}

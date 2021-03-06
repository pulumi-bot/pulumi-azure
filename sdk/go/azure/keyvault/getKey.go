// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keyvault

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Key Vault Key.
func LookupKey(ctx *pulumi.Context, args *LookupKeyArgs, opts ...pulumi.InvokeOption) (*LookupKeyResult, error) {
	var rv LookupKeyResult
	err := ctx.Invoke("azure:keyvault/getKey:getKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKey.
type LookupKeyArgs struct {
	// Specifies the ID of the Key Vault instance where the Secret resides, available on the `keyvault.KeyVault` Data Source / Resource.
	KeyVaultId string `pulumi:"keyVaultId"`
	// Specifies the name of the Key Vault Key.
	Name string `pulumi:"name"`
}

// A collection of values returned by getKey.
type LookupKeyResult struct {
	// The RSA public exponent of this Key Vault Key.
	E string `pulumi:"e"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of JSON web key operations assigned to this Key Vault Key
	KeyOpts []string `pulumi:"keyOpts"`
	// Specifies the Size of this Key Vault Key.
	KeySize int `pulumi:"keySize"`
	// Specifies the Key Type of this Key Vault Key
	KeyType    string `pulumi:"keyType"`
	KeyVaultId string `pulumi:"keyVaultId"`
	// The RSA modulus of this Key Vault Key.
	N    string `pulumi:"n"`
	Name string `pulumi:"name"`
	// A mapping of tags assigned to this Key Vault Key.
	Tags map[string]string `pulumi:"tags"`
	// The current version of the Key Vault Key.
	Version string `pulumi:"version"`
}

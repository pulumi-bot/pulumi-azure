// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keyvault

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Key Vault Certificate.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/keyvault"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleKeyVault, err := keyvault.LookupKeyVault(ctx, &keyvault.LookupKeyVaultArgs{
// 			Name:              "examplekv",
// 			ResourceGroupName: "some-resource-group",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleCertificate, err := keyvault.LookupCertificate(ctx, &keyvault.LookupCertificateArgs{
// 			Name:       "secret-sauce",
// 			KeyVaultId: exampleKeyVault.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("certificateThumbprint", exampleCertificate.Thumbprint)
// 		return nil
// 	})
// }
// ```
func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	var rv LookupCertificateResult
	err := ctx.Invoke("azure:keyvault/getCertificate:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificate.
type LookupCertificateArgs struct {
	// Specifies the ID of the Key Vault instance where the Secret resides, available on the `keyvault.KeyVault` Data Source / Resource.
	KeyVaultId string `pulumi:"keyVaultId"`
	// Specifies the name of the Key Vault Certificate.
	Name string `pulumi:"name"`
	// Specifies the version of the certificate to look up.  (Defaults to latest)
	Version *string `pulumi:"version"`
}

// A collection of values returned by getCertificate.
type LookupCertificateResult struct {
	// The raw Key Vault Certificate data represented as a hexadecimal string.
	CertificateData string `pulumi:"certificateData"`
	// The raw Key Vault Certificate data represented as a base64 string.
	CertificateDataBase64 string `pulumi:"certificateDataBase64"`
	// A `certificatePolicy` block as defined below.
	CertificatePolicies []GetCertificateCertificatePolicy `pulumi:"certificatePolicies"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	KeyVaultId string `pulumi:"keyVaultId"`
	// The name of the Certificate Issuer.
	Name string `pulumi:"name"`
	// The ID of the associated Key Vault Secret.
	SecretId string `pulumi:"secretId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.
	Thumbprint string `pulumi:"thumbprint"`
	// The current version of the Key Vault Certificate.
	Version string `pulumi:"version"`
}

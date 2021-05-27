// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keyvault

import (
	"context"
	"reflect"

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
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/keyvault"
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

func LookupCertificateOutput(ctx *pulumi.Context, args LookupCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCertificateResult, error) {
			args := v.(LookupCertificateArgs)
			r, err := LookupCertificate(ctx, &args, opts...)
			return *r, err
		}).(LookupCertificateResultOutput)
}

// A collection of arguments for invoking getCertificate.
type LookupCertificateOutputArgs struct {
	// Specifies the ID of the Key Vault instance where the Secret resides, available on the `keyvault.KeyVault` Data Source / Resource.
	KeyVaultId pulumi.StringInput `pulumi:"keyVaultId"`
	// Specifies the name of the Key Vault Certificate.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the version of the certificate to look up.  (Defaults to latest)
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (LookupCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateArgs)(nil)).Elem()
}

// A collection of values returned by getCertificate.
type LookupCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateResult)(nil)).Elem()
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutput() LookupCertificateResultOutput {
	return o
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutputWithContext(ctx context.Context) LookupCertificateResultOutput {
	return o
}

// The raw Key Vault Certificate data represented as a hexadecimal string.
func (o LookupCertificateResultOutput) CertificateData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.CertificateData }).(pulumi.StringOutput)
}

// The raw Key Vault Certificate data represented as a base64 string.
func (o LookupCertificateResultOutput) CertificateDataBase64() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.CertificateDataBase64 }).(pulumi.StringOutput)
}

// A `certificatePolicy` block as defined below.
func (o LookupCertificateResultOutput) CertificatePolicies() GetCertificateCertificatePolicyArrayOutput {
	return o.ApplyT(func(v LookupCertificateResult) []GetCertificateCertificatePolicy { return v.CertificatePolicies }).(GetCertificateCertificatePolicyArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) KeyVaultId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.KeyVaultId }).(pulumi.StringOutput)
}

// The name of the Certificate Issuer.
func (o LookupCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

// The ID of the associated Key Vault Secret.
func (o LookupCertificateResultOutput) SecretId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.SecretId }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o LookupCertificateResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCertificateResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.
func (o LookupCertificateResultOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Thumbprint }).(pulumi.StringOutput)
}

// The current version of the Key Vault Certificate.
func (o LookupCertificateResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCertificateResultOutput{})
}

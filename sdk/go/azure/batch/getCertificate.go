// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package batch

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing certificate in a Batch Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/batch"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := batch.LookupCertificate(ctx, &batch.LookupCertificateArgs{
// 			Name:              "SHA1-42C107874FD0E4A9583292A2F1098E8FE4B2EDDA",
// 			AccountName:       "examplebatchaccount",
// 			ResourceGroupName: "example",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("thumbprint", example.Thumbprint)
// 		return nil
// 	})
// }
// ```
func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	var rv LookupCertificateResult
	err := ctx.Invoke("azure:batch/getCertificate:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificate.
type LookupCertificateArgs struct {
	// The name of the Batch account.
	AccountName string `pulumi:"accountName"`
	// The name of the Batch certificate.
	Name string `pulumi:"name"`
	// The Name of the Resource Group where this Batch account exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getCertificate.
type LookupCertificateResult struct {
	AccountName string `pulumi:"accountName"`
	// The format of the certificate, such as `Cer` or `Pfx`.
	Format string `pulumi:"format"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The public key of the certificate.
	PublicData        string `pulumi:"publicData"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The thumbprint of the certificate.
	Thumbprint string `pulumi:"thumbprint"`
	// The algorithm of the certificate thumbprint.
	ThumbprintAlgorithm string `pulumi:"thumbprintAlgorithm"`
}

func LookupCertificateApply(ctx *pulumi.Context, args LookupCertificateApplyInput, opts ...pulumi.InvokeOption) LookupCertificateResultOutput {
	return args.ToLookupCertificateApplyOutput().ApplyT(func(v LookupCertificateArgs) (LookupCertificateResult, error) {
		r, err := LookupCertificate(ctx, &v, opts...)
		return *r, err

	}).(LookupCertificateResultOutput)
}

// LookupCertificateApplyInput is an input type that accepts LookupCertificateApplyArgs and LookupCertificateApplyOutput values.
// You can construct a concrete instance of `LookupCertificateApplyInput` via:
//
//          LookupCertificateApplyArgs{...}
type LookupCertificateApplyInput interface {
	pulumi.Input

	ToLookupCertificateApplyOutput() LookupCertificateApplyOutput
	ToLookupCertificateApplyOutputWithContext(context.Context) LookupCertificateApplyOutput
}

// A collection of arguments for invoking getCertificate.
type LookupCertificateApplyArgs struct {
	// The name of the Batch account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the Batch certificate.
	Name pulumi.StringInput `pulumi:"name"`
	// The Name of the Resource Group where this Batch account exists.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCertificateApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateArgs)(nil)).Elem()
}

func (i LookupCertificateApplyArgs) ToLookupCertificateApplyOutput() LookupCertificateApplyOutput {
	return i.ToLookupCertificateApplyOutputWithContext(context.Background())
}

func (i LookupCertificateApplyArgs) ToLookupCertificateApplyOutputWithContext(ctx context.Context) LookupCertificateApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupCertificateApplyOutput)
}

// A collection of arguments for invoking getCertificate.
type LookupCertificateApplyOutput struct{ *pulumi.OutputState }

func (LookupCertificateApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateArgs)(nil)).Elem()
}

func (o LookupCertificateApplyOutput) ToLookupCertificateApplyOutput() LookupCertificateApplyOutput {
	return o
}

func (o LookupCertificateApplyOutput) ToLookupCertificateApplyOutputWithContext(ctx context.Context) LookupCertificateApplyOutput {
	return o
}

// The name of the Batch account.
func (o LookupCertificateApplyOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateArgs) string { return v.AccountName }).(pulumi.StringOutput)
}

// The name of the Batch certificate.
func (o LookupCertificateApplyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateArgs) string { return v.Name }).(pulumi.StringOutput)
}

// The Name of the Resource Group where this Batch account exists.
func (o LookupCertificateApplyOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateArgs) string { return v.ResourceGroupName }).(pulumi.StringOutput)
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

func (o LookupCertificateResultOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.AccountName }).(pulumi.StringOutput)
}

// The format of the certificate, such as `Cer` or `Pfx`.
func (o LookupCertificateResultOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Format }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

// The public key of the certificate.
func (o LookupCertificateResultOutput) PublicData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.PublicData }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// The thumbprint of the certificate.
func (o LookupCertificateResultOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Thumbprint }).(pulumi.StringOutput)
}

// The algorithm of the certificate thumbprint.
func (o LookupCertificateResultOutput) ThumbprintAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.ThumbprintAlgorithm }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCertificateApplyOutput{})
	pulumi.RegisterOutputType(LookupCertificateResultOutput{})
}

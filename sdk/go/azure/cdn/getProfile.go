// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing CDN Profile.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/cdn"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := cdn.LookupProfile(ctx, &cdn.LookupProfileArgs{
// 			Name:              "myfirstcdnprofile",
// 			ResourceGroupName: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("cdnProfileId", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupProfile(ctx *pulumi.Context, args *LookupProfileArgs, opts ...pulumi.InvokeOption) (*LookupProfileResult, error) {
	var rv LookupProfileResult
	err := ctx.Invoke("azure:cdn/getProfile:getProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProfile.
type LookupProfileArgs struct {
	// The name of the CDN Profile.
	Name string `pulumi:"name"`
	// The name of the resource group in which the CDN Profile exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getProfile.
type LookupProfileResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Azure Region where the resource exists.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The pricing related information of current CDN profile.
	Sku string `pulumi:"sku"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}

func LookupProfileOutput(ctx *pulumi.Context, args LookupProfileOutputArgs, opts ...pulumi.InvokeOption) LookupProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProfileResult, error) {
			args := v.(LookupProfileArgs)
			r, err := LookupProfile(ctx, &args, opts...)
			return *r, err
		}).(LookupProfileResultOutput)
}

// A collection of arguments for invoking getProfile.
type LookupProfileOutputArgs struct {
	// The name of the CDN Profile.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which the CDN Profile exists.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfileArgs)(nil)).Elem()
}

// A collection of values returned by getProfile.
type LookupProfileResultOutput struct{ *pulumi.OutputState }

func (LookupProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfileResult)(nil)).Elem()
}

func (o LookupProfileResultOutput) ToLookupProfileResultOutput() LookupProfileResultOutput {
	return o
}

func (o LookupProfileResultOutput) ToLookupProfileResultOutputWithContext(ctx context.Context) LookupProfileResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Azure Region where the resource exists.
func (o LookupProfileResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProfileResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// The pricing related information of current CDN profile.
func (o LookupProfileResultOutput) Sku() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.Sku }).(pulumi.StringOutput)
}

// A mapping of tags assigned to the resource.
func (o LookupProfileResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProfileResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProfileResultOutput{})
}

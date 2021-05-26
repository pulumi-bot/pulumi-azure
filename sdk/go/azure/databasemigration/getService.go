// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databasemigration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Database Migration Service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/databasemigration"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := databasemigration.LookupService(ctx, &databasemigration.LookupServiceArgs{
// 			Name:              "example-dms",
// 			ResourceGroupName: "example-rg",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("azurermDmsId", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azure:databasemigration/getService:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getService.
type LookupServiceArgs struct {
	// Specify the name of the database migration service.
	Name string `pulumi:"name"`
	// Specifies the Name of the Resource Group within which the database migration service exists
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getService.
type LookupServiceResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Azure location where the resource exists.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku name of database migration service.
	SkuName string `pulumi:"skuName"`
	// The ID of the virtual subnet resource to which the database migration service exists.
	SubnetId string `pulumi:"subnetId"`
	// A mapping of tags to assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}

func LookupServiceApply(ctx *pulumi.Context, args LookupServiceApplyInput, opts ...pulumi.InvokeOption) LookupServiceResultOutput {
	return args.ToLookupServiceApplyOutput().ApplyT(func(v LookupServiceArgs) (LookupServiceResult, error) {
		r, err := LookupService(ctx, &v, opts...)
		return *r, err

	}).(LookupServiceResultOutput)
}

// LookupServiceApplyInput is an input type that accepts LookupServiceApplyArgs and LookupServiceApplyOutput values.
// You can construct a concrete instance of `LookupServiceApplyInput` via:
//
//          LookupServiceApplyArgs{...}
type LookupServiceApplyInput interface {
	pulumi.Input

	ToLookupServiceApplyOutput() LookupServiceApplyOutput
	ToLookupServiceApplyOutputWithContext(context.Context) LookupServiceApplyOutput
}

// A collection of arguments for invoking getService.
type LookupServiceApplyArgs struct {
	// Specify the name of the database migration service.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the Name of the Resource Group within which the database migration service exists
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupServiceApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceArgs)(nil)).Elem()
}

func (i LookupServiceApplyArgs) ToLookupServiceApplyOutput() LookupServiceApplyOutput {
	return i.ToLookupServiceApplyOutputWithContext(context.Background())
}

func (i LookupServiceApplyArgs) ToLookupServiceApplyOutputWithContext(ctx context.Context) LookupServiceApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupServiceApplyOutput)
}

// A collection of arguments for invoking getService.
type LookupServiceApplyOutput struct{ *pulumi.OutputState }

func (LookupServiceApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceArgs)(nil)).Elem()
}

func (o LookupServiceApplyOutput) ToLookupServiceApplyOutput() LookupServiceApplyOutput {
	return o
}

func (o LookupServiceApplyOutput) ToLookupServiceApplyOutputWithContext(ctx context.Context) LookupServiceApplyOutput {
	return o
}

// Specify the name of the database migration service.
func (o LookupServiceApplyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceArgs) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the Name of the Resource Group within which the database migration service exists
func (o LookupServiceApplyOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceArgs) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// A collection of values returned by getService.
type LookupServiceResultOutput struct{ *pulumi.OutputState }

func (LookupServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceResult)(nil)).Elem()
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutput() LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutputWithContext(ctx context.Context) LookupServiceResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Azure location where the resource exists.
func (o LookupServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// The sku name of database migration service.
func (o LookupServiceResultOutput) SkuName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.SkuName }).(pulumi.StringOutput)
}

// The ID of the virtual subnet resource to which the database migration service exists.
func (o LookupServiceResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

// A mapping of tags to assigned to the resource.
func (o LookupServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceApplyOutput{})
	pulumi.RegisterOutputType(LookupServiceResultOutput{})
}

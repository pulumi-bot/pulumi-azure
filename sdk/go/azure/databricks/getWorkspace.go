// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databricks

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Databricks workspace.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/databricks"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := databricks.LookupWorkspace(ctx, &databricks.LookupWorkspaceArgs{
// 			Name:              "example-workspace",
// 			ResourceGroupName: "example-rg",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("databricksWorkspaceId", example.WorkspaceId)
// 		return nil
// 	})
// }
// ```
func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure:databricks/getWorkspace:getWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWorkspace.
type LookupWorkspaceArgs struct {
	// The name of the Databricks Workspace.
	Name string `pulumi:"name"`
	// The Name of the Resource Group where the Databricks Workspace exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Databricks Workspace.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getWorkspace.
type LookupWorkspaceResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                string `pulumi:"id"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SKU of this Databricks Workspace.
	Sku string `pulumi:"sku"`
	// A mapping of tags to assign to the Databricks Workspace.
	Tags map[string]string `pulumi:"tags"`
	// Unique ID of this Databricks Workspace in Databricks management plane.
	WorkspaceId string `pulumi:"workspaceId"`
	// URL this Databricks Workspace is accessible on.
	WorkspaceUrl string `pulumi:"workspaceUrl"`
}

func LookupWorkspaceApply(ctx *pulumi.Context, args LookupWorkspaceApplyInput, opts ...pulumi.InvokeOption) LookupWorkspaceResultOutput {
	return args.ToLookupWorkspaceApplyOutput().ApplyT(func(v LookupWorkspaceArgs) (LookupWorkspaceResult, error) {
		r, err := LookupWorkspace(ctx, &v, opts...)
		return *r, err

	}).(LookupWorkspaceResultOutput)
}

// LookupWorkspaceApplyInput is an input type that accepts LookupWorkspaceApplyArgs and LookupWorkspaceApplyOutput values.
// You can construct a concrete instance of `LookupWorkspaceApplyInput` via:
//
//          LookupWorkspaceApplyArgs{...}
type LookupWorkspaceApplyInput interface {
	pulumi.Input

	ToLookupWorkspaceApplyOutput() LookupWorkspaceApplyOutput
	ToLookupWorkspaceApplyOutputWithContext(context.Context) LookupWorkspaceApplyOutput
}

// A collection of arguments for invoking getWorkspace.
type LookupWorkspaceApplyArgs struct {
	// The name of the Databricks Workspace.
	Name pulumi.StringInput `pulumi:"name"`
	// The Name of the Resource Group where the Databricks Workspace exists.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Databricks Workspace.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupWorkspaceApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceArgs)(nil)).Elem()
}

func (i LookupWorkspaceApplyArgs) ToLookupWorkspaceApplyOutput() LookupWorkspaceApplyOutput {
	return i.ToLookupWorkspaceApplyOutputWithContext(context.Background())
}

func (i LookupWorkspaceApplyArgs) ToLookupWorkspaceApplyOutputWithContext(ctx context.Context) LookupWorkspaceApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupWorkspaceApplyOutput)
}

// A collection of arguments for invoking getWorkspace.
type LookupWorkspaceApplyOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceArgs)(nil)).Elem()
}

func (o LookupWorkspaceApplyOutput) ToLookupWorkspaceApplyOutput() LookupWorkspaceApplyOutput {
	return o
}

func (o LookupWorkspaceApplyOutput) ToLookupWorkspaceApplyOutputWithContext(ctx context.Context) LookupWorkspaceApplyOutput {
	return o
}

// The name of the Databricks Workspace.
func (o LookupWorkspaceApplyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceArgs) string { return v.Name }).(pulumi.StringOutput)
}

// The Name of the Resource Group where the Databricks Workspace exists.
func (o LookupWorkspaceApplyOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceArgs) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the Databricks Workspace.
func (o LookupWorkspaceApplyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkspaceArgs) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// A collection of values returned by getWorkspace.
type LookupWorkspaceResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceResult)(nil)).Elem()
}

func (o LookupWorkspaceResultOutput) ToLookupWorkspaceResultOutput() LookupWorkspaceResultOutput {
	return o
}

func (o LookupWorkspaceResultOutput) ToLookupWorkspaceResultOutputWithContext(ctx context.Context) LookupWorkspaceResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// SKU of this Databricks Workspace.
func (o LookupWorkspaceResultOutput) Sku() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Sku }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the Databricks Workspace.
func (o LookupWorkspaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Unique ID of this Databricks Workspace in Databricks management plane.
func (o LookupWorkspaceResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

// URL this Databricks Workspace is accessible on.
func (o LookupWorkspaceResultOutput) WorkspaceUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.WorkspaceUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceApplyOutput{})
	pulumi.RegisterOutputType(LookupWorkspaceResultOutput{})
}

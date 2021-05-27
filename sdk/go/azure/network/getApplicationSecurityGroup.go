// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Application Security Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := network.LookupApplicationSecurityGroup(ctx, &network.LookupApplicationSecurityGroupArgs{
// 			Name:              "tf-appsecuritygroup",
// 			ResourceGroupName: "my-resource-group",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("applicationSecurityGroupId", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupApplicationSecurityGroup(ctx *pulumi.Context, args *LookupApplicationSecurityGroupArgs, opts ...pulumi.InvokeOption) (*LookupApplicationSecurityGroupResult, error) {
	var rv LookupApplicationSecurityGroupResult
	err := ctx.Invoke("azure:network/getApplicationSecurityGroup:getApplicationSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApplicationSecurityGroup.
type LookupApplicationSecurityGroupArgs struct {
	// The name of the Application Security Group.
	Name string `pulumi:"name"`
	// The name of the resource group in which the Application Security Group exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getApplicationSecurityGroup.
type LookupApplicationSecurityGroupResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The supported Azure location where the Application Security Group exists.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}

func LookupApplicationSecurityGroupOutput(ctx *pulumi.Context, args LookupApplicationSecurityGroupOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationSecurityGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationSecurityGroupResult, error) {
			args := v.(LookupApplicationSecurityGroupArgs)
			r, err := LookupApplicationSecurityGroup(ctx, &args, opts...)
			return *r, err
		}).(LookupApplicationSecurityGroupResultOutput)
}

// A collection of arguments for invoking getApplicationSecurityGroup.
type LookupApplicationSecurityGroupOutputArgs struct {
	// The name of the Application Security Group.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which the Application Security Group exists.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationSecurityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationSecurityGroupArgs)(nil)).Elem()
}

// A collection of values returned by getApplicationSecurityGroup.
type LookupApplicationSecurityGroupResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationSecurityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationSecurityGroupResult)(nil)).Elem()
}

func (o LookupApplicationSecurityGroupResultOutput) ToLookupApplicationSecurityGroupResultOutput() LookupApplicationSecurityGroupResultOutput {
	return o
}

func (o LookupApplicationSecurityGroupResultOutput) ToLookupApplicationSecurityGroupResultOutputWithContext(ctx context.Context) LookupApplicationSecurityGroupResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupApplicationSecurityGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSecurityGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The supported Azure location where the Application Security Group exists.
func (o LookupApplicationSecurityGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSecurityGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupApplicationSecurityGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSecurityGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplicationSecurityGroupResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSecurityGroupResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// A mapping of tags assigned to the resource.
func (o LookupApplicationSecurityGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationSecurityGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationSecurityGroupResultOutput{})
}

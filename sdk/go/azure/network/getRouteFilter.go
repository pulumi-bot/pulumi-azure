// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Route Filter.
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
// 		example, err := network.LookupRouteFilter(ctx, &network.LookupRouteFilterArgs{
// 			Name:              "existing",
// 			ResourceGroupName: "existing",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupRouteFilter(ctx *pulumi.Context, args *LookupRouteFilterArgs, opts ...pulumi.InvokeOption) (*LookupRouteFilterResult, error) {
	var rv LookupRouteFilterResult
	err := ctx.Invoke("azure:network/getRouteFilter:getRouteFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouteFilter.
type LookupRouteFilterArgs struct {
	// The Name of this Route Filter.
	Name string `pulumi:"name"`
	// The name of the Resource Group where the Route Filter exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getRouteFilter.
type LookupRouteFilterResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Azure Region where the Route Filter exists.
	Location string `pulumi:"location"`
	// The Name of Route Filter Rule
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `rule` block as defined below.
	Rules []GetRouteFilterRule `pulumi:"rules"`
	// A mapping of tags assigned to the Route Filter.
	Tags map[string]string `pulumi:"tags"`
}

func LookupRouteFilterOutput(ctx *pulumi.Context, args LookupRouteFilterOutputArgs, opts ...pulumi.InvokeOption) LookupRouteFilterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteFilterResult, error) {
			args := v.(LookupRouteFilterArgs)
			r, err := LookupRouteFilter(ctx, &args, opts...)
			return *r, err
		}).(LookupRouteFilterResultOutput)
}

// A collection of arguments for invoking getRouteFilter.
type LookupRouteFilterOutputArgs struct {
	// The Name of this Route Filter.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the Resource Group where the Route Filter exists.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRouteFilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteFilterArgs)(nil)).Elem()
}

// A collection of values returned by getRouteFilter.
type LookupRouteFilterResultOutput struct{ *pulumi.OutputState }

func (LookupRouteFilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteFilterResult)(nil)).Elem()
}

func (o LookupRouteFilterResultOutput) ToLookupRouteFilterResultOutput() LookupRouteFilterResultOutput {
	return o
}

func (o LookupRouteFilterResultOutput) ToLookupRouteFilterResultOutputWithContext(ctx context.Context) LookupRouteFilterResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouteFilterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Azure Region where the Route Filter exists.
func (o LookupRouteFilterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) string { return v.Location }).(pulumi.StringOutput)
}

// The Name of Route Filter Rule
func (o LookupRouteFilterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRouteFilterResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// A `rule` block as defined below.
func (o LookupRouteFilterResultOutput) Rules() GetRouteFilterRuleArrayOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) []GetRouteFilterRule { return v.Rules }).(GetRouteFilterRuleArrayOutput)
}

// A mapping of tags assigned to the Route Filter.
func (o LookupRouteFilterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteFilterResultOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access the configuration of the AzureRM provider.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("accountId", current.ClientId)
// 		return nil
// 	})
// }
// ```
func GetClientConfig(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetClientConfigResult, error) {
	var rv GetClientConfigResult
	err := ctx.Invoke("azure:core/getClientConfig:getClientConfig", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getClientConfig.
type GetClientConfigResult struct {
	ClientId string `pulumi:"clientId"`
	// The provider-assigned unique ID for this managed resource.
	Id             string `pulumi:"id"`
	ObjectId       string `pulumi:"objectId"`
	SubscriptionId string `pulumi:"subscriptionId"`
	TenantId       string `pulumi:"tenantId"`
}

func GetClientConfigApply(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetClientConfigResultOutput {
	return pulumi.Any(opts).ApplyT(func(v interface{}) (GetClientConfigResult, error) {
		r, err := GetClientConfig(ctx, opts...)
		return *r, err

	}).(GetClientConfigResultOutput)
}

// A collection of values returned by getClientConfig.
type GetClientConfigResultOutput struct{ *pulumi.OutputState }

func (GetClientConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientConfigResult)(nil)).Elem()
}

func (o GetClientConfigResultOutput) ToGetClientConfigResultOutput() GetClientConfigResultOutput {
	return o
}

func (o GetClientConfigResultOutput) ToGetClientConfigResultOutputWithContext(ctx context.Context) GetClientConfigResultOutput {
	return o
}

func (o GetClientConfigResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientConfigResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetClientConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetClientConfigResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientConfigResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

func (o GetClientConfigResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientConfigResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o GetClientConfigResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientConfigResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClientConfigResultOutput{})
}

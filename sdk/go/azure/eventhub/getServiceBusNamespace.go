// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing ServiceBus Namespace.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/servicebus"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := servicebus.LookupNamespace(ctx, &servicebus.LookupNamespaceArgs{
// 			Name:              "examplenamespace",
// 			ResourceGroupName: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("location", example.Location)
// 		return nil
// 	})
// }
// ```
//
// Deprecated: azure.eventhub.getServiceBusNamespace has been deprecated in favor of azure.servicebus.getNamespace
func GetServiceBusNamespace(ctx *pulumi.Context, args *GetServiceBusNamespaceArgs, opts ...pulumi.InvokeOption) (*GetServiceBusNamespaceResult, error) {
	var rv GetServiceBusNamespaceResult
	err := ctx.Invoke("azure:eventhub/getServiceBusNamespace:getServiceBusNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceBusNamespace.
type GetServiceBusNamespaceArgs struct {
	// Specifies the name of the ServiceBus Namespace.
	Name string `pulumi:"name"`
	// Specifies the name of the Resource Group where the ServiceBus Namespace exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getServiceBusNamespace.
type GetServiceBusNamespaceResult struct {
	// The capacity of the ServiceBus Namespace.
	Capacity int `pulumi:"capacity"`
	// The primary connection string for the authorization
	// rule `RootManageSharedAccessKey`.
	DefaultPrimaryConnectionString string `pulumi:"defaultPrimaryConnectionString"`
	// The primary access key for the authorization rule `RootManageSharedAccessKey`.
	DefaultPrimaryKey string `pulumi:"defaultPrimaryKey"`
	// The secondary connection string for the
	// authorization rule `RootManageSharedAccessKey`.
	DefaultSecondaryConnectionString string `pulumi:"defaultSecondaryConnectionString"`
	// The secondary access key for the authorization rule `RootManageSharedAccessKey`.
	DefaultSecondaryKey string `pulumi:"defaultSecondaryKey"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The location of the Resource Group in which the ServiceBus Namespace exists.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Tier used for the ServiceBus Namespace.
	Sku string `pulumi:"sku"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Whether or not this ServiceBus Namespace is zone redundant.
	ZoneRedundant bool `pulumi:"zoneRedundant"`
}

func GetServiceBusNamespaceOutput(ctx *pulumi.Context, args GetServiceBusNamespaceOutputArgs, opts ...pulumi.InvokeOption) GetServiceBusNamespaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServiceBusNamespaceResult, error) {
			args := v.(GetServiceBusNamespaceArgs)
			r, err := GetServiceBusNamespace(ctx, &args, opts...)
			return *r, err
		}).(GetServiceBusNamespaceResultOutput)
}

// A collection of arguments for invoking getServiceBusNamespace.
type GetServiceBusNamespaceOutputArgs struct {
	// Specifies the name of the ServiceBus Namespace.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the name of the Resource Group where the ServiceBus Namespace exists.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetServiceBusNamespaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceBusNamespaceArgs)(nil)).Elem()
}

// A collection of values returned by getServiceBusNamespace.
type GetServiceBusNamespaceResultOutput struct{ *pulumi.OutputState }

func (GetServiceBusNamespaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceBusNamespaceResult)(nil)).Elem()
}

func (o GetServiceBusNamespaceResultOutput) ToGetServiceBusNamespaceResultOutput() GetServiceBusNamespaceResultOutput {
	return o
}

func (o GetServiceBusNamespaceResultOutput) ToGetServiceBusNamespaceResultOutputWithContext(ctx context.Context) GetServiceBusNamespaceResultOutput {
	return o
}

// The capacity of the ServiceBus Namespace.
func (o GetServiceBusNamespaceResultOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v GetServiceBusNamespaceResult) int { return v.Capacity }).(pulumi.IntOutput)
}

// The primary connection string for the authorization
// rule `RootManageSharedAccessKey`.
func (o GetServiceBusNamespaceResultOutput) DefaultPrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceBusNamespaceResult) string { return v.DefaultPrimaryConnectionString }).(pulumi.StringOutput)
}

// The primary access key for the authorization rule `RootManageSharedAccessKey`.
func (o GetServiceBusNamespaceResultOutput) DefaultPrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceBusNamespaceResult) string { return v.DefaultPrimaryKey }).(pulumi.StringOutput)
}

// The secondary connection string for the
// authorization rule `RootManageSharedAccessKey`.
func (o GetServiceBusNamespaceResultOutput) DefaultSecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceBusNamespaceResult) string { return v.DefaultSecondaryConnectionString }).(pulumi.StringOutput)
}

// The secondary access key for the authorization rule `RootManageSharedAccessKey`.
func (o GetServiceBusNamespaceResultOutput) DefaultSecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceBusNamespaceResult) string { return v.DefaultSecondaryKey }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServiceBusNamespaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceBusNamespaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the Resource Group in which the ServiceBus Namespace exists.
func (o GetServiceBusNamespaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceBusNamespaceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetServiceBusNamespaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceBusNamespaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetServiceBusNamespaceResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceBusNamespaceResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// The Tier used for the ServiceBus Namespace.
func (o GetServiceBusNamespaceResultOutput) Sku() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceBusNamespaceResult) string { return v.Sku }).(pulumi.StringOutput)
}

// A mapping of tags assigned to the resource.
func (o GetServiceBusNamespaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetServiceBusNamespaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Whether or not this ServiceBus Namespace is zone redundant.
func (o GetServiceBusNamespaceResultOutput) ZoneRedundant() pulumi.BoolOutput {
	return o.ApplyT(func(v GetServiceBusNamespaceResult) bool { return v.ZoneRedundant }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServiceBusNamespaceResultOutput{})
}

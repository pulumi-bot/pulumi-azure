// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
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
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/servicebus"
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

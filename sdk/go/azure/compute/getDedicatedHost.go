// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Dedicated Host.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := compute.LookupDedicatedHost(ctx, &compute.LookupDedicatedHostArgs{
// 			Name:                   "example-host",
// 			DedicatedHostGroupName: "example-host-group",
// 			ResourceGroupName:      "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("dedicatedHostId", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupDedicatedHost(ctx *pulumi.Context, args *LookupDedicatedHostArgs, opts ...pulumi.InvokeOption) (*LookupDedicatedHostResult, error) {
	var rv LookupDedicatedHostResult
	err := ctx.Invoke("azure:compute/getDedicatedHost:getDedicatedHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDedicatedHost.
type LookupDedicatedHostArgs struct {
	// Specifies the name of the Dedicated Host Group the Dedicated Host is located in.
	DedicatedHostGroupName string `pulumi:"dedicatedHostGroupName"`
	// Specifies the name of the Dedicated Host.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group the Dedicated Host is located in.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getDedicatedHost.
type LookupDedicatedHostResult struct {
	DedicatedHostGroupName string `pulumi:"dedicatedHostGroupName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The location where the Dedicated Host exists.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags assigned to the Dedicated Host.
	Tags map[string]string `pulumi:"tags"`
}

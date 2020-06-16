// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Route Table.
//
// ## Example Usage
//
//
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.LookupRouteTable(ctx, &network.LookupRouteTableArgs{
// 			Name:              "myroutetable",
// 			ResourceGroupName: "some-resource-group",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupRouteTable(ctx *pulumi.Context, args *LookupRouteTableArgs, opts ...pulumi.InvokeOption) (*LookupRouteTableResult, error) {
	var rv LookupRouteTableResult
	err := ctx.Invoke("azure:network/getRouteTable:getRouteTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouteTable.
type LookupRouteTableArgs struct {
	// The name of the Route Table.
	Name string `pulumi:"name"`
	// The name of the Resource Group in which the Route Table exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getRouteTable.
type LookupRouteTableResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Azure Region in which the Route Table exists.
	Location string `pulumi:"location"`
	// The name of the Route.
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// One or more `route` blocks as documented below.
	Routes []GetRouteTableRoute `pulumi:"routes"`
	// The collection of Subnets associated with this route table.
	Subnets []string `pulumi:"subnets"`
	// A mapping of tags assigned to the Route Table.
	Tags map[string]string `pulumi:"tags"`
}

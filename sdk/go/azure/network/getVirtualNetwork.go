// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Virtual Network.
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
// 		example, err := network.LookupVirtualNetwork(ctx, &network.LookupVirtualNetworkArgs{
// 			Name:              "production",
// 			ResourceGroupName: "networking",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("virtualNetworkId", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupVirtualNetwork(ctx *pulumi.Context, args *LookupVirtualNetworkArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkResult, error) {
	var rv LookupVirtualNetworkResult
	err := ctx.Invoke("azure:network/getVirtualNetwork:getVirtualNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualNetwork.
type LookupVirtualNetworkArgs struct {
	// Specifies the name of the Virtual Network.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group the Virtual Network is located in.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getVirtualNetwork.
type LookupVirtualNetworkResult struct {
	// The list of address spaces used by the virtual network.
	AddressSpaces []string `pulumi:"addressSpaces"`
	// The list of DNS servers used by the virtual network.
	DnsServers []string `pulumi:"dnsServers"`
	// The GUID of the virtual network.
	Guid string `pulumi:"guid"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Location of the virtual network.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The list of name of the subnets that are attached to this virtual network.
	Subnets []string `pulumi:"subnets"`
	// A mapping of name - virtual network id of the virtual network peerings.
	VnetPeerings map[string]string `pulumi:"vnetPeerings"`
}

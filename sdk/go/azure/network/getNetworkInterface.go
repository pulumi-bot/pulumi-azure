// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Network Interface.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := network.LookupNetworkInterface(ctx, &network.LookupNetworkInterfaceArgs{
// 			Name:              "acctest-nic",
// 			ResourceGroupName: "networking",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("networkInterfaceId", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupNetworkInterface(ctx *pulumi.Context, args *LookupNetworkInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupNetworkInterfaceResult, error) {
	var rv LookupNetworkInterfaceResult
	err := ctx.Invoke("azure:network/getNetworkInterface:getNetworkInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworkInterface.
type LookupNetworkInterfaceArgs struct {
	// Specifies the name of the Network Interface.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group the Network Interface is located in.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getNetworkInterface.
type LookupNetworkInterfaceResult struct {
	// List of DNS servers applied to the specified Network Interface.
	AppliedDnsServers []string `pulumi:"appliedDnsServers"`
	// The list of DNS servers used by the specified Network Interface.
	DnsServers []string `pulumi:"dnsServers"`
	// Indicates if accelerated networking is set on the specified Network Interface.
	EnableAcceleratedNetworking bool `pulumi:"enableAcceleratedNetworking"`
	// Indicate if IP forwarding is set on the specified Network Interface.
	EnableIpForwarding bool `pulumi:"enableIpForwarding"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The internal dns name label of the specified Network Interface.
	InternalDnsNameLabel string `pulumi:"internalDnsNameLabel"`
	// One or more `ipConfiguration` blocks as defined below.
	IpConfigurations []GetNetworkInterfaceIpConfiguration `pulumi:"ipConfigurations"`
	// The location of the specified Network Interface.
	Location string `pulumi:"location"`
	// The MAC address used by the specified Network Interface.
	MacAddress string `pulumi:"macAddress"`
	// The name of the IP Configuration.
	Name string `pulumi:"name"`
	// The ID of the network security group associated to the specified Network Interface.
	NetworkSecurityGroupId string `pulumi:"networkSecurityGroupId"`
	// The Private IP Address assigned to this Network Interface.
	PrivateIpAddress string `pulumi:"privateIpAddress"`
	// The list of private ip addresses associates to the specified Network Interface.
	PrivateIpAddresses []string `pulumi:"privateIpAddresses"`
	ResourceGroupName  string   `pulumi:"resourceGroupName"`
	// List the tags associated to the specified Network Interface.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the virtual machine that the specified Network Interface is attached to.
	VirtualMachineId string `pulumi:"virtualMachineId"`
}

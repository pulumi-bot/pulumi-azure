// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Traffic Manager Profile.
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
// 		_, err := network.LookupTrafficManagerProfile(ctx, &network.LookupTrafficManagerProfileArgs{
// 			Name:              "test",
// 			ResourceGroupName: "test",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("trafficRoutingMethod", data.Azurerm_traffic_manager_profile.Traffic_routing_method)
// 		return nil
// 	})
// }
// ```
func LookupTrafficManagerProfile(ctx *pulumi.Context, args *LookupTrafficManagerProfileArgs, opts ...pulumi.InvokeOption) (*LookupTrafficManagerProfileResult, error) {
	var rv LookupTrafficManagerProfileResult
	err := ctx.Invoke("azure:network/getTrafficManagerProfile:getTrafficManagerProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTrafficManagerProfile.
type LookupTrafficManagerProfileArgs struct {
	// Specifies the name of the Traffic Manager Profile.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group the Traffic Manager Profile is located in.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Indicates whether Traffic View is enabled for the Traffic Manager profile.
	TrafficViewEnabled *bool `pulumi:"trafficViewEnabled"`
}

// A collection of values returned by getTrafficManagerProfile.
type LookupTrafficManagerProfileResult struct {
	// This block specifies the DNS configuration of the Profile.
	DnsConfigs []GetTrafficManagerProfileDnsConfig `pulumi:"dnsConfigs"`
	// The FQDN of the created Profile.
	Fqdn string `pulumi:"fqdn"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// This block specifies the Endpoint monitoring configuration for the Profile.
	MonitorConfigs []GetTrafficManagerProfileMonitorConfig `pulumi:"monitorConfigs"`
	// The name of the custom header.
	Name string `pulumi:"name"`
	// The status of the profile.
	ProfileStatus     string `pulumi:"profileStatus"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the algorithm used to route traffic.
	TrafficRoutingMethod string `pulumi:"trafficRoutingMethod"`
	// Indicates whether Traffic View is enabled for the Traffic Manager profile.
	TrafficViewEnabled *bool `pulumi:"trafficViewEnabled"`
}

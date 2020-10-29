// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Public IP Prefix.
//
// ## Example Usage
func LookupPublicIpPrefix(ctx *pulumi.Context, args *LookupPublicIpPrefixArgs, opts ...pulumi.InvokeOption) (*LookupPublicIpPrefixResult, error) {
	var rv LookupPublicIpPrefixResult
	err := ctx.Invoke("azure:network/getPublicIpPrefix:getPublicIpPrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPublicIpPrefix.
type LookupPublicIpPrefixArgs struct {
	// Specifies the name of the public IP prefix.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group.
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	Zones             []string `pulumi:"zones"`
}

// A collection of values returned by getPublicIpPrefix.
type LookupPublicIpPrefixResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	IpPrefix string `pulumi:"ipPrefix"`
	// The supported Azure location where the resource exists.
	Location string `pulumi:"location"`
	// The name of the Public IP prefix resource.
	Name string `pulumi:"name"`
	// The number of bits of the prefix.
	PrefixLength int `pulumi:"prefixLength"`
	// The name of the resource group in which to create the public IP.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the Public IP Prefix.
	Sku string `pulumi:"sku"`
	// A mapping of tags to assigned to the resource.
	Tags  map[string]string `pulumi:"tags"`
	Zones []string          `pulumi:"zones"`
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing DNS Zone.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/dns"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "search-service"
// 		example, err := dns.LookupZone(ctx, "azure:dns:getZone", &dns.LookupZoneArgs{
// 			Name:              "search-eventhubns",
// 			ResourceGroupName: &opt0,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("dnsZoneId", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupZone(ctx *pulumi.Context, args *LookupZoneArgs, opts ...pulumi.InvokeOption) (*LookupZoneResult, error) {
	var rv LookupZoneResult
	err := ctx.Invoke("azure:dns/getZone:getZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZone.
type LookupZoneArgs struct {
	// The name of the DNS Zone.
	Name string `pulumi:"name"`
	// The Name of the Resource Group where the DNS Zone exists.
	// If the Name of the Resource Group is not provided, the first DNS Zone from the list of DNS Zones
	// in your subscription that matches `name` will be returned.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getZone.
type LookupZoneResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Maximum number of Records in the zone.
	MaxNumberOfRecordSets int    `pulumi:"maxNumberOfRecordSets"`
	Name                  string `pulumi:"name"`
	// A list of values that make up the NS record for the zone.
	NameServers []string `pulumi:"nameServers"`
	// The number of records already in the zone.
	NumberOfRecordSets int    `pulumi:"numberOfRecordSets"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the EventHub Namespace.
	Tags map[string]string `pulumi:"tags"`
}

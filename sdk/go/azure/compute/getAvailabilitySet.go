// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Availability Set.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := compute.LookupAvailabilitySet(ctx, &compute.LookupAvailabilitySetArgs{
// 			Name:              "tf-appsecuritygroup",
// 			ResourceGroupName: "my-resource-group",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("availabilitySetId", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupAvailabilitySet(ctx *pulumi.Context, args *LookupAvailabilitySetArgs, opts ...pulumi.InvokeOption) (*LookupAvailabilitySetResult, error) {
	var rv LookupAvailabilitySetResult
	err := ctx.Invoke("azure:compute/getAvailabilitySet:getAvailabilitySet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAvailabilitySet.
type LookupAvailabilitySetArgs struct {
	// The name of the Availability Set.
	Name string `pulumi:"name"`
	// The name of the resource group in which the Availability Set exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getAvailabilitySet.
type LookupAvailabilitySetResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The supported Azure location where the Availability Set exists.
	Location string `pulumi:"location"`
	// Whether the availability set is managed or not.
	Managed bool   `pulumi:"managed"`
	Name    string `pulumi:"name"`
	// The number of fault domains that are used.
	PlatformFaultDomainCount string `pulumi:"platformFaultDomainCount"`
	// The number of update domains that are used.
	PlatformUpdateDomainCount string `pulumi:"platformUpdateDomainCount"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerservice

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Container Registry.
//
// ## Example Usage
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
// 		example, err := containerservice.LookupRegistry(ctx, &containerservice.LookupRegistryArgs{
// 			Name:              "testacr",
// 			ResourceGroupName: "test",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("loginServer", example.LoginServer)
// 		return nil
// 	})
// }
// ```
func LookupRegistry(ctx *pulumi.Context, args *LookupRegistryArgs, opts ...pulumi.InvokeOption) (*LookupRegistryResult, error) {
	var rv LookupRegistryResult
	err := ctx.Invoke("azure:containerservice/getRegistry:getRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegistry.
type LookupRegistryArgs struct {
	// The name of the Container Registry.
	Name string `pulumi:"name"`
	// The Name of the Resource Group where this Container Registry exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getRegistry.
type LookupRegistryResult struct {
	// Is the Administrator account enabled for this Container Registry.
	AdminEnabled bool `pulumi:"adminEnabled"`
	// The Password associated with the Container Registry Admin account - if the admin account is enabled.
	AdminPassword string `pulumi:"adminPassword"`
	// The Username associated with the Container Registry Admin account - if the admin account is enabled.
	AdminUsername string `pulumi:"adminUsername"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Azure Region in which this Container Registry exists.
	Location string `pulumi:"location"`
	// The URL that can be used to log into the container registry.
	LoginServer       string `pulumi:"loginServer"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of this Container Registry, such as `Basic`.
	Sku string `pulumi:"sku"`
	// The ID of the Storage Account used for this Container Registry. This is only returned for `Classic` SKU's.
	StorageAccountId string `pulumi:"storageAccountId"`
	// A map of tags assigned to the Container Registry.
	Tags map[string]string `pulumi:"tags"`
}

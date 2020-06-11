// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about a Platform Image.
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
// 		example, err := compute.LookupPlatformImage(ctx, &compute.LookupPlatformImageArgs{
// 			Location:  "West Europe",
// 			Publisher: "Canonical",
// 			Offer:     "UbuntuServer",
// 			Sku:       "16.04-LTS",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", example.Id)
// 		return nil
// 	})
// }
// ```
func GetPlatformImage(ctx *pulumi.Context, args *GetPlatformImageArgs, opts ...pulumi.InvokeOption) (*GetPlatformImageResult, error) {
	var rv GetPlatformImageResult
	err := ctx.Invoke("azure:compute/getPlatformImage:getPlatformImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPlatformImage.
type GetPlatformImageArgs struct {
	// Specifies the Location to pull information about this Platform Image from.
	Location string `pulumi:"location"`
	// Specifies the Offer associated with the Platform Image.
	Offer string `pulumi:"offer"`
	// Specifies the Publisher associated with the Platform Image.
	Publisher string `pulumi:"publisher"`
	// Specifies the SKU of the Platform Image.
	Sku string `pulumi:"sku"`
	// The version of the Platform Image.
	Version *string `pulumi:"version"`
}

// A collection of values returned by getPlatformImage.
type GetPlatformImageResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	Location  string `pulumi:"location"`
	Offer     string `pulumi:"offer"`
	Publisher string `pulumi:"publisher"`
	Sku       string `pulumi:"sku"`
	Version   string `pulumi:"version"`
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Shared Image Gallery.
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
// 		_, err := compute.LookupSharedImageGallery(ctx, "azure:compute:getSharedImageGallery", &compute.LookupSharedImageGalleryArgs{
// 			Name:              "my-image-gallery",
// 			ResourceGroupName: "example-resources",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupSharedImageGallery(ctx *pulumi.Context, args *LookupSharedImageGalleryArgs, opts ...pulumi.InvokeOption) (*LookupSharedImageGalleryResult, error) {
	var rv LookupSharedImageGalleryResult
	err := ctx.Invoke("azure:compute/getSharedImageGallery:getSharedImageGallery", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSharedImageGallery.
type LookupSharedImageGalleryArgs struct {
	// The name of the Shared Image Gallery.
	Name string `pulumi:"name"`
	// The name of the Resource Group in which the Shared Image Gallery exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getSharedImageGallery.
type LookupSharedImageGalleryResult struct {
	// A description for the Shared Image Gallery.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id                string `pulumi:"id"`
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags which are assigned to the Shared Image Gallery.
	Tags map[string]string `pulumi:"tags"`
	// The unique name assigned to the Shared Image Gallery.
	UniqueName string `pulumi:"uniqueName"`
}

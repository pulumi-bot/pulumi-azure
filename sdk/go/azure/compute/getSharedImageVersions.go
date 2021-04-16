// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about existing Versions of a Shared Image within a Shared Image Gallery.
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
// 		_, err := compute.GetSharedImageVersions(ctx, &compute.GetSharedImageVersionsArgs{
// 			GalleryName:       "my-image-gallery",
// 			ImageName:         "my-image",
// 			ResourceGroupName: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetSharedImageVersions(ctx *pulumi.Context, args *GetSharedImageVersionsArgs, opts ...pulumi.InvokeOption) (*GetSharedImageVersionsResult, error) {
	var rv GetSharedImageVersionsResult
	err := ctx.Invoke("azure:compute/getSharedImageVersions:getSharedImageVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSharedImageVersions.
type GetSharedImageVersionsArgs struct {
	// The name of the Shared Image in which the Shared Image exists.
	GalleryName string `pulumi:"galleryName"`
	// The name of the Shared Image in which this Version exists.
	ImageName string `pulumi:"imageName"`
	// The name of the Resource Group in which the Shared Image Gallery exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to filter the list of images against.
	TagsFilter map[string]string `pulumi:"tagsFilter"`
}

// A collection of values returned by getSharedImageVersions.
type GetSharedImageVersionsResult struct {
	GalleryName string `pulumi:"galleryName"`
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	ImageName string `pulumi:"imageName"`
	// An `images` block as defined below:
	Images            []GetSharedImageVersionsImage `pulumi:"images"`
	ResourceGroupName string                        `pulumi:"resourceGroupName"`
	TagsFilter        map[string]string             `pulumi:"tagsFilter"`
}

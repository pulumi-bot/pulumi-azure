// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about existing Images within a Resource Group.
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
// 		_, err := compute.GetImages(ctx, &compute.GetImagesArgs{
// 			ResourceGroupName: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetImages(ctx *pulumi.Context, args *GetImagesArgs, opts ...pulumi.InvokeOption) (*GetImagesResult, error) {
	var rv GetImagesResult
	err := ctx.Invoke("azure:compute/getImages:getImages", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImages.
type GetImagesArgs struct {
	// The name of the Resource Group in which the Image exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to filter the list of images against.
	TagsFilter map[string]string `pulumi:"tagsFilter"`
}

// A collection of values returned by getImages.
type GetImagesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// One or more `images` blocks as defined below:
	Images            []GetImagesImage  `pulumi:"images"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	TagsFilter        map[string]string `pulumi:"tagsFilter"`
}

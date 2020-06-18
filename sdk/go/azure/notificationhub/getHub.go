// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package notificationhub

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Notification Hub within a Notification Hub Namespace.
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
// 		example, err := notificationhub.LookupHub(ctx, &notificationhub.LookupHubArgs{
// 			Name:              "notification-hub",
// 			NamespaceName:     "namespace-name",
// 			ResourceGroupName: "resource-group-name",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupHub(ctx *pulumi.Context, args *LookupHubArgs, opts ...pulumi.InvokeOption) (*LookupHubResult, error) {
	var rv LookupHubResult
	err := ctx.Invoke("azure:notificationhub/getHub:getHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHub.
type LookupHubArgs struct {
	// Specifies the Name of the Notification Hub.
	Name string `pulumi:"name"`
	// Specifies the Name of the Notification Hub Namespace which contains the Notification Hub.
	NamespaceName string `pulumi:"namespaceName"`
	// Specifies the Name of the Resource Group within which the Notification Hub exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getHub.
type LookupHubResult struct {
	// A `apnsCredential` block as defined below.
	ApnsCredentials []GetHubApnsCredential `pulumi:"apnsCredentials"`
	// A `gcmCredential` block as defined below.
	GcmCredentials []GetHubGcmCredential `pulumi:"gcmCredentials"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Azure Region in which this Notification Hub exists.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

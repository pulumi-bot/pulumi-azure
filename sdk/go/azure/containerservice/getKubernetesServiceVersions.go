// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerservice

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to retrieve the version of Kubernetes supported by Azure Kubernetes Service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/containerservice"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := containerservice.GetKubernetesServiceVersions(ctx, "azure:containerservice:getKubernetesServiceVersions", &containerservice.GetKubernetesServiceVersionsArgs{
// 			Location: "West Europe",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("versions", current.Versions)
// 		ctx.Export("latestVersion", current.LatestVersion)
// 		return nil
// 	})
// }
// ```
func GetKubernetesServiceVersions(ctx *pulumi.Context, args *GetKubernetesServiceVersionsArgs, opts ...pulumi.InvokeOption) (*GetKubernetesServiceVersionsResult, error) {
	var rv GetKubernetesServiceVersionsResult
	err := ctx.Invoke("azure:containerservice/getKubernetesServiceVersions:getKubernetesServiceVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKubernetesServiceVersions.
type GetKubernetesServiceVersionsArgs struct {
	// Should Preview versions of Kubernetes in AKS be included? Defaults to `true`
	IncludePreview *bool `pulumi:"includePreview"`
	// Specifies the location in which to query for versions.
	Location string `pulumi:"location"`
	// A prefix filter for the versions of Kubernetes which should be returned; for example `1.` will return `1.9` to `1.14`, whereas `1.12` will return `1.12.2`.
	VersionPrefix *string `pulumi:"versionPrefix"`
}

// A collection of values returned by getKubernetesServiceVersions.
type GetKubernetesServiceVersionsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id             string `pulumi:"id"`
	IncludePreview *bool  `pulumi:"includePreview"`
	// The most recent version available. If `includePreview == false`, this is the most recent non-preview version available.
	LatestVersion string  `pulumi:"latestVersion"`
	Location      string  `pulumi:"location"`
	VersionPrefix *string `pulumi:"versionPrefix"`
	// The list of all supported versions.
	Versions []string `pulumi:"versions"`
}

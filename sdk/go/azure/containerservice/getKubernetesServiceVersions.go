// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve the version of Kubernetes supported by Azure Kubernetes Service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/containerservice"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := containerservice.GetKubernetesServiceVersions(ctx, &containerservice.GetKubernetesServiceVersionsArgs{
// 			Location: "West Europe",
// 		}, nil)
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

func GetKubernetesServiceVersionsOutput(ctx *pulumi.Context, args GetKubernetesServiceVersionsOutputArgs, opts ...pulumi.InvokeOption) GetKubernetesServiceVersionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetKubernetesServiceVersionsResult, error) {
			args := v.(GetKubernetesServiceVersionsArgs)
			r, err := GetKubernetesServiceVersions(ctx, &args, opts...)
			return *r, err
		}).(GetKubernetesServiceVersionsResultOutput)
}

// A collection of arguments for invoking getKubernetesServiceVersions.
type GetKubernetesServiceVersionsOutputArgs struct {
	// Should Preview versions of Kubernetes in AKS be included? Defaults to `true`
	IncludePreview pulumi.BoolPtrInput `pulumi:"includePreview"`
	// Specifies the location in which to query for versions.
	Location pulumi.StringInput `pulumi:"location"`
	// A prefix filter for the versions of Kubernetes which should be returned; for example `1.` will return `1.9` to `1.14`, whereas `1.12` will return `1.12.2`.
	VersionPrefix pulumi.StringPtrInput `pulumi:"versionPrefix"`
}

func (GetKubernetesServiceVersionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubernetesServiceVersionsArgs)(nil)).Elem()
}

// A collection of values returned by getKubernetesServiceVersions.
type GetKubernetesServiceVersionsResultOutput struct{ *pulumi.OutputState }

func (GetKubernetesServiceVersionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubernetesServiceVersionsResult)(nil)).Elem()
}

func (o GetKubernetesServiceVersionsResultOutput) ToGetKubernetesServiceVersionsResultOutput() GetKubernetesServiceVersionsResultOutput {
	return o
}

func (o GetKubernetesServiceVersionsResultOutput) ToGetKubernetesServiceVersionsResultOutputWithContext(ctx context.Context) GetKubernetesServiceVersionsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetKubernetesServiceVersionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubernetesServiceVersionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetKubernetesServiceVersionsResultOutput) IncludePreview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetKubernetesServiceVersionsResult) *bool { return v.IncludePreview }).(pulumi.BoolPtrOutput)
}

// The most recent version available. If `includePreview == false`, this is the most recent non-preview version available.
func (o GetKubernetesServiceVersionsResultOutput) LatestVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubernetesServiceVersionsResult) string { return v.LatestVersion }).(pulumi.StringOutput)
}

func (o GetKubernetesServiceVersionsResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubernetesServiceVersionsResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetKubernetesServiceVersionsResultOutput) VersionPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKubernetesServiceVersionsResult) *string { return v.VersionPrefix }).(pulumi.StringPtrOutput)
}

// The list of all supported versions.
func (o GetKubernetesServiceVersionsResultOutput) Versions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetKubernetesServiceVersionsResult) []string { return v.Versions }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKubernetesServiceVersionsResultOutput{})
}

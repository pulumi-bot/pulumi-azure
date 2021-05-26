// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Healthcare Service
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/healthcare"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := healthcare.LookupService(ctx, &healthcare.LookupServiceArgs{
// 			Name:              "example-healthcare_service",
// 			ResourceGroupName: "example-resources",
// 			Location:          "westus2",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("healthcareServiceId", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azure:healthcare/getService:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getService.
type LookupServiceArgs struct {
	// The Azure Region where the Service is located.
	Location string `pulumi:"location"`
	// Specifies the name of the Healthcare Service.
	Name string `pulumi:"name"`
	// The name of the Resource Group in which the Healthcare Service exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getService.
type LookupServiceResult struct {
	AccessPolicyObjectIds []string `pulumi:"accessPolicyObjectIds"`
	// An `authenticationConfiguration` block as defined below.
	AuthenticationConfigurations []GetServiceAuthenticationConfiguration `pulumi:"authenticationConfigurations"`
	// A `corsConfiguration` block as defined below.
	CorsConfigurations []GetServiceCorsConfiguration `pulumi:"corsConfigurations"`
	// The versionless Key Vault Key ID for CMK encryption of the backing database.
	CosmosdbKeyVaultKeyVersionlessId string `pulumi:"cosmosdbKeyVaultKeyVersionlessId"`
	// The provisioned throughput for the backing database.
	CosmosdbThroughput int `pulumi:"cosmosdbThroughput"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The type of the service.
	Kind string `pulumi:"kind"`
	// The Azure Region where the Service is located.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

func LookupServiceApply(ctx *pulumi.Context, args LookupServiceApplyInput, opts ...pulumi.InvokeOption) LookupServiceResultOutput {
	return args.ToLookupServiceApplyOutput().ApplyT(func(v LookupServiceArgs) (LookupServiceResult, error) {
		r, err := LookupService(ctx, &v, opts...)
		return *r, err

	}).(LookupServiceResultOutput)
}

// LookupServiceApplyInput is an input type that accepts LookupServiceApplyArgs and LookupServiceApplyOutput values.
// You can construct a concrete instance of `LookupServiceApplyInput` via:
//
//          LookupServiceApplyArgs{...}
type LookupServiceApplyInput interface {
	pulumi.Input

	ToLookupServiceApplyOutput() LookupServiceApplyOutput
	ToLookupServiceApplyOutputWithContext(context.Context) LookupServiceApplyOutput
}

// A collection of arguments for invoking getService.
type LookupServiceApplyArgs struct {
	// The Azure Region where the Service is located.
	Location pulumi.StringInput `pulumi:"location"`
	// Specifies the name of the Healthcare Service.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the Resource Group in which the Healthcare Service exists.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupServiceApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceArgs)(nil)).Elem()
}

func (i LookupServiceApplyArgs) ToLookupServiceApplyOutput() LookupServiceApplyOutput {
	return i.ToLookupServiceApplyOutputWithContext(context.Background())
}

func (i LookupServiceApplyArgs) ToLookupServiceApplyOutputWithContext(ctx context.Context) LookupServiceApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupServiceApplyOutput)
}

// A collection of arguments for invoking getService.
type LookupServiceApplyOutput struct{ *pulumi.OutputState }

func (LookupServiceApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceArgs)(nil)).Elem()
}

func (o LookupServiceApplyOutput) ToLookupServiceApplyOutput() LookupServiceApplyOutput {
	return o
}

func (o LookupServiceApplyOutput) ToLookupServiceApplyOutputWithContext(ctx context.Context) LookupServiceApplyOutput {
	return o
}

// The Azure Region where the Service is located.
func (o LookupServiceApplyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceArgs) string { return v.Location }).(pulumi.StringOutput)
}

// Specifies the name of the Healthcare Service.
func (o LookupServiceApplyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceArgs) string { return v.Name }).(pulumi.StringOutput)
}

// The name of the Resource Group in which the Healthcare Service exists.
func (o LookupServiceApplyOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceArgs) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// A collection of values returned by getService.
type LookupServiceResultOutput struct{ *pulumi.OutputState }

func (LookupServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceResult)(nil)).Elem()
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutput() LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutputWithContext(ctx context.Context) LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) AccessPolicyObjectIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServiceResult) []string { return v.AccessPolicyObjectIds }).(pulumi.StringArrayOutput)
}

// An `authenticationConfiguration` block as defined below.
func (o LookupServiceResultOutput) AuthenticationConfigurations() GetServiceAuthenticationConfigurationArrayOutput {
	return o.ApplyT(func(v LookupServiceResult) []GetServiceAuthenticationConfiguration {
		return v.AuthenticationConfigurations
	}).(GetServiceAuthenticationConfigurationArrayOutput)
}

// A `corsConfiguration` block as defined below.
func (o LookupServiceResultOutput) CorsConfigurations() GetServiceCorsConfigurationArrayOutput {
	return o.ApplyT(func(v LookupServiceResult) []GetServiceCorsConfiguration { return v.CorsConfigurations }).(GetServiceCorsConfigurationArrayOutput)
}

// The versionless Key Vault Key ID for CMK encryption of the backing database.
func (o LookupServiceResultOutput) CosmosdbKeyVaultKeyVersionlessId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.CosmosdbKeyVaultKeyVersionlessId }).(pulumi.StringOutput)
}

// The provisioned throughput for the backing database.
func (o LookupServiceResultOutput) CosmosdbThroughput() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServiceResult) int { return v.CosmosdbThroughput }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The type of the service.
func (o LookupServiceResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The Azure Region where the Service is located.
func (o LookupServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o LookupServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceApplyOutput{})
	pulumi.RegisterOutputType(LookupServiceResultOutput{})
}

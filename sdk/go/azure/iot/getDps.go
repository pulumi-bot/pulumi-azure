// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing IotHub Device Provisioning Service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/iot"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iot.GetDps(ctx, &iot.GetDpsArgs{
// 			Name:              "iot_hub_dps_test",
// 			ResourceGroupName: "iothub_dps_rg",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDps(ctx *pulumi.Context, args *GetDpsArgs, opts ...pulumi.InvokeOption) (*GetDpsResult, error) {
	var rv GetDpsResult
	err := ctx.Invoke("azure:iot/getDps:getDps", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDps.
type GetDpsArgs struct {
	// Specifies the name of the Iot Device Provisioning Service resource.
	Name string `pulumi:"name"`
	// The name of the resource group under which the Iot Device Provisioning Service is located in.
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}

// A collection of values returned by getDps.
type GetDpsResult struct {
	// The allocation policy of the IoT Device Provisioning Service.
	AllocationPolicy string `pulumi:"allocationPolicy"`
	// The device endpoint of the IoT Device Provisioning Service.
	DeviceProvisioningHostName string `pulumi:"deviceProvisioningHostName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The unique identifier of the IoT Device Provisioning Service.
	IdScope string `pulumi:"idScope"`
	// Specifies the supported Azure location where the IoT Device Provisioning Service exists.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The service endpoint of the IoT Device Provisioning Service.
	ServiceOperationsHostName string            `pulumi:"serviceOperationsHostName"`
	Tags                      map[string]string `pulumi:"tags"`
}

func GetDpsApply(ctx *pulumi.Context, args GetDpsApplyInput, opts ...pulumi.InvokeOption) GetDpsResultOutput {
	return args.ToGetDpsApplyOutput().ApplyT(func(v GetDpsArgs) (GetDpsResult, error) {
		r, err := GetDps(ctx, &v, opts...)
		return *r, err

	}).(GetDpsResultOutput)
}

// GetDpsApplyInput is an input type that accepts GetDpsApplyArgs and GetDpsApplyOutput values.
// You can construct a concrete instance of `GetDpsApplyInput` via:
//
//          GetDpsApplyArgs{...}
type GetDpsApplyInput interface {
	pulumi.Input

	ToGetDpsApplyOutput() GetDpsApplyOutput
	ToGetDpsApplyOutputWithContext(context.Context) GetDpsApplyOutput
}

// A collection of arguments for invoking getDps.
type GetDpsApplyArgs struct {
	// Specifies the name of the Iot Device Provisioning Service resource.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group under which the Iot Device Provisioning Service is located in.
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	Tags              pulumi.StringMapInput `pulumi:"tags"`
}

func (GetDpsApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDpsArgs)(nil)).Elem()
}

func (i GetDpsApplyArgs) ToGetDpsApplyOutput() GetDpsApplyOutput {
	return i.ToGetDpsApplyOutputWithContext(context.Background())
}

func (i GetDpsApplyArgs) ToGetDpsApplyOutputWithContext(ctx context.Context) GetDpsApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDpsApplyOutput)
}

// A collection of arguments for invoking getDps.
type GetDpsApplyOutput struct{ *pulumi.OutputState }

func (GetDpsApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDpsArgs)(nil)).Elem()
}

func (o GetDpsApplyOutput) ToGetDpsApplyOutput() GetDpsApplyOutput {
	return o
}

func (o GetDpsApplyOutput) ToGetDpsApplyOutputWithContext(ctx context.Context) GetDpsApplyOutput {
	return o
}

// Specifies the name of the Iot Device Provisioning Service resource.
func (o GetDpsApplyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDpsArgs) string { return v.Name }).(pulumi.StringOutput)
}

// The name of the resource group under which the Iot Device Provisioning Service is located in.
func (o GetDpsApplyOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDpsArgs) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

func (o GetDpsApplyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetDpsArgs) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// A collection of values returned by getDps.
type GetDpsResultOutput struct{ *pulumi.OutputState }

func (GetDpsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDpsResult)(nil)).Elem()
}

func (o GetDpsResultOutput) ToGetDpsResultOutput() GetDpsResultOutput {
	return o
}

func (o GetDpsResultOutput) ToGetDpsResultOutputWithContext(ctx context.Context) GetDpsResultOutput {
	return o
}

// The allocation policy of the IoT Device Provisioning Service.
func (o GetDpsResultOutput) AllocationPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v GetDpsResult) string { return v.AllocationPolicy }).(pulumi.StringOutput)
}

// The device endpoint of the IoT Device Provisioning Service.
func (o GetDpsResultOutput) DeviceProvisioningHostName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDpsResult) string { return v.DeviceProvisioningHostName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDpsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDpsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The unique identifier of the IoT Device Provisioning Service.
func (o GetDpsResultOutput) IdScope() pulumi.StringOutput {
	return o.ApplyT(func(v GetDpsResult) string { return v.IdScope }).(pulumi.StringOutput)
}

// Specifies the supported Azure location where the IoT Device Provisioning Service exists.
func (o GetDpsResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetDpsResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetDpsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDpsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetDpsResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDpsResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// The service endpoint of the IoT Device Provisioning Service.
func (o GetDpsResultOutput) ServiceOperationsHostName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDpsResult) string { return v.ServiceOperationsHostName }).(pulumi.StringOutput)
}

func (o GetDpsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetDpsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDpsApplyOutput{})
	pulumi.RegisterOutputType(GetDpsResultOutput{})
}

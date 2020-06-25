// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing IotHub Device Provisioning Service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/iot"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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

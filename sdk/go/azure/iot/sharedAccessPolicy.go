// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an IotHub Shared Access Policy
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/iot"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleIoTHub, err := iot.NewIoTHub(ctx, "exampleIoTHub", &iot.IoTHubArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			Sku: &iot.IoTHubSkuArgs{
// 				Name:     pulumi.String("S1"),
// 				Capacity: pulumi.Int(1),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iot.NewSharedAccessPolicy(ctx, "exampleSharedAccessPolicy", &iot.SharedAccessPolicyArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			IothubName:        exampleIoTHub.Name,
// 			RegistryRead:      pulumi.Bool(true),
// 			RegistryWrite:     pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SharedAccessPolicy struct {
	pulumi.CustomResourceState

	// Adds `DeviceConnect` permission to this Shared Access Account. It allows sending and receiving on the device-side endpoints.
	DeviceConnect pulumi.BoolPtrOutput `pulumi:"deviceConnect"`
	// The name of the IoTHub to which this Shared Access Policy belongs. Changing this forces a new resource to be created.
	IothubName pulumi.StringOutput `pulumi:"iothubName"`
	// Specifies the name of the IotHub Shared Access Policy resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The primary connection string of the Shared Access Policy.
	PrimaryConnectionString pulumi.StringOutput `pulumi:"primaryConnectionString"`
	// The primary key used to create the authentication token.
	PrimaryKey pulumi.StringOutput `pulumi:"primaryKey"`
	// Adds `RegistryRead` permission to this Shared Access Account. It allows read access to the identity registry.
	RegistryRead pulumi.BoolPtrOutput `pulumi:"registryRead"`
	// Adds `RegistryWrite` permission to this Shared Access Account. It allows write access to the identity registry.
	RegistryWrite pulumi.BoolPtrOutput `pulumi:"registryWrite"`
	// The name of the resource group under which the IotHub Shared Access Policy resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The secondary connection string of the Shared Access Policy.
	SecondaryConnectionString pulumi.StringOutput `pulumi:"secondaryConnectionString"`
	// The secondary key used to create the authentication token.
	SecondaryKey pulumi.StringOutput `pulumi:"secondaryKey"`
	// Adds `ServiceConnect` permission to this Shared Access Account. It allows sending and receiving on the cloud-side endpoints.
	ServiceConnect pulumi.BoolPtrOutput `pulumi:"serviceConnect"`
}

// NewSharedAccessPolicy registers a new resource with the given unique name, arguments, and options.
func NewSharedAccessPolicy(ctx *pulumi.Context,
	name string, args *SharedAccessPolicyArgs, opts ...pulumi.ResourceOption) (*SharedAccessPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.IothubName == nil {
		return nil, errors.New("invalid value for required argument 'IothubName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource SharedAccessPolicy
	err := ctx.RegisterResource("azure:iot/sharedAccessPolicy:SharedAccessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSharedAccessPolicy gets an existing SharedAccessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSharedAccessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SharedAccessPolicyState, opts ...pulumi.ResourceOption) (*SharedAccessPolicy, error) {
	var resource SharedAccessPolicy
	err := ctx.ReadResource("azure:iot/sharedAccessPolicy:SharedAccessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SharedAccessPolicy resources.
type sharedAccessPolicyState struct {
	// Adds `DeviceConnect` permission to this Shared Access Account. It allows sending and receiving on the device-side endpoints.
	DeviceConnect *bool `pulumi:"deviceConnect"`
	// The name of the IoTHub to which this Shared Access Policy belongs. Changing this forces a new resource to be created.
	IothubName *string `pulumi:"iothubName"`
	// Specifies the name of the IotHub Shared Access Policy resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The primary connection string of the Shared Access Policy.
	PrimaryConnectionString *string `pulumi:"primaryConnectionString"`
	// The primary key used to create the authentication token.
	PrimaryKey *string `pulumi:"primaryKey"`
	// Adds `RegistryRead` permission to this Shared Access Account. It allows read access to the identity registry.
	RegistryRead *bool `pulumi:"registryRead"`
	// Adds `RegistryWrite` permission to this Shared Access Account. It allows write access to the identity registry.
	RegistryWrite *bool `pulumi:"registryWrite"`
	// The name of the resource group under which the IotHub Shared Access Policy resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The secondary connection string of the Shared Access Policy.
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	// The secondary key used to create the authentication token.
	SecondaryKey *string `pulumi:"secondaryKey"`
	// Adds `ServiceConnect` permission to this Shared Access Account. It allows sending and receiving on the cloud-side endpoints.
	ServiceConnect *bool `pulumi:"serviceConnect"`
}

type SharedAccessPolicyState struct {
	// Adds `DeviceConnect` permission to this Shared Access Account. It allows sending and receiving on the device-side endpoints.
	DeviceConnect pulumi.BoolPtrInput
	// The name of the IoTHub to which this Shared Access Policy belongs. Changing this forces a new resource to be created.
	IothubName pulumi.StringPtrInput
	// Specifies the name of the IotHub Shared Access Policy resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The primary connection string of the Shared Access Policy.
	PrimaryConnectionString pulumi.StringPtrInput
	// The primary key used to create the authentication token.
	PrimaryKey pulumi.StringPtrInput
	// Adds `RegistryRead` permission to this Shared Access Account. It allows read access to the identity registry.
	RegistryRead pulumi.BoolPtrInput
	// Adds `RegistryWrite` permission to this Shared Access Account. It allows write access to the identity registry.
	RegistryWrite pulumi.BoolPtrInput
	// The name of the resource group under which the IotHub Shared Access Policy resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The secondary connection string of the Shared Access Policy.
	SecondaryConnectionString pulumi.StringPtrInput
	// The secondary key used to create the authentication token.
	SecondaryKey pulumi.StringPtrInput
	// Adds `ServiceConnect` permission to this Shared Access Account. It allows sending and receiving on the cloud-side endpoints.
	ServiceConnect pulumi.BoolPtrInput
}

func (SharedAccessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sharedAccessPolicyState)(nil)).Elem()
}

type sharedAccessPolicyArgs struct {
	// Adds `DeviceConnect` permission to this Shared Access Account. It allows sending and receiving on the device-side endpoints.
	DeviceConnect *bool `pulumi:"deviceConnect"`
	// The name of the IoTHub to which this Shared Access Policy belongs. Changing this forces a new resource to be created.
	IothubName string `pulumi:"iothubName"`
	// Specifies the name of the IotHub Shared Access Policy resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Adds `RegistryRead` permission to this Shared Access Account. It allows read access to the identity registry.
	RegistryRead *bool `pulumi:"registryRead"`
	// Adds `RegistryWrite` permission to this Shared Access Account. It allows write access to the identity registry.
	RegistryWrite *bool `pulumi:"registryWrite"`
	// The name of the resource group under which the IotHub Shared Access Policy resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Adds `ServiceConnect` permission to this Shared Access Account. It allows sending and receiving on the cloud-side endpoints.
	ServiceConnect *bool `pulumi:"serviceConnect"`
}

// The set of arguments for constructing a SharedAccessPolicy resource.
type SharedAccessPolicyArgs struct {
	// Adds `DeviceConnect` permission to this Shared Access Account. It allows sending and receiving on the device-side endpoints.
	DeviceConnect pulumi.BoolPtrInput
	// The name of the IoTHub to which this Shared Access Policy belongs. Changing this forces a new resource to be created.
	IothubName pulumi.StringInput
	// Specifies the name of the IotHub Shared Access Policy resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Adds `RegistryRead` permission to this Shared Access Account. It allows read access to the identity registry.
	RegistryRead pulumi.BoolPtrInput
	// Adds `RegistryWrite` permission to this Shared Access Account. It allows write access to the identity registry.
	RegistryWrite pulumi.BoolPtrInput
	// The name of the resource group under which the IotHub Shared Access Policy resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Adds `ServiceConnect` permission to this Shared Access Account. It allows sending and receiving on the cloud-side endpoints.
	ServiceConnect pulumi.BoolPtrInput
}

func (SharedAccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sharedAccessPolicyArgs)(nil)).Elem()
}

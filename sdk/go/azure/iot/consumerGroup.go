// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Consumer Group within an IotHub
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
// 			Tags: pulumi.StringMap{
// 				"purpose": pulumi.String("testing"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iot.NewConsumerGroup(ctx, "exampleConsumerGroup", &iot.ConsumerGroupArgs{
// 			IothubName:           exampleIoTHub.Name,
// 			EventhubEndpointName: pulumi.String("events"),
// 			ResourceGroupName:    exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// IoTHub Consumer Groups can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:iot/consumerGroup:ConsumerGroup group1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/IotHubs/hub1/eventHubEndpoints/events/ConsumerGroups/group1
// ```
type ConsumerGroup struct {
	pulumi.CustomResourceState

	// The name of the Event Hub-compatible endpoint in the IoT hub. Changing this forces a new resource to be created.
	EventhubEndpointName pulumi.StringOutput `pulumi:"eventhubEndpointName"`
	// The name of the IoT Hub. Changing this forces a new resource to be created.
	IothubName pulumi.StringOutput `pulumi:"iothubName"`
	// The name of this Consumer Group. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group that contains the IoT hub. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewConsumerGroup registers a new resource with the given unique name, arguments, and options.
func NewConsumerGroup(ctx *pulumi.Context,
	name string, args *ConsumerGroupArgs, opts ...pulumi.ResourceOption) (*ConsumerGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventhubEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EventhubEndpointName'")
	}
	if args.IothubName == nil {
		return nil, errors.New("invalid value for required argument 'IothubName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ConsumerGroup
	err := ctx.RegisterResource("azure:iot/consumerGroup:ConsumerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConsumerGroup gets an existing ConsumerGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConsumerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsumerGroupState, opts ...pulumi.ResourceOption) (*ConsumerGroup, error) {
	var resource ConsumerGroup
	err := ctx.ReadResource("azure:iot/consumerGroup:ConsumerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConsumerGroup resources.
type consumerGroupState struct {
	// The name of the Event Hub-compatible endpoint in the IoT hub. Changing this forces a new resource to be created.
	EventhubEndpointName *string `pulumi:"eventhubEndpointName"`
	// The name of the IoT Hub. Changing this forces a new resource to be created.
	IothubName *string `pulumi:"iothubName"`
	// The name of this Consumer Group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group that contains the IoT hub. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type ConsumerGroupState struct {
	// The name of the Event Hub-compatible endpoint in the IoT hub. Changing this forces a new resource to be created.
	EventhubEndpointName pulumi.StringPtrInput
	// The name of the IoT Hub. Changing this forces a new resource to be created.
	IothubName pulumi.StringPtrInput
	// The name of this Consumer Group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group that contains the IoT hub. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (ConsumerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerGroupState)(nil)).Elem()
}

type consumerGroupArgs struct {
	// The name of the Event Hub-compatible endpoint in the IoT hub. Changing this forces a new resource to be created.
	EventhubEndpointName string `pulumi:"eventhubEndpointName"`
	// The name of the IoT Hub. Changing this forces a new resource to be created.
	IothubName string `pulumi:"iothubName"`
	// The name of this Consumer Group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group that contains the IoT hub. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ConsumerGroup resource.
type ConsumerGroupArgs struct {
	// The name of the Event Hub-compatible endpoint in the IoT hub. Changing this forces a new resource to be created.
	EventhubEndpointName pulumi.StringInput
	// The name of the IoT Hub. Changing this forces a new resource to be created.
	IothubName pulumi.StringInput
	// The name of this Consumer Group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group that contains the IoT hub. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (ConsumerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerGroupArgs)(nil)).Elem()
}

type ConsumerGroupInput interface {
	pulumi.Input

	ToConsumerGroupOutput() ConsumerGroupOutput
	ToConsumerGroupOutputWithContext(ctx context.Context) ConsumerGroupOutput
}

func (*ConsumerGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsumerGroup)(nil))
}

func (i *ConsumerGroup) ToConsumerGroupOutput() ConsumerGroupOutput {
	return i.ToConsumerGroupOutputWithContext(context.Background())
}

func (i *ConsumerGroup) ToConsumerGroupOutputWithContext(ctx context.Context) ConsumerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerGroupOutput)
}

type ConsumerGroupOutput struct {
	*pulumi.OutputState
}

func (ConsumerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsumerGroup)(nil))
}

func (o ConsumerGroupOutput) ToConsumerGroupOutput() ConsumerGroupOutput {
	return o
}

func (o ConsumerGroupOutput) ToConsumerGroupOutputWithContext(ctx context.Context) ConsumerGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConsumerGroupOutput{})
}

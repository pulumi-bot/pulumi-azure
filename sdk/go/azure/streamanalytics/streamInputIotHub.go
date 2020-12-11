// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Stream Analytics Stream Input IoTHub.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/iot"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/streamanalytics"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := core.LookupResourceGroup(ctx, &core.LookupResourceGroupArgs{
// 			Name: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleJob, err := streamanalytics.LookupJob(ctx, &streamanalytics.LookupJobArgs{
// 			Name:              "example-job",
// 			ResourceGroupName: azurerm_resource_group.Example.Name,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleIoTHub, err := iot.NewIoTHub(ctx, "exampleIoTHub", &iot.IoTHubArgs{
// 			ResourceGroupName: pulumi.Any(azurerm_resource_group.Example.Name),
// 			Location:          pulumi.Any(azurerm_resource_group.Example.Location),
// 			Sku: &iot.IoTHubSkuArgs{
// 				Name:     pulumi.String("S1"),
// 				Capacity: pulumi.Int(1),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = streamanalytics.NewStreamInputIotHub(ctx, "exampleStreamInputIotHub", &streamanalytics.StreamInputIotHubArgs{
// 			StreamAnalyticsJobName:    pulumi.String(exampleJob.Name),
// 			ResourceGroupName:         pulumi.String(exampleJob.ResourceGroupName),
// 			Endpoint:                  pulumi.String("messages/events"),
// 			EventhubConsumerGroupName: pulumi.String(fmt.Sprintf("%v%v", "$", "Default")),
// 			IothubNamespace:           exampleIoTHub.Name,
// 			SharedAccessPolicyKey: pulumi.String(exampleIoTHub.SharedAccessPolicies.ApplyT(func(sharedAccessPolicies []iot.IoTHubSharedAccessPolicy) (string, error) {
// 				return sharedAccessPolicies[0].PrimaryKey, nil
// 			}).(pulumi.StringOutput)),
// 			SharedAccessPolicyName: pulumi.String("iothubowner"),
// 			Serialization: &streamanalytics.StreamInputIotHubSerializationArgs{
// 				Type:     pulumi.String("Json"),
// 				Encoding: pulumi.String("UTF8"),
// 			},
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
// Stream Analytics Stream Input IoTHub's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:streamanalytics/streamInputIotHub:StreamInputIotHub example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/inputs/input1
// ```
type StreamInputIotHub struct {
	pulumi.CustomResourceState

	// The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
	EventhubConsumerGroupName pulumi.StringOutput `pulumi:"eventhubConsumerGroupName"`
	// The name or the URI of the IoT Hub.
	IothubNamespace pulumi.StringOutput `pulumi:"iothubNamespace"`
	// The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization StreamInputIotHubSerializationOutput `pulumi:"serialization"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringOutput `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringOutput `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringOutput `pulumi:"streamAnalyticsJobName"`
}

// NewStreamInputIotHub registers a new resource with the given unique name, arguments, and options.
func NewStreamInputIotHub(ctx *pulumi.Context,
	name string, args *StreamInputIotHubArgs, opts ...pulumi.ResourceOption) (*StreamInputIotHub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Endpoint == nil {
		return nil, errors.New("invalid value for required argument 'Endpoint'")
	}
	if args.EventhubConsumerGroupName == nil {
		return nil, errors.New("invalid value for required argument 'EventhubConsumerGroupName'")
	}
	if args.IothubNamespace == nil {
		return nil, errors.New("invalid value for required argument 'IothubNamespace'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Serialization == nil {
		return nil, errors.New("invalid value for required argument 'Serialization'")
	}
	if args.SharedAccessPolicyKey == nil {
		return nil, errors.New("invalid value for required argument 'SharedAccessPolicyKey'")
	}
	if args.SharedAccessPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'SharedAccessPolicyName'")
	}
	if args.StreamAnalyticsJobName == nil {
		return nil, errors.New("invalid value for required argument 'StreamAnalyticsJobName'")
	}
	var resource StreamInputIotHub
	err := ctx.RegisterResource("azure:streamanalytics/streamInputIotHub:StreamInputIotHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStreamInputIotHub gets an existing StreamInputIotHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamInputIotHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamInputIotHubState, opts ...pulumi.ResourceOption) (*StreamInputIotHub, error) {
	var resource StreamInputIotHub
	err := ctx.ReadResource("azure:streamanalytics/streamInputIotHub:StreamInputIotHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StreamInputIotHub resources.
type streamInputIotHubState struct {
	// The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
	Endpoint *string `pulumi:"endpoint"`
	// The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
	EventhubConsumerGroupName *string `pulumi:"eventhubConsumerGroupName"`
	// The name or the URI of the IoT Hub.
	IothubNamespace *string `pulumi:"iothubNamespace"`
	// The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization *StreamInputIotHubSerialization `pulumi:"serialization"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey *string `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `pulumi:"streamAnalyticsJobName"`
}

type StreamInputIotHubState struct {
	// The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
	Endpoint pulumi.StringPtrInput
	// The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
	EventhubConsumerGroupName pulumi.StringPtrInput
	// The name or the URI of the IoT Hub.
	IothubNamespace pulumi.StringPtrInput
	// The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `serialization` block as defined below.
	Serialization StreamInputIotHubSerializationPtrInput
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringPtrInput
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringPtrInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringPtrInput
}

func (StreamInputIotHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamInputIotHubState)(nil)).Elem()
}

type streamInputIotHubArgs struct {
	// The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
	Endpoint string `pulumi:"endpoint"`
	// The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
	EventhubConsumerGroupName string `pulumi:"eventhubConsumerGroupName"`
	// The name or the URI of the IoT Hub.
	IothubNamespace string `pulumi:"iothubNamespace"`
	// The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization StreamInputIotHubSerialization `pulumi:"serialization"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey string `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName string `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName string `pulumi:"streamAnalyticsJobName"`
}

// The set of arguments for constructing a StreamInputIotHub resource.
type StreamInputIotHubArgs struct {
	// The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
	Endpoint pulumi.StringInput
	// The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
	EventhubConsumerGroupName pulumi.StringInput
	// The name or the URI of the IoT Hub.
	IothubNamespace pulumi.StringInput
	// The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `serialization` block as defined below.
	Serialization StreamInputIotHubSerializationInput
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringInput
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringInput
}

func (StreamInputIotHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamInputIotHubArgs)(nil)).Elem()
}

type StreamInputIotHubInput interface {
	pulumi.Input

	ToStreamInputIotHubOutput() StreamInputIotHubOutput
	ToStreamInputIotHubOutputWithContext(ctx context.Context) StreamInputIotHubOutput
}

type StreamInputIotHubPtrInput interface {
	pulumi.Input

	ToStreamInputIotHubPtrOutput() StreamInputIotHubPtrOutput
	ToStreamInputIotHubPtrOutputWithContext(ctx context.Context) StreamInputIotHubPtrOutput
}

func (StreamInputIotHub) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamInputIotHub)(nil)).Elem()
}

func (i StreamInputIotHub) ToStreamInputIotHubOutput() StreamInputIotHubOutput {
	return i.ToStreamInputIotHubOutputWithContext(context.Background())
}

func (i StreamInputIotHub) ToStreamInputIotHubOutputWithContext(ctx context.Context) StreamInputIotHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamInputIotHubOutput)
}

func (i StreamInputIotHub) ToStreamInputIotHubPtrOutput() StreamInputIotHubPtrOutput {
	return i.ToStreamInputIotHubPtrOutputWithContext(context.Background())
}

func (i StreamInputIotHub) ToStreamInputIotHubPtrOutputWithContext(ctx context.Context) StreamInputIotHubPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamInputIotHubPtrOutput)
}

type StreamInputIotHubOutput struct {
	*pulumi.OutputState
}

func (StreamInputIotHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamInputIotHubOutput)(nil)).Elem()
}

func (o StreamInputIotHubOutput) ToStreamInputIotHubOutput() StreamInputIotHubOutput {
	return o
}

func (o StreamInputIotHubOutput) ToStreamInputIotHubOutputWithContext(ctx context.Context) StreamInputIotHubOutput {
	return o
}

type StreamInputIotHubPtrOutput struct {
	*pulumi.OutputState
}

func (StreamInputIotHubPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamInputIotHub)(nil)).Elem()
}

func (o StreamInputIotHubPtrOutput) ToStreamInputIotHubPtrOutput() StreamInputIotHubPtrOutput {
	return o
}

func (o StreamInputIotHubPtrOutput) ToStreamInputIotHubPtrOutputWithContext(ctx context.Context) StreamInputIotHubPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StreamInputIotHubOutput{})
	pulumi.RegisterOutputType(StreamInputIotHubPtrOutput{})
}

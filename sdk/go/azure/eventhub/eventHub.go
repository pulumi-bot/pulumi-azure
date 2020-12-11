// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Event Hubs as a nested resource within a Event Hubs namespace.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/eventhub"
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
// 		exampleEventHubNamespace, err := eventhub.NewEventHubNamespace(ctx, "exampleEventHubNamespace", &eventhub.EventHubNamespaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Standard"),
// 			Capacity:          pulumi.Int(1),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("Production"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = eventhub.NewEventHub(ctx, "exampleEventHub", &eventhub.EventHubArgs{
// 			NamespaceName:     exampleEventHubNamespace.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			PartitionCount:    pulumi.Int(2),
// 			MessageRetention:  pulumi.Int(1),
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
// EventHubs can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:eventhub/eventHub:EventHub eventhub1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/eventhubs/eventhub1
// ```
type EventHub struct {
	pulumi.CustomResourceState

	// A `captureDescription` block as defined below.
	CaptureDescription EventHubCaptureDescriptionPtrOutput `pulumi:"captureDescription"`
	// Specifies the number of days to retain the events for this Event Hub.
	MessageRetention pulumi.IntOutput `pulumi:"messageRetention"`
	// Specifies the name of the EventHub resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// Specifies the current number of shards on the Event Hub. Changing this forces a new resource to be created.
	PartitionCount pulumi.IntOutput `pulumi:"partitionCount"`
	// The identifiers for partitions created for Event Hubs.
	PartitionIds pulumi.StringArrayOutput `pulumi:"partitionIds"`
	// The name of the resource group in which the EventHub's parent Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewEventHub registers a new resource with the given unique name, arguments, and options.
func NewEventHub(ctx *pulumi.Context,
	name string, args *EventHubArgs, opts ...pulumi.ResourceOption) (*EventHub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MessageRetention == nil {
		return nil, errors.New("invalid value for required argument 'MessageRetention'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.PartitionCount == nil {
		return nil, errors.New("invalid value for required argument 'PartitionCount'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource EventHub
	err := ctx.RegisterResource("azure:eventhub/eventHub:EventHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventHub gets an existing EventHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventHubState, opts ...pulumi.ResourceOption) (*EventHub, error) {
	var resource EventHub
	err := ctx.ReadResource("azure:eventhub/eventHub:EventHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventHub resources.
type eventHubState struct {
	// A `captureDescription` block as defined below.
	CaptureDescription *EventHubCaptureDescription `pulumi:"captureDescription"`
	// Specifies the number of days to retain the events for this Event Hub.
	MessageRetention *int `pulumi:"messageRetention"`
	// Specifies the name of the EventHub resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName *string `pulumi:"namespaceName"`
	// Specifies the current number of shards on the Event Hub. Changing this forces a new resource to be created.
	PartitionCount *int `pulumi:"partitionCount"`
	// The identifiers for partitions created for Event Hubs.
	PartitionIds []string `pulumi:"partitionIds"`
	// The name of the resource group in which the EventHub's parent Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type EventHubState struct {
	// A `captureDescription` block as defined below.
	CaptureDescription EventHubCaptureDescriptionPtrInput
	// Specifies the number of days to retain the events for this Event Hub.
	MessageRetention pulumi.IntPtrInput
	// Specifies the name of the EventHub resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringPtrInput
	// Specifies the current number of shards on the Event Hub. Changing this forces a new resource to be created.
	PartitionCount pulumi.IntPtrInput
	// The identifiers for partitions created for Event Hubs.
	PartitionIds pulumi.StringArrayInput
	// The name of the resource group in which the EventHub's parent Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (EventHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubState)(nil)).Elem()
}

type eventHubArgs struct {
	// A `captureDescription` block as defined below.
	CaptureDescription *EventHubCaptureDescription `pulumi:"captureDescription"`
	// Specifies the number of days to retain the events for this Event Hub.
	MessageRetention int `pulumi:"messageRetention"`
	// Specifies the name of the EventHub resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName string `pulumi:"namespaceName"`
	// Specifies the current number of shards on the Event Hub. Changing this forces a new resource to be created.
	PartitionCount int `pulumi:"partitionCount"`
	// The name of the resource group in which the EventHub's parent Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a EventHub resource.
type EventHubArgs struct {
	// A `captureDescription` block as defined below.
	CaptureDescription EventHubCaptureDescriptionPtrInput
	// Specifies the number of days to retain the events for this Event Hub.
	MessageRetention pulumi.IntInput
	// Specifies the name of the EventHub resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput
	// Specifies the current number of shards on the Event Hub. Changing this forces a new resource to be created.
	PartitionCount pulumi.IntInput
	// The name of the resource group in which the EventHub's parent Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (EventHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubArgs)(nil)).Elem()
}

type EventHubInput interface {
	pulumi.Input

	ToEventHubOutput() EventHubOutput
	ToEventHubOutputWithContext(ctx context.Context) EventHubOutput
}

type EventHubPtrInput interface {
	pulumi.Input

	ToEventHubPtrOutput() EventHubPtrOutput
	ToEventHubPtrOutputWithContext(ctx context.Context) EventHubPtrOutput
}

func (EventHub) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHub)(nil)).Elem()
}

func (i EventHub) ToEventHubOutput() EventHubOutput {
	return i.ToEventHubOutputWithContext(context.Background())
}

func (i EventHub) ToEventHubOutputWithContext(ctx context.Context) EventHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubOutput)
}

func (i EventHub) ToEventHubPtrOutput() EventHubPtrOutput {
	return i.ToEventHubPtrOutputWithContext(context.Background())
}

func (i EventHub) ToEventHubPtrOutputWithContext(ctx context.Context) EventHubPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubPtrOutput)
}

type EventHubOutput struct {
	*pulumi.OutputState
}

func (EventHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubOutput)(nil)).Elem()
}

func (o EventHubOutput) ToEventHubOutput() EventHubOutput {
	return o
}

func (o EventHubOutput) ToEventHubOutputWithContext(ctx context.Context) EventHubOutput {
	return o
}

type EventHubPtrOutput struct {
	*pulumi.OutputState
}

func (EventHubPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHub)(nil)).Elem()
}

func (o EventHubPtrOutput) ToEventHubPtrOutput() EventHubPtrOutput {
	return o
}

func (o EventHubPtrOutput) ToEventHubPtrOutputWithContext(ctx context.Context) EventHubPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EventHubOutput{})
	pulumi.RegisterOutputType(EventHubPtrOutput{})
}

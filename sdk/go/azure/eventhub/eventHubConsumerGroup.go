// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Event Hubs Consumer Group as a nested resource within an Event Hub.
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
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Basic"),
// 			Capacity:          pulumi.Int(2),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("Production"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleEventHub, err := eventhub.NewEventHub(ctx, "exampleEventHub", &eventhub.EventHubArgs{
// 			NamespaceName:     exampleEventHubNamespace.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			PartitionCount:    pulumi.Int(2),
// 			MessageRetention:  pulumi.Int(2),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = eventhub.NewConsumerGroup(ctx, "exampleConsumerGroup", &eventhub.ConsumerGroupArgs{
// 			NamespaceName:     exampleEventHubNamespace.Name,
// 			EventhubName:      exampleEventHub.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			UserMetadata:      pulumi.String("some-meta-data"),
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
// EventHub Consumer Groups can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:eventhub/eventHubConsumerGroup:EventHubConsumerGroup consumerGroup1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/eventhubs/eventhub1/consumergroups/consumerGroup1
// ```
//
// Deprecated: azure.eventhub.EventHubConsumerGroup has been deprecated in favor of azure.eventhub.ConsumerGroup
type EventHubConsumerGroup struct {
	pulumi.CustomResourceState

	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName pulumi.StringOutput `pulumi:"eventhubName"`
	// Specifies the name of the EventHub Consumer Group resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the user metadata.
	UserMetadata pulumi.StringPtrOutput `pulumi:"userMetadata"`
}

// NewEventHubConsumerGroup registers a new resource with the given unique name, arguments, and options.
func NewEventHubConsumerGroup(ctx *pulumi.Context,
	name string, args *EventHubConsumerGroupArgs, opts ...pulumi.ResourceOption) (*EventHubConsumerGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventhubName == nil {
		return nil, errors.New("invalid value for required argument 'EventhubName'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource EventHubConsumerGroup
	err := ctx.RegisterResource("azure:eventhub/eventHubConsumerGroup:EventHubConsumerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventHubConsumerGroup gets an existing EventHubConsumerGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventHubConsumerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventHubConsumerGroupState, opts ...pulumi.ResourceOption) (*EventHubConsumerGroup, error) {
	var resource EventHubConsumerGroup
	err := ctx.ReadResource("azure:eventhub/eventHubConsumerGroup:EventHubConsumerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventHubConsumerGroup resources.
type eventHubConsumerGroupState struct {
	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName *string `pulumi:"eventhubName"`
	// Specifies the name of the EventHub Consumer Group resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName *string `pulumi:"namespaceName"`
	// The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the user metadata.
	UserMetadata *string `pulumi:"userMetadata"`
}

type EventHubConsumerGroupState struct {
	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName pulumi.StringPtrInput
	// Specifies the name of the EventHub Consumer Group resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringPtrInput
	// The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the user metadata.
	UserMetadata pulumi.StringPtrInput
}

func (EventHubConsumerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubConsumerGroupState)(nil)).Elem()
}

type eventHubConsumerGroupArgs struct {
	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName string `pulumi:"eventhubName"`
	// Specifies the name of the EventHub Consumer Group resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the user metadata.
	UserMetadata *string `pulumi:"userMetadata"`
}

// The set of arguments for constructing a EventHubConsumerGroup resource.
type EventHubConsumerGroupArgs struct {
	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName pulumi.StringInput
	// Specifies the name of the EventHub Consumer Group resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput
	// The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the user metadata.
	UserMetadata pulumi.StringPtrInput
}

func (EventHubConsumerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubConsumerGroupArgs)(nil)).Elem()
}

type EventHubConsumerGroupInput interface {
	pulumi.Input

	ToEventHubConsumerGroupOutput() EventHubConsumerGroupOutput
	ToEventHubConsumerGroupOutputWithContext(ctx context.Context) EventHubConsumerGroupOutput
}

func (EventHubConsumerGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubConsumerGroup)(nil)).Elem()
}

func (i EventHubConsumerGroup) ToEventHubConsumerGroupOutput() EventHubConsumerGroupOutput {
	return i.ToEventHubConsumerGroupOutputWithContext(context.Background())
}

func (i EventHubConsumerGroup) ToEventHubConsumerGroupOutputWithContext(ctx context.Context) EventHubConsumerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubConsumerGroupOutput)
}

type EventHubConsumerGroupOutput struct {
	*pulumi.OutputState
}

func (EventHubConsumerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubConsumerGroupOutput)(nil)).Elem()
}

func (o EventHubConsumerGroupOutput) ToEventHubConsumerGroupOutput() EventHubConsumerGroupOutput {
	return o
}

func (o EventHubConsumerGroupOutput) ToEventHubConsumerGroupOutputWithContext(ctx context.Context) EventHubConsumerGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EventHubConsumerGroupOutput{})
}

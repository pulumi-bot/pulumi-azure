// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a ServiceBus Subscription.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/servicebus"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleNamespace, err := servicebus.NewNamespace(ctx, "exampleNamespace", &servicebus.NamespaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Standard"),
// 			Tags: pulumi.StringMap{
// 				"source": pulumi.String("example"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleTopic, err := servicebus.NewTopic(ctx, "exampleTopic", &servicebus.TopicArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			NamespaceName:      exampleNamespace.Name,
// 			EnablePartitioning: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = servicebus.NewSubscription(ctx, "exampleSubscription", &servicebus.SubscriptionArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			NamespaceName:     exampleNamespace.Name,
// 			TopicName:         exampleTopic.Name,
// 			MaxDeliveryCount:  pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Deprecated: azure.eventhub.Subscription has been deprecated in favor of azure.servicebus.Subscription
type Subscription struct {
	pulumi.CustomResourceState

	// The idle interval after which the topic is automatically deleted as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). The minimum duration is `5` minutes or `P5M`.
	AutoDeleteOnIdle pulumi.StringOutput `pulumi:"autoDeleteOnIdle"`
	// Boolean flag which controls whether the Subscription has dead letter support on filter evaluation exceptions. Defaults to `true`.
	DeadLetteringOnFilterEvaluationError pulumi.BoolPtrOutput `pulumi:"deadLetteringOnFilterEvaluationError"`
	// Boolean flag which controls whether the Subscription has dead letter support when a message expires. Defaults to `false`.
	DeadLetteringOnMessageExpiration pulumi.BoolPtrOutput `pulumi:"deadLetteringOnMessageExpiration"`
	// The Default message timespan to live as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTtl pulumi.StringOutput `pulumi:"defaultMessageTtl"`
	// Boolean flag which controls whether the Subscription supports batched operations. Defaults to `false`.
	EnableBatchedOperations pulumi.BoolPtrOutput `pulumi:"enableBatchedOperations"`
	// The name of a Queue or Topic to automatically forward Dead Letter messages to.
	ForwardDeadLetteredMessagesTo pulumi.StringPtrOutput `pulumi:"forwardDeadLetteredMessagesTo"`
	// The name of a Queue or Topic to automatically forward messages to.
	ForwardTo pulumi.StringPtrOutput `pulumi:"forwardTo"`
	// The lock duration for the subscription as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). The default value is `1` minute or `P1M`.
	LockDuration pulumi.StringOutput `pulumi:"lockDuration"`
	// The maximum number of deliveries.
	MaxDeliveryCount pulumi.IntOutput `pulumi:"maxDeliveryCount"`
	// Specifies the name of the ServiceBus Subscription resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the ServiceBus Namespace to create this Subscription in. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// Boolean flag which controls whether this Subscription supports the concept of a session. Defaults to `false`. Changing this forces a new resource to be created.
	RequiresSession pulumi.BoolPtrOutput `pulumi:"requiresSession"`
	// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The status of the Subscription. Possible values are `Active`,`ReceiveDisabled`, or `Disabled`. Defaults to `Active`.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// The name of the ServiceBus Topic to create this Subscription in. Changing this forces a new resource to be created.
	TopicName pulumi.StringOutput `pulumi:"topicName"`
}

// NewSubscription registers a new resource with the given unique name, arguments, and options.
func NewSubscription(ctx *pulumi.Context,
	name string, args *SubscriptionArgs, opts ...pulumi.ResourceOption) (*Subscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MaxDeliveryCount == nil {
		return nil, errors.New("invalid value for required argument 'MaxDeliveryCount'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	var resource Subscription
	err := ctx.RegisterResource("azure:eventhub/subscription:Subscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscription gets an existing Subscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionState, opts ...pulumi.ResourceOption) (*Subscription, error) {
	var resource Subscription
	err := ctx.ReadResource("azure:eventhub/subscription:Subscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subscription resources.
type subscriptionState struct {
	// The idle interval after which the topic is automatically deleted as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). The minimum duration is `5` minutes or `P5M`.
	AutoDeleteOnIdle *string `pulumi:"autoDeleteOnIdle"`
	// Boolean flag which controls whether the Subscription has dead letter support on filter evaluation exceptions. Defaults to `true`.
	DeadLetteringOnFilterEvaluationError *bool `pulumi:"deadLetteringOnFilterEvaluationError"`
	// Boolean flag which controls whether the Subscription has dead letter support when a message expires. Defaults to `false`.
	DeadLetteringOnMessageExpiration *bool `pulumi:"deadLetteringOnMessageExpiration"`
	// The Default message timespan to live as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTtl *string `pulumi:"defaultMessageTtl"`
	// Boolean flag which controls whether the Subscription supports batched operations. Defaults to `false`.
	EnableBatchedOperations *bool `pulumi:"enableBatchedOperations"`
	// The name of a Queue or Topic to automatically forward Dead Letter messages to.
	ForwardDeadLetteredMessagesTo *string `pulumi:"forwardDeadLetteredMessagesTo"`
	// The name of a Queue or Topic to automatically forward messages to.
	ForwardTo *string `pulumi:"forwardTo"`
	// The lock duration for the subscription as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). The default value is `1` minute or `P1M`.
	LockDuration *string `pulumi:"lockDuration"`
	// The maximum number of deliveries.
	MaxDeliveryCount *int `pulumi:"maxDeliveryCount"`
	// Specifies the name of the ServiceBus Subscription resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the ServiceBus Namespace to create this Subscription in. Changing this forces a new resource to be created.
	NamespaceName *string `pulumi:"namespaceName"`
	// Boolean flag which controls whether this Subscription supports the concept of a session. Defaults to `false`. Changing this forces a new resource to be created.
	RequiresSession *bool `pulumi:"requiresSession"`
	// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The status of the Subscription. Possible values are `Active`,`ReceiveDisabled`, or `Disabled`. Defaults to `Active`.
	Status *string `pulumi:"status"`
	// The name of the ServiceBus Topic to create this Subscription in. Changing this forces a new resource to be created.
	TopicName *string `pulumi:"topicName"`
}

type SubscriptionState struct {
	// The idle interval after which the topic is automatically deleted as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). The minimum duration is `5` minutes or `P5M`.
	AutoDeleteOnIdle pulumi.StringPtrInput
	// Boolean flag which controls whether the Subscription has dead letter support on filter evaluation exceptions. Defaults to `true`.
	DeadLetteringOnFilterEvaluationError pulumi.BoolPtrInput
	// Boolean flag which controls whether the Subscription has dead letter support when a message expires. Defaults to `false`.
	DeadLetteringOnMessageExpiration pulumi.BoolPtrInput
	// The Default message timespan to live as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTtl pulumi.StringPtrInput
	// Boolean flag which controls whether the Subscription supports batched operations. Defaults to `false`.
	EnableBatchedOperations pulumi.BoolPtrInput
	// The name of a Queue or Topic to automatically forward Dead Letter messages to.
	ForwardDeadLetteredMessagesTo pulumi.StringPtrInput
	// The name of a Queue or Topic to automatically forward messages to.
	ForwardTo pulumi.StringPtrInput
	// The lock duration for the subscription as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). The default value is `1` minute or `P1M`.
	LockDuration pulumi.StringPtrInput
	// The maximum number of deliveries.
	MaxDeliveryCount pulumi.IntPtrInput
	// Specifies the name of the ServiceBus Subscription resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the ServiceBus Namespace to create this Subscription in. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringPtrInput
	// Boolean flag which controls whether this Subscription supports the concept of a session. Defaults to `false`. Changing this forces a new resource to be created.
	RequiresSession pulumi.BoolPtrInput
	// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The status of the Subscription. Possible values are `Active`,`ReceiveDisabled`, or `Disabled`. Defaults to `Active`.
	Status pulumi.StringPtrInput
	// The name of the ServiceBus Topic to create this Subscription in. Changing this forces a new resource to be created.
	TopicName pulumi.StringPtrInput
}

func (SubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionState)(nil)).Elem()
}

type subscriptionArgs struct {
	// The idle interval after which the topic is automatically deleted as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). The minimum duration is `5` minutes or `P5M`.
	AutoDeleteOnIdle *string `pulumi:"autoDeleteOnIdle"`
	// Boolean flag which controls whether the Subscription has dead letter support on filter evaluation exceptions. Defaults to `true`.
	DeadLetteringOnFilterEvaluationError *bool `pulumi:"deadLetteringOnFilterEvaluationError"`
	// Boolean flag which controls whether the Subscription has dead letter support when a message expires. Defaults to `false`.
	DeadLetteringOnMessageExpiration *bool `pulumi:"deadLetteringOnMessageExpiration"`
	// The Default message timespan to live as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTtl *string `pulumi:"defaultMessageTtl"`
	// Boolean flag which controls whether the Subscription supports batched operations. Defaults to `false`.
	EnableBatchedOperations *bool `pulumi:"enableBatchedOperations"`
	// The name of a Queue or Topic to automatically forward Dead Letter messages to.
	ForwardDeadLetteredMessagesTo *string `pulumi:"forwardDeadLetteredMessagesTo"`
	// The name of a Queue or Topic to automatically forward messages to.
	ForwardTo *string `pulumi:"forwardTo"`
	// The lock duration for the subscription as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). The default value is `1` minute or `P1M`.
	LockDuration *string `pulumi:"lockDuration"`
	// The maximum number of deliveries.
	MaxDeliveryCount int `pulumi:"maxDeliveryCount"`
	// Specifies the name of the ServiceBus Subscription resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the ServiceBus Namespace to create this Subscription in. Changing this forces a new resource to be created.
	NamespaceName string `pulumi:"namespaceName"`
	// Boolean flag which controls whether this Subscription supports the concept of a session. Defaults to `false`. Changing this forces a new resource to be created.
	RequiresSession *bool `pulumi:"requiresSession"`
	// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The status of the Subscription. Possible values are `Active`,`ReceiveDisabled`, or `Disabled`. Defaults to `Active`.
	Status *string `pulumi:"status"`
	// The name of the ServiceBus Topic to create this Subscription in. Changing this forces a new resource to be created.
	TopicName string `pulumi:"topicName"`
}

// The set of arguments for constructing a Subscription resource.
type SubscriptionArgs struct {
	// The idle interval after which the topic is automatically deleted as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). The minimum duration is `5` minutes or `P5M`.
	AutoDeleteOnIdle pulumi.StringPtrInput
	// Boolean flag which controls whether the Subscription has dead letter support on filter evaluation exceptions. Defaults to `true`.
	DeadLetteringOnFilterEvaluationError pulumi.BoolPtrInput
	// Boolean flag which controls whether the Subscription has dead letter support when a message expires. Defaults to `false`.
	DeadLetteringOnMessageExpiration pulumi.BoolPtrInput
	// The Default message timespan to live as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTtl pulumi.StringPtrInput
	// Boolean flag which controls whether the Subscription supports batched operations. Defaults to `false`.
	EnableBatchedOperations pulumi.BoolPtrInput
	// The name of a Queue or Topic to automatically forward Dead Letter messages to.
	ForwardDeadLetteredMessagesTo pulumi.StringPtrInput
	// The name of a Queue or Topic to automatically forward messages to.
	ForwardTo pulumi.StringPtrInput
	// The lock duration for the subscription as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). The default value is `1` minute or `P1M`.
	LockDuration pulumi.StringPtrInput
	// The maximum number of deliveries.
	MaxDeliveryCount pulumi.IntInput
	// Specifies the name of the ServiceBus Subscription resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the ServiceBus Namespace to create this Subscription in. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput
	// Boolean flag which controls whether this Subscription supports the concept of a session. Defaults to `false`. Changing this forces a new resource to be created.
	RequiresSession pulumi.BoolPtrInput
	// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The status of the Subscription. Possible values are `Active`,`ReceiveDisabled`, or `Disabled`. Defaults to `Active`.
	Status pulumi.StringPtrInput
	// The name of the ServiceBus Topic to create this Subscription in. Changing this forces a new resource to be created.
	TopicName pulumi.StringInput
}

func (SubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionArgs)(nil)).Elem()
}

type SubscriptionInput interface {
	pulumi.Input

	ToSubscriptionOutput() SubscriptionOutput
	ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput
}

func (Subscription) ElementType() reflect.Type {
	return reflect.TypeOf((*Subscription)(nil)).Elem()
}

func (i Subscription) ToSubscriptionOutput() SubscriptionOutput {
	return i.ToSubscriptionOutputWithContext(context.Background())
}

func (i Subscription) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionOutput)
}

type SubscriptionOutput struct {
	*pulumi.OutputState
}

func (SubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionOutput)(nil)).Elem()
}

func (o SubscriptionOutput) ToSubscriptionOutput() SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SubscriptionOutput{})
}

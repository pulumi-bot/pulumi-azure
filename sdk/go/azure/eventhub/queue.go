// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a ServiceBus Queue.
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
// 		_, err = servicebus.NewQueue(ctx, "exampleQueue", &servicebus.QueueArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			NamespaceName:      exampleNamespace.Name,
// 			EnablePartitioning: pulumi.Bool(true),
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
// Service Bus Queue can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:eventhub/queue:Queue example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.servicebus/namespaces/sbns1/queues/snqueue1
// ```
//
// Deprecated: azure.eventhub.Queue has been deprecated in favor of azure.servicebus.Queue
type Queue struct {
	pulumi.CustomResourceState

	// The ISO 8601 timespan duration of the idle interval after which the Queue is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle pulumi.StringOutput `pulumi:"autoDeleteOnIdle"`
	// Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
	DeadLetteringOnMessageExpiration pulumi.BoolPtrOutput `pulumi:"deadLetteringOnMessageExpiration"`
	// The ISO 8601 timespan duration of the TTL of messages sent to this queue. This is the default value used when TTL is not set on message itself.
	DefaultMessageTtl pulumi.StringOutput `pulumi:"defaultMessageTtl"`
	// The ISO 8601 timespan duration during which duplicates can be detected. Defaults to 10 minutes (`PT10M`).
	DuplicateDetectionHistoryTimeWindow pulumi.StringOutput `pulumi:"duplicateDetectionHistoryTimeWindow"`
	// Boolean flag which controls whether server-side batched operations are enabled. Defaults to `true`.
	EnableBatchedOperations pulumi.BoolPtrOutput `pulumi:"enableBatchedOperations"`
	// Boolean flag which controls whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `false`.
	EnableExpress pulumi.BoolPtrOutput `pulumi:"enableExpress"`
	// Boolean flag which controls whether to enable the queue to be partitioned across multiple message brokers. Changing this forces a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `true`.
	EnablePartitioning pulumi.BoolPtrOutput `pulumi:"enablePartitioning"`
	// The name of a Queue or Topic to automatically forward dead lettered messages to.
	ForwardDeadLetteredMessagesTo pulumi.StringPtrOutput `pulumi:"forwardDeadLetteredMessagesTo"`
	// The name of a Queue or Topic to automatically forward messages to. Please [see the documentation](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-auto-forwarding) for more information.
	ForwardTo pulumi.StringPtrOutput `pulumi:"forwardTo"`
	// The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute (`PT1M`).
	LockDuration pulumi.StringOutput `pulumi:"lockDuration"`
	// Integer value which controls when a message is automatically dead lettered. Defaults to `10`.
	MaxDeliveryCount pulumi.IntPtrOutput `pulumi:"maxDeliveryCount"`
	// Integer value which controls the size of memory allocated for the queue. For supported values see the "Queue or topic size" section of [Service Bus Quotas](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas). Defaults to `1024`.
	MaxSizeInMegabytes pulumi.IntOutput `pulumi:"maxSizeInMegabytes"`
	// Specifies the name of the ServiceBus Queue resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the ServiceBus Namespace to create this queue in. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// Boolean flag which controls whether the Queue requires duplicate detection. Changing this forces a new resource to be created. Defaults to `false`.
	RequiresDuplicateDetection pulumi.BoolPtrOutput `pulumi:"requiresDuplicateDetection"`
	// Boolean flag which controls whether the Queue requires sessions. This will allow ordered handling of unbounded sequences of related messages. With sessions enabled a queue can guarantee first-in-first-out delivery of messages. Changing this forces a new resource to be created. Defaults to `false`.
	RequiresSession pulumi.BoolPtrOutput `pulumi:"requiresSession"`
	// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The status of the Queue. Possible values are `Active`, `Creating`, `Deleting`, `Disabled`, `ReceiveDisabled`, `Renaming`, `SendDisabled`, `Unknown`. Note that `Restoring` is not accepted. Defaults to `Active`.
	Status pulumi.StringPtrOutput `pulumi:"status"`
}

// NewQueue registers a new resource with the given unique name, arguments, and options.
func NewQueue(ctx *pulumi.Context,
	name string, args *QueueArgs, opts ...pulumi.ResourceOption) (*Queue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Queue
	err := ctx.RegisterResource("azure:eventhub/queue:Queue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueue gets an existing Queue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueState, opts ...pulumi.ResourceOption) (*Queue, error) {
	var resource Queue
	err := ctx.ReadResource("azure:eventhub/queue:Queue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Queue resources.
type queueState struct {
	// The ISO 8601 timespan duration of the idle interval after which the Queue is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle *string `pulumi:"autoDeleteOnIdle"`
	// Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
	DeadLetteringOnMessageExpiration *bool `pulumi:"deadLetteringOnMessageExpiration"`
	// The ISO 8601 timespan duration of the TTL of messages sent to this queue. This is the default value used when TTL is not set on message itself.
	DefaultMessageTtl *string `pulumi:"defaultMessageTtl"`
	// The ISO 8601 timespan duration during which duplicates can be detected. Defaults to 10 minutes (`PT10M`).
	DuplicateDetectionHistoryTimeWindow *string `pulumi:"duplicateDetectionHistoryTimeWindow"`
	// Boolean flag which controls whether server-side batched operations are enabled. Defaults to `true`.
	EnableBatchedOperations *bool `pulumi:"enableBatchedOperations"`
	// Boolean flag which controls whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `false`.
	EnableExpress *bool `pulumi:"enableExpress"`
	// Boolean flag which controls whether to enable the queue to be partitioned across multiple message brokers. Changing this forces a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `true`.
	EnablePartitioning *bool `pulumi:"enablePartitioning"`
	// The name of a Queue or Topic to automatically forward dead lettered messages to.
	ForwardDeadLetteredMessagesTo *string `pulumi:"forwardDeadLetteredMessagesTo"`
	// The name of a Queue or Topic to automatically forward messages to. Please [see the documentation](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-auto-forwarding) for more information.
	ForwardTo *string `pulumi:"forwardTo"`
	// The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute (`PT1M`).
	LockDuration *string `pulumi:"lockDuration"`
	// Integer value which controls when a message is automatically dead lettered. Defaults to `10`.
	MaxDeliveryCount *int `pulumi:"maxDeliveryCount"`
	// Integer value which controls the size of memory allocated for the queue. For supported values see the "Queue or topic size" section of [Service Bus Quotas](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas). Defaults to `1024`.
	MaxSizeInMegabytes *int `pulumi:"maxSizeInMegabytes"`
	// Specifies the name of the ServiceBus Queue resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the ServiceBus Namespace to create this queue in. Changing this forces a new resource to be created.
	NamespaceName *string `pulumi:"namespaceName"`
	// Boolean flag which controls whether the Queue requires duplicate detection. Changing this forces a new resource to be created. Defaults to `false`.
	RequiresDuplicateDetection *bool `pulumi:"requiresDuplicateDetection"`
	// Boolean flag which controls whether the Queue requires sessions. This will allow ordered handling of unbounded sequences of related messages. With sessions enabled a queue can guarantee first-in-first-out delivery of messages. Changing this forces a new resource to be created. Defaults to `false`.
	RequiresSession *bool `pulumi:"requiresSession"`
	// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The status of the Queue. Possible values are `Active`, `Creating`, `Deleting`, `Disabled`, `ReceiveDisabled`, `Renaming`, `SendDisabled`, `Unknown`. Note that `Restoring` is not accepted. Defaults to `Active`.
	Status *string `pulumi:"status"`
}

type QueueState struct {
	// The ISO 8601 timespan duration of the idle interval after which the Queue is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle pulumi.StringPtrInput
	// Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
	DeadLetteringOnMessageExpiration pulumi.BoolPtrInput
	// The ISO 8601 timespan duration of the TTL of messages sent to this queue. This is the default value used when TTL is not set on message itself.
	DefaultMessageTtl pulumi.StringPtrInput
	// The ISO 8601 timespan duration during which duplicates can be detected. Defaults to 10 minutes (`PT10M`).
	DuplicateDetectionHistoryTimeWindow pulumi.StringPtrInput
	// Boolean flag which controls whether server-side batched operations are enabled. Defaults to `true`.
	EnableBatchedOperations pulumi.BoolPtrInput
	// Boolean flag which controls whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `false`.
	EnableExpress pulumi.BoolPtrInput
	// Boolean flag which controls whether to enable the queue to be partitioned across multiple message brokers. Changing this forces a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `true`.
	EnablePartitioning pulumi.BoolPtrInput
	// The name of a Queue or Topic to automatically forward dead lettered messages to.
	ForwardDeadLetteredMessagesTo pulumi.StringPtrInput
	// The name of a Queue or Topic to automatically forward messages to. Please [see the documentation](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-auto-forwarding) for more information.
	ForwardTo pulumi.StringPtrInput
	// The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute (`PT1M`).
	LockDuration pulumi.StringPtrInput
	// Integer value which controls when a message is automatically dead lettered. Defaults to `10`.
	MaxDeliveryCount pulumi.IntPtrInput
	// Integer value which controls the size of memory allocated for the queue. For supported values see the "Queue or topic size" section of [Service Bus Quotas](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas). Defaults to `1024`.
	MaxSizeInMegabytes pulumi.IntPtrInput
	// Specifies the name of the ServiceBus Queue resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the ServiceBus Namespace to create this queue in. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringPtrInput
	// Boolean flag which controls whether the Queue requires duplicate detection. Changing this forces a new resource to be created. Defaults to `false`.
	RequiresDuplicateDetection pulumi.BoolPtrInput
	// Boolean flag which controls whether the Queue requires sessions. This will allow ordered handling of unbounded sequences of related messages. With sessions enabled a queue can guarantee first-in-first-out delivery of messages. Changing this forces a new resource to be created. Defaults to `false`.
	RequiresSession pulumi.BoolPtrInput
	// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The status of the Queue. Possible values are `Active`, `Creating`, `Deleting`, `Disabled`, `ReceiveDisabled`, `Renaming`, `SendDisabled`, `Unknown`. Note that `Restoring` is not accepted. Defaults to `Active`.
	Status pulumi.StringPtrInput
}

func (QueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueState)(nil)).Elem()
}

type queueArgs struct {
	// The ISO 8601 timespan duration of the idle interval after which the Queue is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle *string `pulumi:"autoDeleteOnIdle"`
	// Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
	DeadLetteringOnMessageExpiration *bool `pulumi:"deadLetteringOnMessageExpiration"`
	// The ISO 8601 timespan duration of the TTL of messages sent to this queue. This is the default value used when TTL is not set on message itself.
	DefaultMessageTtl *string `pulumi:"defaultMessageTtl"`
	// The ISO 8601 timespan duration during which duplicates can be detected. Defaults to 10 minutes (`PT10M`).
	DuplicateDetectionHistoryTimeWindow *string `pulumi:"duplicateDetectionHistoryTimeWindow"`
	// Boolean flag which controls whether server-side batched operations are enabled. Defaults to `true`.
	EnableBatchedOperations *bool `pulumi:"enableBatchedOperations"`
	// Boolean flag which controls whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `false`.
	EnableExpress *bool `pulumi:"enableExpress"`
	// Boolean flag which controls whether to enable the queue to be partitioned across multiple message brokers. Changing this forces a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `true`.
	EnablePartitioning *bool `pulumi:"enablePartitioning"`
	// The name of a Queue or Topic to automatically forward dead lettered messages to.
	ForwardDeadLetteredMessagesTo *string `pulumi:"forwardDeadLetteredMessagesTo"`
	// The name of a Queue or Topic to automatically forward messages to. Please [see the documentation](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-auto-forwarding) for more information.
	ForwardTo *string `pulumi:"forwardTo"`
	// The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute (`PT1M`).
	LockDuration *string `pulumi:"lockDuration"`
	// Integer value which controls when a message is automatically dead lettered. Defaults to `10`.
	MaxDeliveryCount *int `pulumi:"maxDeliveryCount"`
	// Integer value which controls the size of memory allocated for the queue. For supported values see the "Queue or topic size" section of [Service Bus Quotas](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas). Defaults to `1024`.
	MaxSizeInMegabytes *int `pulumi:"maxSizeInMegabytes"`
	// Specifies the name of the ServiceBus Queue resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the ServiceBus Namespace to create this queue in. Changing this forces a new resource to be created.
	NamespaceName string `pulumi:"namespaceName"`
	// Boolean flag which controls whether the Queue requires duplicate detection. Changing this forces a new resource to be created. Defaults to `false`.
	RequiresDuplicateDetection *bool `pulumi:"requiresDuplicateDetection"`
	// Boolean flag which controls whether the Queue requires sessions. This will allow ordered handling of unbounded sequences of related messages. With sessions enabled a queue can guarantee first-in-first-out delivery of messages. Changing this forces a new resource to be created. Defaults to `false`.
	RequiresSession *bool `pulumi:"requiresSession"`
	// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The status of the Queue. Possible values are `Active`, `Creating`, `Deleting`, `Disabled`, `ReceiveDisabled`, `Renaming`, `SendDisabled`, `Unknown`. Note that `Restoring` is not accepted. Defaults to `Active`.
	Status *string `pulumi:"status"`
}

// The set of arguments for constructing a Queue resource.
type QueueArgs struct {
	// The ISO 8601 timespan duration of the idle interval after which the Queue is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle pulumi.StringPtrInput
	// Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
	DeadLetteringOnMessageExpiration pulumi.BoolPtrInput
	// The ISO 8601 timespan duration of the TTL of messages sent to this queue. This is the default value used when TTL is not set on message itself.
	DefaultMessageTtl pulumi.StringPtrInput
	// The ISO 8601 timespan duration during which duplicates can be detected. Defaults to 10 minutes (`PT10M`).
	DuplicateDetectionHistoryTimeWindow pulumi.StringPtrInput
	// Boolean flag which controls whether server-side batched operations are enabled. Defaults to `true`.
	EnableBatchedOperations pulumi.BoolPtrInput
	// Boolean flag which controls whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `false`.
	EnableExpress pulumi.BoolPtrInput
	// Boolean flag which controls whether to enable the queue to be partitioned across multiple message brokers. Changing this forces a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `true`.
	EnablePartitioning pulumi.BoolPtrInput
	// The name of a Queue or Topic to automatically forward dead lettered messages to.
	ForwardDeadLetteredMessagesTo pulumi.StringPtrInput
	// The name of a Queue or Topic to automatically forward messages to. Please [see the documentation](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-auto-forwarding) for more information.
	ForwardTo pulumi.StringPtrInput
	// The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute (`PT1M`).
	LockDuration pulumi.StringPtrInput
	// Integer value which controls when a message is automatically dead lettered. Defaults to `10`.
	MaxDeliveryCount pulumi.IntPtrInput
	// Integer value which controls the size of memory allocated for the queue. For supported values see the "Queue or topic size" section of [Service Bus Quotas](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas). Defaults to `1024`.
	MaxSizeInMegabytes pulumi.IntPtrInput
	// Specifies the name of the ServiceBus Queue resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the ServiceBus Namespace to create this queue in. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput
	// Boolean flag which controls whether the Queue requires duplicate detection. Changing this forces a new resource to be created. Defaults to `false`.
	RequiresDuplicateDetection pulumi.BoolPtrInput
	// Boolean flag which controls whether the Queue requires sessions. This will allow ordered handling of unbounded sequences of related messages. With sessions enabled a queue can guarantee first-in-first-out delivery of messages. Changing this forces a new resource to be created. Defaults to `false`.
	RequiresSession pulumi.BoolPtrInput
	// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The status of the Queue. Possible values are `Active`, `Creating`, `Deleting`, `Disabled`, `ReceiveDisabled`, `Renaming`, `SendDisabled`, `Unknown`. Note that `Restoring` is not accepted. Defaults to `Active`.
	Status pulumi.StringPtrInput
}

func (QueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueArgs)(nil)).Elem()
}

type QueueInput interface {
	pulumi.Input

	ToQueueOutput() QueueOutput
	ToQueueOutputWithContext(ctx context.Context) QueueOutput
}

func (*Queue) ElementType() reflect.Type {
	return reflect.TypeOf((*Queue)(nil))
}

func (i *Queue) ToQueueOutput() QueueOutput {
	return i.ToQueueOutputWithContext(context.Background())
}

func (i *Queue) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueOutput)
}

func (i *Queue) ToQueuePtrOutput() QueuePtrOutput {
	return i.ToQueuePtrOutputWithContext(context.Background())
}

func (i *Queue) ToQueuePtrOutputWithContext(ctx context.Context) QueuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuePtrOutput)
}

type QueuePtrInput interface {
	pulumi.Input

	ToQueuePtrOutput() QueuePtrOutput
	ToQueuePtrOutputWithContext(ctx context.Context) QueuePtrOutput
}

type QueueOutput struct {
	*pulumi.OutputState
}

func (QueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Queue)(nil))
}

func (o QueueOutput) ToQueueOutput() QueueOutput {
	return o
}

func (o QueueOutput) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return o
}

type QueuePtrOutput struct {
	*pulumi.OutputState
}

func (QueuePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil))
}

func (o QueuePtrOutput) ToQueuePtrOutput() QueuePtrOutput {
	return o
}

func (o QueuePtrOutput) ToQueuePtrOutputWithContext(ctx context.Context) QueuePtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(QueueOutput{})
	pulumi.RegisterOutputType(QueuePtrOutput{})
}

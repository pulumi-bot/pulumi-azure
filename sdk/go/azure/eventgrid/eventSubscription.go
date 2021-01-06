// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventgrid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an EventGrid Event Subscription
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/eventgrid"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultResourceGroup, err := core.NewResourceGroup(ctx, "defaultResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US 2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultAccount, err := storage.NewAccount(ctx, "defaultAccount", &storage.AccountArgs{
// 			ResourceGroupName:      defaultResourceGroup.Name,
// 			Location:               defaultResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("staging"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultQueue, err := storage.NewQueue(ctx, "defaultQueue", &storage.QueueArgs{
// 			StorageAccountName: defaultAccount.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = eventgrid.NewEventSubscription(ctx, "defaultEventSubscription", &eventgrid.EventSubscriptionArgs{
// 			Scope: defaultResourceGroup.ID(),
// 			StorageQueueEndpoint: &eventgrid.EventSubscriptionStorageQueueEndpointArgs{
// 				StorageAccountId: defaultAccount.ID(),
// 				QueueName:        defaultQueue.Name,
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
// EventGrid Event Subscription's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:eventgrid/eventSubscription:EventSubscription eventSubscription1
// ```
//
//  /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/topics/topic1/providers/Microsoft.EventGrid/eventSubscriptions/eventSubscription1
type EventSubscription struct {
	pulumi.CustomResourceState

	// A `advancedFilter` block as defined below.
	AdvancedFilter EventSubscriptionAdvancedFilterPtrOutput `pulumi:"advancedFilter"`
	// An `azureFunctionEndpoint` block as defined below.
	AzureFunctionEndpoint EventSubscriptionAzureFunctionEndpointPtrOutput `pulumi:"azureFunctionEndpoint"`
	// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	EventDeliverySchema pulumi.StringPtrOutput `pulumi:"eventDeliverySchema"`
	// A `eventhubEndpoint` block as defined below.
	//
	// Deprecated: Deprecated in favour of `eventhub_endpoint_id`
	EventhubEndpoint EventSubscriptionEventhubEndpointOutput `pulumi:"eventhubEndpoint"`
	// Specifies the id where the Event Hub is located.
	EventhubEndpointId pulumi.StringOutput `pulumi:"eventhubEndpointId"`
	// Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
	ExpirationTimeUtc pulumi.StringPtrOutput `pulumi:"expirationTimeUtc"`
	// A `hybridConnectionEndpoint` block as defined below.
	//
	// Deprecated: Deprecated in favour of `hybrid_connection_endpoint_id`
	HybridConnectionEndpoint EventSubscriptionHybridConnectionEndpointOutput `pulumi:"hybridConnectionEndpoint"`
	// Specifies the id where the Hybrid Connection is located.
	HybridConnectionEndpointId pulumi.StringOutput `pulumi:"hybridConnectionEndpointId"`
	// A list of applicable event types that need to be part of the event subscription.
	IncludedEventTypes pulumi.StringArrayOutput `pulumi:"includedEventTypes"`
	// A list of labels to assign to the event subscription.
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `retryPolicy` block as defined below.
	RetryPolicy EventSubscriptionRetryPolicyOutput `pulumi:"retryPolicy"`
	// Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
	Scope pulumi.StringOutput `pulumi:"scope"`
	// Specifies the id where the Service Bus Queue is located.
	ServiceBusQueueEndpointId pulumi.StringPtrOutput `pulumi:"serviceBusQueueEndpointId"`
	// Specifies the id where the Service Bus Topic is located.
	ServiceBusTopicEndpointId pulumi.StringPtrOutput `pulumi:"serviceBusTopicEndpointId"`
	// A `storageBlobDeadLetterDestination` block as defined below.
	StorageBlobDeadLetterDestination EventSubscriptionStorageBlobDeadLetterDestinationPtrOutput `pulumi:"storageBlobDeadLetterDestination"`
	// A `storageQueueEndpoint` block as defined below.
	StorageQueueEndpoint EventSubscriptionStorageQueueEndpointPtrOutput `pulumi:"storageQueueEndpoint"`
	// A `subjectFilter` block as defined below.
	SubjectFilter EventSubscriptionSubjectFilterPtrOutput `pulumi:"subjectFilter"`
	// (Optional/ **Deprecated) Specifies the name of the topic to associate with the event subscription.
	//
	// Deprecated: This field has been updated to readonly field since Apr 25, 2019 so no longer has any affect and will be removed in version 3.0 of the provider.
	TopicName pulumi.StringOutput `pulumi:"topicName"`
	// A `webhookEndpoint` block as defined below.
	WebhookEndpoint EventSubscriptionWebhookEndpointPtrOutput `pulumi:"webhookEndpoint"`
}

// NewEventSubscription registers a new resource with the given unique name, arguments, and options.
func NewEventSubscription(ctx *pulumi.Context,
	name string, args *EventSubscriptionArgs, opts ...pulumi.ResourceOption) (*EventSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure:eventhub/eventSubscription:EventSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource EventSubscription
	err := ctx.RegisterResource("azure:eventgrid/eventSubscription:EventSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventSubscription gets an existing EventSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventSubscriptionState, opts ...pulumi.ResourceOption) (*EventSubscription, error) {
	var resource EventSubscription
	err := ctx.ReadResource("azure:eventgrid/eventSubscription:EventSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventSubscription resources.
type eventSubscriptionState struct {
	// A `advancedFilter` block as defined below.
	AdvancedFilter *EventSubscriptionAdvancedFilter `pulumi:"advancedFilter"`
	// An `azureFunctionEndpoint` block as defined below.
	AzureFunctionEndpoint *EventSubscriptionAzureFunctionEndpoint `pulumi:"azureFunctionEndpoint"`
	// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	EventDeliverySchema *string `pulumi:"eventDeliverySchema"`
	// A `eventhubEndpoint` block as defined below.
	//
	// Deprecated: Deprecated in favour of `eventhub_endpoint_id`
	EventhubEndpoint *EventSubscriptionEventhubEndpoint `pulumi:"eventhubEndpoint"`
	// Specifies the id where the Event Hub is located.
	EventhubEndpointId *string `pulumi:"eventhubEndpointId"`
	// Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
	ExpirationTimeUtc *string `pulumi:"expirationTimeUtc"`
	// A `hybridConnectionEndpoint` block as defined below.
	//
	// Deprecated: Deprecated in favour of `hybrid_connection_endpoint_id`
	HybridConnectionEndpoint *EventSubscriptionHybridConnectionEndpoint `pulumi:"hybridConnectionEndpoint"`
	// Specifies the id where the Hybrid Connection is located.
	HybridConnectionEndpointId *string `pulumi:"hybridConnectionEndpointId"`
	// A list of applicable event types that need to be part of the event subscription.
	IncludedEventTypes []string `pulumi:"includedEventTypes"`
	// A list of labels to assign to the event subscription.
	Labels []string `pulumi:"labels"`
	// Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `retryPolicy` block as defined below.
	RetryPolicy *EventSubscriptionRetryPolicy `pulumi:"retryPolicy"`
	// Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
	Scope *string `pulumi:"scope"`
	// Specifies the id where the Service Bus Queue is located.
	ServiceBusQueueEndpointId *string `pulumi:"serviceBusQueueEndpointId"`
	// Specifies the id where the Service Bus Topic is located.
	ServiceBusTopicEndpointId *string `pulumi:"serviceBusTopicEndpointId"`
	// A `storageBlobDeadLetterDestination` block as defined below.
	StorageBlobDeadLetterDestination *EventSubscriptionStorageBlobDeadLetterDestination `pulumi:"storageBlobDeadLetterDestination"`
	// A `storageQueueEndpoint` block as defined below.
	StorageQueueEndpoint *EventSubscriptionStorageQueueEndpoint `pulumi:"storageQueueEndpoint"`
	// A `subjectFilter` block as defined below.
	SubjectFilter *EventSubscriptionSubjectFilter `pulumi:"subjectFilter"`
	// (Optional/ **Deprecated) Specifies the name of the topic to associate with the event subscription.
	//
	// Deprecated: This field has been updated to readonly field since Apr 25, 2019 so no longer has any affect and will be removed in version 3.0 of the provider.
	TopicName *string `pulumi:"topicName"`
	// A `webhookEndpoint` block as defined below.
	WebhookEndpoint *EventSubscriptionWebhookEndpoint `pulumi:"webhookEndpoint"`
}

type EventSubscriptionState struct {
	// A `advancedFilter` block as defined below.
	AdvancedFilter EventSubscriptionAdvancedFilterPtrInput
	// An `azureFunctionEndpoint` block as defined below.
	AzureFunctionEndpoint EventSubscriptionAzureFunctionEndpointPtrInput
	// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	EventDeliverySchema pulumi.StringPtrInput
	// A `eventhubEndpoint` block as defined below.
	//
	// Deprecated: Deprecated in favour of `eventhub_endpoint_id`
	EventhubEndpoint EventSubscriptionEventhubEndpointPtrInput
	// Specifies the id where the Event Hub is located.
	EventhubEndpointId pulumi.StringPtrInput
	// Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
	ExpirationTimeUtc pulumi.StringPtrInput
	// A `hybridConnectionEndpoint` block as defined below.
	//
	// Deprecated: Deprecated in favour of `hybrid_connection_endpoint_id`
	HybridConnectionEndpoint EventSubscriptionHybridConnectionEndpointPtrInput
	// Specifies the id where the Hybrid Connection is located.
	HybridConnectionEndpointId pulumi.StringPtrInput
	// A list of applicable event types that need to be part of the event subscription.
	IncludedEventTypes pulumi.StringArrayInput
	// A list of labels to assign to the event subscription.
	Labels pulumi.StringArrayInput
	// Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `retryPolicy` block as defined below.
	RetryPolicy EventSubscriptionRetryPolicyPtrInput
	// Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
	Scope pulumi.StringPtrInput
	// Specifies the id where the Service Bus Queue is located.
	ServiceBusQueueEndpointId pulumi.StringPtrInput
	// Specifies the id where the Service Bus Topic is located.
	ServiceBusTopicEndpointId pulumi.StringPtrInput
	// A `storageBlobDeadLetterDestination` block as defined below.
	StorageBlobDeadLetterDestination EventSubscriptionStorageBlobDeadLetterDestinationPtrInput
	// A `storageQueueEndpoint` block as defined below.
	StorageQueueEndpoint EventSubscriptionStorageQueueEndpointPtrInput
	// A `subjectFilter` block as defined below.
	SubjectFilter EventSubscriptionSubjectFilterPtrInput
	// (Optional/ **Deprecated) Specifies the name of the topic to associate with the event subscription.
	//
	// Deprecated: This field has been updated to readonly field since Apr 25, 2019 so no longer has any affect and will be removed in version 3.0 of the provider.
	TopicName pulumi.StringPtrInput
	// A `webhookEndpoint` block as defined below.
	WebhookEndpoint EventSubscriptionWebhookEndpointPtrInput
}

func (EventSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionState)(nil)).Elem()
}

type eventSubscriptionArgs struct {
	// A `advancedFilter` block as defined below.
	AdvancedFilter *EventSubscriptionAdvancedFilter `pulumi:"advancedFilter"`
	// An `azureFunctionEndpoint` block as defined below.
	AzureFunctionEndpoint *EventSubscriptionAzureFunctionEndpoint `pulumi:"azureFunctionEndpoint"`
	// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	EventDeliverySchema *string `pulumi:"eventDeliverySchema"`
	// A `eventhubEndpoint` block as defined below.
	//
	// Deprecated: Deprecated in favour of `eventhub_endpoint_id`
	EventhubEndpoint *EventSubscriptionEventhubEndpoint `pulumi:"eventhubEndpoint"`
	// Specifies the id where the Event Hub is located.
	EventhubEndpointId *string `pulumi:"eventhubEndpointId"`
	// Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
	ExpirationTimeUtc *string `pulumi:"expirationTimeUtc"`
	// A `hybridConnectionEndpoint` block as defined below.
	//
	// Deprecated: Deprecated in favour of `hybrid_connection_endpoint_id`
	HybridConnectionEndpoint *EventSubscriptionHybridConnectionEndpoint `pulumi:"hybridConnectionEndpoint"`
	// Specifies the id where the Hybrid Connection is located.
	HybridConnectionEndpointId *string `pulumi:"hybridConnectionEndpointId"`
	// A list of applicable event types that need to be part of the event subscription.
	IncludedEventTypes []string `pulumi:"includedEventTypes"`
	// A list of labels to assign to the event subscription.
	Labels []string `pulumi:"labels"`
	// Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `retryPolicy` block as defined below.
	RetryPolicy *EventSubscriptionRetryPolicy `pulumi:"retryPolicy"`
	// Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
	Scope string `pulumi:"scope"`
	// Specifies the id where the Service Bus Queue is located.
	ServiceBusQueueEndpointId *string `pulumi:"serviceBusQueueEndpointId"`
	// Specifies the id where the Service Bus Topic is located.
	ServiceBusTopicEndpointId *string `pulumi:"serviceBusTopicEndpointId"`
	// A `storageBlobDeadLetterDestination` block as defined below.
	StorageBlobDeadLetterDestination *EventSubscriptionStorageBlobDeadLetterDestination `pulumi:"storageBlobDeadLetterDestination"`
	// A `storageQueueEndpoint` block as defined below.
	StorageQueueEndpoint *EventSubscriptionStorageQueueEndpoint `pulumi:"storageQueueEndpoint"`
	// A `subjectFilter` block as defined below.
	SubjectFilter *EventSubscriptionSubjectFilter `pulumi:"subjectFilter"`
	// (Optional/ **Deprecated) Specifies the name of the topic to associate with the event subscription.
	//
	// Deprecated: This field has been updated to readonly field since Apr 25, 2019 so no longer has any affect and will be removed in version 3.0 of the provider.
	TopicName *string `pulumi:"topicName"`
	// A `webhookEndpoint` block as defined below.
	WebhookEndpoint *EventSubscriptionWebhookEndpoint `pulumi:"webhookEndpoint"`
}

// The set of arguments for constructing a EventSubscription resource.
type EventSubscriptionArgs struct {
	// A `advancedFilter` block as defined below.
	AdvancedFilter EventSubscriptionAdvancedFilterPtrInput
	// An `azureFunctionEndpoint` block as defined below.
	AzureFunctionEndpoint EventSubscriptionAzureFunctionEndpointPtrInput
	// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	EventDeliverySchema pulumi.StringPtrInput
	// A `eventhubEndpoint` block as defined below.
	//
	// Deprecated: Deprecated in favour of `eventhub_endpoint_id`
	EventhubEndpoint EventSubscriptionEventhubEndpointPtrInput
	// Specifies the id where the Event Hub is located.
	EventhubEndpointId pulumi.StringPtrInput
	// Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
	ExpirationTimeUtc pulumi.StringPtrInput
	// A `hybridConnectionEndpoint` block as defined below.
	//
	// Deprecated: Deprecated in favour of `hybrid_connection_endpoint_id`
	HybridConnectionEndpoint EventSubscriptionHybridConnectionEndpointPtrInput
	// Specifies the id where the Hybrid Connection is located.
	HybridConnectionEndpointId pulumi.StringPtrInput
	// A list of applicable event types that need to be part of the event subscription.
	IncludedEventTypes pulumi.StringArrayInput
	// A list of labels to assign to the event subscription.
	Labels pulumi.StringArrayInput
	// Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `retryPolicy` block as defined below.
	RetryPolicy EventSubscriptionRetryPolicyPtrInput
	// Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
	Scope pulumi.StringInput
	// Specifies the id where the Service Bus Queue is located.
	ServiceBusQueueEndpointId pulumi.StringPtrInput
	// Specifies the id where the Service Bus Topic is located.
	ServiceBusTopicEndpointId pulumi.StringPtrInput
	// A `storageBlobDeadLetterDestination` block as defined below.
	StorageBlobDeadLetterDestination EventSubscriptionStorageBlobDeadLetterDestinationPtrInput
	// A `storageQueueEndpoint` block as defined below.
	StorageQueueEndpoint EventSubscriptionStorageQueueEndpointPtrInput
	// A `subjectFilter` block as defined below.
	SubjectFilter EventSubscriptionSubjectFilterPtrInput
	// (Optional/ **Deprecated) Specifies the name of the topic to associate with the event subscription.
	//
	// Deprecated: This field has been updated to readonly field since Apr 25, 2019 so no longer has any affect and will be removed in version 3.0 of the provider.
	TopicName pulumi.StringPtrInput
	// A `webhookEndpoint` block as defined below.
	WebhookEndpoint EventSubscriptionWebhookEndpointPtrInput
}

func (EventSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionArgs)(nil)).Elem()
}

type EventSubscriptionInput interface {
	pulumi.Input

	ToEventSubscriptionOutput() EventSubscriptionOutput
	ToEventSubscriptionOutputWithContext(ctx context.Context) EventSubscriptionOutput
}

func (*EventSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscription)(nil))
}

func (i *EventSubscription) ToEventSubscriptionOutput() EventSubscriptionOutput {
	return i.ToEventSubscriptionOutputWithContext(context.Background())
}

func (i *EventSubscription) ToEventSubscriptionOutputWithContext(ctx context.Context) EventSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionOutput)
}

func (i *EventSubscription) ToEventSubscriptionPtrOutput() EventSubscriptionPtrOutput {
	return i.ToEventSubscriptionPtrOutputWithContext(context.Background())
}

func (i *EventSubscription) ToEventSubscriptionPtrOutputWithContext(ctx context.Context) EventSubscriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionPtrOutput)
}

type EventSubscriptionPtrInput interface {
	pulumi.Input

	ToEventSubscriptionPtrOutput() EventSubscriptionPtrOutput
	ToEventSubscriptionPtrOutputWithContext(ctx context.Context) EventSubscriptionPtrOutput
}

type eventSubscriptionPtrType EventSubscriptionArgs

func (*eventSubscriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscription)(nil))
}

func (i *eventSubscriptionPtrType) ToEventSubscriptionPtrOutput() EventSubscriptionPtrOutput {
	return i.ToEventSubscriptionPtrOutputWithContext(context.Background())
}

func (i *eventSubscriptionPtrType) ToEventSubscriptionPtrOutputWithContext(ctx context.Context) EventSubscriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionOutput).ToEventSubscriptionPtrOutput()
}

// EventSubscriptionArrayInput is an input type that accepts EventSubscriptionArray and EventSubscriptionArrayOutput values.
// You can construct a concrete instance of `EventSubscriptionArrayInput` via:
//
//          EventSubscriptionArray{ EventSubscriptionArgs{...} }
type EventSubscriptionArrayInput interface {
	pulumi.Input

	ToEventSubscriptionArrayOutput() EventSubscriptionArrayOutput
	ToEventSubscriptionArrayOutputWithContext(context.Context) EventSubscriptionArrayOutput
}

type EventSubscriptionArray []EventSubscriptionInput

func (EventSubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventSubscription)(nil))
}

func (i EventSubscriptionArray) ToEventSubscriptionArrayOutput() EventSubscriptionArrayOutput {
	return i.ToEventSubscriptionArrayOutputWithContext(context.Background())
}

func (i EventSubscriptionArray) ToEventSubscriptionArrayOutputWithContext(ctx context.Context) EventSubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionArrayOutput)
}

// EventSubscriptionMapInput is an input type that accepts EventSubscriptionMap and EventSubscriptionMapOutput values.
// You can construct a concrete instance of `EventSubscriptionMapInput` via:
//
//          EventSubscriptionMap{ "key": EventSubscriptionArgs{...} }
type EventSubscriptionMapInput interface {
	pulumi.Input

	ToEventSubscriptionMapOutput() EventSubscriptionMapOutput
	ToEventSubscriptionMapOutputWithContext(context.Context) EventSubscriptionMapOutput
}

type EventSubscriptionMap map[string]EventSubscriptionInput

func (EventSubscriptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EventSubscription)(nil))
}

func (i EventSubscriptionMap) ToEventSubscriptionMapOutput() EventSubscriptionMapOutput {
	return i.ToEventSubscriptionMapOutputWithContext(context.Background())
}

func (i EventSubscriptionMap) ToEventSubscriptionMapOutputWithContext(ctx context.Context) EventSubscriptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionMapOutput)
}

type EventSubscriptionOutput struct {
	*pulumi.OutputState
}

func (EventSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscription)(nil))
}

func (o EventSubscriptionOutput) ToEventSubscriptionOutput() EventSubscriptionOutput {
	return o
}

func (o EventSubscriptionOutput) ToEventSubscriptionOutputWithContext(ctx context.Context) EventSubscriptionOutput {
	return o
}

func (o EventSubscriptionOutput) ToEventSubscriptionPtrOutput() EventSubscriptionPtrOutput {
	return o.ToEventSubscriptionPtrOutputWithContext(context.Background())
}

func (o EventSubscriptionOutput) ToEventSubscriptionPtrOutputWithContext(ctx context.Context) EventSubscriptionPtrOutput {
	return o.ApplyT(func(v EventSubscription) *EventSubscription {
		return &v
	}).(EventSubscriptionPtrOutput)
}

type EventSubscriptionPtrOutput struct {
	*pulumi.OutputState
}

func (EventSubscriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscription)(nil))
}

func (o EventSubscriptionPtrOutput) ToEventSubscriptionPtrOutput() EventSubscriptionPtrOutput {
	return o
}

func (o EventSubscriptionPtrOutput) ToEventSubscriptionPtrOutputWithContext(ctx context.Context) EventSubscriptionPtrOutput {
	return o
}

type EventSubscriptionArrayOutput struct{ *pulumi.OutputState }

func (EventSubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventSubscription)(nil))
}

func (o EventSubscriptionArrayOutput) ToEventSubscriptionArrayOutput() EventSubscriptionArrayOutput {
	return o
}

func (o EventSubscriptionArrayOutput) ToEventSubscriptionArrayOutputWithContext(ctx context.Context) EventSubscriptionArrayOutput {
	return o
}

func (o EventSubscriptionArrayOutput) Index(i pulumi.IntInput) EventSubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventSubscription {
		return vs[0].([]EventSubscription)[vs[1].(int)]
	}).(EventSubscriptionOutput)
}

type EventSubscriptionMapOutput struct{ *pulumi.OutputState }

func (EventSubscriptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EventSubscription)(nil))
}

func (o EventSubscriptionMapOutput) ToEventSubscriptionMapOutput() EventSubscriptionMapOutput {
	return o
}

func (o EventSubscriptionMapOutput) ToEventSubscriptionMapOutputWithContext(ctx context.Context) EventSubscriptionMapOutput {
	return o
}

func (o EventSubscriptionMapOutput) MapIndex(k pulumi.StringInput) EventSubscriptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EventSubscription {
		return vs[0].(map[string]EventSubscription)[vs[1].(string)]
	}).(EventSubscriptionOutput)
}

func init() {
	pulumi.RegisterOutputType(EventSubscriptionOutput{})
	pulumi.RegisterOutputType(EventSubscriptionPtrOutput{})
	pulumi.RegisterOutputType(EventSubscriptionArrayOutput{})
	pulumi.RegisterOutputType(EventSubscriptionMapOutput{})
}

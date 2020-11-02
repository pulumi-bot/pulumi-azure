// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventHub
{
    /// <summary>
    /// Manages an EventGrid Event Subscription
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var defaultResourceGroup = new Azure.Core.ResourceGroup("defaultResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West US 2",
    ///         });
    ///         var defaultAccount = new Azure.Storage.Account("defaultAccount", new Azure.Storage.AccountArgs
    ///         {
    ///             ResourceGroupName = defaultResourceGroup.Name,
    ///             Location = defaultResourceGroup.Location,
    ///             AccountTier = "Standard",
    ///             AccountReplicationType = "LRS",
    ///             Tags = 
    ///             {
    ///                 { "environment", "staging" },
    ///             },
    ///         });
    ///         var defaultQueue = new Azure.Storage.Queue("defaultQueue", new Azure.Storage.QueueArgs
    ///         {
    ///             StorageAccountName = defaultAccount.Name,
    ///         });
    ///         var defaultEventSubscription = new Azure.EventGrid.EventSubscription("defaultEventSubscription", new Azure.EventGrid.EventSubscriptionArgs
    ///         {
    ///             Scope = defaultResourceGroup.Id,
    ///             StorageQueueEndpoint = new Azure.EventGrid.Inputs.EventSubscriptionStorageQueueEndpointArgs
    ///             {
    ///                 StorageAccountId = defaultAccount.Id,
    ///                 QueueName = defaultQueue.Name,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// EventGrid Event Subscription's can be imported using the `resource id`, e.g.
    /// 
    ///  /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/topics/topic1/providers/Microsoft.EventGrid/eventSubscriptions/eventSubscription1
    /// </summary>
    [Obsolete(@"azure.eventhub.EventSubscription has been deprecated in favor of azure.eventgrid.EventSubscription")]
    public partial class EventSubscription : Pulumi.CustomResource
    {
        /// <summary>
        /// A `advanced_filter` block as defined below.
        /// </summary>
        [Output("advancedFilter")]
        public Output<Outputs.EventSubscriptionAdvancedFilter?> AdvancedFilter { get; private set; } = null!;

        /// <summary>
        /// An `azure_function_endpoint` block as defined below.
        /// </summary>
        [Output("azureFunctionEndpoint")]
        public Output<Outputs.EventSubscriptionAzureFunctionEndpoint?> AzureFunctionEndpoint { get; private set; } = null!;

        /// <summary>
        /// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("eventDeliverySchema")]
        public Output<string?> EventDeliverySchema { get; private set; } = null!;

        /// <summary>
        /// A `eventhub_endpoint` block as defined below.
        /// </summary>
        [Output("eventhubEndpoint")]
        public Output<Outputs.EventSubscriptionEventhubEndpoint> EventhubEndpoint { get; private set; } = null!;

        /// <summary>
        /// Specifies the id where the Event Hub is located.
        /// </summary>
        [Output("eventhubEndpointId")]
        public Output<string> EventhubEndpointId { get; private set; } = null!;

        /// <summary>
        /// Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
        /// </summary>
        [Output("expirationTimeUtc")]
        public Output<string?> ExpirationTimeUtc { get; private set; } = null!;

        /// <summary>
        /// A `hybrid_connection_endpoint` block as defined below.
        /// </summary>
        [Output("hybridConnectionEndpoint")]
        public Output<Outputs.EventSubscriptionHybridConnectionEndpoint> HybridConnectionEndpoint { get; private set; } = null!;

        /// <summary>
        /// Specifies the id where the Hybrid Connection is located.
        /// </summary>
        [Output("hybridConnectionEndpointId")]
        public Output<string> HybridConnectionEndpointId { get; private set; } = null!;

        /// <summary>
        /// A list of applicable event types that need to be part of the event subscription.
        /// </summary>
        [Output("includedEventTypes")]
        public Output<ImmutableArray<string>> IncludedEventTypes { get; private set; } = null!;

        /// <summary>
        /// A list of labels to assign to the event subscription.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A `retry_policy` block as defined below.
        /// </summary>
        [Output("retryPolicy")]
        public Output<Outputs.EventSubscriptionRetryPolicy> RetryPolicy { get; private set; } = null!;

        /// <summary>
        /// Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("scope")]
        public Output<string> Scope { get; private set; } = null!;

        /// <summary>
        /// Specifies the id where the Service Bus Queue is located.
        /// </summary>
        [Output("serviceBusQueueEndpointId")]
        public Output<string?> ServiceBusQueueEndpointId { get; private set; } = null!;

        /// <summary>
        /// Specifies the id where the Service Bus Topic is located.
        /// </summary>
        [Output("serviceBusTopicEndpointId")]
        public Output<string?> ServiceBusTopicEndpointId { get; private set; } = null!;

        /// <summary>
        /// A `storage_blob_dead_letter_destination` block as defined below.
        /// </summary>
        [Output("storageBlobDeadLetterDestination")]
        public Output<Outputs.EventSubscriptionStorageBlobDeadLetterDestination?> StorageBlobDeadLetterDestination { get; private set; } = null!;

        /// <summary>
        /// A `storage_queue_endpoint` block as defined below.
        /// </summary>
        [Output("storageQueueEndpoint")]
        public Output<Outputs.EventSubscriptionStorageQueueEndpoint?> StorageQueueEndpoint { get; private set; } = null!;

        /// <summary>
        /// A `subject_filter` block as defined below.
        /// </summary>
        [Output("subjectFilter")]
        public Output<Outputs.EventSubscriptionSubjectFilter?> SubjectFilter { get; private set; } = null!;

        /// <summary>
        /// (Optional/ **Deprecated) Specifies the name of the topic to associate with the event subscription.
        /// </summary>
        [Output("topicName")]
        public Output<string> TopicName { get; private set; } = null!;

        /// <summary>
        /// A `webhook_endpoint` block as defined below.
        /// </summary>
        [Output("webhookEndpoint")]
        public Output<Outputs.EventSubscriptionWebhookEndpoint?> WebhookEndpoint { get; private set; } = null!;


        /// <summary>
        /// Create a EventSubscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventSubscription(string name, EventSubscriptionArgs args, CustomResourceOptions? options = null)
            : base("azure:eventhub/eventSubscription:EventSubscription", name, args ?? new EventSubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventSubscription(string name, Input<string> id, EventSubscriptionState? state = null, CustomResourceOptions? options = null)
            : base("azure:eventhub/eventSubscription:EventSubscription", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EventSubscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventSubscription Get(string name, Input<string> id, EventSubscriptionState? state = null, CustomResourceOptions? options = null)
        {
            return new EventSubscription(name, id, state, options);
        }
    }

    public sealed class EventSubscriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `advanced_filter` block as defined below.
        /// </summary>
        [Input("advancedFilter")]
        public Input<Inputs.EventSubscriptionAdvancedFilterArgs>? AdvancedFilter { get; set; }

        /// <summary>
        /// An `azure_function_endpoint` block as defined below.
        /// </summary>
        [Input("azureFunctionEndpoint")]
        public Input<Inputs.EventSubscriptionAzureFunctionEndpointArgs>? AzureFunctionEndpoint { get; set; }

        /// <summary>
        /// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("eventDeliverySchema")]
        public Input<string>? EventDeliverySchema { get; set; }

        /// <summary>
        /// A `eventhub_endpoint` block as defined below.
        /// </summary>
        [Input("eventhubEndpoint")]
        public Input<Inputs.EventSubscriptionEventhubEndpointArgs>? EventhubEndpoint { get; set; }

        /// <summary>
        /// Specifies the id where the Event Hub is located.
        /// </summary>
        [Input("eventhubEndpointId")]
        public Input<string>? EventhubEndpointId { get; set; }

        /// <summary>
        /// Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
        /// </summary>
        [Input("expirationTimeUtc")]
        public Input<string>? ExpirationTimeUtc { get; set; }

        /// <summary>
        /// A `hybrid_connection_endpoint` block as defined below.
        /// </summary>
        [Input("hybridConnectionEndpoint")]
        public Input<Inputs.EventSubscriptionHybridConnectionEndpointArgs>? HybridConnectionEndpoint { get; set; }

        /// <summary>
        /// Specifies the id where the Hybrid Connection is located.
        /// </summary>
        [Input("hybridConnectionEndpointId")]
        public Input<string>? HybridConnectionEndpointId { get; set; }

        [Input("includedEventTypes")]
        private InputList<string>? _includedEventTypes;

        /// <summary>
        /// A list of applicable event types that need to be part of the event subscription.
        /// </summary>
        public InputList<string> IncludedEventTypes
        {
            get => _includedEventTypes ?? (_includedEventTypes = new InputList<string>());
            set => _includedEventTypes = value;
        }

        [Input("labels")]
        private InputList<string>? _labels;

        /// <summary>
        /// A list of labels to assign to the event subscription.
        /// </summary>
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `retry_policy` block as defined below.
        /// </summary>
        [Input("retryPolicy")]
        public Input<Inputs.EventSubscriptionRetryPolicyArgs>? RetryPolicy { get; set; }

        /// <summary>
        /// Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        /// <summary>
        /// Specifies the id where the Service Bus Queue is located.
        /// </summary>
        [Input("serviceBusQueueEndpointId")]
        public Input<string>? ServiceBusQueueEndpointId { get; set; }

        /// <summary>
        /// Specifies the id where the Service Bus Topic is located.
        /// </summary>
        [Input("serviceBusTopicEndpointId")]
        public Input<string>? ServiceBusTopicEndpointId { get; set; }

        /// <summary>
        /// A `storage_blob_dead_letter_destination` block as defined below.
        /// </summary>
        [Input("storageBlobDeadLetterDestination")]
        public Input<Inputs.EventSubscriptionStorageBlobDeadLetterDestinationArgs>? StorageBlobDeadLetterDestination { get; set; }

        /// <summary>
        /// A `storage_queue_endpoint` block as defined below.
        /// </summary>
        [Input("storageQueueEndpoint")]
        public Input<Inputs.EventSubscriptionStorageQueueEndpointArgs>? StorageQueueEndpoint { get; set; }

        /// <summary>
        /// A `subject_filter` block as defined below.
        /// </summary>
        [Input("subjectFilter")]
        public Input<Inputs.EventSubscriptionSubjectFilterArgs>? SubjectFilter { get; set; }

        /// <summary>
        /// (Optional/ **Deprecated) Specifies the name of the topic to associate with the event subscription.
        /// </summary>
        [Input("topicName")]
        public Input<string>? TopicName { get; set; }

        /// <summary>
        /// A `webhook_endpoint` block as defined below.
        /// </summary>
        [Input("webhookEndpoint")]
        public Input<Inputs.EventSubscriptionWebhookEndpointArgs>? WebhookEndpoint { get; set; }

        public EventSubscriptionArgs()
        {
        }
    }

    public sealed class EventSubscriptionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `advanced_filter` block as defined below.
        /// </summary>
        [Input("advancedFilter")]
        public Input<Inputs.EventSubscriptionAdvancedFilterGetArgs>? AdvancedFilter { get; set; }

        /// <summary>
        /// An `azure_function_endpoint` block as defined below.
        /// </summary>
        [Input("azureFunctionEndpoint")]
        public Input<Inputs.EventSubscriptionAzureFunctionEndpointGetArgs>? AzureFunctionEndpoint { get; set; }

        /// <summary>
        /// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("eventDeliverySchema")]
        public Input<string>? EventDeliverySchema { get; set; }

        /// <summary>
        /// A `eventhub_endpoint` block as defined below.
        /// </summary>
        [Input("eventhubEndpoint")]
        public Input<Inputs.EventSubscriptionEventhubEndpointGetArgs>? EventhubEndpoint { get; set; }

        /// <summary>
        /// Specifies the id where the Event Hub is located.
        /// </summary>
        [Input("eventhubEndpointId")]
        public Input<string>? EventhubEndpointId { get; set; }

        /// <summary>
        /// Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
        /// </summary>
        [Input("expirationTimeUtc")]
        public Input<string>? ExpirationTimeUtc { get; set; }

        /// <summary>
        /// A `hybrid_connection_endpoint` block as defined below.
        /// </summary>
        [Input("hybridConnectionEndpoint")]
        public Input<Inputs.EventSubscriptionHybridConnectionEndpointGetArgs>? HybridConnectionEndpoint { get; set; }

        /// <summary>
        /// Specifies the id where the Hybrid Connection is located.
        /// </summary>
        [Input("hybridConnectionEndpointId")]
        public Input<string>? HybridConnectionEndpointId { get; set; }

        [Input("includedEventTypes")]
        private InputList<string>? _includedEventTypes;

        /// <summary>
        /// A list of applicable event types that need to be part of the event subscription.
        /// </summary>
        public InputList<string> IncludedEventTypes
        {
            get => _includedEventTypes ?? (_includedEventTypes = new InputList<string>());
            set => _includedEventTypes = value;
        }

        [Input("labels")]
        private InputList<string>? _labels;

        /// <summary>
        /// A list of labels to assign to the event subscription.
        /// </summary>
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `retry_policy` block as defined below.
        /// </summary>
        [Input("retryPolicy")]
        public Input<Inputs.EventSubscriptionRetryPolicyGetArgs>? RetryPolicy { get; set; }

        /// <summary>
        /// Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// Specifies the id where the Service Bus Queue is located.
        /// </summary>
        [Input("serviceBusQueueEndpointId")]
        public Input<string>? ServiceBusQueueEndpointId { get; set; }

        /// <summary>
        /// Specifies the id where the Service Bus Topic is located.
        /// </summary>
        [Input("serviceBusTopicEndpointId")]
        public Input<string>? ServiceBusTopicEndpointId { get; set; }

        /// <summary>
        /// A `storage_blob_dead_letter_destination` block as defined below.
        /// </summary>
        [Input("storageBlobDeadLetterDestination")]
        public Input<Inputs.EventSubscriptionStorageBlobDeadLetterDestinationGetArgs>? StorageBlobDeadLetterDestination { get; set; }

        /// <summary>
        /// A `storage_queue_endpoint` block as defined below.
        /// </summary>
        [Input("storageQueueEndpoint")]
        public Input<Inputs.EventSubscriptionStorageQueueEndpointGetArgs>? StorageQueueEndpoint { get; set; }

        /// <summary>
        /// A `subject_filter` block as defined below.
        /// </summary>
        [Input("subjectFilter")]
        public Input<Inputs.EventSubscriptionSubjectFilterGetArgs>? SubjectFilter { get; set; }

        /// <summary>
        /// (Optional/ **Deprecated) Specifies the name of the topic to associate with the event subscription.
        /// </summary>
        [Input("topicName")]
        public Input<string>? TopicName { get; set; }

        /// <summary>
        /// A `webhook_endpoint` block as defined below.
        /// </summary>
        [Input("webhookEndpoint")]
        public Input<Inputs.EventSubscriptionWebhookEndpointGetArgs>? WebhookEndpoint { get; set; }

        public EventSubscriptionState()
        {
        }
    }
}

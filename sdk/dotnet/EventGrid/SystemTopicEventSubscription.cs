// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventGrid
{
    /// <summary>
    /// Manages an EventGrid System Topic Event Subscription.
    /// 
    /// ## Import
    /// 
    /// EventGrid System Topic Event Subscriptions can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:eventgrid/systemTopicEventSubscription:SystemTopicEventSubscription example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/systemTopics/topic1/eventSubscriptions/subscription1
    /// ```
    /// </summary>
    [AzureResourceType("azure:eventgrid/systemTopicEventSubscription:SystemTopicEventSubscription")]
    public partial class SystemTopicEventSubscription : Pulumi.CustomResource
    {
        /// <summary>
        /// A `advanced_filter` block as defined below.
        /// </summary>
        [Output("advancedFilter")]
        public Output<Outputs.SystemTopicEventSubscriptionAdvancedFilter?> AdvancedFilter { get; private set; } = null!;

        /// <summary>
        /// An `azure_function_endpoint` block as defined below.
        /// </summary>
        [Output("azureFunctionEndpoint")]
        public Output<Outputs.SystemTopicEventSubscriptionAzureFunctionEndpoint?> AzureFunctionEndpoint { get; private set; } = null!;

        /// <summary>
        /// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("eventDeliverySchema")]
        public Output<string?> EventDeliverySchema { get; private set; } = null!;

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
        /// The name which should be used for this Event Subscription. Changing this forces a new Event Subscription to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the System Topic exists. Changing this forces a new Event Subscription to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A `retry_policy` block as defined below.
        /// </summary>
        [Output("retryPolicy")]
        public Output<Outputs.SystemTopicEventSubscriptionRetryPolicy> RetryPolicy { get; private set; } = null!;

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
        public Output<Outputs.SystemTopicEventSubscriptionStorageBlobDeadLetterDestination?> StorageBlobDeadLetterDestination { get; private set; } = null!;

        /// <summary>
        /// A `storage_queue_endpoint` block as defined below.
        /// </summary>
        [Output("storageQueueEndpoint")]
        public Output<Outputs.SystemTopicEventSubscriptionStorageQueueEndpoint?> StorageQueueEndpoint { get; private set; } = null!;

        /// <summary>
        /// A `subject_filter` block as defined below.
        /// </summary>
        [Output("subjectFilter")]
        public Output<Outputs.SystemTopicEventSubscriptionSubjectFilter?> SubjectFilter { get; private set; } = null!;

        /// <summary>
        /// The System Topic where the Event Subscription should be created in. Changing this forces a new Event Subscription to be created.
        /// </summary>
        [Output("systemTopic")]
        public Output<string> SystemTopic { get; private set; } = null!;

        /// <summary>
        /// A `webhook_endpoint` block as defined below.
        /// </summary>
        [Output("webhookEndpoint")]
        public Output<Outputs.SystemTopicEventSubscriptionWebhookEndpoint?> WebhookEndpoint { get; private set; } = null!;


        /// <summary>
        /// Create a SystemTopicEventSubscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemTopicEventSubscription(string name, SystemTopicEventSubscriptionArgs args, CustomResourceOptions? options = null)
            : base("azure:eventgrid/systemTopicEventSubscription:SystemTopicEventSubscription", name, args ?? new SystemTopicEventSubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemTopicEventSubscription(string name, Input<string> id, SystemTopicEventSubscriptionState? state = null, CustomResourceOptions? options = null)
            : base("azure:eventgrid/systemTopicEventSubscription:SystemTopicEventSubscription", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemTopicEventSubscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemTopicEventSubscription Get(string name, Input<string> id, SystemTopicEventSubscriptionState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemTopicEventSubscription(name, id, state, options);
        }
    }

    public sealed class SystemTopicEventSubscriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `advanced_filter` block as defined below.
        /// </summary>
        [Input("advancedFilter")]
        public Input<Inputs.SystemTopicEventSubscriptionAdvancedFilterArgs>? AdvancedFilter { get; set; }

        /// <summary>
        /// An `azure_function_endpoint` block as defined below.
        /// </summary>
        [Input("azureFunctionEndpoint")]
        public Input<Inputs.SystemTopicEventSubscriptionAzureFunctionEndpointArgs>? AzureFunctionEndpoint { get; set; }

        /// <summary>
        /// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("eventDeliverySchema")]
        public Input<string>? EventDeliverySchema { get; set; }

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
        /// The name which should be used for this Event Subscription. Changing this forces a new Event Subscription to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group where the System Topic exists. Changing this forces a new Event Subscription to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `retry_policy` block as defined below.
        /// </summary>
        [Input("retryPolicy")]
        public Input<Inputs.SystemTopicEventSubscriptionRetryPolicyArgs>? RetryPolicy { get; set; }

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
        public Input<Inputs.SystemTopicEventSubscriptionStorageBlobDeadLetterDestinationArgs>? StorageBlobDeadLetterDestination { get; set; }

        /// <summary>
        /// A `storage_queue_endpoint` block as defined below.
        /// </summary>
        [Input("storageQueueEndpoint")]
        public Input<Inputs.SystemTopicEventSubscriptionStorageQueueEndpointArgs>? StorageQueueEndpoint { get; set; }

        /// <summary>
        /// A `subject_filter` block as defined below.
        /// </summary>
        [Input("subjectFilter")]
        public Input<Inputs.SystemTopicEventSubscriptionSubjectFilterArgs>? SubjectFilter { get; set; }

        /// <summary>
        /// The System Topic where the Event Subscription should be created in. Changing this forces a new Event Subscription to be created.
        /// </summary>
        [Input("systemTopic", required: true)]
        public Input<string> SystemTopic { get; set; } = null!;

        /// <summary>
        /// A `webhook_endpoint` block as defined below.
        /// </summary>
        [Input("webhookEndpoint")]
        public Input<Inputs.SystemTopicEventSubscriptionWebhookEndpointArgs>? WebhookEndpoint { get; set; }

        public SystemTopicEventSubscriptionArgs()
        {
        }
    }

    public sealed class SystemTopicEventSubscriptionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `advanced_filter` block as defined below.
        /// </summary>
        [Input("advancedFilter")]
        public Input<Inputs.SystemTopicEventSubscriptionAdvancedFilterGetArgs>? AdvancedFilter { get; set; }

        /// <summary>
        /// An `azure_function_endpoint` block as defined below.
        /// </summary>
        [Input("azureFunctionEndpoint")]
        public Input<Inputs.SystemTopicEventSubscriptionAzureFunctionEndpointGetArgs>? AzureFunctionEndpoint { get; set; }

        /// <summary>
        /// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("eventDeliverySchema")]
        public Input<string>? EventDeliverySchema { get; set; }

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
        /// The name which should be used for this Event Subscription. Changing this forces a new Event Subscription to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group where the System Topic exists. Changing this forces a new Event Subscription to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// A `retry_policy` block as defined below.
        /// </summary>
        [Input("retryPolicy")]
        public Input<Inputs.SystemTopicEventSubscriptionRetryPolicyGetArgs>? RetryPolicy { get; set; }

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
        public Input<Inputs.SystemTopicEventSubscriptionStorageBlobDeadLetterDestinationGetArgs>? StorageBlobDeadLetterDestination { get; set; }

        /// <summary>
        /// A `storage_queue_endpoint` block as defined below.
        /// </summary>
        [Input("storageQueueEndpoint")]
        public Input<Inputs.SystemTopicEventSubscriptionStorageQueueEndpointGetArgs>? StorageQueueEndpoint { get; set; }

        /// <summary>
        /// A `subject_filter` block as defined below.
        /// </summary>
        [Input("subjectFilter")]
        public Input<Inputs.SystemTopicEventSubscriptionSubjectFilterGetArgs>? SubjectFilter { get; set; }

        /// <summary>
        /// The System Topic where the Event Subscription should be created in. Changing this forces a new Event Subscription to be created.
        /// </summary>
        [Input("systemTopic")]
        public Input<string>? SystemTopic { get; set; }

        /// <summary>
        /// A `webhook_endpoint` block as defined below.
        /// </summary>
        [Input("webhookEndpoint")]
        public Input<Inputs.SystemTopicEventSubscriptionWebhookEndpointGetArgs>? WebhookEndpoint { get; set; }

        public SystemTopicEventSubscriptionState()
        {
        }
    }
}

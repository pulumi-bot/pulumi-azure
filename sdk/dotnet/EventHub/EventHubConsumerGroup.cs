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
    /// Manages a Event Hubs Consumer Group as a nested resource within an Event Hub.
    /// </summary>
    [Obsolete(@"azure.eventhub.EventHubConsumerGroup has been deprecated in favor of azure.eventhub.ConsumerGroup")]
    public partial class EventHubConsumerGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the name of the EventHub. Changing this forces a new resource to be created.
        /// </summary>
        [Output("eventhubName")]
        public Output<string> EventhubName { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the EventHub Consumer Group resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Output("namespaceName")]
        public Output<string> NamespaceName { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Specifies the user metadata.
        /// </summary>
        [Output("userMetadata")]
        public Output<string?> UserMetadata { get; private set; } = null!;


        /// <summary>
        /// Create a EventHubConsumerGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventHubConsumerGroup(string name, EventHubConsumerGroupArgs args, CustomResourceOptions? options = null)
            : base("azure:eventhub/eventHubConsumerGroup:EventHubConsumerGroup", name, args ?? new EventHubConsumerGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventHubConsumerGroup(string name, Input<string> id, EventHubConsumerGroupState? state = null, CustomResourceOptions? options = null)
            : base("azure:eventhub/eventHubConsumerGroup:EventHubConsumerGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EventHubConsumerGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventHubConsumerGroup Get(string name, Input<string> id, EventHubConsumerGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new EventHubConsumerGroup(name, id, state, options);
        }
    }

    public sealed class EventHubConsumerGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the EventHub. Changing this forces a new resource to be created.
        /// </summary>
        [Input("eventhubName", required: true)]
        public Input<string> EventhubName { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the EventHub Consumer Group resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies the user metadata.
        /// </summary>
        [Input("userMetadata")]
        public Input<string>? UserMetadata { get; set; }

        public EventHubConsumerGroupArgs()
        {
        }
    }

    public sealed class EventHubConsumerGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the EventHub. Changing this forces a new resource to be created.
        /// </summary>
        [Input("eventhubName")]
        public Input<string>? EventhubName { get; set; }

        /// <summary>
        /// Specifies the name of the EventHub Consumer Group resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("namespaceName")]
        public Input<string>? NamespaceName { get; set; }

        /// <summary>
        /// The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Specifies the user metadata.
        /// </summary>
        [Input("userMetadata")]
        public Input<string>? UserMetadata { get; set; }

        public EventHubConsumerGroupState()
        {
        }
    }
}

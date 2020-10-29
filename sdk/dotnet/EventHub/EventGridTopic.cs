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
    /// Manages an EventGrid Topic
    /// 
    /// &gt; **Note:** at this time EventGrid Topic's are only available in a limited number of regions.
    /// </summary>
    [Obsolete(@"azure.eventhub.EventGridTopic has been deprecated in favor of azure.eventgrid.Topic")]
    public partial class EventGridTopic : Pulumi.CustomResource
    {
        /// <summary>
        /// The Endpoint associated with the EventGrid Topic.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// A `input_mapping_default_values` block as defined below.
        /// </summary>
        [Output("inputMappingDefaultValues")]
        public Output<Outputs.EventGridTopicInputMappingDefaultValues?> InputMappingDefaultValues { get; private set; } = null!;

        /// <summary>
        /// A `input_mapping_fields` block as defined below.
        /// </summary>
        [Output("inputMappingFields")]
        public Output<Outputs.EventGridTopicInputMappingFields?> InputMappingFields { get; private set; } = null!;

        /// <summary>
        /// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("inputSchema")]
        public Output<string?> InputSchema { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Primary Shared Access Key associated with the EventGrid Topic.
        /// </summary>
        [Output("primaryAccessKey")]
        public Output<string> PrimaryAccessKey { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The Secondary Shared Access Key associated with the EventGrid Topic.
        /// </summary>
        [Output("secondaryAccessKey")]
        public Output<string> SecondaryAccessKey { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a EventGridTopic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventGridTopic(string name, EventGridTopicArgs args, CustomResourceOptions? options = null)
            : base("azure:eventhub/eventGridTopic:EventGridTopic", name, args ?? new EventGridTopicArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventGridTopic(string name, Input<string> id, EventGridTopicState? state = null, CustomResourceOptions? options = null)
            : base("azure:eventhub/eventGridTopic:EventGridTopic", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EventGridTopic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventGridTopic Get(string name, Input<string> id, EventGridTopicState? state = null, CustomResourceOptions? options = null)
        {
            return new EventGridTopic(name, id, state, options);
        }
    }

    public sealed class EventGridTopicArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `input_mapping_default_values` block as defined below.
        /// </summary>
        [Input("inputMappingDefaultValues")]
        public Input<Inputs.EventGridTopicInputMappingDefaultValuesArgs>? InputMappingDefaultValues { get; set; }

        /// <summary>
        /// A `input_mapping_fields` block as defined below.
        /// </summary>
        [Input("inputMappingFields")]
        public Input<Inputs.EventGridTopicInputMappingFieldsArgs>? InputMappingFields { get; set; }

        /// <summary>
        /// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("inputSchema")]
        public Input<string>? InputSchema { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public EventGridTopicArgs()
        {
        }
    }

    public sealed class EventGridTopicState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Endpoint associated with the EventGrid Topic.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// A `input_mapping_default_values` block as defined below.
        /// </summary>
        [Input("inputMappingDefaultValues")]
        public Input<Inputs.EventGridTopicInputMappingDefaultValuesGetArgs>? InputMappingDefaultValues { get; set; }

        /// <summary>
        /// A `input_mapping_fields` block as defined below.
        /// </summary>
        [Input("inputMappingFields")]
        public Input<Inputs.EventGridTopicInputMappingFieldsGetArgs>? InputMappingFields { get; set; }

        /// <summary>
        /// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("inputSchema")]
        public Input<string>? InputSchema { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Primary Shared Access Key associated with the EventGrid Topic.
        /// </summary>
        [Input("primaryAccessKey")]
        public Input<string>? PrimaryAccessKey { get; set; }

        /// <summary>
        /// The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The Secondary Shared Access Key associated with the EventGrid Topic.
        /// </summary>
        [Input("secondaryAccessKey")]
        public Input<string>? SecondaryAccessKey { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public EventGridTopicState()
        {
        }
    }
}

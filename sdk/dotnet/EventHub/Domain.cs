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
    /// Manages an EventGrid Domain
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
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West US 2",
    ///         });
    ///         var exampleDomain = new Azure.EventGrid.Domain("exampleDomain", new Azure.EventGrid.DomainArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Tags = 
    ///             {
    ///                 { "environment", "Production" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// EventGrid Domains can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:eventhub/domain:Domain domain1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/domains/domain1
    /// ```
    /// </summary>
    [Obsolete(@"azure.eventhub.Domain has been deprecated in favor of azure.eventgrid.Domain")]
    [AzureResourceType("azure:eventhub/domain:Domain")]
    public partial class Domain : Pulumi.CustomResource
    {
        /// <summary>
        /// The Endpoint associated with the EventGrid Domain.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// One or more `inbound_ip_rule` blocks as defined below.
        /// </summary>
        [Output("inboundIpRules")]
        public Output<ImmutableArray<Outputs.DomainInboundIpRule>> InboundIpRules { get; private set; } = null!;

        /// <summary>
        /// A `input_mapping_default_values` block as defined below.
        /// </summary>
        [Output("inputMappingDefaultValues")]
        public Output<Outputs.DomainInputMappingDefaultValues?> InputMappingDefaultValues { get; private set; } = null!;

        /// <summary>
        /// A `input_mapping_fields` block as defined below.
        /// </summary>
        [Output("inputMappingFields")]
        public Output<Outputs.DomainInputMappingFields?> InputMappingFields { get; private set; } = null!;

        /// <summary>
        /// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("inputSchema")]
        public Output<string?> InputSchema { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Primary Shared Access Key associated with the EventGrid Domain.
        /// </summary>
        [Output("primaryAccessKey")]
        public Output<string> PrimaryAccessKey { get; private set; } = null!;

        /// <summary>
        /// Whether or not public network access is allowed for this server. Defaults to `true`.
        /// </summary>
        [Output("publicNetworkAccessEnabled")]
        public Output<bool?> PublicNetworkAccessEnabled { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The Secondary Shared Access Key associated with the EventGrid Domain.
        /// </summary>
        [Output("secondaryAccessKey")]
        public Output<string> SecondaryAccessKey { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Domain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Domain(string name, DomainArgs args, CustomResourceOptions? options = null)
            : base("azure:eventhub/domain:Domain", name, args ?? new DomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Domain(string name, Input<string> id, DomainState? state = null, CustomResourceOptions? options = null)
            : base("azure:eventhub/domain:Domain", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Domain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Domain Get(string name, Input<string> id, DomainState? state = null, CustomResourceOptions? options = null)
        {
            return new Domain(name, id, state, options);
        }
    }

    public sealed class DomainArgs : Pulumi.ResourceArgs
    {
        [Input("inboundIpRules")]
        private InputList<Inputs.DomainInboundIpRuleArgs>? _inboundIpRules;

        /// <summary>
        /// One or more `inbound_ip_rule` blocks as defined below.
        /// </summary>
        public InputList<Inputs.DomainInboundIpRuleArgs> InboundIpRules
        {
            get => _inboundIpRules ?? (_inboundIpRules = new InputList<Inputs.DomainInboundIpRuleArgs>());
            set => _inboundIpRules = value;
        }

        /// <summary>
        /// A `input_mapping_default_values` block as defined below.
        /// </summary>
        [Input("inputMappingDefaultValues")]
        public Input<Inputs.DomainInputMappingDefaultValuesArgs>? InputMappingDefaultValues { get; set; }

        /// <summary>
        /// A `input_mapping_fields` block as defined below.
        /// </summary>
        [Input("inputMappingFields")]
        public Input<Inputs.DomainInputMappingFieldsArgs>? InputMappingFields { get; set; }

        /// <summary>
        /// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("inputSchema")]
        public Input<string>? InputSchema { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether or not public network access is allowed for this server. Defaults to `true`.
        /// </summary>
        [Input("publicNetworkAccessEnabled")]
        public Input<bool>? PublicNetworkAccessEnabled { get; set; }

        /// <summary>
        /// The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
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

        public DomainArgs()
        {
        }
    }

    public sealed class DomainState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Endpoint associated with the EventGrid Domain.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        [Input("inboundIpRules")]
        private InputList<Inputs.DomainInboundIpRuleGetArgs>? _inboundIpRules;

        /// <summary>
        /// One or more `inbound_ip_rule` blocks as defined below.
        /// </summary>
        public InputList<Inputs.DomainInboundIpRuleGetArgs> InboundIpRules
        {
            get => _inboundIpRules ?? (_inboundIpRules = new InputList<Inputs.DomainInboundIpRuleGetArgs>());
            set => _inboundIpRules = value;
        }

        /// <summary>
        /// A `input_mapping_default_values` block as defined below.
        /// </summary>
        [Input("inputMappingDefaultValues")]
        public Input<Inputs.DomainInputMappingDefaultValuesGetArgs>? InputMappingDefaultValues { get; set; }

        /// <summary>
        /// A `input_mapping_fields` block as defined below.
        /// </summary>
        [Input("inputMappingFields")]
        public Input<Inputs.DomainInputMappingFieldsGetArgs>? InputMappingFields { get; set; }

        /// <summary>
        /// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("inputSchema")]
        public Input<string>? InputSchema { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Primary Shared Access Key associated with the EventGrid Domain.
        /// </summary>
        [Input("primaryAccessKey")]
        public Input<string>? PrimaryAccessKey { get; set; }

        /// <summary>
        /// Whether or not public network access is allowed for this server. Defaults to `true`.
        /// </summary>
        [Input("publicNetworkAccessEnabled")]
        public Input<bool>? PublicNetworkAccessEnabled { get; set; }

        /// <summary>
        /// The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The Secondary Shared Access Key associated with the EventGrid Domain.
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

        public DomainState()
        {
        }
    }
}

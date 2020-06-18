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
    /// Manages an EventHub Namespace.
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
    ///             Location = "West Europe",
    ///         });
    ///         var exampleEventHubNamespace = new Azure.EventHub.EventHubNamespace("exampleEventHubNamespace", new Azure.EventHub.EventHubNamespaceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Sku = "Standard",
    ///             Capacity = 2,
    ///             Tags = 
    ///             {
    ///                 { "environment", "Production" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class EventHubNamespace : Pulumi.CustomResource
    {
        /// <summary>
        /// Is Auto Inflate enabled for the EventHub Namespace?
        /// </summary>
        [Output("autoInflateEnabled")]
        public Output<bool?> AutoInflateEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the Capacity / Throughput Units for a `Standard` SKU namespace. Valid values range from `1` - `20`.
        /// </summary>
        [Output("capacity")]
        public Output<int?> Capacity { get; private set; } = null!;

        /// <summary>
        /// The primary connection string for the authorization
        /// rule `RootManageSharedAccessKey`.
        /// </summary>
        [Output("defaultPrimaryConnectionString")]
        public Output<string> DefaultPrimaryConnectionString { get; private set; } = null!;

        /// <summary>
        /// The alias of the primary connection string for the authorization
        /// rule `RootManageSharedAccessKey`, which is generated when disaster recovery is enabled.
        /// </summary>
        [Output("defaultPrimaryConnectionStringAlias")]
        public Output<string> DefaultPrimaryConnectionStringAlias { get; private set; } = null!;

        /// <summary>
        /// The primary access key for the authorization rule `RootManageSharedAccessKey`.
        /// </summary>
        [Output("defaultPrimaryKey")]
        public Output<string> DefaultPrimaryKey { get; private set; } = null!;

        /// <summary>
        /// The secondary connection string for the
        /// authorization rule `RootManageSharedAccessKey`.
        /// </summary>
        [Output("defaultSecondaryConnectionString")]
        public Output<string> DefaultSecondaryConnectionString { get; private set; } = null!;

        /// <summary>
        /// The alias of the secondary connection string for the
        /// authorization rule `RootManageSharedAccessKey`, which is generated when disaster recovery is enabled.
        /// </summary>
        [Output("defaultSecondaryConnectionStringAlias")]
        public Output<string> DefaultSecondaryConnectionStringAlias { get; private set; } = null!;

        /// <summary>
        /// The secondary access key for the authorization rule `RootManageSharedAccessKey`.
        /// </summary>
        [Output("defaultSecondaryKey")]
        public Output<string> DefaultSecondaryKey { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the maximum number of throughput units when Auto Inflate is Enabled. Valid values range from `1` - `20`.
        /// </summary>
        [Output("maximumThroughputUnits")]
        public Output<int> MaximumThroughputUnits { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A `network_rulesets` block as defined below.
        /// </summary>
        [Output("networkRulesets")]
        public Output<Outputs.EventHubNamespaceNetworkRulesets> NetworkRulesets { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Defines which tier to use. Valid options are `Basic` and `Standard`.
        /// </summary>
        [Output("sku")]
        public Output<string> Sku { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a EventHubNamespace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventHubNamespace(string name, EventHubNamespaceArgs args, CustomResourceOptions? options = null)
            : base("azure:eventhub/eventHubNamespace:EventHubNamespace", name, args ?? new EventHubNamespaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventHubNamespace(string name, Input<string> id, EventHubNamespaceState? state = null, CustomResourceOptions? options = null)
            : base("azure:eventhub/eventHubNamespace:EventHubNamespace", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EventHubNamespace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventHubNamespace Get(string name, Input<string> id, EventHubNamespaceState? state = null, CustomResourceOptions? options = null)
        {
            return new EventHubNamespace(name, id, state, options);
        }
    }

    public sealed class EventHubNamespaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is Auto Inflate enabled for the EventHub Namespace?
        /// </summary>
        [Input("autoInflateEnabled")]
        public Input<bool>? AutoInflateEnabled { get; set; }

        /// <summary>
        /// Specifies the Capacity / Throughput Units for a `Standard` SKU namespace. Valid values range from `1` - `20`.
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the maximum number of throughput units when Auto Inflate is Enabled. Valid values range from `1` - `20`.
        /// </summary>
        [Input("maximumThroughputUnits")]
        public Input<int>? MaximumThroughputUnits { get; set; }

        /// <summary>
        /// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `network_rulesets` block as defined below.
        /// </summary>
        [Input("networkRulesets")]
        public Input<Inputs.EventHubNamespaceNetworkRulesetsArgs>? NetworkRulesets { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Defines which tier to use. Valid options are `Basic` and `Standard`.
        /// </summary>
        [Input("sku", required: true)]
        public Input<string> Sku { get; set; } = null!;

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

        public EventHubNamespaceArgs()
        {
        }
    }

    public sealed class EventHubNamespaceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is Auto Inflate enabled for the EventHub Namespace?
        /// </summary>
        [Input("autoInflateEnabled")]
        public Input<bool>? AutoInflateEnabled { get; set; }

        /// <summary>
        /// Specifies the Capacity / Throughput Units for a `Standard` SKU namespace. Valid values range from `1` - `20`.
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        /// <summary>
        /// The primary connection string for the authorization
        /// rule `RootManageSharedAccessKey`.
        /// </summary>
        [Input("defaultPrimaryConnectionString")]
        public Input<string>? DefaultPrimaryConnectionString { get; set; }

        /// <summary>
        /// The alias of the primary connection string for the authorization
        /// rule `RootManageSharedAccessKey`, which is generated when disaster recovery is enabled.
        /// </summary>
        [Input("defaultPrimaryConnectionStringAlias")]
        public Input<string>? DefaultPrimaryConnectionStringAlias { get; set; }

        /// <summary>
        /// The primary access key for the authorization rule `RootManageSharedAccessKey`.
        /// </summary>
        [Input("defaultPrimaryKey")]
        public Input<string>? DefaultPrimaryKey { get; set; }

        /// <summary>
        /// The secondary connection string for the
        /// authorization rule `RootManageSharedAccessKey`.
        /// </summary>
        [Input("defaultSecondaryConnectionString")]
        public Input<string>? DefaultSecondaryConnectionString { get; set; }

        /// <summary>
        /// The alias of the secondary connection string for the
        /// authorization rule `RootManageSharedAccessKey`, which is generated when disaster recovery is enabled.
        /// </summary>
        [Input("defaultSecondaryConnectionStringAlias")]
        public Input<string>? DefaultSecondaryConnectionStringAlias { get; set; }

        /// <summary>
        /// The secondary access key for the authorization rule `RootManageSharedAccessKey`.
        /// </summary>
        [Input("defaultSecondaryKey")]
        public Input<string>? DefaultSecondaryKey { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the maximum number of throughput units when Auto Inflate is Enabled. Valid values range from `1` - `20`.
        /// </summary>
        [Input("maximumThroughputUnits")]
        public Input<int>? MaximumThroughputUnits { get; set; }

        /// <summary>
        /// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `network_rulesets` block as defined below.
        /// </summary>
        [Input("networkRulesets")]
        public Input<Inputs.EventHubNamespaceNetworkRulesetsGetArgs>? NetworkRulesets { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Defines which tier to use. Valid options are `Basic` and `Standard`.
        /// </summary>
        [Input("sku")]
        public Input<string>? Sku { get; set; }

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

        public EventHubNamespaceState()
        {
        }
    }
}

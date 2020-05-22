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
    /// Manages a Event Hubs authorization Rule within an Event Hub.
    /// 
    /// ## Example Usage
    /// 
    /// 
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
    ///             Location = "West US",
    ///         });
    ///         var exampleEventHubNamespace = new Azure.EventHub.EventHubNamespace("exampleEventHubNamespace", new Azure.EventHub.EventHubNamespaceArgs
    ///         {
    ///             Location = "West US",
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Sku = "Basic",
    ///             Capacity = 2,
    ///             Tags = 
    ///             {
    ///                 { "environment", "Production" },
    ///             },
    ///         });
    ///         var exampleEventHub = new Azure.EventHub.EventHub("exampleEventHub", new Azure.EventHub.EventHubArgs
    ///         {
    ///             NamespaceName = exampleEventHubNamespace.Name,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             PartitionCount = 2,
    ///             MessageRetention = 2,
    ///         });
    ///         var exampleAuthorizationRule = new Azure.EventHub.AuthorizationRule("exampleAuthorizationRule", new Azure.EventHub.AuthorizationRuleArgs
    ///         {
    ///             NamespaceName = exampleEventHubNamespace.Name,
    ///             EventhubName = exampleEventHub.Name,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Listen = true,
    ///             Send = false,
    ///             Manage = false,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [Obsolete(@"azure.eventhub.EventHubAuthorizationRule has been deprecated in favor of azure.eventhub.AuthorizationRule")]
    public partial class EventHubAuthorizationRule : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the name of the EventHub. Changing this forces a new resource to be created.
        /// </summary>
        [Output("eventhubName")]
        public Output<string> EventhubName { get; private set; } = null!;

        /// <summary>
        /// Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
        /// </summary>
        [Output("listen")]
        public Output<bool?> Listen { get; private set; } = null!;

        /// <summary>
        /// Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
        /// </summary>
        [Output("manage")]
        public Output<bool?> Manage { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Output("namespaceName")]
        public Output<string> NamespaceName { get; private set; } = null!;

        /// <summary>
        /// The Primary Connection String for the Event Hubs authorization Rule.
        /// </summary>
        [Output("primaryConnectionString")]
        public Output<string> PrimaryConnectionString { get; private set; } = null!;

        /// <summary>
        /// The alias of the Primary Connection String for the Event Hubs authorization Rule, which is generated when disaster recovery is enabled.
        /// </summary>
        [Output("primaryConnectionStringAlias")]
        public Output<string> PrimaryConnectionStringAlias { get; private set; } = null!;

        /// <summary>
        /// The Primary Key for the Event Hubs authorization Rule.
        /// </summary>
        [Output("primaryKey")]
        public Output<string> PrimaryKey { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The Secondary Connection String for the Event Hubs Authorization Rule.
        /// </summary>
        [Output("secondaryConnectionString")]
        public Output<string> SecondaryConnectionString { get; private set; } = null!;

        /// <summary>
        /// The alias of the Secondary Connection String for the Event Hubs Authorization Rule, which is generated when disaster recovery is enabled.
        /// </summary>
        [Output("secondaryConnectionStringAlias")]
        public Output<string> SecondaryConnectionStringAlias { get; private set; } = null!;

        /// <summary>
        /// The Secondary Key for the Event Hubs Authorization Rule.
        /// </summary>
        [Output("secondaryKey")]
        public Output<string> SecondaryKey { get; private set; } = null!;

        /// <summary>
        /// Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
        /// </summary>
        [Output("send")]
        public Output<bool?> Send { get; private set; } = null!;


        /// <summary>
        /// Create a EventHubAuthorizationRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventHubAuthorizationRule(string name, EventHubAuthorizationRuleArgs args, CustomResourceOptions? options = null)
            : base("azure:eventhub/eventHubAuthorizationRule:EventHubAuthorizationRule", name, args ?? new EventHubAuthorizationRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventHubAuthorizationRule(string name, Input<string> id, EventHubAuthorizationRuleState? state = null, CustomResourceOptions? options = null)
            : base("azure:eventhub/eventHubAuthorizationRule:EventHubAuthorizationRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EventHubAuthorizationRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventHubAuthorizationRule Get(string name, Input<string> id, EventHubAuthorizationRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new EventHubAuthorizationRule(name, id, state, options);
        }
    }

    public sealed class EventHubAuthorizationRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the EventHub. Changing this forces a new resource to be created.
        /// </summary>
        [Input("eventhubName", required: true)]
        public Input<string> EventhubName { get; set; } = null!;

        /// <summary>
        /// Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
        /// </summary>
        [Input("listen")]
        public Input<bool>? Listen { get; set; }

        /// <summary>
        /// Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
        /// </summary>
        [Input("manage")]
        public Input<bool>? Manage { get; set; }

        /// <summary>
        /// Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
        /// </summary>
        [Input("send")]
        public Input<bool>? Send { get; set; }

        public EventHubAuthorizationRuleArgs()
        {
        }
    }

    public sealed class EventHubAuthorizationRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the EventHub. Changing this forces a new resource to be created.
        /// </summary>
        [Input("eventhubName")]
        public Input<string>? EventhubName { get; set; }

        /// <summary>
        /// Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
        /// </summary>
        [Input("listen")]
        public Input<bool>? Listen { get; set; }

        /// <summary>
        /// Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
        /// </summary>
        [Input("manage")]
        public Input<bool>? Manage { get; set; }

        /// <summary>
        /// Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("namespaceName")]
        public Input<string>? NamespaceName { get; set; }

        /// <summary>
        /// The Primary Connection String for the Event Hubs authorization Rule.
        /// </summary>
        [Input("primaryConnectionString")]
        public Input<string>? PrimaryConnectionString { get; set; }

        /// <summary>
        /// The alias of the Primary Connection String for the Event Hubs authorization Rule, which is generated when disaster recovery is enabled.
        /// </summary>
        [Input("primaryConnectionStringAlias")]
        public Input<string>? PrimaryConnectionStringAlias { get; set; }

        /// <summary>
        /// The Primary Key for the Event Hubs authorization Rule.
        /// </summary>
        [Input("primaryKey")]
        public Input<string>? PrimaryKey { get; set; }

        /// <summary>
        /// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The Secondary Connection String for the Event Hubs Authorization Rule.
        /// </summary>
        [Input("secondaryConnectionString")]
        public Input<string>? SecondaryConnectionString { get; set; }

        /// <summary>
        /// The alias of the Secondary Connection String for the Event Hubs Authorization Rule, which is generated when disaster recovery is enabled.
        /// </summary>
        [Input("secondaryConnectionStringAlias")]
        public Input<string>? SecondaryConnectionStringAlias { get; set; }

        /// <summary>
        /// The Secondary Key for the Event Hubs Authorization Rule.
        /// </summary>
        [Input("secondaryKey")]
        public Input<string>? SecondaryKey { get; set; }

        /// <summary>
        /// Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
        /// </summary>
        [Input("send")]
        public Input<bool>? Send { get; set; }

        public EventHubAuthorizationRuleState()
        {
        }
    }
}

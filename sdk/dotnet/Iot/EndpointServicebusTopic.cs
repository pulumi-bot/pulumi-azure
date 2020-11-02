// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Iot
{
    /// <summary>
    /// Manages an IotHub ServiceBus Topic Endpoint
    /// 
    /// &gt; **NOTE:** Endpoints can be defined either directly on the `azure.iot.IoTHub` resource, or using the `azurerm_iothub_endpoint_*` resources - but the two ways of defining the endpoints cannot be used together. If both are used against the same IoTHub, spurious changes will occur. Also, defining a `azurerm_iothub_endpoint_*` resource and another endpoint of a different type directly on the `azure.iot.IoTHub` resource is not supported.
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
    ///             Location = "East US",
    ///         });
    ///         var exampleNamespace = new Azure.ServiceBus.Namespace("exampleNamespace", new Azure.ServiceBus.NamespaceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Sku = "Standard",
    ///         });
    ///         var exampleTopic = new Azure.ServiceBus.Topic("exampleTopic", new Azure.ServiceBus.TopicArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             NamespaceName = exampleNamespace.Name,
    ///         });
    ///         var exampleTopicAuthorizationRule = new Azure.ServiceBus.TopicAuthorizationRule("exampleTopicAuthorizationRule", new Azure.ServiceBus.TopicAuthorizationRuleArgs
    ///         {
    ///             NamespaceName = exampleNamespace.Name,
    ///             TopicName = exampleTopic.Name,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Listen = false,
    ///             Send = true,
    ///             Manage = false,
    ///         });
    ///         var exampleIoTHub = new Azure.Iot.IoTHub("exampleIoTHub", new Azure.Iot.IoTHubArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Sku = new Azure.Iot.Inputs.IoTHubSkuArgs
    ///             {
    ///                 Name = "B1",
    ///                 Tier = "Basic",
    ///                 Capacity = 1,
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "purpose", "example" },
    ///             },
    ///         });
    ///         var exampleEndpointServicebusTopic = new Azure.Iot.EndpointServicebusTopic("exampleEndpointServicebusTopic", new Azure.Iot.EndpointServicebusTopicArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             IothubName = exampleIoTHub.Name,
    ///             ConnectionString = exampleTopicAuthorizationRule.PrimaryConnectionString,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// IoTHub ServiceBus Topic Endpoint can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class EndpointServicebusTopic : Pulumi.CustomResource
    {
        /// <summary>
        /// The connection string for the endpoint.
        /// </summary>
        [Output("connectionString")]
        public Output<string> ConnectionString { get; private set; } = null!;

        [Output("iothubName")]
        public Output<string> IothubName { get; private set; } = null!;

        /// <summary>
        /// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a EndpointServicebusTopic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EndpointServicebusTopic(string name, EndpointServicebusTopicArgs args, CustomResourceOptions? options = null)
            : base("azure:iot/endpointServicebusTopic:EndpointServicebusTopic", name, args ?? new EndpointServicebusTopicArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EndpointServicebusTopic(string name, Input<string> id, EndpointServicebusTopicState? state = null, CustomResourceOptions? options = null)
            : base("azure:iot/endpointServicebusTopic:EndpointServicebusTopic", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EndpointServicebusTopic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EndpointServicebusTopic Get(string name, Input<string> id, EndpointServicebusTopicState? state = null, CustomResourceOptions? options = null)
        {
            return new EndpointServicebusTopic(name, id, state, options);
        }
    }

    public sealed class EndpointServicebusTopicArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The connection string for the endpoint.
        /// </summary>
        [Input("connectionString", required: true)]
        public Input<string> ConnectionString { get; set; } = null!;

        [Input("iothubName", required: true)]
        public Input<string> IothubName { get; set; } = null!;

        /// <summary>
        /// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public EndpointServicebusTopicArgs()
        {
        }
    }

    public sealed class EndpointServicebusTopicState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The connection string for the endpoint.
        /// </summary>
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

        [Input("iothubName")]
        public Input<string>? IothubName { get; set; }

        /// <summary>
        /// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        public EndpointServicebusTopicState()
        {
        }
    }
}

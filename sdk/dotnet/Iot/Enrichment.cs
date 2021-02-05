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
    /// Manages an IotHub Enrichment
    /// 
    /// &gt; **NOTE:** Enrichment can be defined either directly on the `azure.iot.IoTHub` resource, or using the `azure.iot.Enrichment` resources - but the two cannot be used together. If both are used against the same IoTHub, spurious changes will occur.
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
    ///             Location = "West US",
    ///         });
    ///         var exampleAccount = new Azure.Storage.Account("exampleAccount", new Azure.Storage.AccountArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             AccountTier = "Standard",
    ///             AccountReplicationType = "LRS",
    ///         });
    ///         var exampleContainer = new Azure.Storage.Container("exampleContainer", new Azure.Storage.ContainerArgs
    ///         {
    ///             StorageAccountName = exampleAccount.Name,
    ///             ContainerAccessType = "private",
    ///         });
    ///         var exampleIoTHub = new Azure.Iot.IoTHub("exampleIoTHub", new Azure.Iot.IoTHubArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Sku = new Azure.Iot.Inputs.IoTHubSkuArgs
    ///             {
    ///                 Name = "S1",
    ///                 Capacity = 1,
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "purpose", "testing" },
    ///             },
    ///         });
    ///         var exampleEndpointStorageContainer = new Azure.Iot.EndpointStorageContainer("exampleEndpointStorageContainer", new Azure.Iot.EndpointStorageContainerArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             IothubName = exampleIoTHub.Name,
    ///             ConnectionString = exampleAccount.PrimaryBlobConnectionString,
    ///             BatchFrequencyInSeconds = 60,
    ///             MaxChunkSizeInBytes = 10485760,
    ///             ContainerName = exampleContainer.Name,
    ///             Encoding = "Avro",
    ///             FileNameFormat = "{iothub}/{partition}_{YYYY}_{MM}_{DD}_{HH}_{mm}",
    ///         });
    ///         var exampleRoute = new Azure.Iot.Route("exampleRoute", new Azure.Iot.RouteArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             IothubName = exampleIoTHub.Name,
    ///             Source = "DeviceMessages",
    ///             Condition = "true",
    ///             EndpointNames = 
    ///             {
    ///                 exampleEndpointStorageContainer.Name,
    ///             },
    ///             Enabled = true,
    ///         });
    ///         var exampleEnrichment = new Azure.Iot.Enrichment("exampleEnrichment", new Azure.Iot.EnrichmentArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             IothubName = exampleIoTHub.Name,
    ///             Key = "example",
    ///             Value = "my value",
    ///             EndpointNames = 
    ///             {
    ///                 exampleEndpointStorageContainer.Name,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// IoTHub Enrichment can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:iot/enrichment:Enrichment enrichment1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/IotHubs/hub1/Enrichments/enrichment1
    /// ```
    /// </summary>
    [AzureResourceType("azure:iot/enrichment:Enrichment")]
    public partial class Enrichment : Pulumi.CustomResource
    {
        /// <summary>
        /// The list of endpoints which will be enriched.
        /// </summary>
        [Output("endpointNames")]
        public Output<ImmutableArray<string>> EndpointNames { get; private set; } = null!;

        [Output("iothubName")]
        public Output<string> IothubName { get; private set; } = null!;

        /// <summary>
        /// The key of the enrichment.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The value of the enrichment. Value can be any static string, the name of the IoT hub sending the message (use `$iothubname`) or information from the device twin (ex: `$twin.tags.latitude`)
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;


        /// <summary>
        /// Create a Enrichment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Enrichment(string name, EnrichmentArgs args, CustomResourceOptions? options = null)
            : base("azure:iot/enrichment:Enrichment", name, args ?? new EnrichmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Enrichment(string name, Input<string> id, EnrichmentState? state = null, CustomResourceOptions? options = null)
            : base("azure:iot/enrichment:Enrichment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Enrichment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Enrichment Get(string name, Input<string> id, EnrichmentState? state = null, CustomResourceOptions? options = null)
        {
            return new Enrichment(name, id, state, options);
        }
    }

    public sealed class EnrichmentArgs : Pulumi.ResourceArgs
    {
        [Input("endpointNames", required: true)]
        private InputList<string>? _endpointNames;

        /// <summary>
        /// The list of endpoints which will be enriched.
        /// </summary>
        public InputList<string> EndpointNames
        {
            get => _endpointNames ?? (_endpointNames = new InputList<string>());
            set => _endpointNames = value;
        }

        [Input("iothubName", required: true)]
        public Input<string> IothubName { get; set; } = null!;

        /// <summary>
        /// The key of the enrichment.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The value of the enrichment. Value can be any static string, the name of the IoT hub sending the message (use `$iothubname`) or information from the device twin (ex: `$twin.tags.latitude`)
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public EnrichmentArgs()
        {
        }
    }

    public sealed class EnrichmentState : Pulumi.ResourceArgs
    {
        [Input("endpointNames")]
        private InputList<string>? _endpointNames;

        /// <summary>
        /// The list of endpoints which will be enriched.
        /// </summary>
        public InputList<string> EndpointNames
        {
            get => _endpointNames ?? (_endpointNames = new InputList<string>());
            set => _endpointNames = value;
        }

        [Input("iothubName")]
        public Input<string>? IothubName { get; set; }

        /// <summary>
        /// The key of the enrichment.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The value of the enrichment. Value can be any static string, the name of the IoT hub sending the message (use `$iothubname`) or information from the device twin (ex: `$twin.tags.latitude`)
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public EnrichmentState()
        {
        }
    }
}

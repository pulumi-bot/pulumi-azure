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
    /// Manages an IotHub Storage Container Endpoint
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
    ///         });
    ///         var exampleEndpointStorageContainer = new Azure.Iot.EndpointStorageContainer("exampleEndpointStorageContainer", new Azure.Iot.EndpointStorageContainerArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             IothubName = exampleIoTHub.Name,
    ///             ContainerName = "acctestcont",
    ///             ConnectionString = exampleAccount.PrimaryBlobConnectionString,
    ///             FileNameFormat = "{iothub}/{partition}_{YYYY}_{MM}_{DD}_{HH}_{mm}",
    ///             BatchFrequencyInSeconds = 60,
    ///             MaxChunkSizeInBytes = 10485760,
    ///             Encoding = "JSON",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// IoTHub Storage Container Endpoint can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class EndpointStorageContainer : Pulumi.CustomResource
    {
        /// <summary>
        /// Time interval at which blobs are written to storage. Value should be between 60 and 720 seconds. Default value is 300 seconds.
        /// </summary>
        [Output("batchFrequencyInSeconds")]
        public Output<int?> BatchFrequencyInSeconds { get; private set; } = null!;

        /// <summary>
        /// The connection string for the endpoint.
        /// </summary>
        [Output("connectionString")]
        public Output<string> ConnectionString { get; private set; } = null!;

        /// <summary>
        /// The name of storage container in the storage account.
        /// *
        /// </summary>
        [Output("containerName")]
        public Output<string> ContainerName { get; private set; } = null!;

        /// <summary>
        /// Encoding that is used to serialize messages to blobs. Supported values are 'avro' and 'avrodeflate'. Default value is 'avro'.
        /// </summary>
        [Output("encoding")]
        public Output<string?> Encoding { get; private set; } = null!;

        /// <summary>
        /// File name format for the blob. Default format is ``{iothub}/{partition}/{YYYY}/{MM}/{DD}/{HH}/{mm}``. All parameters are mandatory but can be reordered.
        /// </summary>
        [Output("fileNameFormat")]
        public Output<string?> FileNameFormat { get; private set; } = null!;

        /// <summary>
        /// The name of the IoTHub to which this Storage Container Endpoint belongs. Changing this forces a new resource to be created.
        /// </summary>
        [Output("iothubName")]
        public Output<string> IothubName { get; private set; } = null!;

        /// <summary>
        /// Maximum number of bytes for each blob written to storage. Value should be between 10485760(10MB) and 524288000(500MB). Default value is 314572800(300MB).
        /// </summary>
        [Output("maxChunkSizeInBytes")]
        public Output<int?> MaxChunkSizeInBytes { get; private set; } = null!;

        /// <summary>
        /// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group under which the IotHub Storage Container Endpoint resource has to be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a EndpointStorageContainer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EndpointStorageContainer(string name, EndpointStorageContainerArgs args, CustomResourceOptions? options = null)
            : base("azure:iot/endpointStorageContainer:EndpointStorageContainer", name, args ?? new EndpointStorageContainerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EndpointStorageContainer(string name, Input<string> id, EndpointStorageContainerState? state = null, CustomResourceOptions? options = null)
            : base("azure:iot/endpointStorageContainer:EndpointStorageContainer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EndpointStorageContainer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EndpointStorageContainer Get(string name, Input<string> id, EndpointStorageContainerState? state = null, CustomResourceOptions? options = null)
        {
            return new EndpointStorageContainer(name, id, state, options);
        }
    }

    public sealed class EndpointStorageContainerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time interval at which blobs are written to storage. Value should be between 60 and 720 seconds. Default value is 300 seconds.
        /// </summary>
        [Input("batchFrequencyInSeconds")]
        public Input<int>? BatchFrequencyInSeconds { get; set; }

        /// <summary>
        /// The connection string for the endpoint.
        /// </summary>
        [Input("connectionString", required: true)]
        public Input<string> ConnectionString { get; set; } = null!;

        /// <summary>
        /// The name of storage container in the storage account.
        /// *
        /// </summary>
        [Input("containerName", required: true)]
        public Input<string> ContainerName { get; set; } = null!;

        /// <summary>
        /// Encoding that is used to serialize messages to blobs. Supported values are 'avro' and 'avrodeflate'. Default value is 'avro'.
        /// </summary>
        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        /// <summary>
        /// File name format for the blob. Default format is ``{iothub}/{partition}/{YYYY}/{MM}/{DD}/{HH}/{mm}``. All parameters are mandatory but can be reordered.
        /// </summary>
        [Input("fileNameFormat")]
        public Input<string>? FileNameFormat { get; set; }

        /// <summary>
        /// The name of the IoTHub to which this Storage Container Endpoint belongs. Changing this forces a new resource to be created.
        /// </summary>
        [Input("iothubName", required: true)]
        public Input<string> IothubName { get; set; } = null!;

        /// <summary>
        /// Maximum number of bytes for each blob written to storage. Value should be between 10485760(10MB) and 524288000(500MB). Default value is 314572800(300MB).
        /// </summary>
        [Input("maxChunkSizeInBytes")]
        public Input<int>? MaxChunkSizeInBytes { get; set; }

        /// <summary>
        /// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group under which the IotHub Storage Container Endpoint resource has to be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public EndpointStorageContainerArgs()
        {
        }
    }

    public sealed class EndpointStorageContainerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time interval at which blobs are written to storage. Value should be between 60 and 720 seconds. Default value is 300 seconds.
        /// </summary>
        [Input("batchFrequencyInSeconds")]
        public Input<int>? BatchFrequencyInSeconds { get; set; }

        /// <summary>
        /// The connection string for the endpoint.
        /// </summary>
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

        /// <summary>
        /// The name of storage container in the storage account.
        /// *
        /// </summary>
        [Input("containerName")]
        public Input<string>? ContainerName { get; set; }

        /// <summary>
        /// Encoding that is used to serialize messages to blobs. Supported values are 'avro' and 'avrodeflate'. Default value is 'avro'.
        /// </summary>
        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        /// <summary>
        /// File name format for the blob. Default format is ``{iothub}/{partition}/{YYYY}/{MM}/{DD}/{HH}/{mm}``. All parameters are mandatory but can be reordered.
        /// </summary>
        [Input("fileNameFormat")]
        public Input<string>? FileNameFormat { get; set; }

        /// <summary>
        /// The name of the IoTHub to which this Storage Container Endpoint belongs. Changing this forces a new resource to be created.
        /// </summary>
        [Input("iothubName")]
        public Input<string>? IothubName { get; set; }

        /// <summary>
        /// Maximum number of bytes for each blob written to storage. Value should be between 10485760(10MB) and 524288000(500MB). Default value is 314572800(300MB).
        /// </summary>
        [Input("maxChunkSizeInBytes")]
        public Input<int>? MaxChunkSizeInBytes { get; set; }

        /// <summary>
        /// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group under which the IotHub Storage Container Endpoint resource has to be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        public EndpointStorageContainerState()
        {
        }
    }
}

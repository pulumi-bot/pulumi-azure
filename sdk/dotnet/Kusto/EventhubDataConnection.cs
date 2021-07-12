// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Kusto
{
    /// <summary>
    /// Manages a Kusto (also known as Azure Data Explorer) EventHub Data Connection
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
    ///         var rg = new Azure.Core.ResourceGroup("rg", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var cluster = new Azure.Kusto.Cluster("cluster", new Azure.Kusto.ClusterArgs
    ///         {
    ///             Location = rg.Location,
    ///             ResourceGroupName = rg.Name,
    ///             Sku = new Azure.Kusto.Inputs.ClusterSkuArgs
    ///             {
    ///                 Name = "Standard_D13_v2",
    ///                 Capacity = 2,
    ///             },
    ///         });
    ///         var database = new Azure.Kusto.Database("database", new Azure.Kusto.DatabaseArgs
    ///         {
    ///             ResourceGroupName = rg.Name,
    ///             Location = rg.Location,
    ///             ClusterName = cluster.Name,
    ///             HotCachePeriod = "P7D",
    ///             SoftDeletePeriod = "P31D",
    ///         });
    ///         var eventhubNs = new Azure.EventHub.EventHubNamespace("eventhubNs", new Azure.EventHub.EventHubNamespaceArgs
    ///         {
    ///             Location = rg.Location,
    ///             ResourceGroupName = rg.Name,
    ///             Sku = "Standard",
    ///         });
    ///         var eventhub = new Azure.EventHub.EventHub("eventhub", new Azure.EventHub.EventHubArgs
    ///         {
    ///             NamespaceName = eventhubNs.Name,
    ///             ResourceGroupName = rg.Name,
    ///             PartitionCount = 1,
    ///             MessageRetention = 1,
    ///         });
    ///         var consumerGroup = new Azure.EventHub.ConsumerGroup("consumerGroup", new Azure.EventHub.ConsumerGroupArgs
    ///         {
    ///             NamespaceName = eventhubNs.Name,
    ///             EventhubName = eventhub.Name,
    ///             ResourceGroupName = rg.Name,
    ///         });
    ///         var eventhubConnection = new Azure.Kusto.EventhubDataConnection("eventhubConnection", new Azure.Kusto.EventhubDataConnectionArgs
    ///         {
    ///             ResourceGroupName = rg.Name,
    ///             Location = rg.Location,
    ///             ClusterName = cluster.Name,
    ///             DatabaseName = database.Name,
    ///             EventhubId = eventhub.Id,
    ///             ConsumerGroup = consumerGroup.Name,
    ///             TableName = "my-table",
    ///             MappingRuleName = "my-table-mapping",
    ///             DataFormat = "JSON",
    ///         });
    ///         //(Optional)
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Kusto EventHub Data Connections can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:kusto/eventhubDataConnection:EventhubDataConnection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/Databases/database1/DataConnections/eventHubConnection1
    /// ```
    /// </summary>
    [AzureResourceType("azure:kusto/eventhubDataConnection:EventhubDataConnection")]
    public partial class EventhubDataConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// Specifies compression type for the connection. Allowed values: `GZip` and `None`. Defaults to `None`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("compression")]
        public Output<string?> Compression { get; private set; } = null!;

        /// <summary>
        /// Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
        /// </summary>
        [Output("consumerGroup")]
        public Output<string> ConsumerGroup { get; private set; } = null!;

        /// <summary>
        /// Specifies the data format of the EventHub messages. Allowed values: `APACHEAVRO`, `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
        /// </summary>
        [Output("dataFormat")]
        public Output<string?> DataFormat { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("databaseName")]
        public Output<string> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// Specifies a list of system properties for the Event Hub.
        /// </summary>
        [Output("eventSystemProperties")]
        public Output<ImmutableArray<string>> EventSystemProperties { get; private set; } = null!;

        /// <summary>
        /// Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
        /// </summary>
        [Output("eventhubId")]
        public Output<string> EventhubId { get; private set; } = null!;

        /// <summary>
        /// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
        /// </summary>
        [Output("mappingRuleName")]
        public Output<string?> MappingRuleName { get; private set; } = null!;

        /// <summary>
        /// The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
        /// </summary>
        [Output("tableName")]
        public Output<string?> TableName { get; private set; } = null!;


        /// <summary>
        /// Create a EventhubDataConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventhubDataConnection(string name, EventhubDataConnectionArgs args, CustomResourceOptions? options = null)
            : base("azure:kusto/eventhubDataConnection:EventhubDataConnection", name, args ?? new EventhubDataConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventhubDataConnection(string name, Input<string> id, EventhubDataConnectionState? state = null, CustomResourceOptions? options = null)
            : base("azure:kusto/eventhubDataConnection:EventhubDataConnection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EventhubDataConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventhubDataConnection Get(string name, Input<string> id, EventhubDataConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new EventhubDataConnection(name, id, state, options);
        }
    }

    public sealed class EventhubDataConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// Specifies compression type for the connection. Allowed values: `GZip` and `None`. Defaults to `None`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("compression")]
        public Input<string>? Compression { get; set; }

        /// <summary>
        /// Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
        /// </summary>
        [Input("consumerGroup", required: true)]
        public Input<string> ConsumerGroup { get; set; } = null!;

        /// <summary>
        /// Specifies the data format of the EventHub messages. Allowed values: `APACHEAVRO`, `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
        /// </summary>
        [Input("dataFormat")]
        public Input<string>? DataFormat { get; set; }

        /// <summary>
        /// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        [Input("eventSystemProperties")]
        private InputList<string>? _eventSystemProperties;

        /// <summary>
        /// Specifies a list of system properties for the Event Hub.
        /// </summary>
        public InputList<string> EventSystemProperties
        {
            get => _eventSystemProperties ?? (_eventSystemProperties = new InputList<string>());
            set => _eventSystemProperties = value;
        }

        /// <summary>
        /// Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
        /// </summary>
        [Input("eventhubId", required: true)]
        public Input<string> EventhubId { get; set; } = null!;

        /// <summary>
        /// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
        /// </summary>
        [Input("mappingRuleName")]
        public Input<string>? MappingRuleName { get; set; }

        /// <summary>
        /// The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
        /// </summary>
        [Input("tableName")]
        public Input<string>? TableName { get; set; }

        public EventhubDataConnectionArgs()
        {
        }
    }

    public sealed class EventhubDataConnectionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// Specifies compression type for the connection. Allowed values: `GZip` and `None`. Defaults to `None`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("compression")]
        public Input<string>? Compression { get; set; }

        /// <summary>
        /// Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
        /// </summary>
        [Input("consumerGroup")]
        public Input<string>? ConsumerGroup { get; set; }

        /// <summary>
        /// Specifies the data format of the EventHub messages. Allowed values: `APACHEAVRO`, `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
        /// </summary>
        [Input("dataFormat")]
        public Input<string>? DataFormat { get; set; }

        /// <summary>
        /// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        [Input("eventSystemProperties")]
        private InputList<string>? _eventSystemProperties;

        /// <summary>
        /// Specifies a list of system properties for the Event Hub.
        /// </summary>
        public InputList<string> EventSystemProperties
        {
            get => _eventSystemProperties ?? (_eventSystemProperties = new InputList<string>());
            set => _eventSystemProperties = value;
        }

        /// <summary>
        /// Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
        /// </summary>
        [Input("eventhubId")]
        public Input<string>? EventhubId { get; set; }

        /// <summary>
        /// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
        /// </summary>
        [Input("mappingRuleName")]
        public Input<string>? MappingRuleName { get; set; }

        /// <summary>
        /// The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
        /// </summary>
        [Input("tableName")]
        public Input<string>? TableName { get; set; }

        public EventhubDataConnectionState()
        {
        }
    }
}

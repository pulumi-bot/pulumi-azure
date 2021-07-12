// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.HDInsight
{
    /// <summary>
    /// Manages a HDInsight Kafka Cluster.
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
    ///         var exampleKafkaCluster = new Azure.HDInsight.KafkaCluster("exampleKafkaCluster", new Azure.HDInsight.KafkaClusterArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             ClusterVersion = "4.0",
    ///             Tier = "Standard",
    ///             ComponentVersion = new Azure.HDInsight.Inputs.KafkaClusterComponentVersionArgs
    ///             {
    ///                 Kafka = "2.1",
    ///             },
    ///             Gateway = new Azure.HDInsight.Inputs.KafkaClusterGatewayArgs
    ///             {
    ///                 Enabled = true,
    ///                 Username = "acctestusrgw",
    ///                 Password = "Password123!",
    ///             },
    ///             StorageAccounts = 
    ///             {
    ///                 new Azure.HDInsight.Inputs.KafkaClusterStorageAccountArgs
    ///                 {
    ///                     StorageContainerId = exampleContainer.Id,
    ///                     StorageAccountKey = exampleAccount.PrimaryAccessKey,
    ///                     IsDefault = true,
    ///                 },
    ///             },
    ///             Roles = new Azure.HDInsight.Inputs.KafkaClusterRolesArgs
    ///             {
    ///                 HeadNode = new Azure.HDInsight.Inputs.KafkaClusterRolesHeadNodeArgs
    ///                 {
    ///                     VmSize = "Standard_D3_V2",
    ///                     Username = "acctestusrvm",
    ///                     Password = "AccTestvdSC4daf986!",
    ///                 },
    ///                 WorkerNode = new Azure.HDInsight.Inputs.KafkaClusterRolesWorkerNodeArgs
    ///                 {
    ///                     VmSize = "Standard_D3_V2",
    ///                     Username = "acctestusrvm",
    ///                     Password = "AccTestvdSC4daf986!",
    ///                     NumberOfDisksPerNode = 3,
    ///                     TargetInstanceCount = 3,
    ///                 },
    ///                 ZookeeperNode = new Azure.HDInsight.Inputs.KafkaClusterRolesZookeeperNodeArgs
    ///                 {
    ///                     VmSize = "Standard_D3_V2",
    ///                     Username = "acctestusrvm",
    ///                     Password = "AccTestvdSC4daf986!",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// HDInsight Kafka Clusters can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:hdinsight/kafkaCluster:KafkaCluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.HDInsight/clusters/cluster1
    /// ```
    /// </summary>
    [AzureResourceType("azure:hdinsight/kafkaCluster:KafkaCluster")]
    public partial class KafkaCluster : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Output("clusterVersion")]
        public Output<string> ClusterVersion { get; private set; } = null!;

        /// <summary>
        /// A `component_version` block as defined below.
        /// </summary>
        [Output("componentVersion")]
        public Output<Outputs.KafkaClusterComponentVersion> ComponentVersion { get; private set; } = null!;

        /// <summary>
        /// Whether encryption in transit is enabled for this HDInsight Kafka Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Output("encryptionInTransitEnabled")]
        public Output<bool?> EncryptionInTransitEnabled { get; private set; } = null!;

        /// <summary>
        /// A `gateway` block as defined below.
        /// </summary>
        [Output("gateway")]
        public Output<Outputs.KafkaClusterGateway> Gateway { get; private set; } = null!;

        /// <summary>
        /// The HTTPS Connectivity Endpoint for this HDInsight Kafka Cluster.
        /// </summary>
        [Output("httpsEndpoint")]
        public Output<string> HttpsEndpoint { get; private set; } = null!;

        /// <summary>
        /// The Kafka Rest Proxy Endpoint for this HDInsight Kafka Cluster.
        /// </summary>
        [Output("kafkaRestProxyEndpoint")]
        public Output<string> KafkaRestProxyEndpoint { get; private set; } = null!;

        /// <summary>
        /// Specifies the Azure Region which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// A `metastores` block as defined below.
        /// </summary>
        [Output("metastores")]
        public Output<Outputs.KafkaClusterMetastores?> Metastores { get; private set; } = null!;

        /// <summary>
        /// A `monitor` block as defined below.
        /// </summary>
        [Output("monitor")]
        public Output<Outputs.KafkaClusterMonitor?> Monitor { get; private set; } = null!;

        /// <summary>
        /// Specifies the name for this HDInsight Kafka Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group in which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A `rest_proxy` block as defined below.
        /// </summary>
        [Output("restProxy")]
        public Output<Outputs.KafkaClusterRestProxy?> RestProxy { get; private set; } = null!;

        /// <summary>
        /// A `roles` block as defined below.
        /// </summary>
        [Output("roles")]
        public Output<Outputs.KafkaClusterRoles> Roles { get; private set; } = null!;

        /// <summary>
        /// The SSH Connectivity Endpoint for this HDInsight Kafka Cluster.
        /// </summary>
        [Output("sshEndpoint")]
        public Output<string> SshEndpoint { get; private set; } = null!;

        /// <summary>
        /// A `storage_account_gen2` block as defined below.
        /// </summary>
        [Output("storageAccountGen2")]
        public Output<Outputs.KafkaClusterStorageAccountGen2?> StorageAccountGen2 { get; private set; } = null!;

        /// <summary>
        /// One or more `storage_account` block as defined below.
        /// </summary>
        [Output("storageAccounts")]
        public Output<ImmutableArray<Outputs.KafkaClusterStorageAccount>> StorageAccounts { get; private set; } = null!;

        /// <summary>
        /// A map of Tags which should be assigned to this HDInsight Kafka Cluster.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies the Tier which should be used for this HDInsight Kafka Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("tier")]
        public Output<string> Tier { get; private set; } = null!;

        /// <summary>
        /// The minimal supported TLS version. Possible values are `1.0`, `1.1` or `1.2`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("tlsMinVersion")]
        public Output<string?> TlsMinVersion { get; private set; } = null!;


        /// <summary>
        /// Create a KafkaCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KafkaCluster(string name, KafkaClusterArgs args, CustomResourceOptions? options = null)
            : base("azure:hdinsight/kafkaCluster:KafkaCluster", name, args ?? new KafkaClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KafkaCluster(string name, Input<string> id, KafkaClusterState? state = null, CustomResourceOptions? options = null)
            : base("azure:hdinsight/kafkaCluster:KafkaCluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KafkaCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KafkaCluster Get(string name, Input<string> id, KafkaClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new KafkaCluster(name, id, state, options);
        }
    }

    public sealed class KafkaClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("clusterVersion", required: true)]
        public Input<string> ClusterVersion { get; set; } = null!;

        /// <summary>
        /// A `component_version` block as defined below.
        /// </summary>
        [Input("componentVersion", required: true)]
        public Input<Inputs.KafkaClusterComponentVersionArgs> ComponentVersion { get; set; } = null!;

        /// <summary>
        /// Whether encryption in transit is enabled for this HDInsight Kafka Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("encryptionInTransitEnabled")]
        public Input<bool>? EncryptionInTransitEnabled { get; set; }

        /// <summary>
        /// A `gateway` block as defined below.
        /// </summary>
        [Input("gateway", required: true)]
        public Input<Inputs.KafkaClusterGatewayArgs> Gateway { get; set; } = null!;

        /// <summary>
        /// Specifies the Azure Region which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A `metastores` block as defined below.
        /// </summary>
        [Input("metastores")]
        public Input<Inputs.KafkaClusterMetastoresArgs>? Metastores { get; set; }

        /// <summary>
        /// A `monitor` block as defined below.
        /// </summary>
        [Input("monitor")]
        public Input<Inputs.KafkaClusterMonitorArgs>? Monitor { get; set; }

        /// <summary>
        /// Specifies the name for this HDInsight Kafka Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group in which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `rest_proxy` block as defined below.
        /// </summary>
        [Input("restProxy")]
        public Input<Inputs.KafkaClusterRestProxyArgs>? RestProxy { get; set; }

        /// <summary>
        /// A `roles` block as defined below.
        /// </summary>
        [Input("roles", required: true)]
        public Input<Inputs.KafkaClusterRolesArgs> Roles { get; set; } = null!;

        /// <summary>
        /// A `storage_account_gen2` block as defined below.
        /// </summary>
        [Input("storageAccountGen2")]
        public Input<Inputs.KafkaClusterStorageAccountGen2Args>? StorageAccountGen2 { get; set; }

        [Input("storageAccounts")]
        private InputList<Inputs.KafkaClusterStorageAccountArgs>? _storageAccounts;

        /// <summary>
        /// One or more `storage_account` block as defined below.
        /// </summary>
        public InputList<Inputs.KafkaClusterStorageAccountArgs> StorageAccounts
        {
            get => _storageAccounts ?? (_storageAccounts = new InputList<Inputs.KafkaClusterStorageAccountArgs>());
            set => _storageAccounts = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of Tags which should be assigned to this HDInsight Kafka Cluster.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the Tier which should be used for this HDInsight Kafka Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("tier", required: true)]
        public Input<string> Tier { get; set; } = null!;

        /// <summary>
        /// The minimal supported TLS version. Possible values are `1.0`, `1.1` or `1.2`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("tlsMinVersion")]
        public Input<string>? TlsMinVersion { get; set; }

        public KafkaClusterArgs()
        {
        }
    }

    public sealed class KafkaClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("clusterVersion")]
        public Input<string>? ClusterVersion { get; set; }

        /// <summary>
        /// A `component_version` block as defined below.
        /// </summary>
        [Input("componentVersion")]
        public Input<Inputs.KafkaClusterComponentVersionGetArgs>? ComponentVersion { get; set; }

        /// <summary>
        /// Whether encryption in transit is enabled for this HDInsight Kafka Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("encryptionInTransitEnabled")]
        public Input<bool>? EncryptionInTransitEnabled { get; set; }

        /// <summary>
        /// A `gateway` block as defined below.
        /// </summary>
        [Input("gateway")]
        public Input<Inputs.KafkaClusterGatewayGetArgs>? Gateway { get; set; }

        /// <summary>
        /// The HTTPS Connectivity Endpoint for this HDInsight Kafka Cluster.
        /// </summary>
        [Input("httpsEndpoint")]
        public Input<string>? HttpsEndpoint { get; set; }

        /// <summary>
        /// The Kafka Rest Proxy Endpoint for this HDInsight Kafka Cluster.
        /// </summary>
        [Input("kafkaRestProxyEndpoint")]
        public Input<string>? KafkaRestProxyEndpoint { get; set; }

        /// <summary>
        /// Specifies the Azure Region which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A `metastores` block as defined below.
        /// </summary>
        [Input("metastores")]
        public Input<Inputs.KafkaClusterMetastoresGetArgs>? Metastores { get; set; }

        /// <summary>
        /// A `monitor` block as defined below.
        /// </summary>
        [Input("monitor")]
        public Input<Inputs.KafkaClusterMonitorGetArgs>? Monitor { get; set; }

        /// <summary>
        /// Specifies the name for this HDInsight Kafka Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group in which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// A `rest_proxy` block as defined below.
        /// </summary>
        [Input("restProxy")]
        public Input<Inputs.KafkaClusterRestProxyGetArgs>? RestProxy { get; set; }

        /// <summary>
        /// A `roles` block as defined below.
        /// </summary>
        [Input("roles")]
        public Input<Inputs.KafkaClusterRolesGetArgs>? Roles { get; set; }

        /// <summary>
        /// The SSH Connectivity Endpoint for this HDInsight Kafka Cluster.
        /// </summary>
        [Input("sshEndpoint")]
        public Input<string>? SshEndpoint { get; set; }

        /// <summary>
        /// A `storage_account_gen2` block as defined below.
        /// </summary>
        [Input("storageAccountGen2")]
        public Input<Inputs.KafkaClusterStorageAccountGen2GetArgs>? StorageAccountGen2 { get; set; }

        [Input("storageAccounts")]
        private InputList<Inputs.KafkaClusterStorageAccountGetArgs>? _storageAccounts;

        /// <summary>
        /// One or more `storage_account` block as defined below.
        /// </summary>
        public InputList<Inputs.KafkaClusterStorageAccountGetArgs> StorageAccounts
        {
            get => _storageAccounts ?? (_storageAccounts = new InputList<Inputs.KafkaClusterStorageAccountGetArgs>());
            set => _storageAccounts = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of Tags which should be assigned to this HDInsight Kafka Cluster.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the Tier which should be used for this HDInsight Kafka Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        /// <summary>
        /// The minimal supported TLS version. Possible values are `1.0`, `1.1` or `1.2`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("tlsMinVersion")]
        public Input<string>? TlsMinVersion { get; set; }

        public KafkaClusterState()
        {
        }
    }
}

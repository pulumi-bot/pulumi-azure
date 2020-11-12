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
    /// Manages a HDInsight Interactive Query Cluster.
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
    ///         var exampleInteractiveQueryCluster = new Azure.HDInsight.InteractiveQueryCluster("exampleInteractiveQueryCluster", new Azure.HDInsight.InteractiveQueryClusterArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             ClusterVersion = "3.6",
    ///             Tier = "Standard",
    ///             ComponentVersion = new Azure.HDInsight.Inputs.InteractiveQueryClusterComponentVersionArgs
    ///             {
    ///                 InteractiveHive = "2.1",
    ///             },
    ///             Gateway = new Azure.HDInsight.Inputs.InteractiveQueryClusterGatewayArgs
    ///             {
    ///                 Enabled = true,
    ///                 Username = "acctestusrgw",
    ///                 Password = "Password!",
    ///             },
    ///             StorageAccounts = 
    ///             {
    ///                 new Azure.HDInsight.Inputs.InteractiveQueryClusterStorageAccountArgs
    ///                 {
    ///                     StorageContainerId = exampleContainer.Id,
    ///                     StorageAccountKey = exampleAccount.PrimaryAccessKey,
    ///                     IsDefault = true,
    ///                 },
    ///             },
    ///             Roles = new Azure.HDInsight.Inputs.InteractiveQueryClusterRolesArgs
    ///             {
    ///                 HeadNode = new Azure.HDInsight.Inputs.InteractiveQueryClusterRolesHeadNodeArgs
    ///                 {
    ///                     VmSize = "Standard_D13_V2",
    ///                     Username = "acctestusrvm",
    ///                     Password = "AccTestvdSC4daf986!",
    ///                 },
    ///                 WorkerNode = new Azure.HDInsight.Inputs.InteractiveQueryClusterRolesWorkerNodeArgs
    ///                 {
    ///                     VmSize = "Standard_D14_V2",
    ///                     Username = "acctestusrvm",
    ///                     Password = "AccTestvdSC4daf986!",
    ///                     TargetInstanceCount = 3,
    ///                 },
    ///                 ZookeeperNode = new Azure.HDInsight.Inputs.InteractiveQueryClusterRolesZookeeperNodeArgs
    ///                 {
    ///                     VmSize = "Standard_A4_V2",
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
    /// HDInsight Interactive Query Clusters can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:hdinsight/interactiveQueryCluster:InteractiveQueryCluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.HDInsight/clusters/cluster1}
    /// ```
    /// </summary>
    public partial class InteractiveQueryCluster : Pulumi.CustomResource
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
        public Output<Outputs.InteractiveQueryClusterComponentVersion> ComponentVersion { get; private set; } = null!;

        /// <summary>
        /// A `gateway` block as defined below.
        /// </summary>
        [Output("gateway")]
        public Output<Outputs.InteractiveQueryClusterGateway> Gateway { get; private set; } = null!;

        /// <summary>
        /// The HTTPS Connectivity Endpoint for this HDInsight Interactive Query Cluster.
        /// </summary>
        [Output("httpsEndpoint")]
        public Output<string> HttpsEndpoint { get; private set; } = null!;

        /// <summary>
        /// Specifies the Azure Region which this HDInsight Interactive Query Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// A `metastores` block as defined below.
        /// </summary>
        [Output("metastores")]
        public Output<Outputs.InteractiveQueryClusterMetastores?> Metastores { get; private set; } = null!;

        /// <summary>
        /// A `monitor` block as defined below.
        /// </summary>
        [Output("monitor")]
        public Output<Outputs.InteractiveQueryClusterMonitor?> Monitor { get; private set; } = null!;

        /// <summary>
        /// Specifies the name for this HDInsight Interactive Query Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group in which this HDInsight Interactive Query Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A `roles` block as defined below.
        /// </summary>
        [Output("roles")]
        public Output<Outputs.InteractiveQueryClusterRoles> Roles { get; private set; } = null!;

        /// <summary>
        /// The SSH Connectivity Endpoint for this HDInsight Interactive Query Cluster.
        /// </summary>
        [Output("sshEndpoint")]
        public Output<string> SshEndpoint { get; private set; } = null!;

        /// <summary>
        /// A `storage_account_gen2` block as defined below.
        /// </summary>
        [Output("storageAccountGen2")]
        public Output<Outputs.InteractiveQueryClusterStorageAccountGen2?> StorageAccountGen2 { get; private set; } = null!;

        /// <summary>
        /// One or more `storage_account` block as defined below.
        /// </summary>
        [Output("storageAccounts")]
        public Output<ImmutableArray<Outputs.InteractiveQueryClusterStorageAccount>> StorageAccounts { get; private set; } = null!;

        /// <summary>
        /// A map of Tags which should be assigned to this HDInsight Interactive Query Cluster.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies the Tier which should be used for this HDInsight Interactive Query Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("tier")]
        public Output<string> Tier { get; private set; } = null!;

        [Output("tlsMinVersion")]
        public Output<string?> TlsMinVersion { get; private set; } = null!;


        /// <summary>
        /// Create a InteractiveQueryCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InteractiveQueryCluster(string name, InteractiveQueryClusterArgs args, CustomResourceOptions? options = null)
            : base("azure:hdinsight/interactiveQueryCluster:InteractiveQueryCluster", name, args ?? new InteractiveQueryClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InteractiveQueryCluster(string name, Input<string> id, InteractiveQueryClusterState? state = null, CustomResourceOptions? options = null)
            : base("azure:hdinsight/interactiveQueryCluster:InteractiveQueryCluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InteractiveQueryCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InteractiveQueryCluster Get(string name, Input<string> id, InteractiveQueryClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new InteractiveQueryCluster(name, id, state, options);
        }
    }

    public sealed class InteractiveQueryClusterArgs : Pulumi.ResourceArgs
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
        public Input<Inputs.InteractiveQueryClusterComponentVersionArgs> ComponentVersion { get; set; } = null!;

        /// <summary>
        /// A `gateway` block as defined below.
        /// </summary>
        [Input("gateway", required: true)]
        public Input<Inputs.InteractiveQueryClusterGatewayArgs> Gateway { get; set; } = null!;

        /// <summary>
        /// Specifies the Azure Region which this HDInsight Interactive Query Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A `metastores` block as defined below.
        /// </summary>
        [Input("metastores")]
        public Input<Inputs.InteractiveQueryClusterMetastoresArgs>? Metastores { get; set; }

        /// <summary>
        /// A `monitor` block as defined below.
        /// </summary>
        [Input("monitor")]
        public Input<Inputs.InteractiveQueryClusterMonitorArgs>? Monitor { get; set; }

        /// <summary>
        /// Specifies the name for this HDInsight Interactive Query Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group in which this HDInsight Interactive Query Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `roles` block as defined below.
        /// </summary>
        [Input("roles", required: true)]
        public Input<Inputs.InteractiveQueryClusterRolesArgs> Roles { get; set; } = null!;

        /// <summary>
        /// A `storage_account_gen2` block as defined below.
        /// </summary>
        [Input("storageAccountGen2")]
        public Input<Inputs.InteractiveQueryClusterStorageAccountGen2Args>? StorageAccountGen2 { get; set; }

        [Input("storageAccounts")]
        private InputList<Inputs.InteractiveQueryClusterStorageAccountArgs>? _storageAccounts;

        /// <summary>
        /// One or more `storage_account` block as defined below.
        /// </summary>
        public InputList<Inputs.InteractiveQueryClusterStorageAccountArgs> StorageAccounts
        {
            get => _storageAccounts ?? (_storageAccounts = new InputList<Inputs.InteractiveQueryClusterStorageAccountArgs>());
            set => _storageAccounts = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of Tags which should be assigned to this HDInsight Interactive Query Cluster.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the Tier which should be used for this HDInsight Interactive Query Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("tier", required: true)]
        public Input<string> Tier { get; set; } = null!;

        [Input("tlsMinVersion")]
        public Input<string>? TlsMinVersion { get; set; }

        public InteractiveQueryClusterArgs()
        {
        }
    }

    public sealed class InteractiveQueryClusterState : Pulumi.ResourceArgs
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
        public Input<Inputs.InteractiveQueryClusterComponentVersionGetArgs>? ComponentVersion { get; set; }

        /// <summary>
        /// A `gateway` block as defined below.
        /// </summary>
        [Input("gateway")]
        public Input<Inputs.InteractiveQueryClusterGatewayGetArgs>? Gateway { get; set; }

        /// <summary>
        /// The HTTPS Connectivity Endpoint for this HDInsight Interactive Query Cluster.
        /// </summary>
        [Input("httpsEndpoint")]
        public Input<string>? HttpsEndpoint { get; set; }

        /// <summary>
        /// Specifies the Azure Region which this HDInsight Interactive Query Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A `metastores` block as defined below.
        /// </summary>
        [Input("metastores")]
        public Input<Inputs.InteractiveQueryClusterMetastoresGetArgs>? Metastores { get; set; }

        /// <summary>
        /// A `monitor` block as defined below.
        /// </summary>
        [Input("monitor")]
        public Input<Inputs.InteractiveQueryClusterMonitorGetArgs>? Monitor { get; set; }

        /// <summary>
        /// Specifies the name for this HDInsight Interactive Query Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group in which this HDInsight Interactive Query Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// A `roles` block as defined below.
        /// </summary>
        [Input("roles")]
        public Input<Inputs.InteractiveQueryClusterRolesGetArgs>? Roles { get; set; }

        /// <summary>
        /// The SSH Connectivity Endpoint for this HDInsight Interactive Query Cluster.
        /// </summary>
        [Input("sshEndpoint")]
        public Input<string>? SshEndpoint { get; set; }

        /// <summary>
        /// A `storage_account_gen2` block as defined below.
        /// </summary>
        [Input("storageAccountGen2")]
        public Input<Inputs.InteractiveQueryClusterStorageAccountGen2GetArgs>? StorageAccountGen2 { get; set; }

        [Input("storageAccounts")]
        private InputList<Inputs.InteractiveQueryClusterStorageAccountGetArgs>? _storageAccounts;

        /// <summary>
        /// One or more `storage_account` block as defined below.
        /// </summary>
        public InputList<Inputs.InteractiveQueryClusterStorageAccountGetArgs> StorageAccounts
        {
            get => _storageAccounts ?? (_storageAccounts = new InputList<Inputs.InteractiveQueryClusterStorageAccountGetArgs>());
            set => _storageAccounts = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of Tags which should be assigned to this HDInsight Interactive Query Cluster.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the Tier which should be used for this HDInsight Interactive Query Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        [Input("tlsMinVersion")]
        public Input<string>? TlsMinVersion { get; set; }

        public InteractiveQueryClusterState()
        {
        }
    }
}

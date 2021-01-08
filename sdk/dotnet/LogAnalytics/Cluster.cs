// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.LogAnalytics
{
    /// <summary>
    /// !&gt; **Important** Due to capacity constraints, Microsoft requires you to pre-register your subscription IDs before you are allowed to create a Log Analytics cluster. Contact Microsoft, or open a support request to register your subscription IDs.
    /// 
    /// &gt; **Note:** Log Analytics Clusters are subject to 14-day soft delete policy. Clusters created with the same resource group &amp; name as a previously deleted cluster will be recovered rather than creating anew.
    /// 
    /// Manages a Log Analytics Cluster.
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
    ///         var exampleCluster = new Azure.LogAnalytics.Cluster("exampleCluster", new Azure.LogAnalytics.ClusterArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Identity = new Azure.LogAnalytics.Inputs.ClusterIdentityArgs
    ///             {
    ///                 Type = "SystemAssigned",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Log Analytics Clusters can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:loganalytics/cluster:Cluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.OperationalInsights/clusters/cluster1
    /// ```
    /// </summary>
    [AzureResourceType("azure:loganalytics/cluster:Cluster")]
    public partial class Cluster : Pulumi.CustomResource
    {
        /// <summary>
        /// The GUID of the cluster.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// A `identity` block as defined below. Changing this forces a new Log Analytics Cluster to be created.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ClusterIdentity> Identity { get; private set; } = null!;

        /// <summary>
        /// The Azure Region where the Log Analytics Cluster should exist. Changing this forces a new Log Analytics Cluster to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this Log Analytics Cluster. Changing this forces a new Log Analytics Cluster to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Log Analytics Cluster should exist. Changing this forces a new Log Analytics Cluster to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The capacity of the Log Analytics Cluster specified in GB/day. Defaults to 1000.
        /// </summary>
        [Output("sizeGb")]
        public Output<int?> SizeGb { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags which should be assigned to the Log Analytics Cluster.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("azure:loganalytics/cluster:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("azure:loganalytics/cluster:Cluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, state, options);
        }
    }

    public sealed class ClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `identity` block as defined below. Changing this forces a new Log Analytics Cluster to be created.
        /// </summary>
        [Input("identity", required: true)]
        public Input<Inputs.ClusterIdentityArgs> Identity { get; set; } = null!;

        /// <summary>
        /// The Azure Region where the Log Analytics Cluster should exist. Changing this forces a new Log Analytics Cluster to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name which should be used for this Log Analytics Cluster. Changing this forces a new Log Analytics Cluster to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Log Analytics Cluster should exist. Changing this forces a new Log Analytics Cluster to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The capacity of the Log Analytics Cluster specified in GB/day. Defaults to 1000.
        /// </summary>
        [Input("sizeGb")]
        public Input<int>? SizeGb { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Log Analytics Cluster.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ClusterArgs()
        {
        }
    }

    public sealed class ClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The GUID of the cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// A `identity` block as defined below. Changing this forces a new Log Analytics Cluster to be created.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ClusterIdentityGetArgs>? Identity { get; set; }

        /// <summary>
        /// The Azure Region where the Log Analytics Cluster should exist. Changing this forces a new Log Analytics Cluster to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name which should be used for this Log Analytics Cluster. Changing this forces a new Log Analytics Cluster to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Log Analytics Cluster should exist. Changing this forces a new Log Analytics Cluster to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The capacity of the Log Analytics Cluster specified in GB/day. Defaults to 1000.
        /// </summary>
        [Input("sizeGb")]
        public Input<int>? SizeGb { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Log Analytics Cluster.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ClusterState()
        {
        }
    }
}

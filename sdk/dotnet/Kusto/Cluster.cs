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
    /// Manages a Kusto (also known as Azure Data Explorer) Cluster
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
    ///             Location = "East US",
    ///         });
    ///         var example = new Azure.Kusto.Cluster("example", new Azure.Kusto.ClusterArgs
    ///         {
    ///             Location = rg.Location,
    ///             ResourceGroupName = rg.Name,
    ///             Sku = new Azure.Kusto.Inputs.ClusterSkuArgs
    ///             {
    ///                 Name = "Standard_D13_v2",
    ///                 Capacity = 2,
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "Environment", "Production" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Kusto Clusters can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:kusto/cluster:Cluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1
    /// ```
    /// </summary>
    [AzureResourceType("azure:kusto/cluster:Cluster")]
    public partial class Cluster : Pulumi.CustomResource
    {
        /// <summary>
        /// The Kusto Cluster URI to be used for data ingestion.
        /// </summary>
        [Output("dataIngestionUri")]
        public Output<string> DataIngestionUri { get; private set; } = null!;

        /// <summary>
        /// Specifies if the cluster's disks are encrypted.
        /// </summary>
        [Output("enableDiskEncryption")]
        public Output<bool?> EnableDiskEncryption { get; private set; } = null!;

        /// <summary>
        /// Specifies if the purge operations are enabled.
        /// </summary>
        [Output("enablePurge")]
        public Output<bool?> EnablePurge { get; private set; } = null!;

        /// <summary>
        /// Specifies if the streaming ingest is enabled.
        /// </summary>
        [Output("enableStreamingIngest")]
        public Output<bool?> EnableStreamingIngest { get; private set; } = null!;

        /// <summary>
        /// A identity block.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ClusterIdentity> Identity { get; private set; } = null!;

        /// <summary>
        /// An list of `language_extensions` to enable. Valid values are: `PYTHON` and `R`.
        /// </summary>
        [Output("languageExtensions")]
        public Output<ImmutableArray<string>> LanguageExtensions { get; private set; } = null!;

        /// <summary>
        /// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// An `optimized_auto_scale` block as defined below.
        /// </summary>
        [Output("optimizedAutoScale")]
        public Output<Outputs.ClusterOptimizedAutoScale?> OptimizedAutoScale { get; private set; } = null!;

        /// <summary>
        /// Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A `sku` block as defined below.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.ClusterSku> Sku { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies a list of tenant IDs that are trusted by the cluster.
        /// </summary>
        [Output("trustedExternalTenants")]
        public Output<ImmutableArray<string>> TrustedExternalTenants { get; private set; } = null!;

        /// <summary>
        /// The FQDN of the Azure Kusto Cluster.
        /// </summary>
        [Output("uri")]
        public Output<string> Uri { get; private set; } = null!;

        /// <summary>
        /// A `virtual_network_configuration` block as defined below.
        /// </summary>
        [Output("virtualNetworkConfiguration")]
        public Output<Outputs.ClusterVirtualNetworkConfiguration?> VirtualNetworkConfiguration { get; private set; } = null!;

        /// <summary>
        /// A list of Availability Zones in which the cluster instances should be created in. Changing this forces a new resource to be created.
        /// </summary>
        [Output("zones")]
        public Output<ImmutableArray<string>> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("azure:kusto/cluster:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("azure:kusto/cluster:Cluster", name, state, MakeResourceOptions(options, id))
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
        /// Specifies if the cluster's disks are encrypted.
        /// </summary>
        [Input("enableDiskEncryption")]
        public Input<bool>? EnableDiskEncryption { get; set; }

        /// <summary>
        /// Specifies if the purge operations are enabled.
        /// </summary>
        [Input("enablePurge")]
        public Input<bool>? EnablePurge { get; set; }

        /// <summary>
        /// Specifies if the streaming ingest is enabled.
        /// </summary>
        [Input("enableStreamingIngest")]
        public Input<bool>? EnableStreamingIngest { get; set; }

        /// <summary>
        /// A identity block.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ClusterIdentityArgs>? Identity { get; set; }

        [Input("languageExtensions")]
        private InputList<string>? _languageExtensions;

        /// <summary>
        /// An list of `language_extensions` to enable. Valid values are: `PYTHON` and `R`.
        /// </summary>
        public InputList<string> LanguageExtensions
        {
            get => _languageExtensions ?? (_languageExtensions = new InputList<string>());
            set => _languageExtensions = value;
        }

        /// <summary>
        /// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// An `optimized_auto_scale` block as defined below.
        /// </summary>
        [Input("optimizedAutoScale")]
        public Input<Inputs.ClusterOptimizedAutoScaleArgs>? OptimizedAutoScale { get; set; }

        /// <summary>
        /// Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `sku` block as defined below.
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.ClusterSkuArgs> Sku { get; set; } = null!;

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

        [Input("trustedExternalTenants")]
        private InputList<string>? _trustedExternalTenants;

        /// <summary>
        /// Specifies a list of tenant IDs that are trusted by the cluster.
        /// </summary>
        public InputList<string> TrustedExternalTenants
        {
            get => _trustedExternalTenants ?? (_trustedExternalTenants = new InputList<string>());
            set => _trustedExternalTenants = value;
        }

        /// <summary>
        /// A `virtual_network_configuration` block as defined below.
        /// </summary>
        [Input("virtualNetworkConfiguration")]
        public Input<Inputs.ClusterVirtualNetworkConfigurationArgs>? VirtualNetworkConfiguration { get; set; }

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// A list of Availability Zones in which the cluster instances should be created in. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public ClusterArgs()
        {
        }
    }

    public sealed class ClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Kusto Cluster URI to be used for data ingestion.
        /// </summary>
        [Input("dataIngestionUri")]
        public Input<string>? DataIngestionUri { get; set; }

        /// <summary>
        /// Specifies if the cluster's disks are encrypted.
        /// </summary>
        [Input("enableDiskEncryption")]
        public Input<bool>? EnableDiskEncryption { get; set; }

        /// <summary>
        /// Specifies if the purge operations are enabled.
        /// </summary>
        [Input("enablePurge")]
        public Input<bool>? EnablePurge { get; set; }

        /// <summary>
        /// Specifies if the streaming ingest is enabled.
        /// </summary>
        [Input("enableStreamingIngest")]
        public Input<bool>? EnableStreamingIngest { get; set; }

        /// <summary>
        /// A identity block.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ClusterIdentityGetArgs>? Identity { get; set; }

        [Input("languageExtensions")]
        private InputList<string>? _languageExtensions;

        /// <summary>
        /// An list of `language_extensions` to enable. Valid values are: `PYTHON` and `R`.
        /// </summary>
        public InputList<string> LanguageExtensions
        {
            get => _languageExtensions ?? (_languageExtensions = new InputList<string>());
            set => _languageExtensions = value;
        }

        /// <summary>
        /// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// An `optimized_auto_scale` block as defined below.
        /// </summary>
        [Input("optimizedAutoScale")]
        public Input<Inputs.ClusterOptimizedAutoScaleGetArgs>? OptimizedAutoScale { get; set; }

        /// <summary>
        /// Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// A `sku` block as defined below.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.ClusterSkuGetArgs>? Sku { get; set; }

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

        [Input("trustedExternalTenants")]
        private InputList<string>? _trustedExternalTenants;

        /// <summary>
        /// Specifies a list of tenant IDs that are trusted by the cluster.
        /// </summary>
        public InputList<string> TrustedExternalTenants
        {
            get => _trustedExternalTenants ?? (_trustedExternalTenants = new InputList<string>());
            set => _trustedExternalTenants = value;
        }

        /// <summary>
        /// The FQDN of the Azure Kusto Cluster.
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        /// <summary>
        /// A `virtual_network_configuration` block as defined below.
        /// </summary>
        [Input("virtualNetworkConfiguration")]
        public Input<Inputs.ClusterVirtualNetworkConfigurationGetArgs>? VirtualNetworkConfiguration { get; set; }

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// A list of Availability Zones in which the cluster instances should be created in. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public ClusterState()
        {
        }
    }
}

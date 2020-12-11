// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Hpc
{
    /// <summary>
    /// Manages a NFS Target within a HPC Cache.
    /// 
    /// &gt; **NOTE:**: By request of the service team the provider no longer automatically registering the `Microsoft.StorageCache` Resource Provider for this resource. To register it you can run `az provider register --namespace 'Microsoft.StorageCache'`.
    /// 
    /// ## Import
    /// 
    /// NFS Target within a HPC Cache can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:hpc/cacheNfsTarget:CacheNfsTarget example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StorageCache/caches/cache1/storageTargets/target1
    /// ```
    /// </summary>
    [AzureResourceType("azure:hpc/cacheNfsTarget:CacheNfsTarget")]
    public partial class CacheNfsTarget : Pulumi.CustomResource
    {
        /// <summary>
        /// The name HPC Cache, which the HPC Cache NFS Target will be added to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("cacheName")]
        public Output<string> CacheName { get; private set; } = null!;

        /// <summary>
        /// The name of the HPC Cache NFS Target. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Can be specified multiple times to define multiple `namespace_junction`. Each `namespace_juntion` block supports fields documented below.
        /// </summary>
        [Output("namespaceJunctions")]
        public Output<ImmutableArray<Outputs.CacheNfsTargetNamespaceJunction>> NamespaceJunctions { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which to create the HPC Cache NFS Target. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The IP address or fully qualified domain name (FQDN) of the HPC Cache NFS target. Changing this forces a new resource to be created.
        /// </summary>
        [Output("targetHostName")]
        public Output<string> TargetHostName { get; private set; } = null!;

        /// <summary>
        /// The type of usage of the HPC Cache NFS Target.
        /// </summary>
        [Output("usageModel")]
        public Output<string> UsageModel { get; private set; } = null!;


        /// <summary>
        /// Create a CacheNfsTarget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CacheNfsTarget(string name, CacheNfsTargetArgs args, CustomResourceOptions? options = null)
            : base("azure:hpc/cacheNfsTarget:CacheNfsTarget", name, args ?? new CacheNfsTargetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CacheNfsTarget(string name, Input<string> id, CacheNfsTargetState? state = null, CustomResourceOptions? options = null)
            : base("azure:hpc/cacheNfsTarget:CacheNfsTarget", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CacheNfsTarget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CacheNfsTarget Get(string name, Input<string> id, CacheNfsTargetState? state = null, CustomResourceOptions? options = null)
        {
            return new CacheNfsTarget(name, id, state, options);
        }
    }

    public sealed class CacheNfsTargetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name HPC Cache, which the HPC Cache NFS Target will be added to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("cacheName", required: true)]
        public Input<string> CacheName { get; set; } = null!;

        /// <summary>
        /// The name of the HPC Cache NFS Target. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namespaceJunctions", required: true)]
        private InputList<Inputs.CacheNfsTargetNamespaceJunctionArgs>? _namespaceJunctions;

        /// <summary>
        /// Can be specified multiple times to define multiple `namespace_junction`. Each `namespace_juntion` block supports fields documented below.
        /// </summary>
        public InputList<Inputs.CacheNfsTargetNamespaceJunctionArgs> NamespaceJunctions
        {
            get => _namespaceJunctions ?? (_namespaceJunctions = new InputList<Inputs.CacheNfsTargetNamespaceJunctionArgs>());
            set => _namespaceJunctions = value;
        }

        /// <summary>
        /// The name of the Resource Group in which to create the HPC Cache NFS Target. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The IP address or fully qualified domain name (FQDN) of the HPC Cache NFS target. Changing this forces a new resource to be created.
        /// </summary>
        [Input("targetHostName", required: true)]
        public Input<string> TargetHostName { get; set; } = null!;

        /// <summary>
        /// The type of usage of the HPC Cache NFS Target.
        /// </summary>
        [Input("usageModel", required: true)]
        public Input<string> UsageModel { get; set; } = null!;

        public CacheNfsTargetArgs()
        {
        }
    }

    public sealed class CacheNfsTargetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name HPC Cache, which the HPC Cache NFS Target will be added to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("cacheName")]
        public Input<string>? CacheName { get; set; }

        /// <summary>
        /// The name of the HPC Cache NFS Target. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namespaceJunctions")]
        private InputList<Inputs.CacheNfsTargetNamespaceJunctionGetArgs>? _namespaceJunctions;

        /// <summary>
        /// Can be specified multiple times to define multiple `namespace_junction`. Each `namespace_juntion` block supports fields documented below.
        /// </summary>
        public InputList<Inputs.CacheNfsTargetNamespaceJunctionGetArgs> NamespaceJunctions
        {
            get => _namespaceJunctions ?? (_namespaceJunctions = new InputList<Inputs.CacheNfsTargetNamespaceJunctionGetArgs>());
            set => _namespaceJunctions = value;
        }

        /// <summary>
        /// The name of the Resource Group in which to create the HPC Cache NFS Target. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The IP address or fully qualified domain name (FQDN) of the HPC Cache NFS target. Changing this forces a new resource to be created.
        /// </summary>
        [Input("targetHostName")]
        public Input<string>? TargetHostName { get; set; }

        /// <summary>
        /// The type of usage of the HPC Cache NFS Target.
        /// </summary>
        [Input("usageModel")]
        public Input<string>? UsageModel { get; set; }

        public CacheNfsTargetState()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Search
{
    public static class GetService
    {
        /// <summary>
        /// Manages a Search Service.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///         var exampleService = exampleResourceGroup.Name.Apply(name =&gt; Azure.Search.GetService.InvokeAsync(new Azure.Search.GetServiceArgs
        ///         {
        ///             Name = "example-search-service",
        ///             ResourceGroupName = name,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServiceResult> InvokeAsync(GetServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceResult>("azure:search/getService:getService", args ?? new GetServiceArgs(), options.WithVersion());

        public static Output<GetServiceResult> Apply(GetServiceApplyArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Name.Box(),
                args.ResourceGroupName.Box()
            ).Apply(a => {
                    var args = new GetServiceArgs();
                    a[0].Set(args, nameof(args.Name));
                    a[1].Set(args, nameof(args.ResourceGroupName));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetServiceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Name of the Search Service.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Search Service exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetServiceArgs()
        {
        }
    }

    public sealed class GetServiceApplyArgs
    {
        /// <summary>
        /// The Name of the Search Service.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Search Service exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetServiceApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServiceResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServiceIdentityResult> Identities;
        /// <summary>
        /// The name of this Query Key.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The number of partitions which have been created.
        /// </summary>
        public readonly int PartitionCount;
        /// <summary>
        /// The Primary Key used for Search Service Administration.
        /// </summary>
        public readonly string PrimaryKey;
        /// <summary>
        /// Whether or not public network access is enabled for this resource.
        /// </summary>
        public readonly bool PublicNetworkAccessEnabled;
        /// <summary>
        /// A `query_keys` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServiceQueryKeyResult> QueryKeys;
        /// <summary>
        /// The number of replica's which have been created.
        /// </summary>
        public readonly int ReplicaCount;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The Secondary Key used for Search Service Administration.
        /// </summary>
        public readonly string SecondaryKey;

        [OutputConstructor]
        private GetServiceResult(
            string id,

            ImmutableArray<Outputs.GetServiceIdentityResult> identities,

            string name,

            int partitionCount,

            string primaryKey,

            bool publicNetworkAccessEnabled,

            ImmutableArray<Outputs.GetServiceQueryKeyResult> queryKeys,

            int replicaCount,

            string resourceGroupName,

            string secondaryKey)
        {
            Id = id;
            Identities = identities;
            Name = name;
            PartitionCount = partitionCount;
            PrimaryKey = primaryKey;
            PublicNetworkAccessEnabled = publicNetworkAccessEnabled;
            QueryKeys = queryKeys;
            ReplicaCount = replicaCount;
            ResourceGroupName = resourceGroupName;
            SecondaryKey = secondaryKey;
        }
    }
}

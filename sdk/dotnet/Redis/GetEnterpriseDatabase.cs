// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Redis
{
    public static class GetEnterpriseDatabase
    {
        /// <summary>
        /// Use this data source to access information about an existing Redis Enterprise Database
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
        ///         var example = Output.Create(Azure.Redis.GetEnterpriseDatabase.InvokeAsync(new Azure.Redis.GetEnterpriseDatabaseArgs
        ///         {
        ///             Name = "default",
        ///             ResourceGroupName = azurerm_resource_group.Example.Name,
        ///             ClusterId = azurerm_redis_enterprise_cluster.Example.Id,
        ///         }));
        ///         this.RedisEnterpriseDatabasePrimaryKey = example.Apply(example =&gt; example.PrimaryAccessKey);
        ///         this.RedisEnterpriseDatabaseSecondaryKey = example.Apply(example =&gt; example.SecondaryAccessKey);
        ///     }
        /// 
        ///     [Output("redisEnterpriseDatabasePrimaryKey")]
        ///     public Output&lt;string&gt; RedisEnterpriseDatabasePrimaryKey { get; set; }
        ///     [Output("redisEnterpriseDatabaseSecondaryKey")]
        ///     public Output&lt;string&gt; RedisEnterpriseDatabaseSecondaryKey { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEnterpriseDatabaseResult> InvokeAsync(GetEnterpriseDatabaseArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEnterpriseDatabaseResult>("azure:redis/getEnterpriseDatabase:getEnterpriseDatabase", args ?? new GetEnterpriseDatabaseArgs(), options.WithVersion());

        public static Output<GetEnterpriseDatabaseResult> Apply(GetEnterpriseDatabaseApplyArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.ClusterId.Box(),
                args.Name.Box(),
                args.ResourceGroupName.Box()
            ).Apply(a => {
                    var args = new GetEnterpriseDatabaseArgs();
                    a[0].Set(args, nameof(args.ClusterId));
                    a[1].Set(args, nameof(args.Name));
                    a[2].Set(args, nameof(args.ResourceGroupName));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetEnterpriseDatabaseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The resource ID of Redis Enterprise Cluster which hosts the Redis Enterprise Database instance.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// The name of the Redis Enterprise Database.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group the Redis Enterprise Database instance is located in.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetEnterpriseDatabaseArgs()
        {
        }
    }

    public sealed class GetEnterpriseDatabaseApplyArgs
    {
        /// <summary>
        /// The resource ID of Redis Enterprise Cluster which hosts the Redis Enterprise Database instance.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The name of the Redis Enterprise Database.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group the Redis Enterprise Database instance is located in.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetEnterpriseDatabaseApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEnterpriseDatabaseResult
    {
        /// <summary>
        /// The Redis Enterprise Cluster ID that is hosting the Redis Enterprise Database.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Redis Enterprise Database name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Primary Access Key for the Redis Enterprise Database instance.
        /// </summary>
        public readonly string PrimaryAccessKey;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The Secondary Access Key for the Redis Enterprise Database instance.
        /// </summary>
        public readonly string SecondaryAccessKey;

        [OutputConstructor]
        private GetEnterpriseDatabaseResult(
            string clusterId,

            string id,

            string name,

            string primaryAccessKey,

            string resourceGroupName,

            string secondaryAccessKey)
        {
            ClusterId = clusterId;
            Id = id;
            Name = name;
            PrimaryAccessKey = primaryAccessKey;
            ResourceGroupName = resourceGroupName;
            SecondaryAccessKey = secondaryAccessKey;
        }
    }
}

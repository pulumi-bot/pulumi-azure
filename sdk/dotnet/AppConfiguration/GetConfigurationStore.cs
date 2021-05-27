// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppConfiguration
{
    public static class GetConfigurationStore
    {
        /// <summary>
        /// Use this data source to access information about an existing App Configuration.
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
        ///         var example = Output.Create(Azure.AppConfiguration.GetConfigurationStore.InvokeAsync(new Azure.AppConfiguration.GetConfigurationStoreArgs
        ///         {
        ///             Name = "existing",
        ///             ResourceGroupName = "existing",
        ///         }));
        ///         this.Id = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetConfigurationStoreResult> InvokeAsync(GetConfigurationStoreArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConfigurationStoreResult>("azure:appconfiguration/getConfigurationStore:getConfigurationStore", args ?? new GetConfigurationStoreArgs(), options.WithVersion());

        public static Output<GetConfigurationStoreResult> Apply(GetConfigurationStoreApplyArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Name.Box(),
                args.ResourceGroupName.Box()
            ).Apply(a => {
                    var args = new GetConfigurationStoreArgs();
                    a[0].Set(args, nameof(args.Name));
                    a[1].Set(args, nameof(args.ResourceGroupName));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetConfigurationStoreArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Name of this App Configuration.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the App Configuration exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetConfigurationStoreArgs()
        {
        }
    }

    public sealed class GetConfigurationStoreApplyArgs
    {
        /// <summary>
        /// The Name of this App Configuration.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the App Configuration exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetConfigurationStoreApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConfigurationStoreResult
    {
        /// <summary>
        /// The Endpoint used to access this App Configuration.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure Region where the App Configuration exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        /// <summary>
        /// A `primary_read_key` block as defined below containing the primary read access key.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetConfigurationStorePrimaryReadKeyResult> PrimaryReadKeys;
        /// <summary>
        /// A `primary_write_key` block as defined below containing the primary write access key.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetConfigurationStorePrimaryWriteKeyResult> PrimaryWriteKeys;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A `secondary_read_key` block as defined below containing the secondary read access key.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetConfigurationStoreSecondaryReadKeyResult> SecondaryReadKeys;
        /// <summary>
        /// A `secondary_write_key` block as defined below containing the secondary write access key.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetConfigurationStoreSecondaryWriteKeyResult> SecondaryWriteKeys;
        /// <summary>
        /// The name of the SKU used for this App Configuration.
        /// </summary>
        public readonly string Sku;
        /// <summary>
        /// A mapping of tags assigned to the App Configuration.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetConfigurationStoreResult(
            string endpoint,

            string id,

            string location,

            string name,

            ImmutableArray<Outputs.GetConfigurationStorePrimaryReadKeyResult> primaryReadKeys,

            ImmutableArray<Outputs.GetConfigurationStorePrimaryWriteKeyResult> primaryWriteKeys,

            string resourceGroupName,

            ImmutableArray<Outputs.GetConfigurationStoreSecondaryReadKeyResult> secondaryReadKeys,

            ImmutableArray<Outputs.GetConfigurationStoreSecondaryWriteKeyResult> secondaryWriteKeys,

            string sku,

            ImmutableDictionary<string, string> tags)
        {
            Endpoint = endpoint;
            Id = id;
            Location = location;
            Name = name;
            PrimaryReadKeys = primaryReadKeys;
            PrimaryWriteKeys = primaryWriteKeys;
            ResourceGroupName = resourceGroupName;
            SecondaryReadKeys = secondaryReadKeys;
            SecondaryWriteKeys = secondaryWriteKeys;
            Sku = sku;
            Tags = tags;
        }
    }
}

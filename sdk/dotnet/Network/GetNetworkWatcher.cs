// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    public static class GetNetworkWatcher
    {
        /// <summary>
        /// Use this data source to access information about an existing Network Watcher.
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
        ///         var example = Output.Create(Azure.Network.GetNetworkWatcher.InvokeAsync(new Azure.Network.GetNetworkWatcherArgs
        ///         {
        ///             Name = azurerm_network_watcher.Example.Name,
        ///             ResourceGroupName = azurerm_resource_group.Example.Name,
        ///         }));
        ///         this.NetworkWatcherId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("networkWatcherId")]
        ///     public Output&lt;string&gt; NetworkWatcherId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNetworkWatcherResult> InvokeAsync(GetNetworkWatcherArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworkWatcherResult>("azure:network/getNetworkWatcher:getNetworkWatcher", args ?? new GetNetworkWatcherArgs(), options.WithVersion());
    }


    public sealed class GetNetworkWatcherArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the Name of the Network Watcher.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the Name of the Resource Group within which the Network Watcher exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNetworkWatcherArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNetworkWatcherResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The supported Azure location where the resource exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetNetworkWatcherResult(
            string id,

            string location,

            string name,

            string resourceGroupName,

            ImmutableDictionary<string, string> tags)
        {
            Id = id;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
        }
    }
}

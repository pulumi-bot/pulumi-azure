// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    public static class GetVirtualNetwork
    {
        /// <summary>
        /// Use this data source to access information about an existing Virtual Network.
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
        ///         var example = Output.Create(Azure.Network.GetVirtualNetwork.InvokeAsync(new Azure.Network.GetVirtualNetworkArgs
        ///         {
        ///             Name = "production",
        ///             ResourceGroupName = "networking",
        ///         }));
        ///         this.VirtualNetworkId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("virtualNetworkId")]
        ///     public Output&lt;string&gt; VirtualNetworkId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVirtualNetworkResult> InvokeAsync(GetVirtualNetworkArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVirtualNetworkResult>("azure:network/getVirtualNetwork:getVirtualNetwork", args ?? new GetVirtualNetworkArgs(), options.WithVersion());

        public static Output<GetVirtualNetworkResult> Apply(GetVirtualNetworkApplyArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Name.Box(),
                args.ResourceGroupName.Box()
            ).Apply(a => {
                    var args = new GetVirtualNetworkArgs();
                    a[0].Set(args, nameof(args.Name));
                    a[1].Set(args, nameof(args.ResourceGroupName));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetVirtualNetworkArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Virtual Network.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group the Virtual Network is located in.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetVirtualNetworkArgs()
        {
        }
    }

    public sealed class GetVirtualNetworkApplyArgs
    {
        /// <summary>
        /// Specifies the name of the Virtual Network.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group the Virtual Network is located in.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetVirtualNetworkApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVirtualNetworkResult
    {
        /// <summary>
        /// The list of address spaces used by the virtual network.
        /// </summary>
        public readonly ImmutableArray<string> AddressSpaces;
        /// <summary>
        /// The list of DNS servers used by the virtual network.
        /// </summary>
        public readonly ImmutableArray<string> DnsServers;
        /// <summary>
        /// The GUID of the virtual network.
        /// </summary>
        public readonly string Guid;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Location of the virtual network.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The list of name of the subnets that are attached to this virtual network.
        /// </summary>
        public readonly ImmutableArray<string> Subnets;
        /// <summary>
        /// A mapping of name - virtual network id of the virtual network peerings.
        /// </summary>
        public readonly ImmutableDictionary<string, string> VnetPeerings;

        [OutputConstructor]
        private GetVirtualNetworkResult(
            ImmutableArray<string> addressSpaces,

            ImmutableArray<string> dnsServers,

            string guid,

            string id,

            string location,

            string name,

            string resourceGroupName,

            ImmutableArray<string> subnets,

            ImmutableDictionary<string, string> vnetPeerings)
        {
            AddressSpaces = addressSpaces;
            DnsServers = dnsServers;
            Guid = guid;
            Id = id;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Subnets = subnets;
            VnetPeerings = vnetPeerings;
        }
    }
}

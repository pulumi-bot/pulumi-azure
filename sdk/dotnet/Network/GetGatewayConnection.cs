// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    public static class GetGatewayConnection
    {
        /// <summary>
        /// Use this data source to access information about an existing Virtual Network Gateway Connection.
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
        ///         var example = Output.Create(Azure.Network.GetGatewayConnection.InvokeAsync(new Azure.Network.GetGatewayConnectionArgs
        ///         {
        ///             Name = "production",
        ///             ResourceGroupName = "networking",
        ///         }));
        ///         this.VirtualNetworkGatewayConnectionId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("virtualNetworkGatewayConnectionId")]
        ///     public Output&lt;string&gt; VirtualNetworkGatewayConnectionId { get; set; }
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGatewayConnectionResult> InvokeAsync(GetGatewayConnectionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGatewayConnectionResult>("azure:network/getGatewayConnection:getGatewayConnection", args ?? new GetGatewayConnectionArgs(), options.WithVersion());
    }


    public sealed class GetGatewayConnectionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Virtual Network Gateway Connection.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group the Virtual Network Gateway Connection is located in.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetGatewayConnectionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGatewayConnectionResult
    {
        /// <summary>
        /// The authorization key associated with the
        /// Express Route Circuit. This field is present only if the type is an
        /// ExpressRoute connection.
        /// </summary>
        public readonly string AuthorizationKey;
        public readonly string ConnectionProtocol;
        public readonly int EgressBytesTransferred;
        /// <summary>
        /// If `true`, BGP (Border Gateway Protocol) is enabled
        /// for this connection.
        /// </summary>
        public readonly bool EnableBgp;
        /// <summary>
        /// The ID of the Express Route Circuit
        /// (i.e. when `type` is `ExpressRoute`).
        /// </summary>
        public readonly string ExpressRouteCircuitId;
        /// <summary>
        /// If `true`, data packets will bypass ExpressRoute Gateway for data forwarding. This is only valid for ExpressRoute connections.
        /// </summary>
        public readonly bool ExpressRouteGatewayBypass;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int IngressBytesTransferred;
        public readonly ImmutableArray<Outputs.GetGatewayConnectionIpsecPolicyResult> IpsecPolicies;
        /// <summary>
        /// The ID of the local network gateway
        /// when a Site-to-Site connection (i.e. when `type` is `IPsec`).
        /// </summary>
        public readonly string LocalNetworkGatewayId;
        /// <summary>
        /// The location/region where the connection is
        /// located.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        /// <summary>
        /// The ID of the peer virtual
        /// network gateway when a VNet-to-VNet connection (i.e. when `type`
        /// is `Vnet2Vnet`).
        /// </summary>
        public readonly string PeerVirtualNetworkGatewayId;
        public readonly string ResourceGroupName;
        public readonly string ResourceGuid;
        /// <summary>
        /// The routing weight.
        /// </summary>
        public readonly int RoutingWeight;
        /// <summary>
        /// The shared IPSec key. 
        /// </summary>
        public readonly string SharedKey;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The type of connection. Valid options are `IPsec`
        /// (Site-to-Site), `ExpressRoute` (ExpressRoute), and `Vnet2Vnet` (VNet-to-VNet).
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// If `true`, policy-based traffic
        /// selectors are enabled for this connection. Enabling policy-based traffic
        /// selectors requires an `ipsec_policy` block.
        /// </summary>
        public readonly bool UsePolicyBasedTrafficSelectors;
        /// <summary>
        /// The ID of the Virtual Network Gateway
        /// in which the connection is created.
        /// </summary>
        public readonly string VirtualNetworkGatewayId;

        [OutputConstructor]
        private GetGatewayConnectionResult(
            string authorizationKey,

            string connectionProtocol,

            int egressBytesTransferred,

            bool enableBgp,

            string expressRouteCircuitId,

            bool expressRouteGatewayBypass,

            string id,

            int ingressBytesTransferred,

            ImmutableArray<Outputs.GetGatewayConnectionIpsecPolicyResult> ipsecPolicies,

            string localNetworkGatewayId,

            string location,

            string name,

            string peerVirtualNetworkGatewayId,

            string resourceGroupName,

            string resourceGuid,

            int routingWeight,

            string sharedKey,

            ImmutableDictionary<string, string> tags,

            string type,

            bool usePolicyBasedTrafficSelectors,

            string virtualNetworkGatewayId)
        {
            AuthorizationKey = authorizationKey;
            ConnectionProtocol = connectionProtocol;
            EgressBytesTransferred = egressBytesTransferred;
            EnableBgp = enableBgp;
            ExpressRouteCircuitId = expressRouteCircuitId;
            ExpressRouteGatewayBypass = expressRouteGatewayBypass;
            Id = id;
            IngressBytesTransferred = ingressBytesTransferred;
            IpsecPolicies = ipsecPolicies;
            LocalNetworkGatewayId = localNetworkGatewayId;
            Location = location;
            Name = name;
            PeerVirtualNetworkGatewayId = peerVirtualNetworkGatewayId;
            ResourceGroupName = resourceGroupName;
            ResourceGuid = resourceGuid;
            RoutingWeight = routingWeight;
            SharedKey = sharedKey;
            Tags = tags;
            Type = type;
            UsePolicyBasedTrafficSelectors = usePolicyBasedTrafficSelectors;
            VirtualNetworkGatewayId = virtualNetworkGatewayId;
        }
    }
}

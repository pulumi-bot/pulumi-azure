// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    /// <summary>
    /// Manages a connection in an existing Virtual Network Gateway.
    /// 
    /// ## Example Usage
    /// ### Site-to-Site connection
    /// 
    /// The following example shows a connection between an Azure virtual network
    /// and an on-premises VPN device and network.
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
    ///             Location = "West US",
    ///         });
    ///         var exampleVirtualNetwork = new Azure.Network.VirtualNetwork("exampleVirtualNetwork", new Azure.Network.VirtualNetworkArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AddressSpaces = 
    ///             {
    ///                 "10.0.0.0/16",
    ///             },
    ///         });
    ///         var exampleSubnet = new Azure.Network.Subnet("exampleSubnet", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualNetworkName = exampleVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.0.1.0/24",
    ///             },
    ///         });
    ///         var onpremiseLocalNetworkGateway = new Azure.Network.LocalNetworkGateway("onpremiseLocalNetworkGateway", new Azure.Network.LocalNetworkGatewayArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             GatewayAddress = "168.62.225.23",
    ///             AddressSpaces = 
    ///             {
    ///                 "10.1.1.0/24",
    ///             },
    ///         });
    ///         var examplePublicIp = new Azure.Network.PublicIp("examplePublicIp", new Azure.Network.PublicIpArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AllocationMethod = "Dynamic",
    ///         });
    ///         var exampleVirtualNetworkGateway = new Azure.Network.VirtualNetworkGateway("exampleVirtualNetworkGateway", new Azure.Network.VirtualNetworkGatewayArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Type = "Vpn",
    ///             VpnType = "RouteBased",
    ///             ActiveActive = false,
    ///             EnableBgp = false,
    ///             Sku = "Basic",
    ///             IpConfigurations = 
    ///             {
    ///                 new Azure.Network.Inputs.VirtualNetworkGatewayIpConfigurationArgs
    ///                 {
    ///                     PublicIpAddressId = examplePublicIp.Id,
    ///                     PrivateIpAddressAllocation = "Dynamic",
    ///                     SubnetId = exampleSubnet.Id,
    ///                 },
    ///             },
    ///         });
    ///         var onpremiseVirtualNetworkGatewayConnection = new Azure.Network.VirtualNetworkGatewayConnection("onpremiseVirtualNetworkGatewayConnection", new Azure.Network.VirtualNetworkGatewayConnectionArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Type = "IPsec",
    ///             VirtualNetworkGatewayId = exampleVirtualNetworkGateway.Id,
    ///             LocalNetworkGatewayId = onpremiseLocalNetworkGateway.Id,
    ///             SharedKey = "4-v3ry-53cr37-1p53c-5h4r3d-k3y",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### VNet-to-VNet connection
    /// 
    /// The following example shows a connection between two Azure virtual network
    /// in different locations/regions.
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var usResourceGroup = new Azure.Core.ResourceGroup("usResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "East US",
    ///         });
    ///         var usVirtualNetwork = new Azure.Network.VirtualNetwork("usVirtualNetwork", new Azure.Network.VirtualNetworkArgs
    ///         {
    ///             Location = usResourceGroup.Location,
    ///             ResourceGroupName = usResourceGroup.Name,
    ///             AddressSpaces = 
    ///             {
    ///                 "10.0.0.0/16",
    ///             },
    ///         });
    ///         var usGateway = new Azure.Network.Subnet("usGateway", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = usResourceGroup.Name,
    ///             VirtualNetworkName = usVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.0.1.0/24",
    ///             },
    ///         });
    ///         var usPublicIp = new Azure.Network.PublicIp("usPublicIp", new Azure.Network.PublicIpArgs
    ///         {
    ///             Location = usResourceGroup.Location,
    ///             ResourceGroupName = usResourceGroup.Name,
    ///             AllocationMethod = "Dynamic",
    ///         });
    ///         var usVirtualNetworkGateway = new Azure.Network.VirtualNetworkGateway("usVirtualNetworkGateway", new Azure.Network.VirtualNetworkGatewayArgs
    ///         {
    ///             Location = usResourceGroup.Location,
    ///             ResourceGroupName = usResourceGroup.Name,
    ///             Type = "Vpn",
    ///             VpnType = "RouteBased",
    ///             Sku = "Basic",
    ///             IpConfigurations = 
    ///             {
    ///                 new Azure.Network.Inputs.VirtualNetworkGatewayIpConfigurationArgs
    ///                 {
    ///                     PublicIpAddressId = usPublicIp.Id,
    ///                     PrivateIpAddressAllocation = "Dynamic",
    ///                     SubnetId = usGateway.Id,
    ///                 },
    ///             },
    ///         });
    ///         var europeResourceGroup = new Azure.Core.ResourceGroup("europeResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var europeVirtualNetwork = new Azure.Network.VirtualNetwork("europeVirtualNetwork", new Azure.Network.VirtualNetworkArgs
    ///         {
    ///             Location = europeResourceGroup.Location,
    ///             ResourceGroupName = europeResourceGroup.Name,
    ///             AddressSpaces = 
    ///             {
    ///                 "10.1.0.0/16",
    ///             },
    ///         });
    ///         var europeGateway = new Azure.Network.Subnet("europeGateway", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = europeResourceGroup.Name,
    ///             VirtualNetworkName = europeVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.1.1.0/24",
    ///             },
    ///         });
    ///         var europePublicIp = new Azure.Network.PublicIp("europePublicIp", new Azure.Network.PublicIpArgs
    ///         {
    ///             Location = europeResourceGroup.Location,
    ///             ResourceGroupName = europeResourceGroup.Name,
    ///             AllocationMethod = "Dynamic",
    ///         });
    ///         var europeVirtualNetworkGateway = new Azure.Network.VirtualNetworkGateway("europeVirtualNetworkGateway", new Azure.Network.VirtualNetworkGatewayArgs
    ///         {
    ///             Location = europeResourceGroup.Location,
    ///             ResourceGroupName = europeResourceGroup.Name,
    ///             Type = "Vpn",
    ///             VpnType = "RouteBased",
    ///             Sku = "Basic",
    ///             IpConfigurations = 
    ///             {
    ///                 new Azure.Network.Inputs.VirtualNetworkGatewayIpConfigurationArgs
    ///                 {
    ///                     PublicIpAddressId = europePublicIp.Id,
    ///                     PrivateIpAddressAllocation = "Dynamic",
    ///                     SubnetId = europeGateway.Id,
    ///                 },
    ///             },
    ///         });
    ///         var usToEurope = new Azure.Network.VirtualNetworkGatewayConnection("usToEurope", new Azure.Network.VirtualNetworkGatewayConnectionArgs
    ///         {
    ///             Location = usResourceGroup.Location,
    ///             ResourceGroupName = usResourceGroup.Name,
    ///             Type = "Vnet2Vnet",
    ///             VirtualNetworkGatewayId = usVirtualNetworkGateway.Id,
    ///             PeerVirtualNetworkGatewayId = europeVirtualNetworkGateway.Id,
    ///             SharedKey = "4-v3ry-53cr37-1p53c-5h4r3d-k3y",
    ///         });
    ///         var europeToUs = new Azure.Network.VirtualNetworkGatewayConnection("europeToUs", new Azure.Network.VirtualNetworkGatewayConnectionArgs
    ///         {
    ///             Location = europeResourceGroup.Location,
    ///             ResourceGroupName = europeResourceGroup.Name,
    ///             Type = "Vnet2Vnet",
    ///             VirtualNetworkGatewayId = europeVirtualNetworkGateway.Id,
    ///             PeerVirtualNetworkGatewayId = usVirtualNetworkGateway.Id,
    ///             SharedKey = "4-v3ry-53cr37-1p53c-5h4r3d-k3y",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Virtual Network Gateway Connections can be imported using their `resource id`, e.g.
    /// </summary>
    public partial class VirtualNetworkGatewayConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// The authorization key associated with the
        /// Express Route Circuit. This field is required only if the type is an
        /// ExpressRoute connection.
        /// </summary>
        [Output("authorizationKey")]
        public Output<string?> AuthorizationKey { get; private set; } = null!;

        /// <summary>
        /// The IKE protocol version to use. Possible
        /// values are `IKEv1` and `IKEv2`. Defaults to `IKEv2`.
        /// Changing this value will force a resource to be created.
        /// &gt; **Note**: Only valid for `IPSec` connections on virtual network gateways with SKU `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw1AZ`, `VpnGw2AZ` or `VpnGw3AZ`.
        /// </summary>
        [Output("connectionProtocol")]
        public Output<string> ConnectionProtocol { get; private set; } = null!;

        /// <summary>
        /// If `true`, BGP (Border Gateway Protocol) is enabled
        /// for this connection. Defaults to `false`.
        /// </summary>
        [Output("enableBgp")]
        public Output<bool> EnableBgp { get; private set; } = null!;

        /// <summary>
        /// The ID of the Express Route Circuit
        /// when creating an ExpressRoute connection (i.e. when `type` is `ExpressRoute`).
        /// The Express Route Circuit can be in the same or in a different subscription.
        /// </summary>
        [Output("expressRouteCircuitId")]
        public Output<string?> ExpressRouteCircuitId { get; private set; } = null!;

        /// <summary>
        /// If `true`, data packets will bypass ExpressRoute Gateway for data forwarding This is only valid for ExpressRoute connections.
        /// </summary>
        [Output("expressRouteGatewayBypass")]
        public Output<bool> ExpressRouteGatewayBypass { get; private set; } = null!;

        /// <summary>
        /// A `ipsec_policy` block which is documented below.
        /// Only a single policy can be defined for a connection. For details on
        /// custom policies refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-ipsecikepolicy-rm-powershell).
        /// </summary>
        [Output("ipsecPolicy")]
        public Output<Outputs.VirtualNetworkGatewayConnectionIpsecPolicy?> IpsecPolicy { get; private set; } = null!;

        /// <summary>
        /// The ID of the local network gateway
        /// when creating Site-to-Site connection (i.e. when `type` is `IPsec`).
        /// </summary>
        [Output("localNetworkGatewayId")]
        public Output<string?> LocalNetworkGatewayId { get; private set; } = null!;

        /// <summary>
        /// The location/region where the connection is
        /// located. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the connection. Changing the name forces a
        /// new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the peer virtual
        /// network gateway when creating a VNet-to-VNet connection (i.e. when `type`
        /// is `Vnet2Vnet`). The peer Virtual Network Gateway can be in the same or
        /// in a different subscription.
        /// </summary>
        [Output("peerVirtualNetworkGatewayId")]
        public Output<string?> PeerVirtualNetworkGatewayId { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to
        /// create the connection Changing the name forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The routing weight. Defaults to `10`.
        /// </summary>
        [Output("routingWeight")]
        public Output<int> RoutingWeight { get; private set; } = null!;

        /// <summary>
        /// The shared IPSec key. A key could be provided if a
        /// Site-to-Site, VNet-to-VNet or ExpressRoute connection is created.
        /// </summary>
        [Output("sharedKey")]
        public Output<string?> SharedKey { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A `traffic_selector_policy` which allows to specify traffic selector policy proposal to be used in a virtual network gateway connection.
        /// Only one block can be defined for a connection.
        /// For details about traffic selectors refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-connect-multiple-policybased-rm-ps).
        /// </summary>
        [Output("trafficSelectorPolicy")]
        public Output<Outputs.VirtualNetworkGatewayConnectionTrafficSelectorPolicy?> TrafficSelectorPolicy { get; private set; } = null!;

        /// <summary>
        /// The type of connection. Valid options are `IPsec`
        /// (Site-to-Site), `ExpressRoute` (ExpressRoute), and `Vnet2Vnet` (VNet-to-VNet).
        /// Each connection type requires different mandatory arguments (refer to the
        /// examples above). Changing the connection type will force a new connection
        /// to be created.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// If `true`, policy-based traffic
        /// selectors are enabled for this connection. Enabling policy-based traffic
        /// selectors requires an `ipsec_policy` block. Defaults to `false`.
        /// </summary>
        [Output("usePolicyBasedTrafficSelectors")]
        public Output<bool> UsePolicyBasedTrafficSelectors { get; private set; } = null!;

        /// <summary>
        /// The ID of the Virtual Network Gateway
        /// in which the connection will be created. Changing the gateway forces a new
        /// resource to be created.
        /// </summary>
        [Output("virtualNetworkGatewayId")]
        public Output<string> VirtualNetworkGatewayId { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualNetworkGatewayConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualNetworkGatewayConnection(string name, VirtualNetworkGatewayConnectionArgs args, CustomResourceOptions? options = null)
            : base("azure:network/virtualNetworkGatewayConnection:VirtualNetworkGatewayConnection", name, args ?? new VirtualNetworkGatewayConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualNetworkGatewayConnection(string name, Input<string> id, VirtualNetworkGatewayConnectionState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/virtualNetworkGatewayConnection:VirtualNetworkGatewayConnection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualNetworkGatewayConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualNetworkGatewayConnection Get(string name, Input<string> id, VirtualNetworkGatewayConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualNetworkGatewayConnection(name, id, state, options);
        }
    }

    public sealed class VirtualNetworkGatewayConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authorization key associated with the
        /// Express Route Circuit. This field is required only if the type is an
        /// ExpressRoute connection.
        /// </summary>
        [Input("authorizationKey")]
        public Input<string>? AuthorizationKey { get; set; }

        /// <summary>
        /// The IKE protocol version to use. Possible
        /// values are `IKEv1` and `IKEv2`. Defaults to `IKEv2`.
        /// Changing this value will force a resource to be created.
        /// &gt; **Note**: Only valid for `IPSec` connections on virtual network gateways with SKU `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw1AZ`, `VpnGw2AZ` or `VpnGw3AZ`.
        /// </summary>
        [Input("connectionProtocol")]
        public Input<string>? ConnectionProtocol { get; set; }

        /// <summary>
        /// If `true`, BGP (Border Gateway Protocol) is enabled
        /// for this connection. Defaults to `false`.
        /// </summary>
        [Input("enableBgp")]
        public Input<bool>? EnableBgp { get; set; }

        /// <summary>
        /// The ID of the Express Route Circuit
        /// when creating an ExpressRoute connection (i.e. when `type` is `ExpressRoute`).
        /// The Express Route Circuit can be in the same or in a different subscription.
        /// </summary>
        [Input("expressRouteCircuitId")]
        public Input<string>? ExpressRouteCircuitId { get; set; }

        /// <summary>
        /// If `true`, data packets will bypass ExpressRoute Gateway for data forwarding This is only valid for ExpressRoute connections.
        /// </summary>
        [Input("expressRouteGatewayBypass")]
        public Input<bool>? ExpressRouteGatewayBypass { get; set; }

        /// <summary>
        /// A `ipsec_policy` block which is documented below.
        /// Only a single policy can be defined for a connection. For details on
        /// custom policies refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-ipsecikepolicy-rm-powershell).
        /// </summary>
        [Input("ipsecPolicy")]
        public Input<Inputs.VirtualNetworkGatewayConnectionIpsecPolicyArgs>? IpsecPolicy { get; set; }

        /// <summary>
        /// The ID of the local network gateway
        /// when creating Site-to-Site connection (i.e. when `type` is `IPsec`).
        /// </summary>
        [Input("localNetworkGatewayId")]
        public Input<string>? LocalNetworkGatewayId { get; set; }

        /// <summary>
        /// The location/region where the connection is
        /// located. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the connection. Changing the name forces a
        /// new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the peer virtual
        /// network gateway when creating a VNet-to-VNet connection (i.e. when `type`
        /// is `Vnet2Vnet`). The peer Virtual Network Gateway can be in the same or
        /// in a different subscription.
        /// </summary>
        [Input("peerVirtualNetworkGatewayId")]
        public Input<string>? PeerVirtualNetworkGatewayId { get; set; }

        /// <summary>
        /// The name of the resource group in which to
        /// create the connection Changing the name forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The routing weight. Defaults to `10`.
        /// </summary>
        [Input("routingWeight")]
        public Input<int>? RoutingWeight { get; set; }

        /// <summary>
        /// The shared IPSec key. A key could be provided if a
        /// Site-to-Site, VNet-to-VNet or ExpressRoute connection is created.
        /// </summary>
        [Input("sharedKey")]
        public Input<string>? SharedKey { get; set; }

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

        /// <summary>
        /// A `traffic_selector_policy` which allows to specify traffic selector policy proposal to be used in a virtual network gateway connection.
        /// Only one block can be defined for a connection.
        /// For details about traffic selectors refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-connect-multiple-policybased-rm-ps).
        /// </summary>
        [Input("trafficSelectorPolicy")]
        public Input<Inputs.VirtualNetworkGatewayConnectionTrafficSelectorPolicyArgs>? TrafficSelectorPolicy { get; set; }

        /// <summary>
        /// The type of connection. Valid options are `IPsec`
        /// (Site-to-Site), `ExpressRoute` (ExpressRoute), and `Vnet2Vnet` (VNet-to-VNet).
        /// Each connection type requires different mandatory arguments (refer to the
        /// examples above). Changing the connection type will force a new connection
        /// to be created.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// If `true`, policy-based traffic
        /// selectors are enabled for this connection. Enabling policy-based traffic
        /// selectors requires an `ipsec_policy` block. Defaults to `false`.
        /// </summary>
        [Input("usePolicyBasedTrafficSelectors")]
        public Input<bool>? UsePolicyBasedTrafficSelectors { get; set; }

        /// <summary>
        /// The ID of the Virtual Network Gateway
        /// in which the connection will be created. Changing the gateway forces a new
        /// resource to be created.
        /// </summary>
        [Input("virtualNetworkGatewayId", required: true)]
        public Input<string> VirtualNetworkGatewayId { get; set; } = null!;

        public VirtualNetworkGatewayConnectionArgs()
        {
        }
    }

    public sealed class VirtualNetworkGatewayConnectionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authorization key associated with the
        /// Express Route Circuit. This field is required only if the type is an
        /// ExpressRoute connection.
        /// </summary>
        [Input("authorizationKey")]
        public Input<string>? AuthorizationKey { get; set; }

        /// <summary>
        /// The IKE protocol version to use. Possible
        /// values are `IKEv1` and `IKEv2`. Defaults to `IKEv2`.
        /// Changing this value will force a resource to be created.
        /// &gt; **Note**: Only valid for `IPSec` connections on virtual network gateways with SKU `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw1AZ`, `VpnGw2AZ` or `VpnGw3AZ`.
        /// </summary>
        [Input("connectionProtocol")]
        public Input<string>? ConnectionProtocol { get; set; }

        /// <summary>
        /// If `true`, BGP (Border Gateway Protocol) is enabled
        /// for this connection. Defaults to `false`.
        /// </summary>
        [Input("enableBgp")]
        public Input<bool>? EnableBgp { get; set; }

        /// <summary>
        /// The ID of the Express Route Circuit
        /// when creating an ExpressRoute connection (i.e. when `type` is `ExpressRoute`).
        /// The Express Route Circuit can be in the same or in a different subscription.
        /// </summary>
        [Input("expressRouteCircuitId")]
        public Input<string>? ExpressRouteCircuitId { get; set; }

        /// <summary>
        /// If `true`, data packets will bypass ExpressRoute Gateway for data forwarding This is only valid for ExpressRoute connections.
        /// </summary>
        [Input("expressRouteGatewayBypass")]
        public Input<bool>? ExpressRouteGatewayBypass { get; set; }

        /// <summary>
        /// A `ipsec_policy` block which is documented below.
        /// Only a single policy can be defined for a connection. For details on
        /// custom policies refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-ipsecikepolicy-rm-powershell).
        /// </summary>
        [Input("ipsecPolicy")]
        public Input<Inputs.VirtualNetworkGatewayConnectionIpsecPolicyGetArgs>? IpsecPolicy { get; set; }

        /// <summary>
        /// The ID of the local network gateway
        /// when creating Site-to-Site connection (i.e. when `type` is `IPsec`).
        /// </summary>
        [Input("localNetworkGatewayId")]
        public Input<string>? LocalNetworkGatewayId { get; set; }

        /// <summary>
        /// The location/region where the connection is
        /// located. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the connection. Changing the name forces a
        /// new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the peer virtual
        /// network gateway when creating a VNet-to-VNet connection (i.e. when `type`
        /// is `Vnet2Vnet`). The peer Virtual Network Gateway can be in the same or
        /// in a different subscription.
        /// </summary>
        [Input("peerVirtualNetworkGatewayId")]
        public Input<string>? PeerVirtualNetworkGatewayId { get; set; }

        /// <summary>
        /// The name of the resource group in which to
        /// create the connection Changing the name forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The routing weight. Defaults to `10`.
        /// </summary>
        [Input("routingWeight")]
        public Input<int>? RoutingWeight { get; set; }

        /// <summary>
        /// The shared IPSec key. A key could be provided if a
        /// Site-to-Site, VNet-to-VNet or ExpressRoute connection is created.
        /// </summary>
        [Input("sharedKey")]
        public Input<string>? SharedKey { get; set; }

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

        /// <summary>
        /// A `traffic_selector_policy` which allows to specify traffic selector policy proposal to be used in a virtual network gateway connection.
        /// Only one block can be defined for a connection.
        /// For details about traffic selectors refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-connect-multiple-policybased-rm-ps).
        /// </summary>
        [Input("trafficSelectorPolicy")]
        public Input<Inputs.VirtualNetworkGatewayConnectionTrafficSelectorPolicyGetArgs>? TrafficSelectorPolicy { get; set; }

        /// <summary>
        /// The type of connection. Valid options are `IPsec`
        /// (Site-to-Site), `ExpressRoute` (ExpressRoute), and `Vnet2Vnet` (VNet-to-VNet).
        /// Each connection type requires different mandatory arguments (refer to the
        /// examples above). Changing the connection type will force a new connection
        /// to be created.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// If `true`, policy-based traffic
        /// selectors are enabled for this connection. Enabling policy-based traffic
        /// selectors requires an `ipsec_policy` block. Defaults to `false`.
        /// </summary>
        [Input("usePolicyBasedTrafficSelectors")]
        public Input<bool>? UsePolicyBasedTrafficSelectors { get; set; }

        /// <summary>
        /// The ID of the Virtual Network Gateway
        /// in which the connection will be created. Changing the gateway forces a new
        /// resource to be created.
        /// </summary>
        [Input("virtualNetworkGatewayId")]
        public Input<string>? VirtualNetworkGatewayId { get; set; }

        public VirtualNetworkGatewayConnectionState()
        {
        }
    }
}

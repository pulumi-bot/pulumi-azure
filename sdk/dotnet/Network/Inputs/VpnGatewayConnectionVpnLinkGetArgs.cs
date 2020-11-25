// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class VpnGatewayConnectionVpnLinkGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The expected connection bandwidth in MBPS. Defaults to `10`.
        /// </summary>
        [Input("bandwidthMbps")]
        public Input<int>? BandwidthMbps { get; set; }

        /// <summary>
        /// Should the BGP be enabled? Defaults to `false`. Changing this forces a new VPN Gateway Connection to be created.
        /// </summary>
        [Input("bgpEnabled")]
        public Input<bool>? BgpEnabled { get; set; }

        [Input("ipsecPolicies")]
        private InputList<Inputs.VpnGatewayConnectionVpnLinkIpsecPolicyGetArgs>? _ipsecPolicies;

        /// <summary>
        /// One or more `ipsec_policy` blocks as defined above.
        /// </summary>
        public InputList<Inputs.VpnGatewayConnectionVpnLinkIpsecPolicyGetArgs> IpsecPolicies
        {
            get => _ipsecPolicies ?? (_ipsecPolicies = new InputList<Inputs.VpnGatewayConnectionVpnLinkIpsecPolicyGetArgs>());
            set => _ipsecPolicies = value;
        }

        /// <summary>
        /// Whether to use local azure ip to initiate connection? Defaults to `false`.
        /// </summary>
        [Input("localAzureIpAddressEnabled")]
        public Input<bool>? LocalAzureIpAddressEnabled { get; set; }

        /// <summary>
        /// The name which should be used for this VPN Link Connection.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Whether to enable policy-based traffic selectors? Defaults to `false`.
        /// </summary>
        [Input("policyBasedTrafficSelectorEnabled")]
        public Input<bool>? PolicyBasedTrafficSelectorEnabled { get; set; }

        /// <summary>
        /// The protocol used for this VPN Link Connection. Possible values are `IKEv1` and `IKEv2`. Defaults to `IKEv2`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Should the rate limit be enabled? Defaults to `false`.
        /// </summary>
        [Input("ratelimitEnabled")]
        public Input<bool>? RatelimitEnabled { get; set; }

        /// <summary>
        /// Routing weight for this VPN Link Connection. Defaults to `0`.
        /// </summary>
        [Input("routeWeight")]
        public Input<int>? RouteWeight { get; set; }

        /// <summary>
        /// SharedKey for this VPN Link Connection.
        /// </summary>
        [Input("sharedKey")]
        public Input<string>? SharedKey { get; set; }

        /// <summary>
        /// The ID of the connected VPN Site Link. Changing this forces a new VPN Gateway Connection to be created.
        /// </summary>
        [Input("vpnSiteLinkId", required: true)]
        public Input<string> VpnSiteLinkId { get; set; } = null!;

        public VpnGatewayConnectionVpnLinkGetArgs()
        {
        }
    }
}

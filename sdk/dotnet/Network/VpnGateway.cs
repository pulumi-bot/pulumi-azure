// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    /// <summary>
    /// Manages a VPN Gateway within a Virtual Hub, which enables Site-to-Site communication.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/vpn_gateway.html.markdown.
    /// </summary>
    public partial class VpnGateway : Pulumi.CustomResource
    {
        /// <summary>
        /// A `bgp_settings` block as defined below.
        /// </summary>
        [Output("bgpSettings")]
        public Output<ImmutableArray<Outputs.VpnGatewayBgpSettings>> BgpSettings { get; private set; } = null!;

        /// <summary>
        /// The Azure location where this VPN Gateway should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The Name which should be used for this VPN Gateway. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Name of the Resource Group in which this VPN Gateway should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The Scale Unit for this VPN Gateway. Defaults to `1`.
        /// </summary>
        [Output("scaleUnit")]
        public Output<int?> ScaleUnit { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the VPN Gateway.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the Virtual Hub within which this VPN Gateway should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("virtualHubId")]
        public Output<string> VirtualHubId { get; private set; } = null!;


        /// <summary>
        /// Create a VpnGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpnGateway(string name, VpnGatewayArgs args, CustomResourceOptions? options = null)
            : base("azure:network/vpnGateway:VpnGateway", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private VpnGateway(string name, Input<string> id, VpnGatewayState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/vpnGateway:VpnGateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpnGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpnGateway Get(string name, Input<string> id, VpnGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new VpnGateway(name, id, state, options);
        }
    }

    public sealed class VpnGatewayArgs : Pulumi.ResourceArgs
    {
        [Input("bgpSettings")]
        private InputList<Inputs.VpnGatewayBgpSettingsArgs>? _bgpSettings;

        /// <summary>
        /// A `bgp_settings` block as defined below.
        /// </summary>
        public InputList<Inputs.VpnGatewayBgpSettingsArgs> BgpSettings
        {
            get => _bgpSettings ?? (_bgpSettings = new InputList<Inputs.VpnGatewayBgpSettingsArgs>());
            set => _bgpSettings = value;
        }

        /// <summary>
        /// The Azure location where this VPN Gateway should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The Name which should be used for this VPN Gateway. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Name of the Resource Group in which this VPN Gateway should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Scale Unit for this VPN Gateway. Defaults to `1`.
        /// </summary>
        [Input("scaleUnit")]
        public Input<int>? ScaleUnit { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the VPN Gateway.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the Virtual Hub within which this VPN Gateway should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("virtualHubId", required: true)]
        public Input<string> VirtualHubId { get; set; } = null!;

        public VpnGatewayArgs()
        {
        }
    }

    public sealed class VpnGatewayState : Pulumi.ResourceArgs
    {
        [Input("bgpSettings")]
        private InputList<Inputs.VpnGatewayBgpSettingsGetArgs>? _bgpSettings;

        /// <summary>
        /// A `bgp_settings` block as defined below.
        /// </summary>
        public InputList<Inputs.VpnGatewayBgpSettingsGetArgs> BgpSettings
        {
            get => _bgpSettings ?? (_bgpSettings = new InputList<Inputs.VpnGatewayBgpSettingsGetArgs>());
            set => _bgpSettings = value;
        }

        /// <summary>
        /// The Azure location where this VPN Gateway should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The Name which should be used for this VPN Gateway. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Name of the Resource Group in which this VPN Gateway should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The Scale Unit for this VPN Gateway. Defaults to `1`.
        /// </summary>
        [Input("scaleUnit")]
        public Input<int>? ScaleUnit { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the VPN Gateway.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the Virtual Hub within which this VPN Gateway should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("virtualHubId")]
        public Input<string>? VirtualHubId { get; set; }

        public VpnGatewayState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class VpnGatewayBgpSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("asn", required: true)]
        public Input<int> Asn { get; set; } = null!;

        /// <summary>
        /// The Address which should be used for the BGP Peering.
        /// </summary>
        [Input("bgpPeeringAddress")]
        public Input<string>? BgpPeeringAddress { get; set; }

        [Input("peerWeight", required: true)]
        public Input<int> PeerWeight { get; set; } = null!;

        public VpnGatewayBgpSettingsArgs()
        {
        }
    }

    public sealed class VpnGatewayBgpSettingsGetArgs : Pulumi.ResourceArgs
    {
        [Input("asn", required: true)]
        public Input<int> Asn { get; set; } = null!;

        /// <summary>
        /// The Address which should be used for the BGP Peering.
        /// </summary>
        [Input("bgpPeeringAddress")]
        public Input<string>? BgpPeeringAddress { get; set; }

        [Input("peerWeight", required: true)]
        public Input<int> PeerWeight { get; set; } = null!;

        public VpnGatewayBgpSettingsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class VpnGatewayBgpSettings
    {
        public readonly int Asn;
        /// <summary>
        /// The Address which should be used for the BGP Peering.
        /// </summary>
        public readonly string BgpPeeringAddress;
        public readonly int PeerWeight;

        [OutputConstructor]
        private VpnGatewayBgpSettings(
            int asn,
            string bgpPeeringAddress,
            int peerWeight)
        {
            Asn = asn;
            BgpPeeringAddress = bgpPeeringAddress;
            PeerWeight = peerWeight;
        }
    }
    }
}
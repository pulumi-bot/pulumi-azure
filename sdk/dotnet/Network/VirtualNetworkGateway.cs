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
    /// Manages a Virtual Network Gateway to establish secure, cross-premises connectivity.
    /// 
    /// &gt; **Note:** Please be aware that provisioning a Virtual Network Gateway takes a long time (between 30 minutes and 1 hour)
    /// 
    /// ## Example Usage
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
    ///                     Name = "vnetGatewayConfig",
    ///                     PublicIpAddressId = examplePublicIp.Id,
    ///                     PrivateIpAddressAllocation = "Dynamic",
    ///                     SubnetId = exampleSubnet.Id,
    ///                 },
    ///             },
    ///             VpnClientConfiguration = new Azure.Network.Inputs.VirtualNetworkGatewayVpnClientConfigurationArgs
    ///             {
    ///                 AddressSpaces = 
    ///                 {
    ///                     "10.2.0.0/24",
    ///                 },
    ///                 RootCertificates = 
    ///                 {
    ///                     new Azure.Network.Inputs.VirtualNetworkGatewayVpnClientConfigurationRootCertificateArgs
    ///                     {
    ///                         Name = "DigiCert-Federated-ID-Root-CA",
    ///                         PublicCertData = @"MIIDuzCCAqOgAwIBAgIQCHTZWCM+IlfFIRXIvyKSrjANBgkqhkiG9w0BAQsFADBn
    /// MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
    /// d3cuZGlnaWNlcnQuY29tMSYwJAYDVQQDEx1EaWdpQ2VydCBGZWRlcmF0ZWQgSUQg
    /// Um9vdCBDQTAeFw0xMzAxMTUxMjAwMDBaFw0zMzAxMTUxMjAwMDBaMGcxCzAJBgNV
    /// BAYTAlVTMRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdp
    /// Y2VydC5jb20xJjAkBgNVBAMTHURpZ2lDZXJ0IEZlZGVyYXRlZCBJRCBSb290IENB
    /// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvAEB4pcCqnNNOWE6Ur5j
    /// QPUH+1y1F9KdHTRSza6k5iDlXq1kGS1qAkuKtw9JsiNRrjltmFnzMZRBbX8Tlfl8
    /// zAhBmb6dDduDGED01kBsTkgywYPxXVTKec0WxYEEF0oMn4wSYNl0lt2eJAKHXjNf
    /// GTwiibdP8CUR2ghSM2sUTI8Nt1Omfc4SMHhGhYD64uJMbX98THQ/4LMGuYegou+d
    /// GTiahfHtjn7AboSEknwAMJHCh5RlYZZ6B1O4QbKJ+34Q0eKgnI3X6Vc9u0zf6DH8
    /// Dk+4zQDYRRTqTnVO3VT8jzqDlCRuNtq6YvryOWN74/dq8LQhUnXHvFyrsdMaE1X2
    /// DwIDAQABo2MwYTAPBgNVHRMBAf8EBTADAQH/MA4GA1UdDwEB/wQEAwIBhjAdBgNV
    /// HQ4EFgQUGRdkFnbGt1EWjKwbUne+5OaZvRYwHwYDVR0jBBgwFoAUGRdkFnbGt1EW
    /// jKwbUne+5OaZvRYwDQYJKoZIhvcNAQELBQADggEBAHcqsHkrjpESqfuVTRiptJfP
    /// 9JbdtWqRTmOf6uJi2c8YVqI6XlKXsD8C1dUUaaHKLUJzvKiazibVuBwMIT84AyqR
    /// QELn3e0BtgEymEygMU569b01ZPxoFSnNXc7qDZBDef8WfqAV/sxkTi8L9BkmFYfL
    /// uGLOhRJOFprPdoDIUBB+tmCl3oDcBy3vnUeOEioz8zAkprcb3GHwHAK+vHmmfgcn
    /// WsfMLH4JCLa/tRYL+Rw/N3ybCkDp00s0WUZ+AoDywSl0Q/ZEnNY0MsFiw6LyIdbq
    /// M/s/1JRtO3bDSzD9TazRVzn2oBqzSa8VgIo5C1nOnoAKJTlsClJKvIhnRlaLQqk=
    /// ",
    ///                     },
    ///                 },
    ///                 RevokedCertificates = 
    ///                 {
    ///                     new Azure.Network.Inputs.VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateArgs
    ///                     {
    ///                         Name = "Verizon-Global-Root-CA",
    ///                         Thumbprint = "912198EEF23DCAC40939312FEE97DD560BAE49B1",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Virtual Network Gateways can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class VirtualNetworkGateway : Pulumi.CustomResource
    {
        /// <summary>
        /// If `true`, an active-active Virtual Network Gateway
        /// will be created. An active-active gateway requires a `HighPerformance` or an
        /// `UltraPerformance` sku. If `false`, an active-standby gateway will be created.
        /// Defaults to `false`.
        /// </summary>
        [Output("activeActive")]
        public Output<bool> ActiveActive { get; private set; } = null!;

        [Output("bgpSettings")]
        public Output<Outputs.VirtualNetworkGatewayBgpSettings> BgpSettings { get; private set; } = null!;

        /// <summary>
        /// The ID of the local network gateway
        /// through which outbound Internet traffic from the virtual network in which the
        /// gateway is created will be routed (*forced tunnelling*). Refer to the
        /// [Azure documentation on forced tunnelling](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-forced-tunneling-rm).
        /// If not specified, forced tunnelling is disabled.
        /// </summary>
        [Output("defaultLocalNetworkGatewayId")]
        public Output<string?> DefaultLocalNetworkGatewayId { get; private set; } = null!;

        /// <summary>
        /// If `true`, BGP (Border Gateway Protocol) will be enabled
        /// for this Virtual Network Gateway. Defaults to `false`.
        /// </summary>
        [Output("enableBgp")]
        public Output<bool> EnableBgp { get; private set; } = null!;

        /// <summary>
        /// The Generation of the Virtual Network gateway. Possible values include `Generation1`, `Generation2` or `None`.
        /// </summary>
        [Output("generation")]
        public Output<string> Generation { get; private set; } = null!;

        /// <summary>
        /// One or two `ip_configuration` blocks documented below.
        /// An active-standby gateway requires exactly one `ip_configuration` block whereas
        /// an active-active gateway requires exactly two `ip_configuration` blocks.
        /// </summary>
        [Output("ipConfigurations")]
        public Output<ImmutableArray<Outputs.VirtualNetworkGatewayIpConfiguration>> IpConfigurations { get; private set; } = null!;

        /// <summary>
        /// The location/region where the Virtual Network Gateway is
        /// located. Changing the location/region forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// A user-defined name of the revoked certificate.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to
        /// create the Virtual Network Gateway. Changing the resource group name forces
        /// a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Configuration of the size and capacity of the virtual network
        /// gateway. Valid options are `Basic`, `Standard`, `HighPerformance`, `UltraPerformance`,
        /// `ErGw1AZ`, `ErGw2AZ`, `ErGw3AZ`, `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw4`,`VpnGw5`, `VpnGw1AZ`,
        /// `VpnGw2AZ`, `VpnGw3AZ`,`VpnGw4AZ` and `VpnGw5AZ` and depend on the `type`, `vpn_type` and
        /// `generation` arguments.
        /// A `PolicyBased` gateway only supports the `Basic` sku. Further, the `UltraPerformance`
        /// sku is only supported by an `ExpressRoute` gateway.
        /// </summary>
        [Output("sku")]
        public Output<string> Sku { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the Virtual Network Gateway. Valid options are
        /// `Vpn` or `ExpressRoute`. Changing the type forces a new resource to be created.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// A `vpn_client_configuration` block which
        /// is documented below. In this block the Virtual Network Gateway can be configured
        /// to accept IPSec point-to-site connections.
        /// </summary>
        [Output("vpnClientConfiguration")]
        public Output<Outputs.VirtualNetworkGatewayVpnClientConfiguration?> VpnClientConfiguration { get; private set; } = null!;

        /// <summary>
        /// The routing type of the Virtual Network Gateway. Valid
        /// options are `RouteBased` or `PolicyBased`. Defaults to `RouteBased`.
        /// </summary>
        [Output("vpnType")]
        public Output<string?> VpnType { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualNetworkGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualNetworkGateway(string name, VirtualNetworkGatewayArgs args, CustomResourceOptions? options = null)
            : base("azure:network/virtualNetworkGateway:VirtualNetworkGateway", name, args ?? new VirtualNetworkGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualNetworkGateway(string name, Input<string> id, VirtualNetworkGatewayState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/virtualNetworkGateway:VirtualNetworkGateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualNetworkGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualNetworkGateway Get(string name, Input<string> id, VirtualNetworkGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualNetworkGateway(name, id, state, options);
        }
    }

    public sealed class VirtualNetworkGatewayArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If `true`, an active-active Virtual Network Gateway
        /// will be created. An active-active gateway requires a `HighPerformance` or an
        /// `UltraPerformance` sku. If `false`, an active-standby gateway will be created.
        /// Defaults to `false`.
        /// </summary>
        [Input("activeActive")]
        public Input<bool>? ActiveActive { get; set; }

        [Input("bgpSettings")]
        public Input<Inputs.VirtualNetworkGatewayBgpSettingsArgs>? BgpSettings { get; set; }

        /// <summary>
        /// The ID of the local network gateway
        /// through which outbound Internet traffic from the virtual network in which the
        /// gateway is created will be routed (*forced tunnelling*). Refer to the
        /// [Azure documentation on forced tunnelling](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-forced-tunneling-rm).
        /// If not specified, forced tunnelling is disabled.
        /// </summary>
        [Input("defaultLocalNetworkGatewayId")]
        public Input<string>? DefaultLocalNetworkGatewayId { get; set; }

        /// <summary>
        /// If `true`, BGP (Border Gateway Protocol) will be enabled
        /// for this Virtual Network Gateway. Defaults to `false`.
        /// </summary>
        [Input("enableBgp")]
        public Input<bool>? EnableBgp { get; set; }

        /// <summary>
        /// The Generation of the Virtual Network gateway. Possible values include `Generation1`, `Generation2` or `None`.
        /// </summary>
        [Input("generation")]
        public Input<string>? Generation { get; set; }

        [Input("ipConfigurations", required: true)]
        private InputList<Inputs.VirtualNetworkGatewayIpConfigurationArgs>? _ipConfigurations;

        /// <summary>
        /// One or two `ip_configuration` blocks documented below.
        /// An active-standby gateway requires exactly one `ip_configuration` block whereas
        /// an active-active gateway requires exactly two `ip_configuration` blocks.
        /// </summary>
        public InputList<Inputs.VirtualNetworkGatewayIpConfigurationArgs> IpConfigurations
        {
            get => _ipConfigurations ?? (_ipConfigurations = new InputList<Inputs.VirtualNetworkGatewayIpConfigurationArgs>());
            set => _ipConfigurations = value;
        }

        /// <summary>
        /// The location/region where the Virtual Network Gateway is
        /// located. Changing the location/region forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A user-defined name of the revoked certificate.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to
        /// create the Virtual Network Gateway. Changing the resource group name forces
        /// a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Configuration of the size and capacity of the virtual network
        /// gateway. Valid options are `Basic`, `Standard`, `HighPerformance`, `UltraPerformance`,
        /// `ErGw1AZ`, `ErGw2AZ`, `ErGw3AZ`, `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw4`,`VpnGw5`, `VpnGw1AZ`,
        /// `VpnGw2AZ`, `VpnGw3AZ`,`VpnGw4AZ` and `VpnGw5AZ` and depend on the `type`, `vpn_type` and
        /// `generation` arguments.
        /// A `PolicyBased` gateway only supports the `Basic` sku. Further, the `UltraPerformance`
        /// sku is only supported by an `ExpressRoute` gateway.
        /// </summary>
        [Input("sku", required: true)]
        public Input<string> Sku { get; set; } = null!;

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
        /// The type of the Virtual Network Gateway. Valid options are
        /// `Vpn` or `ExpressRoute`. Changing the type forces a new resource to be created.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// A `vpn_client_configuration` block which
        /// is documented below. In this block the Virtual Network Gateway can be configured
        /// to accept IPSec point-to-site connections.
        /// </summary>
        [Input("vpnClientConfiguration")]
        public Input<Inputs.VirtualNetworkGatewayVpnClientConfigurationArgs>? VpnClientConfiguration { get; set; }

        /// <summary>
        /// The routing type of the Virtual Network Gateway. Valid
        /// options are `RouteBased` or `PolicyBased`. Defaults to `RouteBased`.
        /// </summary>
        [Input("vpnType")]
        public Input<string>? VpnType { get; set; }

        public VirtualNetworkGatewayArgs()
        {
        }
    }

    public sealed class VirtualNetworkGatewayState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If `true`, an active-active Virtual Network Gateway
        /// will be created. An active-active gateway requires a `HighPerformance` or an
        /// `UltraPerformance` sku. If `false`, an active-standby gateway will be created.
        /// Defaults to `false`.
        /// </summary>
        [Input("activeActive")]
        public Input<bool>? ActiveActive { get; set; }

        [Input("bgpSettings")]
        public Input<Inputs.VirtualNetworkGatewayBgpSettingsGetArgs>? BgpSettings { get; set; }

        /// <summary>
        /// The ID of the local network gateway
        /// through which outbound Internet traffic from the virtual network in which the
        /// gateway is created will be routed (*forced tunnelling*). Refer to the
        /// [Azure documentation on forced tunnelling](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-forced-tunneling-rm).
        /// If not specified, forced tunnelling is disabled.
        /// </summary>
        [Input("defaultLocalNetworkGatewayId")]
        public Input<string>? DefaultLocalNetworkGatewayId { get; set; }

        /// <summary>
        /// If `true`, BGP (Border Gateway Protocol) will be enabled
        /// for this Virtual Network Gateway. Defaults to `false`.
        /// </summary>
        [Input("enableBgp")]
        public Input<bool>? EnableBgp { get; set; }

        /// <summary>
        /// The Generation of the Virtual Network gateway. Possible values include `Generation1`, `Generation2` or `None`.
        /// </summary>
        [Input("generation")]
        public Input<string>? Generation { get; set; }

        [Input("ipConfigurations")]
        private InputList<Inputs.VirtualNetworkGatewayIpConfigurationGetArgs>? _ipConfigurations;

        /// <summary>
        /// One or two `ip_configuration` blocks documented below.
        /// An active-standby gateway requires exactly one `ip_configuration` block whereas
        /// an active-active gateway requires exactly two `ip_configuration` blocks.
        /// </summary>
        public InputList<Inputs.VirtualNetworkGatewayIpConfigurationGetArgs> IpConfigurations
        {
            get => _ipConfigurations ?? (_ipConfigurations = new InputList<Inputs.VirtualNetworkGatewayIpConfigurationGetArgs>());
            set => _ipConfigurations = value;
        }

        /// <summary>
        /// The location/region where the Virtual Network Gateway is
        /// located. Changing the location/region forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A user-defined name of the revoked certificate.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to
        /// create the Virtual Network Gateway. Changing the resource group name forces
        /// a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Configuration of the size and capacity of the virtual network
        /// gateway. Valid options are `Basic`, `Standard`, `HighPerformance`, `UltraPerformance`,
        /// `ErGw1AZ`, `ErGw2AZ`, `ErGw3AZ`, `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw4`,`VpnGw5`, `VpnGw1AZ`,
        /// `VpnGw2AZ`, `VpnGw3AZ`,`VpnGw4AZ` and `VpnGw5AZ` and depend on the `type`, `vpn_type` and
        /// `generation` arguments.
        /// A `PolicyBased` gateway only supports the `Basic` sku. Further, the `UltraPerformance`
        /// sku is only supported by an `ExpressRoute` gateway.
        /// </summary>
        [Input("sku")]
        public Input<string>? Sku { get; set; }

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
        /// The type of the Virtual Network Gateway. Valid options are
        /// `Vpn` or `ExpressRoute`. Changing the type forces a new resource to be created.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// A `vpn_client_configuration` block which
        /// is documented below. In this block the Virtual Network Gateway can be configured
        /// to accept IPSec point-to-site connections.
        /// </summary>
        [Input("vpnClientConfiguration")]
        public Input<Inputs.VirtualNetworkGatewayVpnClientConfigurationGetArgs>? VpnClientConfiguration { get; set; }

        /// <summary>
        /// The routing type of the Virtual Network Gateway. Valid
        /// options are `RouteBased` or `PolicyBased`. Defaults to `RouteBased`.
        /// </summary>
        [Input("vpnType")]
        public Input<string>? VpnType { get; set; }

        public VirtualNetworkGatewayState()
        {
        }
    }
}

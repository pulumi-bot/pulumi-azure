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
    /// Manages a Point-to-Site VPN Gateway.
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
    ///         var example = new Azure.Network.PointToPointVpnGateway("example", new Azure.Network.PointToPointVpnGatewayArgs
    ///         {
    ///             Location = azurerm_resource_group.Example.Location,
    ///             ResourceGroupName = azurerm_resource_group.Example.Resource_group_name,
    ///             VirtualHubId = azurerm_virtual_hub.Example.Id,
    ///             VpnServerConfigurationId = azurerm_vpn_server_configuration.Example.Id,
    ///             ScaleUnit = 1,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Point-to-Site VPN Gateway's can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class PointToPointVpnGateway : Pulumi.CustomResource
    {
        /// <summary>
        /// A `connection_configuration` block as defined below.
        /// </summary>
        [Output("connectionConfiguration")]
        public Output<Outputs.PointToPointVpnGatewayConnectionConfiguration> ConnectionConfiguration { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The Scale Unit for this Point-to-Site VPN Gateway.
        /// </summary>
        [Output("scaleUnit")]
        public Output<int> ScaleUnit { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the Point-to-Site VPN Gateway.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the Virtual Hub where this Point-to-Site VPN Gateway should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("virtualHubId")]
        public Output<string> VirtualHubId { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPN Server Configuration which this Point-to-Site VPN Gateway should use. Changing this forces a new resource to be created.
        /// </summary>
        [Output("vpnServerConfigurationId")]
        public Output<string> VpnServerConfigurationId { get; private set; } = null!;


        /// <summary>
        /// Create a PointToPointVpnGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PointToPointVpnGateway(string name, PointToPointVpnGatewayArgs args, CustomResourceOptions? options = null)
            : base("azure:network/pointToPointVpnGateway:PointToPointVpnGateway", name, args ?? new PointToPointVpnGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PointToPointVpnGateway(string name, Input<string> id, PointToPointVpnGatewayState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/pointToPointVpnGateway:PointToPointVpnGateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PointToPointVpnGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PointToPointVpnGateway Get(string name, Input<string> id, PointToPointVpnGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new PointToPointVpnGateway(name, id, state, options);
        }
    }

    public sealed class PointToPointVpnGatewayArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `connection_configuration` block as defined below.
        /// </summary>
        [Input("connectionConfiguration", required: true)]
        public Input<Inputs.PointToPointVpnGatewayConnectionConfigurationArgs> ConnectionConfiguration { get; set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Scale Unit for this Point-to-Site VPN Gateway.
        /// </summary>
        [Input("scaleUnit", required: true)]
        public Input<int> ScaleUnit { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the Point-to-Site VPN Gateway.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the Virtual Hub where this Point-to-Site VPN Gateway should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("virtualHubId", required: true)]
        public Input<string> VirtualHubId { get; set; } = null!;

        /// <summary>
        /// The ID of the VPN Server Configuration which this Point-to-Site VPN Gateway should use. Changing this forces a new resource to be created.
        /// </summary>
        [Input("vpnServerConfigurationId", required: true)]
        public Input<string> VpnServerConfigurationId { get; set; } = null!;

        public PointToPointVpnGatewayArgs()
        {
        }
    }

    public sealed class PointToPointVpnGatewayState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `connection_configuration` block as defined below.
        /// </summary>
        [Input("connectionConfiguration")]
        public Input<Inputs.PointToPointVpnGatewayConnectionConfigurationGetArgs>? ConnectionConfiguration { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The Scale Unit for this Point-to-Site VPN Gateway.
        /// </summary>
        [Input("scaleUnit")]
        public Input<int>? ScaleUnit { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the Point-to-Site VPN Gateway.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the Virtual Hub where this Point-to-Site VPN Gateway should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("virtualHubId")]
        public Input<string>? VirtualHubId { get; set; }

        /// <summary>
        /// The ID of the VPN Server Configuration which this Point-to-Site VPN Gateway should use. Changing this forces a new resource to be created.
        /// </summary>
        [Input("vpnServerConfigurationId")]
        public Input<string>? VpnServerConfigurationId { get; set; }

        public PointToPointVpnGatewayState()
        {
        }
    }
}

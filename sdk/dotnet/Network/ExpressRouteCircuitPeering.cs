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
    /// Manages an ExpressRoute Circuit Peering.
    /// 
    /// ## Example Usage
    /// ### Creating A Microsoft Peering)
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
    ///         var exampleExpressRouteCircuit = new Azure.Network.ExpressRouteCircuit("exampleExpressRouteCircuit", new Azure.Network.ExpressRouteCircuitArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             ServiceProviderName = "Equinix",
    ///             PeeringLocation = "Silicon Valley",
    ///             BandwidthInMbps = 50,
    ///             Sku = new Azure.Network.Inputs.ExpressRouteCircuitSkuArgs
    ///             {
    ///                 Tier = "Standard",
    ///                 Family = "MeteredData",
    ///             },
    ///             AllowClassicOperations = false,
    ///             Tags = 
    ///             {
    ///                 { "environment", "Production" },
    ///             },
    ///         });
    ///         var exampleExpressRouteCircuitPeering = new Azure.Network.ExpressRouteCircuitPeering("exampleExpressRouteCircuitPeering", new Azure.Network.ExpressRouteCircuitPeeringArgs
    ///         {
    ///             PeeringType = "MicrosoftPeering",
    ///             ExpressRouteCircuitName = exampleExpressRouteCircuit.Name,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             PeerAsn = 100,
    ///             PrimaryPeerAddressPrefix = "123.0.0.0/30",
    ///             SecondaryPeerAddressPrefix = "123.0.0.4/30",
    ///             VlanId = 300,
    ///             MicrosoftPeeringConfig = new Azure.Network.Inputs.ExpressRouteCircuitPeeringMicrosoftPeeringConfigArgs
    ///             {
    ///                 AdvertisedPublicPrefixes = 
    ///                 {
    ///                     "123.1.0.0/24",
    ///                 },
    ///             },
    ///             Ipv6 = new Azure.Network.Inputs.ExpressRouteCircuitPeeringIpv6Args
    ///             {
    ///                 PrimaryPeerAddressPrefix = "2002:db01::/126",
    ///                 SecondaryPeerAddressPrefix = "2003:db01::/126",
    ///                 MicrosoftPeering = new Azure.Network.Inputs.ExpressRouteCircuitPeeringIpv6MicrosoftPeeringArgs
    ///                 {
    ///                     AdvertisedPublicPrefixes = 
    ///                     {
    ///                         "2002:db01::/126",
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
    /// ExpressRoute Circuit Peerings can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:network/expressRouteCircuitPeering:ExpressRouteCircuitPeering peering1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/expressRouteCircuits/myExpressRoute/peerings/peering1
    /// ```
    /// </summary>
    [AzureResourceType("azure:network/expressRouteCircuitPeering:ExpressRouteCircuitPeering")]
    public partial class ExpressRouteCircuitPeering : Pulumi.CustomResource
    {
        /// <summary>
        /// The ASN used by Azure.
        /// </summary>
        [Output("azureAsn")]
        public Output<int> AzureAsn { get; private set; } = null!;

        /// <summary>
        /// The name of the ExpressRoute Circuit in which to create the Peering.
        /// </summary>
        [Output("expressRouteCircuitName")]
        public Output<string> ExpressRouteCircuitName { get; private set; } = null!;

        /// <summary>
        /// A `ipv6` block as defined below.
        /// </summary>
        [Output("ipv6")]
        public Output<Outputs.ExpressRouteCircuitPeeringIpv6?> Ipv6 { get; private set; } = null!;

        /// <summary>
        /// A `microsoft_peering_config` block as defined below. Required when `peering_type` is set to `MicrosoftPeering`.
        /// </summary>
        [Output("microsoftPeeringConfig")]
        public Output<Outputs.ExpressRouteCircuitPeeringMicrosoftPeeringConfig?> MicrosoftPeeringConfig { get; private set; } = null!;

        /// <summary>
        /// The Either a 16-bit or a 32-bit ASN. Can either be public or private.
        /// </summary>
        [Output("peerAsn")]
        public Output<int> PeerAsn { get; private set; } = null!;

        /// <summary>
        /// The type of the ExpressRoute Circuit Peering. Acceptable values include `AzurePrivatePeering`, `AzurePublicPeering` and `MicrosoftPeering`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("peeringType")]
        public Output<string> PeeringType { get; private set; } = null!;

        /// <summary>
        /// The Primary Port used by Azure for this Peering.
        /// </summary>
        [Output("primaryAzurePort")]
        public Output<string> PrimaryAzurePort { get; private set; } = null!;

        /// <summary>
        /// A subnet for the primary link.
        /// </summary>
        [Output("primaryPeerAddressPrefix")]
        public Output<string> PrimaryPeerAddressPrefix { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Express Route Circuit Peering. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The ID of the Route Filter. Only available when `peering_type` is set to `MicrosoftPeering`.
        /// </summary>
        [Output("routeFilterId")]
        public Output<string?> RouteFilterId { get; private set; } = null!;

        /// <summary>
        /// The Secondary Port used by Azure for this Peering.
        /// </summary>
        [Output("secondaryAzurePort")]
        public Output<string> SecondaryAzurePort { get; private set; } = null!;

        /// <summary>
        /// A subnet for the secondary link.
        /// </summary>
        [Output("secondaryPeerAddressPrefix")]
        public Output<string> SecondaryPeerAddressPrefix { get; private set; } = null!;

        /// <summary>
        /// The shared key. Can be a maximum of 25 characters.
        /// </summary>
        [Output("sharedKey")]
        public Output<string?> SharedKey { get; private set; } = null!;

        /// <summary>
        /// A valid VLAN ID to establish this peering on.
        /// </summary>
        [Output("vlanId")]
        public Output<int> VlanId { get; private set; } = null!;


        /// <summary>
        /// Create a ExpressRouteCircuitPeering resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExpressRouteCircuitPeering(string name, ExpressRouteCircuitPeeringArgs args, CustomResourceOptions? options = null)
            : base("azure:network/expressRouteCircuitPeering:ExpressRouteCircuitPeering", name, args ?? new ExpressRouteCircuitPeeringArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExpressRouteCircuitPeering(string name, Input<string> id, ExpressRouteCircuitPeeringState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/expressRouteCircuitPeering:ExpressRouteCircuitPeering", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ExpressRouteCircuitPeering resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExpressRouteCircuitPeering Get(string name, Input<string> id, ExpressRouteCircuitPeeringState? state = null, CustomResourceOptions? options = null)
        {
            return new ExpressRouteCircuitPeering(name, id, state, options);
        }
    }

    public sealed class ExpressRouteCircuitPeeringArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the ExpressRoute Circuit in which to create the Peering.
        /// </summary>
        [Input("expressRouteCircuitName", required: true)]
        public Input<string> ExpressRouteCircuitName { get; set; } = null!;

        /// <summary>
        /// A `ipv6` block as defined below.
        /// </summary>
        [Input("ipv6")]
        public Input<Inputs.ExpressRouteCircuitPeeringIpv6Args>? Ipv6 { get; set; }

        /// <summary>
        /// A `microsoft_peering_config` block as defined below. Required when `peering_type` is set to `MicrosoftPeering`.
        /// </summary>
        [Input("microsoftPeeringConfig")]
        public Input<Inputs.ExpressRouteCircuitPeeringMicrosoftPeeringConfigArgs>? MicrosoftPeeringConfig { get; set; }

        /// <summary>
        /// The Either a 16-bit or a 32-bit ASN. Can either be public or private.
        /// </summary>
        [Input("peerAsn")]
        public Input<int>? PeerAsn { get; set; }

        /// <summary>
        /// The type of the ExpressRoute Circuit Peering. Acceptable values include `AzurePrivatePeering`, `AzurePublicPeering` and `MicrosoftPeering`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("peeringType", required: true)]
        public Input<string> PeeringType { get; set; } = null!;

        /// <summary>
        /// A subnet for the primary link.
        /// </summary>
        [Input("primaryPeerAddressPrefix", required: true)]
        public Input<string> PrimaryPeerAddressPrefix { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Express Route Circuit Peering. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The ID of the Route Filter. Only available when `peering_type` is set to `MicrosoftPeering`.
        /// </summary>
        [Input("routeFilterId")]
        public Input<string>? RouteFilterId { get; set; }

        /// <summary>
        /// A subnet for the secondary link.
        /// </summary>
        [Input("secondaryPeerAddressPrefix", required: true)]
        public Input<string> SecondaryPeerAddressPrefix { get; set; } = null!;

        /// <summary>
        /// The shared key. Can be a maximum of 25 characters.
        /// </summary>
        [Input("sharedKey")]
        public Input<string>? SharedKey { get; set; }

        /// <summary>
        /// A valid VLAN ID to establish this peering on.
        /// </summary>
        [Input("vlanId", required: true)]
        public Input<int> VlanId { get; set; } = null!;

        public ExpressRouteCircuitPeeringArgs()
        {
        }
    }

    public sealed class ExpressRouteCircuitPeeringState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ASN used by Azure.
        /// </summary>
        [Input("azureAsn")]
        public Input<int>? AzureAsn { get; set; }

        /// <summary>
        /// The name of the ExpressRoute Circuit in which to create the Peering.
        /// </summary>
        [Input("expressRouteCircuitName")]
        public Input<string>? ExpressRouteCircuitName { get; set; }

        /// <summary>
        /// A `ipv6` block as defined below.
        /// </summary>
        [Input("ipv6")]
        public Input<Inputs.ExpressRouteCircuitPeeringIpv6GetArgs>? Ipv6 { get; set; }

        /// <summary>
        /// A `microsoft_peering_config` block as defined below. Required when `peering_type` is set to `MicrosoftPeering`.
        /// </summary>
        [Input("microsoftPeeringConfig")]
        public Input<Inputs.ExpressRouteCircuitPeeringMicrosoftPeeringConfigGetArgs>? MicrosoftPeeringConfig { get; set; }

        /// <summary>
        /// The Either a 16-bit or a 32-bit ASN. Can either be public or private.
        /// </summary>
        [Input("peerAsn")]
        public Input<int>? PeerAsn { get; set; }

        /// <summary>
        /// The type of the ExpressRoute Circuit Peering. Acceptable values include `AzurePrivatePeering`, `AzurePublicPeering` and `MicrosoftPeering`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("peeringType")]
        public Input<string>? PeeringType { get; set; }

        /// <summary>
        /// The Primary Port used by Azure for this Peering.
        /// </summary>
        [Input("primaryAzurePort")]
        public Input<string>? PrimaryAzurePort { get; set; }

        /// <summary>
        /// A subnet for the primary link.
        /// </summary>
        [Input("primaryPeerAddressPrefix")]
        public Input<string>? PrimaryPeerAddressPrefix { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Express Route Circuit Peering. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The ID of the Route Filter. Only available when `peering_type` is set to `MicrosoftPeering`.
        /// </summary>
        [Input("routeFilterId")]
        public Input<string>? RouteFilterId { get; set; }

        /// <summary>
        /// The Secondary Port used by Azure for this Peering.
        /// </summary>
        [Input("secondaryAzurePort")]
        public Input<string>? SecondaryAzurePort { get; set; }

        /// <summary>
        /// A subnet for the secondary link.
        /// </summary>
        [Input("secondaryPeerAddressPrefix")]
        public Input<string>? SecondaryPeerAddressPrefix { get; set; }

        /// <summary>
        /// The shared key. Can be a maximum of 25 characters.
        /// </summary>
        [Input("sharedKey")]
        public Input<string>? SharedKey { get; set; }

        /// <summary>
        /// A valid VLAN ID to establish this peering on.
        /// </summary>
        [Input("vlanId")]
        public Input<int>? VlanId { get; set; }

        public ExpressRouteCircuitPeeringState()
        {
        }
    }
}

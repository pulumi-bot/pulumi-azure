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
    /// Manages the association between a Network Interface and a Load Balancer's NAT Rule.
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
    ///             Location = "West Europe",
    ///         });
    ///         var exampleVirtualNetwork = new Azure.Network.VirtualNetwork("exampleVirtualNetwork", new Azure.Network.VirtualNetworkArgs
    ///         {
    ///             AddressSpaces = 
    ///             {
    ///                 "10.0.0.0/16",
    ///             },
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var exampleSubnet = new Azure.Network.Subnet("exampleSubnet", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualNetworkName = exampleVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.0.2.0/24",
    ///             },
    ///         });
    ///         var examplePublicIp = new Azure.Network.PublicIp("examplePublicIp", new Azure.Network.PublicIpArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AllocationMethod = "Static",
    ///         });
    ///         var exampleLoadBalancer = new Azure.Lb.LoadBalancer("exampleLoadBalancer", new Azure.Lb.LoadBalancerArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             FrontendIpConfigurations = 
    ///             {
    ///                 new Azure.Lb.Inputs.LoadBalancerFrontendIpConfigurationArgs
    ///                 {
    ///                     Name = "primary",
    ///                     PublicIpAddressId = examplePublicIp.Id,
    ///                 },
    ///             },
    ///         });
    ///         var exampleNatRule = new Azure.Lb.NatRule("exampleNatRule", new Azure.Lb.NatRuleArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             LoadbalancerId = exampleLoadBalancer.Id,
    ///             Protocol = "Tcp",
    ///             FrontendPort = 3389,
    ///             BackendPort = 3389,
    ///             FrontendIpConfigurationName = "primary",
    ///         });
    ///         var exampleNetworkInterface = new Azure.Network.NetworkInterface("exampleNetworkInterface", new Azure.Network.NetworkInterfaceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             IpConfigurations = 
    ///             {
    ///                 new Azure.Network.Inputs.NetworkInterfaceIpConfigurationArgs
    ///                 {
    ///                     Name = "testconfiguration1",
    ///                     SubnetId = exampleSubnet.Id,
    ///                     PrivateIpAddressAllocation = "Dynamic",
    ///                 },
    ///             },
    ///         });
    ///         var exampleNetworkInterfaceNatRuleAssociation = new Azure.Network.NetworkInterfaceNatRuleAssociation("exampleNetworkInterfaceNatRuleAssociation", new Azure.Network.NetworkInterfaceNatRuleAssociationArgs
    ///         {
    ///             NetworkInterfaceId = exampleNetworkInterface.Id,
    ///             IpConfigurationName = "testconfiguration1",
    ///             NatRuleId = exampleNatRule.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Associations between Network Interfaces and Load Balancer NAT Rule can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:network/networkInterfaceNatRuleAssociation:NetworkInterfaceNatRuleAssociation association1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.network/networkInterfaces/nic1/ipConfigurations/example|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/rule1
    /// ```
    /// </summary>
    public partial class NetworkInterfaceNatRuleAssociation : Pulumi.CustomResource
    {
        /// <summary>
        /// The Name of the IP Configuration within the Network Interface which should be connected to the NAT Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Output("ipConfigurationName")]
        public Output<string> IpConfigurationName { get; private set; } = null!;

        /// <summary>
        /// The ID of the Load Balancer NAT Rule which this Network Interface which should be connected to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("natRuleId")]
        public Output<string> NatRuleId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Network Interface. Changing this forces a new resource to be created.
        /// </summary>
        [Output("networkInterfaceId")]
        public Output<string> NetworkInterfaceId { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkInterfaceNatRuleAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkInterfaceNatRuleAssociation(string name, NetworkInterfaceNatRuleAssociationArgs args, CustomResourceOptions? options = null)
            : base("azure:network/networkInterfaceNatRuleAssociation:NetworkInterfaceNatRuleAssociation", name, args ?? new NetworkInterfaceNatRuleAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkInterfaceNatRuleAssociation(string name, Input<string> id, NetworkInterfaceNatRuleAssociationState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/networkInterfaceNatRuleAssociation:NetworkInterfaceNatRuleAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkInterfaceNatRuleAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkInterfaceNatRuleAssociation Get(string name, Input<string> id, NetworkInterfaceNatRuleAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkInterfaceNatRuleAssociation(name, id, state, options);
        }
    }

    public sealed class NetworkInterfaceNatRuleAssociationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Name of the IP Configuration within the Network Interface which should be connected to the NAT Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("ipConfigurationName", required: true)]
        public Input<string> IpConfigurationName { get; set; } = null!;

        /// <summary>
        /// The ID of the Load Balancer NAT Rule which this Network Interface which should be connected to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("natRuleId", required: true)]
        public Input<string> NatRuleId { get; set; } = null!;

        /// <summary>
        /// The ID of the Network Interface. Changing this forces a new resource to be created.
        /// </summary>
        [Input("networkInterfaceId", required: true)]
        public Input<string> NetworkInterfaceId { get; set; } = null!;

        public NetworkInterfaceNatRuleAssociationArgs()
        {
        }
    }

    public sealed class NetworkInterfaceNatRuleAssociationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Name of the IP Configuration within the Network Interface which should be connected to the NAT Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("ipConfigurationName")]
        public Input<string>? IpConfigurationName { get; set; }

        /// <summary>
        /// The ID of the Load Balancer NAT Rule which this Network Interface which should be connected to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("natRuleId")]
        public Input<string>? NatRuleId { get; set; }

        /// <summary>
        /// The ID of the Network Interface. Changing this forces a new resource to be created.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        public NetworkInterfaceNatRuleAssociationState()
        {
        }
    }
}

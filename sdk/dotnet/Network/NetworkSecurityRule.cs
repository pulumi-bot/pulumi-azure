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
    /// Manages a Network Security Rule.
    /// 
    /// &gt; **NOTE on Network Security Groups and Network Security Rules:** This provider currently
    /// provides both a standalone Network Security Rule resource, and allows for Network Security Rules to be defined in-line within the Network Security Group resource.
    /// At this time you cannot use a Network Security Group with in-line Network Security Rules in conjunction with any Network Security Rule resources. Doing so will cause a conflict of rule settings and will overwrite rules.
    /// 
    /// ## Example Usage
    /// 
    /// 
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
    ///         var exampleNetworkSecurityGroup = new Azure.Network.NetworkSecurityGroup("exampleNetworkSecurityGroup", new Azure.Network.NetworkSecurityGroupArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var exampleNetworkSecurityRule = new Azure.Network.NetworkSecurityRule("exampleNetworkSecurityRule", new Azure.Network.NetworkSecurityRuleArgs
    ///         {
    ///             Priority = 100,
    ///             Direction = "Outbound",
    ///             Access = "Allow",
    ///             Protocol = "Tcp",
    ///             SourcePortRange = "*",
    ///             DestinationPortRange = "*",
    ///             SourceAddressPrefix = "*",
    ///             DestinationAddressPrefix = "*",
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             NetworkSecurityGroupName = exampleNetworkSecurityGroup.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class NetworkSecurityRule : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
        /// </summary>
        [Output("access")]
        public Output<string> Access { get; private set; } = null!;

        /// <summary>
        /// A description for this rule. Restricted to 140 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// CIDR or destination IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `destination_address_prefixes` is not specified.
        /// </summary>
        [Output("destinationAddressPrefix")]
        public Output<string?> DestinationAddressPrefix { get; private set; } = null!;

        /// <summary>
        /// List of destination address prefixes. Tags may not be used. This is required if `destination_address_prefix` is not specified.
        /// </summary>
        [Output("destinationAddressPrefixes")]
        public Output<ImmutableArray<string>> DestinationAddressPrefixes { get; private set; } = null!;

        /// <summary>
        /// A List of destination Application Security Group ID's
        /// </summary>
        [Output("destinationApplicationSecurityGroupIds")]
        public Output<string?> DestinationApplicationSecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `destination_port_ranges` is not specified.
        /// </summary>
        [Output("destinationPortRange")]
        public Output<string?> DestinationPortRange { get; private set; } = null!;

        /// <summary>
        /// List of destination ports or port ranges. This is required if `destination_port_range` is not specified.
        /// </summary>
        [Output("destinationPortRanges")]
        public Output<ImmutableArray<string>> DestinationPortRanges { get; private set; } = null!;

        /// <summary>
        /// The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
        /// </summary>
        [Output("direction")]
        public Output<string> Direction { get; private set; } = null!;

        /// <summary>
        /// The name of the security rule. This needs to be unique across all Rules in the Network Security Group. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the Network Security Group that we want to attach the rule to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("networkSecurityGroupName")]
        public Output<string> NetworkSecurityGroupName { get; private set; } = null!;

        /// <summary>
        /// Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// Network protocol this rule applies to. Possible values include `Tcp`, `Udp`, `Icmp`, or `*` (which matches all).
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Network Security Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// CIDR or source IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `source_address_prefixes` is not specified.
        /// </summary>
        [Output("sourceAddressPrefix")]
        public Output<string?> SourceAddressPrefix { get; private set; } = null!;

        /// <summary>
        /// List of source address prefixes. Tags may not be used. This is required if `source_address_prefix` is not specified.
        /// </summary>
        [Output("sourceAddressPrefixes")]
        public Output<ImmutableArray<string>> SourceAddressPrefixes { get; private set; } = null!;

        /// <summary>
        /// A List of source Application Security Group ID's
        /// </summary>
        [Output("sourceApplicationSecurityGroupIds")]
        public Output<string?> SourceApplicationSecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// Source Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `source_port_ranges` is not specified.
        /// </summary>
        [Output("sourcePortRange")]
        public Output<string?> SourcePortRange { get; private set; } = null!;

        /// <summary>
        /// List of source ports or port ranges. This is required if `source_port_range` is not specified.
        /// </summary>
        [Output("sourcePortRanges")]
        public Output<ImmutableArray<string>> SourcePortRanges { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkSecurityRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkSecurityRule(string name, NetworkSecurityRuleArgs args, CustomResourceOptions? options = null)
            : base("azure:network/networkSecurityRule:NetworkSecurityRule", name, args ?? new NetworkSecurityRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkSecurityRule(string name, Input<string> id, NetworkSecurityRuleState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/networkSecurityRule:NetworkSecurityRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkSecurityRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkSecurityRule Get(string name, Input<string> id, NetworkSecurityRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkSecurityRule(name, id, state, options);
        }
    }

    public sealed class NetworkSecurityRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
        /// </summary>
        [Input("access", required: true)]
        public Input<string> Access { get; set; } = null!;

        /// <summary>
        /// A description for this rule. Restricted to 140 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// CIDR or destination IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `destination_address_prefixes` is not specified.
        /// </summary>
        [Input("destinationAddressPrefix")]
        public Input<string>? DestinationAddressPrefix { get; set; }

        [Input("destinationAddressPrefixes")]
        private InputList<string>? _destinationAddressPrefixes;

        /// <summary>
        /// List of destination address prefixes. Tags may not be used. This is required if `destination_address_prefix` is not specified.
        /// </summary>
        public InputList<string> DestinationAddressPrefixes
        {
            get => _destinationAddressPrefixes ?? (_destinationAddressPrefixes = new InputList<string>());
            set => _destinationAddressPrefixes = value;
        }

        /// <summary>
        /// A List of destination Application Security Group ID's
        /// </summary>
        [Input("destinationApplicationSecurityGroupIds")]
        public Input<string>? DestinationApplicationSecurityGroupIds { get; set; }

        /// <summary>
        /// Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `destination_port_ranges` is not specified.
        /// </summary>
        [Input("destinationPortRange")]
        public Input<string>? DestinationPortRange { get; set; }

        [Input("destinationPortRanges")]
        private InputList<string>? _destinationPortRanges;

        /// <summary>
        /// List of destination ports or port ranges. This is required if `destination_port_range` is not specified.
        /// </summary>
        public InputList<string> DestinationPortRanges
        {
            get => _destinationPortRanges ?? (_destinationPortRanges = new InputList<string>());
            set => _destinationPortRanges = value;
        }

        /// <summary>
        /// The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
        /// </summary>
        [Input("direction", required: true)]
        public Input<string> Direction { get; set; } = null!;

        /// <summary>
        /// The name of the security rule. This needs to be unique across all Rules in the Network Security Group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Network Security Group that we want to attach the rule to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("networkSecurityGroupName", required: true)]
        public Input<string> NetworkSecurityGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        /// <summary>
        /// Network protocol this rule applies to. Possible values include `Tcp`, `Udp`, `Icmp`, or `*` (which matches all).
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Network Security Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// CIDR or source IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `source_address_prefixes` is not specified.
        /// </summary>
        [Input("sourceAddressPrefix")]
        public Input<string>? SourceAddressPrefix { get; set; }

        [Input("sourceAddressPrefixes")]
        private InputList<string>? _sourceAddressPrefixes;

        /// <summary>
        /// List of source address prefixes. Tags may not be used. This is required if `source_address_prefix` is not specified.
        /// </summary>
        public InputList<string> SourceAddressPrefixes
        {
            get => _sourceAddressPrefixes ?? (_sourceAddressPrefixes = new InputList<string>());
            set => _sourceAddressPrefixes = value;
        }

        /// <summary>
        /// A List of source Application Security Group ID's
        /// </summary>
        [Input("sourceApplicationSecurityGroupIds")]
        public Input<string>? SourceApplicationSecurityGroupIds { get; set; }

        /// <summary>
        /// Source Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `source_port_ranges` is not specified.
        /// </summary>
        [Input("sourcePortRange")]
        public Input<string>? SourcePortRange { get; set; }

        [Input("sourcePortRanges")]
        private InputList<string>? _sourcePortRanges;

        /// <summary>
        /// List of source ports or port ranges. This is required if `source_port_range` is not specified.
        /// </summary>
        public InputList<string> SourcePortRanges
        {
            get => _sourcePortRanges ?? (_sourcePortRanges = new InputList<string>());
            set => _sourcePortRanges = value;
        }

        public NetworkSecurityRuleArgs()
        {
        }
    }

    public sealed class NetworkSecurityRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
        /// </summary>
        [Input("access")]
        public Input<string>? Access { get; set; }

        /// <summary>
        /// A description for this rule. Restricted to 140 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// CIDR or destination IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `destination_address_prefixes` is not specified.
        /// </summary>
        [Input("destinationAddressPrefix")]
        public Input<string>? DestinationAddressPrefix { get; set; }

        [Input("destinationAddressPrefixes")]
        private InputList<string>? _destinationAddressPrefixes;

        /// <summary>
        /// List of destination address prefixes. Tags may not be used. This is required if `destination_address_prefix` is not specified.
        /// </summary>
        public InputList<string> DestinationAddressPrefixes
        {
            get => _destinationAddressPrefixes ?? (_destinationAddressPrefixes = new InputList<string>());
            set => _destinationAddressPrefixes = value;
        }

        /// <summary>
        /// A List of destination Application Security Group ID's
        /// </summary>
        [Input("destinationApplicationSecurityGroupIds")]
        public Input<string>? DestinationApplicationSecurityGroupIds { get; set; }

        /// <summary>
        /// Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `destination_port_ranges` is not specified.
        /// </summary>
        [Input("destinationPortRange")]
        public Input<string>? DestinationPortRange { get; set; }

        [Input("destinationPortRanges")]
        private InputList<string>? _destinationPortRanges;

        /// <summary>
        /// List of destination ports or port ranges. This is required if `destination_port_range` is not specified.
        /// </summary>
        public InputList<string> DestinationPortRanges
        {
            get => _destinationPortRanges ?? (_destinationPortRanges = new InputList<string>());
            set => _destinationPortRanges = value;
        }

        /// <summary>
        /// The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// The name of the security rule. This needs to be unique across all Rules in the Network Security Group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Network Security Group that we want to attach the rule to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("networkSecurityGroupName")]
        public Input<string>? NetworkSecurityGroupName { get; set; }

        /// <summary>
        /// Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Network protocol this rule applies to. Possible values include `Tcp`, `Udp`, `Icmp`, or `*` (which matches all).
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Network Security Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// CIDR or source IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `source_address_prefixes` is not specified.
        /// </summary>
        [Input("sourceAddressPrefix")]
        public Input<string>? SourceAddressPrefix { get; set; }

        [Input("sourceAddressPrefixes")]
        private InputList<string>? _sourceAddressPrefixes;

        /// <summary>
        /// List of source address prefixes. Tags may not be used. This is required if `source_address_prefix` is not specified.
        /// </summary>
        public InputList<string> SourceAddressPrefixes
        {
            get => _sourceAddressPrefixes ?? (_sourceAddressPrefixes = new InputList<string>());
            set => _sourceAddressPrefixes = value;
        }

        /// <summary>
        /// A List of source Application Security Group ID's
        /// </summary>
        [Input("sourceApplicationSecurityGroupIds")]
        public Input<string>? SourceApplicationSecurityGroupIds { get; set; }

        /// <summary>
        /// Source Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `source_port_ranges` is not specified.
        /// </summary>
        [Input("sourcePortRange")]
        public Input<string>? SourcePortRange { get; set; }

        [Input("sourcePortRanges")]
        private InputList<string>? _sourcePortRanges;

        /// <summary>
        /// List of source ports or port ranges. This is required if `source_port_range` is not specified.
        /// </summary>
        public InputList<string> SourcePortRanges
        {
            get => _sourcePortRanges ?? (_sourcePortRanges = new InputList<string>());
            set => _sourcePortRanges = value;
        }

        public NetworkSecurityRuleState()
        {
        }
    }
}

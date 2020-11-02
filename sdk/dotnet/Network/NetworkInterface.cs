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
    /// Manages a Network Interface.
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
    ///         var exampleNetworkInterface = new Azure.Network.NetworkInterface("exampleNetworkInterface", new Azure.Network.NetworkInterfaceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             IpConfigurations = 
    ///             {
    ///                 new Azure.Network.Inputs.NetworkInterfaceIpConfigurationArgs
    ///                 {
    ///                     Name = "internal",
    ///                     SubnetId = exampleSubnet.Id,
    ///                     PrivateIpAddressAllocation = "Dynamic",
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
    /// Network Interfaces can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class NetworkInterface : Pulumi.CustomResource
    {
        /// <summary>
        /// If the Virtual Machine using this Network Interface is part of an Availability Set, then this list will have the union of all DNS servers from all Network Interfaces that are part of the Availability Set.
        /// </summary>
        [Output("appliedDnsServers")]
        public Output<ImmutableArray<string>> AppliedDnsServers { get; private set; } = null!;

        /// <summary>
        /// A list of IP Addresses defining the DNS Servers which should be used for this Network Interface.
        /// </summary>
        [Output("dnsServers")]
        public Output<ImmutableArray<string>> DnsServers { get; private set; } = null!;

        /// <summary>
        /// Should Accelerated Networking be enabled? Defaults to `false`.
        /// </summary>
        [Output("enableAcceleratedNetworking")]
        public Output<bool?> EnableAcceleratedNetworking { get; private set; } = null!;

        /// <summary>
        /// Should IP Forwarding be enabled? Defaults to `false`.
        /// </summary>
        [Output("enableIpForwarding")]
        public Output<bool?> EnableIpForwarding { get; private set; } = null!;

        /// <summary>
        /// The (relative) DNS Name used for internal communications between Virtual Machines in the same Virtual Network.
        /// </summary>
        [Output("internalDnsNameLabel")]
        public Output<string> InternalDnsNameLabel { get; private set; } = null!;

        /// <summary>
        /// Even if `internal_dns_name_label` is not specified, a DNS entry is created for the primary NIC of the VM. This DNS name can be constructed by concatenating the VM name with the value of `internal_domain_name_suffix`.
        /// </summary>
        [Output("internalDomainNameSuffix")]
        public Output<string> InternalDomainNameSuffix { get; private set; } = null!;

        /// <summary>
        /// One or more `ip_configuration` blocks as defined below.
        /// </summary>
        [Output("ipConfigurations")]
        public Output<ImmutableArray<Outputs.NetworkInterfaceIpConfiguration>> IpConfigurations { get; private set; } = null!;

        /// <summary>
        /// The location where the Network Interface should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The Media Access Control (MAC) Address of the Network Interface.
        /// </summary>
        [Output("macAddress")]
        public Output<string> MacAddress { get; private set; } = null!;

        /// <summary>
        /// The name of the Network Interface. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Static IP Address which should be used.
        /// </summary>
        [Output("privateIpAddress")]
        public Output<string> PrivateIpAddress { get; private set; } = null!;

        /// <summary>
        /// The private IP addresses of the network interface.
        /// </summary>
        [Output("privateIpAddresses")]
        public Output<ImmutableArray<string>> PrivateIpAddresses { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which to create the Network Interface. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the Virtual Machine which this Network Interface is connected to.
        /// </summary>
        [Output("virtualMachineId")]
        public Output<string> VirtualMachineId { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkInterface resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkInterface(string name, NetworkInterfaceArgs args, CustomResourceOptions? options = null)
            : base("azure:network/networkInterface:NetworkInterface", name, args ?? new NetworkInterfaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkInterface(string name, Input<string> id, NetworkInterfaceState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/networkInterface:NetworkInterface", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkInterface resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkInterface Get(string name, Input<string> id, NetworkInterfaceState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkInterface(name, id, state, options);
        }
    }

    public sealed class NetworkInterfaceArgs : Pulumi.ResourceArgs
    {
        [Input("dnsServers")]
        private InputList<string>? _dnsServers;

        /// <summary>
        /// A list of IP Addresses defining the DNS Servers which should be used for this Network Interface.
        /// </summary>
        public InputList<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new InputList<string>());
            set => _dnsServers = value;
        }

        /// <summary>
        /// Should Accelerated Networking be enabled? Defaults to `false`.
        /// </summary>
        [Input("enableAcceleratedNetworking")]
        public Input<bool>? EnableAcceleratedNetworking { get; set; }

        /// <summary>
        /// Should IP Forwarding be enabled? Defaults to `false`.
        /// </summary>
        [Input("enableIpForwarding")]
        public Input<bool>? EnableIpForwarding { get; set; }

        /// <summary>
        /// The (relative) DNS Name used for internal communications between Virtual Machines in the same Virtual Network.
        /// </summary>
        [Input("internalDnsNameLabel")]
        public Input<string>? InternalDnsNameLabel { get; set; }

        [Input("ipConfigurations", required: true)]
        private InputList<Inputs.NetworkInterfaceIpConfigurationArgs>? _ipConfigurations;

        /// <summary>
        /// One or more `ip_configuration` blocks as defined below.
        /// </summary>
        public InputList<Inputs.NetworkInterfaceIpConfigurationArgs> IpConfigurations
        {
            get => _ipConfigurations ?? (_ipConfigurations = new InputList<Inputs.NetworkInterfaceIpConfigurationArgs>());
            set => _ipConfigurations = value;
        }

        /// <summary>
        /// The location where the Network Interface should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Network Interface. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group in which to create the Network Interface. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

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

        public NetworkInterfaceArgs()
        {
        }
    }

    public sealed class NetworkInterfaceState : Pulumi.ResourceArgs
    {
        [Input("appliedDnsServers")]
        private InputList<string>? _appliedDnsServers;

        /// <summary>
        /// If the Virtual Machine using this Network Interface is part of an Availability Set, then this list will have the union of all DNS servers from all Network Interfaces that are part of the Availability Set.
        /// </summary>
        public InputList<string> AppliedDnsServers
        {
            get => _appliedDnsServers ?? (_appliedDnsServers = new InputList<string>());
            set => _appliedDnsServers = value;
        }

        [Input("dnsServers")]
        private InputList<string>? _dnsServers;

        /// <summary>
        /// A list of IP Addresses defining the DNS Servers which should be used for this Network Interface.
        /// </summary>
        public InputList<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new InputList<string>());
            set => _dnsServers = value;
        }

        /// <summary>
        /// Should Accelerated Networking be enabled? Defaults to `false`.
        /// </summary>
        [Input("enableAcceleratedNetworking")]
        public Input<bool>? EnableAcceleratedNetworking { get; set; }

        /// <summary>
        /// Should IP Forwarding be enabled? Defaults to `false`.
        /// </summary>
        [Input("enableIpForwarding")]
        public Input<bool>? EnableIpForwarding { get; set; }

        /// <summary>
        /// The (relative) DNS Name used for internal communications between Virtual Machines in the same Virtual Network.
        /// </summary>
        [Input("internalDnsNameLabel")]
        public Input<string>? InternalDnsNameLabel { get; set; }

        /// <summary>
        /// Even if `internal_dns_name_label` is not specified, a DNS entry is created for the primary NIC of the VM. This DNS name can be constructed by concatenating the VM name with the value of `internal_domain_name_suffix`.
        /// </summary>
        [Input("internalDomainNameSuffix")]
        public Input<string>? InternalDomainNameSuffix { get; set; }

        [Input("ipConfigurations")]
        private InputList<Inputs.NetworkInterfaceIpConfigurationGetArgs>? _ipConfigurations;

        /// <summary>
        /// One or more `ip_configuration` blocks as defined below.
        /// </summary>
        public InputList<Inputs.NetworkInterfaceIpConfigurationGetArgs> IpConfigurations
        {
            get => _ipConfigurations ?? (_ipConfigurations = new InputList<Inputs.NetworkInterfaceIpConfigurationGetArgs>());
            set => _ipConfigurations = value;
        }

        /// <summary>
        /// The location where the Network Interface should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The Media Access Control (MAC) Address of the Network Interface.
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// The name of the Network Interface. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Static IP Address which should be used.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        [Input("privateIpAddresses")]
        private InputList<string>? _privateIpAddresses;

        /// <summary>
        /// The private IP addresses of the network interface.
        /// </summary>
        public InputList<string> PrivateIpAddresses
        {
            get => _privateIpAddresses ?? (_privateIpAddresses = new InputList<string>());
            set => _privateIpAddresses = value;
        }

        /// <summary>
        /// The name of the Resource Group in which to create the Network Interface. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

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
        /// The ID of the Virtual Machine which this Network Interface is connected to.
        /// </summary>
        [Input("virtualMachineId")]
        public Input<string>? VirtualMachineId { get; set; }

        public NetworkInterfaceState()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.PrivateLink
{
    /// <summary>
    /// Manages a Private Endpoint.
    /// 
    /// Azure Private Endpoint is a network interface that connects you privately and securely to a service powered by Azure Private Link. Private Endpoint uses a private IP address from your VNet, effectively bringing the service into your VNet. The service could be an Azure service such as Azure Storage, SQL, etc. or your own Private Link Service.
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
    ///         var service = new Azure.Network.Subnet("service", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualNetworkName = exampleVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.0.1.0/24",
    ///             },
    ///             EnforcePrivateLinkServiceNetworkPolicies = true,
    ///         });
    ///         var endpoint = new Azure.Network.Subnet("endpoint", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualNetworkName = exampleVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.0.2.0/24",
    ///             },
    ///             EnforcePrivateLinkEndpointNetworkPolicies = true,
    ///         });
    ///         var examplePublicIp = new Azure.Network.PublicIp("examplePublicIp", new Azure.Network.PublicIpArgs
    ///         {
    ///             Sku = "Standard",
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AllocationMethod = "Static",
    ///         });
    ///         var exampleLoadBalancer = new Azure.Lb.LoadBalancer("exampleLoadBalancer", new Azure.Lb.LoadBalancerArgs
    ///         {
    ///             Sku = "Standard",
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             FrontendIpConfigurations = 
    ///             {
    ///                 new Azure.Lb.Inputs.LoadBalancerFrontendIpConfigurationArgs
    ///                 {
    ///                     Name = examplePublicIp.Name,
    ///                     PublicIpAddressId = examplePublicIp.Id,
    ///                 },
    ///             },
    ///         });
    ///         var exampleLinkService = new Azure.PrivateDns.LinkService("exampleLinkService", new Azure.PrivateDns.LinkServiceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             NatIpConfigurations = 
    ///             {
    ///                 new Azure.PrivateDns.Inputs.LinkServiceNatIpConfigurationArgs
    ///                 {
    ///                     Name = examplePublicIp.Name,
    ///                     Primary = true,
    ///                     SubnetId = service.Id,
    ///                 },
    ///             },
    ///             LoadBalancerFrontendIpConfigurationIds = 
    ///             {
    ///                 exampleLoadBalancer.FrontendIpConfigurations.Apply(frontendIpConfigurations =&gt; frontendIpConfigurations?[0]?.Id),
    ///             },
    ///         });
    ///         var exampleEndpoint = new Azure.PrivateLink.Endpoint("exampleEndpoint", new Azure.PrivateLink.EndpointArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             SubnetId = endpoint.Id,
    ///             PrivateServiceConnection = new Azure.PrivateLink.Inputs.EndpointPrivateServiceConnectionArgs
    ///             {
    ///                 Name = "example-privateserviceconnection",
    ///                 PrivateConnectionResourceId = exampleLinkService.Id,
    ///                 IsManualConnection = false,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// Using a Private Link Service Alias with existing resources:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var rg = Output.Create(Azure.Core.GetResourceGroup.InvokeAsync(new Azure.Core.GetResourceGroupArgs
    ///         {
    ///             Name = "example-resources",
    ///         }));
    ///         var vnet = rg.Apply(rg =&gt; Output.Create(Azure.Network.GetVirtualNetwork.InvokeAsync(new Azure.Network.GetVirtualNetworkArgs
    ///         {
    ///             Name = "example-network",
    ///             ResourceGroupName = rg.Name,
    ///         })));
    ///         var subnet = Output.Tuple(vnet, rg).Apply(values =&gt;
    ///         {
    ///             var vnet = values.Item1;
    ///             var rg = values.Item2;
    ///             return Output.Create(Azure.Network.GetSubnet.InvokeAsync(new Azure.Network.GetSubnetArgs
    ///             {
    ///                 Name = "default",
    ///                 VirtualNetworkName = vnet.Name,
    ///                 ResourceGroupName = rg.Name,
    ///             }));
    ///         });
    ///         var example = new Azure.PrivateLink.Endpoint("example", new Azure.PrivateLink.EndpointArgs
    ///         {
    ///             Location = rg.Apply(rg =&gt; rg.Location),
    ///             ResourceGroupName = rg.Apply(rg =&gt; rg.Name),
    ///             SubnetId = subnet.Apply(subnet =&gt; subnet.Id),
    ///             PrivateServiceConnection = new Azure.PrivateLink.Inputs.EndpointPrivateServiceConnectionArgs
    ///             {
    ///                 Name = "example-privateserviceconnection",
    ///                 PrivateConnectionResourceAlias = "example-privatelinkservice.d20286c8-4ea5-11eb-9584-8f53157226c6.centralus.azure.privatelinkservice",
    ///                 IsManualConnection = true,
    ///                 RequestMessage = "PL",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Private Endpoints can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:privatelink/endpoint:Endpoint example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/privateEndpoints/endpoint1
    /// ```
    /// </summary>
    [AzureResourceType("azure:privatelink/endpoint:Endpoint")]
    public partial class Endpoint : Pulumi.CustomResource
    {
        [Output("customDnsConfigs")]
        public Output<ImmutableArray<Outputs.EndpointCustomDnsConfig>> CustomDnsConfigs { get; private set; } = null!;

        /// <summary>
        /// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the Name of the Private Endpoint. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("privateDnsZoneConfigs")]
        public Output<ImmutableArray<Outputs.EndpointPrivateDnsZoneConfig>> PrivateDnsZoneConfigs { get; private set; } = null!;

        /// <summary>
        /// A `private_dns_zone_group` block as defined below.
        /// </summary>
        [Output("privateDnsZoneGroup")]
        public Output<Outputs.EndpointPrivateDnsZoneGroup?> PrivateDnsZoneGroup { get; private set; } = null!;

        /// <summary>
        /// A `private_service_connection` block as defined below.
        /// </summary>
        [Output("privateServiceConnection")]
        public Output<Outputs.EndpointPrivateServiceConnection> PrivateServiceConnection { get; private set; } = null!;

        /// <summary>
        /// Specifies the Name of the Resource Group within which the Private Endpoint should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The ID of the Subnet from which Private IP Addresses will be allocated for this Private Endpoint. Changing this forces a new resource to be created.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Endpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Endpoint(string name, EndpointArgs args, CustomResourceOptions? options = null)
            : base("azure:privatelink/endpoint:Endpoint", name, args ?? new EndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Endpoint(string name, Input<string> id, EndpointState? state = null, CustomResourceOptions? options = null)
            : base("azure:privatelink/endpoint:Endpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Endpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Endpoint Get(string name, Input<string> id, EndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new Endpoint(name, id, state, options);
        }
    }

    public sealed class EndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the Name of the Private Endpoint. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `private_dns_zone_group` block as defined below.
        /// </summary>
        [Input("privateDnsZoneGroup")]
        public Input<Inputs.EndpointPrivateDnsZoneGroupArgs>? PrivateDnsZoneGroup { get; set; }

        /// <summary>
        /// A `private_service_connection` block as defined below.
        /// </summary>
        [Input("privateServiceConnection", required: true)]
        public Input<Inputs.EndpointPrivateServiceConnectionArgs> PrivateServiceConnection { get; set; } = null!;

        /// <summary>
        /// Specifies the Name of the Resource Group within which the Private Endpoint should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The ID of the Subnet from which Private IP Addresses will be allocated for this Private Endpoint. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

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

        public EndpointArgs()
        {
        }
    }

    public sealed class EndpointState : Pulumi.ResourceArgs
    {
        [Input("customDnsConfigs")]
        private InputList<Inputs.EndpointCustomDnsConfigGetArgs>? _customDnsConfigs;
        public InputList<Inputs.EndpointCustomDnsConfigGetArgs> CustomDnsConfigs
        {
            get => _customDnsConfigs ?? (_customDnsConfigs = new InputList<Inputs.EndpointCustomDnsConfigGetArgs>());
            set => _customDnsConfigs = value;
        }

        /// <summary>
        /// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the Name of the Private Endpoint. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("privateDnsZoneConfigs")]
        private InputList<Inputs.EndpointPrivateDnsZoneConfigGetArgs>? _privateDnsZoneConfigs;
        public InputList<Inputs.EndpointPrivateDnsZoneConfigGetArgs> PrivateDnsZoneConfigs
        {
            get => _privateDnsZoneConfigs ?? (_privateDnsZoneConfigs = new InputList<Inputs.EndpointPrivateDnsZoneConfigGetArgs>());
            set => _privateDnsZoneConfigs = value;
        }

        /// <summary>
        /// A `private_dns_zone_group` block as defined below.
        /// </summary>
        [Input("privateDnsZoneGroup")]
        public Input<Inputs.EndpointPrivateDnsZoneGroupGetArgs>? PrivateDnsZoneGroup { get; set; }

        /// <summary>
        /// A `private_service_connection` block as defined below.
        /// </summary>
        [Input("privateServiceConnection")]
        public Input<Inputs.EndpointPrivateServiceConnectionGetArgs>? PrivateServiceConnection { get; set; }

        /// <summary>
        /// Specifies the Name of the Resource Group within which the Private Endpoint should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The ID of the Subnet from which Private IP Addresses will be allocated for this Private Endpoint. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

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

        public EndpointState()
        {
        }
    }
}

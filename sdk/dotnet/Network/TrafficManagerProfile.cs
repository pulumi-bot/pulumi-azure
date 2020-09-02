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
    /// Manages a Traffic Manager Profile to which multiple endpoints can be attached.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// using Random = Pulumi.Random;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var server = new Random.RandomId("server", new Random.RandomIdArgs
    ///         {
    ///             Keepers = 
    ///             {
    ///                 { "azi_id", 1 },
    ///             },
    ///             ByteLength = 8,
    ///         });
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West US",
    ///         });
    ///         var exampleTrafficManagerProfile = new Azure.Network.TrafficManagerProfile("exampleTrafficManagerProfile", new Azure.Network.TrafficManagerProfileArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             TrafficRoutingMethod = "Weighted",
    ///             DnsConfig = new Azure.Network.Inputs.TrafficManagerProfileDnsConfigArgs
    ///             {
    ///                 RelativeName = server.Hex,
    ///                 Ttl = 100,
    ///             },
    ///             MonitorConfig = new Azure.Network.Inputs.TrafficManagerProfileMonitorConfigArgs
    ///             {
    ///                 Protocol = "http",
    ///                 Port = 80,
    ///                 Path = "/",
    ///                 IntervalInSeconds = 30,
    ///                 TimeoutInSeconds = 9,
    ///                 ToleratedNumberOfFailures = 3,
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "environment", "Production" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class TrafficManagerProfile : Pulumi.CustomResource
    {
        /// <summary>
        /// This block specifies the DNS configuration of the Profile, it supports the fields documented below.
        /// </summary>
        [Output("dnsConfig")]
        public Output<Outputs.TrafficManagerProfileDnsConfig> DnsConfig { get; private set; } = null!;

        /// <summary>
        /// The FQDN of the created Profile.
        /// </summary>
        [Output("fqdn")]
        public Output<string> Fqdn { get; private set; } = null!;

        /// <summary>
        /// This block specifies the Endpoint monitoring configuration for the Profile, it supports the fields documented below.
        /// </summary>
        [Output("monitorConfig")]
        public Output<Outputs.TrafficManagerProfileMonitorConfig> MonitorConfig { get; private set; } = null!;

        /// <summary>
        /// The name of the Traffic Manager profile. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The status of the profile, can be set to either `Enabled` or `Disabled`. Defaults to `Enabled`.
        /// </summary>
        [Output("profileStatus")]
        public Output<string> ProfileStatus { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Traffic Manager profile.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies the algorithm used to route traffic, possible values are:
        /// </summary>
        [Output("trafficRoutingMethod")]
        public Output<string> TrafficRoutingMethod { get; private set; } = null!;


        /// <summary>
        /// Create a TrafficManagerProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TrafficManagerProfile(string name, TrafficManagerProfileArgs args, CustomResourceOptions? options = null)
            : base("azure:network/trafficManagerProfile:TrafficManagerProfile", name, args ?? new TrafficManagerProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TrafficManagerProfile(string name, Input<string> id, TrafficManagerProfileState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/trafficManagerProfile:TrafficManagerProfile", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure:trafficmanager/profile:Profile"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TrafficManagerProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TrafficManagerProfile Get(string name, Input<string> id, TrafficManagerProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new TrafficManagerProfile(name, id, state, options);
        }
    }

    public sealed class TrafficManagerProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// This block specifies the DNS configuration of the Profile, it supports the fields documented below.
        /// </summary>
        [Input("dnsConfig", required: true)]
        public Input<Inputs.TrafficManagerProfileDnsConfigArgs> DnsConfig { get; set; } = null!;

        /// <summary>
        /// This block specifies the Endpoint monitoring configuration for the Profile, it supports the fields documented below.
        /// </summary>
        [Input("monitorConfig", required: true)]
        public Input<Inputs.TrafficManagerProfileMonitorConfigArgs> MonitorConfig { get; set; } = null!;

        /// <summary>
        /// The name of the Traffic Manager profile. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The status of the profile, can be set to either `Enabled` or `Disabled`. Defaults to `Enabled`.
        /// </summary>
        [Input("profileStatus")]
        public Input<string>? ProfileStatus { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Traffic Manager profile.
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

        /// <summary>
        /// Specifies the algorithm used to route traffic, possible values are:
        /// </summary>
        [Input("trafficRoutingMethod", required: true)]
        public Input<string> TrafficRoutingMethod { get; set; } = null!;

        public TrafficManagerProfileArgs()
        {
        }
    }

    public sealed class TrafficManagerProfileState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// This block specifies the DNS configuration of the Profile, it supports the fields documented below.
        /// </summary>
        [Input("dnsConfig")]
        public Input<Inputs.TrafficManagerProfileDnsConfigGetArgs>? DnsConfig { get; set; }

        /// <summary>
        /// The FQDN of the created Profile.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// This block specifies the Endpoint monitoring configuration for the Profile, it supports the fields documented below.
        /// </summary>
        [Input("monitorConfig")]
        public Input<Inputs.TrafficManagerProfileMonitorConfigGetArgs>? MonitorConfig { get; set; }

        /// <summary>
        /// The name of the Traffic Manager profile. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The status of the profile, can be set to either `Enabled` or `Disabled`. Defaults to `Enabled`.
        /// </summary>
        [Input("profileStatus")]
        public Input<string>? ProfileStatus { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Traffic Manager profile.
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
        /// Specifies the algorithm used to route traffic, possible values are:
        /// </summary>
        [Input("trafficRoutingMethod")]
        public Input<string>? TrafficRoutingMethod { get; set; }

        public TrafficManagerProfileState()
        {
        }
    }
}

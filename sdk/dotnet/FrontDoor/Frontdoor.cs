// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.FrontDoor
{
    /// <summary>
    /// Manages an Azure Front Door instance.
    /// 
    /// Azure Front Door Service is Microsoft's highly available and scalable web application acceleration platform and global HTTP(s) load balancer. It provides built-in DDoS protection and application layer security and caching. Front Door enables you to build applications that maximize and automate high-availability and performance for your end-users. Use Front Door with Azure services including Web/Mobile Apps, Cloud Services and Virtual Machines – or combine it with on-premises services for hybrid deployments and smooth cloud migration.
    /// 
    /// Below are some of the key scenarios that Azure Front Door Service addresses:
    /// * Use Front Door to improve application scale and availability with instant multi-region failover
    /// * Use Front Door to improve application performance with SSL offload and routing requests to the fastest available application backend.
    /// * Use Front Door for application layer security and DDoS protection for your application.
    /// </summary>
    public partial class Frontdoor : Pulumi.CustomResource
    {
        /// <summary>
        /// A `backend_pool_health_probe` block as defined below.
        /// </summary>
        [Output("backendPoolHealthProbes")]
        public Output<ImmutableArray<Outputs.FrontdoorBackendPoolHealthProbe>> BackendPoolHealthProbes { get; private set; } = null!;

        /// <summary>
        /// A `backend_pool_load_balancing` block as defined below.
        /// </summary>
        [Output("backendPoolLoadBalancings")]
        public Output<ImmutableArray<Outputs.FrontdoorBackendPoolLoadBalancing>> BackendPoolLoadBalancings { get; private set; } = null!;

        /// <summary>
        /// A `backend_pool` block as defined below.
        /// </summary>
        [Output("backendPools")]
        public Output<ImmutableArray<Outputs.FrontdoorBackendPool>> BackendPools { get; private set; } = null!;

        /// <summary>
        /// Specifies the send and receive timeout on forwarding request to the backend. When the timeout is reached, the request fails and returns. Possible values are between `0` - `240`. Defaults to `60`.
        /// </summary>
        [Output("backendPoolsSendReceiveTimeoutSeconds")]
        public Output<int?> BackendPoolsSendReceiveTimeoutSeconds { get; private set; } = null!;

        /// <summary>
        /// The host that each frontendEndpoint must CNAME to.
        /// </summary>
        [Output("cname")]
        public Output<string> Cname { get; private set; } = null!;

        /// <summary>
        /// Enforce certificate name check on `HTTPS` requests to all backend pools, this setting will have no effect on `HTTP` requests. Permitted values are `true` or `false`.
        /// </summary>
        [Output("enforceBackendPoolsCertificateNameCheck")]
        public Output<bool> EnforceBackendPoolsCertificateNameCheck { get; private set; } = null!;

        /// <summary>
        /// A friendly name for the Front Door service.
        /// </summary>
        [Output("friendlyName")]
        public Output<string?> FriendlyName { get; private set; } = null!;

        /// <summary>
        /// A `frontend_endpoint` block as defined below.
        /// </summary>
        [Output("frontendEndpoints")]
        public Output<ImmutableArray<Outputs.FrontdoorFrontendEndpoint>> FrontendEndpoints { get; private set; } = null!;

        /// <summary>
        /// The unique ID of the Front Door which is embedded into the incoming headers `X-Azure-FDID` attribute and maybe used to filter traffic sent by the Front Door to your backend.
        /// </summary>
        [Output("headerFrontdoorId")]
        public Output<string> HeaderFrontdoorId { get; private set; } = null!;

        /// <summary>
        /// Should the Front Door Load Balancer be Enabled? Defaults to `true`.
        /// </summary>
        [Output("loadBalancerEnabled")]
        public Output<bool?> LoadBalancerEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Front Door service. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group in which the Front Door service should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A `routing_rule` block as defined below.
        /// </summary>
        [Output("routingRules")]
        public Output<ImmutableArray<Outputs.FrontdoorRoutingRule>> RoutingRules { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Frontdoor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Frontdoor(string name, FrontdoorArgs args, CustomResourceOptions? options = null)
            : base("azure:frontdoor/frontdoor:Frontdoor", name, args ?? new FrontdoorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Frontdoor(string name, Input<string> id, FrontdoorState? state = null, CustomResourceOptions? options = null)
            : base("azure:frontdoor/frontdoor:Frontdoor", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Frontdoor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Frontdoor Get(string name, Input<string> id, FrontdoorState? state = null, CustomResourceOptions? options = null)
        {
            return new Frontdoor(name, id, state, options);
        }
    }

    public sealed class FrontdoorArgs : Pulumi.ResourceArgs
    {
        [Input("backendPoolHealthProbes", required: true)]
        private InputList<Inputs.FrontdoorBackendPoolHealthProbeArgs>? _backendPoolHealthProbes;

        /// <summary>
        /// A `backend_pool_health_probe` block as defined below.
        /// </summary>
        public InputList<Inputs.FrontdoorBackendPoolHealthProbeArgs> BackendPoolHealthProbes
        {
            get => _backendPoolHealthProbes ?? (_backendPoolHealthProbes = new InputList<Inputs.FrontdoorBackendPoolHealthProbeArgs>());
            set => _backendPoolHealthProbes = value;
        }

        [Input("backendPoolLoadBalancings", required: true)]
        private InputList<Inputs.FrontdoorBackendPoolLoadBalancingArgs>? _backendPoolLoadBalancings;

        /// <summary>
        /// A `backend_pool_load_balancing` block as defined below.
        /// </summary>
        public InputList<Inputs.FrontdoorBackendPoolLoadBalancingArgs> BackendPoolLoadBalancings
        {
            get => _backendPoolLoadBalancings ?? (_backendPoolLoadBalancings = new InputList<Inputs.FrontdoorBackendPoolLoadBalancingArgs>());
            set => _backendPoolLoadBalancings = value;
        }

        [Input("backendPools", required: true)]
        private InputList<Inputs.FrontdoorBackendPoolArgs>? _backendPools;

        /// <summary>
        /// A `backend_pool` block as defined below.
        /// </summary>
        public InputList<Inputs.FrontdoorBackendPoolArgs> BackendPools
        {
            get => _backendPools ?? (_backendPools = new InputList<Inputs.FrontdoorBackendPoolArgs>());
            set => _backendPools = value;
        }

        /// <summary>
        /// Specifies the send and receive timeout on forwarding request to the backend. When the timeout is reached, the request fails and returns. Possible values are between `0` - `240`. Defaults to `60`.
        /// </summary>
        [Input("backendPoolsSendReceiveTimeoutSeconds")]
        public Input<int>? BackendPoolsSendReceiveTimeoutSeconds { get; set; }

        /// <summary>
        /// Enforce certificate name check on `HTTPS` requests to all backend pools, this setting will have no effect on `HTTP` requests. Permitted values are `true` or `false`.
        /// </summary>
        [Input("enforceBackendPoolsCertificateNameCheck", required: true)]
        public Input<bool> EnforceBackendPoolsCertificateNameCheck { get; set; } = null!;

        /// <summary>
        /// A friendly name for the Front Door service.
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        [Input("frontendEndpoints", required: true)]
        private InputList<Inputs.FrontdoorFrontendEndpointArgs>? _frontendEndpoints;

        /// <summary>
        /// A `frontend_endpoint` block as defined below.
        /// </summary>
        public InputList<Inputs.FrontdoorFrontendEndpointArgs> FrontendEndpoints
        {
            get => _frontendEndpoints ?? (_frontendEndpoints = new InputList<Inputs.FrontdoorFrontendEndpointArgs>());
            set => _frontendEndpoints = value;
        }

        /// <summary>
        /// Should the Front Door Load Balancer be Enabled? Defaults to `true`.
        /// </summary>
        [Input("loadBalancerEnabled")]
        public Input<bool>? LoadBalancerEnabled { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Front Door service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group in which the Front Door service should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("routingRules", required: true)]
        private InputList<Inputs.FrontdoorRoutingRuleArgs>? _routingRules;

        /// <summary>
        /// A `routing_rule` block as defined below.
        /// </summary>
        public InputList<Inputs.FrontdoorRoutingRuleArgs> RoutingRules
        {
            get => _routingRules ?? (_routingRules = new InputList<Inputs.FrontdoorRoutingRuleArgs>());
            set => _routingRules = value;
        }

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

        public FrontdoorArgs()
        {
        }
    }

    public sealed class FrontdoorState : Pulumi.ResourceArgs
    {
        [Input("backendPoolHealthProbes")]
        private InputList<Inputs.FrontdoorBackendPoolHealthProbeGetArgs>? _backendPoolHealthProbes;

        /// <summary>
        /// A `backend_pool_health_probe` block as defined below.
        /// </summary>
        public InputList<Inputs.FrontdoorBackendPoolHealthProbeGetArgs> BackendPoolHealthProbes
        {
            get => _backendPoolHealthProbes ?? (_backendPoolHealthProbes = new InputList<Inputs.FrontdoorBackendPoolHealthProbeGetArgs>());
            set => _backendPoolHealthProbes = value;
        }

        [Input("backendPoolLoadBalancings")]
        private InputList<Inputs.FrontdoorBackendPoolLoadBalancingGetArgs>? _backendPoolLoadBalancings;

        /// <summary>
        /// A `backend_pool_load_balancing` block as defined below.
        /// </summary>
        public InputList<Inputs.FrontdoorBackendPoolLoadBalancingGetArgs> BackendPoolLoadBalancings
        {
            get => _backendPoolLoadBalancings ?? (_backendPoolLoadBalancings = new InputList<Inputs.FrontdoorBackendPoolLoadBalancingGetArgs>());
            set => _backendPoolLoadBalancings = value;
        }

        [Input("backendPools")]
        private InputList<Inputs.FrontdoorBackendPoolGetArgs>? _backendPools;

        /// <summary>
        /// A `backend_pool` block as defined below.
        /// </summary>
        public InputList<Inputs.FrontdoorBackendPoolGetArgs> BackendPools
        {
            get => _backendPools ?? (_backendPools = new InputList<Inputs.FrontdoorBackendPoolGetArgs>());
            set => _backendPools = value;
        }

        /// <summary>
        /// Specifies the send and receive timeout on forwarding request to the backend. When the timeout is reached, the request fails and returns. Possible values are between `0` - `240`. Defaults to `60`.
        /// </summary>
        [Input("backendPoolsSendReceiveTimeoutSeconds")]
        public Input<int>? BackendPoolsSendReceiveTimeoutSeconds { get; set; }

        /// <summary>
        /// The host that each frontendEndpoint must CNAME to.
        /// </summary>
        [Input("cname")]
        public Input<string>? Cname { get; set; }

        /// <summary>
        /// Enforce certificate name check on `HTTPS` requests to all backend pools, this setting will have no effect on `HTTP` requests. Permitted values are `true` or `false`.
        /// </summary>
        [Input("enforceBackendPoolsCertificateNameCheck")]
        public Input<bool>? EnforceBackendPoolsCertificateNameCheck { get; set; }

        /// <summary>
        /// A friendly name for the Front Door service.
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        [Input("frontendEndpoints")]
        private InputList<Inputs.FrontdoorFrontendEndpointGetArgs>? _frontendEndpoints;

        /// <summary>
        /// A `frontend_endpoint` block as defined below.
        /// </summary>
        public InputList<Inputs.FrontdoorFrontendEndpointGetArgs> FrontendEndpoints
        {
            get => _frontendEndpoints ?? (_frontendEndpoints = new InputList<Inputs.FrontdoorFrontendEndpointGetArgs>());
            set => _frontendEndpoints = value;
        }

        /// <summary>
        /// The unique ID of the Front Door which is embedded into the incoming headers `X-Azure-FDID` attribute and maybe used to filter traffic sent by the Front Door to your backend.
        /// </summary>
        [Input("headerFrontdoorId")]
        public Input<string>? HeaderFrontdoorId { get; set; }

        /// <summary>
        /// Should the Front Door Load Balancer be Enabled? Defaults to `true`.
        /// </summary>
        [Input("loadBalancerEnabled")]
        public Input<bool>? LoadBalancerEnabled { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Front Door service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group in which the Front Door service should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("routingRules")]
        private InputList<Inputs.FrontdoorRoutingRuleGetArgs>? _routingRules;

        /// <summary>
        /// A `routing_rule` block as defined below.
        /// </summary>
        public InputList<Inputs.FrontdoorRoutingRuleGetArgs> RoutingRules
        {
            get => _routingRules ?? (_routingRules = new InputList<Inputs.FrontdoorRoutingRuleGetArgs>());
            set => _routingRules = value;
        }

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

        public FrontdoorState()
        {
        }
    }
}

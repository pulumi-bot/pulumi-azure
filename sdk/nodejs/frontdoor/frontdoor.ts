// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an Azure Front Door instance.
 *
 * Azure Front Door Service is Microsoft's highly available and scalable web application acceleration platform and global HTTP(s) load balancer. It provides built-in DDoS protection and application layer security and caching. Front Door enables you to build applications that maximize and automate high-availability and performance for your end-users. Use Front Door with Azure services including Web/Mobile Apps, Cloud Services and Virtual Machines – or combine it with on-premises services for hybrid deployments and smooth cloud migration.
 *
 * Below are some of the key scenarios that Azure Front Door Service addresses:
 * * Use Front Door to improve application scale and availability with instant multi-region failover
 * * Use Front Door to improve application performance with SSL offload and routing requests to the fastest available application backend.
 * * Use Front Door for application layer security and DDoS protection for your application.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "EastUS2"});
 * const exampleFrontdoor = new azure.frontdoor.Frontdoor("exampleFrontdoor", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     enforceBackendPoolsCertificateNameCheck: false,
 *     routing_rule: [{
 *         name: "exampleRoutingRule1",
 *         acceptedProtocols: [
 *             "Http",
 *             "Https",
 *         ],
 *         patternsToMatches: ["/*"],
 *         frontendEndpoints: ["exampleFrontendEndpoint1"],
 *         forwarding_configuration: {
 *             forwardingProtocol: "MatchRequest",
 *             backendPoolName: "exampleBackendBing",
 *         },
 *     }],
 *     backend_pool_load_balancing: [{
 *         name: "exampleLoadBalancingSettings1",
 *     }],
 *     backend_pool_health_probe: [{
 *         name: "exampleHealthProbeSetting1",
 *     }],
 *     backend_pool: [{
 *         name: "exampleBackendBing",
 *         backend: [{
 *             hostHeader: "www.bing.com",
 *             address: "www.bing.com",
 *             httpPort: 80,
 *             httpsPort: 443,
 *         }],
 *         loadBalancingName: "exampleLoadBalancingSettings1",
 *         healthProbeName: "exampleHealthProbeSetting1",
 *     }],
 *     frontend_endpoint: [{
 *         name: "exampleFrontendEndpoint1",
 *         hostName: "example-FrontDoor.azurefd.net",
 *         customHttpsProvisioningEnabled: false,
 *     }],
 * });
 * ```
 */
export class Frontdoor extends pulumi.CustomResource {
    /**
     * Get an existing Frontdoor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FrontdoorState, opts?: pulumi.CustomResourceOptions): Frontdoor {
        return new Frontdoor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:frontdoor/frontdoor:Frontdoor';

    /**
     * Returns true if the given object is an instance of Frontdoor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Frontdoor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Frontdoor.__pulumiType;
    }

    /**
     * A `backendPoolHealthProbe` block as defined below.
     */
    public readonly backendPoolHealthProbes!: pulumi.Output<outputs.frontdoor.FrontdoorBackendPoolHealthProbe[]>;
    /**
     * A `backendPoolLoadBalancing` block as defined below.
     */
    public readonly backendPoolLoadBalancings!: pulumi.Output<outputs.frontdoor.FrontdoorBackendPoolLoadBalancing[]>;
    /**
     * A `backendPool` block as defined below.
     */
    public readonly backendPools!: pulumi.Output<outputs.frontdoor.FrontdoorBackendPool[]>;
    /**
     * Specifies the send and receive timeout on forwarding request to the backend. When the timeout is reached, the request fails and returns. Possible values are between `0` - `240`. Defaults to `60`.
     */
    public readonly backendPoolsSendReceiveTimeoutSeconds!: pulumi.Output<number | undefined>;
    /**
     * The host that each frontendEndpoint must CNAME to.
     */
    public /*out*/ readonly cname!: pulumi.Output<string>;
    /**
     * Enforce certificate name check on `HTTPS` requests to all backend pools, this setting will have no effect on `HTTP` requests. Permitted values are `true` or `false`.
     */
    public readonly enforceBackendPoolsCertificateNameCheck!: pulumi.Output<boolean>;
    /**
     * A friendly name for the Front Door service.
     */
    public readonly friendlyName!: pulumi.Output<string | undefined>;
    /**
     * A `frontendEndpoint` block as defined below.
     */
    public readonly frontendEndpoints!: pulumi.Output<outputs.frontdoor.FrontdoorFrontendEndpoint[]>;
    /**
     * The unique ID of the Front Door which is embedded into the incoming headers `X-Azure-FDID` attribute and maybe used to filter traffic sent by the Front Door to your backend.
     */
    public /*out*/ readonly headerFrontdoorId!: pulumi.Output<string>;
    /**
     * Should the Front Door Load Balancer be Enabled? Defaults to `true`.
     */
    public readonly loadBalancerEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * @deprecated Due to the service's API changing 'location' must now always be set to 'Global' for new resources, however if the Front Door service was created prior 2020/03/10 it may continue to exist in a specific current location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Front Door service. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the name of the Resource Group in which the Front Door service should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `routingRule` block as defined below.
     */
    public readonly routingRules!: pulumi.Output<outputs.frontdoor.FrontdoorRoutingRule[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Frontdoor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FrontdoorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FrontdoorArgs | FrontdoorState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FrontdoorState | undefined;
            inputs["backendPoolHealthProbes"] = state ? state.backendPoolHealthProbes : undefined;
            inputs["backendPoolLoadBalancings"] = state ? state.backendPoolLoadBalancings : undefined;
            inputs["backendPools"] = state ? state.backendPools : undefined;
            inputs["backendPoolsSendReceiveTimeoutSeconds"] = state ? state.backendPoolsSendReceiveTimeoutSeconds : undefined;
            inputs["cname"] = state ? state.cname : undefined;
            inputs["enforceBackendPoolsCertificateNameCheck"] = state ? state.enforceBackendPoolsCertificateNameCheck : undefined;
            inputs["friendlyName"] = state ? state.friendlyName : undefined;
            inputs["frontendEndpoints"] = state ? state.frontendEndpoints : undefined;
            inputs["headerFrontdoorId"] = state ? state.headerFrontdoorId : undefined;
            inputs["loadBalancerEnabled"] = state ? state.loadBalancerEnabled : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["routingRules"] = state ? state.routingRules : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as FrontdoorArgs | undefined;
            if (!args || args.backendPoolHealthProbes === undefined) {
                throw new Error("Missing required property 'backendPoolHealthProbes'");
            }
            if (!args || args.backendPoolLoadBalancings === undefined) {
                throw new Error("Missing required property 'backendPoolLoadBalancings'");
            }
            if (!args || args.backendPools === undefined) {
                throw new Error("Missing required property 'backendPools'");
            }
            if (!args || args.enforceBackendPoolsCertificateNameCheck === undefined) {
                throw new Error("Missing required property 'enforceBackendPoolsCertificateNameCheck'");
            }
            if (!args || args.frontendEndpoints === undefined) {
                throw new Error("Missing required property 'frontendEndpoints'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.routingRules === undefined) {
                throw new Error("Missing required property 'routingRules'");
            }
            inputs["backendPoolHealthProbes"] = args ? args.backendPoolHealthProbes : undefined;
            inputs["backendPoolLoadBalancings"] = args ? args.backendPoolLoadBalancings : undefined;
            inputs["backendPools"] = args ? args.backendPools : undefined;
            inputs["backendPoolsSendReceiveTimeoutSeconds"] = args ? args.backendPoolsSendReceiveTimeoutSeconds : undefined;
            inputs["enforceBackendPoolsCertificateNameCheck"] = args ? args.enforceBackendPoolsCertificateNameCheck : undefined;
            inputs["friendlyName"] = args ? args.friendlyName : undefined;
            inputs["frontendEndpoints"] = args ? args.frontendEndpoints : undefined;
            inputs["loadBalancerEnabled"] = args ? args.loadBalancerEnabled : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["routingRules"] = args ? args.routingRules : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["cname"] = undefined /*out*/;
            inputs["headerFrontdoorId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Frontdoor.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Frontdoor resources.
 */
export interface FrontdoorState {
    /**
     * A `backendPoolHealthProbe` block as defined below.
     */
    readonly backendPoolHealthProbes?: pulumi.Input<pulumi.Input<inputs.frontdoor.FrontdoorBackendPoolHealthProbe>[]>;
    /**
     * A `backendPoolLoadBalancing` block as defined below.
     */
    readonly backendPoolLoadBalancings?: pulumi.Input<pulumi.Input<inputs.frontdoor.FrontdoorBackendPoolLoadBalancing>[]>;
    /**
     * A `backendPool` block as defined below.
     */
    readonly backendPools?: pulumi.Input<pulumi.Input<inputs.frontdoor.FrontdoorBackendPool>[]>;
    /**
     * Specifies the send and receive timeout on forwarding request to the backend. When the timeout is reached, the request fails and returns. Possible values are between `0` - `240`. Defaults to `60`.
     */
    readonly backendPoolsSendReceiveTimeoutSeconds?: pulumi.Input<number>;
    /**
     * The host that each frontendEndpoint must CNAME to.
     */
    readonly cname?: pulumi.Input<string>;
    /**
     * Enforce certificate name check on `HTTPS` requests to all backend pools, this setting will have no effect on `HTTP` requests. Permitted values are `true` or `false`.
     */
    readonly enforceBackendPoolsCertificateNameCheck?: pulumi.Input<boolean>;
    /**
     * A friendly name for the Front Door service.
     */
    readonly friendlyName?: pulumi.Input<string>;
    /**
     * A `frontendEndpoint` block as defined below.
     */
    readonly frontendEndpoints?: pulumi.Input<pulumi.Input<inputs.frontdoor.FrontdoorFrontendEndpoint>[]>;
    /**
     * The unique ID of the Front Door which is embedded into the incoming headers `X-Azure-FDID` attribute and maybe used to filter traffic sent by the Front Door to your backend.
     */
    readonly headerFrontdoorId?: pulumi.Input<string>;
    /**
     * Should the Front Door Load Balancer be Enabled? Defaults to `true`.
     */
    readonly loadBalancerEnabled?: pulumi.Input<boolean>;
    /**
     * @deprecated Due to the service's API changing 'location' must now always be set to 'Global' for new resources, however if the Front Door service was created prior 2020/03/10 it may continue to exist in a specific current location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Front Door service. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the name of the Resource Group in which the Front Door service should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A `routingRule` block as defined below.
     */
    readonly routingRules?: pulumi.Input<pulumi.Input<inputs.frontdoor.FrontdoorRoutingRule>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Frontdoor resource.
 */
export interface FrontdoorArgs {
    /**
     * A `backendPoolHealthProbe` block as defined below.
     */
    readonly backendPoolHealthProbes: pulumi.Input<pulumi.Input<inputs.frontdoor.FrontdoorBackendPoolHealthProbe>[]>;
    /**
     * A `backendPoolLoadBalancing` block as defined below.
     */
    readonly backendPoolLoadBalancings: pulumi.Input<pulumi.Input<inputs.frontdoor.FrontdoorBackendPoolLoadBalancing>[]>;
    /**
     * A `backendPool` block as defined below.
     */
    readonly backendPools: pulumi.Input<pulumi.Input<inputs.frontdoor.FrontdoorBackendPool>[]>;
    /**
     * Specifies the send and receive timeout on forwarding request to the backend. When the timeout is reached, the request fails and returns. Possible values are between `0` - `240`. Defaults to `60`.
     */
    readonly backendPoolsSendReceiveTimeoutSeconds?: pulumi.Input<number>;
    /**
     * Enforce certificate name check on `HTTPS` requests to all backend pools, this setting will have no effect on `HTTP` requests. Permitted values are `true` or `false`.
     */
    readonly enforceBackendPoolsCertificateNameCheck: pulumi.Input<boolean>;
    /**
     * A friendly name for the Front Door service.
     */
    readonly friendlyName?: pulumi.Input<string>;
    /**
     * A `frontendEndpoint` block as defined below.
     */
    readonly frontendEndpoints: pulumi.Input<pulumi.Input<inputs.frontdoor.FrontdoorFrontendEndpoint>[]>;
    /**
     * Should the Front Door Load Balancer be Enabled? Defaults to `true`.
     */
    readonly loadBalancerEnabled?: pulumi.Input<boolean>;
    /**
     * @deprecated Due to the service's API changing 'location' must now always be set to 'Global' for new resources, however if the Front Door service was created prior 2020/03/10 it may continue to exist in a specific current location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Front Door service. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the name of the Resource Group in which the Front Door service should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `routingRule` block as defined below.
     */
    readonly routingRules: pulumi.Input<pulumi.Input<inputs.frontdoor.FrontdoorRoutingRule>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

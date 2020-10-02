// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Manages a Load Balancer Resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US"});
 * const examplePublicIp = new azure.network.PublicIp("examplePublicIp", {
 *     location: "West US",
 *     resourceGroupName: exampleResourceGroup.name,
 *     allocationMethod: "Static",
 * });
 * const exampleLoadBalancer = new azure.lb.LoadBalancer("exampleLoadBalancer", {
 *     location: "West US",
 *     resourceGroupName: exampleResourceGroup.name,
 *     frontendIpConfigurations: [{
 *         name: "PublicIPAddress",
 *         publicIpAddressId: examplePublicIp.id,
 *     }],
 * });
 * ```
 */
export class LoadBalancer extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadBalancerState, opts?: pulumi.CustomResourceOptions): LoadBalancer {
        return new LoadBalancer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:lb/loadBalancer:LoadBalancer';

    /**
     * Returns true if the given object is an instance of LoadBalancer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancer.__pulumiType;
    }

    /**
     * One or multiple `frontendIpConfiguration` blocks as documented below.
     */
    public readonly frontendIpConfigurations!: pulumi.Output<outputs.lb.LoadBalancerFrontendIpConfiguration[] | undefined>;
    /**
     * Specifies the supported Azure Region where the Load Balancer should be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Load Balancer.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
     */
    public /*out*/ readonly privateIpAddress!: pulumi.Output<string>;
    /**
     * The list of private IP address assigned to the load balancer in `frontendIpConfiguration` blocks, if any.
     */
    public /*out*/ readonly privateIpAddresses!: pulumi.Output<string[]>;
    /**
     * The name of the Resource Group in which to create the Load Balancer.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The SKU of the Azure Load Balancer. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
     */
    public readonly sku!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a LoadBalancer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadBalancerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoadBalancerArgs | LoadBalancerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LoadBalancerState | undefined;
            inputs["frontendIpConfigurations"] = state ? state.frontendIpConfigurations : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["privateIpAddress"] = state ? state.privateIpAddress : undefined;
            inputs["privateIpAddresses"] = state ? state.privateIpAddresses : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as LoadBalancerArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["frontendIpConfigurations"] = args ? args.frontendIpConfigurations : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["privateIpAddress"] = undefined /*out*/;
            inputs["privateIpAddresses"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LoadBalancer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadBalancer resources.
 */
export interface LoadBalancerState {
    /**
     * One or multiple `frontendIpConfiguration` blocks as documented below.
     */
    readonly frontendIpConfigurations?: pulumi.Input<pulumi.Input<inputs.lb.LoadBalancerFrontendIpConfiguration>[]>;
    /**
     * Specifies the supported Azure Region where the Load Balancer should be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Load Balancer.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
     */
    readonly privateIpAddress?: pulumi.Input<string>;
    /**
     * The list of private IP address assigned to the load balancer in `frontendIpConfiguration` blocks, if any.
     */
    readonly privateIpAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the Resource Group in which to create the Load Balancer.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The SKU of the Azure Load Balancer. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
     */
    readonly sku?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a LoadBalancer resource.
 */
export interface LoadBalancerArgs {
    /**
     * One or multiple `frontendIpConfiguration` blocks as documented below.
     */
    readonly frontendIpConfigurations?: pulumi.Input<pulumi.Input<inputs.lb.LoadBalancerFrontendIpConfiguration>[]>;
    /**
     * Specifies the supported Azure Region where the Load Balancer should be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Load Balancer.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which to create the Load Balancer.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The SKU of the Azure Load Balancer. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
     */
    readonly sku?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

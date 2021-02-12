// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Load Balancer Backend Address Pool.
 *
 * > **NOTE:** When using this resource, the Load Balancer needs to have a FrontEnd IP Configuration Attached
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
 * const exampleBackendAddressPool = new azure.lb.BackendAddressPool("exampleBackendAddressPool", {loadbalancerId: exampleLoadBalancer.id});
 * ```
 *
 * ## Import
 *
 * Load Balancer Backend Address Pools can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:lb/backendAddressPool:BackendAddressPool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/backendAddressPools/pool1
 * ```
 */
export class BackendAddressPool extends pulumi.CustomResource {
    /**
     * Get an existing BackendAddressPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackendAddressPoolState, opts?: pulumi.CustomResourceOptions): BackendAddressPool {
        return new BackendAddressPool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:lb/backendAddressPool:BackendAddressPool';

    /**
     * Returns true if the given object is an instance of BackendAddressPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackendAddressPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackendAddressPool.__pulumiType;
    }

    /**
     * An array of `backendAddress` block as defined below.
     */
    public readonly backendAddresses!: pulumi.Output<outputs.lb.BackendAddressPoolBackendAddress[]>;
    /**
     * The Backend IP Configurations associated with this Backend Address Pool.
     */
    public /*out*/ readonly backendIpConfigurations!: pulumi.Output<string[]>;
    /**
     * The Load Balancing Rules associated with this Backend Address Pool.
     */
    public /*out*/ readonly loadBalancingRules!: pulumi.Output<string[]>;
    /**
     * The ID of the Load Balancer in which to create the Backend Address Pool.
     */
    public readonly loadbalancerId!: pulumi.Output<string>;
    /**
     * Specifies the name of the Backend Address Pool.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * An array of the Load Balancing Outbound Rules associated with this Backend Address Pool.
     */
    public /*out*/ readonly outboundRules!: pulumi.Output<string[]>;
    /**
     * @deprecated This field is no longer used and will be removed in the next major version of the Azure Provider
     */
    public readonly resourceGroupName!: pulumi.Output<string>;

    /**
     * Create a BackendAddressPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackendAddressPoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackendAddressPoolArgs | BackendAddressPoolState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BackendAddressPoolState | undefined;
            inputs["backendAddresses"] = state ? state.backendAddresses : undefined;
            inputs["backendIpConfigurations"] = state ? state.backendIpConfigurations : undefined;
            inputs["loadBalancingRules"] = state ? state.loadBalancingRules : undefined;
            inputs["loadbalancerId"] = state ? state.loadbalancerId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["outboundRules"] = state ? state.outboundRules : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
        } else {
            const args = argsOrState as BackendAddressPoolArgs | undefined;
            if ((!args || args.loadbalancerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loadbalancerId'");
            }
            inputs["backendAddresses"] = args ? args.backendAddresses : undefined;
            inputs["loadbalancerId"] = args ? args.loadbalancerId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["backendIpConfigurations"] = undefined /*out*/;
            inputs["loadBalancingRules"] = undefined /*out*/;
            inputs["outboundRules"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(BackendAddressPool.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackendAddressPool resources.
 */
export interface BackendAddressPoolState {
    /**
     * An array of `backendAddress` block as defined below.
     */
    readonly backendAddresses?: pulumi.Input<pulumi.Input<inputs.lb.BackendAddressPoolBackendAddress>[]>;
    /**
     * The Backend IP Configurations associated with this Backend Address Pool.
     */
    readonly backendIpConfigurations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Load Balancing Rules associated with this Backend Address Pool.
     */
    readonly loadBalancingRules?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the Load Balancer in which to create the Backend Address Pool.
     */
    readonly loadbalancerId?: pulumi.Input<string>;
    /**
     * Specifies the name of the Backend Address Pool.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * An array of the Load Balancing Outbound Rules associated with this Backend Address Pool.
     */
    readonly outboundRules?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * @deprecated This field is no longer used and will be removed in the next major version of the Azure Provider
     */
    readonly resourceGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BackendAddressPool resource.
 */
export interface BackendAddressPoolArgs {
    /**
     * An array of `backendAddress` block as defined below.
     */
    readonly backendAddresses?: pulumi.Input<pulumi.Input<inputs.lb.BackendAddressPoolBackendAddress>[]>;
    /**
     * The ID of the Load Balancer in which to create the Backend Address Pool.
     */
    readonly loadbalancerId: pulumi.Input<string>;
    /**
     * Specifies the name of the Backend Address Pool.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * @deprecated This field is no longer used and will be removed in the next major version of the Azure Provider
     */
    readonly resourceGroupName?: pulumi.Input<string>;
}

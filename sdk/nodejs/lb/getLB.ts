// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Load Balancer
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.lb.getLB({
 *     name: "example-lb",
 *     resourceGroupName: "example-resources",
 * });
 * export const loadbalancerId = example.then(example => example.id);
 * ```
 */
export function getLB(args: GetLBArgs, opts?: pulumi.InvokeOptions): Promise<GetLBResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:lb/getLB:getLB", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getLB.
 */
export interface GetLBArgs {
    /**
     * Specifies the name of the Load Balancer.
     */
    name: string;
    /**
     * The name of the Resource Group in which the Load Balancer exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getLB.
 */
export interface GetLBResult {
    /**
     * A `frontendIpConfiguration` block as documented below.
     */
    readonly frontendIpConfigurations: outputs.lb.GetLBFrontendIpConfiguration[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Azure location where the Load Balancer exists.
     */
    readonly location: string;
    /**
     * The name of the Frontend IP Configuration.
     */
    readonly name: string;
    /**
     * Private IP Address to assign to the Load Balancer.
     */
    readonly privateIpAddress: string;
    /**
     * The list of private IP address assigned to the load balancer in `frontendIpConfiguration` blocks, if any.
     */
    readonly privateIpAddresses: string[];
    readonly resourceGroupName: string;
    /**
     * The SKU of the Load Balancer.
     */
    readonly sku: string;
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags: {[key: string]: string};
}

export function getLBApply(args: GetLBApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLBResult> {
    return pulumi.output(args).apply(a => getLB(a, opts))
}

/**
 * A collection of arguments for invoking getLB.
 */
export interface GetLBApplyArgs {
    /**
     * Specifies the name of the Load Balancer.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the Load Balancer exists.
     */
    resourceGroupName: pulumi.Input<string>;
}

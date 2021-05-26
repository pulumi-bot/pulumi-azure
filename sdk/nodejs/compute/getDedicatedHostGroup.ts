// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Dedicated Host Group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.compute.getDedicatedHostGroup({
 *     name: "example-dedicated-host-group",
 *     resourceGroupName: "example-rg",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getDedicatedHostGroup(args: GetDedicatedHostGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetDedicatedHostGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:compute/getDedicatedHostGroup:getDedicatedHostGroup", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getDedicatedHostGroup.
 */
export interface GetDedicatedHostGroupArgs {
    /**
     * Specifies the name of the Dedicated Host Group.
     */
    name: string;
    /**
     * Specifies the name of the resource group the Dedicated Host Group is located in.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getDedicatedHostGroup.
 */
export interface GetDedicatedHostGroupResult {
    /**
     * Whether virtual machines or virtual machine scale sets be placed automatically on this Dedicated Host Group.
     */
    readonly automaticPlacementEnabled: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Azure location where the Dedicated Host Group exists.
     */
    readonly location: string;
    readonly name: string;
    /**
     * The number of fault domains that the Dedicated Host Group spans.
     */
    readonly platformFaultDomainCount: number;
    readonly resourceGroupName: string;
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags: {[key: string]: string};
    /**
     * The Availability Zones in which this Dedicated Host Group is located.
     */
    readonly zones: string[];
}

export function getDedicatedHostGroupApply(args: GetDedicatedHostGroupApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDedicatedHostGroupResult> {
    return pulumi.output(args).apply(a => getDedicatedHostGroup(a, opts))
}

/**
 * A collection of arguments for invoking getDedicatedHostGroup.
 */
export interface GetDedicatedHostGroupApplyArgs {
    /**
     * Specifies the name of the Dedicated Host Group.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the name of the resource group the Dedicated Host Group is located in.
     */
    resourceGroupName: pulumi.Input<string>;
}

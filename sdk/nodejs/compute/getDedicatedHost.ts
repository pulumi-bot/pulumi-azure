// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Dedicated Host.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.compute.getDedicatedHost({
 *     name: "example-host",
 *     dedicatedHostGroupName: "example-host-group",
 *     resourceGroupName: "example-resources",
 * });
 * export const dedicatedHostId = example.then(example => example.id);
 * ```
 */
export function getDedicatedHost(args: GetDedicatedHostArgs, opts?: pulumi.InvokeOptions): Promise<GetDedicatedHostResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:compute/getDedicatedHost:getDedicatedHost", {
        "dedicatedHostGroupName": args.dedicatedHostGroupName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getDedicatedHost.
 */
export interface GetDedicatedHostArgs {
    /**
     * Specifies the name of the Dedicated Host Group the Dedicated Host is located in.
     */
    readonly dedicatedHostGroupName: string;
    /**
     * Specifies the name of the Dedicated Host.
     */
    readonly name: string;
    /**
     * Specifies the name of the resource group the Dedicated Host is located in.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getDedicatedHost.
 */
export interface GetDedicatedHostResult {
    readonly dedicatedHostGroupName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The location where the Dedicated Host exists.
     */
    readonly location: string;
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * A mapping of tags assigned to the Dedicated Host.
     */
    readonly tags: {[key: string]: string};
}

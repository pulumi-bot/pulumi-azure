// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Route Filter.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.network.getRouteFilter({
 *     name: "existing",
 *     resourceGroupName: "existing",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getRouteFilter(args: GetRouteFilterArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteFilterResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:network/getRouteFilter:getRouteFilter", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouteFilter.
 */
export interface GetRouteFilterArgs {
    /**
     * The Name of this Route Filter.
     */
    readonly name: string;
    /**
     * The name of the Resource Group where the Route Filter exists.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getRouteFilter.
 */
export interface GetRouteFilterResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Azure Region where the Route Filter exists.
     */
    readonly location: string;
    /**
     * The Name of Route Filter Rule
     */
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * A `rule` block as defined below.
     */
    readonly rules: outputs.network.GetRouteFilterRule[];
    /**
     * A mapping of tags assigned to the Route Filter.
     */
    readonly tags: {[key: string]: string};
}

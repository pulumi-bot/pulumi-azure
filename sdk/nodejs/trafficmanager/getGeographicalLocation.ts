// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access the ID of a specified Traffic Manager Geographical Location within the Geographical Hierarchy.
 *
 * ## Example Usage
 * ### World)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.network.getTrafficManager({
 *     name: "World",
 * });
 * export const locationCode = example.then(example => example.id);
 * ```
 */
/** @deprecated azure.trafficmanager.getGeographicalLocation has been deprecated in favor of azure.network.getTrafficManager */
export function getGeographicalLocation(args: GetGeographicalLocationArgs, opts?: pulumi.InvokeOptions): Promise<GetGeographicalLocationResult> {
    pulumi.log.warn("getGeographicalLocation is deprecated: azure.trafficmanager.getGeographicalLocation has been deprecated in favor of azure.network.getTrafficManager")
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:trafficmanager/getGeographicalLocation:getGeographicalLocation", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getGeographicalLocation.
 */
export interface GetGeographicalLocationArgs {
    /**
     * Specifies the name of the Location, for example `World`, `Europe` or `Germany`.
     */
    name: string;
}

/**
 * A collection of values returned by getGeographicalLocation.
 */
export interface GetGeographicalLocationResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
}

export function getGeographicalLocationOutput(args: GetGeographicalLocationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGeographicalLocationResult> {
    return pulumi.output(args).apply(a => getGeographicalLocation(a, opts))
}

/**
 * A collection of arguments for invoking getGeographicalLocation.
 */
export interface GetGeographicalLocationOutputArgs {
    /**
     * Specifies the name of the Location, for example `World`, `Europe` or `Germany`.
     */
    name: pulumi.Input<string>;
}

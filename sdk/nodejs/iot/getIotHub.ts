// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing IoTHub.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.iot.getIotHub({
 *     name: "existing",
 *     resourceGroupName: "existing",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getIotHub(args: GetIotHubArgs, opts?: pulumi.InvokeOptions): Promise<GetIotHubResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:iot/getIotHub:getIotHub", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getIotHub.
 */
export interface GetIotHubArgs {
    /**
     * The name of this IoTHub.
     */
    name: string;
    /**
     * The name of the Resource Group where the IoTHub exists.
     */
    resourceGroupName: string;
    /**
     * A mapping of tags which should be assigned to the IoTHub.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getIotHub.
 */
export interface GetIotHubResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly resourceGroupName: string;
    readonly tags?: {[key: string]: string};
}

export function getIotHubApply(args: GetIotHubApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIotHubResult> {
    return pulumi.output(args).apply(a => getIotHub(a, opts))
}

/**
 * A collection of arguments for invoking getIotHub.
 */
export interface GetIotHubApplyArgs {
    /**
     * The name of this IoTHub.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the IoTHub exists.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the IoTHub.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

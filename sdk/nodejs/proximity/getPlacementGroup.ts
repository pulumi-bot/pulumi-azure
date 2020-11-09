// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Proximity Placement Group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.proximity.getPlacementGroup({
 *     name: "tf-appsecuritygroup",
 *     resourceGroupName: "my-resource-group",
 * });
 * export const proximityPlacementGroupId = example.then(example => example.id);
 * ```
 */
export function getPlacementGroup(args: GetPlacementGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetPlacementGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:proximity/getPlacementGroup:getPlacementGroup", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getPlacementGroup.
 */
export interface GetPlacementGroupArgs {
    /**
     * The name of the Proximity Placement Group.
     */
    readonly name: string;
    /**
     * The name of the resource group in which the Proximity Placement Group exists.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getPlacementGroup.
 */
export interface GetPlacementGroupResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly location: string;
    readonly name: string;
    readonly resourceGroupName: string;
    readonly tags: {[key: string]: string};
}

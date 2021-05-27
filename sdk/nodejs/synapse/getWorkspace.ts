// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Synapse Workspace.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.synapse.getWorkspace({
 *     name: "existing",
 *     resourceGroupName: "example-resource-group",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getWorkspace(args: GetWorkspaceArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkspaceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:synapse/getWorkspace:getWorkspace", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getWorkspace.
 */
export interface GetWorkspaceArgs {
    /**
     * The name of this Synapse Workspace.
     */
    name: string;
    /**
     * The name of the Resource Group where the Synapse Workspace exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getWorkspace.
 */
export interface GetWorkspaceResult {
    /**
     * A list of Connectivity endpoints for this Synapse Workspace.
     */
    readonly connectivityEndpoints: {[key: string]: string};
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Azure location where the Synapse Workspace exists.
     */
    readonly location: string;
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags: {[key: string]: string};
}

export function getWorkspaceApply(args: GetWorkspaceApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWorkspaceResult> {
    return pulumi.output(args).apply(a => getWorkspace(a, opts))
}

/**
 * A collection of arguments for invoking getWorkspace.
 */
export interface GetWorkspaceApplyArgs {
    /**
     * The name of this Synapse Workspace.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Synapse Workspace exists.
     */
    resourceGroupName: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Databricks workspace.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.databricks.getWorkspace({
 *     name: "example-workspace",
 *     resourceGroupName: "example-rg",
 * });
 * export const databricksWorkspaceId = example.then(example => example.workspaceId);
 * ```
 */
export function getWorkspace(args: GetWorkspaceArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkspaceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:databricks/getWorkspace:getWorkspace", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getWorkspace.
 */
export interface GetWorkspaceArgs {
    /**
     * The name of the Databricks Workspace.
     */
    name: string;
    /**
     * The Name of the Resource Group where the Databricks Workspace exists.
     */
    resourceGroupName: string;
    /**
     * A mapping of tags to assign to the Databricks Workspace.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getWorkspace.
 */
export interface GetWorkspaceResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * SKU of this Databricks Workspace.
     */
    readonly sku: string;
    /**
     * A mapping of tags to assign to the Databricks Workspace.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Unique ID of this Databricks Workspace in Databricks management plane.
     */
    readonly workspaceId: string;
    /**
     * URL this Databricks Workspace is accessible on.
     */
    readonly workspaceUrl: string;
}

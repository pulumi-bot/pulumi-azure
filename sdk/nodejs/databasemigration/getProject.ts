// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Database Migration Project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = pulumi.output(azure.databasemigration.getProject({
 *     name: "example-dbms-project",
 *     resourceGroupName: "example-rg",
 *     serviceName: "example-dbms",
 * }, { async: true }));
 *
 * export const name = example.name;
 * ```
 */
export function getProject(args: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:databasemigration/getProject:getProject", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectArgs {
    /**
     * Name of the database migration project.
     */
    name: string;
    /**
     * Name of the resource group where resource belongs to.
     */
    resourceGroupName: string;
    /**
     * Name of the database migration service where resource belongs to.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getProject.
 */
export interface GetProjectResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Azure location where the resource exists.
     */
    readonly location: string;
    readonly name: string;
    readonly resourceGroupName: string;
    readonly serviceName: string;
    /**
     * The platform type of the migration source.
     */
    readonly sourcePlatform: string;
    /**
     * A mapping of tags to assigned to the resource.
     */
    readonly tags: {[key: string]: string};
    /**
     * The platform type of the migration target.
     */
    readonly targetPlatform: string;
}

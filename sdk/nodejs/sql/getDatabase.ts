// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing SQL Azure Database.
 */
export function getDatabase(args: GetDatabaseArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:sql/getDatabase:getDatabase", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabase.
 */
export interface GetDatabaseArgs {
    /**
     * The name of the SQL Database.
     */
    readonly name: string;
    /**
     * Specifies the name of the Resource Group where the Azure SQL Database exists.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the SQL Server.
     */
    readonly serverName: string;
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getDatabase.
 */
export interface GetDatabaseResult {
    /**
     * The name of the collation.
     */
    readonly collation: string;
    /**
     * The default secondary location of the SQL Database.
     */
    readonly defaultSecondaryLocation: string;
    /**
     * The edition of the database.
     */
    readonly edition: string;
    /**
     * The name of the elastic database pool the database belongs to.
     */
    readonly elasticPoolName: string;
    /**
     * The ID of the failover group the database belongs to.
     */
    readonly failoverGroupId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The location of the Resource Group in which the SQL Server exists.
     */
    readonly location: string;
    /**
     * The name of the database.
     */
    readonly name: string;
    /**
     * Indicate if read-only connections will be redirected to a high-available replica.
     */
    readonly readScale: boolean;
    /**
     * The name of the resource group in which the database resides. This will always be the same resource group as the Database Server.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the SQL Server on which to create the database.
     */
    readonly serverName: string;
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags?: {[key: string]: string};
}

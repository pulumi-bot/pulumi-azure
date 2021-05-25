// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to obtain a Shared Access Signature (SAS Token) for an existing Storage Account Blob Container.
 *
 * Shared access signatures allow fine-grained, ephemeral access control to various aspects of an Azure Storage Account Blob Container.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const rg = new azure.core.ResourceGroup("rg", {location: "West Europe"});
 * const storage = new azure.storage.Account("storage", {
 *     resourceGroupName: rg.name,
 *     location: rg.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const container = new azure.storage.Container("container", {
 *     storageAccountName: storage.name,
 *     containerAccessType: "private",
 * });
 * const example = pulumi.all([storage.primaryConnectionString, container.name]).apply(([primaryConnectionString, name]) => azure.storage.getAccountBlobContainerSAS({
 *     connectionString: primaryConnectionString,
 *     containerName: name,
 *     httpsOnly: true,
 *     ipAddress: "168.1.5.65",
 *     start: "2018-03-21",
 *     expiry: "2018-03-21",
 *     permissions: {
 *         read: true,
 *         add: true,
 *         create: false,
 *         write: false,
 *         "delete": true,
 *         list: true,
 *     },
 *     cacheControl: "max-age=5",
 *     contentDisposition: "inline",
 *     contentEncoding: "deflate",
 *     contentLanguage: "en-US",
 *     contentType: "application/json",
 * }));
 * export const sasUrlQueryString = example.sas;
 * ```
 */
export function getAccountBlobContainerSAS(args: GetAccountBlobContainerSASArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountBlobContainerSASResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:storage/getAccountBlobContainerSAS:getAccountBlobContainerSAS", {
        "cacheControl": args.cacheControl,
        "connectionString": args.connectionString,
        "containerName": args.containerName,
        "contentDisposition": args.contentDisposition,
        "contentEncoding": args.contentEncoding,
        "contentLanguage": args.contentLanguage,
        "contentType": args.contentType,
        "expiry": args.expiry,
        "httpsOnly": args.httpsOnly,
        "ipAddress": args.ipAddress,
        "permissions": args.permissions,
        "start": args.start,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccountBlobContainerSAS.
 */
export interface GetAccountBlobContainerSASArgs {
    /**
     * The `Cache-Control` response header that is sent when this SAS token is used.
     */
    cacheControl?: string;
    /**
     * The connection string for the storage account to which this SAS applies. Typically directly from the `primaryConnectionString` attribute of an `azure.storage.Account` resource.
     */
    connectionString: string;
    /**
     * Name of the container.
     */
    containerName: string;
    /**
     * The `Content-Disposition` response header that is sent when this SAS token is used.
     */
    contentDisposition?: string;
    /**
     * The `Content-Encoding` response header that is sent when this SAS token is used.
     */
    contentEncoding?: string;
    /**
     * The `Content-Language` response header that is sent when this SAS token is used.
     */
    contentLanguage?: string;
    /**
     * The `Content-Type` response header that is sent when this SAS token is used.
     */
    contentType?: string;
    /**
     * The expiration time and date of this SAS. Must be a valid ISO-8601 format time/date string.
     */
    expiry: string;
    /**
     * Only permit `https` access. If `false`, both `http` and `https` are permitted. Defaults to `true`.
     */
    httpsOnly?: boolean;
    /**
     * Single ipv4 address or range (connected with a dash) of ipv4 addresses.
     */
    ipAddress?: string;
    /**
     * A `permissions` block as defined below.
     */
    permissions: inputs.storage.GetAccountBlobContainerSASPermissions;
    /**
     * The starting time and date of validity of this SAS. Must be a valid ISO-8601 format time/date string.
     */
    start: string;
}

/**
 * A collection of values returned by getAccountBlobContainerSAS.
 */
export interface GetAccountBlobContainerSASResult {
    readonly cacheControl?: string;
    readonly connectionString: string;
    readonly containerName: string;
    readonly contentDisposition?: string;
    readonly contentEncoding?: string;
    readonly contentLanguage?: string;
    readonly contentType?: string;
    readonly expiry: string;
    readonly httpsOnly?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipAddress?: string;
    readonly permissions: outputs.storage.GetAccountBlobContainerSASPermissions;
    /**
     * The computed Blob Container Shared Access Signature (SAS).
     */
    readonly sas: string;
    readonly start: string;
}

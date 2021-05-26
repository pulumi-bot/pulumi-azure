// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Storage Blob.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = pulumi.output(azure.storage.getBlob({
 *     name: "example-blob-name",
 *     storageAccountName: "example-storage-account-name",
 *     storageContainerName: "example-storage-container-name",
 * }, { async: true }));
 * ```
 */
export function getBlob(args: GetBlobArgs, opts?: pulumi.InvokeOptions): Promise<GetBlobResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:storage/getBlob:getBlob", {
        "metadata": args.metadata,
        "name": args.name,
        "storageAccountName": args.storageAccountName,
        "storageContainerName": args.storageContainerName,
    }, opts);
}

/**
 * A collection of arguments for invoking getBlob.
 */
export interface GetBlobArgs {
    /**
     * A map of custom blob metadata.
     */
    metadata?: {[key: string]: string};
    /**
     * The name of the Blob.
     */
    name: string;
    /**
     * The name of the Storage Account where the Container exists.
     */
    storageAccountName: string;
    /**
     * The name of the Storage Container where the Blob exists.
     */
    storageContainerName: string;
}

/**
 * A collection of values returned by getBlob.
 */
export interface GetBlobResult {
    /**
     * The access tier of the storage blob.
     */
    readonly accessTier: string;
    /**
     * The MD5 sum of the blob contents.
     */
    readonly contentMd5: string;
    /**
     * The content type of the storage blob.
     */
    readonly contentType: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A map of custom blob metadata.
     */
    readonly metadata: {[key: string]: string};
    readonly name: string;
    readonly storageAccountName: string;
    readonly storageContainerName: string;
    /**
     * The type of the storage blob
     */
    readonly type: string;
    /**
     * The URL of the storage blob.
     */
    readonly url: string;
}

export function getBlobApply(args: GetBlobApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBlobResult> {
    return pulumi.output(args).apply(a => getBlob(a, opts))
}

/**
 * A collection of arguments for invoking getBlob.
 */
export interface GetBlobApplyArgs {
    /**
     * A map of custom blob metadata.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the Blob.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the Storage Account where the Container exists.
     */
    storageAccountName: pulumi.Input<string>;
    /**
     * The name of the Storage Container where the Blob exists.
     */
    storageContainerName: pulumi.Input<string>;
}

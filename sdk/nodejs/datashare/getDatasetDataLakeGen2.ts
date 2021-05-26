// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Data Share Data Lake Gen2 Dataset.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.datashare.getDatasetDataLakeGen2({
 *     name: "example-dsdlg2ds",
 *     shareId: "example-share-id",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getDatasetDataLakeGen2(args: GetDatasetDataLakeGen2Args, opts?: pulumi.InvokeOptions): Promise<GetDatasetDataLakeGen2Result> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:datashare/getDatasetDataLakeGen2:getDatasetDataLakeGen2", {
        "name": args.name,
        "shareId": args.shareId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatasetDataLakeGen2.
 */
export interface GetDatasetDataLakeGen2Args {
    /**
     * The name of this Data Share Data Lake Gen2 Dataset.
     */
    name: string;
    /**
     * The resource ID of the Data Share where this Data Share Data Lake Gen2 Dataset should be created.
     */
    shareId: string;
}

/**
 * A collection of values returned by getDatasetDataLakeGen2.
 */
export interface GetDatasetDataLakeGen2Result {
    /**
     * The name of the Data Share Dataset.
     */
    readonly displayName: string;
    /**
     * The path of the file in the data lake file system to be shared with the receiver.
     */
    readonly filePath: string;
    /**
     * The name of the data lake file system to be shared with the receiver.
     */
    readonly fileSystemName: string;
    /**
     * The folder path in the data lake file system to be shared with the receiver.
     */
    readonly folderPath: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly shareId: string;
    /**
     * The resource ID of the storage account of the data lake file system to be shared with the receiver.
     */
    readonly storageAccountId: string;
}

export function getDatasetDataLakeGen2Apply(args: GetDatasetDataLakeGen2ApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatasetDataLakeGen2Result> {
    return pulumi.output(args).apply(a => getDatasetDataLakeGen2(a, opts))
}

/**
 * A collection of arguments for invoking getDatasetDataLakeGen2.
 */
export interface GetDatasetDataLakeGen2ApplyArgs {
    /**
     * The name of this Data Share Data Lake Gen2 Dataset.
     */
    name: pulumi.Input<string>;
    /**
     * The resource ID of the Data Share where this Data Share Data Lake Gen2 Dataset should be created.
     */
    shareId: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Storage Container.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = pulumi.output(azure.storage.getStorageContainer({
 *     name: "example-container-name",
 *     storageAccountName: "example-storage-account-name",
 * }, { async: true }));
 * ```
 */
export function getStorageContainer(args: GetStorageContainerArgs, opts?: pulumi.InvokeOptions): Promise<GetStorageContainerResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:storage/getStorageContainer:getStorageContainer", {
        "metadata": args.metadata,
        "name": args.name,
        "storageAccountName": args.storageAccountName,
    }, opts);
}

/**
 * A collection of arguments for invoking getStorageContainer.
 */
export interface GetStorageContainerArgs {
    /**
     * A mapping of MetaData for this Container.
     */
    readonly metadata?: {[key: string]: string};
    /**
     * The name of the Container.
     */
    readonly name: string;
    /**
     * The name of the Storage Account where the Container exists.
     */
    readonly storageAccountName: string;
}

/**
 * A collection of values returned by getStorageContainer.
 */
export interface GetStorageContainerResult {
    /**
     * The Access Level configured for this Container.
     */
    readonly containerAccessType: string;
    /**
     * Is there an Immutability Policy configured on this Storage Container?
     */
    readonly hasImmutabilityPolicy: boolean;
    /**
     * Is there a Legal Hold configured on this Storage Container?
     */
    readonly hasLegalHold: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A mapping of MetaData for this Container.
     */
    readonly metadata: {[key: string]: string};
    readonly name: string;
    /**
     * The Resource Manager ID of this Storage Container.
     */
    readonly resourceManagerId: string;
    readonly storageAccountName: string;
}

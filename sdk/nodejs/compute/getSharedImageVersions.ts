// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about existing Versions of a Shared Image within a Shared Image Gallery.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = pulumi.output(azure.compute.getSharedImageVersions({
 *     galleryName: "my-image-gallery",
 *     imageName: "my-image",
 *     resourceGroupName: "example-resources",
 * }, { async: true }));
 * ```
 */
export function getSharedImageVersions(args: GetSharedImageVersionsArgs, opts?: pulumi.InvokeOptions): Promise<GetSharedImageVersionsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:compute/getSharedImageVersions:getSharedImageVersions", {
        "galleryName": args.galleryName,
        "imageName": args.imageName,
        "resourceGroupName": args.resourceGroupName,
        "tagsFilter": args.tagsFilter,
    }, opts);
}

/**
 * A collection of arguments for invoking getSharedImageVersions.
 */
export interface GetSharedImageVersionsArgs {
    /**
     * The name of the Shared Image in which the Shared Image exists.
     */
    readonly galleryName: string;
    /**
     * The name of the Shared Image in which this Version exists.
     */
    readonly imageName: string;
    /**
     * The name of the Resource Group in which the Shared Image Gallery exists.
     */
    readonly resourceGroupName: string;
    /**
     * A mapping of tags to filter the list of images against.
     */
    readonly tagsFilter?: {[key: string]: string};
}

/**
 * A collection of values returned by getSharedImageVersions.
 */
export interface GetSharedImageVersionsResult {
    readonly galleryName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly imageName: string;
    /**
     * An `images` block as defined below:
     */
    readonly images: outputs.compute.GetSharedImageVersionsImage[];
    readonly resourceGroupName: string;
    readonly tagsFilter?: {[key: string]: string};
}

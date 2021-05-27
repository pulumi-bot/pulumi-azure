// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Uses this data source to access information about an existing NetApp Volume.
 *
 * ## NetApp Volume Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.netapp.getVolume({
 *     resourceGroupName: "acctestRG",
 *     accountName: "acctestnetappaccount",
 *     poolName: "acctestnetapppool",
 *     name: "example-volume",
 * });
 * export const netappVolumeId = example.then(example => example.id);
 * ```
 */
export function getVolume(args: GetVolumeArgs, opts?: pulumi.InvokeOptions): Promise<GetVolumeResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:netapp/getVolume:getVolume", {
        "accountName": args.accountName,
        "name": args.name,
        "poolName": args.poolName,
        "resourceGroupName": args.resourceGroupName,
        "securityStyle": args.securityStyle,
    }, opts);
}

/**
 * A collection of arguments for invoking getVolume.
 */
export interface GetVolumeArgs {
    /**
     * The name of the NetApp account where the NetApp pool exists.
     */
    accountName: string;
    /**
     * The name of the NetApp Volume.
     */
    name: string;
    /**
     * The name of the NetApp pool where the NetApp volume exists.
     */
    poolName: string;
    /**
     * The Name of the Resource Group where the NetApp Volume exists.
     */
    resourceGroupName: string;
    /**
     * Volume security style
     */
    securityStyle?: string;
}

/**
 * A collection of values returned by getVolume.
 */
export interface GetVolumeResult {
    readonly accountName: string;
    /**
     * Volume data protection block
     * *
     */
    readonly dataProtectionReplications: outputs.netapp.GetVolumeDataProtectionReplication[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Azure Region where the NetApp Volume exists.
     */
    readonly location: string;
    /**
     * A list of IPv4 Addresses which should be used to mount the volume.
     */
    readonly mountIpAddresses: string[];
    readonly name: string;
    readonly poolName: string;
    /**
     * A list of protocol types enabled on volume.
     */
    readonly protocols: string[];
    readonly resourceGroupName: string;
    /**
     * Volume security style
     */
    readonly securityStyle?: string;
    /**
     * The service level of the file system.
     */
    readonly serviceLevel: string;
    /**
     * The maximum Storage Quota in Gigabytes allowed for a file system.
     */
    readonly storageQuotaInGb: number;
    /**
     * The ID of a Subnet in which the NetApp Volume resides.
     */
    readonly subnetId: string;
    /**
     * The unique file path of the volume.
     */
    readonly volumePath: string;
}

export function getVolumeApply(args: GetVolumeApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVolumeResult> {
    return pulumi.output(args).apply(a => getVolume(a, opts))
}

/**
 * A collection of arguments for invoking getVolume.
 */
export interface GetVolumeApplyArgs {
    /**
     * The name of the NetApp account where the NetApp pool exists.
     */
    accountName: pulumi.Input<string>;
    /**
     * The name of the NetApp Volume.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the NetApp pool where the NetApp volume exists.
     */
    poolName: pulumi.Input<string>;
    /**
     * The Name of the Resource Group where the NetApp Volume exists.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Volume security style
     */
    securityStyle?: pulumi.Input<string>;
}

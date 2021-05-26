// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing DNS Zone.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.dns.getZone({
 *     name: "search-eventhubns",
 *     resourceGroupName: "search-service",
 * });
 * export const dnsZoneId = example.then(example => example.id);
 * ```
 */
export function getZone(args: GetZoneArgs, opts?: pulumi.InvokeOptions): Promise<GetZoneResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:dns/getZone:getZone", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getZone.
 */
export interface GetZoneArgs {
    /**
     * The name of the DNS Zone.
     */
    name: string;
    /**
     * The Name of the Resource Group where the DNS Zone exists.
     * If the Name of the Resource Group is not provided, the first DNS Zone from the list of DNS Zones
     * in your subscription that matches `name` will be returned.
     */
    resourceGroupName?: string;
}

/**
 * A collection of values returned by getZone.
 */
export interface GetZoneResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Maximum number of Records in the zone.
     */
    readonly maxNumberOfRecordSets: number;
    readonly name: string;
    /**
     * A list of values that make up the NS record for the zone.
     */
    readonly nameServers: string[];
    /**
     * The number of records already in the zone.
     */
    readonly numberOfRecordSets: number;
    readonly resourceGroupName: string;
    /**
     * A mapping of tags to assign to the EventHub Namespace.
     */
    readonly tags: {[key: string]: string};
}

export function getZoneApply(args: GetZoneApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetZoneResult> {
    return pulumi.output(args).apply(a => getZone(a, opts))
}

/**
 * A collection of arguments for invoking getZone.
 */
export interface GetZoneApplyArgs {
    /**
     * The name of the DNS Zone.
     */
    name: pulumi.Input<string>;
    /**
     * The Name of the Resource Group where the DNS Zone exists.
     * If the Name of the Resource Group is not provided, the first DNS Zone from the list of DNS Zones
     * in your subscription that matches `name` will be returned.
     */
    resourceGroupName?: pulumi.Input<string>;
}

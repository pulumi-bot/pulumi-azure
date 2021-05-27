// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Virtual Wan.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.network.getVirtualWan({
 *     name: "existing",
 *     resourceGroupName: "existing",
 * });
 * export const id = example.then(example => example.id);
 * export const allowBranchToBranchTraffic = example.then(example => example.allowBranchToBranchTraffic);
 * export const disableVpnEncryption = data.azurerm_virtual_wan.exemple.disable_vpn_encryption;
 * export const location = data.azurerm_virtual_wan.exemple.location;
 * export const office365LocalBreakoutCategory = data.azurerm_virtual_wan.exemple.office365_local_breakout_category;
 * export const sku = data.azurerm_virtual_wan.exemple.sku;
 * export const tags = data.azurerm_virtual_wan.exemple.tags;
 * export const virtualHubs = data.azurerm_virtual_wan.exemple.virtual_hubs;
 * export const vpnSites = data.azurerm_virtual_wan.exemple.vpn_sites;
 * ```
 */
export function getVirtualWan(args: GetVirtualWanArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualWanResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:network/getVirtualWan:getVirtualWan", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getVirtualWan.
 */
export interface GetVirtualWanArgs {
    /**
     * The name of this Virtual Wan.
     */
    name: string;
    /**
     * The name of the Resource Group where the Virtual Wan exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getVirtualWan.
 */
export interface GetVirtualWanResult {
    /**
     * Is branch to branch traffic is allowed?
     */
    readonly allowBranchToBranchTraffic: boolean;
    /**
     * Is VPN Encryption disabled?
     */
    readonly disableVpnEncryption: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Azure Region where the Virtual Wan exists.
     */
    readonly location: string;
    readonly name: string;
    /**
     * The Office365 Local Breakout Category.
     */
    readonly office365LocalBreakoutCategory: string;
    readonly resourceGroupName: string;
    /**
     * Type of Virtual Wan (Basic or Standard).
     */
    readonly sku: string;
    /**
     * A mapping of tags assigned to the Virtual Wan.
     */
    readonly tags: {[key: string]: string};
    /**
     * A list of Virtual Hubs ID's attached to this Virtual WAN.
     */
    readonly virtualHubIds: string[];
    /**
     * A list of VPN Site ID's attached to this Virtual WAN.
     */
    readonly vpnSiteIds: string[];
}

export function getVirtualWanApply(args: GetVirtualWanApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVirtualWanResult> {
    return pulumi.output(args).apply(a => getVirtualWan(a, opts))
}

/**
 * A collection of arguments for invoking getVirtualWan.
 */
export interface GetVirtualWanApplyArgs {
    /**
     * The name of this Virtual Wan.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Virtual Wan exists.
     */
    resourceGroupName: pulumi.Input<string>;
}

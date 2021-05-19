// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Network Watcher.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.network.getNetworkWatcher({
 *     name: azurerm_network_watcher.example.name,
 *     resourceGroupName: azurerm_resource_group.example.name,
 * });
 * export const networkWatcherId = example.then(example => example.id);
 * ```
 */
export function getNetworkWatcher(args: GetNetworkWatcherArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkWatcherResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:network/getNetworkWatcher:getNetworkWatcher", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetworkWatcher.
 */
export interface GetNetworkWatcherArgs {
    /**
     * Specifies the Name of the Network Watcher.
     */
    name: string;
    /**
     * Specifies the Name of the Resource Group within which the Network Watcher exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getNetworkWatcher.
 */
export interface GetNetworkWatcherResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The supported Azure location where the resource exists.
     */
    readonly location: string;
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags: {[key: string]: string};
}

export function getNetworkWatcherOutput(args: GetNetworkWatcherOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetworkWatcherResult> {
    return pulumi.output(args).apply(a => getNetworkWatcher(a, opts))
}

/**
 * A collection of arguments for invoking getNetworkWatcher.
 */
export interface GetNetworkWatcherOutputArgs {
    /**
     * Specifies the Name of the Network Watcher.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the Name of the Resource Group within which the Network Watcher exists.
     */
    resourceGroupName: pulumi.Input<string>;
}

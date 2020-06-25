// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Dev Test Lab Virtual Network.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.devtest.getVirtualNetwork({
 *     name: "example-network",
 *     labName: "examplelab",
 *     resourceGroupName: "example-resource",
 * });
 * export const labSubnetName = example.then(example => example.allowedSubnets[0].labSubnetName);
 * ```
 */
export function getVirtualNetwork(args: GetVirtualNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualNetworkResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:devtest/getVirtualNetwork:getVirtualNetwork", {
        "labName": args.labName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getVirtualNetwork.
 */
export interface GetVirtualNetworkArgs {
    /**
     * Specifies the name of the Dev Test Lab.
     */
    readonly labName: string;
    /**
     * Specifies the name of the Virtual Network.
     */
    readonly name: string;
    /**
     * Specifies the name of the resource group that contains the Virtual Network.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getVirtualNetwork.
 */
export interface GetVirtualNetworkResult {
    /**
     * The list of subnets enabled for the virtual network as defined below.
     */
    readonly allowedSubnets: outputs.devtest.GetVirtualNetworkAllowedSubnet[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly labName: string;
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * The list of permission overrides for the subnets as defined below.
     */
    readonly subnetOverrides: outputs.devtest.GetVirtualNetworkSubnetOverride[];
    /**
     * The unique immutable identifier of the virtual network.
     */
    readonly uniqueIdentifier: string;
}

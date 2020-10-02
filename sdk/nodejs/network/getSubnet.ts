// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Subnet within a Virtual Network.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.network.getSubnet({
 *     name: "backend",
 *     virtualNetworkName: "production",
 *     resourceGroupName: "networking",
 * });
 * export const subnetId = example.then(example => example.id);
 * ```
 */
export function getSubnet(args: GetSubnetArgs, opts?: pulumi.InvokeOptions): Promise<GetSubnetResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:network/getSubnet:getSubnet", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "virtualNetworkName": args.virtualNetworkName,
    }, opts);
}

/**
 * A collection of arguments for invoking getSubnet.
 */
export interface GetSubnetArgs {
    /**
     * Specifies the name of the Subnet.
     */
    readonly name: string;
    /**
     * Specifies the name of the resource group the Virtual Network is located in.
     */
    readonly resourceGroupName: string;
    /**
     * Specifies the name of the Virtual Network this Subnet is located within.
     */
    readonly virtualNetworkName: string;
}

/**
 * A collection of values returned by getSubnet.
 */
export interface GetSubnetResult {
    /**
     * (Deprecated) The address prefix used for the subnet.
     */
    readonly addressPrefix: string;
    /**
     * The address prefixes for the subnet.
     */
    readonly addressPrefixes: string[];
    /**
     * Enable or Disable network policies for the private link endpoint on the subnet.
     */
    readonly enforcePrivateLinkEndpointNetworkPolicies: boolean;
    /**
     * Enable or Disable network policies for the private link service on the subnet.
     */
    readonly enforcePrivateLinkServiceNetworkPolicies: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * The ID of the Network Security Group associated with the subnet.
     */
    readonly networkSecurityGroupId: string;
    readonly resourceGroupName: string;
    /**
     * The ID of the Route Table associated with this subnet.
     */
    readonly routeTableId: string;
    /**
     * A list of Service Endpoints within this subnet.
     */
    readonly serviceEndpoints: string[];
    readonly virtualNetworkName: string;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Azure Network DDoS Protection Plan.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.network.getNetworkDdosProtectionPlan({
 *     name: azurerm_network_ddos_protection_plan.example.name,
 *     resourceGroupName: azurerm_network_ddos_protection_plan.example.resource_group_name,
 * });
 * export const ddosProtectionPlanId = example.then(example => example.id);
 * ```
 */
export function getNetworkDdosProtectionPlan(args: GetNetworkDdosProtectionPlanArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkDdosProtectionPlanResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:network/getNetworkDdosProtectionPlan:getNetworkDdosProtectionPlan", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetworkDdosProtectionPlan.
 */
export interface GetNetworkDdosProtectionPlanArgs {
    /**
     * The name of the Network DDoS Protection Plan.
     */
    readonly name: string;
    /**
     * The name of the resource group where the Network DDoS Protection Plan exists.
     */
    readonly resourceGroupName: string;
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getNetworkDdosProtectionPlan.
 */
export interface GetNetworkDdosProtectionPlanResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Specifies the supported Azure location where the resource exists.
     */
    readonly location: string;
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags?: {[key: string]: string};
    /**
     * A list of ID's of the Virtual Networks associated with this DDoS Protection Plan.
     */
    readonly virtualNetworkIds: string[];
}

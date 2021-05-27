// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Event Hubs Consumer Group within an Event Hub.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const test = pulumi.all([azurerm_eventhub_test.name, azurerm_eventhub_consumer_group_test.name, azurerm_eventhub_namespace_test.name, azurerm_resource_group_test.name]).apply(([azurerm_eventhub_testName, azurerm_eventhub_consumer_group_testName, azurerm_eventhub_namespace_testName, azurerm_resource_group_testName]) => azure.eventhub.getConsumeGroup({
 *     eventhubName: azurerm_eventhub_testName,
 *     name: azurerm_eventhub_consumer_group_testName,
 *     namespaceName: azurerm_eventhub_namespace_testName,
 *     resourceGroupName: azurerm_resource_group_testName,
 * }));
 * ```
 */
export function getConsumeGroup(args: GetConsumeGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetConsumeGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:eventhub/getConsumeGroup:getConsumeGroup", {
        "eventhubName": args.eventhubName,
        "name": args.name,
        "namespaceName": args.namespaceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getConsumeGroup.
 */
export interface GetConsumeGroupArgs {
    /**
     * Specifies the name of the EventHub.
     */
    eventhubName: string;
    /**
     * Specifies the name of the EventHub Consumer Group resource.
     */
    name: string;
    /**
     * Specifies the name of the grandparent EventHub Namespace.
     */
    namespaceName: string;
    /**
     * The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getConsumeGroup.
 */
export interface GetConsumeGroupResult {
    readonly eventhubName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly location: string;
    readonly name: string;
    readonly namespaceName: string;
    readonly resourceGroupName: string;
    /**
     * Specifies the user metadata.
     */
    readonly userMetadata: string;
}

export function getConsumeGroupApply(args: GetConsumeGroupApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConsumeGroupResult> {
    return pulumi.output(args).apply(a => getConsumeGroup(a, opts))
}

/**
 * A collection of arguments for invoking getConsumeGroup.
 */
export interface GetConsumeGroupApplyArgs {
    /**
     * Specifies the name of the EventHub.
     */
    eventhubName: pulumi.Input<string>;
    /**
     * Specifies the name of the EventHub Consumer Group resource.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the name of the grandparent EventHub Namespace.
     */
    namespaceName: pulumi.Input<string>;
    /**
     * The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists.
     */
    resourceGroupName: pulumi.Input<string>;
}

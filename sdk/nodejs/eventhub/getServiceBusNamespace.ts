// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing ServiceBus Namespace.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.servicebus.getNamespace({
 *     name: "examplenamespace",
 *     resourceGroupName: "example-resources",
 * });
 * export const location = example.then(example => example.location);
 * ```
 */
/** @deprecated azure.eventhub.getServiceBusNamespace has been deprecated in favor of azure.servicebus.getNamespace */
export function getServiceBusNamespace(args: GetServiceBusNamespaceArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceBusNamespaceResult> {
    pulumi.log.warn("getServiceBusNamespace is deprecated: azure.eventhub.getServiceBusNamespace has been deprecated in favor of azure.servicebus.getNamespace")
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:eventhub/getServiceBusNamespace:getServiceBusNamespace", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceBusNamespace.
 */
export interface GetServiceBusNamespaceArgs {
    /**
     * Specifies the name of the ServiceBus Namespace.
     */
    readonly name: string;
    /**
     * Specifies the name of the Resource Group where the ServiceBus Namespace exists.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getServiceBusNamespace.
 */
export interface GetServiceBusNamespaceResult {
    /**
     * The capacity of the ServiceBus Namespace.
     */
    readonly capacity: number;
    /**
     * The primary connection string for the authorization
     * rule `RootManageSharedAccessKey`.
     */
    readonly defaultPrimaryConnectionString: string;
    /**
     * The primary access key for the authorization rule `RootManageSharedAccessKey`.
     */
    readonly defaultPrimaryKey: string;
    /**
     * The secondary connection string for the
     * authorization rule `RootManageSharedAccessKey`.
     */
    readonly defaultSecondaryConnectionString: string;
    /**
     * The secondary access key for the authorization rule `RootManageSharedAccessKey`.
     */
    readonly defaultSecondaryKey: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The location of the Resource Group in which the ServiceBus Namespace exists.
     */
    readonly location: string;
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * The Tier used for the ServiceBus Namespace.
     */
    readonly sku: string;
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags: {[key: string]: string};
    /**
     * Whether or not this ServiceBus Namespace is zone redundant.
     */
    readonly zoneRedundant: boolean;
}

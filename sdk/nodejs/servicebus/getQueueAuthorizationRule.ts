// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing ServiceBus Queue Authorisation Rule within a ServiceBus Queue.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.servicebus.getQueueAuthorizationRule({
 *     name: "example-tfex_name",
 *     resourceGroupName: "example-resources",
 *     queueName: "example-servicebus_queue",
 *     namespaceName: "example-namespace",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getQueueAuthorizationRule(args: GetQueueAuthorizationRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetQueueAuthorizationRuleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:servicebus/getQueueAuthorizationRule:getQueueAuthorizationRule", {
        "name": args.name,
        "namespaceName": args.namespaceName,
        "queueName": args.queueName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getQueueAuthorizationRule.
 */
export interface GetQueueAuthorizationRuleArgs {
    /**
     * The name of this ServiceBus Queue Authorisation Rule.
     */
    name: string;
    /**
     * The name of the ServiceBus Namespace.
     */
    namespaceName: string;
    /**
     * The name of the ServiceBus Queue.
     */
    queueName: string;
    /**
     * The name of the Resource Group where the ServiceBus Queue Authorisation Rule exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getQueueAuthorizationRule.
 */
export interface GetQueueAuthorizationRuleResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly listen: boolean;
    readonly manage: boolean;
    readonly name: string;
    readonly namespaceName: string;
    /**
     * The Primary Connection String for the ServiceBus Queue authorization Rule.
     */
    readonly primaryConnectionString: string;
    /**
     * The Primary Key for the ServiceBus Queue authorization Rule.
     */
    readonly primaryKey: string;
    readonly queueName: string;
    readonly resourceGroupName: string;
    /**
     * The Secondary Connection String for the ServiceBus Queue authorization Rule.
     */
    readonly secondaryConnectionString: string;
    /**
     * The Secondary Key for the ServiceBus Queue authorization Rule.
     */
    readonly secondaryKey: string;
    readonly send: boolean;
}

export function getQueueAuthorizationRuleOutput(args: GetQueueAuthorizationRuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetQueueAuthorizationRuleResult> {
    return pulumi.output(args).apply(a => getQueueAuthorizationRule(a, opts))
}

/**
 * A collection of arguments for invoking getQueueAuthorizationRule.
 */
export interface GetQueueAuthorizationRuleOutputArgs {
    /**
     * The name of this ServiceBus Queue Authorisation Rule.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the ServiceBus Namespace.
     */
    namespaceName: pulumi.Input<string>;
    /**
     * The name of the ServiceBus Queue.
     */
    queueName: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the ServiceBus Queue Authorisation Rule exists.
     */
    resourceGroupName: pulumi.Input<string>;
}

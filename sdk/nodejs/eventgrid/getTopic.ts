// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing EventGrid Topic
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = pulumi.output(azure.eventgrid.getTopic({
 *     name: "my-eventgrid-topic",
 *     resourceGroupName: "example-resources",
 * }, { async: true }));
 * ```
 */
export function getTopic(args: GetTopicArgs, opts?: pulumi.InvokeOptions): Promise<GetTopicResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:eventgrid/getTopic:getTopic", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getTopic.
 */
export interface GetTopicArgs {
    /**
     * The name of the EventGrid Topic resource.
     */
    name: string;
    /**
     * The name of the resource group in which the EventGrid Topic exists.
     */
    resourceGroupName: string;
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getTopic.
 */
export interface GetTopicResult {
    /**
     * The Endpoint associated with the EventGrid Topic.
     */
    readonly endpoint: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly location: string;
    readonly name: string;
    /**
     * The Primary Shared Access Key associated with the EventGrid Topic.
     */
    readonly primaryAccessKey: string;
    readonly resourceGroupName: string;
    /**
     * The Secondary Shared Access Key associated with the EventGrid Topic.
     */
    readonly secondaryAccessKey: string;
    readonly tags?: {[key: string]: string};
}

export function getTopicOutput(args: GetTopicOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTopicResult> {
    return pulumi.output(args).apply(a => getTopic(a, opts))
}

/**
 * A collection of arguments for invoking getTopic.
 */
export interface GetTopicOutputArgs {
    /**
     * The name of the EventGrid Topic resource.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the resource group in which the EventGrid Topic exists.
     */
    resourceGroupName: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

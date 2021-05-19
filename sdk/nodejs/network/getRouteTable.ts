// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Route Table.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = pulumi.output(azure.network.getRouteTable({
 *     name: "myroutetable",
 *     resourceGroupName: "some-resource-group",
 * }, { async: true }));
 * ```
 */
export function getRouteTable(args: GetRouteTableArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteTableResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:network/getRouteTable:getRouteTable", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouteTable.
 */
export interface GetRouteTableArgs {
    /**
     * The name of the Route Table.
     */
    name: string;
    /**
     * The name of the Resource Group in which the Route Table exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getRouteTable.
 */
export interface GetRouteTableResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Azure Region in which the Route Table exists.
     */
    readonly location: string;
    /**
     * The name of the Route.
     */
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * One or more `route` blocks as documented below.
     */
    readonly routes: outputs.network.GetRouteTableRoute[];
    /**
     * The collection of Subnets associated with this route table.
     */
    readonly subnets: string[];
    /**
     * A mapping of tags assigned to the Route Table.
     */
    readonly tags: {[key: string]: string};
}

export function getRouteTableOutput(args: GetRouteTableOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouteTableResult> {
    return pulumi.output(args).apply(a => getRouteTable(a, opts))
}

/**
 * A collection of arguments for invoking getRouteTable.
 */
export interface GetRouteTableOutputArgs {
    /**
     * The name of the Route Table.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the Route Table exists.
     */
    resourceGroupName: pulumi.Input<string>;
}

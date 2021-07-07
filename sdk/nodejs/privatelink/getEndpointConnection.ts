// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access the connection status information about an existing Private Endpoint Connection.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.privatelink.getEndpointConnection({
 *     name: "example-private-endpoint",
 *     resourceGroupName: "example-rg",
 * });
 * export const privateEndpointStatus = example.then(example => example.privateServiceConnections[0].status);
 * ```
 */
export function getEndpointConnection(args: GetEndpointConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetEndpointConnectionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:privatelink/getEndpointConnection:getEndpointConnection", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getEndpointConnection.
 */
export interface GetEndpointConnectionArgs {
    /**
     * Specifies the Name of the private endpoint.
     */
    name: string;
    /**
     * Specifies the Name of the Resource Group within which the private endpoint exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getEndpointConnection.
 */
export interface GetEndpointConnectionResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The supported Azure location where the resource exists.
     */
    readonly location: string;
    /**
     * The name of the private endpoint.
     */
    readonly name: string;
    readonly privateServiceConnections: outputs.privatelink.GetEndpointConnectionPrivateServiceConnection[];
    readonly resourceGroupName: string;
}

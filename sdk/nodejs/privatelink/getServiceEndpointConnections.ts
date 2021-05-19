// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access endpoint connection information about an existing Private Link Service.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.privatelink.getServiceEndpointConnections({
 *     serviceId: azurerm_private_link_service.example.id,
 *     resourceGroupName: azurerm_resource_group.example.name,
 * });
 * export const privateEndpointStatus = example.then(example => example.privateEndpointConnections[0].status);
 * ```
 */
export function getServiceEndpointConnections(args: GetServiceEndpointConnectionsArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceEndpointConnectionsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:privatelink/getServiceEndpointConnections:getServiceEndpointConnections", {
        "resourceGroupName": args.resourceGroupName,
        "serviceId": args.serviceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceEndpointConnections.
 */
export interface GetServiceEndpointConnectionsArgs {
    /**
     * The name of the resource group in which the private link service resides.
     */
    resourceGroupName: string;
    /**
     * The resource ID of the private link service.
     */
    serviceId: string;
}

/**
 * A collection of values returned by getServiceEndpointConnections.
 */
export interface GetServiceEndpointConnectionsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly location: string;
    readonly privateEndpointConnections: outputs.privatelink.GetServiceEndpointConnectionsPrivateEndpointConnection[];
    readonly resourceGroupName: string;
    readonly serviceId: string;
    /**
     * The name of the private link service.
     */
    readonly serviceName: string;
}

export function getServiceEndpointConnectionsOutput(args: GetServiceEndpointConnectionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceEndpointConnectionsResult> {
    return pulumi.output(args).apply(a => getServiceEndpointConnections(a, opts))
}

/**
 * A collection of arguments for invoking getServiceEndpointConnections.
 */
export interface GetServiceEndpointConnectionsOutputArgs {
    /**
     * The name of the resource group in which the private link service resides.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The resource ID of the private link service.
     */
    serviceId: pulumi.Input<string>;
}

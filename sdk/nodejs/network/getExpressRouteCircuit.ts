// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing ExpressRoute circuit.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.network.getExpressRouteCircuit({
 *     resourceGroupName: azurerm_resource_group.example.name,
 *     name: azurerm_express_route_circuit.example.name,
 * });
 * export const expressRouteCircuitId = example.then(example => example.id);
 * export const serviceKey = example.then(example => example.serviceKey);
 * ```
 */
export function getExpressRouteCircuit(args: GetExpressRouteCircuitArgs, opts?: pulumi.InvokeOptions): Promise<GetExpressRouteCircuitResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:network/getExpressRouteCircuit:getExpressRouteCircuit", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getExpressRouteCircuit.
 */
export interface GetExpressRouteCircuitArgs {
    /**
     * The name of the ExpressRoute circuit.
     */
    name: string;
    /**
     * The Name of the Resource Group where the ExpressRoute circuit exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getExpressRouteCircuit.
 */
export interface GetExpressRouteCircuitResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Azure location where the ExpressRoute circuit exists
     */
    readonly location: string;
    readonly name: string;
    /**
     * A `peerings` block for the ExpressRoute circuit as documented below
     */
    readonly peerings: outputs.network.GetExpressRouteCircuitPeering[];
    readonly resourceGroupName: string;
    /**
     * The string needed by the service provider to provision the ExpressRoute circuit.
     */
    readonly serviceKey: string;
    /**
     * A `serviceProviderProperties` block for the ExpressRoute circuit as documented below
     */
    readonly serviceProviderProperties: outputs.network.GetExpressRouteCircuitServiceProviderProperty[];
    /**
     * The ExpressRoute circuit provisioning state from your chosen service provider. Possible values are "NotProvisioned", "Provisioning", "Provisioned", and "Deprovisioning".
     */
    readonly serviceProviderProvisioningState: string;
    /**
     * A `sku` block for the ExpressRoute circuit as documented below.
     */
    readonly sku: outputs.network.GetExpressRouteCircuitSku;
}

export function getExpressRouteCircuitApply(args: GetExpressRouteCircuitApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetExpressRouteCircuitResult> {
    return pulumi.output(args).apply(a => getExpressRouteCircuit(a, opts))
}

/**
 * A collection of arguments for invoking getExpressRouteCircuit.
 */
export interface GetExpressRouteCircuitApplyArgs {
    /**
     * The name of the ExpressRoute circuit.
     */
    name: pulumi.Input<string>;
    /**
     * The Name of the Resource Group where the ExpressRoute circuit exists.
     */
    resourceGroupName: pulumi.Input<string>;
}

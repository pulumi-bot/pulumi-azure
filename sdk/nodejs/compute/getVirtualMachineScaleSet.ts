// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Virtual Machine Scale Set.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.compute.getVirtualMachineScaleSet({
 *     name: "existing",
 *     resourceGroupName: "existing",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getVirtualMachineScaleSet(args: GetVirtualMachineScaleSetArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualMachineScaleSetResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:compute/getVirtualMachineScaleSet:getVirtualMachineScaleSet", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getVirtualMachineScaleSet.
 */
export interface GetVirtualMachineScaleSetArgs {
    /**
     * The name of this Virtual Machine Scale Set.
     */
    readonly name: string;
    /**
     * The name of the Resource Group where the Virtual Machine Scale Set exists.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getVirtualMachineScaleSet.
 */
export interface GetVirtualMachineScaleSetResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A `identity` block as defined below.
     */
    readonly identities: outputs.compute.GetVirtualMachineScaleSetIdentity[];
    readonly location: string;
    readonly name: string;
    readonly resourceGroupName: string;
}

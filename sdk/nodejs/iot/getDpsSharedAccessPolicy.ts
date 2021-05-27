// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing IotHub Device Provisioning Service Shared Access Policy
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.iot.getDpsSharedAccessPolicy({
 *     name: "example",
 *     resourceGroupName: azurerm_resource_group.example.name,
 *     iothubDpsName: azurerm_iothub_dps.example.name,
 * });
 * ```
 */
export function getDpsSharedAccessPolicy(args: GetDpsSharedAccessPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetDpsSharedAccessPolicyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:iot/getDpsSharedAccessPolicy:getDpsSharedAccessPolicy", {
        "iothubDpsName": args.iothubDpsName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getDpsSharedAccessPolicy.
 */
export interface GetDpsSharedAccessPolicyArgs {
    /**
     * Specifies the name of the IoT Hub Device Provisioning service to which the Shared Access Policy belongs.
     */
    iothubDpsName: string;
    /**
     * Specifies the name of the IotHub Shared Access Policy.
     */
    name: string;
    /**
     * Specifies the name of the resource group under which the IotHub Shared Access Policy resource exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getDpsSharedAccessPolicy.
 */
export interface GetDpsSharedAccessPolicyResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly iothubDpsName: string;
    readonly name: string;
    /**
     * The primary connection string of the Shared Access Policy.
     */
    readonly primaryConnectionString: string;
    /**
     * The primary key used to create the authentication token.
     */
    readonly primaryKey: string;
    readonly resourceGroupName: string;
    /**
     * The secondary connection string of the Shared Access Policy.
     */
    readonly secondaryConnectionString: string;
    /**
     * The secondary key used to create the authentication token.
     */
    readonly secondaryKey: string;
}

export function getDpsSharedAccessPolicyApply(args: GetDpsSharedAccessPolicyApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDpsSharedAccessPolicyResult> {
    return pulumi.output(args).apply(a => getDpsSharedAccessPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getDpsSharedAccessPolicy.
 */
export interface GetDpsSharedAccessPolicyApplyArgs {
    /**
     * Specifies the name of the IoT Hub Device Provisioning service to which the Shared Access Policy belongs.
     */
    iothubDpsName: pulumi.Input<string>;
    /**
     * Specifies the name of the IotHub Shared Access Policy.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the name of the resource group under which the IotHub Shared Access Policy resource exists.
     */
    resourceGroupName: pulumi.Input<string>;
}

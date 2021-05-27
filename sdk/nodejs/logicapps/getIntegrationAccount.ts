// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Logic App Integration Account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.logicapps.getIntegrationAccount({
 *     name: "example-account",
 *     resourceGroupName: "example-resource-group",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getIntegrationAccount(args: GetIntegrationAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetIntegrationAccountResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:logicapps/getIntegrationAccount:getIntegrationAccount", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getIntegrationAccount.
 */
export interface GetIntegrationAccountArgs {
    /**
     * The name of this Logic App Integration Account.
     */
    name: string;
    /**
     * The name of the Resource Group where the Logic App Integration Account exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getIntegrationAccount.
 */
export interface GetIntegrationAccountResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Azure Region where the Logic App Integration Account exists.
     */
    readonly location: string;
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * The sku name of the Logic App Integration Account.
     */
    readonly skuName: string;
    /**
     * A mapping of tags assigned to the Logic App Integration Account.
     */
    readonly tags: {[key: string]: string};
}

export function getIntegrationAccountApply(args: GetIntegrationAccountApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIntegrationAccountResult> {
    return pulumi.output(args).apply(a => getIntegrationAccount(a, opts))
}

/**
 * A collection of arguments for invoking getIntegrationAccount.
 */
export interface GetIntegrationAccountApplyArgs {
    /**
     * The name of this Logic App Integration Account.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Logic App Integration Account exists.
     */
    resourceGroupName: pulumi.Input<string>;
}

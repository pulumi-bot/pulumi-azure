// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Storage Management Policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleAccount = azure.storage.getAccount({
 *     name: "storageaccountname",
 *     resourceGroupName: "resourcegroupname",
 * });
 * const examplePolicy = azure.storage.getPolicy({
 *     storageAccountId: azurerm_storage_account.example.id,
 * });
 * ```
 */
export function getPolicy(args: GetPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:storage/getPolicy:getPolicy", {
        "storageAccountId": args.storageAccountId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPolicy.
 */
export interface GetPolicyArgs {
    /**
     * Specifies the id of the storage account to retrieve the management policy for.
     */
    storageAccountId: string;
}

/**
 * A collection of values returned by getPolicy.
 */
export interface GetPolicyResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A `rule` block as documented below.
     */
    readonly rules: outputs.storage.GetPolicyRule[];
    readonly storageAccountId: string;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access an ID for your MCA Account billing scope.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.billing.getMcaAccountScope({
 *     billingAccountName: "e879cf0f-2b4d-5431-109a-f72fc9868693:024cabf4-7321-4cf9-be59-df0c77ca51de_2019-05-31",
 *     billingProfileName: "PE2Q-NOIT-BG7-TGB",
 *     invoiceSectionName: "MTT4-OBS7-PJA-TGB",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getMcaAccountScope(args: GetMcaAccountScopeArgs, opts?: pulumi.InvokeOptions): Promise<GetMcaAccountScopeResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:billing/getMcaAccountScope:getMcaAccountScope", {
        "billingAccountName": args.billingAccountName,
        "billingProfileName": args.billingProfileName,
        "invoiceSectionName": args.invoiceSectionName,
    }, opts);
}

/**
 * A collection of arguments for invoking getMcaAccountScope.
 */
export interface GetMcaAccountScopeArgs {
    /**
     * The Billing Account Name of the MCA account.
     */
    billingAccountName: string;
    /**
     * The Billing Profile Name in the above Billing Account.
     */
    billingProfileName: string;
    /**
     * The Invoice Section Name in the above Billing Profile.
     */
    invoiceSectionName: string;
}

/**
 * A collection of values returned by getMcaAccountScope.
 */
export interface GetMcaAccountScopeResult {
    readonly billingAccountName: string;
    readonly billingProfileName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly invoiceSectionName: string;
}

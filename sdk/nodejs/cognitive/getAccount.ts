// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Cognitive Services Account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const test = pulumi.output(azure.cognitive.getAccount({
 *     name: "example-account",
 *     resourceGroupName: "cognitive_account_rg",
 * }));
 *
 * export const primaryAccessKey = test.primaryAccessKey;
 * ```
 */
export function getAccount(args: GetAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:cognitive/getAccount:getAccount", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccount.
 */
export interface GetAccountArgs {
    /**
     * Specifies the name of the Cognitive Services Account.
     */
    name: string;
    /**
     * Specifies the name of the resource group where the Cognitive Services Account resides.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getAccount.
 */
export interface GetAccountResult {
    /**
     * The endpoint of the Cognitive Services Account
     */
    readonly endpoint: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The kind of the Cognitive Services Account
     */
    readonly kind: string;
    /**
     * The Azure location where the Cognitive Services Account exists
     */
    readonly location: string;
    readonly name: string;
    /**
     * The primary access key of the Cognitive Services Account
     */
    readonly primaryAccessKey: string;
    /**
     * If `kind` is `QnAMaker` the link to the QNA runtime.
     */
    readonly qnaRuntimeEndpoint: string;
    readonly resourceGroupName: string;
    /**
     * The secondary access key of the Cognitive Services Account
     */
    readonly secondaryAccessKey: string;
    /**
     * The sku name of the Cognitive Services Account
     */
    readonly skuName: string;
    /**
     * A mapping of tags to assigned to the resource.
     */
    readonly tags: {[key: string]: string};
}

export function getAccountApply(args: GetAccountApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccountResult> {
    return pulumi.output(args).apply(a => getAccount(a, opts))
}

/**
 * A collection of arguments for invoking getAccount.
 */
export interface GetAccountApplyArgs {
    /**
     * Specifies the name of the Cognitive Services Account.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the name of the resource group where the Cognitive Services Account resides.
     */
    resourceGroupName: pulumi.Input<string>;
}

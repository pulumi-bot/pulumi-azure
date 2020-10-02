// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Use this data source to access the configuration of the AzureRM provider.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const current = azure.core.getClientConfig({});
 * export const accountId = current.then(current => current.clientId);
 * ```
 */
export function getClientConfig(opts?: pulumi.InvokeOptions): Promise<GetClientConfigResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:core/getClientConfig:getClientConfig", {
    }, opts);
}

/**
 * A collection of values returned by getClientConfig.
 */
export interface GetClientConfigResult {
    readonly clientId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly objectId: string;
    readonly subscriptionId: string;
    readonly tenantId: string;
}

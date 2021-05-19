// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Key Vault Managed Hardware Security Module.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.keyvault.getManagedHardwareSecurityModule({
 *     name: "mykeyvaultHsm",
 *     resourceGroupName: "some-resource-group",
 * });
 * export const hsmUri = example.then(example => example.hsmUri);
 * ```
 */
export function getManagedHardwareSecurityModule(args: GetManagedHardwareSecurityModuleArgs, opts?: pulumi.InvokeOptions): Promise<GetManagedHardwareSecurityModuleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:keyvault/getManagedHardwareSecurityModule:getManagedHardwareSecurityModule", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getManagedHardwareSecurityModule.
 */
export interface GetManagedHardwareSecurityModuleArgs {
    /**
     * The name of the Key Vault Managed Hardware Security Module.
     */
    name: string;
    /**
     * The name of the Resource Group in which the Key Vault Managed Hardware Security Module exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getManagedHardwareSecurityModule.
 */
export interface GetManagedHardwareSecurityModuleResult {
    /**
     * Specifies a list of administrators object IDs for the key vault Managed Hardware Security Module.
     */
    readonly adminObjectIds: string[];
    /**
     * The URI of the Hardware Security Module for performing operations on keys and secrets.
     */
    readonly hsmUri: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Azure Region in which the Key Vault managed Hardware Security Module exists.
     */
    readonly location: string;
    readonly name: string;
    /**
     * Is purge protection enabled on this Key Vault Managed Hardware Security Module?
     */
    readonly purgeProtectionEnabled: boolean;
    readonly resourceGroupName: string;
    /**
     * The Name of the SKU used for this Key Vault Managed Hardware Security Module.
     */
    readonly skuName: string;
    /**
     * The number of days that items should be retained for soft-deleted.
     */
    readonly softDeleteRetentionDays: number;
    /**
     * A mapping of tags assigned to the Key Vault Managed Hardware Security Module.
     */
    readonly tags: {[key: string]: string};
    /**
     * The Azure Active Directory Tenant ID used for authenticating requests to the Key Vault Managed Hardware Security Module.
     */
    readonly tenantId: string;
}

export function getManagedHardwareSecurityModuleOutput(args: GetManagedHardwareSecurityModuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetManagedHardwareSecurityModuleResult> {
    return pulumi.output(args).apply(a => getManagedHardwareSecurityModule(a, opts))
}

/**
 * A collection of arguments for invoking getManagedHardwareSecurityModule.
 */
export interface GetManagedHardwareSecurityModuleOutputArgs {
    /**
     * The name of the Key Vault Managed Hardware Security Module.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the Key Vault Managed Hardware Security Module exists.
     */
    resourceGroupName: pulumi.Input<string>;
}

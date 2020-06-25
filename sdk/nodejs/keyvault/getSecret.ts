// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Key Vault Secret.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.keyvault.getSecret({
 *     name: "secret-sauce",
 *     keyVaultId: data.azurerm_key_vault.existing.id,
 * });
 * export const secretValue = example.then(example => example.value);
 * ```
 */
export function getSecret(args: GetSecretArgs, opts?: pulumi.InvokeOptions): Promise<GetSecretResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:keyvault/getSecret:getSecret", {
        "keyVaultId": args.keyVaultId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecret.
 */
export interface GetSecretArgs {
    /**
     * Specifies the ID of the Key Vault instance where the Secret resides, available on the `azure.keyvault.KeyVault` Data Source / Resource.
     */
    readonly keyVaultId: string;
    /**
     * Specifies the name of the Key Vault Secret.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getSecret.
 */
export interface GetSecretResult {
    /**
     * The content type for the Key Vault Secret.
     */
    readonly contentType: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly keyVaultId: string;
    readonly name: string;
    /**
     * Any tags assigned to this resource.
     */
    readonly tags: {[key: string]: string};
    /**
     * The value of the Key Vault Secret.
     */
    readonly value: string;
    /**
     * The current version of the Key Vault Secret.
     */
    readonly version: string;
}

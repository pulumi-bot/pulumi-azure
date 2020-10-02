// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Key Vault Key.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.keyvault.getKey({
 *     name: "secret-sauce",
 *     keyVaultId: data.azurerm_key_vault.existing.id,
 * });
 * export const keyType = example.then(example => example.keyType);
 * ```
 */
export function getKey(args: GetKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetKeyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:keyvault/getKey:getKey", {
        "keyVaultId": args.keyVaultId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getKey.
 */
export interface GetKeyArgs {
    /**
     * Specifies the ID of the Key Vault instance where the Secret resides, available on the `azure.keyvault.KeyVault` Data Source / Resource.
     */
    readonly keyVaultId: string;
    /**
     * Specifies the name of the Key Vault Key.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getKey.
 */
export interface GetKeyResult {
    /**
     * The RSA public exponent of this Key Vault Key.
     */
    readonly e: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of JSON web key operations assigned to this Key Vault Key
     */
    readonly keyOpts: string[];
    /**
     * Specifies the Size of this Key Vault Key.
     */
    readonly keySize: number;
    /**
     * Specifies the Key Type of this Key Vault Key
     */
    readonly keyType: string;
    readonly keyVaultId: string;
    /**
     * The RSA modulus of this Key Vault Key.
     */
    readonly n: string;
    readonly name: string;
    /**
     * A mapping of tags assigned to this Key Vault Key.
     */
    readonly tags: {[key: string]: string};
    /**
     * The current version of the Key Vault Key.
     */
    readonly version: string;
}

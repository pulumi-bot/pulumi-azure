// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Storage Encryption Scope.
 *
 * > **Note:** Storage Encryption Scopes are in Preview [more information can be found here](https://docs.microsoft.com/en-us/azure/storage/blobs/encryption-scope-manage).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 *     identity: {
 *         type: "SystemAssigned",
 *     },
 * });
 * const exampleEncryptionScope = new azure.storage.EncryptionScope("exampleEncryptionScope", {
 *     storageAccountId: exampleAccount.id,
 *     source: "Microsoft.Storage",
 * });
 * ```
 *
 * ## Import
 *
 * Storage Encryption Scopes can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:storage/encryptionScope:EncryptionScope example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Storage/storageAccounts/account1/encryptionScopes/scope1
 * ```
 */
export class EncryptionScope extends pulumi.CustomResource {
    /**
     * Get an existing EncryptionScope resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EncryptionScopeState, opts?: pulumi.CustomResourceOptions): EncryptionScope {
        return new EncryptionScope(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:storage/encryptionScope:EncryptionScope';

    /**
     * Returns true if the given object is an instance of EncryptionScope.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EncryptionScope {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EncryptionScope.__pulumiType;
    }

    /**
     * The ID of the Key Vault Key. Required when `source` is `Microsoft.KeyVault`.
     */
    public readonly keyVaultKeyId!: pulumi.Output<string | undefined>;
    /**
     * The name which should be used for this Storage Encryption Scope. Changing this forces a new Storage Encryption Scope to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The source of the Storage Encryption Scope. Possible values are `Microsoft.KeyVault` and `Microsoft.Storage`.
     */
    public readonly source!: pulumi.Output<string>;
    /**
     * The ID of the Storage Account where this Storage Encryption Scope is created. Changing this forces a new Storage Encryption Scope to be created.
     */
    public readonly storageAccountId!: pulumi.Output<string>;

    /**
     * Create a EncryptionScope resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EncryptionScopeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EncryptionScopeArgs | EncryptionScopeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EncryptionScopeState | undefined;
            inputs["keyVaultKeyId"] = state ? state.keyVaultKeyId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["source"] = state ? state.source : undefined;
            inputs["storageAccountId"] = state ? state.storageAccountId : undefined;
        } else {
            const args = argsOrState as EncryptionScopeArgs | undefined;
            if (!args || args.source === undefined) {
                throw new Error("Missing required property 'source'");
            }
            if (!args || args.storageAccountId === undefined) {
                throw new Error("Missing required property 'storageAccountId'");
            }
            inputs["keyVaultKeyId"] = args ? args.keyVaultKeyId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["storageAccountId"] = args ? args.storageAccountId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EncryptionScope.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EncryptionScope resources.
 */
export interface EncryptionScopeState {
    /**
     * The ID of the Key Vault Key. Required when `source` is `Microsoft.KeyVault`.
     */
    readonly keyVaultKeyId?: pulumi.Input<string>;
    /**
     * The name which should be used for this Storage Encryption Scope. Changing this forces a new Storage Encryption Scope to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The source of the Storage Encryption Scope. Possible values are `Microsoft.KeyVault` and `Microsoft.Storage`.
     */
    readonly source?: pulumi.Input<string>;
    /**
     * The ID of the Storage Account where this Storage Encryption Scope is created. Changing this forces a new Storage Encryption Scope to be created.
     */
    readonly storageAccountId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EncryptionScope resource.
 */
export interface EncryptionScopeArgs {
    /**
     * The ID of the Key Vault Key. Required when `source` is `Microsoft.KeyVault`.
     */
    readonly keyVaultKeyId?: pulumi.Input<string>;
    /**
     * The name which should be used for this Storage Encryption Scope. Changing this forces a new Storage Encryption Scope to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The source of the Storage Encryption Scope. Possible values are `Microsoft.KeyVault` and `Microsoft.Storage`.
     */
    readonly source: pulumi.Input<string>;
    /**
     * The ID of the Storage Account where this Storage Encryption Scope is created. Changing this forces a new Storage Encryption Scope to be created.
     */
    readonly storageAccountId: pulumi.Input<string>;
}

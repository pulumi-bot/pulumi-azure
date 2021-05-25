// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Key Vault Access Policy.
 *
 * > **NOTE:** It's possible to define Key Vault Access Policies both within the `azure.keyvault.KeyVault` resource via the `accessPolicy` block and by using the `azure.keyvault.AccessPolicy` resource. However it's not possible to use both methods to manage Access Policies within a KeyVault, since there'll be conflicts.
 *
 * > **NOTE:** Azure permits a maximum of 1024 Access Policies per Key Vault - [more information can be found in this document](https://docs.microsoft.com/en-us/azure/key-vault/key-vault-secure-your-key-vault#data-plane-access-control).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const current = azure.core.getClientConfig({});
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleKeyVault = new azure.keyvault.KeyVault("exampleKeyVault", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     tenantId: current.then(current => current.tenantId),
 *     skuName: "premium",
 * });
 * const exampleAccessPolicy = new azure.keyvault.AccessPolicy("exampleAccessPolicy", {
 *     keyVaultId: exampleKeyVault.id,
 *     tenantId: current.then(current => current.tenantId),
 *     objectId: current.then(current => current.objectId),
 *     keyPermissions: ["Get"],
 *     secretPermissions: ["Get"],
 * });
 * ```
 *
 * ## Import
 *
 * Key Vault Access Policies can be imported using the Resource ID of the Key Vault, plus some additional metadata. If both an `object_id` and `application_id` are specified, then the Access Policy can be imported using the following code
 *
 * ```sh
 *  $ pulumi import azure:keyvault/accessPolicy:AccessPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.KeyVault/vaults/test-vault/objectId/11111111-1111-1111-1111-111111111111/applicationId/22222222-2222-2222-2222-222222222222
 * ```
 *
 *  where `11111111-1111-1111-1111-111111111111` is the `object_id` and `22222222-2222-2222-2222-222222222222` is the `application_id`. --- Access Policies with an `object_id` but no `application_id` can be imported using the following command
 *
 * ```sh
 *  $ pulumi import azure:keyvault/accessPolicy:AccessPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.KeyVault/vaults/test-vault/objectId/11111111-1111-1111-1111-111111111111
 * ```
 *
 *  where `11111111-1111-1111-1111-111111111111` is the `object_id`.
 */
export class AccessPolicy extends pulumi.CustomResource {
    /**
     * Get an existing AccessPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccessPolicyState, opts?: pulumi.CustomResourceOptions): AccessPolicy {
        return new AccessPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:keyvault/accessPolicy:AccessPolicy';

    /**
     * Returns true if the given object is an instance of AccessPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessPolicy.__pulumiType;
    }

    /**
     * The object ID of an Application in Azure Active Directory.
     */
    public readonly applicationId!: pulumi.Output<string | undefined>;
    /**
     * List of certificate permissions, must be one or more from the following: `Backup`, `Create`, `Delete`, `DeleteIssuers`, `Get`, `GetIssuers`, `Import`, `List`, `ListIssuers`, `ManageContacts`, `ManageIssuers`, `Purge`, `Recover`, `Restore`, `SetIssuers` and `Update`.
     */
    public readonly certificatePermissions!: pulumi.Output<string[] | undefined>;
    /**
     * List of key permissions, must be one or more from the following: `Backup`, `Create`, `Decrypt`, `Delete`, `Encrypt`, `Get`, `Import`, `List`, `Purge`, `Recover`, `Restore`, `Sign`, `UnwrapKey`, `Update`, `Verify` and `WrapKey`.
     */
    public readonly keyPermissions!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the id of the Key Vault resource. Changing this
     * forces a new resource to be created.
     */
    public readonly keyVaultId!: pulumi.Output<string>;
    /**
     * The object ID of a user, service principal or security
     * group in the Azure Active Directory tenant for the vault. The object ID must
     * be unique for the list of access policies. Changing this forces a new resource
     * to be created.
     */
    public readonly objectId!: pulumi.Output<string>;
    /**
     * List of secret permissions, must be one or more from the following: `Backup`, `Delete`, `get`, `list`, `purge`, `recover`, `restore` and `set`.
     */
    public readonly secretPermissions!: pulumi.Output<string[] | undefined>;
    /**
     * List of storage permissions, must be one or more from the following: `Backup`, `Delete`, `DeleteSAS`, `Get`, `GetSAS`, `List`, `ListSAS`, `Purge`, `Recover`, `RegenerateKey`, `Restore`, `Set`, `SetSAS` and `Update`.
     */
    public readonly storagePermissions!: pulumi.Output<string[] | undefined>;
    /**
     * The Azure Active Directory tenant ID that should be used
     * for authenticating requests to the key vault. Changing this forces a new resource
     * to be created.
     */
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a AccessPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessPolicyArgs | AccessPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccessPolicyState | undefined;
            inputs["applicationId"] = state ? state.applicationId : undefined;
            inputs["certificatePermissions"] = state ? state.certificatePermissions : undefined;
            inputs["keyPermissions"] = state ? state.keyPermissions : undefined;
            inputs["keyVaultId"] = state ? state.keyVaultId : undefined;
            inputs["objectId"] = state ? state.objectId : undefined;
            inputs["secretPermissions"] = state ? state.secretPermissions : undefined;
            inputs["storagePermissions"] = state ? state.storagePermissions : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as AccessPolicyArgs | undefined;
            if ((!args || args.keyVaultId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyVaultId'");
            }
            if ((!args || args.objectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'objectId'");
            }
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            inputs["applicationId"] = args ? args.applicationId : undefined;
            inputs["certificatePermissions"] = args ? args.certificatePermissions : undefined;
            inputs["keyPermissions"] = args ? args.keyPermissions : undefined;
            inputs["keyVaultId"] = args ? args.keyVaultId : undefined;
            inputs["objectId"] = args ? args.objectId : undefined;
            inputs["secretPermissions"] = args ? args.secretPermissions : undefined;
            inputs["storagePermissions"] = args ? args.storagePermissions : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AccessPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccessPolicy resources.
 */
export interface AccessPolicyState {
    /**
     * The object ID of an Application in Azure Active Directory.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * List of certificate permissions, must be one or more from the following: `Backup`, `Create`, `Delete`, `DeleteIssuers`, `Get`, `GetIssuers`, `Import`, `List`, `ListIssuers`, `ManageContacts`, `ManageIssuers`, `Purge`, `Recover`, `Restore`, `SetIssuers` and `Update`.
     */
    certificatePermissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of key permissions, must be one or more from the following: `Backup`, `Create`, `Decrypt`, `Delete`, `Encrypt`, `Get`, `Import`, `List`, `Purge`, `Recover`, `Restore`, `Sign`, `UnwrapKey`, `Update`, `Verify` and `WrapKey`.
     */
    keyPermissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the id of the Key Vault resource. Changing this
     * forces a new resource to be created.
     */
    keyVaultId?: pulumi.Input<string>;
    /**
     * The object ID of a user, service principal or security
     * group in the Azure Active Directory tenant for the vault. The object ID must
     * be unique for the list of access policies. Changing this forces a new resource
     * to be created.
     */
    objectId?: pulumi.Input<string>;
    /**
     * List of secret permissions, must be one or more from the following: `Backup`, `Delete`, `get`, `list`, `purge`, `recover`, `restore` and `set`.
     */
    secretPermissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of storage permissions, must be one or more from the following: `Backup`, `Delete`, `DeleteSAS`, `Get`, `GetSAS`, `List`, `ListSAS`, `Purge`, `Recover`, `RegenerateKey`, `Restore`, `Set`, `SetSAS` and `Update`.
     */
    storagePermissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Azure Active Directory tenant ID that should be used
     * for authenticating requests to the key vault. Changing this forces a new resource
     * to be created.
     */
    tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccessPolicy resource.
 */
export interface AccessPolicyArgs {
    /**
     * The object ID of an Application in Azure Active Directory.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * List of certificate permissions, must be one or more from the following: `Backup`, `Create`, `Delete`, `DeleteIssuers`, `Get`, `GetIssuers`, `Import`, `List`, `ListIssuers`, `ManageContacts`, `ManageIssuers`, `Purge`, `Recover`, `Restore`, `SetIssuers` and `Update`.
     */
    certificatePermissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of key permissions, must be one or more from the following: `Backup`, `Create`, `Decrypt`, `Delete`, `Encrypt`, `Get`, `Import`, `List`, `Purge`, `Recover`, `Restore`, `Sign`, `UnwrapKey`, `Update`, `Verify` and `WrapKey`.
     */
    keyPermissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the id of the Key Vault resource. Changing this
     * forces a new resource to be created.
     */
    keyVaultId: pulumi.Input<string>;
    /**
     * The object ID of a user, service principal or security
     * group in the Azure Active Directory tenant for the vault. The object ID must
     * be unique for the list of access policies. Changing this forces a new resource
     * to be created.
     */
    objectId: pulumi.Input<string>;
    /**
     * List of secret permissions, must be one or more from the following: `Backup`, `Delete`, `get`, `list`, `purge`, `recover`, `restore` and `set`.
     */
    secretPermissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of storage permissions, must be one or more from the following: `Backup`, `Delete`, `DeleteSAS`, `Get`, `GetSAS`, `List`, `ListSAS`, `Purge`, `Recover`, `RegenerateKey`, `Restore`, `Set`, `SetSAS` and `Update`.
     */
    storagePermissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Azure Active Directory tenant ID that should be used
     * for authenticating requests to the key vault. Changing this forces a new resource
     * to be created.
     */
    tenantId: pulumi.Input<string>;
}

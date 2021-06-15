// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Customer Managed Key for a PostgreSQL Server.
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
 *     purgeProtectionEnabled: true,
 * });
 * const exampleServer = new azure.postgresql.Server("exampleServer", {
 *     location: azurerm_resource_group.test.location,
 *     resourceGroupName: azurerm_resource_group.test.name,
 *     administratorLogin: "psqladminun",
 *     administratorLoginPassword: "H@Sh1CoR3!",
 *     skuName: "GP_Gen5_2",
 *     version: "11",
 *     storageMb: 51200,
 *     sslEnforcementEnabled: true,
 *     identity: {
 *         type: "SystemAssigned",
 *     },
 * });
 * const server = new azure.keyvault.AccessPolicy("server", {
 *     keyVaultId: exampleKeyVault.id,
 *     tenantId: current.then(current => current.tenantId),
 *     objectId: exampleServer.identity.apply(identity => identity.principalId),
 *     keyPermissions: [
 *         "get",
 *         "unwrapkey",
 *         "wrapkey",
 *     ],
 *     secretPermissions: ["get"],
 * });
 * const client = new azure.keyvault.AccessPolicy("client", {
 *     keyVaultId: exampleKeyVault.id,
 *     tenantId: current.then(current => current.tenantId),
 *     objectId: current.then(current => current.objectId),
 *     keyPermissions: [
 *         "get",
 *         "create",
 *         "delete",
 *         "list",
 *         "restore",
 *         "recover",
 *         "unwrapkey",
 *         "wrapkey",
 *         "purge",
 *         "encrypt",
 *         "decrypt",
 *         "sign",
 *         "verify",
 *     ],
 *     secretPermissions: ["get"],
 * });
 * const exampleKey = new azure.keyvault.Key("exampleKey", {
 *     keyVaultId: exampleKeyVault.id,
 *     keyType: "RSA",
 *     keySize: 2048,
 *     keyOpts: [
 *         "decrypt",
 *         "encrypt",
 *         "sign",
 *         "unwrapKey",
 *         "verify",
 *         "wrapKey",
 *     ],
 * }, {
 *     dependsOn: [
 *         client,
 *         server,
 *     ],
 * });
 * const exampleServerKey = new azure.postgresql.ServerKey("exampleServerKey", {
 *     serverId: exampleServer.id,
 *     keyVaultKeyId: exampleKey.id,
 * });
 * ```
 *
 * ## Import
 *
 * A PostgreSQL Server Key can be imported using the `resource id` of the PostgreSQL Server Key, e.g.
 *
 * ```sh
 *  $ pulumi import azure:postgresql/serverKey:ServerKey example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforPostgreSQL/servers/server1/keys/keyvaultname_key-name_keyversion
 * ```
 */
export class ServerKey extends pulumi.CustomResource {
    /**
     * Get an existing ServerKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerKeyState, opts?: pulumi.CustomResourceOptions): ServerKey {
        return new ServerKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:postgresql/serverKey:ServerKey';

    /**
     * Returns true if the given object is an instance of ServerKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerKey.__pulumiType;
    }

    /**
     * The URL to a Key Vault Key.
     */
    public readonly keyVaultKeyId!: pulumi.Output<string>;
    /**
     * The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
     */
    public readonly serverId!: pulumi.Output<string>;

    /**
     * Create a ServerKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerKeyArgs | ServerKeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServerKeyState | undefined;
            inputs["keyVaultKeyId"] = state ? state.keyVaultKeyId : undefined;
            inputs["serverId"] = state ? state.serverId : undefined;
        } else {
            const args = argsOrState as ServerKeyArgs | undefined;
            if ((!args || args.keyVaultKeyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyVaultKeyId'");
            }
            if ((!args || args.serverId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverId'");
            }
            inputs["keyVaultKeyId"] = args ? args.keyVaultKeyId : undefined;
            inputs["serverId"] = args ? args.serverId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ServerKey.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServerKey resources.
 */
export interface ServerKeyState {
    /**
     * The URL to a Key Vault Key.
     */
    keyVaultKeyId?: pulumi.Input<string>;
    /**
     * The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
     */
    serverId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServerKey resource.
 */
export interface ServerKeyArgs {
    /**
     * The URL to a Key Vault Key.
     */
    keyVaultKeyId: pulumi.Input<string>;
    /**
     * The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
     */
    serverId: pulumi.Input<string>;
}

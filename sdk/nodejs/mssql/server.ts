// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Manages a Microsoft SQL Azure Database Server.
 *
 * > **Note:** All arguments including the administrator login and password will be stored in the raw state as plain-text.
 * [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US"});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const exampleServer = new azure.mssql.Server("exampleServer", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     version: "12.0",
 *     administratorLogin: "missadministrator",
 *     administratorLoginPassword: "thisIsKat11",
 *     azureadAdministrator: {
 *         loginUsername: "AzureAD Admin",
 *         objectId: "00000000-0000-0000-0000-000000000000",
 *     },
 *     extendedAuditingPolicy: {
 *         storageEndpoint: exampleAccount.primaryBlobEndpoint,
 *         storageAccountAccessKey: exampleAccount.primaryAccessKey,
 *         storageAccountAccessKeyIsSecondary: true,
 *         retentionInDays: 6,
 *     },
 *     tags: {
 *         environment: "production",
 *     },
 * });
 * ```
 */
export class Server extends pulumi.CustomResource {
    /**
     * Get an existing Server resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerState, opts?: pulumi.CustomResourceOptions): Server {
        return new Server(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:mssql/server:Server';

    /**
     * Returns true if the given object is an instance of Server.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Server {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Server.__pulumiType;
    }

    /**
     * The administrator's login name for the new server. Changing this forces a new resource to be created.
     */
    public readonly administratorLogin!: pulumi.Output<string>;
    /**
     * The password associated with the `administratorLogin` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
     */
    public readonly administratorLoginPassword!: pulumi.Output<string>;
    /**
     * An `azureadAdministrator` block as defined below.
     */
    public readonly azureadAdministrator!: pulumi.Output<outputs.mssql.ServerAzureadAdministrator | undefined>;
    /**
     * The connection policy the server will use. Possible values are `Default`, `Proxy`, and `Redirect`. Defaults to `Default`.
     */
    public readonly connectionPolicy!: pulumi.Output<string | undefined>;
    /**
     * A `extendedAuditingPolicy` block as defined below.
     *
     * @deprecated the `extended_auditing_policy` block has been moved to `azurerm_mssql_server_extended_auditing_policy` and `azurerm_mssql_database_extended_auditing_policy`. This block will be removed in version 3.0 of the provider.
     */
    public readonly extendedAuditingPolicy!: pulumi.Output<outputs.mssql.ServerExtendedAuditingPolicy>;
    /**
     * The fully qualified domain name of the Azure SQL Server (e.g. myServerName.database.windows.net)
     */
    public /*out*/ readonly fullyQualifiedDomainName!: pulumi.Output<string>;
    /**
     * An `identity` block as defined below.
     */
    public readonly identity!: pulumi.Output<outputs.mssql.ServerIdentity | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the Microsoft SQL Server. This needs to be globally unique within Azure.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Whether or not public network access is allowed for this server. Defaults to `true`.
     */
    public readonly publicNetworkAccessEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the resource group in which to create the Microsoft SQL Server.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A list of dropped restorable database IDs on the server.
     */
    public /*out*/ readonly restorableDroppedDatabaseIds!: pulumi.Output<string[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * This servers MS SQL version. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server).
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a Server resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerArgs | ServerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServerState | undefined;
            inputs["administratorLogin"] = state ? state.administratorLogin : undefined;
            inputs["administratorLoginPassword"] = state ? state.administratorLoginPassword : undefined;
            inputs["azureadAdministrator"] = state ? state.azureadAdministrator : undefined;
            inputs["connectionPolicy"] = state ? state.connectionPolicy : undefined;
            inputs["extendedAuditingPolicy"] = state ? state.extendedAuditingPolicy : undefined;
            inputs["fullyQualifiedDomainName"] = state ? state.fullyQualifiedDomainName : undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["publicNetworkAccessEnabled"] = state ? state.publicNetworkAccessEnabled : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["restorableDroppedDatabaseIds"] = state ? state.restorableDroppedDatabaseIds : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as ServerArgs | undefined;
            if (!args || args.administratorLogin === undefined) {
                throw new Error("Missing required property 'administratorLogin'");
            }
            if (!args || args.administratorLoginPassword === undefined) {
                throw new Error("Missing required property 'administratorLoginPassword'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.version === undefined) {
                throw new Error("Missing required property 'version'");
            }
            inputs["administratorLogin"] = args ? args.administratorLogin : undefined;
            inputs["administratorLoginPassword"] = args ? args.administratorLoginPassword : undefined;
            inputs["azureadAdministrator"] = args ? args.azureadAdministrator : undefined;
            inputs["connectionPolicy"] = args ? args.connectionPolicy : undefined;
            inputs["extendedAuditingPolicy"] = args ? args.extendedAuditingPolicy : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["publicNetworkAccessEnabled"] = args ? args.publicNetworkAccessEnabled : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["fullyQualifiedDomainName"] = undefined /*out*/;
            inputs["restorableDroppedDatabaseIds"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Server.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Server resources.
 */
export interface ServerState {
    /**
     * The administrator's login name for the new server. Changing this forces a new resource to be created.
     */
    readonly administratorLogin?: pulumi.Input<string>;
    /**
     * The password associated with the `administratorLogin` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
     */
    readonly administratorLoginPassword?: pulumi.Input<string>;
    /**
     * An `azureadAdministrator` block as defined below.
     */
    readonly azureadAdministrator?: pulumi.Input<inputs.mssql.ServerAzureadAdministrator>;
    /**
     * The connection policy the server will use. Possible values are `Default`, `Proxy`, and `Redirect`. Defaults to `Default`.
     */
    readonly connectionPolicy?: pulumi.Input<string>;
    /**
     * A `extendedAuditingPolicy` block as defined below.
     *
     * @deprecated the `extended_auditing_policy` block has been moved to `azurerm_mssql_server_extended_auditing_policy` and `azurerm_mssql_database_extended_auditing_policy`. This block will be removed in version 3.0 of the provider.
     */
    readonly extendedAuditingPolicy?: pulumi.Input<inputs.mssql.ServerExtendedAuditingPolicy>;
    /**
     * The fully qualified domain name of the Azure SQL Server (e.g. myServerName.database.windows.net)
     */
    readonly fullyQualifiedDomainName?: pulumi.Input<string>;
    /**
     * An `identity` block as defined below.
     */
    readonly identity?: pulumi.Input<inputs.mssql.ServerIdentity>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the Microsoft SQL Server. This needs to be globally unique within Azure.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Whether or not public network access is allowed for this server. Defaults to `true`.
     */
    readonly publicNetworkAccessEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which to create the Microsoft SQL Server.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A list of dropped restorable database IDs on the server.
     */
    readonly restorableDroppedDatabaseIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * This servers MS SQL version. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server).
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Server resource.
 */
export interface ServerArgs {
    /**
     * The administrator's login name for the new server. Changing this forces a new resource to be created.
     */
    readonly administratorLogin: pulumi.Input<string>;
    /**
     * The password associated with the `administratorLogin` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
     */
    readonly administratorLoginPassword: pulumi.Input<string>;
    /**
     * An `azureadAdministrator` block as defined below.
     */
    readonly azureadAdministrator?: pulumi.Input<inputs.mssql.ServerAzureadAdministrator>;
    /**
     * The connection policy the server will use. Possible values are `Default`, `Proxy`, and `Redirect`. Defaults to `Default`.
     */
    readonly connectionPolicy?: pulumi.Input<string>;
    /**
     * A `extendedAuditingPolicy` block as defined below.
     *
     * @deprecated the `extended_auditing_policy` block has been moved to `azurerm_mssql_server_extended_auditing_policy` and `azurerm_mssql_database_extended_auditing_policy`. This block will be removed in version 3.0 of the provider.
     */
    readonly extendedAuditingPolicy?: pulumi.Input<inputs.mssql.ServerExtendedAuditingPolicy>;
    /**
     * An `identity` block as defined below.
     */
    readonly identity?: pulumi.Input<inputs.mssql.ServerIdentity>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the Microsoft SQL Server. This needs to be globally unique within Azure.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Whether or not public network access is allowed for this server. Defaults to `true`.
     */
    readonly publicNetworkAccessEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which to create the Microsoft SQL Server.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * This servers MS SQL version. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server).
     */
    readonly version: pulumi.Input<string>;
}

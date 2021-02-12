// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Ms Sql Server Extended Auditing Policy.
 *
 * > **NOTE:** The Server Extended Auditing Policy Can be set inline here as well as with the mssqlServerExtendedAuditingPolicy resource resource. You can only use one or the other and using both will cause a conflict.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleServer = new azure.mssql.Server("exampleServer", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     version: "12.0",
 *     administratorLogin: "missadministrator",
 *     administratorLoginPassword: "AdminPassword123!",
 * });
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const exampleServerExtendedAuditingPolicy = new azure.mssql.ServerExtendedAuditingPolicy("exampleServerExtendedAuditingPolicy", {
 *     serverId: exampleServer.id,
 *     storageEndpoint: exampleAccount.primaryBlobEndpoint,
 *     storageAccountAccessKey: exampleAccount.primaryAccessKey,
 *     storageAccountAccessKeyIsSecondary: false,
 *     retentionInDays: 6,
 * });
 * ```
 *
 * ## Import
 *
 * Ms Sql Server Extended Auditing Policys can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:mssql/serverExtendedAuditingPolicy:ServerExtendedAuditingPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlServer1/extendedAuditingSettings/default
 * ```
 */
export class ServerExtendedAuditingPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ServerExtendedAuditingPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerExtendedAuditingPolicyState, opts?: pulumi.CustomResourceOptions): ServerExtendedAuditingPolicy {
        return new ServerExtendedAuditingPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:mssql/serverExtendedAuditingPolicy:ServerExtendedAuditingPolicy';

    /**
     * Returns true if the given object is an instance of ServerExtendedAuditingPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerExtendedAuditingPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerExtendedAuditingPolicy.__pulumiType;
    }

    /**
     * The number of days to retain logs for in the storage account.
     */
    public readonly retentionInDays!: pulumi.Output<number | undefined>;
    /**
     * The ID of the sql server to set the extended auditing policy. Changing this forces a new resource to be created.
     */
    public readonly serverId!: pulumi.Output<string>;
    /**
     * The access key to use for the auditing storage account.
     */
    public readonly storageAccountAccessKey!: pulumi.Output<string | undefined>;
    /**
     * Is `storageAccountAccessKey` value the storage's secondary key?
     */
    public readonly storageAccountAccessKeyIsSecondary!: pulumi.Output<boolean | undefined>;
    /**
     * The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
     */
    public readonly storageEndpoint!: pulumi.Output<string>;

    /**
     * Create a ServerExtendedAuditingPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerExtendedAuditingPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerExtendedAuditingPolicyArgs | ServerExtendedAuditingPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServerExtendedAuditingPolicyState | undefined;
            inputs["retentionInDays"] = state ? state.retentionInDays : undefined;
            inputs["serverId"] = state ? state.serverId : undefined;
            inputs["storageAccountAccessKey"] = state ? state.storageAccountAccessKey : undefined;
            inputs["storageAccountAccessKeyIsSecondary"] = state ? state.storageAccountAccessKeyIsSecondary : undefined;
            inputs["storageEndpoint"] = state ? state.storageEndpoint : undefined;
        } else {
            const args = argsOrState as ServerExtendedAuditingPolicyArgs | undefined;
            if ((!args || args.serverId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverId'");
            }
            if ((!args || args.storageEndpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageEndpoint'");
            }
            inputs["retentionInDays"] = args ? args.retentionInDays : undefined;
            inputs["serverId"] = args ? args.serverId : undefined;
            inputs["storageAccountAccessKey"] = args ? args.storageAccountAccessKey : undefined;
            inputs["storageAccountAccessKeyIsSecondary"] = args ? args.storageAccountAccessKeyIsSecondary : undefined;
            inputs["storageEndpoint"] = args ? args.storageEndpoint : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ServerExtendedAuditingPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServerExtendedAuditingPolicy resources.
 */
export interface ServerExtendedAuditingPolicyState {
    /**
     * The number of days to retain logs for in the storage account.
     */
    readonly retentionInDays?: pulumi.Input<number>;
    /**
     * The ID of the sql server to set the extended auditing policy. Changing this forces a new resource to be created.
     */
    readonly serverId?: pulumi.Input<string>;
    /**
     * The access key to use for the auditing storage account.
     */
    readonly storageAccountAccessKey?: pulumi.Input<string>;
    /**
     * Is `storageAccountAccessKey` value the storage's secondary key?
     */
    readonly storageAccountAccessKeyIsSecondary?: pulumi.Input<boolean>;
    /**
     * The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
     */
    readonly storageEndpoint?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServerExtendedAuditingPolicy resource.
 */
export interface ServerExtendedAuditingPolicyArgs {
    /**
     * The number of days to retain logs for in the storage account.
     */
    readonly retentionInDays?: pulumi.Input<number>;
    /**
     * The ID of the sql server to set the extended auditing policy. Changing this forces a new resource to be created.
     */
    readonly serverId: pulumi.Input<string>;
    /**
     * The access key to use for the auditing storage account.
     */
    readonly storageAccountAccessKey?: pulumi.Input<string>;
    /**
     * Is `storageAccountAccessKey` value the storage's secondary key?
     */
    readonly storageAccountAccessKeyIsSecondary?: pulumi.Input<boolean>;
    /**
     * The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
     */
    readonly storageEndpoint: pulumi.Input<string>;
}

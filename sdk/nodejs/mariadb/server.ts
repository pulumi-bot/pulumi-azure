// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a MariaDB Server.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleServer = new azure.mariadb.Server("exampleServer", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     administratorLogin: "mariadbadmin",
 *     administratorLoginPassword: "H@Sh1CoR3!",
 *     skuName: "B_Gen5_2",
 *     storageMb: 5120,
 *     version: "10.2",
 *     autoGrowEnabled: true,
 *     backupRetentionDays: 7,
 *     geoRedundantBackupEnabled: false,
 *     publicNetworkAccessEnabled: false,
 *     sslEnforcementEnabled: true,
 * });
 * ```
 *
 * ## Import
 *
 * MariaDB Server's can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:mariadb/server:Server server1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforMariaDB/servers/server1
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
    public static readonly __pulumiType = 'azure:mariadb/server:Server';

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
     * The Administrator Login for the MariaDB Server. Changing this forces a new resource to be created.
     */
    public readonly administratorLogin!: pulumi.Output<string>;
    /**
     * The Password associated with the `administratorLogin` for the MariaDB Server.
     */
    public readonly administratorLoginPassword!: pulumi.Output<string | undefined>;
    /**
     * Enable/Disable auto-growing of the storage. Storage auto-grow prevents your server from running out of storage and becoming read-only. If storage auto grow is enabled, the storage automatically grows without impacting the workload. The default value if not explicitly specified is `true`.
     */
    public readonly autoGrowEnabled!: pulumi.Output<boolean>;
    /**
     * Backup retention days for the server, supported values are between `7` and `35` days.
     */
    public readonly backupRetentionDays!: pulumi.Output<number>;
    /**
     * The creation mode. Can be used to restore or replicate existing servers. Possible values are `Default`, `Replica`, `GeoRestore`, and `PointInTimeRestore`. Defaults to `Default`.
     */
    public readonly createMode!: pulumi.Output<string | undefined>;
    /**
     * For creation modes other than `Default`, the source server ID to use.
     */
    public readonly creationSourceServerId!: pulumi.Output<string | undefined>;
    /**
     * The FQDN of the MariaDB Server.
     */
    public /*out*/ readonly fqdn!: pulumi.Output<string>;
    /**
     * Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not supported for the Basic tier.
     */
    public readonly geoRedundantBackupEnabled!: pulumi.Output<boolean>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Whether or not public network access is allowed for this server. Defaults to `true`.
     */
    public readonly publicNetworkAccessEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the resource group in which to create the MariaDB Server. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * When `createMode` is `PointInTimeRestore`, specifies the point in time to restore from `creationSourceServerId`.
     */
    public readonly restorePointInTime!: pulumi.Output<string | undefined>;
    /**
     * Specifies the SKU Name for this MariaDB Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mariadb/servers/create#sku).
     */
    public readonly skuName!: pulumi.Output<string>;
    /**
     * @deprecated this has been moved to the boolean attribute `ssl_enforcement_enabled` and will be removed in version 3.0 of the provider.
     */
    public readonly sslEnforcement!: pulumi.Output<string>;
    /**
     * Specifies if SSL should be enforced on connections. Possible values are `true` and `false`.
     */
    public readonly sslEnforcementEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Max storage allowed for a server. Possible values are between `5120` MB (5GB) and `1024000`MB (1TB) for the Basic SKU and between `5120` MB (5GB) and `4096000` MB (4TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mariadb/servers/create#storageprofile).
     */
    public readonly storageMb!: pulumi.Output<number>;
    /**
     * @deprecated all storage_profile properties have been moved to the top level. This block will be removed in version 3.0 of the provider.
     */
    public readonly storageProfile!: pulumi.Output<outputs.mariadb.ServerStorageProfile>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the version of MariaDB to use. Possible values are `10.2` and `10.3`. Changing this forces a new resource to be created.
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
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServerState | undefined;
            inputs["administratorLogin"] = state ? state.administratorLogin : undefined;
            inputs["administratorLoginPassword"] = state ? state.administratorLoginPassword : undefined;
            inputs["autoGrowEnabled"] = state ? state.autoGrowEnabled : undefined;
            inputs["backupRetentionDays"] = state ? state.backupRetentionDays : undefined;
            inputs["createMode"] = state ? state.createMode : undefined;
            inputs["creationSourceServerId"] = state ? state.creationSourceServerId : undefined;
            inputs["fqdn"] = state ? state.fqdn : undefined;
            inputs["geoRedundantBackupEnabled"] = state ? state.geoRedundantBackupEnabled : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["publicNetworkAccessEnabled"] = state ? state.publicNetworkAccessEnabled : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["restorePointInTime"] = state ? state.restorePointInTime : undefined;
            inputs["skuName"] = state ? state.skuName : undefined;
            inputs["sslEnforcement"] = state ? state.sslEnforcement : undefined;
            inputs["sslEnforcementEnabled"] = state ? state.sslEnforcementEnabled : undefined;
            inputs["storageMb"] = state ? state.storageMb : undefined;
            inputs["storageProfile"] = state ? state.storageProfile : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as ServerArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.skuName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'skuName'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            inputs["administratorLogin"] = args ? args.administratorLogin : undefined;
            inputs["administratorLoginPassword"] = args ? args.administratorLoginPassword : undefined;
            inputs["autoGrowEnabled"] = args ? args.autoGrowEnabled : undefined;
            inputs["backupRetentionDays"] = args ? args.backupRetentionDays : undefined;
            inputs["createMode"] = args ? args.createMode : undefined;
            inputs["creationSourceServerId"] = args ? args.creationSourceServerId : undefined;
            inputs["geoRedundantBackupEnabled"] = args ? args.geoRedundantBackupEnabled : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["publicNetworkAccessEnabled"] = args ? args.publicNetworkAccessEnabled : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["restorePointInTime"] = args ? args.restorePointInTime : undefined;
            inputs["skuName"] = args ? args.skuName : undefined;
            inputs["sslEnforcement"] = args ? args.sslEnforcement : undefined;
            inputs["sslEnforcementEnabled"] = args ? args.sslEnforcementEnabled : undefined;
            inputs["storageMb"] = args ? args.storageMb : undefined;
            inputs["storageProfile"] = args ? args.storageProfile : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["fqdn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Server.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Server resources.
 */
export interface ServerState {
    /**
     * The Administrator Login for the MariaDB Server. Changing this forces a new resource to be created.
     */
    administratorLogin?: pulumi.Input<string>;
    /**
     * The Password associated with the `administratorLogin` for the MariaDB Server.
     */
    administratorLoginPassword?: pulumi.Input<string>;
    /**
     * Enable/Disable auto-growing of the storage. Storage auto-grow prevents your server from running out of storage and becoming read-only. If storage auto grow is enabled, the storage automatically grows without impacting the workload. The default value if not explicitly specified is `true`.
     */
    autoGrowEnabled?: pulumi.Input<boolean>;
    /**
     * Backup retention days for the server, supported values are between `7` and `35` days.
     */
    backupRetentionDays?: pulumi.Input<number>;
    /**
     * The creation mode. Can be used to restore or replicate existing servers. Possible values are `Default`, `Replica`, `GeoRestore`, and `PointInTimeRestore`. Defaults to `Default`.
     */
    createMode?: pulumi.Input<string>;
    /**
     * For creation modes other than `Default`, the source server ID to use.
     */
    creationSourceServerId?: pulumi.Input<string>;
    /**
     * The FQDN of the MariaDB Server.
     */
    fqdn?: pulumi.Input<string>;
    /**
     * Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not supported for the Basic tier.
     */
    geoRedundantBackupEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether or not public network access is allowed for this server. Defaults to `true`.
     */
    publicNetworkAccessEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which to create the MariaDB Server. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * When `createMode` is `PointInTimeRestore`, specifies the point in time to restore from `creationSourceServerId`.
     */
    restorePointInTime?: pulumi.Input<string>;
    /**
     * Specifies the SKU Name for this MariaDB Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mariadb/servers/create#sku).
     */
    skuName?: pulumi.Input<string>;
    /**
     * @deprecated this has been moved to the boolean attribute `ssl_enforcement_enabled` and will be removed in version 3.0 of the provider.
     */
    sslEnforcement?: pulumi.Input<string>;
    /**
     * Specifies if SSL should be enforced on connections. Possible values are `true` and `false`.
     */
    sslEnforcementEnabled?: pulumi.Input<boolean>;
    /**
     * Max storage allowed for a server. Possible values are between `5120` MB (5GB) and `1024000`MB (1TB) for the Basic SKU and between `5120` MB (5GB) and `4096000` MB (4TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mariadb/servers/create#storageprofile).
     */
    storageMb?: pulumi.Input<number>;
    /**
     * @deprecated all storage_profile properties have been moved to the top level. This block will be removed in version 3.0 of the provider.
     */
    storageProfile?: pulumi.Input<inputs.mariadb.ServerStorageProfile>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the version of MariaDB to use. Possible values are `10.2` and `10.3`. Changing this forces a new resource to be created.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Server resource.
 */
export interface ServerArgs {
    /**
     * The Administrator Login for the MariaDB Server. Changing this forces a new resource to be created.
     */
    administratorLogin?: pulumi.Input<string>;
    /**
     * The Password associated with the `administratorLogin` for the MariaDB Server.
     */
    administratorLoginPassword?: pulumi.Input<string>;
    /**
     * Enable/Disable auto-growing of the storage. Storage auto-grow prevents your server from running out of storage and becoming read-only. If storage auto grow is enabled, the storage automatically grows without impacting the workload. The default value if not explicitly specified is `true`.
     */
    autoGrowEnabled?: pulumi.Input<boolean>;
    /**
     * Backup retention days for the server, supported values are between `7` and `35` days.
     */
    backupRetentionDays?: pulumi.Input<number>;
    /**
     * The creation mode. Can be used to restore or replicate existing servers. Possible values are `Default`, `Replica`, `GeoRestore`, and `PointInTimeRestore`. Defaults to `Default`.
     */
    createMode?: pulumi.Input<string>;
    /**
     * For creation modes other than `Default`, the source server ID to use.
     */
    creationSourceServerId?: pulumi.Input<string>;
    /**
     * Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not supported for the Basic tier.
     */
    geoRedundantBackupEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether or not public network access is allowed for this server. Defaults to `true`.
     */
    publicNetworkAccessEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which to create the MariaDB Server. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * When `createMode` is `PointInTimeRestore`, specifies the point in time to restore from `creationSourceServerId`.
     */
    restorePointInTime?: pulumi.Input<string>;
    /**
     * Specifies the SKU Name for this MariaDB Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mariadb/servers/create#sku).
     */
    skuName: pulumi.Input<string>;
    /**
     * @deprecated this has been moved to the boolean attribute `ssl_enforcement_enabled` and will be removed in version 3.0 of the provider.
     */
    sslEnforcement?: pulumi.Input<string>;
    /**
     * Specifies if SSL should be enforced on connections. Possible values are `true` and `false`.
     */
    sslEnforcementEnabled?: pulumi.Input<boolean>;
    /**
     * Max storage allowed for a server. Possible values are between `5120` MB (5GB) and `1024000`MB (1TB) for the Basic SKU and between `5120` MB (5GB) and `4096000` MB (4TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mariadb/servers/create#storageprofile).
     */
    storageMb?: pulumi.Input<number>;
    /**
     * @deprecated all storage_profile properties have been moved to the top level. This block will be removed in version 3.0 of the provider.
     */
    storageProfile?: pulumi.Input<inputs.mariadb.ServerStorageProfile>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the version of MariaDB to use. Possible values are `10.2` and `10.3`. Changing this forces a new resource to be created.
     */
    version: pulumi.Input<string>;
}

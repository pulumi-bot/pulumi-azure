// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Allows you to manage an Azure SQL Elastic Pool via the `v3.0` API which allows for `vCore` and `DTU` based configurations.
 */
export class ElasticPool extends pulumi.CustomResource {
    /**
     * Get an existing ElasticPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ElasticPoolState, opts?: pulumi.CustomResourceOptions): ElasticPool {
        return new ElasticPool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:mssql/elasticPool:ElasticPool';

    /**
     * Returns true if the given object is an instance of ElasticPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ElasticPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ElasticPool.__pulumiType;
    }

    /**
     * Specifies the license type applied to this database. Possible values are `LicenseIncluded` and `BasePrice`.
     */
    public readonly licenseType!: pulumi.Output<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The max data size of the elastic pool in bytes. Conflicts with `maxSizeGb`.
     */
    public readonly maxSizeBytes!: pulumi.Output<number>;
    /**
     * The max data size of the elastic pool in gigabytes. Conflicts with `maxSizeBytes`.
     */
    public readonly maxSizeGb!: pulumi.Output<number>;
    /**
     * The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A `perDatabaseSettings` block as defined below.
     */
    public readonly perDatabaseSettings!: pulumi.Output<outputs.mssql.ElasticPoolPerDatabaseSettings>;
    /**
     * The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
     */
    public readonly serverName!: pulumi.Output<string>;
    /**
     * A `sku` block as defined below.
     */
    public readonly sku!: pulumi.Output<outputs.mssql.ElasticPoolSku>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Whether or not this elastic pool is zone redundant. `tier` needs to be `Premium` for `DTU` based  or `BusinessCritical` for `vCore` based `sku`. Defaults to `false`.
     */
    public readonly zoneRedundant!: pulumi.Output<boolean | undefined>;

    /**
     * Create a ElasticPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ElasticPoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ElasticPoolArgs | ElasticPoolState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ElasticPoolState | undefined;
            inputs["licenseType"] = state ? state.licenseType : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["maxSizeBytes"] = state ? state.maxSizeBytes : undefined;
            inputs["maxSizeGb"] = state ? state.maxSizeGb : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["perDatabaseSettings"] = state ? state.perDatabaseSettings : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["serverName"] = state ? state.serverName : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["zoneRedundant"] = state ? state.zoneRedundant : undefined;
        } else {
            const args = argsOrState as ElasticPoolArgs | undefined;
            if (!args || args.perDatabaseSettings === undefined) {
                throw new Error("Missing required property 'perDatabaseSettings'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serverName === undefined) {
                throw new Error("Missing required property 'serverName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["licenseType"] = args ? args.licenseType : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["maxSizeBytes"] = args ? args.maxSizeBytes : undefined;
            inputs["maxSizeGb"] = args ? args.maxSizeGb : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["perDatabaseSettings"] = args ? args.perDatabaseSettings : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["zoneRedundant"] = args ? args.zoneRedundant : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ElasticPool.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ElasticPool resources.
 */
export interface ElasticPoolState {
    /**
     * Specifies the license type applied to this database. Possible values are `LicenseIncluded` and `BasePrice`.
     */
    readonly licenseType?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The max data size of the elastic pool in bytes. Conflicts with `maxSizeGb`.
     */
    readonly maxSizeBytes?: pulumi.Input<number>;
    /**
     * The max data size of the elastic pool in gigabytes. Conflicts with `maxSizeBytes`.
     */
    readonly maxSizeGb?: pulumi.Input<number>;
    /**
     * The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `perDatabaseSettings` block as defined below.
     */
    readonly perDatabaseSettings?: pulumi.Input<inputs.mssql.ElasticPoolPerDatabaseSettings>;
    /**
     * The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
     */
    readonly serverName?: pulumi.Input<string>;
    /**
     * A `sku` block as defined below.
     */
    readonly sku?: pulumi.Input<inputs.mssql.ElasticPoolSku>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether or not this elastic pool is zone redundant. `tier` needs to be `Premium` for `DTU` based  or `BusinessCritical` for `vCore` based `sku`. Defaults to `false`.
     */
    readonly zoneRedundant?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ElasticPool resource.
 */
export interface ElasticPoolArgs {
    /**
     * Specifies the license type applied to this database. Possible values are `LicenseIncluded` and `BasePrice`.
     */
    readonly licenseType?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The max data size of the elastic pool in bytes. Conflicts with `maxSizeGb`.
     */
    readonly maxSizeBytes?: pulumi.Input<number>;
    /**
     * The max data size of the elastic pool in gigabytes. Conflicts with `maxSizeBytes`.
     */
    readonly maxSizeGb?: pulumi.Input<number>;
    /**
     * The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `perDatabaseSettings` block as defined below.
     */
    readonly perDatabaseSettings: pulumi.Input<inputs.mssql.ElasticPoolPerDatabaseSettings>;
    /**
     * The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
     */
    readonly serverName: pulumi.Input<string>;
    /**
     * A `sku` block as defined below.
     */
    readonly sku: pulumi.Input<inputs.mssql.ElasticPoolSku>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether or not this elastic pool is zone redundant. `tier` needs to be `Premium` for `DTU` based  or `BusinessCritical` for `vCore` based `sku`. Defaults to `false`.
     */
    readonly zoneRedundant?: pulumi.Input<boolean>;
}

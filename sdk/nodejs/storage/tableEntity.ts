// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Entity within a Table in an Azure Storage Account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "westus"});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const exampleTable = new azure.storage.Table("exampleTable", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     storageAccountName: exampleAccount.name,
 * });
 * const exampleTableEntity = new azure.storage.TableEntity("exampleTableEntity", {
 *     storageAccountName: exampleAccount.name,
 *     tableName: exampleTable.name,
 *     partitionKey: "examplepartition",
 *     rowKey: "examplerow",
 *     entity: {
 *         example: "example",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Entities within a Table in an Azure Storage Account can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:storage/tableEntity:TableEntity entity1 https://example.table.core.windows.net/table1(PartitionKey='samplepartition',RowKey='samplerow')
 * ```
 */
export class TableEntity extends pulumi.CustomResource {
    /**
     * Get an existing TableEntity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TableEntityState, opts?: pulumi.CustomResourceOptions): TableEntity {
        return new TableEntity(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:storage/tableEntity:TableEntity';

    /**
     * Returns true if the given object is an instance of TableEntity.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TableEntity {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TableEntity.__pulumiType;
    }

    /**
     * A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
     */
    public readonly entity!: pulumi.Output<{[key: string]: string}>;
    /**
     * The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
     */
    public readonly partitionKey!: pulumi.Output<string>;
    /**
     * The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
     */
    public readonly rowKey!: pulumi.Output<string>;
    /**
     * Specifies the storage account in which to create the storage table entity.
     * Changing this forces a new resource to be created.
     */
    public readonly storageAccountName!: pulumi.Output<string>;
    /**
     * The name of the storage table in which to create the storage table entity.
     * Changing this forces a new resource to be created.
     */
    public readonly tableName!: pulumi.Output<string>;

    /**
     * Create a TableEntity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TableEntityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TableEntityArgs | TableEntityState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TableEntityState | undefined;
            inputs["entity"] = state ? state.entity : undefined;
            inputs["partitionKey"] = state ? state.partitionKey : undefined;
            inputs["rowKey"] = state ? state.rowKey : undefined;
            inputs["storageAccountName"] = state ? state.storageAccountName : undefined;
            inputs["tableName"] = state ? state.tableName : undefined;
        } else {
            const args = argsOrState as TableEntityArgs | undefined;
            if ((!args || args.entity === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'entity'");
            }
            if ((!args || args.partitionKey === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'partitionKey'");
            }
            if ((!args || args.rowKey === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'rowKey'");
            }
            if ((!args || args.storageAccountName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'storageAccountName'");
            }
            if ((!args || args.tableName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'tableName'");
            }
            inputs["entity"] = args ? args.entity : undefined;
            inputs["partitionKey"] = args ? args.partitionKey : undefined;
            inputs["rowKey"] = args ? args.rowKey : undefined;
            inputs["storageAccountName"] = args ? args.storageAccountName : undefined;
            inputs["tableName"] = args ? args.tableName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TableEntity.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TableEntity resources.
 */
export interface TableEntityState {
    /**
     * A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
     */
    readonly entity?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
     */
    readonly partitionKey?: pulumi.Input<string>;
    /**
     * The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
     */
    readonly rowKey?: pulumi.Input<string>;
    /**
     * Specifies the storage account in which to create the storage table entity.
     * Changing this forces a new resource to be created.
     */
    readonly storageAccountName?: pulumi.Input<string>;
    /**
     * The name of the storage table in which to create the storage table entity.
     * Changing this forces a new resource to be created.
     */
    readonly tableName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TableEntity resource.
 */
export interface TableEntityArgs {
    /**
     * A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
     */
    readonly entity: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
     */
    readonly partitionKey: pulumi.Input<string>;
    /**
     * The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
     */
    readonly rowKey: pulumi.Input<string>;
    /**
     * Specifies the storage account in which to create the storage table entity.
     * Changing this forces a new resource to be created.
     */
    readonly storageAccountName: pulumi.Input<string>;
    /**
     * The name of the storage table in which to create the storage table entity.
     * Changing this forces a new resource to be created.
     */
    readonly tableName: pulumi.Input<string>;
}

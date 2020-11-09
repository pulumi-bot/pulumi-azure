// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Gremlin Graph within a Cosmos DB Account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleAccount = azure.cosmosdb.getAccount({
 *     name: "tfex-cosmosdb-account",
 *     resourceGroupName: "tfex-cosmosdb-account-rg",
 * });
 * const exampleGremlinDatabase = new azure.cosmosdb.GremlinDatabase("exampleGremlinDatabase", {
 *     resourceGroupName: exampleAccount.then(exampleAccount => exampleAccount.resourceGroupName),
 *     accountName: exampleAccount.then(exampleAccount => exampleAccount.name),
 * });
 * const exampleGremlinGraph = new azure.cosmosdb.GremlinGraph("exampleGremlinGraph", {
 *     resourceGroupName: azurerm_cosmosdb_account.example.resource_group_name,
 *     accountName: azurerm_cosmosdb_account.example.name,
 *     databaseName: exampleGremlinDatabase.name,
 *     partitionKeyPath: "/Example",
 *     throughput: 400,
 *     indexPolicies: [{
 *         automatic: true,
 *         indexingMode: "Consistent",
 *         includedPaths: ["/*"],
 *         excludedPaths: ["/\"_etag\"/?"],
 *     }],
 *     conflictResolutionPolicies: [{
 *         mode: "LastWriterWins",
 *         conflictResolutionPath: "/_ts",
 *     }],
 *     uniqueKeys: [{
 *         paths: [
 *             "/definition/id1",
 *             "/definition/id2",
 *         ],
 *     }],
 * });
 * ```
 *
 * > **NOTE:** The CosmosDB Account needs to have the `EnableGremlin` capability enabled to use this resource - which can be done by adding this to the `capabilities` list within the `azure.cosmosdb.Account` resource.
 */
export class GremlinGraph extends pulumi.CustomResource {
    /**
     * Get an existing GremlinGraph resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GremlinGraphState, opts?: pulumi.CustomResourceOptions): GremlinGraph {
        return new GremlinGraph(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:cosmosdb/gremlinGraph:GremlinGraph';

    /**
     * Returns true if the given object is an instance of GremlinGraph.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GremlinGraph {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GremlinGraph.__pulumiType;
    }

    /**
     * The name of the CosmosDB Account to create the Gremlin Graph within. Changing this forces a new resource to be created.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply. Requires `partitionKeyPath` to be set.
     */
    public readonly autoscaleSettings!: pulumi.Output<outputs.cosmosdb.GremlinGraphAutoscaleSettings | undefined>;
    /**
     * The conflict resolution policy for the graph. One or more `conflictResolutionPolicy` blocks as defined below. Changing this forces a new resource to be created.
     */
    public readonly conflictResolutionPolicies!: pulumi.Output<outputs.cosmosdb.GremlinGraphConflictResolutionPolicy[]>;
    /**
     * The name of the Cosmos DB Graph Database in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
     */
    public readonly databaseName!: pulumi.Output<string>;
    /**
     * The configuration of the indexing policy. One or more `indexPolicy` blocks as defined below. Changing this forces a new resource to be created.
     */
    public readonly indexPolicies!: pulumi.Output<outputs.cosmosdb.GremlinGraphIndexPolicy[]>;
    /**
     * Specifies the name of the Cosmos DB Gremlin Graph. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Define a partition key. Changing this forces a new resource to be created.
     */
    public readonly partitionKeyPath!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource group in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The throughput of the Gremlin graph (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply.
     */
    public readonly throughput!: pulumi.Output<number>;
    /**
     * One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
     */
    public readonly uniqueKeys!: pulumi.Output<outputs.cosmosdb.GremlinGraphUniqueKey[] | undefined>;

    /**
     * Create a GremlinGraph resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GremlinGraphArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GremlinGraphArgs | GremlinGraphState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GremlinGraphState | undefined;
            inputs["accountName"] = state ? state.accountName : undefined;
            inputs["autoscaleSettings"] = state ? state.autoscaleSettings : undefined;
            inputs["conflictResolutionPolicies"] = state ? state.conflictResolutionPolicies : undefined;
            inputs["databaseName"] = state ? state.databaseName : undefined;
            inputs["indexPolicies"] = state ? state.indexPolicies : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["partitionKeyPath"] = state ? state.partitionKeyPath : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["throughput"] = state ? state.throughput : undefined;
            inputs["uniqueKeys"] = state ? state.uniqueKeys : undefined;
        } else {
            const args = argsOrState as GremlinGraphArgs | undefined;
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.conflictResolutionPolicies === undefined) {
                throw new Error("Missing required property 'conflictResolutionPolicies'");
            }
            if (!args || args.databaseName === undefined) {
                throw new Error("Missing required property 'databaseName'");
            }
            if (!args || args.indexPolicies === undefined) {
                throw new Error("Missing required property 'indexPolicies'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["autoscaleSettings"] = args ? args.autoscaleSettings : undefined;
            inputs["conflictResolutionPolicies"] = args ? args.conflictResolutionPolicies : undefined;
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["indexPolicies"] = args ? args.indexPolicies : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["partitionKeyPath"] = args ? args.partitionKeyPath : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["throughput"] = args ? args.throughput : undefined;
            inputs["uniqueKeys"] = args ? args.uniqueKeys : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GremlinGraph.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GremlinGraph resources.
 */
export interface GremlinGraphState {
    /**
     * The name of the CosmosDB Account to create the Gremlin Graph within. Changing this forces a new resource to be created.
     */
    readonly accountName?: pulumi.Input<string>;
    /**
     * An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply. Requires `partitionKeyPath` to be set.
     */
    readonly autoscaleSettings?: pulumi.Input<inputs.cosmosdb.GremlinGraphAutoscaleSettings>;
    /**
     * The conflict resolution policy for the graph. One or more `conflictResolutionPolicy` blocks as defined below. Changing this forces a new resource to be created.
     */
    readonly conflictResolutionPolicies?: pulumi.Input<pulumi.Input<inputs.cosmosdb.GremlinGraphConflictResolutionPolicy>[]>;
    /**
     * The name of the Cosmos DB Graph Database in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
     */
    readonly databaseName?: pulumi.Input<string>;
    /**
     * The configuration of the indexing policy. One or more `indexPolicy` blocks as defined below. Changing this forces a new resource to be created.
     */
    readonly indexPolicies?: pulumi.Input<pulumi.Input<inputs.cosmosdb.GremlinGraphIndexPolicy>[]>;
    /**
     * Specifies the name of the Cosmos DB Gremlin Graph. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Define a partition key. Changing this forces a new resource to be created.
     */
    readonly partitionKeyPath?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The throughput of the Gremlin graph (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply.
     */
    readonly throughput?: pulumi.Input<number>;
    /**
     * One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
     */
    readonly uniqueKeys?: pulumi.Input<pulumi.Input<inputs.cosmosdb.GremlinGraphUniqueKey>[]>;
}

/**
 * The set of arguments for constructing a GremlinGraph resource.
 */
export interface GremlinGraphArgs {
    /**
     * The name of the CosmosDB Account to create the Gremlin Graph within. Changing this forces a new resource to be created.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply. Requires `partitionKeyPath` to be set.
     */
    readonly autoscaleSettings?: pulumi.Input<inputs.cosmosdb.GremlinGraphAutoscaleSettings>;
    /**
     * The conflict resolution policy for the graph. One or more `conflictResolutionPolicy` blocks as defined below. Changing this forces a new resource to be created.
     */
    readonly conflictResolutionPolicies: pulumi.Input<pulumi.Input<inputs.cosmosdb.GremlinGraphConflictResolutionPolicy>[]>;
    /**
     * The name of the Cosmos DB Graph Database in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
     */
    readonly databaseName: pulumi.Input<string>;
    /**
     * The configuration of the indexing policy. One or more `indexPolicy` blocks as defined below. Changing this forces a new resource to be created.
     */
    readonly indexPolicies: pulumi.Input<pulumi.Input<inputs.cosmosdb.GremlinGraphIndexPolicy>[]>;
    /**
     * Specifies the name of the Cosmos DB Gremlin Graph. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Define a partition key. Changing this forces a new resource to be created.
     */
    readonly partitionKeyPath?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The throughput of the Gremlin graph (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply.
     */
    readonly throughput?: pulumi.Input<number>;
    /**
     * One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
     */
    readonly uniqueKeys?: pulumi.Input<pulumi.Input<inputs.cosmosdb.GremlinGraphUniqueKey>[]>;
}

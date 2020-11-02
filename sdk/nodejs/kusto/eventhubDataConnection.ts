// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Kusto (also known as Azure Data Explorer) EventHub Data Connection
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const rg = new azure.core.ResourceGroup("rg", {location: "East US"});
 * const cluster = new azure.kusto.Cluster("cluster", {
 *     location: rg.location,
 *     resourceGroupName: rg.name,
 *     sku: {
 *         name: "Standard_D13_v2",
 *         capacity: 2,
 *     },
 * });
 * const database = new azure.kusto.Database("database", {
 *     resourceGroupName: rg.name,
 *     location: rg.location,
 *     clusterName: cluster.name,
 *     hotCachePeriod: "P7D",
 *     softDeletePeriod: "P31D",
 * });
 * const eventhubNs = new azure.eventhub.EventHubNamespace("eventhubNs", {
 *     location: rg.location,
 *     resourceGroupName: rg.name,
 *     sku: "Standard",
 * });
 * const eventhub = new azure.eventhub.EventHub("eventhub", {
 *     namespaceName: eventhubNs.name,
 *     resourceGroupName: rg.name,
 *     partitionCount: 1,
 *     messageRetention: 1,
 * });
 * const consumerGroup = new azure.eventhub.ConsumerGroup("consumerGroup", {
 *     namespaceName: eventhubNs.name,
 *     eventhubName: eventhub.name,
 *     resourceGroupName: rg.name,
 * });
 * const eventhubConnection = new azure.kusto.EventhubDataConnection("eventhubConnection", {
 *     resourceGroupName: rg.name,
 *     location: rg.location,
 *     clusterName: cluster.name,
 *     databaseName: database.name,
 *     eventhubId: azurerm_eventhub.evenhub.id,
 *     consumerGroup: consumerGroup.name,
 *     tableName: "my-table",
 *     mappingRuleName: "my-table-mapping",
 *     dataFormat: "JSON",
 * });
 * //(Optional)
 * ```
 *
 * ## Import
 *
 * Kusto EventHub Data Connections can be imported using the `resource id`, e.g.
 */
export class EventhubDataConnection extends pulumi.CustomResource {
    /**
     * Get an existing EventhubDataConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventhubDataConnectionState, opts?: pulumi.CustomResourceOptions): EventhubDataConnection {
        return new EventhubDataConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:kusto/eventhubDataConnection:EventhubDataConnection';

    /**
     * Returns true if the given object is an instance of EventhubDataConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventhubDataConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventhubDataConnection.__pulumiType;
    }

    /**
     * Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
     */
    public readonly clusterName!: pulumi.Output<string>;
    /**
     * Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
     */
    public readonly consumerGroup!: pulumi.Output<string>;
    /**
     * Specifies the data format of the EventHub messages. Allowed values: `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
     */
    public readonly dataFormat!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
     */
    public readonly databaseName!: pulumi.Output<string>;
    /**
     * Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
     */
    public readonly eventhubId!: pulumi.Output<string>;
    /**
     * The location where the Kusto Database should be created. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
     */
    public readonly mappingRuleName!: pulumi.Output<string | undefined>;
    /**
     * The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Specifies the target table name used for the message ingestion. Table must exist before resource is created.
     */
    public readonly tableName!: pulumi.Output<string | undefined>;

    /**
     * Create a EventhubDataConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventhubDataConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventhubDataConnectionArgs | EventhubDataConnectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EventhubDataConnectionState | undefined;
            inputs["clusterName"] = state ? state.clusterName : undefined;
            inputs["consumerGroup"] = state ? state.consumerGroup : undefined;
            inputs["dataFormat"] = state ? state.dataFormat : undefined;
            inputs["databaseName"] = state ? state.databaseName : undefined;
            inputs["eventhubId"] = state ? state.eventhubId : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["mappingRuleName"] = state ? state.mappingRuleName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tableName"] = state ? state.tableName : undefined;
        } else {
            const args = argsOrState as EventhubDataConnectionArgs | undefined;
            if (!args || args.clusterName === undefined) {
                throw new Error("Missing required property 'clusterName'");
            }
            if (!args || args.consumerGroup === undefined) {
                throw new Error("Missing required property 'consumerGroup'");
            }
            if (!args || args.databaseName === undefined) {
                throw new Error("Missing required property 'databaseName'");
            }
            if (!args || args.eventhubId === undefined) {
                throw new Error("Missing required property 'eventhubId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["consumerGroup"] = args ? args.consumerGroup : undefined;
            inputs["dataFormat"] = args ? args.dataFormat : undefined;
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["eventhubId"] = args ? args.eventhubId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["mappingRuleName"] = args ? args.mappingRuleName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tableName"] = args ? args.tableName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EventhubDataConnection.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventhubDataConnection resources.
 */
export interface EventhubDataConnectionState {
    /**
     * Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
     */
    readonly clusterName?: pulumi.Input<string>;
    /**
     * Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
     */
    readonly consumerGroup?: pulumi.Input<string>;
    /**
     * Specifies the data format of the EventHub messages. Allowed values: `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
     */
    readonly dataFormat?: pulumi.Input<string>;
    /**
     * Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
     */
    readonly databaseName?: pulumi.Input<string>;
    /**
     * Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
     */
    readonly eventhubId?: pulumi.Input<string>;
    /**
     * The location where the Kusto Database should be created. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
     */
    readonly mappingRuleName?: pulumi.Input<string>;
    /**
     * The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies the target table name used for the message ingestion. Table must exist before resource is created.
     */
    readonly tableName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EventhubDataConnection resource.
 */
export interface EventhubDataConnectionArgs {
    /**
     * Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
     */
    readonly clusterName: pulumi.Input<string>;
    /**
     * Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
     */
    readonly consumerGroup: pulumi.Input<string>;
    /**
     * Specifies the data format of the EventHub messages. Allowed values: `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
     */
    readonly dataFormat?: pulumi.Input<string>;
    /**
     * Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
     */
    readonly databaseName: pulumi.Input<string>;
    /**
     * Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
     */
    readonly eventhubId: pulumi.Input<string>;
    /**
     * The location where the Kusto Database should be created. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
     */
    readonly mappingRuleName?: pulumi.Input<string>;
    /**
     * The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the target table name used for the message ingestion. Table must exist before resource is created.
     */
    readonly tableName?: pulumi.Input<string>;
}

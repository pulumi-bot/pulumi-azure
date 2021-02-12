// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a HDInsight Kafka Cluster.
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
 * });
 * const exampleContainer = new azure.storage.Container("exampleContainer", {
 *     storageAccountName: exampleAccount.name,
 *     containerAccessType: "private",
 * });
 * const exampleKafkaCluster = new azure.hdinsight.KafkaCluster("exampleKafkaCluster", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     clusterVersion: "4.0",
 *     tier: "Standard",
 *     componentVersion: {
 *         kafka: "2.1",
 *     },
 *     gateway: {
 *         enabled: true,
 *         username: "acctestusrgw",
 *         password: "Password123!",
 *     },
 *     storageAccounts: [{
 *         storageContainerId: exampleContainer.id,
 *         storageAccountKey: exampleAccount.primaryAccessKey,
 *         isDefault: true,
 *     }],
 *     roles: {
 *         headNode: {
 *             vmSize: "Standard_D3_V2",
 *             username: "acctestusrvm",
 *             password: "AccTestvdSC4daf986!",
 *         },
 *         workerNode: {
 *             vmSize: "Standard_D3_V2",
 *             username: "acctestusrvm",
 *             password: "AccTestvdSC4daf986!",
 *             numberOfDisksPerNode: 3,
 *             targetInstanceCount: 3,
 *         },
 *         zookeeperNode: {
 *             vmSize: "Standard_D3_V2",
 *             username: "acctestusrvm",
 *             password: "AccTestvdSC4daf986!",
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * HDInsight Kafka Clusters can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:hdinsight/kafkaCluster:KafkaCluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.HDInsight/clusters/cluster1
 * ```
 */
export class KafkaCluster extends pulumi.CustomResource {
    /**
     * Get an existing KafkaCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KafkaClusterState, opts?: pulumi.CustomResourceOptions): KafkaCluster {
        return new KafkaCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:hdinsight/kafkaCluster:KafkaCluster';

    /**
     * Returns true if the given object is an instance of KafkaCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KafkaCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KafkaCluster.__pulumiType;
    }

    /**
     * Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
     */
    public readonly clusterVersion!: pulumi.Output<string>;
    /**
     * A `componentVersion` block as defined below.
     */
    public readonly componentVersion!: pulumi.Output<outputs.hdinsight.KafkaClusterComponentVersion>;
    /**
     * A `gateway` block as defined below.
     */
    public readonly gateway!: pulumi.Output<outputs.hdinsight.KafkaClusterGateway>;
    /**
     * The HTTPS Connectivity Endpoint for this HDInsight Kafka Cluster.
     */
    public /*out*/ readonly httpsEndpoint!: pulumi.Output<string>;
    /**
     * The Kafka Rest Proxy Endpoint for this HDInsight Kafka Cluster.
     */
    public /*out*/ readonly kafkaRestProxyEndpoint!: pulumi.Output<string>;
    /**
     * Specifies the Azure Region which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * A `metastores` block as defined below.
     */
    public readonly metastores!: pulumi.Output<outputs.hdinsight.KafkaClusterMetastores | undefined>;
    /**
     * A `monitor` block as defined below.
     */
    public readonly monitor!: pulumi.Output<outputs.hdinsight.KafkaClusterMonitor | undefined>;
    /**
     * Specifies the name for this HDInsight Kafka Cluster. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the name of the Resource Group in which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `restProxy` block as defined below.
     */
    public readonly restProxy!: pulumi.Output<outputs.hdinsight.KafkaClusterRestProxy | undefined>;
    /**
     * A `roles` block as defined below.
     */
    public readonly roles!: pulumi.Output<outputs.hdinsight.KafkaClusterRoles>;
    /**
     * The SSH Connectivity Endpoint for this HDInsight Kafka Cluster.
     */
    public /*out*/ readonly sshEndpoint!: pulumi.Output<string>;
    /**
     * A `storageAccountGen2` block as defined below.
     */
    public readonly storageAccountGen2!: pulumi.Output<outputs.hdinsight.KafkaClusterStorageAccountGen2 | undefined>;
    /**
     * One or more `storageAccount` block as defined below.
     */
    public readonly storageAccounts!: pulumi.Output<outputs.hdinsight.KafkaClusterStorageAccount[] | undefined>;
    /**
     * A map of Tags which should be assigned to this HDInsight Kafka Cluster.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the Tier which should be used for this HDInsight Kafka Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
     */
    public readonly tier!: pulumi.Output<string>;
    public readonly tlsMinVersion!: pulumi.Output<string | undefined>;

    /**
     * Create a KafkaCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KafkaClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KafkaClusterArgs | KafkaClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KafkaClusterState | undefined;
            inputs["clusterVersion"] = state ? state.clusterVersion : undefined;
            inputs["componentVersion"] = state ? state.componentVersion : undefined;
            inputs["gateway"] = state ? state.gateway : undefined;
            inputs["httpsEndpoint"] = state ? state.httpsEndpoint : undefined;
            inputs["kafkaRestProxyEndpoint"] = state ? state.kafkaRestProxyEndpoint : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["metastores"] = state ? state.metastores : undefined;
            inputs["monitor"] = state ? state.monitor : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["restProxy"] = state ? state.restProxy : undefined;
            inputs["roles"] = state ? state.roles : undefined;
            inputs["sshEndpoint"] = state ? state.sshEndpoint : undefined;
            inputs["storageAccountGen2"] = state ? state.storageAccountGen2 : undefined;
            inputs["storageAccounts"] = state ? state.storageAccounts : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tier"] = state ? state.tier : undefined;
            inputs["tlsMinVersion"] = state ? state.tlsMinVersion : undefined;
        } else {
            const args = argsOrState as KafkaClusterArgs | undefined;
            if ((!args || args.clusterVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterVersion'");
            }
            if ((!args || args.componentVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'componentVersion'");
            }
            if ((!args || args.gateway === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gateway'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.roles === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roles'");
            }
            if ((!args || args.tier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tier'");
            }
            inputs["clusterVersion"] = args ? args.clusterVersion : undefined;
            inputs["componentVersion"] = args ? args.componentVersion : undefined;
            inputs["gateway"] = args ? args.gateway : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["metastores"] = args ? args.metastores : undefined;
            inputs["monitor"] = args ? args.monitor : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["restProxy"] = args ? args.restProxy : undefined;
            inputs["roles"] = args ? args.roles : undefined;
            inputs["storageAccountGen2"] = args ? args.storageAccountGen2 : undefined;
            inputs["storageAccounts"] = args ? args.storageAccounts : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tier"] = args ? args.tier : undefined;
            inputs["tlsMinVersion"] = args ? args.tlsMinVersion : undefined;
            inputs["httpsEndpoint"] = undefined /*out*/;
            inputs["kafkaRestProxyEndpoint"] = undefined /*out*/;
            inputs["sshEndpoint"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(KafkaCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KafkaCluster resources.
 */
export interface KafkaClusterState {
    /**
     * Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
     */
    readonly clusterVersion?: pulumi.Input<string>;
    /**
     * A `componentVersion` block as defined below.
     */
    readonly componentVersion?: pulumi.Input<inputs.hdinsight.KafkaClusterComponentVersion>;
    /**
     * A `gateway` block as defined below.
     */
    readonly gateway?: pulumi.Input<inputs.hdinsight.KafkaClusterGateway>;
    /**
     * The HTTPS Connectivity Endpoint for this HDInsight Kafka Cluster.
     */
    readonly httpsEndpoint?: pulumi.Input<string>;
    /**
     * The Kafka Rest Proxy Endpoint for this HDInsight Kafka Cluster.
     */
    readonly kafkaRestProxyEndpoint?: pulumi.Input<string>;
    /**
     * Specifies the Azure Region which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * A `metastores` block as defined below.
     */
    readonly metastores?: pulumi.Input<inputs.hdinsight.KafkaClusterMetastores>;
    /**
     * A `monitor` block as defined below.
     */
    readonly monitor?: pulumi.Input<inputs.hdinsight.KafkaClusterMonitor>;
    /**
     * Specifies the name for this HDInsight Kafka Cluster. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the name of the Resource Group in which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A `restProxy` block as defined below.
     */
    readonly restProxy?: pulumi.Input<inputs.hdinsight.KafkaClusterRestProxy>;
    /**
     * A `roles` block as defined below.
     */
    readonly roles?: pulumi.Input<inputs.hdinsight.KafkaClusterRoles>;
    /**
     * The SSH Connectivity Endpoint for this HDInsight Kafka Cluster.
     */
    readonly sshEndpoint?: pulumi.Input<string>;
    /**
     * A `storageAccountGen2` block as defined below.
     */
    readonly storageAccountGen2?: pulumi.Input<inputs.hdinsight.KafkaClusterStorageAccountGen2>;
    /**
     * One or more `storageAccount` block as defined below.
     */
    readonly storageAccounts?: pulumi.Input<pulumi.Input<inputs.hdinsight.KafkaClusterStorageAccount>[]>;
    /**
     * A map of Tags which should be assigned to this HDInsight Kafka Cluster.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the Tier which should be used for this HDInsight Kafka Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
     */
    readonly tier?: pulumi.Input<string>;
    readonly tlsMinVersion?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KafkaCluster resource.
 */
export interface KafkaClusterArgs {
    /**
     * Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
     */
    readonly clusterVersion: pulumi.Input<string>;
    /**
     * A `componentVersion` block as defined below.
     */
    readonly componentVersion: pulumi.Input<inputs.hdinsight.KafkaClusterComponentVersion>;
    /**
     * A `gateway` block as defined below.
     */
    readonly gateway: pulumi.Input<inputs.hdinsight.KafkaClusterGateway>;
    /**
     * Specifies the Azure Region which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * A `metastores` block as defined below.
     */
    readonly metastores?: pulumi.Input<inputs.hdinsight.KafkaClusterMetastores>;
    /**
     * A `monitor` block as defined below.
     */
    readonly monitor?: pulumi.Input<inputs.hdinsight.KafkaClusterMonitor>;
    /**
     * Specifies the name for this HDInsight Kafka Cluster. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the name of the Resource Group in which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `restProxy` block as defined below.
     */
    readonly restProxy?: pulumi.Input<inputs.hdinsight.KafkaClusterRestProxy>;
    /**
     * A `roles` block as defined below.
     */
    readonly roles: pulumi.Input<inputs.hdinsight.KafkaClusterRoles>;
    /**
     * A `storageAccountGen2` block as defined below.
     */
    readonly storageAccountGen2?: pulumi.Input<inputs.hdinsight.KafkaClusterStorageAccountGen2>;
    /**
     * One or more `storageAccount` block as defined below.
     */
    readonly storageAccounts?: pulumi.Input<pulumi.Input<inputs.hdinsight.KafkaClusterStorageAccount>[]>;
    /**
     * A map of Tags which should be assigned to this HDInsight Kafka Cluster.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the Tier which should be used for this HDInsight Kafka Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
     */
    readonly tier: pulumi.Input<string>;
    readonly tlsMinVersion?: pulumi.Input<string>;
}

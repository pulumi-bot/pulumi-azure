// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a HDInsight Hadoop Cluster.
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
 * const exampleHadoopCluster = new azure.hdinsight.HadoopCluster("exampleHadoopCluster", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     clusterVersion: "3.6",
 *     tier: "Standard",
 *     componentVersion: {
 *         hadoop: "2.7",
 *     },
 *     gateway: {
 *         enabled: true,
 *         username: "acctestusrgw",
 *         password: "PAssword123!",
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
 *             vmSize: "Standard_D4_V2",
 *             username: "acctestusrvm",
 *             password: "AccTestvdSC4daf986!",
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
 * HDInsight Hadoop Clusters can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:hdinsight/hadoopCluster:HadoopCluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.HDInsight/clusters/cluster1
 * ```
 */
export class HadoopCluster extends pulumi.CustomResource {
    /**
     * Get an existing HadoopCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HadoopClusterState, opts?: pulumi.CustomResourceOptions): HadoopCluster {
        return new HadoopCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:hdinsight/hadoopCluster:HadoopCluster';

    /**
     * Returns true if the given object is an instance of HadoopCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HadoopCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HadoopCluster.__pulumiType;
    }

    /**
     * Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
     */
    public readonly clusterVersion!: pulumi.Output<string>;
    /**
     * A `componentVersion` block as defined below.
     */
    public readonly componentVersion!: pulumi.Output<outputs.hdinsight.HadoopClusterComponentVersion>;
    /**
     * A `gateway` block as defined below.
     */
    public readonly gateway!: pulumi.Output<outputs.hdinsight.HadoopClusterGateway>;
    /**
     * The HTTPS Connectivity Endpoint for this HDInsight Hadoop Cluster.
     */
    public /*out*/ readonly httpsEndpoint!: pulumi.Output<string>;
    /**
     * Specifies the Azure Region which this HDInsight Hadoop Cluster should exist. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * A `metastores` block as defined below.
     */
    public readonly metastores!: pulumi.Output<outputs.hdinsight.HadoopClusterMetastores | undefined>;
    /**
     * A `monitor` block as defined below.
     */
    public readonly monitor!: pulumi.Output<outputs.hdinsight.HadoopClusterMonitor | undefined>;
    /**
     * Specifies the name for this HDInsight Hadoop Cluster. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A `network` block as defined below.
     */
    public readonly network!: pulumi.Output<outputs.hdinsight.HadoopClusterNetwork | undefined>;
    /**
     * Specifies the name of the Resource Group in which this HDInsight Hadoop Cluster should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `roles` block as defined below.
     */
    public readonly roles!: pulumi.Output<outputs.hdinsight.HadoopClusterRoles>;
    /**
     * The SSH Connectivity Endpoint for this HDInsight Hadoop Cluster.
     */
    public /*out*/ readonly sshEndpoint!: pulumi.Output<string>;
    /**
     * A `storageAccountGen2` block as defined below.
     */
    public readonly storageAccountGen2!: pulumi.Output<outputs.hdinsight.HadoopClusterStorageAccountGen2 | undefined>;
    /**
     * One or more `storageAccount` block as defined below.
     */
    public readonly storageAccounts!: pulumi.Output<outputs.hdinsight.HadoopClusterStorageAccount[] | undefined>;
    /**
     * A map of Tags which should be assigned to this HDInsight Hadoop Cluster.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the Tier which should be used for this HDInsight Hadoop Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
     */
    public readonly tier!: pulumi.Output<string>;
    public readonly tlsMinVersion!: pulumi.Output<string | undefined>;

    /**
     * Create a HadoopCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HadoopClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HadoopClusterArgs | HadoopClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HadoopClusterState | undefined;
            inputs["clusterVersion"] = state ? state.clusterVersion : undefined;
            inputs["componentVersion"] = state ? state.componentVersion : undefined;
            inputs["gateway"] = state ? state.gateway : undefined;
            inputs["httpsEndpoint"] = state ? state.httpsEndpoint : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["metastores"] = state ? state.metastores : undefined;
            inputs["monitor"] = state ? state.monitor : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["roles"] = state ? state.roles : undefined;
            inputs["sshEndpoint"] = state ? state.sshEndpoint : undefined;
            inputs["storageAccountGen2"] = state ? state.storageAccountGen2 : undefined;
            inputs["storageAccounts"] = state ? state.storageAccounts : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tier"] = state ? state.tier : undefined;
            inputs["tlsMinVersion"] = state ? state.tlsMinVersion : undefined;
        } else {
            const args = argsOrState as HadoopClusterArgs | undefined;
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
            inputs["network"] = args ? args.network : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["roles"] = args ? args.roles : undefined;
            inputs["storageAccountGen2"] = args ? args.storageAccountGen2 : undefined;
            inputs["storageAccounts"] = args ? args.storageAccounts : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tier"] = args ? args.tier : undefined;
            inputs["tlsMinVersion"] = args ? args.tlsMinVersion : undefined;
            inputs["httpsEndpoint"] = undefined /*out*/;
            inputs["sshEndpoint"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(HadoopCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HadoopCluster resources.
 */
export interface HadoopClusterState {
    /**
     * Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
     */
    clusterVersion?: pulumi.Input<string>;
    /**
     * A `componentVersion` block as defined below.
     */
    componentVersion?: pulumi.Input<inputs.hdinsight.HadoopClusterComponentVersion>;
    /**
     * A `gateway` block as defined below.
     */
    gateway?: pulumi.Input<inputs.hdinsight.HadoopClusterGateway>;
    /**
     * The HTTPS Connectivity Endpoint for this HDInsight Hadoop Cluster.
     */
    httpsEndpoint?: pulumi.Input<string>;
    /**
     * Specifies the Azure Region which this HDInsight Hadoop Cluster should exist. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * A `metastores` block as defined below.
     */
    metastores?: pulumi.Input<inputs.hdinsight.HadoopClusterMetastores>;
    /**
     * A `monitor` block as defined below.
     */
    monitor?: pulumi.Input<inputs.hdinsight.HadoopClusterMonitor>;
    /**
     * Specifies the name for this HDInsight Hadoop Cluster. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A `network` block as defined below.
     */
    network?: pulumi.Input<inputs.hdinsight.HadoopClusterNetwork>;
    /**
     * Specifies the name of the Resource Group in which this HDInsight Hadoop Cluster should exist. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * A `roles` block as defined below.
     */
    roles?: pulumi.Input<inputs.hdinsight.HadoopClusterRoles>;
    /**
     * The SSH Connectivity Endpoint for this HDInsight Hadoop Cluster.
     */
    sshEndpoint?: pulumi.Input<string>;
    /**
     * A `storageAccountGen2` block as defined below.
     */
    storageAccountGen2?: pulumi.Input<inputs.hdinsight.HadoopClusterStorageAccountGen2>;
    /**
     * One or more `storageAccount` block as defined below.
     */
    storageAccounts?: pulumi.Input<pulumi.Input<inputs.hdinsight.HadoopClusterStorageAccount>[]>;
    /**
     * A map of Tags which should be assigned to this HDInsight Hadoop Cluster.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the Tier which should be used for this HDInsight Hadoop Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
     */
    tier?: pulumi.Input<string>;
    tlsMinVersion?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HadoopCluster resource.
 */
export interface HadoopClusterArgs {
    /**
     * Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
     */
    clusterVersion: pulumi.Input<string>;
    /**
     * A `componentVersion` block as defined below.
     */
    componentVersion: pulumi.Input<inputs.hdinsight.HadoopClusterComponentVersion>;
    /**
     * A `gateway` block as defined below.
     */
    gateway: pulumi.Input<inputs.hdinsight.HadoopClusterGateway>;
    /**
     * Specifies the Azure Region which this HDInsight Hadoop Cluster should exist. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * A `metastores` block as defined below.
     */
    metastores?: pulumi.Input<inputs.hdinsight.HadoopClusterMetastores>;
    /**
     * A `monitor` block as defined below.
     */
    monitor?: pulumi.Input<inputs.hdinsight.HadoopClusterMonitor>;
    /**
     * Specifies the name for this HDInsight Hadoop Cluster. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A `network` block as defined below.
     */
    network?: pulumi.Input<inputs.hdinsight.HadoopClusterNetwork>;
    /**
     * Specifies the name of the Resource Group in which this HDInsight Hadoop Cluster should exist. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * A `roles` block as defined below.
     */
    roles: pulumi.Input<inputs.hdinsight.HadoopClusterRoles>;
    /**
     * A `storageAccountGen2` block as defined below.
     */
    storageAccountGen2?: pulumi.Input<inputs.hdinsight.HadoopClusterStorageAccountGen2>;
    /**
     * One or more `storageAccount` block as defined below.
     */
    storageAccounts?: pulumi.Input<pulumi.Input<inputs.hdinsight.HadoopClusterStorageAccount>[]>;
    /**
     * A map of Tags which should be assigned to this HDInsight Hadoop Cluster.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the Tier which should be used for this HDInsight Hadoop Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
     */
    tier: pulumi.Input<string>;
    tlsMinVersion?: pulumi.Input<string>;
}

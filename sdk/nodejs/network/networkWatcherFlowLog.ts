// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Network Watcher Flow Log.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const testResourceGroup = new azure.core.ResourceGroup("testResourceGroup", {location: "eastus"});
 * const testNetworkSecurityGroup = new azure.network.NetworkSecurityGroup("testNetworkSecurityGroup", {
 *     location: testResourceGroup.location,
 *     resourceGroupName: testResourceGroup.name,
 * });
 * const testNetworkWatcher = new azure.network.NetworkWatcher("testNetworkWatcher", {
 *     location: testResourceGroup.location,
 *     resourceGroupName: testResourceGroup.name,
 * });
 * const testAccount = new azure.storage.Account("testAccount", {
 *     resourceGroupName: testResourceGroup.name,
 *     location: testResourceGroup.location,
 *     accountTier: "Standard",
 *     accountKind: "StorageV2",
 *     accountReplicationType: "LRS",
 *     enableHttpsTrafficOnly: true,
 * });
 * const testAnalyticsWorkspace = new azure.operationalinsights.AnalyticsWorkspace("testAnalyticsWorkspace", {
 *     location: testResourceGroup.location,
 *     resourceGroupName: testResourceGroup.name,
 *     sku: "PerGB2018",
 * });
 * const testNetworkWatcherFlowLog = new azure.network.NetworkWatcherFlowLog("testNetworkWatcherFlowLog", {
 *     networkWatcherName: testNetworkWatcher.name,
 *     resourceGroupName: testResourceGroup.name,
 *     networkSecurityGroupId: testNetworkSecurityGroup.id,
 *     storageAccountId: testAccount.id,
 *     enabled: true,
 *     retention_policy: {
 *         enabled: true,
 *         days: 7,
 *     },
 *     traffic_analytics: {
 *         enabled: true,
 *         workspaceId: testAnalyticsWorkspace.workspaceId,
 *         workspaceRegion: testAnalyticsWorkspace.location,
 *         workspaceResourceId: testAnalyticsWorkspace.id,
 *         intervalInMinutes: 10,
 *     },
 * });
 * ```
 */
export class NetworkWatcherFlowLog extends pulumi.CustomResource {
    /**
     * Get an existing NetworkWatcherFlowLog resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkWatcherFlowLogState, opts?: pulumi.CustomResourceOptions): NetworkWatcherFlowLog {
        return new NetworkWatcherFlowLog(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/networkWatcherFlowLog:NetworkWatcherFlowLog';

    /**
     * Returns true if the given object is an instance of NetworkWatcherFlowLog.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkWatcherFlowLog {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkWatcherFlowLog.__pulumiType;
    }

    /**
     * Boolean flag to enable/disable traffic analytics.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * The ID of the Network Security Group for which to enable flow logs for. Changing this forces a new resource to be created.
     */
    public readonly networkSecurityGroupId!: pulumi.Output<string>;
    /**
     * The name of the Network Watcher. Changing this forces a new resource to be created.
     */
    public readonly networkWatcherName!: pulumi.Output<string>;
    /**
     * The name of the resource group in which the Network Watcher was deployed. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `retentionPolicy` block as documented below.
     */
    public readonly retentionPolicy!: pulumi.Output<outputs.network.NetworkWatcherFlowLogRetentionPolicy>;
    /**
     * The ID of the Storage Account where flow logs are stored.
     */
    public readonly storageAccountId!: pulumi.Output<string>;
    /**
     * A `trafficAnalytics` block as documented below.
     */
    public readonly trafficAnalytics!: pulumi.Output<outputs.network.NetworkWatcherFlowLogTrafficAnalytics | undefined>;
    /**
     * The version (revision) of the flow log. Possible values are `1` and `2`.
     */
    public readonly version!: pulumi.Output<number>;

    /**
     * Create a NetworkWatcherFlowLog resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkWatcherFlowLogArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkWatcherFlowLogArgs | NetworkWatcherFlowLogState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NetworkWatcherFlowLogState | undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["networkSecurityGroupId"] = state ? state.networkSecurityGroupId : undefined;
            inputs["networkWatcherName"] = state ? state.networkWatcherName : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["retentionPolicy"] = state ? state.retentionPolicy : undefined;
            inputs["storageAccountId"] = state ? state.storageAccountId : undefined;
            inputs["trafficAnalytics"] = state ? state.trafficAnalytics : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as NetworkWatcherFlowLogArgs | undefined;
            if (!args || args.enabled === undefined) {
                throw new Error("Missing required property 'enabled'");
            }
            if (!args || args.networkSecurityGroupId === undefined) {
                throw new Error("Missing required property 'networkSecurityGroupId'");
            }
            if (!args || args.networkWatcherName === undefined) {
                throw new Error("Missing required property 'networkWatcherName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.retentionPolicy === undefined) {
                throw new Error("Missing required property 'retentionPolicy'");
            }
            if (!args || args.storageAccountId === undefined) {
                throw new Error("Missing required property 'storageAccountId'");
            }
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["networkSecurityGroupId"] = args ? args.networkSecurityGroupId : undefined;
            inputs["networkWatcherName"] = args ? args.networkWatcherName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["retentionPolicy"] = args ? args.retentionPolicy : undefined;
            inputs["storageAccountId"] = args ? args.storageAccountId : undefined;
            inputs["trafficAnalytics"] = args ? args.trafficAnalytics : undefined;
            inputs["version"] = args ? args.version : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NetworkWatcherFlowLog.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkWatcherFlowLog resources.
 */
export interface NetworkWatcherFlowLogState {
    /**
     * Boolean flag to enable/disable traffic analytics.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The ID of the Network Security Group for which to enable flow logs for. Changing this forces a new resource to be created.
     */
    readonly networkSecurityGroupId?: pulumi.Input<string>;
    /**
     * The name of the Network Watcher. Changing this forces a new resource to be created.
     */
    readonly networkWatcherName?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the Network Watcher was deployed. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A `retentionPolicy` block as documented below.
     */
    readonly retentionPolicy?: pulumi.Input<inputs.network.NetworkWatcherFlowLogRetentionPolicy>;
    /**
     * The ID of the Storage Account where flow logs are stored.
     */
    readonly storageAccountId?: pulumi.Input<string>;
    /**
     * A `trafficAnalytics` block as documented below.
     */
    readonly trafficAnalytics?: pulumi.Input<inputs.network.NetworkWatcherFlowLogTrafficAnalytics>;
    /**
     * The version (revision) of the flow log. Possible values are `1` and `2`.
     */
    readonly version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a NetworkWatcherFlowLog resource.
 */
export interface NetworkWatcherFlowLogArgs {
    /**
     * Boolean flag to enable/disable traffic analytics.
     */
    readonly enabled: pulumi.Input<boolean>;
    /**
     * The ID of the Network Security Group for which to enable flow logs for. Changing this forces a new resource to be created.
     */
    readonly networkSecurityGroupId: pulumi.Input<string>;
    /**
     * The name of the Network Watcher. Changing this forces a new resource to be created.
     */
    readonly networkWatcherName: pulumi.Input<string>;
    /**
     * The name of the resource group in which the Network Watcher was deployed. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `retentionPolicy` block as documented below.
     */
    readonly retentionPolicy: pulumi.Input<inputs.network.NetworkWatcherFlowLogRetentionPolicy>;
    /**
     * The ID of the Storage Account where flow logs are stored.
     */
    readonly storageAccountId: pulumi.Input<string>;
    /**
     * A `trafficAnalytics` block as documented below.
     */
    readonly trafficAnalytics?: pulumi.Input<inputs.network.NetworkWatcherFlowLogTrafficAnalytics>;
    /**
     * The version (revision) of the flow log. Possible values are `1` and `2`.
     */
    readonly version?: pulumi.Input<number>;
}

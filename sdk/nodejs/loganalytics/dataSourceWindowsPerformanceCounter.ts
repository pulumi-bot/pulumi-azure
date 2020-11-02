// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Log Analytics (formally Operational Insights) Windows Performance Counter DataSource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAnalyticsWorkspace = new azure.operationalinsights.AnalyticsWorkspace("exampleAnalyticsWorkspace", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "PerGB2018",
 * });
 * const exampleDataSourceWindowsPerformanceCounter = new azure.loganalytics.DataSourceWindowsPerformanceCounter("exampleDataSourceWindowsPerformanceCounter", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     workspaceName: exampleAnalyticsWorkspace.name,
 *     objectName: "CPU",
 *     instanceName: "*",
 *     counterName: "CPU",
 *     intervalSeconds: 10,
 * });
 * ```
 *
 * ## Import
 *
 * Log Analytics Windows Performance Counter DataSources can be imported using the `resource id`, e.g.
 */
export class DataSourceWindowsPerformanceCounter extends pulumi.CustomResource {
    /**
     * Get an existing DataSourceWindowsPerformanceCounter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataSourceWindowsPerformanceCounterState, opts?: pulumi.CustomResourceOptions): DataSourceWindowsPerformanceCounter {
        return new DataSourceWindowsPerformanceCounter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:loganalytics/dataSourceWindowsPerformanceCounter:DataSourceWindowsPerformanceCounter';

    /**
     * Returns true if the given object is an instance of DataSourceWindowsPerformanceCounter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataSourceWindowsPerformanceCounter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataSourceWindowsPerformanceCounter.__pulumiType;
    }

    /**
     * The friendly name of the performance counter.
     */
    public readonly counterName!: pulumi.Output<string>;
    /**
     * The name of the virtual machine instance to which the Windows Performance Counter DataSource be applied. Specify a `*` will apply to all instances.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * The time of sample interval in seconds. Supports values between 10 and 2147483647.
     */
    public readonly intervalSeconds!: pulumi.Output<number>;
    /**
     * The Name which should be used for this Log Analytics Windows Performance Counter DataSource. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The object name of the Log Analytics Windows Performance Counter DataSource.
     */
    public readonly objectName!: pulumi.Output<string>;
    /**
     * The name of the Resource Group where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The name of the Log Analytics Workspace where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
     */
    public readonly workspaceName!: pulumi.Output<string>;

    /**
     * Create a DataSourceWindowsPerformanceCounter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataSourceWindowsPerformanceCounterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataSourceWindowsPerformanceCounterArgs | DataSourceWindowsPerformanceCounterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DataSourceWindowsPerformanceCounterState | undefined;
            inputs["counterName"] = state ? state.counterName : undefined;
            inputs["instanceName"] = state ? state.instanceName : undefined;
            inputs["intervalSeconds"] = state ? state.intervalSeconds : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["objectName"] = state ? state.objectName : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["workspaceName"] = state ? state.workspaceName : undefined;
        } else {
            const args = argsOrState as DataSourceWindowsPerformanceCounterArgs | undefined;
            if (!args || args.counterName === undefined) {
                throw new Error("Missing required property 'counterName'");
            }
            if (!args || args.instanceName === undefined) {
                throw new Error("Missing required property 'instanceName'");
            }
            if (!args || args.intervalSeconds === undefined) {
                throw new Error("Missing required property 'intervalSeconds'");
            }
            if (!args || args.objectName === undefined) {
                throw new Error("Missing required property 'objectName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["counterName"] = args ? args.counterName : undefined;
            inputs["instanceName"] = args ? args.instanceName : undefined;
            inputs["intervalSeconds"] = args ? args.intervalSeconds : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["objectName"] = args ? args.objectName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DataSourceWindowsPerformanceCounter.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataSourceWindowsPerformanceCounter resources.
 */
export interface DataSourceWindowsPerformanceCounterState {
    /**
     * The friendly name of the performance counter.
     */
    readonly counterName?: pulumi.Input<string>;
    /**
     * The name of the virtual machine instance to which the Windows Performance Counter DataSource be applied. Specify a `*` will apply to all instances.
     */
    readonly instanceName?: pulumi.Input<string>;
    /**
     * The time of sample interval in seconds. Supports values between 10 and 2147483647.
     */
    readonly intervalSeconds?: pulumi.Input<number>;
    /**
     * The Name which should be used for this Log Analytics Windows Performance Counter DataSource. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The object name of the Log Analytics Windows Performance Counter DataSource.
     */
    readonly objectName?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The name of the Log Analytics Workspace where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
     */
    readonly workspaceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DataSourceWindowsPerformanceCounter resource.
 */
export interface DataSourceWindowsPerformanceCounterArgs {
    /**
     * The friendly name of the performance counter.
     */
    readonly counterName: pulumi.Input<string>;
    /**
     * The name of the virtual machine instance to which the Windows Performance Counter DataSource be applied. Specify a `*` will apply to all instances.
     */
    readonly instanceName: pulumi.Input<string>;
    /**
     * The time of sample interval in seconds. Supports values between 10 and 2147483647.
     */
    readonly intervalSeconds: pulumi.Input<number>;
    /**
     * The Name which should be used for this Log Analytics Windows Performance Counter DataSource. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The object name of the Log Analytics Windows Performance Counter DataSource.
     */
    readonly objectName: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Log Analytics Workspace where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
     */
    readonly workspaceName: pulumi.Input<string>;
}

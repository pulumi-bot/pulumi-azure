// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Log Analytics Data Export Rule.
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
 *     retentionInDays: 30,
 * });
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const exampleDataExportRule = new azure.loganalytics.DataExportRule("exampleDataExportRule", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     workspaceResourceId: exampleAnalyticsWorkspace.id,
 *     destinationResourceId: exampleAccount.id,
 *     tableNames: ["Heartbeat"],
 * });
 * ```
 *
 * ## Import
 *
 * Log Analytics Data Export Rule can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:loganalytics/dataExportRule:DataExportRule example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/dataExports/dataExport1
 * ```
 */
export class DataExportRule extends pulumi.CustomResource {
    /**
     * Get an existing DataExportRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataExportRuleState, opts?: pulumi.CustomResourceOptions): DataExportRule {
        return new DataExportRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:loganalytics/dataExportRule:DataExportRule';

    /**
     * Returns true if the given object is an instance of DataExportRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataExportRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataExportRule.__pulumiType;
    }

    /**
     * The destination resource ID. It should be a storage account, an event hub namespace or an event hub. If the destination is an event hub namespace, an event hub would be created for each table automatically.
     */
    public readonly destinationResourceId!: pulumi.Output<string>;
    /**
     * Is this Log Analytics Data Export Rule when enabled? Possible values include `true` or `false`. Defaults to `false`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the created Data Export Rule.
     */
    public /*out*/ readonly exportRuleId!: pulumi.Output<string>;
    /**
     * The name of the Log Analytics Data Export Rule. Changing this forces a new Log Analytics Data Export Rule to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Resource Group where the Log Analytics Data Export should exist. Changing this forces a new Log Analytics Data Export Rule to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A list of table names to export to the destination resource, for example: `["Heartbeat", "SecurityEvent"]`.
     */
    public readonly tableNames!: pulumi.Output<string[]>;
    /**
     * The resource ID of the workspace. Changing this forces a new Log Analytics Data Export Rule to be created.
     */
    public readonly workspaceResourceId!: pulumi.Output<string>;

    /**
     * Create a DataExportRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataExportRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataExportRuleArgs | DataExportRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DataExportRuleState | undefined;
            inputs["destinationResourceId"] = state ? state.destinationResourceId : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["exportRuleId"] = state ? state.exportRuleId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tableNames"] = state ? state.tableNames : undefined;
            inputs["workspaceResourceId"] = state ? state.workspaceResourceId : undefined;
        } else {
            const args = argsOrState as DataExportRuleArgs | undefined;
            if (!args || args.destinationResourceId === undefined) {
                throw new Error("Missing required property 'destinationResourceId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.tableNames === undefined) {
                throw new Error("Missing required property 'tableNames'");
            }
            if (!args || args.workspaceResourceId === undefined) {
                throw new Error("Missing required property 'workspaceResourceId'");
            }
            inputs["destinationResourceId"] = args ? args.destinationResourceId : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tableNames"] = args ? args.tableNames : undefined;
            inputs["workspaceResourceId"] = args ? args.workspaceResourceId : undefined;
            inputs["exportRuleId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DataExportRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataExportRule resources.
 */
export interface DataExportRuleState {
    /**
     * The destination resource ID. It should be a storage account, an event hub namespace or an event hub. If the destination is an event hub namespace, an event hub would be created for each table automatically.
     */
    readonly destinationResourceId?: pulumi.Input<string>;
    /**
     * Is this Log Analytics Data Export Rule when enabled? Possible values include `true` or `false`. Defaults to `false`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The ID of the created Data Export Rule.
     */
    readonly exportRuleId?: pulumi.Input<string>;
    /**
     * The name of the Log Analytics Data Export Rule. Changing this forces a new Log Analytics Data Export Rule to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Log Analytics Data Export should exist. Changing this forces a new Log Analytics Data Export Rule to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A list of table names to export to the destination resource, for example: `["Heartbeat", "SecurityEvent"]`.
     */
    readonly tableNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The resource ID of the workspace. Changing this forces a new Log Analytics Data Export Rule to be created.
     */
    readonly workspaceResourceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DataExportRule resource.
 */
export interface DataExportRuleArgs {
    /**
     * The destination resource ID. It should be a storage account, an event hub namespace or an event hub. If the destination is an event hub namespace, an event hub would be created for each table automatically.
     */
    readonly destinationResourceId: pulumi.Input<string>;
    /**
     * Is this Log Analytics Data Export Rule when enabled? Possible values include `true` or `false`. Defaults to `false`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The name of the Log Analytics Data Export Rule. Changing this forces a new Log Analytics Data Export Rule to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Log Analytics Data Export should exist. Changing this forces a new Log Analytics Data Export Rule to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A list of table names to export to the destination resource, for example: `["Heartbeat", "SecurityEvent"]`.
     */
    readonly tableNames: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The resource ID of the workspace. Changing this forces a new Log Analytics Data Export Rule to be created.
     */
    readonly workspaceResourceId: pulumi.Input<string>;
}

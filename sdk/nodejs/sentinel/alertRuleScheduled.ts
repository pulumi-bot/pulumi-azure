// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Sentinel Scheduled Alert Rule.
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
 *     sku: "pergb2018",
 * });
 * const exampleAlertRuleScheduled = new azure.sentinel.AlertRuleScheduled("exampleAlertRuleScheduled", {
 *     logAnalyticsWorkspaceId: exampleAnalyticsWorkspace.id,
 *     displayName: "example",
 *     severity: "High",
 *     query: `AzureActivity |
 *   where OperationName == "Create or Update Virtual Machine" or OperationName =="Create Deployment" |
 *   where ActivityStatus == "Succeeded" |
 *   make-series dcount(ResourceId) default=0 on EventSubmissionTimestamp in range(ago(7d), now(), 1d) by Caller
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * Sentinel Scheduled Alert Rules can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:sentinel/alertRuleScheduled:AlertRuleScheduled example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
 * ```
 */
export class AlertRuleScheduled extends pulumi.CustomResource {
    /**
     * Get an existing AlertRuleScheduled resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlertRuleScheduledState, opts?: pulumi.CustomResourceOptions): AlertRuleScheduled {
        return new AlertRuleScheduled(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:sentinel/alertRuleScheduled:AlertRuleScheduled';

    /**
     * Returns true if the given object is an instance of AlertRuleScheduled.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlertRuleScheduled {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlertRuleScheduled.__pulumiType;
    }

    /**
     * The description of this Sentinel Scheduled Alert Rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The friendly name of this Sentinel Scheduled Alert Rule.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Should the Sentinel Scheduled Alert Rule be enabled? Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the Log Analytics Workspace this Sentinel Scheduled Alert Rule belongs to. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
     */
    public readonly logAnalyticsWorkspaceId!: pulumi.Output<string>;
    /**
     * The name which should be used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The query of this Sentinel Scheduled Alert Rule.
     */
    public readonly query!: pulumi.Output<string>;
    /**
     * The ISO 8601 timespan duration between two consecutive queries. Defaults to `PT5H`.
     */
    public readonly queryFrequency!: pulumi.Output<string | undefined>;
    /**
     * The ISO 8601 timespan duration, which determine the time period of the data covered by the query. For example, it can query the past 10 minutes of data, or the past 6 hours of data. Defaults to `PT5H`.
     */
    public readonly queryPeriod!: pulumi.Output<string | undefined>;
    /**
     * The alert severity of this Sentinel Scheduled Alert Rule. Possible values are `High`, `Medium`, `Low` and `Informational`.
     */
    public readonly severity!: pulumi.Output<string>;
    /**
     * If `suppressionEnabled` is `true`, this is ISO 8601 timespan duration, which specifies the amount of time the query should stop running after alert is generated. Defaults to `PT5H`.
     */
    public readonly suppressionDuration!: pulumi.Output<string | undefined>;
    /**
     * Should the Sentinel Scheduled Alert Rulea stop running query after alert is generated? Defaults to `false`.
     */
    public readonly suppressionEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * A list of categories of attacks by which to classify the rule. Possible values are `Collection`, `CommandAndControl`, `CredentialAccess`, `DefenseEvasion`, `Discovery`, `Execution`, `Exfiltration`, `Impact`, `InitialAccess`, `LateralMovement`, `Persistence` and `PrivilegeEscalation`.
     */
    public readonly tactics!: pulumi.Output<string[] | undefined>;
    /**
     * The alert trigger operator, combined with `triggerThreshold`, setting alert threshold of this Sentinel Scheduled Alert Rule. Possible values are `Equal`, `GreaterThan`, `LessThan`, `NotEqual`.
     */
    public readonly triggerOperator!: pulumi.Output<string | undefined>;
    /**
     * The baseline number of query results generated, combined with `triggerOperator`, setting alert threshold of this Sentinel Scheduled Alert Rule.
     */
    public readonly triggerThreshold!: pulumi.Output<number | undefined>;

    /**
     * Create a AlertRuleScheduled resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlertRuleScheduledArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlertRuleScheduledArgs | AlertRuleScheduledState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AlertRuleScheduledState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["logAnalyticsWorkspaceId"] = state ? state.logAnalyticsWorkspaceId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["query"] = state ? state.query : undefined;
            inputs["queryFrequency"] = state ? state.queryFrequency : undefined;
            inputs["queryPeriod"] = state ? state.queryPeriod : undefined;
            inputs["severity"] = state ? state.severity : undefined;
            inputs["suppressionDuration"] = state ? state.suppressionDuration : undefined;
            inputs["suppressionEnabled"] = state ? state.suppressionEnabled : undefined;
            inputs["tactics"] = state ? state.tactics : undefined;
            inputs["triggerOperator"] = state ? state.triggerOperator : undefined;
            inputs["triggerThreshold"] = state ? state.triggerThreshold : undefined;
        } else {
            const args = argsOrState as AlertRuleScheduledArgs | undefined;
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            if (!args || args.logAnalyticsWorkspaceId === undefined) {
                throw new Error("Missing required property 'logAnalyticsWorkspaceId'");
            }
            if (!args || args.query === undefined) {
                throw new Error("Missing required property 'query'");
            }
            if (!args || args.severity === undefined) {
                throw new Error("Missing required property 'severity'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["logAnalyticsWorkspaceId"] = args ? args.logAnalyticsWorkspaceId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["query"] = args ? args.query : undefined;
            inputs["queryFrequency"] = args ? args.queryFrequency : undefined;
            inputs["queryPeriod"] = args ? args.queryPeriod : undefined;
            inputs["severity"] = args ? args.severity : undefined;
            inputs["suppressionDuration"] = args ? args.suppressionDuration : undefined;
            inputs["suppressionEnabled"] = args ? args.suppressionEnabled : undefined;
            inputs["tactics"] = args ? args.tactics : undefined;
            inputs["triggerOperator"] = args ? args.triggerOperator : undefined;
            inputs["triggerThreshold"] = args ? args.triggerThreshold : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AlertRuleScheduled.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AlertRuleScheduled resources.
 */
export interface AlertRuleScheduledState {
    /**
     * The description of this Sentinel Scheduled Alert Rule.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The friendly name of this Sentinel Scheduled Alert Rule.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Should the Sentinel Scheduled Alert Rule be enabled? Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The ID of the Log Analytics Workspace this Sentinel Scheduled Alert Rule belongs to. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
     */
    readonly logAnalyticsWorkspaceId?: pulumi.Input<string>;
    /**
     * The name which should be used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The query of this Sentinel Scheduled Alert Rule.
     */
    readonly query?: pulumi.Input<string>;
    /**
     * The ISO 8601 timespan duration between two consecutive queries. Defaults to `PT5H`.
     */
    readonly queryFrequency?: pulumi.Input<string>;
    /**
     * The ISO 8601 timespan duration, which determine the time period of the data covered by the query. For example, it can query the past 10 minutes of data, or the past 6 hours of data. Defaults to `PT5H`.
     */
    readonly queryPeriod?: pulumi.Input<string>;
    /**
     * The alert severity of this Sentinel Scheduled Alert Rule. Possible values are `High`, `Medium`, `Low` and `Informational`.
     */
    readonly severity?: pulumi.Input<string>;
    /**
     * If `suppressionEnabled` is `true`, this is ISO 8601 timespan duration, which specifies the amount of time the query should stop running after alert is generated. Defaults to `PT5H`.
     */
    readonly suppressionDuration?: pulumi.Input<string>;
    /**
     * Should the Sentinel Scheduled Alert Rulea stop running query after alert is generated? Defaults to `false`.
     */
    readonly suppressionEnabled?: pulumi.Input<boolean>;
    /**
     * A list of categories of attacks by which to classify the rule. Possible values are `Collection`, `CommandAndControl`, `CredentialAccess`, `DefenseEvasion`, `Discovery`, `Execution`, `Exfiltration`, `Impact`, `InitialAccess`, `LateralMovement`, `Persistence` and `PrivilegeEscalation`.
     */
    readonly tactics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The alert trigger operator, combined with `triggerThreshold`, setting alert threshold of this Sentinel Scheduled Alert Rule. Possible values are `Equal`, `GreaterThan`, `LessThan`, `NotEqual`.
     */
    readonly triggerOperator?: pulumi.Input<string>;
    /**
     * The baseline number of query results generated, combined with `triggerOperator`, setting alert threshold of this Sentinel Scheduled Alert Rule.
     */
    readonly triggerThreshold?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AlertRuleScheduled resource.
 */
export interface AlertRuleScheduledArgs {
    /**
     * The description of this Sentinel Scheduled Alert Rule.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The friendly name of this Sentinel Scheduled Alert Rule.
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * Should the Sentinel Scheduled Alert Rule be enabled? Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The ID of the Log Analytics Workspace this Sentinel Scheduled Alert Rule belongs to. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
     */
    readonly logAnalyticsWorkspaceId: pulumi.Input<string>;
    /**
     * The name which should be used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The query of this Sentinel Scheduled Alert Rule.
     */
    readonly query: pulumi.Input<string>;
    /**
     * The ISO 8601 timespan duration between two consecutive queries. Defaults to `PT5H`.
     */
    readonly queryFrequency?: pulumi.Input<string>;
    /**
     * The ISO 8601 timespan duration, which determine the time period of the data covered by the query. For example, it can query the past 10 minutes of data, or the past 6 hours of data. Defaults to `PT5H`.
     */
    readonly queryPeriod?: pulumi.Input<string>;
    /**
     * The alert severity of this Sentinel Scheduled Alert Rule. Possible values are `High`, `Medium`, `Low` and `Informational`.
     */
    readonly severity: pulumi.Input<string>;
    /**
     * If `suppressionEnabled` is `true`, this is ISO 8601 timespan duration, which specifies the amount of time the query should stop running after alert is generated. Defaults to `PT5H`.
     */
    readonly suppressionDuration?: pulumi.Input<string>;
    /**
     * Should the Sentinel Scheduled Alert Rulea stop running query after alert is generated? Defaults to `false`.
     */
    readonly suppressionEnabled?: pulumi.Input<boolean>;
    /**
     * A list of categories of attacks by which to classify the rule. Possible values are `Collection`, `CommandAndControl`, `CredentialAccess`, `DefenseEvasion`, `Discovery`, `Execution`, `Exfiltration`, `Impact`, `InitialAccess`, `LateralMovement`, `Persistence` and `PrivilegeEscalation`.
     */
    readonly tactics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The alert trigger operator, combined with `triggerThreshold`, setting alert threshold of this Sentinel Scheduled Alert Rule. Possible values are `Equal`, `GreaterThan`, `LessThan`, `NotEqual`.
     */
    readonly triggerOperator?: pulumi.Input<string>;
    /**
     * The baseline number of query results generated, combined with `triggerOperator`, setting alert threshold of this Sentinel Scheduled Alert Rule.
     */
    readonly triggerThreshold?: pulumi.Input<number>;
}

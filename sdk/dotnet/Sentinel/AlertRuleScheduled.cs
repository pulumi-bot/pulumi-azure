// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Sentinel
{
    /// <summary>
    /// Manages a Sentinel Scheduled Alert Rule.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var exampleAnalyticsWorkspace = new Azure.OperationalInsights.AnalyticsWorkspace("exampleAnalyticsWorkspace", new Azure.OperationalInsights.AnalyticsWorkspaceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Sku = "pergb2018",
    ///         });
    ///         var exampleAlertRuleScheduled = new Azure.Sentinel.AlertRuleScheduled("exampleAlertRuleScheduled", new Azure.Sentinel.AlertRuleScheduledArgs
    ///         {
    ///             LogAnalyticsWorkspaceId = exampleAnalyticsWorkspace.Id,
    ///             DisplayName = "example",
    ///             Severity = "High",
    ///             Query = @"AzureActivity |
    ///   where OperationName == ""Create or Update Virtual Machine"" or OperationName ==""Create Deployment"" |
    ///   where ActivityStatus == ""Succeeded"" |
    ///   make-series dcount(ResourceId) default=0 on EventSubmissionTimestamp in range(ago(7d), now(), 1d) by Caller
    /// ",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Sentinel Scheduled Alert Rules can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class AlertRuleScheduled : Pulumi.CustomResource
    {
        /// <summary>
        /// The description of this Sentinel Scheduled Alert Rule.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The friendly name of this Sentinel Scheduled Alert Rule.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Should the Sentinel Scheduled Alert Rule be enabled? Defaults to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The ID of the Log Analytics Workspace this Sentinel Scheduled Alert Rule belongs to. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        /// </summary>
        [Output("logAnalyticsWorkspaceId")]
        public Output<string> LogAnalyticsWorkspaceId { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The query of this Sentinel Scheduled Alert Rule.
        /// </summary>
        [Output("query")]
        public Output<string> Query { get; private set; } = null!;

        /// <summary>
        /// The ISO 8601 timespan duration between two consecutive queries. Defaults to `PT5H`.
        /// </summary>
        [Output("queryFrequency")]
        public Output<string?> QueryFrequency { get; private set; } = null!;

        /// <summary>
        /// The ISO 8601 timespan duration, which determine the time period of the data covered by the query. For example, it can query the past 10 minutes of data, or the past 6 hours of data. Defaults to `PT5H`.
        /// </summary>
        [Output("queryPeriod")]
        public Output<string?> QueryPeriod { get; private set; } = null!;

        /// <summary>
        /// The alert severity of this Sentinel Scheduled Alert Rule. Possible values are `High`, `Medium`, `Low` and `Informational`.
        /// </summary>
        [Output("severity")]
        public Output<string> Severity { get; private set; } = null!;

        /// <summary>
        /// If `suppression_enabled` is `true`, this is ISO 8601 timespan duration, which specifies the amount of time the query should stop running after alert is generated. Defaults to `PT5H`.
        /// </summary>
        [Output("suppressionDuration")]
        public Output<string?> SuppressionDuration { get; private set; } = null!;

        /// <summary>
        /// Should the Sentinel Scheduled Alert Rulea stop running query after alert is generated? Defaults to `false`.
        /// </summary>
        [Output("suppressionEnabled")]
        public Output<bool?> SuppressionEnabled { get; private set; } = null!;

        /// <summary>
        /// A list of categories of attacks by which to classify the rule. Possible values are `Collection`, `CommandAndControl`, `CredentialAccess`, `DefenseEvasion`, `Discovery`, `Execution`, `Exfiltration`, `Impact`, `InitialAccess`, `LateralMovement`, `Persistence` and `PrivilegeEscalation`.
        /// </summary>
        [Output("tactics")]
        public Output<ImmutableArray<string>> Tactics { get; private set; } = null!;

        /// <summary>
        /// The alert trigger operator, combined with `trigger_threshold`, setting alert threshold of this Sentinel Scheduled Alert Rule. Possible values are `Equal`, `GreaterThan`, `LessThan`, `NotEqual`.
        /// </summary>
        [Output("triggerOperator")]
        public Output<string?> TriggerOperator { get; private set; } = null!;

        /// <summary>
        /// The baseline number of query results generated, combined with `trigger_operator`, setting alert threshold of this Sentinel Scheduled Alert Rule.
        /// </summary>
        [Output("triggerThreshold")]
        public Output<int?> TriggerThreshold { get; private set; } = null!;


        /// <summary>
        /// Create a AlertRuleScheduled resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AlertRuleScheduled(string name, AlertRuleScheduledArgs args, CustomResourceOptions? options = null)
            : base("azure:sentinel/alertRuleScheduled:AlertRuleScheduled", name, args ?? new AlertRuleScheduledArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AlertRuleScheduled(string name, Input<string> id, AlertRuleScheduledState? state = null, CustomResourceOptions? options = null)
            : base("azure:sentinel/alertRuleScheduled:AlertRuleScheduled", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AlertRuleScheduled resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AlertRuleScheduled Get(string name, Input<string> id, AlertRuleScheduledState? state = null, CustomResourceOptions? options = null)
        {
            return new AlertRuleScheduled(name, id, state, options);
        }
    }

    public sealed class AlertRuleScheduledArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this Sentinel Scheduled Alert Rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The friendly name of this Sentinel Scheduled Alert Rule.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Should the Sentinel Scheduled Alert Rule be enabled? Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The ID of the Log Analytics Workspace this Sentinel Scheduled Alert Rule belongs to. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        /// </summary>
        [Input("logAnalyticsWorkspaceId", required: true)]
        public Input<string> LogAnalyticsWorkspaceId { get; set; } = null!;

        /// <summary>
        /// The name which should be used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The query of this Sentinel Scheduled Alert Rule.
        /// </summary>
        [Input("query", required: true)]
        public Input<string> Query { get; set; } = null!;

        /// <summary>
        /// The ISO 8601 timespan duration between two consecutive queries. Defaults to `PT5H`.
        /// </summary>
        [Input("queryFrequency")]
        public Input<string>? QueryFrequency { get; set; }

        /// <summary>
        /// The ISO 8601 timespan duration, which determine the time period of the data covered by the query. For example, it can query the past 10 minutes of data, or the past 6 hours of data. Defaults to `PT5H`.
        /// </summary>
        [Input("queryPeriod")]
        public Input<string>? QueryPeriod { get; set; }

        /// <summary>
        /// The alert severity of this Sentinel Scheduled Alert Rule. Possible values are `High`, `Medium`, `Low` and `Informational`.
        /// </summary>
        [Input("severity", required: true)]
        public Input<string> Severity { get; set; } = null!;

        /// <summary>
        /// If `suppression_enabled` is `true`, this is ISO 8601 timespan duration, which specifies the amount of time the query should stop running after alert is generated. Defaults to `PT5H`.
        /// </summary>
        [Input("suppressionDuration")]
        public Input<string>? SuppressionDuration { get; set; }

        /// <summary>
        /// Should the Sentinel Scheduled Alert Rulea stop running query after alert is generated? Defaults to `false`.
        /// </summary>
        [Input("suppressionEnabled")]
        public Input<bool>? SuppressionEnabled { get; set; }

        [Input("tactics")]
        private InputList<string>? _tactics;

        /// <summary>
        /// A list of categories of attacks by which to classify the rule. Possible values are `Collection`, `CommandAndControl`, `CredentialAccess`, `DefenseEvasion`, `Discovery`, `Execution`, `Exfiltration`, `Impact`, `InitialAccess`, `LateralMovement`, `Persistence` and `PrivilegeEscalation`.
        /// </summary>
        public InputList<string> Tactics
        {
            get => _tactics ?? (_tactics = new InputList<string>());
            set => _tactics = value;
        }

        /// <summary>
        /// The alert trigger operator, combined with `trigger_threshold`, setting alert threshold of this Sentinel Scheduled Alert Rule. Possible values are `Equal`, `GreaterThan`, `LessThan`, `NotEqual`.
        /// </summary>
        [Input("triggerOperator")]
        public Input<string>? TriggerOperator { get; set; }

        /// <summary>
        /// The baseline number of query results generated, combined with `trigger_operator`, setting alert threshold of this Sentinel Scheduled Alert Rule.
        /// </summary>
        [Input("triggerThreshold")]
        public Input<int>? TriggerThreshold { get; set; }

        public AlertRuleScheduledArgs()
        {
        }
    }

    public sealed class AlertRuleScheduledState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this Sentinel Scheduled Alert Rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The friendly name of this Sentinel Scheduled Alert Rule.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Should the Sentinel Scheduled Alert Rule be enabled? Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The ID of the Log Analytics Workspace this Sentinel Scheduled Alert Rule belongs to. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        /// </summary>
        [Input("logAnalyticsWorkspaceId")]
        public Input<string>? LogAnalyticsWorkspaceId { get; set; }

        /// <summary>
        /// The name which should be used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The query of this Sentinel Scheduled Alert Rule.
        /// </summary>
        [Input("query")]
        public Input<string>? Query { get; set; }

        /// <summary>
        /// The ISO 8601 timespan duration between two consecutive queries. Defaults to `PT5H`.
        /// </summary>
        [Input("queryFrequency")]
        public Input<string>? QueryFrequency { get; set; }

        /// <summary>
        /// The ISO 8601 timespan duration, which determine the time period of the data covered by the query. For example, it can query the past 10 minutes of data, or the past 6 hours of data. Defaults to `PT5H`.
        /// </summary>
        [Input("queryPeriod")]
        public Input<string>? QueryPeriod { get; set; }

        /// <summary>
        /// The alert severity of this Sentinel Scheduled Alert Rule. Possible values are `High`, `Medium`, `Low` and `Informational`.
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        /// <summary>
        /// If `suppression_enabled` is `true`, this is ISO 8601 timespan duration, which specifies the amount of time the query should stop running after alert is generated. Defaults to `PT5H`.
        /// </summary>
        [Input("suppressionDuration")]
        public Input<string>? SuppressionDuration { get; set; }

        /// <summary>
        /// Should the Sentinel Scheduled Alert Rulea stop running query after alert is generated? Defaults to `false`.
        /// </summary>
        [Input("suppressionEnabled")]
        public Input<bool>? SuppressionEnabled { get; set; }

        [Input("tactics")]
        private InputList<string>? _tactics;

        /// <summary>
        /// A list of categories of attacks by which to classify the rule. Possible values are `Collection`, `CommandAndControl`, `CredentialAccess`, `DefenseEvasion`, `Discovery`, `Execution`, `Exfiltration`, `Impact`, `InitialAccess`, `LateralMovement`, `Persistence` and `PrivilegeEscalation`.
        /// </summary>
        public InputList<string> Tactics
        {
            get => _tactics ?? (_tactics = new InputList<string>());
            set => _tactics = value;
        }

        /// <summary>
        /// The alert trigger operator, combined with `trigger_threshold`, setting alert threshold of this Sentinel Scheduled Alert Rule. Possible values are `Equal`, `GreaterThan`, `LessThan`, `NotEqual`.
        /// </summary>
        [Input("triggerOperator")]
        public Input<string>? TriggerOperator { get; set; }

        /// <summary>
        /// The baseline number of query results generated, combined with `trigger_operator`, setting alert threshold of this Sentinel Scheduled Alert Rule.
        /// </summary>
        [Input("triggerThreshold")]
        public Input<int>? TriggerThreshold { get; set; }

        public AlertRuleScheduledState()
        {
        }
    }
}

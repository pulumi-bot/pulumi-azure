// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring
{
    /// <summary>
    /// Manages a Diagnostic Setting for an existing Resource.
    /// 
    /// ## Example Usage
    /// 
    /// 
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
    ///         var exampleAccount = exampleResourceGroup.Name.Apply(name =&gt; Azure.Storage.GetAccount.InvokeAsync(new Azure.Storage.GetAccountArgs
    ///         {
    ///             Name = "examplestoracc",
    ///             ResourceGroupName = name,
    ///         }));
    ///         var exampleKeyVault = exampleResourceGroup.Name.Apply(name =&gt; Azure.KeyVault.GetKeyVault.InvokeAsync(new Azure.KeyVault.GetKeyVaultArgs
    ///         {
    ///             Name = "example-vault",
    ///             ResourceGroupName = name,
    ///         }));
    ///         var exampleDiagnosticSetting = new Azure.Monitoring.DiagnosticSetting("exampleDiagnosticSetting", new Azure.Monitoring.DiagnosticSettingArgs
    ///         {
    ///             Logs = 
    ///             {
    ///                 new Azure.Monitoring.Inputs.DiagnosticSettingLogArgs
    ///                 {
    ///                     Category = "AuditEvent",
    ///                     Enabled = false,
    ///                     RetentionPolicy = new Azure.Monitoring.Inputs.DiagnosticSettingLogRetentionPolicyArgs
    ///                     {
    ///                         Enabled = false,
    ///                     },
    ///                 },
    ///             },
    ///             Metrics = 
    ///             {
    ///                 new Azure.Monitoring.Inputs.DiagnosticSettingMetricArgs
    ///                 {
    ///                     Category = "AllMetrics",
    ///                     RetentionPolicy = new Azure.Monitoring.Inputs.DiagnosticSettingMetricRetentionPolicyArgs
    ///                     {
    ///                         Enabled = false,
    ///                     },
    ///                 },
    ///             },
    ///             StorageAccountId = exampleAccount.Apply(exampleAccount =&gt; exampleAccount.Id),
    ///             TargetResourceId = exampleKeyVault.Apply(exampleKeyVault =&gt; exampleKeyVault.Id),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class DiagnosticSetting : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the ID of an Event Hub Namespace Authorization Rule used to send Diagnostics Data. Changing this forces a new resource to be created.
        /// </summary>
        [Output("eventhubAuthorizationRuleId")]
        public Output<string?> EventhubAuthorizationRuleId { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Event Hub where Diagnostics Data should be sent. Changing this forces a new resource to be created.
        /// </summary>
        [Output("eventhubName")]
        public Output<string?> EventhubName { get; private set; } = null!;

        /// <summary>
        /// When set to 'Dedicated' logs sent to a Log Analytics workspace will go into resource specific tables, instead of the legacy AzureDiagnostics table.
        /// </summary>
        [Output("logAnalyticsDestinationType")]
        public Output<string?> LogAnalyticsDestinationType { get; private set; } = null!;

        /// <summary>
        /// Specifies the ID of a Log Analytics Workspace where Diagnostics Data should be sent. Changing this forces a new resource to be created.
        /// </summary>
        [Output("logAnalyticsWorkspaceId")]
        public Output<string?> LogAnalyticsWorkspaceId { get; private set; } = null!;

        /// <summary>
        /// One or more `log` blocks as defined below.
        /// </summary>
        [Output("logs")]
        public Output<ImmutableArray<Outputs.DiagnosticSettingLog>> Logs { get; private set; } = null!;

        /// <summary>
        /// One or more `metric` blocks as defined below.
        /// </summary>
        [Output("metrics")]
        public Output<ImmutableArray<Outputs.DiagnosticSettingMetric>> Metrics { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Diagnostic Setting. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// With this parameter you can specify a storage account which should be used to send the logs to. Parameter must be a valid Azure Resource ID. Changing this forces a new resource to be created.
        /// </summary>
        [Output("storageAccountId")]
        public Output<string?> StorageAccountId { get; private set; } = null!;

        /// <summary>
        /// The ID of an existing Resource on which to configure Diagnostic Settings. Changing this forces a new resource to be created.
        /// </summary>
        [Output("targetResourceId")]
        public Output<string> TargetResourceId { get; private set; } = null!;


        /// <summary>
        /// Create a DiagnosticSetting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DiagnosticSetting(string name, DiagnosticSettingArgs args, CustomResourceOptions? options = null)
            : base("azure:monitoring/diagnosticSetting:DiagnosticSetting", name, args ?? new DiagnosticSettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DiagnosticSetting(string name, Input<string> id, DiagnosticSettingState? state = null, CustomResourceOptions? options = null)
            : base("azure:monitoring/diagnosticSetting:DiagnosticSetting", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DiagnosticSetting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DiagnosticSetting Get(string name, Input<string> id, DiagnosticSettingState? state = null, CustomResourceOptions? options = null)
        {
            return new DiagnosticSetting(name, id, state, options);
        }
    }

    public sealed class DiagnosticSettingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the ID of an Event Hub Namespace Authorization Rule used to send Diagnostics Data. Changing this forces a new resource to be created.
        /// </summary>
        [Input("eventhubAuthorizationRuleId")]
        public Input<string>? EventhubAuthorizationRuleId { get; set; }

        /// <summary>
        /// Specifies the name of the Event Hub where Diagnostics Data should be sent. Changing this forces a new resource to be created.
        /// </summary>
        [Input("eventhubName")]
        public Input<string>? EventhubName { get; set; }

        /// <summary>
        /// When set to 'Dedicated' logs sent to a Log Analytics workspace will go into resource specific tables, instead of the legacy AzureDiagnostics table.
        /// </summary>
        [Input("logAnalyticsDestinationType")]
        public Input<string>? LogAnalyticsDestinationType { get; set; }

        /// <summary>
        /// Specifies the ID of a Log Analytics Workspace where Diagnostics Data should be sent. Changing this forces a new resource to be created.
        /// </summary>
        [Input("logAnalyticsWorkspaceId")]
        public Input<string>? LogAnalyticsWorkspaceId { get; set; }

        [Input("logs")]
        private InputList<Inputs.DiagnosticSettingLogArgs>? _logs;

        /// <summary>
        /// One or more `log` blocks as defined below.
        /// </summary>
        public InputList<Inputs.DiagnosticSettingLogArgs> Logs
        {
            get => _logs ?? (_logs = new InputList<Inputs.DiagnosticSettingLogArgs>());
            set => _logs = value;
        }

        [Input("metrics")]
        private InputList<Inputs.DiagnosticSettingMetricArgs>? _metrics;

        /// <summary>
        /// One or more `metric` blocks as defined below.
        /// </summary>
        public InputList<Inputs.DiagnosticSettingMetricArgs> Metrics
        {
            get => _metrics ?? (_metrics = new InputList<Inputs.DiagnosticSettingMetricArgs>());
            set => _metrics = value;
        }

        /// <summary>
        /// Specifies the name of the Diagnostic Setting. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// With this parameter you can specify a storage account which should be used to send the logs to. Parameter must be a valid Azure Resource ID. Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageAccountId")]
        public Input<string>? StorageAccountId { get; set; }

        /// <summary>
        /// The ID of an existing Resource on which to configure Diagnostic Settings. Changing this forces a new resource to be created.
        /// </summary>
        [Input("targetResourceId", required: true)]
        public Input<string> TargetResourceId { get; set; } = null!;

        public DiagnosticSettingArgs()
        {
        }
    }

    public sealed class DiagnosticSettingState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the ID of an Event Hub Namespace Authorization Rule used to send Diagnostics Data. Changing this forces a new resource to be created.
        /// </summary>
        [Input("eventhubAuthorizationRuleId")]
        public Input<string>? EventhubAuthorizationRuleId { get; set; }

        /// <summary>
        /// Specifies the name of the Event Hub where Diagnostics Data should be sent. Changing this forces a new resource to be created.
        /// </summary>
        [Input("eventhubName")]
        public Input<string>? EventhubName { get; set; }

        /// <summary>
        /// When set to 'Dedicated' logs sent to a Log Analytics workspace will go into resource specific tables, instead of the legacy AzureDiagnostics table.
        /// </summary>
        [Input("logAnalyticsDestinationType")]
        public Input<string>? LogAnalyticsDestinationType { get; set; }

        /// <summary>
        /// Specifies the ID of a Log Analytics Workspace where Diagnostics Data should be sent. Changing this forces a new resource to be created.
        /// </summary>
        [Input("logAnalyticsWorkspaceId")]
        public Input<string>? LogAnalyticsWorkspaceId { get; set; }

        [Input("logs")]
        private InputList<Inputs.DiagnosticSettingLogGetArgs>? _logs;

        /// <summary>
        /// One or more `log` blocks as defined below.
        /// </summary>
        public InputList<Inputs.DiagnosticSettingLogGetArgs> Logs
        {
            get => _logs ?? (_logs = new InputList<Inputs.DiagnosticSettingLogGetArgs>());
            set => _logs = value;
        }

        [Input("metrics")]
        private InputList<Inputs.DiagnosticSettingMetricGetArgs>? _metrics;

        /// <summary>
        /// One or more `metric` blocks as defined below.
        /// </summary>
        public InputList<Inputs.DiagnosticSettingMetricGetArgs> Metrics
        {
            get => _metrics ?? (_metrics = new InputList<Inputs.DiagnosticSettingMetricGetArgs>());
            set => _metrics = value;
        }

        /// <summary>
        /// Specifies the name of the Diagnostic Setting. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// With this parameter you can specify a storage account which should be used to send the logs to. Parameter must be a valid Azure Resource ID. Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageAccountId")]
        public Input<string>? StorageAccountId { get; set; }

        /// <summary>
        /// The ID of an existing Resource on which to configure Diagnostic Settings. Changing this forces a new resource to be created.
        /// </summary>
        [Input("targetResourceId")]
        public Input<string>? TargetResourceId { get; set; }

        public DiagnosticSettingState()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.LogicApps
{
    /// <summary>
    /// Manages a Recurrence Trigger within a Logic App Workflow
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
    ///             Location = "East US",
    ///         });
    ///         var exampleWorkflow = new Azure.LogicApps.Workflow("exampleWorkflow", new Azure.LogicApps.WorkflowArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var exampleTriggerRecurrence = new Azure.LogicApps.TriggerRecurrence("exampleTriggerRecurrence", new Azure.LogicApps.TriggerRecurrenceArgs
    ///         {
    ///             LogicAppId = exampleWorkflow.Id,
    ///             Frequency = "Day",
    ///             Interval = 1,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class TriggerRecurrence : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the Frequency at which this Trigger should be run. Possible values include `Month`, `Week`, `Day`, `Hour`, `Minute` and `Second`.
        /// </summary>
        [Output("frequency")]
        public Output<string> Frequency { get; private set; } = null!;

        /// <summary>
        /// Specifies interval used for the Frequency, for example a value of `4` for `interval` and `hour` for `frequency` would run the Trigger every 4 hours.
        /// </summary>
        [Output("interval")]
        public Output<int> Interval { get; private set; } = null!;

        /// <summary>
        /// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        /// </summary>
        [Output("logicAppId")]
        public Output<string> LogicAppId { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Recurrence Triggers to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the start date and time for this trigger in RFC3339 format: `2000-01-02T03:04:05Z`.
        /// </summary>
        [Output("startTime")]
        public Output<string?> StartTime { get; private set; } = null!;


        /// <summary>
        /// Create a TriggerRecurrence resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TriggerRecurrence(string name, TriggerRecurrenceArgs args, CustomResourceOptions? options = null)
            : base("azure:logicapps/triggerRecurrence:TriggerRecurrence", name, args ?? new TriggerRecurrenceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TriggerRecurrence(string name, Input<string> id, TriggerRecurrenceState? state = null, CustomResourceOptions? options = null)
            : base("azure:logicapps/triggerRecurrence:TriggerRecurrence", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TriggerRecurrence resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TriggerRecurrence Get(string name, Input<string> id, TriggerRecurrenceState? state = null, CustomResourceOptions? options = null)
        {
            return new TriggerRecurrence(name, id, state, options);
        }
    }

    public sealed class TriggerRecurrenceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the Frequency at which this Trigger should be run. Possible values include `Month`, `Week`, `Day`, `Hour`, `Minute` and `Second`.
        /// </summary>
        [Input("frequency", required: true)]
        public Input<string> Frequency { get; set; } = null!;

        /// <summary>
        /// Specifies interval used for the Frequency, for example a value of `4` for `interval` and `hour` for `frequency` would run the Trigger every 4 hours.
        /// </summary>
        [Input("interval", required: true)]
        public Input<int> Interval { get; set; } = null!;

        /// <summary>
        /// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        /// </summary>
        [Input("logicAppId", required: true)]
        public Input<string> LogicAppId { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Recurrence Triggers to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the start date and time for this trigger in RFC3339 format: `2000-01-02T03:04:05Z`.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public TriggerRecurrenceArgs()
        {
        }
    }

    public sealed class TriggerRecurrenceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the Frequency at which this Trigger should be run. Possible values include `Month`, `Week`, `Day`, `Hour`, `Minute` and `Second`.
        /// </summary>
        [Input("frequency")]
        public Input<string>? Frequency { get; set; }

        /// <summary>
        /// Specifies interval used for the Frequency, for example a value of `4` for `interval` and `hour` for `frequency` would run the Trigger every 4 hours.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        /// </summary>
        [Input("logicAppId")]
        public Input<string>? LogicAppId { get; set; }

        /// <summary>
        /// Specifies the name of the Recurrence Triggers to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the start date and time for this trigger in RFC3339 format: `2000-01-02T03:04:05Z`.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public TriggerRecurrenceState()
        {
        }
    }
}

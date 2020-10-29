// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DevTest
{
    /// <summary>
    /// Manages automated startup and shutdown schedules for Azure Dev Test Lab.
    /// </summary>
    public partial class Schedule : Pulumi.CustomResource
    {
        [Output("dailyRecurrence")]
        public Output<Outputs.ScheduleDailyRecurrence?> DailyRecurrence { get; private set; } = null!;

        [Output("hourlyRecurrence")]
        public Output<Outputs.ScheduleHourlyRecurrence?> HourlyRecurrence { get; private set; } = null!;

        /// <summary>
        /// The name of the dev test lab. Changing this forces a new resource to be created.
        /// </summary>
        [Output("labName")]
        public Output<string> LabName { get; private set; } = null!;

        /// <summary>
        /// The location where the schedule is created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the dev test lab schedule. Valid value for name depends on the `task_type`. For instance for task_type `LabVmsStartupTask` the name needs to be `LabVmAutoStart`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("notificationSettings")]
        public Output<Outputs.ScheduleNotificationSettings> NotificationSettings { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the dev test lab schedule. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The status of this schedule. Possible values are `Enabled` and `Disabled`. Defaults to `Disabled`.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The task type of the schedule. Possible values include `LabVmsShutdownTask` and `LabVmAutoStart`.
        /// </summary>
        [Output("taskType")]
        public Output<string> TaskType { get; private set; } = null!;

        /// <summary>
        /// The time zone ID (e.g. Pacific Standard time).
        /// </summary>
        [Output("timeZoneId")]
        public Output<string> TimeZoneId { get; private set; } = null!;

        [Output("weeklyRecurrence")]
        public Output<Outputs.ScheduleWeeklyRecurrence?> WeeklyRecurrence { get; private set; } = null!;


        /// <summary>
        /// Create a Schedule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Schedule(string name, ScheduleArgs args, CustomResourceOptions? options = null)
            : base("azure:devtest/schedule:Schedule", name, args ?? new ScheduleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Schedule(string name, Input<string> id, ScheduleState? state = null, CustomResourceOptions? options = null)
            : base("azure:devtest/schedule:Schedule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Schedule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Schedule Get(string name, Input<string> id, ScheduleState? state = null, CustomResourceOptions? options = null)
        {
            return new Schedule(name, id, state, options);
        }
    }

    public sealed class ScheduleArgs : Pulumi.ResourceArgs
    {
        [Input("dailyRecurrence")]
        public Input<Inputs.ScheduleDailyRecurrenceArgs>? DailyRecurrence { get; set; }

        [Input("hourlyRecurrence")]
        public Input<Inputs.ScheduleHourlyRecurrenceArgs>? HourlyRecurrence { get; set; }

        /// <summary>
        /// The name of the dev test lab. Changing this forces a new resource to be created.
        /// </summary>
        [Input("labName", required: true)]
        public Input<string> LabName { get; set; } = null!;

        /// <summary>
        /// The location where the schedule is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the dev test lab schedule. Valid value for name depends on the `task_type`. For instance for task_type `LabVmsStartupTask` the name needs to be `LabVmAutoStart`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationSettings", required: true)]
        public Input<Inputs.ScheduleNotificationSettingsArgs> NotificationSettings { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the dev test lab schedule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The status of this schedule. Possible values are `Enabled` and `Disabled`. Defaults to `Disabled`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The task type of the schedule. Possible values include `LabVmsShutdownTask` and `LabVmAutoStart`.
        /// </summary>
        [Input("taskType", required: true)]
        public Input<string> TaskType { get; set; } = null!;

        /// <summary>
        /// The time zone ID (e.g. Pacific Standard time).
        /// </summary>
        [Input("timeZoneId", required: true)]
        public Input<string> TimeZoneId { get; set; } = null!;

        [Input("weeklyRecurrence")]
        public Input<Inputs.ScheduleWeeklyRecurrenceArgs>? WeeklyRecurrence { get; set; }

        public ScheduleArgs()
        {
        }
    }

    public sealed class ScheduleState : Pulumi.ResourceArgs
    {
        [Input("dailyRecurrence")]
        public Input<Inputs.ScheduleDailyRecurrenceGetArgs>? DailyRecurrence { get; set; }

        [Input("hourlyRecurrence")]
        public Input<Inputs.ScheduleHourlyRecurrenceGetArgs>? HourlyRecurrence { get; set; }

        /// <summary>
        /// The name of the dev test lab. Changing this forces a new resource to be created.
        /// </summary>
        [Input("labName")]
        public Input<string>? LabName { get; set; }

        /// <summary>
        /// The location where the schedule is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the dev test lab schedule. Valid value for name depends on the `task_type`. For instance for task_type `LabVmsStartupTask` the name needs to be `LabVmAutoStart`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationSettings")]
        public Input<Inputs.ScheduleNotificationSettingsGetArgs>? NotificationSettings { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the dev test lab schedule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The status of this schedule. Possible values are `Enabled` and `Disabled`. Defaults to `Disabled`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The task type of the schedule. Possible values include `LabVmsShutdownTask` and `LabVmAutoStart`.
        /// </summary>
        [Input("taskType")]
        public Input<string>? TaskType { get; set; }

        /// <summary>
        /// The time zone ID (e.g. Pacific Standard time).
        /// </summary>
        [Input("timeZoneId")]
        public Input<string>? TimeZoneId { get; set; }

        [Input("weeklyRecurrence")]
        public Input<Inputs.ScheduleWeeklyRecurrenceGetArgs>? WeeklyRecurrence { get; set; }

        public ScheduleState()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DevTest.Outputs
{

    [OutputType]
    public sealed class GlobalVMShutdownScheduleNotificationSettings
    {
        /// <summary>
        /// Whether to enable pre-shutdown notifications. Possible values are `true` and `false`. Defaults to `false`
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// Time in minutes between 15 and 120 before a shutdown event at which a notification will be sent. Defaults to `30`.
        /// </summary>
        public readonly int? TimeInMinutes;
        /// <summary>
        /// The webhook URL to which the notification will be sent. Required if `enabled` is `true`. Optional otherwise.
        /// </summary>
        public readonly string? WebhookUrl;

        [OutputConstructor]
        private GlobalVMShutdownScheduleNotificationSettings(
            bool enabled,

            int? timeInMinutes,

            string? webhookUrl)
        {
            Enabled = enabled;
            TimeInMinutes = timeInMinutes;
            WebhookUrl = webhookUrl;
        }
    }
}

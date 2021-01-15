// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Sentinel.Outputs
{

    [OutputType]
    public sealed class AlertRuleScheduledIncidentConfigurationGrouping
    {
        /// <summary>
        /// Enable grouping incidents created from alerts triggered by this Sentinel Scheduled Alert Rule. Defaults to `true`.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The method used to group incidents. Possible values are `All`, `Custom` and `None`. Defaults to `None`.
        /// </summary>
        public readonly string? EntityMatchingMethod;
        /// <summary>
        /// A list of entity types to group by, only when the `entity_matching_method` is `Custom`. Possible values are `Account`, `Host`, `Url`, `Ip`.
        /// </summary>
        public readonly ImmutableArray<string> GroupBies;
        /// <summary>
        /// Limit the group to alerts created within the lookback duration (in ISO 8601 duration format). Defaults to `PT5M`.
        /// </summary>
        public readonly string? LookbackDuration;
        /// <summary>
        /// Whether to re-open closed matching incidents? Defaults to `false`.
        /// </summary>
        public readonly bool? ReopenClosedIncidents;

        [OutputConstructor]
        private AlertRuleScheduledIncidentConfigurationGrouping(
            bool? enabled,

            string? entityMatchingMethod,

            ImmutableArray<string> groupBies,

            string? lookbackDuration,

            bool? reopenClosedIncidents)
        {
            Enabled = enabled;
            EntityMatchingMethod = entityMatchingMethod;
            GroupBies = groupBies;
            LookbackDuration = lookbackDuration;
            ReopenClosedIncidents = reopenClosedIncidents;
        }
    }
}
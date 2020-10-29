// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ManagementGroups
{
    [Obsolete(@"azure.managementgroups.getManagementGroup has been deprecated in favor of azure.management.getGroup")]
    public static class GetManagementGroup
    {
        /// <summary>
        /// Use this data source to access information about an existing Management Group.
        /// </summary>
        public static Task<GetManagementGroupResult> InvokeAsync(GetManagementGroupArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetManagementGroupResult>("azure:managementgroups/getManagementGroup:getManagementGroup", args ?? new GetManagementGroupArgs(), options.WithVersion());
    }


    public sealed class GetManagementGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the display name of this Management Group.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        /// <summary>
        /// Specifies the name or UUID of this Management Group.
        /// </summary>
        [Input("groupId")]
        public string? GroupId { get; set; }

        /// <summary>
        /// Specifies the name or UUID of this Management Group.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetManagementGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetManagementGroupResult
    {
        public readonly string DisplayName;
        public readonly string GroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// The ID of any Parent Management Group.
        /// </summary>
        public readonly string ParentManagementGroupId;
        /// <summary>
        /// A list of Subscription IDs which are assigned to the Management Group.
        /// </summary>
        public readonly ImmutableArray<string> SubscriptionIds;

        [OutputConstructor]
        private GetManagementGroupResult(
            string displayName,

            string groupId,

            string id,

            string name,

            string parentManagementGroupId,

            ImmutableArray<string> subscriptionIds)
        {
            DisplayName = displayName;
            GroupId = groupId;
            Id = id;
            Name = name;
            ParentManagementGroupId = parentManagementGroupId;
            SubscriptionIds = subscriptionIds;
        }
    }
}

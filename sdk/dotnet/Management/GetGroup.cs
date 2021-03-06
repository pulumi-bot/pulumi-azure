// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Management
{
    public static class GetGroup
    {
        /// <summary>
        /// Use this data source to access information about an existing Management Group.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.Management.GetGroup.InvokeAsync(new Azure.Management.GetGroupArgs
        ///         {
        ///             Name = "00000000-0000-0000-0000-000000000000",
        ///         }));
        ///         this.DisplayName = example.Apply(example =&gt; example.DisplayName);
        ///     }
        /// 
        ///     [Output("displayName")]
        ///     public Output&lt;string&gt; DisplayName { get; set; }
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGroupResult> InvokeAsync(GetGroupArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("azure:management/getGroup:getGroup", args ?? new GetGroupArgs(), options.WithVersion());
    }


    public sealed class GetGroupArgs : Pulumi.InvokeArgs
    {
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

        public GetGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGroupResult
    {
        /// <summary>
        /// A friendly name for the Management Group.
        /// </summary>
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
        /// A list of Subscription ID's which are assigned to the Management Group.
        /// </summary>
        public readonly ImmutableArray<string> SubscriptionIds;

        [OutputConstructor]
        private GetGroupResult(
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

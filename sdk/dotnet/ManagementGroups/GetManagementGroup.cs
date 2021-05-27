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
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetManagementGroupResult> InvokeAsync(GetManagementGroupArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetManagementGroupResult>("azure:managementgroups/getManagementGroup:getManagementGroup", args ?? new GetManagementGroupArgs(), options.WithVersion());

        public static Output<GetManagementGroupResult> Invoke(GetManagementGroupOutputArgs? args = null, InvokeOptions? options = null)
        {
            args = args ?? new GetManagementGroupOutputArgs();
            return Pulumi.Output.All(
                args.DisplayName.Box(),
                args.GroupId.Box(),
                args.Name.Box()
            ).Apply(a => {
                    var args = new GetManagementGroupArgs();
                    a[0].Set(args, nameof(args.DisplayName));
                    a[1].Set(args, nameof(args.GroupId));
                    a[2].Set(args, nameof(args.Name));
                    return InvokeAsync(args, options);
            });
        }
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

    public sealed class GetManagementGroupOutputArgs
    {
        /// <summary>
        /// Specifies the display name of this Management Group.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Specifies the name or UUID of this Management Group.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// Specifies the name or UUID of this Management Group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetManagementGroupOutputArgs()
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

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ManagementGroups
{
    /// <summary>
    /// Manages a Management Group.
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
    ///         var current = Output.Create(Azure.Core.GetSubscription.InvokeAsync());
    ///         var exampleParent = new Azure.Management.Group("exampleParent", new Azure.Management.GroupArgs
    ///         {
    ///             DisplayName = "ParentGroup",
    ///             SubscriptionIds = 
    ///             {
    ///                 current.Apply(current =&gt; current.SubscriptionId),
    ///             },
    ///         });
    ///         var exampleChild = new Azure.Management.Group("exampleChild", new Azure.Management.GroupArgs
    ///         {
    ///             DisplayName = "ChildGroup",
    ///             ParentManagementGroupId = exampleParent.Id,
    ///             SubscriptionIds = 
    ///             {
    ///                 current.Apply(current =&gt; current.SubscriptionId),
    ///             },
    ///         });
    ///         // other subscription IDs can go here
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Management Groups can be imported using the `management group resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:managementgroups/managementGroup:ManagementGroup example /providers/Microsoft.Management/managementGroups/group1
    /// ```
    /// </summary>
    [Obsolete(@"azure.managementgroups.ManagementGroup has been deprecated in favor of azure.management.Group")]
    [AzureResourceType("azure:managementgroups/managementGroup:ManagementGroup")]
    public partial class ManagementGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// A friendly name for this Management Group. If not specified, this'll be the same as the `name`.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the Parent Management Group. Changing this forces a new resource to be created.
        /// </summary>
        [Output("parentManagementGroupId")]
        public Output<string> ParentManagementGroupId { get; private set; } = null!;

        /// <summary>
        /// A list of Subscription GUIDs which should be assigned to the Management Group.
        /// </summary>
        [Output("subscriptionIds")]
        public Output<ImmutableArray<string>> SubscriptionIds { get; private set; } = null!;


        /// <summary>
        /// Create a ManagementGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagementGroup(string name, ManagementGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("azure:managementgroups/managementGroup:ManagementGroup", name, args ?? new ManagementGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagementGroup(string name, Input<string> id, ManagementGroupState? state = null, CustomResourceOptions? options = null)
            : base("azure:managementgroups/managementGroup:ManagementGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ManagementGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagementGroup Get(string name, Input<string> id, ManagementGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ManagementGroup(name, id, state, options);
        }
    }

    public sealed class ManagementGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A friendly name for this Management Group. If not specified, this'll be the same as the `name`.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the Parent Management Group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("parentManagementGroupId")]
        public Input<string>? ParentManagementGroupId { get; set; }

        [Input("subscriptionIds")]
        private InputList<string>? _subscriptionIds;

        /// <summary>
        /// A list of Subscription GUIDs which should be assigned to the Management Group.
        /// </summary>
        public InputList<string> SubscriptionIds
        {
            get => _subscriptionIds ?? (_subscriptionIds = new InputList<string>());
            set => _subscriptionIds = value;
        }

        public ManagementGroupArgs()
        {
        }
    }

    public sealed class ManagementGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A friendly name for this Management Group. If not specified, this'll be the same as the `name`.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the Parent Management Group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("parentManagementGroupId")]
        public Input<string>? ParentManagementGroupId { get; set; }

        [Input("subscriptionIds")]
        private InputList<string>? _subscriptionIds;

        /// <summary>
        /// A list of Subscription GUIDs which should be assigned to the Management Group.
        /// </summary>
        public InputList<string> SubscriptionIds
        {
            get => _subscriptionIds ?? (_subscriptionIds = new InputList<string>());
            set => _subscriptionIds = value;
        }

        public ManagementGroupState()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement
{
    /// <summary>
    /// Manages an API Management User Assignment to a Group.
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
    ///         var exampleUser = Output.Create(Azure.ApiManagement.GetUser.InvokeAsync(new Azure.ApiManagement.GetUserArgs
    ///         {
    ///             UserId = "my-user",
    ///             ApiManagementName = "example-apim",
    ///             ResourceGroupName = "search-service",
    ///         }));
    ///         var exampleGroupUser = new Azure.ApiManagement.GroupUser("exampleGroupUser", new Azure.ApiManagement.GroupUserArgs
    ///         {
    ///             UserId = exampleUser.Apply(exampleUser =&gt; exampleUser.Id),
    ///             GroupName = "example-group",
    ///             ResourceGroupName = exampleUser.Apply(exampleUser =&gt; exampleUser.ResourceGroupName),
    ///             ApiManagementName = exampleUser.Apply(exampleUser =&gt; exampleUser.ApiManagementName),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class GroupUser : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Output("apiManagementName")]
        public Output<string> ApiManagementName { get; private set; } = null!;

        /// <summary>
        /// The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Output("groupName")]
        public Output<string> GroupName { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The ID of the API Management User which should be assigned to this API Management Group. Changing this forces a new resource to be created.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a GroupUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupUser(string name, GroupUserArgs args, CustomResourceOptions? options = null)
            : base("azure:apimanagement/groupUser:GroupUser", name, args ?? new GroupUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupUser(string name, Input<string> id, GroupUserState? state = null, CustomResourceOptions? options = null)
            : base("azure:apimanagement/groupUser:GroupUser", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupUser Get(string name, Input<string> id, GroupUserState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupUser(name, id, state, options);
        }
    }

    public sealed class GroupUserArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiManagementName", required: true)]
        public Input<string> ApiManagementName { get; set; } = null!;

        /// <summary>
        /// The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The ID of the API Management User which should be assigned to this API Management Group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public GroupUserArgs()
        {
        }
    }

    public sealed class GroupUserState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiManagementName")]
        public Input<string>? ApiManagementName { get; set; }

        /// <summary>
        /// The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        /// <summary>
        /// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The ID of the API Management User which should be assigned to this API Management Group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public GroupUserState()
        {
        }
    }
}

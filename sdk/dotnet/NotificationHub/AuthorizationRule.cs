// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.NotificationHub
{
    /// <summary>
    /// Manages an Authorization Rule associated with a Notification Hub within a Notification Hub Namespace.
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
    ///             Location = "Australia East",
    ///         });
    ///         var exampleNamespace = new Azure.NotificationHub.Namespace("exampleNamespace", new Azure.NotificationHub.NamespaceArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             NamespaceType = "NotificationHub",
    ///             SkuName = "Free",
    ///         });
    ///         var exampleHub = new Azure.NotificationHub.Hub("exampleHub", new Azure.NotificationHub.HubArgs
    ///         {
    ///             NamespaceName = exampleNamespace.Name,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///         });
    ///         var exampleAuthorizationRule = new Azure.NotificationHub.AuthorizationRule("exampleAuthorizationRule", new Azure.NotificationHub.AuthorizationRuleArgs
    ///         {
    ///             NotificationHubName = exampleHub.Name,
    ///             NamespaceName = exampleNamespace.Name,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Manage = true,
    ///             Send = true,
    ///             Listen = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class AuthorizationRule : Pulumi.CustomResource
    {
        /// <summary>
        /// Does this Authorization Rule have Listen access to the Notification Hub? Defaults to `false`.
        /// </summary>
        [Output("listen")]
        public Output<bool?> Listen { get; private set; } = null!;

        /// <summary>
        /// Does this Authorization Rule have Manage access to the Notification Hub? Defaults to `false`.
        /// </summary>
        [Output("manage")]
        public Output<bool?> Manage { get; private set; } = null!;

        /// <summary>
        /// The name to use for this Authorization Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the Notification Hub Namespace in which the Notification Hub exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("namespaceName")]
        public Output<string> NamespaceName { get; private set; } = null!;

        /// <summary>
        /// The name of the Notification Hub for which the Authorization Rule should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("notificationHubName")]
        public Output<string> NotificationHubName { get; private set; } = null!;

        /// <summary>
        /// The Primary Access Key associated with this Authorization Rule.
        /// </summary>
        [Output("primaryAccessKey")]
        public Output<string> PrimaryAccessKey { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Notification Hub Namespace exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The Secondary Access Key associated with this Authorization Rule.
        /// </summary>
        [Output("secondaryAccessKey")]
        public Output<string> SecondaryAccessKey { get; private set; } = null!;

        /// <summary>
        /// Does this Authorization Rule have Send access to the Notification Hub? Defaults to `false`.
        /// </summary>
        [Output("send")]
        public Output<bool?> Send { get; private set; } = null!;


        /// <summary>
        /// Create a AuthorizationRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthorizationRule(string name, AuthorizationRuleArgs args, CustomResourceOptions? options = null)
            : base("azure:notificationhub/authorizationRule:AuthorizationRule", name, args ?? new AuthorizationRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthorizationRule(string name, Input<string> id, AuthorizationRuleState? state = null, CustomResourceOptions? options = null)
            : base("azure:notificationhub/authorizationRule:AuthorizationRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuthorizationRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthorizationRule Get(string name, Input<string> id, AuthorizationRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthorizationRule(name, id, state, options);
        }
    }

    public sealed class AuthorizationRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Does this Authorization Rule have Listen access to the Notification Hub? Defaults to `false`.
        /// </summary>
        [Input("listen")]
        public Input<bool>? Listen { get; set; }

        /// <summary>
        /// Does this Authorization Rule have Manage access to the Notification Hub? Defaults to `false`.
        /// </summary>
        [Input("manage")]
        public Input<bool>? Manage { get; set; }

        /// <summary>
        /// The name to use for this Authorization Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Notification Hub Namespace in which the Notification Hub exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// The name of the Notification Hub for which the Authorization Rule should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("notificationHubName", required: true)]
        public Input<string> NotificationHubName { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Notification Hub Namespace exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Does this Authorization Rule have Send access to the Notification Hub? Defaults to `false`.
        /// </summary>
        [Input("send")]
        public Input<bool>? Send { get; set; }

        public AuthorizationRuleArgs()
        {
        }
    }

    public sealed class AuthorizationRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Does this Authorization Rule have Listen access to the Notification Hub? Defaults to `false`.
        /// </summary>
        [Input("listen")]
        public Input<bool>? Listen { get; set; }

        /// <summary>
        /// Does this Authorization Rule have Manage access to the Notification Hub? Defaults to `false`.
        /// </summary>
        [Input("manage")]
        public Input<bool>? Manage { get; set; }

        /// <summary>
        /// The name to use for this Authorization Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Notification Hub Namespace in which the Notification Hub exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("namespaceName")]
        public Input<string>? NamespaceName { get; set; }

        /// <summary>
        /// The name of the Notification Hub for which the Authorization Rule should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("notificationHubName")]
        public Input<string>? NotificationHubName { get; set; }

        /// <summary>
        /// The Primary Access Key associated with this Authorization Rule.
        /// </summary>
        [Input("primaryAccessKey")]
        public Input<string>? PrimaryAccessKey { get; set; }

        /// <summary>
        /// The name of the Resource Group in which the Notification Hub Namespace exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The Secondary Access Key associated with this Authorization Rule.
        /// </summary>
        [Input("secondaryAccessKey")]
        public Input<string>? SecondaryAccessKey { get; set; }

        /// <summary>
        /// Does this Authorization Rule have Send access to the Notification Hub? Defaults to `false`.
        /// </summary>
        [Input("send")]
        public Input<bool>? Send { get; set; }

        public AuthorizationRuleState()
        {
        }
    }
}

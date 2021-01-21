// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Policy
{
    /// <summary>
    /// Manages a policy rule definition on a management group or your provider subscription.
    /// 
    /// Policy definitions do not take effect until they are assigned to a scope using a Policy Assignment.
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
    ///         var policy = new Azure.Policy.Definition("policy", new Azure.Policy.DefinitionArgs
    ///         {
    ///             DisplayName = "acceptance test policy definition",
    ///             Metadata = @"    {
    ///     ""category"": ""General""
    ///     }
    /// 
    /// 
    /// ",
    ///             Mode = "Indexed",
    ///             Parameters = @"	{
    ///     ""allowedLocations"": {
    ///       ""type"": ""Array"",
    ///       ""metadata"": {
    ///         ""description"": ""The list of allowed locations for resources."",
    ///         ""displayName"": ""Allowed locations"",
    ///         ""strongType"": ""location""
    ///       }
    ///     }
    ///   }
    /// 
    /// ",
    ///             PolicyRule = @"	{
    ///     ""if"": {
    ///       ""not"": {
    ///         ""field"": ""location"",
    ///         ""in"": ""[parameters('allowedLocations')]""
    ///       }
    ///     },
    ///     ""then"": {
    ///       ""effect"": ""audit""
    ///     }
    ///   }
    /// 
    /// ",
    ///             PolicyType = "Custom",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Policy Definitions can be imported using the `policy name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:policy/definition:Definition examplePolicy /subscriptions/&lt;SUBSCRIPTION_ID&gt;/providers/Microsoft.Authorization/policyDefinitions/&lt;POLICY_NAME&gt;
    /// ```
    /// 
    ///  or
    /// 
    /// ```sh
    ///  $ pulumi import azure:policy/definition:Definition examplePolicy /providers/Microsoft.Management/managementgroups/&lt;MANGAGEMENT_GROUP_ID&gt;/providers/Microsoft.Authorization/policyDefinitions/&lt;POLICY_NAME&gt;
    /// ```
    /// </summary>
    [AzureResourceType("azure:policy/definition:Definition")]
    public partial class Definition : Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the policy definition.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The display name of the policy definition.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
        /// </summary>
        [Output("managementGroupId")]
        public Output<string> ManagementGroupId { get; private set; } = null!;

        /// <summary>
        /// The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
        /// </summary>
        [Output("managementGroupName")]
        public Output<string> ManagementGroupName { get; private set; } = null!;

        /// <summary>
        /// The metadata for the policy definition. This
        /// is a JSON string representing additional metadata that should be stored
        /// with the policy definition.
        /// </summary>
        [Output("metadata")]
        public Output<string> Metadata { get; private set; } = null!;

        /// <summary>
        /// The policy mode that allows you to specify which resource
        /// types will be evaluated. Possible values are `All`, `Indexed`, `Microsoft.ContainerService.Data`, `Microsoft.CustomerLockbox.Data`, `Microsoft.DataCatalog.Data`, `Microsoft.KeyVault.Data`, `Microsoft.Kubernetes.Data`, `Microsoft.MachineLearningServices.Data`, `Microsoft.Network.Data` and `Microsoft.Synapse.Data`.
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// The name of the policy definition. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Parameters for the policy definition. This field
        /// is a JSON string that allows you to parameterize your policy definition.
        /// </summary>
        [Output("parameters")]
        public Output<string?> Parameters { get; private set; } = null!;

        /// <summary>
        /// The policy rule for the policy definition. This
        /// is a JSON string representing the rule that contains an if and
        /// a then block.
        /// </summary>
        [Output("policyRule")]
        public Output<string?> PolicyRule { get; private set; } = null!;

        /// <summary>
        /// The policy type. Possible values are `BuiltIn`, `Custom` and `NotSpecified`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("policyType")]
        public Output<string> PolicyType { get; private set; } = null!;


        /// <summary>
        /// Create a Definition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Definition(string name, DefinitionArgs args, CustomResourceOptions? options = null)
            : base("azure:policy/definition:Definition", name, args ?? new DefinitionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Definition(string name, Input<string> id, DefinitionState? state = null, CustomResourceOptions? options = null)
            : base("azure:policy/definition:Definition", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Definition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Definition Get(string name, Input<string> id, DefinitionState? state = null, CustomResourceOptions? options = null)
        {
            return new Definition(name, id, state, options);
        }
    }

    public sealed class DefinitionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the policy definition.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the policy definition.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
        /// </summary>
        [Input("managementGroupId")]
        public Input<string>? ManagementGroupId { get; set; }

        /// <summary>
        /// The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
        /// </summary>
        [Input("managementGroupName")]
        public Input<string>? ManagementGroupName { get; set; }

        /// <summary>
        /// The metadata for the policy definition. This
        /// is a JSON string representing additional metadata that should be stored
        /// with the policy definition.
        /// </summary>
        [Input("metadata")]
        public Input<string>? Metadata { get; set; }

        /// <summary>
        /// The policy mode that allows you to specify which resource
        /// types will be evaluated. Possible values are `All`, `Indexed`, `Microsoft.ContainerService.Data`, `Microsoft.CustomerLockbox.Data`, `Microsoft.DataCatalog.Data`, `Microsoft.KeyVault.Data`, `Microsoft.Kubernetes.Data`, `Microsoft.MachineLearningServices.Data`, `Microsoft.Network.Data` and `Microsoft.Synapse.Data`.
        /// </summary>
        [Input("mode", required: true)]
        public Input<string> Mode { get; set; } = null!;

        /// <summary>
        /// The name of the policy definition. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Parameters for the policy definition. This field
        /// is a JSON string that allows you to parameterize your policy definition.
        /// </summary>
        [Input("parameters")]
        public Input<string>? Parameters { get; set; }

        /// <summary>
        /// The policy rule for the policy definition. This
        /// is a JSON string representing the rule that contains an if and
        /// a then block.
        /// </summary>
        [Input("policyRule")]
        public Input<string>? PolicyRule { get; set; }

        /// <summary>
        /// The policy type. Possible values are `BuiltIn`, `Custom` and `NotSpecified`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("policyType", required: true)]
        public Input<string> PolicyType { get; set; } = null!;

        public DefinitionArgs()
        {
        }
    }

    public sealed class DefinitionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the policy definition.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the policy definition.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
        /// </summary>
        [Input("managementGroupId")]
        public Input<string>? ManagementGroupId { get; set; }

        /// <summary>
        /// The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
        /// </summary>
        [Input("managementGroupName")]
        public Input<string>? ManagementGroupName { get; set; }

        /// <summary>
        /// The metadata for the policy definition. This
        /// is a JSON string representing additional metadata that should be stored
        /// with the policy definition.
        /// </summary>
        [Input("metadata")]
        public Input<string>? Metadata { get; set; }

        /// <summary>
        /// The policy mode that allows you to specify which resource
        /// types will be evaluated. Possible values are `All`, `Indexed`, `Microsoft.ContainerService.Data`, `Microsoft.CustomerLockbox.Data`, `Microsoft.DataCatalog.Data`, `Microsoft.KeyVault.Data`, `Microsoft.Kubernetes.Data`, `Microsoft.MachineLearningServices.Data`, `Microsoft.Network.Data` and `Microsoft.Synapse.Data`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// The name of the policy definition. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Parameters for the policy definition. This field
        /// is a JSON string that allows you to parameterize your policy definition.
        /// </summary>
        [Input("parameters")]
        public Input<string>? Parameters { get; set; }

        /// <summary>
        /// The policy rule for the policy definition. This
        /// is a JSON string representing the rule that contains an if and
        /// a then block.
        /// </summary>
        [Input("policyRule")]
        public Input<string>? PolicyRule { get; set; }

        /// <summary>
        /// The policy type. Possible values are `BuiltIn`, `Custom` and `NotSpecified`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        public DefinitionState()
        {
        }
    }
}

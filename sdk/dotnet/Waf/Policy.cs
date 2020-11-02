// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Waf
{
    /// <summary>
    /// Manages a Azure Web Application Firewall Policy instance.
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
    ///             Location = "West US 2",
    ///         });
    ///         var examplePolicy = new Azure.Waf.Policy("examplePolicy", new Azure.Waf.PolicyArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             CustomRules = 
    ///             {
    ///                 new Azure.Waf.Inputs.PolicyCustomRuleArgs
    ///                 {
    ///                     Name = "Rule1",
    ///                     Priority = 1,
    ///                     RuleType = "MatchRule",
    ///                     MatchConditions = 
    ///                     {
    ///                         new Azure.Waf.Inputs.PolicyCustomRuleMatchConditionArgs
    ///                         {
    ///                             MatchVariables = 
    ///                             {
    ///                                 new Azure.Waf.Inputs.PolicyCustomRuleMatchConditionMatchVariableArgs
    ///                                 {
    ///                                     VariableName = "RemoteAddr",
    ///                                 },
    ///                             },
    ///                             Operator = "IPMatch",
    ///                             NegationCondition = false,
    ///                             MatchValues = 
    ///                             {
    ///                                 "192.168.1.0/24",
    ///                                 "10.0.0.0/24",
    ///                             },
    ///                         },
    ///                     },
    ///                     Action = "Block",
    ///                 },
    ///                 new Azure.Waf.Inputs.PolicyCustomRuleArgs
    ///                 {
    ///                     Name = "Rule2",
    ///                     Priority = 2,
    ///                     RuleType = "MatchRule",
    ///                     MatchConditions = 
    ///                     {
    ///                         new Azure.Waf.Inputs.PolicyCustomRuleMatchConditionArgs
    ///                         {
    ///                             MatchVariables = 
    ///                             {
    ///                                 new Azure.Waf.Inputs.PolicyCustomRuleMatchConditionMatchVariableArgs
    ///                                 {
    ///                                     VariableName = "RemoteAddr",
    ///                                 },
    ///                             },
    ///                             Operator = "IPMatch",
    ///                             NegationCondition = false,
    ///                             MatchValues = 
    ///                             {
    ///                                 "192.168.1.0/24",
    ///                             },
    ///                         },
    ///                         new Azure.Waf.Inputs.PolicyCustomRuleMatchConditionArgs
    ///                         {
    ///                             MatchVariables = 
    ///                             {
    ///                                 new Azure.Waf.Inputs.PolicyCustomRuleMatchConditionMatchVariableArgs
    ///                                 {
    ///                                     VariableName = "RequestHeaders",
    ///                                     Selector = "UserAgent",
    ///                                 },
    ///                             },
    ///                             Operator = "Contains",
    ///                             NegationCondition = false,
    ///                             MatchValues = 
    ///                             {
    ///                                 "Windows",
    ///                             },
    ///                         },
    ///                     },
    ///                     Action = "Block",
    ///                 },
    ///             },
    ///             PolicySettings = new Azure.Waf.Inputs.PolicyPolicySettingsArgs
    ///             {
    ///                 Enabled = true,
    ///                 Mode = "Prevention",
    ///                 RequestBodyCheck = true,
    ///                 FileUploadLimitInMb = 100,
    ///                 MaxRequestBodySizeInKb = 128,
    ///             },
    ///             ManagedRules = new Azure.Waf.Inputs.PolicyManagedRulesArgs
    ///             {
    ///                 Exclusions = 
    ///                 {
    ///                     new Azure.Waf.Inputs.PolicyManagedRulesExclusionArgs
    ///                     {
    ///                         MatchVariable = "RequestHeaderNames",
    ///                         Selector = "x-company-secret-header",
    ///                         SelectorMatchOperator = "Equals",
    ///                     },
    ///                     new Azure.Waf.Inputs.PolicyManagedRulesExclusionArgs
    ///                     {
    ///                         MatchVariable = "RequestCookieNames",
    ///                         Selector = "too-tasty",
    ///                         SelectorMatchOperator = "EndsWith",
    ///                     },
    ///                 },
    ///                 ManagedRuleSets = 
    ///                 {
    ///                     new Azure.Waf.Inputs.PolicyManagedRulesManagedRuleSetArgs
    ///                     {
    ///                         Type = "OWASP",
    ///                         Version = "3.1",
    ///                         RuleGroupOverrides = 
    ///                         {
    ///                             new Azure.Waf.Inputs.PolicyManagedRulesManagedRuleSetRuleGroupOverrideArgs
    ///                             {
    ///                                 RuleGroupName = "REQUEST-920-PROTOCOL-ENFORCEMENT",
    ///                                 DisabledRules = 
    ///                                 {
    ///                                     "920300",
    ///                                     "920440",
    ///                                 },
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Web Application Firewall Policy can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class Policy : Pulumi.CustomResource
    {
        /// <summary>
        /// One or more `custom_rules` blocks as defined below.
        /// </summary>
        [Output("customRules")]
        public Output<ImmutableArray<Outputs.PolicyCustomRule>> CustomRules { get; private set; } = null!;

        /// <summary>
        /// Resource location. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// A `managed_rules` blocks as defined below.
        /// </summary>
        [Output("managedRules")]
        public Output<Outputs.PolicyManagedRules> ManagedRules { get; private set; } = null!;

        /// <summary>
        /// The name of the policy. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A `policy_settings` block as defined below.
        /// </summary>
        [Output("policySettings")]
        public Output<Outputs.PolicyPolicySettings?> PolicySettings { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the Web Application Firewall Policy.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Policy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Policy(string name, PolicyArgs args, CustomResourceOptions? options = null)
            : base("azure:waf/policy:Policy", name, args ?? new PolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Policy(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
            : base("azure:waf/policy:Policy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Policy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Policy Get(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new Policy(name, id, state, options);
        }
    }

    public sealed class PolicyArgs : Pulumi.ResourceArgs
    {
        [Input("customRules")]
        private InputList<Inputs.PolicyCustomRuleArgs>? _customRules;

        /// <summary>
        /// One or more `custom_rules` blocks as defined below.
        /// </summary>
        public InputList<Inputs.PolicyCustomRuleArgs> CustomRules
        {
            get => _customRules ?? (_customRules = new InputList<Inputs.PolicyCustomRuleArgs>());
            set => _customRules = value;
        }

        /// <summary>
        /// Resource location. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A `managed_rules` blocks as defined below.
        /// </summary>
        [Input("managedRules", required: true)]
        public Input<Inputs.PolicyManagedRulesArgs> ManagedRules { get; set; } = null!;

        /// <summary>
        /// The name of the policy. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `policy_settings` block as defined below.
        /// </summary>
        [Input("policySettings")]
        public Input<Inputs.PolicyPolicySettingsArgs>? PolicySettings { get; set; }

        /// <summary>
        /// The name of the resource group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the Web Application Firewall Policy.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public PolicyArgs()
        {
        }
    }

    public sealed class PolicyState : Pulumi.ResourceArgs
    {
        [Input("customRules")]
        private InputList<Inputs.PolicyCustomRuleGetArgs>? _customRules;

        /// <summary>
        /// One or more `custom_rules` blocks as defined below.
        /// </summary>
        public InputList<Inputs.PolicyCustomRuleGetArgs> CustomRules
        {
            get => _customRules ?? (_customRules = new InputList<Inputs.PolicyCustomRuleGetArgs>());
            set => _customRules = value;
        }

        /// <summary>
        /// Resource location. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A `managed_rules` blocks as defined below.
        /// </summary>
        [Input("managedRules")]
        public Input<Inputs.PolicyManagedRulesGetArgs>? ManagedRules { get; set; }

        /// <summary>
        /// The name of the policy. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `policy_settings` block as defined below.
        /// </summary>
        [Input("policySettings")]
        public Input<Inputs.PolicyPolicySettingsGetArgs>? PolicySettings { get; set; }

        /// <summary>
        /// The name of the resource group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the Web Application Firewall Policy.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public PolicyState()
        {
        }
    }
}

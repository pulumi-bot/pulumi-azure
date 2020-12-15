// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.SecurityCenter
{
    /// <summary>
    /// Manages Security Center Automation and Continuous Export. This resource supports three types of destination in the `action`, Logic Apps, Log Analytics and Event Hubs
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
    ///         var current = Output.Create(Azure.Core.GetClientConfig.InvokeAsync());
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "westeurope",
    ///         });
    ///         var exampleEventHubNamespace = new Azure.EventHub.EventHubNamespace("exampleEventHubNamespace", new Azure.EventHub.EventHubNamespaceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Sku = "Standard",
    ///             Capacity = 2,
    ///         });
    ///         var exampleAutomation = new Azure.SecurityCenter.Automation("exampleAutomation", new Azure.SecurityCenter.AutomationArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Actions = 
    ///             {
    ///                 new Azure.SecurityCenter.Inputs.AutomationActionArgs
    ///                 {
    ///                     Type = "EventHub",
    ///                     ResourceId = exampleEventHubNamespace.Id,
    ///                     ConnectionString = exampleEventHubNamespace.DefaultPrimaryConnectionString,
    ///                 },
    ///             },
    ///             Sources = 
    ///             {
    ///                 new Azure.SecurityCenter.Inputs.AutomationSourceArgs
    ///                 {
    ///                     EventSource = "Alerts",
    ///                     RuleSets = 
    ///                     {
    ///                         new Azure.SecurityCenter.Inputs.AutomationSourceRuleSetArgs
    ///                         {
    ///                             Rules = 
    ///                             {
    ///                                 new Azure.SecurityCenter.Inputs.AutomationSourceRuleSetRuleArgs
    ///                                 {
    ///                                     PropertyPath = "properties.metadata.severity",
    ///                                     Operator = "Equals",
    ///                                     ExpectedValue = "High",
    ///                                     PropertyType = "String",
    ///                                 },
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///             Scopes = 
    ///             {
    ///                 current.Apply(current =&gt; $"/subscriptions/{current.SubscriptionId}"),
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Security Center Automations can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:securitycenter/automation:Automation example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Security/automations/automation1
    /// ```
    /// </summary>
    [AzureResourceType("azure:securitycenter/automation:Automation")]
    public partial class Automation : Pulumi.CustomResource
    {
        /// <summary>
        /// One or more `action` blocks as defined below. An `action` tells this automation where the data is to be sent to upon being evaluated by the rules in the `source`.
        /// </summary>
        [Output("actions")]
        public Output<ImmutableArray<Outputs.AutomationAction>> Actions { get; private set; } = null!;

        /// <summary>
        /// Boolean to enable or disable this Security Center Automation
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The Azure Region where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this Security Center Automation. Changing this forces a new Security Center Automation to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A list of scopes on which the automation logic is applied, at least one is required. Supported scopes are a subscription (in this format `/subscriptions/00000000-0000-0000-0000-000000000000`) or a resource group under that subscription (in the format `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example`). The automation will only apply on defined scopes.
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;

        /// <summary>
        /// One or more `source` blocks as defined below. A `source` defines what data types will be processed and a set of rules to filter that data.
        /// </summary>
        [Output("sources")]
        public Output<ImmutableArray<Outputs.AutomationSource>> Sources { get; private set; } = null!;


        /// <summary>
        /// Create a Automation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Automation(string name, AutomationArgs args, CustomResourceOptions? options = null)
            : base("azure:securitycenter/automation:Automation", name, args ?? new AutomationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Automation(string name, Input<string> id, AutomationState? state = null, CustomResourceOptions? options = null)
            : base("azure:securitycenter/automation:Automation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Automation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Automation Get(string name, Input<string> id, AutomationState? state = null, CustomResourceOptions? options = null)
        {
            return new Automation(name, id, state, options);
        }
    }

    public sealed class AutomationArgs : Pulumi.ResourceArgs
    {
        [Input("actions", required: true)]
        private InputList<Inputs.AutomationActionArgs>? _actions;

        /// <summary>
        /// One or more `action` blocks as defined below. An `action` tells this automation where the data is to be sent to upon being evaluated by the rules in the `source`.
        /// </summary>
        public InputList<Inputs.AutomationActionArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.AutomationActionArgs>());
            set => _actions = value;
        }

        /// <summary>
        /// Boolean to enable or disable this Security Center Automation
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The Azure Region where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name which should be used for this Security Center Automation. Changing this forces a new Security Center Automation to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("scopes", required: true)]
        private InputList<string>? _scopes;

        /// <summary>
        /// A list of scopes on which the automation logic is applied, at least one is required. Supported scopes are a subscription (in this format `/subscriptions/00000000-0000-0000-0000-000000000000`) or a resource group under that subscription (in the format `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example`). The automation will only apply on defined scopes.
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        [Input("sources", required: true)]
        private InputList<Inputs.AutomationSourceArgs>? _sources;

        /// <summary>
        /// One or more `source` blocks as defined below. A `source` defines what data types will be processed and a set of rules to filter that data.
        /// </summary>
        public InputList<Inputs.AutomationSourceArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.AutomationSourceArgs>());
            set => _sources = value;
        }

        public AutomationArgs()
        {
        }
    }

    public sealed class AutomationState : Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<Inputs.AutomationActionGetArgs>? _actions;

        /// <summary>
        /// One or more `action` blocks as defined below. An `action` tells this automation where the data is to be sent to upon being evaluated by the rules in the `source`.
        /// </summary>
        public InputList<Inputs.AutomationActionGetArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.AutomationActionGetArgs>());
            set => _actions = value;
        }

        /// <summary>
        /// Boolean to enable or disable this Security Center Automation
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The Azure Region where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name which should be used for this Security Center Automation. Changing this forces a new Security Center Automation to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// A list of scopes on which the automation logic is applied, at least one is required. Supported scopes are a subscription (in this format `/subscriptions/00000000-0000-0000-0000-000000000000`) or a resource group under that subscription (in the format `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example`). The automation will only apply on defined scopes.
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        [Input("sources")]
        private InputList<Inputs.AutomationSourceGetArgs>? _sources;

        /// <summary>
        /// One or more `source` blocks as defined below. A `source` defines what data types will be processed and a set of rules to filter that data.
        /// </summary>
        public InputList<Inputs.AutomationSourceGetArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.AutomationSourceGetArgs>());
            set => _sources = value;
        }

        public AutomationState()
        {
        }
    }
}

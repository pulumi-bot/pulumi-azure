// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppInsights
{
    /// <summary>
    /// Manages an Application Insights Smart Detection Rule.
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
    ///             Location = "West Europe",
    ///         });
    ///         var exampleInsights = new Azure.AppInsights.Insights("exampleInsights", new Azure.AppInsights.InsightsArgs
    ///         {
    ///             Location = "West Europe",
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             ApplicationType = "web",
    ///         });
    ///         var exampleSmartDetectionRule = new Azure.AppInsights.SmartDetectionRule("exampleSmartDetectionRule", new Azure.AppInsights.SmartDetectionRuleArgs
    ///         {
    ///             ApplicationInsightsId = exampleInsights.Id,
    ///             Enabled = false,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class SmartDetectionRule : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies a list of additional recipients that will be sent emails on this Application Insights Smart Detection Rule.
        /// </summary>
        [Output("additionalEmailRecipients")]
        public Output<ImmutableArray<string>> AdditionalEmailRecipients { get; private set; } = null!;

        /// <summary>
        /// The ID of the Application Insights component on which the Smart Detection Rule operates. Changing this forces a new resource to be created.
        /// </summary>
        [Output("applicationInsightsId")]
        public Output<string> ApplicationInsightsId { get; private set; } = null!;

        /// <summary>
        /// Is the Application Insights Smart Detection Rule enabled? Defaults to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Application Insights Smart Detection Rule. Valid values include `Slow page load time`, `Slow server response time`, 
        /// `Long dependency duration`.  Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Do emails get sent to subscription owners? Defaults to `true`.
        /// </summary>
        [Output("sendEmailsToSubscriptionOwners")]
        public Output<bool?> SendEmailsToSubscriptionOwners { get; private set; } = null!;


        /// <summary>
        /// Create a SmartDetectionRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SmartDetectionRule(string name, SmartDetectionRuleArgs args, CustomResourceOptions? options = null)
            : base("azure:appinsights/smartDetectionRule:SmartDetectionRule", name, args ?? new SmartDetectionRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SmartDetectionRule(string name, Input<string> id, SmartDetectionRuleState? state = null, CustomResourceOptions? options = null)
            : base("azure:appinsights/smartDetectionRule:SmartDetectionRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SmartDetectionRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SmartDetectionRule Get(string name, Input<string> id, SmartDetectionRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new SmartDetectionRule(name, id, state, options);
        }
    }

    public sealed class SmartDetectionRuleArgs : Pulumi.ResourceArgs
    {
        [Input("additionalEmailRecipients")]
        private InputList<string>? _additionalEmailRecipients;

        /// <summary>
        /// Specifies a list of additional recipients that will be sent emails on this Application Insights Smart Detection Rule.
        /// </summary>
        public InputList<string> AdditionalEmailRecipients
        {
            get => _additionalEmailRecipients ?? (_additionalEmailRecipients = new InputList<string>());
            set => _additionalEmailRecipients = value;
        }

        /// <summary>
        /// The ID of the Application Insights component on which the Smart Detection Rule operates. Changing this forces a new resource to be created.
        /// </summary>
        [Input("applicationInsightsId", required: true)]
        public Input<string> ApplicationInsightsId { get; set; } = null!;

        /// <summary>
        /// Is the Application Insights Smart Detection Rule enabled? Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Specifies the name of the Application Insights Smart Detection Rule. Valid values include `Slow page load time`, `Slow server response time`, 
        /// `Long dependency duration`.  Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Do emails get sent to subscription owners? Defaults to `true`.
        /// </summary>
        [Input("sendEmailsToSubscriptionOwners")]
        public Input<bool>? SendEmailsToSubscriptionOwners { get; set; }

        public SmartDetectionRuleArgs()
        {
        }
    }

    public sealed class SmartDetectionRuleState : Pulumi.ResourceArgs
    {
        [Input("additionalEmailRecipients")]
        private InputList<string>? _additionalEmailRecipients;

        /// <summary>
        /// Specifies a list of additional recipients that will be sent emails on this Application Insights Smart Detection Rule.
        /// </summary>
        public InputList<string> AdditionalEmailRecipients
        {
            get => _additionalEmailRecipients ?? (_additionalEmailRecipients = new InputList<string>());
            set => _additionalEmailRecipients = value;
        }

        /// <summary>
        /// The ID of the Application Insights component on which the Smart Detection Rule operates. Changing this forces a new resource to be created.
        /// </summary>
        [Input("applicationInsightsId")]
        public Input<string>? ApplicationInsightsId { get; set; }

        /// <summary>
        /// Is the Application Insights Smart Detection Rule enabled? Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Specifies the name of the Application Insights Smart Detection Rule. Valid values include `Slow page load time`, `Slow server response time`, 
        /// `Long dependency duration`.  Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Do emails get sent to subscription owners? Defaults to `true`.
        /// </summary>
        [Input("sendEmailsToSubscriptionOwners")]
        public Input<bool>? SendEmailsToSubscriptionOwners { get; set; }

        public SmartDetectionRuleState()
        {
        }
    }
}
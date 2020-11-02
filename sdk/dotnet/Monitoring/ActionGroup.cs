// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring
{
    /// <summary>
    /// Manages an Action Group within Azure Monitor.
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
    ///             Location = "West US",
    ///         });
    ///         var exampleActionGroup = new Azure.Monitoring.ActionGroup("exampleActionGroup", new Azure.Monitoring.ActionGroupArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             ShortName = "p0action",
    ///             ArmRoleReceivers = 
    ///             {
    ///                 new Azure.Monitoring.Inputs.ActionGroupArmRoleReceiverArgs
    ///                 {
    ///                     Name = "armroleaction",
    ///                     RoleId = "de139f84-1756-47ae-9be6-808fbbe84772",
    ///                     UseCommonAlertSchema = true,
    ///                 },
    ///             },
    ///             AutomationRunbookReceivers = 
    ///             {
    ///                 new Azure.Monitoring.Inputs.ActionGroupAutomationRunbookReceiverArgs
    ///                 {
    ///                     Name = "action_name_1",
    ///                     AutomationAccountId = "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg-runbooks/providers/microsoft.automation/automationaccounts/aaa001",
    ///                     RunbookName = "my runbook",
    ///                     WebhookResourceId = "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg-runbooks/providers/microsoft.automation/automationaccounts/aaa001/webhooks/webhook_alert",
    ///                     IsGlobalRunbook = true,
    ///                     ServiceUri = "https://s13events.azure-automation.net/webhooks?token=randomtoken",
    ///                     UseCommonAlertSchema = true,
    ///                 },
    ///             },
    ///             AzureAppPushReceivers = 
    ///             {
    ///                 new Azure.Monitoring.Inputs.ActionGroupAzureAppPushReceiverArgs
    ///                 {
    ///                     Name = "pushtoadmin",
    ///                     EmailAddress = "admin@contoso.com",
    ///                 },
    ///             },
    ///             AzureFunctionReceivers = 
    ///             {
    ///                 new Azure.Monitoring.Inputs.ActionGroupAzureFunctionReceiverArgs
    ///                 {
    ///                     Name = "funcaction",
    ///                     FunctionAppResourceId = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-funcapp/providers/Microsoft.Web/sites/funcapp",
    ///                     FunctionName = "myfunc",
    ///                     HttpTriggerUrl = "https://example.com/trigger",
    ///                     UseCommonAlertSchema = true,
    ///                 },
    ///             },
    ///             EmailReceivers = 
    ///             {
    ///                 new Azure.Monitoring.Inputs.ActionGroupEmailReceiverArgs
    ///                 {
    ///                     Name = "sendtoadmin",
    ///                     EmailAddress = "admin@contoso.com",
    ///                 },
    ///                 new Azure.Monitoring.Inputs.ActionGroupEmailReceiverArgs
    ///                 {
    ///                     Name = "sendtodevops",
    ///                     EmailAddress = "devops@contoso.com",
    ///                     UseCommonAlertSchema = true,
    ///                 },
    ///             },
    ///             ItsmReceivers = 
    ///             {
    ///                 new Azure.Monitoring.Inputs.ActionGroupItsmReceiverArgs
    ///                 {
    ///                     Name = "createorupdateticket",
    ///                     WorkspaceId = "6eee3a18-aac3-40e4-b98e-1f309f329816",
    ///                     ConnectionId = "53de6956-42b4-41ba-be3c-b154cdf17b13",
    ///                     TicketConfiguration = "{}",
    ///                     Region = "southcentralus",
    ///                 },
    ///             },
    ///             LogicAppReceivers = 
    ///             {
    ///                 new Azure.Monitoring.Inputs.ActionGroupLogicAppReceiverArgs
    ///                 {
    ///                     Name = "logicappaction",
    ///                     ResourceId = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-logicapp/providers/Microsoft.Logic/workflows/logicapp",
    ///                     CallbackUrl = "https://logicapptriggerurl/...",
    ///                     UseCommonAlertSchema = true,
    ///                 },
    ///             },
    ///             SmsReceivers = 
    ///             {
    ///                 new Azure.Monitoring.Inputs.ActionGroupSmsReceiverArgs
    ///                 {
    ///                     Name = "oncallmsg",
    ///                     CountryCode = "1",
    ///                     PhoneNumber = "1231231234",
    ///                 },
    ///             },
    ///             VoiceReceivers = 
    ///             {
    ///                 new Azure.Monitoring.Inputs.ActionGroupVoiceReceiverArgs
    ///                 {
    ///                     Name = "remotesupport",
    ///                     CountryCode = "86",
    ///                     PhoneNumber = "13888888888",
    ///                 },
    ///             },
    ///             WebhookReceivers = 
    ///             {
    ///                 new Azure.Monitoring.Inputs.ActionGroupWebhookReceiverArgs
    ///                 {
    ///                     Name = "callmyapiaswell",
    ///                     ServiceUri = "http://example.com/alert",
    ///                     UseCommonAlertSchema = true,
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
    /// Action Groups can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class ActionGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// One or more `arm_role_receiver` blocks as defined below.
        /// </summary>
        [Output("armRoleReceivers")]
        public Output<ImmutableArray<Outputs.ActionGroupArmRoleReceiver>> ArmRoleReceivers { get; private set; } = null!;

        /// <summary>
        /// One or more `automation_runbook_receiver` blocks as defined below.
        /// </summary>
        [Output("automationRunbookReceivers")]
        public Output<ImmutableArray<Outputs.ActionGroupAutomationRunbookReceiver>> AutomationRunbookReceivers { get; private set; } = null!;

        /// <summary>
        /// One or more `azure_app_push_receiver` blocks as defined below.
        /// </summary>
        [Output("azureAppPushReceivers")]
        public Output<ImmutableArray<Outputs.ActionGroupAzureAppPushReceiver>> AzureAppPushReceivers { get; private set; } = null!;

        /// <summary>
        /// One or more `azure_function_receiver` blocks as defined below.
        /// </summary>
        [Output("azureFunctionReceivers")]
        public Output<ImmutableArray<Outputs.ActionGroupAzureFunctionReceiver>> AzureFunctionReceivers { get; private set; } = null!;

        /// <summary>
        /// One or more `email_receiver` blocks as defined below.
        /// </summary>
        [Output("emailReceivers")]
        public Output<ImmutableArray<Outputs.ActionGroupEmailReceiver>> EmailReceivers { get; private set; } = null!;

        /// <summary>
        /// Whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications. Defaults to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// One or more `itsm_receiver` blocks as defined below.
        /// </summary>
        [Output("itsmReceivers")]
        public Output<ImmutableArray<Outputs.ActionGroupItsmReceiver>> ItsmReceivers { get; private set; } = null!;

        /// <summary>
        /// One or more `logic_app_receiver` blocks as defined below.
        /// </summary>
        [Output("logicAppReceivers")]
        public Output<ImmutableArray<Outputs.ActionGroupLogicAppReceiver>> LogicAppReceivers { get; private set; } = null!;

        /// <summary>
        /// The name of the Action Group. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Action Group instance.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The short name of the action group. This will be used in SMS messages.
        /// </summary>
        [Output("shortName")]
        public Output<string> ShortName { get; private set; } = null!;

        /// <summary>
        /// One or more `sms_receiver` blocks as defined below.
        /// </summary>
        [Output("smsReceivers")]
        public Output<ImmutableArray<Outputs.ActionGroupSmsReceiver>> SmsReceivers { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// One or more `voice_receiver` blocks as defined below.
        /// </summary>
        [Output("voiceReceivers")]
        public Output<ImmutableArray<Outputs.ActionGroupVoiceReceiver>> VoiceReceivers { get; private set; } = null!;

        /// <summary>
        /// One or more `webhook_receiver` blocks as defined below.
        /// </summary>
        [Output("webhookReceivers")]
        public Output<ImmutableArray<Outputs.ActionGroupWebhookReceiver>> WebhookReceivers { get; private set; } = null!;


        /// <summary>
        /// Create a ActionGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ActionGroup(string name, ActionGroupArgs args, CustomResourceOptions? options = null)
            : base("azure:monitoring/actionGroup:ActionGroup", name, args ?? new ActionGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ActionGroup(string name, Input<string> id, ActionGroupState? state = null, CustomResourceOptions? options = null)
            : base("azure:monitoring/actionGroup:ActionGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ActionGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ActionGroup Get(string name, Input<string> id, ActionGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ActionGroup(name, id, state, options);
        }
    }

    public sealed class ActionGroupArgs : Pulumi.ResourceArgs
    {
        [Input("armRoleReceivers")]
        private InputList<Inputs.ActionGroupArmRoleReceiverArgs>? _armRoleReceivers;

        /// <summary>
        /// One or more `arm_role_receiver` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ActionGroupArmRoleReceiverArgs> ArmRoleReceivers
        {
            get => _armRoleReceivers ?? (_armRoleReceivers = new InputList<Inputs.ActionGroupArmRoleReceiverArgs>());
            set => _armRoleReceivers = value;
        }

        [Input("automationRunbookReceivers")]
        private InputList<Inputs.ActionGroupAutomationRunbookReceiverArgs>? _automationRunbookReceivers;

        /// <summary>
        /// One or more `automation_runbook_receiver` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ActionGroupAutomationRunbookReceiverArgs> AutomationRunbookReceivers
        {
            get => _automationRunbookReceivers ?? (_automationRunbookReceivers = new InputList<Inputs.ActionGroupAutomationRunbookReceiverArgs>());
            set => _automationRunbookReceivers = value;
        }

        [Input("azureAppPushReceivers")]
        private InputList<Inputs.ActionGroupAzureAppPushReceiverArgs>? _azureAppPushReceivers;

        /// <summary>
        /// One or more `azure_app_push_receiver` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ActionGroupAzureAppPushReceiverArgs> AzureAppPushReceivers
        {
            get => _azureAppPushReceivers ?? (_azureAppPushReceivers = new InputList<Inputs.ActionGroupAzureAppPushReceiverArgs>());
            set => _azureAppPushReceivers = value;
        }

        [Input("azureFunctionReceivers")]
        private InputList<Inputs.ActionGroupAzureFunctionReceiverArgs>? _azureFunctionReceivers;

        /// <summary>
        /// One or more `azure_function_receiver` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ActionGroupAzureFunctionReceiverArgs> AzureFunctionReceivers
        {
            get => _azureFunctionReceivers ?? (_azureFunctionReceivers = new InputList<Inputs.ActionGroupAzureFunctionReceiverArgs>());
            set => _azureFunctionReceivers = value;
        }

        [Input("emailReceivers")]
        private InputList<Inputs.ActionGroupEmailReceiverArgs>? _emailReceivers;

        /// <summary>
        /// One or more `email_receiver` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ActionGroupEmailReceiverArgs> EmailReceivers
        {
            get => _emailReceivers ?? (_emailReceivers = new InputList<Inputs.ActionGroupEmailReceiverArgs>());
            set => _emailReceivers = value;
        }

        /// <summary>
        /// Whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("itsmReceivers")]
        private InputList<Inputs.ActionGroupItsmReceiverArgs>? _itsmReceivers;

        /// <summary>
        /// One or more `itsm_receiver` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ActionGroupItsmReceiverArgs> ItsmReceivers
        {
            get => _itsmReceivers ?? (_itsmReceivers = new InputList<Inputs.ActionGroupItsmReceiverArgs>());
            set => _itsmReceivers = value;
        }

        [Input("logicAppReceivers")]
        private InputList<Inputs.ActionGroupLogicAppReceiverArgs>? _logicAppReceivers;

        /// <summary>
        /// One or more `logic_app_receiver` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ActionGroupLogicAppReceiverArgs> LogicAppReceivers
        {
            get => _logicAppReceivers ?? (_logicAppReceivers = new InputList<Inputs.ActionGroupLogicAppReceiverArgs>());
            set => _logicAppReceivers = value;
        }

        /// <summary>
        /// The name of the Action Group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Action Group instance.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The short name of the action group. This will be used in SMS messages.
        /// </summary>
        [Input("shortName", required: true)]
        public Input<string> ShortName { get; set; } = null!;

        [Input("smsReceivers")]
        private InputList<Inputs.ActionGroupSmsReceiverArgs>? _smsReceivers;

        /// <summary>
        /// One or more `sms_receiver` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ActionGroupSmsReceiverArgs> SmsReceivers
        {
            get => _smsReceivers ?? (_smsReceivers = new InputList<Inputs.ActionGroupSmsReceiverArgs>());
            set => _smsReceivers = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("voiceReceivers")]
        private InputList<Inputs.ActionGroupVoiceReceiverArgs>? _voiceReceivers;

        /// <summary>
        /// One or more `voice_receiver` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ActionGroupVoiceReceiverArgs> VoiceReceivers
        {
            get => _voiceReceivers ?? (_voiceReceivers = new InputList<Inputs.ActionGroupVoiceReceiverArgs>());
            set => _voiceReceivers = value;
        }

        [Input("webhookReceivers")]
        private InputList<Inputs.ActionGroupWebhookReceiverArgs>? _webhookReceivers;

        /// <summary>
        /// One or more `webhook_receiver` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ActionGroupWebhookReceiverArgs> WebhookReceivers
        {
            get => _webhookReceivers ?? (_webhookReceivers = new InputList<Inputs.ActionGroupWebhookReceiverArgs>());
            set => _webhookReceivers = value;
        }

        public ActionGroupArgs()
        {
        }
    }

    public sealed class ActionGroupState : Pulumi.ResourceArgs
    {
        [Input("armRoleReceivers")]
        private InputList<Inputs.ActionGroupArmRoleReceiverGetArgs>? _armRoleReceivers;

        /// <summary>
        /// One or more `arm_role_receiver` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ActionGroupArmRoleReceiverGetArgs> ArmRoleReceivers
        {
            get => _armRoleReceivers ?? (_armRoleReceivers = new InputList<Inputs.ActionGroupArmRoleReceiverGetArgs>());
            set => _armRoleReceivers = value;
        }

        [Input("automationRunbookReceivers")]
        private InputList<Inputs.ActionGroupAutomationRunbookReceiverGetArgs>? _automationRunbookReceivers;

        /// <summary>
        /// One or more `automation_runbook_receiver` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ActionGroupAutomationRunbookReceiverGetArgs> AutomationRunbookReceivers
        {
            get => _automationRunbookReceivers ?? (_automationRunbookReceivers = new InputList<Inputs.ActionGroupAutomationRunbookReceiverGetArgs>());
            set => _automationRunbookReceivers = value;
        }

        [Input("azureAppPushReceivers")]
        private InputList<Inputs.ActionGroupAzureAppPushReceiverGetArgs>? _azureAppPushReceivers;

        /// <summary>
        /// One or more `azure_app_push_receiver` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ActionGroupAzureAppPushReceiverGetArgs> AzureAppPushReceivers
        {
            get => _azureAppPushReceivers ?? (_azureAppPushReceivers = new InputList<Inputs.ActionGroupAzureAppPushReceiverGetArgs>());
            set => _azureAppPushReceivers = value;
        }

        [Input("azureFunctionReceivers")]
        private InputList<Inputs.ActionGroupAzureFunctionReceiverGetArgs>? _azureFunctionReceivers;

        /// <summary>
        /// One or more `azure_function_receiver` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ActionGroupAzureFunctionReceiverGetArgs> AzureFunctionReceivers
        {
            get => _azureFunctionReceivers ?? (_azureFunctionReceivers = new InputList<Inputs.ActionGroupAzureFunctionReceiverGetArgs>());
            set => _azureFunctionReceivers = value;
        }

        [Input("emailReceivers")]
        private InputList<Inputs.ActionGroupEmailReceiverGetArgs>? _emailReceivers;

        /// <summary>
        /// One or more `email_receiver` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ActionGroupEmailReceiverGetArgs> EmailReceivers
        {
            get => _emailReceivers ?? (_emailReceivers = new InputList<Inputs.ActionGroupEmailReceiverGetArgs>());
            set => _emailReceivers = value;
        }

        /// <summary>
        /// Whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("itsmReceivers")]
        private InputList<Inputs.ActionGroupItsmReceiverGetArgs>? _itsmReceivers;

        /// <summary>
        /// One or more `itsm_receiver` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ActionGroupItsmReceiverGetArgs> ItsmReceivers
        {
            get => _itsmReceivers ?? (_itsmReceivers = new InputList<Inputs.ActionGroupItsmReceiverGetArgs>());
            set => _itsmReceivers = value;
        }

        [Input("logicAppReceivers")]
        private InputList<Inputs.ActionGroupLogicAppReceiverGetArgs>? _logicAppReceivers;

        /// <summary>
        /// One or more `logic_app_receiver` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ActionGroupLogicAppReceiverGetArgs> LogicAppReceivers
        {
            get => _logicAppReceivers ?? (_logicAppReceivers = new InputList<Inputs.ActionGroupLogicAppReceiverGetArgs>());
            set => _logicAppReceivers = value;
        }

        /// <summary>
        /// The name of the Action Group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Action Group instance.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The short name of the action group. This will be used in SMS messages.
        /// </summary>
        [Input("shortName")]
        public Input<string>? ShortName { get; set; }

        [Input("smsReceivers")]
        private InputList<Inputs.ActionGroupSmsReceiverGetArgs>? _smsReceivers;

        /// <summary>
        /// One or more `sms_receiver` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ActionGroupSmsReceiverGetArgs> SmsReceivers
        {
            get => _smsReceivers ?? (_smsReceivers = new InputList<Inputs.ActionGroupSmsReceiverGetArgs>());
            set => _smsReceivers = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("voiceReceivers")]
        private InputList<Inputs.ActionGroupVoiceReceiverGetArgs>? _voiceReceivers;

        /// <summary>
        /// One or more `voice_receiver` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ActionGroupVoiceReceiverGetArgs> VoiceReceivers
        {
            get => _voiceReceivers ?? (_voiceReceivers = new InputList<Inputs.ActionGroupVoiceReceiverGetArgs>());
            set => _voiceReceivers = value;
        }

        [Input("webhookReceivers")]
        private InputList<Inputs.ActionGroupWebhookReceiverGetArgs>? _webhookReceivers;

        /// <summary>
        /// One or more `webhook_receiver` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ActionGroupWebhookReceiverGetArgs> WebhookReceivers
        {
            get => _webhookReceivers ?? (_webhookReceivers = new InputList<Inputs.ActionGroupWebhookReceiverGetArgs>());
            set => _webhookReceivers = value;
        }

        public ActionGroupState()
        {
        }
    }
}

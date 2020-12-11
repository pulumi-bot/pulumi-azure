// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Automation
{
    /// <summary>
    /// Manages a Automation Runbook.
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
    ///         var exampleAccount = new Azure.Automation.Account("exampleAccount", new Azure.Automation.AccountArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             SkuName = "Basic",
    ///         });
    ///         var exampleRunBook = new Azure.Automation.RunBook("exampleRunBook", new Azure.Automation.RunBookArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AutomationAccountName = exampleAccount.Name,
    ///             LogVerbose = true,
    ///             LogProgress = true,
    ///             Description = "This is an example runbook",
    ///             RunbookType = "PowerShellWorkflow",
    ///             PublishContentLink = new Azure.Automation.Inputs.RunBookPublishContentLinkArgs
    ///             {
    ///                 Uri = "https://raw.githubusercontent.com/Azure/azure-quickstart-templates/c4935ffb69246a6058eb24f54640f53f69d3ac9f/101-automation-runbook-getvms/Runbooks/Get-AzureVMTutorial.ps1",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Automation Runbooks can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:automation/runBook:RunBook Get-AzureVMTutorial /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/runbooks/Get-AzureVMTutorial
    /// ```
    /// </summary>
    [AzureResourceType("azure:automation/runBook:RunBook")]
    public partial class RunBook : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("automationAccountName")]
        public Output<string> AutomationAccountName { get; private set; } = null!;

        /// <summary>
        /// The desired content of the runbook.
        /// </summary>
        [Output("content")]
        public Output<string> Content { get; private set; } = null!;

        /// <summary>
        /// A description for this credential.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("jobSchedules")]
        public Output<ImmutableArray<Outputs.RunBookJobSchedule>> JobSchedules { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Progress log option.
        /// </summary>
        [Output("logProgress")]
        public Output<bool> LogProgress { get; private set; } = null!;

        /// <summary>
        /// Verbose log option.
        /// </summary>
        [Output("logVerbose")]
        public Output<bool> LogVerbose { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Runbook. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The published runbook content link.
        /// </summary>
        [Output("publishContentLink")]
        public Output<Outputs.RunBookPublishContentLink?> PublishContentLink { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The type of the runbook - can be either `Graph`, `GraphPowerShell`, `GraphPowerShellWorkflow`, `PowerShellWorkflow`, `PowerShell` or `Script`.
        /// </summary>
        [Output("runbookType")]
        public Output<string> RunbookType { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a RunBook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RunBook(string name, RunBookArgs args, CustomResourceOptions? options = null)
            : base("azure:automation/runBook:RunBook", name, args ?? new RunBookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RunBook(string name, Input<string> id, RunBookState? state = null, CustomResourceOptions? options = null)
            : base("azure:automation/runBook:RunBook", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RunBook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RunBook Get(string name, Input<string> id, RunBookState? state = null, CustomResourceOptions? options = null)
        {
            return new RunBook(name, id, state, options);
        }
    }

    public sealed class RunBookArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public Input<string> AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// The desired content of the runbook.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// A description for this credential.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("jobSchedules")]
        private InputList<Inputs.RunBookJobScheduleArgs>? _jobSchedules;
        public InputList<Inputs.RunBookJobScheduleArgs> JobSchedules
        {
            get => _jobSchedules ?? (_jobSchedules = new InputList<Inputs.RunBookJobScheduleArgs>());
            set => _jobSchedules = value;
        }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Progress log option.
        /// </summary>
        [Input("logProgress", required: true)]
        public Input<bool> LogProgress { get; set; } = null!;

        /// <summary>
        /// Verbose log option.
        /// </summary>
        [Input("logVerbose", required: true)]
        public Input<bool> LogVerbose { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Runbook. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The published runbook content link.
        /// </summary>
        [Input("publishContentLink")]
        public Input<Inputs.RunBookPublishContentLinkArgs>? PublishContentLink { get; set; }

        /// <summary>
        /// The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The type of the runbook - can be either `Graph`, `GraphPowerShell`, `GraphPowerShellWorkflow`, `PowerShellWorkflow`, `PowerShell` or `Script`.
        /// </summary>
        [Input("runbookType", required: true)]
        public Input<string> RunbookType { get; set; } = null!;

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

        public RunBookArgs()
        {
        }
    }

    public sealed class RunBookState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("automationAccountName")]
        public Input<string>? AutomationAccountName { get; set; }

        /// <summary>
        /// The desired content of the runbook.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// A description for this credential.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("jobSchedules")]
        private InputList<Inputs.RunBookJobScheduleGetArgs>? _jobSchedules;
        public InputList<Inputs.RunBookJobScheduleGetArgs> JobSchedules
        {
            get => _jobSchedules ?? (_jobSchedules = new InputList<Inputs.RunBookJobScheduleGetArgs>());
            set => _jobSchedules = value;
        }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Progress log option.
        /// </summary>
        [Input("logProgress")]
        public Input<bool>? LogProgress { get; set; }

        /// <summary>
        /// Verbose log option.
        /// </summary>
        [Input("logVerbose")]
        public Input<bool>? LogVerbose { get; set; }

        /// <summary>
        /// Specifies the name of the Runbook. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The published runbook content link.
        /// </summary>
        [Input("publishContentLink")]
        public Input<Inputs.RunBookPublishContentLinkGetArgs>? PublishContentLink { get; set; }

        /// <summary>
        /// The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The type of the runbook - can be either `Graph`, `GraphPowerShell`, `GraphPowerShellWorkflow`, `PowerShellWorkflow`, `PowerShell` or `Script`.
        /// </summary>
        [Input("runbookType")]
        public Input<string>? RunbookType { get; set; }

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

        public RunBookState()
        {
        }
    }
}

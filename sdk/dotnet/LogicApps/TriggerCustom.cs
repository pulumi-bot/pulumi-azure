// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.LogicApps
{
    /// <summary>
    /// Manages a Custom Trigger within a Logic App Workflow
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
    ///             Location = "East US",
    ///         });
    ///         var exampleWorkflow = new Azure.LogicApps.Workflow("exampleWorkflow", new Azure.LogicApps.WorkflowArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var exampleTriggerCustom = new Azure.LogicApps.TriggerCustom("exampleTriggerCustom", new Azure.LogicApps.TriggerCustomArgs
    ///         {
    ///             LogicAppId = exampleWorkflow.Id,
    ///             Body = @"{
    ///   ""recurrence"": {
    ///     ""frequency"": ""Day"",
    ///     ""interval"": 1
    ///   },
    ///   ""type"": ""Recurrence""
    /// }
    /// ",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class TriggerCustom : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the JSON Blob defining the Body of this Custom Trigger.
        /// </summary>
        [Output("body")]
        public Output<string> Body { get; private set; } = null!;

        /// <summary>
        /// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        /// </summary>
        [Output("logicAppId")]
        public Output<string> LogicAppId { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the HTTP Trigger to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a TriggerCustom resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TriggerCustom(string name, TriggerCustomArgs args, CustomResourceOptions? options = null)
            : base("azure:logicapps/triggerCustom:TriggerCustom", name, args ?? new TriggerCustomArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TriggerCustom(string name, Input<string> id, TriggerCustomState? state = null, CustomResourceOptions? options = null)
            : base("azure:logicapps/triggerCustom:TriggerCustom", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TriggerCustom resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TriggerCustom Get(string name, Input<string> id, TriggerCustomState? state = null, CustomResourceOptions? options = null)
        {
            return new TriggerCustom(name, id, state, options);
        }
    }

    public sealed class TriggerCustomArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the JSON Blob defining the Body of this Custom Trigger.
        /// </summary>
        [Input("body", required: true)]
        public Input<string> Body { get; set; } = null!;

        /// <summary>
        /// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        /// </summary>
        [Input("logicAppId", required: true)]
        public Input<string> LogicAppId { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the HTTP Trigger to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public TriggerCustomArgs()
        {
        }
    }

    public sealed class TriggerCustomState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the JSON Blob defining the Body of this Custom Trigger.
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        /// <summary>
        /// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        /// </summary>
        [Input("logicAppId")]
        public Input<string>? LogicAppId { get; set; }

        /// <summary>
        /// Specifies the name of the HTTP Trigger to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public TriggerCustomState()
        {
        }
    }
}

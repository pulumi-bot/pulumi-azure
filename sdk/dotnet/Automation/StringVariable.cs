// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Automation
{
    /// <summary>
    /// Manages a string variable in Azure Automation
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/automation_variable_string.html.markdown.
    /// </summary>
    public partial class StringVariable : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("automationAccountName")]
        public Output<string> AutomationAccountName { get; private set; } = null!;

        /// <summary>
        /// The description of the Automation Variable.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies if the Automation Variable is encrypted. Defaults to `false`.
        /// </summary>
        [Output("encrypted")]
        public Output<bool?> Encrypted { get; private set; } = null!;

        /// <summary>
        /// The name of the Automation Variable. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The value of the Automation Variable as a `string`.
        /// </summary>
        [Output("value")]
        public Output<string?> Value { get; private set; } = null!;


        /// <summary>
        /// Create a StringVariable resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StringVariable(string name, StringVariableArgs args, CustomResourceOptions? options = null)
            : base("azure:automation/stringVariable:StringVariable", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private StringVariable(string name, Input<string> id, StringVariableState? state = null, CustomResourceOptions? options = null)
            : base("azure:automation/stringVariable:StringVariable", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StringVariable resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StringVariable Get(string name, Input<string> id, StringVariableState? state = null, CustomResourceOptions? options = null)
        {
            return new StringVariable(name, id, state, options);
        }
    }

    public sealed class StringVariableArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public Input<string> AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// The description of the Automation Variable.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies if the Automation Variable is encrypted. Defaults to `false`.
        /// </summary>
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        /// <summary>
        /// The name of the Automation Variable. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The value of the Automation Variable as a `string`.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public StringVariableArgs()
        {
        }
    }

    public sealed class StringVariableState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("automationAccountName")]
        public Input<string>? AutomationAccountName { get; set; }

        /// <summary>
        /// The description of the Automation Variable.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies if the Automation Variable is encrypted. Defaults to `false`.
        /// </summary>
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        /// <summary>
        /// The name of the Automation Variable. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The value of the Automation Variable as a `string`.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public StringVariableState()
        {
        }
    }
}

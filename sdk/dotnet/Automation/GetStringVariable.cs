// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Automation
{
    public static class GetStringVariable
    {
        /// <summary>
        /// Use this data source to access information about an existing Automation String Variable.
        /// 
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.Automation.GetStringVariable.InvokeAsync(new Azure.Automation.GetStringVariableArgs
        ///         {
        ///             Name = "tfex-example-var",
        ///             ResourceGroupName = "tfex-example-rg",
        ///             AutomationAccountName = "tfex-example-account",
        ///         }));
        ///         this.VariableId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("variableId")]
        ///     public Output&lt;string&gt; VariableId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetStringVariableResult> InvokeAsync(GetStringVariableArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStringVariableResult>("azure:automation/getStringVariable:getStringVariable", args ?? new GetStringVariableArgs(), options.WithVersion());

        public static Output<GetStringVariableResult> Invoke(GetStringVariableOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.AutomationAccountName.Box(),
                args.Name.Box(),
                args.ResourceGroupName.Box()
            ).Apply(a => {
                    var args = new GetStringVariableArgs();
                    a[0].Set(args, nameof(args.AutomationAccountName));
                    a[1].Set(args, nameof(args.Name));
                    a[2].Set(args, nameof(args.ResourceGroupName));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetStringVariableArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the automation account in which the Automation Variable exists.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public string AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// The name of the Automation Variable.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the automation account exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetStringVariableArgs()
        {
        }
    }

    public sealed class GetStringVariableOutputArgs
    {
        /// <summary>
        /// The name of the automation account in which the Automation Variable exists.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public Input<string> AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// The name of the Automation Variable.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the automation account exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetStringVariableOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetStringVariableResult
    {
        public readonly string AutomationAccountName;
        /// <summary>
        /// The description of the Automation Variable.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Specifies if the Automation Variable is encrypted. Defaults to `false`.
        /// </summary>
        public readonly bool Encrypted;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The value of the Automation Variable as a `string`.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetStringVariableResult(
            string automationAccountName,

            string description,

            bool encrypted,

            string id,

            string name,

            string resourceGroupName,

            string value)
        {
            AutomationAccountName = automationAccountName;
            Description = description;
            Encrypted = encrypted;
            Id = id;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Value = value;
        }
    }
}

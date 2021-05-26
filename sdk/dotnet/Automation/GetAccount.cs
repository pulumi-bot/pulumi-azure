// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Automation
{
    public static class GetAccount
    {
        /// <summary>
        /// Use this data source to access information about an existing Automation Account.
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
        ///         var example = Output.Create(Azure.Automation.GetAccount.InvokeAsync(new Azure.Automation.GetAccountArgs
        ///         {
        ///             Name = "example-account",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///         this.AutomationAccountId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("automationAccountId")]
        ///     public Output&lt;string&gt; AutomationAccountId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAccountResult> InvokeAsync(GetAccountArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccountResult>("azure:automation/getAccount:getAccount", args ?? new GetAccountArgs(), options.WithVersion());

        public static Output<GetAccountResult> Apply(GetAccountApplyArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Name.Box(),
                args.ResourceGroupName.Box()
            ).Apply(a => {
                    var args = new GetAccountArgs();
                    a[0].Set(args, nameof(args.Name));
                    a[1].Set(args, nameof(args.ResourceGroupName));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Automation Account.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group where the Automation Account exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAccountArgs()
        {
        }
    }

    public sealed class GetAccountApplyArgs
    {
        /// <summary>
        /// The name of the Automation Account.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group where the Automation Account exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetAccountApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAccountResult
    {
        /// <summary>
        /// The Endpoint for this Automation Account.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// The Primary Access Key for the Automation Account.
        /// </summary>
        public readonly string PrimaryKey;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The Secondary Access Key for the Automation Account.
        /// </summary>
        public readonly string SecondaryKey;

        [OutputConstructor]
        private GetAccountResult(
            string endpoint,

            string id,

            string name,

            string primaryKey,

            string resourceGroupName,

            string secondaryKey)
        {
            Endpoint = endpoint;
            Id = id;
            Name = name;
            PrimaryKey = primaryKey;
            ResourceGroupName = resourceGroupName;
            SecondaryKey = secondaryKey;
        }
    }
}

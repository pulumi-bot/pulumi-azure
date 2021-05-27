// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Billing
{
    public static class GetMcaAccountScope
    {
        /// <summary>
        /// Use this data source to access an ID for your MCA Account billing scope.
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
        ///         var example = Output.Create(Azure.Billing.GetMcaAccountScope.InvokeAsync(new Azure.Billing.GetMcaAccountScopeArgs
        ///         {
        ///             BillingAccountName = "e879cf0f-2b4d-5431-109a-f72fc9868693:024cabf4-7321-4cf9-be59-df0c77ca51de_2019-05-31",
        ///             BillingProfileName = "PE2Q-NOIT-BG7-TGB",
        ///             InvoiceSectionName = "MTT4-OBS7-PJA-TGB",
        ///         }));
        ///         this.Id = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMcaAccountScopeResult> InvokeAsync(GetMcaAccountScopeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMcaAccountScopeResult>("azure:billing/getMcaAccountScope:getMcaAccountScope", args ?? new GetMcaAccountScopeArgs(), options.WithVersion());

        public static Output<GetMcaAccountScopeResult> Invoke(GetMcaAccountScopeOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.BillingAccountName.Box(),
                args.BillingProfileName.Box(),
                args.InvoiceSectionName.Box()
            ).Apply(a => {
                    var args = new GetMcaAccountScopeArgs();
                    a[0].Set(args, nameof(args.BillingAccountName));
                    a[1].Set(args, nameof(args.BillingProfileName));
                    a[2].Set(args, nameof(args.InvoiceSectionName));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetMcaAccountScopeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Billing Account Name of the MCA account.
        /// </summary>
        [Input("billingAccountName", required: true)]
        public string BillingAccountName { get; set; } = null!;

        /// <summary>
        /// The Billing Profile Name in the above Billing Account.
        /// </summary>
        [Input("billingProfileName", required: true)]
        public string BillingProfileName { get; set; } = null!;

        /// <summary>
        /// The Invoice Section Name in the above Billing Profile.
        /// </summary>
        [Input("invoiceSectionName", required: true)]
        public string InvoiceSectionName { get; set; } = null!;

        public GetMcaAccountScopeArgs()
        {
        }
    }

    public sealed class GetMcaAccountScopeOutputArgs
    {
        /// <summary>
        /// The Billing Account Name of the MCA account.
        /// </summary>
        [Input("billingAccountName", required: true)]
        public Input<string> BillingAccountName { get; set; } = null!;

        /// <summary>
        /// The Billing Profile Name in the above Billing Account.
        /// </summary>
        [Input("billingProfileName", required: true)]
        public Input<string> BillingProfileName { get; set; } = null!;

        /// <summary>
        /// The Invoice Section Name in the above Billing Profile.
        /// </summary>
        [Input("invoiceSectionName", required: true)]
        public Input<string> InvoiceSectionName { get; set; } = null!;

        public GetMcaAccountScopeOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMcaAccountScopeResult
    {
        public readonly string BillingAccountName;
        public readonly string BillingProfileName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InvoiceSectionName;

        [OutputConstructor]
        private GetMcaAccountScopeResult(
            string billingAccountName,

            string billingProfileName,

            string id,

            string invoiceSectionName)
        {
            BillingAccountName = billingAccountName;
            BillingProfileName = billingProfileName;
            Id = id;
            InvoiceSectionName = invoiceSectionName;
        }
    }
}

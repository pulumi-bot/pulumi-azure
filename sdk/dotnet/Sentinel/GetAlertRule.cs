// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Sentinel
{
    public static class GetAlertRule
    {
        /// <summary>
        /// Use this data source to access information about an existing Sentinel Alert Rule.
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
        ///         var exampleAnalyticsWorkspace = Output.Create(Azure.OperationalInsights.GetAnalyticsWorkspace.InvokeAsync(new Azure.OperationalInsights.GetAnalyticsWorkspaceArgs
        ///         {
        ///             Name = "example",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///         var exampleAlertRule = exampleAnalyticsWorkspace.Apply(exampleAnalyticsWorkspace =&gt; Output.Create(Azure.Sentinel.GetAlertRule.InvokeAsync(new Azure.Sentinel.GetAlertRuleArgs
        ///         {
        ///             Name = "existing",
        ///             LogAnalyticsWorkspaceId = exampleAnalyticsWorkspace.Id,
        ///         })));
        ///         this.Id = exampleAlertRule.Apply(exampleAlertRule =&gt; exampleAlertRule.Id);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAlertRuleResult> InvokeAsync(GetAlertRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAlertRuleResult>("azure:sentinel/getAlertRule:getAlertRule", args ?? new GetAlertRuleArgs(), options.WithVersion());

        public static Output<GetAlertRuleResult> Invoke(GetAlertRuleOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.LogAnalyticsWorkspaceId.Box(),
                args.Name.Box()
            ).Apply(a => {
                    var args = new GetAlertRuleArgs();
                    a[0].Set(args, nameof(args.LogAnalyticsWorkspaceId));
                    a[1].Set(args, nameof(args.Name));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetAlertRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Log Analytics Workspace this Sentinel Alert Rule belongs to.
        /// </summary>
        [Input("logAnalyticsWorkspaceId", required: true)]
        public string LogAnalyticsWorkspaceId { get; set; } = null!;

        /// <summary>
        /// The name which should be used for this Sentinel Alert Rule.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetAlertRuleArgs()
        {
        }
    }

    public sealed class GetAlertRuleOutputArgs
    {
        /// <summary>
        /// The ID of the Log Analytics Workspace this Sentinel Alert Rule belongs to.
        /// </summary>
        [Input("logAnalyticsWorkspaceId", required: true)]
        public Input<string> LogAnalyticsWorkspaceId { get; set; } = null!;

        /// <summary>
        /// The name which should be used for this Sentinel Alert Rule.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetAlertRuleOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAlertRuleResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string LogAnalyticsWorkspaceId;
        public readonly string Name;

        [OutputConstructor]
        private GetAlertRuleResult(
            string id,

            string logAnalyticsWorkspaceId,

            string name)
        {
            Id = id;
            LogAnalyticsWorkspaceId = logAnalyticsWorkspaceId;
            Name = name;
        }
    }
}

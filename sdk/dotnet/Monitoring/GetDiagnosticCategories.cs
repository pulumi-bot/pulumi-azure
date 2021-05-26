// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring
{
    public static class GetDiagnosticCategories
    {
        /// <summary>
        /// Use this data source to access information about the Monitor Diagnostics Categories supported by an existing Resource.
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
        ///         var exampleKeyVault = Output.Create(Azure.KeyVault.GetKeyVault.InvokeAsync(new Azure.KeyVault.GetKeyVaultArgs
        ///         {
        ///             Name = azurerm_key_vault.Example.Name,
        ///             ResourceGroupName = azurerm_key_vault.Example.Resource_group_name,
        ///         }));
        ///         var exampleDiagnosticCategories = exampleKeyVault.Apply(exampleKeyVault =&gt; Output.Create(Azure.Monitoring.GetDiagnosticCategories.InvokeAsync(new Azure.Monitoring.GetDiagnosticCategoriesArgs
        ///         {
        ///             ResourceId = exampleKeyVault.Id,
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDiagnosticCategoriesResult> InvokeAsync(GetDiagnosticCategoriesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDiagnosticCategoriesResult>("azure:monitoring/getDiagnosticCategories:getDiagnosticCategories", args ?? new GetDiagnosticCategoriesArgs(), options.WithVersion());

        public static Output<GetDiagnosticCategoriesResult> Apply(GetDiagnosticCategoriesApplyArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.ResourceId.Box()
            ).Apply(a => {
                    var args = new GetDiagnosticCategoriesArgs();
                    a[0].Set(args, nameof(args.ResourceId));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetDiagnosticCategoriesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of an existing Resource which Monitor Diagnostics Categories should be retrieved for.
        /// </summary>
        [Input("resourceId", required: true)]
        public string ResourceId { get; set; } = null!;

        public GetDiagnosticCategoriesArgs()
        {
        }
    }

    public sealed class GetDiagnosticCategoriesApplyArgs
    {
        /// <summary>
        /// The ID of an existing Resource which Monitor Diagnostics Categories should be retrieved for.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        public GetDiagnosticCategoriesApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDiagnosticCategoriesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the Log Categories supported for this Resource.
        /// </summary>
        public readonly ImmutableArray<string> Logs;
        /// <summary>
        /// A list of the Metric Categories supported for this Resource.
        /// </summary>
        public readonly ImmutableArray<string> Metrics;
        public readonly string ResourceId;

        [OutputConstructor]
        private GetDiagnosticCategoriesResult(
            string id,

            ImmutableArray<string> logs,

            ImmutableArray<string> metrics,

            string resourceId)
        {
            Id = id;
            Logs = logs;
            Metrics = metrics;
            ResourceId = resourceId;
        }
    }
}

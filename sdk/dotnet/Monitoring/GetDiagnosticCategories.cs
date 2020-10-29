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
        /// </summary>
        public static Task<GetDiagnosticCategoriesResult> InvokeAsync(GetDiagnosticCategoriesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDiagnosticCategoriesResult>("azure:monitoring/getDiagnosticCategories:getDiagnosticCategories", args ?? new GetDiagnosticCategoriesArgs(), options.WithVersion());
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

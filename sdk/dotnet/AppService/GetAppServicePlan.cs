// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService
{
    public static class GetAppServicePlan
    {
        /// <summary>
        /// Use this data source to access information about an existing App Service Plan (formerly known as a `Server Farm`).
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
        ///         var example = Output.Create(Azure.AppService.GetAppServicePlan.InvokeAsync(new Azure.AppService.GetAppServicePlanArgs
        ///         {
        ///             Name = "search-app-service-plan",
        ///             ResourceGroupName = "search-service",
        ///         }));
        ///         this.AppServicePlanId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("appServicePlanId")]
        ///     public Output&lt;string&gt; AppServicePlanId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppServicePlanResult> InvokeAsync(GetAppServicePlanArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppServicePlanResult>("azure:appservice/getAppServicePlan:getAppServicePlan", args ?? new GetAppServicePlanArgs(), options.WithVersion());
    }


    public sealed class GetAppServicePlanArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the App Service Plan.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the App Service Plan exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAppServicePlanArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppServicePlanResult
    {
        /// <summary>
        /// The ID of the App Service Environment where the App Service Plan is located.
        /// </summary>
        public readonly string AppServiceEnvironmentId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A flag that indicates if it's a xenon plan (support for Windows Container)
        /// </summary>
        public readonly bool IsXenon;
        /// <summary>
        /// The Operating System type of the App Service Plan
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The Azure location where the App Service Plan exists
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan.
        /// </summary>
        public readonly int MaximumElasticWorkerCount;
        /// <summary>
        /// The maximum number of workers supported with the App Service Plan's sku.
        /// </summary>
        public readonly int MaximumNumberOfWorkers;
        public readonly string Name;
        /// <summary>
        /// Can Apps assigned to this App Service Plan be scaled independently?
        /// </summary>
        public readonly bool PerSiteScaling;
        /// <summary>
        /// Is this App Service Plan `Reserved`?
        /// </summary>
        public readonly bool Reserved;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A `sku` block as documented below.
        /// </summary>
        public readonly Outputs.GetAppServicePlanSkuResult Sku;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetAppServicePlanResult(
            string appServiceEnvironmentId,

            string id,

            bool isXenon,

            string kind,

            string location,

            int maximumElasticWorkerCount,

            int maximumNumberOfWorkers,

            string name,

            bool perSiteScaling,

            bool reserved,

            string resourceGroupName,

            Outputs.GetAppServicePlanSkuResult sku,

            ImmutableDictionary<string, string> tags)
        {
            AppServiceEnvironmentId = appServiceEnvironmentId;
            Id = id;
            IsXenon = isXenon;
            Kind = kind;
            Location = location;
            MaximumElasticWorkerCount = maximumElasticWorkerCount;
            MaximumNumberOfWorkers = maximumNumberOfWorkers;
            Name = name;
            PerSiteScaling = perSiteScaling;
            Reserved = reserved;
            ResourceGroupName = resourceGroupName;
            Sku = sku;
            Tags = tags;
        }
    }
}

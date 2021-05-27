// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring
{
    public static class GetLogProfile
    {
        /// <summary>
        /// Use this data source to access the properties of a Log Profile.
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
        ///         var example = Output.Create(Azure.Monitoring.GetLogProfile.InvokeAsync(new Azure.Monitoring.GetLogProfileArgs
        ///         {
        ///             Name = "test-logprofile",
        ///         }));
        ///         this.LogProfileStorageAccountId = example.Apply(example =&gt; example.StorageAccountId);
        ///     }
        /// 
        ///     [Output("logProfileStorageAccountId")]
        ///     public Output&lt;string&gt; LogProfileStorageAccountId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetLogProfileResult> InvokeAsync(GetLogProfileArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLogProfileResult>("azure:monitoring/getLogProfile:getLogProfile", args ?? new GetLogProfileArgs(), options.WithVersion());

        public static Output<GetLogProfileResult> Apply(GetLogProfileApplyArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Name.Box()
            ).Apply(a => {
                    var args = new GetLogProfileArgs();
                    a[0].Set(args, nameof(args.Name));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetLogProfileArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the Name of the Log Profile.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetLogProfileArgs()
        {
        }
    }

    public sealed class GetLogProfileApplyArgs
    {
        /// <summary>
        /// Specifies the Name of the Log Profile.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetLogProfileApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLogProfileResult
    {
        /// <summary>
        /// List of categories of the logs.
        /// </summary>
        public readonly ImmutableArray<string> Categories;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of regions for which Activity Log events are stored or streamed.
        /// </summary>
        public readonly ImmutableArray<string> Locations;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetLogProfileRetentionPolicyResult> RetentionPolicies;
        /// <summary>
        /// The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to.
        /// </summary>
        public readonly string ServicebusRuleId;
        /// <summary>
        /// The resource id of the storage account in which the Activity Log is stored.
        /// </summary>
        public readonly string StorageAccountId;

        [OutputConstructor]
        private GetLogProfileResult(
            ImmutableArray<string> categories,

            string id,

            ImmutableArray<string> locations,

            string name,

            ImmutableArray<Outputs.GetLogProfileRetentionPolicyResult> retentionPolicies,

            string servicebusRuleId,

            string storageAccountId)
        {
            Categories = categories;
            Id = id;
            Locations = locations;
            Name = name;
            RetentionPolicies = retentionPolicies;
            ServicebusRuleId = servicebusRuleId;
            StorageAccountId = storageAccountId;
        }
    }
}

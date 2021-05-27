// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DatabaseMigration
{
    public static class GetService
    {
        /// <summary>
        /// Use this data source to access information about an existing Database Migration Service.
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
        ///         var example = Output.Create(Azure.DatabaseMigration.GetService.InvokeAsync(new Azure.DatabaseMigration.GetServiceArgs
        ///         {
        ///             Name = "example-dms",
        ///             ResourceGroupName = "example-rg",
        ///         }));
        ///         this.AzurermDmsId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("azurermDmsId")]
        ///     public Output&lt;string&gt; AzurermDmsId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServiceResult> InvokeAsync(GetServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceResult>("azure:databasemigration/getService:getService", args ?? new GetServiceArgs(), options.WithVersion());

        public static Output<GetServiceResult> Invoke(GetServiceOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Name.Box(),
                args.ResourceGroupName.Box()
            ).Apply(a => {
                    var args = new GetServiceArgs();
                    a[0].Set(args, nameof(args.Name));
                    a[1].Set(args, nameof(args.ResourceGroupName));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetServiceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the database migration service.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the Name of the Resource Group within which the database migration service exists
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetServiceArgs()
        {
        }
    }

    public sealed class GetServiceOutputArgs
    {
        /// <summary>
        /// Specify the name of the database migration service.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the Name of the Resource Group within which the database migration service exists
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetServiceOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServiceResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Azure location where the resource exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The sku name of database migration service.
        /// </summary>
        public readonly string SkuName;
        /// <summary>
        /// The ID of the virtual subnet resource to which the database migration service exists.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// A mapping of tags to assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetServiceResult(
            string id,

            string location,

            string name,

            string resourceGroupName,

            string skuName,

            string subnetId,

            ImmutableDictionary<string, string> tags)
        {
            Id = id;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            SkuName = skuName;
            SubnetId = subnetId;
            Tags = tags;
        }
    }
}

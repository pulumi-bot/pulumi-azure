// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute
{
    public static class GetSharedImageGallery
    {
        /// <summary>
        /// Use this data source to access information about an existing Shared Image Gallery.
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
        ///         var example = Output.Create(Azure.Compute.GetSharedImageGallery.InvokeAsync(new Azure.Compute.GetSharedImageGalleryArgs
        ///         {
        ///             Name = "my-image-gallery",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSharedImageGalleryResult> InvokeAsync(GetSharedImageGalleryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSharedImageGalleryResult>("azure:compute/getSharedImageGallery:getSharedImageGallery", args ?? new GetSharedImageGalleryArgs(), options.WithVersion());
    }


    public sealed class GetSharedImageGalleryArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Shared Image Gallery.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Shared Image Gallery exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetSharedImageGalleryArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSharedImageGalleryResult
    {
        /// <summary>
        /// A description for the Shared Image Gallery.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A mapping of tags which are assigned to the Shared Image Gallery.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The unique name assigned to the Shared Image Gallery.
        /// </summary>
        public readonly string UniqueName;

        [OutputConstructor]
        private GetSharedImageGalleryResult(
            string description,

            string id,

            string location,

            string name,

            string resourceGroupName,

            ImmutableDictionary<string, string> tags,

            string uniqueName)
        {
            Description = description;
            Id = id;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
            UniqueName = uniqueName;
        }
    }
}

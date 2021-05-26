// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute
{
    public static class GetSharedImageVersions
    {
        /// <summary>
        /// Use this data source to access information about existing Versions of a Shared Image within a Shared Image Gallery.
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
        ///         var example = Output.Create(Azure.Compute.GetSharedImageVersions.InvokeAsync(new Azure.Compute.GetSharedImageVersionsArgs
        ///         {
        ///             GalleryName = "my-image-gallery",
        ///             ImageName = "my-image",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSharedImageVersionsResult> InvokeAsync(GetSharedImageVersionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSharedImageVersionsResult>("azure:compute/getSharedImageVersions:getSharedImageVersions", args ?? new GetSharedImageVersionsArgs(), options.WithVersion());

        public static Output<GetSharedImageVersionsResult> Apply(GetSharedImageVersionsApplyArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.GalleryName.Box(),
                args.ImageName.Box(),
                args.ResourceGroupName.Box(),
                args.TagsFilter.Box()
            ).Apply(a => {
                    var args = new GetSharedImageVersionsArgs();
                    a[0].Set(args, nameof(args.GalleryName));
                    a[1].Set(args, nameof(args.ImageName));
                    a[2].Set(args, nameof(args.ResourceGroupName));
                    a[3].Set(args, nameof(args.TagsFilter));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetSharedImageVersionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Shared Image in which the Shared Image exists.
        /// </summary>
        [Input("galleryName", required: true)]
        public string GalleryName { get; set; } = null!;

        /// <summary>
        /// The name of the Shared Image in which this Version exists.
        /// </summary>
        [Input("imageName", required: true)]
        public string ImageName { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Shared Image Gallery exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("tagsFilter")]
        private Dictionary<string, string>? _tagsFilter;

        /// <summary>
        /// A mapping of tags to filter the list of images against.
        /// </summary>
        public Dictionary<string, string> TagsFilter
        {
            get => _tagsFilter ?? (_tagsFilter = new Dictionary<string, string>());
            set => _tagsFilter = value;
        }

        public GetSharedImageVersionsArgs()
        {
        }
    }

    public sealed class GetSharedImageVersionsApplyArgs
    {
        /// <summary>
        /// The name of the Shared Image in which the Shared Image exists.
        /// </summary>
        [Input("galleryName", required: true)]
        public Input<string> GalleryName { get; set; } = null!;

        /// <summary>
        /// The name of the Shared Image in which this Version exists.
        /// </summary>
        [Input("imageName", required: true)]
        public Input<string> ImageName { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Shared Image Gallery exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tagsFilter")]
        private InputMap<string>? _tagsFilter;

        /// <summary>
        /// A mapping of tags to filter the list of images against.
        /// </summary>
        public InputMap<string> TagsFilter
        {
            get => _tagsFilter ?? (_tagsFilter = new InputMap<string>());
            set => _tagsFilter = value;
        }

        public GetSharedImageVersionsApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSharedImageVersionsResult
    {
        public readonly string GalleryName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ImageName;
        /// <summary>
        /// An `images` block as defined below:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSharedImageVersionsImageResult> Images;
        public readonly string ResourceGroupName;
        public readonly ImmutableDictionary<string, string>? TagsFilter;

        [OutputConstructor]
        private GetSharedImageVersionsResult(
            string galleryName,

            string id,

            string imageName,

            ImmutableArray<Outputs.GetSharedImageVersionsImageResult> images,

            string resourceGroupName,

            ImmutableDictionary<string, string>? tagsFilter)
        {
            GalleryName = galleryName;
            Id = id;
            ImageName = imageName;
            Images = images;
            ResourceGroupName = resourceGroupName;
            TagsFilter = tagsFilter;
        }
    }
}

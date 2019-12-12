// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about an existing Image.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/image.html.markdown.
        /// </summary>
        public static Task<GetImageResult> GetImage(GetImageArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetImageResult>("azure:compute/getImage:getImage", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetImageArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Image.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Regex pattern of the image to match.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// The Name of the Resource Group where this Image exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// By default when matching by regex, images are sorted by name in ascending order and the first match is chosen, to sort descending, set this flag.
        /// </summary>
        [Input("sortDescending")]
        public bool? SortDescending { get; set; }

        public GetImageArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetImageResult
    {
        /// <summary>
        /// a collection of `data_disk` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetImageDataDisksResult> DataDisks;
        /// <summary>
        /// the Azure Location where this Image exists.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// the name of the Image.
        /// </summary>
        public readonly string? Name;
        public readonly string? NameRegex;
        /// <summary>
        /// a `os_disk` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetImageOsDisksResult> OsDisks;
        public readonly string ResourceGroupName;
        public readonly bool? SortDescending;
        /// <summary>
        /// a mapping of tags to assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// is zone resiliency enabled?
        /// </summary>
        public readonly bool ZoneResilient;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetImageResult(
            ImmutableArray<Outputs.GetImageDataDisksResult> dataDisks,
            string location,
            string? name,
            string? nameRegex,
            ImmutableArray<Outputs.GetImageOsDisksResult> osDisks,
            string resourceGroupName,
            bool? sortDescending,
            ImmutableDictionary<string, string> tags,
            bool zoneResilient,
            string id)
        {
            DataDisks = dataDisks;
            Location = location;
            Name = name;
            NameRegex = nameRegex;
            OsDisks = osDisks;
            ResourceGroupName = resourceGroupName;
            SortDescending = sortDescending;
            Tags = tags;
            ZoneResilient = zoneResilient;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetImageDataDisksResult
    {
        /// <summary>
        /// the URI in Azure storage of the blob used to create the image.
        /// </summary>
        public readonly string BlobUri;
        /// <summary>
        /// the caching mode for the Data Disk, such as `ReadWrite`, `ReadOnly`, or `None`.
        /// </summary>
        public readonly string Caching;
        /// <summary>
        /// the logical unit number of the data disk.
        /// </summary>
        public readonly int Lun;
        /// <summary>
        /// the ID of the Managed Disk used as the Data Disk Image.
        /// </summary>
        public readonly string ManagedDiskId;
        /// <summary>
        /// the size of this Data Disk in GB.
        /// </summary>
        public readonly int SizeGb;

        [OutputConstructor]
        private GetImageDataDisksResult(
            string blobUri,
            string caching,
            int lun,
            string managedDiskId,
            int sizeGb)
        {
            BlobUri = blobUri;
            Caching = caching;
            Lun = lun;
            ManagedDiskId = managedDiskId;
            SizeGb = sizeGb;
        }
    }

    [OutputType]
    public sealed class GetImageOsDisksResult
    {
        /// <summary>
        /// the URI in Azure storage of the blob used to create the image.
        /// </summary>
        public readonly string BlobUri;
        /// <summary>
        /// the caching mode for the Data Disk, such as `ReadWrite`, `ReadOnly`, or `None`.
        /// </summary>
        public readonly string Caching;
        /// <summary>
        /// the ID of the Managed Disk used as the Data Disk Image.
        /// </summary>
        public readonly string ManagedDiskId;
        /// <summary>
        /// the State of the OS used in the Image, such as `Generalized`.
        /// </summary>
        public readonly string OsState;
        /// <summary>
        /// the type of Operating System used on the OS Disk. such as `Linux` or `Windows`.
        /// </summary>
        public readonly string OsType;
        /// <summary>
        /// the size of this Data Disk in GB.
        /// </summary>
        public readonly int SizeGb;

        [OutputConstructor]
        private GetImageOsDisksResult(
            string blobUri,
            string caching,
            string managedDiskId,
            string osState,
            string osType,
            int sizeGb)
        {
            BlobUri = blobUri;
            Caching = caching;
            ManagedDiskId = managedDiskId;
            OsState = osState;
            OsType = osType;
            SizeGb = sizeGb;
        }
    }
    }
}

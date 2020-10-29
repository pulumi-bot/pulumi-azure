// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataShare
{
    public static class GetDatasetBlobStorage
    {
        /// <summary>
        /// Use this data source to access information about an existing Data Share Blob Storage Dataset.
        /// </summary>
        public static Task<GetDatasetBlobStorageResult> InvokeAsync(GetDatasetBlobStorageArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatasetBlobStorageResult>("azure:datashare/getDatasetBlobStorage:getDatasetBlobStorage", args ?? new GetDatasetBlobStorageArgs(), options.WithVersion());
    }


    public sealed class GetDatasetBlobStorageArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Data Share in which this Data Share Blob Storage Dataset should be created.
        /// </summary>
        [Input("dataShareId", required: true)]
        public string DataShareId { get; set; } = null!;

        /// <summary>
        /// The name of this Data Share Blob Storage Dataset.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetDatasetBlobStorageArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatasetBlobStorageResult
    {
        /// <summary>
        /// The name of the storage account container to be shared with the receiver.
        /// </summary>
        public readonly string ContainerName;
        public readonly string DataShareId;
        /// <summary>
        /// The name of the Data Share Dataset.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The path of the file in the storage container to be shared with the receiver.
        /// </summary>
        public readonly string FilePath;
        /// <summary>
        /// The folder path of the file in the storage container to be shared with the receiver.
        /// </summary>
        public readonly string FolderPath;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the storage account to be shared with the receiver.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A `storage_account` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatasetBlobStorageStorageAccountResult> StorageAccounts;

        [OutputConstructor]
        private GetDatasetBlobStorageResult(
            string containerName,

            string dataShareId,

            string displayName,

            string filePath,

            string folderPath,

            string id,

            string name,

            ImmutableArray<Outputs.GetDatasetBlobStorageStorageAccountResult> storageAccounts)
        {
            ContainerName = containerName;
            DataShareId = dataShareId;
            DisplayName = displayName;
            FilePath = filePath;
            FolderPath = folderPath;
            Id = id;
            Name = name;
            StorageAccounts = storageAccounts;
        }
    }
}

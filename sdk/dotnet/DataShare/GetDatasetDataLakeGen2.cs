// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataShare
{
    public static class GetDatasetDataLakeGen2
    {
        /// <summary>
        /// Use this data source to access information about an existing Data Share Data Lake Gen2 Dataset.
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
        ///         var example = Output.Create(Azure.DataShare.GetDatasetDataLakeGen2.InvokeAsync(new Azure.DataShare.GetDatasetDataLakeGen2Args
        ///         {
        ///             Name = "example-dsdlg2ds",
        ///             ShareId = "example-share-id",
        ///         }));
        ///         this.Id = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDatasetDataLakeGen2Result> InvokeAsync(GetDatasetDataLakeGen2Args args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatasetDataLakeGen2Result>("azure:datashare/getDatasetDataLakeGen2:getDatasetDataLakeGen2", args ?? new GetDatasetDataLakeGen2Args(), options.WithVersion());
    }


    public sealed class GetDatasetDataLakeGen2Args : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this Data Share Data Lake Gen2 Dataset.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The resource ID of the Data Share where this Data Share Data Lake Gen2 Dataset should be created.
        /// </summary>
        [Input("shareId", required: true)]
        public string ShareId { get; set; } = null!;

        public GetDatasetDataLakeGen2Args()
        {
        }
    }


    [OutputType]
    public sealed class GetDatasetDataLakeGen2Result
    {
        /// <summary>
        /// The name of the Data Share Dataset.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The path of the file in the data lake file system to be shared with the receiver.
        /// </summary>
        public readonly string FilePath;
        /// <summary>
        /// The name of the data lake file system to be shared with the receiver.
        /// </summary>
        public readonly string FileSystemName;
        /// <summary>
        /// The folder path in the data lake file system to be shared with the receiver.
        /// </summary>
        public readonly string FolderPath;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string ShareId;
        /// <summary>
        /// The resource ID of the storage account of the data lake file system to be shared with the receiver.
        /// </summary>
        public readonly string StorageAccountId;

        [OutputConstructor]
        private GetDatasetDataLakeGen2Result(
            string displayName,

            string filePath,

            string fileSystemName,

            string folderPath,

            string id,

            string name,

            string shareId,

            string storageAccountId)
        {
            DisplayName = displayName;
            FilePath = filePath;
            FileSystemName = fileSystemName;
            FolderPath = folderPath;
            Id = id;
            Name = name;
            ShareId = shareId;
            StorageAccountId = storageAccountId;
        }
    }
}

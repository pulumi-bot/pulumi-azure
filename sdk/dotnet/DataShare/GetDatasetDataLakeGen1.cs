// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataShare
{
    public static class GetDatasetDataLakeGen1
    {
        /// <summary>
        /// Use this data source to access information about an existing DataShareDataLakeGen1Dataset.
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
        ///         var example = Output.Create(Azure.DataShare.GetDatasetDataLakeGen1.InvokeAsync(new Azure.DataShare.GetDatasetDataLakeGen1Args
        ///         {
        ///             Name = "example-dsdsdlg1",
        ///             DataShareId = "example-share-id",
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
        public static Task<GetDatasetDataLakeGen1Result> InvokeAsync(GetDatasetDataLakeGen1Args args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatasetDataLakeGen1Result>("azure:datashare/getDatasetDataLakeGen1:getDatasetDataLakeGen1", args ?? new GetDatasetDataLakeGen1Args(), options.WithVersion());

        public static Output<GetDatasetDataLakeGen1Result> Invoke(GetDatasetDataLakeGen1OutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.DataShareId.Box(),
                args.Name.Box()
            ).Apply(a => {
                    var args = new GetDatasetDataLakeGen1Args();
                    a[0].Set(args, nameof(args.DataShareId));
                    a[1].Set(args, nameof(args.Name));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetDatasetDataLakeGen1Args : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The resource ID of the Data Share where this Data Share Data Lake Gen1 Dataset should be created.
        /// </summary>
        [Input("dataShareId", required: true)]
        public string DataShareId { get; set; } = null!;

        /// <summary>
        /// The name of the Data Share Data Lake Gen1 Dataset.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetDatasetDataLakeGen1Args()
        {
        }
    }

    public sealed class GetDatasetDataLakeGen1OutputArgs
    {
        /// <summary>
        /// The resource ID of the Data Share where this Data Share Data Lake Gen1 Dataset should be created.
        /// </summary>
        [Input("dataShareId", required: true)]
        public Input<string> DataShareId { get; set; } = null!;

        /// <summary>
        /// The name of the Data Share Data Lake Gen1 Dataset.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetDatasetDataLakeGen1OutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatasetDataLakeGen1Result
    {
        /// <summary>
        /// The resource ID of the Data Lake Store to be shared with the receiver.
        /// </summary>
        public readonly string DataLakeStoreId;
        public readonly string DataShareId;
        /// <summary>
        /// The displayed name of the Data Share Dataset.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The file name of the data lake store to be shared with the receiver.
        /// </summary>
        public readonly string FileName;
        /// <summary>
        /// The folder path of the data lake store to be shared with the receiver.
        /// </summary>
        public readonly string FolderPath;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetDatasetDataLakeGen1Result(
            string dataLakeStoreId,

            string dataShareId,

            string displayName,

            string fileName,

            string folderPath,

            string id,

            string name)
        {
            DataLakeStoreId = dataLakeStoreId;
            DataShareId = dataShareId;
            DisplayName = displayName;
            FileName = fileName;
            FolderPath = folderPath;
            Id = id;
            Name = name;
        }
    }
}

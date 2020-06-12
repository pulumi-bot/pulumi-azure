// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Storage
{
    public static class GetStorageContainer
    {
        /// <summary>
        /// Use this data source to access information about an existing Storage Container.
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
        ///         var example = Output.Create(Azure.Storage.GetStorageContainer.InvokeAsync(new Azure.Storage.GetStorageContainerArgs
        ///         {
        ///             Name = "example-container-name",
        ///             StorageAccountName = "example-storage-account-name",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetStorageContainerResult> InvokeAsync(GetStorageContainerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStorageContainerResult>("azure:storage/getStorageContainer:getStorageContainer", args ?? new GetStorageContainerArgs(), options.WithVersion());
    }


    public sealed class GetStorageContainerArgs : Pulumi.InvokeArgs
    {
        [Input("metadata")]
        private Dictionary<string, string>? _metadata;

        /// <summary>
        /// A mapping of MetaData for this Container.
        /// </summary>
        public Dictionary<string, string> Metadata
        {
            get => _metadata ?? (_metadata = new Dictionary<string, string>());
            set => _metadata = value;
        }

        /// <summary>
        /// The name of the Container.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Storage Account where the Container exists.
        /// </summary>
        [Input("storageAccountName", required: true)]
        public string StorageAccountName { get; set; } = null!;

        public GetStorageContainerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetStorageContainerResult
    {
        /// <summary>
        /// The Access Level configured for this Container.
        /// </summary>
        public readonly string ContainerAccessType;
        /// <summary>
        /// Is there an Immutability Policy configured on this Storage Container?
        /// </summary>
        public readonly bool HasImmutabilityPolicy;
        /// <summary>
        /// Is there a Legal Hold configured on this Storage Container?
        /// </summary>
        public readonly bool HasLegalHold;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A mapping of MetaData for this Container.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metadata;
        public readonly string Name;
        /// <summary>
        /// The Resource Manager ID of this Storage Container.
        /// </summary>
        public readonly string ResourceManagerId;
        public readonly string StorageAccountName;

        [OutputConstructor]
        private GetStorageContainerResult(
            string containerAccessType,

            bool hasImmutabilityPolicy,

            bool hasLegalHold,

            string id,

            ImmutableDictionary<string, string> metadata,

            string name,

            string resourceManagerId,

            string storageAccountName)
        {
            ContainerAccessType = containerAccessType;
            HasImmutabilityPolicy = hasImmutabilityPolicy;
            HasLegalHold = hasLegalHold;
            Id = id;
            Metadata = metadata;
            Name = name;
            ResourceManagerId = resourceManagerId;
            StorageAccountName = storageAccountName;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataShare
{
    public static class GetDatasetKustoCluster
    {
        /// <summary>
        /// Use this data source to access information about an existing Data Share Kusto Cluster Dataset.
        /// </summary>
        public static Task<GetDatasetKustoClusterResult> InvokeAsync(GetDatasetKustoClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatasetKustoClusterResult>("azure:datashare/getDatasetKustoCluster:getDatasetKustoCluster", args ?? new GetDatasetKustoClusterArgs(), options.WithVersion());
    }


    public sealed class GetDatasetKustoClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this Data Share Kusto Cluster Dataset.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The resource ID of the Data Share where this Data Share Kusto Cluster Dataset should be created.
        /// </summary>
        [Input("shareId", required: true)]
        public string ShareId { get; set; } = null!;

        public GetDatasetKustoClusterArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatasetKustoClusterResult
    {
        /// <summary>
        /// The name of the Data Share Dataset.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The resource ID of the Kusto Cluster to be shared with the receiver.
        /// </summary>
        public readonly string KustoClusterId;
        /// <summary>
        /// The location of the Kusto Cluster.
        /// </summary>
        public readonly string KustoClusterLocation;
        public readonly string Name;
        public readonly string ShareId;

        [OutputConstructor]
        private GetDatasetKustoClusterResult(
            string displayName,

            string id,

            string kustoClusterId,

            string kustoClusterLocation,

            string name,

            string shareId)
        {
            DisplayName = displayName;
            Id = id;
            KustoClusterId = kustoClusterId;
            KustoClusterLocation = kustoClusterLocation;
            Name = name;
            ShareId = shareId;
        }
    }
}

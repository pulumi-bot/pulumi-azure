// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataLake
{
    public static class GetStore
    {
        /// <summary>
        /// Use this data source to access information about an existing Data Lake Store.
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
        ///         var example = Output.Create(Azure.DataLake.GetStore.InvokeAsync(new Azure.DataLake.GetStoreArgs
        ///         {
        ///             Name = "testdatalake",
        ///             ResourceGroupName = "testdatalake",
        ///         }));
        ///         this.DataLakeStoreId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("dataLakeStoreId")]
        ///     public Output&lt;string&gt; DataLakeStoreId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetStoreResult> InvokeAsync(GetStoreArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStoreResult>("azure:datalake/getStore:getStore", args ?? new GetStoreArgs(), options.WithVersion());
    }


    public sealed class GetStoreArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Data Lake Store.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the Data Lake Store exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetStoreArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetStoreResult
    {
        /// <summary>
        /// the Encryption State of this Data Lake Store Account, such as `Enabled` or `Disabled`.
        /// </summary>
        public readonly string EncryptionState;
        /// <summary>
        /// the Encryption Type used for this Data Lake Store Account.
        /// </summary>
        public readonly string EncryptionType;
        /// <summary>
        /// are Azure Service IP's allowed through the firewall?
        /// </summary>
        public readonly string FirewallAllowAzureIps;
        /// <summary>
        /// the state of the firewall, such as `Enabled` or `Disabled`.
        /// </summary>
        public readonly string FirewallState;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A mapping of tags to assign to the Data Lake Store.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// Current monthly commitment tier for the account.
        /// </summary>
        public readonly string Tier;

        [OutputConstructor]
        private GetStoreResult(
            string encryptionState,

            string encryptionType,

            string firewallAllowAzureIps,

            string firewallState,

            string id,

            string location,

            string name,

            string resourceGroupName,

            ImmutableDictionary<string, string> tags,

            string tier)
        {
            EncryptionState = encryptionState;
            EncryptionType = encryptionType;
            FirewallAllowAzureIps = firewallAllowAzureIps;
            FirewallState = firewallState;
            Id = id;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
            Tier = tier;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Batch
{
    public static class GetAccount
    {
        /// <summary>
        /// Use this data source to access information about an existing Batch Account.
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
        ///         var example = Output.Create(Azure.Batch.GetAccount.InvokeAsync(new Azure.Batch.GetAccountArgs
        ///         {
        ///             Name = "testbatchaccount",
        ///             ResourceGroupName = "test",
        ///         }));
        ///         this.PoolAllocationMode = example.Apply(example =&gt; example.PoolAllocationMode);
        ///     }
        /// 
        ///     [Output("poolAllocationMode")]
        ///     public Output&lt;string&gt; PoolAllocationMode { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAccountResult> InvokeAsync(GetAccountArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccountResult>("azure:batch/getAccount:getAccount", args ?? new GetAccountArgs(), options.WithVersion());

        public static Output<GetAccountResult> Invoke(GetAccountOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Name.Box(),
                args.ResourceGroupName.Box()
            ).Apply(a => {
                    var args = new GetAccountArgs();
                    a[0].Set(args, nameof(args.Name));
                    a[1].Set(args, nameof(args.ResourceGroupName));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Batch account.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where this Batch account exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAccountArgs()
        {
        }
    }

    public sealed class GetAccountOutputArgs
    {
        /// <summary>
        /// The name of the Batch account.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where this Batch account exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetAccountOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAccountResult
    {
        /// <summary>
        /// The account endpoint used to interact with the Batch service.
        /// </summary>
        public readonly string AccountEndpoint;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The `key_vault_reference` block that describes the Azure KeyVault reference to use when deploying the Azure Batch account using the `UserSubscription` pool allocation mode.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccountKeyVaultReferenceResult> KeyVaultReferences;
        /// <summary>
        /// The Azure Region in which this Batch account exists.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The Batch account name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The pool allocation mode configured for this Batch account.
        /// </summary>
        public readonly string PoolAllocationMode;
        /// <summary>
        /// The Batch account primary access key.
        /// </summary>
        public readonly string PrimaryAccessKey;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The Batch account secondary access key.
        /// </summary>
        public readonly string SecondaryAccessKey;
        /// <summary>
        /// The ID of the Storage Account used for this Batch account.
        /// </summary>
        public readonly string StorageAccountId;
        /// <summary>
        /// A map of tags assigned to the Batch account.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetAccountResult(
            string accountEndpoint,

            string id,

            ImmutableArray<Outputs.GetAccountKeyVaultReferenceResult> keyVaultReferences,

            string location,

            string name,

            string poolAllocationMode,

            string primaryAccessKey,

            string resourceGroupName,

            string secondaryAccessKey,

            string storageAccountId,

            ImmutableDictionary<string, string> tags)
        {
            AccountEndpoint = accountEndpoint;
            Id = id;
            KeyVaultReferences = keyVaultReferences;
            Location = location;
            Name = name;
            PoolAllocationMode = poolAllocationMode;
            PrimaryAccessKey = primaryAccessKey;
            ResourceGroupName = resourceGroupName;
            SecondaryAccessKey = secondaryAccessKey;
            StorageAccountId = storageAccountId;
            Tags = tags;
        }
    }
}

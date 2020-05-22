// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.KeyVault
{
    public static class GetKey
    {
        /// <summary>
        /// Use this data source to access information about an existing Key Vault Key.
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
        ///         var example = Output.Create(Azure.KeyVault.GetKey.InvokeAsync(new Azure.KeyVault.GetKeyArgs
        ///         {
        ///             Name = "secret-sauce",
        ///             KeyVaultId = data.Azurerm_key_vault.Existing.Id,
        ///         }));
        ///         this.KeyType = example.Apply(example =&gt; example.KeyType);
        ///     }
        /// 
        ///     [Output("keyType")]
        ///     public Output&lt;string&gt; KeyType { get; set; }
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetKeyResult> InvokeAsync(GetKeyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKeyResult>("azure:keyvault/getKey:getKey", args ?? new GetKeyArgs(), options.WithVersion());
    }


    public sealed class GetKeyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the ID of the Key Vault instance where the Secret resides, available on the `azure.keyvault.KeyVault` Data Source / Resource. 
        /// </summary>
        [Input("keyVaultId", required: true)]
        public string KeyVaultId { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Key Vault Key.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetKeyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetKeyResult
    {
        /// <summary>
        /// The RSA public exponent of this Key Vault Key.
        /// </summary>
        public readonly string E;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of JSON web key operations assigned to this Key Vault Key
        /// </summary>
        public readonly ImmutableArray<string> KeyOpts;
        /// <summary>
        /// Specifies the Size of this Key Vault Key.
        /// </summary>
        public readonly int KeySize;
        /// <summary>
        /// Specifies the Key Type of this Key Vault Key
        /// </summary>
        public readonly string KeyType;
        public readonly string KeyVaultId;
        /// <summary>
        /// The RSA modulus of this Key Vault Key.
        /// </summary>
        public readonly string N;
        public readonly string Name;
        /// <summary>
        /// A mapping of tags assigned to this Key Vault Key.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The current version of the Key Vault Key.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetKeyResult(
            string e,

            string id,

            ImmutableArray<string> keyOpts,

            int keySize,

            string keyType,

            string keyVaultId,

            string n,

            string name,

            ImmutableDictionary<string, string> tags,

            string version)
        {
            E = e;
            Id = id;
            KeyOpts = keyOpts;
            KeySize = keySize;
            KeyType = keyType;
            KeyVaultId = keyVaultId;
            N = n;
            Name = name;
            Tags = tags;
            Version = version;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.KeyVault
{
    public static class GetKeyVault
    {
        /// <summary>
        /// Use this data source to access information about an existing Key Vault.
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
        ///         var example = Output.Create(Azure.KeyVault.GetKeyVault.InvokeAsync(new Azure.KeyVault.GetKeyVaultArgs
        ///         {
        ///             Name = "mykeyvault",
        ///             ResourceGroupName = "some-resource-group",
        ///         }));
        ///         this.VaultUri = example.Apply(example =&gt; example.VaultUri);
        ///     }
        /// 
        ///     [Output("vaultUri")]
        ///     public Output&lt;string&gt; VaultUri { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetKeyVaultResult> InvokeAsync(GetKeyVaultArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKeyVaultResult>("azure:keyvault/getKeyVault:getKeyVault", args ?? new GetKeyVaultArgs(), options.WithVersion());
    }


    public sealed class GetKeyVaultArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Key Vault.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Key Vault exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetKeyVaultArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetKeyVaultResult
    {
        /// <summary>
        /// One or more `access_policy` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKeyVaultAccessPolicyResult> AccessPolicies;
        /// <summary>
        /// Can Azure Virtual Machines retrieve certificates stored as secrets from the Key Vault?
        /// </summary>
        public readonly bool EnabledForDeployment;
        /// <summary>
        /// Can Azure Disk Encryption retrieve secrets from the Key Vault?
        /// </summary>
        public readonly bool EnabledForDiskEncryption;
        /// <summary>
        /// Can Azure Resource Manager retrieve secrets from the Key Vault?
        /// </summary>
        public readonly bool EnabledForTemplateDeployment;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure Region in which the Key Vault exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetKeyVaultNetworkAclResult> NetworkAcls;
        /// <summary>
        /// Is purge protection enabled on this Key Vault?
        /// </summary>
        public readonly bool PurgeProtectionEnabled;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The Name of the SKU used for this Key Vault.
        /// </summary>
        public readonly string SkuName;
        /// <summary>
        /// Is soft delete enabled on this Key Vault?
        /// </summary>
        public readonly bool SoftDeleteEnabled;
        /// <summary>
        /// A mapping of tags assigned to the Key Vault.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The Azure Active Directory Tenant ID used to authenticate requests for this Key Vault.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The URI of the vault for performing operations on keys and secrets.
        /// </summary>
        public readonly string VaultUri;

        [OutputConstructor]
        private GetKeyVaultResult(
            ImmutableArray<Outputs.GetKeyVaultAccessPolicyResult> accessPolicies,

            bool enabledForDeployment,

            bool enabledForDiskEncryption,

            bool enabledForTemplateDeployment,

            string id,

            string location,

            string name,

            ImmutableArray<Outputs.GetKeyVaultNetworkAclResult> networkAcls,

            bool purgeProtectionEnabled,

            string resourceGroupName,

            string skuName,

            bool softDeleteEnabled,

            ImmutableDictionary<string, string> tags,

            string tenantId,

            string vaultUri)
        {
            AccessPolicies = accessPolicies;
            EnabledForDeployment = enabledForDeployment;
            EnabledForDiskEncryption = enabledForDiskEncryption;
            EnabledForTemplateDeployment = enabledForTemplateDeployment;
            Id = id;
            Location = location;
            Name = name;
            NetworkAcls = networkAcls;
            PurgeProtectionEnabled = purgeProtectionEnabled;
            ResourceGroupName = resourceGroupName;
            SkuName = skuName;
            SoftDeleteEnabled = softDeleteEnabled;
            Tags = tags;
            TenantId = tenantId;
            VaultUri = vaultUri;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MSSql
{
    /// <summary>
    /// Manages the transparent data encryption configuration for a MSSQL Server
    /// 
    /// &gt; **NOTE:** Once transparent data encryption is enabled on a MS SQL instance, it is not possible to remove TDE. You will be able to switch between 'ServiceManaged' and 'CustomerManaged' keys, but will not be able to remove encryption. For safety when this resource is deleted, the TDE mode will automatically be set to 'ServiceManaged'. See `key_vault_uri` for more information on how to specify the key types. As SQL Server only supports a single configuration for encryption settings, this resource will replace the current encryption settings on the server.
    /// 
    /// &gt; **Note:** See [documentation](https://docs.microsoft.com/en-us/azure/azure-sql/database/transparent-data-encryption-byok-overview) for important information on how handle lifecycle management of the keys to prevent data lockout.
    /// 
    /// ## Example Usage
    /// ### With Service Managed Key
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "EastUs",
    ///         });
    ///         var exampleServer = new Azure.MSSql.Server("exampleServer", new Azure.MSSql.ServerArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Version = "12.0",
    ///             AdministratorLogin = "missadministrator",
    ///             AdministratorLoginPassword = "thisIsKat11",
    ///             MinimumTlsVersion = "1.2",
    ///             AzureadAdministrator = new Azure.MSSql.Inputs.ServerAzureadAdministratorArgs
    ///             {
    ///                 LoginUsername = "AzureAD Admin",
    ///                 ObjectId = "00000000-0000-0000-0000-000000000000",
    ///             },
    ///             ExtendedAuditingPolicy = new Azure.MSSql.Inputs.ServerExtendedAuditingPolicyArgs
    ///             {
    ///                 StorageEndpoint = azurerm_storage_account.Example.Primary_blob_endpoint,
    ///                 StorageAccountAccessKey = azurerm_storage_account.Example.Primary_access_key,
    ///                 StorageAccountAccessKeyIsSecondary = true,
    ///                 RetentionInDays = 6,
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "environment", "production" },
    ///             },
    ///         });
    ///         var exampleServerTransparentDataEncryption = new Azure.MSSql.ServerTransparentDataEncryption("exampleServerTransparentDataEncryption", new Azure.MSSql.ServerTransparentDataEncryptionArgs
    ///         {
    ///             ServerId = exampleServer.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### With Customer Managed Key
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var current = Output.Create(Azure.Core.GetClientConfig.InvokeAsync());
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "EastUs",
    ///         });
    ///         var exampleServer = new Azure.MSSql.Server("exampleServer", new Azure.MSSql.ServerArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Version = "12.0",
    ///             AdministratorLogin = "missadministrator",
    ///             AdministratorLoginPassword = "thisIsKat11",
    ///             MinimumTlsVersion = "1.2",
    ///             AzureadAdministrator = new Azure.MSSql.Inputs.ServerAzureadAdministratorArgs
    ///             {
    ///                 LoginUsername = "AzureAD Admin",
    ///                 ObjectId = "00000000-0000-0000-0000-000000000000",
    ///             },
    ///             ExtendedAuditingPolicy = new Azure.MSSql.Inputs.ServerExtendedAuditingPolicyArgs
    ///             {
    ///                 StorageEndpoint = azurerm_storage_account.Example.Primary_blob_endpoint,
    ///                 StorageAccountAccessKey = azurerm_storage_account.Example.Primary_access_key,
    ///                 StorageAccountAccessKeyIsSecondary = true,
    ///                 RetentionInDays = 6,
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "environment", "production" },
    ///             },
    ///             Identity = new Azure.MSSql.Inputs.ServerIdentityArgs
    ///             {
    ///                 Type = "SystemAssigned",
    ///             },
    ///         });
    ///         // Create a key vault with policies for the deployer to create a key &amp; SQL Server to wrap/unwrap/get key
    ///         var exampleKeyVault = new Azure.KeyVault.KeyVault("exampleKeyVault", new Azure.KeyVault.KeyVaultArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             EnabledForDiskEncryption = true,
    ///             TenantId = current.Apply(current =&gt; current.TenantId),
    ///             SoftDeleteRetentionDays = 7,
    ///             PurgeProtectionEnabled = false,
    ///             SkuName = "standard",
    ///             AccessPolicies = 
    ///             {
    ///                 new Azure.KeyVault.Inputs.KeyVaultAccessPolicyArgs
    ///                 {
    ///                     TenantId = current.Apply(current =&gt; current.TenantId),
    ///                     ObjectId = current.Apply(current =&gt; current.ObjectId),
    ///                     KeyPermissions = 
    ///                     {
    ///                         "Get",
    ///                         "List",
    ///                         "Create",
    ///                         "Delete",
    ///                         "Update",
    ///                         "Recover",
    ///                         "Purge",
    ///                     },
    ///                 },
    ///                 new Azure.KeyVault.Inputs.KeyVaultAccessPolicyArgs
    ///                 {
    ///                     TenantId = exampleServer.Identity.Apply(identity =&gt; identity.TenantId),
    ///                     ObjectId = exampleServer.Identity.Apply(identity =&gt; identity.PrincipalId),
    ///                     KeyPermissions = 
    ///                     {
    ///                         "Get",
    ///                         "WrapKey",
    ///                         "UnwrapKey",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///         var exampleKey = new Azure.KeyVault.Key("exampleKey", new Azure.KeyVault.KeyArgs
    ///         {
    ///             KeyVaultId = exampleKeyVault.Id,
    ///             KeyType = "RSA",
    ///             KeySize = 2048,
    ///             KeyOpts = 
    ///             {
    ///                 "unwrapKey",
    ///                 "wrapKey",
    ///             },
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 exampleKeyVault,
    ///             },
    ///         });
    ///         var exampleServerTransparentDataEncryption = new Azure.MSSql.ServerTransparentDataEncryption("exampleServerTransparentDataEncryption", new Azure.MSSql.ServerTransparentDataEncryptionArgs
    ///         {
    ///             ServerId = exampleServer.Id,
    ///             KeyVaultKeyId = exampleKey.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// SQL Server Transparent Data Encryption can be imported using the resource id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:mssql/serverTransparentDataEncryption:ServerTransparentDataEncryption example /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/encryptionProtector/current
    /// ```
    /// </summary>
    [AzureResourceType("azure:mssql/serverTransparentDataEncryption:ServerTransparentDataEncryption")]
    public partial class ServerTransparentDataEncryption : Pulumi.CustomResource
    {
        /// <summary>
        /// To use customer managed keys from Azure Key Vault, provide the AKV Key ID. To use service managed keys, omit this field.
        /// </summary>
        [Output("keyVaultKeyId")]
        public Output<string?> KeyVaultKeyId { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the MS SQL Server.
        /// </summary>
        [Output("serverId")]
        public Output<string> ServerId { get; private set; } = null!;


        /// <summary>
        /// Create a ServerTransparentDataEncryption resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerTransparentDataEncryption(string name, ServerTransparentDataEncryptionArgs args, CustomResourceOptions? options = null)
            : base("azure:mssql/serverTransparentDataEncryption:ServerTransparentDataEncryption", name, args ?? new ServerTransparentDataEncryptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerTransparentDataEncryption(string name, Input<string> id, ServerTransparentDataEncryptionState? state = null, CustomResourceOptions? options = null)
            : base("azure:mssql/serverTransparentDataEncryption:ServerTransparentDataEncryption", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServerTransparentDataEncryption resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerTransparentDataEncryption Get(string name, Input<string> id, ServerTransparentDataEncryptionState? state = null, CustomResourceOptions? options = null)
        {
            return new ServerTransparentDataEncryption(name, id, state, options);
        }
    }

    public sealed class ServerTransparentDataEncryptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// To use customer managed keys from Azure Key Vault, provide the AKV Key ID. To use service managed keys, omit this field.
        /// </summary>
        [Input("keyVaultKeyId")]
        public Input<string>? KeyVaultKeyId { get; set; }

        /// <summary>
        /// Specifies the name of the MS SQL Server.
        /// </summary>
        [Input("serverId", required: true)]
        public Input<string> ServerId { get; set; } = null!;

        public ServerTransparentDataEncryptionArgs()
        {
        }
    }

    public sealed class ServerTransparentDataEncryptionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// To use customer managed keys from Azure Key Vault, provide the AKV Key ID. To use service managed keys, omit this field.
        /// </summary>
        [Input("keyVaultKeyId")]
        public Input<string>? KeyVaultKeyId { get; set; }

        /// <summary>
        /// Specifies the name of the MS SQL Server.
        /// </summary>
        [Input("serverId")]
        public Input<string>? ServerId { get; set; }

        public ServerTransparentDataEncryptionState()
        {
        }
    }
}

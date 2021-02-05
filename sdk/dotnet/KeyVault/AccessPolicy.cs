// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.KeyVault
{
    /// <summary>
    /// Manages a Key Vault Access Policy.
    /// 
    /// &gt; **NOTE:** It's possible to define Key Vault Access Policies both within the `azure.keyvault.KeyVault` resource via the `access_policy` block and by using the `azure.keyvault.AccessPolicy` resource. However it's not possible to use both methods to manage Access Policies within a KeyVault, since there'll be conflicts.
    /// 
    /// &gt; **NOTE:** Azure permits a maximum of 1024 Access Policies per Key Vault - [more information can be found in this document](https://docs.microsoft.com/en-us/azure/key-vault/key-vault-secure-your-key-vault#data-plane-access-control).
    /// 
    /// ## Example Usage
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
    ///             Location = "West Europe",
    ///         });
    ///         var exampleKeyVault = new Azure.KeyVault.KeyVault("exampleKeyVault", new Azure.KeyVault.KeyVaultArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             TenantId = current.Apply(current =&gt; current.TenantId),
    ///             SkuName = "premium",
    ///         });
    ///         var exampleAccessPolicy = new Azure.KeyVault.AccessPolicy("exampleAccessPolicy", new Azure.KeyVault.AccessPolicyArgs
    ///         {
    ///             KeyVaultId = exampleKeyVault.Id,
    ///             TenantId = current.Apply(current =&gt; current.TenantId),
    ///             ObjectId = current.Apply(current =&gt; current.ObjectId),
    ///             KeyPermissions = 
    ///             {
    ///                 "get",
    ///             },
    ///             SecretPermissions = 
    ///             {
    ///                 "get",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Key Vault Access Policies can be imported using the Resource ID of the Key Vault, plus some additional metadata. If both an `object_id` and `application_id` are specified, then the Access Policy can be imported using the following code
    /// 
    /// ```sh
    ///  $ pulumi import azure:keyvault/accessPolicy:AccessPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.KeyVault/vaults/test-vault/objectId/11111111-1111-1111-1111-111111111111/applicationId/22222222-2222-2222-2222-222222222222
    /// ```
    /// 
    ///  where `11111111-1111-1111-1111-111111111111` is the `object_id` and `22222222-2222-2222-2222-222222222222` is the `application_id`. --- Access Policies with an `object_id` but no `application_id` can be imported using the following command
    /// 
    /// ```sh
    ///  $ pulumi import azure:keyvault/accessPolicy:AccessPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.KeyVault/vaults/test-vault/objectId/11111111-1111-1111-1111-111111111111
    /// ```
    /// 
    ///  where `11111111-1111-1111-1111-111111111111` is the `object_id`.
    /// </summary>
    [AzureResourceType("azure:keyvault/accessPolicy:AccessPolicy")]
    public partial class AccessPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The object ID of an Application in Azure Active Directory.
        /// </summary>
        [Output("applicationId")]
        public Output<string?> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// List of certificate permissions, must be one or more from
        /// the following: `backup`, `create`, `delete`, `deleteissuers`, `get`, `getissuers`, `import`, `list`, `listissuers`,
        /// `managecontacts`, `manageissuers`, `purge`, `recover`, `restore`, `setissuers` and `update`.
        /// </summary>
        [Output("certificatePermissions")]
        public Output<ImmutableArray<string>> CertificatePermissions { get; private set; } = null!;

        /// <summary>
        /// List of key permissions, must be one or more from
        /// the following: `backup`, `create`, `decrypt`, `delete`, `encrypt`, `get`, `import`, `list`, `purge`,
        /// `recover`, `restore`, `sign`, `unwrapKey`, `update`, `verify` and `wrapKey`.
        /// </summary>
        [Output("keyPermissions")]
        public Output<ImmutableArray<string>> KeyPermissions { get; private set; } = null!;

        /// <summary>
        /// Specifies the id of the Key Vault resource. Changing this
        /// forces a new resource to be created.
        /// </summary>
        [Output("keyVaultId")]
        public Output<string> KeyVaultId { get; private set; } = null!;

        /// <summary>
        /// The object ID of a user, service principal or security
        /// group in the Azure Active Directory tenant for the vault. The object ID must
        /// be unique for the list of access policies. Changing this forces a new resource
        /// to be created.
        /// </summary>
        [Output("objectId")]
        public Output<string> ObjectId { get; private set; } = null!;

        /// <summary>
        /// List of secret permissions, must be one or more
        /// from the following: `backup`, `delete`, `get`, `list`, `purge`, `recover`, `restore` and `set`.
        /// </summary>
        [Output("secretPermissions")]
        public Output<ImmutableArray<string>> SecretPermissions { get; private set; } = null!;

        /// <summary>
        /// List of storage permissions, must be one or more from the following: `backup`, `delete`, `deletesas`, `get`, `getsas`, `list`, `listsas`, `purge`, `recover`, `regeneratekey`, `restore`, `set`, `setsas` and `update`.
        /// </summary>
        [Output("storagePermissions")]
        public Output<ImmutableArray<string>> StoragePermissions { get; private set; } = null!;

        /// <summary>
        /// The Azure Active Directory tenant ID that should be used
        /// for authenticating requests to the key vault. Changing this forces a new resource
        /// to be created.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a AccessPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessPolicy(string name, AccessPolicyArgs args, CustomResourceOptions? options = null)
            : base("azure:keyvault/accessPolicy:AccessPolicy", name, args ?? new AccessPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessPolicy(string name, Input<string> id, AccessPolicyState? state = null, CustomResourceOptions? options = null)
            : base("azure:keyvault/accessPolicy:AccessPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccessPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessPolicy Get(string name, Input<string> id, AccessPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new AccessPolicy(name, id, state, options);
        }
    }

    public sealed class AccessPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The object ID of an Application in Azure Active Directory.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        [Input("certificatePermissions")]
        private InputList<string>? _certificatePermissions;

        /// <summary>
        /// List of certificate permissions, must be one or more from
        /// the following: `backup`, `create`, `delete`, `deleteissuers`, `get`, `getissuers`, `import`, `list`, `listissuers`,
        /// `managecontacts`, `manageissuers`, `purge`, `recover`, `restore`, `setissuers` and `update`.
        /// </summary>
        public InputList<string> CertificatePermissions
        {
            get => _certificatePermissions ?? (_certificatePermissions = new InputList<string>());
            set => _certificatePermissions = value;
        }

        [Input("keyPermissions")]
        private InputList<string>? _keyPermissions;

        /// <summary>
        /// List of key permissions, must be one or more from
        /// the following: `backup`, `create`, `decrypt`, `delete`, `encrypt`, `get`, `import`, `list`, `purge`,
        /// `recover`, `restore`, `sign`, `unwrapKey`, `update`, `verify` and `wrapKey`.
        /// </summary>
        public InputList<string> KeyPermissions
        {
            get => _keyPermissions ?? (_keyPermissions = new InputList<string>());
            set => _keyPermissions = value;
        }

        /// <summary>
        /// Specifies the id of the Key Vault resource. Changing this
        /// forces a new resource to be created.
        /// </summary>
        [Input("keyVaultId", required: true)]
        public Input<string> KeyVaultId { get; set; } = null!;

        /// <summary>
        /// The object ID of a user, service principal or security
        /// group in the Azure Active Directory tenant for the vault. The object ID must
        /// be unique for the list of access policies. Changing this forces a new resource
        /// to be created.
        /// </summary>
        [Input("objectId", required: true)]
        public Input<string> ObjectId { get; set; } = null!;

        [Input("secretPermissions")]
        private InputList<string>? _secretPermissions;

        /// <summary>
        /// List of secret permissions, must be one or more
        /// from the following: `backup`, `delete`, `get`, `list`, `purge`, `recover`, `restore` and `set`.
        /// </summary>
        public InputList<string> SecretPermissions
        {
            get => _secretPermissions ?? (_secretPermissions = new InputList<string>());
            set => _secretPermissions = value;
        }

        [Input("storagePermissions")]
        private InputList<string>? _storagePermissions;

        /// <summary>
        /// List of storage permissions, must be one or more from the following: `backup`, `delete`, `deletesas`, `get`, `getsas`, `list`, `listsas`, `purge`, `recover`, `regeneratekey`, `restore`, `set`, `setsas` and `update`.
        /// </summary>
        public InputList<string> StoragePermissions
        {
            get => _storagePermissions ?? (_storagePermissions = new InputList<string>());
            set => _storagePermissions = value;
        }

        /// <summary>
        /// The Azure Active Directory tenant ID that should be used
        /// for authenticating requests to the key vault. Changing this forces a new resource
        /// to be created.
        /// </summary>
        [Input("tenantId", required: true)]
        public Input<string> TenantId { get; set; } = null!;

        public AccessPolicyArgs()
        {
        }
    }

    public sealed class AccessPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The object ID of an Application in Azure Active Directory.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        [Input("certificatePermissions")]
        private InputList<string>? _certificatePermissions;

        /// <summary>
        /// List of certificate permissions, must be one or more from
        /// the following: `backup`, `create`, `delete`, `deleteissuers`, `get`, `getissuers`, `import`, `list`, `listissuers`,
        /// `managecontacts`, `manageissuers`, `purge`, `recover`, `restore`, `setissuers` and `update`.
        /// </summary>
        public InputList<string> CertificatePermissions
        {
            get => _certificatePermissions ?? (_certificatePermissions = new InputList<string>());
            set => _certificatePermissions = value;
        }

        [Input("keyPermissions")]
        private InputList<string>? _keyPermissions;

        /// <summary>
        /// List of key permissions, must be one or more from
        /// the following: `backup`, `create`, `decrypt`, `delete`, `encrypt`, `get`, `import`, `list`, `purge`,
        /// `recover`, `restore`, `sign`, `unwrapKey`, `update`, `verify` and `wrapKey`.
        /// </summary>
        public InputList<string> KeyPermissions
        {
            get => _keyPermissions ?? (_keyPermissions = new InputList<string>());
            set => _keyPermissions = value;
        }

        /// <summary>
        /// Specifies the id of the Key Vault resource. Changing this
        /// forces a new resource to be created.
        /// </summary>
        [Input("keyVaultId")]
        public Input<string>? KeyVaultId { get; set; }

        /// <summary>
        /// The object ID of a user, service principal or security
        /// group in the Azure Active Directory tenant for the vault. The object ID must
        /// be unique for the list of access policies. Changing this forces a new resource
        /// to be created.
        /// </summary>
        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        [Input("secretPermissions")]
        private InputList<string>? _secretPermissions;

        /// <summary>
        /// List of secret permissions, must be one or more
        /// from the following: `backup`, `delete`, `get`, `list`, `purge`, `recover`, `restore` and `set`.
        /// </summary>
        public InputList<string> SecretPermissions
        {
            get => _secretPermissions ?? (_secretPermissions = new InputList<string>());
            set => _secretPermissions = value;
        }

        [Input("storagePermissions")]
        private InputList<string>? _storagePermissions;

        /// <summary>
        /// List of storage permissions, must be one or more from the following: `backup`, `delete`, `deletesas`, `get`, `getsas`, `list`, `listsas`, `purge`, `recover`, `regeneratekey`, `restore`, `set`, `setsas` and `update`.
        /// </summary>
        public InputList<string> StoragePermissions
        {
            get => _storagePermissions ?? (_storagePermissions = new InputList<string>());
            set => _storagePermissions = value;
        }

        /// <summary>
        /// The Azure Active Directory tenant ID that should be used
        /// for authenticating requests to the key vault. Changing this forces a new resource
        /// to be created.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public AccessPolicyState()
        {
        }
    }
}

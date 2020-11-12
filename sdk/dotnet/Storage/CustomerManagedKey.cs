// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Storage
{
    /// <summary>
    /// Manages a Customer Managed Key for a Storage Account.
    /// 
    /// ## Import
    /// 
    /// Customer Managed Keys for a Storage Account can be imported using the `resource id` of the Storage Account, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:storage/customerManagedKey:CustomerManagedKey example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Storage/storageAccounts/myaccount
    /// ```
    /// </summary>
    public partial class CustomerManagedKey : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of Key Vault Key.
        /// </summary>
        [Output("keyName")]
        public Output<string> KeyName { get; private set; } = null!;

        /// <summary>
        /// The ID of the Key Vault. Changing this forces a new resource to be created.
        /// </summary>
        [Output("keyVaultId")]
        public Output<string> KeyVaultId { get; private set; } = null!;

        /// <summary>
        /// The version of Key Vault Key. Remove or omit this argument to enable Automatic Key Rotation.
        /// </summary>
        [Output("keyVersion")]
        public Output<string?> KeyVersion { get; private set; } = null!;

        /// <summary>
        /// The ID of the Storage Account. Changing this forces a new resource to be created.
        /// </summary>
        [Output("storageAccountId")]
        public Output<string> StorageAccountId { get; private set; } = null!;


        /// <summary>
        /// Create a CustomerManagedKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomerManagedKey(string name, CustomerManagedKeyArgs args, CustomResourceOptions? options = null)
            : base("azure:storage/customerManagedKey:CustomerManagedKey", name, args ?? new CustomerManagedKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomerManagedKey(string name, Input<string> id, CustomerManagedKeyState? state = null, CustomResourceOptions? options = null)
            : base("azure:storage/customerManagedKey:CustomerManagedKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CustomerManagedKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomerManagedKey Get(string name, Input<string> id, CustomerManagedKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new CustomerManagedKey(name, id, state, options);
        }
    }

    public sealed class CustomerManagedKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of Key Vault Key.
        /// </summary>
        [Input("keyName", required: true)]
        public Input<string> KeyName { get; set; } = null!;

        /// <summary>
        /// The ID of the Key Vault. Changing this forces a new resource to be created.
        /// </summary>
        [Input("keyVaultId", required: true)]
        public Input<string> KeyVaultId { get; set; } = null!;

        /// <summary>
        /// The version of Key Vault Key. Remove or omit this argument to enable Automatic Key Rotation.
        /// </summary>
        [Input("keyVersion")]
        public Input<string>? KeyVersion { get; set; }

        /// <summary>
        /// The ID of the Storage Account. Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageAccountId", required: true)]
        public Input<string> StorageAccountId { get; set; } = null!;

        public CustomerManagedKeyArgs()
        {
        }
    }

    public sealed class CustomerManagedKeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of Key Vault Key.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// The ID of the Key Vault. Changing this forces a new resource to be created.
        /// </summary>
        [Input("keyVaultId")]
        public Input<string>? KeyVaultId { get; set; }

        /// <summary>
        /// The version of Key Vault Key. Remove or omit this argument to enable Automatic Key Rotation.
        /// </summary>
        [Input("keyVersion")]
        public Input<string>? KeyVersion { get; set; }

        /// <summary>
        /// The ID of the Storage Account. Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageAccountId")]
        public Input<string>? StorageAccountId { get; set; }

        public CustomerManagedKeyState()
        {
        }
    }
}

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
    /// Manages a Key Vault Certificate Issuer.
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
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West US",
    ///         });
    ///         var current = Output.Create(Azure.Core.GetClientConfig.InvokeAsync());
    ///         var exampleKeyVault = new Azure.KeyVault.KeyVault("exampleKeyVault", new Azure.KeyVault.KeyVaultArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             SkuName = "standard",
    ///             TenantId = current.Apply(current =&gt; current.TenantId),
    ///         });
    ///         var exampleCertificateIssuer = new Azure.KeyVault.CertificateIssuer("exampleCertificateIssuer", new Azure.KeyVault.CertificateIssuerArgs
    ///         {
    ///             OrgId = "ExampleOrgName",
    ///             KeyVaultId = exampleKeyVault.Id,
    ///             ProviderName = "DigiCert",
    ///             AccountId = "0000",
    ///             Password = "example-password",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Key Vault Certificate Issuers can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:keyvault/certificateIssuer:CertificateIssuer example "https://key-vault-name.vault.azure.net/certificates/issuers/example"
    /// ```
    /// </summary>
    [AzureResourceType("azure:keyvault/certificateIssuer:CertificateIssuer")]
    public partial class CertificateIssuer : Pulumi.CustomResource
    {
        /// <summary>
        /// The account number with the third-party Certificate Issuer.
        /// </summary>
        [Output("accountId")]
        public Output<string?> AccountId { get; private set; } = null!;

        /// <summary>
        /// One or more `admin` blocks as defined below.
        /// </summary>
        [Output("admins")]
        public Output<ImmutableArray<Outputs.CertificateIssuerAdmin>> Admins { get; private set; } = null!;

        /// <summary>
        /// The ID of the Key Vault in which to create the Certificate Issuer.
        /// </summary>
        [Output("keyVaultId")]
        public Output<string> KeyVaultId { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this Key Vault Certificate Issuer. Changing this forces a new Key Vault Certificate Issuer to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the organization as provided to the issuer.
        /// </summary>
        [Output("orgId")]
        public Output<string?> OrgId { get; private set; } = null!;

        /// <summary>
        /// The password associated with the account and organization ID at the third-party Certificate Issuer. If not specified, will not overwrite any previous value.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// The name of the third-party Certificate Issuer. Possible values are: `DigiCert`, `GlobalSign`, `OneCertV2-PrivateCA`, `OneCertV2-PublicCA` and `SslAdminV2`.
        /// </summary>
        [Output("providerName")]
        public Output<string> ProviderName { get; private set; } = null!;


        /// <summary>
        /// Create a CertificateIssuer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CertificateIssuer(string name, CertificateIssuerArgs args, CustomResourceOptions? options = null)
            : base("azure:keyvault/certificateIssuer:CertificateIssuer", name, args ?? new CertificateIssuerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CertificateIssuer(string name, Input<string> id, CertificateIssuerState? state = null, CustomResourceOptions? options = null)
            : base("azure:keyvault/certificateIssuer:CertificateIssuer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CertificateIssuer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CertificateIssuer Get(string name, Input<string> id, CertificateIssuerState? state = null, CustomResourceOptions? options = null)
        {
            return new CertificateIssuer(name, id, state, options);
        }
    }

    public sealed class CertificateIssuerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account number with the third-party Certificate Issuer.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        [Input("admins")]
        private InputList<Inputs.CertificateIssuerAdminArgs>? _admins;

        /// <summary>
        /// One or more `admin` blocks as defined below.
        /// </summary>
        public InputList<Inputs.CertificateIssuerAdminArgs> Admins
        {
            get => _admins ?? (_admins = new InputList<Inputs.CertificateIssuerAdminArgs>());
            set => _admins = value;
        }

        /// <summary>
        /// The ID of the Key Vault in which to create the Certificate Issuer.
        /// </summary>
        [Input("keyVaultId", required: true)]
        public Input<string> KeyVaultId { get; set; } = null!;

        /// <summary>
        /// The name which should be used for this Key Vault Certificate Issuer. Changing this forces a new Key Vault Certificate Issuer to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the organization as provided to the issuer.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// The password associated with the account and organization ID at the third-party Certificate Issuer. If not specified, will not overwrite any previous value.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The name of the third-party Certificate Issuer. Possible values are: `DigiCert`, `GlobalSign`, `OneCertV2-PrivateCA`, `OneCertV2-PublicCA` and `SslAdminV2`.
        /// </summary>
        [Input("providerName", required: true)]
        public Input<string> ProviderName { get; set; } = null!;

        public CertificateIssuerArgs()
        {
        }
    }

    public sealed class CertificateIssuerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account number with the third-party Certificate Issuer.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        [Input("admins")]
        private InputList<Inputs.CertificateIssuerAdminGetArgs>? _admins;

        /// <summary>
        /// One or more `admin` blocks as defined below.
        /// </summary>
        public InputList<Inputs.CertificateIssuerAdminGetArgs> Admins
        {
            get => _admins ?? (_admins = new InputList<Inputs.CertificateIssuerAdminGetArgs>());
            set => _admins = value;
        }

        /// <summary>
        /// The ID of the Key Vault in which to create the Certificate Issuer.
        /// </summary>
        [Input("keyVaultId")]
        public Input<string>? KeyVaultId { get; set; }

        /// <summary>
        /// The name which should be used for this Key Vault Certificate Issuer. Changing this forces a new Key Vault Certificate Issuer to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the organization as provided to the issuer.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// The password associated with the account and organization ID at the third-party Certificate Issuer. If not specified, will not overwrite any previous value.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The name of the third-party Certificate Issuer. Possible values are: `DigiCert`, `GlobalSign`, `OneCertV2-PrivateCA`, `OneCertV2-PublicCA` and `SslAdminV2`.
        /// </summary>
        [Input("providerName")]
        public Input<string>? ProviderName { get; set; }

        public CertificateIssuerState()
        {
        }
    }
}

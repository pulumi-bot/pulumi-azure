// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService
{
    /// <summary>
    /// Manages an App Service Certificate Order.
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
    ///             Location = "West Europe",
    ///         });
    ///         var exampleCertificateOrder = new Azure.AppService.CertificateOrder("exampleCertificateOrder", new Azure.AppService.CertificateOrderArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = "global",
    ///             DistinguishedName = "CN=example.com",
    ///             ProductType = "Standard",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class CertificateOrder : Pulumi.CustomResource
    {
        /// <summary>
        /// Reasons why App Service Certificate is not renewable at the current moment.
        /// </summary>
        [Output("appServiceCertificateNotRenewableReasons")]
        public Output<ImmutableArray<string>> AppServiceCertificateNotRenewableReasons { get; private set; } = null!;

        /// <summary>
        /// true if the certificate should be automatically renewed when it expires; otherwise, false. Defaults to true.
        /// </summary>
        [Output("autoRenew")]
        public Output<bool?> AutoRenew { get; private set; } = null!;

        /// <summary>
        /// State of the Key Vault secret. A `certificates` block as defined below.
        /// </summary>
        [Output("certificates")]
        public Output<ImmutableArray<Outputs.CertificateOrderCertificate>> Certificates { get; private set; } = null!;

        /// <summary>
        /// Last CSR that was created for this order.
        /// </summary>
        [Output("csr")]
        public Output<string> Csr { get; private set; } = null!;

        /// <summary>
        /// The Distinguished Name for the App Service Certificate Order.
        /// </summary>
        [Output("distinguishedName")]
        public Output<string> DistinguishedName { get; private set; } = null!;

        /// <summary>
        /// Domain verification token.
        /// </summary>
        [Output("domainVerificationToken")]
        public Output<string> DomainVerificationToken { get; private set; } = null!;

        /// <summary>
        /// Certificate expiration time.
        /// </summary>
        [Output("expirationTime")]
        public Output<string> ExpirationTime { get; private set; } = null!;

        /// <summary>
        /// Certificate thumbprint intermediate certificate.
        /// </summary>
        [Output("intermediateThumbprint")]
        public Output<string> IntermediateThumbprint { get; private set; } = null!;

        /// <summary>
        /// Whether the private key is external or not.
        /// </summary>
        [Output("isPrivateKeyExternal")]
        public Output<bool> IsPrivateKeyExternal { get; private set; } = null!;

        /// <summary>
        /// Certificate key size.  Defaults to 2048.
        /// </summary>
        [Output("keySize")]
        public Output<int?> KeySize { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. Currently the only valid value is `global`.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Certificate product type, such as `Standard` or `WildCard`.
        /// </summary>
        [Output("productType")]
        public Output<string?> ProductType { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Certificate thumbprint for root certificate.
        /// </summary>
        [Output("rootThumbprint")]
        public Output<string> RootThumbprint { get; private set; } = null!;

        /// <summary>
        /// Certificate thumbprint for signed certificate.
        /// </summary>
        [Output("signedCertificateThumbprint")]
        public Output<string> SignedCertificateThumbprint { get; private set; } = null!;

        /// <summary>
        /// Current order status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Duration in years (must be between `1` and `3`).  Defaults to `1`.
        /// </summary>
        [Output("validityInYears")]
        public Output<int?> ValidityInYears { get; private set; } = null!;


        /// <summary>
        /// Create a CertificateOrder resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CertificateOrder(string name, CertificateOrderArgs args, CustomResourceOptions? options = null)
            : base("azure:appservice/certificateOrder:CertificateOrder", name, args ?? new CertificateOrderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CertificateOrder(string name, Input<string> id, CertificateOrderState? state = null, CustomResourceOptions? options = null)
            : base("azure:appservice/certificateOrder:CertificateOrder", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CertificateOrder resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CertificateOrder Get(string name, Input<string> id, CertificateOrderState? state = null, CustomResourceOptions? options = null)
        {
            return new CertificateOrder(name, id, state, options);
        }
    }

    public sealed class CertificateOrderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// true if the certificate should be automatically renewed when it expires; otherwise, false. Defaults to true.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// Last CSR that was created for this order.
        /// </summary>
        [Input("csr")]
        public Input<string>? Csr { get; set; }

        /// <summary>
        /// The Distinguished Name for the App Service Certificate Order.
        /// </summary>
        [Input("distinguishedName")]
        public Input<string>? DistinguishedName { get; set; }

        /// <summary>
        /// Certificate key size.  Defaults to 2048.
        /// </summary>
        [Input("keySize")]
        public Input<int>? KeySize { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. Currently the only valid value is `global`.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Certificate product type, such as `Standard` or `WildCard`.
        /// </summary>
        [Input("productType")]
        public Input<string>? ProductType { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Duration in years (must be between `1` and `3`).  Defaults to `1`.
        /// </summary>
        [Input("validityInYears")]
        public Input<int>? ValidityInYears { get; set; }

        public CertificateOrderArgs()
        {
        }
    }

    public sealed class CertificateOrderState : Pulumi.ResourceArgs
    {
        [Input("appServiceCertificateNotRenewableReasons")]
        private InputList<string>? _appServiceCertificateNotRenewableReasons;

        /// <summary>
        /// Reasons why App Service Certificate is not renewable at the current moment.
        /// </summary>
        public InputList<string> AppServiceCertificateNotRenewableReasons
        {
            get => _appServiceCertificateNotRenewableReasons ?? (_appServiceCertificateNotRenewableReasons = new InputList<string>());
            set => _appServiceCertificateNotRenewableReasons = value;
        }

        /// <summary>
        /// true if the certificate should be automatically renewed when it expires; otherwise, false. Defaults to true.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        [Input("certificates")]
        private InputList<Inputs.CertificateOrderCertificateGetArgs>? _certificates;

        /// <summary>
        /// State of the Key Vault secret. A `certificates` block as defined below.
        /// </summary>
        public InputList<Inputs.CertificateOrderCertificateGetArgs> Certificates
        {
            get => _certificates ?? (_certificates = new InputList<Inputs.CertificateOrderCertificateGetArgs>());
            set => _certificates = value;
        }

        /// <summary>
        /// Last CSR that was created for this order.
        /// </summary>
        [Input("csr")]
        public Input<string>? Csr { get; set; }

        /// <summary>
        /// The Distinguished Name for the App Service Certificate Order.
        /// </summary>
        [Input("distinguishedName")]
        public Input<string>? DistinguishedName { get; set; }

        /// <summary>
        /// Domain verification token.
        /// </summary>
        [Input("domainVerificationToken")]
        public Input<string>? DomainVerificationToken { get; set; }

        /// <summary>
        /// Certificate expiration time.
        /// </summary>
        [Input("expirationTime")]
        public Input<string>? ExpirationTime { get; set; }

        /// <summary>
        /// Certificate thumbprint intermediate certificate.
        /// </summary>
        [Input("intermediateThumbprint")]
        public Input<string>? IntermediateThumbprint { get; set; }

        /// <summary>
        /// Whether the private key is external or not.
        /// </summary>
        [Input("isPrivateKeyExternal")]
        public Input<bool>? IsPrivateKeyExternal { get; set; }

        /// <summary>
        /// Certificate key size.  Defaults to 2048.
        /// </summary>
        [Input("keySize")]
        public Input<int>? KeySize { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. Currently the only valid value is `global`.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Certificate product type, such as `Standard` or `WildCard`.
        /// </summary>
        [Input("productType")]
        public Input<string>? ProductType { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Certificate thumbprint for root certificate.
        /// </summary>
        [Input("rootThumbprint")]
        public Input<string>? RootThumbprint { get; set; }

        /// <summary>
        /// Certificate thumbprint for signed certificate.
        /// </summary>
        [Input("signedCertificateThumbprint")]
        public Input<string>? SignedCertificateThumbprint { get; set; }

        /// <summary>
        /// Current order status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Duration in years (must be between `1` and `3`).  Defaults to `1`.
        /// </summary>
        [Input("validityInYears")]
        public Input<int>? ValidityInYears { get; set; }

        public CertificateOrderState()
        {
        }
    }
}

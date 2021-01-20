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
    /// Manages an App Service certificate.
    /// 
    /// ## Import
    /// 
    /// App Service Certificates can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:appservice/certificate:Certificate example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/certificates/certificate1
    /// ```
    /// </summary>
    [AzureResourceType("azure:appservice/certificate:Certificate")]
    public partial class Certificate : Pulumi.CustomResource
    {
        /// <summary>
        /// The expiration date for the certificate.
        /// </summary>
        [Output("expirationDate")]
        public Output<string> ExpirationDate { get; private set; } = null!;

        /// <summary>
        /// The friendly name of the certificate.
        /// </summary>
        [Output("friendlyName")]
        public Output<string> FriendlyName { get; private set; } = null!;

        /// <summary>
        /// List of host names the certificate applies to.
        /// </summary>
        [Output("hostNames")]
        public Output<ImmutableArray<string>> HostNames { get; private set; } = null!;

        /// <summary>
        /// Must be specified when the certificate is for an App Service Environment hosted App Service. Changing this forces a new resource to be created.
        /// </summary>
        [Output("hostingEnvironmentProfileId")]
        public Output<string?> HostingEnvironmentProfileId { get; private set; } = null!;

        /// <summary>
        /// The issue date for the certificate.
        /// </summary>
        [Output("issueDate")]
        public Output<string> IssueDate { get; private set; } = null!;

        /// <summary>
        /// The name of the certificate issuer.
        /// </summary>
        [Output("issuer")]
        public Output<string> Issuer { get; private set; } = null!;

        /// <summary>
        /// The ID of the Key Vault secret. Changing this forces a new resource to be created.
        /// </summary>
        [Output("keyVaultSecretId")]
        public Output<string?> KeyVaultSecretId { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The password to access the certificate's private key. Changing this forces a new resource to be created.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// The base64-encoded contents of the certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Output("pfxBlob")]
        public Output<string?> PfxBlob { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The subject name of the certificate.
        /// </summary>
        [Output("subjectName")]
        public Output<string> SubjectName { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The thumbprint for the certificate.
        /// </summary>
        [Output("thumbprint")]
        public Output<string> Thumbprint { get; private set; } = null!;


        /// <summary>
        /// Create a Certificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Certificate(string name, CertificateArgs args, CustomResourceOptions? options = null)
            : base("azure:appservice/certificate:Certificate", name, args ?? new CertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Certificate(string name, Input<string> id, CertificateState? state = null, CustomResourceOptions? options = null)
            : base("azure:appservice/certificate:Certificate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Certificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Certificate Get(string name, Input<string> id, CertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new Certificate(name, id, state, options);
        }
    }

    public sealed class CertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Must be specified when the certificate is for an App Service Environment hosted App Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("hostingEnvironmentProfileId")]
        public Input<string>? HostingEnvironmentProfileId { get; set; }

        /// <summary>
        /// The ID of the Key Vault secret. Changing this forces a new resource to be created.
        /// </summary>
        [Input("keyVaultSecretId")]
        public Input<string>? KeyVaultSecretId { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The password to access the certificate's private key. Changing this forces a new resource to be created.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The base64-encoded contents of the certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("pfxBlob")]
        public Input<string>? PfxBlob { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public CertificateArgs()
        {
        }
    }

    public sealed class CertificateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The expiration date for the certificate.
        /// </summary>
        [Input("expirationDate")]
        public Input<string>? ExpirationDate { get; set; }

        /// <summary>
        /// The friendly name of the certificate.
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        [Input("hostNames")]
        private InputList<string>? _hostNames;

        /// <summary>
        /// List of host names the certificate applies to.
        /// </summary>
        public InputList<string> HostNames
        {
            get => _hostNames ?? (_hostNames = new InputList<string>());
            set => _hostNames = value;
        }

        /// <summary>
        /// Must be specified when the certificate is for an App Service Environment hosted App Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("hostingEnvironmentProfileId")]
        public Input<string>? HostingEnvironmentProfileId { get; set; }

        /// <summary>
        /// The issue date for the certificate.
        /// </summary>
        [Input("issueDate")]
        public Input<string>? IssueDate { get; set; }

        /// <summary>
        /// The name of the certificate issuer.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// The ID of the Key Vault secret. Changing this forces a new resource to be created.
        /// </summary>
        [Input("keyVaultSecretId")]
        public Input<string>? KeyVaultSecretId { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The password to access the certificate's private key. Changing this forces a new resource to be created.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The base64-encoded contents of the certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("pfxBlob")]
        public Input<string>? PfxBlob { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The subject name of the certificate.
        /// </summary>
        [Input("subjectName")]
        public Input<string>? SubjectName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The thumbprint for the certificate.
        /// </summary>
        [Input("thumbprint")]
        public Input<string>? Thumbprint { get; set; }

        public CertificateState()
        {
        }
    }
}

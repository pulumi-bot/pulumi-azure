// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppPlatform
{
    /// <summary>
    /// Manages an Azure Spring Cloud Certificate.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "Southeast Asia",
    ///         });
    ///         var current = Output.Create(Azure.Core.GetClientConfig.InvokeAsync());
    ///         var exampleServicePrincipal = Output.Create(AzureAD.GetServicePrincipal.InvokeAsync(new AzureAD.GetServicePrincipalArgs
    ///         {
    ///             DisplayName = "Azure Spring Cloud Domain-Management",
    ///         }));
    ///         var exampleKeyVault = new Azure.KeyVault.KeyVault("exampleKeyVault", new Azure.KeyVault.KeyVaultArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             TenantId = current.Apply(current =&gt; current.TenantId),
    ///             SkuName = "standard",
    ///             AccessPolicies = 
    ///             {
    ///                 new Azure.KeyVault.Inputs.KeyVaultAccessPolicyArgs
    ///                 {
    ///                     TenantId = current.Apply(current =&gt; current.TenantId),
    ///                     ObjectId = current.Apply(current =&gt; current.ObjectId),
    ///                     SecretPermissions = 
    ///                     {
    ///                         "set",
    ///                     },
    ///                     CertificatePermissions = 
    ///                     {
    ///                         "create",
    ///                         "delete",
    ///                         "get",
    ///                         "update",
    ///                     },
    ///                 },
    ///                 new Azure.KeyVault.Inputs.KeyVaultAccessPolicyArgs
    ///                 {
    ///                     TenantId = current.Apply(current =&gt; current.TenantId),
    ///                     ObjectId = exampleServicePrincipal.Apply(exampleServicePrincipal =&gt; exampleServicePrincipal.ObjectId),
    ///                     SecretPermissions = 
    ///                     {
    ///                         "get",
    ///                         "list",
    ///                     },
    ///                     CertificatePermissions = 
    ///                     {
    ///                         "get",
    ///                         "list",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///         var exampleCertificate = new Azure.KeyVault.Certificate("exampleCertificate", new Azure.KeyVault.CertificateArgs
    ///         {
    ///             KeyVaultId = exampleKeyVault.Id,
    ///             CertificatePolicy = new Azure.KeyVault.Inputs.CertificateCertificatePolicyArgs
    ///             {
    ///                 IssuerParameters = new Azure.KeyVault.Inputs.CertificateCertificatePolicyIssuerParametersArgs
    ///                 {
    ///                     Name = "Self",
    ///                 },
    ///                 KeyProperties = new Azure.KeyVault.Inputs.CertificateCertificatePolicyKeyPropertiesArgs
    ///                 {
    ///                     Exportable = true,
    ///                     KeySize = 2048,
    ///                     KeyType = "RSA",
    ///                     ReuseKey = true,
    ///                 },
    ///                 LifetimeActions = 
    ///                 {
    ///                     new Azure.KeyVault.Inputs.CertificateCertificatePolicyLifetimeActionArgs
    ///                     {
    ///                         Action = new Azure.KeyVault.Inputs.CertificateCertificatePolicyLifetimeActionActionArgs
    ///                         {
    ///                             ActionType = "AutoRenew",
    ///                         },
    ///                         Trigger = new Azure.KeyVault.Inputs.CertificateCertificatePolicyLifetimeActionTriggerArgs
    ///                         {
    ///                             DaysBeforeExpiry = 30,
    ///                         },
    ///                     },
    ///                 },
    ///                 SecretProperties = new Azure.KeyVault.Inputs.CertificateCertificatePolicySecretPropertiesArgs
    ///                 {
    ///                     ContentType = "application/x-pkcs12",
    ///                 },
    ///                 X509CertificateProperties = new Azure.KeyVault.Inputs.CertificateCertificatePolicyX509CertificatePropertiesArgs
    ///                 {
    ///                     KeyUsages = 
    ///                     {
    ///                         "cRLSign",
    ///                         "dataEncipherment",
    ///                         "digitalSignature",
    ///                         "keyAgreement",
    ///                         "keyCertSign",
    ///                         "keyEncipherment",
    ///                     },
    ///                     Subject = "CN=contoso.com",
    ///                     ValidityInMonths = 12,
    ///                 },
    ///             },
    ///         });
    ///         var exampleSpringCloudService = new Azure.AppPlatform.SpringCloudService("exampleSpringCloudService", new Azure.AppPlatform.SpringCloudServiceArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///         });
    ///         var exampleSpringCloudCertificate = new Azure.AppPlatform.SpringCloudCertificate("exampleSpringCloudCertificate", new Azure.AppPlatform.SpringCloudCertificateArgs
    ///         {
    ///             ResourceGroupName = exampleSpringCloudService.ResourceGroupName,
    ///             ServiceName = exampleSpringCloudService.Name,
    ///             KeyVaultCertificateId = exampleCertificate.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Spring Cloud Certificate can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:appplatform/springCloudCertificate:SpringCloudCertificate example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/Spring/spring1/certificates/cert1
    /// ```
    /// </summary>
    public partial class SpringCloudCertificate : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the ID of the Key Vault Certificate resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("keyVaultCertificateId")]
        public Output<string> KeyVaultCertificateId { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Spring Cloud Certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group in which to create the Spring Cloud Certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;


        /// <summary>
        /// Create a SpringCloudCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SpringCloudCertificate(string name, SpringCloudCertificateArgs args, CustomResourceOptions? options = null)
            : base("azure:appplatform/springCloudCertificate:SpringCloudCertificate", name, args ?? new SpringCloudCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SpringCloudCertificate(string name, Input<string> id, SpringCloudCertificateState? state = null, CustomResourceOptions? options = null)
            : base("azure:appplatform/springCloudCertificate:SpringCloudCertificate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SpringCloudCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SpringCloudCertificate Get(string name, Input<string> id, SpringCloudCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new SpringCloudCertificate(name, id, state, options);
        }
    }

    public sealed class SpringCloudCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the ID of the Key Vault Certificate resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("keyVaultCertificateId", required: true)]
        public Input<string> KeyVaultCertificateId { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Spring Cloud Certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the resource group in which to create the Spring Cloud Certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public SpringCloudCertificateArgs()
        {
        }
    }

    public sealed class SpringCloudCertificateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the ID of the Key Vault Certificate resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("keyVaultCertificateId")]
        public Input<string>? KeyVaultCertificateId { get; set; }

        /// <summary>
        /// Specifies the name of the Spring Cloud Certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the resource group in which to create the Spring Cloud Certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public SpringCloudCertificateState()
        {
        }
    }
}

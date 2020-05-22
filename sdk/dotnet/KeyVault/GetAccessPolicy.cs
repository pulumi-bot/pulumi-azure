// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.KeyVault
{
    public static class GetAccessPolicy
    {
        /// <summary>
        /// Use this data source to access information about the permissions from the Management Key Vault Templates.
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
        ///         var contributor = Output.Create(Azure.KeyVault.GetAccessPolicy.InvokeAsync(new Azure.KeyVault.GetAccessPolicyArgs
        ///         {
        ///             Name = "Key Management",
        ///         }));
        ///         this.AccessPolicyKeyPermissions = contributor.Apply(contributor =&gt; contributor.KeyPermissions);
        ///     }
        /// 
        ///     [Output("accessPolicyKeyPermissions")]
        ///     public Output&lt;string&gt; AccessPolicyKeyPermissions { get; set; }
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAccessPolicyResult> InvokeAsync(GetAccessPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccessPolicyResult>("azure:keyvault/getAccessPolicy:getAccessPolicy", args ?? new GetAccessPolicyArgs(), options.WithVersion());
    }


    public sealed class GetAccessPolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Management Template. Possible values are: `Key Management`,
        /// `Secret Management`, `Certificate Management`, `Key &amp; Secret Management`, `Key &amp; Certificate Management`,
        /// `Secret &amp; Certificate Management`,  `Key, Secret, &amp; Certificate Management`
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetAccessPolicyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAccessPolicyResult
    {
        /// <summary>
        /// the certificate permissions for the access policy
        /// </summary>
        public readonly ImmutableArray<string> CertificatePermissions;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// the key permissions for the access policy
        /// </summary>
        public readonly ImmutableArray<string> KeyPermissions;
        public readonly string Name;
        /// <summary>
        /// the secret permissions for the access policy
        /// </summary>
        public readonly ImmutableArray<string> SecretPermissions;

        [OutputConstructor]
        private GetAccessPolicyResult(
            ImmutableArray<string> certificatePermissions,

            string id,

            ImmutableArray<string> keyPermissions,

            string name,

            ImmutableArray<string> secretPermissions)
        {
            CertificatePermissions = certificatePermissions;
            Id = id;
            KeyPermissions = keyPermissions;
            Name = name;
            SecretPermissions = secretPermissions;
        }
    }
}

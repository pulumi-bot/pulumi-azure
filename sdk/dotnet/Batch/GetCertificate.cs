// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Batch
{
    public static class GetCertificate
    {
        /// <summary>
        /// Use this data source to access information about an existing certificate in a Batch Account.
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
        ///         var example = Output.Create(Azure.Batch.GetCertificate.InvokeAsync(new Azure.Batch.GetCertificateArgs
        ///         {
        ///             Name = "SHA1-42C107874FD0E4A9583292A2F1098E8FE4B2EDDA",
        ///             AccountName = "examplebatchaccount",
        ///             ResourceGroupName = "example",
        ///         }));
        ///         this.Thumbprint = example.Apply(example =&gt; example.Thumbprint);
        ///     }
        /// 
        ///     [Output("thumbprint")]
        ///     public Output&lt;string&gt; Thumbprint { get; set; }
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCertificateResult> InvokeAsync(GetCertificateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCertificateResult>("azure:batch/getCertificate:getCertificate", args ?? new GetCertificateArgs(), options.WithVersion());
    }


    public sealed class GetCertificateArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Batch account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the Batch certificate.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where this Batch account exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetCertificateArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCertificateResult
    {
        public readonly string AccountName;
        /// <summary>
        /// The format of the certificate, such as `Cer` or `Pfx`.
        /// </summary>
        public readonly string Format;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// The public key of the certificate.
        /// </summary>
        public readonly string PublicData;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The thumbprint of the certificate.
        /// </summary>
        public readonly string Thumbprint;
        /// <summary>
        /// The algorithm of the certificate thumbprint.
        /// </summary>
        public readonly string ThumbprintAlgorithm;

        [OutputConstructor]
        private GetCertificateResult(
            string accountName,

            string format,

            string id,

            string name,

            string publicData,

            string resourceGroupName,

            string thumbprint,

            string thumbprintAlgorithm)
        {
            AccountName = accountName;
            Format = format;
            Id = id;
            Name = name;
            PublicData = publicData;
            ResourceGroupName = resourceGroupName;
            Thumbprint = thumbprint;
            ThumbprintAlgorithm = thumbprintAlgorithm;
        }
    }
}

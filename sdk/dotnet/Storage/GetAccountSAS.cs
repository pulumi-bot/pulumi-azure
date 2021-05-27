// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Storage
{
    public static class GetAccountSAS
    {
        /// <summary>
        /// Use this data source to obtain a Shared Access Signature (SAS Token) for an existing Storage Account.
        /// 
        /// Shared access signatures allow fine-grained, ephemeral access control to various aspects of an Azure Storage Account.
        /// 
        /// Note that this is an [Account SAS](https://docs.microsoft.com/en-us/rest/api/storageservices/constructing-an-account-sas)
        /// and *not* a [Service SAS](https://docs.microsoft.com/en-us/rest/api/storageservices/constructing-a-service-sas).
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
        ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
        ///         {
        ///             Location = "West Europe",
        ///         });
        ///         var exampleAccount = new Azure.Storage.Account("exampleAccount", new Azure.Storage.AccountArgs
        ///         {
        ///             ResourceGroupName = exampleResourceGroup.Name,
        ///             Location = "westus",
        ///             AccountTier = "Standard",
        ///             AccountReplicationType = "GRS",
        ///             Tags = 
        ///             {
        ///                 { "environment", "staging" },
        ///             },
        ///         });
        ///         var exampleAccountSAS = exampleAccount.PrimaryConnectionString.Apply(primaryConnectionString =&gt; Azure.Storage.GetAccountSAS.InvokeAsync(new Azure.Storage.GetAccountSASArgs
        ///         {
        ///             ConnectionString = primaryConnectionString,
        ///             HttpsOnly = true,
        ///             SignedVersion = "2017-07-29",
        ///             ResourceTypes = new Azure.Storage.Inputs.GetAccountSASResourceTypesArgs
        ///             {
        ///                 Service = true,
        ///                 Container = false,
        ///                 Object = false,
        ///             },
        ///             Services = new Azure.Storage.Inputs.GetAccountSASServicesArgs
        ///             {
        ///                 Blob = true,
        ///                 Queue = false,
        ///                 Table = false,
        ///                 File = false,
        ///             },
        ///             Start = "2018-03-21T00:00:00Z",
        ///             Expiry = "2020-03-21T00:00:00Z",
        ///             Permissions = new Azure.Storage.Inputs.GetAccountSASPermissionsArgs
        ///             {
        ///                 Read = true,
        ///                 Write = true,
        ///                 Delete = false,
        ///                 List = false,
        ///                 Add = true,
        ///                 Create = true,
        ///                 Update = false,
        ///                 Process = false,
        ///             },
        ///         }));
        ///         this.SasUrlQueryString = exampleAccountSAS.Apply(exampleAccountSAS =&gt; exampleAccountSAS.Sas);
        ///     }
        /// 
        ///     [Output("sasUrlQueryString")]
        ///     public Output&lt;string&gt; SasUrlQueryString { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAccountSASResult> InvokeAsync(GetAccountSASArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccountSASResult>("azure:storage/getAccountSAS:getAccountSAS", args ?? new GetAccountSASArgs(), options.WithVersion());

        public static Output<GetAccountSASResult> Invoke(GetAccountSASOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.ConnectionString.Box(),
                args.Expiry.Box(),
                args.HttpsOnly.Box(),
                args.Permissions.Box(),
                args.ResourceTypes.Box(),
                args.Services.Box(),
                args.SignedVersion.Box(),
                args.Start.Box()
            ).Apply(a => {
                    var args = new GetAccountSASArgs();
                    a[0].Set(args, nameof(args.ConnectionString));
                    a[1].Set(args, nameof(args.Expiry));
                    a[2].Set(args, nameof(args.HttpsOnly));
                    a[3].Set(args, nameof(args.Permissions));
                    a[4].Set(args, nameof(args.ResourceTypes));
                    a[5].Set(args, nameof(args.Services));
                    a[6].Set(args, nameof(args.SignedVersion));
                    a[7].Set(args, nameof(args.Start));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetAccountSASArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The connection string for the storage account to which this SAS applies. Typically directly from the `primary_connection_string` attribute of a `azure.storage.Account` resource.
        /// </summary>
        [Input("connectionString", required: true)]
        public string ConnectionString { get; set; } = null!;

        /// <summary>
        /// The expiration time and date of this SAS. Must be a valid ISO-8601 format time/date string.
        /// </summary>
        [Input("expiry", required: true)]
        public string Expiry { get; set; } = null!;

        /// <summary>
        /// Only permit `https` access. If `false`, both `http` and `https` are permitted. Defaults to `true`.
        /// </summary>
        [Input("httpsOnly")]
        public bool? HttpsOnly { get; set; }

        /// <summary>
        /// A `permissions` block as defined below.
        /// </summary>
        [Input("permissions", required: true)]
        public Inputs.GetAccountSASPermissionsArgs Permissions { get; set; } = null!;

        /// <summary>
        /// A `resource_types` block as defined below.
        /// </summary>
        [Input("resourceTypes", required: true)]
        public Inputs.GetAccountSASResourceTypesArgs ResourceTypes { get; set; } = null!;

        /// <summary>
        /// A `services` block as defined below.
        /// </summary>
        [Input("services", required: true)]
        public Inputs.GetAccountSASServicesArgs Services { get; set; } = null!;

        /// <summary>
        /// Specifies the signed storage service version to use to authorize requests made with this account SAS. Defaults to `2017-07-29`.
        /// </summary>
        [Input("signedVersion")]
        public string? SignedVersion { get; set; }

        /// <summary>
        /// The starting time and date of validity of this SAS. Must be a valid ISO-8601 format time/date string.
        /// </summary>
        [Input("start", required: true)]
        public string Start { get; set; } = null!;

        public GetAccountSASArgs()
        {
        }
    }

    public sealed class GetAccountSASOutputArgs
    {
        /// <summary>
        /// The connection string for the storage account to which this SAS applies. Typically directly from the `primary_connection_string` attribute of a `azure.storage.Account` resource.
        /// </summary>
        [Input("connectionString", required: true)]
        public Input<string> ConnectionString { get; set; } = null!;

        /// <summary>
        /// The expiration time and date of this SAS. Must be a valid ISO-8601 format time/date string.
        /// </summary>
        [Input("expiry", required: true)]
        public Input<string> Expiry { get; set; } = null!;

        /// <summary>
        /// Only permit `https` access. If `false`, both `http` and `https` are permitted. Defaults to `true`.
        /// </summary>
        [Input("httpsOnly")]
        public Input<bool>? HttpsOnly { get; set; }

        /// <summary>
        /// A `permissions` block as defined below.
        /// </summary>
        [Input("permissions", required: true)]
        public Input<Inputs.GetAccountSASPermissionsArgs> Permissions { get; set; } = null!;

        /// <summary>
        /// A `resource_types` block as defined below.
        /// </summary>
        [Input("resourceTypes", required: true)]
        public Input<Inputs.GetAccountSASResourceTypesArgs> ResourceTypes { get; set; } = null!;

        /// <summary>
        /// A `services` block as defined below.
        /// </summary>
        [Input("services", required: true)]
        public Input<Inputs.GetAccountSASServicesArgs> Services { get; set; } = null!;

        /// <summary>
        /// Specifies the signed storage service version to use to authorize requests made with this account SAS. Defaults to `2017-07-29`.
        /// </summary>
        [Input("signedVersion")]
        public Input<string>? SignedVersion { get; set; }

        /// <summary>
        /// The starting time and date of validity of this SAS. Must be a valid ISO-8601 format time/date string.
        /// </summary>
        [Input("start", required: true)]
        public Input<string> Start { get; set; } = null!;

        public GetAccountSASOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAccountSASResult
    {
        public readonly string ConnectionString;
        public readonly string Expiry;
        public readonly bool? HttpsOnly;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly Outputs.GetAccountSASPermissionsResult Permissions;
        public readonly Outputs.GetAccountSASResourceTypesResult ResourceTypes;
        /// <summary>
        /// The computed Account Shared Access Signature (SAS).
        /// </summary>
        public readonly string Sas;
        public readonly Outputs.GetAccountSASServicesResult Services;
        public readonly string? SignedVersion;
        public readonly string Start;

        [OutputConstructor]
        private GetAccountSASResult(
            string connectionString,

            string expiry,

            bool? httpsOnly,

            string id,

            Outputs.GetAccountSASPermissionsResult permissions,

            Outputs.GetAccountSASResourceTypesResult resourceTypes,

            string sas,

            Outputs.GetAccountSASServicesResult services,

            string? signedVersion,

            string start)
        {
            ConnectionString = connectionString;
            Expiry = expiry;
            HttpsOnly = httpsOnly;
            Id = id;
            Permissions = permissions;
            ResourceTypes = resourceTypes;
            Sas = sas;
            Services = services;
            SignedVersion = signedVersion;
            Start = start;
        }
    }
}

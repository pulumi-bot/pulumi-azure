// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Sql
{
    public static class GetServer
    {
        /// <summary>
        /// Use this data source to access information about an existing SQL Azure Database Server.
        /// 
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
        ///         var example = Output.Create(Azure.Sql.GetServer.InvokeAsync(new Azure.Sql.GetServerArgs
        ///         {
        ///             Name = "examplesqlservername",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///         this.SqlServerId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("sqlServerId")]
        ///     public Output&lt;string&gt; SqlServerId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServerResult> InvokeAsync(GetServerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServerResult>("azure:sql/getServer:getServer", args ?? new GetServerArgs(), options.WithVersion());
    }


    public sealed class GetServerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the SQL Server.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group where the SQL Server exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetServerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServerResult
    {
        /// <summary>
        /// The administrator username of the SQL Server.
        /// </summary>
        public readonly string AdministratorLogin;
        /// <summary>
        /// The fully qualified domain name of the SQL Server.
        /// </summary>
        public readonly string Fqdn;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServerIdentityResult> Identities;
        /// <summary>
        /// The location of the Resource Group in which the SQL Server exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The version of the SQL Server.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetServerResult(
            string administratorLogin,

            string fqdn,

            string id,

            ImmutableArray<Outputs.GetServerIdentityResult> identities,

            string location,

            string name,

            string resourceGroupName,

            ImmutableDictionary<string, string> tags,

            string version)
        {
            AdministratorLogin = administratorLogin;
            Fqdn = fqdn;
            Id = id;
            Identities = identities;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
            Version = version;
        }
    }
}

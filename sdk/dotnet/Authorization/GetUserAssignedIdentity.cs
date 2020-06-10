// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Authorization
{
    public static class GetUserAssignedIdentity
    {
        /// <summary>
        /// Use this data source to access information about an existing User Assigned Identity.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// 
        /// {{% example %}}
        /// ### Reference An Existing)
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.Authorization.GetUserAssignedIdentity.InvokeAsync(new Azure.Authorization.GetUserAssignedIdentityArgs
        ///         {
        ///             Name = "name_of_user_assigned_identity",
        ///             ResourceGroupName = "name_of_resource_group",
        ///         }));
        ///         this.UaiClientId = example.Apply(example =&gt; example.ClientId);
        ///         this.UaiPrincipalId = example.Apply(example =&gt; example.PrincipalId);
        ///     }
        /// 
        ///     [Output("uaiClientId")]
        ///     public Output&lt;string&gt; UaiClientId { get; set; }
        ///     [Output("uaiPrincipalId")]
        ///     public Output&lt;string&gt; UaiPrincipalId { get; set; }
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUserAssignedIdentityResult> InvokeAsync(GetUserAssignedIdentityArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUserAssignedIdentityResult>("azure:authorization/getUserAssignedIdentity:getUserAssignedIdentity", args ?? new GetUserAssignedIdentityArgs(), options.WithVersion());
    }


    public sealed class GetUserAssignedIdentityArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the User Assigned Identity.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the User Assigned Identity exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetUserAssignedIdentityArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUserAssignedIdentityResult
    {
        /// <summary>
        /// The Client ID of the User Assigned Identity.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure location where the User Assigned Identity exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        /// <summary>
        /// The Service Principal ID of the User Assigned Identity.
        /// </summary>
        public readonly string PrincipalId;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A mapping of tags assigned to the User Assigned Identity.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetUserAssignedIdentityResult(
            string clientId,

            string id,

            string location,

            string name,

            string principalId,

            string resourceGroupName,

            ImmutableDictionary<string, string> tags)
        {
            ClientId = clientId;
            Id = id;
            Location = location;
            Name = name;
            PrincipalId = principalId;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
        }
    }
}

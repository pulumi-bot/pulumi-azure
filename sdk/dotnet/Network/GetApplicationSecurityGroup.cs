// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    public static class GetApplicationSecurityGroup
    {
        /// <summary>
        /// Use this data source to access information about an existing Application Security Group.
        /// </summary>
        public static Task<GetApplicationSecurityGroupResult> InvokeAsync(GetApplicationSecurityGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApplicationSecurityGroupResult>("azure:network/getApplicationSecurityGroup:getApplicationSecurityGroup", args ?? new GetApplicationSecurityGroupArgs(), options.WithVersion());
    }


    public sealed class GetApplicationSecurityGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Application Security Group.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the Application Security Group exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetApplicationSecurityGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApplicationSecurityGroupResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The supported Azure location where the Application Security Group exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetApplicationSecurityGroupResult(
            string id,

            string location,

            string name,

            string resourceGroupName,

            ImmutableDictionary<string, string> tags)
        {
            Id = id;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
        }
    }
}

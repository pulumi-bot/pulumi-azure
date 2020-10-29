// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Lb
{
    public static class GetBackendAddressPool
    {
        /// <summary>
        /// Use this data source to access information about an existing Load Balancer's Backend Address Pool.
        /// </summary>
        public static Task<GetBackendAddressPoolResult> InvokeAsync(GetBackendAddressPoolArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBackendAddressPoolResult>("azure:lb/getBackendAddressPool:getBackendAddressPool", args ?? new GetBackendAddressPoolArgs(), options.WithVersion());
    }


    public sealed class GetBackendAddressPoolArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Load Balancer in which the Backend Address Pool exists.
        /// </summary>
        [Input("loadbalancerId", required: true)]
        public string LoadbalancerId { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Backend Address Pool.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetBackendAddressPoolArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBackendAddressPoolResult
    {
        /// <summary>
        /// An array of references to IP addresses defined in network interfaces.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBackendAddressPoolBackendIpConfigurationResult> BackendIpConfigurations;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string LoadbalancerId;
        /// <summary>
        /// The name of the Backend Address Pool.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetBackendAddressPoolResult(
            ImmutableArray<Outputs.GetBackendAddressPoolBackendIpConfigurationResult> backendIpConfigurations,

            string id,

            string loadbalancerId,

            string name)
        {
            BackendIpConfigurations = backendIpConfigurations;
            Id = id;
            LoadbalancerId = loadbalancerId;
            Name = name;
        }
    }
}

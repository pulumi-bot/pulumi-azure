// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.PrivateLink
{
    public static class GetService
    {
        /// <summary>
        /// Use this data source to access information about an existing Private Link Service.
        /// </summary>
        public static Task<GetServiceResult> InvokeAsync(GetServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceResult>("azure:privatelink/getService:getService", args ?? new GetServiceArgs(), options.WithVersion());
    }


    public sealed class GetServiceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the private link service.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the private link service resides.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetServiceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServiceResult
    {
        /// <summary>
        /// The alias is a globally unique name for your private link service which Azure generates for you. Your can use this alias to request a connection to your private link service.
        /// </summary>
        public readonly string Alias;
        /// <summary>
        /// The list of subscription(s) globally unique identifiers that will be auto approved to use the private link service.
        /// </summary>
        public readonly ImmutableArray<string> AutoApprovalSubscriptionIds;
        /// <summary>
        /// Does the Private Link Service support the Proxy Protocol?
        /// </summary>
        public readonly bool EnableProxyProtocol;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of Standard Load Balancer(SLB) resource IDs. The Private Link service is tied to the frontend IP address of a SLB. All traffic destined for the private link service will reach the frontend of the SLB. You can configure SLB rules to direct this traffic to appropriate backend pools where your applications are running.
        /// </summary>
        public readonly ImmutableArray<string> LoadBalancerFrontendIpConfigurationIds;
        /// <summary>
        /// The supported Azure location where the resource exists.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of private link service NAT IP configuration.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The `nat_ip_configuration` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServiceNatIpConfigurationResult> NatIpConfigurations;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The list of subscription(s) globally unique identifiers(GUID) that will be able to see the private link service.
        /// </summary>
        public readonly ImmutableArray<string> VisibilitySubscriptionIds;

        [OutputConstructor]
        private GetServiceResult(
            string alias,

            ImmutableArray<string> autoApprovalSubscriptionIds,

            bool enableProxyProtocol,

            string id,

            ImmutableArray<string> loadBalancerFrontendIpConfigurationIds,

            string location,

            string name,

            ImmutableArray<Outputs.GetServiceNatIpConfigurationResult> natIpConfigurations,

            string resourceGroupName,

            ImmutableDictionary<string, string> tags,

            ImmutableArray<string> visibilitySubscriptionIds)
        {
            Alias = alias;
            AutoApprovalSubscriptionIds = autoApprovalSubscriptionIds;
            EnableProxyProtocol = enableProxyProtocol;
            Id = id;
            LoadBalancerFrontendIpConfigurationIds = loadBalancerFrontendIpConfigurationIds;
            Location = location;
            Name = name;
            NatIpConfigurations = natIpConfigurations;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
            VisibilitySubscriptionIds = visibilitySubscriptionIds;
        }
    }
}

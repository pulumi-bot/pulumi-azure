// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Outputs
{

    [OutputType]
    public sealed class GetFirewallVirtualHubResult
    {
        /// <summary>
        /// The private IP address associated with the Azure Firewall.
        /// </summary>
        public readonly string PrivateIpAddress;
        /// <summary>
        /// The list of public IP addresses associated with the Azure Firewall.
        /// </summary>
        public readonly ImmutableArray<string> PublicIpAddresses;
        /// <summary>
        /// The number of public IPs assigned to the Azure Firewall.
        /// </summary>
        public readonly int PublicIpCount;
        /// <summary>
        /// The ID of the Virtual Hub where the Azure Firewall resides in.
        /// </summary>
        public readonly string VirtualHubId;

        [OutputConstructor]
        private GetFirewallVirtualHubResult(
            string privateIpAddress,

            ImmutableArray<string> publicIpAddresses,

            int publicIpCount,

            string virtualHubId)
        {
            PrivateIpAddress = privateIpAddress;
            PublicIpAddresses = publicIpAddresses;
            PublicIpCount = publicIpCount;
            VirtualHubId = virtualHubId;
        }
    }
}

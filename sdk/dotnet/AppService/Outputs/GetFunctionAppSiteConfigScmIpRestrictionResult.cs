// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService.Outputs
{

    [OutputType]
    public sealed class GetFunctionAppSiteConfigScmIpRestrictionResult
    {
        /// <summary>
        /// Allow or Deny access for this IP range. Defaults to Allow.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// The IP Address used for this IP Restriction in CIDR notation.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// The name of the Function App resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The priority for this IP Restriction.
        /// </summary>
        public readonly int Priority;
        public readonly string SubnetId;
        /// <summary>
        /// The Virtual Network Subnet ID used for this IP Restriction.
        /// </summary>
        public readonly string VirtualNetworkSubnetId;

        [OutputConstructor]
        private GetFunctionAppSiteConfigScmIpRestrictionResult(
            string action,

            string ipAddress,

            string name,

            int priority,

            string subnetId,

            string virtualNetworkSubnetId)
        {
            Action = action;
            IpAddress = ipAddress;
            Name = name;
            Priority = priority;
            SubnetId = subnetId;
            VirtualNetworkSubnetId = virtualNetworkSubnetId;
        }
    }
}

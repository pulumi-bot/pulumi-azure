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
    public sealed class FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRule
    {
        /// <summary>
        /// Specifies a list of destination IP addresses (including CIDR and `*`) or Service Tags.
        /// </summary>
        public readonly ImmutableArray<string> DestinationAddresses;
        /// <summary>
        /// Specifies a list of destination FQDNs.
        /// </summary>
        public readonly ImmutableArray<string> DestinationFqdns;
        /// <summary>
        /// Specifies a list of destination IP groups.
        /// </summary>
        public readonly ImmutableArray<string> DestinationIpGroups;
        /// <summary>
        /// Specifies a list of destination ports.
        /// </summary>
        public readonly ImmutableArray<string> DestinationPorts;
        /// <summary>
        /// The name which should be used for this rule.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies a list of network protocols this rule applies to. Possible values are `TCP`, `UDP`.
        /// </summary>
        public readonly ImmutableArray<string> Protocols;
        /// <summary>
        /// Specifies a list of source IP addresses (including CIDR and `*`).
        /// </summary>
        public readonly ImmutableArray<string> SourceAddresses;
        /// <summary>
        /// Specifies a list of source IP groups.
        /// </summary>
        public readonly ImmutableArray<string> SourceIpGroups;

        [OutputConstructor]
        private FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRule(
            ImmutableArray<string> destinationAddresses,

            ImmutableArray<string> destinationFqdns,

            ImmutableArray<string> destinationIpGroups,

            ImmutableArray<string> destinationPorts,

            string name,

            ImmutableArray<string> protocols,

            ImmutableArray<string> sourceAddresses,

            ImmutableArray<string> sourceIpGroups)
        {
            DestinationAddresses = destinationAddresses;
            DestinationFqdns = destinationFqdns;
            DestinationIpGroups = destinationIpGroups;
            DestinationPorts = destinationPorts;
            Name = name;
            Protocols = protocols;
            SourceAddresses = sourceAddresses;
            SourceIpGroups = sourceIpGroups;
        }
    }
}
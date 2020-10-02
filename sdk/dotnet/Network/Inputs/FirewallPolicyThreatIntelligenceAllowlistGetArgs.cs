// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class FirewallPolicyThreatIntelligenceAllowlistGetArgs : Pulumi.ResourceArgs
    {
        [Input("fqdns")]
        private InputList<string>? _fqdns;

        /// <summary>
        /// A list of FQDNs that will be skipped for threat detection.
        /// </summary>
        public InputList<string> Fqdns
        {
            get => _fqdns ?? (_fqdns = new InputList<string>());
            set => _fqdns = value;
        }

        [Input("ipAddresses")]
        private InputList<string>? _ipAddresses;

        /// <summary>
        /// A list of IP addresses or IP address ranges that will be skipped for threat detection.
        /// </summary>
        public InputList<string> IpAddresses
        {
            get => _ipAddresses ?? (_ipAddresses = new InputList<string>());
            set => _ipAddresses = value;
        }

        public FirewallPolicyThreatIntelligenceAllowlistGetArgs()
        {
        }
    }
}
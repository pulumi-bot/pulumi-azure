// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class FirewallIpConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the IP Configuration.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The private IP address associated with the Firewall.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        /// <summary>
        /// The ID of the Public IP Address associated with the firewall.
        /// </summary>
        [Input("publicIpAddressId", required: true)]
        public Input<string> PublicIpAddressId { get; set; } = null!;

        /// <summary>
        /// Reference to the subnet associated with the IP Configuration.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public FirewallIpConfigurationGetArgs()
        {
        }
    }
}

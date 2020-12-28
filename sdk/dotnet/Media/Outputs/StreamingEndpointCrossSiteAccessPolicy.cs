// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Media.Outputs
{

    [OutputType]
    public sealed class StreamingEndpointCrossSiteAccessPolicy
    {
        /// <summary>
        /// The content of clientaccesspolicy.xml used by Silverlight.
        /// </summary>
        public readonly string? ClientAccessPolicy;
        /// <summary>
        /// The content of crossdomain.xml used by Silverlight.
        /// </summary>
        public readonly string? CrossDomainPolicy;

        [OutputConstructor]
        private StreamingEndpointCrossSiteAccessPolicy(
            string? clientAccessPolicy,

            string? crossDomainPolicy)
        {
            ClientAccessPolicy = clientAccessPolicy;
            CrossDomainPolicy = crossDomainPolicy;
        }
    }
}
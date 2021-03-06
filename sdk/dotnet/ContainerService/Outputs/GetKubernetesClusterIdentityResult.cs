// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService.Outputs
{

    [OutputType]
    public sealed class GetKubernetesClusterIdentityResult
    {
        /// <summary>
        /// The principal id of the system assigned identity which is used by master components.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The tenant id of the system assigned identity which is used by master components.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The type of identity used for the managed cluster.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetKubernetesClusterIdentityResult(
            string principalId,

            string tenantId,

            string type)
        {
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
        }
    }
}

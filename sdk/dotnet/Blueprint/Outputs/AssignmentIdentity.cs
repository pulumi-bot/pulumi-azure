// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Blueprint.Outputs
{

    [OutputType]
    public sealed class AssignmentIdentity
    {
        public readonly ImmutableArray<string> IdentityIds;
        public readonly string? PrincipalId;
        public readonly string? TenantId;
        /// <summary>
        /// The Identity type for the Managed Service Identity. Currently only `UserAssigned` is supported.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private AssignmentIdentity(
            ImmutableArray<string> identityIds,

            string? principalId,

            string? tenantId,

            string type)
        {
            IdentityIds = identityIds;
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
        }
    }
}
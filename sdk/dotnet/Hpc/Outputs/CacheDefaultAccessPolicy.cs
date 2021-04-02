// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Hpc.Outputs
{

    [OutputType]
    public sealed class CacheDefaultAccessPolicy
    {
        /// <summary>
        /// One to three `access_rule` blocks as defined above.
        /// </summary>
        public readonly ImmutableArray<Outputs.CacheDefaultAccessPolicyAccessRule> AccessRules;

        [OutputConstructor]
        private CacheDefaultAccessPolicy(ImmutableArray<Outputs.CacheDefaultAccessPolicyAccessRule> accessRules)
        {
            AccessRules = accessRules;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Hpc.Inputs
{

    public sealed class CacheDefaultAccessPolicyAccessRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access level for this rule. Possible values are: `rw`, `ro`, `no`.
        /// </summary>
        [Input("access", required: true)]
        public Input<string> Access { get; set; } = null!;

        /// <summary>
        /// The anonymous GID used when `root_squash_enabled` is `true`.
        /// </summary>
        [Input("anonymousGid")]
        public Input<int>? AnonymousGid { get; set; }

        /// <summary>
        /// The anonymous UID used when `root_squash_enabled` is `true`.
        /// </summary>
        [Input("anonymousUid")]
        public Input<int>? AnonymousUid { get; set; }

        /// <summary>
        /// The filter applied to the `scope` for this rule. The filter's format depends on its scope: `default` scope matches all clients and has no filter value; `network` scope takes a CIDR format; `host` takes an IP address or fully qualified domain name. If a client does not match any filter rule and there is no default rule, access is denied.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// Whether to enable [root squash](https://docs.microsoft.com/en-us/azure/hpc-cache/access-policies#root-squash)? Defaults to `false`.
        /// </summary>
        [Input("rootSquashEnabled")]
        public Input<bool>? RootSquashEnabled { get; set; }

        /// <summary>
        /// The scope of this rule. The `scope` and (potentially) the `filter` determine which clients match the rule. Possible values are: `default`, `network`, `host`.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        /// <summary>
        /// Whether allow access to subdirectories under the root export? Defaults to `false`.
        /// </summary>
        [Input("submountAccessEnabled")]
        public Input<bool>? SubmountAccessEnabled { get; set; }

        /// <summary>
        /// Whether [SUID](https://docs.microsoft.com/en-us/azure/hpc-cache/access-policies#suid) is allowed? Defaults to `false`.
        /// </summary>
        [Input("suidEnabled")]
        public Input<bool>? SuidEnabled { get; set; }

        public CacheDefaultAccessPolicyAccessRuleArgs()
        {
        }
    }
}

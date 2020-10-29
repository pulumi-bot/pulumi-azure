// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Core
{
    public static class GetSubscriptions
    {
        /// <summary>
        /// Use this data source to access information about all the Subscriptions currently available.
        /// </summary>
        public static Task<GetSubscriptionsResult> InvokeAsync(GetSubscriptionsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSubscriptionsResult>("azure:core/getSubscriptions:getSubscriptions", args ?? new GetSubscriptionsArgs(), options.WithVersion());
    }


    public sealed class GetSubscriptionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A case-insensitive value which must be contained within the `display_name` field, used to filter the results
        /// </summary>
        [Input("displayNameContains")]
        public string? DisplayNameContains { get; set; }

        /// <summary>
        /// A case-insensitive prefix which can be used to filter on the `display_name` field
        /// </summary>
        [Input("displayNamePrefix")]
        public string? DisplayNamePrefix { get; set; }

        public GetSubscriptionsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSubscriptionsResult
    {
        public readonly string? DisplayNameContains;
        public readonly string? DisplayNamePrefix;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// One or more `subscription` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSubscriptionsSubscriptionResult> Subscriptions;

        [OutputConstructor]
        private GetSubscriptionsResult(
            string? displayNameContains,

            string? displayNamePrefix,

            string id,

            ImmutableArray<Outputs.GetSubscriptionsSubscriptionResult> subscriptions)
        {
            DisplayNameContains = displayNameContains;
            DisplayNamePrefix = displayNamePrefix;
            Id = id;
            Subscriptions = subscriptions;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Cdn
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about an existing CDN Profile.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/cdn_profile.html.markdown.
        /// </summary>
        public static Task<GetProfileResult> GetProfile(GetProfileArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProfileResult>("azure:cdn/getProfile:getProfile", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetProfileArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the CDN Profile.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the CDN Profile exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetProfileArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetProfileResult
    {
        /// <summary>
        /// The Azure Region where the resource exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The pricing related information of current CDN profile.
        /// </summary>
        public readonly string Sku;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetProfileResult(
            string location,
            string name,
            string resourceGroupName,
            string sku,
            ImmutableDictionary<string, string> tags,
            string id)
        {
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Sku = sku;
            Tags = tags;
            Id = id;
        }
    }
}

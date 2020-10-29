// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.NotificationHub
{
    public static class GetNamespace
    {
        /// <summary>
        /// Use this data source to access information about an existing Notification Hub Namespace.
        /// </summary>
        public static Task<GetNamespaceResult> InvokeAsync(GetNamespaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNamespaceResult>("azure:notificationhub/getNamespace:getNamespace", args ?? new GetNamespaceArgs(), options.WithVersion());
    }


    public sealed class GetNamespaceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the Name of the Notification Hub Namespace.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the Name of the Resource Group within which the Notification Hub exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNamespaceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNamespaceResult
    {
        /// <summary>
        /// Is this Notification Hub Namespace enabled?
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure Region in which this Notification Hub Namespace exists.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the SKU to use for this Notification Hub Namespace. Possible values are `Free`, `Basic` or `Standard.`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Type of Namespace, such as `Messaging` or `NotificationHub`.
        /// </summary>
        public readonly string NamespaceType;
        public readonly string ResourceGroupName;
        public readonly string ServicebusEndpoint;
        /// <summary>
        /// A `sku` block as defined below.
        /// </summary>
        public readonly Outputs.GetNamespaceSkuResult Sku;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetNamespaceResult(
            bool enabled,

            string id,

            string location,

            string name,

            string namespaceType,

            string resourceGroupName,

            string servicebusEndpoint,

            Outputs.GetNamespaceSkuResult sku,

            ImmutableDictionary<string, string> tags)
        {
            Enabled = enabled;
            Id = id;
            Location = location;
            Name = name;
            NamespaceType = namespaceType;
            ResourceGroupName = resourceGroupName;
            ServicebusEndpoint = servicebusEndpoint;
            Sku = sku;
            Tags = tags;
        }
    }
}

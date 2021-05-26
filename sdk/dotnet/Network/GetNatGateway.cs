// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    public static class GetNatGateway
    {
        /// <summary>
        /// Use this data source to access information about an existing NAT Gateway.
        /// </summary>
        public static Task<GetNatGatewayResult> InvokeAsync(GetNatGatewayArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNatGatewayResult>("azure:network/getNatGateway:getNatGateway", args ?? new GetNatGatewayArgs(), options.WithVersion());

        public static Output<GetNatGatewayResult> Apply(GetNatGatewayApplyArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Name.Box(),
                args.PublicIpAddressIds.Box(),
                args.PublicIpPrefixIds.Box(),
                args.ResourceGroupName.Box()
            ).Apply(a => {
                    var args = new GetNatGatewayArgs();
                    a[0].Set(args, nameof(args.Name));
                    a[1].Set(args, nameof(args.PublicIpAddressIds));
                    a[2].Set(args, nameof(args.PublicIpPrefixIds));
                    a[3].Set(args, nameof(args.ResourceGroupName));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetNatGatewayArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the Name of the NAT Gateway.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("publicIpAddressIds")]
        private List<string>? _publicIpAddressIds;

        /// <summary>
        /// A list of existing Public IP Address resource IDs which the NAT Gateway is using.
        /// </summary>
        public List<string> PublicIpAddressIds
        {
            get => _publicIpAddressIds ?? (_publicIpAddressIds = new List<string>());
            set => _publicIpAddressIds = value;
        }

        [Input("publicIpPrefixIds")]
        private List<string>? _publicIpPrefixIds;

        /// <summary>
        /// A list of existing Public IP Prefix resource IDs which the NAT Gateway is using.
        /// </summary>
        public List<string> PublicIpPrefixIds
        {
            get => _publicIpPrefixIds ?? (_publicIpPrefixIds = new List<string>());
            set => _publicIpPrefixIds = value;
        }

        /// <summary>
        /// Specifies the name of the Resource Group where the NAT Gateway exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNatGatewayArgs()
        {
        }
    }

    public sealed class GetNatGatewayApplyArgs
    {
        /// <summary>
        /// Specifies the Name of the NAT Gateway.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("publicIpAddressIds")]
        private InputList<string>? _publicIpAddressIds;

        /// <summary>
        /// A list of existing Public IP Address resource IDs which the NAT Gateway is using.
        /// </summary>
        public InputList<string> PublicIpAddressIds
        {
            get => _publicIpAddressIds ?? (_publicIpAddressIds = new InputList<string>());
            set => _publicIpAddressIds = value;
        }

        [Input("publicIpPrefixIds")]
        private InputList<string>? _publicIpPrefixIds;

        /// <summary>
        /// A list of existing Public IP Prefix resource IDs which the NAT Gateway is using.
        /// </summary>
        public InputList<string> PublicIpPrefixIds
        {
            get => _publicIpPrefixIds ?? (_publicIpPrefixIds = new InputList<string>());
            set => _publicIpPrefixIds = value;
        }

        /// <summary>
        /// Specifies the name of the Resource Group where the NAT Gateway exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetNatGatewayApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNatGatewayResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The idle timeout in minutes which is used for the NAT Gateway.
        /// </summary>
        public readonly int IdleTimeoutInMinutes;
        /// <summary>
        /// The location where the NAT Gateway exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        /// <summary>
        /// A list of existing Public IP Address resource IDs which the NAT Gateway is using.
        /// </summary>
        public readonly ImmutableArray<string> PublicIpAddressIds;
        /// <summary>
        /// A list of existing Public IP Prefix resource IDs which the NAT Gateway is using.
        /// </summary>
        public readonly ImmutableArray<string> PublicIpPrefixIds;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The Resource GUID of the NAT Gateway.
        /// </summary>
        public readonly string ResourceGuid;
        /// <summary>
        /// The SKU used by the NAT Gateway.
        /// </summary>
        public readonly string SkuName;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// A list of Availability Zones which the NAT Gateway exists in.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetNatGatewayResult(
            string id,

            int idleTimeoutInMinutes,

            string location,

            string name,

            ImmutableArray<string> publicIpAddressIds,

            ImmutableArray<string> publicIpPrefixIds,

            string resourceGroupName,

            string resourceGuid,

            string skuName,

            ImmutableDictionary<string, string> tags,

            ImmutableArray<string> zones)
        {
            Id = id;
            IdleTimeoutInMinutes = idleTimeoutInMinutes;
            Location = location;
            Name = name;
            PublicIpAddressIds = publicIpAddressIds;
            PublicIpPrefixIds = publicIpPrefixIds;
            ResourceGroupName = resourceGroupName;
            ResourceGuid = resourceGuid;
            SkuName = skuName;
            Tags = tags;
            Zones = zones;
        }
    }
}

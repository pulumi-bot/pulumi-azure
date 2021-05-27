// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    public static class GetVirtualWan
    {
        /// <summary>
        /// Use this data source to access information about an existing Virtual Wan.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.Network.GetVirtualWan.InvokeAsync(new Azure.Network.GetVirtualWanArgs
        ///         {
        ///             Name = "existing",
        ///             ResourceGroupName = "existing",
        ///         }));
        ///         this.Id = example.Apply(example =&gt; example.Id);
        ///         this.AllowBranchToBranchTraffic = example.Apply(example =&gt; example.AllowBranchToBranchTraffic);
        ///         this.DisableVpnEncryption = data.Azurerm_virtual_wan.Exemple.Disable_vpn_encryption;
        ///         this.Location = data.Azurerm_virtual_wan.Exemple.Location;
        ///         this.Office365LocalBreakoutCategory = data.Azurerm_virtual_wan.Exemple.Office365_local_breakout_category;
        ///         this.Sku = data.Azurerm_virtual_wan.Exemple.Sku;
        ///         this.Tags = data.Azurerm_virtual_wan.Exemple.Tags;
        ///         this.VirtualHubs = data.Azurerm_virtual_wan.Exemple.Virtual_hubs;
        ///         this.VpnSites = data.Azurerm_virtual_wan.Exemple.Vpn_sites;
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        ///     [Output("allowBranchToBranchTraffic")]
        ///     public Output&lt;string&gt; AllowBranchToBranchTraffic { get; set; }
        ///     [Output("disableVpnEncryption")]
        ///     public Output&lt;string&gt; DisableVpnEncryption { get; set; }
        ///     [Output("location")]
        ///     public Output&lt;string&gt; Location { get; set; }
        ///     [Output("office365LocalBreakoutCategory")]
        ///     public Output&lt;string&gt; Office365LocalBreakoutCategory { get; set; }
        ///     [Output("sku")]
        ///     public Output&lt;string&gt; Sku { get; set; }
        ///     [Output("tags")]
        ///     public Output&lt;string&gt; Tags { get; set; }
        ///     [Output("virtualHubs")]
        ///     public Output&lt;string&gt; VirtualHubs { get; set; }
        ///     [Output("vpnSites")]
        ///     public Output&lt;string&gt; VpnSites { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVirtualWanResult> InvokeAsync(GetVirtualWanArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVirtualWanResult>("azure:network/getVirtualWan:getVirtualWan", args ?? new GetVirtualWanArgs(), options.WithVersion());

        public static Output<GetVirtualWanResult> Invoke(GetVirtualWanOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Name.Box(),
                args.ResourceGroupName.Box()
            ).Apply(a => {
                    var args = new GetVirtualWanArgs();
                    a[0].Set(args, nameof(args.Name));
                    a[1].Set(args, nameof(args.ResourceGroupName));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetVirtualWanArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this Virtual Wan.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Virtual Wan exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetVirtualWanArgs()
        {
        }
    }

    public sealed class GetVirtualWanOutputArgs
    {
        /// <summary>
        /// The name of this Virtual Wan.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Virtual Wan exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetVirtualWanOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVirtualWanResult
    {
        /// <summary>
        /// Is branch to branch traffic is allowed?
        /// </summary>
        public readonly bool AllowBranchToBranchTraffic;
        /// <summary>
        /// Is VPN Encryption disabled?
        /// </summary>
        public readonly bool DisableVpnEncryption;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure Region where the Virtual Wan exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        /// <summary>
        /// The Office365 Local Breakout Category.
        /// </summary>
        public readonly string Office365LocalBreakoutCategory;
        public readonly string ResourceGroupName;
        /// <summary>
        /// Type of Virtual Wan (Basic or Standard).
        /// </summary>
        public readonly string Sku;
        /// <summary>
        /// A mapping of tags assigned to the Virtual Wan.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// A list of Virtual Hubs ID's attached to this Virtual WAN.
        /// </summary>
        public readonly ImmutableArray<string> VirtualHubIds;
        /// <summary>
        /// A list of VPN Site ID's attached to this Virtual WAN.
        /// </summary>
        public readonly ImmutableArray<string> VpnSiteIds;

        [OutputConstructor]
        private GetVirtualWanResult(
            bool allowBranchToBranchTraffic,

            bool disableVpnEncryption,

            string id,

            string location,

            string name,

            string office365LocalBreakoutCategory,

            string resourceGroupName,

            string sku,

            ImmutableDictionary<string, string> tags,

            ImmutableArray<string> virtualHubIds,

            ImmutableArray<string> vpnSiteIds)
        {
            AllowBranchToBranchTraffic = allowBranchToBranchTraffic;
            DisableVpnEncryption = disableVpnEncryption;
            Id = id;
            Location = location;
            Name = name;
            Office365LocalBreakoutCategory = office365LocalBreakoutCategory;
            ResourceGroupName = resourceGroupName;
            Sku = sku;
            Tags = tags;
            VirtualHubIds = virtualHubIds;
            VpnSiteIds = vpnSiteIds;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    /// <summary>
    /// Manages a Virtual WAN.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var exampleVirtualWan = new Azure.Network.VirtualWan("exampleVirtualWan", new Azure.Network.VirtualWanArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Virtual WAN can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:network/virtualWan:VirtualWan example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualWans/testvwan
    /// ```
    /// </summary>
    [AzureResourceType("azure:network/virtualWan:VirtualWan")]
    public partial class VirtualWan : Pulumi.CustomResource
    {
        /// <summary>
        /// Boolean flag to specify whether branch to branch traffic is allowed. Defaults to `true`.
        /// </summary>
        [Output("allowBranchToBranchTraffic")]
        public Output<bool?> AllowBranchToBranchTraffic { get; private set; } = null!;

        [Output("allowVnetToVnetTraffic")]
        public Output<bool?> AllowVnetToVnetTraffic { get; private set; } = null!;

        /// <summary>
        /// Boolean flag to specify whether VPN encryption is disabled. Defaults to `false`.
        /// </summary>
        [Output("disableVpnEncryption")]
        public Output<bool?> DisableVpnEncryption { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Virtual WAN. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the Office365 local breakout category. Possible values include: `Optimize`, `OptimizeAndAllow`, `All`, `None`. Defaults to `None`.
        /// </summary>
        [Output("office365LocalBreakoutCategory")]
        public Output<string?> Office365LocalBreakoutCategory { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Virtual WAN. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the Virtual WAN.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies the Virtual WAN type. Possible Values include: `Basic` and `Standard`. Defaults to `Standard`.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualWan resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualWan(string name, VirtualWanArgs args, CustomResourceOptions? options = null)
            : base("azure:network/virtualWan:VirtualWan", name, args ?? new VirtualWanArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualWan(string name, Input<string> id, VirtualWanState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/virtualWan:VirtualWan", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualWan resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualWan Get(string name, Input<string> id, VirtualWanState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualWan(name, id, state, options);
        }
    }

    public sealed class VirtualWanArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean flag to specify whether branch to branch traffic is allowed. Defaults to `true`.
        /// </summary>
        [Input("allowBranchToBranchTraffic")]
        public Input<bool>? AllowBranchToBranchTraffic { get; set; }

        [Input("allowVnetToVnetTraffic")]
        public Input<bool>? AllowVnetToVnetTraffic { get; set; }

        /// <summary>
        /// Boolean flag to specify whether VPN encryption is disabled. Defaults to `false`.
        /// </summary>
        [Input("disableVpnEncryption")]
        public Input<bool>? DisableVpnEncryption { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Virtual WAN. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the Office365 local breakout category. Possible values include: `Optimize`, `OptimizeAndAllow`, `All`, `None`. Defaults to `None`.
        /// </summary>
        [Input("office365LocalBreakoutCategory")]
        public Input<string>? Office365LocalBreakoutCategory { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Virtual WAN. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the Virtual WAN.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the Virtual WAN type. Possible Values include: `Basic` and `Standard`. Defaults to `Standard`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public VirtualWanArgs()
        {
        }
    }

    public sealed class VirtualWanState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean flag to specify whether branch to branch traffic is allowed. Defaults to `true`.
        /// </summary>
        [Input("allowBranchToBranchTraffic")]
        public Input<bool>? AllowBranchToBranchTraffic { get; set; }

        [Input("allowVnetToVnetTraffic")]
        public Input<bool>? AllowVnetToVnetTraffic { get; set; }

        /// <summary>
        /// Boolean flag to specify whether VPN encryption is disabled. Defaults to `false`.
        /// </summary>
        [Input("disableVpnEncryption")]
        public Input<bool>? DisableVpnEncryption { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Virtual WAN. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the Office365 local breakout category. Possible values include: `Optimize`, `OptimizeAndAllow`, `All`, `None`. Defaults to `None`.
        /// </summary>
        [Input("office365LocalBreakoutCategory")]
        public Input<string>? Office365LocalBreakoutCategory { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Virtual WAN. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the Virtual WAN.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the Virtual WAN type. Possible Values include: `Basic` and `Standard`. Defaults to `Standard`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public VirtualWanState()
        {
        }
    }
}

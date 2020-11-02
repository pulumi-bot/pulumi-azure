// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.PrivateDns
{
    /// <summary>
    /// Enables you to manage Private DNS zones within Azure DNS. These zones are hosted on Azure's name servers.
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
    ///             Location = "West US",
    ///         });
    ///         var exampleZone = new Azure.PrivateDns.Zone("exampleZone", new Azure.PrivateDns.ZoneArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Private DNS Zones can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class Zone : Pulumi.CustomResource
    {
        /// <summary>
        /// The maximum number of record sets that can be created in this Private DNS zone.
        /// </summary>
        [Output("maxNumberOfRecordSets")]
        public Output<int> MaxNumberOfRecordSets { get; private set; } = null!;

        /// <summary>
        /// The maximum number of virtual networks that can be linked to this Private DNS zone.
        /// </summary>
        [Output("maxNumberOfVirtualNetworkLinks")]
        public Output<int> MaxNumberOfVirtualNetworkLinks { get; private set; } = null!;

        /// <summary>
        /// The maximum number of virtual networks that can be linked to this Private DNS zone with registration enabled.
        /// </summary>
        [Output("maxNumberOfVirtualNetworkLinksWithRegistration")]
        public Output<int> MaxNumberOfVirtualNetworkLinksWithRegistration { get; private set; } = null!;

        /// <summary>
        /// The name of the Private DNS Zone. Must be a valid domain name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The current number of record sets in this Private DNS zone.
        /// </summary>
        [Output("numberOfRecordSets")]
        public Output<int> NumberOfRecordSets { get; private set; } = null!;

        /// <summary>
        /// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Zone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Zone(string name, ZoneArgs args, CustomResourceOptions? options = null)
            : base("azure:privatedns/zone:Zone", name, args ?? new ZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Zone(string name, Input<string> id, ZoneState? state = null, CustomResourceOptions? options = null)
            : base("azure:privatedns/zone:Zone", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Zone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Zone Get(string name, Input<string> id, ZoneState? state = null, CustomResourceOptions? options = null)
        {
            return new Zone(name, id, state, options);
        }
    }

    public sealed class ZoneArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Private DNS Zone. Must be a valid domain name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ZoneArgs()
        {
        }
    }

    public sealed class ZoneState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum number of record sets that can be created in this Private DNS zone.
        /// </summary>
        [Input("maxNumberOfRecordSets")]
        public Input<int>? MaxNumberOfRecordSets { get; set; }

        /// <summary>
        /// The maximum number of virtual networks that can be linked to this Private DNS zone.
        /// </summary>
        [Input("maxNumberOfVirtualNetworkLinks")]
        public Input<int>? MaxNumberOfVirtualNetworkLinks { get; set; }

        /// <summary>
        /// The maximum number of virtual networks that can be linked to this Private DNS zone with registration enabled.
        /// </summary>
        [Input("maxNumberOfVirtualNetworkLinksWithRegistration")]
        public Input<int>? MaxNumberOfVirtualNetworkLinksWithRegistration { get; set; }

        /// <summary>
        /// The name of the Private DNS Zone. Must be a valid domain name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The current number of record sets in this Private DNS zone.
        /// </summary>
        [Input("numberOfRecordSets")]
        public Input<int>? NumberOfRecordSets { get; set; }

        /// <summary>
        /// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ZoneState()
        {
        }
    }
}

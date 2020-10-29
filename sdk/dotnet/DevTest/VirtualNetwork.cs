// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DevTest
{
    /// <summary>
    /// Manages a Virtual Network within a DevTest Lab.
    /// </summary>
    public partial class VirtualNetwork : Pulumi.CustomResource
    {
        /// <summary>
        /// A description for the Virtual Network.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Dev Test Lab in which the Virtual Network should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("labName")]
        public Output<string> LabName { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A `subnet` block as defined below.
        /// </summary>
        [Output("subnet")]
        public Output<Outputs.VirtualNetworkSubnet> Subnet { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The unique immutable identifier of the Dev Test Virtual Network.
        /// </summary>
        [Output("uniqueIdentifier")]
        public Output<string> UniqueIdentifier { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualNetwork resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualNetwork(string name, VirtualNetworkArgs args, CustomResourceOptions? options = null)
            : base("azure:devtest/virtualNetwork:VirtualNetwork", name, args ?? new VirtualNetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualNetwork(string name, Input<string> id, VirtualNetworkState? state = null, CustomResourceOptions? options = null)
            : base("azure:devtest/virtualNetwork:VirtualNetwork", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualNetwork resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualNetwork Get(string name, Input<string> id, VirtualNetworkState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualNetwork(name, id, state, options);
        }
    }

    public sealed class VirtualNetworkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description for the Virtual Network.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the name of the Dev Test Lab in which the Virtual Network should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("labName", required: true)]
        public Input<string> LabName { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `subnet` block as defined below.
        /// </summary>
        [Input("subnet")]
        public Input<Inputs.VirtualNetworkSubnetArgs>? Subnet { get; set; }

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

        public VirtualNetworkArgs()
        {
        }
    }

    public sealed class VirtualNetworkState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description for the Virtual Network.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the name of the Dev Test Lab in which the Virtual Network should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("labName")]
        public Input<string>? LabName { get; set; }

        /// <summary>
        /// Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// A `subnet` block as defined below.
        /// </summary>
        [Input("subnet")]
        public Input<Inputs.VirtualNetworkSubnetGetArgs>? Subnet { get; set; }

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

        /// <summary>
        /// The unique immutable identifier of the Dev Test Virtual Network.
        /// </summary>
        [Input("uniqueIdentifier")]
        public Input<string>? UniqueIdentifier { get; set; }

        public VirtualNetworkState()
        {
        }
    }
}

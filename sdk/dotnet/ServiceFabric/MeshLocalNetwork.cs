// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ServiceFabric
{
    /// <summary>
    /// Manages a Service Fabric Mesh Local Network.
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
    ///         var exampleMeshLocalNetwork = new Azure.ServiceFabric.MeshLocalNetwork("exampleMeshLocalNetwork", new Azure.ServiceFabric.MeshLocalNetworkArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             NetworkAddressPrefix = "10.0.0.0/22",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Service Fabric Mesh Local Network can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:servicefabric/meshLocalNetwork:MeshLocalNetwork network1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceFabricMesh/networks/network1
    /// ```
    /// </summary>
    [AzureResourceType("azure:servicefabric/meshLocalNetwork:MeshLocalNetwork")]
    public partial class MeshLocalNetwork : Pulumi.CustomResource
    {
        /// <summary>
        /// A description of this Service Fabric Mesh Local Network.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies the Azure Region where the Service Fabric Mesh Local Network should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the Service Fabric Mesh Local Network. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The address space for the local container network.
        /// </summary>
        [Output("networkAddressPrefix")]
        public Output<string> NetworkAddressPrefix { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Service Fabric Mesh Local Network exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a MeshLocalNetwork resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MeshLocalNetwork(string name, MeshLocalNetworkArgs args, CustomResourceOptions? options = null)
            : base("azure:servicefabric/meshLocalNetwork:MeshLocalNetwork", name, args ?? new MeshLocalNetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MeshLocalNetwork(string name, Input<string> id, MeshLocalNetworkState? state = null, CustomResourceOptions? options = null)
            : base("azure:servicefabric/meshLocalNetwork:MeshLocalNetwork", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MeshLocalNetwork resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MeshLocalNetwork Get(string name, Input<string> id, MeshLocalNetworkState? state = null, CustomResourceOptions? options = null)
        {
            return new MeshLocalNetwork(name, id, state, options);
        }
    }

    public sealed class MeshLocalNetworkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of this Service Fabric Mesh Local Network.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the Azure Region where the Service Fabric Mesh Local Network should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Service Fabric Mesh Local Network. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The address space for the local container network.
        /// </summary>
        [Input("networkAddressPrefix", required: true)]
        public Input<string> NetworkAddressPrefix { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Service Fabric Mesh Local Network exists. Changing this forces a new resource to be created.
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

        public MeshLocalNetworkArgs()
        {
        }
    }

    public sealed class MeshLocalNetworkState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of this Service Fabric Mesh Local Network.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the Azure Region where the Service Fabric Mesh Local Network should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Service Fabric Mesh Local Network. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The address space for the local container network.
        /// </summary>
        [Input("networkAddressPrefix")]
        public Input<string>? NetworkAddressPrefix { get; set; }

        /// <summary>
        /// The name of the Resource Group in which the Service Fabric Mesh Local Network exists. Changing this forces a new resource to be created.
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

        public MeshLocalNetworkState()
        {
        }
    }
}

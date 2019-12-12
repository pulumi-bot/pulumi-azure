// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    /// <summary>
    /// Manages a Virtual Hub within a Virtual WAN.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/virtual_hub.html.markdown.
    /// </summary>
    public partial class VirtualHub : Pulumi.CustomResource
    {
        /// <summary>
        /// The Address Prefix which should be used for this Virtual Hub.
        /// </summary>
        [Output("addressPrefix")]
        public Output<string> AddressPrefix { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the Virtual Hub should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the Virtual Hub. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group where the Virtual Hub should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// One or more `route` blocks as defined below.
        /// </summary>
        [Output("routes")]
        public Output<ImmutableArray<Outputs.VirtualHubRoutes>> Routes { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the Virtual Hub.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of a Virtual WAN within which the Virtual Hub should be created.
        /// </summary>
        [Output("virtualWanId")]
        public Output<string> VirtualWanId { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualHub resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualHub(string name, VirtualHubArgs args, CustomResourceOptions? options = null)
            : base("azure:network/virtualHub:VirtualHub", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private VirtualHub(string name, Input<string> id, VirtualHubState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/virtualHub:VirtualHub", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualHub resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualHub Get(string name, Input<string> id, VirtualHubState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualHub(name, id, state, options);
        }
    }

    public sealed class VirtualHubArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Address Prefix which should be used for this Virtual Hub.
        /// </summary>
        [Input("addressPrefix", required: true)]
        public Input<string> AddressPrefix { get; set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the Virtual Hub should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Virtual Hub. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group where the Virtual Hub should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("routes")]
        private InputList<Inputs.VirtualHubRoutesArgs>? _routes;

        /// <summary>
        /// One or more `route` blocks as defined below.
        /// </summary>
        public InputList<Inputs.VirtualHubRoutesArgs> Routes
        {
            get => _routes ?? (_routes = new InputList<Inputs.VirtualHubRoutesArgs>());
            set => _routes = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the Virtual Hub.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of a Virtual WAN within which the Virtual Hub should be created.
        /// </summary>
        [Input("virtualWanId", required: true)]
        public Input<string> VirtualWanId { get; set; } = null!;

        public VirtualHubArgs()
        {
        }
    }

    public sealed class VirtualHubState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Address Prefix which should be used for this Virtual Hub.
        /// </summary>
        [Input("addressPrefix")]
        public Input<string>? AddressPrefix { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the Virtual Hub should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Virtual Hub. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group where the Virtual Hub should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("routes")]
        private InputList<Inputs.VirtualHubRoutesGetArgs>? _routes;

        /// <summary>
        /// One or more `route` blocks as defined below.
        /// </summary>
        public InputList<Inputs.VirtualHubRoutesGetArgs> Routes
        {
            get => _routes ?? (_routes = new InputList<Inputs.VirtualHubRoutesGetArgs>());
            set => _routes = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the Virtual Hub.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of a Virtual WAN within which the Virtual Hub should be created.
        /// </summary>
        [Input("virtualWanId")]
        public Input<string>? VirtualWanId { get; set; }

        public VirtualHubState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class VirtualHubRoutesArgs : Pulumi.ResourceArgs
    {
        [Input("addressPrefixes", required: true)]
        private InputList<string>? _addressPrefixes;
        public InputList<string> AddressPrefixes
        {
            get => _addressPrefixes ?? (_addressPrefixes = new InputList<string>());
            set => _addressPrefixes = value;
        }

        [Input("nextHopIpAddress", required: true)]
        public Input<string> NextHopIpAddress { get; set; } = null!;

        public VirtualHubRoutesArgs()
        {
        }
    }

    public sealed class VirtualHubRoutesGetArgs : Pulumi.ResourceArgs
    {
        [Input("addressPrefixes", required: true)]
        private InputList<string>? _addressPrefixes;
        public InputList<string> AddressPrefixes
        {
            get => _addressPrefixes ?? (_addressPrefixes = new InputList<string>());
            set => _addressPrefixes = value;
        }

        [Input("nextHopIpAddress", required: true)]
        public Input<string> NextHopIpAddress { get; set; } = null!;

        public VirtualHubRoutesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class VirtualHubRoutes
    {
        public readonly ImmutableArray<string> AddressPrefixes;
        public readonly string NextHopIpAddress;

        [OutputConstructor]
        private VirtualHubRoutes(
            ImmutableArray<string> addressPrefixes,
            string nextHopIpAddress)
        {
            AddressPrefixes = addressPrefixes;
            NextHopIpAddress = nextHopIpAddress;
        }
    }
    }
}

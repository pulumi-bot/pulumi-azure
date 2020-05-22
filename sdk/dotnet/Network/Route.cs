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
    /// Manages a Route within a Route Table.
    /// 
    /// &gt; **NOTE on Route Tables and Routes:** This provider currently
    /// provides both a standalone Route resource, and allows for Routes to be defined in-line within the Route Table resource.
    /// At this time you cannot use a Route Table with in-line Routes in conjunction with any Route resources. Doing so will cause a conflict of Route configurations and will overwrite Routes.
    /// 
    /// ## Example Usage
    /// 
    /// 
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
    ///         var exampleRouteTable = new Azure.Network.RouteTable("exampleRouteTable", new Azure.Network.RouteTableArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var exampleRoute = new Azure.Network.Route("exampleRoute", new Azure.Network.RouteArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             RouteTableName = exampleRouteTable.Name,
    ///             AddressPrefix = "10.1.0.0/16",
    ///             NextHopType = "vnetlocal",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Route : Pulumi.CustomResource
    {
        /// <summary>
        /// The destination CIDR to which the route applies, such as `10.1.0.0/16`
        /// </summary>
        [Output("addressPrefix")]
        public Output<string> AddressPrefix { get; private set; } = null!;

        /// <summary>
        /// The name of the route. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is `VirtualAppliance`.
        /// </summary>
        [Output("nextHopInIpAddress")]
        public Output<string?> NextHopInIpAddress { get; private set; } = null!;

        /// <summary>
        /// The type of Azure hop the packet should be sent to. Possible values are `VirtualNetworkGateway`, `VnetLocal`, `Internet`, `VirtualAppliance` and `None`
        /// </summary>
        [Output("nextHopType")]
        public Output<string> NextHopType { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the route. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The name of the route table within which create the route. Changing this forces a new resource to be created.
        /// </summary>
        [Output("routeTableName")]
        public Output<string> RouteTableName { get; private set; } = null!;


        /// <summary>
        /// Create a Route resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Route(string name, RouteArgs args, CustomResourceOptions? options = null)
            : base("azure:network/route:Route", name, args ?? new RouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Route(string name, Input<string> id, RouteState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/route:Route", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Route resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Route Get(string name, Input<string> id, RouteState? state = null, CustomResourceOptions? options = null)
        {
            return new Route(name, id, state, options);
        }
    }

    public sealed class RouteArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination CIDR to which the route applies, such as `10.1.0.0/16`
        /// </summary>
        [Input("addressPrefix", required: true)]
        public Input<string> AddressPrefix { get; set; } = null!;

        /// <summary>
        /// The name of the route. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is `VirtualAppliance`.
        /// </summary>
        [Input("nextHopInIpAddress")]
        public Input<string>? NextHopInIpAddress { get; set; }

        /// <summary>
        /// The type of Azure hop the packet should be sent to. Possible values are `VirtualNetworkGateway`, `VnetLocal`, `Internet`, `VirtualAppliance` and `None`
        /// </summary>
        [Input("nextHopType", required: true)]
        public Input<string> NextHopType { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the route. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the route table within which create the route. Changing this forces a new resource to be created.
        /// </summary>
        [Input("routeTableName", required: true)]
        public Input<string> RouteTableName { get; set; } = null!;

        public RouteArgs()
        {
        }
    }

    public sealed class RouteState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination CIDR to which the route applies, such as `10.1.0.0/16`
        /// </summary>
        [Input("addressPrefix")]
        public Input<string>? AddressPrefix { get; set; }

        /// <summary>
        /// The name of the route. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is `VirtualAppliance`.
        /// </summary>
        [Input("nextHopInIpAddress")]
        public Input<string>? NextHopInIpAddress { get; set; }

        /// <summary>
        /// The type of Azure hop the packet should be sent to. Possible values are `VirtualNetworkGateway`, `VnetLocal`, `Internet`, `VirtualAppliance` and `None`
        /// </summary>
        [Input("nextHopType")]
        public Input<string>? NextHopType { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the route. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The name of the route table within which create the route. Changing this forces a new resource to be created.
        /// </summary>
        [Input("routeTableName")]
        public Input<string>? RouteTableName { get; set; }

        public RouteState()
        {
        }
    }
}

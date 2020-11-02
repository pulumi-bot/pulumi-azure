// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    public static class GetRouteTable
    {
        /// <summary>
        /// Use this data source to access information about an existing Route Table.
        /// 
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
        ///         var example = Output.Create(Azure.Network.GetRouteTable.InvokeAsync(new Azure.Network.GetRouteTableArgs
        ///         {
        ///             Name = "myroutetable",
        ///             ResourceGroupName = "some-resource-group",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRouteTableResult> InvokeAsync(GetRouteTableArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRouteTableResult>("azure:network/getRouteTable:getRouteTable", args ?? new GetRouteTableArgs(), options.WithVersion());
    }


    public sealed class GetRouteTableArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Route Table.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Route Table exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetRouteTableArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRouteTableResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure Region in which the Route Table exists.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the Route.
        /// </summary>
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// One or more `route` blocks as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouteTableRouteResult> Routes;
        /// <summary>
        /// The collection of Subnets associated with this route table.
        /// </summary>
        public readonly ImmutableArray<string> Subnets;
        /// <summary>
        /// A mapping of tags assigned to the Route Table.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetRouteTableResult(
            string id,

            string location,

            string name,

            string resourceGroupName,

            ImmutableArray<Outputs.GetRouteTableRouteResult> routes,

            ImmutableArray<string> subnets,

            ImmutableDictionary<string, string> tags)
        {
            Id = id;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Routes = routes;
            Subnets = subnets;
            Tags = tags;
        }
    }
}

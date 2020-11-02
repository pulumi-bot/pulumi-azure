// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    public static class GetExpressRouteCircuit
    {
        /// <summary>
        /// Use this data source to access information about an existing ExpressRoute circuit.
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
        ///         var example = Output.Create(Azure.Network.GetExpressRouteCircuit.InvokeAsync(new Azure.Network.GetExpressRouteCircuitArgs
        ///         {
        ///             ResourceGroupName = azurerm_resource_group.Example.Name,
        ///             Name = azurerm_express_route_circuit.Example.Name,
        ///         }));
        ///         this.ExpressRouteCircuitId = example.Apply(example =&gt; example.Id);
        ///         this.ServiceKey = example.Apply(example =&gt; example.ServiceKey);
        ///     }
        /// 
        ///     [Output("expressRouteCircuitId")]
        ///     public Output&lt;string&gt; ExpressRouteCircuitId { get; set; }
        ///     [Output("serviceKey")]
        ///     public Output&lt;string&gt; ServiceKey { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetExpressRouteCircuitResult> InvokeAsync(GetExpressRouteCircuitArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetExpressRouteCircuitResult>("azure:network/getExpressRouteCircuit:getExpressRouteCircuit", args ?? new GetExpressRouteCircuitArgs(), options.WithVersion());
    }


    public sealed class GetExpressRouteCircuitArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the ExpressRoute circuit.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the ExpressRoute circuit exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetExpressRouteCircuitArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetExpressRouteCircuitResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure location where the ExpressRoute circuit exists
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        /// <summary>
        /// A `peerings` block for the ExpressRoute circuit as documented below
        /// </summary>
        public readonly ImmutableArray<Outputs.GetExpressRouteCircuitPeeringResult> Peerings;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The string needed by the service provider to provision the ExpressRoute circuit.
        /// </summary>
        public readonly string ServiceKey;
        /// <summary>
        /// A `service_provider_properties` block for the ExpressRoute circuit as documented below
        /// </summary>
        public readonly ImmutableArray<Outputs.GetExpressRouteCircuitServiceProviderPropertyResult> ServiceProviderProperties;
        /// <summary>
        /// The ExpressRoute circuit provisioning state from your chosen service provider. Possible values are "NotProvisioned", "Provisioning", "Provisioned", and "Deprovisioning".
        /// </summary>
        public readonly string ServiceProviderProvisioningState;
        /// <summary>
        /// A `sku` block for the ExpressRoute circuit as documented below.
        /// </summary>
        public readonly Outputs.GetExpressRouteCircuitSkuResult Sku;

        [OutputConstructor]
        private GetExpressRouteCircuitResult(
            string id,

            string location,

            string name,

            ImmutableArray<Outputs.GetExpressRouteCircuitPeeringResult> peerings,

            string resourceGroupName,

            string serviceKey,

            ImmutableArray<Outputs.GetExpressRouteCircuitServiceProviderPropertyResult> serviceProviderProperties,

            string serviceProviderProvisioningState,

            Outputs.GetExpressRouteCircuitSkuResult sku)
        {
            Id = id;
            Location = location;
            Name = name;
            Peerings = peerings;
            ResourceGroupName = resourceGroupName;
            ServiceKey = serviceKey;
            ServiceProviderProperties = serviceProviderProperties;
            ServiceProviderProvisioningState = serviceProviderProvisioningState;
            Sku = sku;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Lb
{
    public static class GetLBRule
    {
        /// <summary>
        /// Use this data source to access information about an existing Load Balancer Rule.
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
        ///         var exampleLB = Output.Create(Azure.Lb.GetLB.InvokeAsync(new Azure.Lb.GetLBArgs
        ///         {
        ///             Name = "example-lb",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///         var exampleLBRule = exampleLB.Apply(exampleLB =&gt; Output.Create(Azure.Lb.GetLBRule.InvokeAsync(new Azure.Lb.GetLBRuleArgs
        ///         {
        ///             Name = "first",
        ///             ResourceGroupName = "example-resources",
        ///             LoadbalancerId = exampleLB.Id,
        ///         })));
        ///         this.LbRuleId = exampleLBRule.Apply(exampleLBRule =&gt; exampleLBRule.Id);
        ///     }
        /// 
        ///     [Output("lbRuleId")]
        ///     public Output&lt;string&gt; LbRuleId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetLBRuleResult> InvokeAsync(GetLBRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLBRuleResult>("azure:lb/getLBRule:getLBRule", args ?? new GetLBRuleArgs(), options.WithVersion());
    }


    public sealed class GetLBRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Load Balancer Rule.
        /// </summary>
        [Input("loadbalancerId", required: true)]
        public string LoadbalancerId { get; set; } = null!;

        /// <summary>
        /// The name of this Load Balancer Rule.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Load Balancer Rule exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetLBRuleArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLBRuleResult
    {
        /// <summary>
        /// A reference to a Backend Address Pool over which this Load Balancing Rule operates.
        /// </summary>
        public readonly string BackendAddressPoolId;
        /// <summary>
        /// The port used for internal connections on the endpoint.
        /// </summary>
        public readonly int BackendPort;
        /// <summary>
        /// If outbound SNAT is enabled for this Load Balancer Rule.
        /// </summary>
        public readonly bool DisableOutboundSnat;
        /// <summary>
        /// If Floating IPs are enabled for this Load Balancer Rule
        /// </summary>
        public readonly bool EnableFloatingIp;
        /// <summary>
        /// If TCP Reset is enabled for this Load Balancer Rule.
        /// </summary>
        public readonly bool EnableTcpReset;
        /// <summary>
        /// The name of the frontend IP configuration to which the rule is associated.
        /// </summary>
        public readonly string FrontendIpConfigurationName;
        /// <summary>
        /// The port for the external endpoint.
        /// </summary>
        public readonly int FrontendPort;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Specifies the idle timeout in minutes for TCP connections.
        /// </summary>
        public readonly int IdleTimeoutInMinutes;
        /// <summary>
        /// Specifies the load balancing distribution type used by the Load Balancer.
        /// </summary>
        public readonly string LoadDistribution;
        public readonly string LoadbalancerId;
        public readonly string Name;
        /// <summary>
        /// A reference to a Probe used by this Load Balancing Rule.
        /// </summary>
        public readonly string ProbeId;
        /// <summary>
        /// The transport protocol for the external endpoint.
        /// </summary>
        public readonly string Protocol;
        public readonly string ResourceGroupName;

        [OutputConstructor]
        private GetLBRuleResult(
            string backendAddressPoolId,

            int backendPort,

            bool disableOutboundSnat,

            bool enableFloatingIp,

            bool enableTcpReset,

            string frontendIpConfigurationName,

            int frontendPort,

            string id,

            int idleTimeoutInMinutes,

            string loadDistribution,

            string loadbalancerId,

            string name,

            string probeId,

            string protocol,

            string resourceGroupName)
        {
            BackendAddressPoolId = backendAddressPoolId;
            BackendPort = backendPort;
            DisableOutboundSnat = disableOutboundSnat;
            EnableFloatingIp = enableFloatingIp;
            EnableTcpReset = enableTcpReset;
            FrontendIpConfigurationName = frontendIpConfigurationName;
            FrontendPort = frontendPort;
            Id = id;
            IdleTimeoutInMinutes = idleTimeoutInMinutes;
            LoadDistribution = loadDistribution;
            LoadbalancerId = loadbalancerId;
            Name = name;
            ProbeId = probeId;
            Protocol = protocol;
            ResourceGroupName = resourceGroupName;
        }
    }
}

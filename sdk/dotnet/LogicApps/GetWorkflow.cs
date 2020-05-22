// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.LogicApps
{
    public static class GetWorkflow
    {
        /// <summary>
        /// Use this data source to access information about an existing Logic App Workflow.
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
        ///         var example = Output.Create(Azure.LogicApps.GetWorkflow.InvokeAsync(new Azure.LogicApps.GetWorkflowArgs
        ///         {
        ///             Name = "workflow1",
        ///             ResourceGroupName = "my-resource-group",
        ///         }));
        ///         this.AccessEndpoint = example.Apply(example =&gt; example.AccessEndpoint);
        ///     }
        /// 
        ///     [Output("accessEndpoint")]
        ///     public Output&lt;string&gt; AccessEndpoint { get; set; }
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetWorkflowResult> InvokeAsync(GetWorkflowArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWorkflowResult>("azure:logicapps/getWorkflow:getWorkflow", args ?? new GetWorkflowArgs(), options.WithVersion());
    }


    public sealed class GetWorkflowArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Logic App Workflow.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Logic App Workflow exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetWorkflowArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWorkflowResult
    {
        /// <summary>
        /// The Access Endpoint for the Logic App Workflow
        /// </summary>
        public readonly string AccessEndpoint;
        /// <summary>
        /// The list of access endpoint ip addresses of connector.
        /// </summary>
        public readonly ImmutableArray<string> ConnectorEndpointIpAddresses;
        /// <summary>
        /// The list of outgoing ip addresses of connector.
        /// </summary>
        public readonly ImmutableArray<string> ConnectorOutboundIpAddresses;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure location where the Logic App Workflow exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        /// <summary>
        /// A map of Key-Value pairs.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Parameters;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The list of access endpoint ip addresses of workflow.
        /// </summary>
        public readonly ImmutableArray<string> WorkflowEndpointIpAddresses;
        /// <summary>
        /// The list of outgoing ip addresses of workflow.
        /// </summary>
        public readonly ImmutableArray<string> WorkflowOutboundIpAddresses;
        /// <summary>
        /// The Schema used for this Logic App Workflow.
        /// </summary>
        public readonly string WorkflowSchema;
        /// <summary>
        /// The version of the Schema used for this Logic App Workflow. Defaults to `1.0.0.0`.
        /// </summary>
        public readonly string WorkflowVersion;

        [OutputConstructor]
        private GetWorkflowResult(
            string accessEndpoint,

            ImmutableArray<string> connectorEndpointIpAddresses,

            ImmutableArray<string> connectorOutboundIpAddresses,

            string id,

            string location,

            string name,

            ImmutableDictionary<string, string> parameters,

            string resourceGroupName,

            ImmutableDictionary<string, string> tags,

            ImmutableArray<string> workflowEndpointIpAddresses,

            ImmutableArray<string> workflowOutboundIpAddresses,

            string workflowSchema,

            string workflowVersion)
        {
            AccessEndpoint = accessEndpoint;
            ConnectorEndpointIpAddresses = connectorEndpointIpAddresses;
            ConnectorOutboundIpAddresses = connectorOutboundIpAddresses;
            Id = id;
            Location = location;
            Name = name;
            Parameters = parameters;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
            WorkflowEndpointIpAddresses = workflowEndpointIpAddresses;
            WorkflowOutboundIpAddresses = workflowOutboundIpAddresses;
            WorkflowSchema = workflowSchema;
            WorkflowVersion = workflowVersion;
        }
    }
}

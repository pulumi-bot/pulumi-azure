// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataBricks
{
    public static class GetWorkspace
    {
        /// <summary>
        /// Use this data source to access information about an existing Databricks workspace.
        /// </summary>
        public static Task<GetWorkspaceResult> InvokeAsync(GetWorkspaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWorkspaceResult>("azure:databricks/getWorkspace:getWorkspace", args ?? new GetWorkspaceArgs(), options.WithVersion());
    }


    public sealed class GetWorkspaceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Databricks Workspace.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the Databricks Workspace exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetWorkspaceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWorkspaceResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// SKU of this Databricks Workspace.
        /// </summary>
        public readonly string Sku;
        /// <summary>
        /// Unique ID of this Databricks Workspace in Databricks management plane.
        /// </summary>
        public readonly string WorkspaceId;
        /// <summary>
        /// URL this Databricks Workspace is accessible on.
        /// </summary>
        public readonly string WorkspaceUrl;

        [OutputConstructor]
        private GetWorkspaceResult(
            string id,

            string name,

            string resourceGroupName,

            string sku,

            string workspaceId,

            string workspaceUrl)
        {
            Id = id;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Sku = sku;
            WorkspaceId = workspaceId;
            WorkspaceUrl = workspaceUrl;
        }
    }
}

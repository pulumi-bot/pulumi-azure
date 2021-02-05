// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Synapse
{
    /// <summary>
    /// Allows you to Manages a Synapse Managed Private Endpoint.
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
    ///         var exampleAccount = new Azure.Storage.Account("exampleAccount", new Azure.Storage.AccountArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             AccountTier = "Standard",
    ///             AccountReplicationType = "LRS",
    ///             AccountKind = "StorageV2",
    ///             IsHnsEnabled = true,
    ///         });
    ///         var exampleDataLakeGen2Filesystem = new Azure.Storage.DataLakeGen2Filesystem("exampleDataLakeGen2Filesystem", new Azure.Storage.DataLakeGen2FilesystemArgs
    ///         {
    ///             StorageAccountId = exampleAccount.Id,
    ///         });
    ///         var exampleWorkspace = new Azure.Synapse.Workspace("exampleWorkspace", new Azure.Synapse.WorkspaceArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             StorageDataLakeGen2FilesystemId = exampleDataLakeGen2Filesystem.Id,
    ///             SqlAdministratorLogin = "sqladminuser",
    ///             SqlAdministratorLoginPassword = "H@Sh1CoR3!",
    ///             ManagedVirtualNetworkEnabled = true,
    ///         });
    ///         var exampleFirewallRule = new Azure.Synapse.FirewallRule("exampleFirewallRule", new Azure.Synapse.FirewallRuleArgs
    ///         {
    ///             SynapseWorkspaceId = azurerm_synapse_workspace.Test.Id,
    ///             StartIpAddress = "0.0.0.0",
    ///             EndIpAddress = "255.255.255.255",
    ///         });
    ///         var exampleConnect = new Azure.Storage.Account("exampleConnect", new Azure.Storage.AccountArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             AccountTier = "Standard",
    ///             AccountReplicationType = "LRS",
    ///             AccountKind = "BlobStorage",
    ///         });
    ///         var exampleManagedPrivateEndpoint = new Azure.Synapse.ManagedPrivateEndpoint("exampleManagedPrivateEndpoint", new Azure.Synapse.ManagedPrivateEndpointArgs
    ///         {
    ///             SynapseWorkspaceId = exampleWorkspace.Id,
    ///             TargetResourceId = exampleConnect.Id,
    ///             SubresourceName = "blob",
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 exampleFirewallRule,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Synapse Managed Private Endpoint can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:synapse/managedPrivateEndpoint:ManagedPrivateEndpoint example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/managedVirtualNetworks/default/managedPrivateEndpoints/endpoint1
    /// ```
    /// </summary>
    [AzureResourceType("azure:synapse/managedPrivateEndpoint:ManagedPrivateEndpoint")]
    public partial class ManagedPrivateEndpoint : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the sub resource name which the Synapse Private Endpoint is able to connect to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("subresourceName")]
        public Output<string> SubresourceName { get; private set; } = null!;

        /// <summary>
        /// The ID of the Synapse Workspace on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
        /// </summary>
        [Output("synapseWorkspaceId")]
        public Output<string> SynapseWorkspaceId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Private Link Enabled Remote Resource which this Synapse Private Endpoint should be connected to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("targetResourceId")]
        public Output<string> TargetResourceId { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedPrivateEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedPrivateEndpoint(string name, ManagedPrivateEndpointArgs args, CustomResourceOptions? options = null)
            : base("azure:synapse/managedPrivateEndpoint:ManagedPrivateEndpoint", name, args ?? new ManagedPrivateEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedPrivateEndpoint(string name, Input<string> id, ManagedPrivateEndpointState? state = null, CustomResourceOptions? options = null)
            : base("azure:synapse/managedPrivateEndpoint:ManagedPrivateEndpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ManagedPrivateEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedPrivateEndpoint Get(string name, Input<string> id, ManagedPrivateEndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new ManagedPrivateEndpoint(name, id, state, options);
        }
    }

    public sealed class ManagedPrivateEndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the sub resource name which the Synapse Private Endpoint is able to connect to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subresourceName", required: true)]
        public Input<string> SubresourceName { get; set; } = null!;

        /// <summary>
        /// The ID of the Synapse Workspace on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
        /// </summary>
        [Input("synapseWorkspaceId", required: true)]
        public Input<string> SynapseWorkspaceId { get; set; } = null!;

        /// <summary>
        /// The ID of the Private Link Enabled Remote Resource which this Synapse Private Endpoint should be connected to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("targetResourceId", required: true)]
        public Input<string> TargetResourceId { get; set; } = null!;

        public ManagedPrivateEndpointArgs()
        {
        }
    }

    public sealed class ManagedPrivateEndpointState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the sub resource name which the Synapse Private Endpoint is able to connect to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subresourceName")]
        public Input<string>? SubresourceName { get; set; }

        /// <summary>
        /// The ID of the Synapse Workspace on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
        /// </summary>
        [Input("synapseWorkspaceId")]
        public Input<string>? SynapseWorkspaceId { get; set; }

        /// <summary>
        /// The ID of the Private Link Enabled Remote Resource which this Synapse Private Endpoint should be connected to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("targetResourceId")]
        public Input<string>? TargetResourceId { get; set; }

        public ManagedPrivateEndpointState()
        {
        }
    }
}

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
    /// Manages a Synapse Workspace.
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
    ///             AadAdmin = new Azure.Synapse.Inputs.WorkspaceAadAdminArgs
    ///             {
    ///                 Login = "AzureAD Admin",
    ///                 ObjectId = "00000000-0000-0000-0000-000000000000",
    ///                 TenantId = "00000000-0000-0000-0000-000000000000",
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "Env", "production" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Synapse Workspace can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:synapse/workspace:Workspace example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1
    /// ```
    /// </summary>
    [AzureResourceType("azure:synapse/workspace:Workspace")]
    public partial class Workspace : Pulumi.CustomResource
    {
        /// <summary>
        /// An `aad_admin` block as defined below.
        /// </summary>
        [Output("aadAdmin")]
        public Output<Outputs.WorkspaceAadAdmin> AadAdmin { get; private set; } = null!;

        /// <summary>
        /// A list of Connectivity endpoints for this Synapse Workspace.
        /// </summary>
        [Output("connectivityEndpoints")]
        public Output<ImmutableDictionary<string, string>> ConnectivityEndpoints { get; private set; } = null!;

        /// <summary>
        /// An `identity` block as defined below, which contains the Managed Service Identity information for this Synapse Workspace.
        /// </summary>
        [Output("identities")]
        public Output<ImmutableArray<Outputs.WorkspaceIdentity>> Identities { get; private set; } = null!;

        /// <summary>
        /// Specifies the Azure Region where the synapse Workspace should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Workspace managed resource group.
        /// </summary>
        [Output("managedResourceGroupName")]
        public Output<string> ManagedResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Is Virtual Network enabled for all computes in this workspace? Changing this forces a new resource to be created.
        /// </summary>
        [Output("managedVirtualNetworkEnabled")]
        public Output<bool?> ManagedVirtualNetworkEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the name which should be used for this synapse Workspace. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group where the synapse Workspace should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Specifies The Login Name of the SQL administrator. Changing this forces a new resource to be created.
        /// </summary>
        [Output("sqlAdministratorLogin")]
        public Output<string> SqlAdministratorLogin { get; private set; } = null!;

        /// <summary>
        /// The Password associated with the `sql_administrator_login` for the SQL administrator.
        /// </summary>
        [Output("sqlAdministratorLoginPassword")]
        public Output<string> SqlAdministratorLoginPassword { get; private set; } = null!;

        /// <summary>
        /// Are pipelines (running as workspace's system assigned identity) allowed to access SQL pools?
        /// </summary>
        [Output("sqlIdentityControlEnabled")]
        public Output<bool?> SqlIdentityControlEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the ID of storage data lake gen2 filesystem resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("storageDataLakeGen2FilesystemId")]
        public Output<string> StorageDataLakeGen2FilesystemId { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags which should be assigned to the Synapse Workspace.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Workspace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Workspace(string name, WorkspaceArgs args, CustomResourceOptions? options = null)
            : base("azure:synapse/workspace:Workspace", name, args ?? new WorkspaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Workspace(string name, Input<string> id, WorkspaceState? state = null, CustomResourceOptions? options = null)
            : base("azure:synapse/workspace:Workspace", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Workspace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Workspace Get(string name, Input<string> id, WorkspaceState? state = null, CustomResourceOptions? options = null)
        {
            return new Workspace(name, id, state, options);
        }
    }

    public sealed class WorkspaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `aad_admin` block as defined below.
        /// </summary>
        [Input("aadAdmin")]
        public Input<Inputs.WorkspaceAadAdminArgs>? AadAdmin { get; set; }

        /// <summary>
        /// Specifies the Azure Region where the synapse Workspace should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Workspace managed resource group.
        /// </summary>
        [Input("managedResourceGroupName")]
        public Input<string>? ManagedResourceGroupName { get; set; }

        /// <summary>
        /// Is Virtual Network enabled for all computes in this workspace? Changing this forces a new resource to be created.
        /// </summary>
        [Input("managedVirtualNetworkEnabled")]
        public Input<bool>? ManagedVirtualNetworkEnabled { get; set; }

        /// <summary>
        /// Specifies the name which should be used for this synapse Workspace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group where the synapse Workspace should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies The Login Name of the SQL administrator. Changing this forces a new resource to be created.
        /// </summary>
        [Input("sqlAdministratorLogin", required: true)]
        public Input<string> SqlAdministratorLogin { get; set; } = null!;

        /// <summary>
        /// The Password associated with the `sql_administrator_login` for the SQL administrator.
        /// </summary>
        [Input("sqlAdministratorLoginPassword", required: true)]
        public Input<string> SqlAdministratorLoginPassword { get; set; } = null!;

        /// <summary>
        /// Are pipelines (running as workspace's system assigned identity) allowed to access SQL pools?
        /// </summary>
        [Input("sqlIdentityControlEnabled")]
        public Input<bool>? SqlIdentityControlEnabled { get; set; }

        /// <summary>
        /// Specifies the ID of storage data lake gen2 filesystem resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageDataLakeGen2FilesystemId", required: true)]
        public Input<string> StorageDataLakeGen2FilesystemId { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Synapse Workspace.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public WorkspaceArgs()
        {
        }
    }

    public sealed class WorkspaceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `aad_admin` block as defined below.
        /// </summary>
        [Input("aadAdmin")]
        public Input<Inputs.WorkspaceAadAdminGetArgs>? AadAdmin { get; set; }

        [Input("connectivityEndpoints")]
        private InputMap<string>? _connectivityEndpoints;

        /// <summary>
        /// A list of Connectivity endpoints for this Synapse Workspace.
        /// </summary>
        public InputMap<string> ConnectivityEndpoints
        {
            get => _connectivityEndpoints ?? (_connectivityEndpoints = new InputMap<string>());
            set => _connectivityEndpoints = value;
        }

        [Input("identities")]
        private InputList<Inputs.WorkspaceIdentityGetArgs>? _identities;

        /// <summary>
        /// An `identity` block as defined below, which contains the Managed Service Identity information for this Synapse Workspace.
        /// </summary>
        public InputList<Inputs.WorkspaceIdentityGetArgs> Identities
        {
            get => _identities ?? (_identities = new InputList<Inputs.WorkspaceIdentityGetArgs>());
            set => _identities = value;
        }

        /// <summary>
        /// Specifies the Azure Region where the synapse Workspace should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Workspace managed resource group.
        /// </summary>
        [Input("managedResourceGroupName")]
        public Input<string>? ManagedResourceGroupName { get; set; }

        /// <summary>
        /// Is Virtual Network enabled for all computes in this workspace? Changing this forces a new resource to be created.
        /// </summary>
        [Input("managedVirtualNetworkEnabled")]
        public Input<bool>? ManagedVirtualNetworkEnabled { get; set; }

        /// <summary>
        /// Specifies the name which should be used for this synapse Workspace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group where the synapse Workspace should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Specifies The Login Name of the SQL administrator. Changing this forces a new resource to be created.
        /// </summary>
        [Input("sqlAdministratorLogin")]
        public Input<string>? SqlAdministratorLogin { get; set; }

        /// <summary>
        /// The Password associated with the `sql_administrator_login` for the SQL administrator.
        /// </summary>
        [Input("sqlAdministratorLoginPassword")]
        public Input<string>? SqlAdministratorLoginPassword { get; set; }

        /// <summary>
        /// Are pipelines (running as workspace's system assigned identity) allowed to access SQL pools?
        /// </summary>
        [Input("sqlIdentityControlEnabled")]
        public Input<bool>? SqlIdentityControlEnabled { get; set; }

        /// <summary>
        /// Specifies the ID of storage data lake gen2 filesystem resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageDataLakeGen2FilesystemId")]
        public Input<string>? StorageDataLakeGen2FilesystemId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Synapse Workspace.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public WorkspaceState()
        {
        }
    }
}

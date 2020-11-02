// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Kusto
{
    /// <summary>
    /// Manages a Kusto (also known as Azure Data Explorer) Database Principal Assignment.
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
    ///         var current = Output.Create(Azure.Core.GetClientConfig.InvokeAsync());
    ///         var rg = new Azure.Core.ResourceGroup("rg", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "East US",
    ///         });
    ///         var exampleCluster = new Azure.Kusto.Cluster("exampleCluster", new Azure.Kusto.ClusterArgs
    ///         {
    ///             Location = rg.Location,
    ///             ResourceGroupName = rg.Name,
    ///             Sku = new Azure.Kusto.Inputs.ClusterSkuArgs
    ///             {
    ///                 Name = "Standard_D13_v2",
    ///                 Capacity = 2,
    ///             },
    ///         });
    ///         var exampleDatabase = new Azure.Kusto.Database("exampleDatabase", new Azure.Kusto.DatabaseArgs
    ///         {
    ///             ResourceGroupName = rg.Name,
    ///             Location = rg.Location,
    ///             ClusterName = exampleCluster.Name,
    ///             HotCachePeriod = "P7D",
    ///             SoftDeletePeriod = "P31D",
    ///         });
    ///         var exampleDatabasePrincipalAssignment = new Azure.Kusto.DatabasePrincipalAssignment("exampleDatabasePrincipalAssignment", new Azure.Kusto.DatabasePrincipalAssignmentArgs
    ///         {
    ///             ResourceGroupName = rg.Name,
    ///             ClusterName = exampleCluster.Name,
    ///             DatabaseName = exampleDatabase.Name,
    ///             TenantId = current.Apply(current =&gt; current.TenantId),
    ///             PrincipalId = current.Apply(current =&gt; current.ClientId),
    ///             PrincipalType = "App",
    ///             Role = "Viewer",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Kusto Database Principal Assignment can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class DatabasePrincipalAssignment : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// The name of the database in which to create the resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("databaseName")]
        public Output<string> DatabaseName { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The object id of the principal. Changing this forces a new resource to be created.
        /// </summary>
        [Output("principalId")]
        public Output<string> PrincipalId { get; private set; } = null!;

        /// <summary>
        /// The name of the principal.
        /// </summary>
        [Output("principalName")]
        public Output<string> PrincipalName { get; private set; } = null!;

        /// <summary>
        /// The type of the principal. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("principalType")]
        public Output<string> PrincipalType { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The database role assigned to the principal. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User` and `Viewer`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// The tenant id in which the principal resides. Changing this forces a new resource to be created.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// The name of the tenant.
        /// </summary>
        [Output("tenantName")]
        public Output<string> TenantName { get; private set; } = null!;


        /// <summary>
        /// Create a DatabasePrincipalAssignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatabasePrincipalAssignment(string name, DatabasePrincipalAssignmentArgs args, CustomResourceOptions? options = null)
            : base("azure:kusto/databasePrincipalAssignment:DatabasePrincipalAssignment", name, args ?? new DatabasePrincipalAssignmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatabasePrincipalAssignment(string name, Input<string> id, DatabasePrincipalAssignmentState? state = null, CustomResourceOptions? options = null)
            : base("azure:kusto/databasePrincipalAssignment:DatabasePrincipalAssignment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DatabasePrincipalAssignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatabasePrincipalAssignment Get(string name, Input<string> id, DatabasePrincipalAssignmentState? state = null, CustomResourceOptions? options = null)
        {
            return new DatabasePrincipalAssignment(name, id, state, options);
        }
    }

    public sealed class DatabasePrincipalAssignmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the database in which to create the resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The object id of the principal. Changing this forces a new resource to be created.
        /// </summary>
        [Input("principalId", required: true)]
        public Input<string> PrincipalId { get; set; } = null!;

        /// <summary>
        /// The type of the principal. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("principalType", required: true)]
        public Input<string> PrincipalType { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The database role assigned to the principal. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User` and `Viewer`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// The tenant id in which the principal resides. Changing this forces a new resource to be created.
        /// </summary>
        [Input("tenantId", required: true)]
        public Input<string> TenantId { get; set; } = null!;

        public DatabasePrincipalAssignmentArgs()
        {
        }
    }

    public sealed class DatabasePrincipalAssignmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// The name of the database in which to create the resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The object id of the principal. Changing this forces a new resource to be created.
        /// </summary>
        [Input("principalId")]
        public Input<string>? PrincipalId { get; set; }

        /// <summary>
        /// The name of the principal.
        /// </summary>
        [Input("principalName")]
        public Input<string>? PrincipalName { get; set; }

        /// <summary>
        /// The type of the principal. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("principalType")]
        public Input<string>? PrincipalType { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The database role assigned to the principal. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User` and `Viewer`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The tenant id in which the principal resides. Changing this forces a new resource to be created.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The name of the tenant.
        /// </summary>
        [Input("tenantName")]
        public Input<string>? TenantName { get; set; }

        public DatabasePrincipalAssignmentState()
        {
        }
    }
}

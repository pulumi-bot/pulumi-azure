// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Sql
{
    /// <summary>
    /// Allows you to manage an Azure SQL Elastic Pool.
    /// 
    /// &gt; **NOTE:** -  This version of the `Elasticpool` resource is being **deprecated** and should no longer be used. Please use the azure.mssql.ElasticPool version instead.
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
    ///             Location = "West US",
    ///         });
    ///         var exampleSqlServer = new Azure.Sql.SqlServer("exampleSqlServer", new Azure.Sql.SqlServerArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Version = "12.0",
    ///             AdministratorLogin = "4dm1n157r470r",
    ///             AdministratorLoginPassword = "4-v3ry-53cr37-p455w0rd",
    ///         });
    ///         var exampleElasticPool = new Azure.Sql.ElasticPool("exampleElasticPool", new Azure.Sql.ElasticPoolArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             ServerName = exampleSqlServer.Name,
    ///             Edition = "Basic",
    ///             Dtu = 50,
    ///             DbDtuMin = 0,
    ///             DbDtuMax = 5,
    ///             PoolSize = 5000,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// &gt; **NOTE on `azure.sql.ElasticPool`:** -  The values of `edition`, `dtu`, and `pool_size` must be consistent with the [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus). Any inconsistent argument configuration will be rejected.
    /// </summary>
    public partial class ElasticPool : Pulumi.CustomResource
    {
        /// <summary>
        /// The creation date of the SQL Elastic Pool.
        /// </summary>
        [Output("creationDate")]
        public Output<string> CreationDate { get; private set; } = null!;

        /// <summary>
        /// The maximum DTU which will be guaranteed to all databases in the elastic pool to be created.
        /// </summary>
        [Output("dbDtuMax")]
        public Output<int> DbDtuMax { get; private set; } = null!;

        /// <summary>
        /// The minimum DTU which will be guaranteed to all databases in the elastic pool to be created.
        /// </summary>
        [Output("dbDtuMin")]
        public Output<int> DbDtuMin { get; private set; } = null!;

        /// <summary>
        /// The total shared DTU for the elastic pool. Valid values depend on the `edition` which has been defined. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for valid combinations.
        /// </summary>
        [Output("dtu")]
        public Output<int> Dtu { get; private set; } = null!;

        /// <summary>
        /// The edition of the elastic pool to be created. Valid values are `Basic`, `Standard`, and `Premium`. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for details. Changing this forces a new resource to be created.
        /// </summary>
        [Output("edition")]
        public Output<string> Edition { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The maximum size in MB that all databases in the elastic pool can grow to. The maximum size must be consistent with combination of `edition` and `dtu` and the limits documented in [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus). If not defined when creating an elastic pool, the value is set to the size implied by `edition` and `dtu`.
        /// </summary>
        [Output("poolSize")]
        public Output<int> PoolSize { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
        /// </summary>
        [Output("serverName")]
        public Output<string> ServerName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ElasticPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ElasticPool(string name, ElasticPoolArgs args, CustomResourceOptions? options = null)
            : base("azure:sql/elasticPool:ElasticPool", name, args ?? new ElasticPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ElasticPool(string name, Input<string> id, ElasticPoolState? state = null, CustomResourceOptions? options = null)
            : base("azure:sql/elasticPool:ElasticPool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ElasticPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ElasticPool Get(string name, Input<string> id, ElasticPoolState? state = null, CustomResourceOptions? options = null)
        {
            return new ElasticPool(name, id, state, options);
        }
    }

    public sealed class ElasticPoolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum DTU which will be guaranteed to all databases in the elastic pool to be created.
        /// </summary>
        [Input("dbDtuMax")]
        public Input<int>? DbDtuMax { get; set; }

        /// <summary>
        /// The minimum DTU which will be guaranteed to all databases in the elastic pool to be created.
        /// </summary>
        [Input("dbDtuMin")]
        public Input<int>? DbDtuMin { get; set; }

        /// <summary>
        /// The total shared DTU for the elastic pool. Valid values depend on the `edition` which has been defined. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for valid combinations.
        /// </summary>
        [Input("dtu", required: true)]
        public Input<int> Dtu { get; set; } = null!;

        /// <summary>
        /// The edition of the elastic pool to be created. Valid values are `Basic`, `Standard`, and `Premium`. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for details. Changing this forces a new resource to be created.
        /// </summary>
        [Input("edition", required: true)]
        public Input<string> Edition { get; set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The maximum size in MB that all databases in the elastic pool can grow to. The maximum size must be consistent with combination of `edition` and `dtu` and the limits documented in [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus). If not defined when creating an elastic pool, the value is set to the size implied by `edition` and `dtu`.
        /// </summary>
        [Input("poolSize")]
        public Input<int>? PoolSize { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ElasticPoolArgs()
        {
        }
    }

    public sealed class ElasticPoolState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The creation date of the SQL Elastic Pool.
        /// </summary>
        [Input("creationDate")]
        public Input<string>? CreationDate { get; set; }

        /// <summary>
        /// The maximum DTU which will be guaranteed to all databases in the elastic pool to be created.
        /// </summary>
        [Input("dbDtuMax")]
        public Input<int>? DbDtuMax { get; set; }

        /// <summary>
        /// The minimum DTU which will be guaranteed to all databases in the elastic pool to be created.
        /// </summary>
        [Input("dbDtuMin")]
        public Input<int>? DbDtuMin { get; set; }

        /// <summary>
        /// The total shared DTU for the elastic pool. Valid values depend on the `edition` which has been defined. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for valid combinations.
        /// </summary>
        [Input("dtu")]
        public Input<int>? Dtu { get; set; }

        /// <summary>
        /// The edition of the elastic pool to be created. Valid values are `Basic`, `Standard`, and `Premium`. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for details. Changing this forces a new resource to be created.
        /// </summary>
        [Input("edition")]
        public Input<string>? Edition { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The maximum size in MB that all databases in the elastic pool can grow to. The maximum size must be consistent with combination of `edition` and `dtu` and the limits documented in [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus). If not defined when creating an elastic pool, the value is set to the size implied by `edition` and `dtu`.
        /// </summary>
        [Input("poolSize")]
        public Input<int>? PoolSize { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
        /// </summary>
        [Input("serverName")]
        public Input<string>? ServerName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ElasticPoolState()
        {
        }
    }
}

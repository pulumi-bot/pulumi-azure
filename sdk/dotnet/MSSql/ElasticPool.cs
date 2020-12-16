// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MSSql
{
    /// <summary>
    /// Allows you to manage an Azure SQL Elastic Pool via the `v3.0` API which allows for `vCore` and `DTU` based configurations.
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
    ///             Location = "westeurope",
    ///         });
    ///         var exampleSqlServer = new Azure.Sql.SqlServer("exampleSqlServer", new Azure.Sql.SqlServerArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Version = "12.0",
    ///             AdministratorLogin = "4dm1n157r470r",
    ///             AdministratorLoginPassword = "4-v3ry-53cr37-p455w0rd",
    ///         });
    ///         var exampleElasticPool = new Azure.MSSql.ElasticPool("exampleElasticPool", new Azure.MSSql.ElasticPoolArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             ServerName = exampleSqlServer.Name,
    ///             LicenseType = "LicenseIncluded",
    ///             MaxSizeGb = 756,
    ///             Sku = new Azure.MSSql.Inputs.ElasticPoolSkuArgs
    ///             {
    ///                 Name = "GP_Gen5",
    ///                 Tier = "GeneralPurpose",
    ///                 Family = "Gen5",
    ///                 Capacity = 4,
    ///             },
    ///             PerDatabaseSettings = new Azure.MSSql.Inputs.ElasticPoolPerDatabaseSettingsArgs
    ///             {
    ///                 MinCapacity = 0.25,
    ///                 MaxCapacity = 4,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// SQL Elastic Pool can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:mssql/elasticPool:ElasticPool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/elasticPools/myelasticpoolname
    /// ```
    /// </summary>
    [AzureResourceType("azure:mssql/elasticPool:ElasticPool")]
    public partial class ElasticPool : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the license type applied to this database. Possible values are `LicenseIncluded` and `BasePrice`.
        /// </summary>
        [Output("licenseType")]
        public Output<string> LicenseType { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The max data size of the elastic pool in bytes. Conflicts with `max_size_gb`.
        /// </summary>
        [Output("maxSizeBytes")]
        public Output<int> MaxSizeBytes { get; private set; } = null!;

        /// <summary>
        /// The max data size of the elastic pool in gigabytes. Conflicts with `max_size_bytes`.
        /// </summary>
        [Output("maxSizeGb")]
        public Output<double> MaxSizeGb { get; private set; } = null!;

        /// <summary>
        /// The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A `per_database_settings` block as defined below.
        /// </summary>
        [Output("perDatabaseSettings")]
        public Output<Outputs.ElasticPoolPerDatabaseSettings> PerDatabaseSettings { get; private set; } = null!;

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
        /// A `sku` block as defined below.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.ElasticPoolSku> Sku { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Whether or not this elastic pool is zone redundant. `tier` needs to be `Premium` for `DTU` based  or `BusinessCritical` for `vCore` based `sku`. Defaults to `false`.
        /// </summary>
        [Output("zoneRedundant")]
        public Output<bool?> ZoneRedundant { get; private set; } = null!;


        /// <summary>
        /// Create a ElasticPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ElasticPool(string name, ElasticPoolArgs args, CustomResourceOptions? options = null)
            : base("azure:mssql/elasticPool:ElasticPool", name, args ?? new ElasticPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ElasticPool(string name, Input<string> id, ElasticPoolState? state = null, CustomResourceOptions? options = null)
            : base("azure:mssql/elasticPool:ElasticPool", name, state, MakeResourceOptions(options, id))
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
        /// Specifies the license type applied to this database. Possible values are `LicenseIncluded` and `BasePrice`.
        /// </summary>
        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The max data size of the elastic pool in bytes. Conflicts with `max_size_gb`.
        /// </summary>
        [Input("maxSizeBytes")]
        public Input<int>? MaxSizeBytes { get; set; }

        /// <summary>
        /// The max data size of the elastic pool in gigabytes. Conflicts with `max_size_bytes`.
        /// </summary>
        [Input("maxSizeGb")]
        public Input<double>? MaxSizeGb { get; set; }

        /// <summary>
        /// The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `per_database_settings` block as defined below.
        /// </summary>
        [Input("perDatabaseSettings", required: true)]
        public Input<Inputs.ElasticPoolPerDatabaseSettingsArgs> PerDatabaseSettings { get; set; } = null!;

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

        /// <summary>
        /// A `sku` block as defined below.
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.ElasticPoolSkuArgs> Sku { get; set; } = null!;

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

        /// <summary>
        /// Whether or not this elastic pool is zone redundant. `tier` needs to be `Premium` for `DTU` based  or `BusinessCritical` for `vCore` based `sku`. Defaults to `false`.
        /// </summary>
        [Input("zoneRedundant")]
        public Input<bool>? ZoneRedundant { get; set; }

        public ElasticPoolArgs()
        {
        }
    }

    public sealed class ElasticPoolState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the license type applied to this database. Possible values are `LicenseIncluded` and `BasePrice`.
        /// </summary>
        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The max data size of the elastic pool in bytes. Conflicts with `max_size_gb`.
        /// </summary>
        [Input("maxSizeBytes")]
        public Input<int>? MaxSizeBytes { get; set; }

        /// <summary>
        /// The max data size of the elastic pool in gigabytes. Conflicts with `max_size_bytes`.
        /// </summary>
        [Input("maxSizeGb")]
        public Input<double>? MaxSizeGb { get; set; }

        /// <summary>
        /// The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `per_database_settings` block as defined below.
        /// </summary>
        [Input("perDatabaseSettings")]
        public Input<Inputs.ElasticPoolPerDatabaseSettingsGetArgs>? PerDatabaseSettings { get; set; }

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

        /// <summary>
        /// A `sku` block as defined below.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.ElasticPoolSkuGetArgs>? Sku { get; set; }

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

        /// <summary>
        /// Whether or not this elastic pool is zone redundant. `tier` needs to be `Premium` for `DTU` based  or `BusinessCritical` for `vCore` based `sku`. Defaults to `false`.
        /// </summary>
        [Input("zoneRedundant")]
        public Input<bool>? ZoneRedundant { get; set; }

        public ElasticPoolState()
        {
        }
    }
}

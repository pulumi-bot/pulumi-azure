// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataFactory
{
    /// <summary>
    /// Manages an Azure Data Factory Managed Integration Runtime.
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
    ///             Location = "northeurope",
    ///         });
    ///         var exampleFactory = new Azure.DataFactory.Factory("exampleFactory", new Azure.DataFactory.FactoryArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var exampleIntegrationRuntimeManaged = new Azure.DataFactory.IntegrationRuntimeManaged("exampleIntegrationRuntimeManaged", new Azure.DataFactory.IntegrationRuntimeManagedArgs
    ///         {
    ///             DataFactoryName = exampleFactory.Name,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             NodeSize = "Standard_D8_v3",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Data Factory Integration Managed Runtimes can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:datafactory/integrationRuntimeManaged:IntegrationRuntimeManaged example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/integrationruntimes/example
    /// ```
    /// </summary>
    [AzureResourceType("azure:datafactory/integrationRuntimeManaged:IntegrationRuntimeManaged")]
    public partial class IntegrationRuntimeManaged : Pulumi.CustomResource
    {
        /// <summary>
        /// A `catalog_info` block as defined below.
        /// </summary>
        [Output("catalogInfo")]
        public Output<Outputs.IntegrationRuntimeManagedCatalogInfo?> CatalogInfo { get; private set; } = null!;

        /// <summary>
        /// A `custom_setup_script` block as defined below.
        /// </summary>
        [Output("customSetupScript")]
        public Output<Outputs.IntegrationRuntimeManagedCustomSetupScript?> CustomSetupScript { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Data Factory the Managed Integration Runtime belongs to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("dataFactoryName")]
        public Output<string> DataFactoryName { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Managed Integration Runtime edition. Valid values are `Standard` and `Enterprise`. Defaults to `Standard`.
        /// </summary>
        [Output("edition")]
        public Output<string?> Edition { get; private set; } = null!;

        /// <summary>
        /// The type of the license that is used. Valid values are `LicenseIncluded` and `BasePrize`. Defaults to `LicenseIncluded`.
        /// </summary>
        [Output("licenseType")]
        public Output<string?> LicenseType { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Defines the maximum parallel executions per node. Defaults to `1`. Max is `16`.
        /// </summary>
        [Output("maxParallelExecutionsPerNode")]
        public Output<int?> MaxParallelExecutionsPerNode { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Managed Integration Runtime. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The size of the nodes on which the Managed Integration Runtime runs. Valid values are: `Standard_D2_v3`, `Standard_D4_v3`, `Standard_D8_v3`, `Standard_D16_v3`, `Standard_D32_v3`, `Standard_D64_v3`, `Standard_E2_v3`, `Standard_E4_v3`, `Standard_E8_v3`, `Standard_E16_v3`, `Standard_E32_v3`, `Standard_E64_v3`, `Standard_D1_v2`, `Standard_D2_v2`, `Standard_D3_v2`, `Standard_D4_v2`, `Standard_A4_v2` and `Standard_A8_v2`
        /// </summary>
        [Output("nodeSize")]
        public Output<string> NodeSize { get; private set; } = null!;

        /// <summary>
        /// Number of nodes for the Managed Integration Runtime. Max is `10`. Defaults to `1`.
        /// </summary>
        [Output("numberOfNodes")]
        public Output<int?> NumberOfNodes { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Managed Integration Runtime. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A `vnet_integration` block as defined below.
        /// </summary>
        [Output("vnetIntegration")]
        public Output<Outputs.IntegrationRuntimeManagedVnetIntegration?> VnetIntegration { get; private set; } = null!;


        /// <summary>
        /// Create a IntegrationRuntimeManaged resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IntegrationRuntimeManaged(string name, IntegrationRuntimeManagedArgs args, CustomResourceOptions? options = null)
            : base("azure:datafactory/integrationRuntimeManaged:IntegrationRuntimeManaged", name, args ?? new IntegrationRuntimeManagedArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IntegrationRuntimeManaged(string name, Input<string> id, IntegrationRuntimeManagedState? state = null, CustomResourceOptions? options = null)
            : base("azure:datafactory/integrationRuntimeManaged:IntegrationRuntimeManaged", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IntegrationRuntimeManaged resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IntegrationRuntimeManaged Get(string name, Input<string> id, IntegrationRuntimeManagedState? state = null, CustomResourceOptions? options = null)
        {
            return new IntegrationRuntimeManaged(name, id, state, options);
        }
    }

    public sealed class IntegrationRuntimeManagedArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `catalog_info` block as defined below.
        /// </summary>
        [Input("catalogInfo")]
        public Input<Inputs.IntegrationRuntimeManagedCatalogInfoArgs>? CatalogInfo { get; set; }

        /// <summary>
        /// A `custom_setup_script` block as defined below.
        /// </summary>
        [Input("customSetupScript")]
        public Input<Inputs.IntegrationRuntimeManagedCustomSetupScriptArgs>? CustomSetupScript { get; set; }

        /// <summary>
        /// Specifies the name of the Data Factory the Managed Integration Runtime belongs to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dataFactoryName", required: true)]
        public Input<string> DataFactoryName { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Managed Integration Runtime edition. Valid values are `Standard` and `Enterprise`. Defaults to `Standard`.
        /// </summary>
        [Input("edition")]
        public Input<string>? Edition { get; set; }

        /// <summary>
        /// The type of the license that is used. Valid values are `LicenseIncluded` and `BasePrize`. Defaults to `LicenseIncluded`.
        /// </summary>
        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Defines the maximum parallel executions per node. Defaults to `1`. Max is `16`.
        /// </summary>
        [Input("maxParallelExecutionsPerNode")]
        public Input<int>? MaxParallelExecutionsPerNode { get; set; }

        /// <summary>
        /// Specifies the name of the Managed Integration Runtime. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The size of the nodes on which the Managed Integration Runtime runs. Valid values are: `Standard_D2_v3`, `Standard_D4_v3`, `Standard_D8_v3`, `Standard_D16_v3`, `Standard_D32_v3`, `Standard_D64_v3`, `Standard_E2_v3`, `Standard_E4_v3`, `Standard_E8_v3`, `Standard_E16_v3`, `Standard_E32_v3`, `Standard_E64_v3`, `Standard_D1_v2`, `Standard_D2_v2`, `Standard_D3_v2`, `Standard_D4_v2`, `Standard_A4_v2` and `Standard_A8_v2`
        /// </summary>
        [Input("nodeSize", required: true)]
        public Input<string> NodeSize { get; set; } = null!;

        /// <summary>
        /// Number of nodes for the Managed Integration Runtime. Max is `10`. Defaults to `1`.
        /// </summary>
        [Input("numberOfNodes")]
        public Input<int>? NumberOfNodes { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Managed Integration Runtime. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `vnet_integration` block as defined below.
        /// </summary>
        [Input("vnetIntegration")]
        public Input<Inputs.IntegrationRuntimeManagedVnetIntegrationArgs>? VnetIntegration { get; set; }

        public IntegrationRuntimeManagedArgs()
        {
        }
    }

    public sealed class IntegrationRuntimeManagedState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `catalog_info` block as defined below.
        /// </summary>
        [Input("catalogInfo")]
        public Input<Inputs.IntegrationRuntimeManagedCatalogInfoGetArgs>? CatalogInfo { get; set; }

        /// <summary>
        /// A `custom_setup_script` block as defined below.
        /// </summary>
        [Input("customSetupScript")]
        public Input<Inputs.IntegrationRuntimeManagedCustomSetupScriptGetArgs>? CustomSetupScript { get; set; }

        /// <summary>
        /// Specifies the name of the Data Factory the Managed Integration Runtime belongs to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dataFactoryName")]
        public Input<string>? DataFactoryName { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Managed Integration Runtime edition. Valid values are `Standard` and `Enterprise`. Defaults to `Standard`.
        /// </summary>
        [Input("edition")]
        public Input<string>? Edition { get; set; }

        /// <summary>
        /// The type of the license that is used. Valid values are `LicenseIncluded` and `BasePrize`. Defaults to `LicenseIncluded`.
        /// </summary>
        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Defines the maximum parallel executions per node. Defaults to `1`. Max is `16`.
        /// </summary>
        [Input("maxParallelExecutionsPerNode")]
        public Input<int>? MaxParallelExecutionsPerNode { get; set; }

        /// <summary>
        /// Specifies the name of the Managed Integration Runtime. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The size of the nodes on which the Managed Integration Runtime runs. Valid values are: `Standard_D2_v3`, `Standard_D4_v3`, `Standard_D8_v3`, `Standard_D16_v3`, `Standard_D32_v3`, `Standard_D64_v3`, `Standard_E2_v3`, `Standard_E4_v3`, `Standard_E8_v3`, `Standard_E16_v3`, `Standard_E32_v3`, `Standard_E64_v3`, `Standard_D1_v2`, `Standard_D2_v2`, `Standard_D3_v2`, `Standard_D4_v2`, `Standard_A4_v2` and `Standard_A8_v2`
        /// </summary>
        [Input("nodeSize")]
        public Input<string>? NodeSize { get; set; }

        /// <summary>
        /// Number of nodes for the Managed Integration Runtime. Max is `10`. Defaults to `1`.
        /// </summary>
        [Input("numberOfNodes")]
        public Input<int>? NumberOfNodes { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Managed Integration Runtime. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// A `vnet_integration` block as defined below.
        /// </summary>
        [Input("vnetIntegration")]
        public Input<Inputs.IntegrationRuntimeManagedVnetIntegrationGetArgs>? VnetIntegration { get; set; }

        public IntegrationRuntimeManagedState()
        {
        }
    }
}

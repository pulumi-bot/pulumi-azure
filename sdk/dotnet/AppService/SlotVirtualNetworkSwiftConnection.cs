// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService
{
    /// <summary>
    /// Manages an App Service Slot's Virtual Network Association (this is for the [Regional VNet Integration](https://docs.microsoft.com/en-us/azure/app-service/web-sites-integrate-with-vnet#regional-vnet-integration) which is still in preview).
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
    ///             Location = "uksouth",
    ///         });
    ///         var exampleVirtualNetwork = new Azure.Network.VirtualNetwork("exampleVirtualNetwork", new Azure.Network.VirtualNetworkArgs
    ///         {
    ///             AddressSpaces = 
    ///             {
    ///                 "10.0.0.0/16",
    ///             },
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var exampleSubnet = new Azure.Network.Subnet("exampleSubnet", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualNetworkName = exampleVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.0.1.0/24",
    ///             },
    ///             Delegations = 
    ///             {
    ///                 new Azure.Network.Inputs.SubnetDelegationArgs
    ///                 {
    ///                     Name = "example-delegation",
    ///                     ServiceDelegation = new Azure.Network.Inputs.SubnetDelegationServiceDelegationArgs
    ///                     {
    ///                         Name = "Microsoft.Web/serverFarms",
    ///                         Actions = 
    ///                         {
    ///                             "Microsoft.Network/virtualNetworks/subnets/action",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///         var examplePlan = new Azure.AppService.Plan("examplePlan", new Azure.AppService.PlanArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Sku = new Azure.AppService.Inputs.PlanSkuArgs
    ///             {
    ///                 Tier = "Standard",
    ///                 Size = "S1",
    ///             },
    ///         });
    ///         var exampleAppService = new Azure.AppService.AppService("exampleAppService", new Azure.AppService.AppServiceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AppServicePlanId = examplePlan.Id,
    ///         });
    ///         var example_staging = new Azure.AppService.Slot("example-staging", new Azure.AppService.SlotArgs
    ///         {
    ///             AppServiceName = exampleAppService.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AppServicePlanId = examplePlan.Id,
    ///         });
    ///         var exampleSlotVirtualNetworkSwiftConnection = new Azure.AppService.SlotVirtualNetworkSwiftConnection("exampleSlotVirtualNetworkSwiftConnection", new Azure.AppService.SlotVirtualNetworkSwiftConnectionArgs
    ///         {
    ///             SlotName = example_staging.Name,
    ///             AppServiceId = exampleAppService.Id,
    ///             SubnetId = exampleSubnet.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// App Service Slot Virtual Network Associations can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class SlotVirtualNetworkSwiftConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the App Service or Function App to associate to the VNet. Changing this forces a new resource to be created.
        /// </summary>
        [Output("appServiceId")]
        public Output<string> AppServiceId { get; private set; } = null!;

        /// <summary>
        /// The name of the App Service Slot or Function App Slot. Changing this forces a new resource to be created.
        /// </summary>
        [Output("slotName")]
        public Output<string> SlotName { get; private set; } = null!;

        /// <summary>
        /// The ID of the subnet the app service will be associated to (the subnet must have a `service_delegation` configured for `Microsoft.Web/serverFarms`).
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;


        /// <summary>
        /// Create a SlotVirtualNetworkSwiftConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SlotVirtualNetworkSwiftConnection(string name, SlotVirtualNetworkSwiftConnectionArgs args, CustomResourceOptions? options = null)
            : base("azure:appservice/slotVirtualNetworkSwiftConnection:SlotVirtualNetworkSwiftConnection", name, args ?? new SlotVirtualNetworkSwiftConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SlotVirtualNetworkSwiftConnection(string name, Input<string> id, SlotVirtualNetworkSwiftConnectionState? state = null, CustomResourceOptions? options = null)
            : base("azure:appservice/slotVirtualNetworkSwiftConnection:SlotVirtualNetworkSwiftConnection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SlotVirtualNetworkSwiftConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SlotVirtualNetworkSwiftConnection Get(string name, Input<string> id, SlotVirtualNetworkSwiftConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new SlotVirtualNetworkSwiftConnection(name, id, state, options);
        }
    }

    public sealed class SlotVirtualNetworkSwiftConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the App Service or Function App to associate to the VNet. Changing this forces a new resource to be created.
        /// </summary>
        [Input("appServiceId", required: true)]
        public Input<string> AppServiceId { get; set; } = null!;

        /// <summary>
        /// The name of the App Service Slot or Function App Slot. Changing this forces a new resource to be created.
        /// </summary>
        [Input("slotName", required: true)]
        public Input<string> SlotName { get; set; } = null!;

        /// <summary>
        /// The ID of the subnet the app service will be associated to (the subnet must have a `service_delegation` configured for `Microsoft.Web/serverFarms`).
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public SlotVirtualNetworkSwiftConnectionArgs()
        {
        }
    }

    public sealed class SlotVirtualNetworkSwiftConnectionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the App Service or Function App to associate to the VNet. Changing this forces a new resource to be created.
        /// </summary>
        [Input("appServiceId")]
        public Input<string>? AppServiceId { get; set; }

        /// <summary>
        /// The name of the App Service Slot or Function App Slot. Changing this forces a new resource to be created.
        /// </summary>
        [Input("slotName")]
        public Input<string>? SlotName { get; set; }

        /// <summary>
        /// The ID of the subnet the app service will be associated to (the subnet must have a `service_delegation` configured for `Microsoft.Web/serverFarms`).
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public SlotVirtualNetworkSwiftConnectionState()
        {
        }
    }
}

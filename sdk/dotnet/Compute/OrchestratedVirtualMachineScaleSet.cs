// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute
{
    /// <summary>
    /// Manages an Orchestrated Virtual Machine Scale Set.
    /// 
    /// &gt; **Note:** Orchestrated Virtual Machine Scale Sets are in Public Preview and it may receive breaking changes - [more details can be found in the Azure Documentation](https://docs.microsoft.com/en-us/azure/virtual-machine-scale-sets/orchestration-modes).
    /// 
    /// &gt; **Note:** Azure is planning to deprecate the `single_placement_group` attribute in the Orchestrated Virtual Machine Scale Set starting from api-version `2019-12-01` and there will be a breaking change in the Orchestrated Virtual Machine Scale Set.
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
    ///         var exampleOrchestratedVirtualMachineScaleSet = new Azure.Compute.OrchestratedVirtualMachineScaleSet("exampleOrchestratedVirtualMachineScaleSet", new Azure.Compute.OrchestratedVirtualMachineScaleSetArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             PlatformFaultDomainCount = 1,
    ///             Zones = 
    ///             {
    ///                 "1",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// An Orchestrated Virtual Machine Scale Set can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class OrchestratedVirtualMachineScaleSet : Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        /// </summary>
        [Output("platformFaultDomainCount")]
        public Output<int> PlatformFaultDomainCount { get; private set; } = null!;

        /// <summary>
        /// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("proximityPlacementGroupId")]
        public Output<string?> ProximityPlacementGroupId { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Should the Orchestrated Virtual Machine Scale Set use single placement group? Defaults to `false`.
        /// </summary>
        [Output("singlePlacementGroup")]
        public Output<bool?> SinglePlacementGroup { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The Unique ID for the Orchestrated Virtual Machine Scale Set.
        /// </summary>
        [Output("uniqueId")]
        public Output<string> UniqueId { get; private set; } = null!;

        /// <summary>
        /// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
        /// </summary>
        [Output("zones")]
        public Output<string?> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a OrchestratedVirtualMachineScaleSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrchestratedVirtualMachineScaleSet(string name, OrchestratedVirtualMachineScaleSetArgs args, CustomResourceOptions? options = null)
            : base("azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet", name, args ?? new OrchestratedVirtualMachineScaleSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrchestratedVirtualMachineScaleSet(string name, Input<string> id, OrchestratedVirtualMachineScaleSetState? state = null, CustomResourceOptions? options = null)
            : base("azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OrchestratedVirtualMachineScaleSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrchestratedVirtualMachineScaleSet Get(string name, Input<string> id, OrchestratedVirtualMachineScaleSetState? state = null, CustomResourceOptions? options = null)
        {
            return new OrchestratedVirtualMachineScaleSet(name, id, state, options);
        }
    }

    public sealed class OrchestratedVirtualMachineScaleSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("platformFaultDomainCount", required: true)]
        public Input<int> PlatformFaultDomainCount { get; set; } = null!;

        /// <summary>
        /// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("proximityPlacementGroupId")]
        public Input<string>? ProximityPlacementGroupId { get; set; }

        /// <summary>
        /// The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Should the Orchestrated Virtual Machine Scale Set use single placement group? Defaults to `false`.
        /// </summary>
        [Input("singlePlacementGroup")]
        public Input<bool>? SinglePlacementGroup { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
        /// </summary>
        [Input("zones")]
        public Input<string>? Zones { get; set; }

        public OrchestratedVirtualMachineScaleSetArgs()
        {
        }
    }

    public sealed class OrchestratedVirtualMachineScaleSetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("platformFaultDomainCount")]
        public Input<int>? PlatformFaultDomainCount { get; set; }

        /// <summary>
        /// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("proximityPlacementGroupId")]
        public Input<string>? ProximityPlacementGroupId { get; set; }

        /// <summary>
        /// The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Should the Orchestrated Virtual Machine Scale Set use single placement group? Defaults to `false`.
        /// </summary>
        [Input("singlePlacementGroup")]
        public Input<bool>? SinglePlacementGroup { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The Unique ID for the Orchestrated Virtual Machine Scale Set.
        /// </summary>
        [Input("uniqueId")]
        public Input<string>? UniqueId { get; set; }

        /// <summary>
        /// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
        /// </summary>
        [Input("zones")]
        public Input<string>? Zones { get; set; }

        public OrchestratedVirtualMachineScaleSetState()
        {
        }
    }
}

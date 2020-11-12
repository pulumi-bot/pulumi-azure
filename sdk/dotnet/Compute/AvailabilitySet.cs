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
    /// Manages an Availability Set for Virtual Machines.
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
    ///         var exampleAvailabilitySet = new Azure.Compute.AvailabilitySet("exampleAvailabilitySet", new Azure.Compute.AvailabilitySetArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Tags = 
    ///             {
    ///                 { "environment", "Production" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Availability Sets can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:compute/availabilitySet:AvailabilitySet group1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/availabilitySets/webAvailSet
    /// ```
    /// </summary>
    public partial class AvailabilitySet : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
        /// </summary>
        [Output("managed")]
        public Output<bool?> Managed { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the availability set. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of fault domains that are used. Defaults to `3`.
        /// </summary>
        [Output("platformFaultDomainCount")]
        public Output<int?> PlatformFaultDomainCount { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of update domains that are used. Defaults to `5`.
        /// </summary>
        [Output("platformUpdateDomainCount")]
        public Output<int?> PlatformUpdateDomainCount { get; private set; } = null!;

        /// <summary>
        /// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
        /// </summary>
        [Output("proximityPlacementGroupId")]
        public Output<string?> ProximityPlacementGroupId { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a AvailabilitySet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AvailabilitySet(string name, AvailabilitySetArgs args, CustomResourceOptions? options = null)
            : base("azure:compute/availabilitySet:AvailabilitySet", name, args ?? new AvailabilitySetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AvailabilitySet(string name, Input<string> id, AvailabilitySetState? state = null, CustomResourceOptions? options = null)
            : base("azure:compute/availabilitySet:AvailabilitySet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AvailabilitySet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AvailabilitySet Get(string name, Input<string> id, AvailabilitySetState? state = null, CustomResourceOptions? options = null)
        {
            return new AvailabilitySet(name, id, state, options);
        }
    }

    public sealed class AvailabilitySetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
        /// </summary>
        [Input("managed")]
        public Input<bool>? Managed { get; set; }

        /// <summary>
        /// Specifies the name of the availability set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the number of fault domains that are used. Defaults to `3`.
        /// </summary>
        [Input("platformFaultDomainCount")]
        public Input<int>? PlatformFaultDomainCount { get; set; }

        /// <summary>
        /// Specifies the number of update domains that are used. Defaults to `5`.
        /// </summary>
        [Input("platformUpdateDomainCount")]
        public Input<int>? PlatformUpdateDomainCount { get; set; }

        /// <summary>
        /// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
        /// </summary>
        [Input("proximityPlacementGroupId")]
        public Input<string>? ProximityPlacementGroupId { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

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

        public AvailabilitySetArgs()
        {
        }
    }

    public sealed class AvailabilitySetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
        /// </summary>
        [Input("managed")]
        public Input<bool>? Managed { get; set; }

        /// <summary>
        /// Specifies the name of the availability set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the number of fault domains that are used. Defaults to `3`.
        /// </summary>
        [Input("platformFaultDomainCount")]
        public Input<int>? PlatformFaultDomainCount { get; set; }

        /// <summary>
        /// Specifies the number of update domains that are used. Defaults to `5`.
        /// </summary>
        [Input("platformUpdateDomainCount")]
        public Input<int>? PlatformUpdateDomainCount { get; set; }

        /// <summary>
        /// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
        /// </summary>
        [Input("proximityPlacementGroupId")]
        public Input<string>? ProximityPlacementGroupId { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

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

        public AvailabilitySetState()
        {
        }
    }
}

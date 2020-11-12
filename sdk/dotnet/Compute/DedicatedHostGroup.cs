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
    /// Manage a Dedicated Host Group.
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
    ///         var exampleDedicatedHostGroup = new Azure.Compute.DedicatedHostGroup("exampleDedicatedHostGroup", new Azure.Compute.DedicatedHostGroupArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             PlatformFaultDomainCount = 1,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Dedicated Host Group can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:compute/dedicatedHostGroup:DedicatedHostGroup example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/hostGroups/group1
    /// ```
    /// </summary>
    public partial class DedicatedHostGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure location where the Dedicated Host Group exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Dedicated Host Group. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The number of fault domains that the Dedicated Host Group spans. Changing this forces a new resource to be created.
        /// </summary>
        [Output("platformFaultDomainCount")]
        public Output<int> PlatformFaultDomainCount { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group the Dedicated Host Group is located in. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A list of Availability Zones in which the Dedicated Host Group should be located. Changing this forces a new resource to be created.
        /// </summary>
        [Output("zones")]
        public Output<string?> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a DedicatedHostGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DedicatedHostGroup(string name, DedicatedHostGroupArgs args, CustomResourceOptions? options = null)
            : base("azure:compute/dedicatedHostGroup:DedicatedHostGroup", name, args ?? new DedicatedHostGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DedicatedHostGroup(string name, Input<string> id, DedicatedHostGroupState? state = null, CustomResourceOptions? options = null)
            : base("azure:compute/dedicatedHostGroup:DedicatedHostGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DedicatedHostGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DedicatedHostGroup Get(string name, Input<string> id, DedicatedHostGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new DedicatedHostGroup(name, id, state, options);
        }
    }

    public sealed class DedicatedHostGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Azure location where the Dedicated Host Group exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Dedicated Host Group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The number of fault domains that the Dedicated Host Group spans. Changing this forces a new resource to be created.
        /// </summary>
        [Input("platformFaultDomainCount", required: true)]
        public Input<int> PlatformFaultDomainCount { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group the Dedicated Host Group is located in. Changing this forces a new resource to be created.
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

        /// <summary>
        /// A list of Availability Zones in which the Dedicated Host Group should be located. Changing this forces a new resource to be created.
        /// </summary>
        [Input("zones")]
        public Input<string>? Zones { get; set; }

        public DedicatedHostGroupArgs()
        {
        }
    }

    public sealed class DedicatedHostGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Azure location where the Dedicated Host Group exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Dedicated Host Group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The number of fault domains that the Dedicated Host Group spans. Changing this forces a new resource to be created.
        /// </summary>
        [Input("platformFaultDomainCount")]
        public Input<int>? PlatformFaultDomainCount { get; set; }

        /// <summary>
        /// Specifies the name of the resource group the Dedicated Host Group is located in. Changing this forces a new resource to be created.
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

        /// <summary>
        /// A list of Availability Zones in which the Dedicated Host Group should be located. Changing this forces a new resource to be created.
        /// </summary>
        [Input("zones")]
        public Input<string>? Zones { get; set; }

        public DedicatedHostGroupState()
        {
        }
    }
}

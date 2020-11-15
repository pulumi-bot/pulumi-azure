// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    /// <summary>
    /// Manages an AzureNetwork DDoS Protection Plan.
    /// 
    /// &gt; **NOTE** Azure only allows `one` DDoS Protection Plan per region.
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
    ///         var exampleDdosProtectionPlan = new Azure.Network.DdosProtectionPlan("exampleDdosProtectionPlan", new Azure.Network.DdosProtectionPlanArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Azure DDoS Protection Plan can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:network/ddosProtectionPlan:DdosProtectionPlan example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/ddosProtectionPlans/testddospplan
    /// ```
    /// </summary>
    public partial class DdosProtectionPlan : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Network DDoS Protection Plan. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A list of Virtual Network ID's associated with the DDoS Protection Plan.
        /// </summary>
        [Output("virtualNetworkIds")]
        public Output<ImmutableArray<string>> VirtualNetworkIds { get; private set; } = null!;


        /// <summary>
        /// Create a DdosProtectionPlan resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DdosProtectionPlan(string name, DdosProtectionPlanArgs args, CustomResourceOptions? options = null)
            : base("azure:network/ddosProtectionPlan:DdosProtectionPlan", name, args ?? new DdosProtectionPlanArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DdosProtectionPlan(string name, Input<string> id, DdosProtectionPlanState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/ddosProtectionPlan:DdosProtectionPlan", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DdosProtectionPlan resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DdosProtectionPlan Get(string name, Input<string> id, DdosProtectionPlanState? state = null, CustomResourceOptions? options = null)
        {
            return new DdosProtectionPlan(name, id, state, options);
        }
    }

    public sealed class DdosProtectionPlanArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Network DDoS Protection Plan. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
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

        public DdosProtectionPlanArgs()
        {
        }
    }

    public sealed class DdosProtectionPlanState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Network DDoS Protection Plan. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
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

        [Input("virtualNetworkIds")]
        private InputList<string>? _virtualNetworkIds;

        /// <summary>
        /// A list of Virtual Network ID's associated with the DDoS Protection Plan.
        /// </summary>
        public InputList<string> VirtualNetworkIds
        {
            get => _virtualNetworkIds ?? (_virtualNetworkIds = new InputList<string>());
            set => _virtualNetworkIds = value;
        }

        public DdosProtectionPlanState()
        {
        }
    }
}

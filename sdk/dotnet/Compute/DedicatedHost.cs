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
    /// Manage a Dedicated Host within a Dedicated Host Group.
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
    ///             PlatformFaultDomainCount = 2,
    ///         });
    ///         var exampleDedicatedHost = new Azure.Compute.DedicatedHost("exampleDedicatedHost", new Azure.Compute.DedicatedHostArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             DedicatedHostGroupId = exampleDedicatedHostGroup.Id,
    ///             SkuName = "DSv3-Type1",
    ///             PlatformFaultDomain = 1,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Dedicated Hosts can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class DedicatedHost : Pulumi.CustomResource
    {
        /// <summary>
        /// Should the Dedicated Host automatically be replaced in case of a Hardware Failure? Defaults to `true`.
        /// </summary>
        [Output("autoReplaceOnFailure")]
        public Output<bool?> AutoReplaceOnFailure { get; private set; } = null!;

        /// <summary>
        /// Specifies the ID of the Dedicated Host Group where the Dedicated Host should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("dedicatedHostGroupId")]
        public Output<string> DedicatedHostGroupId { get; private set; } = null!;

        /// <summary>
        /// Specifies the software license type that will be applied to the VMs deployed on the Dedicated Host. Possible values are `None`, `Windows_Server_Hybrid` and `Windows_Server_Perpetual`. Defaults to `None`.
        /// </summary>
        [Output("licenseType")]
        public Output<string?> LicenseType { get; private set; } = null!;

        /// <summary>
        /// Specify the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of this Dedicated Host. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specify the fault domain of the Dedicated Host Group in which to create the Dedicated Host. Changing this forces a new resource to be created.
        /// </summary>
        [Output("platformFaultDomain")]
        public Output<int> PlatformFaultDomain { get; private set; } = null!;

        /// <summary>
        /// Specify the sku name of the Dedicated Host. Possible values are `DSv3-Type1`, `DSv3-Type2`, `DSv4-Type1`, `ESv3-Type1`, `ESv3-Type2`,`FSv2-Type2`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("skuName")]
        public Output<string> SkuName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DedicatedHost resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DedicatedHost(string name, DedicatedHostArgs args, CustomResourceOptions? options = null)
            : base("azure:compute/dedicatedHost:DedicatedHost", name, args ?? new DedicatedHostArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DedicatedHost(string name, Input<string> id, DedicatedHostState? state = null, CustomResourceOptions? options = null)
            : base("azure:compute/dedicatedHost:DedicatedHost", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DedicatedHost resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DedicatedHost Get(string name, Input<string> id, DedicatedHostState? state = null, CustomResourceOptions? options = null)
        {
            return new DedicatedHost(name, id, state, options);
        }
    }

    public sealed class DedicatedHostArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Should the Dedicated Host automatically be replaced in case of a Hardware Failure? Defaults to `true`.
        /// </summary>
        [Input("autoReplaceOnFailure")]
        public Input<bool>? AutoReplaceOnFailure { get; set; }

        /// <summary>
        /// Specifies the ID of the Dedicated Host Group where the Dedicated Host should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dedicatedHostGroupId", required: true)]
        public Input<string> DedicatedHostGroupId { get; set; } = null!;

        /// <summary>
        /// Specifies the software license type that will be applied to the VMs deployed on the Dedicated Host. Possible values are `None`, `Windows_Server_Hybrid` and `Windows_Server_Perpetual`. Defaults to `None`.
        /// </summary>
        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// Specify the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of this Dedicated Host. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specify the fault domain of the Dedicated Host Group in which to create the Dedicated Host. Changing this forces a new resource to be created.
        /// </summary>
        [Input("platformFaultDomain", required: true)]
        public Input<int> PlatformFaultDomain { get; set; } = null!;

        /// <summary>
        /// Specify the sku name of the Dedicated Host. Possible values are `DSv3-Type1`, `DSv3-Type2`, `DSv4-Type1`, `ESv3-Type1`, `ESv3-Type2`,`FSv2-Type2`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("skuName", required: true)]
        public Input<string> SkuName { get; set; } = null!;

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

        public DedicatedHostArgs()
        {
        }
    }

    public sealed class DedicatedHostState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Should the Dedicated Host automatically be replaced in case of a Hardware Failure? Defaults to `true`.
        /// </summary>
        [Input("autoReplaceOnFailure")]
        public Input<bool>? AutoReplaceOnFailure { get; set; }

        /// <summary>
        /// Specifies the ID of the Dedicated Host Group where the Dedicated Host should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dedicatedHostGroupId")]
        public Input<string>? DedicatedHostGroupId { get; set; }

        /// <summary>
        /// Specifies the software license type that will be applied to the VMs deployed on the Dedicated Host. Possible values are `None`, `Windows_Server_Hybrid` and `Windows_Server_Perpetual`. Defaults to `None`.
        /// </summary>
        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// Specify the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of this Dedicated Host. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specify the fault domain of the Dedicated Host Group in which to create the Dedicated Host. Changing this forces a new resource to be created.
        /// </summary>
        [Input("platformFaultDomain")]
        public Input<int>? PlatformFaultDomain { get; set; }

        /// <summary>
        /// Specify the sku name of the Dedicated Host. Possible values are `DSv3-Type1`, `DSv3-Type2`, `DSv4-Type1`, `ESv3-Type1`, `ESv3-Type2`,`FSv2-Type2`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

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

        public DedicatedHostState()
        {
        }
    }
}

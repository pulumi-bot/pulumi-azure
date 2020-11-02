// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Maintenance
{
    /// <summary>
    /// Manages a maintenance assignment to Dedicated Host.
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
    ///         var exampleConfiguration = new Azure.Maintenance.Configuration("exampleConfiguration", new Azure.Maintenance.ConfigurationArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Scope = "All",
    ///         });
    ///         var exampleAssignmentDedicatedHost = new Azure.Maintenance.AssignmentDedicatedHost("exampleAssignmentDedicatedHost", new Azure.Maintenance.AssignmentDedicatedHostArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             MaintenanceConfigurationId = exampleConfiguration.Id,
    ///             DedicatedHostId = exampleDedicatedHost.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Maintenance Assignment can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class AssignmentDedicatedHost : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the Dedicated Host ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
        /// </summary>
        [Output("dedicatedHostId")]
        public Output<string> DedicatedHostId { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("maintenanceConfigurationId")]
        public Output<string> MaintenanceConfigurationId { get; private set; } = null!;


        /// <summary>
        /// Create a AssignmentDedicatedHost resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AssignmentDedicatedHost(string name, AssignmentDedicatedHostArgs args, CustomResourceOptions? options = null)
            : base("azure:maintenance/assignmentDedicatedHost:AssignmentDedicatedHost", name, args ?? new AssignmentDedicatedHostArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AssignmentDedicatedHost(string name, Input<string> id, AssignmentDedicatedHostState? state = null, CustomResourceOptions? options = null)
            : base("azure:maintenance/assignmentDedicatedHost:AssignmentDedicatedHost", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AssignmentDedicatedHost resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AssignmentDedicatedHost Get(string name, Input<string> id, AssignmentDedicatedHostState? state = null, CustomResourceOptions? options = null)
        {
            return new AssignmentDedicatedHost(name, id, state, options);
        }
    }

    public sealed class AssignmentDedicatedHostArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the Dedicated Host ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dedicatedHostId", required: true)]
        public Input<string> DedicatedHostId { get; set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("maintenanceConfigurationId", required: true)]
        public Input<string> MaintenanceConfigurationId { get; set; } = null!;

        public AssignmentDedicatedHostArgs()
        {
        }
    }

    public sealed class AssignmentDedicatedHostState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the Dedicated Host ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dedicatedHostId")]
        public Input<string>? DedicatedHostId { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("maintenanceConfigurationId")]
        public Input<string>? MaintenanceConfigurationId { get; set; }

        public AssignmentDedicatedHostState()
        {
        }
    }
}

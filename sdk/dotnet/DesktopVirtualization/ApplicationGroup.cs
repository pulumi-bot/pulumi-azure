// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DesktopVirtualization
{
    /// <summary>
    /// Manages a Virtual Desktop Application Group.
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
    ///         var example = new Azure.Core.ResourceGroup("example", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "eastus",
    ///         });
    ///         var pooledbreadthfirst = new Azure.DesktopVirtualization.HostPool("pooledbreadthfirst", new Azure.DesktopVirtualization.HostPoolArgs
    ///         {
    ///             Location = example.Location,
    ///             ResourceGroupName = example.Name,
    ///             Type = "Pooled",
    ///             LoadBalancerType = "BreadthFirst",
    ///         });
    ///         var personalautomatic = new Azure.DesktopVirtualization.HostPool("personalautomatic", new Azure.DesktopVirtualization.HostPoolArgs
    ///         {
    ///             Location = example.Location,
    ///             ResourceGroupName = example.Name,
    ///             Type = "Personal",
    ///             PersonalDesktopAssignmentType = "Automatic",
    ///         });
    ///         var remoteapp = new Azure.DesktopVirtualization.ApplicationGroup("remoteapp", new Azure.DesktopVirtualization.ApplicationGroupArgs
    ///         {
    ///             Location = example.Location,
    ///             ResourceGroupName = example.Name,
    ///             Type = "RemoteApp",
    ///             HostPoolId = pooledbreadthfirst.Id,
    ///             FriendlyName = "TestAppGroup",
    ///             Description = "Acceptance Test: An application group",
    ///         });
    ///         var desktopapp = new Azure.DesktopVirtualization.ApplicationGroup("desktopapp", new Azure.DesktopVirtualization.ApplicationGroupArgs
    ///         {
    ///             Location = example.Location,
    ///             ResourceGroupName = example.Name,
    ///             Type = "Desktop",
    ///             HostPoolId = personalautomatic.Id,
    ///             FriendlyName = "TestAppGroup",
    ///             Description = "Acceptance Test: An application group",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Virtual Desktop Application Groups can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class ApplicationGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Option to set a description for the Virtual Desktop Application Group.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Option to set a friendly name for the Virtual Desktop Application Group.
        /// </summary>
        [Output("friendlyName")]
        public Output<string?> FriendlyName { get; private set; } = null!;

        /// <summary>
        /// Resource ID for a Virtual Desktop Host Pool to associate with the
        /// Virtual Desktop Application Group.
        /// </summary>
        [Output("hostPoolId")]
        public Output<string> HostPoolId { get; private set; } = null!;

        /// <summary>
        /// The location/region where the Virtual Desktop Application Group is
        /// located. Changing the location/region forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the Virtual Desktop Application Group. Changing the name forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to
        /// create the Virtual Desktop Application Group. Changing the resource group name forces
        /// a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Type of Virtual Desktop Application Group.
        /// Valid options are `RemoteApp` or `Desktop` application groups.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationGroup(string name, ApplicationGroupArgs args, CustomResourceOptions? options = null)
            : base("azure:desktopvirtualization/applicationGroup:ApplicationGroup", name, args ?? new ApplicationGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationGroup(string name, Input<string> id, ApplicationGroupState? state = null, CustomResourceOptions? options = null)
            : base("azure:desktopvirtualization/applicationGroup:ApplicationGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationGroup Get(string name, Input<string> id, ApplicationGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationGroup(name, id, state, options);
        }
    }

    public sealed class ApplicationGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Option to set a description for the Virtual Desktop Application Group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Option to set a friendly name for the Virtual Desktop Application Group.
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        /// <summary>
        /// Resource ID for a Virtual Desktop Host Pool to associate with the
        /// Virtual Desktop Application Group.
        /// </summary>
        [Input("hostPoolId", required: true)]
        public Input<string> HostPoolId { get; set; } = null!;

        /// <summary>
        /// The location/region where the Virtual Desktop Application Group is
        /// located. Changing the location/region forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Virtual Desktop Application Group. Changing the name forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to
        /// create the Virtual Desktop Application Group. Changing the resource group name forces
        /// a new resource to be created.
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
        /// Type of Virtual Desktop Application Group.
        /// Valid options are `RemoteApp` or `Desktop` application groups.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ApplicationGroupArgs()
        {
        }
    }

    public sealed class ApplicationGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Option to set a description for the Virtual Desktop Application Group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Option to set a friendly name for the Virtual Desktop Application Group.
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        /// <summary>
        /// Resource ID for a Virtual Desktop Host Pool to associate with the
        /// Virtual Desktop Application Group.
        /// </summary>
        [Input("hostPoolId")]
        public Input<string>? HostPoolId { get; set; }

        /// <summary>
        /// The location/region where the Virtual Desktop Application Group is
        /// located. Changing the location/region forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Virtual Desktop Application Group. Changing the name forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to
        /// create the Virtual Desktop Application Group. Changing the resource group name forces
        /// a new resource to be created.
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
        /// Type of Virtual Desktop Application Group.
        /// Valid options are `RemoteApp` or `Desktop` application groups.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ApplicationGroupState()
        {
        }
    }
}

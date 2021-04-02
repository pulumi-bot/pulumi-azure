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
    /// Manages a Virtual Desktop Host Pool.
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
    ///         var exampleHostPool = new Azure.DesktopVirtualization.HostPool("exampleHostPool", new Azure.DesktopVirtualization.HostPoolArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             FriendlyName = "pooleddepthfirst",
    ///             ValidateEnvironment = true,
    ///             CustomRdpProperties = "audiocapturemode:i:1;audiomode:i:0;",
    ///             Description = "Acceptance Test: A pooled host pool - pooleddepthfirst",
    ///             Type = "Pooled",
    ///             MaximumSessionsAllowed = 50,
    ///             LoadBalancerType = "DepthFirst",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Virtual Desktop Host Pools can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:desktopvirtualization/hostPool:HostPool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.DesktopVirtualization/hostpools/myhostpool
    /// ```
    /// </summary>
    [AzureResourceType("azure:desktopvirtualization/hostPool:HostPool")]
    public partial class HostPool : Pulumi.CustomResource
    {
        /// <summary>
        /// A valid custom RDP properties string for the Virtual Desktop Host Pool, available properties can be [found in this article](https://docs.microsoft.com/en-us/windows-server/remote/remote-desktop-services/clients/rdp-files).
        /// </summary>
        [Output("customRdpProperties")]
        public Output<string?> CustomRdpProperties { get; private set; } = null!;

        /// <summary>
        /// A description for the Virtual Desktop Host Pool.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A friendly name for the Virtual Desktop Host Pool.
        /// </summary>
        [Output("friendlyName")]
        public Output<string?> FriendlyName { get; private set; } = null!;

        /// <summary>
        /// `BreadthFirst` load balancing distributes new user sessions across all available session hosts in the host pool.
        /// `DepthFirst` load balancing distributes new user sessions to an available session host with the highest number of connections but has not reached its maximum session limit threshold.
        /// `Persistent` should be used if the host pool type is `Personal`
        /// </summary>
        [Output("loadBalancerType")]
        public Output<string> LoadBalancerType { get; private set; } = null!;

        /// <summary>
        /// The location/region where the Virtual Desktop Host Pool is
        /// located. Changing the location/region forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// A valid integer value from 0 to 999999 for the maximum number of users that have concurrent sessions on a session host.
        /// Should only be set if the `type` of your Virtual Desktop Host Pool is `Pooled`.
        /// </summary>
        [Output("maximumSessionsAllowed")]
        public Output<int?> MaximumSessionsAllowed { get; private set; } = null!;

        /// <summary>
        /// The name of the Virtual Desktop Host Pool. Changing the name
        /// forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// `Automatic` assignment – The service will select an available host and assign it to an user.
        /// `Direct` Assignment – Admin selects a specific host to assign to an user.
        /// </summary>
        [Output("personalDesktopAssignmentType")]
        public Output<string?> PersonalDesktopAssignmentType { get; private set; } = null!;

        /// <summary>
        /// Option to specify the preferred Application Group type for the Virtual Desktop Host Pool.
        /// Valid options are `None`, `Desktop` or `RailApplications`. Default is `None`.
        /// </summary>
        [Output("preferredAppGroupType")]
        public Output<string?> PreferredAppGroupType { get; private set; } = null!;

        /// <summary>
        /// A `registration_info` block which is documented below. Specifies configuration on the registration information of the Virtual Desktop Host Pool.
        /// </summary>
        [Output("registrationInfo")]
        public Output<Outputs.HostPoolRegistrationInfo?> RegistrationInfo { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to
        /// create the Virtual Desktop Host Pool. Changing the resource group name forces
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
        /// The type of the Virtual Desktop Host Pool. Valid options are
        /// `Personal` or `Pooled`. Changing the type forces a new resource to be created.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Allows you to test service changes before they are deployed to production. Defaults to `false`.
        /// </summary>
        [Output("validateEnvironment")]
        public Output<bool?> ValidateEnvironment { get; private set; } = null!;


        /// <summary>
        /// Create a HostPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HostPool(string name, HostPoolArgs args, CustomResourceOptions? options = null)
            : base("azure:desktopvirtualization/hostPool:HostPool", name, args ?? new HostPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HostPool(string name, Input<string> id, HostPoolState? state = null, CustomResourceOptions? options = null)
            : base("azure:desktopvirtualization/hostPool:HostPool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HostPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HostPool Get(string name, Input<string> id, HostPoolState? state = null, CustomResourceOptions? options = null)
        {
            return new HostPool(name, id, state, options);
        }
    }

    public sealed class HostPoolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A valid custom RDP properties string for the Virtual Desktop Host Pool, available properties can be [found in this article](https://docs.microsoft.com/en-us/windows-server/remote/remote-desktop-services/clients/rdp-files).
        /// </summary>
        [Input("customRdpProperties")]
        public Input<string>? CustomRdpProperties { get; set; }

        /// <summary>
        /// A description for the Virtual Desktop Host Pool.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A friendly name for the Virtual Desktop Host Pool.
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        /// <summary>
        /// `BreadthFirst` load balancing distributes new user sessions across all available session hosts in the host pool.
        /// `DepthFirst` load balancing distributes new user sessions to an available session host with the highest number of connections but has not reached its maximum session limit threshold.
        /// `Persistent` should be used if the host pool type is `Personal`
        /// </summary>
        [Input("loadBalancerType", required: true)]
        public Input<string> LoadBalancerType { get; set; } = null!;

        /// <summary>
        /// The location/region where the Virtual Desktop Host Pool is
        /// located. Changing the location/region forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A valid integer value from 0 to 999999 for the maximum number of users that have concurrent sessions on a session host.
        /// Should only be set if the `type` of your Virtual Desktop Host Pool is `Pooled`.
        /// </summary>
        [Input("maximumSessionsAllowed")]
        public Input<int>? MaximumSessionsAllowed { get; set; }

        /// <summary>
        /// The name of the Virtual Desktop Host Pool. Changing the name
        /// forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `Automatic` assignment – The service will select an available host and assign it to an user.
        /// `Direct` Assignment – Admin selects a specific host to assign to an user.
        /// </summary>
        [Input("personalDesktopAssignmentType")]
        public Input<string>? PersonalDesktopAssignmentType { get; set; }

        /// <summary>
        /// Option to specify the preferred Application Group type for the Virtual Desktop Host Pool.
        /// Valid options are `None`, `Desktop` or `RailApplications`. Default is `None`.
        /// </summary>
        [Input("preferredAppGroupType")]
        public Input<string>? PreferredAppGroupType { get; set; }

        /// <summary>
        /// A `registration_info` block which is documented below. Specifies configuration on the registration information of the Virtual Desktop Host Pool.
        /// </summary>
        [Input("registrationInfo")]
        public Input<Inputs.HostPoolRegistrationInfoArgs>? RegistrationInfo { get; set; }

        /// <summary>
        /// The name of the resource group in which to
        /// create the Virtual Desktop Host Pool. Changing the resource group name forces
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
        /// The type of the Virtual Desktop Host Pool. Valid options are
        /// `Personal` or `Pooled`. Changing the type forces a new resource to be created.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Allows you to test service changes before they are deployed to production. Defaults to `false`.
        /// </summary>
        [Input("validateEnvironment")]
        public Input<bool>? ValidateEnvironment { get; set; }

        public HostPoolArgs()
        {
        }
    }

    public sealed class HostPoolState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A valid custom RDP properties string for the Virtual Desktop Host Pool, available properties can be [found in this article](https://docs.microsoft.com/en-us/windows-server/remote/remote-desktop-services/clients/rdp-files).
        /// </summary>
        [Input("customRdpProperties")]
        public Input<string>? CustomRdpProperties { get; set; }

        /// <summary>
        /// A description for the Virtual Desktop Host Pool.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A friendly name for the Virtual Desktop Host Pool.
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        /// <summary>
        /// `BreadthFirst` load balancing distributes new user sessions across all available session hosts in the host pool.
        /// `DepthFirst` load balancing distributes new user sessions to an available session host with the highest number of connections but has not reached its maximum session limit threshold.
        /// `Persistent` should be used if the host pool type is `Personal`
        /// </summary>
        [Input("loadBalancerType")]
        public Input<string>? LoadBalancerType { get; set; }

        /// <summary>
        /// The location/region where the Virtual Desktop Host Pool is
        /// located. Changing the location/region forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A valid integer value from 0 to 999999 for the maximum number of users that have concurrent sessions on a session host.
        /// Should only be set if the `type` of your Virtual Desktop Host Pool is `Pooled`.
        /// </summary>
        [Input("maximumSessionsAllowed")]
        public Input<int>? MaximumSessionsAllowed { get; set; }

        /// <summary>
        /// The name of the Virtual Desktop Host Pool. Changing the name
        /// forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `Automatic` assignment – The service will select an available host and assign it to an user.
        /// `Direct` Assignment – Admin selects a specific host to assign to an user.
        /// </summary>
        [Input("personalDesktopAssignmentType")]
        public Input<string>? PersonalDesktopAssignmentType { get; set; }

        /// <summary>
        /// Option to specify the preferred Application Group type for the Virtual Desktop Host Pool.
        /// Valid options are `None`, `Desktop` or `RailApplications`. Default is `None`.
        /// </summary>
        [Input("preferredAppGroupType")]
        public Input<string>? PreferredAppGroupType { get; set; }

        /// <summary>
        /// A `registration_info` block which is documented below. Specifies configuration on the registration information of the Virtual Desktop Host Pool.
        /// </summary>
        [Input("registrationInfo")]
        public Input<Inputs.HostPoolRegistrationInfoGetArgs>? RegistrationInfo { get; set; }

        /// <summary>
        /// The name of the resource group in which to
        /// create the Virtual Desktop Host Pool. Changing the resource group name forces
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
        /// The type of the Virtual Desktop Host Pool. Valid options are
        /// `Personal` or `Pooled`. Changing the type forces a new resource to be created.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Allows you to test service changes before they are deployed to production. Defaults to `false`.
        /// </summary>
        [Input("validateEnvironment")]
        public Input<bool>? ValidateEnvironment { get; set; }

        public HostPoolState()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Core
{
    /// <summary>
    /// Manages a Resource Group.
    /// 
    /// ## Example Usage
    /// 
    /// 
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
    ///             Location = "West Europe",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class ResourceGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure Region where the Resource Group should exist. Changing this forces a new Resource Group to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The Name which should be used for this Resource Group. Changing this forces a new Resource Group to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags which should be assigned to the Resource Group.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ResourceGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourceGroup(string name, ResourceGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("azure:core/resourceGroup:ResourceGroup", name, args ?? new ResourceGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourceGroup(string name, Input<string> id, ResourceGroupState? state = null, CustomResourceOptions? options = null)
            : base("azure:core/resourceGroup:ResourceGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ResourceGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourceGroup Get(string name, Input<string> id, ResourceGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ResourceGroup(name, id, state, options);
        }
    }

    public sealed class ResourceGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Azure Region where the Resource Group should exist. Changing this forces a new Resource Group to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The Name which should be used for this Resource Group. Changing this forces a new Resource Group to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Resource Group.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ResourceGroupArgs()
        {
        }
    }

    public sealed class ResourceGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Azure Region where the Resource Group should exist. Changing this forces a new Resource Group to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The Name which should be used for this Resource Group. Changing this forces a new Resource Group to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Resource Group.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ResourceGroupState()
        {
        }
    }
}

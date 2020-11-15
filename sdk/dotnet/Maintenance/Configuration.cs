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
    /// Manages a maintenance configuration.
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
    ///         var exampleConfiguration = new Azure.Maintenance.Configuration("exampleConfiguration", new Azure.Maintenance.ConfigurationArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Scope = "All",
    ///             Tags = 
    ///             {
    ///                 { "Env", "prod" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Maintenance Configuration can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:maintenance/configuration:Configuration example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.maintenance/maintenanceconfigurations/example-mc
    /// ```
    /// </summary>
    public partial class Configuration : Pulumi.CustomResource
    {
        /// <summary>
        /// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Maintenance Configuration. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Maintenance Configuration should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The scope of the Maintenance Configuration. Possible values are `All`, `Host`, `Resource` or `InResource`. Default to `All`.
        /// </summary>
        [Output("scope")]
        public Output<string?> Scope { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource. The key could not contain upper case letter.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Configuration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Configuration(string name, ConfigurationArgs args, CustomResourceOptions? options = null)
            : base("azure:maintenance/configuration:Configuration", name, args ?? new ConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Configuration(string name, Input<string> id, ConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("azure:maintenance/configuration:Configuration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Configuration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Configuration Get(string name, Input<string> id, ConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new Configuration(name, id, state, options);
        }
    }

    public sealed class ConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Maintenance Configuration. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Maintenance Configuration should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The scope of the Maintenance Configuration. Possible values are `All`, `Host`, `Resource` or `InResource`. Default to `All`.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource. The key could not contain upper case letter.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ConfigurationArgs()
        {
        }
    }

    public sealed class ConfigurationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Maintenance Configuration. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Maintenance Configuration should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The scope of the Maintenance Configuration. Possible values are `All`, `Host`, `Resource` or `InResource`. Default to `All`.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource. The key could not contain upper case letter.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ConfigurationState()
        {
        }
    }
}

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
    /// Manages an Azure Custom Provider.
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
    ///             Location = "northeurope",
    ///         });
    ///         var exampleCustomProvider = new Azure.Core.CustomProvider("exampleCustomProvider", new Azure.Core.CustomProviderArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             ResourceTypes = 
    ///             {
    ///                 new Azure.Core.Inputs.CustomProviderResourceTypeArgs
    ///                 {
    ///                     Name = "dEf1",
    ///                     Endpoint = "https://testendpoint.com/",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Custom Provider can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class CustomProvider : Pulumi.CustomResource
    {
        /// <summary>
        /// Any number of `action` block as defined below. One of `resource_type` or `action` must be specified.
        /// </summary>
        [Output("actions")]
        public Output<ImmutableArray<Outputs.CustomProviderAction>> Actions { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Custom Provider. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Custom Provider.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Any number of `resource_type` block as defined below. One of `resource_type` or `action` must be specified.
        /// </summary>
        [Output("resourceTypes")]
        public Output<ImmutableArray<Outputs.CustomProviderResourceType>> ResourceTypes { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Any number of `validation` block as defined below.
        /// </summary>
        [Output("validations")]
        public Output<ImmutableArray<Outputs.CustomProviderValidation>> Validations { get; private set; } = null!;


        /// <summary>
        /// Create a CustomProvider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomProvider(string name, CustomProviderArgs args, CustomResourceOptions? options = null)
            : base("azure:core/customProvider:CustomProvider", name, args ?? new CustomProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomProvider(string name, Input<string> id, CustomProviderState? state = null, CustomResourceOptions? options = null)
            : base("azure:core/customProvider:CustomProvider", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CustomProvider resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomProvider Get(string name, Input<string> id, CustomProviderState? state = null, CustomResourceOptions? options = null)
        {
            return new CustomProvider(name, id, state, options);
        }
    }

    public sealed class CustomProviderArgs : Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<Inputs.CustomProviderActionArgs>? _actions;

        /// <summary>
        /// Any number of `action` block as defined below. One of `resource_type` or `action` must be specified.
        /// </summary>
        public InputList<Inputs.CustomProviderActionArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.CustomProviderActionArgs>());
            set => _actions = value;
        }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Custom Provider. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Custom Provider.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("resourceTypes")]
        private InputList<Inputs.CustomProviderResourceTypeArgs>? _resourceTypes;

        /// <summary>
        /// Any number of `resource_type` block as defined below. One of `resource_type` or `action` must be specified.
        /// </summary>
        public InputList<Inputs.CustomProviderResourceTypeArgs> ResourceTypes
        {
            get => _resourceTypes ?? (_resourceTypes = new InputList<Inputs.CustomProviderResourceTypeArgs>());
            set => _resourceTypes = value;
        }

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

        [Input("validations")]
        private InputList<Inputs.CustomProviderValidationArgs>? _validations;

        /// <summary>
        /// Any number of `validation` block as defined below.
        /// </summary>
        public InputList<Inputs.CustomProviderValidationArgs> Validations
        {
            get => _validations ?? (_validations = new InputList<Inputs.CustomProviderValidationArgs>());
            set => _validations = value;
        }

        public CustomProviderArgs()
        {
        }
    }

    public sealed class CustomProviderState : Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<Inputs.CustomProviderActionGetArgs>? _actions;

        /// <summary>
        /// Any number of `action` block as defined below. One of `resource_type` or `action` must be specified.
        /// </summary>
        public InputList<Inputs.CustomProviderActionGetArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.CustomProviderActionGetArgs>());
            set => _actions = value;
        }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Custom Provider. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Custom Provider.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("resourceTypes")]
        private InputList<Inputs.CustomProviderResourceTypeGetArgs>? _resourceTypes;

        /// <summary>
        /// Any number of `resource_type` block as defined below. One of `resource_type` or `action` must be specified.
        /// </summary>
        public InputList<Inputs.CustomProviderResourceTypeGetArgs> ResourceTypes
        {
            get => _resourceTypes ?? (_resourceTypes = new InputList<Inputs.CustomProviderResourceTypeGetArgs>());
            set => _resourceTypes = value;
        }

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

        [Input("validations")]
        private InputList<Inputs.CustomProviderValidationGetArgs>? _validations;

        /// <summary>
        /// Any number of `validation` block as defined below.
        /// </summary>
        public InputList<Inputs.CustomProviderValidationGetArgs> Validations
        {
            get => _validations ?? (_validations = new InputList<Inputs.CustomProviderValidationGetArgs>());
            set => _validations = value;
        }

        public CustomProviderState()
        {
        }
    }
}

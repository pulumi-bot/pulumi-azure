// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ManagedApplication
{
    /// <summary>
    /// Manages a Managed Application Definition.
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
    ///         var current = Output.Create(Azure.Core.GetClientConfig.InvokeAsync());
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var exampleDefinition = new Azure.ManagedApplication.Definition("exampleDefinition", new Azure.ManagedApplication.DefinitionArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             LockLevel = "ReadOnly",
    ///             PackageFileUri = "https://github.com/Azure/azure-managedapp-samples/raw/master/Managed Application Sample Packages/201-managed-storage-account/managedstorage.zip",
    ///             DisplayName = "TestManagedApplicationDefinition",
    ///             Description = "Test Managed Application Definition",
    ///             Authorizations = 
    ///             {
    ///                 new Azure.ManagedApplication.Inputs.DefinitionAuthorizationArgs
    ///                 {
    ///                     ServicePrincipalId = current.Apply(current =&gt; current.ObjectId),
    ///                     RoleDefinitionId = "a094b430-dad3-424d-ae58-13f72fd72591",
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
    /// Managed Application Definition can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class Definition : Pulumi.CustomResource
    {
        /// <summary>
        /// One or more `authorization` block defined below.
        /// </summary>
        [Output("authorizations")]
        public Output<ImmutableArray<Outputs.DefinitionAuthorization>> Authorizations { get; private set; } = null!;

        /// <summary>
        /// Specifies the `createUiDefinition` json for the backing template with `Microsoft.Solutions/applications` resource.
        /// </summary>
        [Output("createUiDefinition")]
        public Output<string?> CreateUiDefinition { get; private set; } = null!;

        /// <summary>
        /// Specifies the managed application definition description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies the managed application definition display name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the managed application lock level. Valid values include `CanNotDelete`, `None`, `ReadOnly`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("lockLevel")]
        public Output<string> LockLevel { get; private set; } = null!;

        /// <summary>
        /// Specifies the inline main template json which has resources to be provisioned.
        /// </summary>
        [Output("mainTemplate")]
        public Output<string?> MainTemplate { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Managed Application Definition. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Is the package enabled? Defaults to `true`.
        /// </summary>
        [Output("packageEnabled")]
        public Output<bool?> PackageEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the managed application definition package file Uri.
        /// </summary>
        [Output("packageFileUri")]
        public Output<string?> PackageFileUri { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Managed Application Definition should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Definition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Definition(string name, DefinitionArgs args, CustomResourceOptions? options = null)
            : base("azure:managedapplication/definition:Definition", name, args ?? new DefinitionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Definition(string name, Input<string> id, DefinitionState? state = null, CustomResourceOptions? options = null)
            : base("azure:managedapplication/definition:Definition", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Definition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Definition Get(string name, Input<string> id, DefinitionState? state = null, CustomResourceOptions? options = null)
        {
            return new Definition(name, id, state, options);
        }
    }

    public sealed class DefinitionArgs : Pulumi.ResourceArgs
    {
        [Input("authorizations")]
        private InputList<Inputs.DefinitionAuthorizationArgs>? _authorizations;

        /// <summary>
        /// One or more `authorization` block defined below.
        /// </summary>
        public InputList<Inputs.DefinitionAuthorizationArgs> Authorizations
        {
            get => _authorizations ?? (_authorizations = new InputList<Inputs.DefinitionAuthorizationArgs>());
            set => _authorizations = value;
        }

        /// <summary>
        /// Specifies the `createUiDefinition` json for the backing template with `Microsoft.Solutions/applications` resource.
        /// </summary>
        [Input("createUiDefinition")]
        public Input<string>? CreateUiDefinition { get; set; }

        /// <summary>
        /// Specifies the managed application definition description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the managed application definition display name.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the managed application lock level. Valid values include `CanNotDelete`, `None`, `ReadOnly`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("lockLevel", required: true)]
        public Input<string> LockLevel { get; set; } = null!;

        /// <summary>
        /// Specifies the inline main template json which has resources to be provisioned.
        /// </summary>
        [Input("mainTemplate")]
        public Input<string>? MainTemplate { get; set; }

        /// <summary>
        /// Specifies the name of the Managed Application Definition. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Is the package enabled? Defaults to `true`.
        /// </summary>
        [Input("packageEnabled")]
        public Input<bool>? PackageEnabled { get; set; }

        /// <summary>
        /// Specifies the managed application definition package file Uri.
        /// </summary>
        [Input("packageFileUri")]
        public Input<string>? PackageFileUri { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Managed Application Definition should exist. Changing this forces a new resource to be created.
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

        public DefinitionArgs()
        {
        }
    }

    public sealed class DefinitionState : Pulumi.ResourceArgs
    {
        [Input("authorizations")]
        private InputList<Inputs.DefinitionAuthorizationGetArgs>? _authorizations;

        /// <summary>
        /// One or more `authorization` block defined below.
        /// </summary>
        public InputList<Inputs.DefinitionAuthorizationGetArgs> Authorizations
        {
            get => _authorizations ?? (_authorizations = new InputList<Inputs.DefinitionAuthorizationGetArgs>());
            set => _authorizations = value;
        }

        /// <summary>
        /// Specifies the `createUiDefinition` json for the backing template with `Microsoft.Solutions/applications` resource.
        /// </summary>
        [Input("createUiDefinition")]
        public Input<string>? CreateUiDefinition { get; set; }

        /// <summary>
        /// Specifies the managed application definition description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the managed application definition display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the managed application lock level. Valid values include `CanNotDelete`, `None`, `ReadOnly`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("lockLevel")]
        public Input<string>? LockLevel { get; set; }

        /// <summary>
        /// Specifies the inline main template json which has resources to be provisioned.
        /// </summary>
        [Input("mainTemplate")]
        public Input<string>? MainTemplate { get; set; }

        /// <summary>
        /// Specifies the name of the Managed Application Definition. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Is the package enabled? Defaults to `true`.
        /// </summary>
        [Input("packageEnabled")]
        public Input<bool>? PackageEnabled { get; set; }

        /// <summary>
        /// Specifies the managed application definition package file Uri.
        /// </summary>
        [Input("packageFileUri")]
        public Input<string>? PackageFileUri { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Managed Application Definition should exist. Changing this forces a new resource to be created.
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

        public DefinitionState()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Role
{
    /// <summary>
    /// Manages a custom Role Definition, used to assign Roles to Users/Principals. See ['Understand role definitions'](https://docs.microsoft.com/en-us/azure/role-based-access-control/role-definitions) in the Azure documentation for more details.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/role_definition_legacy.html.markdown.
    /// </summary>
    public partial class Definition : Pulumi.CustomResource
    {
        /// <summary>
        /// One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
        /// </summary>
        [Output("assignableScopes")]
        public Output<ImmutableArray<string>> AssignableScopes { get; private set; } = null!;

        /// <summary>
        /// A description of the Role Definition.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the Role Definition. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A `permissions` block as defined below.
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableArray<Outputs.DefinitionPermissions>> Permissions { get; private set; } = null!;

        /// <summary>
        /// A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
        /// </summary>
        [Output("roleDefinitionId")]
        public Output<string> RoleDefinitionId { get; private set; } = null!;

        /// <summary>
        /// The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("scope")]
        public Output<string> Scope { get; private set; } = null!;


        /// <summary>
        /// Create a Definition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Definition(string name, DefinitionArgs args, CustomResourceOptions? options = null)
            : base("azure:role/definition:Definition", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Definition(string name, Input<string> id, DefinitionState? state = null, CustomResourceOptions? options = null)
            : base("azure:role/definition:Definition", name, state, MakeResourceOptions(options, id))
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
        [Input("assignableScopes", required: true)]
        private InputList<string>? _assignableScopes;

        /// <summary>
        /// One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
        /// </summary>
        public InputList<string> AssignableScopes
        {
            get => _assignableScopes ?? (_assignableScopes = new InputList<string>());
            set => _assignableScopes = value;
        }

        /// <summary>
        /// A description of the Role Definition.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the Role Definition. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("permissions", required: true)]
        private InputList<Inputs.DefinitionPermissionsArgs>? _permissions;

        /// <summary>
        /// A `permissions` block as defined below.
        /// </summary>
        public InputList<Inputs.DefinitionPermissionsArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.DefinitionPermissionsArgs>());
            set => _permissions = value;
        }

        /// <summary>
        /// A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
        /// </summary>
        [Input("roleDefinitionId")]
        public Input<string>? RoleDefinitionId { get; set; }

        /// <summary>
        /// The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        public DefinitionArgs()
        {
        }
    }

    public sealed class DefinitionState : Pulumi.ResourceArgs
    {
        [Input("assignableScopes")]
        private InputList<string>? _assignableScopes;

        /// <summary>
        /// One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
        /// </summary>
        public InputList<string> AssignableScopes
        {
            get => _assignableScopes ?? (_assignableScopes = new InputList<string>());
            set => _assignableScopes = value;
        }

        /// <summary>
        /// A description of the Role Definition.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the Role Definition. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("permissions")]
        private InputList<Inputs.DefinitionPermissionsGetArgs>? _permissions;

        /// <summary>
        /// A `permissions` block as defined below.
        /// </summary>
        public InputList<Inputs.DefinitionPermissionsGetArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.DefinitionPermissionsGetArgs>());
            set => _permissions = value;
        }

        /// <summary>
        /// A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
        /// </summary>
        [Input("roleDefinitionId")]
        public Input<string>? RoleDefinitionId { get; set; }

        /// <summary>
        /// The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        public DefinitionState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class DefinitionPermissionsArgs : Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<string>? _actions;
        public InputList<string> Actions
        {
            get => _actions ?? (_actions = new InputList<string>());
            set => _actions = value;
        }

        [Input("dataActions")]
        private InputList<string>? _dataActions;
        public InputList<string> DataActions
        {
            get => _dataActions ?? (_dataActions = new InputList<string>());
            set => _dataActions = value;
        }

        [Input("notActions")]
        private InputList<string>? _notActions;
        public InputList<string> NotActions
        {
            get => _notActions ?? (_notActions = new InputList<string>());
            set => _notActions = value;
        }

        [Input("notDataActions")]
        private InputList<string>? _notDataActions;
        public InputList<string> NotDataActions
        {
            get => _notDataActions ?? (_notDataActions = new InputList<string>());
            set => _notDataActions = value;
        }

        public DefinitionPermissionsArgs()
        {
        }
    }

    public sealed class DefinitionPermissionsGetArgs : Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<string>? _actions;
        public InputList<string> Actions
        {
            get => _actions ?? (_actions = new InputList<string>());
            set => _actions = value;
        }

        [Input("dataActions")]
        private InputList<string>? _dataActions;
        public InputList<string> DataActions
        {
            get => _dataActions ?? (_dataActions = new InputList<string>());
            set => _dataActions = value;
        }

        [Input("notActions")]
        private InputList<string>? _notActions;
        public InputList<string> NotActions
        {
            get => _notActions ?? (_notActions = new InputList<string>());
            set => _notActions = value;
        }

        [Input("notDataActions")]
        private InputList<string>? _notDataActions;
        public InputList<string> NotDataActions
        {
            get => _notDataActions ?? (_notDataActions = new InputList<string>());
            set => _notDataActions = value;
        }

        public DefinitionPermissionsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class DefinitionPermissions
    {
        public readonly ImmutableArray<string> Actions;
        public readonly ImmutableArray<string> DataActions;
        public readonly ImmutableArray<string> NotActions;
        public readonly ImmutableArray<string> NotDataActions;

        [OutputConstructor]
        private DefinitionPermissions(
            ImmutableArray<string> actions,
            ImmutableArray<string> dataActions,
            ImmutableArray<string> notActions,
            ImmutableArray<string> notDataActions)
        {
            Actions = actions;
            DataActions = dataActions;
            NotActions = notActions;
            NotDataActions = notDataActions;
        }
    }
    }
}
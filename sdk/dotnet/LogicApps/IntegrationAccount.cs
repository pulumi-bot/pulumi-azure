// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.LogicApps
{
    /// <summary>
    /// Manages a Logic App Integration Account.
    /// </summary>
    public partial class IntegrationAccount : Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure Region where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this Logic App Integration Account. Changing this forces a new Logic App Integration Account to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The sku name of the Logic App Integration Account. Possible Values are `Basic`, `Free` and `Standard`.
        /// </summary>
        [Output("skuName")]
        public Output<string> SkuName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags which should be assigned to the Logic App Integration Account.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a IntegrationAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IntegrationAccount(string name, IntegrationAccountArgs args, CustomResourceOptions? options = null)
            : base("azure:logicapps/integrationAccount:IntegrationAccount", name, args ?? new IntegrationAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IntegrationAccount(string name, Input<string> id, IntegrationAccountState? state = null, CustomResourceOptions? options = null)
            : base("azure:logicapps/integrationAccount:IntegrationAccount", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IntegrationAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IntegrationAccount Get(string name, Input<string> id, IntegrationAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new IntegrationAccount(name, id, state, options);
        }
    }

    public sealed class IntegrationAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Azure Region where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name which should be used for this Logic App Integration Account. Changing this forces a new Logic App Integration Account to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The sku name of the Logic App Integration Account. Possible Values are `Basic`, `Free` and `Standard`.
        /// </summary>
        [Input("skuName", required: true)]
        public Input<string> SkuName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Logic App Integration Account.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public IntegrationAccountArgs()
        {
        }
    }

    public sealed class IntegrationAccountState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Azure Region where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name which should be used for this Logic App Integration Account. Changing this forces a new Logic App Integration Account to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Logic App Integration Account should exist. Changing this forces a new Logic App Integration Account to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The sku name of the Logic App Integration Account. Possible Values are `Basic`, `Free` and `Standard`.
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Logic App Integration Account.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public IntegrationAccountState()
        {
        }
    }
}

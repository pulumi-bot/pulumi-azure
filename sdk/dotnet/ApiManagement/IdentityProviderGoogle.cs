// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement
{
    /// <summary>
    /// Manages an API Management Google Identity Provider.
    /// </summary>
    public partial class IdentityProviderGoogle : Pulumi.CustomResource
    {
        /// <summary>
        /// The Name of the API Management Service where this Google Identity Provider should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("apiManagementName")]
        public Output<string> ApiManagementName { get; private set; } = null!;

        /// <summary>
        /// Client Id for Google Sign-in.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// Client secret for Google Sign-in.
        /// </summary>
        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a IdentityProviderGoogle resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdentityProviderGoogle(string name, IdentityProviderGoogleArgs args, CustomResourceOptions? options = null)
            : base("azure:apimanagement/identityProviderGoogle:IdentityProviderGoogle", name, args ?? new IdentityProviderGoogleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IdentityProviderGoogle(string name, Input<string> id, IdentityProviderGoogleState? state = null, CustomResourceOptions? options = null)
            : base("azure:apimanagement/identityProviderGoogle:IdentityProviderGoogle", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IdentityProviderGoogle resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdentityProviderGoogle Get(string name, Input<string> id, IdentityProviderGoogleState? state = null, CustomResourceOptions? options = null)
        {
            return new IdentityProviderGoogle(name, id, state, options);
        }
    }

    public sealed class IdentityProviderGoogleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Name of the API Management Service where this Google Identity Provider should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiManagementName", required: true)]
        public Input<string> ApiManagementName { get; set; } = null!;

        /// <summary>
        /// Client Id for Google Sign-in.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// Client secret for Google Sign-in.
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public IdentityProviderGoogleArgs()
        {
        }
    }

    public sealed class IdentityProviderGoogleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Name of the API Management Service where this Google Identity Provider should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiManagementName")]
        public Input<string>? ApiManagementName { get; set; }

        /// <summary>
        /// Client Id for Google Sign-in.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// Client secret for Google Sign-in.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        public IdentityProviderGoogleState()
        {
        }
    }
}

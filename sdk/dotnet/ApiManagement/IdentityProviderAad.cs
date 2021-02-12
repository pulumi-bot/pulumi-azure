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
    /// Manages an API Management AAD Identity Provider.
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
    ///         var exampleService = new Azure.ApiManagement.Service("exampleService", new Azure.ApiManagement.ServiceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             PublisherName = "My Company",
    ///             PublisherEmail = "company@mycompany.io",
    ///             SkuName = "Developer_1",
    ///         });
    ///         var exampleIdentityProviderAad = new Azure.ApiManagement.IdentityProviderAad("exampleIdentityProviderAad", new Azure.ApiManagement.IdentityProviderAadArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             ApiManagementName = exampleService.Name,
    ///             ClientId = "00000000-0000-0000-0000-000000000000",
    ///             ClientSecret = "00000000000000000000000000000000",
    ///             AllowedTenants = 
    ///             {
    ///                 "00000000-0000-0000-0000-000000000000",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// API Management AAD Identity Provider can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:apimanagement/identityProviderAad:IdentityProviderAad example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/aad
    /// ```
    /// </summary>
    [AzureResourceType("azure:apimanagement/identityProviderAad:IdentityProviderAad")]
    public partial class IdentityProviderAad : Pulumi.CustomResource
    {
        /// <summary>
        /// List of allowed AAD Tenants.
        /// </summary>
        [Output("allowedTenants")]
        public Output<ImmutableArray<string>> AllowedTenants { get; private set; } = null!;

        /// <summary>
        /// The Name of the API Management Service where this AAD Identity Provider should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("apiManagementName")]
        public Output<string> ApiManagementName { get; private set; } = null!;

        /// <summary>
        /// Client Id of the Application in the AAD Identity Provider.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// Client secret of the Application in the AAD Identity Provider.
        /// </summary>
        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The AAD Tenant to use instead of Common when logging into Active Directory
        /// </summary>
        [Output("signinTenant")]
        public Output<string?> SigninTenant { get; private set; } = null!;


        /// <summary>
        /// Create a IdentityProviderAad resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdentityProviderAad(string name, IdentityProviderAadArgs args, CustomResourceOptions? options = null)
            : base("azure:apimanagement/identityProviderAad:IdentityProviderAad", name, args ?? new IdentityProviderAadArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IdentityProviderAad(string name, Input<string> id, IdentityProviderAadState? state = null, CustomResourceOptions? options = null)
            : base("azure:apimanagement/identityProviderAad:IdentityProviderAad", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IdentityProviderAad resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdentityProviderAad Get(string name, Input<string> id, IdentityProviderAadState? state = null, CustomResourceOptions? options = null)
        {
            return new IdentityProviderAad(name, id, state, options);
        }
    }

    public sealed class IdentityProviderAadArgs : Pulumi.ResourceArgs
    {
        [Input("allowedTenants", required: true)]
        private InputList<string>? _allowedTenants;

        /// <summary>
        /// List of allowed AAD Tenants.
        /// </summary>
        public InputList<string> AllowedTenants
        {
            get => _allowedTenants ?? (_allowedTenants = new InputList<string>());
            set => _allowedTenants = value;
        }

        /// <summary>
        /// The Name of the API Management Service where this AAD Identity Provider should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiManagementName", required: true)]
        public Input<string> ApiManagementName { get; set; } = null!;

        /// <summary>
        /// Client Id of the Application in the AAD Identity Provider.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// Client secret of the Application in the AAD Identity Provider.
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The AAD Tenant to use instead of Common when logging into Active Directory
        /// </summary>
        [Input("signinTenant")]
        public Input<string>? SigninTenant { get; set; }

        public IdentityProviderAadArgs()
        {
        }
    }

    public sealed class IdentityProviderAadState : Pulumi.ResourceArgs
    {
        [Input("allowedTenants")]
        private InputList<string>? _allowedTenants;

        /// <summary>
        /// List of allowed AAD Tenants.
        /// </summary>
        public InputList<string> AllowedTenants
        {
            get => _allowedTenants ?? (_allowedTenants = new InputList<string>());
            set => _allowedTenants = value;
        }

        /// <summary>
        /// The Name of the API Management Service where this AAD Identity Provider should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiManagementName")]
        public Input<string>? ApiManagementName { get; set; }

        /// <summary>
        /// Client Id of the Application in the AAD Identity Provider.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// Client secret of the Application in the AAD Identity Provider.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The AAD Tenant to use instead of Common when logging into Active Directory
        /// </summary>
        [Input("signinTenant")]
        public Input<string>? SigninTenant { get; set; }

        public IdentityProviderAadState()
        {
        }
    }
}

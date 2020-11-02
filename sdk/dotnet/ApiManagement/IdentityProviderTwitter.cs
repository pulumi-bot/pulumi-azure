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
    /// Manages an API Management Twitter Identity Provider.
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
    ///         var exampleIdentityProviderTwitter = new Azure.ApiManagement.IdentityProviderTwitter("exampleIdentityProviderTwitter", new Azure.ApiManagement.IdentityProviderTwitterArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             ApiManagementName = exampleService.Name,
    ///             ApiKey = "00000000000000000000000000000000",
    ///             ApiSecretKey = "00000000000000000000000000000000",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// API Management Twitter Identity Provider can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class IdentityProviderTwitter : Pulumi.CustomResource
    {
        /// <summary>
        /// App Consumer API key for Twitter.
        /// </summary>
        [Output("apiKey")]
        public Output<string> ApiKey { get; private set; } = null!;

        /// <summary>
        /// The Name of the API Management Service where this Twitter Identity Provider should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("apiManagementName")]
        public Output<string> ApiManagementName { get; private set; } = null!;

        /// <summary>
        /// App Consumer API secret key for Twitter.
        /// </summary>
        [Output("apiSecretKey")]
        public Output<string> ApiSecretKey { get; private set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a IdentityProviderTwitter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdentityProviderTwitter(string name, IdentityProviderTwitterArgs args, CustomResourceOptions? options = null)
            : base("azure:apimanagement/identityProviderTwitter:IdentityProviderTwitter", name, args ?? new IdentityProviderTwitterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IdentityProviderTwitter(string name, Input<string> id, IdentityProviderTwitterState? state = null, CustomResourceOptions? options = null)
            : base("azure:apimanagement/identityProviderTwitter:IdentityProviderTwitter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IdentityProviderTwitter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdentityProviderTwitter Get(string name, Input<string> id, IdentityProviderTwitterState? state = null, CustomResourceOptions? options = null)
        {
            return new IdentityProviderTwitter(name, id, state, options);
        }
    }

    public sealed class IdentityProviderTwitterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// App Consumer API key for Twitter.
        /// </summary>
        [Input("apiKey", required: true)]
        public Input<string> ApiKey { get; set; } = null!;

        /// <summary>
        /// The Name of the API Management Service where this Twitter Identity Provider should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiManagementName", required: true)]
        public Input<string> ApiManagementName { get; set; } = null!;

        /// <summary>
        /// App Consumer API secret key for Twitter.
        /// </summary>
        [Input("apiSecretKey", required: true)]
        public Input<string> ApiSecretKey { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public IdentityProviderTwitterArgs()
        {
        }
    }

    public sealed class IdentityProviderTwitterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// App Consumer API key for Twitter.
        /// </summary>
        [Input("apiKey")]
        public Input<string>? ApiKey { get; set; }

        /// <summary>
        /// The Name of the API Management Service where this Twitter Identity Provider should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiManagementName")]
        public Input<string>? ApiManagementName { get; set; }

        /// <summary>
        /// App Consumer API secret key for Twitter.
        /// </summary>
        [Input("apiSecretKey")]
        public Input<string>? ApiSecretKey { get; set; }

        /// <summary>
        /// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        public IdentityProviderTwitterState()
        {
        }
    }
}

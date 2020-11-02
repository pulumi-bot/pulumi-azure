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
    /// Manages an API within an API Management Service.
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
    ///             PublisherEmail = "company@exmaple.com",
    ///             SkuName = "Developer_1",
    ///         });
    ///         var exampleApi = new Azure.ApiManagement.Api("exampleApi", new Azure.ApiManagement.ApiArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             ApiManagementName = exampleService.Name,
    ///             Revision = "1",
    ///             DisplayName = "Example API",
    ///             Path = "example",
    ///             Protocols = 
    ///             {
    ///                 "https",
    ///             },
    ///             Import = new Azure.ApiManagement.Inputs.ApiImportArgs
    ///             {
    ///                 ContentFormat = "swagger-link-json",
    ///                 ContentValue = "http://conferenceapi.azurewebsites.net/?format=json",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// API Management API's can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class Api : Pulumi.CustomResource
    {
        /// <summary>
        /// The Name of the API Management Service where this API should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("apiManagementName")]
        public Output<string> ApiManagementName { get; private set; } = null!;

        /// <summary>
        /// A description of the API Management API, which may include HTML formatting tags.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The display name of the API.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// A `import` block as documented below.
        /// </summary>
        [Output("import")]
        public Output<Outputs.ApiImport?> Import { get; private set; } = null!;

        /// <summary>
        /// Is this the current API Revision?
        /// </summary>
        [Output("isCurrent")]
        public Output<bool> IsCurrent { get; private set; } = null!;

        /// <summary>
        /// Is this API Revision online/accessible via the Gateway?
        /// </summary>
        [Output("isOnline")]
        public Output<bool> IsOnline { get; private set; } = null!;

        /// <summary>
        /// The name of the API Management API. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// An `oauth2_authorization` block as documented below.
        /// </summary>
        [Output("oauth2Authorization")]
        public Output<Outputs.ApiOauth2Authorization?> Oauth2Authorization { get; private set; } = null!;

        /// <summary>
        /// An `openid_authentication` block as documented below.
        /// </summary>
        [Output("openidAuthentication")]
        public Output<Outputs.ApiOpenidAuthentication?> OpenidAuthentication { get; private set; } = null!;

        /// <summary>
        /// The Path for this API Management API, which is a relative URL which uniquely identifies this API and all of its resource paths within the API Management Service.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// A list of protocols the operations in this API can be invoked. Possible values are `http` and `https`.
        /// </summary>
        [Output("protocols")]
        public Output<ImmutableArray<string>> Protocols { get; private set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the API Management API exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The Revision which used for this API.
        /// </summary>
        [Output("revision")]
        public Output<string> Revision { get; private set; } = null!;

        /// <summary>
        /// Absolute URL of the backend service implementing this API.
        /// </summary>
        [Output("serviceUrl")]
        public Output<string> ServiceUrl { get; private set; } = null!;

        /// <summary>
        /// Should this API expose a SOAP frontend, rather than a HTTP frontend? Defaults to `false`.
        /// </summary>
        [Output("soapPassThrough")]
        public Output<bool?> SoapPassThrough { get; private set; } = null!;

        /// <summary>
        /// A `subscription_key_parameter_names` block as documented below.
        /// </summary>
        [Output("subscriptionKeyParameterNames")]
        public Output<Outputs.ApiSubscriptionKeyParameterNames> SubscriptionKeyParameterNames { get; private set; } = null!;

        /// <summary>
        /// Should this API require a subscription key?
        /// </summary>
        [Output("subscriptionRequired")]
        public Output<bool?> SubscriptionRequired { get; private set; } = null!;

        /// <summary>
        /// The Version number of this API, if this API is versioned.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// The ID of the Version Set which this API is associated with.
        /// </summary>
        [Output("versionSetId")]
        public Output<string> VersionSetId { get; private set; } = null!;


        /// <summary>
        /// Create a Api resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Api(string name, ApiArgs args, CustomResourceOptions? options = null)
            : base("azure:apimanagement/api:Api", name, args ?? new ApiArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Api(string name, Input<string> id, ApiState? state = null, CustomResourceOptions? options = null)
            : base("azure:apimanagement/api:Api", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Api resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Api Get(string name, Input<string> id, ApiState? state = null, CustomResourceOptions? options = null)
        {
            return new Api(name, id, state, options);
        }
    }

    public sealed class ApiArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Name of the API Management Service where this API should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiManagementName", required: true)]
        public Input<string> ApiManagementName { get; set; } = null!;

        /// <summary>
        /// A description of the API Management API, which may include HTML formatting tags.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the API.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// A `import` block as documented below.
        /// </summary>
        [Input("import")]
        public Input<Inputs.ApiImportArgs>? Import { get; set; }

        /// <summary>
        /// The name of the API Management API. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// An `oauth2_authorization` block as documented below.
        /// </summary>
        [Input("oauth2Authorization")]
        public Input<Inputs.ApiOauth2AuthorizationArgs>? Oauth2Authorization { get; set; }

        /// <summary>
        /// An `openid_authentication` block as documented below.
        /// </summary>
        [Input("openidAuthentication")]
        public Input<Inputs.ApiOpenidAuthenticationArgs>? OpenidAuthentication { get; set; }

        /// <summary>
        /// The Path for this API Management API, which is a relative URL which uniquely identifies this API and all of its resource paths within the API Management Service.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        [Input("protocols", required: true)]
        private InputList<string>? _protocols;

        /// <summary>
        /// A list of protocols the operations in this API can be invoked. Possible values are `http` and `https`.
        /// </summary>
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
            set => _protocols = value;
        }

        /// <summary>
        /// The Name of the Resource Group where the API Management API exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Revision which used for this API.
        /// </summary>
        [Input("revision", required: true)]
        public Input<string> Revision { get; set; } = null!;

        /// <summary>
        /// Absolute URL of the backend service implementing this API.
        /// </summary>
        [Input("serviceUrl")]
        public Input<string>? ServiceUrl { get; set; }

        /// <summary>
        /// Should this API expose a SOAP frontend, rather than a HTTP frontend? Defaults to `false`.
        /// </summary>
        [Input("soapPassThrough")]
        public Input<bool>? SoapPassThrough { get; set; }

        /// <summary>
        /// A `subscription_key_parameter_names` block as documented below.
        /// </summary>
        [Input("subscriptionKeyParameterNames")]
        public Input<Inputs.ApiSubscriptionKeyParameterNamesArgs>? SubscriptionKeyParameterNames { get; set; }

        /// <summary>
        /// Should this API require a subscription key?
        /// </summary>
        [Input("subscriptionRequired")]
        public Input<bool>? SubscriptionRequired { get; set; }

        /// <summary>
        /// The Version number of this API, if this API is versioned.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The ID of the Version Set which this API is associated with.
        /// </summary>
        [Input("versionSetId")]
        public Input<string>? VersionSetId { get; set; }

        public ApiArgs()
        {
        }
    }

    public sealed class ApiState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Name of the API Management Service where this API should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiManagementName")]
        public Input<string>? ApiManagementName { get; set; }

        /// <summary>
        /// A description of the API Management API, which may include HTML formatting tags.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the API.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// A `import` block as documented below.
        /// </summary>
        [Input("import")]
        public Input<Inputs.ApiImportGetArgs>? Import { get; set; }

        /// <summary>
        /// Is this the current API Revision?
        /// </summary>
        [Input("isCurrent")]
        public Input<bool>? IsCurrent { get; set; }

        /// <summary>
        /// Is this API Revision online/accessible via the Gateway?
        /// </summary>
        [Input("isOnline")]
        public Input<bool>? IsOnline { get; set; }

        /// <summary>
        /// The name of the API Management API. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// An `oauth2_authorization` block as documented below.
        /// </summary>
        [Input("oauth2Authorization")]
        public Input<Inputs.ApiOauth2AuthorizationGetArgs>? Oauth2Authorization { get; set; }

        /// <summary>
        /// An `openid_authentication` block as documented below.
        /// </summary>
        [Input("openidAuthentication")]
        public Input<Inputs.ApiOpenidAuthenticationGetArgs>? OpenidAuthentication { get; set; }

        /// <summary>
        /// The Path for this API Management API, which is a relative URL which uniquely identifies this API and all of its resource paths within the API Management Service.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("protocols")]
        private InputList<string>? _protocols;

        /// <summary>
        /// A list of protocols the operations in this API can be invoked. Possible values are `http` and `https`.
        /// </summary>
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
            set => _protocols = value;
        }

        /// <summary>
        /// The Name of the Resource Group where the API Management API exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The Revision which used for this API.
        /// </summary>
        [Input("revision")]
        public Input<string>? Revision { get; set; }

        /// <summary>
        /// Absolute URL of the backend service implementing this API.
        /// </summary>
        [Input("serviceUrl")]
        public Input<string>? ServiceUrl { get; set; }

        /// <summary>
        /// Should this API expose a SOAP frontend, rather than a HTTP frontend? Defaults to `false`.
        /// </summary>
        [Input("soapPassThrough")]
        public Input<bool>? SoapPassThrough { get; set; }

        /// <summary>
        /// A `subscription_key_parameter_names` block as documented below.
        /// </summary>
        [Input("subscriptionKeyParameterNames")]
        public Input<Inputs.ApiSubscriptionKeyParameterNamesGetArgs>? SubscriptionKeyParameterNames { get; set; }

        /// <summary>
        /// Should this API require a subscription key?
        /// </summary>
        [Input("subscriptionRequired")]
        public Input<bool>? SubscriptionRequired { get; set; }

        /// <summary>
        /// The Version number of this API, if this API is versioned.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The ID of the Version Set which this API is associated with.
        /// </summary>
        [Input("versionSetId")]
        public Input<string>? VersionSetId { get; set; }

        public ApiState()
        {
        }
    }
}

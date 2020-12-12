// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.FrontDoor
{
    /// <summary>
    /// Manages the Custom Https Configuration for an Azure Front Door Frontend Endpoint..
    /// 
    /// &gt; **NOTE:** Custom https configurations for a Front Door Frontend Endpoint can be defined both within the `azure.frontdoor.Frontdoor` resource via the `custom_https_configuration` block and by using a separate resource, as described in the following sections.
    /// 
    /// &gt; **NOTE:** Defining custom https configurations using a separate `azure.frontdoor.CustomHttpsConfiguration` resource allows for parallel creation/update.
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
    ///             Location = "EastUS2",
    ///         });
    ///         var vault = Output.Create(Azure.KeyVault.GetKeyVault.InvokeAsync(new Azure.KeyVault.GetKeyVaultArgs
    ///         {
    ///             Name = "example-vault",
    ///             ResourceGroupName = "example-vault-rg",
    ///         }));
    ///         var exampleFrontdoor = new Azure.FrontDoor.Frontdoor("exampleFrontdoor", new Azure.FrontDoor.FrontdoorArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             EnforceBackendPoolsCertificateNameCheck = false,
    ///             RoutingRules = 
    ///             {
    ///                 new Azure.FrontDoor.Inputs.FrontdoorRoutingRuleArgs
    ///                 {
    ///                     Name = "exampleRoutingRule1",
    ///                     AcceptedProtocols = 
    ///                     {
    ///                         "Http",
    ///                         "Https",
    ///                     },
    ///                     PatternsToMatches = 
    ///                     {
    ///                         "/*",
    ///                     },
    ///                     FrontendEndpoints = 
    ///                     {
    ///                         "exampleFrontendEndpoint1",
    ///                     },
    ///                     ForwardingConfiguration = new Azure.FrontDoor.Inputs.FrontdoorRoutingRuleForwardingConfigurationArgs
    ///                     {
    ///                         ForwardingProtocol = "MatchRequest",
    ///                         BackendPoolName = "exampleBackendBing",
    ///                     },
    ///                 },
    ///             },
    ///             BackendPoolLoadBalancings = 
    ///             {
    ///                 new Azure.FrontDoor.Inputs.FrontdoorBackendPoolLoadBalancingArgs
    ///                 {
    ///                     Name = "exampleLoadBalancingSettings1",
    ///                 },
    ///             },
    ///             BackendPoolHealthProbes = 
    ///             {
    ///                 new Azure.FrontDoor.Inputs.FrontdoorBackendPoolHealthProbeArgs
    ///                 {
    ///                     Name = "exampleHealthProbeSetting1",
    ///                 },
    ///             },
    ///             BackendPools = 
    ///             {
    ///                 new Azure.FrontDoor.Inputs.FrontdoorBackendPoolArgs
    ///                 {
    ///                     Name = "exampleBackendBing",
    ///                     Backends = 
    ///                     {
    ///                         new Azure.FrontDoor.Inputs.FrontdoorBackendPoolBackendArgs
    ///                         {
    ///                             HostHeader = "www.bing.com",
    ///                             Address = "www.bing.com",
    ///                             HttpPort = 80,
    ///                             HttpsPort = 443,
    ///                         },
    ///                     },
    ///                     LoadBalancingName = "exampleLoadBalancingSettings1",
    ///                     HealthProbeName = "exampleHealthProbeSetting1",
    ///                 },
    ///             },
    ///             FrontendEndpoints = 
    ///             {
    ///                 new Azure.FrontDoor.Inputs.FrontdoorFrontendEndpointArgs
    ///                 {
    ///                     Name = "exampleFrontendEndpoint1",
    ///                     HostName = "example-FrontDoor.azurefd.net",
    ///                 },
    ///                 new Azure.FrontDoor.Inputs.FrontdoorFrontendEndpointArgs
    ///                 {
    ///                     Name = "exampleFrontendEndpoint2",
    ///                     HostName = "examplefd1.examplefd.net",
    ///                 },
    ///             },
    ///         });
    ///         var exampleCustomHttps0 = new Azure.FrontDoor.CustomHttpsConfiguration("exampleCustomHttps0", new Azure.FrontDoor.CustomHttpsConfigurationArgs
    ///         {
    ///             FrontendEndpointId = exampleFrontdoor.FrontendEndpoints.Apply(frontendEndpoints =&gt; frontendEndpoints[0].Id),
    ///             CustomHttpsProvisioningEnabled = false,
    ///         });
    ///         var exampleCustomHttps1 = new Azure.FrontDoor.CustomHttpsConfiguration("exampleCustomHttps1", new Azure.FrontDoor.CustomHttpsConfigurationArgs
    ///         {
    ///             FrontendEndpointId = exampleFrontdoor.FrontendEndpoints.Apply(frontendEndpoints =&gt; frontendEndpoints[1].Id),
    ///             CustomHttpsProvisioningEnabled = true,
    ///             CustomHttpsConfiguration = new Azure.FrontDoor.Inputs.CustomHttpsConfigurationCustomHttpsConfigurationArgs
    ///             {
    ///                 CertificateSource = "AzureKeyVault",
    ///                 AzureKeyVaultCertificateSecretName = "examplefd1",
    ///                 AzureKeyVaultCertificateSecretVersion = "ec8d0737e0df4f4gb52ecea858e97a73",
    ///                 AzureKeyVaultCertificateVaultId = vault.Apply(vault =&gt; vault.Id),
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Front Door Custom Https Configurations can be imported using the `resource id` of the Frontend Endpoint, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:frontdoor/customHttpsConfiguration:CustomHttpsConfiguration example_custom_https_1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/frontdoors/frontdoor1/frontendEndpoints/endpoint1
    /// ```
    /// </summary>
    [AzureResourceType("azure:frontdoor/customHttpsConfiguration:CustomHttpsConfiguration")]
    public partial class CustomHttpsConfiguration : Pulumi.CustomResource
    {
        /// <summary>
        /// A `custom_https_configuration` block as defined below.
        /// </summary>
        [Output("customHttpsConfiguration")]
        public Output<Outputs.CustomHttpsConfigurationCustomHttpsConfiguration?> CustomHttpsConfigurationConfig { get; private set; } = null!;

        /// <summary>
        /// Should the HTTPS protocol be enabled for this custom domain associated with the Front Door?
        /// </summary>
        [Output("customHttpsProvisioningEnabled")]
        public Output<bool> CustomHttpsProvisioningEnabled { get; private set; } = null!;

        /// <summary>
        /// The ID of the FrontDoor Frontend Endpoint which this configuration refers to.
        /// </summary>
        [Output("frontendEndpointId")]
        public Output<string> FrontendEndpointId { get; private set; } = null!;

        [Output("resourceGroupName")]
        public Output<string?> ResourceGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a CustomHttpsConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomHttpsConfiguration(string name, CustomHttpsConfigurationArgs args, CustomResourceOptions? options = null)
            : base("azure:frontdoor/customHttpsConfiguration:CustomHttpsConfiguration", name, args ?? new CustomHttpsConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomHttpsConfiguration(string name, Input<string> id, CustomHttpsConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("azure:frontdoor/customHttpsConfiguration:CustomHttpsConfiguration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CustomHttpsConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomHttpsConfiguration Get(string name, Input<string> id, CustomHttpsConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new CustomHttpsConfiguration(name, id, state, options);
        }
    }

    public sealed class CustomHttpsConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `custom_https_configuration` block as defined below.
        /// </summary>
        [Input("customHttpsConfiguration")]
        public Input<Inputs.CustomHttpsConfigurationCustomHttpsConfigurationArgs>? CustomHttpsConfigurationConfig { get; set; }

        /// <summary>
        /// Should the HTTPS protocol be enabled for this custom domain associated with the Front Door?
        /// </summary>
        [Input("customHttpsProvisioningEnabled", required: true)]
        public Input<bool> CustomHttpsProvisioningEnabled { get; set; } = null!;

        /// <summary>
        /// The ID of the FrontDoor Frontend Endpoint which this configuration refers to.
        /// </summary>
        [Input("frontendEndpointId", required: true)]
        public Input<string> FrontendEndpointId { get; set; } = null!;

        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        public CustomHttpsConfigurationArgs()
        {
        }
    }

    public sealed class CustomHttpsConfigurationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `custom_https_configuration` block as defined below.
        /// </summary>
        [Input("customHttpsConfiguration")]
        public Input<Inputs.CustomHttpsConfigurationCustomHttpsConfigurationGetArgs>? CustomHttpsConfigurationConfig { get; set; }

        /// <summary>
        /// Should the HTTPS protocol be enabled for this custom domain associated with the Front Door?
        /// </summary>
        [Input("customHttpsProvisioningEnabled")]
        public Input<bool>? CustomHttpsProvisioningEnabled { get; set; }

        /// <summary>
        /// The ID of the FrontDoor Frontend Endpoint which this configuration refers to.
        /// </summary>
        [Input("frontendEndpointId")]
        public Input<string>? FrontendEndpointId { get; set; }

        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        public CustomHttpsConfigurationState()
        {
        }
    }
}

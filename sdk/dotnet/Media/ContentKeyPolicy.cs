// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Media
{
    /// <summary>
    /// Manages a Content Key Policy.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
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
    ///         var exampleAccount = new Azure.Storage.Account("exampleAccount", new Azure.Storage.AccountArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             AccountTier = "Standard",
    ///             AccountReplicationType = "GRS",
    ///         });
    ///         var exampleServiceAccount = new Azure.Media.ServiceAccount("exampleServiceAccount", new Azure.Media.ServiceAccountArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             StorageAccounts = 
    ///             {
    ///                 new Azure.Media.Inputs.ServiceAccountStorageAccountArgs
    ///                 {
    ///                     Id = exampleAccount.Id,
    ///                     IsPrimary = true,
    ///                 },
    ///             },
    ///         });
    ///         var exampleContentKeyPolicy = new Azure.Media.ContentKeyPolicy("exampleContentKeyPolicy", new Azure.Media.ContentKeyPolicyArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             MediaServicesAccountName = exampleServiceAccount.Name,
    ///             PolicyOptions = 
    ///             {
    ///                 new Azure.Media.Inputs.ContentKeyPolicyPolicyOptionArgs
    ///                 {
    ///                     Name = "fairPlay",
    ///                     FairplayConfiguration = new Azure.Media.Inputs.ContentKeyPolicyPolicyOptionFairplayConfigurationArgs
    ///                     {
    ///                         Ask = "myKeyForFairPlay",
    ///                         Pfx = "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",
    ///                         PfxPassword = "password",
    ///                         RentalDurationSeconds = 2249,
    ///                         RentalAndLeaseKeyType = "PersistentUnlimited",
    ///                     },
    ///                     OpenRestrictionEnabled = true,
    ///                 },
    ///                 new Azure.Media.Inputs.ContentKeyPolicyPolicyOptionArgs
    ///                 {
    ///                     Name = "playReady",
    ///                     PlayreadyConfigurationLicenses = 
    ///                     {
    ///                         new Azure.Media.Inputs.ContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseArgs
    ///                         {
    ///                             AllowTestDevices = true,
    ///                             BeginDate = "2017-10-16T18:22:53Z",
    ///                             PlayRight = new Azure.Media.Inputs.ContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightArgs
    ///                             {
    ///                                 ScmsRestriction = 2,
    ///                                 DigitalVideoOnlyContentRestriction = false,
    ///                                 ImageConstraintForAnalogComponentVideoRestriction = false,
    ///                                 ImageConstraintForAnalogComputerMonitorRestriction = false,
    ///                                 AllowPassingVideoContentToUnknownOutput = "NotAllowed",
    ///                                 UncompressedDigitalVideoOpl = 100,
    ///                                 UncompressedDigitalAudioOpl = 100,
    ///                                 AnalogVideoOpl = 150,
    ///                                 CompressedDigitalAudioOpl = 150,
    ///                             },
    ///                             LicenseType = "Persistent",
    ///                             ContentType = "UltraVioletDownload",
    ///                             ContentKeyLocationFromHeaderEnabled = true,
    ///                         },
    ///                     },
    ///                     OpenRestrictionEnabled = true,
    ///                 },
    ///                 new Azure.Media.Inputs.ContentKeyPolicyPolicyOptionArgs
    ///                 {
    ///                     Name = "clearKey",
    ///                     ClearKeyConfigurationEnabled = true,
    ///                     TokenRestriction = new Azure.Media.Inputs.ContentKeyPolicyPolicyOptionTokenRestrictionArgs
    ///                     {
    ///                         Issuer = "urn:issuer",
    ///                         Audience = "urn:audience",
    ///                         TokenType = "Swt",
    ///                         PrimarySymmetricTokenKey = "AAAAAAAAAAAAAAAAAAAAAA==",
    ///                     },
    ///                 },
    ///                 new Azure.Media.Inputs.ContentKeyPolicyPolicyOptionArgs
    ///                 {
    ///                     Name = "widevine",
    ///                     WidevineConfigurationTemplate = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         { "allowed_track_types", "SD_HD" },
    ///                         { "content_key_specs", new[]
    ///                             {
    ///                                 new Dictionary&lt;string, object?&gt;
    ///                                 {
    ///                                     { "track_type", "SD" },
    ///                                     { "security_level", 1 },
    ///                                     { "required_output_protection", new Dictionary&lt;string, object?&gt;
    ///                                     {
    ///                                         { "hdcp", "HDCP_V2" },
    ///                                     } },
    ///                                 },
    ///                             }
    ///                          },
    ///                         { "policy_overrides", new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             { "can_play", true },
    ///                             { "can_persist", true },
    ///                             { "can_renew", false },
    ///                         } },
    ///                     }),
    ///                     OpenRestrictionEnabled = true,
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
    /// Resource Groups can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:media/contentKeyPolicy:ContentKeyPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Media/mediaservices/account1/contentkeypolicies/policy1
    /// ```
    /// </summary>
    [AzureResourceType("azure:media/contentKeyPolicy:ContentKeyPolicy")]
    public partial class ContentKeyPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// A description for the Policy.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Media Services account name. Changing this forces a new Content Key Policy to be created.
        /// </summary>
        [Output("mediaServicesAccountName")]
        public Output<string> MediaServicesAccountName { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this Content Key Policy. Changing this forces a new Content Key Policy to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// One or more `policy_option` blocks as defined below.
        /// </summary>
        [Output("policyOptions")]
        public Output<ImmutableArray<Outputs.ContentKeyPolicyPolicyOption>> PolicyOptions { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Content Key Policy should exist. Changing this forces a new Content Key Policy to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a ContentKeyPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ContentKeyPolicy(string name, ContentKeyPolicyArgs args, CustomResourceOptions? options = null)
            : base("azure:media/contentKeyPolicy:ContentKeyPolicy", name, args ?? new ContentKeyPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ContentKeyPolicy(string name, Input<string> id, ContentKeyPolicyState? state = null, CustomResourceOptions? options = null)
            : base("azure:media/contentKeyPolicy:ContentKeyPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ContentKeyPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ContentKeyPolicy Get(string name, Input<string> id, ContentKeyPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new ContentKeyPolicy(name, id, state, options);
        }
    }

    public sealed class ContentKeyPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description for the Policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Media Services account name. Changing this forces a new Content Key Policy to be created.
        /// </summary>
        [Input("mediaServicesAccountName", required: true)]
        public Input<string> MediaServicesAccountName { get; set; } = null!;

        /// <summary>
        /// The name which should be used for this Content Key Policy. Changing this forces a new Content Key Policy to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policyOptions", required: true)]
        private InputList<Inputs.ContentKeyPolicyPolicyOptionArgs>? _policyOptions;

        /// <summary>
        /// One or more `policy_option` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ContentKeyPolicyPolicyOptionArgs> PolicyOptions
        {
            get => _policyOptions ?? (_policyOptions = new InputList<Inputs.ContentKeyPolicyPolicyOptionArgs>());
            set => _policyOptions = value;
        }

        /// <summary>
        /// The name of the Resource Group where the Content Key Policy should exist. Changing this forces a new Content Key Policy to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ContentKeyPolicyArgs()
        {
        }
    }

    public sealed class ContentKeyPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description for the Policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Media Services account name. Changing this forces a new Content Key Policy to be created.
        /// </summary>
        [Input("mediaServicesAccountName")]
        public Input<string>? MediaServicesAccountName { get; set; }

        /// <summary>
        /// The name which should be used for this Content Key Policy. Changing this forces a new Content Key Policy to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policyOptions")]
        private InputList<Inputs.ContentKeyPolicyPolicyOptionGetArgs>? _policyOptions;

        /// <summary>
        /// One or more `policy_option` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ContentKeyPolicyPolicyOptionGetArgs> PolicyOptions
        {
            get => _policyOptions ?? (_policyOptions = new InputList<Inputs.ContentKeyPolicyPolicyOptionGetArgs>());
            set => _policyOptions = value;
        }

        /// <summary>
        /// The name of the Resource Group where the Content Key Policy should exist. Changing this forces a new Content Key Policy to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        public ContentKeyPolicyState()
        {
        }
    }
}

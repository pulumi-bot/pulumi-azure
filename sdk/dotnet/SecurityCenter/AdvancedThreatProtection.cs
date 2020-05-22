// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.SecurityCenter
{
    /// <summary>
    /// Manages a resources Advanced Threat Protection setting.
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var rg = new Azure.Core.ResourceGroup("rg", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "northeurope",
    ///         });
    ///         var exampleAccount = new Azure.Storage.Account("exampleAccount", new Azure.Storage.AccountArgs
    ///         {
    ///             ResourceGroupName = azurerm_resource_group.Example.Name,
    ///             Location = azurerm_resource_group.Example.Location,
    ///             AccountTier = "Standard",
    ///             AccountReplicationType = "LRS",
    ///             Tags = 
    ///             {
    ///                 { "environment", "example" },
    ///             },
    ///         });
    ///         var exampleAdvancedThreatProtection = new Azure.SecurityCenter.AdvancedThreatProtection("exampleAdvancedThreatProtection", new Azure.SecurityCenter.AdvancedThreatProtectionArgs
    ///         {
    ///             TargetResourceId = exampleAccount.Id,
    ///             Enabled = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class AdvancedThreatProtection : Pulumi.CustomResource
    {
        /// <summary>
        /// Should Advanced Threat Protection be enabled on this resource?
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// The ID of the Azure Resource which to enable Advanced Threat Protection on. Changing this forces a new resource to be created.
        /// </summary>
        [Output("targetResourceId")]
        public Output<string> TargetResourceId { get; private set; } = null!;


        /// <summary>
        /// Create a AdvancedThreatProtection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AdvancedThreatProtection(string name, AdvancedThreatProtectionArgs args, CustomResourceOptions? options = null)
            : base("azure:securitycenter/advancedThreatProtection:AdvancedThreatProtection", name, args ?? new AdvancedThreatProtectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AdvancedThreatProtection(string name, Input<string> id, AdvancedThreatProtectionState? state = null, CustomResourceOptions? options = null)
            : base("azure:securitycenter/advancedThreatProtection:AdvancedThreatProtection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AdvancedThreatProtection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AdvancedThreatProtection Get(string name, Input<string> id, AdvancedThreatProtectionState? state = null, CustomResourceOptions? options = null)
        {
            return new AdvancedThreatProtection(name, id, state, options);
        }
    }

    public sealed class AdvancedThreatProtectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Should Advanced Threat Protection be enabled on this resource?
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// The ID of the Azure Resource which to enable Advanced Threat Protection on. Changing this forces a new resource to be created.
        /// </summary>
        [Input("targetResourceId", required: true)]
        public Input<string> TargetResourceId { get; set; } = null!;

        public AdvancedThreatProtectionArgs()
        {
        }
    }

    public sealed class AdvancedThreatProtectionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Should Advanced Threat Protection be enabled on this resource?
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The ID of the Azure Resource which to enable Advanced Threat Protection on. Changing this forces a new resource to be created.
        /// </summary>
        [Input("targetResourceId")]
        public Input<string>? TargetResourceId { get; set; }

        public AdvancedThreatProtectionState()
        {
        }
    }
}

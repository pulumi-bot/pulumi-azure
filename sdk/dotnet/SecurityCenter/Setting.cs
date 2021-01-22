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
    /// Manages the Data Access Settings for Azure Security Center.
    /// 
    /// &gt; **NOTE:** This resource requires the `Owner` permission on the Subscription.
    /// 
    /// &gt; **NOTE:** Deletion of this resource does not change or reset the data access settings
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
    ///         var example = new Azure.SecurityCenter.Setting("example", new Azure.SecurityCenter.SettingArgs
    ///         {
    ///             Enabled = true,
    ///             SettingName = "MCAS",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// The setting can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:securitycenter/setting:Setting example /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Security/settings/&lt;setting_name&gt;
    /// ```
    /// </summary>
    [AzureResourceType("azure:securitycenter/setting:Setting")]
    public partial class Setting : Pulumi.CustomResource
    {
        /// <summary>
        /// Boolean flag to enable/disable data access.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// The setting to manage. Possible values are `MCAS` and `WDATP`.
        /// </summary>
        [Output("settingName")]
        public Output<string> SettingName { get; private set; } = null!;


        /// <summary>
        /// Create a Setting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Setting(string name, SettingArgs args, CustomResourceOptions? options = null)
            : base("azure:securitycenter/setting:Setting", name, args ?? new SettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Setting(string name, Input<string> id, SettingState? state = null, CustomResourceOptions? options = null)
            : base("azure:securitycenter/setting:Setting", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Setting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Setting Get(string name, Input<string> id, SettingState? state = null, CustomResourceOptions? options = null)
        {
            return new Setting(name, id, state, options);
        }
    }

    public sealed class SettingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean flag to enable/disable data access.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// The setting to manage. Possible values are `MCAS` and `WDATP`.
        /// </summary>
        [Input("settingName", required: true)]
        public Input<string> SettingName { get; set; } = null!;

        public SettingArgs()
        {
        }
    }

    public sealed class SettingState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean flag to enable/disable data access.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The setting to manage. Possible values are `MCAS` and `WDATP`.
        /// </summary>
        [Input("settingName")]
        public Input<string>? SettingName { get; set; }

        public SettingState()
        {
        }
    }
}

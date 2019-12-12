// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService
{
    /// <summary>
    /// Promotes an App Service Slot to Production within an App Service.
    /// 
    /// &gt; **Note:** When using Slots - the `app_settings`, `connection_string` and `site_config` blocks on the `azure.appservice.AppService` resource will be overwritten when promoting a Slot using the `azure.appservice.ActiveSlot` resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/app_service_active_slot.html.markdown.
    /// </summary>
    public partial class ActiveSlot : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the App Service within which the Slot exists.  Changing this forces a new resource to be created.
        /// </summary>
        [Output("appServiceName")]
        public Output<string> AppServiceName { get; private set; } = null!;

        /// <summary>
        /// The name of the App Service Slot which should be promoted to the Production Slot within the App Service.
        /// </summary>
        [Output("appServiceSlotName")]
        public Output<string> AppServiceSlotName { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a ActiveSlot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ActiveSlot(string name, ActiveSlotArgs args, CustomResourceOptions? options = null)
            : base("azure:appservice/activeSlot:ActiveSlot", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ActiveSlot(string name, Input<string> id, ActiveSlotState? state = null, CustomResourceOptions? options = null)
            : base("azure:appservice/activeSlot:ActiveSlot", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ActiveSlot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ActiveSlot Get(string name, Input<string> id, ActiveSlotState? state = null, CustomResourceOptions? options = null)
        {
            return new ActiveSlot(name, id, state, options);
        }
    }

    public sealed class ActiveSlotArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the App Service within which the Slot exists.  Changing this forces a new resource to be created.
        /// </summary>
        [Input("appServiceName", required: true)]
        public Input<string> AppServiceName { get; set; } = null!;

        /// <summary>
        /// The name of the App Service Slot which should be promoted to the Production Slot within the App Service.
        /// </summary>
        [Input("appServiceSlotName", required: true)]
        public Input<string> AppServiceSlotName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ActiveSlotArgs()
        {
        }
    }

    public sealed class ActiveSlotState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the App Service within which the Slot exists.  Changing this forces a new resource to be created.
        /// </summary>
        [Input("appServiceName")]
        public Input<string>? AppServiceName { get; set; }

        /// <summary>
        /// The name of the App Service Slot which should be promoted to the Production Slot within the App Service.
        /// </summary>
        [Input("appServiceSlotName")]
        public Input<string>? AppServiceSlotName { get; set; }

        /// <summary>
        /// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        public ActiveSlotState()
        {
        }
    }
}

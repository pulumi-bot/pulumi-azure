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
    /// Manages the Pricing Tier for Azure Security Center in the current subscription.
    /// 
    /// &gt; **NOTE:** This resource requires the `Owner` permission on the Subscription.
    /// 
    /// &gt; **NOTE:** Deletion of this resource does not change or reset the pricing tier to `Free`
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
    ///         var example = new Azure.SecurityCenter.SubscriptionPricing("example", new Azure.SecurityCenter.SubscriptionPricingArgs
    ///         {
    ///             ResourceType = "VirtualMachines",
    ///             Tier = "Standard",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// The pricing tier can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class SubscriptionPricing : Pulumi.CustomResource
    {
        /// <summary>
        /// The resource type this setting affects. Possible values are `AppServices`, `ContainerRegistry`, `KeyVaults`, `KubernetesService`, `SqlServers`, `SqlServerVirtualMachines`, `StorageAccounts`, and `VirtualMachines`.
        /// </summary>
        [Output("resourceType")]
        public Output<string?> ResourceType { get; private set; } = null!;

        /// <summary>
        /// The pricing tier to use. Possible values are `Free` and `Standard`.
        /// </summary>
        [Output("tier")]
        public Output<string> Tier { get; private set; } = null!;


        /// <summary>
        /// Create a SubscriptionPricing resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SubscriptionPricing(string name, SubscriptionPricingArgs args, CustomResourceOptions? options = null)
            : base("azure:securitycenter/subscriptionPricing:SubscriptionPricing", name, args ?? new SubscriptionPricingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SubscriptionPricing(string name, Input<string> id, SubscriptionPricingState? state = null, CustomResourceOptions? options = null)
            : base("azure:securitycenter/subscriptionPricing:SubscriptionPricing", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SubscriptionPricing resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SubscriptionPricing Get(string name, Input<string> id, SubscriptionPricingState? state = null, CustomResourceOptions? options = null)
        {
            return new SubscriptionPricing(name, id, state, options);
        }
    }

    public sealed class SubscriptionPricingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource type this setting affects. Possible values are `AppServices`, `ContainerRegistry`, `KeyVaults`, `KubernetesService`, `SqlServers`, `SqlServerVirtualMachines`, `StorageAccounts`, and `VirtualMachines`.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        /// <summary>
        /// The pricing tier to use. Possible values are `Free` and `Standard`.
        /// </summary>
        [Input("tier", required: true)]
        public Input<string> Tier { get; set; } = null!;

        public SubscriptionPricingArgs()
        {
        }
    }

    public sealed class SubscriptionPricingState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource type this setting affects. Possible values are `AppServices`, `ContainerRegistry`, `KeyVaults`, `KubernetesService`, `SqlServers`, `SqlServerVirtualMachines`, `StorageAccounts`, and `VirtualMachines`.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        /// <summary>
        /// The pricing tier to use. Possible values are `Free` and `Standard`.
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        public SubscriptionPricingState()
        {
        }
    }
}

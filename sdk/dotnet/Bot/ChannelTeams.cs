// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Bot
{
    /// <summary>
    /// Manages a MS Teams integration for a Bot Channel
    /// 
    /// &gt; **Note** A bot can only have a single MS Teams Channel associated with it.
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
    ///         var current = Output.Create(Azure.Core.GetClientConfig.InvokeAsync());
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "northeurope",
    ///         });
    ///         var exampleChannelsRegistration = new Azure.Bot.ChannelsRegistration("exampleChannelsRegistration", new Azure.Bot.ChannelsRegistrationArgs
    ///         {
    ///             Location = "global",
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Sku = "F0",
    ///             MicrosoftAppId = current.Apply(current =&gt; current.ClientId),
    ///         });
    ///         var exampleChannelTeams = new Azure.Bot.ChannelTeams("exampleChannelTeams", new Azure.Bot.ChannelTeamsArgs
    ///         {
    ///             BotName = exampleChannelsRegistration.Name,
    ///             Location = exampleChannelsRegistration.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// The Microsoft Teams Integration for a Bot Channel can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:bot/channelTeams:ChannelTeams example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example/channels/MsTeamsChannel
    /// ```
    /// </summary>
    [AzureResourceType("azure:bot/channelTeams:ChannelTeams")]
    public partial class ChannelTeams : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
        /// </summary>
        [Output("botName")]
        public Output<string> BotName { get; private set; } = null!;

        /// <summary>
        /// Specifies the webhook for Microsoft Teams channel calls.
        /// </summary>
        [Output("callingWebHook")]
        public Output<string> CallingWebHook { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to enable Microsoft Teams channel calls. This defaults to `false`.
        /// </summary>
        [Output("enableCalling")]
        public Output<bool?> EnableCalling { get; private set; } = null!;

        /// <summary>
        /// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a ChannelTeams resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ChannelTeams(string name, ChannelTeamsArgs args, CustomResourceOptions? options = null)
            : base("azure:bot/channelTeams:ChannelTeams", name, args ?? new ChannelTeamsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ChannelTeams(string name, Input<string> id, ChannelTeamsState? state = null, CustomResourceOptions? options = null)
            : base("azure:bot/channelTeams:ChannelTeams", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ChannelTeams resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ChannelTeams Get(string name, Input<string> id, ChannelTeamsState? state = null, CustomResourceOptions? options = null)
        {
            return new ChannelTeams(name, id, state, options);
        }
    }

    public sealed class ChannelTeamsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
        /// </summary>
        [Input("botName", required: true)]
        public Input<string> BotName { get; set; } = null!;

        /// <summary>
        /// Specifies the webhook for Microsoft Teams channel calls.
        /// </summary>
        [Input("callingWebHook")]
        public Input<string>? CallingWebHook { get; set; }

        /// <summary>
        /// Specifies whether to enable Microsoft Teams channel calls. This defaults to `false`.
        /// </summary>
        [Input("enableCalling")]
        public Input<bool>? EnableCalling { get; set; }

        /// <summary>
        /// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ChannelTeamsArgs()
        {
        }
    }

    public sealed class ChannelTeamsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
        /// </summary>
        [Input("botName")]
        public Input<string>? BotName { get; set; }

        /// <summary>
        /// Specifies the webhook for Microsoft Teams channel calls.
        /// </summary>
        [Input("callingWebHook")]
        public Input<string>? CallingWebHook { get; set; }

        /// <summary>
        /// Specifies whether to enable Microsoft Teams channel calls. This defaults to `false`.
        /// </summary>
        [Input("enableCalling")]
        public Input<bool>? EnableCalling { get; set; }

        /// <summary>
        /// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        public ChannelTeamsState()
        {
        }
    }
}

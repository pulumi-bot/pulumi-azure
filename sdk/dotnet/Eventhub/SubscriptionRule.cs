// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventHub
{
    /// <summary>
    /// Manages a ServiceBus Subscription Rule.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/servicebus_subscription_rule_legacy.html.markdown.
    /// </summary>
    public partial class SubscriptionRule : Pulumi.CustomResource
    {
        /// <summary>
        /// Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
        /// </summary>
        [Output("action")]
        public Output<string?> Action { get; private set; } = null!;

        /// <summary>
        /// A `correlation_filter` block as documented below to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `CorrelationFilter`.
        /// </summary>
        [Output("correlationFilter")]
        public Output<Outputs.SubscriptionRuleCorrelationFilter?> CorrelationFilter { get; private set; } = null!;

        /// <summary>
        /// Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
        /// </summary>
        [Output("filterType")]
        public Output<string> FilterType { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the ServiceBus Namespace in which the ServiceBus Topic exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("namespaceName")]
        public Output<string> NamespaceName { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in the ServiceBus Namespace exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `SqlFilter`.
        /// </summary>
        [Output("sqlFilter")]
        public Output<string?> SqlFilter { get; private set; } = null!;

        /// <summary>
        /// The name of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("subscriptionName")]
        public Output<string> SubscriptionName { get; private set; } = null!;

        /// <summary>
        /// The name of the ServiceBus Topic in which the ServiceBus Subscription exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("topicName")]
        public Output<string> TopicName { get; private set; } = null!;


        /// <summary>
        /// Create a SubscriptionRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SubscriptionRule(string name, SubscriptionRuleArgs args, CustomResourceOptions? options = null)
            : base("azure:eventhub/subscriptionRule:SubscriptionRule", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SubscriptionRule(string name, Input<string> id, SubscriptionRuleState? state = null, CustomResourceOptions? options = null)
            : base("azure:eventhub/subscriptionRule:SubscriptionRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SubscriptionRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SubscriptionRule Get(string name, Input<string> id, SubscriptionRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new SubscriptionRule(name, id, state, options);
        }
    }

    public sealed class SubscriptionRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// A `correlation_filter` block as documented below to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `CorrelationFilter`.
        /// </summary>
        [Input("correlationFilter")]
        public Input<Inputs.SubscriptionRuleCorrelationFilterArgs>? CorrelationFilter { get; set; }

        /// <summary>
        /// Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
        /// </summary>
        [Input("filterType", required: true)]
        public Input<string> FilterType { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the ServiceBus Namespace in which the ServiceBus Topic exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in the ServiceBus Namespace exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `SqlFilter`.
        /// </summary>
        [Input("sqlFilter")]
        public Input<string>? SqlFilter { get; set; }

        /// <summary>
        /// The name of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subscriptionName", required: true)]
        public Input<string> SubscriptionName { get; set; } = null!;

        /// <summary>
        /// The name of the ServiceBus Topic in which the ServiceBus Subscription exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("topicName", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        public SubscriptionRuleArgs()
        {
        }
    }

    public sealed class SubscriptionRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// A `correlation_filter` block as documented below to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `CorrelationFilter`.
        /// </summary>
        [Input("correlationFilter")]
        public Input<Inputs.SubscriptionRuleCorrelationFilterGetArgs>? CorrelationFilter { get; set; }

        /// <summary>
        /// Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
        /// </summary>
        [Input("filterType")]
        public Input<string>? FilterType { get; set; }

        /// <summary>
        /// Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the ServiceBus Namespace in which the ServiceBus Topic exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("namespaceName")]
        public Input<string>? NamespaceName { get; set; }

        /// <summary>
        /// The name of the resource group in the ServiceBus Namespace exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `SqlFilter`.
        /// </summary>
        [Input("sqlFilter")]
        public Input<string>? SqlFilter { get; set; }

        /// <summary>
        /// The name of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subscriptionName")]
        public Input<string>? SubscriptionName { get; set; }

        /// <summary>
        /// The name of the ServiceBus Topic in which the ServiceBus Subscription exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("topicName")]
        public Input<string>? TopicName { get; set; }

        public SubscriptionRuleState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class SubscriptionRuleCorrelationFilterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Content type of the message.
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// Identifier of the correlation.
        /// </summary>
        [Input("correlationId")]
        public Input<string>? CorrelationId { get; set; }

        /// <summary>
        /// Application specific label.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Identifier of the message.
        /// </summary>
        [Input("messageId")]
        public Input<string>? MessageId { get; set; }

        /// <summary>
        /// Address of the queue to reply to.
        /// </summary>
        [Input("replyTo")]
        public Input<string>? ReplyTo { get; set; }

        /// <summary>
        /// Session identifier to reply to.
        /// </summary>
        [Input("replyToSessionId")]
        public Input<string>? ReplyToSessionId { get; set; }

        /// <summary>
        /// Session identifier.
        /// </summary>
        [Input("sessionId")]
        public Input<string>? SessionId { get; set; }

        /// <summary>
        /// Address to send to.
        /// </summary>
        [Input("to")]
        public Input<string>? To { get; set; }

        public SubscriptionRuleCorrelationFilterArgs()
        {
        }
    }

    public sealed class SubscriptionRuleCorrelationFilterGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Content type of the message.
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// Identifier of the correlation.
        /// </summary>
        [Input("correlationId")]
        public Input<string>? CorrelationId { get; set; }

        /// <summary>
        /// Application specific label.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Identifier of the message.
        /// </summary>
        [Input("messageId")]
        public Input<string>? MessageId { get; set; }

        /// <summary>
        /// Address of the queue to reply to.
        /// </summary>
        [Input("replyTo")]
        public Input<string>? ReplyTo { get; set; }

        /// <summary>
        /// Session identifier to reply to.
        /// </summary>
        [Input("replyToSessionId")]
        public Input<string>? ReplyToSessionId { get; set; }

        /// <summary>
        /// Session identifier.
        /// </summary>
        [Input("sessionId")]
        public Input<string>? SessionId { get; set; }

        /// <summary>
        /// Address to send to.
        /// </summary>
        [Input("to")]
        public Input<string>? To { get; set; }

        public SubscriptionRuleCorrelationFilterGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class SubscriptionRuleCorrelationFilter
    {
        /// <summary>
        /// Content type of the message.
        /// </summary>
        public readonly string? ContentType;
        /// <summary>
        /// Identifier of the correlation.
        /// </summary>
        public readonly string? CorrelationId;
        /// <summary>
        /// Application specific label.
        /// </summary>
        public readonly string? Label;
        /// <summary>
        /// Identifier of the message.
        /// </summary>
        public readonly string? MessageId;
        /// <summary>
        /// Address of the queue to reply to.
        /// </summary>
        public readonly string? ReplyTo;
        /// <summary>
        /// Session identifier to reply to.
        /// </summary>
        public readonly string? ReplyToSessionId;
        /// <summary>
        /// Session identifier.
        /// </summary>
        public readonly string? SessionId;
        /// <summary>
        /// Address to send to.
        /// </summary>
        public readonly string? To;

        [OutputConstructor]
        private SubscriptionRuleCorrelationFilter(
            string? contentType,
            string? correlationId,
            string? label,
            string? messageId,
            string? replyTo,
            string? replyToSessionId,
            string? sessionId,
            string? to)
        {
            ContentType = contentType;
            CorrelationId = correlationId;
            Label = label;
            MessageId = messageId;
            ReplyTo = replyTo;
            ReplyToSessionId = replyToSessionId;
            SessionId = sessionId;
            To = to;
        }
    }
    }
}

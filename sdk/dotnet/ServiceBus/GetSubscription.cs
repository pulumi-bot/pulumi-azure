// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ServiceBus
{
    public static class GetSubscription
    {
        /// <summary>
        /// Use this data source to access information about an existing ServiceBus Subscription.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.ServiceBus.GetSubscription.InvokeAsync(new Azure.ServiceBus.GetSubscriptionArgs
        ///         {
        ///             Name = "examplesubscription",
        ///             ResourceGroupName = "exampleresources",
        ///             NamespaceName = "examplenamespace",
        ///             TopicName = "exampletopic",
        ///         }));
        ///         this.ServicebusSubscription = data.Azurerm_servicebus_namespace.Example;
        ///     }
        /// 
        ///     [Output("servicebusSubscription")]
        ///     public Output&lt;string&gt; ServicebusSubscription { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSubscriptionResult> InvokeAsync(GetSubscriptionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSubscriptionResult>("azure:servicebus/getSubscription:getSubscription", args ?? new GetSubscriptionArgs(), options.WithVersion());

        public static Output<GetSubscriptionResult> Invoke(GetSubscriptionOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Name.Box(),
                args.NamespaceName.Box(),
                args.ResourceGroupName.Box(),
                args.TopicName.Box()
            ).Apply(a => {
                    var args = new GetSubscriptionArgs();
                    a[0].Set(args, nameof(args.Name));
                    a[1].Set(args, nameof(args.NamespaceName));
                    a[2].Set(args, nameof(args.ResourceGroupName));
                    a[3].Set(args, nameof(args.TopicName));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetSubscriptionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the ServiceBus Subscription.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the ServiceBus Namespace.
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group where the ServiceBus Namespace exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the ServiceBus Topic.
        /// </summary>
        [Input("topicName", required: true)]
        public string TopicName { get; set; } = null!;

        public GetSubscriptionArgs()
        {
        }
    }

    public sealed class GetSubscriptionOutputArgs
    {
        /// <summary>
        /// Specifies the name of the ServiceBus Subscription.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the ServiceBus Namespace.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group where the ServiceBus Namespace exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the ServiceBus Topic.
        /// </summary>
        [Input("topicName", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        public GetSubscriptionOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSubscriptionResult
    {
        /// <summary>
        /// The idle interval after which the topic is automatically deleted.
        /// </summary>
        public readonly string AutoDeleteOnIdle;
        /// <summary>
        /// Does the ServiceBus Subscription have dead letter support on filter evaluation exceptions?
        /// </summary>
        public readonly bool DeadLetteringOnFilterEvaluationError;
        /// <summary>
        /// Does the Service Bus Subscription have dead letter support when a message expires?
        /// </summary>
        public readonly bool DeadLetteringOnMessageExpiration;
        /// <summary>
        /// The Default message timespan to live. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
        /// </summary>
        public readonly string DefaultMessageTtl;
        /// <summary>
        /// Are batched operations enabled on this ServiceBus Subscription?
        /// </summary>
        public readonly bool EnableBatchedOperations;
        /// <summary>
        /// The name of a Queue or Topic to automatically forward Dead Letter messages to.
        /// </summary>
        public readonly string ForwardDeadLetteredMessagesTo;
        /// <summary>
        /// The name of a ServiceBus Queue or ServiceBus Topic where messages are automatically forwarded.
        /// </summary>
        public readonly string ForwardTo;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The lock duration for the subscription.
        /// </summary>
        public readonly string LockDuration;
        /// <summary>
        /// The maximum number of deliveries.
        /// </summary>
        public readonly int MaxDeliveryCount;
        public readonly string Name;
        public readonly string NamespaceName;
        /// <summary>
        /// Whether or not this ServiceBus Subscription supports session.
        /// </summary>
        public readonly bool RequiresSession;
        public readonly string ResourceGroupName;
        public readonly string TopicName;

        [OutputConstructor]
        private GetSubscriptionResult(
            string autoDeleteOnIdle,

            bool deadLetteringOnFilterEvaluationError,

            bool deadLetteringOnMessageExpiration,

            string defaultMessageTtl,

            bool enableBatchedOperations,

            string forwardDeadLetteredMessagesTo,

            string forwardTo,

            string id,

            string lockDuration,

            int maxDeliveryCount,

            string name,

            string namespaceName,

            bool requiresSession,

            string resourceGroupName,

            string topicName)
        {
            AutoDeleteOnIdle = autoDeleteOnIdle;
            DeadLetteringOnFilterEvaluationError = deadLetteringOnFilterEvaluationError;
            DeadLetteringOnMessageExpiration = deadLetteringOnMessageExpiration;
            DefaultMessageTtl = defaultMessageTtl;
            EnableBatchedOperations = enableBatchedOperations;
            ForwardDeadLetteredMessagesTo = forwardDeadLetteredMessagesTo;
            ForwardTo = forwardTo;
            Id = id;
            LockDuration = lockDuration;
            MaxDeliveryCount = maxDeliveryCount;
            Name = name;
            NamespaceName = namespaceName;
            RequiresSession = requiresSession;
            ResourceGroupName = resourceGroupName;
            TopicName = topicName;
        }
    }
}

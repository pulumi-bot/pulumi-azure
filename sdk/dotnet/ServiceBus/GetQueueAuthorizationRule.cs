// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ServiceBus
{
    public static class GetQueueAuthorizationRule
    {
        /// <summary>
        /// Use this data source to access information about an existing ServiceBus Queue Authorisation Rule within a ServiceBus Queue.
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
        ///         var example = Output.Create(Azure.ServiceBus.GetQueueAuthorizationRule.InvokeAsync(new Azure.ServiceBus.GetQueueAuthorizationRuleArgs
        ///         {
        ///             Name = "example-tfex_name",
        ///             ResourceGroupName = "example-resources",
        ///             QueueName = "example-servicebus_queue",
        ///             NamespaceName = "example-namespace",
        ///         }));
        ///         this.Id = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetQueueAuthorizationRuleResult> InvokeAsync(GetQueueAuthorizationRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetQueueAuthorizationRuleResult>("azure:servicebus/getQueueAuthorizationRule:getQueueAuthorizationRule", args ?? new GetQueueAuthorizationRuleArgs(), options.WithVersion());

        public static Output<GetQueueAuthorizationRuleResult> Invoke(GetQueueAuthorizationRuleOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Name.Box(),
                args.NamespaceName.Box(),
                args.QueueName.Box(),
                args.ResourceGroupName.Box()
            ).Apply(a => {
                    var args = new GetQueueAuthorizationRuleArgs();
                    a[0].Set(args, nameof(args.Name));
                    a[1].Set(args, nameof(args.NamespaceName));
                    a[2].Set(args, nameof(args.QueueName));
                    a[3].Set(args, nameof(args.ResourceGroupName));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetQueueAuthorizationRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this ServiceBus Queue Authorisation Rule.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the ServiceBus Namespace.
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// The name of the ServiceBus Queue.
        /// </summary>
        [Input("queueName", required: true)]
        public string QueueName { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the ServiceBus Queue Authorisation Rule exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetQueueAuthorizationRuleArgs()
        {
        }
    }

    public sealed class GetQueueAuthorizationRuleOutputArgs
    {
        /// <summary>
        /// The name of this ServiceBus Queue Authorisation Rule.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the ServiceBus Namespace.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// The name of the ServiceBus Queue.
        /// </summary>
        [Input("queueName", required: true)]
        public Input<string> QueueName { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the ServiceBus Queue Authorisation Rule exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetQueueAuthorizationRuleOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetQueueAuthorizationRuleResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool Listen;
        public readonly bool Manage;
        public readonly string Name;
        public readonly string NamespaceName;
        /// <summary>
        /// The Primary Connection String for the ServiceBus Queue authorization Rule.
        /// </summary>
        public readonly string PrimaryConnectionString;
        /// <summary>
        /// The Primary Key for the ServiceBus Queue authorization Rule.
        /// </summary>
        public readonly string PrimaryKey;
        public readonly string QueueName;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The Secondary Connection String for the ServiceBus Queue authorization Rule.
        /// </summary>
        public readonly string SecondaryConnectionString;
        /// <summary>
        /// The Secondary Key for the ServiceBus Queue authorization Rule.
        /// </summary>
        public readonly string SecondaryKey;
        public readonly bool Send;

        [OutputConstructor]
        private GetQueueAuthorizationRuleResult(
            string id,

            bool listen,

            bool manage,

            string name,

            string namespaceName,

            string primaryConnectionString,

            string primaryKey,

            string queueName,

            string resourceGroupName,

            string secondaryConnectionString,

            string secondaryKey,

            bool send)
        {
            Id = id;
            Listen = listen;
            Manage = manage;
            Name = name;
            NamespaceName = namespaceName;
            PrimaryConnectionString = primaryConnectionString;
            PrimaryKey = primaryKey;
            QueueName = queueName;
            ResourceGroupName = resourceGroupName;
            SecondaryConnectionString = secondaryConnectionString;
            SecondaryKey = secondaryKey;
            Send = send;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventHub
{
    public static class GetNamespaceAuthorizationRule
    {
        /// <summary>
        /// Use this data source to access information about an Authorization Rule for an Event Hub Namespace.
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
        ///         var example = Output.Create(Azure.EventHub.GetNamespaceAuthorizationRule.InvokeAsync(new Azure.EventHub.GetNamespaceAuthorizationRuleArgs
        ///         {
        ///             Name = "navi",
        ///             ResourceGroupName = "example-resources",
        ///             NamespaceName = "example-ns",
        ///         }));
        ///         this.EventhubAuthorizationRuleId = data.Azurem_eventhub_namespace_authorization_rule.Example.Id;
        ///     }
        /// 
        ///     [Output("eventhubAuthorizationRuleId")]
        ///     public Output&lt;string&gt; EventhubAuthorizationRuleId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNamespaceAuthorizationRuleResult> InvokeAsync(GetNamespaceAuthorizationRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNamespaceAuthorizationRuleResult>("azure:eventhub/getNamespaceAuthorizationRule:getNamespaceAuthorizationRule", args ?? new GetNamespaceAuthorizationRuleArgs(), options.WithVersion());

        public static Output<GetNamespaceAuthorizationRuleResult> Apply(GetNamespaceAuthorizationRuleApplyArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Name.Box(),
                args.NamespaceName.Box(),
                args.ResourceGroupName.Box()
            ).Apply(a => {
                    var args = new GetNamespaceAuthorizationRuleArgs();
                    a[0].Set(args, nameof(args.Name));
                    a[1].Set(args, nameof(args.NamespaceName));
                    a[2].Set(args, nameof(args.ResourceGroupName));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetNamespaceAuthorizationRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the EventHub Authorization Rule resource.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the EventHub Namespace.
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the EventHub Namespace exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNamespaceAuthorizationRuleArgs()
        {
        }
    }

    public sealed class GetNamespaceAuthorizationRuleApplyArgs
    {
        /// <summary>
        /// The name of the EventHub Authorization Rule resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the EventHub Namespace.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the EventHub Namespace exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetNamespaceAuthorizationRuleApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNamespaceAuthorizationRuleResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Does this Authorization Rule have permissions to Listen to the Event Hub?
        /// </summary>
        public readonly bool Listen;
        /// <summary>
        /// Does this Authorization Rule have permissions to Manage to the Event Hub?
        /// </summary>
        public readonly bool Manage;
        public readonly string Name;
        public readonly string NamespaceName;
        /// <summary>
        /// The Primary Connection String for the Event Hubs authorization Rule.
        /// </summary>
        public readonly string PrimaryConnectionString;
        /// <summary>
        /// The alias of the Primary Connection String for the Event Hubs authorization Rule.
        /// </summary>
        public readonly string PrimaryConnectionStringAlias;
        /// <summary>
        /// The Primary Key for the Event Hubs authorization Rule.
        /// </summary>
        public readonly string PrimaryKey;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The Secondary Connection String for the Event Hubs authorization Rule.
        /// </summary>
        public readonly string SecondaryConnectionString;
        /// <summary>
        /// The alias of the Secondary Connection String for the Event Hubs authorization Rule.
        /// </summary>
        public readonly string SecondaryConnectionStringAlias;
        /// <summary>
        /// The Secondary Key for the Event Hubs authorization Rule.
        /// </summary>
        public readonly string SecondaryKey;
        /// <summary>
        /// Does this Authorization Rule have permissions to Send to the Event Hub?
        /// </summary>
        public readonly bool Send;

        [OutputConstructor]
        private GetNamespaceAuthorizationRuleResult(
            string id,

            bool listen,

            bool manage,

            string name,

            string namespaceName,

            string primaryConnectionString,

            string primaryConnectionStringAlias,

            string primaryKey,

            string resourceGroupName,

            string secondaryConnectionString,

            string secondaryConnectionStringAlias,

            string secondaryKey,

            bool send)
        {
            Id = id;
            Listen = listen;
            Manage = manage;
            Name = name;
            NamespaceName = namespaceName;
            PrimaryConnectionString = primaryConnectionString;
            PrimaryConnectionStringAlias = primaryConnectionStringAlias;
            PrimaryKey = primaryKey;
            ResourceGroupName = resourceGroupName;
            SecondaryConnectionString = secondaryConnectionString;
            SecondaryConnectionStringAlias = secondaryConnectionStringAlias;
            SecondaryKey = secondaryKey;
            Send = send;
        }
    }
}

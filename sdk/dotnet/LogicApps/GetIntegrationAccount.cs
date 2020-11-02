// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.LogicApps
{
    public static class GetIntegrationAccount
    {
        /// <summary>
        /// Use this data source to access information about an existing Logic App Integration Account.
        /// 
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
        ///         var example = Output.Create(Azure.LogicApps.GetIntegrationAccount.InvokeAsync(new Azure.LogicApps.GetIntegrationAccountArgs
        ///         {
        ///             Name = "example-account",
        ///             ResourceGroupName = "example-resource-group",
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
        public static Task<GetIntegrationAccountResult> InvokeAsync(GetIntegrationAccountArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIntegrationAccountResult>("azure:logicapps/getIntegrationAccount:getIntegrationAccount", args ?? new GetIntegrationAccountArgs(), options.WithVersion());
    }


    public sealed class GetIntegrationAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this Logic App Integration Account.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Logic App Integration Account exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetIntegrationAccountArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIntegrationAccountResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure Region where the Logic App Integration Account exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The sku name of the Logic App Integration Account.
        /// </summary>
        public readonly string SkuName;
        /// <summary>
        /// A mapping of tags assigned to the Logic App Integration Account.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetIntegrationAccountResult(
            string id,

            string location,

            string name,

            string resourceGroupName,

            string skuName,

            ImmutableDictionary<string, string> tags)
        {
            Id = id;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            SkuName = skuName;
            Tags = tags;
        }
    }
}

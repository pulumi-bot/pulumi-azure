// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement
{
    public static class GetGroup
    {
        /// <summary>
        /// Use this data source to access information about an existing API Management Group.
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
        ///         var example = Output.Create(Azure.ApiManagement.GetGroup.InvokeAsync(new Azure.ApiManagement.GetGroupArgs
        ///         {
        ///             Name = "my-group",
        ///             ApiManagementName = "example-apim",
        ///             ResourceGroupName = "search-service",
        ///         }));
        ///         this.GroupType = example.Apply(example =&gt; example.Type);
        ///     }
        /// 
        ///     [Output("groupType")]
        ///     public Output&lt;string&gt; GroupType { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGroupResult> InvokeAsync(GetGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("azure:apimanagement/getGroup:getGroup", args ?? new GetGroupArgs(), options.WithVersion());

        public static Output<GetGroupResult> Invoke(GetGroupOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.ApiManagementName.Box(),
                args.Name.Box(),
                args.ResourceGroupName.Box()
            ).Apply(a => {
                    var args = new GetGroupArgs();
                    a[0].Set(args, nameof(args.ApiManagementName));
                    a[1].Set(args, nameof(args.Name));
                    a[2].Set(args, nameof(args.ResourceGroupName));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Name of the API Management Service in which this Group exists.
        /// </summary>
        [Input("apiManagementName", required: true)]
        public string ApiManagementName { get; set; } = null!;

        /// <summary>
        /// The Name of the API Management Group.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group in which the API Management Service exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetGroupArgs()
        {
        }
    }

    public sealed class GetGroupOutputArgs
    {
        /// <summary>
        /// The Name of the API Management Service in which this Group exists.
        /// </summary>
        [Input("apiManagementName", required: true)]
        public Input<string> ApiManagementName { get; set; } = null!;

        /// <summary>
        /// The Name of the API Management Group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group in which the API Management Service exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetGroupOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGroupResult
    {
        public readonly string ApiManagementName;
        /// <summary>
        /// The description of this API Management Group.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The display name of this API Management Group.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The identifier of the external Group.
        /// </summary>
        public readonly string ExternalId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The type of this API Management Group, such as `custom` or `external`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetGroupResult(
            string apiManagementName,

            string description,

            string displayName,

            string externalId,

            string id,

            string name,

            string resourceGroupName,

            string type)
        {
            ApiManagementName = apiManagementName;
            Description = description;
            DisplayName = displayName;
            ExternalId = externalId;
            Id = id;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Type = type;
        }
    }
}

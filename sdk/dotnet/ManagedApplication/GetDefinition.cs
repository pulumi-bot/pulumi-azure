// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ManagedApplication
{
    public static class GetDefinition
    {
        /// <summary>
        /// Uses this data source to access information about an existing Managed Application Definition.
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
        ///         var example = Output.Create(Azure.ManagedApplication.GetDefinition.InvokeAsync(new Azure.ManagedApplication.GetDefinitionArgs
        ///         {
        ///             Name = "example-managedappdef",
        ///             ResourceGroupName = "example-resources",
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
        public static Task<GetDefinitionResult> InvokeAsync(GetDefinitionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDefinitionResult>("azure:managedapplication/getDefinition:getDefinition", args ?? new GetDefinitionArgs(), options.WithVersion());

        public static Output<GetDefinitionResult> Apply(GetDefinitionApplyArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Name.Box(),
                args.ResourceGroupName.Box()
            ).Apply(a => {
                    var args = new GetDefinitionArgs();
                    a[0].Set(args, nameof(args.Name));
                    a[1].Set(args, nameof(args.ResourceGroupName));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetDefinitionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Managed Application Definition.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group where this Managed Application Definition exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDefinitionArgs()
        {
        }
    }

    public sealed class GetDefinitionApplyArgs
    {
        /// <summary>
        /// Specifies the name of the Managed Application Definition.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group where this Managed Application Definition exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetDefinitionApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDefinitionResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;

        [OutputConstructor]
        private GetDefinitionResult(
            string id,

            string location,

            string name,

            string resourceGroupName)
        {
            Id = id;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
        }
    }
}

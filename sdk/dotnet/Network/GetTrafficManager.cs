// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    public static class GetTrafficManager
    {
        /// <summary>
        /// Use this data source to access the ID of a specified Traffic Manager Geographical Location within the Geographical Hierarchy.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### World)
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.Network.GetTrafficManager.InvokeAsync(new Azure.Network.GetTrafficManagerArgs
        ///         {
        ///             Name = "World",
        ///         }));
        ///         this.LocationCode = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("locationCode")]
        ///     public Output&lt;string&gt; LocationCode { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTrafficManagerResult> InvokeAsync(GetTrafficManagerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTrafficManagerResult>("azure:network/getTrafficManager:getTrafficManager", args ?? new GetTrafficManagerArgs(), options.WithVersion());

        public static Output<GetTrafficManagerResult> Apply(GetTrafficManagerApplyArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Name.Box()
            ).Apply(a => {
                    var args = new GetTrafficManagerArgs();
                    a[0].Set(args, nameof(args.Name));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetTrafficManagerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Location, for example `World`, `Europe` or `Germany`.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetTrafficManagerArgs()
        {
        }
    }

    public sealed class GetTrafficManagerApplyArgs
    {
        /// <summary>
        /// Specifies the name of the Location, for example `World`, `Europe` or `Germany`.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetTrafficManagerApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTrafficManagerResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetTrafficManagerResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}

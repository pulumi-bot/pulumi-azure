// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.TrafficManager
{
    [Obsolete(@"azure.trafficmanager.getGeographicalLocation has been deprecated in favor of azure.network.getTrafficManager")]
    public static class GetGeographicalLocation
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
        public static Task<GetGeographicalLocationResult> InvokeAsync(GetGeographicalLocationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGeographicalLocationResult>("azure:trafficmanager/getGeographicalLocation:getGeographicalLocation", args ?? new GetGeographicalLocationArgs(), options.WithVersion());

        public static Output<GetGeographicalLocationResult> Invoke(GetGeographicalLocationOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Name.Box()
            ).Apply(a => {
                    var args = new GetGeographicalLocationArgs();
                    a[0].Set(args, nameof(args.Name));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetGeographicalLocationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Location, for example `World`, `Europe` or `Germany`.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetGeographicalLocationArgs()
        {
        }
    }

    public sealed class GetGeographicalLocationOutputArgs
    {
        /// <summary>
        /// Specifies the name of the Location, for example `World`, `Europe` or `Germany`.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetGeographicalLocationOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGeographicalLocationResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetGeographicalLocationResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}

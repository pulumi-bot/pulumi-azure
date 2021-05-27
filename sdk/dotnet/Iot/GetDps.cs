// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Iot
{
    public static class GetDps
    {
        /// <summary>
        /// Use this data source to access information about an existing IotHub Device Provisioning Service.
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
        ///         var example = Output.Create(Azure.Iot.GetDps.InvokeAsync(new Azure.Iot.GetDpsArgs
        ///         {
        ///             Name = "iot_hub_dps_test",
        ///             ResourceGroupName = "iothub_dps_rg",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDpsResult> InvokeAsync(GetDpsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDpsResult>("azure:iot/getDps:getDps", args ?? new GetDpsArgs(), options.WithVersion());

        public static Output<GetDpsResult> Apply(GetDpsApplyArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Name.Box(),
                args.ResourceGroupName.Box(),
                args.Tags.ToDict().Box()
            ).Apply(a => {
                    var args = new GetDpsArgs();
                    a[0].Set(args, nameof(args.Name));
                    a[1].Set(args, nameof(args.ResourceGroupName));
                    a[2].Set(args, nameof(args.Tags));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetDpsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Iot Device Provisioning Service resource.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group under which the Iot Device Provisioning Service is located in.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetDpsArgs()
        {
        }
    }

    public sealed class GetDpsApplyArgs
    {
        /// <summary>
        /// Specifies the name of the Iot Device Provisioning Service resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group under which the Iot Device Provisioning Service is located in.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetDpsApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDpsResult
    {
        /// <summary>
        /// The allocation policy of the IoT Device Provisioning Service.
        /// </summary>
        public readonly string AllocationPolicy;
        /// <summary>
        /// The device endpoint of the IoT Device Provisioning Service.
        /// </summary>
        public readonly string DeviceProvisioningHostName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The unique identifier of the IoT Device Provisioning Service.
        /// </summary>
        public readonly string IdScope;
        /// <summary>
        /// Specifies the supported Azure location where the IoT Device Provisioning Service exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The service endpoint of the IoT Device Provisioning Service.
        /// </summary>
        public readonly string ServiceOperationsHostName;
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private GetDpsResult(
            string allocationPolicy,

            string deviceProvisioningHostName,

            string id,

            string idScope,

            string location,

            string name,

            string resourceGroupName,

            string serviceOperationsHostName,

            ImmutableDictionary<string, string>? tags)
        {
            AllocationPolicy = allocationPolicy;
            DeviceProvisioningHostName = deviceProvisioningHostName;
            Id = id;
            IdScope = idScope;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            ServiceOperationsHostName = serviceOperationsHostName;
            Tags = tags;
        }
    }
}

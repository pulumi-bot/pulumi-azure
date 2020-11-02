// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Iot
{
    /// <summary>
    /// Manages an IotHub Device Provisioning Service Certificate.
    /// 
    /// ## Import
    /// 
    /// IoTHub Device Provisioning Service Certificates can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class IotHubCertificate : Pulumi.CustomResource
    {
        /// <summary>
        /// The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
        /// </summary>
        [Output("certificateContent")]
        public Output<string> CertificateContent { get; private set; } = null!;

        /// <summary>
        /// The name of the IoT Device Provisioning Service that this certificate will be attached to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("iotDpsName")]
        public Output<string> IotDpsName { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Iot Device Provisioning Service Certificate resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group under which the Iot Device Provisioning Service Certificate resource has to be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a IotHubCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IotHubCertificate(string name, IotHubCertificateArgs args, CustomResourceOptions? options = null)
            : base("azure:iot/iotHubCertificate:IotHubCertificate", name, args ?? new IotHubCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IotHubCertificate(string name, Input<string> id, IotHubCertificateState? state = null, CustomResourceOptions? options = null)
            : base("azure:iot/iotHubCertificate:IotHubCertificate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IotHubCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IotHubCertificate Get(string name, Input<string> id, IotHubCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new IotHubCertificate(name, id, state, options);
        }
    }

    public sealed class IotHubCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
        /// </summary>
        [Input("certificateContent", required: true)]
        public Input<string> CertificateContent { get; set; } = null!;

        /// <summary>
        /// The name of the IoT Device Provisioning Service that this certificate will be attached to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("iotDpsName", required: true)]
        public Input<string> IotDpsName { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Iot Device Provisioning Service Certificate resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group under which the Iot Device Provisioning Service Certificate resource has to be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public IotHubCertificateArgs()
        {
        }
    }

    public sealed class IotHubCertificateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
        /// </summary>
        [Input("certificateContent")]
        public Input<string>? CertificateContent { get; set; }

        /// <summary>
        /// The name of the IoT Device Provisioning Service that this certificate will be attached to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("iotDpsName")]
        public Input<string>? IotDpsName { get; set; }

        /// <summary>
        /// Specifies the name of the Iot Device Provisioning Service Certificate resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group under which the Iot Device Provisioning Service Certificate resource has to be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        public IotHubCertificateState()
        {
        }
    }
}

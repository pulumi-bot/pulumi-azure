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
    /// Manages an IotHub Device Provisioning Service Shared Access Policy
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var exampleIotHubDps = new Azure.Iot.IotHubDps("exampleIotHubDps", new Azure.Iot.IotHubDpsArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Sku = new Azure.Iot.Inputs.IotHubDpsSkuArgs
    ///             {
    ///                 Name = "S1",
    ///                 Capacity = 1,
    ///             },
    ///         });
    ///         var exampleDpsSharedAccessPolicy = new Azure.Iot.DpsSharedAccessPolicy("exampleDpsSharedAccessPolicy", new Azure.Iot.DpsSharedAccessPolicyArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             IothubDpsName = exampleIotHubDps.Name,
    ///             EnrollmentWrite = true,
    ///             EnrollmentRead = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class DpsSharedAccessPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Adds `EnrollmentRead` permission to this Shared Access Account. It allows read access to enrollment data.
        /// </summary>
        [Output("enrollmentRead")]
        public Output<bool?> EnrollmentRead { get; private set; } = null!;

        /// <summary>
        /// Adds `EnrollmentWrite` permission to this Shared Access Account. It allows write access to enrollment data.
        /// </summary>
        [Output("enrollmentWrite")]
        public Output<bool?> EnrollmentWrite { get; private set; } = null!;

        /// <summary>
        /// The name of the IoT Hub Device Provisioning service to which this Shared Access Policy belongs. Changing this forces a new resource to be created.
        /// </summary>
        [Output("iothubDpsName")]
        public Output<string> IothubDpsName { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the IotHub Shared Access Policy resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The primary connection string of the Shared Access Policy.
        /// </summary>
        [Output("primaryConnectionString")]
        public Output<string> PrimaryConnectionString { get; private set; } = null!;

        /// <summary>
        /// The primary key used to create the authentication token.
        /// </summary>
        [Output("primaryKey")]
        public Output<string> PrimaryKey { get; private set; } = null!;

        /// <summary>
        /// Adds `RegistrationStatusRead` permission to this Shared Access Account. It allows read access to device registrations.
        /// </summary>
        [Output("registrationRead")]
        public Output<bool?> RegistrationRead { get; private set; } = null!;

        /// <summary>
        /// Adds `RegistrationStatusWrite` permission to this Shared Access Account. It allows write access to device registrations.
        /// </summary>
        [Output("registrationWrite")]
        public Output<bool?> RegistrationWrite { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group under which the IotHub Shared Access Policy resource has to be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The secondary connection string of the Shared Access Policy.
        /// </summary>
        [Output("secondaryConnectionString")]
        public Output<string> SecondaryConnectionString { get; private set; } = null!;

        /// <summary>
        /// The secondary key used to create the authentication token.
        /// </summary>
        [Output("secondaryKey")]
        public Output<string> SecondaryKey { get; private set; } = null!;

        /// <summary>
        /// Adds `ServiceConfig` permission to this Shared Access Account. It allows configuration of the Device Provisioning Service.
        /// </summary>
        [Output("serviceConfig")]
        public Output<bool?> ServiceConfig { get; private set; } = null!;


        /// <summary>
        /// Create a DpsSharedAccessPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DpsSharedAccessPolicy(string name, DpsSharedAccessPolicyArgs args, CustomResourceOptions? options = null)
            : base("azure:iot/dpsSharedAccessPolicy:DpsSharedAccessPolicy", name, args ?? new DpsSharedAccessPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DpsSharedAccessPolicy(string name, Input<string> id, DpsSharedAccessPolicyState? state = null, CustomResourceOptions? options = null)
            : base("azure:iot/dpsSharedAccessPolicy:DpsSharedAccessPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DpsSharedAccessPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DpsSharedAccessPolicy Get(string name, Input<string> id, DpsSharedAccessPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new DpsSharedAccessPolicy(name, id, state, options);
        }
    }

    public sealed class DpsSharedAccessPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Adds `EnrollmentRead` permission to this Shared Access Account. It allows read access to enrollment data.
        /// </summary>
        [Input("enrollmentRead")]
        public Input<bool>? EnrollmentRead { get; set; }

        /// <summary>
        /// Adds `EnrollmentWrite` permission to this Shared Access Account. It allows write access to enrollment data.
        /// </summary>
        [Input("enrollmentWrite")]
        public Input<bool>? EnrollmentWrite { get; set; }

        /// <summary>
        /// The name of the IoT Hub Device Provisioning service to which this Shared Access Policy belongs. Changing this forces a new resource to be created.
        /// </summary>
        [Input("iothubDpsName", required: true)]
        public Input<string> IothubDpsName { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the IotHub Shared Access Policy resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Adds `RegistrationStatusRead` permission to this Shared Access Account. It allows read access to device registrations.
        /// </summary>
        [Input("registrationRead")]
        public Input<bool>? RegistrationRead { get; set; }

        /// <summary>
        /// Adds `RegistrationStatusWrite` permission to this Shared Access Account. It allows write access to device registrations.
        /// </summary>
        [Input("registrationWrite")]
        public Input<bool>? RegistrationWrite { get; set; }

        /// <summary>
        /// The name of the resource group under which the IotHub Shared Access Policy resource has to be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Adds `ServiceConfig` permission to this Shared Access Account. It allows configuration of the Device Provisioning Service.
        /// </summary>
        [Input("serviceConfig")]
        public Input<bool>? ServiceConfig { get; set; }

        public DpsSharedAccessPolicyArgs()
        {
        }
    }

    public sealed class DpsSharedAccessPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Adds `EnrollmentRead` permission to this Shared Access Account. It allows read access to enrollment data.
        /// </summary>
        [Input("enrollmentRead")]
        public Input<bool>? EnrollmentRead { get; set; }

        /// <summary>
        /// Adds `EnrollmentWrite` permission to this Shared Access Account. It allows write access to enrollment data.
        /// </summary>
        [Input("enrollmentWrite")]
        public Input<bool>? EnrollmentWrite { get; set; }

        /// <summary>
        /// The name of the IoT Hub Device Provisioning service to which this Shared Access Policy belongs. Changing this forces a new resource to be created.
        /// </summary>
        [Input("iothubDpsName")]
        public Input<string>? IothubDpsName { get; set; }

        /// <summary>
        /// Specifies the name of the IotHub Shared Access Policy resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The primary connection string of the Shared Access Policy.
        /// </summary>
        [Input("primaryConnectionString")]
        public Input<string>? PrimaryConnectionString { get; set; }

        /// <summary>
        /// The primary key used to create the authentication token.
        /// </summary>
        [Input("primaryKey")]
        public Input<string>? PrimaryKey { get; set; }

        /// <summary>
        /// Adds `RegistrationStatusRead` permission to this Shared Access Account. It allows read access to device registrations.
        /// </summary>
        [Input("registrationRead")]
        public Input<bool>? RegistrationRead { get; set; }

        /// <summary>
        /// Adds `RegistrationStatusWrite` permission to this Shared Access Account. It allows write access to device registrations.
        /// </summary>
        [Input("registrationWrite")]
        public Input<bool>? RegistrationWrite { get; set; }

        /// <summary>
        /// The name of the resource group under which the IotHub Shared Access Policy resource has to be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The secondary connection string of the Shared Access Policy.
        /// </summary>
        [Input("secondaryConnectionString")]
        public Input<string>? SecondaryConnectionString { get; set; }

        /// <summary>
        /// The secondary key used to create the authentication token.
        /// </summary>
        [Input("secondaryKey")]
        public Input<string>? SecondaryKey { get; set; }

        /// <summary>
        /// Adds `ServiceConfig` permission to this Shared Access Account. It allows configuration of the Device Provisioning Service.
        /// </summary>
        [Input("serviceConfig")]
        public Input<bool>? ServiceConfig { get; set; }

        public DpsSharedAccessPolicyState()
        {
        }
    }
}

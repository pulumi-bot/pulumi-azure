// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppPlatform
{
    /// <summary>
    /// Manage an Azure Spring Cloud Application.
    /// </summary>
    public partial class SpringCloudApp : Pulumi.CustomResource
    {
        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.SpringCloudAppIdentity?> Identity { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Spring Cloud Application. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group in which to create the Spring Cloud Application. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;


        /// <summary>
        /// Create a SpringCloudApp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SpringCloudApp(string name, SpringCloudAppArgs args, CustomResourceOptions? options = null)
            : base("azure:appplatform/springCloudApp:SpringCloudApp", name, args ?? new SpringCloudAppArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SpringCloudApp(string name, Input<string> id, SpringCloudAppState? state = null, CustomResourceOptions? options = null)
            : base("azure:appplatform/springCloudApp:SpringCloudApp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SpringCloudApp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SpringCloudApp Get(string name, Input<string> id, SpringCloudAppState? state = null, CustomResourceOptions? options = null)
        {
            return new SpringCloudApp(name, id, state, options);
        }
    }

    public sealed class SpringCloudAppArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.SpringCloudAppIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Specifies the name of the Spring Cloud Application. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the resource group in which to create the Spring Cloud Application. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public SpringCloudAppArgs()
        {
        }
    }

    public sealed class SpringCloudAppState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.SpringCloudAppIdentityGetArgs>? Identity { get; set; }

        /// <summary>
        /// Specifies the name of the Spring Cloud Application. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the resource group in which to create the Spring Cloud Application. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public SpringCloudAppState()
        {
        }
    }
}

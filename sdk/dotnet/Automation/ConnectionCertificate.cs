// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Automation
{
    /// <summary>
    /// Manages an Automation Connection with type `Azure`.
    /// 
    /// ## Import
    /// 
    /// Automation Connection can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class ConnectionCertificate : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("automationAccountName")]
        public Output<string> AutomationAccountName { get; private set; } = null!;

        /// <summary>
        /// The name of the automation certificate.
        /// </summary>
        [Output("automationCertificateName")]
        public Output<string> AutomationCertificateName { get; private set; } = null!;

        /// <summary>
        /// A description for this Connection.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Connection. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The id of subscription where the automation certificate exists.
        /// </summary>
        [Output("subscriptionId")]
        public Output<string> SubscriptionId { get; private set; } = null!;


        /// <summary>
        /// Create a ConnectionCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConnectionCertificate(string name, ConnectionCertificateArgs args, CustomResourceOptions? options = null)
            : base("azure:automation/connectionCertificate:ConnectionCertificate", name, args ?? new ConnectionCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConnectionCertificate(string name, Input<string> id, ConnectionCertificateState? state = null, CustomResourceOptions? options = null)
            : base("azure:automation/connectionCertificate:ConnectionCertificate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConnectionCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConnectionCertificate Get(string name, Input<string> id, ConnectionCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new ConnectionCertificate(name, id, state, options);
        }
    }

    public sealed class ConnectionCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public Input<string> AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// The name of the automation certificate.
        /// </summary>
        [Input("automationCertificateName", required: true)]
        public Input<string> AutomationCertificateName { get; set; } = null!;

        /// <summary>
        /// A description for this Connection.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the name of the Connection. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The id of subscription where the automation certificate exists.
        /// </summary>
        [Input("subscriptionId", required: true)]
        public Input<string> SubscriptionId { get; set; } = null!;

        public ConnectionCertificateArgs()
        {
        }
    }

    public sealed class ConnectionCertificateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("automationAccountName")]
        public Input<string>? AutomationAccountName { get; set; }

        /// <summary>
        /// The name of the automation certificate.
        /// </summary>
        [Input("automationCertificateName")]
        public Input<string>? AutomationCertificateName { get; set; }

        /// <summary>
        /// A description for this Connection.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the name of the Connection. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The id of subscription where the automation certificate exists.
        /// </summary>
        [Input("subscriptionId")]
        public Input<string>? SubscriptionId { get; set; }

        public ConnectionCertificateState()
        {
        }
    }
}

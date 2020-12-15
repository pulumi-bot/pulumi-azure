// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement
{
    /// <summary>
    /// Manages a API Management service Policy.
    /// 
    /// &gt; **NOTE:** This resource will, upon creation, **overwrite any existing policy in the API Management service**, as there is no feasible way to test whether the policy has been modified from the default. Similarly, when this resource is destroyed, the API Management service will revert to its default policy.
    /// 
    /// ## Import
    /// 
    /// API Management service Policys can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:apimanagement/policy:Policy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/instance1/policies/policy
    /// ```
    /// </summary>
    [AzureResourceType("azure:apimanagement/policy:Policy")]
    public partial class Policy : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the API Management service. Changing this forces a new API Management service Policy to be created.
        /// </summary>
        [Output("apiManagementId")]
        public Output<string> ApiManagementId { get; private set; } = null!;

        [Output("xmlContent")]
        public Output<string> XmlContent { get; private set; } = null!;

        /// <summary>
        /// A link to a Policy XML Document, which must be publicly available.
        /// </summary>
        [Output("xmlLink")]
        public Output<string?> XmlLink { get; private set; } = null!;


        /// <summary>
        /// Create a Policy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Policy(string name, PolicyArgs args, CustomResourceOptions? options = null)
            : base("azure:apimanagement/policy:Policy", name, args ?? new PolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Policy(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
            : base("azure:apimanagement/policy:Policy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Policy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Policy Get(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new Policy(name, id, state, options);
        }
    }

    public sealed class PolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the API Management service. Changing this forces a new API Management service Policy to be created.
        /// </summary>
        [Input("apiManagementId", required: true)]
        public Input<string> ApiManagementId { get; set; } = null!;

        [Input("xmlContent")]
        public Input<string>? XmlContent { get; set; }

        /// <summary>
        /// A link to a Policy XML Document, which must be publicly available.
        /// </summary>
        [Input("xmlLink")]
        public Input<string>? XmlLink { get; set; }

        public PolicyArgs()
        {
        }
    }

    public sealed class PolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the API Management service. Changing this forces a new API Management service Policy to be created.
        /// </summary>
        [Input("apiManagementId")]
        public Input<string>? ApiManagementId { get; set; }

        [Input("xmlContent")]
        public Input<string>? XmlContent { get; set; }

        /// <summary>
        /// A link to a Policy XML Document, which must be publicly available.
        /// </summary>
        [Input("xmlLink")]
        public Input<string>? XmlLink { get; set; }

        public PolicyState()
        {
        }
    }
}

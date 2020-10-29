// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Policy
{
    /// <summary>
    /// Manages an Azure Policy Remediation at the specified Scope.
    /// </summary>
    public partial class Remediation : Pulumi.CustomResource
    {
        /// <summary>
        /// A list of the resource locations that will be remediated.
        /// </summary>
        [Output("locationFilters")]
        public Output<ImmutableArray<string>> LocationFilters { get; private set; } = null!;

        /// <summary>
        /// The name of the Policy Remediation. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the Policy Assignment that should be remediated.
        /// </summary>
        [Output("policyAssignmentId")]
        public Output<string> PolicyAssignmentId { get; private set; } = null!;

        /// <summary>
        /// The unique ID for the policy definition within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
        /// </summary>
        [Output("policyDefinitionReferenceId")]
        public Output<string?> PolicyDefinitionReferenceId { get; private set; } = null!;

        /// <summary>
        /// The Scope at which the Policy Remediation should be applied. Changing this forces a new resource to be created. A scope must be a Resource ID out of one of the following list:
        /// </summary>
        [Output("scope")]
        public Output<string> Scope { get; private set; } = null!;


        /// <summary>
        /// Create a Remediation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Remediation(string name, RemediationArgs args, CustomResourceOptions? options = null)
            : base("azure:policy/remediation:Remediation", name, args ?? new RemediationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Remediation(string name, Input<string> id, RemediationState? state = null, CustomResourceOptions? options = null)
            : base("azure:policy/remediation:Remediation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Remediation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Remediation Get(string name, Input<string> id, RemediationState? state = null, CustomResourceOptions? options = null)
        {
            return new Remediation(name, id, state, options);
        }
    }

    public sealed class RemediationArgs : Pulumi.ResourceArgs
    {
        [Input("locationFilters")]
        private InputList<string>? _locationFilters;

        /// <summary>
        /// A list of the resource locations that will be remediated.
        /// </summary>
        public InputList<string> LocationFilters
        {
            get => _locationFilters ?? (_locationFilters = new InputList<string>());
            set => _locationFilters = value;
        }

        /// <summary>
        /// The name of the Policy Remediation. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the Policy Assignment that should be remediated.
        /// </summary>
        [Input("policyAssignmentId", required: true)]
        public Input<string> PolicyAssignmentId { get; set; } = null!;

        /// <summary>
        /// The unique ID for the policy definition within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
        /// </summary>
        [Input("policyDefinitionReferenceId")]
        public Input<string>? PolicyDefinitionReferenceId { get; set; }

        /// <summary>
        /// The Scope at which the Policy Remediation should be applied. Changing this forces a new resource to be created. A scope must be a Resource ID out of one of the following list:
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        public RemediationArgs()
        {
        }
    }

    public sealed class RemediationState : Pulumi.ResourceArgs
    {
        [Input("locationFilters")]
        private InputList<string>? _locationFilters;

        /// <summary>
        /// A list of the resource locations that will be remediated.
        /// </summary>
        public InputList<string> LocationFilters
        {
            get => _locationFilters ?? (_locationFilters = new InputList<string>());
            set => _locationFilters = value;
        }

        /// <summary>
        /// The name of the Policy Remediation. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the Policy Assignment that should be remediated.
        /// </summary>
        [Input("policyAssignmentId")]
        public Input<string>? PolicyAssignmentId { get; set; }

        /// <summary>
        /// The unique ID for the policy definition within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
        /// </summary>
        [Input("policyDefinitionReferenceId")]
        public Input<string>? PolicyDefinitionReferenceId { get; set; }

        /// <summary>
        /// The Scope at which the Policy Remediation should be applied. Changing this forces a new resource to be created. A scope must be a Resource ID out of one of the following list:
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        public RemediationState()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.SecurityCenter
{
    /// <summary>
    /// Manages the Security Center Assessment Metadata for Azure Security Center.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Azure.SecurityCenter.AssessmentMetadata("example", new Azure.SecurityCenter.AssessmentMetadataArgs
    ///         {
    ///             Description = "Test Description",
    ///             DisplayName = "Test Display Name",
    ///             Severity = "Medium",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Security Assessments Metadata can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:securitycenter/assessmentMetadata:AssessmentMetadata example /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Security/assessmentMetadata/metadata1
    /// ```
    /// </summary>
    [AzureResourceType("azure:securitycenter/assessmentMetadata:AssessmentMetadata")]
    public partial class AssessmentMetadata : Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the Security Center Assessment.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The user-friendly display name of the Security Center Assessment.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The implementation effort which is used to remediate the Security Center Assessment. Possible values are `Low`, `Moderate` and `High`.
        /// </summary>
        [Output("implementationEffort")]
        public Output<string?> ImplementationEffort { get; private set; } = null!;

        /// <summary>
        /// The GUID as the name of the Security Center Assessment Metadata.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The description which is used to mitigate the security issue.
        /// </summary>
        [Output("remediationDescription")]
        public Output<string?> RemediationDescription { get; private set; } = null!;

        /// <summary>
        /// The severity level of the Security Center Assessment. Possible values are `Low`, `Medium` and `High`. Defaults to `Medium`.
        /// </summary>
        [Output("severity")]
        public Output<string?> Severity { get; private set; } = null!;

        /// <summary>
        /// A list of the threat impacts for the Security Center Assessment. Possible values are `AccountBreach`, `DataExfiltration`, `DataSpillage`, `DenialOfService`, `ElevationOfPrivilege`, `MaliciousInsider`, `MissingCoverage` and `ThreatResistance`.
        /// </summary>
        [Output("threats")]
        public Output<ImmutableArray<string>> Threats { get; private set; } = null!;

        /// <summary>
        /// The user impact of the Security Center Assessment. Possible values are `Low`, `Moderate` and `High`.
        /// </summary>
        [Output("userImpact")]
        public Output<string?> UserImpact { get; private set; } = null!;


        /// <summary>
        /// Create a AssessmentMetadata resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AssessmentMetadata(string name, AssessmentMetadataArgs args, CustomResourceOptions? options = null)
            : base("azure:securitycenter/assessmentMetadata:AssessmentMetadata", name, args ?? new AssessmentMetadataArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AssessmentMetadata(string name, Input<string> id, AssessmentMetadataState? state = null, CustomResourceOptions? options = null)
            : base("azure:securitycenter/assessmentMetadata:AssessmentMetadata", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AssessmentMetadata resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AssessmentMetadata Get(string name, Input<string> id, AssessmentMetadataState? state = null, CustomResourceOptions? options = null)
        {
            return new AssessmentMetadata(name, id, state, options);
        }
    }

    public sealed class AssessmentMetadataArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the Security Center Assessment.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The user-friendly display name of the Security Center Assessment.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The implementation effort which is used to remediate the Security Center Assessment. Possible values are `Low`, `Moderate` and `High`.
        /// </summary>
        [Input("implementationEffort")]
        public Input<string>? ImplementationEffort { get; set; }

        /// <summary>
        /// The description which is used to mitigate the security issue.
        /// </summary>
        [Input("remediationDescription")]
        public Input<string>? RemediationDescription { get; set; }

        /// <summary>
        /// The severity level of the Security Center Assessment. Possible values are `Low`, `Medium` and `High`. Defaults to `Medium`.
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        [Input("threats")]
        private InputList<string>? _threats;

        /// <summary>
        /// A list of the threat impacts for the Security Center Assessment. Possible values are `AccountBreach`, `DataExfiltration`, `DataSpillage`, `DenialOfService`, `ElevationOfPrivilege`, `MaliciousInsider`, `MissingCoverage` and `ThreatResistance`.
        /// </summary>
        public InputList<string> Threats
        {
            get => _threats ?? (_threats = new InputList<string>());
            set => _threats = value;
        }

        /// <summary>
        /// The user impact of the Security Center Assessment. Possible values are `Low`, `Moderate` and `High`.
        /// </summary>
        [Input("userImpact")]
        public Input<string>? UserImpact { get; set; }

        public AssessmentMetadataArgs()
        {
        }
    }

    public sealed class AssessmentMetadataState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the Security Center Assessment.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The user-friendly display name of the Security Center Assessment.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The implementation effort which is used to remediate the Security Center Assessment. Possible values are `Low`, `Moderate` and `High`.
        /// </summary>
        [Input("implementationEffort")]
        public Input<string>? ImplementationEffort { get; set; }

        /// <summary>
        /// The GUID as the name of the Security Center Assessment Metadata.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The description which is used to mitigate the security issue.
        /// </summary>
        [Input("remediationDescription")]
        public Input<string>? RemediationDescription { get; set; }

        /// <summary>
        /// The severity level of the Security Center Assessment. Possible values are `Low`, `Medium` and `High`. Defaults to `Medium`.
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        [Input("threats")]
        private InputList<string>? _threats;

        /// <summary>
        /// A list of the threat impacts for the Security Center Assessment. Possible values are `AccountBreach`, `DataExfiltration`, `DataSpillage`, `DenialOfService`, `ElevationOfPrivilege`, `MaliciousInsider`, `MissingCoverage` and `ThreatResistance`.
        /// </summary>
        public InputList<string> Threats
        {
            get => _threats ?? (_threats = new InputList<string>());
            set => _threats = value;
        }

        /// <summary>
        /// The user impact of the Security Center Assessment. Possible values are `Low`, `Moderate` and `High`.
        /// </summary>
        [Input("userImpact")]
        public Input<string>? UserImpact { get; set; }

        public AssessmentMetadataState()
        {
        }
    }
}
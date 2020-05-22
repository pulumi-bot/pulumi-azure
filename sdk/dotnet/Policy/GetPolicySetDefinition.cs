// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Policy
{
    public static class GetPolicySetDefinition
    {
        /// <summary>
        /// Use this data source to access information about an existing Policy Set Definition.
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
        ///         var example = Output.Create(Azure.Policy.GetPolicySetDefinition.InvokeAsync(new Azure.Policy.GetPolicySetDefinitionArgs
        ///         {
        ///             DisplayName = "Policy Set Definition Example",
        ///         }));
        ///         this.Id = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPolicySetDefinitionResult> InvokeAsync(GetPolicySetDefinitionArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicySetDefinitionResult>("azure:policy/getPolicySetDefinition:getPolicySetDefinition", args ?? new GetPolicySetDefinitionArgs(), options.WithVersion());
    }


    public sealed class GetPolicySetDefinitionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the display name of the Policy Set Definition. Conflicts with `name`.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        /// <summary>
        /// Only retrieve Policy Set Definitions from this Management Group.
        /// </summary>
        [Input("managementGroupName")]
        public string? ManagementGroupName { get; set; }

        /// <summary>
        /// Specifies the name of the Policy Set Definition. Conflicts with `display_name`.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetPolicySetDefinitionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPolicySetDefinitionResult
    {
        /// <summary>
        /// The Description of the Policy Set Definition.
        /// </summary>
        public readonly string Description;
        public readonly string DisplayName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ManagementGroupName;
        /// <summary>
        /// Any Metadata defined in the Policy Set Definition.
        /// </summary>
        public readonly string Metadata;
        public readonly string Name;
        /// <summary>
        /// Any Parameters defined in the Policy Set Definition.
        /// </summary>
        public readonly string Parameters;
        /// <summary>
        /// The policy definitions contained within the policy set definition.
        /// </summary>
        public readonly string PolicyDefinitions;
        /// <summary>
        /// The Type of the Policy Set Definition.
        /// </summary>
        public readonly string PolicyType;

        [OutputConstructor]
        private GetPolicySetDefinitionResult(
            string description,

            string displayName,

            string id,

            string? managementGroupName,

            string metadata,

            string name,

            string parameters,

            string policyDefinitions,

            string policyType)
        {
            Description = description;
            DisplayName = displayName;
            Id = id;
            ManagementGroupName = managementGroupName;
            Metadata = metadata;
            Name = name;
            Parameters = parameters;
            PolicyDefinitions = policyDefinitions;
            PolicyType = policyType;
        }
    }
}

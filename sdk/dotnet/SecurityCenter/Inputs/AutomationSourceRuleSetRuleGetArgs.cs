// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.SecurityCenter.Inputs
{

    public sealed class AutomationSourceRuleSetRuleGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A value that will be compared with the value in `property_path`.
        /// </summary>
        [Input("expectedValue", required: true)]
        public Input<string> ExpectedValue { get; set; } = null!;

        /// <summary>
        /// The comparison operator to use, must be one of: `Contains`, `EndsWith`, `Equals`, `GreaterThan`, `GreaterThanOrEqualTo`, `LesserThan`, `LesserThanOrEqualTo`, `NotEquals`, `StartsWith`
        /// </summary>
        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        /// <summary>
        /// The JPath of the entity model property that should be checked.
        /// </summary>
        [Input("propertyPath", required: true)]
        public Input<string> PropertyPath { get; set; } = null!;

        /// <summary>
        /// The data type of the compared operands, must be one of: `Integer`, `String`, `Boolean` or `Number`.
        /// </summary>
        [Input("propertyType", required: true)]
        public Input<string> PropertyType { get; set; } = null!;

        public AutomationSourceRuleSetRuleGetArgs()
        {
        }
    }
}

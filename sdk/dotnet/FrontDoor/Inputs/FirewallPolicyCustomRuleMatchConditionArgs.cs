// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.FrontDoor.Inputs
{

    public sealed class FirewallPolicyCustomRuleMatchConditionArgs : Pulumi.ResourceArgs
    {
        [Input("matchValues", required: true)]
        private InputList<string>? _matchValues;

        /// <summary>
        /// Up to `100` possible values to match.
        /// </summary>
        public InputList<string> MatchValues
        {
            get => _matchValues ?? (_matchValues = new InputList<string>());
            set => _matchValues = value;
        }

        /// <summary>
        /// The request variable to compare with. Possible values are `Cookies`, `PostArgs`, `QueryString`, `RemoteAddr`, `RequestBody`, `RequestHeader`, `RequestMethod`, `RequestUri`, or `SocketAddr`.
        /// </summary>
        [Input("matchVariable", required: true)]
        public Input<string> MatchVariable { get; set; } = null!;

        /// <summary>
        /// Should the result of the condition be negated.
        /// </summary>
        [Input("negationCondition")]
        public Input<bool>? NegationCondition { get; set; }

        /// <summary>
        /// Comparison type to use for matching with the variable value. Possible values are `Any`, `BeginsWith`, `Contains`, `EndsWith`, `Equal`, `GeoMatch`, `GreaterThan`, `GreaterThanOrEqual`, `IPMatch`, `LessThan`, `LessThanOrEqual` or `RegEx`.
        /// </summary>
        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        /// <summary>
        /// Match against a specific key if the `match_variable` is `QueryString`, `PostArgs`, `RequestHeader` or `Cookies`.
        /// </summary>
        [Input("selector")]
        public Input<string>? Selector { get; set; }

        [Input("transforms")]
        private InputList<string>? _transforms;

        /// <summary>
        /// Up to `5` transforms to apply. Possible values are `Lowercase`, `RemoveNulls`, `Trim`, `Uppercase`, `URLDecode` or`URLEncode`.
        /// </summary>
        public InputList<string> Transforms
        {
            get => _transforms ?? (_transforms = new InputList<string>());
            set => _transforms = value;
        }

        public FirewallPolicyCustomRuleMatchConditionArgs()
        {
        }
    }
}

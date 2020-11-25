// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement.Inputs
{

    public sealed class ApiDiagnosticFrontendResponseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of payload bytes to log (up to 8192).
        /// </summary>
        [Input("bodyBytes")]
        public Input<int>? BodyBytes { get; set; }

        [Input("headersToLogs")]
        private InputList<string>? _headersToLogs;

        /// <summary>
        /// Specifies a list of headers to log.
        /// </summary>
        public InputList<string> HeadersToLogs
        {
            get => _headersToLogs ?? (_headersToLogs = new InputList<string>());
            set => _headersToLogs = value;
        }

        public ApiDiagnosticFrontendResponseArgs()
        {
        }
    }
}

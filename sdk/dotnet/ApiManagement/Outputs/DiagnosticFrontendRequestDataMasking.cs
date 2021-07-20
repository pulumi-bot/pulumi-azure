// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement.Outputs
{

    [OutputType]
    public sealed class DiagnosticFrontendRequestDataMasking
    {
        public readonly ImmutableArray<Outputs.DiagnosticFrontendRequestDataMaskingHeader> Headers;
        public readonly ImmutableArray<Outputs.DiagnosticFrontendRequestDataMaskingQueryParam> QueryParams;

        [OutputConstructor]
        private DiagnosticFrontendRequestDataMasking(
            ImmutableArray<Outputs.DiagnosticFrontendRequestDataMaskingHeader> headers,

            ImmutableArray<Outputs.DiagnosticFrontendRequestDataMaskingQueryParam> queryParams)
        {
            Headers = headers;
            QueryParams = queryParams;
        }
    }
}
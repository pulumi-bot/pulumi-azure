// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Storage.Inputs
{

    public sealed class AccountBlobPropertiesContainerDeleteRetentionPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the number of days that the container should be retained, between `1` and `365` days. Defaults to `7`.
        /// </summary>
        [Input("days")]
        public Input<int>? Days { get; set; }

        public AccountBlobPropertiesContainerDeleteRetentionPolicyArgs()
        {
        }
    }
}

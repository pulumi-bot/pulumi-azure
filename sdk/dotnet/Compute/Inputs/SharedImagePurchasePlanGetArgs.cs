// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class SharedImagePurchasePlanGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Purchase Plan Name for this Shared Image. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Purchase Plan Product for this Gallery Image. Changing this forces a new resource to be created.
        /// </summary>
        [Input("product")]
        public Input<string>? Product { get; set; }

        /// <summary>
        /// The Purchase Plan Publisher for this Gallery Image. Changing this forces a new resource to be created.
        /// </summary>
        [Input("publisher")]
        public Input<string>? Publisher { get; set; }

        public SharedImagePurchasePlanGetArgs()
        {
        }
    }
}

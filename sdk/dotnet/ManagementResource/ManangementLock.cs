// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ManagementResource
{
    /// <summary>
    /// Manages a Management Lock which is scoped to a Subscription, Resource Group or Resource.
    /// 
    /// ## Example Usage
    /// </summary>
    [Obsolete(@"azure.managementresource.ManangementLock has been deprecated in favor of azure.management.Lock")]
    public partial class ManangementLock : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the Level to be used for this Lock. Possible values are `CanNotDelete` and `ReadOnly`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("lockLevel")]
        public Output<string> LockLevel { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Management Lock. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies some notes about the lock. Maximum of 512 characters. Changing this forces a new resource to be created.
        /// </summary>
        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        /// <summary>
        /// Specifies the scope at which the Management Lock should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("scope")]
        public Output<string> Scope { get; private set; } = null!;


        /// <summary>
        /// Create a ManangementLock resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManangementLock(string name, ManangementLockArgs args, CustomResourceOptions? options = null)
            : base("azure:managementresource/manangementLock:ManangementLock", name, args ?? new ManangementLockArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManangementLock(string name, Input<string> id, ManangementLockState? state = null, CustomResourceOptions? options = null)
            : base("azure:managementresource/manangementLock:ManangementLock", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ManangementLock resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManangementLock Get(string name, Input<string> id, ManangementLockState? state = null, CustomResourceOptions? options = null)
        {
            return new ManangementLock(name, id, state, options);
        }
    }

    public sealed class ManangementLockArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the Level to be used for this Lock. Possible values are `CanNotDelete` and `ReadOnly`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("lockLevel", required: true)]
        public Input<string> LockLevel { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Management Lock. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies some notes about the lock. Maximum of 512 characters. Changing this forces a new resource to be created.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// Specifies the scope at which the Management Lock should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        public ManangementLockArgs()
        {
        }
    }

    public sealed class ManangementLockState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the Level to be used for this Lock. Possible values are `CanNotDelete` and `ReadOnly`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("lockLevel")]
        public Input<string>? LockLevel { get; set; }

        /// <summary>
        /// Specifies the name of the Management Lock. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies some notes about the lock. Maximum of 512 characters. Changing this forces a new resource to be created.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// Specifies the scope at which the Management Lock should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        public ManangementLockState()
        {
        }
    }
}

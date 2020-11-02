// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.PostgreSql
{
    /// <summary>
    /// Manages a Customer Managed Key for a PostgreSQL Server.
    /// 
    /// ## Import
    /// 
    /// A PostgreSQL Server Key can be imported using the `resource id` of the PostgreSQL Server Key, e.g.
    /// </summary>
    public partial class ServerKey : Pulumi.CustomResource
    {
        /// <summary>
        /// The URL to a Key Vault Key.
        /// </summary>
        [Output("keyVaultKeyId")]
        public Output<string> KeyVaultKeyId { get; private set; } = null!;

        /// <summary>
        /// The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
        /// </summary>
        [Output("serverId")]
        public Output<string> ServerId { get; private set; } = null!;


        /// <summary>
        /// Create a ServerKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerKey(string name, ServerKeyArgs args, CustomResourceOptions? options = null)
            : base("azure:postgresql/serverKey:ServerKey", name, args ?? new ServerKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerKey(string name, Input<string> id, ServerKeyState? state = null, CustomResourceOptions? options = null)
            : base("azure:postgresql/serverKey:ServerKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServerKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerKey Get(string name, Input<string> id, ServerKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new ServerKey(name, id, state, options);
        }
    }

    public sealed class ServerKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL to a Key Vault Key.
        /// </summary>
        [Input("keyVaultKeyId", required: true)]
        public Input<string> KeyVaultKeyId { get; set; } = null!;

        /// <summary>
        /// The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
        /// </summary>
        [Input("serverId", required: true)]
        public Input<string> ServerId { get; set; } = null!;

        public ServerKeyArgs()
        {
        }
    }

    public sealed class ServerKeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL to a Key Vault Key.
        /// </summary>
        [Input("keyVaultKeyId")]
        public Input<string>? KeyVaultKeyId { get; set; }

        /// <summary>
        /// The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
        /// </summary>
        [Input("serverId")]
        public Input<string>? ServerId { get; set; }

        public ServerKeyState()
        {
        }
    }
}

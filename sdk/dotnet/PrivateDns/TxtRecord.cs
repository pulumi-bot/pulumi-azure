// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.PrivateDns
{
    /// <summary>
    /// Enables you to manage DNS TXT Records within Azure Private DNS.
    /// </summary>
    public partial class TxtRecord : Pulumi.CustomResource
    {
        /// <summary>
        /// The FQDN of the DNS TXT Record.
        /// </summary>
        [Output("fqdn")]
        public Output<string> Fqdn { get; private set; } = null!;

        /// <summary>
        /// The name of the DNS TXT Record. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// One or more `record` blocks as defined below.
        /// </summary>
        [Output("records")]
        public Output<ImmutableArray<Outputs.TxtRecordRecord>> Records { get; private set; } = null!;

        /// <summary>
        /// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("ttl")]
        public Output<int> Ttl { get; private set; } = null!;

        /// <summary>
        /// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("zoneName")]
        public Output<string> ZoneName { get; private set; } = null!;


        /// <summary>
        /// Create a TxtRecord resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TxtRecord(string name, TxtRecordArgs args, CustomResourceOptions? options = null)
            : base("azure:privatedns/txtRecord:TxtRecord", name, args ?? new TxtRecordArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TxtRecord(string name, Input<string> id, TxtRecordState? state = null, CustomResourceOptions? options = null)
            : base("azure:privatedns/txtRecord:TxtRecord", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TxtRecord resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TxtRecord Get(string name, Input<string> id, TxtRecordState? state = null, CustomResourceOptions? options = null)
        {
            return new TxtRecord(name, id, state, options);
        }
    }

    public sealed class TxtRecordArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the DNS TXT Record. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("records", required: true)]
        private InputList<Inputs.TxtRecordRecordArgs>? _records;

        /// <summary>
        /// One or more `record` blocks as defined below.
        /// </summary>
        public InputList<Inputs.TxtRecordRecordArgs> Records
        {
            get => _records ?? (_records = new InputList<Inputs.TxtRecordRecordArgs>());
            set => _records = value;
        }

        /// <summary>
        /// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("ttl", required: true)]
        public Input<int> Ttl { get; set; } = null!;

        /// <summary>
        /// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("zoneName", required: true)]
        public Input<string> ZoneName { get; set; } = null!;

        public TxtRecordArgs()
        {
        }
    }

    public sealed class TxtRecordState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The FQDN of the DNS TXT Record.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// The name of the DNS TXT Record. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("records")]
        private InputList<Inputs.TxtRecordRecordGetArgs>? _records;

        /// <summary>
        /// One or more `record` blocks as defined below.
        /// </summary>
        public InputList<Inputs.TxtRecordRecordGetArgs> Records
        {
            get => _records ?? (_records = new InputList<Inputs.TxtRecordRecordGetArgs>());
            set => _records = value;
        }

        /// <summary>
        /// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("zoneName")]
        public Input<string>? ZoneName { get; set; }

        public TxtRecordState()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Dns
{
    /// <summary>
    /// Enables you to manage DNS SRV Records within Azure DNS.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West US",
    ///         });
    ///         var exampleZone = new Azure.Dns.Zone("exampleZone", new Azure.Dns.ZoneArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var exampleSrvRecord = new Azure.Dns.SrvRecord("exampleSrvRecord", new Azure.Dns.SrvRecordArgs
    ///         {
    ///             ZoneName = exampleZone.Name,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Ttl = 300,
    ///             Records = 
    ///             {
    ///                 new Azure.Dns.Inputs.SrvRecordRecordArgs
    ///                 {
    ///                     Priority = 1,
    ///                     Weight = 5,
    ///                     Port = 8080,
    ///                     Target = "target1.contoso.com",
    ///                 },
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "Environment", "Production" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// SRV records can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:dns/srvRecord:SrvRecord example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnszones/zone1/SRV/myrecord1
    /// ```
    /// </summary>
    [AzureResourceType("azure:dns/srvRecord:SrvRecord")]
    public partial class SrvRecord : Pulumi.CustomResource
    {
        /// <summary>
        /// The FQDN of the DNS SRV Record.
        /// </summary>
        [Output("fqdn")]
        public Output<string> Fqdn { get; private set; } = null!;

        /// <summary>
        /// The name of the DNS SRV Record.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of values that make up the SRV record. Each `record` block supports fields documented below.
        /// </summary>
        [Output("records")]
        public Output<ImmutableArray<Outputs.SrvRecordRecord>> Records { get; private set; } = null!;

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

        /// <summary>
        /// The Time To Live (TTL) of the DNS record in seconds.
        /// </summary>
        [Output("ttl")]
        public Output<int> Ttl { get; private set; } = null!;

        /// <summary>
        /// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("zoneName")]
        public Output<string> ZoneName { get; private set; } = null!;


        /// <summary>
        /// Create a SrvRecord resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SrvRecord(string name, SrvRecordArgs args, CustomResourceOptions? options = null)
            : base("azure:dns/srvRecord:SrvRecord", name, args ?? new SrvRecordArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SrvRecord(string name, Input<string> id, SrvRecordState? state = null, CustomResourceOptions? options = null)
            : base("azure:dns/srvRecord:SrvRecord", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SrvRecord resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SrvRecord Get(string name, Input<string> id, SrvRecordState? state = null, CustomResourceOptions? options = null)
        {
            return new SrvRecord(name, id, state, options);
        }
    }

    public sealed class SrvRecordArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the DNS SRV Record.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("records", required: true)]
        private InputList<Inputs.SrvRecordRecordArgs>? _records;

        /// <summary>
        /// A list of values that make up the SRV record. Each `record` block supports fields documented below.
        /// </summary>
        public InputList<Inputs.SrvRecordRecordArgs> Records
        {
            get => _records ?? (_records = new InputList<Inputs.SrvRecordRecordArgs>());
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

        /// <summary>
        /// The Time To Live (TTL) of the DNS record in seconds.
        /// </summary>
        [Input("ttl", required: true)]
        public Input<int> Ttl { get; set; } = null!;

        /// <summary>
        /// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("zoneName", required: true)]
        public Input<string> ZoneName { get; set; } = null!;

        public SrvRecordArgs()
        {
        }
    }

    public sealed class SrvRecordState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The FQDN of the DNS SRV Record.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// The name of the DNS SRV Record.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("records")]
        private InputList<Inputs.SrvRecordRecordGetArgs>? _records;

        /// <summary>
        /// A list of values that make up the SRV record. Each `record` block supports fields documented below.
        /// </summary>
        public InputList<Inputs.SrvRecordRecordGetArgs> Records
        {
            get => _records ?? (_records = new InputList<Inputs.SrvRecordRecordGetArgs>());
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

        /// <summary>
        /// The Time To Live (TTL) of the DNS record in seconds.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("zoneName")]
        public Input<string>? ZoneName { get; set; }

        public SrvRecordState()
        {
        }
    }
}

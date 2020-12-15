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
    /// Enables you to manage DNS PTR Records within Azure DNS.
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
    ///         var examplePtrRecord = new Azure.Dns.PtrRecord("examplePtrRecord", new Azure.Dns.PtrRecordArgs
    ///         {
    ///             ZoneName = exampleZone.Name,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Ttl = 300,
    ///             Records = 
    ///             {
    ///                 "yourdomain.com",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// PTR records can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:dns/ptrRecord:PtrRecord example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnszones/zone1/PTR/myrecord1
    /// ```
    /// </summary>
    [AzureResourceType("azure:dns/ptrRecord:PtrRecord")]
    public partial class PtrRecord : Pulumi.CustomResource
    {
        /// <summary>
        /// The FQDN of the DNS PTR Record.
        /// </summary>
        [Output("fqdn")]
        public Output<string> Fqdn { get; private set; } = null!;

        /// <summary>
        /// The name of the DNS PTR Record.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of Fully Qualified Domain Names.
        /// </summary>
        [Output("records")]
        public Output<ImmutableArray<string>> Records { get; private set; } = null!;

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
        /// Create a PtrRecord resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PtrRecord(string name, PtrRecordArgs args, CustomResourceOptions? options = null)
            : base("azure:dns/ptrRecord:PtrRecord", name, args ?? new PtrRecordArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PtrRecord(string name, Input<string> id, PtrRecordState? state = null, CustomResourceOptions? options = null)
            : base("azure:dns/ptrRecord:PtrRecord", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PtrRecord resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PtrRecord Get(string name, Input<string> id, PtrRecordState? state = null, CustomResourceOptions? options = null)
        {
            return new PtrRecord(name, id, state, options);
        }
    }

    public sealed class PtrRecordArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the DNS PTR Record.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("records", required: true)]
        private InputList<string>? _records;

        /// <summary>
        /// List of Fully Qualified Domain Names.
        /// </summary>
        public InputList<string> Records
        {
            get => _records ?? (_records = new InputList<string>());
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

        public PtrRecordArgs()
        {
        }
    }

    public sealed class PtrRecordState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The FQDN of the DNS PTR Record.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// The name of the DNS PTR Record.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("records")]
        private InputList<string>? _records;

        /// <summary>
        /// List of Fully Qualified Domain Names.
        /// </summary>
        public InputList<string> Records
        {
            get => _records ?? (_records = new InputList<string>());
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

        public PtrRecordState()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Storage
{
    /// <summary>
    /// Manages a Directory within an Azure Storage File Share.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/storage_share_directory.html.markdown.
    /// </summary>
    public partial class ShareDirectory : Pulumi.CustomResource
    {
        /// <summary>
        /// A mapping of metadata to assign to this Directory.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, object>?> Metadata { get; private set; } = null!;

        /// <summary>
        /// The name (or path) of the Directory that should be created within this File Share. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the File Share where this Directory should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("shareName")]
        public Output<string> ShareName { get; private set; } = null!;

        /// <summary>
        /// The name of the Storage Account within which the File Share is located. Changing this forces a new resource to be created.
        /// </summary>
        [Output("storageAccountName")]
        public Output<string> StorageAccountName { get; private set; } = null!;


        /// <summary>
        /// Create a ShareDirectory resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ShareDirectory(string name, ShareDirectoryArgs args, CustomResourceOptions? options = null)
            : base("azure:storage/shareDirectory:ShareDirectory", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ShareDirectory(string name, Input<string> id, ShareDirectoryState? state = null, CustomResourceOptions? options = null)
            : base("azure:storage/shareDirectory:ShareDirectory", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ShareDirectory resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ShareDirectory Get(string name, Input<string> id, ShareDirectoryState? state = null, CustomResourceOptions? options = null)
        {
            return new ShareDirectory(name, id, state, options);
        }
    }

    public sealed class ShareDirectoryArgs : Pulumi.ResourceArgs
    {
        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// A mapping of metadata to assign to this Directory.
        /// </summary>
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        /// <summary>
        /// The name (or path) of the Directory that should be created within this File Share. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the File Share where this Directory should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("shareName", required: true)]
        public Input<string> ShareName { get; set; } = null!;

        /// <summary>
        /// The name of the Storage Account within which the File Share is located. Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageAccountName", required: true)]
        public Input<string> StorageAccountName { get; set; } = null!;

        public ShareDirectoryArgs()
        {
        }
    }

    public sealed class ShareDirectoryState : Pulumi.ResourceArgs
    {
        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// A mapping of metadata to assign to this Directory.
        /// </summary>
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        /// <summary>
        /// The name (or path) of the Directory that should be created within this File Share. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the File Share where this Directory should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("shareName")]
        public Input<string>? ShareName { get; set; }

        /// <summary>
        /// The name of the Storage Account within which the File Share is located. Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageAccountName")]
        public Input<string>? StorageAccountName { get; set; }

        public ShareDirectoryState()
        {
        }
    }
}

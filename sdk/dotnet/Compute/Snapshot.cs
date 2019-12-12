// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute
{
    /// <summary>
    /// Manages a Disk Snapshot.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/snapshot.html.markdown.
    /// </summary>
    public partial class Snapshot : Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates how the snapshot is to be created. Possible values are `Copy` or `Import`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("createOption")]
        public Output<string> CreateOption { get; private set; } = null!;

        /// <summary>
        /// The size of the Snapshotted Disk in GB.
        /// </summary>
        [Output("diskSizeGb")]
        public Output<int> DiskSizeGb { get; private set; } = null!;

        [Output("encryptionSettings")]
        public Output<Outputs.SnapshotEncryptionSettings?> EncryptionSettings { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Snapshot resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Specifies a reference to an existing snapshot, when `create_option` is `Copy`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("sourceResourceId")]
        public Output<string?> SourceResourceId { get; private set; } = null!;

        /// <summary>
        /// Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
        /// </summary>
        [Output("sourceUri")]
        public Output<string?> SourceUri { get; private set; } = null!;

        /// <summary>
        /// Specifies the ID of an storage account. Used with `source_uri` to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
        /// </summary>
        [Output("storageAccountId")]
        public Output<string?> StorageAccountId { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Snapshot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Snapshot(string name, SnapshotArgs args, CustomResourceOptions? options = null)
            : base("azure:compute/snapshot:Snapshot", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Snapshot(string name, Input<string> id, SnapshotState? state = null, CustomResourceOptions? options = null)
            : base("azure:compute/snapshot:Snapshot", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Snapshot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Snapshot Get(string name, Input<string> id, SnapshotState? state = null, CustomResourceOptions? options = null)
        {
            return new Snapshot(name, id, state, options);
        }
    }

    public sealed class SnapshotArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates how the snapshot is to be created. Possible values are `Copy` or `Import`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("createOption", required: true)]
        public Input<string> CreateOption { get; set; } = null!;

        /// <summary>
        /// The size of the Snapshotted Disk in GB.
        /// </summary>
        [Input("diskSizeGb")]
        public Input<int>? DiskSizeGb { get; set; }

        [Input("encryptionSettings")]
        public Input<Inputs.SnapshotEncryptionSettingsArgs>? EncryptionSettings { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Snapshot resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies a reference to an existing snapshot, when `create_option` is `Copy`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("sourceResourceId")]
        public Input<string>? SourceResourceId { get; set; }

        /// <summary>
        /// Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
        /// </summary>
        [Input("sourceUri")]
        public Input<string>? SourceUri { get; set; }

        /// <summary>
        /// Specifies the ID of an storage account. Used with `source_uri` to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageAccountId")]
        public Input<string>? StorageAccountId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public SnapshotArgs()
        {
        }
    }

    public sealed class SnapshotState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates how the snapshot is to be created. Possible values are `Copy` or `Import`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("createOption")]
        public Input<string>? CreateOption { get; set; }

        /// <summary>
        /// The size of the Snapshotted Disk in GB.
        /// </summary>
        [Input("diskSizeGb")]
        public Input<int>? DiskSizeGb { get; set; }

        [Input("encryptionSettings")]
        public Input<Inputs.SnapshotEncryptionSettingsGetArgs>? EncryptionSettings { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Snapshot resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Specifies a reference to an existing snapshot, when `create_option` is `Copy`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("sourceResourceId")]
        public Input<string>? SourceResourceId { get; set; }

        /// <summary>
        /// Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
        /// </summary>
        [Input("sourceUri")]
        public Input<string>? SourceUri { get; set; }

        /// <summary>
        /// Specifies the ID of an storage account. Used with `source_uri` to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageAccountId")]
        public Input<string>? StorageAccountId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public SnapshotState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class SnapshotEncryptionSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("diskEncryptionKey")]
        public Input<SnapshotEncryptionSettingsDiskEncryptionKeyArgs>? DiskEncryptionKey { get; set; }

        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("keyEncryptionKey")]
        public Input<SnapshotEncryptionSettingsKeyEncryptionKeyArgs>? KeyEncryptionKey { get; set; }

        public SnapshotEncryptionSettingsArgs()
        {
        }
    }

    public sealed class SnapshotEncryptionSettingsDiskEncryptionKeyArgs : Pulumi.ResourceArgs
    {
        [Input("secretUrl", required: true)]
        public Input<string> SecretUrl { get; set; } = null!;

        [Input("sourceVaultId", required: true)]
        public Input<string> SourceVaultId { get; set; } = null!;

        public SnapshotEncryptionSettingsDiskEncryptionKeyArgs()
        {
        }
    }

    public sealed class SnapshotEncryptionSettingsDiskEncryptionKeyGetArgs : Pulumi.ResourceArgs
    {
        [Input("secretUrl", required: true)]
        public Input<string> SecretUrl { get; set; } = null!;

        [Input("sourceVaultId", required: true)]
        public Input<string> SourceVaultId { get; set; } = null!;

        public SnapshotEncryptionSettingsDiskEncryptionKeyGetArgs()
        {
        }
    }

    public sealed class SnapshotEncryptionSettingsGetArgs : Pulumi.ResourceArgs
    {
        [Input("diskEncryptionKey")]
        public Input<SnapshotEncryptionSettingsDiskEncryptionKeyGetArgs>? DiskEncryptionKey { get; set; }

        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("keyEncryptionKey")]
        public Input<SnapshotEncryptionSettingsKeyEncryptionKeyGetArgs>? KeyEncryptionKey { get; set; }

        public SnapshotEncryptionSettingsGetArgs()
        {
        }
    }

    public sealed class SnapshotEncryptionSettingsKeyEncryptionKeyArgs : Pulumi.ResourceArgs
    {
        [Input("keyUrl", required: true)]
        public Input<string> KeyUrl { get; set; } = null!;

        [Input("sourceVaultId", required: true)]
        public Input<string> SourceVaultId { get; set; } = null!;

        public SnapshotEncryptionSettingsKeyEncryptionKeyArgs()
        {
        }
    }

    public sealed class SnapshotEncryptionSettingsKeyEncryptionKeyGetArgs : Pulumi.ResourceArgs
    {
        [Input("keyUrl", required: true)]
        public Input<string> KeyUrl { get; set; } = null!;

        [Input("sourceVaultId", required: true)]
        public Input<string> SourceVaultId { get; set; } = null!;

        public SnapshotEncryptionSettingsKeyEncryptionKeyGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class SnapshotEncryptionSettings
    {
        public readonly SnapshotEncryptionSettingsDiskEncryptionKey? DiskEncryptionKey;
        public readonly bool Enabled;
        public readonly SnapshotEncryptionSettingsKeyEncryptionKey? KeyEncryptionKey;

        [OutputConstructor]
        private SnapshotEncryptionSettings(
            SnapshotEncryptionSettingsDiskEncryptionKey? diskEncryptionKey,
            bool enabled,
            SnapshotEncryptionSettingsKeyEncryptionKey? keyEncryptionKey)
        {
            DiskEncryptionKey = diskEncryptionKey;
            Enabled = enabled;
            KeyEncryptionKey = keyEncryptionKey;
        }
    }

    [OutputType]
    public sealed class SnapshotEncryptionSettingsDiskEncryptionKey
    {
        public readonly string SecretUrl;
        public readonly string SourceVaultId;

        [OutputConstructor]
        private SnapshotEncryptionSettingsDiskEncryptionKey(
            string secretUrl,
            string sourceVaultId)
        {
            SecretUrl = secretUrl;
            SourceVaultId = sourceVaultId;
        }
    }

    [OutputType]
    public sealed class SnapshotEncryptionSettingsKeyEncryptionKey
    {
        public readonly string KeyUrl;
        public readonly string SourceVaultId;

        [OutputConstructor]
        private SnapshotEncryptionSettingsKeyEncryptionKey(
            string keyUrl,
            string sourceVaultId)
        {
            KeyUrl = keyUrl;
            SourceVaultId = sourceVaultId;
        }
    }
    }
}

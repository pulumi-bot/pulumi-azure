// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService.Outputs
{

    [OutputType]
    public sealed class GroupContainerVolume
    {
        /// <summary>
        /// A `git_repo` block as defined below.
        /// </summary>
        public readonly Outputs.GroupContainerVolumeGitRepo? GitRepo;
        /// <summary>
        /// The path on which this volume is to be mounted. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string MountPath;
        /// <summary>
        /// Specifies the name of the Container Group. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specify if the volume is to be mounted as read only or not. The default value is `false`. Changing this forces a new resource to be created.
        /// </summary>
        public readonly bool? ReadOnly;
        /// <summary>
        /// A map of secrets that will be mounted as files in the volume. Changing this forces a new resource to be created.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Secret;
        /// <summary>
        /// The Azure storage share that is to be mounted as a volume. This must be created on the storage account specified as above. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? ShareName;
        /// <summary>
        /// The access key for the Azure Storage account specified as above. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? StorageAccountKey;
        /// <summary>
        /// The Azure storage account from which the volume is to be mounted. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? StorageAccountName;

        [OutputConstructor]
        private GroupContainerVolume(
            Outputs.GroupContainerVolumeGitRepo? gitRepo,

            string mountPath,

            string name,

            bool? readOnly,

            ImmutableDictionary<string, string>? secret,

            string? shareName,

            string? storageAccountKey,

            string? storageAccountName)
        {
            GitRepo = gitRepo;
            MountPath = mountPath;
            Name = name;
            ReadOnly = readOnly;
            Secret = secret;
            ShareName = shareName;
            StorageAccountKey = storageAccountKey;
            StorageAccountName = storageAccountName;
        }
    }
}

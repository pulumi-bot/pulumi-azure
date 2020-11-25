// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class WindowsVirtualMachineScaleSetDataDiskGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of Caching which should be used for this Data Disk. Possible values are `None`, `ReadOnly` and `ReadWrite`.
        /// </summary>
        [Input("caching", required: true)]
        public Input<string> Caching { get; set; } = null!;

        /// <summary>
        /// The create option which should be used for this Data Disk. Possible values are `Empty` and `FromImage`. Defaults to `Empty`. (`FromImage` should only be used if the source image includes data disks).
        /// </summary>
        [Input("createOption")]
        public Input<string>? CreateOption { get; set; }

        /// <summary>
        /// The ID of the Disk Encryption Set which should be used to encrypt this Data Disk.
        /// </summary>
        [Input("diskEncryptionSetId")]
        public Input<string>? DiskEncryptionSetId { get; set; }

        /// <summary>
        /// Specifies the Read-Write IOPS for this Data Disk. Only settable for UltraSSD disks.
        /// </summary>
        [Input("diskIopsReadWrite")]
        public Input<int>? DiskIopsReadWrite { get; set; }

        /// <summary>
        /// Specifies the bandwidth in MB per second for this Data Disk. Only settable for UltraSSD disks.
        /// </summary>
        [Input("diskMbpsReadWrite")]
        public Input<int>? DiskMbpsReadWrite { get; set; }

        /// <summary>
        /// The size of the Data Disk which should be created.
        /// </summary>
        [Input("diskSizeGb", required: true)]
        public Input<int> DiskSizeGb { get; set; } = null!;

        /// <summary>
        /// The Logical Unit Number of the Data Disk, which must be unique within the Virtual Machine.
        /// </summary>
        [Input("lun", required: true)]
        public Input<int> Lun { get; set; } = null!;

        /// <summary>
        /// The Type of Storage Account which should back this Data Disk. Possible values include `Standard_LRS`, `StandardSSD_LRS`, `Premium_LRS` and `UltraSSD_LRS`.
        /// </summary>
        [Input("storageAccountType", required: true)]
        public Input<string> StorageAccountType { get; set; } = null!;

        /// <summary>
        /// Should Write Accelerator be enabled for this Data Disk? Defaults to `false`.
        /// </summary>
        [Input("writeAcceleratorEnabled")]
        public Input<bool>? WriteAcceleratorEnabled { get; set; }

        public WindowsVirtualMachineScaleSetDataDiskGetArgs()
        {
        }
    }
}

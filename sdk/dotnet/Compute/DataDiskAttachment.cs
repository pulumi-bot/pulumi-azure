// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute
{
    /// <summary>
    /// Manages attaching a Disk to a Virtual Machine.
    /// 
    /// &gt; **NOTE:** Data Disks can be attached either directly on the `azure.compute.VirtualMachine` resource, or using the `azure.compute.DataDiskAttachment` resource - but the two cannot be used together. If both are used against the same Virtual Machine, spurious changes will occur.
    /// 
    /// &gt; **Please Note:** only Managed Disks are supported via this separate resource, Unmanaged Disks can be attached using the `storage_data_disk` block in the `azure.compute.VirtualMachine` resource.
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
    ///         var config = new Config();
    ///         var prefix = config.Get("prefix") ?? "example";
    ///         var vmName = $"{prefix}-vm";
    ///         var mainResourceGroup = new Azure.Core.ResourceGroup("mainResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var mainVirtualNetwork = new Azure.Network.VirtualNetwork("mainVirtualNetwork", new Azure.Network.VirtualNetworkArgs
    ///         {
    ///             AddressSpaces = 
    ///             {
    ///                 "10.0.0.0/16",
    ///             },
    ///             Location = mainResourceGroup.Location,
    ///             ResourceGroupName = mainResourceGroup.Name,
    ///         });
    ///         var @internal = new Azure.Network.Subnet("internal", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = mainResourceGroup.Name,
    ///             VirtualNetworkName = mainVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.0.2.0/24",
    ///             },
    ///         });
    ///         var mainNetworkInterface = new Azure.Network.NetworkInterface("mainNetworkInterface", new Azure.Network.NetworkInterfaceArgs
    ///         {
    ///             Location = mainResourceGroup.Location,
    ///             ResourceGroupName = mainResourceGroup.Name,
    ///             IpConfigurations = 
    ///             {
    ///                 new Azure.Network.Inputs.NetworkInterfaceIpConfigurationArgs
    ///                 {
    ///                     Name = "internal",
    ///                     SubnetId = @internal.Id,
    ///                     PrivateIpAddressAllocation = "Dynamic",
    ///                 },
    ///             },
    ///         });
    ///         var exampleVirtualMachine = new Azure.Compute.VirtualMachine("exampleVirtualMachine", new Azure.Compute.VirtualMachineArgs
    ///         {
    ///             Location = mainResourceGroup.Location,
    ///             ResourceGroupName = mainResourceGroup.Name,
    ///             NetworkInterfaceIds = 
    ///             {
    ///                 mainNetworkInterface.Id,
    ///             },
    ///             VmSize = "Standard_F2",
    ///             StorageImageReference = new Azure.Compute.Inputs.VirtualMachineStorageImageReferenceArgs
    ///             {
    ///                 Publisher = "Canonical",
    ///                 Offer = "UbuntuServer",
    ///                 Sku = "16.04-LTS",
    ///                 Version = "latest",
    ///             },
    ///             StorageOsDisk = new Azure.Compute.Inputs.VirtualMachineStorageOsDiskArgs
    ///             {
    ///                 Name = "myosdisk1",
    ///                 Caching = "ReadWrite",
    ///                 CreateOption = "FromImage",
    ///                 ManagedDiskType = "Standard_LRS",
    ///             },
    ///             OsProfile = new Azure.Compute.Inputs.VirtualMachineOsProfileArgs
    ///             {
    ///                 ComputerName = vmName,
    ///                 AdminUsername = "testadmin",
    ///                 AdminPassword = "Password1234!",
    ///             },
    ///             OsProfileLinuxConfig = new Azure.Compute.Inputs.VirtualMachineOsProfileLinuxConfigArgs
    ///             {
    ///                 DisablePasswordAuthentication = false,
    ///             },
    ///         });
    ///         var exampleManagedDisk = new Azure.Compute.ManagedDisk("exampleManagedDisk", new Azure.Compute.ManagedDiskArgs
    ///         {
    ///             Location = mainResourceGroup.Location,
    ///             ResourceGroupName = mainResourceGroup.Name,
    ///             StorageAccountType = "Standard_LRS",
    ///             CreateOption = "Empty",
    ///             DiskSizeGb = 10,
    ///         });
    ///         var exampleDataDiskAttachment = new Azure.Compute.DataDiskAttachment("exampleDataDiskAttachment", new Azure.Compute.DataDiskAttachmentArgs
    ///         {
    ///             ManagedDiskId = exampleManagedDisk.Id,
    ///             VirtualMachineId = exampleVirtualMachine.Id,
    ///             Lun = 10,
    ///             Caching = "ReadWrite",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Virtual Machines Data Disk Attachments can be imported using the `resource id`, e.g.
    /// 
    ///  -&gt; **Please Note:** This is a ID (specific to this provider) matching the format`{virtualMachineID}/dataDisks/{diskName}`
    /// </summary>
    public partial class DataDiskAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the caching requirements for this Data Disk. Possible values include `None`, `ReadOnly` and `ReadWrite`.
        /// </summary>
        [Output("caching")]
        public Output<string> Caching { get; private set; } = null!;

        /// <summary>
        /// The Create Option of the Data Disk, such as `Empty` or `Attach`. Defaults to `Attach`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("createOption")]
        public Output<string?> CreateOption { get; private set; } = null!;

        /// <summary>
        /// The Logical Unit Number of the Data Disk, which needs to be unique within the Virtual Machine. Changing this forces a new resource to be created.
        /// </summary>
        [Output("lun")]
        public Output<int> Lun { get; private set; } = null!;

        /// <summary>
        /// The ID of an existing Managed Disk which should be attached. Changing this forces a new resource to be created.
        /// </summary>
        [Output("managedDiskId")]
        public Output<string> ManagedDiskId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Virtual Machine to which the Data Disk should be attached. Changing this forces a new resource to be created.
        /// </summary>
        [Output("virtualMachineId")]
        public Output<string> VirtualMachineId { get; private set; } = null!;

        /// <summary>
        /// Specifies if Write Accelerator is enabled on the disk. This can only be enabled on `Premium_LRS` managed disks with no caching and [M-Series VMs](https://docs.microsoft.com/en-us/azure/virtual-machines/workloads/sap/how-to-enable-write-accelerator). Defaults to `false`.
        /// </summary>
        [Output("writeAcceleratorEnabled")]
        public Output<bool?> WriteAcceleratorEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a DataDiskAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataDiskAttachment(string name, DataDiskAttachmentArgs args, CustomResourceOptions? options = null)
            : base("azure:compute/dataDiskAttachment:DataDiskAttachment", name, args ?? new DataDiskAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataDiskAttachment(string name, Input<string> id, DataDiskAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("azure:compute/dataDiskAttachment:DataDiskAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DataDiskAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataDiskAttachment Get(string name, Input<string> id, DataDiskAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new DataDiskAttachment(name, id, state, options);
        }
    }

    public sealed class DataDiskAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the caching requirements for this Data Disk. Possible values include `None`, `ReadOnly` and `ReadWrite`.
        /// </summary>
        [Input("caching", required: true)]
        public Input<string> Caching { get; set; } = null!;

        /// <summary>
        /// The Create Option of the Data Disk, such as `Empty` or `Attach`. Defaults to `Attach`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("createOption")]
        public Input<string>? CreateOption { get; set; }

        /// <summary>
        /// The Logical Unit Number of the Data Disk, which needs to be unique within the Virtual Machine. Changing this forces a new resource to be created.
        /// </summary>
        [Input("lun", required: true)]
        public Input<int> Lun { get; set; } = null!;

        /// <summary>
        /// The ID of an existing Managed Disk which should be attached. Changing this forces a new resource to be created.
        /// </summary>
        [Input("managedDiskId", required: true)]
        public Input<string> ManagedDiskId { get; set; } = null!;

        /// <summary>
        /// The ID of the Virtual Machine to which the Data Disk should be attached. Changing this forces a new resource to be created.
        /// </summary>
        [Input("virtualMachineId", required: true)]
        public Input<string> VirtualMachineId { get; set; } = null!;

        /// <summary>
        /// Specifies if Write Accelerator is enabled on the disk. This can only be enabled on `Premium_LRS` managed disks with no caching and [M-Series VMs](https://docs.microsoft.com/en-us/azure/virtual-machines/workloads/sap/how-to-enable-write-accelerator). Defaults to `false`.
        /// </summary>
        [Input("writeAcceleratorEnabled")]
        public Input<bool>? WriteAcceleratorEnabled { get; set; }

        public DataDiskAttachmentArgs()
        {
        }
    }

    public sealed class DataDiskAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the caching requirements for this Data Disk. Possible values include `None`, `ReadOnly` and `ReadWrite`.
        /// </summary>
        [Input("caching")]
        public Input<string>? Caching { get; set; }

        /// <summary>
        /// The Create Option of the Data Disk, such as `Empty` or `Attach`. Defaults to `Attach`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("createOption")]
        public Input<string>? CreateOption { get; set; }

        /// <summary>
        /// The Logical Unit Number of the Data Disk, which needs to be unique within the Virtual Machine. Changing this forces a new resource to be created.
        /// </summary>
        [Input("lun")]
        public Input<int>? Lun { get; set; }

        /// <summary>
        /// The ID of an existing Managed Disk which should be attached. Changing this forces a new resource to be created.
        /// </summary>
        [Input("managedDiskId")]
        public Input<string>? ManagedDiskId { get; set; }

        /// <summary>
        /// The ID of the Virtual Machine to which the Data Disk should be attached. Changing this forces a new resource to be created.
        /// </summary>
        [Input("virtualMachineId")]
        public Input<string>? VirtualMachineId { get; set; }

        /// <summary>
        /// Specifies if Write Accelerator is enabled on the disk. This can only be enabled on `Premium_LRS` managed disks with no caching and [M-Series VMs](https://docs.microsoft.com/en-us/azure/virtual-machines/workloads/sap/how-to-enable-write-accelerator). Defaults to `false`.
        /// </summary>
        [Input("writeAcceleratorEnabled")]
        public Input<bool>? WriteAcceleratorEnabled { get; set; }

        public DataDiskAttachmentState()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages attaching a Disk to a Virtual Machine.
//
// > **NOTE:** Data Disks can be attached either directly on the `compute.VirtualMachine` resource, or using the `compute.DataDiskAttachment` resource - but the two cannot be used together. If both are used against the same Virtual Machine, spurious changes will occur.
//
// > **Please Note:** only Managed Disks are supported via this separate resource, Unmanaged Disks can be attached using the `storageDataDisk` block in the `compute.VirtualMachine` resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		prefix := "example"
// 		if param := cfg.Get("prefix"); param != "" {
// 			prefix = param
// 		}
// 		vmName := fmt.Sprintf("%v%v", prefix, "-vm")
// 		mainResourceGroup, err := core.NewResourceGroup(ctx, "mainResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		mainVirtualNetwork, err := network.NewVirtualNetwork(ctx, "mainVirtualNetwork", &network.VirtualNetworkArgs{
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
// 			},
// 			Location:          mainResourceGroup.Location,
// 			ResourceGroupName: mainResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		internal, err := network.NewSubnet(ctx, "internal", &network.SubnetArgs{
// 			ResourceGroupName:  mainResourceGroup.Name,
// 			VirtualNetworkName: mainVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.2.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		mainNetworkInterface, err := network.NewNetworkInterface(ctx, "mainNetworkInterface", &network.NetworkInterfaceArgs{
// 			Location:          mainResourceGroup.Location,
// 			ResourceGroupName: mainResourceGroup.Name,
// 			IpConfigurations: network.NetworkInterfaceIpConfigurationArray{
// 				&network.NetworkInterfaceIpConfigurationArgs{
// 					Name:                       pulumi.String("internal"),
// 					SubnetId:                   internal.ID(),
// 					PrivateIpAddressAllocation: pulumi.String("Dynamic"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualMachine, err := compute.NewVirtualMachine(ctx, "exampleVirtualMachine", &compute.VirtualMachineArgs{
// 			Location:          mainResourceGroup.Location,
// 			ResourceGroupName: mainResourceGroup.Name,
// 			NetworkInterfaceIds: pulumi.StringArray{
// 				mainNetworkInterface.ID(),
// 			},
// 			VmSize: pulumi.String("Standard_F2"),
// 			StorageImageReference: &compute.VirtualMachineStorageImageReferenceArgs{
// 				Publisher: pulumi.String("Canonical"),
// 				Offer:     pulumi.String("UbuntuServer"),
// 				Sku:       pulumi.String("16.04-LTS"),
// 				Version:   pulumi.String("latest"),
// 			},
// 			StorageOsDisk: &compute.VirtualMachineStorageOsDiskArgs{
// 				Name:            pulumi.String("myosdisk1"),
// 				Caching:         pulumi.String("ReadWrite"),
// 				CreateOption:    pulumi.String("FromImage"),
// 				ManagedDiskType: pulumi.String("Standard_LRS"),
// 			},
// 			OsProfile: &compute.VirtualMachineOsProfileArgs{
// 				ComputerName:  pulumi.String(vmName),
// 				AdminUsername: pulumi.String("testadmin"),
// 				AdminPassword: pulumi.String("Password1234!"),
// 			},
// 			OsProfileLinuxConfig: &compute.VirtualMachineOsProfileLinuxConfigArgs{
// 				DisablePasswordAuthentication: pulumi.Bool(false),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleManagedDisk, err := compute.NewManagedDisk(ctx, "exampleManagedDisk", &compute.ManagedDiskArgs{
// 			Location:           mainResourceGroup.Location,
// 			ResourceGroupName:  mainResourceGroup.Name,
// 			StorageAccountType: pulumi.String("Standard_LRS"),
// 			CreateOption:       pulumi.String("Empty"),
// 			DiskSizeGb:         pulumi.Int(10),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewDataDiskAttachment(ctx, "exampleDataDiskAttachment", &compute.DataDiskAttachmentArgs{
// 			ManagedDiskId:    exampleManagedDisk.ID(),
// 			VirtualMachineId: exampleVirtualMachine.ID(),
// 			Lun:              pulumi.Int(10),
// 			Caching:          pulumi.String("ReadWrite"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DataDiskAttachment struct {
	pulumi.CustomResourceState

	// Specifies the caching requirements for this Data Disk. Possible values include `None`, `ReadOnly` and `ReadWrite`.
	Caching pulumi.StringOutput `pulumi:"caching"`
	// The Create Option of the Data Disk, such as `Empty` or `Attach`. Defaults to `Attach`. Changing this forces a new resource to be created.
	CreateOption pulumi.StringPtrOutput `pulumi:"createOption"`
	// The Logical Unit Number of the Data Disk, which needs to be unique within the Virtual Machine. Changing this forces a new resource to be created.
	Lun pulumi.IntOutput `pulumi:"lun"`
	// The ID of an existing Managed Disk which should be attached. Changing this forces a new resource to be created.
	ManagedDiskId pulumi.StringOutput `pulumi:"managedDiskId"`
	// The ID of the Virtual Machine to which the Data Disk should be attached. Changing this forces a new resource to be created.
	VirtualMachineId pulumi.StringOutput `pulumi:"virtualMachineId"`
	// Specifies if Write Accelerator is enabled on the disk. This can only be enabled on `Premium_LRS` managed disks with no caching and [M-Series VMs](https://docs.microsoft.com/en-us/azure/virtual-machines/workloads/sap/how-to-enable-write-accelerator). Defaults to `false`.
	WriteAcceleratorEnabled pulumi.BoolPtrOutput `pulumi:"writeAcceleratorEnabled"`
}

// NewDataDiskAttachment registers a new resource with the given unique name, arguments, and options.
func NewDataDiskAttachment(ctx *pulumi.Context,
	name string, args *DataDiskAttachmentArgs, opts ...pulumi.ResourceOption) (*DataDiskAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Caching == nil {
		return nil, errors.New("invalid value for required argument 'Caching'")
	}
	if args.Lun == nil {
		return nil, errors.New("invalid value for required argument 'Lun'")
	}
	if args.ManagedDiskId == nil {
		return nil, errors.New("invalid value for required argument 'ManagedDiskId'")
	}
	if args.VirtualMachineId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineId'")
	}
	var resource DataDiskAttachment
	err := ctx.RegisterResource("azure:compute/dataDiskAttachment:DataDiskAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataDiskAttachment gets an existing DataDiskAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataDiskAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataDiskAttachmentState, opts ...pulumi.ResourceOption) (*DataDiskAttachment, error) {
	var resource DataDiskAttachment
	err := ctx.ReadResource("azure:compute/dataDiskAttachment:DataDiskAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataDiskAttachment resources.
type dataDiskAttachmentState struct {
	// Specifies the caching requirements for this Data Disk. Possible values include `None`, `ReadOnly` and `ReadWrite`.
	Caching *string `pulumi:"caching"`
	// The Create Option of the Data Disk, such as `Empty` or `Attach`. Defaults to `Attach`. Changing this forces a new resource to be created.
	CreateOption *string `pulumi:"createOption"`
	// The Logical Unit Number of the Data Disk, which needs to be unique within the Virtual Machine. Changing this forces a new resource to be created.
	Lun *int `pulumi:"lun"`
	// The ID of an existing Managed Disk which should be attached. Changing this forces a new resource to be created.
	ManagedDiskId *string `pulumi:"managedDiskId"`
	// The ID of the Virtual Machine to which the Data Disk should be attached. Changing this forces a new resource to be created.
	VirtualMachineId *string `pulumi:"virtualMachineId"`
	// Specifies if Write Accelerator is enabled on the disk. This can only be enabled on `Premium_LRS` managed disks with no caching and [M-Series VMs](https://docs.microsoft.com/en-us/azure/virtual-machines/workloads/sap/how-to-enable-write-accelerator). Defaults to `false`.
	WriteAcceleratorEnabled *bool `pulumi:"writeAcceleratorEnabled"`
}

type DataDiskAttachmentState struct {
	// Specifies the caching requirements for this Data Disk. Possible values include `None`, `ReadOnly` and `ReadWrite`.
	Caching pulumi.StringPtrInput
	// The Create Option of the Data Disk, such as `Empty` or `Attach`. Defaults to `Attach`. Changing this forces a new resource to be created.
	CreateOption pulumi.StringPtrInput
	// The Logical Unit Number of the Data Disk, which needs to be unique within the Virtual Machine. Changing this forces a new resource to be created.
	Lun pulumi.IntPtrInput
	// The ID of an existing Managed Disk which should be attached. Changing this forces a new resource to be created.
	ManagedDiskId pulumi.StringPtrInput
	// The ID of the Virtual Machine to which the Data Disk should be attached. Changing this forces a new resource to be created.
	VirtualMachineId pulumi.StringPtrInput
	// Specifies if Write Accelerator is enabled on the disk. This can only be enabled on `Premium_LRS` managed disks with no caching and [M-Series VMs](https://docs.microsoft.com/en-us/azure/virtual-machines/workloads/sap/how-to-enable-write-accelerator). Defaults to `false`.
	WriteAcceleratorEnabled pulumi.BoolPtrInput
}

func (DataDiskAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataDiskAttachmentState)(nil)).Elem()
}

type dataDiskAttachmentArgs struct {
	// Specifies the caching requirements for this Data Disk. Possible values include `None`, `ReadOnly` and `ReadWrite`.
	Caching string `pulumi:"caching"`
	// The Create Option of the Data Disk, such as `Empty` or `Attach`. Defaults to `Attach`. Changing this forces a new resource to be created.
	CreateOption *string `pulumi:"createOption"`
	// The Logical Unit Number of the Data Disk, which needs to be unique within the Virtual Machine. Changing this forces a new resource to be created.
	Lun int `pulumi:"lun"`
	// The ID of an existing Managed Disk which should be attached. Changing this forces a new resource to be created.
	ManagedDiskId string `pulumi:"managedDiskId"`
	// The ID of the Virtual Machine to which the Data Disk should be attached. Changing this forces a new resource to be created.
	VirtualMachineId string `pulumi:"virtualMachineId"`
	// Specifies if Write Accelerator is enabled on the disk. This can only be enabled on `Premium_LRS` managed disks with no caching and [M-Series VMs](https://docs.microsoft.com/en-us/azure/virtual-machines/workloads/sap/how-to-enable-write-accelerator). Defaults to `false`.
	WriteAcceleratorEnabled *bool `pulumi:"writeAcceleratorEnabled"`
}

// The set of arguments for constructing a DataDiskAttachment resource.
type DataDiskAttachmentArgs struct {
	// Specifies the caching requirements for this Data Disk. Possible values include `None`, `ReadOnly` and `ReadWrite`.
	Caching pulumi.StringInput
	// The Create Option of the Data Disk, such as `Empty` or `Attach`. Defaults to `Attach`. Changing this forces a new resource to be created.
	CreateOption pulumi.StringPtrInput
	// The Logical Unit Number of the Data Disk, which needs to be unique within the Virtual Machine. Changing this forces a new resource to be created.
	Lun pulumi.IntInput
	// The ID of an existing Managed Disk which should be attached. Changing this forces a new resource to be created.
	ManagedDiskId pulumi.StringInput
	// The ID of the Virtual Machine to which the Data Disk should be attached. Changing this forces a new resource to be created.
	VirtualMachineId pulumi.StringInput
	// Specifies if Write Accelerator is enabled on the disk. This can only be enabled on `Premium_LRS` managed disks with no caching and [M-Series VMs](https://docs.microsoft.com/en-us/azure/virtual-machines/workloads/sap/how-to-enable-write-accelerator). Defaults to `false`.
	WriteAcceleratorEnabled pulumi.BoolPtrInput
}

func (DataDiskAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataDiskAttachmentArgs)(nil)).Elem()
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Linux Virtual Machine.
//
// ## Disclaimers
//
// > **Note** This provider will automatically remove the OS Disk by default - this behaviour can be configured using the `features` configuration within the Provider configuration block.
//
// > **Note** This resource does not support Unmanaged Disks. If you need to use Unmanaged Disks you can continue to use the `compute.VirtualMachine` resource instead.
//
// > **Note** This resource does not support attaching existing OS Disks. You can instead capture an image of the OS Disk or continue to use the `compute.VirtualMachine` resource instead.
//
// > In this release there's a known issue where the `publicIpAddress` and `publicIpAddresses` fields may not be fully populated for Dynamic Public IP's.
//
// ## Import
//
// Linux Virtual Machines can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:compute/linuxVirtualMachine:LinuxVirtualMachine example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachines/machine1
// ```
type LinuxVirtualMachine struct {
	pulumi.CustomResourceState

	// A `additionalCapabilities` block as defined below.
	AdditionalCapabilities LinuxVirtualMachineAdditionalCapabilitiesPtrOutput `pulumi:"additionalCapabilities"`
	// The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
	AdminPassword pulumi.StringPtrOutput `pulumi:"adminPassword"`
	// One or more `adminSshKey` blocks as defined below.
	AdminSshKeys LinuxVirtualMachineAdminSshKeyArrayOutput `pulumi:"adminSshKeys"`
	// The username of the local administrator used for the Virtual Machine. Changing this forces a new resource to be created.
	AdminUsername pulumi.StringOutput `pulumi:"adminUsername"`
	// Should Extension Operations be allowed on this Virtual Machine?
	AllowExtensionOperations pulumi.BoolPtrOutput `pulumi:"allowExtensionOperations"`
	// Specifies the ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
	AvailabilitySetId pulumi.StringPtrOutput `pulumi:"availabilitySetId"`
	// A `bootDiagnostics` block as defined below.
	BootDiagnostics LinuxVirtualMachineBootDiagnosticsPtrOutput `pulumi:"bootDiagnostics"`
	// Specifies the Hostname which should be used for this Virtual Machine. If unspecified this defaults to the value for the `name` field. If the value of the `name` field is not a valid `computerName`, then you must specify `computerName`. Changing this forces a new resource to be created.
	ComputerName pulumi.StringOutput `pulumi:"computerName"`
	// The Base64-Encoded Custom Data which should be used for this Virtual Machine. Changing this forces a new resource to be created.
	CustomData pulumi.StringPtrOutput `pulumi:"customData"`
	// The ID of a Dedicated Host where this machine should be run on. Changing this forces a new resource to be created.
	DedicatedHostId pulumi.StringPtrOutput `pulumi:"dedicatedHostId"`
	// Should Password Authentication be disabled on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
	DisablePasswordAuthentication pulumi.BoolPtrOutput `pulumi:"disablePasswordAuthentication"`
	// Should all of the disks (including the temp disk) attached to this Virtual Machine be encrypted by enabling Encryption at Host?
	EncryptionAtHostEnabled pulumi.BoolPtrOutput `pulumi:"encryptionAtHostEnabled"`
	// Specifies what should happen when the Virtual Machine is evicted for price reasons when using a Spot instance. At this time the only supported value is `Deallocate`. Changing this forces a new resource to be created.
	EvictionPolicy pulumi.StringPtrOutput `pulumi:"evictionPolicy"`
	// An `identity` block as defined below.
	Identity LinuxVirtualMachineIdentityPtrOutput `pulumi:"identity"`
	// The Azure location where the Linux Virtual Machine should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The maximum price you're willing to pay for this Virtual Machine, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machine will be evicted using the `evictionPolicy`. Defaults to `-1`, which means that the Virtual Machine should not be evicted for price reasons.
	MaxBidPrice pulumi.Float64PtrOutput `pulumi:"maxBidPrice"`
	// The name of the Linux Virtual Machine. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// . A list of Network Interface ID's which should be attached to this Virtual Machine. The first Network Interface ID in this list will be the Primary Network Interface on the Virtual Machine.
	NetworkInterfaceIds pulumi.StringArrayOutput `pulumi:"networkInterfaceIds"`
	// A `osDisk` block as defined below.
	OsDisk LinuxVirtualMachineOsDiskOutput `pulumi:"osDisk"`
	// A `plan` block as defined below. Changing this forces a new resource to be created.
	Plan LinuxVirtualMachinePlanPtrOutput `pulumi:"plan"`
	// Specifies the priority of this Virtual Machine. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
	Priority pulumi.StringPtrOutput `pulumi:"priority"`
	// The Primary Private IP Address assigned to this Virtual Machine.
	PrivateIpAddress pulumi.StringOutput `pulumi:"privateIpAddress"`
	// A list of Private IP Addresses assigned to this Virtual Machine.
	PrivateIpAddresses pulumi.StringArrayOutput `pulumi:"privateIpAddresses"`
	// Should the Azure VM Agent be provisioned on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
	ProvisionVmAgent pulumi.BoolPtrOutput `pulumi:"provisionVmAgent"`
	// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId pulumi.StringPtrOutput `pulumi:"proximityPlacementGroupId"`
	// The Primary Public IP Address assigned to this Virtual Machine.
	PublicIpAddress pulumi.StringOutput `pulumi:"publicIpAddress"`
	// A list of the Public IP Addresses assigned to this Virtual Machine.
	PublicIpAddresses pulumi.StringArrayOutput `pulumi:"publicIpAddresses"`
	// The name of the Resource Group in which the Linux Virtual Machine should be exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// One or more `secret` blocks as defined below.
	Secrets LinuxVirtualMachineSecretArrayOutput `pulumi:"secrets"`
	// The SKU which should be used for this Virtual Machine, such as `Standard_F2`.
	Size pulumi.StringOutput `pulumi:"size"`
	// The ID of the Image which this Virtual Machine should be created from. Changing this forces a new resource to be created.
	SourceImageId pulumi.StringPtrOutput `pulumi:"sourceImageId"`
	// A `sourceImageReference` block as defined below. Changing this forces a new resource to be created.
	SourceImageReference LinuxVirtualMachineSourceImageReferencePtrOutput `pulumi:"sourceImageReference"`
	// A mapping of tags which should be assigned to this Virtual Machine.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A 128-bit identifier which uniquely identifies this Virtual Machine.
	VirtualMachineId pulumi.StringOutput `pulumi:"virtualMachineId"`
	// Specifies the Orchestrated Virtual Machine Scale Set that this Virtual Machine should be created within. Changing this forces a new resource to be created.
	VirtualMachineScaleSetId pulumi.StringPtrOutput `pulumi:"virtualMachineScaleSetId"`
	// The Zone in which this Virtual Machine should be created. Changing this forces a new resource to be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewLinuxVirtualMachine registers a new resource with the given unique name, arguments, and options.
func NewLinuxVirtualMachine(ctx *pulumi.Context,
	name string, args *LinuxVirtualMachineArgs, opts ...pulumi.ResourceOption) (*LinuxVirtualMachine, error) {
	if args == nil || args.AdminUsername == nil {
		return nil, errors.New("missing required argument 'AdminUsername'")
	}
	if args == nil || args.NetworkInterfaceIds == nil {
		return nil, errors.New("missing required argument 'NetworkInterfaceIds'")
	}
	if args == nil || args.OsDisk == nil {
		return nil, errors.New("missing required argument 'OsDisk'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Size == nil {
		return nil, errors.New("missing required argument 'Size'")
	}
	if args == nil {
		args = &LinuxVirtualMachineArgs{}
	}
	var resource LinuxVirtualMachine
	err := ctx.RegisterResource("azure:compute/linuxVirtualMachine:LinuxVirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinuxVirtualMachine gets an existing LinuxVirtualMachine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinuxVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinuxVirtualMachineState, opts ...pulumi.ResourceOption) (*LinuxVirtualMachine, error) {
	var resource LinuxVirtualMachine
	err := ctx.ReadResource("azure:compute/linuxVirtualMachine:LinuxVirtualMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinuxVirtualMachine resources.
type linuxVirtualMachineState struct {
	// A `additionalCapabilities` block as defined below.
	AdditionalCapabilities *LinuxVirtualMachineAdditionalCapabilities `pulumi:"additionalCapabilities"`
	// The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
	AdminPassword *string `pulumi:"adminPassword"`
	// One or more `adminSshKey` blocks as defined below.
	AdminSshKeys []LinuxVirtualMachineAdminSshKey `pulumi:"adminSshKeys"`
	// The username of the local administrator used for the Virtual Machine. Changing this forces a new resource to be created.
	AdminUsername *string `pulumi:"adminUsername"`
	// Should Extension Operations be allowed on this Virtual Machine?
	AllowExtensionOperations *bool `pulumi:"allowExtensionOperations"`
	// Specifies the ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
	AvailabilitySetId *string `pulumi:"availabilitySetId"`
	// A `bootDiagnostics` block as defined below.
	BootDiagnostics *LinuxVirtualMachineBootDiagnostics `pulumi:"bootDiagnostics"`
	// Specifies the Hostname which should be used for this Virtual Machine. If unspecified this defaults to the value for the `name` field. If the value of the `name` field is not a valid `computerName`, then you must specify `computerName`. Changing this forces a new resource to be created.
	ComputerName *string `pulumi:"computerName"`
	// The Base64-Encoded Custom Data which should be used for this Virtual Machine. Changing this forces a new resource to be created.
	CustomData *string `pulumi:"customData"`
	// The ID of a Dedicated Host where this machine should be run on. Changing this forces a new resource to be created.
	DedicatedHostId *string `pulumi:"dedicatedHostId"`
	// Should Password Authentication be disabled on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
	DisablePasswordAuthentication *bool `pulumi:"disablePasswordAuthentication"`
	// Should all of the disks (including the temp disk) attached to this Virtual Machine be encrypted by enabling Encryption at Host?
	EncryptionAtHostEnabled *bool `pulumi:"encryptionAtHostEnabled"`
	// Specifies what should happen when the Virtual Machine is evicted for price reasons when using a Spot instance. At this time the only supported value is `Deallocate`. Changing this forces a new resource to be created.
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// An `identity` block as defined below.
	Identity *LinuxVirtualMachineIdentity `pulumi:"identity"`
	// The Azure location where the Linux Virtual Machine should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The maximum price you're willing to pay for this Virtual Machine, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machine will be evicted using the `evictionPolicy`. Defaults to `-1`, which means that the Virtual Machine should not be evicted for price reasons.
	MaxBidPrice *float64 `pulumi:"maxBidPrice"`
	// The name of the Linux Virtual Machine. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// . A list of Network Interface ID's which should be attached to this Virtual Machine. The first Network Interface ID in this list will be the Primary Network Interface on the Virtual Machine.
	NetworkInterfaceIds []string `pulumi:"networkInterfaceIds"`
	// A `osDisk` block as defined below.
	OsDisk *LinuxVirtualMachineOsDisk `pulumi:"osDisk"`
	// A `plan` block as defined below. Changing this forces a new resource to be created.
	Plan *LinuxVirtualMachinePlan `pulumi:"plan"`
	// Specifies the priority of this Virtual Machine. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
	Priority *string `pulumi:"priority"`
	// The Primary Private IP Address assigned to this Virtual Machine.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// A list of Private IP Addresses assigned to this Virtual Machine.
	PrivateIpAddresses []string `pulumi:"privateIpAddresses"`
	// Should the Azure VM Agent be provisioned on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
	ProvisionVmAgent *bool `pulumi:"provisionVmAgent"`
	// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId *string `pulumi:"proximityPlacementGroupId"`
	// The Primary Public IP Address assigned to this Virtual Machine.
	PublicIpAddress *string `pulumi:"publicIpAddress"`
	// A list of the Public IP Addresses assigned to this Virtual Machine.
	PublicIpAddresses []string `pulumi:"publicIpAddresses"`
	// The name of the Resource Group in which the Linux Virtual Machine should be exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// One or more `secret` blocks as defined below.
	Secrets []LinuxVirtualMachineSecret `pulumi:"secrets"`
	// The SKU which should be used for this Virtual Machine, such as `Standard_F2`.
	Size *string `pulumi:"size"`
	// The ID of the Image which this Virtual Machine should be created from. Changing this forces a new resource to be created.
	SourceImageId *string `pulumi:"sourceImageId"`
	// A `sourceImageReference` block as defined below. Changing this forces a new resource to be created.
	SourceImageReference *LinuxVirtualMachineSourceImageReference `pulumi:"sourceImageReference"`
	// A mapping of tags which should be assigned to this Virtual Machine.
	Tags map[string]string `pulumi:"tags"`
	// A 128-bit identifier which uniquely identifies this Virtual Machine.
	VirtualMachineId *string `pulumi:"virtualMachineId"`
	// Specifies the Orchestrated Virtual Machine Scale Set that this Virtual Machine should be created within. Changing this forces a new resource to be created.
	VirtualMachineScaleSetId *string `pulumi:"virtualMachineScaleSetId"`
	// The Zone in which this Virtual Machine should be created. Changing this forces a new resource to be created.
	Zone *string `pulumi:"zone"`
}

type LinuxVirtualMachineState struct {
	// A `additionalCapabilities` block as defined below.
	AdditionalCapabilities LinuxVirtualMachineAdditionalCapabilitiesPtrInput
	// The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
	AdminPassword pulumi.StringPtrInput
	// One or more `adminSshKey` blocks as defined below.
	AdminSshKeys LinuxVirtualMachineAdminSshKeyArrayInput
	// The username of the local administrator used for the Virtual Machine. Changing this forces a new resource to be created.
	AdminUsername pulumi.StringPtrInput
	// Should Extension Operations be allowed on this Virtual Machine?
	AllowExtensionOperations pulumi.BoolPtrInput
	// Specifies the ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
	AvailabilitySetId pulumi.StringPtrInput
	// A `bootDiagnostics` block as defined below.
	BootDiagnostics LinuxVirtualMachineBootDiagnosticsPtrInput
	// Specifies the Hostname which should be used for this Virtual Machine. If unspecified this defaults to the value for the `name` field. If the value of the `name` field is not a valid `computerName`, then you must specify `computerName`. Changing this forces a new resource to be created.
	ComputerName pulumi.StringPtrInput
	// The Base64-Encoded Custom Data which should be used for this Virtual Machine. Changing this forces a new resource to be created.
	CustomData pulumi.StringPtrInput
	// The ID of a Dedicated Host where this machine should be run on. Changing this forces a new resource to be created.
	DedicatedHostId pulumi.StringPtrInput
	// Should Password Authentication be disabled on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
	DisablePasswordAuthentication pulumi.BoolPtrInput
	// Should all of the disks (including the temp disk) attached to this Virtual Machine be encrypted by enabling Encryption at Host?
	EncryptionAtHostEnabled pulumi.BoolPtrInput
	// Specifies what should happen when the Virtual Machine is evicted for price reasons when using a Spot instance. At this time the only supported value is `Deallocate`. Changing this forces a new resource to be created.
	EvictionPolicy pulumi.StringPtrInput
	// An `identity` block as defined below.
	Identity LinuxVirtualMachineIdentityPtrInput
	// The Azure location where the Linux Virtual Machine should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The maximum price you're willing to pay for this Virtual Machine, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machine will be evicted using the `evictionPolicy`. Defaults to `-1`, which means that the Virtual Machine should not be evicted for price reasons.
	MaxBidPrice pulumi.Float64PtrInput
	// The name of the Linux Virtual Machine. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// . A list of Network Interface ID's which should be attached to this Virtual Machine. The first Network Interface ID in this list will be the Primary Network Interface on the Virtual Machine.
	NetworkInterfaceIds pulumi.StringArrayInput
	// A `osDisk` block as defined below.
	OsDisk LinuxVirtualMachineOsDiskPtrInput
	// A `plan` block as defined below. Changing this forces a new resource to be created.
	Plan LinuxVirtualMachinePlanPtrInput
	// Specifies the priority of this Virtual Machine. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
	Priority pulumi.StringPtrInput
	// The Primary Private IP Address assigned to this Virtual Machine.
	PrivateIpAddress pulumi.StringPtrInput
	// A list of Private IP Addresses assigned to this Virtual Machine.
	PrivateIpAddresses pulumi.StringArrayInput
	// Should the Azure VM Agent be provisioned on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
	ProvisionVmAgent pulumi.BoolPtrInput
	// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId pulumi.StringPtrInput
	// The Primary Public IP Address assigned to this Virtual Machine.
	PublicIpAddress pulumi.StringPtrInput
	// A list of the Public IP Addresses assigned to this Virtual Machine.
	PublicIpAddresses pulumi.StringArrayInput
	// The name of the Resource Group in which the Linux Virtual Machine should be exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// One or more `secret` blocks as defined below.
	Secrets LinuxVirtualMachineSecretArrayInput
	// The SKU which should be used for this Virtual Machine, such as `Standard_F2`.
	Size pulumi.StringPtrInput
	// The ID of the Image which this Virtual Machine should be created from. Changing this forces a new resource to be created.
	SourceImageId pulumi.StringPtrInput
	// A `sourceImageReference` block as defined below. Changing this forces a new resource to be created.
	SourceImageReference LinuxVirtualMachineSourceImageReferencePtrInput
	// A mapping of tags which should be assigned to this Virtual Machine.
	Tags pulumi.StringMapInput
	// A 128-bit identifier which uniquely identifies this Virtual Machine.
	VirtualMachineId pulumi.StringPtrInput
	// Specifies the Orchestrated Virtual Machine Scale Set that this Virtual Machine should be created within. Changing this forces a new resource to be created.
	VirtualMachineScaleSetId pulumi.StringPtrInput
	// The Zone in which this Virtual Machine should be created. Changing this forces a new resource to be created.
	Zone pulumi.StringPtrInput
}

func (LinuxVirtualMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*linuxVirtualMachineState)(nil)).Elem()
}

type linuxVirtualMachineArgs struct {
	// A `additionalCapabilities` block as defined below.
	AdditionalCapabilities *LinuxVirtualMachineAdditionalCapabilities `pulumi:"additionalCapabilities"`
	// The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
	AdminPassword *string `pulumi:"adminPassword"`
	// One or more `adminSshKey` blocks as defined below.
	AdminSshKeys []LinuxVirtualMachineAdminSshKey `pulumi:"adminSshKeys"`
	// The username of the local administrator used for the Virtual Machine. Changing this forces a new resource to be created.
	AdminUsername string `pulumi:"adminUsername"`
	// Should Extension Operations be allowed on this Virtual Machine?
	AllowExtensionOperations *bool `pulumi:"allowExtensionOperations"`
	// Specifies the ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
	AvailabilitySetId *string `pulumi:"availabilitySetId"`
	// A `bootDiagnostics` block as defined below.
	BootDiagnostics *LinuxVirtualMachineBootDiagnostics `pulumi:"bootDiagnostics"`
	// Specifies the Hostname which should be used for this Virtual Machine. If unspecified this defaults to the value for the `name` field. If the value of the `name` field is not a valid `computerName`, then you must specify `computerName`. Changing this forces a new resource to be created.
	ComputerName *string `pulumi:"computerName"`
	// The Base64-Encoded Custom Data which should be used for this Virtual Machine. Changing this forces a new resource to be created.
	CustomData *string `pulumi:"customData"`
	// The ID of a Dedicated Host where this machine should be run on. Changing this forces a new resource to be created.
	DedicatedHostId *string `pulumi:"dedicatedHostId"`
	// Should Password Authentication be disabled on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
	DisablePasswordAuthentication *bool `pulumi:"disablePasswordAuthentication"`
	// Should all of the disks (including the temp disk) attached to this Virtual Machine be encrypted by enabling Encryption at Host?
	EncryptionAtHostEnabled *bool `pulumi:"encryptionAtHostEnabled"`
	// Specifies what should happen when the Virtual Machine is evicted for price reasons when using a Spot instance. At this time the only supported value is `Deallocate`. Changing this forces a new resource to be created.
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// An `identity` block as defined below.
	Identity *LinuxVirtualMachineIdentity `pulumi:"identity"`
	// The Azure location where the Linux Virtual Machine should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The maximum price you're willing to pay for this Virtual Machine, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machine will be evicted using the `evictionPolicy`. Defaults to `-1`, which means that the Virtual Machine should not be evicted for price reasons.
	MaxBidPrice *float64 `pulumi:"maxBidPrice"`
	// The name of the Linux Virtual Machine. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// . A list of Network Interface ID's which should be attached to this Virtual Machine. The first Network Interface ID in this list will be the Primary Network Interface on the Virtual Machine.
	NetworkInterfaceIds []string `pulumi:"networkInterfaceIds"`
	// A `osDisk` block as defined below.
	OsDisk LinuxVirtualMachineOsDisk `pulumi:"osDisk"`
	// A `plan` block as defined below. Changing this forces a new resource to be created.
	Plan *LinuxVirtualMachinePlan `pulumi:"plan"`
	// Specifies the priority of this Virtual Machine. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
	Priority *string `pulumi:"priority"`
	// Should the Azure VM Agent be provisioned on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
	ProvisionVmAgent *bool `pulumi:"provisionVmAgent"`
	// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId *string `pulumi:"proximityPlacementGroupId"`
	// The name of the Resource Group in which the Linux Virtual Machine should be exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// One or more `secret` blocks as defined below.
	Secrets []LinuxVirtualMachineSecret `pulumi:"secrets"`
	// The SKU which should be used for this Virtual Machine, such as `Standard_F2`.
	Size string `pulumi:"size"`
	// The ID of the Image which this Virtual Machine should be created from. Changing this forces a new resource to be created.
	SourceImageId *string `pulumi:"sourceImageId"`
	// A `sourceImageReference` block as defined below. Changing this forces a new resource to be created.
	SourceImageReference *LinuxVirtualMachineSourceImageReference `pulumi:"sourceImageReference"`
	// A mapping of tags which should be assigned to this Virtual Machine.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Orchestrated Virtual Machine Scale Set that this Virtual Machine should be created within. Changing this forces a new resource to be created.
	VirtualMachineScaleSetId *string `pulumi:"virtualMachineScaleSetId"`
	// The Zone in which this Virtual Machine should be created. Changing this forces a new resource to be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a LinuxVirtualMachine resource.
type LinuxVirtualMachineArgs struct {
	// A `additionalCapabilities` block as defined below.
	AdditionalCapabilities LinuxVirtualMachineAdditionalCapabilitiesPtrInput
	// The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
	AdminPassword pulumi.StringPtrInput
	// One or more `adminSshKey` blocks as defined below.
	AdminSshKeys LinuxVirtualMachineAdminSshKeyArrayInput
	// The username of the local administrator used for the Virtual Machine. Changing this forces a new resource to be created.
	AdminUsername pulumi.StringInput
	// Should Extension Operations be allowed on this Virtual Machine?
	AllowExtensionOperations pulumi.BoolPtrInput
	// Specifies the ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
	AvailabilitySetId pulumi.StringPtrInput
	// A `bootDiagnostics` block as defined below.
	BootDiagnostics LinuxVirtualMachineBootDiagnosticsPtrInput
	// Specifies the Hostname which should be used for this Virtual Machine. If unspecified this defaults to the value for the `name` field. If the value of the `name` field is not a valid `computerName`, then you must specify `computerName`. Changing this forces a new resource to be created.
	ComputerName pulumi.StringPtrInput
	// The Base64-Encoded Custom Data which should be used for this Virtual Machine. Changing this forces a new resource to be created.
	CustomData pulumi.StringPtrInput
	// The ID of a Dedicated Host where this machine should be run on. Changing this forces a new resource to be created.
	DedicatedHostId pulumi.StringPtrInput
	// Should Password Authentication be disabled on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
	DisablePasswordAuthentication pulumi.BoolPtrInput
	// Should all of the disks (including the temp disk) attached to this Virtual Machine be encrypted by enabling Encryption at Host?
	EncryptionAtHostEnabled pulumi.BoolPtrInput
	// Specifies what should happen when the Virtual Machine is evicted for price reasons when using a Spot instance. At this time the only supported value is `Deallocate`. Changing this forces a new resource to be created.
	EvictionPolicy pulumi.StringPtrInput
	// An `identity` block as defined below.
	Identity LinuxVirtualMachineIdentityPtrInput
	// The Azure location where the Linux Virtual Machine should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The maximum price you're willing to pay for this Virtual Machine, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machine will be evicted using the `evictionPolicy`. Defaults to `-1`, which means that the Virtual Machine should not be evicted for price reasons.
	MaxBidPrice pulumi.Float64PtrInput
	// The name of the Linux Virtual Machine. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// . A list of Network Interface ID's which should be attached to this Virtual Machine. The first Network Interface ID in this list will be the Primary Network Interface on the Virtual Machine.
	NetworkInterfaceIds pulumi.StringArrayInput
	// A `osDisk` block as defined below.
	OsDisk LinuxVirtualMachineOsDiskInput
	// A `plan` block as defined below. Changing this forces a new resource to be created.
	Plan LinuxVirtualMachinePlanPtrInput
	// Specifies the priority of this Virtual Machine. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
	Priority pulumi.StringPtrInput
	// Should the Azure VM Agent be provisioned on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
	ProvisionVmAgent pulumi.BoolPtrInput
	// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId pulumi.StringPtrInput
	// The name of the Resource Group in which the Linux Virtual Machine should be exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// One or more `secret` blocks as defined below.
	Secrets LinuxVirtualMachineSecretArrayInput
	// The SKU which should be used for this Virtual Machine, such as `Standard_F2`.
	Size pulumi.StringInput
	// The ID of the Image which this Virtual Machine should be created from. Changing this forces a new resource to be created.
	SourceImageId pulumi.StringPtrInput
	// A `sourceImageReference` block as defined below. Changing this forces a new resource to be created.
	SourceImageReference LinuxVirtualMachineSourceImageReferencePtrInput
	// A mapping of tags which should be assigned to this Virtual Machine.
	Tags pulumi.StringMapInput
	// Specifies the Orchestrated Virtual Machine Scale Set that this Virtual Machine should be created within. Changing this forces a new resource to be created.
	VirtualMachineScaleSetId pulumi.StringPtrInput
	// The Zone in which this Virtual Machine should be created. Changing this forces a new resource to be created.
	Zone pulumi.StringPtrInput
}

func (LinuxVirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linuxVirtualMachineArgs)(nil)).Elem()
}

type LinuxVirtualMachineInput interface {
	pulumi.Input

	ToLinuxVirtualMachineOutput() LinuxVirtualMachineOutput
	ToLinuxVirtualMachineOutputWithContext(ctx context.Context) LinuxVirtualMachineOutput
}

func (LinuxVirtualMachine) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxVirtualMachine)(nil)).Elem()
}

func (i LinuxVirtualMachine) ToLinuxVirtualMachineOutput() LinuxVirtualMachineOutput {
	return i.ToLinuxVirtualMachineOutputWithContext(context.Background())
}

func (i LinuxVirtualMachine) ToLinuxVirtualMachineOutputWithContext(ctx context.Context) LinuxVirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxVirtualMachineOutput)
}

type LinuxVirtualMachineOutput struct {
	*pulumi.OutputState
}

func (LinuxVirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxVirtualMachineOutput)(nil)).Elem()
}

func (o LinuxVirtualMachineOutput) ToLinuxVirtualMachineOutput() LinuxVirtualMachineOutput {
	return o
}

func (o LinuxVirtualMachineOutput) ToLinuxVirtualMachineOutputWithContext(ctx context.Context) LinuxVirtualMachineOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LinuxVirtualMachineOutput{})
}

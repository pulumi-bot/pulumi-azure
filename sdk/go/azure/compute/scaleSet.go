// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a virtual machine scale set.
//
// ## Disclaimers
//
// > **Note:** The `compute.ScaleSet` resource has been superseded by the `compute.LinuxVirtualMachineScaleSet` and `compute.WindowsVirtualMachineScaleSet` resources. The existing `compute.ScaleSet` resource will continue to be available throughout the 2.x releases however is in a feature-frozen state to maintain compatibility - new functionality will instead be added to the `compute.LinuxVirtualMachineScaleSet` and `compute.WindowsVirtualMachineScaleSet` resources.
//
// ## Example Usage
// ## Example of storageProfileImageReference with id
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleImage, err := compute.NewImage(ctx, "exampleImage", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewScaleSet(ctx, "exampleScaleSet", &compute.ScaleSetArgs{
// 			StorageProfileImageReference: &compute.ScaleSetStorageProfileImageReferenceArgs{
// 				Id: exampleImage.ID(),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Virtual Machine Scale Sets can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:compute/scaleSet:ScaleSet scaleset1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachineScaleSets/scaleset1
// ```
type ScaleSet struct {
	pulumi.CustomResourceState

	// Automatic OS patches can be applied by Azure to your scaleset. This is particularly useful when `upgradePolicyMode` is set to `Rolling`. Defaults to `false`.
	AutomaticOsUpgrade pulumi.BoolPtrOutput `pulumi:"automaticOsUpgrade"`
	// A boot diagnostics profile block as referenced below.
	BootDiagnostics ScaleSetBootDiagnosticsPtrOutput `pulumi:"bootDiagnostics"`
	// Specifies the eviction policy for Virtual Machines in this Scale Set. Possible values are `Deallocate` and `Delete`.
	EvictionPolicy pulumi.StringPtrOutput `pulumi:"evictionPolicy"`
	// Can be specified multiple times to add extension profiles to the scale set. Each `extension` block supports the fields documented below.
	Extensions ScaleSetExtensionArrayOutput `pulumi:"extensions"`
	// Specifies the identifier for the load balancer health probe. Required when using `Rolling` as your `upgradePolicyMode`.
	HealthProbeId pulumi.StringPtrOutput `pulumi:"healthProbeId"`
	Identity      ScaleSetIdentityOutput `pulumi:"identity"`
	// Specifies the Windows OS license type. If supplied, the only allowed values are `Windows_Client` and `Windows_Server`.
	LicenseType pulumi.StringOutput `pulumi:"licenseType"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the virtual machine scale set resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A collection of network profile block as documented below.
	NetworkProfiles ScaleSetNetworkProfileArrayOutput `pulumi:"networkProfiles"`
	// A Virtual Machine OS Profile block as documented below.
	OsProfile ScaleSetOsProfileOutput `pulumi:"osProfile"`
	// A Linux config block as documented below.
	OsProfileLinuxConfig ScaleSetOsProfileLinuxConfigOutput `pulumi:"osProfileLinuxConfig"`
	// A collection of Secret blocks as documented below.
	OsProfileSecrets ScaleSetOsProfileSecretArrayOutput `pulumi:"osProfileSecrets"`
	// A Windows config block as documented below.
	OsProfileWindowsConfig ScaleSetOsProfileWindowsConfigPtrOutput `pulumi:"osProfileWindowsConfig"`
	// Specifies whether the virtual machine scale set should be overprovisioned. Defaults to `true`.
	Overprovision pulumi.BoolPtrOutput `pulumi:"overprovision"`
	// A plan block as documented below.
	Plan ScaleSetPlanPtrOutput `pulumi:"plan"`
	// Specifies the priority for the Virtual Machines in the Scale Set. Defaults to `Regular`. Possible values are `Low` and `Regular`.
	Priority pulumi.StringPtrOutput `pulumi:"priority"`
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
	ProximityPlacementGroupId pulumi.StringPtrOutput `pulumi:"proximityPlacementGroupId"`
	// The name of the resource group in which to create the virtual machine scale set. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `rollingUpgradePolicy` block as defined below. This is only applicable when the `upgradePolicyMode` is `Rolling`.
	RollingUpgradePolicy ScaleSetRollingUpgradePolicyPtrOutput `pulumi:"rollingUpgradePolicy"`
	// Specifies whether the scale set is limited to a single placement group with a maximum size of 100 virtual machines. If set to false, managed disks must be used. Default is true. Changing this forces a new resource to be created. See [documentation](http://docs.microsoft.com/en-us/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-placement-groups) for more information.
	SinglePlacementGroup pulumi.BoolPtrOutput `pulumi:"singlePlacementGroup"`
	// A sku block as documented below.
	Sku ScaleSetSkuOutput `pulumi:"sku"`
	// A storage profile data disk block as documented below
	StorageProfileDataDisks ScaleSetStorageProfileDataDiskArrayOutput `pulumi:"storageProfileDataDisks"`
	// A storage profile image reference block as documented below.
	StorageProfileImageReference ScaleSetStorageProfileImageReferenceOutput `pulumi:"storageProfileImageReference"`
	// A storage profile os disk block as documented below
	StorageProfileOsDisk ScaleSetStorageProfileOsDiskOutput `pulumi:"storageProfileOsDisk"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the mode of an upgrade to virtual machines in the scale set. Possible values, `Rolling`, `Manual`, or `Automatic`. When choosing `Rolling`, you will need to set a health probe.
	UpgradePolicyMode pulumi.StringOutput `pulumi:"upgradePolicyMode"`
	// A collection of availability zones to spread the Virtual Machines over.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewScaleSet registers a new resource with the given unique name, arguments, and options.
func NewScaleSet(ctx *pulumi.Context,
	name string, args *ScaleSetArgs, opts ...pulumi.ResourceOption) (*ScaleSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkProfiles == nil {
		return nil, errors.New("invalid value for required argument 'NetworkProfiles'")
	}
	if args.OsProfile == nil {
		return nil, errors.New("invalid value for required argument 'OsProfile'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.StorageProfileOsDisk == nil {
		return nil, errors.New("invalid value for required argument 'StorageProfileOsDisk'")
	}
	if args.UpgradePolicyMode == nil {
		return nil, errors.New("invalid value for required argument 'UpgradePolicyMode'")
	}
	var resource ScaleSet
	err := ctx.RegisterResource("azure:compute/scaleSet:ScaleSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScaleSet gets an existing ScaleSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScaleSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScaleSetState, opts ...pulumi.ResourceOption) (*ScaleSet, error) {
	var resource ScaleSet
	err := ctx.ReadResource("azure:compute/scaleSet:ScaleSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScaleSet resources.
type scaleSetState struct {
	// Automatic OS patches can be applied by Azure to your scaleset. This is particularly useful when `upgradePolicyMode` is set to `Rolling`. Defaults to `false`.
	AutomaticOsUpgrade *bool `pulumi:"automaticOsUpgrade"`
	// A boot diagnostics profile block as referenced below.
	BootDiagnostics *ScaleSetBootDiagnostics `pulumi:"bootDiagnostics"`
	// Specifies the eviction policy for Virtual Machines in this Scale Set. Possible values are `Deallocate` and `Delete`.
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// Can be specified multiple times to add extension profiles to the scale set. Each `extension` block supports the fields documented below.
	Extensions []ScaleSetExtension `pulumi:"extensions"`
	// Specifies the identifier for the load balancer health probe. Required when using `Rolling` as your `upgradePolicyMode`.
	HealthProbeId *string           `pulumi:"healthProbeId"`
	Identity      *ScaleSetIdentity `pulumi:"identity"`
	// Specifies the Windows OS license type. If supplied, the only allowed values are `Windows_Client` and `Windows_Server`.
	LicenseType *string `pulumi:"licenseType"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the virtual machine scale set resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A collection of network profile block as documented below.
	NetworkProfiles []ScaleSetNetworkProfile `pulumi:"networkProfiles"`
	// A Virtual Machine OS Profile block as documented below.
	OsProfile *ScaleSetOsProfile `pulumi:"osProfile"`
	// A Linux config block as documented below.
	OsProfileLinuxConfig *ScaleSetOsProfileLinuxConfig `pulumi:"osProfileLinuxConfig"`
	// A collection of Secret blocks as documented below.
	OsProfileSecrets []ScaleSetOsProfileSecret `pulumi:"osProfileSecrets"`
	// A Windows config block as documented below.
	OsProfileWindowsConfig *ScaleSetOsProfileWindowsConfig `pulumi:"osProfileWindowsConfig"`
	// Specifies whether the virtual machine scale set should be overprovisioned. Defaults to `true`.
	Overprovision *bool `pulumi:"overprovision"`
	// A plan block as documented below.
	Plan *ScaleSetPlan `pulumi:"plan"`
	// Specifies the priority for the Virtual Machines in the Scale Set. Defaults to `Regular`. Possible values are `Low` and `Regular`.
	Priority *string `pulumi:"priority"`
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
	ProximityPlacementGroupId *string `pulumi:"proximityPlacementGroupId"`
	// The name of the resource group in which to create the virtual machine scale set. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `rollingUpgradePolicy` block as defined below. This is only applicable when the `upgradePolicyMode` is `Rolling`.
	RollingUpgradePolicy *ScaleSetRollingUpgradePolicy `pulumi:"rollingUpgradePolicy"`
	// Specifies whether the scale set is limited to a single placement group with a maximum size of 100 virtual machines. If set to false, managed disks must be used. Default is true. Changing this forces a new resource to be created. See [documentation](http://docs.microsoft.com/en-us/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-placement-groups) for more information.
	SinglePlacementGroup *bool `pulumi:"singlePlacementGroup"`
	// A sku block as documented below.
	Sku *ScaleSetSku `pulumi:"sku"`
	// A storage profile data disk block as documented below
	StorageProfileDataDisks []ScaleSetStorageProfileDataDisk `pulumi:"storageProfileDataDisks"`
	// A storage profile image reference block as documented below.
	StorageProfileImageReference *ScaleSetStorageProfileImageReference `pulumi:"storageProfileImageReference"`
	// A storage profile os disk block as documented below
	StorageProfileOsDisk *ScaleSetStorageProfileOsDisk `pulumi:"storageProfileOsDisk"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the mode of an upgrade to virtual machines in the scale set. Possible values, `Rolling`, `Manual`, or `Automatic`. When choosing `Rolling`, you will need to set a health probe.
	UpgradePolicyMode *string `pulumi:"upgradePolicyMode"`
	// A collection of availability zones to spread the Virtual Machines over.
	Zones []string `pulumi:"zones"`
}

type ScaleSetState struct {
	// Automatic OS patches can be applied by Azure to your scaleset. This is particularly useful when `upgradePolicyMode` is set to `Rolling`. Defaults to `false`.
	AutomaticOsUpgrade pulumi.BoolPtrInput
	// A boot diagnostics profile block as referenced below.
	BootDiagnostics ScaleSetBootDiagnosticsPtrInput
	// Specifies the eviction policy for Virtual Machines in this Scale Set. Possible values are `Deallocate` and `Delete`.
	EvictionPolicy pulumi.StringPtrInput
	// Can be specified multiple times to add extension profiles to the scale set. Each `extension` block supports the fields documented below.
	Extensions ScaleSetExtensionArrayInput
	// Specifies the identifier for the load balancer health probe. Required when using `Rolling` as your `upgradePolicyMode`.
	HealthProbeId pulumi.StringPtrInput
	Identity      ScaleSetIdentityPtrInput
	// Specifies the Windows OS license type. If supplied, the only allowed values are `Windows_Client` and `Windows_Server`.
	LicenseType pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the virtual machine scale set resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A collection of network profile block as documented below.
	NetworkProfiles ScaleSetNetworkProfileArrayInput
	// A Virtual Machine OS Profile block as documented below.
	OsProfile ScaleSetOsProfilePtrInput
	// A Linux config block as documented below.
	OsProfileLinuxConfig ScaleSetOsProfileLinuxConfigPtrInput
	// A collection of Secret blocks as documented below.
	OsProfileSecrets ScaleSetOsProfileSecretArrayInput
	// A Windows config block as documented below.
	OsProfileWindowsConfig ScaleSetOsProfileWindowsConfigPtrInput
	// Specifies whether the virtual machine scale set should be overprovisioned. Defaults to `true`.
	Overprovision pulumi.BoolPtrInput
	// A plan block as documented below.
	Plan ScaleSetPlanPtrInput
	// Specifies the priority for the Virtual Machines in the Scale Set. Defaults to `Regular`. Possible values are `Low` and `Regular`.
	Priority pulumi.StringPtrInput
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
	ProximityPlacementGroupId pulumi.StringPtrInput
	// The name of the resource group in which to create the virtual machine scale set. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `rollingUpgradePolicy` block as defined below. This is only applicable when the `upgradePolicyMode` is `Rolling`.
	RollingUpgradePolicy ScaleSetRollingUpgradePolicyPtrInput
	// Specifies whether the scale set is limited to a single placement group with a maximum size of 100 virtual machines. If set to false, managed disks must be used. Default is true. Changing this forces a new resource to be created. See [documentation](http://docs.microsoft.com/en-us/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-placement-groups) for more information.
	SinglePlacementGroup pulumi.BoolPtrInput
	// A sku block as documented below.
	Sku ScaleSetSkuPtrInput
	// A storage profile data disk block as documented below
	StorageProfileDataDisks ScaleSetStorageProfileDataDiskArrayInput
	// A storage profile image reference block as documented below.
	StorageProfileImageReference ScaleSetStorageProfileImageReferencePtrInput
	// A storage profile os disk block as documented below
	StorageProfileOsDisk ScaleSetStorageProfileOsDiskPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the mode of an upgrade to virtual machines in the scale set. Possible values, `Rolling`, `Manual`, or `Automatic`. When choosing `Rolling`, you will need to set a health probe.
	UpgradePolicyMode pulumi.StringPtrInput
	// A collection of availability zones to spread the Virtual Machines over.
	Zones pulumi.StringArrayInput
}

func (ScaleSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*scaleSetState)(nil)).Elem()
}

type scaleSetArgs struct {
	// Automatic OS patches can be applied by Azure to your scaleset. This is particularly useful when `upgradePolicyMode` is set to `Rolling`. Defaults to `false`.
	AutomaticOsUpgrade *bool `pulumi:"automaticOsUpgrade"`
	// A boot diagnostics profile block as referenced below.
	BootDiagnostics *ScaleSetBootDiagnostics `pulumi:"bootDiagnostics"`
	// Specifies the eviction policy for Virtual Machines in this Scale Set. Possible values are `Deallocate` and `Delete`.
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// Can be specified multiple times to add extension profiles to the scale set. Each `extension` block supports the fields documented below.
	Extensions []ScaleSetExtension `pulumi:"extensions"`
	// Specifies the identifier for the load balancer health probe. Required when using `Rolling` as your `upgradePolicyMode`.
	HealthProbeId *string           `pulumi:"healthProbeId"`
	Identity      *ScaleSetIdentity `pulumi:"identity"`
	// Specifies the Windows OS license type. If supplied, the only allowed values are `Windows_Client` and `Windows_Server`.
	LicenseType *string `pulumi:"licenseType"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the virtual machine scale set resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A collection of network profile block as documented below.
	NetworkProfiles []ScaleSetNetworkProfile `pulumi:"networkProfiles"`
	// A Virtual Machine OS Profile block as documented below.
	OsProfile ScaleSetOsProfile `pulumi:"osProfile"`
	// A Linux config block as documented below.
	OsProfileLinuxConfig *ScaleSetOsProfileLinuxConfig `pulumi:"osProfileLinuxConfig"`
	// A collection of Secret blocks as documented below.
	OsProfileSecrets []ScaleSetOsProfileSecret `pulumi:"osProfileSecrets"`
	// A Windows config block as documented below.
	OsProfileWindowsConfig *ScaleSetOsProfileWindowsConfig `pulumi:"osProfileWindowsConfig"`
	// Specifies whether the virtual machine scale set should be overprovisioned. Defaults to `true`.
	Overprovision *bool `pulumi:"overprovision"`
	// A plan block as documented below.
	Plan *ScaleSetPlan `pulumi:"plan"`
	// Specifies the priority for the Virtual Machines in the Scale Set. Defaults to `Regular`. Possible values are `Low` and `Regular`.
	Priority *string `pulumi:"priority"`
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
	ProximityPlacementGroupId *string `pulumi:"proximityPlacementGroupId"`
	// The name of the resource group in which to create the virtual machine scale set. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `rollingUpgradePolicy` block as defined below. This is only applicable when the `upgradePolicyMode` is `Rolling`.
	RollingUpgradePolicy *ScaleSetRollingUpgradePolicy `pulumi:"rollingUpgradePolicy"`
	// Specifies whether the scale set is limited to a single placement group with a maximum size of 100 virtual machines. If set to false, managed disks must be used. Default is true. Changing this forces a new resource to be created. See [documentation](http://docs.microsoft.com/en-us/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-placement-groups) for more information.
	SinglePlacementGroup *bool `pulumi:"singlePlacementGroup"`
	// A sku block as documented below.
	Sku ScaleSetSku `pulumi:"sku"`
	// A storage profile data disk block as documented below
	StorageProfileDataDisks []ScaleSetStorageProfileDataDisk `pulumi:"storageProfileDataDisks"`
	// A storage profile image reference block as documented below.
	StorageProfileImageReference *ScaleSetStorageProfileImageReference `pulumi:"storageProfileImageReference"`
	// A storage profile os disk block as documented below
	StorageProfileOsDisk ScaleSetStorageProfileOsDisk `pulumi:"storageProfileOsDisk"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the mode of an upgrade to virtual machines in the scale set. Possible values, `Rolling`, `Manual`, or `Automatic`. When choosing `Rolling`, you will need to set a health probe.
	UpgradePolicyMode string `pulumi:"upgradePolicyMode"`
	// A collection of availability zones to spread the Virtual Machines over.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a ScaleSet resource.
type ScaleSetArgs struct {
	// Automatic OS patches can be applied by Azure to your scaleset. This is particularly useful when `upgradePolicyMode` is set to `Rolling`. Defaults to `false`.
	AutomaticOsUpgrade pulumi.BoolPtrInput
	// A boot diagnostics profile block as referenced below.
	BootDiagnostics ScaleSetBootDiagnosticsPtrInput
	// Specifies the eviction policy for Virtual Machines in this Scale Set. Possible values are `Deallocate` and `Delete`.
	EvictionPolicy pulumi.StringPtrInput
	// Can be specified multiple times to add extension profiles to the scale set. Each `extension` block supports the fields documented below.
	Extensions ScaleSetExtensionArrayInput
	// Specifies the identifier for the load balancer health probe. Required when using `Rolling` as your `upgradePolicyMode`.
	HealthProbeId pulumi.StringPtrInput
	Identity      ScaleSetIdentityPtrInput
	// Specifies the Windows OS license type. If supplied, the only allowed values are `Windows_Client` and `Windows_Server`.
	LicenseType pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the virtual machine scale set resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A collection of network profile block as documented below.
	NetworkProfiles ScaleSetNetworkProfileArrayInput
	// A Virtual Machine OS Profile block as documented below.
	OsProfile ScaleSetOsProfileInput
	// A Linux config block as documented below.
	OsProfileLinuxConfig ScaleSetOsProfileLinuxConfigPtrInput
	// A collection of Secret blocks as documented below.
	OsProfileSecrets ScaleSetOsProfileSecretArrayInput
	// A Windows config block as documented below.
	OsProfileWindowsConfig ScaleSetOsProfileWindowsConfigPtrInput
	// Specifies whether the virtual machine scale set should be overprovisioned. Defaults to `true`.
	Overprovision pulumi.BoolPtrInput
	// A plan block as documented below.
	Plan ScaleSetPlanPtrInput
	// Specifies the priority for the Virtual Machines in the Scale Set. Defaults to `Regular`. Possible values are `Low` and `Regular`.
	Priority pulumi.StringPtrInput
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
	ProximityPlacementGroupId pulumi.StringPtrInput
	// The name of the resource group in which to create the virtual machine scale set. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `rollingUpgradePolicy` block as defined below. This is only applicable when the `upgradePolicyMode` is `Rolling`.
	RollingUpgradePolicy ScaleSetRollingUpgradePolicyPtrInput
	// Specifies whether the scale set is limited to a single placement group with a maximum size of 100 virtual machines. If set to false, managed disks must be used. Default is true. Changing this forces a new resource to be created. See [documentation](http://docs.microsoft.com/en-us/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-placement-groups) for more information.
	SinglePlacementGroup pulumi.BoolPtrInput
	// A sku block as documented below.
	Sku ScaleSetSkuInput
	// A storage profile data disk block as documented below
	StorageProfileDataDisks ScaleSetStorageProfileDataDiskArrayInput
	// A storage profile image reference block as documented below.
	StorageProfileImageReference ScaleSetStorageProfileImageReferencePtrInput
	// A storage profile os disk block as documented below
	StorageProfileOsDisk ScaleSetStorageProfileOsDiskInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the mode of an upgrade to virtual machines in the scale set. Possible values, `Rolling`, `Manual`, or `Automatic`. When choosing `Rolling`, you will need to set a health probe.
	UpgradePolicyMode pulumi.StringInput
	// A collection of availability zones to spread the Virtual Machines over.
	Zones pulumi.StringArrayInput
}

func (ScaleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scaleSetArgs)(nil)).Elem()
}

type ScaleSetInput interface {
	pulumi.Input

	ToScaleSetOutput() ScaleSetOutput
	ToScaleSetOutputWithContext(ctx context.Context) ScaleSetOutput
}

func (ScaleSet) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleSet)(nil)).Elem()
}

func (i ScaleSet) ToScaleSetOutput() ScaleSetOutput {
	return i.ToScaleSetOutputWithContext(context.Background())
}

func (i ScaleSet) ToScaleSetOutputWithContext(ctx context.Context) ScaleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleSetOutput)
}

type ScaleSetOutput struct {
	*pulumi.OutputState
}

func (ScaleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleSetOutput)(nil)).Elem()
}

func (o ScaleSetOutput) ToScaleSetOutput() ScaleSetOutput {
	return o
}

func (o ScaleSetOutput) ToScaleSetOutputWithContext(ctx context.Context) ScaleSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ScaleSetOutput{})
}

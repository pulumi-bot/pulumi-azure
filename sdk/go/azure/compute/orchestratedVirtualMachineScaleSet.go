// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Orchestrated Virtual Machine Scale Set.
//
// > **Note:** Orchestrated Virtual Machine Scale Sets are in Public Preview and it may receive breaking changes - [more details can be found in the Azure Documentation](https://docs.microsoft.com/en-us/azure/virtual-machine-scale-sets/orchestration-modes).
//
// > **Note:** Azure is planning to deprecate the `singlePlacementGroup` attribute in the Orchestrated Virtual Machine Scale Set starting from api-version `2019-12-01` and there will be a breaking change in the Orchestrated Virtual Machine Scale Set.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewOrchestratedVirtualMachineScaleSet(ctx, "exampleOrchestratedVirtualMachineScaleSet", &compute.OrchestratedVirtualMachineScaleSetArgs{
// 			Location:                 exampleResourceGroup.Location,
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			PlatformFaultDomainCount: pulumi.Int(1),
// 			Zones: pulumi.String(pulumi.String{
// 				pulumi.String("1"),
// 			}),
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
// An Orchestrated Virtual Machine Scale Set can be imported using the `resource id`, e.g.
type OrchestratedVirtualMachineScaleSet struct {
	pulumi.CustomResourceState

	// The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	PlatformFaultDomainCount pulumi.IntOutput `pulumi:"platformFaultDomainCount"`
	// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId pulumi.StringPtrOutput `pulumi:"proximityPlacementGroupId"`
	// The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Should the Orchestrated Virtual Machine Scale Set use single placement group? Defaults to `false`.
	SinglePlacementGroup pulumi.BoolPtrOutput `pulumi:"singlePlacementGroup"`
	// A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Unique ID for the Orchestrated Virtual Machine Scale Set.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
	// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
	Zones pulumi.StringPtrOutput `pulumi:"zones"`
}

// NewOrchestratedVirtualMachineScaleSet registers a new resource with the given unique name, arguments, and options.
func NewOrchestratedVirtualMachineScaleSet(ctx *pulumi.Context,
	name string, args *OrchestratedVirtualMachineScaleSetArgs, opts ...pulumi.ResourceOption) (*OrchestratedVirtualMachineScaleSet, error) {
	if args == nil || args.PlatformFaultDomainCount == nil {
		return nil, errors.New("missing required argument 'PlatformFaultDomainCount'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &OrchestratedVirtualMachineScaleSetArgs{}
	}
	var resource OrchestratedVirtualMachineScaleSet
	err := ctx.RegisterResource("azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrchestratedVirtualMachineScaleSet gets an existing OrchestratedVirtualMachineScaleSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrchestratedVirtualMachineScaleSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrchestratedVirtualMachineScaleSetState, opts ...pulumi.ResourceOption) (*OrchestratedVirtualMachineScaleSet, error) {
	var resource OrchestratedVirtualMachineScaleSet
	err := ctx.ReadResource("azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrchestratedVirtualMachineScaleSet resources.
type orchestratedVirtualMachineScaleSetState struct {
	// The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	PlatformFaultDomainCount *int `pulumi:"platformFaultDomainCount"`
	// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId *string `pulumi:"proximityPlacementGroupId"`
	// The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Should the Orchestrated Virtual Machine Scale Set use single placement group? Defaults to `false`.
	SinglePlacementGroup *bool `pulumi:"singlePlacementGroup"`
	// A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
	Tags map[string]string `pulumi:"tags"`
	// The Unique ID for the Orchestrated Virtual Machine Scale Set.
	UniqueId *string `pulumi:"uniqueId"`
	// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
	Zones *string `pulumi:"zones"`
}

type OrchestratedVirtualMachineScaleSetState struct {
	// The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	PlatformFaultDomainCount pulumi.IntPtrInput
	// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId pulumi.StringPtrInput
	// The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Should the Orchestrated Virtual Machine Scale Set use single placement group? Defaults to `false`.
	SinglePlacementGroup pulumi.BoolPtrInput
	// A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
	Tags pulumi.StringMapInput
	// The Unique ID for the Orchestrated Virtual Machine Scale Set.
	UniqueId pulumi.StringPtrInput
	// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
	Zones pulumi.StringPtrInput
}

func (OrchestratedVirtualMachineScaleSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*orchestratedVirtualMachineScaleSetState)(nil)).Elem()
}

type orchestratedVirtualMachineScaleSetArgs struct {
	// The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	PlatformFaultDomainCount int `pulumi:"platformFaultDomainCount"`
	// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId *string `pulumi:"proximityPlacementGroupId"`
	// The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Should the Orchestrated Virtual Machine Scale Set use single placement group? Defaults to `false`.
	SinglePlacementGroup *bool `pulumi:"singlePlacementGroup"`
	// A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
	Tags map[string]string `pulumi:"tags"`
	// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
	Zones *string `pulumi:"zones"`
}

// The set of arguments for constructing a OrchestratedVirtualMachineScaleSet resource.
type OrchestratedVirtualMachineScaleSetArgs struct {
	// The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	PlatformFaultDomainCount pulumi.IntInput
	// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId pulumi.StringPtrInput
	// The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Should the Orchestrated Virtual Machine Scale Set use single placement group? Defaults to `false`.
	SinglePlacementGroup pulumi.BoolPtrInput
	// A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
	Tags pulumi.StringMapInput
	// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
	Zones pulumi.StringPtrInput
}

func (OrchestratedVirtualMachineScaleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orchestratedVirtualMachineScaleSetArgs)(nil)).Elem()
}

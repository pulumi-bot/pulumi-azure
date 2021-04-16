// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
//
// ```sh
//  $ pulumi import azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/Microsoft.Compute/virtualMachineScaleSets/scaleset1
// ```
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PlatformFaultDomainCount == nil {
		return nil, errors.New("invalid value for required argument 'PlatformFaultDomainCount'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
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

type OrchestratedVirtualMachineScaleSetInput interface {
	pulumi.Input

	ToOrchestratedVirtualMachineScaleSetOutput() OrchestratedVirtualMachineScaleSetOutput
	ToOrchestratedVirtualMachineScaleSetOutputWithContext(ctx context.Context) OrchestratedVirtualMachineScaleSetOutput
}

func (*OrchestratedVirtualMachineScaleSet) ElementType() reflect.Type {
	return reflect.TypeOf((*OrchestratedVirtualMachineScaleSet)(nil))
}

func (i *OrchestratedVirtualMachineScaleSet) ToOrchestratedVirtualMachineScaleSetOutput() OrchestratedVirtualMachineScaleSetOutput {
	return i.ToOrchestratedVirtualMachineScaleSetOutputWithContext(context.Background())
}

func (i *OrchestratedVirtualMachineScaleSet) ToOrchestratedVirtualMachineScaleSetOutputWithContext(ctx context.Context) OrchestratedVirtualMachineScaleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrchestratedVirtualMachineScaleSetOutput)
}

func (i *OrchestratedVirtualMachineScaleSet) ToOrchestratedVirtualMachineScaleSetPtrOutput() OrchestratedVirtualMachineScaleSetPtrOutput {
	return i.ToOrchestratedVirtualMachineScaleSetPtrOutputWithContext(context.Background())
}

func (i *OrchestratedVirtualMachineScaleSet) ToOrchestratedVirtualMachineScaleSetPtrOutputWithContext(ctx context.Context) OrchestratedVirtualMachineScaleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrchestratedVirtualMachineScaleSetPtrOutput)
}

type OrchestratedVirtualMachineScaleSetPtrInput interface {
	pulumi.Input

	ToOrchestratedVirtualMachineScaleSetPtrOutput() OrchestratedVirtualMachineScaleSetPtrOutput
	ToOrchestratedVirtualMachineScaleSetPtrOutputWithContext(ctx context.Context) OrchestratedVirtualMachineScaleSetPtrOutput
}

type orchestratedVirtualMachineScaleSetPtrType OrchestratedVirtualMachineScaleSetArgs

func (*orchestratedVirtualMachineScaleSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrchestratedVirtualMachineScaleSet)(nil))
}

func (i *orchestratedVirtualMachineScaleSetPtrType) ToOrchestratedVirtualMachineScaleSetPtrOutput() OrchestratedVirtualMachineScaleSetPtrOutput {
	return i.ToOrchestratedVirtualMachineScaleSetPtrOutputWithContext(context.Background())
}

func (i *orchestratedVirtualMachineScaleSetPtrType) ToOrchestratedVirtualMachineScaleSetPtrOutputWithContext(ctx context.Context) OrchestratedVirtualMachineScaleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrchestratedVirtualMachineScaleSetPtrOutput)
}

// OrchestratedVirtualMachineScaleSetArrayInput is an input type that accepts OrchestratedVirtualMachineScaleSetArray and OrchestratedVirtualMachineScaleSetArrayOutput values.
// You can construct a concrete instance of `OrchestratedVirtualMachineScaleSetArrayInput` via:
//
//          OrchestratedVirtualMachineScaleSetArray{ OrchestratedVirtualMachineScaleSetArgs{...} }
type OrchestratedVirtualMachineScaleSetArrayInput interface {
	pulumi.Input

	ToOrchestratedVirtualMachineScaleSetArrayOutput() OrchestratedVirtualMachineScaleSetArrayOutput
	ToOrchestratedVirtualMachineScaleSetArrayOutputWithContext(context.Context) OrchestratedVirtualMachineScaleSetArrayOutput
}

type OrchestratedVirtualMachineScaleSetArray []OrchestratedVirtualMachineScaleSetInput

func (OrchestratedVirtualMachineScaleSetArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*OrchestratedVirtualMachineScaleSet)(nil))
}

func (i OrchestratedVirtualMachineScaleSetArray) ToOrchestratedVirtualMachineScaleSetArrayOutput() OrchestratedVirtualMachineScaleSetArrayOutput {
	return i.ToOrchestratedVirtualMachineScaleSetArrayOutputWithContext(context.Background())
}

func (i OrchestratedVirtualMachineScaleSetArray) ToOrchestratedVirtualMachineScaleSetArrayOutputWithContext(ctx context.Context) OrchestratedVirtualMachineScaleSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrchestratedVirtualMachineScaleSetArrayOutput)
}

// OrchestratedVirtualMachineScaleSetMapInput is an input type that accepts OrchestratedVirtualMachineScaleSetMap and OrchestratedVirtualMachineScaleSetMapOutput values.
// You can construct a concrete instance of `OrchestratedVirtualMachineScaleSetMapInput` via:
//
//          OrchestratedVirtualMachineScaleSetMap{ "key": OrchestratedVirtualMachineScaleSetArgs{...} }
type OrchestratedVirtualMachineScaleSetMapInput interface {
	pulumi.Input

	ToOrchestratedVirtualMachineScaleSetMapOutput() OrchestratedVirtualMachineScaleSetMapOutput
	ToOrchestratedVirtualMachineScaleSetMapOutputWithContext(context.Context) OrchestratedVirtualMachineScaleSetMapOutput
}

type OrchestratedVirtualMachineScaleSetMap map[string]OrchestratedVirtualMachineScaleSetInput

func (OrchestratedVirtualMachineScaleSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*OrchestratedVirtualMachineScaleSet)(nil))
}

func (i OrchestratedVirtualMachineScaleSetMap) ToOrchestratedVirtualMachineScaleSetMapOutput() OrchestratedVirtualMachineScaleSetMapOutput {
	return i.ToOrchestratedVirtualMachineScaleSetMapOutputWithContext(context.Background())
}

func (i OrchestratedVirtualMachineScaleSetMap) ToOrchestratedVirtualMachineScaleSetMapOutputWithContext(ctx context.Context) OrchestratedVirtualMachineScaleSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrchestratedVirtualMachineScaleSetMapOutput)
}

type OrchestratedVirtualMachineScaleSetOutput struct {
	*pulumi.OutputState
}

func (OrchestratedVirtualMachineScaleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrchestratedVirtualMachineScaleSet)(nil))
}

func (o OrchestratedVirtualMachineScaleSetOutput) ToOrchestratedVirtualMachineScaleSetOutput() OrchestratedVirtualMachineScaleSetOutput {
	return o
}

func (o OrchestratedVirtualMachineScaleSetOutput) ToOrchestratedVirtualMachineScaleSetOutputWithContext(ctx context.Context) OrchestratedVirtualMachineScaleSetOutput {
	return o
}

func (o OrchestratedVirtualMachineScaleSetOutput) ToOrchestratedVirtualMachineScaleSetPtrOutput() OrchestratedVirtualMachineScaleSetPtrOutput {
	return o.ToOrchestratedVirtualMachineScaleSetPtrOutputWithContext(context.Background())
}

func (o OrchestratedVirtualMachineScaleSetOutput) ToOrchestratedVirtualMachineScaleSetPtrOutputWithContext(ctx context.Context) OrchestratedVirtualMachineScaleSetPtrOutput {
	return o.ApplyT(func(v OrchestratedVirtualMachineScaleSet) *OrchestratedVirtualMachineScaleSet {
		return &v
	}).(OrchestratedVirtualMachineScaleSetPtrOutput)
}

type OrchestratedVirtualMachineScaleSetPtrOutput struct {
	*pulumi.OutputState
}

func (OrchestratedVirtualMachineScaleSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrchestratedVirtualMachineScaleSet)(nil))
}

func (o OrchestratedVirtualMachineScaleSetPtrOutput) ToOrchestratedVirtualMachineScaleSetPtrOutput() OrchestratedVirtualMachineScaleSetPtrOutput {
	return o
}

func (o OrchestratedVirtualMachineScaleSetPtrOutput) ToOrchestratedVirtualMachineScaleSetPtrOutputWithContext(ctx context.Context) OrchestratedVirtualMachineScaleSetPtrOutput {
	return o
}

type OrchestratedVirtualMachineScaleSetArrayOutput struct{ *pulumi.OutputState }

func (OrchestratedVirtualMachineScaleSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OrchestratedVirtualMachineScaleSet)(nil))
}

func (o OrchestratedVirtualMachineScaleSetArrayOutput) ToOrchestratedVirtualMachineScaleSetArrayOutput() OrchestratedVirtualMachineScaleSetArrayOutput {
	return o
}

func (o OrchestratedVirtualMachineScaleSetArrayOutput) ToOrchestratedVirtualMachineScaleSetArrayOutputWithContext(ctx context.Context) OrchestratedVirtualMachineScaleSetArrayOutput {
	return o
}

func (o OrchestratedVirtualMachineScaleSetArrayOutput) Index(i pulumi.IntInput) OrchestratedVirtualMachineScaleSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OrchestratedVirtualMachineScaleSet {
		return vs[0].([]OrchestratedVirtualMachineScaleSet)[vs[1].(int)]
	}).(OrchestratedVirtualMachineScaleSetOutput)
}

type OrchestratedVirtualMachineScaleSetMapOutput struct{ *pulumi.OutputState }

func (OrchestratedVirtualMachineScaleSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OrchestratedVirtualMachineScaleSet)(nil))
}

func (o OrchestratedVirtualMachineScaleSetMapOutput) ToOrchestratedVirtualMachineScaleSetMapOutput() OrchestratedVirtualMachineScaleSetMapOutput {
	return o
}

func (o OrchestratedVirtualMachineScaleSetMapOutput) ToOrchestratedVirtualMachineScaleSetMapOutputWithContext(ctx context.Context) OrchestratedVirtualMachineScaleSetMapOutput {
	return o
}

func (o OrchestratedVirtualMachineScaleSetMapOutput) MapIndex(k pulumi.StringInput) OrchestratedVirtualMachineScaleSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) OrchestratedVirtualMachineScaleSet {
		return vs[0].(map[string]OrchestratedVirtualMachineScaleSet)[vs[1].(string)]
	}).(OrchestratedVirtualMachineScaleSetOutput)
}

func init() {
	pulumi.RegisterOutputType(OrchestratedVirtualMachineScaleSetOutput{})
	pulumi.RegisterOutputType(OrchestratedVirtualMachineScaleSetPtrOutput{})
	pulumi.RegisterOutputType(OrchestratedVirtualMachineScaleSetArrayOutput{})
	pulumi.RegisterOutputType(OrchestratedVirtualMachineScaleSetMapOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manage a Dedicated Host Group.
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
// 		_, err = compute.NewDedicatedHostGroup(ctx, "exampleDedicatedHostGroup", &compute.DedicatedHostGroupArgs{
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			Location:                 exampleResourceGroup.Location,
// 			PlatformFaultDomainCount: pulumi.Int(1),
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
// Dedicated Host Group can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:compute/dedicatedHostGroup:DedicatedHostGroup example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/hostGroups/group1
// ```
type DedicatedHostGroup struct {
	pulumi.CustomResourceState

	// The Azure location where the Dedicated Host Group exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Dedicated Host Group. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of fault domains that the Dedicated Host Group spans. Changing this forces a new resource to be created.
	PlatformFaultDomainCount pulumi.IntOutput `pulumi:"platformFaultDomainCount"`
	// Specifies the name of the resource group the Dedicated Host Group is located in. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A list of Availability Zones in which the Dedicated Host Group should be located. Changing this forces a new resource to be created.
	Zones pulumi.StringPtrOutput `pulumi:"zones"`
}

// NewDedicatedHostGroup registers a new resource with the given unique name, arguments, and options.
func NewDedicatedHostGroup(ctx *pulumi.Context,
	name string, args *DedicatedHostGroupArgs, opts ...pulumi.ResourceOption) (*DedicatedHostGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PlatformFaultDomainCount == nil {
		return nil, errors.New("invalid value for required argument 'PlatformFaultDomainCount'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource DedicatedHostGroup
	err := ctx.RegisterResource("azure:compute/dedicatedHostGroup:DedicatedHostGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedHostGroup gets an existing DedicatedHostGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedHostGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedHostGroupState, opts ...pulumi.ResourceOption) (*DedicatedHostGroup, error) {
	var resource DedicatedHostGroup
	err := ctx.ReadResource("azure:compute/dedicatedHostGroup:DedicatedHostGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedHostGroup resources.
type dedicatedHostGroupState struct {
	// The Azure location where the Dedicated Host Group exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Dedicated Host Group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The number of fault domains that the Dedicated Host Group spans. Changing this forces a new resource to be created.
	PlatformFaultDomainCount *int `pulumi:"platformFaultDomainCount"`
	// Specifies the name of the resource group the Dedicated Host Group is located in. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A list of Availability Zones in which the Dedicated Host Group should be located. Changing this forces a new resource to be created.
	Zones *string `pulumi:"zones"`
}

type DedicatedHostGroupState struct {
	// The Azure location where the Dedicated Host Group exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Dedicated Host Group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The number of fault domains that the Dedicated Host Group spans. Changing this forces a new resource to be created.
	PlatformFaultDomainCount pulumi.IntPtrInput
	// Specifies the name of the resource group the Dedicated Host Group is located in. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A list of Availability Zones in which the Dedicated Host Group should be located. Changing this forces a new resource to be created.
	Zones pulumi.StringPtrInput
}

func (DedicatedHostGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostGroupState)(nil)).Elem()
}

type dedicatedHostGroupArgs struct {
	// The Azure location where the Dedicated Host Group exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Dedicated Host Group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The number of fault domains that the Dedicated Host Group spans. Changing this forces a new resource to be created.
	PlatformFaultDomainCount int `pulumi:"platformFaultDomainCount"`
	// Specifies the name of the resource group the Dedicated Host Group is located in. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A list of Availability Zones in which the Dedicated Host Group should be located. Changing this forces a new resource to be created.
	Zones *string `pulumi:"zones"`
}

// The set of arguments for constructing a DedicatedHostGroup resource.
type DedicatedHostGroupArgs struct {
	// The Azure location where the Dedicated Host Group exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Dedicated Host Group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The number of fault domains that the Dedicated Host Group spans. Changing this forces a new resource to be created.
	PlatformFaultDomainCount pulumi.IntInput
	// Specifies the name of the resource group the Dedicated Host Group is located in. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A list of Availability Zones in which the Dedicated Host Group should be located. Changing this forces a new resource to be created.
	Zones pulumi.StringPtrInput
}

func (DedicatedHostGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostGroupArgs)(nil)).Elem()
}

type DedicatedHostGroupInput interface {
	pulumi.Input

	ToDedicatedHostGroupOutput() DedicatedHostGroupOutput
	ToDedicatedHostGroupOutputWithContext(ctx context.Context) DedicatedHostGroupOutput
}

type DedicatedHostGroupPtrInput interface {
	pulumi.Input

	ToDedicatedHostGroupPtrOutput() DedicatedHostGroupPtrOutput
	ToDedicatedHostGroupPtrOutputWithContext(ctx context.Context) DedicatedHostGroupPtrOutput
}

func (DedicatedHostGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedHostGroup)(nil)).Elem()
}

func (i DedicatedHostGroup) ToDedicatedHostGroupOutput() DedicatedHostGroupOutput {
	return i.ToDedicatedHostGroupOutputWithContext(context.Background())
}

func (i DedicatedHostGroup) ToDedicatedHostGroupOutputWithContext(ctx context.Context) DedicatedHostGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostGroupOutput)
}

func (i DedicatedHostGroup) ToDedicatedHostGroupPtrOutput() DedicatedHostGroupPtrOutput {
	return i.ToDedicatedHostGroupPtrOutputWithContext(context.Background())
}

func (i DedicatedHostGroup) ToDedicatedHostGroupPtrOutputWithContext(ctx context.Context) DedicatedHostGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostGroupPtrOutput)
}

type DedicatedHostGroupOutput struct {
	*pulumi.OutputState
}

func (DedicatedHostGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedHostGroupOutput)(nil)).Elem()
}

func (o DedicatedHostGroupOutput) ToDedicatedHostGroupOutput() DedicatedHostGroupOutput {
	return o
}

func (o DedicatedHostGroupOutput) ToDedicatedHostGroupOutputWithContext(ctx context.Context) DedicatedHostGroupOutput {
	return o
}

type DedicatedHostGroupPtrOutput struct {
	*pulumi.OutputState
}

func (DedicatedHostGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHostGroup)(nil)).Elem()
}

func (o DedicatedHostGroupPtrOutput) ToDedicatedHostGroupPtrOutput() DedicatedHostGroupPtrOutput {
	return o
}

func (o DedicatedHostGroupPtrOutput) ToDedicatedHostGroupPtrOutputWithContext(ctx context.Context) DedicatedHostGroupPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DedicatedHostGroupOutput{})
	pulumi.RegisterOutputType(DedicatedHostGroupPtrOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mixedreality

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure Spatial Anchors Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/mixedreality"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("northeurope"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mixedreality.NewSpatialAnchorsAccount(ctx, "exampleSpatialAnchorsAccount", &mixedreality.SpatialAnchorsAccountArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
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
// Spatial Anchors Account can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:mixedreality/spatialAnchorsAccount:SpatialAnchorsAccount example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.MixedReality/spatialAnchorsAccounts/example
// ```
type SpatialAnchorsAccount struct {
	pulumi.CustomResourceState

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Spatial Anchors Account. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Spatial Anchors Account.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewSpatialAnchorsAccount registers a new resource with the given unique name, arguments, and options.
func NewSpatialAnchorsAccount(ctx *pulumi.Context,
	name string, args *SpatialAnchorsAccountArgs, opts ...pulumi.ResourceOption) (*SpatialAnchorsAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource SpatialAnchorsAccount
	err := ctx.RegisterResource("azure:mixedreality/spatialAnchorsAccount:SpatialAnchorsAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpatialAnchorsAccount gets an existing SpatialAnchorsAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpatialAnchorsAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpatialAnchorsAccountState, opts ...pulumi.ResourceOption) (*SpatialAnchorsAccount, error) {
	var resource SpatialAnchorsAccount
	err := ctx.ReadResource("azure:mixedreality/spatialAnchorsAccount:SpatialAnchorsAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpatialAnchorsAccount resources.
type spatialAnchorsAccountState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Spatial Anchors Account. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Spatial Anchors Account.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type SpatialAnchorsAccountState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Spatial Anchors Account. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Spatial Anchors Account.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SpatialAnchorsAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*spatialAnchorsAccountState)(nil)).Elem()
}

type spatialAnchorsAccountArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Spatial Anchors Account. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Spatial Anchors Account.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SpatialAnchorsAccount resource.
type SpatialAnchorsAccountArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Spatial Anchors Account. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Spatial Anchors Account.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SpatialAnchorsAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spatialAnchorsAccountArgs)(nil)).Elem()
}

type SpatialAnchorsAccountInput interface {
	pulumi.Input

	ToSpatialAnchorsAccountOutput() SpatialAnchorsAccountOutput
	ToSpatialAnchorsAccountOutputWithContext(ctx context.Context) SpatialAnchorsAccountOutput
}

func (*SpatialAnchorsAccount) ElementType() reflect.Type {
	return reflect.TypeOf((*SpatialAnchorsAccount)(nil))
}

func (i *SpatialAnchorsAccount) ToSpatialAnchorsAccountOutput() SpatialAnchorsAccountOutput {
	return i.ToSpatialAnchorsAccountOutputWithContext(context.Background())
}

func (i *SpatialAnchorsAccount) ToSpatialAnchorsAccountOutputWithContext(ctx context.Context) SpatialAnchorsAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpatialAnchorsAccountOutput)
}

func (i *SpatialAnchorsAccount) ToSpatialAnchorsAccountPtrOutput() SpatialAnchorsAccountPtrOutput {
	return i.ToSpatialAnchorsAccountPtrOutputWithContext(context.Background())
}

func (i *SpatialAnchorsAccount) ToSpatialAnchorsAccountPtrOutputWithContext(ctx context.Context) SpatialAnchorsAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpatialAnchorsAccountPtrOutput)
}

type SpatialAnchorsAccountPtrInput interface {
	pulumi.Input

	ToSpatialAnchorsAccountPtrOutput() SpatialAnchorsAccountPtrOutput
	ToSpatialAnchorsAccountPtrOutputWithContext(ctx context.Context) SpatialAnchorsAccountPtrOutput
}

type spatialAnchorsAccountPtrType SpatialAnchorsAccountArgs

func (*spatialAnchorsAccountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SpatialAnchorsAccount)(nil))
}

func (i *spatialAnchorsAccountPtrType) ToSpatialAnchorsAccountPtrOutput() SpatialAnchorsAccountPtrOutput {
	return i.ToSpatialAnchorsAccountPtrOutputWithContext(context.Background())
}

func (i *spatialAnchorsAccountPtrType) ToSpatialAnchorsAccountPtrOutputWithContext(ctx context.Context) SpatialAnchorsAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpatialAnchorsAccountPtrOutput)
}

// SpatialAnchorsAccountArrayInput is an input type that accepts SpatialAnchorsAccountArray and SpatialAnchorsAccountArrayOutput values.
// You can construct a concrete instance of `SpatialAnchorsAccountArrayInput` via:
//
//          SpatialAnchorsAccountArray{ SpatialAnchorsAccountArgs{...} }
type SpatialAnchorsAccountArrayInput interface {
	pulumi.Input

	ToSpatialAnchorsAccountArrayOutput() SpatialAnchorsAccountArrayOutput
	ToSpatialAnchorsAccountArrayOutputWithContext(context.Context) SpatialAnchorsAccountArrayOutput
}

type SpatialAnchorsAccountArray []SpatialAnchorsAccountInput

func (SpatialAnchorsAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SpatialAnchorsAccount)(nil))
}

func (i SpatialAnchorsAccountArray) ToSpatialAnchorsAccountArrayOutput() SpatialAnchorsAccountArrayOutput {
	return i.ToSpatialAnchorsAccountArrayOutputWithContext(context.Background())
}

func (i SpatialAnchorsAccountArray) ToSpatialAnchorsAccountArrayOutputWithContext(ctx context.Context) SpatialAnchorsAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpatialAnchorsAccountArrayOutput)
}

// SpatialAnchorsAccountMapInput is an input type that accepts SpatialAnchorsAccountMap and SpatialAnchorsAccountMapOutput values.
// You can construct a concrete instance of `SpatialAnchorsAccountMapInput` via:
//
//          SpatialAnchorsAccountMap{ "key": SpatialAnchorsAccountArgs{...} }
type SpatialAnchorsAccountMapInput interface {
	pulumi.Input

	ToSpatialAnchorsAccountMapOutput() SpatialAnchorsAccountMapOutput
	ToSpatialAnchorsAccountMapOutputWithContext(context.Context) SpatialAnchorsAccountMapOutput
}

type SpatialAnchorsAccountMap map[string]SpatialAnchorsAccountInput

func (SpatialAnchorsAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SpatialAnchorsAccount)(nil))
}

func (i SpatialAnchorsAccountMap) ToSpatialAnchorsAccountMapOutput() SpatialAnchorsAccountMapOutput {
	return i.ToSpatialAnchorsAccountMapOutputWithContext(context.Background())
}

func (i SpatialAnchorsAccountMap) ToSpatialAnchorsAccountMapOutputWithContext(ctx context.Context) SpatialAnchorsAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpatialAnchorsAccountMapOutput)
}

type SpatialAnchorsAccountOutput struct {
	*pulumi.OutputState
}

func (SpatialAnchorsAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SpatialAnchorsAccount)(nil))
}

func (o SpatialAnchorsAccountOutput) ToSpatialAnchorsAccountOutput() SpatialAnchorsAccountOutput {
	return o
}

func (o SpatialAnchorsAccountOutput) ToSpatialAnchorsAccountOutputWithContext(ctx context.Context) SpatialAnchorsAccountOutput {
	return o
}

func (o SpatialAnchorsAccountOutput) ToSpatialAnchorsAccountPtrOutput() SpatialAnchorsAccountPtrOutput {
	return o.ToSpatialAnchorsAccountPtrOutputWithContext(context.Background())
}

func (o SpatialAnchorsAccountOutput) ToSpatialAnchorsAccountPtrOutputWithContext(ctx context.Context) SpatialAnchorsAccountPtrOutput {
	return o.ApplyT(func(v SpatialAnchorsAccount) *SpatialAnchorsAccount {
		return &v
	}).(SpatialAnchorsAccountPtrOutput)
}

type SpatialAnchorsAccountPtrOutput struct {
	*pulumi.OutputState
}

func (SpatialAnchorsAccountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpatialAnchorsAccount)(nil))
}

func (o SpatialAnchorsAccountPtrOutput) ToSpatialAnchorsAccountPtrOutput() SpatialAnchorsAccountPtrOutput {
	return o
}

func (o SpatialAnchorsAccountPtrOutput) ToSpatialAnchorsAccountPtrOutputWithContext(ctx context.Context) SpatialAnchorsAccountPtrOutput {
	return o
}

type SpatialAnchorsAccountArrayOutput struct{ *pulumi.OutputState }

func (SpatialAnchorsAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SpatialAnchorsAccount)(nil))
}

func (o SpatialAnchorsAccountArrayOutput) ToSpatialAnchorsAccountArrayOutput() SpatialAnchorsAccountArrayOutput {
	return o
}

func (o SpatialAnchorsAccountArrayOutput) ToSpatialAnchorsAccountArrayOutputWithContext(ctx context.Context) SpatialAnchorsAccountArrayOutput {
	return o
}

func (o SpatialAnchorsAccountArrayOutput) Index(i pulumi.IntInput) SpatialAnchorsAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SpatialAnchorsAccount {
		return vs[0].([]SpatialAnchorsAccount)[vs[1].(int)]
	}).(SpatialAnchorsAccountOutput)
}

type SpatialAnchorsAccountMapOutput struct{ *pulumi.OutputState }

func (SpatialAnchorsAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SpatialAnchorsAccount)(nil))
}

func (o SpatialAnchorsAccountMapOutput) ToSpatialAnchorsAccountMapOutput() SpatialAnchorsAccountMapOutput {
	return o
}

func (o SpatialAnchorsAccountMapOutput) ToSpatialAnchorsAccountMapOutputWithContext(ctx context.Context) SpatialAnchorsAccountMapOutput {
	return o
}

func (o SpatialAnchorsAccountMapOutput) MapIndex(k pulumi.StringInput) SpatialAnchorsAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SpatialAnchorsAccount {
		return vs[0].(map[string]SpatialAnchorsAccount)[vs[1].(string)]
	}).(SpatialAnchorsAccountOutput)
}

func init() {
	pulumi.RegisterOutputType(SpatialAnchorsAccountOutput{})
	pulumi.RegisterOutputType(SpatialAnchorsAccountPtrOutput{})
	pulumi.RegisterOutputType(SpatialAnchorsAccountArrayOutput{})
	pulumi.RegisterOutputType(SpatialAnchorsAccountMapOutput{})
}

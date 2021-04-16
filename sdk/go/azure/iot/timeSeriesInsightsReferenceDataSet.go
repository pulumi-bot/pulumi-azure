// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure IoT Time Series Insights Reference Data Set.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/iot"
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
// 		exampleTimeSeriesInsightsStandardEnvironment, err := iot.NewTimeSeriesInsightsStandardEnvironment(ctx, "exampleTimeSeriesInsightsStandardEnvironment", &iot.TimeSeriesInsightsStandardEnvironmentArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			SkuName:           pulumi.String("S1_1"),
// 			DataRetentionTime: pulumi.String("P30D"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iot.NewTimeSeriesInsightsReferenceDataSet(ctx, "exampleTimeSeriesInsightsReferenceDataSet", &iot.TimeSeriesInsightsReferenceDataSetArgs{
// 			TimeSeriesInsightsEnvironmentId: exampleTimeSeriesInsightsStandardEnvironment.ID(),
// 			Location:                        exampleResourceGroup.Location,
// 			KeyProperties: iot.TimeSeriesInsightsReferenceDataSetKeyPropertyArray{
// 				&iot.TimeSeriesInsightsReferenceDataSetKeyPropertyArgs{
// 					Name: pulumi.String("keyProperty1"),
// 					Type: pulumi.String("String"),
// 				},
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
// Azure IoT Time Series Insights Reference Data Set can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:iot/timeSeriesInsightsReferenceDataSet:TimeSeriesInsightsReferenceDataSet example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/example/referenceDataSets/example
// ```
type TimeSeriesInsightsReferenceDataSet struct {
	pulumi.CustomResourceState

	// The comparison behavior that will be used to compare keys. Valid values include `Ordinal` and `OrdinalIgnoreCase`. Defaults to `Ordinal`. Changing this forces a new resource to be created.
	DataStringComparisonBehavior pulumi.StringPtrOutput `pulumi:"dataStringComparisonBehavior"`
	// A `keyProperty` block as defined below. Changing this forces a new resource to be created.
	KeyProperties TimeSeriesInsightsReferenceDataSetKeyPropertyArrayOutput `pulumi:"keyProperties"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
	TimeSeriesInsightsEnvironmentId pulumi.StringOutput `pulumi:"timeSeriesInsightsEnvironmentId"`
}

// NewTimeSeriesInsightsReferenceDataSet registers a new resource with the given unique name, arguments, and options.
func NewTimeSeriesInsightsReferenceDataSet(ctx *pulumi.Context,
	name string, args *TimeSeriesInsightsReferenceDataSetArgs, opts ...pulumi.ResourceOption) (*TimeSeriesInsightsReferenceDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyProperties == nil {
		return nil, errors.New("invalid value for required argument 'KeyProperties'")
	}
	if args.TimeSeriesInsightsEnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'TimeSeriesInsightsEnvironmentId'")
	}
	var resource TimeSeriesInsightsReferenceDataSet
	err := ctx.RegisterResource("azure:iot/timeSeriesInsightsReferenceDataSet:TimeSeriesInsightsReferenceDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTimeSeriesInsightsReferenceDataSet gets an existing TimeSeriesInsightsReferenceDataSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTimeSeriesInsightsReferenceDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TimeSeriesInsightsReferenceDataSetState, opts ...pulumi.ResourceOption) (*TimeSeriesInsightsReferenceDataSet, error) {
	var resource TimeSeriesInsightsReferenceDataSet
	err := ctx.ReadResource("azure:iot/timeSeriesInsightsReferenceDataSet:TimeSeriesInsightsReferenceDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TimeSeriesInsightsReferenceDataSet resources.
type timeSeriesInsightsReferenceDataSetState struct {
	// The comparison behavior that will be used to compare keys. Valid values include `Ordinal` and `OrdinalIgnoreCase`. Defaults to `Ordinal`. Changing this forces a new resource to be created.
	DataStringComparisonBehavior *string `pulumi:"dataStringComparisonBehavior"`
	// A `keyProperty` block as defined below. Changing this forces a new resource to be created.
	KeyProperties []TimeSeriesInsightsReferenceDataSetKeyProperty `pulumi:"keyProperties"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `pulumi:"name"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
	TimeSeriesInsightsEnvironmentId *string `pulumi:"timeSeriesInsightsEnvironmentId"`
}

type TimeSeriesInsightsReferenceDataSetState struct {
	// The comparison behavior that will be used to compare keys. Valid values include `Ordinal` and `OrdinalIgnoreCase`. Defaults to `Ordinal`. Changing this forces a new resource to be created.
	DataStringComparisonBehavior pulumi.StringPtrInput
	// A `keyProperty` block as defined below. Changing this forces a new resource to be created.
	KeyProperties TimeSeriesInsightsReferenceDataSetKeyPropertyArrayInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
	TimeSeriesInsightsEnvironmentId pulumi.StringPtrInput
}

func (TimeSeriesInsightsReferenceDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*timeSeriesInsightsReferenceDataSetState)(nil)).Elem()
}

type timeSeriesInsightsReferenceDataSetArgs struct {
	// The comparison behavior that will be used to compare keys. Valid values include `Ordinal` and `OrdinalIgnoreCase`. Defaults to `Ordinal`. Changing this forces a new resource to be created.
	DataStringComparisonBehavior *string `pulumi:"dataStringComparisonBehavior"`
	// A `keyProperty` block as defined below. Changing this forces a new resource to be created.
	KeyProperties []TimeSeriesInsightsReferenceDataSetKeyProperty `pulumi:"keyProperties"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `pulumi:"name"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
	TimeSeriesInsightsEnvironmentId string `pulumi:"timeSeriesInsightsEnvironmentId"`
}

// The set of arguments for constructing a TimeSeriesInsightsReferenceDataSet resource.
type TimeSeriesInsightsReferenceDataSetArgs struct {
	// The comparison behavior that will be used to compare keys. Valid values include `Ordinal` and `OrdinalIgnoreCase`. Defaults to `Ordinal`. Changing this forces a new resource to be created.
	DataStringComparisonBehavior pulumi.StringPtrInput
	// A `keyProperty` block as defined below. Changing this forces a new resource to be created.
	KeyProperties TimeSeriesInsightsReferenceDataSetKeyPropertyArrayInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
	TimeSeriesInsightsEnvironmentId pulumi.StringInput
}

func (TimeSeriesInsightsReferenceDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*timeSeriesInsightsReferenceDataSetArgs)(nil)).Elem()
}

type TimeSeriesInsightsReferenceDataSetInput interface {
	pulumi.Input

	ToTimeSeriesInsightsReferenceDataSetOutput() TimeSeriesInsightsReferenceDataSetOutput
	ToTimeSeriesInsightsReferenceDataSetOutputWithContext(ctx context.Context) TimeSeriesInsightsReferenceDataSetOutput
}

func (*TimeSeriesInsightsReferenceDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeSeriesInsightsReferenceDataSet)(nil))
}

func (i *TimeSeriesInsightsReferenceDataSet) ToTimeSeriesInsightsReferenceDataSetOutput() TimeSeriesInsightsReferenceDataSetOutput {
	return i.ToTimeSeriesInsightsReferenceDataSetOutputWithContext(context.Background())
}

func (i *TimeSeriesInsightsReferenceDataSet) ToTimeSeriesInsightsReferenceDataSetOutputWithContext(ctx context.Context) TimeSeriesInsightsReferenceDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsReferenceDataSetOutput)
}

func (i *TimeSeriesInsightsReferenceDataSet) ToTimeSeriesInsightsReferenceDataSetPtrOutput() TimeSeriesInsightsReferenceDataSetPtrOutput {
	return i.ToTimeSeriesInsightsReferenceDataSetPtrOutputWithContext(context.Background())
}

func (i *TimeSeriesInsightsReferenceDataSet) ToTimeSeriesInsightsReferenceDataSetPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsReferenceDataSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsReferenceDataSetPtrOutput)
}

type TimeSeriesInsightsReferenceDataSetPtrInput interface {
	pulumi.Input

	ToTimeSeriesInsightsReferenceDataSetPtrOutput() TimeSeriesInsightsReferenceDataSetPtrOutput
	ToTimeSeriesInsightsReferenceDataSetPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsReferenceDataSetPtrOutput
}

type timeSeriesInsightsReferenceDataSetPtrType TimeSeriesInsightsReferenceDataSetArgs

func (*timeSeriesInsightsReferenceDataSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeSeriesInsightsReferenceDataSet)(nil))
}

func (i *timeSeriesInsightsReferenceDataSetPtrType) ToTimeSeriesInsightsReferenceDataSetPtrOutput() TimeSeriesInsightsReferenceDataSetPtrOutput {
	return i.ToTimeSeriesInsightsReferenceDataSetPtrOutputWithContext(context.Background())
}

func (i *timeSeriesInsightsReferenceDataSetPtrType) ToTimeSeriesInsightsReferenceDataSetPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsReferenceDataSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsReferenceDataSetPtrOutput)
}

// TimeSeriesInsightsReferenceDataSetArrayInput is an input type that accepts TimeSeriesInsightsReferenceDataSetArray and TimeSeriesInsightsReferenceDataSetArrayOutput values.
// You can construct a concrete instance of `TimeSeriesInsightsReferenceDataSetArrayInput` via:
//
//          TimeSeriesInsightsReferenceDataSetArray{ TimeSeriesInsightsReferenceDataSetArgs{...} }
type TimeSeriesInsightsReferenceDataSetArrayInput interface {
	pulumi.Input

	ToTimeSeriesInsightsReferenceDataSetArrayOutput() TimeSeriesInsightsReferenceDataSetArrayOutput
	ToTimeSeriesInsightsReferenceDataSetArrayOutputWithContext(context.Context) TimeSeriesInsightsReferenceDataSetArrayOutput
}

type TimeSeriesInsightsReferenceDataSetArray []TimeSeriesInsightsReferenceDataSetInput

func (TimeSeriesInsightsReferenceDataSetArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*TimeSeriesInsightsReferenceDataSet)(nil))
}

func (i TimeSeriesInsightsReferenceDataSetArray) ToTimeSeriesInsightsReferenceDataSetArrayOutput() TimeSeriesInsightsReferenceDataSetArrayOutput {
	return i.ToTimeSeriesInsightsReferenceDataSetArrayOutputWithContext(context.Background())
}

func (i TimeSeriesInsightsReferenceDataSetArray) ToTimeSeriesInsightsReferenceDataSetArrayOutputWithContext(ctx context.Context) TimeSeriesInsightsReferenceDataSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsReferenceDataSetArrayOutput)
}

// TimeSeriesInsightsReferenceDataSetMapInput is an input type that accepts TimeSeriesInsightsReferenceDataSetMap and TimeSeriesInsightsReferenceDataSetMapOutput values.
// You can construct a concrete instance of `TimeSeriesInsightsReferenceDataSetMapInput` via:
//
//          TimeSeriesInsightsReferenceDataSetMap{ "key": TimeSeriesInsightsReferenceDataSetArgs{...} }
type TimeSeriesInsightsReferenceDataSetMapInput interface {
	pulumi.Input

	ToTimeSeriesInsightsReferenceDataSetMapOutput() TimeSeriesInsightsReferenceDataSetMapOutput
	ToTimeSeriesInsightsReferenceDataSetMapOutputWithContext(context.Context) TimeSeriesInsightsReferenceDataSetMapOutput
}

type TimeSeriesInsightsReferenceDataSetMap map[string]TimeSeriesInsightsReferenceDataSetInput

func (TimeSeriesInsightsReferenceDataSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*TimeSeriesInsightsReferenceDataSet)(nil))
}

func (i TimeSeriesInsightsReferenceDataSetMap) ToTimeSeriesInsightsReferenceDataSetMapOutput() TimeSeriesInsightsReferenceDataSetMapOutput {
	return i.ToTimeSeriesInsightsReferenceDataSetMapOutputWithContext(context.Background())
}

func (i TimeSeriesInsightsReferenceDataSetMap) ToTimeSeriesInsightsReferenceDataSetMapOutputWithContext(ctx context.Context) TimeSeriesInsightsReferenceDataSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsReferenceDataSetMapOutput)
}

type TimeSeriesInsightsReferenceDataSetOutput struct {
	*pulumi.OutputState
}

func (TimeSeriesInsightsReferenceDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeSeriesInsightsReferenceDataSet)(nil))
}

func (o TimeSeriesInsightsReferenceDataSetOutput) ToTimeSeriesInsightsReferenceDataSetOutput() TimeSeriesInsightsReferenceDataSetOutput {
	return o
}

func (o TimeSeriesInsightsReferenceDataSetOutput) ToTimeSeriesInsightsReferenceDataSetOutputWithContext(ctx context.Context) TimeSeriesInsightsReferenceDataSetOutput {
	return o
}

func (o TimeSeriesInsightsReferenceDataSetOutput) ToTimeSeriesInsightsReferenceDataSetPtrOutput() TimeSeriesInsightsReferenceDataSetPtrOutput {
	return o.ToTimeSeriesInsightsReferenceDataSetPtrOutputWithContext(context.Background())
}

func (o TimeSeriesInsightsReferenceDataSetOutput) ToTimeSeriesInsightsReferenceDataSetPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsReferenceDataSetPtrOutput {
	return o.ApplyT(func(v TimeSeriesInsightsReferenceDataSet) *TimeSeriesInsightsReferenceDataSet {
		return &v
	}).(TimeSeriesInsightsReferenceDataSetPtrOutput)
}

type TimeSeriesInsightsReferenceDataSetPtrOutput struct {
	*pulumi.OutputState
}

func (TimeSeriesInsightsReferenceDataSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeSeriesInsightsReferenceDataSet)(nil))
}

func (o TimeSeriesInsightsReferenceDataSetPtrOutput) ToTimeSeriesInsightsReferenceDataSetPtrOutput() TimeSeriesInsightsReferenceDataSetPtrOutput {
	return o
}

func (o TimeSeriesInsightsReferenceDataSetPtrOutput) ToTimeSeriesInsightsReferenceDataSetPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsReferenceDataSetPtrOutput {
	return o
}

type TimeSeriesInsightsReferenceDataSetArrayOutput struct{ *pulumi.OutputState }

func (TimeSeriesInsightsReferenceDataSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeSeriesInsightsReferenceDataSet)(nil))
}

func (o TimeSeriesInsightsReferenceDataSetArrayOutput) ToTimeSeriesInsightsReferenceDataSetArrayOutput() TimeSeriesInsightsReferenceDataSetArrayOutput {
	return o
}

func (o TimeSeriesInsightsReferenceDataSetArrayOutput) ToTimeSeriesInsightsReferenceDataSetArrayOutputWithContext(ctx context.Context) TimeSeriesInsightsReferenceDataSetArrayOutput {
	return o
}

func (o TimeSeriesInsightsReferenceDataSetArrayOutput) Index(i pulumi.IntInput) TimeSeriesInsightsReferenceDataSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TimeSeriesInsightsReferenceDataSet {
		return vs[0].([]TimeSeriesInsightsReferenceDataSet)[vs[1].(int)]
	}).(TimeSeriesInsightsReferenceDataSetOutput)
}

type TimeSeriesInsightsReferenceDataSetMapOutput struct{ *pulumi.OutputState }

func (TimeSeriesInsightsReferenceDataSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TimeSeriesInsightsReferenceDataSet)(nil))
}

func (o TimeSeriesInsightsReferenceDataSetMapOutput) ToTimeSeriesInsightsReferenceDataSetMapOutput() TimeSeriesInsightsReferenceDataSetMapOutput {
	return o
}

func (o TimeSeriesInsightsReferenceDataSetMapOutput) ToTimeSeriesInsightsReferenceDataSetMapOutputWithContext(ctx context.Context) TimeSeriesInsightsReferenceDataSetMapOutput {
	return o
}

func (o TimeSeriesInsightsReferenceDataSetMapOutput) MapIndex(k pulumi.StringInput) TimeSeriesInsightsReferenceDataSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TimeSeriesInsightsReferenceDataSet {
		return vs[0].(map[string]TimeSeriesInsightsReferenceDataSet)[vs[1].(string)]
	}).(TimeSeriesInsightsReferenceDataSetOutput)
}

func init() {
	pulumi.RegisterOutputType(TimeSeriesInsightsReferenceDataSetOutput{})
	pulumi.RegisterOutputType(TimeSeriesInsightsReferenceDataSetPtrOutput{})
	pulumi.RegisterOutputType(TimeSeriesInsightsReferenceDataSetArrayOutput{})
	pulumi.RegisterOutputType(TimeSeriesInsightsReferenceDataSetMapOutput{})
}

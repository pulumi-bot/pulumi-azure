// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure IoT Time Series Insights Reference Data Set.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/iot"
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
type TimeSeriesInsightsReferenceDataSet struct {
	pulumi.CustomResourceState

	// The comparison behavior that will be used to compare keys. Valid values include `Ordinal` and `OrdinalIgnoreCase`. Defaults to `Ordinal`.
	DataStringComparisonBehavior pulumi.StringPtrOutput `pulumi:"dataStringComparisonBehavior"`
	// A `keyProperty` block as defined below.
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
	// The comparison behavior that will be used to compare keys. Valid values include `Ordinal` and `OrdinalIgnoreCase`. Defaults to `Ordinal`.
	DataStringComparisonBehavior *string `pulumi:"dataStringComparisonBehavior"`
	// A `keyProperty` block as defined below.
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
	// The comparison behavior that will be used to compare keys. Valid values include `Ordinal` and `OrdinalIgnoreCase`. Defaults to `Ordinal`.
	DataStringComparisonBehavior pulumi.StringPtrInput
	// A `keyProperty` block as defined below.
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
	// The comparison behavior that will be used to compare keys. Valid values include `Ordinal` and `OrdinalIgnoreCase`. Defaults to `Ordinal`.
	DataStringComparisonBehavior *string `pulumi:"dataStringComparisonBehavior"`
	// A `keyProperty` block as defined below.
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
	// The comparison behavior that will be used to compare keys. Valid values include `Ordinal` and `OrdinalIgnoreCase`. Defaults to `Ordinal`.
	DataStringComparisonBehavior pulumi.StringPtrInput
	// A `keyProperty` block as defined below.
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

func (TimeSeriesInsightsReferenceDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeSeriesInsightsReferenceDataSet)(nil)).Elem()
}

func (i TimeSeriesInsightsReferenceDataSet) ToTimeSeriesInsightsReferenceDataSetOutput() TimeSeriesInsightsReferenceDataSetOutput {
	return i.ToTimeSeriesInsightsReferenceDataSetOutputWithContext(context.Background())
}

func (i TimeSeriesInsightsReferenceDataSet) ToTimeSeriesInsightsReferenceDataSetOutputWithContext(ctx context.Context) TimeSeriesInsightsReferenceDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsReferenceDataSetOutput)
}

type TimeSeriesInsightsReferenceDataSetOutput struct {
	*pulumi.OutputState
}

func (TimeSeriesInsightsReferenceDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeSeriesInsightsReferenceDataSetOutput)(nil)).Elem()
}

func (o TimeSeriesInsightsReferenceDataSetOutput) ToTimeSeriesInsightsReferenceDataSetOutput() TimeSeriesInsightsReferenceDataSetOutput {
	return o
}

func (o TimeSeriesInsightsReferenceDataSetOutput) ToTimeSeriesInsightsReferenceDataSetOutputWithContext(ctx context.Context) TimeSeriesInsightsReferenceDataSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TimeSeriesInsightsReferenceDataSetOutput{})
}

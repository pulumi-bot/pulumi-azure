// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure IoT Time Series Insights Standard Environment.
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
// 		_, err = iot.NewTimeSeriesInsightsStandardEnvironment(ctx, "exampleTimeSeriesInsightsStandardEnvironment", &iot.TimeSeriesInsightsStandardEnvironmentArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			SkuName:           pulumi.String("S1_1"),
// 			DataRetentionTime: pulumi.String("P30D"),
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
// Azure IoT Time Series Insights Standard Environment can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:iot/timeSeriesInsightsStandardEnvironment:TimeSeriesInsightsStandardEnvironment example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/example
// ```
type TimeSeriesInsightsStandardEnvironment struct {
	pulumi.CustomResourceState

	// Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
	DataRetentionTime pulumi.StringOutput `pulumi:"dataRetentionTime"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Azure IoT Time Series Insights Standard Environment. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the event property which will be used to partition data. Changing this forces a new resource to be created.
	PartitionKey pulumi.StringPtrOutput `pulumi:"partitionKey"`
	// The name of the resource group in which to create the Azure IoT Time Series Insights Standard Environment.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the SKU Name for this IoT Time Series Insights Standard Environment. It is string consisting of two parts separated by an underscore(\_).The fist part is the `name`, valid values include: `S1` and `S2`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `S1_1`). Changing this forces a new resource to be created.
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// Specifies the behaviour the IoT Time Series Insights service should take when the environment's capacity has been exceeded. Valid values include `PauseIngress` and `PurgeOldData`. Defaults to `PurgeOldData`.
	StorageLimitExceededBehavior pulumi.StringPtrOutput `pulumi:"storageLimitExceededBehavior"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewTimeSeriesInsightsStandardEnvironment registers a new resource with the given unique name, arguments, and options.
func NewTimeSeriesInsightsStandardEnvironment(ctx *pulumi.Context,
	name string, args *TimeSeriesInsightsStandardEnvironmentArgs, opts ...pulumi.ResourceOption) (*TimeSeriesInsightsStandardEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataRetentionTime == nil {
		return nil, errors.New("invalid value for required argument 'DataRetentionTime'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SkuName == nil {
		return nil, errors.New("invalid value for required argument 'SkuName'")
	}
	var resource TimeSeriesInsightsStandardEnvironment
	err := ctx.RegisterResource("azure:iot/timeSeriesInsightsStandardEnvironment:TimeSeriesInsightsStandardEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTimeSeriesInsightsStandardEnvironment gets an existing TimeSeriesInsightsStandardEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTimeSeriesInsightsStandardEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TimeSeriesInsightsStandardEnvironmentState, opts ...pulumi.ResourceOption) (*TimeSeriesInsightsStandardEnvironment, error) {
	var resource TimeSeriesInsightsStandardEnvironment
	err := ctx.ReadResource("azure:iot/timeSeriesInsightsStandardEnvironment:TimeSeriesInsightsStandardEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TimeSeriesInsightsStandardEnvironment resources.
type timeSeriesInsightsStandardEnvironmentState struct {
	// Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
	DataRetentionTime *string `pulumi:"dataRetentionTime"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Azure IoT Time Series Insights Standard Environment. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `pulumi:"name"`
	// The name of the event property which will be used to partition data. Changing this forces a new resource to be created.
	PartitionKey *string `pulumi:"partitionKey"`
	// The name of the resource group in which to create the Azure IoT Time Series Insights Standard Environment.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the SKU Name for this IoT Time Series Insights Standard Environment. It is string consisting of two parts separated by an underscore(\_).The fist part is the `name`, valid values include: `S1` and `S2`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `S1_1`). Changing this forces a new resource to be created.
	SkuName *string `pulumi:"skuName"`
	// Specifies the behaviour the IoT Time Series Insights service should take when the environment's capacity has been exceeded. Valid values include `PauseIngress` and `PurgeOldData`. Defaults to `PurgeOldData`.
	StorageLimitExceededBehavior *string `pulumi:"storageLimitExceededBehavior"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type TimeSeriesInsightsStandardEnvironmentState struct {
	// Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
	DataRetentionTime pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Azure IoT Time Series Insights Standard Environment. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringPtrInput
	// The name of the event property which will be used to partition data. Changing this forces a new resource to be created.
	PartitionKey pulumi.StringPtrInput
	// The name of the resource group in which to create the Azure IoT Time Series Insights Standard Environment.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the SKU Name for this IoT Time Series Insights Standard Environment. It is string consisting of two parts separated by an underscore(\_).The fist part is the `name`, valid values include: `S1` and `S2`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `S1_1`). Changing this forces a new resource to be created.
	SkuName pulumi.StringPtrInput
	// Specifies the behaviour the IoT Time Series Insights service should take when the environment's capacity has been exceeded. Valid values include `PauseIngress` and `PurgeOldData`. Defaults to `PurgeOldData`.
	StorageLimitExceededBehavior pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (TimeSeriesInsightsStandardEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*timeSeriesInsightsStandardEnvironmentState)(nil)).Elem()
}

type timeSeriesInsightsStandardEnvironmentArgs struct {
	// Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
	DataRetentionTime string `pulumi:"dataRetentionTime"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Azure IoT Time Series Insights Standard Environment. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `pulumi:"name"`
	// The name of the event property which will be used to partition data. Changing this forces a new resource to be created.
	PartitionKey *string `pulumi:"partitionKey"`
	// The name of the resource group in which to create the Azure IoT Time Series Insights Standard Environment.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the SKU Name for this IoT Time Series Insights Standard Environment. It is string consisting of two parts separated by an underscore(\_).The fist part is the `name`, valid values include: `S1` and `S2`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `S1_1`). Changing this forces a new resource to be created.
	SkuName string `pulumi:"skuName"`
	// Specifies the behaviour the IoT Time Series Insights service should take when the environment's capacity has been exceeded. Valid values include `PauseIngress` and `PurgeOldData`. Defaults to `PurgeOldData`.
	StorageLimitExceededBehavior *string `pulumi:"storageLimitExceededBehavior"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a TimeSeriesInsightsStandardEnvironment resource.
type TimeSeriesInsightsStandardEnvironmentArgs struct {
	// Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
	DataRetentionTime pulumi.StringInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Azure IoT Time Series Insights Standard Environment. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringPtrInput
	// The name of the event property which will be used to partition data. Changing this forces a new resource to be created.
	PartitionKey pulumi.StringPtrInput
	// The name of the resource group in which to create the Azure IoT Time Series Insights Standard Environment.
	ResourceGroupName pulumi.StringInput
	// Specifies the SKU Name for this IoT Time Series Insights Standard Environment. It is string consisting of two parts separated by an underscore(\_).The fist part is the `name`, valid values include: `S1` and `S2`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `S1_1`). Changing this forces a new resource to be created.
	SkuName pulumi.StringInput
	// Specifies the behaviour the IoT Time Series Insights service should take when the environment's capacity has been exceeded. Valid values include `PauseIngress` and `PurgeOldData`. Defaults to `PurgeOldData`.
	StorageLimitExceededBehavior pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (TimeSeriesInsightsStandardEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*timeSeriesInsightsStandardEnvironmentArgs)(nil)).Elem()
}

type TimeSeriesInsightsStandardEnvironmentInput interface {
	pulumi.Input

	ToTimeSeriesInsightsStandardEnvironmentOutput() TimeSeriesInsightsStandardEnvironmentOutput
	ToTimeSeriesInsightsStandardEnvironmentOutputWithContext(ctx context.Context) TimeSeriesInsightsStandardEnvironmentOutput
}

func (*TimeSeriesInsightsStandardEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeSeriesInsightsStandardEnvironment)(nil))
}

func (i *TimeSeriesInsightsStandardEnvironment) ToTimeSeriesInsightsStandardEnvironmentOutput() TimeSeriesInsightsStandardEnvironmentOutput {
	return i.ToTimeSeriesInsightsStandardEnvironmentOutputWithContext(context.Background())
}

func (i *TimeSeriesInsightsStandardEnvironment) ToTimeSeriesInsightsStandardEnvironmentOutputWithContext(ctx context.Context) TimeSeriesInsightsStandardEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsStandardEnvironmentOutput)
}

func (i *TimeSeriesInsightsStandardEnvironment) ToTimeSeriesInsightsStandardEnvironmentPtrOutput() TimeSeriesInsightsStandardEnvironmentPtrOutput {
	return i.ToTimeSeriesInsightsStandardEnvironmentPtrOutputWithContext(context.Background())
}

func (i *TimeSeriesInsightsStandardEnvironment) ToTimeSeriesInsightsStandardEnvironmentPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsStandardEnvironmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsStandardEnvironmentPtrOutput)
}

type TimeSeriesInsightsStandardEnvironmentPtrInput interface {
	pulumi.Input

	ToTimeSeriesInsightsStandardEnvironmentPtrOutput() TimeSeriesInsightsStandardEnvironmentPtrOutput
	ToTimeSeriesInsightsStandardEnvironmentPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsStandardEnvironmentPtrOutput
}

type TimeSeriesInsightsStandardEnvironmentOutput struct {
	*pulumi.OutputState
}

func (TimeSeriesInsightsStandardEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeSeriesInsightsStandardEnvironment)(nil))
}

func (o TimeSeriesInsightsStandardEnvironmentOutput) ToTimeSeriesInsightsStandardEnvironmentOutput() TimeSeriesInsightsStandardEnvironmentOutput {
	return o
}

func (o TimeSeriesInsightsStandardEnvironmentOutput) ToTimeSeriesInsightsStandardEnvironmentOutputWithContext(ctx context.Context) TimeSeriesInsightsStandardEnvironmentOutput {
	return o
}

type TimeSeriesInsightsStandardEnvironmentPtrOutput struct {
	*pulumi.OutputState
}

func (TimeSeriesInsightsStandardEnvironmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeSeriesInsightsStandardEnvironment)(nil))
}

func (o TimeSeriesInsightsStandardEnvironmentPtrOutput) ToTimeSeriesInsightsStandardEnvironmentPtrOutput() TimeSeriesInsightsStandardEnvironmentPtrOutput {
	return o
}

func (o TimeSeriesInsightsStandardEnvironmentPtrOutput) ToTimeSeriesInsightsStandardEnvironmentPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsStandardEnvironmentPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TimeSeriesInsightsStandardEnvironmentOutput{})
	pulumi.RegisterOutputType(TimeSeriesInsightsStandardEnvironmentPtrOutput{})
}

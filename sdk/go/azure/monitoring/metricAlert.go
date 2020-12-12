// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Metric Alert within Azure Monitor.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/monitoring"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		mainResourceGroup, err := core.NewResourceGroup(ctx, "mainResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		toMonitor, err := storage.NewAccount(ctx, "toMonitor", &storage.AccountArgs{
// 			ResourceGroupName:      mainResourceGroup.Name,
// 			Location:               mainResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		mainActionGroup, err := monitoring.NewActionGroup(ctx, "mainActionGroup", &monitoring.ActionGroupArgs{
// 			ResourceGroupName: mainResourceGroup.Name,
// 			ShortName:         pulumi.String("exampleact"),
// 			WebhookReceivers: monitoring.ActionGroupWebhookReceiverArray{
// 				&monitoring.ActionGroupWebhookReceiverArgs{
// 					Name:       pulumi.String("callmyapi"),
// 					ServiceUri: pulumi.String("http://example.com/alert"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = monitoring.NewMetricAlert(ctx, "example", &monitoring.MetricAlertArgs{
// 			ResourceGroupName: mainResourceGroup.Name,
// 			Scopes: pulumi.StringArray{
// 				toMonitor.ID(),
// 			},
// 			Description: pulumi.String("Action will be triggered when Transactions count is greater than 50."),
// 			Criterias: monitoring.MetricAlertCriteriaArray{
// 				&monitoring.MetricAlertCriteriaArgs{
// 					MetricNamespace: pulumi.String("Microsoft.Storage/storageAccounts"),
// 					MetricName:      pulumi.String("Transactions"),
// 					Aggregation:     pulumi.String("Total"),
// 					Operator:        pulumi.String("GreaterThan"),
// 					Threshold:       pulumi.Float64(50),
// 					Dimensions: monitoring.MetricAlertCriteriaDimensionArray{
// 						&monitoring.MetricAlertCriteriaDimensionArgs{
// 							Name:     pulumi.String("ApiName"),
// 							Operator: pulumi.String("Include"),
// 							Values: pulumi.StringArray{
// 								pulumi.String("*"),
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Actions: monitoring.MetricAlertActionArray{
// 				&monitoring.MetricAlertActionArgs{
// 					ActionGroupId: mainActionGroup.ID(),
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
// Metric Alerts can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:monitoring/metricAlert:MetricAlert main /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/microsoft.insights/metricalerts/example-metricalert
// ```
type MetricAlert struct {
	pulumi.CustomResourceState

	// One or more `action` blocks as defined below.
	Actions MetricAlertActionArrayOutput `pulumi:"actions"`
	// A `applicationInsightsWebTestLocationAvailabilityCriteria` block as defined below.
	ApplicationInsightsWebTestLocationAvailabilityCriteria MetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaPtrOutput `pulumi:"applicationInsightsWebTestLocationAvailabilityCriteria"`
	// Should the alerts in this Metric Alert be auto resolved? Defaults to `true`.
	AutoMitigate pulumi.BoolPtrOutput `pulumi:"autoMitigate"`
	// One or more (static) `criteria` blocks as defined below.
	Criterias MetricAlertCriteriaArrayOutput `pulumi:"criterias"`
	// The description of this Metric Alert.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A `dynamicCriteria` block as defined below.
	DynamicCriteria MetricAlertDynamicCriteriaPtrOutput `pulumi:"dynamicCriteria"`
	// Should this Metric Alert be enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The evaluation frequency of this Metric Alert, represented in ISO 8601 duration format. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M` and `PT1H`. Defaults to `PT1M`.
	Frequency pulumi.StringPtrOutput `pulumi:"frequency"`
	// The name of the Metric Alert. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Metric Alert instance.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A set of strings of resource IDs at which the metric criteria should be applied.
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// The severity of this Metric Alert. Possible values are `0`, `1`, `2`, `3` and `4`. Defaults to `3`.
	Severity pulumi.IntPtrOutput `pulumi:"severity"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The location of the target resource.
	TargetResourceLocation pulumi.StringOutput `pulumi:"targetResourceLocation"`
	// The resource type (e.g. `Microsoft.Compute/virtualMachines`) of the target resource.
	TargetResourceType pulumi.StringOutput `pulumi:"targetResourceType"`
	// The period of time that is used to monitor alert activity, represented in ISO 8601 duration format. This value must be greater than `frequency`. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M`, `PT1H`, `PT6H`, `PT12H` and `P1D`. Defaults to `PT5M`.
	WindowSize pulumi.StringPtrOutput `pulumi:"windowSize"`
}

// NewMetricAlert registers a new resource with the given unique name, arguments, and options.
func NewMetricAlert(ctx *pulumi.Context,
	name string, args *MetricAlertArgs, opts ...pulumi.ResourceOption) (*MetricAlert, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Scopes == nil {
		return nil, errors.New("invalid value for required argument 'Scopes'")
	}
	var resource MetricAlert
	err := ctx.RegisterResource("azure:monitoring/metricAlert:MetricAlert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetricAlert gets an existing MetricAlert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetricAlert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetricAlertState, opts ...pulumi.ResourceOption) (*MetricAlert, error) {
	var resource MetricAlert
	err := ctx.ReadResource("azure:monitoring/metricAlert:MetricAlert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetricAlert resources.
type metricAlertState struct {
	// One or more `action` blocks as defined below.
	Actions []MetricAlertAction `pulumi:"actions"`
	// A `applicationInsightsWebTestLocationAvailabilityCriteria` block as defined below.
	ApplicationInsightsWebTestLocationAvailabilityCriteria *MetricAlertApplicationInsightsWebTestLocationAvailabilityCriteria `pulumi:"applicationInsightsWebTestLocationAvailabilityCriteria"`
	// Should the alerts in this Metric Alert be auto resolved? Defaults to `true`.
	AutoMitigate *bool `pulumi:"autoMitigate"`
	// One or more (static) `criteria` blocks as defined below.
	Criterias []MetricAlertCriteria `pulumi:"criterias"`
	// The description of this Metric Alert.
	Description *string `pulumi:"description"`
	// A `dynamicCriteria` block as defined below.
	DynamicCriteria *MetricAlertDynamicCriteria `pulumi:"dynamicCriteria"`
	// Should this Metric Alert be enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The evaluation frequency of this Metric Alert, represented in ISO 8601 duration format. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M` and `PT1H`. Defaults to `PT1M`.
	Frequency *string `pulumi:"frequency"`
	// The name of the Metric Alert. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Metric Alert instance.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A set of strings of resource IDs at which the metric criteria should be applied.
	Scopes []string `pulumi:"scopes"`
	// The severity of this Metric Alert. Possible values are `0`, `1`, `2`, `3` and `4`. Defaults to `3`.
	Severity *int `pulumi:"severity"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The location of the target resource.
	TargetResourceLocation *string `pulumi:"targetResourceLocation"`
	// The resource type (e.g. `Microsoft.Compute/virtualMachines`) of the target resource.
	TargetResourceType *string `pulumi:"targetResourceType"`
	// The period of time that is used to monitor alert activity, represented in ISO 8601 duration format. This value must be greater than `frequency`. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M`, `PT1H`, `PT6H`, `PT12H` and `P1D`. Defaults to `PT5M`.
	WindowSize *string `pulumi:"windowSize"`
}

type MetricAlertState struct {
	// One or more `action` blocks as defined below.
	Actions MetricAlertActionArrayInput
	// A `applicationInsightsWebTestLocationAvailabilityCriteria` block as defined below.
	ApplicationInsightsWebTestLocationAvailabilityCriteria MetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaPtrInput
	// Should the alerts in this Metric Alert be auto resolved? Defaults to `true`.
	AutoMitigate pulumi.BoolPtrInput
	// One or more (static) `criteria` blocks as defined below.
	Criterias MetricAlertCriteriaArrayInput
	// The description of this Metric Alert.
	Description pulumi.StringPtrInput
	// A `dynamicCriteria` block as defined below.
	DynamicCriteria MetricAlertDynamicCriteriaPtrInput
	// Should this Metric Alert be enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The evaluation frequency of this Metric Alert, represented in ISO 8601 duration format. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M` and `PT1H`. Defaults to `PT1M`.
	Frequency pulumi.StringPtrInput
	// The name of the Metric Alert. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Metric Alert instance.
	ResourceGroupName pulumi.StringPtrInput
	// A set of strings of resource IDs at which the metric criteria should be applied.
	Scopes pulumi.StringArrayInput
	// The severity of this Metric Alert. Possible values are `0`, `1`, `2`, `3` and `4`. Defaults to `3`.
	Severity pulumi.IntPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The location of the target resource.
	TargetResourceLocation pulumi.StringPtrInput
	// The resource type (e.g. `Microsoft.Compute/virtualMachines`) of the target resource.
	TargetResourceType pulumi.StringPtrInput
	// The period of time that is used to monitor alert activity, represented in ISO 8601 duration format. This value must be greater than `frequency`. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M`, `PT1H`, `PT6H`, `PT12H` and `P1D`. Defaults to `PT5M`.
	WindowSize pulumi.StringPtrInput
}

func (MetricAlertState) ElementType() reflect.Type {
	return reflect.TypeOf((*metricAlertState)(nil)).Elem()
}

type metricAlertArgs struct {
	// One or more `action` blocks as defined below.
	Actions []MetricAlertAction `pulumi:"actions"`
	// A `applicationInsightsWebTestLocationAvailabilityCriteria` block as defined below.
	ApplicationInsightsWebTestLocationAvailabilityCriteria *MetricAlertApplicationInsightsWebTestLocationAvailabilityCriteria `pulumi:"applicationInsightsWebTestLocationAvailabilityCriteria"`
	// Should the alerts in this Metric Alert be auto resolved? Defaults to `true`.
	AutoMitigate *bool `pulumi:"autoMitigate"`
	// One or more (static) `criteria` blocks as defined below.
	Criterias []MetricAlertCriteria `pulumi:"criterias"`
	// The description of this Metric Alert.
	Description *string `pulumi:"description"`
	// A `dynamicCriteria` block as defined below.
	DynamicCriteria *MetricAlertDynamicCriteria `pulumi:"dynamicCriteria"`
	// Should this Metric Alert be enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The evaluation frequency of this Metric Alert, represented in ISO 8601 duration format. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M` and `PT1H`. Defaults to `PT1M`.
	Frequency *string `pulumi:"frequency"`
	// The name of the Metric Alert. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Metric Alert instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A set of strings of resource IDs at which the metric criteria should be applied.
	Scopes []string `pulumi:"scopes"`
	// The severity of this Metric Alert. Possible values are `0`, `1`, `2`, `3` and `4`. Defaults to `3`.
	Severity *int `pulumi:"severity"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The location of the target resource.
	TargetResourceLocation *string `pulumi:"targetResourceLocation"`
	// The resource type (e.g. `Microsoft.Compute/virtualMachines`) of the target resource.
	TargetResourceType *string `pulumi:"targetResourceType"`
	// The period of time that is used to monitor alert activity, represented in ISO 8601 duration format. This value must be greater than `frequency`. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M`, `PT1H`, `PT6H`, `PT12H` and `P1D`. Defaults to `PT5M`.
	WindowSize *string `pulumi:"windowSize"`
}

// The set of arguments for constructing a MetricAlert resource.
type MetricAlertArgs struct {
	// One or more `action` blocks as defined below.
	Actions MetricAlertActionArrayInput
	// A `applicationInsightsWebTestLocationAvailabilityCriteria` block as defined below.
	ApplicationInsightsWebTestLocationAvailabilityCriteria MetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaPtrInput
	// Should the alerts in this Metric Alert be auto resolved? Defaults to `true`.
	AutoMitigate pulumi.BoolPtrInput
	// One or more (static) `criteria` blocks as defined below.
	Criterias MetricAlertCriteriaArrayInput
	// The description of this Metric Alert.
	Description pulumi.StringPtrInput
	// A `dynamicCriteria` block as defined below.
	DynamicCriteria MetricAlertDynamicCriteriaPtrInput
	// Should this Metric Alert be enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The evaluation frequency of this Metric Alert, represented in ISO 8601 duration format. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M` and `PT1H`. Defaults to `PT1M`.
	Frequency pulumi.StringPtrInput
	// The name of the Metric Alert. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Metric Alert instance.
	ResourceGroupName pulumi.StringInput
	// A set of strings of resource IDs at which the metric criteria should be applied.
	Scopes pulumi.StringArrayInput
	// The severity of this Metric Alert. Possible values are `0`, `1`, `2`, `3` and `4`. Defaults to `3`.
	Severity pulumi.IntPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The location of the target resource.
	TargetResourceLocation pulumi.StringPtrInput
	// The resource type (e.g. `Microsoft.Compute/virtualMachines`) of the target resource.
	TargetResourceType pulumi.StringPtrInput
	// The period of time that is used to monitor alert activity, represented in ISO 8601 duration format. This value must be greater than `frequency`. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M`, `PT1H`, `PT6H`, `PT12H` and `P1D`. Defaults to `PT5M`.
	WindowSize pulumi.StringPtrInput
}

func (MetricAlertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metricAlertArgs)(nil)).Elem()
}

type MetricAlertInput interface {
	pulumi.Input

	ToMetricAlertOutput() MetricAlertOutput
	ToMetricAlertOutputWithContext(ctx context.Context) MetricAlertOutput
}

func (*MetricAlert) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricAlert)(nil))
}

func (i *MetricAlert) ToMetricAlertOutput() MetricAlertOutput {
	return i.ToMetricAlertOutputWithContext(context.Background())
}

func (i *MetricAlert) ToMetricAlertOutputWithContext(ctx context.Context) MetricAlertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricAlertOutput)
}

func (i *MetricAlert) ToMetricAlertPtrOutput() MetricAlertPtrOutput {
	return i.ToMetricAlertPtrOutputWithContext(context.Background())
}

func (i *MetricAlert) ToMetricAlertPtrOutputWithContext(ctx context.Context) MetricAlertPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricAlertPtrOutput)
}

type MetricAlertPtrInput interface {
	pulumi.Input

	ToMetricAlertPtrOutput() MetricAlertPtrOutput
	ToMetricAlertPtrOutputWithContext(ctx context.Context) MetricAlertPtrOutput
}

type MetricAlertOutput struct {
	*pulumi.OutputState
}

func (MetricAlertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricAlert)(nil))
}

func (o MetricAlertOutput) ToMetricAlertOutput() MetricAlertOutput {
	return o
}

func (o MetricAlertOutput) ToMetricAlertOutputWithContext(ctx context.Context) MetricAlertOutput {
	return o
}

type MetricAlertPtrOutput struct {
	*pulumi.OutputState
}

func (MetricAlertPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricAlert)(nil))
}

func (o MetricAlertPtrOutput) ToMetricAlertPtrOutput() MetricAlertPtrOutput {
	return o
}

func (o MetricAlertPtrOutput) ToMetricAlertPtrOutputWithContext(ctx context.Context) MetricAlertPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MetricAlertOutput{})
	pulumi.RegisterOutputType(MetricAlertPtrOutput{})
}

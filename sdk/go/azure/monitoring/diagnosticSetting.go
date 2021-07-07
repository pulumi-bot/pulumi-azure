// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Diagnostic Setting for an existing Resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/keyvault"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/monitoring"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
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
// 		_, err = monitoring.NewDiagnosticSetting(ctx, "exampleDiagnosticSetting", &monitoring.DiagnosticSettingArgs{
// 			TargetResourceId: exampleKeyVault.ApplyT(func(exampleKeyVault keyvault.LookupKeyVaultResult) (string, error) {
// 				return exampleKeyVault.Id, nil
// 			}).(pulumi.StringOutput),
// 			StorageAccountId: exampleAccount.ApplyT(func(exampleAccount storage.LookupAccountResult) (string, error) {
// 				return exampleAccount.Id, nil
// 			}).(pulumi.StringOutput),
// 			Logs: monitoring.DiagnosticSettingLogArray{
// 				&monitoring.DiagnosticSettingLogArgs{
// 					Category: pulumi.String("AuditEvent"),
// 					Enabled:  pulumi.Bool(false),
// 					RetentionPolicy: &monitoring.DiagnosticSettingLogRetentionPolicyArgs{
// 						Enabled: pulumi.Bool(false),
// 					},
// 				},
// 			},
// 			Metrics: monitoring.DiagnosticSettingMetricArray{
// 				&monitoring.DiagnosticSettingMetricArgs{
// 					Category: pulumi.String("AllMetrics"),
// 					RetentionPolicy: &monitoring.DiagnosticSettingMetricRetentionPolicyArgs{
// 						Enabled: pulumi.Bool(false),
// 					},
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
// Diagnostic Settings can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:monitoring/diagnosticSetting:DiagnosticSetting example "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.KeyVault/vaults/vault1|logMonitoring1"
// ```
type DiagnosticSetting struct {
	pulumi.CustomResourceState

	// Specifies the ID of an Event Hub Namespace Authorization Rule used to send Diagnostics Data. Changing this forces a new resource to be created.
	EventhubAuthorizationRuleId pulumi.StringPtrOutput `pulumi:"eventhubAuthorizationRuleId"`
	// Specifies the name of the Event Hub where Diagnostics Data should be sent. Changing this forces a new resource to be created.
	EventhubName pulumi.StringPtrOutput `pulumi:"eventhubName"`
	// When set to 'Dedicated' logs sent to a Log Analytics workspace will go into resource specific tables, instead of the legacy AzureDiagnostics table.
	LogAnalyticsDestinationType pulumi.StringPtrOutput `pulumi:"logAnalyticsDestinationType"`
	// Specifies the ID of a Log Analytics Workspace where Diagnostics Data should be sent.
	LogAnalyticsWorkspaceId pulumi.StringPtrOutput `pulumi:"logAnalyticsWorkspaceId"`
	// One or more `log` blocks as defined below.
	Logs DiagnosticSettingLogArrayOutput `pulumi:"logs"`
	// One or more `metric` blocks as defined below.
	Metrics DiagnosticSettingMetricArrayOutput `pulumi:"metrics"`
	// Specifies the name of the Diagnostic Setting. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the Storage Account where logs should be sent. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringPtrOutput `pulumi:"storageAccountId"`
	// The ID of an existing Resource on which to configure Diagnostic Settings. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringOutput `pulumi:"targetResourceId"`
}

// NewDiagnosticSetting registers a new resource with the given unique name, arguments, and options.
func NewDiagnosticSetting(ctx *pulumi.Context,
	name string, args *DiagnosticSettingArgs, opts ...pulumi.ResourceOption) (*DiagnosticSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TargetResourceId == nil {
		return nil, errors.New("invalid value for required argument 'TargetResourceId'")
	}
	var resource DiagnosticSetting
	err := ctx.RegisterResource("azure:monitoring/diagnosticSetting:DiagnosticSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiagnosticSetting gets an existing DiagnosticSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiagnosticSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiagnosticSettingState, opts ...pulumi.ResourceOption) (*DiagnosticSetting, error) {
	var resource DiagnosticSetting
	err := ctx.ReadResource("azure:monitoring/diagnosticSetting:DiagnosticSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DiagnosticSetting resources.
type diagnosticSettingState struct {
	// Specifies the ID of an Event Hub Namespace Authorization Rule used to send Diagnostics Data. Changing this forces a new resource to be created.
	EventhubAuthorizationRuleId *string `pulumi:"eventhubAuthorizationRuleId"`
	// Specifies the name of the Event Hub where Diagnostics Data should be sent. Changing this forces a new resource to be created.
	EventhubName *string `pulumi:"eventhubName"`
	// When set to 'Dedicated' logs sent to a Log Analytics workspace will go into resource specific tables, instead of the legacy AzureDiagnostics table.
	LogAnalyticsDestinationType *string `pulumi:"logAnalyticsDestinationType"`
	// Specifies the ID of a Log Analytics Workspace where Diagnostics Data should be sent.
	LogAnalyticsWorkspaceId *string `pulumi:"logAnalyticsWorkspaceId"`
	// One or more `log` blocks as defined below.
	Logs []DiagnosticSettingLog `pulumi:"logs"`
	// One or more `metric` blocks as defined below.
	Metrics []DiagnosticSettingMetric `pulumi:"metrics"`
	// Specifies the name of the Diagnostic Setting. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the Storage Account where logs should be sent. Changing this forces a new resource to be created.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// The ID of an existing Resource on which to configure Diagnostic Settings. Changing this forces a new resource to be created.
	TargetResourceId *string `pulumi:"targetResourceId"`
}

type DiagnosticSettingState struct {
	// Specifies the ID of an Event Hub Namespace Authorization Rule used to send Diagnostics Data. Changing this forces a new resource to be created.
	EventhubAuthorizationRuleId pulumi.StringPtrInput
	// Specifies the name of the Event Hub where Diagnostics Data should be sent. Changing this forces a new resource to be created.
	EventhubName pulumi.StringPtrInput
	// When set to 'Dedicated' logs sent to a Log Analytics workspace will go into resource specific tables, instead of the legacy AzureDiagnostics table.
	LogAnalyticsDestinationType pulumi.StringPtrInput
	// Specifies the ID of a Log Analytics Workspace where Diagnostics Data should be sent.
	LogAnalyticsWorkspaceId pulumi.StringPtrInput
	// One or more `log` blocks as defined below.
	Logs DiagnosticSettingLogArrayInput
	// One or more `metric` blocks as defined below.
	Metrics DiagnosticSettingMetricArrayInput
	// Specifies the name of the Diagnostic Setting. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the Storage Account where logs should be sent. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringPtrInput
	// The ID of an existing Resource on which to configure Diagnostic Settings. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringPtrInput
}

func (DiagnosticSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticSettingState)(nil)).Elem()
}

type diagnosticSettingArgs struct {
	// Specifies the ID of an Event Hub Namespace Authorization Rule used to send Diagnostics Data. Changing this forces a new resource to be created.
	EventhubAuthorizationRuleId *string `pulumi:"eventhubAuthorizationRuleId"`
	// Specifies the name of the Event Hub where Diagnostics Data should be sent. Changing this forces a new resource to be created.
	EventhubName *string `pulumi:"eventhubName"`
	// When set to 'Dedicated' logs sent to a Log Analytics workspace will go into resource specific tables, instead of the legacy AzureDiagnostics table.
	LogAnalyticsDestinationType *string `pulumi:"logAnalyticsDestinationType"`
	// Specifies the ID of a Log Analytics Workspace where Diagnostics Data should be sent.
	LogAnalyticsWorkspaceId *string `pulumi:"logAnalyticsWorkspaceId"`
	// One or more `log` blocks as defined below.
	Logs []DiagnosticSettingLog `pulumi:"logs"`
	// One or more `metric` blocks as defined below.
	Metrics []DiagnosticSettingMetric `pulumi:"metrics"`
	// Specifies the name of the Diagnostic Setting. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the Storage Account where logs should be sent. Changing this forces a new resource to be created.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// The ID of an existing Resource on which to configure Diagnostic Settings. Changing this forces a new resource to be created.
	TargetResourceId string `pulumi:"targetResourceId"`
}

// The set of arguments for constructing a DiagnosticSetting resource.
type DiagnosticSettingArgs struct {
	// Specifies the ID of an Event Hub Namespace Authorization Rule used to send Diagnostics Data. Changing this forces a new resource to be created.
	EventhubAuthorizationRuleId pulumi.StringPtrInput
	// Specifies the name of the Event Hub where Diagnostics Data should be sent. Changing this forces a new resource to be created.
	EventhubName pulumi.StringPtrInput
	// When set to 'Dedicated' logs sent to a Log Analytics workspace will go into resource specific tables, instead of the legacy AzureDiagnostics table.
	LogAnalyticsDestinationType pulumi.StringPtrInput
	// Specifies the ID of a Log Analytics Workspace where Diagnostics Data should be sent.
	LogAnalyticsWorkspaceId pulumi.StringPtrInput
	// One or more `log` blocks as defined below.
	Logs DiagnosticSettingLogArrayInput
	// One or more `metric` blocks as defined below.
	Metrics DiagnosticSettingMetricArrayInput
	// Specifies the name of the Diagnostic Setting. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the Storage Account where logs should be sent. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringPtrInput
	// The ID of an existing Resource on which to configure Diagnostic Settings. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringInput
}

func (DiagnosticSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticSettingArgs)(nil)).Elem()
}

type DiagnosticSettingInput interface {
	pulumi.Input

	ToDiagnosticSettingOutput() DiagnosticSettingOutput
	ToDiagnosticSettingOutputWithContext(ctx context.Context) DiagnosticSettingOutput
}

func (*DiagnosticSetting) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticSetting)(nil))
}

func (i *DiagnosticSetting) ToDiagnosticSettingOutput() DiagnosticSettingOutput {
	return i.ToDiagnosticSettingOutputWithContext(context.Background())
}

func (i *DiagnosticSetting) ToDiagnosticSettingOutputWithContext(ctx context.Context) DiagnosticSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticSettingOutput)
}

func (i *DiagnosticSetting) ToDiagnosticSettingPtrOutput() DiagnosticSettingPtrOutput {
	return i.ToDiagnosticSettingPtrOutputWithContext(context.Background())
}

func (i *DiagnosticSetting) ToDiagnosticSettingPtrOutputWithContext(ctx context.Context) DiagnosticSettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticSettingPtrOutput)
}

type DiagnosticSettingPtrInput interface {
	pulumi.Input

	ToDiagnosticSettingPtrOutput() DiagnosticSettingPtrOutput
	ToDiagnosticSettingPtrOutputWithContext(ctx context.Context) DiagnosticSettingPtrOutput
}

type diagnosticSettingPtrType DiagnosticSettingArgs

func (*diagnosticSettingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticSetting)(nil))
}

func (i *diagnosticSettingPtrType) ToDiagnosticSettingPtrOutput() DiagnosticSettingPtrOutput {
	return i.ToDiagnosticSettingPtrOutputWithContext(context.Background())
}

func (i *diagnosticSettingPtrType) ToDiagnosticSettingPtrOutputWithContext(ctx context.Context) DiagnosticSettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticSettingPtrOutput)
}

// DiagnosticSettingArrayInput is an input type that accepts DiagnosticSettingArray and DiagnosticSettingArrayOutput values.
// You can construct a concrete instance of `DiagnosticSettingArrayInput` via:
//
//          DiagnosticSettingArray{ DiagnosticSettingArgs{...} }
type DiagnosticSettingArrayInput interface {
	pulumi.Input

	ToDiagnosticSettingArrayOutput() DiagnosticSettingArrayOutput
	ToDiagnosticSettingArrayOutputWithContext(context.Context) DiagnosticSettingArrayOutput
}

type DiagnosticSettingArray []DiagnosticSettingInput

func (DiagnosticSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DiagnosticSetting)(nil))
}

func (i DiagnosticSettingArray) ToDiagnosticSettingArrayOutput() DiagnosticSettingArrayOutput {
	return i.ToDiagnosticSettingArrayOutputWithContext(context.Background())
}

func (i DiagnosticSettingArray) ToDiagnosticSettingArrayOutputWithContext(ctx context.Context) DiagnosticSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticSettingArrayOutput)
}

// DiagnosticSettingMapInput is an input type that accepts DiagnosticSettingMap and DiagnosticSettingMapOutput values.
// You can construct a concrete instance of `DiagnosticSettingMapInput` via:
//
//          DiagnosticSettingMap{ "key": DiagnosticSettingArgs{...} }
type DiagnosticSettingMapInput interface {
	pulumi.Input

	ToDiagnosticSettingMapOutput() DiagnosticSettingMapOutput
	ToDiagnosticSettingMapOutputWithContext(context.Context) DiagnosticSettingMapOutput
}

type DiagnosticSettingMap map[string]DiagnosticSettingInput

func (DiagnosticSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DiagnosticSetting)(nil))
}

func (i DiagnosticSettingMap) ToDiagnosticSettingMapOutput() DiagnosticSettingMapOutput {
	return i.ToDiagnosticSettingMapOutputWithContext(context.Background())
}

func (i DiagnosticSettingMap) ToDiagnosticSettingMapOutputWithContext(ctx context.Context) DiagnosticSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticSettingMapOutput)
}

type DiagnosticSettingOutput struct {
	*pulumi.OutputState
}

func (DiagnosticSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticSetting)(nil))
}

func (o DiagnosticSettingOutput) ToDiagnosticSettingOutput() DiagnosticSettingOutput {
	return o
}

func (o DiagnosticSettingOutput) ToDiagnosticSettingOutputWithContext(ctx context.Context) DiagnosticSettingOutput {
	return o
}

func (o DiagnosticSettingOutput) ToDiagnosticSettingPtrOutput() DiagnosticSettingPtrOutput {
	return o.ToDiagnosticSettingPtrOutputWithContext(context.Background())
}

func (o DiagnosticSettingOutput) ToDiagnosticSettingPtrOutputWithContext(ctx context.Context) DiagnosticSettingPtrOutput {
	return o.ApplyT(func(v DiagnosticSetting) *DiagnosticSetting {
		return &v
	}).(DiagnosticSettingPtrOutput)
}

type DiagnosticSettingPtrOutput struct {
	*pulumi.OutputState
}

func (DiagnosticSettingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticSetting)(nil))
}

func (o DiagnosticSettingPtrOutput) ToDiagnosticSettingPtrOutput() DiagnosticSettingPtrOutput {
	return o
}

func (o DiagnosticSettingPtrOutput) ToDiagnosticSettingPtrOutputWithContext(ctx context.Context) DiagnosticSettingPtrOutput {
	return o
}

type DiagnosticSettingArrayOutput struct{ *pulumi.OutputState }

func (DiagnosticSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiagnosticSetting)(nil))
}

func (o DiagnosticSettingArrayOutput) ToDiagnosticSettingArrayOutput() DiagnosticSettingArrayOutput {
	return o
}

func (o DiagnosticSettingArrayOutput) ToDiagnosticSettingArrayOutputWithContext(ctx context.Context) DiagnosticSettingArrayOutput {
	return o
}

func (o DiagnosticSettingArrayOutput) Index(i pulumi.IntInput) DiagnosticSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiagnosticSetting {
		return vs[0].([]DiagnosticSetting)[vs[1].(int)]
	}).(DiagnosticSettingOutput)
}

type DiagnosticSettingMapOutput struct{ *pulumi.OutputState }

func (DiagnosticSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DiagnosticSetting)(nil))
}

func (o DiagnosticSettingMapOutput) ToDiagnosticSettingMapOutput() DiagnosticSettingMapOutput {
	return o
}

func (o DiagnosticSettingMapOutput) ToDiagnosticSettingMapOutputWithContext(ctx context.Context) DiagnosticSettingMapOutput {
	return o
}

func (o DiagnosticSettingMapOutput) MapIndex(k pulumi.StringInput) DiagnosticSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DiagnosticSetting {
		return vs[0].(map[string]DiagnosticSetting)[vs[1].(string)]
	}).(DiagnosticSettingOutput)
}

func init() {
	pulumi.RegisterOutputType(DiagnosticSettingOutput{})
	pulumi.RegisterOutputType(DiagnosticSettingPtrOutput{})
	pulumi.RegisterOutputType(DiagnosticSettingArrayOutput{})
	pulumi.RegisterOutputType(DiagnosticSettingMapOutput{})
}

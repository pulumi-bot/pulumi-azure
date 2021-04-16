// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Activity Log Alert within Azure Monitor.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/monitoring"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		mainResourceGroup, err := core.NewResourceGroup(ctx, "mainResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		mainActionGroup, err := monitoring.NewActionGroup(ctx, "mainActionGroup", &monitoring.ActionGroupArgs{
// 			ResourceGroupName: mainResourceGroup.Name,
// 			ShortName:         pulumi.String("p0action"),
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
// 		toMonitor, err := storage.NewAccount(ctx, "toMonitor", &storage.AccountArgs{
// 			ResourceGroupName:      mainResourceGroup.Name,
// 			Location:               mainResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("GRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = monitoring.NewActivityLogAlert(ctx, "mainActivityLogAlert", &monitoring.ActivityLogAlertArgs{
// 			ResourceGroupName: mainResourceGroup.Name,
// 			Scopes: pulumi.StringArray{
// 				mainResourceGroup.ID(),
// 			},
// 			Description: pulumi.String("This alert will monitor a specific storage account updates."),
// 			Criteria: &monitoring.ActivityLogAlertCriteriaArgs{
// 				ResourceId:    toMonitor.ID(),
// 				OperationName: pulumi.String("Microsoft.Storage/storageAccounts/write"),
// 				Category:      pulumi.String("Recommendation"),
// 			},
// 			Actions: monitoring.ActivityLogAlertActionArray{
// 				&monitoring.ActivityLogAlertActionArgs{
// 					ActionGroupId: mainActionGroup.ID(),
// 					WebhookProperties: pulumi.StringMap{
// 						"from": pulumi.String("source"),
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
// Activity log alerts can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:monitoring/activityLogAlert:ActivityLogAlert example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/microsoft.insights/activityLogAlerts/myalertname
// ```
type ActivityLogAlert struct {
	pulumi.CustomResourceState

	// One or more `action` blocks as defined below.
	Actions ActivityLogAlertActionArrayOutput `pulumi:"actions"`
	// A `criteria` block as defined below.
	Criteria ActivityLogAlertCriteriaOutput `pulumi:"criteria"`
	// The description of this activity log alert.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Should this Activity Log Alert be enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The name of the activity log alert. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the activity log alert instance.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Scope at which the Activity Log should be applied, for example a the Resource ID of a Subscription or a Resource (such as a Storage Account).
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewActivityLogAlert registers a new resource with the given unique name, arguments, and options.
func NewActivityLogAlert(ctx *pulumi.Context,
	name string, args *ActivityLogAlertArgs, opts ...pulumi.ResourceOption) (*ActivityLogAlert, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Criteria == nil {
		return nil, errors.New("invalid value for required argument 'Criteria'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Scopes == nil {
		return nil, errors.New("invalid value for required argument 'Scopes'")
	}
	var resource ActivityLogAlert
	err := ctx.RegisterResource("azure:monitoring/activityLogAlert:ActivityLogAlert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActivityLogAlert gets an existing ActivityLogAlert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActivityLogAlert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActivityLogAlertState, opts ...pulumi.ResourceOption) (*ActivityLogAlert, error) {
	var resource ActivityLogAlert
	err := ctx.ReadResource("azure:monitoring/activityLogAlert:ActivityLogAlert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActivityLogAlert resources.
type activityLogAlertState struct {
	// One or more `action` blocks as defined below.
	Actions []ActivityLogAlertAction `pulumi:"actions"`
	// A `criteria` block as defined below.
	Criteria *ActivityLogAlertCriteria `pulumi:"criteria"`
	// The description of this activity log alert.
	Description *string `pulumi:"description"`
	// Should this Activity Log Alert be enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The name of the activity log alert. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the activity log alert instance.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Scope at which the Activity Log should be applied, for example a the Resource ID of a Subscription or a Resource (such as a Storage Account).
	Scopes []string `pulumi:"scopes"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ActivityLogAlertState struct {
	// One or more `action` blocks as defined below.
	Actions ActivityLogAlertActionArrayInput
	// A `criteria` block as defined below.
	Criteria ActivityLogAlertCriteriaPtrInput
	// The description of this activity log alert.
	Description pulumi.StringPtrInput
	// Should this Activity Log Alert be enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The name of the activity log alert. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the activity log alert instance.
	ResourceGroupName pulumi.StringPtrInput
	// The Scope at which the Activity Log should be applied, for example a the Resource ID of a Subscription or a Resource (such as a Storage Account).
	Scopes pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ActivityLogAlertState) ElementType() reflect.Type {
	return reflect.TypeOf((*activityLogAlertState)(nil)).Elem()
}

type activityLogAlertArgs struct {
	// One or more `action` blocks as defined below.
	Actions []ActivityLogAlertAction `pulumi:"actions"`
	// A `criteria` block as defined below.
	Criteria ActivityLogAlertCriteria `pulumi:"criteria"`
	// The description of this activity log alert.
	Description *string `pulumi:"description"`
	// Should this Activity Log Alert be enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The name of the activity log alert. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the activity log alert instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Scope at which the Activity Log should be applied, for example a the Resource ID of a Subscription or a Resource (such as a Storage Account).
	Scopes []string `pulumi:"scopes"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ActivityLogAlert resource.
type ActivityLogAlertArgs struct {
	// One or more `action` blocks as defined below.
	Actions ActivityLogAlertActionArrayInput
	// A `criteria` block as defined below.
	Criteria ActivityLogAlertCriteriaInput
	// The description of this activity log alert.
	Description pulumi.StringPtrInput
	// Should this Activity Log Alert be enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The name of the activity log alert. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the activity log alert instance.
	ResourceGroupName pulumi.StringInput
	// The Scope at which the Activity Log should be applied, for example a the Resource ID of a Subscription or a Resource (such as a Storage Account).
	Scopes pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ActivityLogAlertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*activityLogAlertArgs)(nil)).Elem()
}

type ActivityLogAlertInput interface {
	pulumi.Input

	ToActivityLogAlertOutput() ActivityLogAlertOutput
	ToActivityLogAlertOutputWithContext(ctx context.Context) ActivityLogAlertOutput
}

func (*ActivityLogAlert) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityLogAlert)(nil))
}

func (i *ActivityLogAlert) ToActivityLogAlertOutput() ActivityLogAlertOutput {
	return i.ToActivityLogAlertOutputWithContext(context.Background())
}

func (i *ActivityLogAlert) ToActivityLogAlertOutputWithContext(ctx context.Context) ActivityLogAlertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityLogAlertOutput)
}

func (i *ActivityLogAlert) ToActivityLogAlertPtrOutput() ActivityLogAlertPtrOutput {
	return i.ToActivityLogAlertPtrOutputWithContext(context.Background())
}

func (i *ActivityLogAlert) ToActivityLogAlertPtrOutputWithContext(ctx context.Context) ActivityLogAlertPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityLogAlertPtrOutput)
}

type ActivityLogAlertPtrInput interface {
	pulumi.Input

	ToActivityLogAlertPtrOutput() ActivityLogAlertPtrOutput
	ToActivityLogAlertPtrOutputWithContext(ctx context.Context) ActivityLogAlertPtrOutput
}

type activityLogAlertPtrType ActivityLogAlertArgs

func (*activityLogAlertPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActivityLogAlert)(nil))
}

func (i *activityLogAlertPtrType) ToActivityLogAlertPtrOutput() ActivityLogAlertPtrOutput {
	return i.ToActivityLogAlertPtrOutputWithContext(context.Background())
}

func (i *activityLogAlertPtrType) ToActivityLogAlertPtrOutputWithContext(ctx context.Context) ActivityLogAlertPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityLogAlertPtrOutput)
}

// ActivityLogAlertArrayInput is an input type that accepts ActivityLogAlertArray and ActivityLogAlertArrayOutput values.
// You can construct a concrete instance of `ActivityLogAlertArrayInput` via:
//
//          ActivityLogAlertArray{ ActivityLogAlertArgs{...} }
type ActivityLogAlertArrayInput interface {
	pulumi.Input

	ToActivityLogAlertArrayOutput() ActivityLogAlertArrayOutput
	ToActivityLogAlertArrayOutputWithContext(context.Context) ActivityLogAlertArrayOutput
}

type ActivityLogAlertArray []ActivityLogAlertInput

func (ActivityLogAlertArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ActivityLogAlert)(nil))
}

func (i ActivityLogAlertArray) ToActivityLogAlertArrayOutput() ActivityLogAlertArrayOutput {
	return i.ToActivityLogAlertArrayOutputWithContext(context.Background())
}

func (i ActivityLogAlertArray) ToActivityLogAlertArrayOutputWithContext(ctx context.Context) ActivityLogAlertArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityLogAlertArrayOutput)
}

// ActivityLogAlertMapInput is an input type that accepts ActivityLogAlertMap and ActivityLogAlertMapOutput values.
// You can construct a concrete instance of `ActivityLogAlertMapInput` via:
//
//          ActivityLogAlertMap{ "key": ActivityLogAlertArgs{...} }
type ActivityLogAlertMapInput interface {
	pulumi.Input

	ToActivityLogAlertMapOutput() ActivityLogAlertMapOutput
	ToActivityLogAlertMapOutputWithContext(context.Context) ActivityLogAlertMapOutput
}

type ActivityLogAlertMap map[string]ActivityLogAlertInput

func (ActivityLogAlertMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ActivityLogAlert)(nil))
}

func (i ActivityLogAlertMap) ToActivityLogAlertMapOutput() ActivityLogAlertMapOutput {
	return i.ToActivityLogAlertMapOutputWithContext(context.Background())
}

func (i ActivityLogAlertMap) ToActivityLogAlertMapOutputWithContext(ctx context.Context) ActivityLogAlertMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityLogAlertMapOutput)
}

type ActivityLogAlertOutput struct {
	*pulumi.OutputState
}

func (ActivityLogAlertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityLogAlert)(nil))
}

func (o ActivityLogAlertOutput) ToActivityLogAlertOutput() ActivityLogAlertOutput {
	return o
}

func (o ActivityLogAlertOutput) ToActivityLogAlertOutputWithContext(ctx context.Context) ActivityLogAlertOutput {
	return o
}

func (o ActivityLogAlertOutput) ToActivityLogAlertPtrOutput() ActivityLogAlertPtrOutput {
	return o.ToActivityLogAlertPtrOutputWithContext(context.Background())
}

func (o ActivityLogAlertOutput) ToActivityLogAlertPtrOutputWithContext(ctx context.Context) ActivityLogAlertPtrOutput {
	return o.ApplyT(func(v ActivityLogAlert) *ActivityLogAlert {
		return &v
	}).(ActivityLogAlertPtrOutput)
}

type ActivityLogAlertPtrOutput struct {
	*pulumi.OutputState
}

func (ActivityLogAlertPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActivityLogAlert)(nil))
}

func (o ActivityLogAlertPtrOutput) ToActivityLogAlertPtrOutput() ActivityLogAlertPtrOutput {
	return o
}

func (o ActivityLogAlertPtrOutput) ToActivityLogAlertPtrOutputWithContext(ctx context.Context) ActivityLogAlertPtrOutput {
	return o
}

type ActivityLogAlertArrayOutput struct{ *pulumi.OutputState }

func (ActivityLogAlertArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActivityLogAlert)(nil))
}

func (o ActivityLogAlertArrayOutput) ToActivityLogAlertArrayOutput() ActivityLogAlertArrayOutput {
	return o
}

func (o ActivityLogAlertArrayOutput) ToActivityLogAlertArrayOutputWithContext(ctx context.Context) ActivityLogAlertArrayOutput {
	return o
}

func (o ActivityLogAlertArrayOutput) Index(i pulumi.IntInput) ActivityLogAlertOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActivityLogAlert {
		return vs[0].([]ActivityLogAlert)[vs[1].(int)]
	}).(ActivityLogAlertOutput)
}

type ActivityLogAlertMapOutput struct{ *pulumi.OutputState }

func (ActivityLogAlertMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ActivityLogAlert)(nil))
}

func (o ActivityLogAlertMapOutput) ToActivityLogAlertMapOutput() ActivityLogAlertMapOutput {
	return o
}

func (o ActivityLogAlertMapOutput) ToActivityLogAlertMapOutputWithContext(ctx context.Context) ActivityLogAlertMapOutput {
	return o
}

func (o ActivityLogAlertMapOutput) MapIndex(k pulumi.StringInput) ActivityLogAlertOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ActivityLogAlert {
		return vs[0].(map[string]ActivityLogAlert)[vs[1].(string)]
	}).(ActivityLogAlertOutput)
}

func init() {
	pulumi.RegisterOutputType(ActivityLogAlertOutput{})
	pulumi.RegisterOutputType(ActivityLogAlertPtrOutput{})
	pulumi.RegisterOutputType(ActivityLogAlertArrayOutput{})
	pulumi.RegisterOutputType(ActivityLogAlertMapOutput{})
}

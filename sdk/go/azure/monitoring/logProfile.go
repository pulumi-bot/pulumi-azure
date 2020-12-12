// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a [Log Profile](https://docs.microsoft.com/en-us/azure/monitoring-and-diagnostics/monitoring-overview-activity-logs#export-the-activity-log-with-a-log-profile). A Log Profile configures how Activity Logs are exported.
//
// > **NOTE:** It's only possible to configure one Log Profile per Subscription. If you are trying to create more than one Log Profile, an error with `StatusCode=409` will occur.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/eventhub"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/monitoring"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("eastus"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("GRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleEventHubNamespace, err := eventhub.NewEventHubNamespace(ctx, "exampleEventHubNamespace", &eventhub.EventHubNamespaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Standard"),
// 			Capacity:          pulumi.Int(2),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = monitoring.NewLogProfile(ctx, "exampleLogProfile", &monitoring.LogProfileArgs{
// 			Categories: pulumi.StringArray{
// 				pulumi.String("Action"),
// 				pulumi.String("Delete"),
// 				pulumi.String("Write"),
// 			},
// 			Locations: pulumi.StringArray{
// 				pulumi.String("westus"),
// 				pulumi.String("global"),
// 			},
// 			ServicebusRuleId: exampleEventHubNamespace.ID().ApplyT(func(id string) (string, error) {
// 				return fmt.Sprintf("%v%v", id, "/authorizationrules/RootManageSharedAccessKey"), nil
// 			}).(pulumi.StringOutput),
// 			StorageAccountId: exampleAccount.ID(),
// 			RetentionPolicy: &monitoring.LogProfileRetentionPolicyArgs{
// 				Enabled: pulumi.Bool(true),
// 				Days:    pulumi.Int(7),
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
// A Log Profile can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:monitoring/logProfile:LogProfile example /subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.insights/logprofiles/test
// ```
type LogProfile struct {
	pulumi.CustomResourceState

	// List of categories of the logs.
	Categories pulumi.StringArrayOutput `pulumi:"categories"`
	// List of regions for which Activity Log events are stored or streamed.
	Locations pulumi.StringArrayOutput `pulumi:"locations"`
	// The name of the Log Profile. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `retentionPolicy` block as documented below. A retention policy for how long Activity Logs are retained in the storage account.
	RetentionPolicy LogProfileRetentionPolicyOutput `pulumi:"retentionPolicy"`
	// The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to. At least one of `storageAccountId` or `servicebusRuleId` must be set.
	ServicebusRuleId pulumi.StringPtrOutput `pulumi:"servicebusRuleId"`
	// The resource ID of the storage account in which the Activity Log is stored. At least one of `storageAccountId` or `servicebusRuleId` must be set.
	StorageAccountId pulumi.StringPtrOutput `pulumi:"storageAccountId"`
}

// NewLogProfile registers a new resource with the given unique name, arguments, and options.
func NewLogProfile(ctx *pulumi.Context,
	name string, args *LogProfileArgs, opts ...pulumi.ResourceOption) (*LogProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Categories == nil {
		return nil, errors.New("invalid value for required argument 'Categories'")
	}
	if args.Locations == nil {
		return nil, errors.New("invalid value for required argument 'Locations'")
	}
	if args.RetentionPolicy == nil {
		return nil, errors.New("invalid value for required argument 'RetentionPolicy'")
	}
	var resource LogProfile
	err := ctx.RegisterResource("azure:monitoring/logProfile:LogProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogProfile gets an existing LogProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogProfileState, opts ...pulumi.ResourceOption) (*LogProfile, error) {
	var resource LogProfile
	err := ctx.ReadResource("azure:monitoring/logProfile:LogProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogProfile resources.
type logProfileState struct {
	// List of categories of the logs.
	Categories []string `pulumi:"categories"`
	// List of regions for which Activity Log events are stored or streamed.
	Locations []string `pulumi:"locations"`
	// The name of the Log Profile. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// A `retentionPolicy` block as documented below. A retention policy for how long Activity Logs are retained in the storage account.
	RetentionPolicy *LogProfileRetentionPolicy `pulumi:"retentionPolicy"`
	// The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to. At least one of `storageAccountId` or `servicebusRuleId` must be set.
	ServicebusRuleId *string `pulumi:"servicebusRuleId"`
	// The resource ID of the storage account in which the Activity Log is stored. At least one of `storageAccountId` or `servicebusRuleId` must be set.
	StorageAccountId *string `pulumi:"storageAccountId"`
}

type LogProfileState struct {
	// List of categories of the logs.
	Categories pulumi.StringArrayInput
	// List of regions for which Activity Log events are stored or streamed.
	Locations pulumi.StringArrayInput
	// The name of the Log Profile. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// A `retentionPolicy` block as documented below. A retention policy for how long Activity Logs are retained in the storage account.
	RetentionPolicy LogProfileRetentionPolicyPtrInput
	// The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to. At least one of `storageAccountId` or `servicebusRuleId` must be set.
	ServicebusRuleId pulumi.StringPtrInput
	// The resource ID of the storage account in which the Activity Log is stored. At least one of `storageAccountId` or `servicebusRuleId` must be set.
	StorageAccountId pulumi.StringPtrInput
}

func (LogProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*logProfileState)(nil)).Elem()
}

type logProfileArgs struct {
	// List of categories of the logs.
	Categories []string `pulumi:"categories"`
	// List of regions for which Activity Log events are stored or streamed.
	Locations []string `pulumi:"locations"`
	// The name of the Log Profile. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// A `retentionPolicy` block as documented below. A retention policy for how long Activity Logs are retained in the storage account.
	RetentionPolicy LogProfileRetentionPolicy `pulumi:"retentionPolicy"`
	// The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to. At least one of `storageAccountId` or `servicebusRuleId` must be set.
	ServicebusRuleId *string `pulumi:"servicebusRuleId"`
	// The resource ID of the storage account in which the Activity Log is stored. At least one of `storageAccountId` or `servicebusRuleId` must be set.
	StorageAccountId *string `pulumi:"storageAccountId"`
}

// The set of arguments for constructing a LogProfile resource.
type LogProfileArgs struct {
	// List of categories of the logs.
	Categories pulumi.StringArrayInput
	// List of regions for which Activity Log events are stored or streamed.
	Locations pulumi.StringArrayInput
	// The name of the Log Profile. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// A `retentionPolicy` block as documented below. A retention policy for how long Activity Logs are retained in the storage account.
	RetentionPolicy LogProfileRetentionPolicyInput
	// The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to. At least one of `storageAccountId` or `servicebusRuleId` must be set.
	ServicebusRuleId pulumi.StringPtrInput
	// The resource ID of the storage account in which the Activity Log is stored. At least one of `storageAccountId` or `servicebusRuleId` must be set.
	StorageAccountId pulumi.StringPtrInput
}

func (LogProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logProfileArgs)(nil)).Elem()
}

type LogProfileInput interface {
	pulumi.Input

	ToLogProfileOutput() LogProfileOutput
	ToLogProfileOutputWithContext(ctx context.Context) LogProfileOutput
}

func (*LogProfile) ElementType() reflect.Type {
	return reflect.TypeOf((*LogProfile)(nil))
}

func (i *LogProfile) ToLogProfileOutput() LogProfileOutput {
	return i.ToLogProfileOutputWithContext(context.Background())
}

func (i *LogProfile) ToLogProfileOutputWithContext(ctx context.Context) LogProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogProfileOutput)
}

func (i *LogProfile) ToLogProfilePtrOutput() LogProfilePtrOutput {
	return i.ToLogProfilePtrOutputWithContext(context.Background())
}

func (i *LogProfile) ToLogProfilePtrOutputWithContext(ctx context.Context) LogProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogProfilePtrOutput)
}

type LogProfilePtrInput interface {
	pulumi.Input

	ToLogProfilePtrOutput() LogProfilePtrOutput
	ToLogProfilePtrOutputWithContext(ctx context.Context) LogProfilePtrOutput
}

type LogProfileOutput struct {
	*pulumi.OutputState
}

func (LogProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogProfile)(nil))
}

func (o LogProfileOutput) ToLogProfileOutput() LogProfileOutput {
	return o
}

func (o LogProfileOutput) ToLogProfileOutputWithContext(ctx context.Context) LogProfileOutput {
	return o
}

type LogProfilePtrOutput struct {
	*pulumi.OutputState
}

func (LogProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogProfile)(nil))
}

func (o LogProfilePtrOutput) ToLogProfilePtrOutput() LogProfilePtrOutput {
	return o
}

func (o LogProfilePtrOutput) ToLogProfilePtrOutputWithContext(ctx context.Context) LogProfilePtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LogProfileOutput{})
	pulumi.RegisterOutputType(LogProfilePtrOutput{})
}

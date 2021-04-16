// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loganalytics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Log Analytics Storage Insights resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/loganalytics"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/operationalinsights"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
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
// 		exampleAnalyticsWorkspace, err := operationalinsights.NewAnalyticsWorkspace(ctx, "exampleAnalyticsWorkspace", &operationalinsights.AnalyticsWorkspaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("PerGB2018"),
// 			RetentionInDays:   pulumi.Int(30),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = loganalytics.NewStorageInsights(ctx, "exampleStorageInsights", &loganalytics.StorageInsightsArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			WorkspaceId:       exampleAnalyticsWorkspace.ID(),
// 			StorageAccountId:  exampleAccount.ID(),
// 			StorageAccountKey: exampleAccount.PrimaryAccessKey,
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
// Log Analytics Storage Insight Configs can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:loganalytics/storageInsights:StorageInsights example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/storageInsightConfigs/storageInsight1
// ```
type StorageInsights struct {
	pulumi.CustomResourceState

	// The names of the blob containers that the workspace should read.
	BlobContainerNames pulumi.StringArrayOutput `pulumi:"blobContainerNames"`
	// The name which should be used for this Log Analytics Storage Insights. Changing this forces a new Log Analytics Storage Insights to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Log Analytics Storage Insights should exist. Changing this forces a new Log Analytics Storage Insights to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The ID of the Storage Account used by this Log Analytics Storage Insights.
	StorageAccountId pulumi.StringOutput `pulumi:"storageAccountId"`
	// The storage access key to be used to connect to the storage account.
	StorageAccountKey pulumi.StringOutput `pulumi:"storageAccountKey"`
	// The names of the Azure tables that the workspace should read.
	TableNames pulumi.StringArrayOutput `pulumi:"tableNames"`
	// A mapping of tags which should be assigned to the Log Analytics Storage Insights.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The ID of the Log Analytics Workspace within which the Storage Insights should exist. Changing this forces a new Log Analytics Storage Insights to be created.
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
}

// NewStorageInsights registers a new resource with the given unique name, arguments, and options.
func NewStorageInsights(ctx *pulumi.Context,
	name string, args *StorageInsightsArgs, opts ...pulumi.ResourceOption) (*StorageInsights, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccountId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountId'")
	}
	if args.StorageAccountKey == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountKey'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	var resource StorageInsights
	err := ctx.RegisterResource("azure:loganalytics/storageInsights:StorageInsights", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageInsights gets an existing StorageInsights resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageInsights(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageInsightsState, opts ...pulumi.ResourceOption) (*StorageInsights, error) {
	var resource StorageInsights
	err := ctx.ReadResource("azure:loganalytics/storageInsights:StorageInsights", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageInsights resources.
type storageInsightsState struct {
	// The names of the blob containers that the workspace should read.
	BlobContainerNames []string `pulumi:"blobContainerNames"`
	// The name which should be used for this Log Analytics Storage Insights. Changing this forces a new Log Analytics Storage Insights to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Log Analytics Storage Insights should exist. Changing this forces a new Log Analytics Storage Insights to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The ID of the Storage Account used by this Log Analytics Storage Insights.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// The storage access key to be used to connect to the storage account.
	StorageAccountKey *string `pulumi:"storageAccountKey"`
	// The names of the Azure tables that the workspace should read.
	TableNames []string `pulumi:"tableNames"`
	// A mapping of tags which should be assigned to the Log Analytics Storage Insights.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Log Analytics Workspace within which the Storage Insights should exist. Changing this forces a new Log Analytics Storage Insights to be created.
	WorkspaceId *string `pulumi:"workspaceId"`
}

type StorageInsightsState struct {
	// The names of the blob containers that the workspace should read.
	BlobContainerNames pulumi.StringArrayInput
	// The name which should be used for this Log Analytics Storage Insights. Changing this forces a new Log Analytics Storage Insights to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Log Analytics Storage Insights should exist. Changing this forces a new Log Analytics Storage Insights to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The ID of the Storage Account used by this Log Analytics Storage Insights.
	StorageAccountId pulumi.StringPtrInput
	// The storage access key to be used to connect to the storage account.
	StorageAccountKey pulumi.StringPtrInput
	// The names of the Azure tables that the workspace should read.
	TableNames pulumi.StringArrayInput
	// A mapping of tags which should be assigned to the Log Analytics Storage Insights.
	Tags pulumi.StringMapInput
	// The ID of the Log Analytics Workspace within which the Storage Insights should exist. Changing this forces a new Log Analytics Storage Insights to be created.
	WorkspaceId pulumi.StringPtrInput
}

func (StorageInsightsState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageInsightsState)(nil)).Elem()
}

type storageInsightsArgs struct {
	// The names of the blob containers that the workspace should read.
	BlobContainerNames []string `pulumi:"blobContainerNames"`
	// The name which should be used for this Log Analytics Storage Insights. Changing this forces a new Log Analytics Storage Insights to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Log Analytics Storage Insights should exist. Changing this forces a new Log Analytics Storage Insights to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ID of the Storage Account used by this Log Analytics Storage Insights.
	StorageAccountId string `pulumi:"storageAccountId"`
	// The storage access key to be used to connect to the storage account.
	StorageAccountKey string `pulumi:"storageAccountKey"`
	// The names of the Azure tables that the workspace should read.
	TableNames []string `pulumi:"tableNames"`
	// A mapping of tags which should be assigned to the Log Analytics Storage Insights.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Log Analytics Workspace within which the Storage Insights should exist. Changing this forces a new Log Analytics Storage Insights to be created.
	WorkspaceId string `pulumi:"workspaceId"`
}

// The set of arguments for constructing a StorageInsights resource.
type StorageInsightsArgs struct {
	// The names of the blob containers that the workspace should read.
	BlobContainerNames pulumi.StringArrayInput
	// The name which should be used for this Log Analytics Storage Insights. Changing this forces a new Log Analytics Storage Insights to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Log Analytics Storage Insights should exist. Changing this forces a new Log Analytics Storage Insights to be created.
	ResourceGroupName pulumi.StringInput
	// The ID of the Storage Account used by this Log Analytics Storage Insights.
	StorageAccountId pulumi.StringInput
	// The storage access key to be used to connect to the storage account.
	StorageAccountKey pulumi.StringInput
	// The names of the Azure tables that the workspace should read.
	TableNames pulumi.StringArrayInput
	// A mapping of tags which should be assigned to the Log Analytics Storage Insights.
	Tags pulumi.StringMapInput
	// The ID of the Log Analytics Workspace within which the Storage Insights should exist. Changing this forces a new Log Analytics Storage Insights to be created.
	WorkspaceId pulumi.StringInput
}

func (StorageInsightsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageInsightsArgs)(nil)).Elem()
}

type StorageInsightsInput interface {
	pulumi.Input

	ToStorageInsightsOutput() StorageInsightsOutput
	ToStorageInsightsOutputWithContext(ctx context.Context) StorageInsightsOutput
}

func (*StorageInsights) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageInsights)(nil))
}

func (i *StorageInsights) ToStorageInsightsOutput() StorageInsightsOutput {
	return i.ToStorageInsightsOutputWithContext(context.Background())
}

func (i *StorageInsights) ToStorageInsightsOutputWithContext(ctx context.Context) StorageInsightsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageInsightsOutput)
}

func (i *StorageInsights) ToStorageInsightsPtrOutput() StorageInsightsPtrOutput {
	return i.ToStorageInsightsPtrOutputWithContext(context.Background())
}

func (i *StorageInsights) ToStorageInsightsPtrOutputWithContext(ctx context.Context) StorageInsightsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageInsightsPtrOutput)
}

type StorageInsightsPtrInput interface {
	pulumi.Input

	ToStorageInsightsPtrOutput() StorageInsightsPtrOutput
	ToStorageInsightsPtrOutputWithContext(ctx context.Context) StorageInsightsPtrOutput
}

type storageInsightsPtrType StorageInsightsArgs

func (*storageInsightsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageInsights)(nil))
}

func (i *storageInsightsPtrType) ToStorageInsightsPtrOutput() StorageInsightsPtrOutput {
	return i.ToStorageInsightsPtrOutputWithContext(context.Background())
}

func (i *storageInsightsPtrType) ToStorageInsightsPtrOutputWithContext(ctx context.Context) StorageInsightsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageInsightsPtrOutput)
}

// StorageInsightsArrayInput is an input type that accepts StorageInsightsArray and StorageInsightsArrayOutput values.
// You can construct a concrete instance of `StorageInsightsArrayInput` via:
//
//          StorageInsightsArray{ StorageInsightsArgs{...} }
type StorageInsightsArrayInput interface {
	pulumi.Input

	ToStorageInsightsArrayOutput() StorageInsightsArrayOutput
	ToStorageInsightsArrayOutputWithContext(context.Context) StorageInsightsArrayOutput
}

type StorageInsightsArray []StorageInsightsInput

func (StorageInsightsArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*StorageInsights)(nil))
}

func (i StorageInsightsArray) ToStorageInsightsArrayOutput() StorageInsightsArrayOutput {
	return i.ToStorageInsightsArrayOutputWithContext(context.Background())
}

func (i StorageInsightsArray) ToStorageInsightsArrayOutputWithContext(ctx context.Context) StorageInsightsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageInsightsArrayOutput)
}

// StorageInsightsMapInput is an input type that accepts StorageInsightsMap and StorageInsightsMapOutput values.
// You can construct a concrete instance of `StorageInsightsMapInput` via:
//
//          StorageInsightsMap{ "key": StorageInsightsArgs{...} }
type StorageInsightsMapInput interface {
	pulumi.Input

	ToStorageInsightsMapOutput() StorageInsightsMapOutput
	ToStorageInsightsMapOutputWithContext(context.Context) StorageInsightsMapOutput
}

type StorageInsightsMap map[string]StorageInsightsInput

func (StorageInsightsMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*StorageInsights)(nil))
}

func (i StorageInsightsMap) ToStorageInsightsMapOutput() StorageInsightsMapOutput {
	return i.ToStorageInsightsMapOutputWithContext(context.Background())
}

func (i StorageInsightsMap) ToStorageInsightsMapOutputWithContext(ctx context.Context) StorageInsightsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageInsightsMapOutput)
}

type StorageInsightsOutput struct {
	*pulumi.OutputState
}

func (StorageInsightsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageInsights)(nil))
}

func (o StorageInsightsOutput) ToStorageInsightsOutput() StorageInsightsOutput {
	return o
}

func (o StorageInsightsOutput) ToStorageInsightsOutputWithContext(ctx context.Context) StorageInsightsOutput {
	return o
}

func (o StorageInsightsOutput) ToStorageInsightsPtrOutput() StorageInsightsPtrOutput {
	return o.ToStorageInsightsPtrOutputWithContext(context.Background())
}

func (o StorageInsightsOutput) ToStorageInsightsPtrOutputWithContext(ctx context.Context) StorageInsightsPtrOutput {
	return o.ApplyT(func(v StorageInsights) *StorageInsights {
		return &v
	}).(StorageInsightsPtrOutput)
}

type StorageInsightsPtrOutput struct {
	*pulumi.OutputState
}

func (StorageInsightsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageInsights)(nil))
}

func (o StorageInsightsPtrOutput) ToStorageInsightsPtrOutput() StorageInsightsPtrOutput {
	return o
}

func (o StorageInsightsPtrOutput) ToStorageInsightsPtrOutputWithContext(ctx context.Context) StorageInsightsPtrOutput {
	return o
}

type StorageInsightsArrayOutput struct{ *pulumi.OutputState }

func (StorageInsightsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageInsights)(nil))
}

func (o StorageInsightsArrayOutput) ToStorageInsightsArrayOutput() StorageInsightsArrayOutput {
	return o
}

func (o StorageInsightsArrayOutput) ToStorageInsightsArrayOutputWithContext(ctx context.Context) StorageInsightsArrayOutput {
	return o
}

func (o StorageInsightsArrayOutput) Index(i pulumi.IntInput) StorageInsightsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageInsights {
		return vs[0].([]StorageInsights)[vs[1].(int)]
	}).(StorageInsightsOutput)
}

type StorageInsightsMapOutput struct{ *pulumi.OutputState }

func (StorageInsightsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StorageInsights)(nil))
}

func (o StorageInsightsMapOutput) ToStorageInsightsMapOutput() StorageInsightsMapOutput {
	return o
}

func (o StorageInsightsMapOutput) ToStorageInsightsMapOutputWithContext(ctx context.Context) StorageInsightsMapOutput {
	return o
}

func (o StorageInsightsMapOutput) MapIndex(k pulumi.StringInput) StorageInsightsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) StorageInsights {
		return vs[0].(map[string]StorageInsights)[vs[1].(string)]
	}).(StorageInsightsOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageInsightsOutput{})
	pulumi.RegisterOutputType(StorageInsightsPtrOutput{})
	pulumi.RegisterOutputType(StorageInsightsArrayOutput{})
	pulumi.RegisterOutputType(StorageInsightsMapOutput{})
}

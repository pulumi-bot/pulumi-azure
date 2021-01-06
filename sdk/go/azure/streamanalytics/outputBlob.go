// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Stream Analytics Output to Blob Storage.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/streamanalytics"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.LookupResourceGroup(ctx, &core.LookupResourceGroupArgs{
// 			Name: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleJob, err := streamanalytics.LookupJob(ctx, &streamanalytics.LookupJobArgs{
// 			Name:              "example-job",
// 			ResourceGroupName: azurerm_resource_group.Example.Name,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      pulumi.String(exampleResourceGroup.Name),
// 			Location:               pulumi.String(exampleResourceGroup.Location),
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleContainer, err := storage.NewContainer(ctx, "exampleContainer", &storage.ContainerArgs{
// 			StorageAccountName:  exampleAccount.Name,
// 			ContainerAccessType: pulumi.String("private"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = streamanalytics.NewOutputBlob(ctx, "exampleOutputBlob", &streamanalytics.OutputBlobArgs{
// 			StreamAnalyticsJobName: pulumi.String(exampleJob.Name),
// 			ResourceGroupName:      pulumi.String(exampleJob.ResourceGroupName),
// 			StorageAccountName:     exampleAccount.Name,
// 			StorageAccountKey:      exampleAccount.PrimaryAccessKey,
// 			StorageContainerName:   exampleContainer.Name,
// 			PathPattern:            pulumi.String("some-pattern"),
// 			DateFormat:             pulumi.String("yyyy-MM-dd"),
// 			TimeFormat:             pulumi.String("HH"),
// 			Serialization: &streamanalytics.OutputBlobSerializationArgs{
// 				Type:           pulumi.String("Csv"),
// 				Encoding:       pulumi.String("UTF8"),
// 				FieldDelimiter: pulumi.String(","),
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
// Stream Analytics Outputs to Blob Storage can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:streamanalytics/outputBlob:OutputBlob example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/outputs/output1
// ```
type OutputBlob struct {
	pulumi.CustomResourceState

	// The date format. Wherever `{date}` appears in `pathPattern`, the value of this property is used as the date format instead.
	DateFormat pulumi.StringOutput `pulumi:"dateFormat"`
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
	PathPattern pulumi.StringOutput `pulumi:"pathPattern"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization OutputBlobSerializationOutput `pulumi:"serialization"`
	// The Access Key which should be used to connect to this Storage Account.
	StorageAccountKey pulumi.StringOutput `pulumi:"storageAccountKey"`
	// The name of the Storage Account.
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
	// The name of the Container within the Storage Account.
	StorageContainerName pulumi.StringOutput `pulumi:"storageContainerName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringOutput `pulumi:"streamAnalyticsJobName"`
	// The time format. Wherever `{time}` appears in `pathPattern`, the value of this property is used as the time format instead.
	TimeFormat pulumi.StringOutput `pulumi:"timeFormat"`
}

// NewOutputBlob registers a new resource with the given unique name, arguments, and options.
func NewOutputBlob(ctx *pulumi.Context,
	name string, args *OutputBlobArgs, opts ...pulumi.ResourceOption) (*OutputBlob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DateFormat == nil {
		return nil, errors.New("invalid value for required argument 'DateFormat'")
	}
	if args.PathPattern == nil {
		return nil, errors.New("invalid value for required argument 'PathPattern'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Serialization == nil {
		return nil, errors.New("invalid value for required argument 'Serialization'")
	}
	if args.StorageAccountKey == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountKey'")
	}
	if args.StorageAccountName == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountName'")
	}
	if args.StorageContainerName == nil {
		return nil, errors.New("invalid value for required argument 'StorageContainerName'")
	}
	if args.StreamAnalyticsJobName == nil {
		return nil, errors.New("invalid value for required argument 'StreamAnalyticsJobName'")
	}
	if args.TimeFormat == nil {
		return nil, errors.New("invalid value for required argument 'TimeFormat'")
	}
	var resource OutputBlob
	err := ctx.RegisterResource("azure:streamanalytics/outputBlob:OutputBlob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOutputBlob gets an existing OutputBlob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOutputBlob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OutputBlobState, opts ...pulumi.ResourceOption) (*OutputBlob, error) {
	var resource OutputBlob
	err := ctx.ReadResource("azure:streamanalytics/outputBlob:OutputBlob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OutputBlob resources.
type outputBlobState struct {
	// The date format. Wherever `{date}` appears in `pathPattern`, the value of this property is used as the date format instead.
	DateFormat *string `pulumi:"dateFormat"`
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
	PathPattern *string `pulumi:"pathPattern"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization *OutputBlobSerialization `pulumi:"serialization"`
	// The Access Key which should be used to connect to this Storage Account.
	StorageAccountKey *string `pulumi:"storageAccountKey"`
	// The name of the Storage Account.
	StorageAccountName *string `pulumi:"storageAccountName"`
	// The name of the Container within the Storage Account.
	StorageContainerName *string `pulumi:"storageContainerName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `pulumi:"streamAnalyticsJobName"`
	// The time format. Wherever `{time}` appears in `pathPattern`, the value of this property is used as the time format instead.
	TimeFormat *string `pulumi:"timeFormat"`
}

type OutputBlobState struct {
	// The date format. Wherever `{date}` appears in `pathPattern`, the value of this property is used as the date format instead.
	DateFormat pulumi.StringPtrInput
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
	PathPattern pulumi.StringPtrInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `serialization` block as defined below.
	Serialization OutputBlobSerializationPtrInput
	// The Access Key which should be used to connect to this Storage Account.
	StorageAccountKey pulumi.StringPtrInput
	// The name of the Storage Account.
	StorageAccountName pulumi.StringPtrInput
	// The name of the Container within the Storage Account.
	StorageContainerName pulumi.StringPtrInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringPtrInput
	// The time format. Wherever `{time}` appears in `pathPattern`, the value of this property is used as the time format instead.
	TimeFormat pulumi.StringPtrInput
}

func (OutputBlobState) ElementType() reflect.Type {
	return reflect.TypeOf((*outputBlobState)(nil)).Elem()
}

type outputBlobArgs struct {
	// The date format. Wherever `{date}` appears in `pathPattern`, the value of this property is used as the date format instead.
	DateFormat string `pulumi:"dateFormat"`
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
	PathPattern string `pulumi:"pathPattern"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization OutputBlobSerialization `pulumi:"serialization"`
	// The Access Key which should be used to connect to this Storage Account.
	StorageAccountKey string `pulumi:"storageAccountKey"`
	// The name of the Storage Account.
	StorageAccountName string `pulumi:"storageAccountName"`
	// The name of the Container within the Storage Account.
	StorageContainerName string `pulumi:"storageContainerName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName string `pulumi:"streamAnalyticsJobName"`
	// The time format. Wherever `{time}` appears in `pathPattern`, the value of this property is used as the time format instead.
	TimeFormat string `pulumi:"timeFormat"`
}

// The set of arguments for constructing a OutputBlob resource.
type OutputBlobArgs struct {
	// The date format. Wherever `{date}` appears in `pathPattern`, the value of this property is used as the date format instead.
	DateFormat pulumi.StringInput
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
	PathPattern pulumi.StringInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `serialization` block as defined below.
	Serialization OutputBlobSerializationInput
	// The Access Key which should be used to connect to this Storage Account.
	StorageAccountKey pulumi.StringInput
	// The name of the Storage Account.
	StorageAccountName pulumi.StringInput
	// The name of the Container within the Storage Account.
	StorageContainerName pulumi.StringInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringInput
	// The time format. Wherever `{time}` appears in `pathPattern`, the value of this property is used as the time format instead.
	TimeFormat pulumi.StringInput
}

func (OutputBlobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*outputBlobArgs)(nil)).Elem()
}

type OutputBlobInput interface {
	pulumi.Input

	ToOutputBlobOutput() OutputBlobOutput
	ToOutputBlobOutputWithContext(ctx context.Context) OutputBlobOutput
}

func (*OutputBlob) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputBlob)(nil))
}

func (i *OutputBlob) ToOutputBlobOutput() OutputBlobOutput {
	return i.ToOutputBlobOutputWithContext(context.Background())
}

func (i *OutputBlob) ToOutputBlobOutputWithContext(ctx context.Context) OutputBlobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputBlobOutput)
}

func (i *OutputBlob) ToOutputBlobPtrOutput() OutputBlobPtrOutput {
	return i.ToOutputBlobPtrOutputWithContext(context.Background())
}

func (i *OutputBlob) ToOutputBlobPtrOutputWithContext(ctx context.Context) OutputBlobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputBlobPtrOutput)
}

type OutputBlobPtrInput interface {
	pulumi.Input

	ToOutputBlobPtrOutput() OutputBlobPtrOutput
	ToOutputBlobPtrOutputWithContext(ctx context.Context) OutputBlobPtrOutput
}

type outputBlobPtrType OutputBlobArgs

func (*outputBlobPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OutputBlob)(nil))
}

func (i *outputBlobPtrType) ToOutputBlobPtrOutput() OutputBlobPtrOutput {
	return i.ToOutputBlobPtrOutputWithContext(context.Background())
}

func (i *outputBlobPtrType) ToOutputBlobPtrOutputWithContext(ctx context.Context) OutputBlobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputBlobOutput).ToOutputBlobPtrOutput()
}

// OutputBlobArrayInput is an input type that accepts OutputBlobArray and OutputBlobArrayOutput values.
// You can construct a concrete instance of `OutputBlobArrayInput` via:
//
//          OutputBlobArray{ OutputBlobArgs{...} }
type OutputBlobArrayInput interface {
	pulumi.Input

	ToOutputBlobArrayOutput() OutputBlobArrayOutput
	ToOutputBlobArrayOutputWithContext(context.Context) OutputBlobArrayOutput
}

type OutputBlobArray []OutputBlobInput

func (OutputBlobArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutputBlob)(nil))
}

func (i OutputBlobArray) ToOutputBlobArrayOutput() OutputBlobArrayOutput {
	return i.ToOutputBlobArrayOutputWithContext(context.Background())
}

func (i OutputBlobArray) ToOutputBlobArrayOutputWithContext(ctx context.Context) OutputBlobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputBlobArrayOutput)
}

// OutputBlobMapInput is an input type that accepts OutputBlobMap and OutputBlobMapOutput values.
// You can construct a concrete instance of `OutputBlobMapInput` via:
//
//          OutputBlobMap{ "key": OutputBlobArgs{...} }
type OutputBlobMapInput interface {
	pulumi.Input

	ToOutputBlobMapOutput() OutputBlobMapOutput
	ToOutputBlobMapOutputWithContext(context.Context) OutputBlobMapOutput
}

type OutputBlobMap map[string]OutputBlobInput

func (OutputBlobMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OutputBlob)(nil))
}

func (i OutputBlobMap) ToOutputBlobMapOutput() OutputBlobMapOutput {
	return i.ToOutputBlobMapOutputWithContext(context.Background())
}

func (i OutputBlobMap) ToOutputBlobMapOutputWithContext(ctx context.Context) OutputBlobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputBlobMapOutput)
}

type OutputBlobOutput struct {
	*pulumi.OutputState
}

func (OutputBlobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputBlob)(nil))
}

func (o OutputBlobOutput) ToOutputBlobOutput() OutputBlobOutput {
	return o
}

func (o OutputBlobOutput) ToOutputBlobOutputWithContext(ctx context.Context) OutputBlobOutput {
	return o
}

func (o OutputBlobOutput) ToOutputBlobPtrOutput() OutputBlobPtrOutput {
	return o.ToOutputBlobPtrOutputWithContext(context.Background())
}

func (o OutputBlobOutput) ToOutputBlobPtrOutputWithContext(ctx context.Context) OutputBlobPtrOutput {
	return o.ApplyT(func(v OutputBlob) *OutputBlob {
		return &v
	}).(OutputBlobPtrOutput)
}

type OutputBlobPtrOutput struct {
	*pulumi.OutputState
}

func (OutputBlobPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutputBlob)(nil))
}

func (o OutputBlobPtrOutput) ToOutputBlobPtrOutput() OutputBlobPtrOutput {
	return o
}

func (o OutputBlobPtrOutput) ToOutputBlobPtrOutputWithContext(ctx context.Context) OutputBlobPtrOutput {
	return o
}

type OutputBlobArrayOutput struct{ *pulumi.OutputState }

func (OutputBlobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutputBlob)(nil))
}

func (o OutputBlobArrayOutput) ToOutputBlobArrayOutput() OutputBlobArrayOutput {
	return o
}

func (o OutputBlobArrayOutput) ToOutputBlobArrayOutputWithContext(ctx context.Context) OutputBlobArrayOutput {
	return o
}

func (o OutputBlobArrayOutput) Index(i pulumi.IntInput) OutputBlobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OutputBlob {
		return vs[0].([]OutputBlob)[vs[1].(int)]
	}).(OutputBlobOutput)
}

type OutputBlobMapOutput struct{ *pulumi.OutputState }

func (OutputBlobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OutputBlob)(nil))
}

func (o OutputBlobMapOutput) ToOutputBlobMapOutput() OutputBlobMapOutput {
	return o
}

func (o OutputBlobMapOutput) ToOutputBlobMapOutputWithContext(ctx context.Context) OutputBlobMapOutput {
	return o
}

func (o OutputBlobMapOutput) MapIndex(k pulumi.StringInput) OutputBlobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) OutputBlob {
		return vs[0].(map[string]OutputBlob)[vs[1].(string)]
	}).(OutputBlobOutput)
}

func init() {
	pulumi.RegisterOutputType(OutputBlobOutput{})
	pulumi.RegisterOutputType(OutputBlobPtrOutput{})
	pulumi.RegisterOutputType(OutputBlobArrayOutput{})
	pulumi.RegisterOutputType(OutputBlobMapOutput{})
}

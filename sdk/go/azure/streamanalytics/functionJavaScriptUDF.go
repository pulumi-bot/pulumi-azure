// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a JavaScript UDF Function within Stream Analytics Streaming Job.
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
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/streamanalytics"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := core.LookupResourceGroup(ctx, &core.LookupResourceGroupArgs{
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
// 		_, err = streamanalytics.NewFunctionJavaScriptUDF(ctx, "exampleFunctionJavaScriptUDF", &streamanalytics.FunctionJavaScriptUDFArgs{
// 			StreamAnalyticsJobName: pulumi.String(exampleJob.Name),
// 			ResourceGroupName:      pulumi.String(exampleJob.ResourceGroupName),
// 			Script:                 pulumi.String(fmt.Sprintf("%v%v%v", "function getRandomNumber(in) {\n", "  return in;\n", "}\n")),
// 			Inputs: streamanalytics.FunctionJavaScriptUDFInputArray{
// 				&streamanalytics.FunctionJavaScriptUDFInputArgs{
// 					Type: pulumi.String("bigint"),
// 				},
// 			},
// 			Output: &streamanalytics.FunctionJavaScriptUDFOutputArgs{
// 				Type: pulumi.String("bigint"),
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
// Stream Analytics JavaScript UDF Functions can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:streamanalytics/functionJavaScriptUDF:FunctionJavaScriptUDF example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/outputs/output1
// ```
type FunctionJavaScriptUDF struct {
	pulumi.CustomResourceState

	// One or more `input` blocks as defined below.
	Inputs FunctionJavaScriptUDFInputArrayOutput `pulumi:"inputs"`
	// The name of the JavaScript UDF Function. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// An `output` blocks as defined below.
	Output FunctionJavaScriptUDFOutputOutput `pulumi:"output"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The JavaScript of this UDF Function.
	Script pulumi.StringOutput `pulumi:"script"`
	// The name of the Stream Analytics Job where this Function should be created. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringOutput `pulumi:"streamAnalyticsJobName"`
}

// NewFunctionJavaScriptUDF registers a new resource with the given unique name, arguments, and options.
func NewFunctionJavaScriptUDF(ctx *pulumi.Context,
	name string, args *FunctionJavaScriptUDFArgs, opts ...pulumi.ResourceOption) (*FunctionJavaScriptUDF, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Inputs == nil {
		return nil, errors.New("invalid value for required argument 'Inputs'")
	}
	if args.Output == nil {
		return nil, errors.New("invalid value for required argument 'Output'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Script == nil {
		return nil, errors.New("invalid value for required argument 'Script'")
	}
	if args.StreamAnalyticsJobName == nil {
		return nil, errors.New("invalid value for required argument 'StreamAnalyticsJobName'")
	}
	var resource FunctionJavaScriptUDF
	err := ctx.RegisterResource("azure:streamanalytics/functionJavaScriptUDF:FunctionJavaScriptUDF", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionJavaScriptUDF gets an existing FunctionJavaScriptUDF resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionJavaScriptUDF(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionJavaScriptUDFState, opts ...pulumi.ResourceOption) (*FunctionJavaScriptUDF, error) {
	var resource FunctionJavaScriptUDF
	err := ctx.ReadResource("azure:streamanalytics/functionJavaScriptUDF:FunctionJavaScriptUDF", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionJavaScriptUDF resources.
type functionJavaScriptUDFState struct {
	// One or more `input` blocks as defined below.
	Inputs []FunctionJavaScriptUDFInput `pulumi:"inputs"`
	// The name of the JavaScript UDF Function. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// An `output` blocks as defined below.
	Output *FunctionJavaScriptUDFOutput `pulumi:"output"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The JavaScript of this UDF Function.
	Script *string `pulumi:"script"`
	// The name of the Stream Analytics Job where this Function should be created. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `pulumi:"streamAnalyticsJobName"`
}

type FunctionJavaScriptUDFState struct {
	// One or more `input` blocks as defined below.
	Inputs FunctionJavaScriptUDFInputArrayInput
	// The name of the JavaScript UDF Function. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// An `output` blocks as defined below.
	Output FunctionJavaScriptUDFOutputPtrInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The JavaScript of this UDF Function.
	Script pulumi.StringPtrInput
	// The name of the Stream Analytics Job where this Function should be created. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringPtrInput
}

func (FunctionJavaScriptUDFState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionJavaScriptUDFState)(nil)).Elem()
}

type functionJavaScriptUDFArgs struct {
	// One or more `input` blocks as defined below.
	Inputs []FunctionJavaScriptUDFInput `pulumi:"inputs"`
	// The name of the JavaScript UDF Function. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// An `output` blocks as defined below.
	Output FunctionJavaScriptUDFOutput `pulumi:"output"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The JavaScript of this UDF Function.
	Script string `pulumi:"script"`
	// The name of the Stream Analytics Job where this Function should be created. Changing this forces a new resource to be created.
	StreamAnalyticsJobName string `pulumi:"streamAnalyticsJobName"`
}

// The set of arguments for constructing a FunctionJavaScriptUDF resource.
type FunctionJavaScriptUDFArgs struct {
	// One or more `input` blocks as defined below.
	Inputs FunctionJavaScriptUDFInputArrayInput
	// The name of the JavaScript UDF Function. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// An `output` blocks as defined below.
	Output FunctionJavaScriptUDFOutputInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The JavaScript of this UDF Function.
	Script pulumi.StringInput
	// The name of the Stream Analytics Job where this Function should be created. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringInput
}

func (FunctionJavaScriptUDFArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionJavaScriptUDFArgs)(nil)).Elem()
}

type FunctionJavaScriptUDFInput interface {
	pulumi.Input

	ToFunctionJavaScriptUDFOutput() FunctionJavaScriptUDFOutput
	ToFunctionJavaScriptUDFOutputWithContext(ctx context.Context) FunctionJavaScriptUDFOutput
}

func (*FunctionJavaScriptUDF) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionJavaScriptUDF)(nil))
}

func (i *FunctionJavaScriptUDF) ToFunctionJavaScriptUDFOutput() FunctionJavaScriptUDFOutput {
	return i.ToFunctionJavaScriptUDFOutputWithContext(context.Background())
}

func (i *FunctionJavaScriptUDF) ToFunctionJavaScriptUDFOutputWithContext(ctx context.Context) FunctionJavaScriptUDFOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionJavaScriptUDFOutput)
}

func (i *FunctionJavaScriptUDF) ToFunctionJavaScriptUDFPtrOutput() FunctionJavaScriptUDFPtrOutput {
	return i.ToFunctionJavaScriptUDFPtrOutputWithContext(context.Background())
}

func (i *FunctionJavaScriptUDF) ToFunctionJavaScriptUDFPtrOutputWithContext(ctx context.Context) FunctionJavaScriptUDFPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionJavaScriptUDFPtrOutput)
}

type FunctionJavaScriptUDFPtrInput interface {
	pulumi.Input

	ToFunctionJavaScriptUDFPtrOutput() FunctionJavaScriptUDFPtrOutput
	ToFunctionJavaScriptUDFPtrOutputWithContext(ctx context.Context) FunctionJavaScriptUDFPtrOutput
}

type functionJavaScriptUDFPtrType FunctionJavaScriptUDFArgs

func (*functionJavaScriptUDFPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionJavaScriptUDF)(nil))
}

func (i *functionJavaScriptUDFPtrType) ToFunctionJavaScriptUDFPtrOutput() FunctionJavaScriptUDFPtrOutput {
	return i.ToFunctionJavaScriptUDFPtrOutputWithContext(context.Background())
}

func (i *functionJavaScriptUDFPtrType) ToFunctionJavaScriptUDFPtrOutputWithContext(ctx context.Context) FunctionJavaScriptUDFPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionJavaScriptUDFOutput).ToFunctionJavaScriptUDFPtrOutput()
}

// FunctionJavaScriptUDFArrayInput is an input type that accepts FunctionJavaScriptUDFArray and FunctionJavaScriptUDFArrayOutput values.
// You can construct a concrete instance of `FunctionJavaScriptUDFArrayInput` via:
//
//          FunctionJavaScriptUDFArray{ FunctionJavaScriptUDFArgs{...} }
type FunctionJavaScriptUDFArrayInput interface {
	pulumi.Input

	ToFunctionJavaScriptUDFArrayOutput() FunctionJavaScriptUDFArrayOutput
	ToFunctionJavaScriptUDFArrayOutputWithContext(context.Context) FunctionJavaScriptUDFArrayOutput
}

type FunctionJavaScriptUDFArray []FunctionJavaScriptUDFInput

func (FunctionJavaScriptUDFArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionJavaScriptUDF)(nil))
}

func (i FunctionJavaScriptUDFArray) ToFunctionJavaScriptUDFArrayOutput() FunctionJavaScriptUDFArrayOutput {
	return i.ToFunctionJavaScriptUDFArrayOutputWithContext(context.Background())
}

func (i FunctionJavaScriptUDFArray) ToFunctionJavaScriptUDFArrayOutputWithContext(ctx context.Context) FunctionJavaScriptUDFArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionJavaScriptUDFArrayOutput)
}

// FunctionJavaScriptUDFMapInput is an input type that accepts FunctionJavaScriptUDFMap and FunctionJavaScriptUDFMapOutput values.
// You can construct a concrete instance of `FunctionJavaScriptUDFMapInput` via:
//
//          FunctionJavaScriptUDFMap{ "key": FunctionJavaScriptUDFArgs{...} }
type FunctionJavaScriptUDFMapInput interface {
	pulumi.Input

	ToFunctionJavaScriptUDFMapOutput() FunctionJavaScriptUDFMapOutput
	ToFunctionJavaScriptUDFMapOutputWithContext(context.Context) FunctionJavaScriptUDFMapOutput
}

type FunctionJavaScriptUDFMap map[string]FunctionJavaScriptUDFInput

func (FunctionJavaScriptUDFMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FunctionJavaScriptUDF)(nil))
}

func (i FunctionJavaScriptUDFMap) ToFunctionJavaScriptUDFMapOutput() FunctionJavaScriptUDFMapOutput {
	return i.ToFunctionJavaScriptUDFMapOutputWithContext(context.Background())
}

func (i FunctionJavaScriptUDFMap) ToFunctionJavaScriptUDFMapOutputWithContext(ctx context.Context) FunctionJavaScriptUDFMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionJavaScriptUDFMapOutput)
}

type FunctionJavaScriptUDFOutput struct {
	*pulumi.OutputState
}

func (FunctionJavaScriptUDFOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionJavaScriptUDF)(nil))
}

func (o FunctionJavaScriptUDFOutput) ToFunctionJavaScriptUDFOutput() FunctionJavaScriptUDFOutput {
	return o
}

func (o FunctionJavaScriptUDFOutput) ToFunctionJavaScriptUDFOutputWithContext(ctx context.Context) FunctionJavaScriptUDFOutput {
	return o
}

func (o FunctionJavaScriptUDFOutput) ToFunctionJavaScriptUDFPtrOutput() FunctionJavaScriptUDFPtrOutput {
	return o.ToFunctionJavaScriptUDFPtrOutputWithContext(context.Background())
}

func (o FunctionJavaScriptUDFOutput) ToFunctionJavaScriptUDFPtrOutputWithContext(ctx context.Context) FunctionJavaScriptUDFPtrOutput {
	return o.ApplyT(func(v FunctionJavaScriptUDF) *FunctionJavaScriptUDF {
		return &v
	}).(FunctionJavaScriptUDFPtrOutput)
}

type FunctionJavaScriptUDFPtrOutput struct {
	*pulumi.OutputState
}

func (FunctionJavaScriptUDFPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionJavaScriptUDF)(nil))
}

func (o FunctionJavaScriptUDFPtrOutput) ToFunctionJavaScriptUDFPtrOutput() FunctionJavaScriptUDFPtrOutput {
	return o
}

func (o FunctionJavaScriptUDFPtrOutput) ToFunctionJavaScriptUDFPtrOutputWithContext(ctx context.Context) FunctionJavaScriptUDFPtrOutput {
	return o
}

type FunctionJavaScriptUDFArrayOutput struct{ *pulumi.OutputState }

func (FunctionJavaScriptUDFArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionJavaScriptUDF)(nil))
}

func (o FunctionJavaScriptUDFArrayOutput) ToFunctionJavaScriptUDFArrayOutput() FunctionJavaScriptUDFArrayOutput {
	return o
}

func (o FunctionJavaScriptUDFArrayOutput) ToFunctionJavaScriptUDFArrayOutputWithContext(ctx context.Context) FunctionJavaScriptUDFArrayOutput {
	return o
}

func (o FunctionJavaScriptUDFArrayOutput) Index(i pulumi.IntInput) FunctionJavaScriptUDFOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FunctionJavaScriptUDF {
		return vs[0].([]FunctionJavaScriptUDF)[vs[1].(int)]
	}).(FunctionJavaScriptUDFOutput)
}

type FunctionJavaScriptUDFMapOutput struct{ *pulumi.OutputState }

func (FunctionJavaScriptUDFMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FunctionJavaScriptUDF)(nil))
}

func (o FunctionJavaScriptUDFMapOutput) ToFunctionJavaScriptUDFMapOutput() FunctionJavaScriptUDFMapOutput {
	return o
}

func (o FunctionJavaScriptUDFMapOutput) ToFunctionJavaScriptUDFMapOutputWithContext(ctx context.Context) FunctionJavaScriptUDFMapOutput {
	return o
}

func (o FunctionJavaScriptUDFMapOutput) MapIndex(k pulumi.StringInput) FunctionJavaScriptUDFOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FunctionJavaScriptUDF {
		return vs[0].(map[string]FunctionJavaScriptUDF)[vs[1].(string)]
	}).(FunctionJavaScriptUDFOutput)
}

func init() {
	pulumi.RegisterOutputType(FunctionJavaScriptUDFOutput{})
	pulumi.RegisterOutputType(FunctionJavaScriptUDFPtrOutput{})
	pulumi.RegisterOutputType(FunctionJavaScriptUDFArrayOutput{})
	pulumi.RegisterOutputType(FunctionJavaScriptUDFMapOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Pipeline inside a Azure Data Factory.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/datafactory"
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
// 		exampleFactory, err := datafactory.NewFactory(ctx, "exampleFactory", &datafactory.FactoryArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewPipeline(ctx, "examplePipeline", &datafactory.PipelineArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryName:   exampleFactory.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### With Activities
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/datafactory"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := datafactory.NewPipeline(ctx, "test", &datafactory.PipelineArgs{
// 			ResourceGroupName: pulumi.Any(azurerm_resource_group.Test.Name),
// 			DataFactoryName:   pulumi.Any(azurerm_data_factory.Test.Name),
// 			Variables: pulumi.StringMap{
// 				"bob": pulumi.String("item1"),
// 			},
// 			ActivitiesJson: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v", "[\n", "	{\n", "		\"name\": \"Append variable1\",\n", "		\"type\": \"AppendVariable\",\n", "		\"dependsOn\": [],\n", "		\"userProperties\": [],\n", "		\"typeProperties\": {\n", "			\"variableName\": \"bob\",\n", "			\"value\": \"something\"\n", "		}\n", "	}\n", "]\n")),
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
// Data Factory Pipeline's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/pipeline:Pipeline example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/pipelines/example
// ```
type Pipeline struct {
	pulumi.CustomResourceState

	// A JSON object that contains the activities that will be associated with the Data Factory Pipeline.
	ActivitiesJson pulumi.StringPtrOutput `pulumi:"activitiesJson"`
	// List of tags that can be used for describing the Data Factory Pipeline.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The Data Factory name in which to associate the Pipeline with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Pipeline.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the name of the Data Factory Pipeline. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Pipeline.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Pipeline. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A map of variables to associate with the Data Factory Pipeline.
	Variables pulumi.StringMapOutput `pulumi:"variables"`
}

// NewPipeline registers a new resource with the given unique name, arguments, and options.
func NewPipeline(ctx *pulumi.Context,
	name string, args *PipelineArgs, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataFactoryName == nil {
		return nil, errors.New("invalid value for required argument 'DataFactoryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Pipeline
	err := ctx.RegisterResource("azure:datafactory/pipeline:Pipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipeline gets an existing Pipeline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineState, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	var resource Pipeline
	err := ctx.ReadResource("azure:datafactory/pipeline:Pipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pipeline resources.
type pipelineState struct {
	// A JSON object that contains the activities that will be associated with the Data Factory Pipeline.
	ActivitiesJson *string `pulumi:"activitiesJson"`
	// List of tags that can be used for describing the Data Factory Pipeline.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory name in which to associate the Pipeline with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Pipeline.
	Description *string `pulumi:"description"`
	// Specifies the name of the Data Factory Pipeline. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Pipeline.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Pipeline. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A map of variables to associate with the Data Factory Pipeline.
	Variables map[string]string `pulumi:"variables"`
}

type PipelineState struct {
	// A JSON object that contains the activities that will be associated with the Data Factory Pipeline.
	ActivitiesJson pulumi.StringPtrInput
	// List of tags that can be used for describing the Data Factory Pipeline.
	Annotations pulumi.StringArrayInput
	// The Data Factory name in which to associate the Pipeline with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Pipeline.
	Description pulumi.StringPtrInput
	// Specifies the name of the Data Factory Pipeline. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Pipeline.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Pipeline. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
	// A map of variables to associate with the Data Factory Pipeline.
	Variables pulumi.StringMapInput
}

func (PipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineState)(nil)).Elem()
}

type pipelineArgs struct {
	// A JSON object that contains the activities that will be associated with the Data Factory Pipeline.
	ActivitiesJson *string `pulumi:"activitiesJson"`
	// List of tags that can be used for describing the Data Factory Pipeline.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory name in which to associate the Pipeline with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Pipeline.
	Description *string `pulumi:"description"`
	// Specifies the name of the Data Factory Pipeline. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Pipeline.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Pipeline. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A map of variables to associate with the Data Factory Pipeline.
	Variables map[string]string `pulumi:"variables"`
}

// The set of arguments for constructing a Pipeline resource.
type PipelineArgs struct {
	// A JSON object that contains the activities that will be associated with the Data Factory Pipeline.
	ActivitiesJson pulumi.StringPtrInput
	// List of tags that can be used for describing the Data Factory Pipeline.
	Annotations pulumi.StringArrayInput
	// The Data Factory name in which to associate the Pipeline with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The description for the Data Factory Pipeline.
	Description pulumi.StringPtrInput
	// Specifies the name of the Data Factory Pipeline. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Pipeline.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Pipeline. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
	// A map of variables to associate with the Data Factory Pipeline.
	Variables pulumi.StringMapInput
}

func (PipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineArgs)(nil)).Elem()
}

type PipelineInput interface {
	pulumi.Input

	ToPipelineOutput() PipelineOutput
	ToPipelineOutputWithContext(ctx context.Context) PipelineOutput
}

func (*Pipeline) ElementType() reflect.Type {
	return reflect.TypeOf((*Pipeline)(nil))
}

func (i *Pipeline) ToPipelineOutput() PipelineOutput {
	return i.ToPipelineOutputWithContext(context.Background())
}

func (i *Pipeline) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineOutput)
}

func (i *Pipeline) ToPipelinePtrOutput() PipelinePtrOutput {
	return i.ToPipelinePtrOutputWithContext(context.Background())
}

func (i *Pipeline) ToPipelinePtrOutputWithContext(ctx context.Context) PipelinePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelinePtrOutput)
}

type PipelinePtrInput interface {
	pulumi.Input

	ToPipelinePtrOutput() PipelinePtrOutput
	ToPipelinePtrOutputWithContext(ctx context.Context) PipelinePtrOutput
}

type pipelinePtrType PipelineArgs

func (*pipelinePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Pipeline)(nil))
}

func (i *pipelinePtrType) ToPipelinePtrOutput() PipelinePtrOutput {
	return i.ToPipelinePtrOutputWithContext(context.Background())
}

func (i *pipelinePtrType) ToPipelinePtrOutputWithContext(ctx context.Context) PipelinePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineOutput).ToPipelinePtrOutput()
}

// PipelineArrayInput is an input type that accepts PipelineArray and PipelineArrayOutput values.
// You can construct a concrete instance of `PipelineArrayInput` via:
//
//          PipelineArray{ PipelineArgs{...} }
type PipelineArrayInput interface {
	pulumi.Input

	ToPipelineArrayOutput() PipelineArrayOutput
	ToPipelineArrayOutputWithContext(context.Context) PipelineArrayOutput
}

type PipelineArray []PipelineInput

func (PipelineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Pipeline)(nil))
}

func (i PipelineArray) ToPipelineArrayOutput() PipelineArrayOutput {
	return i.ToPipelineArrayOutputWithContext(context.Background())
}

func (i PipelineArray) ToPipelineArrayOutputWithContext(ctx context.Context) PipelineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineArrayOutput)
}

// PipelineMapInput is an input type that accepts PipelineMap and PipelineMapOutput values.
// You can construct a concrete instance of `PipelineMapInput` via:
//
//          PipelineMap{ "key": PipelineArgs{...} }
type PipelineMapInput interface {
	pulumi.Input

	ToPipelineMapOutput() PipelineMapOutput
	ToPipelineMapOutputWithContext(context.Context) PipelineMapOutput
}

type PipelineMap map[string]PipelineInput

func (PipelineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Pipeline)(nil))
}

func (i PipelineMap) ToPipelineMapOutput() PipelineMapOutput {
	return i.ToPipelineMapOutputWithContext(context.Background())
}

func (i PipelineMap) ToPipelineMapOutputWithContext(ctx context.Context) PipelineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineMapOutput)
}

type PipelineOutput struct {
	*pulumi.OutputState
}

func (PipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Pipeline)(nil))
}

func (o PipelineOutput) ToPipelineOutput() PipelineOutput {
	return o
}

func (o PipelineOutput) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return o
}

func (o PipelineOutput) ToPipelinePtrOutput() PipelinePtrOutput {
	return o.ToPipelinePtrOutputWithContext(context.Background())
}

func (o PipelineOutput) ToPipelinePtrOutputWithContext(ctx context.Context) PipelinePtrOutput {
	return o.ApplyT(func(v Pipeline) *Pipeline {
		return &v
	}).(PipelinePtrOutput)
}

type PipelinePtrOutput struct {
	*pulumi.OutputState
}

func (PipelinePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pipeline)(nil))
}

func (o PipelinePtrOutput) ToPipelinePtrOutput() PipelinePtrOutput {
	return o
}

func (o PipelinePtrOutput) ToPipelinePtrOutputWithContext(ctx context.Context) PipelinePtrOutput {
	return o
}

type PipelineArrayOutput struct{ *pulumi.OutputState }

func (PipelineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Pipeline)(nil))
}

func (o PipelineArrayOutput) ToPipelineArrayOutput() PipelineArrayOutput {
	return o
}

func (o PipelineArrayOutput) ToPipelineArrayOutputWithContext(ctx context.Context) PipelineArrayOutput {
	return o
}

func (o PipelineArrayOutput) Index(i pulumi.IntInput) PipelineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Pipeline {
		return vs[0].([]Pipeline)[vs[1].(int)]
	}).(PipelineOutput)
}

type PipelineMapOutput struct{ *pulumi.OutputState }

func (PipelineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Pipeline)(nil))
}

func (o PipelineMapOutput) ToPipelineMapOutput() PipelineMapOutput {
	return o
}

func (o PipelineMapOutput) ToPipelineMapOutputWithContext(ctx context.Context) PipelineMapOutput {
	return o
}

func (o PipelineMapOutput) MapIndex(k pulumi.StringInput) PipelineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Pipeline {
		return vs[0].(map[string]Pipeline)[vs[1].(string)]
	}).(PipelineOutput)
}

func init() {
	pulumi.RegisterOutputType(PipelineOutput{})
	pulumi.RegisterOutputType(PipelinePtrOutput{})
	pulumi.RegisterOutputType(PipelineArrayOutput{})
	pulumi.RegisterOutputType(PipelineMapOutput{})
}

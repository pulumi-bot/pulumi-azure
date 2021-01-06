// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logicapps

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Custom Trigger within a Logic App Workflow
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
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/logicapps"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("East US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleWorkflow, err := logicapps.NewWorkflow(ctx, "exampleWorkflow", &logicapps.WorkflowArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = logicapps.NewTriggerCustom(ctx, "exampleTriggerCustom", &logicapps.TriggerCustomArgs{
// 			LogicAppId: exampleWorkflow.ID(),
// 			Body:       pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v", "{\n", "  \"recurrence\": {\n", "    \"frequency\": \"Day\",\n", "    \"interval\": 1\n", "  },\n", "  \"type\": \"Recurrence\"\n", "}\n")),
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
// Logic App Custom Triggers can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:logicapps/triggerCustom:TriggerCustom custom1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1/triggers/custom1
// ```
type TriggerCustom struct {
	pulumi.CustomResourceState

	// Specifies the JSON Blob defining the Body of this Custom Trigger.
	Body pulumi.StringOutput `pulumi:"body"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringOutput `pulumi:"logicAppId"`
	// Specifies the name of the HTTP Trigger to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewTriggerCustom registers a new resource with the given unique name, arguments, and options.
func NewTriggerCustom(ctx *pulumi.Context,
	name string, args *TriggerCustomArgs, opts ...pulumi.ResourceOption) (*TriggerCustom, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Body == nil {
		return nil, errors.New("invalid value for required argument 'Body'")
	}
	if args.LogicAppId == nil {
		return nil, errors.New("invalid value for required argument 'LogicAppId'")
	}
	var resource TriggerCustom
	err := ctx.RegisterResource("azure:logicapps/triggerCustom:TriggerCustom", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTriggerCustom gets an existing TriggerCustom resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTriggerCustom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerCustomState, opts ...pulumi.ResourceOption) (*TriggerCustom, error) {
	var resource TriggerCustom
	err := ctx.ReadResource("azure:logicapps/triggerCustom:TriggerCustom", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TriggerCustom resources.
type triggerCustomState struct {
	// Specifies the JSON Blob defining the Body of this Custom Trigger.
	Body *string `pulumi:"body"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId *string `pulumi:"logicAppId"`
	// Specifies the name of the HTTP Trigger to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
}

type TriggerCustomState struct {
	// Specifies the JSON Blob defining the Body of this Custom Trigger.
	Body pulumi.StringPtrInput
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringPtrInput
	// Specifies the name of the HTTP Trigger to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
}

func (TriggerCustomState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerCustomState)(nil)).Elem()
}

type triggerCustomArgs struct {
	// Specifies the JSON Blob defining the Body of this Custom Trigger.
	Body string `pulumi:"body"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId string `pulumi:"logicAppId"`
	// Specifies the name of the HTTP Trigger to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a TriggerCustom resource.
type TriggerCustomArgs struct {
	// Specifies the JSON Blob defining the Body of this Custom Trigger.
	Body pulumi.StringInput
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringInput
	// Specifies the name of the HTTP Trigger to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
}

func (TriggerCustomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerCustomArgs)(nil)).Elem()
}

type TriggerCustomInput interface {
	pulumi.Input

	ToTriggerCustomOutput() TriggerCustomOutput
	ToTriggerCustomOutputWithContext(ctx context.Context) TriggerCustomOutput
}

func (*TriggerCustom) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerCustom)(nil))
}

func (i *TriggerCustom) ToTriggerCustomOutput() TriggerCustomOutput {
	return i.ToTriggerCustomOutputWithContext(context.Background())
}

func (i *TriggerCustom) ToTriggerCustomOutputWithContext(ctx context.Context) TriggerCustomOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerCustomOutput)
}

func (i *TriggerCustom) ToTriggerCustomPtrOutput() TriggerCustomPtrOutput {
	return i.ToTriggerCustomPtrOutputWithContext(context.Background())
}

func (i *TriggerCustom) ToTriggerCustomPtrOutputWithContext(ctx context.Context) TriggerCustomPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerCustomPtrOutput)
}

type TriggerCustomPtrInput interface {
	pulumi.Input

	ToTriggerCustomPtrOutput() TriggerCustomPtrOutput
	ToTriggerCustomPtrOutputWithContext(ctx context.Context) TriggerCustomPtrOutput
}

type triggerCustomPtrType TriggerCustomArgs

func (*triggerCustomPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerCustom)(nil))
}

func (i *triggerCustomPtrType) ToTriggerCustomPtrOutput() TriggerCustomPtrOutput {
	return i.ToTriggerCustomPtrOutputWithContext(context.Background())
}

func (i *triggerCustomPtrType) ToTriggerCustomPtrOutputWithContext(ctx context.Context) TriggerCustomPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerCustomOutput).ToTriggerCustomPtrOutput()
}

// TriggerCustomArrayInput is an input type that accepts TriggerCustomArray and TriggerCustomArrayOutput values.
// You can construct a concrete instance of `TriggerCustomArrayInput` via:
//
//          TriggerCustomArray{ TriggerCustomArgs{...} }
type TriggerCustomArrayInput interface {
	pulumi.Input

	ToTriggerCustomArrayOutput() TriggerCustomArrayOutput
	ToTriggerCustomArrayOutputWithContext(context.Context) TriggerCustomArrayOutput
}

type TriggerCustomArray []TriggerCustomInput

func (TriggerCustomArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TriggerCustom)(nil))
}

func (i TriggerCustomArray) ToTriggerCustomArrayOutput() TriggerCustomArrayOutput {
	return i.ToTriggerCustomArrayOutputWithContext(context.Background())
}

func (i TriggerCustomArray) ToTriggerCustomArrayOutputWithContext(ctx context.Context) TriggerCustomArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerCustomArrayOutput)
}

// TriggerCustomMapInput is an input type that accepts TriggerCustomMap and TriggerCustomMapOutput values.
// You can construct a concrete instance of `TriggerCustomMapInput` via:
//
//          TriggerCustomMap{ "key": TriggerCustomArgs{...} }
type TriggerCustomMapInput interface {
	pulumi.Input

	ToTriggerCustomMapOutput() TriggerCustomMapOutput
	ToTriggerCustomMapOutputWithContext(context.Context) TriggerCustomMapOutput
}

type TriggerCustomMap map[string]TriggerCustomInput

func (TriggerCustomMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TriggerCustom)(nil))
}

func (i TriggerCustomMap) ToTriggerCustomMapOutput() TriggerCustomMapOutput {
	return i.ToTriggerCustomMapOutputWithContext(context.Background())
}

func (i TriggerCustomMap) ToTriggerCustomMapOutputWithContext(ctx context.Context) TriggerCustomMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerCustomMapOutput)
}

type TriggerCustomOutput struct {
	*pulumi.OutputState
}

func (TriggerCustomOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerCustom)(nil))
}

func (o TriggerCustomOutput) ToTriggerCustomOutput() TriggerCustomOutput {
	return o
}

func (o TriggerCustomOutput) ToTriggerCustomOutputWithContext(ctx context.Context) TriggerCustomOutput {
	return o
}

func (o TriggerCustomOutput) ToTriggerCustomPtrOutput() TriggerCustomPtrOutput {
	return o.ToTriggerCustomPtrOutputWithContext(context.Background())
}

func (o TriggerCustomOutput) ToTriggerCustomPtrOutputWithContext(ctx context.Context) TriggerCustomPtrOutput {
	return o.ApplyT(func(v TriggerCustom) *TriggerCustom {
		return &v
	}).(TriggerCustomPtrOutput)
}

type TriggerCustomPtrOutput struct {
	*pulumi.OutputState
}

func (TriggerCustomPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerCustom)(nil))
}

func (o TriggerCustomPtrOutput) ToTriggerCustomPtrOutput() TriggerCustomPtrOutput {
	return o
}

func (o TriggerCustomPtrOutput) ToTriggerCustomPtrOutputWithContext(ctx context.Context) TriggerCustomPtrOutput {
	return o
}

type TriggerCustomArrayOutput struct{ *pulumi.OutputState }

func (TriggerCustomArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TriggerCustom)(nil))
}

func (o TriggerCustomArrayOutput) ToTriggerCustomArrayOutput() TriggerCustomArrayOutput {
	return o
}

func (o TriggerCustomArrayOutput) ToTriggerCustomArrayOutputWithContext(ctx context.Context) TriggerCustomArrayOutput {
	return o
}

func (o TriggerCustomArrayOutput) Index(i pulumi.IntInput) TriggerCustomOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TriggerCustom {
		return vs[0].([]TriggerCustom)[vs[1].(int)]
	}).(TriggerCustomOutput)
}

type TriggerCustomMapOutput struct{ *pulumi.OutputState }

func (TriggerCustomMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TriggerCustom)(nil))
}

func (o TriggerCustomMapOutput) ToTriggerCustomMapOutput() TriggerCustomMapOutput {
	return o
}

func (o TriggerCustomMapOutput) ToTriggerCustomMapOutputWithContext(ctx context.Context) TriggerCustomMapOutput {
	return o
}

func (o TriggerCustomMapOutput) MapIndex(k pulumi.StringInput) TriggerCustomOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TriggerCustom {
		return vs[0].(map[string]TriggerCustom)[vs[1].(string)]
	}).(TriggerCustomOutput)
}

func init() {
	pulumi.RegisterOutputType(TriggerCustomOutput{})
	pulumi.RegisterOutputType(TriggerCustomPtrOutput{})
	pulumi.RegisterOutputType(TriggerCustomArrayOutput{})
	pulumi.RegisterOutputType(TriggerCustomMapOutput{})
}

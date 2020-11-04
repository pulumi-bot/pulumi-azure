// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logicapps

import (
	"fmt"
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

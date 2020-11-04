// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logicapps

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Custom Action within a Logic App Workflow
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
// 		_, err = logicapps.NewActionCustom(ctx, "exampleActionCustom", &logicapps.ActionCustomArgs{
// 			LogicAppId: exampleWorkflow.ID(),
// 			Body:       pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"description\": \"A variable to configure the auto expiration age in days. Configured in negative number. Default is -30 (30 days old).\",\n", "    \"inputs\": {\n", "        \"variables\": [\n", "            {\n", "                \"name\": \"ExpirationAgeInDays\",\n", "                \"type\": \"Integer\",\n", "                \"value\": -30\n", "            }\n", "        ]\n", "    },\n", "    \"runAfter\": {},\n", "    \"type\": \"InitializeVariable\"\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ActionCustom struct {
	pulumi.CustomResourceState

	// Specifies the JSON Blob defining the Body of this Custom Action.
	Body pulumi.StringOutput `pulumi:"body"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringOutput `pulumi:"logicAppId"`
	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewActionCustom registers a new resource with the given unique name, arguments, and options.
func NewActionCustom(ctx *pulumi.Context,
	name string, args *ActionCustomArgs, opts ...pulumi.ResourceOption) (*ActionCustom, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Body == nil {
		return nil, errors.New("invalid value for required argument 'Body'")
	}
	if args.LogicAppId == nil {
		return nil, errors.New("invalid value for required argument 'LogicAppId'")
	}
	var resource ActionCustom
	err := ctx.RegisterResource("azure:logicapps/actionCustom:ActionCustom", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionCustom gets an existing ActionCustom resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionCustom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionCustomState, opts ...pulumi.ResourceOption) (*ActionCustom, error) {
	var resource ActionCustom
	err := ctx.ReadResource("azure:logicapps/actionCustom:ActionCustom", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionCustom resources.
type actionCustomState struct {
	// Specifies the JSON Blob defining the Body of this Custom Action.
	Body *string `pulumi:"body"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId *string `pulumi:"logicAppId"`
	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
}

type ActionCustomState struct {
	// Specifies the JSON Blob defining the Body of this Custom Action.
	Body pulumi.StringPtrInput
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringPtrInput
	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
}

func (ActionCustomState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionCustomState)(nil)).Elem()
}

type actionCustomArgs struct {
	// Specifies the JSON Blob defining the Body of this Custom Action.
	Body string `pulumi:"body"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId string `pulumi:"logicAppId"`
	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ActionCustom resource.
type ActionCustomArgs struct {
	// Specifies the JSON Blob defining the Body of this Custom Action.
	Body pulumi.StringInput
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringInput
	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
}

func (ActionCustomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionCustomArgs)(nil)).Elem()
}

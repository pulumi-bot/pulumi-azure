// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an App Service Plan component.
//
// ## Example Usage
// ### Dedicated)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/appservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 		_, err = appservice.NewPlan(ctx, "examplePlan", &appservice.PlanArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku: &appservice.PlanSkuArgs{
// 				Tier: pulumi.String("Standard"),
// 				Size: pulumi.String("S1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Shared / Consumption Plan)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/appservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 		_, err = appservice.NewPlan(ctx, "examplePlan", &appservice.PlanArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Kind:              pulumi.String("FunctionApp"),
// 			Sku: &appservice.PlanSkuArgs{
// 				Tier: pulumi.String("Dynamic"),
// 				Size: pulumi.String("Y1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Linux)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/appservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 		_, err = appservice.NewPlan(ctx, "examplePlan", &appservice.PlanArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Kind:              pulumi.String("Linux"),
// 			Reserved:          pulumi.Bool(true),
// 			Sku: &appservice.PlanSkuArgs{
// 				Tier: pulumi.String("Standard"),
// 				Size: pulumi.String("S1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Windows Container)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/appservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 		_, err = appservice.NewPlan(ctx, "examplePlan", &appservice.PlanArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Kind:              pulumi.String("xenon"),
// 			IsXenon:           pulumi.Bool(true),
// 			Sku: &appservice.PlanSkuArgs{
// 				Tier: pulumi.String("PremiumContainer"),
// 				Size: pulumi.String("PC2"),
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
// App Service Plan instances can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:appservice/plan:Plan instance1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/serverfarms/instance1
// ```
type Plan struct {
	pulumi.CustomResourceState

	// The ID of the App Service Environment where the App Service Plan should be located. Changing forces a new resource to be created.
	AppServiceEnvironmentId pulumi.StringPtrOutput `pulumi:"appServiceEnvironmentId"`
	IsXenon                 pulumi.BoolPtrOutput   `pulumi:"isXenon"`
	// The kind of the App Service Plan to create. Possible values are `Windows` (also available as `App`), `Linux`, `elastic` (for Premium Consumption) and `FunctionApp` (for a Consumption Plan). Defaults to `Windows`. Changing this forces a new resource to be created.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan.
	MaximumElasticWorkerCount pulumi.IntOutput `pulumi:"maximumElasticWorkerCount"`
	// The maximum number of workers supported with the App Service Plan's sku.
	MaximumNumberOfWorkers pulumi.IntOutput `pulumi:"maximumNumberOfWorkers"`
	// Specifies the name of the App Service Plan component. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Can Apps assigned to this App Service Plan be scaled independently? If set to `false` apps assigned to this plan will scale to all instances of the plan.  Defaults to `false`.
	PerSiteScaling pulumi.BoolPtrOutput `pulumi:"perSiteScaling"`
	// Is this App Service Plan `Reserved`. Defaults to `false`.
	Reserved pulumi.BoolPtrOutput `pulumi:"reserved"`
	// The name of the resource group in which to create the App Service Plan component.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `sku` block as documented below.
	Sku PlanSkuOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewPlan registers a new resource with the given unique name, arguments, and options.
func NewPlan(ctx *pulumi.Context,
	name string, args *PlanArgs, opts ...pulumi.ResourceOption) (*Plan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	var resource Plan
	err := ctx.RegisterResource("azure:appservice/plan:Plan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPlan gets an existing Plan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PlanState, opts ...pulumi.ResourceOption) (*Plan, error) {
	var resource Plan
	err := ctx.ReadResource("azure:appservice/plan:Plan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Plan resources.
type planState struct {
	// The ID of the App Service Environment where the App Service Plan should be located. Changing forces a new resource to be created.
	AppServiceEnvironmentId *string `pulumi:"appServiceEnvironmentId"`
	IsXenon                 *bool   `pulumi:"isXenon"`
	// The kind of the App Service Plan to create. Possible values are `Windows` (also available as `App`), `Linux`, `elastic` (for Premium Consumption) and `FunctionApp` (for a Consumption Plan). Defaults to `Windows`. Changing this forces a new resource to be created.
	Kind *string `pulumi:"kind"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan.
	MaximumElasticWorkerCount *int `pulumi:"maximumElasticWorkerCount"`
	// The maximum number of workers supported with the App Service Plan's sku.
	MaximumNumberOfWorkers *int `pulumi:"maximumNumberOfWorkers"`
	// Specifies the name of the App Service Plan component. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Can Apps assigned to this App Service Plan be scaled independently? If set to `false` apps assigned to this plan will scale to all instances of the plan.  Defaults to `false`.
	PerSiteScaling *bool `pulumi:"perSiteScaling"`
	// Is this App Service Plan `Reserved`. Defaults to `false`.
	Reserved *bool `pulumi:"reserved"`
	// The name of the resource group in which to create the App Service Plan component.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `sku` block as documented below.
	Sku *PlanSku `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type PlanState struct {
	// The ID of the App Service Environment where the App Service Plan should be located. Changing forces a new resource to be created.
	AppServiceEnvironmentId pulumi.StringPtrInput
	IsXenon                 pulumi.BoolPtrInput
	// The kind of the App Service Plan to create. Possible values are `Windows` (also available as `App`), `Linux`, `elastic` (for Premium Consumption) and `FunctionApp` (for a Consumption Plan). Defaults to `Windows`. Changing this forces a new resource to be created.
	Kind pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan.
	MaximumElasticWorkerCount pulumi.IntPtrInput
	// The maximum number of workers supported with the App Service Plan's sku.
	MaximumNumberOfWorkers pulumi.IntPtrInput
	// Specifies the name of the App Service Plan component. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Can Apps assigned to this App Service Plan be scaled independently? If set to `false` apps assigned to this plan will scale to all instances of the plan.  Defaults to `false`.
	PerSiteScaling pulumi.BoolPtrInput
	// Is this App Service Plan `Reserved`. Defaults to `false`.
	Reserved pulumi.BoolPtrInput
	// The name of the resource group in which to create the App Service Plan component.
	ResourceGroupName pulumi.StringPtrInput
	// A `sku` block as documented below.
	Sku PlanSkuPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (PlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*planState)(nil)).Elem()
}

type planArgs struct {
	// The ID of the App Service Environment where the App Service Plan should be located. Changing forces a new resource to be created.
	AppServiceEnvironmentId *string `pulumi:"appServiceEnvironmentId"`
	IsXenon                 *bool   `pulumi:"isXenon"`
	// The kind of the App Service Plan to create. Possible values are `Windows` (also available as `App`), `Linux`, `elastic` (for Premium Consumption) and `FunctionApp` (for a Consumption Plan). Defaults to `Windows`. Changing this forces a new resource to be created.
	Kind interface{} `pulumi:"kind"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan.
	MaximumElasticWorkerCount *int `pulumi:"maximumElasticWorkerCount"`
	// Specifies the name of the App Service Plan component. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Can Apps assigned to this App Service Plan be scaled independently? If set to `false` apps assigned to this plan will scale to all instances of the plan.  Defaults to `false`.
	PerSiteScaling *bool `pulumi:"perSiteScaling"`
	// Is this App Service Plan `Reserved`. Defaults to `false`.
	Reserved *bool `pulumi:"reserved"`
	// The name of the resource group in which to create the App Service Plan component.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `sku` block as documented below.
	Sku PlanSku `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Plan resource.
type PlanArgs struct {
	// The ID of the App Service Environment where the App Service Plan should be located. Changing forces a new resource to be created.
	AppServiceEnvironmentId pulumi.StringPtrInput
	IsXenon                 pulumi.BoolPtrInput
	// The kind of the App Service Plan to create. Possible values are `Windows` (also available as `App`), `Linux`, `elastic` (for Premium Consumption) and `FunctionApp` (for a Consumption Plan). Defaults to `Windows`. Changing this forces a new resource to be created.
	Kind pulumi.Input
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan.
	MaximumElasticWorkerCount pulumi.IntPtrInput
	// Specifies the name of the App Service Plan component. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Can Apps assigned to this App Service Plan be scaled independently? If set to `false` apps assigned to this plan will scale to all instances of the plan.  Defaults to `false`.
	PerSiteScaling pulumi.BoolPtrInput
	// Is this App Service Plan `Reserved`. Defaults to `false`.
	Reserved pulumi.BoolPtrInput
	// The name of the resource group in which to create the App Service Plan component.
	ResourceGroupName pulumi.StringInput
	// A `sku` block as documented below.
	Sku PlanSkuInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (PlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*planArgs)(nil)).Elem()
}

type PlanInput interface {
	pulumi.Input

	ToPlanOutput() PlanOutput
	ToPlanOutputWithContext(ctx context.Context) PlanOutput
}

func (Plan) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (i Plan) ToPlanOutput() PlanOutput {
	return i.ToPlanOutputWithContext(context.Background())
}

func (i Plan) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput)
}

type PlanOutput struct {
	*pulumi.OutputState
}

func (PlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanOutput)(nil)).Elem()
}

func (o PlanOutput) ToPlanOutput() PlanOutput {
	return o
}

func (o PlanOutput) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PlanOutput{})
}

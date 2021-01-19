// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure Data Factory Managed Integration Runtime.
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
// 		_, err = datafactory.NewIntegrationRuntimeManaged(ctx, "exampleIntegrationRuntimeManaged", &datafactory.IntegrationRuntimeManagedArgs{
// 			DataFactoryName:   exampleFactory.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			NodeSize:          pulumi.String("Standard_D8_v3"),
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
// Data Factory Integration Managed Runtimes can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/integrationRuntimeManaged:IntegrationRuntimeManaged example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/integrationruntimes/example
// ```
type IntegrationRuntimeManaged struct {
	pulumi.CustomResourceState

	// A `catalogInfo` block as defined below.
	CatalogInfo IntegrationRuntimeManagedCatalogInfoPtrOutput `pulumi:"catalogInfo"`
	// A `customSetupScript` block as defined below.
	CustomSetupScript IntegrationRuntimeManagedCustomSetupScriptPtrOutput `pulumi:"customSetupScript"`
	// Specifies the name of the Data Factory the Managed Integration Runtime belongs to. Changing this forces a new resource to be created.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// Integration runtime description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Managed Integration Runtime edition. Valid values are `Standard` and `Enterprise`. Defaults to `Standard`.
	Edition pulumi.StringPtrOutput `pulumi:"edition"`
	// The type of the license that is used. Valid values are `LicenseIncluded` and `BasePrize`. Defaults to `LicenseIncluded`.
	LicenseType pulumi.StringPtrOutput `pulumi:"licenseType"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Defines the maximum parallel executions per node. Defaults to `1`. Max is `16`.
	MaxParallelExecutionsPerNode pulumi.IntPtrOutput `pulumi:"maxParallelExecutionsPerNode"`
	// Specifies the name of the Managed Integration Runtime. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// The size of the nodes on which the Managed Integration Runtime runs. Valid values are: `Standard_D2_v3`, `Standard_D4_v3`, `Standard_D8_v3`, `Standard_D16_v3`, `Standard_D32_v3`, `Standard_D64_v3`, `Standard_E2_v3`, `Standard_E4_v3`, `Standard_E8_v3`, `Standard_E16_v3`, `Standard_E32_v3`, `Standard_E64_v3`, `Standard_D1_v2`, `Standard_D2_v2`, `Standard_D3_v2`, `Standard_D4_v2`, `Standard_A4_v2` and `Standard_A8_v2`
	NodeSize pulumi.StringOutput `pulumi:"nodeSize"`
	// Number of nodes for the Managed Integration Runtime. Max is `10`. Defaults to `1`.
	NumberOfNodes pulumi.IntPtrOutput `pulumi:"numberOfNodes"`
	// The name of the resource group in which to create the Managed Integration Runtime. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `vnetIntegration` block as defined below.
	VnetIntegration IntegrationRuntimeManagedVnetIntegrationPtrOutput `pulumi:"vnetIntegration"`
}

// NewIntegrationRuntimeManaged registers a new resource with the given unique name, arguments, and options.
func NewIntegrationRuntimeManaged(ctx *pulumi.Context,
	name string, args *IntegrationRuntimeManagedArgs, opts ...pulumi.ResourceOption) (*IntegrationRuntimeManaged, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataFactoryName == nil {
		return nil, errors.New("invalid value for required argument 'DataFactoryName'")
	}
	if args.NodeSize == nil {
		return nil, errors.New("invalid value for required argument 'NodeSize'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource IntegrationRuntimeManaged
	err := ctx.RegisterResource("azure:datafactory/integrationRuntimeManaged:IntegrationRuntimeManaged", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationRuntimeManaged gets an existing IntegrationRuntimeManaged resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationRuntimeManaged(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationRuntimeManagedState, opts ...pulumi.ResourceOption) (*IntegrationRuntimeManaged, error) {
	var resource IntegrationRuntimeManaged
	err := ctx.ReadResource("azure:datafactory/integrationRuntimeManaged:IntegrationRuntimeManaged", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationRuntimeManaged resources.
type integrationRuntimeManagedState struct {
	// A `catalogInfo` block as defined below.
	CatalogInfo *IntegrationRuntimeManagedCatalogInfo `pulumi:"catalogInfo"`
	// A `customSetupScript` block as defined below.
	CustomSetupScript *IntegrationRuntimeManagedCustomSetupScript `pulumi:"customSetupScript"`
	// Specifies the name of the Data Factory the Managed Integration Runtime belongs to. Changing this forces a new resource to be created.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// Integration runtime description.
	Description *string `pulumi:"description"`
	// The Managed Integration Runtime edition. Valid values are `Standard` and `Enterprise`. Defaults to `Standard`.
	Edition *string `pulumi:"edition"`
	// The type of the license that is used. Valid values are `LicenseIncluded` and `BasePrize`. Defaults to `LicenseIncluded`.
	LicenseType *string `pulumi:"licenseType"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Defines the maximum parallel executions per node. Defaults to `1`. Max is `16`.
	MaxParallelExecutionsPerNode *int `pulumi:"maxParallelExecutionsPerNode"`
	// Specifies the name of the Managed Integration Runtime. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// The size of the nodes on which the Managed Integration Runtime runs. Valid values are: `Standard_D2_v3`, `Standard_D4_v3`, `Standard_D8_v3`, `Standard_D16_v3`, `Standard_D32_v3`, `Standard_D64_v3`, `Standard_E2_v3`, `Standard_E4_v3`, `Standard_E8_v3`, `Standard_E16_v3`, `Standard_E32_v3`, `Standard_E64_v3`, `Standard_D1_v2`, `Standard_D2_v2`, `Standard_D3_v2`, `Standard_D4_v2`, `Standard_A4_v2` and `Standard_A8_v2`
	NodeSize *string `pulumi:"nodeSize"`
	// Number of nodes for the Managed Integration Runtime. Max is `10`. Defaults to `1`.
	NumberOfNodes *int `pulumi:"numberOfNodes"`
	// The name of the resource group in which to create the Managed Integration Runtime. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `vnetIntegration` block as defined below.
	VnetIntegration *IntegrationRuntimeManagedVnetIntegration `pulumi:"vnetIntegration"`
}

type IntegrationRuntimeManagedState struct {
	// A `catalogInfo` block as defined below.
	CatalogInfo IntegrationRuntimeManagedCatalogInfoPtrInput
	// A `customSetupScript` block as defined below.
	CustomSetupScript IntegrationRuntimeManagedCustomSetupScriptPtrInput
	// Specifies the name of the Data Factory the Managed Integration Runtime belongs to. Changing this forces a new resource to be created.
	DataFactoryName pulumi.StringPtrInput
	// Integration runtime description.
	Description pulumi.StringPtrInput
	// The Managed Integration Runtime edition. Valid values are `Standard` and `Enterprise`. Defaults to `Standard`.
	Edition pulumi.StringPtrInput
	// The type of the license that is used. Valid values are `LicenseIncluded` and `BasePrize`. Defaults to `LicenseIncluded`.
	LicenseType pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Defines the maximum parallel executions per node. Defaults to `1`. Max is `16`.
	MaxParallelExecutionsPerNode pulumi.IntPtrInput
	// Specifies the name of the Managed Integration Runtime. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// The size of the nodes on which the Managed Integration Runtime runs. Valid values are: `Standard_D2_v3`, `Standard_D4_v3`, `Standard_D8_v3`, `Standard_D16_v3`, `Standard_D32_v3`, `Standard_D64_v3`, `Standard_E2_v3`, `Standard_E4_v3`, `Standard_E8_v3`, `Standard_E16_v3`, `Standard_E32_v3`, `Standard_E64_v3`, `Standard_D1_v2`, `Standard_D2_v2`, `Standard_D3_v2`, `Standard_D4_v2`, `Standard_A4_v2` and `Standard_A8_v2`
	NodeSize pulumi.StringPtrInput
	// Number of nodes for the Managed Integration Runtime. Max is `10`. Defaults to `1`.
	NumberOfNodes pulumi.IntPtrInput
	// The name of the resource group in which to create the Managed Integration Runtime. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `vnetIntegration` block as defined below.
	VnetIntegration IntegrationRuntimeManagedVnetIntegrationPtrInput
}

func (IntegrationRuntimeManagedState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationRuntimeManagedState)(nil)).Elem()
}

type integrationRuntimeManagedArgs struct {
	// A `catalogInfo` block as defined below.
	CatalogInfo *IntegrationRuntimeManagedCatalogInfo `pulumi:"catalogInfo"`
	// A `customSetupScript` block as defined below.
	CustomSetupScript *IntegrationRuntimeManagedCustomSetupScript `pulumi:"customSetupScript"`
	// Specifies the name of the Data Factory the Managed Integration Runtime belongs to. Changing this forces a new resource to be created.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// Integration runtime description.
	Description *string `pulumi:"description"`
	// The Managed Integration Runtime edition. Valid values are `Standard` and `Enterprise`. Defaults to `Standard`.
	Edition *string `pulumi:"edition"`
	// The type of the license that is used. Valid values are `LicenseIncluded` and `BasePrize`. Defaults to `LicenseIncluded`.
	LicenseType *string `pulumi:"licenseType"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Defines the maximum parallel executions per node. Defaults to `1`. Max is `16`.
	MaxParallelExecutionsPerNode *int `pulumi:"maxParallelExecutionsPerNode"`
	// Specifies the name of the Managed Integration Runtime. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// The size of the nodes on which the Managed Integration Runtime runs. Valid values are: `Standard_D2_v3`, `Standard_D4_v3`, `Standard_D8_v3`, `Standard_D16_v3`, `Standard_D32_v3`, `Standard_D64_v3`, `Standard_E2_v3`, `Standard_E4_v3`, `Standard_E8_v3`, `Standard_E16_v3`, `Standard_E32_v3`, `Standard_E64_v3`, `Standard_D1_v2`, `Standard_D2_v2`, `Standard_D3_v2`, `Standard_D4_v2`, `Standard_A4_v2` and `Standard_A8_v2`
	NodeSize string `pulumi:"nodeSize"`
	// Number of nodes for the Managed Integration Runtime. Max is `10`. Defaults to `1`.
	NumberOfNodes *int `pulumi:"numberOfNodes"`
	// The name of the resource group in which to create the Managed Integration Runtime. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `vnetIntegration` block as defined below.
	VnetIntegration *IntegrationRuntimeManagedVnetIntegration `pulumi:"vnetIntegration"`
}

// The set of arguments for constructing a IntegrationRuntimeManaged resource.
type IntegrationRuntimeManagedArgs struct {
	// A `catalogInfo` block as defined below.
	CatalogInfo IntegrationRuntimeManagedCatalogInfoPtrInput
	// A `customSetupScript` block as defined below.
	CustomSetupScript IntegrationRuntimeManagedCustomSetupScriptPtrInput
	// Specifies the name of the Data Factory the Managed Integration Runtime belongs to. Changing this forces a new resource to be created.
	DataFactoryName pulumi.StringInput
	// Integration runtime description.
	Description pulumi.StringPtrInput
	// The Managed Integration Runtime edition. Valid values are `Standard` and `Enterprise`. Defaults to `Standard`.
	Edition pulumi.StringPtrInput
	// The type of the license that is used. Valid values are `LicenseIncluded` and `BasePrize`. Defaults to `LicenseIncluded`.
	LicenseType pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Defines the maximum parallel executions per node. Defaults to `1`. Max is `16`.
	MaxParallelExecutionsPerNode pulumi.IntPtrInput
	// Specifies the name of the Managed Integration Runtime. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// The size of the nodes on which the Managed Integration Runtime runs. Valid values are: `Standard_D2_v3`, `Standard_D4_v3`, `Standard_D8_v3`, `Standard_D16_v3`, `Standard_D32_v3`, `Standard_D64_v3`, `Standard_E2_v3`, `Standard_E4_v3`, `Standard_E8_v3`, `Standard_E16_v3`, `Standard_E32_v3`, `Standard_E64_v3`, `Standard_D1_v2`, `Standard_D2_v2`, `Standard_D3_v2`, `Standard_D4_v2`, `Standard_A4_v2` and `Standard_A8_v2`
	NodeSize pulumi.StringInput
	// Number of nodes for the Managed Integration Runtime. Max is `10`. Defaults to `1`.
	NumberOfNodes pulumi.IntPtrInput
	// The name of the resource group in which to create the Managed Integration Runtime. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `vnetIntegration` block as defined below.
	VnetIntegration IntegrationRuntimeManagedVnetIntegrationPtrInput
}

func (IntegrationRuntimeManagedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationRuntimeManagedArgs)(nil)).Elem()
}

type IntegrationRuntimeManagedInput interface {
	pulumi.Input

	ToIntegrationRuntimeManagedOutput() IntegrationRuntimeManagedOutput
	ToIntegrationRuntimeManagedOutputWithContext(ctx context.Context) IntegrationRuntimeManagedOutput
}

func (*IntegrationRuntimeManaged) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationRuntimeManaged)(nil))
}

func (i *IntegrationRuntimeManaged) ToIntegrationRuntimeManagedOutput() IntegrationRuntimeManagedOutput {
	return i.ToIntegrationRuntimeManagedOutputWithContext(context.Background())
}

func (i *IntegrationRuntimeManaged) ToIntegrationRuntimeManagedOutputWithContext(ctx context.Context) IntegrationRuntimeManagedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRuntimeManagedOutput)
}

func (i *IntegrationRuntimeManaged) ToIntegrationRuntimeManagedPtrOutput() IntegrationRuntimeManagedPtrOutput {
	return i.ToIntegrationRuntimeManagedPtrOutputWithContext(context.Background())
}

func (i *IntegrationRuntimeManaged) ToIntegrationRuntimeManagedPtrOutputWithContext(ctx context.Context) IntegrationRuntimeManagedPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRuntimeManagedPtrOutput)
}

type IntegrationRuntimeManagedPtrInput interface {
	pulumi.Input

	ToIntegrationRuntimeManagedPtrOutput() IntegrationRuntimeManagedPtrOutput
	ToIntegrationRuntimeManagedPtrOutputWithContext(ctx context.Context) IntegrationRuntimeManagedPtrOutput
}

type integrationRuntimeManagedPtrType IntegrationRuntimeManagedArgs

func (*integrationRuntimeManagedPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationRuntimeManaged)(nil))
}

func (i *integrationRuntimeManagedPtrType) ToIntegrationRuntimeManagedPtrOutput() IntegrationRuntimeManagedPtrOutput {
	return i.ToIntegrationRuntimeManagedPtrOutputWithContext(context.Background())
}

func (i *integrationRuntimeManagedPtrType) ToIntegrationRuntimeManagedPtrOutputWithContext(ctx context.Context) IntegrationRuntimeManagedPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRuntimeManagedPtrOutput)
}

// IntegrationRuntimeManagedArrayInput is an input type that accepts IntegrationRuntimeManagedArray and IntegrationRuntimeManagedArrayOutput values.
// You can construct a concrete instance of `IntegrationRuntimeManagedArrayInput` via:
//
//          IntegrationRuntimeManagedArray{ IntegrationRuntimeManagedArgs{...} }
type IntegrationRuntimeManagedArrayInput interface {
	pulumi.Input

	ToIntegrationRuntimeManagedArrayOutput() IntegrationRuntimeManagedArrayOutput
	ToIntegrationRuntimeManagedArrayOutputWithContext(context.Context) IntegrationRuntimeManagedArrayOutput
}

type IntegrationRuntimeManagedArray []IntegrationRuntimeManagedInput

func (IntegrationRuntimeManagedArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*IntegrationRuntimeManaged)(nil))
}

func (i IntegrationRuntimeManagedArray) ToIntegrationRuntimeManagedArrayOutput() IntegrationRuntimeManagedArrayOutput {
	return i.ToIntegrationRuntimeManagedArrayOutputWithContext(context.Background())
}

func (i IntegrationRuntimeManagedArray) ToIntegrationRuntimeManagedArrayOutputWithContext(ctx context.Context) IntegrationRuntimeManagedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRuntimeManagedArrayOutput)
}

// IntegrationRuntimeManagedMapInput is an input type that accepts IntegrationRuntimeManagedMap and IntegrationRuntimeManagedMapOutput values.
// You can construct a concrete instance of `IntegrationRuntimeManagedMapInput` via:
//
//          IntegrationRuntimeManagedMap{ "key": IntegrationRuntimeManagedArgs{...} }
type IntegrationRuntimeManagedMapInput interface {
	pulumi.Input

	ToIntegrationRuntimeManagedMapOutput() IntegrationRuntimeManagedMapOutput
	ToIntegrationRuntimeManagedMapOutputWithContext(context.Context) IntegrationRuntimeManagedMapOutput
}

type IntegrationRuntimeManagedMap map[string]IntegrationRuntimeManagedInput

func (IntegrationRuntimeManagedMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*IntegrationRuntimeManaged)(nil))
}

func (i IntegrationRuntimeManagedMap) ToIntegrationRuntimeManagedMapOutput() IntegrationRuntimeManagedMapOutput {
	return i.ToIntegrationRuntimeManagedMapOutputWithContext(context.Background())
}

func (i IntegrationRuntimeManagedMap) ToIntegrationRuntimeManagedMapOutputWithContext(ctx context.Context) IntegrationRuntimeManagedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRuntimeManagedMapOutput)
}

type IntegrationRuntimeManagedOutput struct {
	*pulumi.OutputState
}

func (IntegrationRuntimeManagedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationRuntimeManaged)(nil))
}

func (o IntegrationRuntimeManagedOutput) ToIntegrationRuntimeManagedOutput() IntegrationRuntimeManagedOutput {
	return o
}

func (o IntegrationRuntimeManagedOutput) ToIntegrationRuntimeManagedOutputWithContext(ctx context.Context) IntegrationRuntimeManagedOutput {
	return o
}

func (o IntegrationRuntimeManagedOutput) ToIntegrationRuntimeManagedPtrOutput() IntegrationRuntimeManagedPtrOutput {
	return o.ToIntegrationRuntimeManagedPtrOutputWithContext(context.Background())
}

func (o IntegrationRuntimeManagedOutput) ToIntegrationRuntimeManagedPtrOutputWithContext(ctx context.Context) IntegrationRuntimeManagedPtrOutput {
	return o.ApplyT(func(v IntegrationRuntimeManaged) *IntegrationRuntimeManaged {
		return &v
	}).(IntegrationRuntimeManagedPtrOutput)
}

type IntegrationRuntimeManagedPtrOutput struct {
	*pulumi.OutputState
}

func (IntegrationRuntimeManagedPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationRuntimeManaged)(nil))
}

func (o IntegrationRuntimeManagedPtrOutput) ToIntegrationRuntimeManagedPtrOutput() IntegrationRuntimeManagedPtrOutput {
	return o
}

func (o IntegrationRuntimeManagedPtrOutput) ToIntegrationRuntimeManagedPtrOutputWithContext(ctx context.Context) IntegrationRuntimeManagedPtrOutput {
	return o
}

type IntegrationRuntimeManagedArrayOutput struct{ *pulumi.OutputState }

func (IntegrationRuntimeManagedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IntegrationRuntimeManaged)(nil))
}

func (o IntegrationRuntimeManagedArrayOutput) ToIntegrationRuntimeManagedArrayOutput() IntegrationRuntimeManagedArrayOutput {
	return o
}

func (o IntegrationRuntimeManagedArrayOutput) ToIntegrationRuntimeManagedArrayOutputWithContext(ctx context.Context) IntegrationRuntimeManagedArrayOutput {
	return o
}

func (o IntegrationRuntimeManagedArrayOutput) Index(i pulumi.IntInput) IntegrationRuntimeManagedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IntegrationRuntimeManaged {
		return vs[0].([]IntegrationRuntimeManaged)[vs[1].(int)]
	}).(IntegrationRuntimeManagedOutput)
}

type IntegrationRuntimeManagedMapOutput struct{ *pulumi.OutputState }

func (IntegrationRuntimeManagedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IntegrationRuntimeManaged)(nil))
}

func (o IntegrationRuntimeManagedMapOutput) ToIntegrationRuntimeManagedMapOutput() IntegrationRuntimeManagedMapOutput {
	return o
}

func (o IntegrationRuntimeManagedMapOutput) ToIntegrationRuntimeManagedMapOutputWithContext(ctx context.Context) IntegrationRuntimeManagedMapOutput {
	return o
}

func (o IntegrationRuntimeManagedMapOutput) MapIndex(k pulumi.StringInput) IntegrationRuntimeManagedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IntegrationRuntimeManaged {
		return vs[0].(map[string]IntegrationRuntimeManaged)[vs[1].(string)]
	}).(IntegrationRuntimeManagedOutput)
}

func init() {
	pulumi.RegisterOutputType(IntegrationRuntimeManagedOutput{})
	pulumi.RegisterOutputType(IntegrationRuntimeManagedPtrOutput{})
	pulumi.RegisterOutputType(IntegrationRuntimeManagedArrayOutput{})
	pulumi.RegisterOutputType(IntegrationRuntimeManagedMapOutput{})
}

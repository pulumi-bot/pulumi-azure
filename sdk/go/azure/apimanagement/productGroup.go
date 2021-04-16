// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an API Management Product Assignment to a Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/apimanagement"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleService, err := apimanagement.LookupService(ctx, &apimanagement.LookupServiceArgs{
// 			Name:              "example-api",
// 			ResourceGroupName: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleProduct, err := apimanagement.LookupProduct(ctx, &apimanagement.LookupProductArgs{
// 			ProductId:         "my-product",
// 			ApiManagementName: exampleService.Name,
// 			ResourceGroupName: exampleService.ResourceGroupName,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleGroup, err := apimanagement.LookupGroup(ctx, &apimanagement.LookupGroupArgs{
// 			Name:              "my-group",
// 			ApiManagementName: exampleService.Name,
// 			ResourceGroupName: exampleService.ResourceGroupName,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apimanagement.NewProductGroup(ctx, "exampleProductGroup", &apimanagement.ProductGroupArgs{
// 			ProductId:         pulumi.String(exampleProduct.ProductId),
// 			GroupName:         pulumi.String(exampleGroup.Name),
// 			ApiManagementName: pulumi.String(exampleService.Name),
// 			ResourceGroupName: pulumi.String(exampleService.ResourceGroupName),
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
// API Management Product Groups can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/productGroup:ProductGroup example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/products/exampleId/groups/groupId
// ```
type ProductGroup struct {
	pulumi.CustomResourceState

	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	ProductId pulumi.StringOutput `pulumi:"productId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewProductGroup registers a new resource with the given unique name, arguments, and options.
func NewProductGroup(ctx *pulumi.Context,
	name string, args *ProductGroupArgs, opts ...pulumi.ResourceOption) (*ProductGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiManagementName == nil {
		return nil, errors.New("invalid value for required argument 'ApiManagementName'")
	}
	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ProductGroup
	err := ctx.RegisterResource("azure:apimanagement/productGroup:ProductGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProductGroup gets an existing ProductGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProductGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductGroupState, opts ...pulumi.ResourceOption) (*ProductGroup, error) {
	var resource ProductGroup
	err := ctx.ReadResource("azure:apimanagement/productGroup:ProductGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProductGroup resources.
type productGroupState struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
	GroupName *string `pulumi:"groupName"`
	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	ProductId *string `pulumi:"productId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type ProductGroupState struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
	GroupName pulumi.StringPtrInput
	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	ProductId pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (ProductGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*productGroupState)(nil)).Elem()
}

type productGroupArgs struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
	GroupName string `pulumi:"groupName"`
	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	ProductId string `pulumi:"productId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ProductGroup resource.
type ProductGroupArgs struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
	GroupName pulumi.StringInput
	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	ProductId pulumi.StringInput
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (ProductGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productGroupArgs)(nil)).Elem()
}

type ProductGroupInput interface {
	pulumi.Input

	ToProductGroupOutput() ProductGroupOutput
	ToProductGroupOutputWithContext(ctx context.Context) ProductGroupOutput
}

func (*ProductGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ProductGroup)(nil))
}

func (i *ProductGroup) ToProductGroupOutput() ProductGroupOutput {
	return i.ToProductGroupOutputWithContext(context.Background())
}

func (i *ProductGroup) ToProductGroupOutputWithContext(ctx context.Context) ProductGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductGroupOutput)
}

func (i *ProductGroup) ToProductGroupPtrOutput() ProductGroupPtrOutput {
	return i.ToProductGroupPtrOutputWithContext(context.Background())
}

func (i *ProductGroup) ToProductGroupPtrOutputWithContext(ctx context.Context) ProductGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductGroupPtrOutput)
}

type ProductGroupPtrInput interface {
	pulumi.Input

	ToProductGroupPtrOutput() ProductGroupPtrOutput
	ToProductGroupPtrOutputWithContext(ctx context.Context) ProductGroupPtrOutput
}

type productGroupPtrType ProductGroupArgs

func (*productGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductGroup)(nil))
}

func (i *productGroupPtrType) ToProductGroupPtrOutput() ProductGroupPtrOutput {
	return i.ToProductGroupPtrOutputWithContext(context.Background())
}

func (i *productGroupPtrType) ToProductGroupPtrOutputWithContext(ctx context.Context) ProductGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductGroupPtrOutput)
}

// ProductGroupArrayInput is an input type that accepts ProductGroupArray and ProductGroupArrayOutput values.
// You can construct a concrete instance of `ProductGroupArrayInput` via:
//
//          ProductGroupArray{ ProductGroupArgs{...} }
type ProductGroupArrayInput interface {
	pulumi.Input

	ToProductGroupArrayOutput() ProductGroupArrayOutput
	ToProductGroupArrayOutputWithContext(context.Context) ProductGroupArrayOutput
}

type ProductGroupArray []ProductGroupInput

func (ProductGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ProductGroup)(nil))
}

func (i ProductGroupArray) ToProductGroupArrayOutput() ProductGroupArrayOutput {
	return i.ToProductGroupArrayOutputWithContext(context.Background())
}

func (i ProductGroupArray) ToProductGroupArrayOutputWithContext(ctx context.Context) ProductGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductGroupArrayOutput)
}

// ProductGroupMapInput is an input type that accepts ProductGroupMap and ProductGroupMapOutput values.
// You can construct a concrete instance of `ProductGroupMapInput` via:
//
//          ProductGroupMap{ "key": ProductGroupArgs{...} }
type ProductGroupMapInput interface {
	pulumi.Input

	ToProductGroupMapOutput() ProductGroupMapOutput
	ToProductGroupMapOutputWithContext(context.Context) ProductGroupMapOutput
}

type ProductGroupMap map[string]ProductGroupInput

func (ProductGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ProductGroup)(nil))
}

func (i ProductGroupMap) ToProductGroupMapOutput() ProductGroupMapOutput {
	return i.ToProductGroupMapOutputWithContext(context.Background())
}

func (i ProductGroupMap) ToProductGroupMapOutputWithContext(ctx context.Context) ProductGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductGroupMapOutput)
}

type ProductGroupOutput struct {
	*pulumi.OutputState
}

func (ProductGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProductGroup)(nil))
}

func (o ProductGroupOutput) ToProductGroupOutput() ProductGroupOutput {
	return o
}

func (o ProductGroupOutput) ToProductGroupOutputWithContext(ctx context.Context) ProductGroupOutput {
	return o
}

func (o ProductGroupOutput) ToProductGroupPtrOutput() ProductGroupPtrOutput {
	return o.ToProductGroupPtrOutputWithContext(context.Background())
}

func (o ProductGroupOutput) ToProductGroupPtrOutputWithContext(ctx context.Context) ProductGroupPtrOutput {
	return o.ApplyT(func(v ProductGroup) *ProductGroup {
		return &v
	}).(ProductGroupPtrOutput)
}

type ProductGroupPtrOutput struct {
	*pulumi.OutputState
}

func (ProductGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductGroup)(nil))
}

func (o ProductGroupPtrOutput) ToProductGroupPtrOutput() ProductGroupPtrOutput {
	return o
}

func (o ProductGroupPtrOutput) ToProductGroupPtrOutputWithContext(ctx context.Context) ProductGroupPtrOutput {
	return o
}

type ProductGroupArrayOutput struct{ *pulumi.OutputState }

func (ProductGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProductGroup)(nil))
}

func (o ProductGroupArrayOutput) ToProductGroupArrayOutput() ProductGroupArrayOutput {
	return o
}

func (o ProductGroupArrayOutput) ToProductGroupArrayOutputWithContext(ctx context.Context) ProductGroupArrayOutput {
	return o
}

func (o ProductGroupArrayOutput) Index(i pulumi.IntInput) ProductGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProductGroup {
		return vs[0].([]ProductGroup)[vs[1].(int)]
	}).(ProductGroupOutput)
}

type ProductGroupMapOutput struct{ *pulumi.OutputState }

func (ProductGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ProductGroup)(nil))
}

func (o ProductGroupMapOutput) ToProductGroupMapOutput() ProductGroupMapOutput {
	return o
}

func (o ProductGroupMapOutput) ToProductGroupMapOutputWithContext(ctx context.Context) ProductGroupMapOutput {
	return o
}

func (o ProductGroupMapOutput) MapIndex(k pulumi.StringInput) ProductGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ProductGroup {
		return vs[0].(map[string]ProductGroup)[vs[1].(string)]
	}).(ProductGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(ProductGroupOutput{})
	pulumi.RegisterOutputType(ProductGroupPtrOutput{})
	pulumi.RegisterOutputType(ProductGroupArrayOutput{})
	pulumi.RegisterOutputType(ProductGroupMapOutput{})
}

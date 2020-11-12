// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an API Management API Assignment to a Product.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/apimanagement"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 		exampleApi, err := apimanagement.LookupApi(ctx, &apimanagement.LookupApiArgs{
// 			Name:              "search-api",
// 			ApiManagementName: exampleService.Name,
// 			ResourceGroupName: exampleService.ResourceGroupName,
// 			Revision:          "2",
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
// 		_, err = apimanagement.NewProductApi(ctx, "exampleProductApi", &apimanagement.ProductApiArgs{
// 			ApiName:           pulumi.String(exampleApi.Name),
// 			ProductId:         pulumi.String(exampleProduct.ProductId),
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
// API Management Product API's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/productApi:ProductApi example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/products/exampleId/apis/apiId
// ```
type ProductApi struct {
	pulumi.CustomResourceState

	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The Name of the API Management API within the API Management Service. Changing this forces a new resource to be created.
	ApiName pulumi.StringOutput `pulumi:"apiName"`
	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	ProductId pulumi.StringOutput `pulumi:"productId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewProductApi registers a new resource with the given unique name, arguments, and options.
func NewProductApi(ctx *pulumi.Context,
	name string, args *ProductApiArgs, opts ...pulumi.ResourceOption) (*ProductApi, error) {
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.ApiName == nil {
		return nil, errors.New("missing required argument 'ApiName'")
	}
	if args == nil || args.ProductId == nil {
		return nil, errors.New("missing required argument 'ProductId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ProductApiArgs{}
	}
	var resource ProductApi
	err := ctx.RegisterResource("azure:apimanagement/productApi:ProductApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProductApi gets an existing ProductApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProductApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductApiState, opts ...pulumi.ResourceOption) (*ProductApi, error) {
	var resource ProductApi
	err := ctx.ReadResource("azure:apimanagement/productApi:ProductApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProductApi resources.
type productApiState struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The Name of the API Management API within the API Management Service. Changing this forces a new resource to be created.
	ApiName *string `pulumi:"apiName"`
	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	ProductId *string `pulumi:"productId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type ProductApiState struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The Name of the API Management API within the API Management Service. Changing this forces a new resource to be created.
	ApiName pulumi.StringPtrInput
	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	ProductId pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (ProductApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*productApiState)(nil)).Elem()
}

type productApiArgs struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The Name of the API Management API within the API Management Service. Changing this forces a new resource to be created.
	ApiName string `pulumi:"apiName"`
	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	ProductId string `pulumi:"productId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ProductApi resource.
type ProductApiArgs struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The Name of the API Management API within the API Management Service. Changing this forces a new resource to be created.
	ApiName pulumi.StringInput
	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	ProductId pulumi.StringInput
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (ProductApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productApiArgs)(nil)).Elem()
}

type ProductApiInput interface {
	pulumi.Input

	ToProductApiOutput() ProductApiOutput
	ToProductApiOutputWithContext(ctx context.Context) ProductApiOutput
}

func (ProductApi) ElementType() reflect.Type {
	return reflect.TypeOf((*ProductApi)(nil)).Elem()
}

func (i ProductApi) ToProductApiOutput() ProductApiOutput {
	return i.ToProductApiOutputWithContext(context.Background())
}

func (i ProductApi) ToProductApiOutputWithContext(ctx context.Context) ProductApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductApiOutput)
}

type ProductApiOutput struct {
	*pulumi.OutputState
}

func (ProductApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProductApiOutput)(nil)).Elem()
}

func (o ProductApiOutput) ToProductApiOutput() ProductApiOutput {
	return o
}

func (o ProductApiOutput) ToProductApiOutputWithContext(ctx context.Context) ProductApiOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProductApiOutput{})
}

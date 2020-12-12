// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an API Management Product.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/apimanagement"
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
// 		exampleService, err := apimanagement.NewService(ctx, "exampleService", &apimanagement.ServiceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			PublisherName:     pulumi.String("My Company"),
// 			PublisherEmail:    pulumi.String("company@exmaple.com"),
// 			SkuName:           pulumi.String("Developer_1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apimanagement.NewProduct(ctx, "exampleProduct", &apimanagement.ProductArgs{
// 			ProductId:            pulumi.String("test-product"),
// 			ApiManagementName:    exampleService.Name,
// 			ResourceGroupName:    exampleResourceGroup.Name,
// 			DisplayName:          pulumi.String("Test Product"),
// 			SubscriptionRequired: pulumi.Bool(true),
// 			ApprovalRequired:     pulumi.Bool(true),
// 			Published:            pulumi.Bool(true),
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
// API Management Products can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/product:Product example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/products/myproduct
// ```
type Product struct {
	pulumi.CustomResourceState

	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// Do subscribers need to be approved prior to being able to use the Product?
	ApprovalRequired pulumi.BoolPtrOutput `pulumi:"approvalRequired"`
	// A description of this Product, which may include HTML formatting tags.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Display Name for this API Management Product.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The Identifier for this Product, which must be unique within the API Management Service. Changing this forces a new resource to be created.
	ProductId pulumi.StringOutput `pulumi:"productId"`
	// Is this Product Published?
	Published pulumi.BoolOutput `pulumi:"published"`
	// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Is a Subscription required to access API's included in this Product?
	SubscriptionRequired pulumi.BoolOutput `pulumi:"subscriptionRequired"`
	// The number of subscriptions a user can have to this Product at the same time.
	SubscriptionsLimit pulumi.IntPtrOutput `pulumi:"subscriptionsLimit"`
	// The Terms and Conditions for this Product, which must be accepted by Developers before they can begin the Subscription process.
	Terms pulumi.StringPtrOutput `pulumi:"terms"`
}

// NewProduct registers a new resource with the given unique name, arguments, and options.
func NewProduct(ctx *pulumi.Context,
	name string, args *ProductArgs, opts ...pulumi.ResourceOption) (*Product, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiManagementName == nil {
		return nil, errors.New("invalid value for required argument 'ApiManagementName'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	if args.Published == nil {
		return nil, errors.New("invalid value for required argument 'Published'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SubscriptionRequired == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionRequired'")
	}
	var resource Product
	err := ctx.RegisterResource("azure:apimanagement/product:Product", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProduct gets an existing Product resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProduct(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductState, opts ...pulumi.ResourceOption) (*Product, error) {
	var resource Product
	err := ctx.ReadResource("azure:apimanagement/product:Product", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Product resources.
type productState struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// Do subscribers need to be approved prior to being able to use the Product?
	ApprovalRequired *bool `pulumi:"approvalRequired"`
	// A description of this Product, which may include HTML formatting tags.
	Description *string `pulumi:"description"`
	// The Display Name for this API Management Product.
	DisplayName *string `pulumi:"displayName"`
	// The Identifier for this Product, which must be unique within the API Management Service. Changing this forces a new resource to be created.
	ProductId *string `pulumi:"productId"`
	// Is this Product Published?
	Published *bool `pulumi:"published"`
	// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Is a Subscription required to access API's included in this Product?
	SubscriptionRequired *bool `pulumi:"subscriptionRequired"`
	// The number of subscriptions a user can have to this Product at the same time.
	SubscriptionsLimit *int `pulumi:"subscriptionsLimit"`
	// The Terms and Conditions for this Product, which must be accepted by Developers before they can begin the Subscription process.
	Terms *string `pulumi:"terms"`
}

type ProductState struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// Do subscribers need to be approved prior to being able to use the Product?
	ApprovalRequired pulumi.BoolPtrInput
	// A description of this Product, which may include HTML formatting tags.
	Description pulumi.StringPtrInput
	// The Display Name for this API Management Product.
	DisplayName pulumi.StringPtrInput
	// The Identifier for this Product, which must be unique within the API Management Service. Changing this forces a new resource to be created.
	ProductId pulumi.StringPtrInput
	// Is this Product Published?
	Published pulumi.BoolPtrInput
	// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Is a Subscription required to access API's included in this Product?
	SubscriptionRequired pulumi.BoolPtrInput
	// The number of subscriptions a user can have to this Product at the same time.
	SubscriptionsLimit pulumi.IntPtrInput
	// The Terms and Conditions for this Product, which must be accepted by Developers before they can begin the Subscription process.
	Terms pulumi.StringPtrInput
}

func (ProductState) ElementType() reflect.Type {
	return reflect.TypeOf((*productState)(nil)).Elem()
}

type productArgs struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// Do subscribers need to be approved prior to being able to use the Product?
	ApprovalRequired *bool `pulumi:"approvalRequired"`
	// A description of this Product, which may include HTML formatting tags.
	Description *string `pulumi:"description"`
	// The Display Name for this API Management Product.
	DisplayName string `pulumi:"displayName"`
	// The Identifier for this Product, which must be unique within the API Management Service. Changing this forces a new resource to be created.
	ProductId string `pulumi:"productId"`
	// Is this Product Published?
	Published bool `pulumi:"published"`
	// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Is a Subscription required to access API's included in this Product?
	SubscriptionRequired bool `pulumi:"subscriptionRequired"`
	// The number of subscriptions a user can have to this Product at the same time.
	SubscriptionsLimit *int `pulumi:"subscriptionsLimit"`
	// The Terms and Conditions for this Product, which must be accepted by Developers before they can begin the Subscription process.
	Terms *string `pulumi:"terms"`
}

// The set of arguments for constructing a Product resource.
type ProductArgs struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// Do subscribers need to be approved prior to being able to use the Product?
	ApprovalRequired pulumi.BoolPtrInput
	// A description of this Product, which may include HTML formatting tags.
	Description pulumi.StringPtrInput
	// The Display Name for this API Management Product.
	DisplayName pulumi.StringInput
	// The Identifier for this Product, which must be unique within the API Management Service. Changing this forces a new resource to be created.
	ProductId pulumi.StringInput
	// Is this Product Published?
	Published pulumi.BoolInput
	// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Is a Subscription required to access API's included in this Product?
	SubscriptionRequired pulumi.BoolInput
	// The number of subscriptions a user can have to this Product at the same time.
	SubscriptionsLimit pulumi.IntPtrInput
	// The Terms and Conditions for this Product, which must be accepted by Developers before they can begin the Subscription process.
	Terms pulumi.StringPtrInput
}

func (ProductArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productArgs)(nil)).Elem()
}

type ProductInput interface {
	pulumi.Input

	ToProductOutput() ProductOutput
	ToProductOutputWithContext(ctx context.Context) ProductOutput
}

func (*Product) ElementType() reflect.Type {
	return reflect.TypeOf((*Product)(nil))
}

func (i *Product) ToProductOutput() ProductOutput {
	return i.ToProductOutputWithContext(context.Background())
}

func (i *Product) ToProductOutputWithContext(ctx context.Context) ProductOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductOutput)
}

func (i *Product) ToProductPtrOutput() ProductPtrOutput {
	return i.ToProductPtrOutputWithContext(context.Background())
}

func (i *Product) ToProductPtrOutputWithContext(ctx context.Context) ProductPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductPtrOutput)
}

type ProductPtrInput interface {
	pulumi.Input

	ToProductPtrOutput() ProductPtrOutput
	ToProductPtrOutputWithContext(ctx context.Context) ProductPtrOutput
}

type ProductOutput struct {
	*pulumi.OutputState
}

func (ProductOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Product)(nil))
}

func (o ProductOutput) ToProductOutput() ProductOutput {
	return o
}

func (o ProductOutput) ToProductOutputWithContext(ctx context.Context) ProductOutput {
	return o
}

type ProductPtrOutput struct {
	*pulumi.OutputState
}

func (ProductPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Product)(nil))
}

func (o ProductPtrOutput) ToProductPtrOutput() ProductPtrOutput {
	return o
}

func (o ProductPtrOutput) ToProductPtrOutputWithContext(ctx context.Context) ProductPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProductOutput{})
	pulumi.RegisterOutputType(ProductPtrOutput{})
}

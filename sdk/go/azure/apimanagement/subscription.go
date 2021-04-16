// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Subscription within a API Management Service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/apimanagement"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleService, err := apimanagement.LookupService(ctx, &apimanagement.LookupServiceArgs{
// 			Name:              "example-apim",
// 			ResourceGroupName: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleProduct, err := apimanagement.LookupProduct(ctx, &apimanagement.LookupProductArgs{
// 			ProductId:         "00000000-0000-0000-0000-000000000000",
// 			ApiManagementName: exampleService.Name,
// 			ResourceGroupName: exampleService.ResourceGroupName,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleUser, err := apimanagement.LookupUser(ctx, &apimanagement.LookupUserArgs{
// 			UserId:            "11111111-1111-1111-1111-111111111111",
// 			ApiManagementName: exampleService.Name,
// 			ResourceGroupName: exampleService.ResourceGroupName,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apimanagement.NewSubscription(ctx, "exampleSubscription", &apimanagement.SubscriptionArgs{
// 			ApiManagementName: pulumi.String(exampleService.Name),
// 			ResourceGroupName: pulumi.String(exampleService.ResourceGroupName),
// 			UserId:            pulumi.String(exampleUser.Id),
// 			ProductId:         pulumi.String(exampleProduct.Id),
// 			DisplayName:       pulumi.String("Parser API"),
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
// API Management Subscriptions can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/subscription:Subscription example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.ApiManagement/service/example-apim/subscriptions/subscription-name
// ```
type Subscription struct {
	pulumi.CustomResourceState

	// Determines whether tracing can be enabled.  Defaults to `true`.
	AllowTracing pulumi.BoolPtrOutput `pulumi:"allowTracing"`
	// The name of the API Management Service where this Subscription should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The display name of this Subscription.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	PrimaryKey  pulumi.StringOutput `pulumi:"primaryKey"`
	// The ID of the Product which should be assigned to this Subscription. Changing this forces a new resource to be created.
	ProductId pulumi.StringPtrOutput `pulumi:"productId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	SecondaryKey      pulumi.StringOutput `pulumi:"secondaryKey"`
	// The state of this Subscription. Possible values are `active`, `cancelled`, `expired`, `rejected`, `submitted` and `suspended`. Defaults to `submitted`.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// An Identifier which should used as the ID of this Subscription. If not specified a new Subscription ID will be generated. Changing this forces a new resource to be created.
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// The ID of the User which should be assigned to this Subscription. Changing this forces a new resource to be created.
	UserId pulumi.StringPtrOutput `pulumi:"userId"`
}

// NewSubscription registers a new resource with the given unique name, arguments, and options.
func NewSubscription(ctx *pulumi.Context,
	name string, args *SubscriptionArgs, opts ...pulumi.ResourceOption) (*Subscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiManagementName == nil {
		return nil, errors.New("invalid value for required argument 'ApiManagementName'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Subscription
	err := ctx.RegisterResource("azure:apimanagement/subscription:Subscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscription gets an existing Subscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionState, opts ...pulumi.ResourceOption) (*Subscription, error) {
	var resource Subscription
	err := ctx.ReadResource("azure:apimanagement/subscription:Subscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subscription resources.
type subscriptionState struct {
	// Determines whether tracing can be enabled.  Defaults to `true`.
	AllowTracing *bool `pulumi:"allowTracing"`
	// The name of the API Management Service where this Subscription should be created. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The display name of this Subscription.
	DisplayName *string `pulumi:"displayName"`
	PrimaryKey  *string `pulumi:"primaryKey"`
	// The ID of the Product which should be assigned to this Subscription. Changing this forces a new resource to be created.
	ProductId *string `pulumi:"productId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	SecondaryKey      *string `pulumi:"secondaryKey"`
	// The state of this Subscription. Possible values are `active`, `cancelled`, `expired`, `rejected`, `submitted` and `suspended`. Defaults to `submitted`.
	State *string `pulumi:"state"`
	// An Identifier which should used as the ID of this Subscription. If not specified a new Subscription ID will be generated. Changing this forces a new resource to be created.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The ID of the User which should be assigned to this Subscription. Changing this forces a new resource to be created.
	UserId *string `pulumi:"userId"`
}

type SubscriptionState struct {
	// Determines whether tracing can be enabled.  Defaults to `true`.
	AllowTracing pulumi.BoolPtrInput
	// The name of the API Management Service where this Subscription should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The display name of this Subscription.
	DisplayName pulumi.StringPtrInput
	PrimaryKey  pulumi.StringPtrInput
	// The ID of the Product which should be assigned to this Subscription. Changing this forces a new resource to be created.
	ProductId pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	SecondaryKey      pulumi.StringPtrInput
	// The state of this Subscription. Possible values are `active`, `cancelled`, `expired`, `rejected`, `submitted` and `suspended`. Defaults to `submitted`.
	State pulumi.StringPtrInput
	// An Identifier which should used as the ID of this Subscription. If not specified a new Subscription ID will be generated. Changing this forces a new resource to be created.
	SubscriptionId pulumi.StringPtrInput
	// The ID of the User which should be assigned to this Subscription. Changing this forces a new resource to be created.
	UserId pulumi.StringPtrInput
}

func (SubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionState)(nil)).Elem()
}

type subscriptionArgs struct {
	// Determines whether tracing can be enabled.  Defaults to `true`.
	AllowTracing *bool `pulumi:"allowTracing"`
	// The name of the API Management Service where this Subscription should be created. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The display name of this Subscription.
	DisplayName string  `pulumi:"displayName"`
	PrimaryKey  *string `pulumi:"primaryKey"`
	// The ID of the Product which should be assigned to this Subscription. Changing this forces a new resource to be created.
	ProductId *string `pulumi:"productId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	SecondaryKey      *string `pulumi:"secondaryKey"`
	// The state of this Subscription. Possible values are `active`, `cancelled`, `expired`, `rejected`, `submitted` and `suspended`. Defaults to `submitted`.
	State *string `pulumi:"state"`
	// An Identifier which should used as the ID of this Subscription. If not specified a new Subscription ID will be generated. Changing this forces a new resource to be created.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The ID of the User which should be assigned to this Subscription. Changing this forces a new resource to be created.
	UserId *string `pulumi:"userId"`
}

// The set of arguments for constructing a Subscription resource.
type SubscriptionArgs struct {
	// Determines whether tracing can be enabled.  Defaults to `true`.
	AllowTracing pulumi.BoolPtrInput
	// The name of the API Management Service where this Subscription should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The display name of this Subscription.
	DisplayName pulumi.StringInput
	PrimaryKey  pulumi.StringPtrInput
	// The ID of the Product which should be assigned to this Subscription. Changing this forces a new resource to be created.
	ProductId pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	SecondaryKey      pulumi.StringPtrInput
	// The state of this Subscription. Possible values are `active`, `cancelled`, `expired`, `rejected`, `submitted` and `suspended`. Defaults to `submitted`.
	State pulumi.StringPtrInput
	// An Identifier which should used as the ID of this Subscription. If not specified a new Subscription ID will be generated. Changing this forces a new resource to be created.
	SubscriptionId pulumi.StringPtrInput
	// The ID of the User which should be assigned to this Subscription. Changing this forces a new resource to be created.
	UserId pulumi.StringPtrInput
}

func (SubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionArgs)(nil)).Elem()
}

type SubscriptionInput interface {
	pulumi.Input

	ToSubscriptionOutput() SubscriptionOutput
	ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput
}

func (*Subscription) ElementType() reflect.Type {
	return reflect.TypeOf((*Subscription)(nil))
}

func (i *Subscription) ToSubscriptionOutput() SubscriptionOutput {
	return i.ToSubscriptionOutputWithContext(context.Background())
}

func (i *Subscription) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionOutput)
}

func (i *Subscription) ToSubscriptionPtrOutput() SubscriptionPtrOutput {
	return i.ToSubscriptionPtrOutputWithContext(context.Background())
}

func (i *Subscription) ToSubscriptionPtrOutputWithContext(ctx context.Context) SubscriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionPtrOutput)
}

type SubscriptionPtrInput interface {
	pulumi.Input

	ToSubscriptionPtrOutput() SubscriptionPtrOutput
	ToSubscriptionPtrOutputWithContext(ctx context.Context) SubscriptionPtrOutput
}

type subscriptionPtrType SubscriptionArgs

func (*subscriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil))
}

func (i *subscriptionPtrType) ToSubscriptionPtrOutput() SubscriptionPtrOutput {
	return i.ToSubscriptionPtrOutputWithContext(context.Background())
}

func (i *subscriptionPtrType) ToSubscriptionPtrOutputWithContext(ctx context.Context) SubscriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionPtrOutput)
}

// SubscriptionArrayInput is an input type that accepts SubscriptionArray and SubscriptionArrayOutput values.
// You can construct a concrete instance of `SubscriptionArrayInput` via:
//
//          SubscriptionArray{ SubscriptionArgs{...} }
type SubscriptionArrayInput interface {
	pulumi.Input

	ToSubscriptionArrayOutput() SubscriptionArrayOutput
	ToSubscriptionArrayOutputWithContext(context.Context) SubscriptionArrayOutput
}

type SubscriptionArray []SubscriptionInput

func (SubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Subscription)(nil))
}

func (i SubscriptionArray) ToSubscriptionArrayOutput() SubscriptionArrayOutput {
	return i.ToSubscriptionArrayOutputWithContext(context.Background())
}

func (i SubscriptionArray) ToSubscriptionArrayOutputWithContext(ctx context.Context) SubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionArrayOutput)
}

// SubscriptionMapInput is an input type that accepts SubscriptionMap and SubscriptionMapOutput values.
// You can construct a concrete instance of `SubscriptionMapInput` via:
//
//          SubscriptionMap{ "key": SubscriptionArgs{...} }
type SubscriptionMapInput interface {
	pulumi.Input

	ToSubscriptionMapOutput() SubscriptionMapOutput
	ToSubscriptionMapOutputWithContext(context.Context) SubscriptionMapOutput
}

type SubscriptionMap map[string]SubscriptionInput

func (SubscriptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Subscription)(nil))
}

func (i SubscriptionMap) ToSubscriptionMapOutput() SubscriptionMapOutput {
	return i.ToSubscriptionMapOutputWithContext(context.Background())
}

func (i SubscriptionMap) ToSubscriptionMapOutputWithContext(ctx context.Context) SubscriptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionMapOutput)
}

type SubscriptionOutput struct {
	*pulumi.OutputState
}

func (SubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Subscription)(nil))
}

func (o SubscriptionOutput) ToSubscriptionOutput() SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) ToSubscriptionPtrOutput() SubscriptionPtrOutput {
	return o.ToSubscriptionPtrOutputWithContext(context.Background())
}

func (o SubscriptionOutput) ToSubscriptionPtrOutputWithContext(ctx context.Context) SubscriptionPtrOutput {
	return o.ApplyT(func(v Subscription) *Subscription {
		return &v
	}).(SubscriptionPtrOutput)
}

type SubscriptionPtrOutput struct {
	*pulumi.OutputState
}

func (SubscriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil))
}

func (o SubscriptionPtrOutput) ToSubscriptionPtrOutput() SubscriptionPtrOutput {
	return o
}

func (o SubscriptionPtrOutput) ToSubscriptionPtrOutputWithContext(ctx context.Context) SubscriptionPtrOutput {
	return o
}

type SubscriptionArrayOutput struct{ *pulumi.OutputState }

func (SubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Subscription)(nil))
}

func (o SubscriptionArrayOutput) ToSubscriptionArrayOutput() SubscriptionArrayOutput {
	return o
}

func (o SubscriptionArrayOutput) ToSubscriptionArrayOutputWithContext(ctx context.Context) SubscriptionArrayOutput {
	return o
}

func (o SubscriptionArrayOutput) Index(i pulumi.IntInput) SubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Subscription {
		return vs[0].([]Subscription)[vs[1].(int)]
	}).(SubscriptionOutput)
}

type SubscriptionMapOutput struct{ *pulumi.OutputState }

func (SubscriptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Subscription)(nil))
}

func (o SubscriptionMapOutput) ToSubscriptionMapOutput() SubscriptionMapOutput {
	return o
}

func (o SubscriptionMapOutput) ToSubscriptionMapOutputWithContext(ctx context.Context) SubscriptionMapOutput {
	return o
}

func (o SubscriptionMapOutput) MapIndex(k pulumi.StringInput) SubscriptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Subscription {
		return vs[0].(map[string]Subscription)[vs[1].(string)]
	}).(SubscriptionOutput)
}

func init() {
	pulumi.RegisterOutputType(SubscriptionOutput{})
	pulumi.RegisterOutputType(SubscriptionPtrOutput{})
	pulumi.RegisterOutputType(SubscriptionArrayOutput{})
	pulumi.RegisterOutputType(SubscriptionMapOutput{})
}

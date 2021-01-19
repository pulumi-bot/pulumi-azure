// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an API Management Facebook Identity Provider.
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
// 			PublisherEmail:    pulumi.String("company@mycompany.io"),
// 			SkuName:           pulumi.String("Developer_1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apimanagement.NewIdentityProviderFacebook(ctx, "exampleIdentityProviderFacebook", &apimanagement.IdentityProviderFacebookArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApiManagementName: exampleService.Name,
// 			AppId:             pulumi.String("00000000000000000000000000000000"),
// 			AppSecret:         pulumi.String("00000000000000000000000000000000"),
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
// API Management Facebook Identity Provider can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/identityProviderFacebook:IdentityProviderFacebook example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/facebook
// ```
type IdentityProviderFacebook struct {
	pulumi.CustomResourceState

	// The Name of the API Management Service where this Facebook Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// App ID for Facebook.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// App Secret for Facebook.
	AppSecret pulumi.StringOutput `pulumi:"appSecret"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewIdentityProviderFacebook registers a new resource with the given unique name, arguments, and options.
func NewIdentityProviderFacebook(ctx *pulumi.Context,
	name string, args *IdentityProviderFacebookArgs, opts ...pulumi.ResourceOption) (*IdentityProviderFacebook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiManagementName == nil {
		return nil, errors.New("invalid value for required argument 'ApiManagementName'")
	}
	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.AppSecret == nil {
		return nil, errors.New("invalid value for required argument 'AppSecret'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource IdentityProviderFacebook
	err := ctx.RegisterResource("azure:apimanagement/identityProviderFacebook:IdentityProviderFacebook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityProviderFacebook gets an existing IdentityProviderFacebook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityProviderFacebook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityProviderFacebookState, opts ...pulumi.ResourceOption) (*IdentityProviderFacebook, error) {
	var resource IdentityProviderFacebook
	err := ctx.ReadResource("azure:apimanagement/identityProviderFacebook:IdentityProviderFacebook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityProviderFacebook resources.
type identityProviderFacebookState struct {
	// The Name of the API Management Service where this Facebook Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// App ID for Facebook.
	AppId *string `pulumi:"appId"`
	// App Secret for Facebook.
	AppSecret *string `pulumi:"appSecret"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type IdentityProviderFacebookState struct {
	// The Name of the API Management Service where this Facebook Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// App ID for Facebook.
	AppId pulumi.StringPtrInput
	// App Secret for Facebook.
	AppSecret pulumi.StringPtrInput
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (IdentityProviderFacebookState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderFacebookState)(nil)).Elem()
}

type identityProviderFacebookArgs struct {
	// The Name of the API Management Service where this Facebook Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// App ID for Facebook.
	AppId string `pulumi:"appId"`
	// App Secret for Facebook.
	AppSecret string `pulumi:"appSecret"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a IdentityProviderFacebook resource.
type IdentityProviderFacebookArgs struct {
	// The Name of the API Management Service where this Facebook Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// App ID for Facebook.
	AppId pulumi.StringInput
	// App Secret for Facebook.
	AppSecret pulumi.StringInput
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (IdentityProviderFacebookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderFacebookArgs)(nil)).Elem()
}

type IdentityProviderFacebookInput interface {
	pulumi.Input

	ToIdentityProviderFacebookOutput() IdentityProviderFacebookOutput
	ToIdentityProviderFacebookOutputWithContext(ctx context.Context) IdentityProviderFacebookOutput
}

func (*IdentityProviderFacebook) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProviderFacebook)(nil))
}

func (i *IdentityProviderFacebook) ToIdentityProviderFacebookOutput() IdentityProviderFacebookOutput {
	return i.ToIdentityProviderFacebookOutputWithContext(context.Background())
}

func (i *IdentityProviderFacebook) ToIdentityProviderFacebookOutputWithContext(ctx context.Context) IdentityProviderFacebookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderFacebookOutput)
}

func (i *IdentityProviderFacebook) ToIdentityProviderFacebookPtrOutput() IdentityProviderFacebookPtrOutput {
	return i.ToIdentityProviderFacebookPtrOutputWithContext(context.Background())
}

func (i *IdentityProviderFacebook) ToIdentityProviderFacebookPtrOutputWithContext(ctx context.Context) IdentityProviderFacebookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderFacebookPtrOutput)
}

type IdentityProviderFacebookPtrInput interface {
	pulumi.Input

	ToIdentityProviderFacebookPtrOutput() IdentityProviderFacebookPtrOutput
	ToIdentityProviderFacebookPtrOutputWithContext(ctx context.Context) IdentityProviderFacebookPtrOutput
}

type identityProviderFacebookPtrType IdentityProviderFacebookArgs

func (*identityProviderFacebookPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProviderFacebook)(nil))
}

func (i *identityProviderFacebookPtrType) ToIdentityProviderFacebookPtrOutput() IdentityProviderFacebookPtrOutput {
	return i.ToIdentityProviderFacebookPtrOutputWithContext(context.Background())
}

func (i *identityProviderFacebookPtrType) ToIdentityProviderFacebookPtrOutputWithContext(ctx context.Context) IdentityProviderFacebookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderFacebookPtrOutput)
}

// IdentityProviderFacebookArrayInput is an input type that accepts IdentityProviderFacebookArray and IdentityProviderFacebookArrayOutput values.
// You can construct a concrete instance of `IdentityProviderFacebookArrayInput` via:
//
//          IdentityProviderFacebookArray{ IdentityProviderFacebookArgs{...} }
type IdentityProviderFacebookArrayInput interface {
	pulumi.Input

	ToIdentityProviderFacebookArrayOutput() IdentityProviderFacebookArrayOutput
	ToIdentityProviderFacebookArrayOutputWithContext(context.Context) IdentityProviderFacebookArrayOutput
}

type IdentityProviderFacebookArray []IdentityProviderFacebookInput

func (IdentityProviderFacebookArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*IdentityProviderFacebook)(nil))
}

func (i IdentityProviderFacebookArray) ToIdentityProviderFacebookArrayOutput() IdentityProviderFacebookArrayOutput {
	return i.ToIdentityProviderFacebookArrayOutputWithContext(context.Background())
}

func (i IdentityProviderFacebookArray) ToIdentityProviderFacebookArrayOutputWithContext(ctx context.Context) IdentityProviderFacebookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderFacebookArrayOutput)
}

// IdentityProviderFacebookMapInput is an input type that accepts IdentityProviderFacebookMap and IdentityProviderFacebookMapOutput values.
// You can construct a concrete instance of `IdentityProviderFacebookMapInput` via:
//
//          IdentityProviderFacebookMap{ "key": IdentityProviderFacebookArgs{...} }
type IdentityProviderFacebookMapInput interface {
	pulumi.Input

	ToIdentityProviderFacebookMapOutput() IdentityProviderFacebookMapOutput
	ToIdentityProviderFacebookMapOutputWithContext(context.Context) IdentityProviderFacebookMapOutput
}

type IdentityProviderFacebookMap map[string]IdentityProviderFacebookInput

func (IdentityProviderFacebookMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*IdentityProviderFacebook)(nil))
}

func (i IdentityProviderFacebookMap) ToIdentityProviderFacebookMapOutput() IdentityProviderFacebookMapOutput {
	return i.ToIdentityProviderFacebookMapOutputWithContext(context.Background())
}

func (i IdentityProviderFacebookMap) ToIdentityProviderFacebookMapOutputWithContext(ctx context.Context) IdentityProviderFacebookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderFacebookMapOutput)
}

type IdentityProviderFacebookOutput struct {
	*pulumi.OutputState
}

func (IdentityProviderFacebookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProviderFacebook)(nil))
}

func (o IdentityProviderFacebookOutput) ToIdentityProviderFacebookOutput() IdentityProviderFacebookOutput {
	return o
}

func (o IdentityProviderFacebookOutput) ToIdentityProviderFacebookOutputWithContext(ctx context.Context) IdentityProviderFacebookOutput {
	return o
}

func (o IdentityProviderFacebookOutput) ToIdentityProviderFacebookPtrOutput() IdentityProviderFacebookPtrOutput {
	return o.ToIdentityProviderFacebookPtrOutputWithContext(context.Background())
}

func (o IdentityProviderFacebookOutput) ToIdentityProviderFacebookPtrOutputWithContext(ctx context.Context) IdentityProviderFacebookPtrOutput {
	return o.ApplyT(func(v IdentityProviderFacebook) *IdentityProviderFacebook {
		return &v
	}).(IdentityProviderFacebookPtrOutput)
}

type IdentityProviderFacebookPtrOutput struct {
	*pulumi.OutputState
}

func (IdentityProviderFacebookPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProviderFacebook)(nil))
}

func (o IdentityProviderFacebookPtrOutput) ToIdentityProviderFacebookPtrOutput() IdentityProviderFacebookPtrOutput {
	return o
}

func (o IdentityProviderFacebookPtrOutput) ToIdentityProviderFacebookPtrOutputWithContext(ctx context.Context) IdentityProviderFacebookPtrOutput {
	return o
}

type IdentityProviderFacebookArrayOutput struct{ *pulumi.OutputState }

func (IdentityProviderFacebookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentityProviderFacebook)(nil))
}

func (o IdentityProviderFacebookArrayOutput) ToIdentityProviderFacebookArrayOutput() IdentityProviderFacebookArrayOutput {
	return o
}

func (o IdentityProviderFacebookArrayOutput) ToIdentityProviderFacebookArrayOutputWithContext(ctx context.Context) IdentityProviderFacebookArrayOutput {
	return o
}

func (o IdentityProviderFacebookArrayOutput) Index(i pulumi.IntInput) IdentityProviderFacebookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IdentityProviderFacebook {
		return vs[0].([]IdentityProviderFacebook)[vs[1].(int)]
	}).(IdentityProviderFacebookOutput)
}

type IdentityProviderFacebookMapOutput struct{ *pulumi.OutputState }

func (IdentityProviderFacebookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IdentityProviderFacebook)(nil))
}

func (o IdentityProviderFacebookMapOutput) ToIdentityProviderFacebookMapOutput() IdentityProviderFacebookMapOutput {
	return o
}

func (o IdentityProviderFacebookMapOutput) ToIdentityProviderFacebookMapOutputWithContext(ctx context.Context) IdentityProviderFacebookMapOutput {
	return o
}

func (o IdentityProviderFacebookMapOutput) MapIndex(k pulumi.StringInput) IdentityProviderFacebookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IdentityProviderFacebook {
		return vs[0].(map[string]IdentityProviderFacebook)[vs[1].(string)]
	}).(IdentityProviderFacebookOutput)
}

func init() {
	pulumi.RegisterOutputType(IdentityProviderFacebookOutput{})
	pulumi.RegisterOutputType(IdentityProviderFacebookPtrOutput{})
	pulumi.RegisterOutputType(IdentityProviderFacebookArrayOutput{})
	pulumi.RegisterOutputType(IdentityProviderFacebookMapOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an API Management Google Identity Provider.
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
// 		_, err = apimanagement.NewIdentityProviderGoogle(ctx, "exampleIdentityProviderGoogle", &apimanagement.IdentityProviderGoogleArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApiManagementName: exampleService.Name,
// 			ClientId:          pulumi.String("00000000.apps.googleusercontent.com"),
// 			ClientSecret:      pulumi.String("00000000000000000000000000000000"),
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
// API Management Google Identity Provider can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/identityProviderGoogle:IdentityProviderGoogle example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/google
// ```
type IdentityProviderGoogle struct {
	pulumi.CustomResourceState

	// The Name of the API Management Service where this Google Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// Client Id for Google Sign-in.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// Client secret for Google Sign-in.
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewIdentityProviderGoogle registers a new resource with the given unique name, arguments, and options.
func NewIdentityProviderGoogle(ctx *pulumi.Context,
	name string, args *IdentityProviderGoogleArgs, opts ...pulumi.ResourceOption) (*IdentityProviderGoogle, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiManagementName == nil {
		return nil, errors.New("invalid value for required argument 'ApiManagementName'")
	}
	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClientSecret == nil {
		return nil, errors.New("invalid value for required argument 'ClientSecret'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource IdentityProviderGoogle
	err := ctx.RegisterResource("azure:apimanagement/identityProviderGoogle:IdentityProviderGoogle", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityProviderGoogle gets an existing IdentityProviderGoogle resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityProviderGoogle(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityProviderGoogleState, opts ...pulumi.ResourceOption) (*IdentityProviderGoogle, error) {
	var resource IdentityProviderGoogle
	err := ctx.ReadResource("azure:apimanagement/identityProviderGoogle:IdentityProviderGoogle", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityProviderGoogle resources.
type identityProviderGoogleState struct {
	// The Name of the API Management Service where this Google Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// Client Id for Google Sign-in.
	ClientId *string `pulumi:"clientId"`
	// Client secret for Google Sign-in.
	ClientSecret *string `pulumi:"clientSecret"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type IdentityProviderGoogleState struct {
	// The Name of the API Management Service where this Google Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// Client Id for Google Sign-in.
	ClientId pulumi.StringPtrInput
	// Client secret for Google Sign-in.
	ClientSecret pulumi.StringPtrInput
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (IdentityProviderGoogleState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderGoogleState)(nil)).Elem()
}

type identityProviderGoogleArgs struct {
	// The Name of the API Management Service where this Google Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// Client Id for Google Sign-in.
	ClientId string `pulumi:"clientId"`
	// Client secret for Google Sign-in.
	ClientSecret string `pulumi:"clientSecret"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a IdentityProviderGoogle resource.
type IdentityProviderGoogleArgs struct {
	// The Name of the API Management Service where this Google Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// Client Id for Google Sign-in.
	ClientId pulumi.StringInput
	// Client secret for Google Sign-in.
	ClientSecret pulumi.StringInput
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (IdentityProviderGoogleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderGoogleArgs)(nil)).Elem()
}

type IdentityProviderGoogleInput interface {
	pulumi.Input

	ToIdentityProviderGoogleOutput() IdentityProviderGoogleOutput
	ToIdentityProviderGoogleOutputWithContext(ctx context.Context) IdentityProviderGoogleOutput
}

func (*IdentityProviderGoogle) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProviderGoogle)(nil))
}

func (i *IdentityProviderGoogle) ToIdentityProviderGoogleOutput() IdentityProviderGoogleOutput {
	return i.ToIdentityProviderGoogleOutputWithContext(context.Background())
}

func (i *IdentityProviderGoogle) ToIdentityProviderGoogleOutputWithContext(ctx context.Context) IdentityProviderGoogleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderGoogleOutput)
}

func (i *IdentityProviderGoogle) ToIdentityProviderGooglePtrOutput() IdentityProviderGooglePtrOutput {
	return i.ToIdentityProviderGooglePtrOutputWithContext(context.Background())
}

func (i *IdentityProviderGoogle) ToIdentityProviderGooglePtrOutputWithContext(ctx context.Context) IdentityProviderGooglePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderGooglePtrOutput)
}

type IdentityProviderGooglePtrInput interface {
	pulumi.Input

	ToIdentityProviderGooglePtrOutput() IdentityProviderGooglePtrOutput
	ToIdentityProviderGooglePtrOutputWithContext(ctx context.Context) IdentityProviderGooglePtrOutput
}

type identityProviderGooglePtrType IdentityProviderGoogleArgs

func (*identityProviderGooglePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProviderGoogle)(nil))
}

func (i *identityProviderGooglePtrType) ToIdentityProviderGooglePtrOutput() IdentityProviderGooglePtrOutput {
	return i.ToIdentityProviderGooglePtrOutputWithContext(context.Background())
}

func (i *identityProviderGooglePtrType) ToIdentityProviderGooglePtrOutputWithContext(ctx context.Context) IdentityProviderGooglePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderGooglePtrOutput)
}

// IdentityProviderGoogleArrayInput is an input type that accepts IdentityProviderGoogleArray and IdentityProviderGoogleArrayOutput values.
// You can construct a concrete instance of `IdentityProviderGoogleArrayInput` via:
//
//          IdentityProviderGoogleArray{ IdentityProviderGoogleArgs{...} }
type IdentityProviderGoogleArrayInput interface {
	pulumi.Input

	ToIdentityProviderGoogleArrayOutput() IdentityProviderGoogleArrayOutput
	ToIdentityProviderGoogleArrayOutputWithContext(context.Context) IdentityProviderGoogleArrayOutput
}

type IdentityProviderGoogleArray []IdentityProviderGoogleInput

func (IdentityProviderGoogleArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*IdentityProviderGoogle)(nil))
}

func (i IdentityProviderGoogleArray) ToIdentityProviderGoogleArrayOutput() IdentityProviderGoogleArrayOutput {
	return i.ToIdentityProviderGoogleArrayOutputWithContext(context.Background())
}

func (i IdentityProviderGoogleArray) ToIdentityProviderGoogleArrayOutputWithContext(ctx context.Context) IdentityProviderGoogleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderGoogleArrayOutput)
}

// IdentityProviderGoogleMapInput is an input type that accepts IdentityProviderGoogleMap and IdentityProviderGoogleMapOutput values.
// You can construct a concrete instance of `IdentityProviderGoogleMapInput` via:
//
//          IdentityProviderGoogleMap{ "key": IdentityProviderGoogleArgs{...} }
type IdentityProviderGoogleMapInput interface {
	pulumi.Input

	ToIdentityProviderGoogleMapOutput() IdentityProviderGoogleMapOutput
	ToIdentityProviderGoogleMapOutputWithContext(context.Context) IdentityProviderGoogleMapOutput
}

type IdentityProviderGoogleMap map[string]IdentityProviderGoogleInput

func (IdentityProviderGoogleMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*IdentityProviderGoogle)(nil))
}

func (i IdentityProviderGoogleMap) ToIdentityProviderGoogleMapOutput() IdentityProviderGoogleMapOutput {
	return i.ToIdentityProviderGoogleMapOutputWithContext(context.Background())
}

func (i IdentityProviderGoogleMap) ToIdentityProviderGoogleMapOutputWithContext(ctx context.Context) IdentityProviderGoogleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderGoogleMapOutput)
}

type IdentityProviderGoogleOutput struct {
	*pulumi.OutputState
}

func (IdentityProviderGoogleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProviderGoogle)(nil))
}

func (o IdentityProviderGoogleOutput) ToIdentityProviderGoogleOutput() IdentityProviderGoogleOutput {
	return o
}

func (o IdentityProviderGoogleOutput) ToIdentityProviderGoogleOutputWithContext(ctx context.Context) IdentityProviderGoogleOutput {
	return o
}

func (o IdentityProviderGoogleOutput) ToIdentityProviderGooglePtrOutput() IdentityProviderGooglePtrOutput {
	return o.ToIdentityProviderGooglePtrOutputWithContext(context.Background())
}

func (o IdentityProviderGoogleOutput) ToIdentityProviderGooglePtrOutputWithContext(ctx context.Context) IdentityProviderGooglePtrOutput {
	return o.ApplyT(func(v IdentityProviderGoogle) *IdentityProviderGoogle {
		return &v
	}).(IdentityProviderGooglePtrOutput)
}

type IdentityProviderGooglePtrOutput struct {
	*pulumi.OutputState
}

func (IdentityProviderGooglePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProviderGoogle)(nil))
}

func (o IdentityProviderGooglePtrOutput) ToIdentityProviderGooglePtrOutput() IdentityProviderGooglePtrOutput {
	return o
}

func (o IdentityProviderGooglePtrOutput) ToIdentityProviderGooglePtrOutputWithContext(ctx context.Context) IdentityProviderGooglePtrOutput {
	return o
}

type IdentityProviderGoogleArrayOutput struct{ *pulumi.OutputState }

func (IdentityProviderGoogleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentityProviderGoogle)(nil))
}

func (o IdentityProviderGoogleArrayOutput) ToIdentityProviderGoogleArrayOutput() IdentityProviderGoogleArrayOutput {
	return o
}

func (o IdentityProviderGoogleArrayOutput) ToIdentityProviderGoogleArrayOutputWithContext(ctx context.Context) IdentityProviderGoogleArrayOutput {
	return o
}

func (o IdentityProviderGoogleArrayOutput) Index(i pulumi.IntInput) IdentityProviderGoogleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IdentityProviderGoogle {
		return vs[0].([]IdentityProviderGoogle)[vs[1].(int)]
	}).(IdentityProviderGoogleOutput)
}

type IdentityProviderGoogleMapOutput struct{ *pulumi.OutputState }

func (IdentityProviderGoogleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IdentityProviderGoogle)(nil))
}

func (o IdentityProviderGoogleMapOutput) ToIdentityProviderGoogleMapOutput() IdentityProviderGoogleMapOutput {
	return o
}

func (o IdentityProviderGoogleMapOutput) ToIdentityProviderGoogleMapOutputWithContext(ctx context.Context) IdentityProviderGoogleMapOutput {
	return o
}

func (o IdentityProviderGoogleMapOutput) MapIndex(k pulumi.StringInput) IdentityProviderGoogleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IdentityProviderGoogle {
		return vs[0].(map[string]IdentityProviderGoogle)[vs[1].(string)]
	}).(IdentityProviderGoogleOutput)
}

func init() {
	pulumi.RegisterOutputType(IdentityProviderGoogleOutput{})
	pulumi.RegisterOutputType(IdentityProviderGooglePtrOutput{})
	pulumi.RegisterOutputType(IdentityProviderGoogleArrayOutput{})
	pulumi.RegisterOutputType(IdentityProviderGoogleMapOutput{})
}

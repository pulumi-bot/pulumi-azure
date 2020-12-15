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

func init() {
	pulumi.RegisterOutputType(IdentityProviderGoogleOutput{})
	pulumi.RegisterOutputType(IdentityProviderGooglePtrOutput{})
}

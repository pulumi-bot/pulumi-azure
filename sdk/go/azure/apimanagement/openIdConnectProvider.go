// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an OpenID Connect Provider within a API Management Service.
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
// 		_, err = apimanagement.NewOpenIdConnectProvider(ctx, "exampleOpenIdConnectProvider", &apimanagement.OpenIdConnectProviderArgs{
// 			ApiManagementName: exampleService.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ClientId:          pulumi.String("00001111-2222-3333-4444-555566667777"),
// 			DisplayName:       pulumi.String("Example Provider"),
// 			MetadataEndpoint:  pulumi.String("https://example.com/example"),
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
// API Management OpenID Connect Providers can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/openIdConnectProvider:OpenIdConnectProvider example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/openidConnectProviders/provider1
// ```
type OpenIdConnectProvider struct {
	pulumi.CustomResourceState

	// The name of the API Management Service in which this OpenID Connect Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The Client ID used for the Client Application.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The Client Secret used for the Client Application.
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// A description of this OpenID Connect Provider.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A user-friendly name for this OpenID Connect Provider.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The URI of the Metadata endpoint.
	MetadataEndpoint pulumi.StringOutput `pulumi:"metadataEndpoint"`
	// the Name of the OpenID Connect Provider which should be created within the API Management Service. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewOpenIdConnectProvider registers a new resource with the given unique name, arguments, and options.
func NewOpenIdConnectProvider(ctx *pulumi.Context,
	name string, args *OpenIdConnectProviderArgs, opts ...pulumi.ResourceOption) (*OpenIdConnectProvider, error) {
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
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.MetadataEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'MetadataEndpoint'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource OpenIdConnectProvider
	err := ctx.RegisterResource("azure:apimanagement/openIdConnectProvider:OpenIdConnectProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOpenIdConnectProvider gets an existing OpenIdConnectProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOpenIdConnectProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OpenIdConnectProviderState, opts ...pulumi.ResourceOption) (*OpenIdConnectProvider, error) {
	var resource OpenIdConnectProvider
	err := ctx.ReadResource("azure:apimanagement/openIdConnectProvider:OpenIdConnectProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OpenIdConnectProvider resources.
type openIdConnectProviderState struct {
	// The name of the API Management Service in which this OpenID Connect Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The Client ID used for the Client Application.
	ClientId *string `pulumi:"clientId"`
	// The Client Secret used for the Client Application.
	ClientSecret *string `pulumi:"clientSecret"`
	// A description of this OpenID Connect Provider.
	Description *string `pulumi:"description"`
	// A user-friendly name for this OpenID Connect Provider.
	DisplayName *string `pulumi:"displayName"`
	// The URI of the Metadata endpoint.
	MetadataEndpoint *string `pulumi:"metadataEndpoint"`
	// the Name of the OpenID Connect Provider which should be created within the API Management Service. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type OpenIdConnectProviderState struct {
	// The name of the API Management Service in which this OpenID Connect Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The Client ID used for the Client Application.
	ClientId pulumi.StringPtrInput
	// The Client Secret used for the Client Application.
	ClientSecret pulumi.StringPtrInput
	// A description of this OpenID Connect Provider.
	Description pulumi.StringPtrInput
	// A user-friendly name for this OpenID Connect Provider.
	DisplayName pulumi.StringPtrInput
	// The URI of the Metadata endpoint.
	MetadataEndpoint pulumi.StringPtrInput
	// the Name of the OpenID Connect Provider which should be created within the API Management Service. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (OpenIdConnectProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*openIdConnectProviderState)(nil)).Elem()
}

type openIdConnectProviderArgs struct {
	// The name of the API Management Service in which this OpenID Connect Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The Client ID used for the Client Application.
	ClientId string `pulumi:"clientId"`
	// The Client Secret used for the Client Application.
	ClientSecret string `pulumi:"clientSecret"`
	// A description of this OpenID Connect Provider.
	Description *string `pulumi:"description"`
	// A user-friendly name for this OpenID Connect Provider.
	DisplayName string `pulumi:"displayName"`
	// The URI of the Metadata endpoint.
	MetadataEndpoint string `pulumi:"metadataEndpoint"`
	// the Name of the OpenID Connect Provider which should be created within the API Management Service. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a OpenIdConnectProvider resource.
type OpenIdConnectProviderArgs struct {
	// The name of the API Management Service in which this OpenID Connect Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The Client ID used for the Client Application.
	ClientId pulumi.StringInput
	// The Client Secret used for the Client Application.
	ClientSecret pulumi.StringInput
	// A description of this OpenID Connect Provider.
	Description pulumi.StringPtrInput
	// A user-friendly name for this OpenID Connect Provider.
	DisplayName pulumi.StringInput
	// The URI of the Metadata endpoint.
	MetadataEndpoint pulumi.StringInput
	// the Name of the OpenID Connect Provider which should be created within the API Management Service. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (OpenIdConnectProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*openIdConnectProviderArgs)(nil)).Elem()
}

type OpenIdConnectProviderInput interface {
	pulumi.Input

	ToOpenIdConnectProviderOutput() OpenIdConnectProviderOutput
	ToOpenIdConnectProviderOutputWithContext(ctx context.Context) OpenIdConnectProviderOutput
}

func (OpenIdConnectProvider) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectProvider)(nil)).Elem()
}

func (i OpenIdConnectProvider) ToOpenIdConnectProviderOutput() OpenIdConnectProviderOutput {
	return i.ToOpenIdConnectProviderOutputWithContext(context.Background())
}

func (i OpenIdConnectProvider) ToOpenIdConnectProviderOutputWithContext(ctx context.Context) OpenIdConnectProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectProviderOutput)
}

type OpenIdConnectProviderOutput struct {
	*pulumi.OutputState
}

func (OpenIdConnectProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectProviderOutput)(nil)).Elem()
}

func (o OpenIdConnectProviderOutput) ToOpenIdConnectProviderOutput() OpenIdConnectProviderOutput {
	return o
}

func (o OpenIdConnectProviderOutput) ToOpenIdConnectProviderOutputWithContext(ctx context.Context) OpenIdConnectProviderOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OpenIdConnectProviderOutput{})
}

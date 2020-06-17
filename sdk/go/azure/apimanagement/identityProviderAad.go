// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an API Management AAD Identity Provider.
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
// 		_, err = apimanagement.NewIdentityProviderAad(ctx, "exampleIdentityProviderAad", &apimanagement.IdentityProviderAadArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApiManagementName: exampleService.Name,
// 			ClientId:          pulumi.String("00000000-0000-0000-0000-000000000000"),
// 			ClientSecret:      pulumi.String("00000000000000000000000000000000"),
// 			AllowedTenants: pulumi.StringArray{
// 				pulumi.String("00000000-0000-0000-0000-000000000000"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type IdentityProviderAad struct {
	pulumi.CustomResourceState

	// List of allowed AAD Tenants.
	AllowedTenants pulumi.StringArrayOutput `pulumi:"allowedTenants"`
	// The Name of the API Management Service where this AAD Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// Client Id of the Application in the AAD Identity Provider.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// Client secret of the Application in the AAD Identity Provider.
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewIdentityProviderAad registers a new resource with the given unique name, arguments, and options.
func NewIdentityProviderAad(ctx *pulumi.Context,
	name string, args *IdentityProviderAadArgs, opts ...pulumi.ResourceOption) (*IdentityProviderAad, error) {
	if args == nil || args.AllowedTenants == nil {
		return nil, errors.New("missing required argument 'AllowedTenants'")
	}
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.ClientId == nil {
		return nil, errors.New("missing required argument 'ClientId'")
	}
	if args == nil || args.ClientSecret == nil {
		return nil, errors.New("missing required argument 'ClientSecret'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &IdentityProviderAadArgs{}
	}
	var resource IdentityProviderAad
	err := ctx.RegisterResource("azure:apimanagement/identityProviderAad:IdentityProviderAad", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityProviderAad gets an existing IdentityProviderAad resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityProviderAad(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityProviderAadState, opts ...pulumi.ResourceOption) (*IdentityProviderAad, error) {
	var resource IdentityProviderAad
	err := ctx.ReadResource("azure:apimanagement/identityProviderAad:IdentityProviderAad", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityProviderAad resources.
type identityProviderAadState struct {
	// List of allowed AAD Tenants.
	AllowedTenants []string `pulumi:"allowedTenants"`
	// The Name of the API Management Service where this AAD Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// Client Id of the Application in the AAD Identity Provider.
	ClientId *string `pulumi:"clientId"`
	// Client secret of the Application in the AAD Identity Provider.
	ClientSecret *string `pulumi:"clientSecret"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type IdentityProviderAadState struct {
	// List of allowed AAD Tenants.
	AllowedTenants pulumi.StringArrayInput
	// The Name of the API Management Service where this AAD Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// Client Id of the Application in the AAD Identity Provider.
	ClientId pulumi.StringPtrInput
	// Client secret of the Application in the AAD Identity Provider.
	ClientSecret pulumi.StringPtrInput
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (IdentityProviderAadState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderAadState)(nil)).Elem()
}

type identityProviderAadArgs struct {
	// List of allowed AAD Tenants.
	AllowedTenants []string `pulumi:"allowedTenants"`
	// The Name of the API Management Service where this AAD Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// Client Id of the Application in the AAD Identity Provider.
	ClientId string `pulumi:"clientId"`
	// Client secret of the Application in the AAD Identity Provider.
	ClientSecret string `pulumi:"clientSecret"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a IdentityProviderAad resource.
type IdentityProviderAadArgs struct {
	// List of allowed AAD Tenants.
	AllowedTenants pulumi.StringArrayInput
	// The Name of the API Management Service where this AAD Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// Client Id of the Application in the AAD Identity Provider.
	ClientId pulumi.StringInput
	// Client secret of the Application in the AAD Identity Provider.
	ClientSecret pulumi.StringInput
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (IdentityProviderAadArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderAadArgs)(nil)).Elem()
}

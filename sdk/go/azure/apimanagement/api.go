// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an API within an API Management Service.
//
// ## Example Usage
//
//
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
// 		_, err = apimanagement.NewApi(ctx, "exampleApi", &apimanagement.ApiArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApiManagementName: exampleService.Name,
// 			Revision:          pulumi.String("1"),
// 			DisplayName:       pulumi.String("Example API"),
// 			Path:              pulumi.String("example"),
// 			Protocols: pulumi.StringArray{
// 				pulumi.String("https"),
// 			},
// 			Import: &apimanagement.ApiImportArgs{
// 				ContentFormat: pulumi.String("swagger-link-json"),
// 				ContentValue:  pulumi.String("http://conferenceapi.azurewebsites.net/?format=json"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Api struct {
	pulumi.CustomResourceState

	// The Name of the API Management Service where this API should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// A description of the API Management API, which may include HTML formatting tags.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the API.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// A `import` block as documented below.
	Import ApiImportPtrOutput `pulumi:"import"`
	// Is this the current API Revision?
	IsCurrent pulumi.BoolOutput `pulumi:"isCurrent"`
	// Is this API Revision online/accessible via the Gateway?
	IsOnline pulumi.BoolOutput `pulumi:"isOnline"`
	// The name of the API Management API. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Path for this API Management API, which is a relative URL which uniquely identifies this API and all of it's resource paths within the API Management Service.
	Path pulumi.StringOutput `pulumi:"path"`
	// A list of protocols the operations in this API can be invoked. Possible values are `http` and `https`.
	Protocols pulumi.StringArrayOutput `pulumi:"protocols"`
	// The Name of the Resource Group where the API Management API exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Revision which used for this API.
	Revision pulumi.StringOutput `pulumi:"revision"`
	// Absolute URL of the backend service implementing this API.
	ServiceUrl pulumi.StringOutput `pulumi:"serviceUrl"`
	// Should this API expose a SOAP frontend, rather than a HTTP frontend? Defaults to `false`.
	SoapPassThrough pulumi.BoolPtrOutput `pulumi:"soapPassThrough"`
	// A `subscriptionKeyParameterNames` block as documented below.
	SubscriptionKeyParameterNames ApiSubscriptionKeyParameterNamesOutput `pulumi:"subscriptionKeyParameterNames"`
	// The Version number of this API, if this API is versioned.
	Version pulumi.StringOutput `pulumi:"version"`
	// The ID of the Version Set which this API is associated with.
	VersionSetId pulumi.StringOutput `pulumi:"versionSetId"`
}

// NewApi registers a new resource with the given unique name, arguments, and options.
func NewApi(ctx *pulumi.Context,
	name string, args *ApiArgs, opts ...pulumi.ResourceOption) (*Api, error) {
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.Path == nil {
		return nil, errors.New("missing required argument 'Path'")
	}
	if args == nil || args.Protocols == nil {
		return nil, errors.New("missing required argument 'Protocols'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Revision == nil {
		return nil, errors.New("missing required argument 'Revision'")
	}
	if args == nil {
		args = &ApiArgs{}
	}
	var resource Api
	err := ctx.RegisterResource("azure:apimanagement/api:Api", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApi gets an existing Api resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiState, opts ...pulumi.ResourceOption) (*Api, error) {
	var resource Api
	err := ctx.ReadResource("azure:apimanagement/api:Api", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Api resources.
type apiState struct {
	// The Name of the API Management Service where this API should be created. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// A description of the API Management API, which may include HTML formatting tags.
	Description *string `pulumi:"description"`
	// The display name of the API.
	DisplayName *string `pulumi:"displayName"`
	// A `import` block as documented below.
	Import *ApiImport `pulumi:"import"`
	// Is this the current API Revision?
	IsCurrent *bool `pulumi:"isCurrent"`
	// Is this API Revision online/accessible via the Gateway?
	IsOnline *bool `pulumi:"isOnline"`
	// The name of the API Management API. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The Path for this API Management API, which is a relative URL which uniquely identifies this API and all of it's resource paths within the API Management Service.
	Path *string `pulumi:"path"`
	// A list of protocols the operations in this API can be invoked. Possible values are `http` and `https`.
	Protocols []string `pulumi:"protocols"`
	// The Name of the Resource Group where the API Management API exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Revision which used for this API.
	Revision *string `pulumi:"revision"`
	// Absolute URL of the backend service implementing this API.
	ServiceUrl *string `pulumi:"serviceUrl"`
	// Should this API expose a SOAP frontend, rather than a HTTP frontend? Defaults to `false`.
	SoapPassThrough *bool `pulumi:"soapPassThrough"`
	// A `subscriptionKeyParameterNames` block as documented below.
	SubscriptionKeyParameterNames *ApiSubscriptionKeyParameterNames `pulumi:"subscriptionKeyParameterNames"`
	// The Version number of this API, if this API is versioned.
	Version *string `pulumi:"version"`
	// The ID of the Version Set which this API is associated with.
	VersionSetId *string `pulumi:"versionSetId"`
}

type ApiState struct {
	// The Name of the API Management Service where this API should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// A description of the API Management API, which may include HTML formatting tags.
	Description pulumi.StringPtrInput
	// The display name of the API.
	DisplayName pulumi.StringPtrInput
	// A `import` block as documented below.
	Import ApiImportPtrInput
	// Is this the current API Revision?
	IsCurrent pulumi.BoolPtrInput
	// Is this API Revision online/accessible via the Gateway?
	IsOnline pulumi.BoolPtrInput
	// The name of the API Management API. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The Path for this API Management API, which is a relative URL which uniquely identifies this API and all of it's resource paths within the API Management Service.
	Path pulumi.StringPtrInput
	// A list of protocols the operations in this API can be invoked. Possible values are `http` and `https`.
	Protocols pulumi.StringArrayInput
	// The Name of the Resource Group where the API Management API exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Revision which used for this API.
	Revision pulumi.StringPtrInput
	// Absolute URL of the backend service implementing this API.
	ServiceUrl pulumi.StringPtrInput
	// Should this API expose a SOAP frontend, rather than a HTTP frontend? Defaults to `false`.
	SoapPassThrough pulumi.BoolPtrInput
	// A `subscriptionKeyParameterNames` block as documented below.
	SubscriptionKeyParameterNames ApiSubscriptionKeyParameterNamesPtrInput
	// The Version number of this API, if this API is versioned.
	Version pulumi.StringPtrInput
	// The ID of the Version Set which this API is associated with.
	VersionSetId pulumi.StringPtrInput
}

func (ApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiState)(nil)).Elem()
}

type apiArgs struct {
	// The Name of the API Management Service where this API should be created. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// A description of the API Management API, which may include HTML formatting tags.
	Description *string `pulumi:"description"`
	// The display name of the API.
	DisplayName string `pulumi:"displayName"`
	// A `import` block as documented below.
	Import *ApiImport `pulumi:"import"`
	// The name of the API Management API. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The Path for this API Management API, which is a relative URL which uniquely identifies this API and all of it's resource paths within the API Management Service.
	Path string `pulumi:"path"`
	// A list of protocols the operations in this API can be invoked. Possible values are `http` and `https`.
	Protocols []string `pulumi:"protocols"`
	// The Name of the Resource Group where the API Management API exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Revision which used for this API.
	Revision string `pulumi:"revision"`
	// Absolute URL of the backend service implementing this API.
	ServiceUrl *string `pulumi:"serviceUrl"`
	// Should this API expose a SOAP frontend, rather than a HTTP frontend? Defaults to `false`.
	SoapPassThrough *bool `pulumi:"soapPassThrough"`
	// A `subscriptionKeyParameterNames` block as documented below.
	SubscriptionKeyParameterNames *ApiSubscriptionKeyParameterNames `pulumi:"subscriptionKeyParameterNames"`
	// The Version number of this API, if this API is versioned.
	Version *string `pulumi:"version"`
	// The ID of the Version Set which this API is associated with.
	VersionSetId *string `pulumi:"versionSetId"`
}

// The set of arguments for constructing a Api resource.
type ApiArgs struct {
	// The Name of the API Management Service where this API should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// A description of the API Management API, which may include HTML formatting tags.
	Description pulumi.StringPtrInput
	// The display name of the API.
	DisplayName pulumi.StringInput
	// A `import` block as documented below.
	Import ApiImportPtrInput
	// The name of the API Management API. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The Path for this API Management API, which is a relative URL which uniquely identifies this API and all of it's resource paths within the API Management Service.
	Path pulumi.StringInput
	// A list of protocols the operations in this API can be invoked. Possible values are `http` and `https`.
	Protocols pulumi.StringArrayInput
	// The Name of the Resource Group where the API Management API exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The Revision which used for this API.
	Revision pulumi.StringInput
	// Absolute URL of the backend service implementing this API.
	ServiceUrl pulumi.StringPtrInput
	// Should this API expose a SOAP frontend, rather than a HTTP frontend? Defaults to `false`.
	SoapPassThrough pulumi.BoolPtrInput
	// A `subscriptionKeyParameterNames` block as documented below.
	SubscriptionKeyParameterNames ApiSubscriptionKeyParameterNamesPtrInput
	// The Version number of this API, if this API is versioned.
	Version pulumi.StringPtrInput
	// The ID of the Version Set which this API is associated with.
	VersionSetId pulumi.StringPtrInput
}

func (ApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiArgs)(nil)).Elem()
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an API Version Set within an API Management Service.
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
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleService, err := apimanagement.NewService(ctx, "exampleService", &apimanagement.ServiceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			PublisherName:     pulumi.String("pub1"),
// 			PublisherEmail:    pulumi.String("pub1@email.com"),
// 			SkuName:           pulumi.String("Developer_1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apimanagement.NewApiVersionSet(ctx, "exampleApiVersionSet", &apimanagement.ApiVersionSetArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApiManagementName: exampleService.Name,
// 			DisplayName:       pulumi.String("ExampleAPIVersionSet"),
// 			VersioningScheme:  pulumi.String("Segment"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ApiVersionSet struct {
	pulumi.CustomResourceState

	// The name of the API Management Service in which the API Version Set should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The description of API Version Set.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of this API Version Set.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The name of the API Version Set. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group in which the parent API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The name of the Header which should be read from Inbound Requests which defines the API Version.
	VersionHeaderName pulumi.StringPtrOutput `pulumi:"versionHeaderName"`
	// The name of the Query String which should be read from Inbound Requests which defines the API Version.
	VersionQueryName pulumi.StringPtrOutput `pulumi:"versionQueryName"`
	// Specifies where in an Inbound HTTP Request that the API Version should be read from. Possible values are `Header`, `Query` and `Segment`.
	VersioningScheme pulumi.StringOutput `pulumi:"versioningScheme"`
}

// NewApiVersionSet registers a new resource with the given unique name, arguments, and options.
func NewApiVersionSet(ctx *pulumi.Context,
	name string, args *ApiVersionSetArgs, opts ...pulumi.ResourceOption) (*ApiVersionSet, error) {
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VersioningScheme == nil {
		return nil, errors.New("missing required argument 'VersioningScheme'")
	}
	if args == nil {
		args = &ApiVersionSetArgs{}
	}
	var resource ApiVersionSet
	err := ctx.RegisterResource("azure:apimanagement/apiVersionSet:ApiVersionSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiVersionSet gets an existing ApiVersionSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiVersionSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiVersionSetState, opts ...pulumi.ResourceOption) (*ApiVersionSet, error) {
	var resource ApiVersionSet
	err := ctx.ReadResource("azure:apimanagement/apiVersionSet:ApiVersionSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiVersionSet resources.
type apiVersionSetState struct {
	// The name of the API Management Service in which the API Version Set should exist. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The description of API Version Set.
	Description *string `pulumi:"description"`
	// The display name of this API Version Set.
	DisplayName *string `pulumi:"displayName"`
	// The name of the API Version Set. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the parent API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The name of the Header which should be read from Inbound Requests which defines the API Version.
	VersionHeaderName *string `pulumi:"versionHeaderName"`
	// The name of the Query String which should be read from Inbound Requests which defines the API Version.
	VersionQueryName *string `pulumi:"versionQueryName"`
	// Specifies where in an Inbound HTTP Request that the API Version should be read from. Possible values are `Header`, `Query` and `Segment`.
	VersioningScheme *string `pulumi:"versioningScheme"`
}

type ApiVersionSetState struct {
	// The name of the API Management Service in which the API Version Set should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The description of API Version Set.
	Description pulumi.StringPtrInput
	// The display name of this API Version Set.
	DisplayName pulumi.StringPtrInput
	// The name of the API Version Set. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the parent API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The name of the Header which should be read from Inbound Requests which defines the API Version.
	VersionHeaderName pulumi.StringPtrInput
	// The name of the Query String which should be read from Inbound Requests which defines the API Version.
	VersionQueryName pulumi.StringPtrInput
	// Specifies where in an Inbound HTTP Request that the API Version should be read from. Possible values are `Header`, `Query` and `Segment`.
	VersioningScheme pulumi.StringPtrInput
}

func (ApiVersionSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiVersionSetState)(nil)).Elem()
}

type apiVersionSetArgs struct {
	// The name of the API Management Service in which the API Version Set should exist. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The description of API Version Set.
	Description *string `pulumi:"description"`
	// The display name of this API Version Set.
	DisplayName string `pulumi:"displayName"`
	// The name of the API Version Set. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the parent API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Header which should be read from Inbound Requests which defines the API Version.
	VersionHeaderName *string `pulumi:"versionHeaderName"`
	// The name of the Query String which should be read from Inbound Requests which defines the API Version.
	VersionQueryName *string `pulumi:"versionQueryName"`
	// Specifies where in an Inbound HTTP Request that the API Version should be read from. Possible values are `Header`, `Query` and `Segment`.
	VersioningScheme string `pulumi:"versioningScheme"`
}

// The set of arguments for constructing a ApiVersionSet resource.
type ApiVersionSetArgs struct {
	// The name of the API Management Service in which the API Version Set should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The description of API Version Set.
	Description pulumi.StringPtrInput
	// The display name of this API Version Set.
	DisplayName pulumi.StringInput
	// The name of the API Version Set. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the parent API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The name of the Header which should be read from Inbound Requests which defines the API Version.
	VersionHeaderName pulumi.StringPtrInput
	// The name of the Query String which should be read from Inbound Requests which defines the API Version.
	VersionQueryName pulumi.StringPtrInput
	// Specifies where in an Inbound HTTP Request that the API Version should be read from. Possible values are `Header`, `Query` and `Segment`.
	VersioningScheme pulumi.StringInput
}

func (ApiVersionSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiVersionSetArgs)(nil)).Elem()
}

type ApiVersionSetInput interface {
	pulumi.Input

	ToApiVersionSetOutput() ApiVersionSetOutput
	ToApiVersionSetOutputWithContext(ctx context.Context) ApiVersionSetOutput
}

func (ApiVersionSet) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiVersionSet)(nil)).Elem()
}

func (i ApiVersionSet) ToApiVersionSetOutput() ApiVersionSetOutput {
	return i.ToApiVersionSetOutputWithContext(context.Background())
}

func (i ApiVersionSet) ToApiVersionSetOutputWithContext(ctx context.Context) ApiVersionSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiVersionSetOutput)
}

type ApiVersionSetOutput struct {
	*pulumi.OutputState
}

func (ApiVersionSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiVersionSetOutput)(nil)).Elem()
}

func (o ApiVersionSetOutput) ToApiVersionSetOutput() ApiVersionSetOutput {
	return o
}

func (o ApiVersionSetOutput) ToApiVersionSetOutputWithContext(ctx context.Context) ApiVersionSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApiVersionSetOutput{})
}

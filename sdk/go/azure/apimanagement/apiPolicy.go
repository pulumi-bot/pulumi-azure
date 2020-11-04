// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an API Management API Policy
type ApiPolicy struct {
	pulumi.CustomResourceState

	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The ID of the API Management API within the API Management Service. Changing this forces a new resource to be created.
	ApiName pulumi.StringOutput `pulumi:"apiName"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The XML Content for this Policy as a string.
	XmlContent pulumi.StringOutput `pulumi:"xmlContent"`
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink pulumi.StringPtrOutput `pulumi:"xmlLink"`
}

// NewApiPolicy registers a new resource with the given unique name, arguments, and options.
func NewApiPolicy(ctx *pulumi.Context,
	name string, args *ApiPolicyArgs, opts ...pulumi.ResourceOption) (*ApiPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.ApiManagementName == nil {
		return nil, errors.New("invalid value for required argument 'ApiManagementName'")
	}
	if args.ApiName == nil {
		return nil, errors.New("invalid value for required argument 'ApiName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ApiPolicy
	err := ctx.RegisterResource("azure:apimanagement/apiPolicy:ApiPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiPolicy gets an existing ApiPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiPolicyState, opts ...pulumi.ResourceOption) (*ApiPolicy, error) {
	var resource ApiPolicy
	err := ctx.ReadResource("azure:apimanagement/apiPolicy:ApiPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiPolicy resources.
type apiPolicyState struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The ID of the API Management API within the API Management Service. Changing this forces a new resource to be created.
	ApiName *string `pulumi:"apiName"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The XML Content for this Policy as a string.
	XmlContent *string `pulumi:"xmlContent"`
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink *string `pulumi:"xmlLink"`
}

type ApiPolicyState struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The ID of the API Management API within the API Management Service. Changing this forces a new resource to be created.
	ApiName pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The XML Content for this Policy as a string.
	XmlContent pulumi.StringPtrInput
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink pulumi.StringPtrInput
}

func (ApiPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiPolicyState)(nil)).Elem()
}

type apiPolicyArgs struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The ID of the API Management API within the API Management Service. Changing this forces a new resource to be created.
	ApiName string `pulumi:"apiName"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The XML Content for this Policy as a string.
	XmlContent *string `pulumi:"xmlContent"`
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink *string `pulumi:"xmlLink"`
}

// The set of arguments for constructing a ApiPolicy resource.
type ApiPolicyArgs struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The ID of the API Management API within the API Management Service. Changing this forces a new resource to be created.
	ApiName pulumi.StringInput
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The XML Content for this Policy as a string.
	XmlContent pulumi.StringPtrInput
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink pulumi.StringPtrInput
}

func (ApiPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiPolicyArgs)(nil)).Elem()
}

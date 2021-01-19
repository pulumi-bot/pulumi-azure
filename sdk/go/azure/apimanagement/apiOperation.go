// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an API Operation within an API Management Service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/apimanagement"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleApi, err := apimanagement.LookupApi(ctx, &apimanagement.LookupApiArgs{
// 			Name:              "search-api",
// 			ApiManagementName: "search-api-management",
// 			ResourceGroupName: "search-service",
// 			Revision:          "2",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apimanagement.NewApiOperation(ctx, "exampleApiOperation", &apimanagement.ApiOperationArgs{
// 			OperationId:       pulumi.String("user-delete"),
// 			ApiName:           pulumi.String(exampleApi.Name),
// 			ApiManagementName: pulumi.String(exampleApi.ApiManagementName),
// 			ResourceGroupName: pulumi.String(exampleApi.ResourceGroupName),
// 			DisplayName:       pulumi.String("Delete User Operation"),
// 			Method:            pulumi.String("DELETE"),
// 			UrlTemplate:       pulumi.String("/users/{id}/delete"),
// 			Description:       pulumi.String("This can only be done by the logged in user."),
// 			Responses: apimanagement.ApiOperationResponseArray{
// 				&apimanagement.ApiOperationResponseArgs{
// 					StatusCode: pulumi.Int(200),
// 				},
// 			},
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
// API Management API Operation's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/apiOperation:ApiOperation example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/apis/api1/operations/operation1
// ```
type ApiOperation struct {
	pulumi.CustomResourceState

	// The Name of the API Management Service where the API exists. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The name of the API within the API Management Service where this API Operation should be created. Changing this forces a new resource to be created.
	ApiName pulumi.StringOutput `pulumi:"apiName"`
	// A description for this API Operation, which may include HTML formatting tags.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Display Name for this API Management Operation.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The HTTP Method used for this API Management Operation, like `GET`, `DELETE`, `PUT` or `POST` - but not limited to these values.
	Method pulumi.StringOutput `pulumi:"method"`
	// A unique identifier for this API Operation. Changing this forces a new resource to be created.
	OperationId pulumi.StringOutput `pulumi:"operationId"`
	// A `request` block as defined below.
	Request ApiOperationRequestOutput `pulumi:"request"`
	// The Name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// One or more `response` blocks as defined below.
	Responses ApiOperationResponseArrayOutput `pulumi:"responses"`
	// One or more `templateParameter` blocks as defined below.
	TemplateParameters ApiOperationTemplateParameterArrayOutput `pulumi:"templateParameters"`
	// The relative URL Template identifying the target resource for this operation, which may include parameters.
	UrlTemplate pulumi.StringOutput `pulumi:"urlTemplate"`
}

// NewApiOperation registers a new resource with the given unique name, arguments, and options.
func NewApiOperation(ctx *pulumi.Context,
	name string, args *ApiOperationArgs, opts ...pulumi.ResourceOption) (*ApiOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiManagementName == nil {
		return nil, errors.New("invalid value for required argument 'ApiManagementName'")
	}
	if args.ApiName == nil {
		return nil, errors.New("invalid value for required argument 'ApiName'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Method == nil {
		return nil, errors.New("invalid value for required argument 'Method'")
	}
	if args.OperationId == nil {
		return nil, errors.New("invalid value for required argument 'OperationId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.UrlTemplate == nil {
		return nil, errors.New("invalid value for required argument 'UrlTemplate'")
	}
	var resource ApiOperation
	err := ctx.RegisterResource("azure:apimanagement/apiOperation:ApiOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiOperation gets an existing ApiOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiOperationState, opts ...pulumi.ResourceOption) (*ApiOperation, error) {
	var resource ApiOperation
	err := ctx.ReadResource("azure:apimanagement/apiOperation:ApiOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiOperation resources.
type apiOperationState struct {
	// The Name of the API Management Service where the API exists. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The name of the API within the API Management Service where this API Operation should be created. Changing this forces a new resource to be created.
	ApiName *string `pulumi:"apiName"`
	// A description for this API Operation, which may include HTML formatting tags.
	Description *string `pulumi:"description"`
	// The Display Name for this API Management Operation.
	DisplayName *string `pulumi:"displayName"`
	// The HTTP Method used for this API Management Operation, like `GET`, `DELETE`, `PUT` or `POST` - but not limited to these values.
	Method *string `pulumi:"method"`
	// A unique identifier for this API Operation. Changing this forces a new resource to be created.
	OperationId *string `pulumi:"operationId"`
	// A `request` block as defined below.
	Request *ApiOperationRequest `pulumi:"request"`
	// The Name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// One or more `response` blocks as defined below.
	Responses []ApiOperationResponse `pulumi:"responses"`
	// One or more `templateParameter` blocks as defined below.
	TemplateParameters []ApiOperationTemplateParameter `pulumi:"templateParameters"`
	// The relative URL Template identifying the target resource for this operation, which may include parameters.
	UrlTemplate *string `pulumi:"urlTemplate"`
}

type ApiOperationState struct {
	// The Name of the API Management Service where the API exists. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The name of the API within the API Management Service where this API Operation should be created. Changing this forces a new resource to be created.
	ApiName pulumi.StringPtrInput
	// A description for this API Operation, which may include HTML formatting tags.
	Description pulumi.StringPtrInput
	// The Display Name for this API Management Operation.
	DisplayName pulumi.StringPtrInput
	// The HTTP Method used for this API Management Operation, like `GET`, `DELETE`, `PUT` or `POST` - but not limited to these values.
	Method pulumi.StringPtrInput
	// A unique identifier for this API Operation. Changing this forces a new resource to be created.
	OperationId pulumi.StringPtrInput
	// A `request` block as defined below.
	Request ApiOperationRequestPtrInput
	// The Name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// One or more `response` blocks as defined below.
	Responses ApiOperationResponseArrayInput
	// One or more `templateParameter` blocks as defined below.
	TemplateParameters ApiOperationTemplateParameterArrayInput
	// The relative URL Template identifying the target resource for this operation, which may include parameters.
	UrlTemplate pulumi.StringPtrInput
}

func (ApiOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOperationState)(nil)).Elem()
}

type apiOperationArgs struct {
	// The Name of the API Management Service where the API exists. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The name of the API within the API Management Service where this API Operation should be created. Changing this forces a new resource to be created.
	ApiName string `pulumi:"apiName"`
	// A description for this API Operation, which may include HTML formatting tags.
	Description *string `pulumi:"description"`
	// The Display Name for this API Management Operation.
	DisplayName string `pulumi:"displayName"`
	// The HTTP Method used for this API Management Operation, like `GET`, `DELETE`, `PUT` or `POST` - but not limited to these values.
	Method string `pulumi:"method"`
	// A unique identifier for this API Operation. Changing this forces a new resource to be created.
	OperationId string `pulumi:"operationId"`
	// A `request` block as defined below.
	Request *ApiOperationRequest `pulumi:"request"`
	// The Name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// One or more `response` blocks as defined below.
	Responses []ApiOperationResponse `pulumi:"responses"`
	// One or more `templateParameter` blocks as defined below.
	TemplateParameters []ApiOperationTemplateParameter `pulumi:"templateParameters"`
	// The relative URL Template identifying the target resource for this operation, which may include parameters.
	UrlTemplate string `pulumi:"urlTemplate"`
}

// The set of arguments for constructing a ApiOperation resource.
type ApiOperationArgs struct {
	// The Name of the API Management Service where the API exists. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The name of the API within the API Management Service where this API Operation should be created. Changing this forces a new resource to be created.
	ApiName pulumi.StringInput
	// A description for this API Operation, which may include HTML formatting tags.
	Description pulumi.StringPtrInput
	// The Display Name for this API Management Operation.
	DisplayName pulumi.StringInput
	// The HTTP Method used for this API Management Operation, like `GET`, `DELETE`, `PUT` or `POST` - but not limited to these values.
	Method pulumi.StringInput
	// A unique identifier for this API Operation. Changing this forces a new resource to be created.
	OperationId pulumi.StringInput
	// A `request` block as defined below.
	Request ApiOperationRequestPtrInput
	// The Name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// One or more `response` blocks as defined below.
	Responses ApiOperationResponseArrayInput
	// One or more `templateParameter` blocks as defined below.
	TemplateParameters ApiOperationTemplateParameterArrayInput
	// The relative URL Template identifying the target resource for this operation, which may include parameters.
	UrlTemplate pulumi.StringInput
}

func (ApiOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOperationArgs)(nil)).Elem()
}

type ApiOperationInput interface {
	pulumi.Input

	ToApiOperationOutput() ApiOperationOutput
	ToApiOperationOutputWithContext(ctx context.Context) ApiOperationOutput
}

func (*ApiOperation) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOperation)(nil))
}

func (i *ApiOperation) ToApiOperationOutput() ApiOperationOutput {
	return i.ToApiOperationOutputWithContext(context.Background())
}

func (i *ApiOperation) ToApiOperationOutputWithContext(ctx context.Context) ApiOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOperationOutput)
}

func (i *ApiOperation) ToApiOperationPtrOutput() ApiOperationPtrOutput {
	return i.ToApiOperationPtrOutputWithContext(context.Background())
}

func (i *ApiOperation) ToApiOperationPtrOutputWithContext(ctx context.Context) ApiOperationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOperationPtrOutput)
}

type ApiOperationPtrInput interface {
	pulumi.Input

	ToApiOperationPtrOutput() ApiOperationPtrOutput
	ToApiOperationPtrOutputWithContext(ctx context.Context) ApiOperationPtrOutput
}

type apiOperationPtrType ApiOperationArgs

func (*apiOperationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiOperation)(nil))
}

func (i *apiOperationPtrType) ToApiOperationPtrOutput() ApiOperationPtrOutput {
	return i.ToApiOperationPtrOutputWithContext(context.Background())
}

func (i *apiOperationPtrType) ToApiOperationPtrOutputWithContext(ctx context.Context) ApiOperationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOperationPtrOutput)
}

// ApiOperationArrayInput is an input type that accepts ApiOperationArray and ApiOperationArrayOutput values.
// You can construct a concrete instance of `ApiOperationArrayInput` via:
//
//          ApiOperationArray{ ApiOperationArgs{...} }
type ApiOperationArrayInput interface {
	pulumi.Input

	ToApiOperationArrayOutput() ApiOperationArrayOutput
	ToApiOperationArrayOutputWithContext(context.Context) ApiOperationArrayOutput
}

type ApiOperationArray []ApiOperationInput

func (ApiOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ApiOperation)(nil))
}

func (i ApiOperationArray) ToApiOperationArrayOutput() ApiOperationArrayOutput {
	return i.ToApiOperationArrayOutputWithContext(context.Background())
}

func (i ApiOperationArray) ToApiOperationArrayOutputWithContext(ctx context.Context) ApiOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOperationArrayOutput)
}

// ApiOperationMapInput is an input type that accepts ApiOperationMap and ApiOperationMapOutput values.
// You can construct a concrete instance of `ApiOperationMapInput` via:
//
//          ApiOperationMap{ "key": ApiOperationArgs{...} }
type ApiOperationMapInput interface {
	pulumi.Input

	ToApiOperationMapOutput() ApiOperationMapOutput
	ToApiOperationMapOutputWithContext(context.Context) ApiOperationMapOutput
}

type ApiOperationMap map[string]ApiOperationInput

func (ApiOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ApiOperation)(nil))
}

func (i ApiOperationMap) ToApiOperationMapOutput() ApiOperationMapOutput {
	return i.ToApiOperationMapOutputWithContext(context.Background())
}

func (i ApiOperationMap) ToApiOperationMapOutputWithContext(ctx context.Context) ApiOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOperationMapOutput)
}

type ApiOperationOutput struct {
	*pulumi.OutputState
}

func (ApiOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOperation)(nil))
}

func (o ApiOperationOutput) ToApiOperationOutput() ApiOperationOutput {
	return o
}

func (o ApiOperationOutput) ToApiOperationOutputWithContext(ctx context.Context) ApiOperationOutput {
	return o
}

func (o ApiOperationOutput) ToApiOperationPtrOutput() ApiOperationPtrOutput {
	return o.ToApiOperationPtrOutputWithContext(context.Background())
}

func (o ApiOperationOutput) ToApiOperationPtrOutputWithContext(ctx context.Context) ApiOperationPtrOutput {
	return o.ApplyT(func(v ApiOperation) *ApiOperation {
		return &v
	}).(ApiOperationPtrOutput)
}

type ApiOperationPtrOutput struct {
	*pulumi.OutputState
}

func (ApiOperationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiOperation)(nil))
}

func (o ApiOperationPtrOutput) ToApiOperationPtrOutput() ApiOperationPtrOutput {
	return o
}

func (o ApiOperationPtrOutput) ToApiOperationPtrOutputWithContext(ctx context.Context) ApiOperationPtrOutput {
	return o
}

type ApiOperationArrayOutput struct{ *pulumi.OutputState }

func (ApiOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiOperation)(nil))
}

func (o ApiOperationArrayOutput) ToApiOperationArrayOutput() ApiOperationArrayOutput {
	return o
}

func (o ApiOperationArrayOutput) ToApiOperationArrayOutputWithContext(ctx context.Context) ApiOperationArrayOutput {
	return o
}

func (o ApiOperationArrayOutput) Index(i pulumi.IntInput) ApiOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiOperation {
		return vs[0].([]ApiOperation)[vs[1].(int)]
	}).(ApiOperationOutput)
}

type ApiOperationMapOutput struct{ *pulumi.OutputState }

func (ApiOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApiOperation)(nil))
}

func (o ApiOperationMapOutput) ToApiOperationMapOutput() ApiOperationMapOutput {
	return o
}

func (o ApiOperationMapOutput) ToApiOperationMapOutputWithContext(ctx context.Context) ApiOperationMapOutput {
	return o
}

func (o ApiOperationMapOutput) MapIndex(k pulumi.StringInput) ApiOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApiOperation {
		return vs[0].(map[string]ApiOperation)[vs[1].(string)]
	}).(ApiOperationOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiOperationOutput{})
	pulumi.RegisterOutputType(ApiOperationPtrOutput{})
	pulumi.RegisterOutputType(ApiOperationArrayOutput{})
	pulumi.RegisterOutputType(ApiOperationMapOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Authorization Server within an API Management Service.
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
// 		_, err := apimanagement.LookupApi(ctx, &apimanagement.LookupApiArgs{
// 			Name:              "search-api",
// 			ApiManagementName: "search-api-management",
// 			ResourceGroupName: "search-service",
// 			Revision:          "2",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apimanagement.NewAuthorizationServer(ctx, "exampleAuthorizationServer", &apimanagement.AuthorizationServerArgs{
// 			ApiManagementName:          pulumi.Any(data.Azurerm_api_management.Example.Name),
// 			ResourceGroupName:          pulumi.Any(data.Azurerm_api_management.Example.Resource_group_name),
// 			DisplayName:                pulumi.String("Test Server"),
// 			AuthorizationEndpoint:      pulumi.String("https://example.mydomain.com/client/authorize"),
// 			ClientId:                   pulumi.String("42424242-4242-4242-4242-424242424242"),
// 			ClientRegistrationEndpoint: pulumi.String("https://example.mydomain.com/client/register"),
// 			GrantTypes: pulumi.StringArray{
// 				pulumi.String("authorizationCode"),
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
// API Management Authorization Servers can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/authorizationServer:AuthorizationServer example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/authorizationServers/server1
// ```
type AuthorizationServer struct {
	pulumi.CustomResourceState

	// The name of the API Management Service in which this Authorization Server should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The OAUTH Authorization Endpoint.
	AuthorizationEndpoint pulumi.StringOutput `pulumi:"authorizationEndpoint"`
	// The HTTP Verbs supported by the Authorization Endpoint. Possible values are `DELETE`, `GET`, `HEAD`, `OPTIONS`, `PATCH`, `POST`, `PUT` and `TRACE`.
	AuthorizationMethods pulumi.StringArrayOutput `pulumi:"authorizationMethods"`
	// The mechanism by which Access Tokens are passed to the API. Possible values are `authorizationHeader` and `query`.
	BearerTokenSendingMethods pulumi.StringArrayOutput `pulumi:"bearerTokenSendingMethods"`
	// The Authentication Methods supported by the Token endpoint of this Authorization Server.. Possible values are `Basic` and `Body`.
	ClientAuthenticationMethods pulumi.StringArrayOutput `pulumi:"clientAuthenticationMethods"`
	// The Client/App ID registered with this Authorization Server.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The URI of page where Client/App Registration is performed for this Authorization Server.
	ClientRegistrationEndpoint pulumi.StringOutput `pulumi:"clientRegistrationEndpoint"`
	// The Client/App Secret registered with this Authorization Server.
	ClientSecret pulumi.StringPtrOutput `pulumi:"clientSecret"`
	// The Default Scope used when requesting an Access Token, specified as a string containing space-delimited values.
	DefaultScope pulumi.StringPtrOutput `pulumi:"defaultScope"`
	// A description of the Authorization Server, which may contain HTML formatting tags.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The user-friendly name of this Authorization Server.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Form of Authorization Grants required when requesting an Access Token. Possible values are `authorizationCode`, `clientCredentials`, `implicit` and `resourceOwnerPassword`.
	GrantTypes pulumi.StringArrayOutput `pulumi:"grantTypes"`
	// The name of this Authorization Server. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The password associated with the Resource Owner.
	ResourceOwnerPassword pulumi.StringPtrOutput `pulumi:"resourceOwnerPassword"`
	// The username associated with the Resource Owner.
	ResourceOwnerUsername pulumi.StringPtrOutput `pulumi:"resourceOwnerUsername"`
	// Does this Authorization Server support State? If this is set to `true` the client may use the state parameter to raise protocol security.
	SupportState pulumi.BoolPtrOutput `pulumi:"supportState"`
	// A `tokenBodyParameter` block as defined below.
	TokenBodyParameters AuthorizationServerTokenBodyParameterArrayOutput `pulumi:"tokenBodyParameters"`
	// The OAUTH Token Endpoint.
	TokenEndpoint pulumi.StringPtrOutput `pulumi:"tokenEndpoint"`
}

// NewAuthorizationServer registers a new resource with the given unique name, arguments, and options.
func NewAuthorizationServer(ctx *pulumi.Context,
	name string, args *AuthorizationServerArgs, opts ...pulumi.ResourceOption) (*AuthorizationServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiManagementName == nil {
		return nil, errors.New("invalid value for required argument 'ApiManagementName'")
	}
	if args.AuthorizationEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizationEndpoint'")
	}
	if args.AuthorizationMethods == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizationMethods'")
	}
	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClientRegistrationEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'ClientRegistrationEndpoint'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.GrantTypes == nil {
		return nil, errors.New("invalid value for required argument 'GrantTypes'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource AuthorizationServer
	err := ctx.RegisterResource("azure:apimanagement/authorizationServer:AuthorizationServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthorizationServer gets an existing AuthorizationServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthorizationServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorizationServerState, opts ...pulumi.ResourceOption) (*AuthorizationServer, error) {
	var resource AuthorizationServer
	err := ctx.ReadResource("azure:apimanagement/authorizationServer:AuthorizationServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthorizationServer resources.
type authorizationServerState struct {
	// The name of the API Management Service in which this Authorization Server should be created. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The OAUTH Authorization Endpoint.
	AuthorizationEndpoint *string `pulumi:"authorizationEndpoint"`
	// The HTTP Verbs supported by the Authorization Endpoint. Possible values are `DELETE`, `GET`, `HEAD`, `OPTIONS`, `PATCH`, `POST`, `PUT` and `TRACE`.
	AuthorizationMethods []string `pulumi:"authorizationMethods"`
	// The mechanism by which Access Tokens are passed to the API. Possible values are `authorizationHeader` and `query`.
	BearerTokenSendingMethods []string `pulumi:"bearerTokenSendingMethods"`
	// The Authentication Methods supported by the Token endpoint of this Authorization Server.. Possible values are `Basic` and `Body`.
	ClientAuthenticationMethods []string `pulumi:"clientAuthenticationMethods"`
	// The Client/App ID registered with this Authorization Server.
	ClientId *string `pulumi:"clientId"`
	// The URI of page where Client/App Registration is performed for this Authorization Server.
	ClientRegistrationEndpoint *string `pulumi:"clientRegistrationEndpoint"`
	// The Client/App Secret registered with this Authorization Server.
	ClientSecret *string `pulumi:"clientSecret"`
	// The Default Scope used when requesting an Access Token, specified as a string containing space-delimited values.
	DefaultScope *string `pulumi:"defaultScope"`
	// A description of the Authorization Server, which may contain HTML formatting tags.
	Description *string `pulumi:"description"`
	// The user-friendly name of this Authorization Server.
	DisplayName *string `pulumi:"displayName"`
	// Form of Authorization Grants required when requesting an Access Token. Possible values are `authorizationCode`, `clientCredentials`, `implicit` and `resourceOwnerPassword`.
	GrantTypes []string `pulumi:"grantTypes"`
	// The name of this Authorization Server. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The password associated with the Resource Owner.
	ResourceOwnerPassword *string `pulumi:"resourceOwnerPassword"`
	// The username associated with the Resource Owner.
	ResourceOwnerUsername *string `pulumi:"resourceOwnerUsername"`
	// Does this Authorization Server support State? If this is set to `true` the client may use the state parameter to raise protocol security.
	SupportState *bool `pulumi:"supportState"`
	// A `tokenBodyParameter` block as defined below.
	TokenBodyParameters []AuthorizationServerTokenBodyParameter `pulumi:"tokenBodyParameters"`
	// The OAUTH Token Endpoint.
	TokenEndpoint *string `pulumi:"tokenEndpoint"`
}

type AuthorizationServerState struct {
	// The name of the API Management Service in which this Authorization Server should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The OAUTH Authorization Endpoint.
	AuthorizationEndpoint pulumi.StringPtrInput
	// The HTTP Verbs supported by the Authorization Endpoint. Possible values are `DELETE`, `GET`, `HEAD`, `OPTIONS`, `PATCH`, `POST`, `PUT` and `TRACE`.
	AuthorizationMethods pulumi.StringArrayInput
	// The mechanism by which Access Tokens are passed to the API. Possible values are `authorizationHeader` and `query`.
	BearerTokenSendingMethods pulumi.StringArrayInput
	// The Authentication Methods supported by the Token endpoint of this Authorization Server.. Possible values are `Basic` and `Body`.
	ClientAuthenticationMethods pulumi.StringArrayInput
	// The Client/App ID registered with this Authorization Server.
	ClientId pulumi.StringPtrInput
	// The URI of page where Client/App Registration is performed for this Authorization Server.
	ClientRegistrationEndpoint pulumi.StringPtrInput
	// The Client/App Secret registered with this Authorization Server.
	ClientSecret pulumi.StringPtrInput
	// The Default Scope used when requesting an Access Token, specified as a string containing space-delimited values.
	DefaultScope pulumi.StringPtrInput
	// A description of the Authorization Server, which may contain HTML formatting tags.
	Description pulumi.StringPtrInput
	// The user-friendly name of this Authorization Server.
	DisplayName pulumi.StringPtrInput
	// Form of Authorization Grants required when requesting an Access Token. Possible values are `authorizationCode`, `clientCredentials`, `implicit` and `resourceOwnerPassword`.
	GrantTypes pulumi.StringArrayInput
	// The name of this Authorization Server. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The password associated with the Resource Owner.
	ResourceOwnerPassword pulumi.StringPtrInput
	// The username associated with the Resource Owner.
	ResourceOwnerUsername pulumi.StringPtrInput
	// Does this Authorization Server support State? If this is set to `true` the client may use the state parameter to raise protocol security.
	SupportState pulumi.BoolPtrInput
	// A `tokenBodyParameter` block as defined below.
	TokenBodyParameters AuthorizationServerTokenBodyParameterArrayInput
	// The OAUTH Token Endpoint.
	TokenEndpoint pulumi.StringPtrInput
}

func (AuthorizationServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationServerState)(nil)).Elem()
}

type authorizationServerArgs struct {
	// The name of the API Management Service in which this Authorization Server should be created. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The OAUTH Authorization Endpoint.
	AuthorizationEndpoint string `pulumi:"authorizationEndpoint"`
	// The HTTP Verbs supported by the Authorization Endpoint. Possible values are `DELETE`, `GET`, `HEAD`, `OPTIONS`, `PATCH`, `POST`, `PUT` and `TRACE`.
	AuthorizationMethods []string `pulumi:"authorizationMethods"`
	// The mechanism by which Access Tokens are passed to the API. Possible values are `authorizationHeader` and `query`.
	BearerTokenSendingMethods []string `pulumi:"bearerTokenSendingMethods"`
	// The Authentication Methods supported by the Token endpoint of this Authorization Server.. Possible values are `Basic` and `Body`.
	ClientAuthenticationMethods []string `pulumi:"clientAuthenticationMethods"`
	// The Client/App ID registered with this Authorization Server.
	ClientId string `pulumi:"clientId"`
	// The URI of page where Client/App Registration is performed for this Authorization Server.
	ClientRegistrationEndpoint string `pulumi:"clientRegistrationEndpoint"`
	// The Client/App Secret registered with this Authorization Server.
	ClientSecret *string `pulumi:"clientSecret"`
	// The Default Scope used when requesting an Access Token, specified as a string containing space-delimited values.
	DefaultScope *string `pulumi:"defaultScope"`
	// A description of the Authorization Server, which may contain HTML formatting tags.
	Description *string `pulumi:"description"`
	// The user-friendly name of this Authorization Server.
	DisplayName string `pulumi:"displayName"`
	// Form of Authorization Grants required when requesting an Access Token. Possible values are `authorizationCode`, `clientCredentials`, `implicit` and `resourceOwnerPassword`.
	GrantTypes []string `pulumi:"grantTypes"`
	// The name of this Authorization Server. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The password associated with the Resource Owner.
	ResourceOwnerPassword *string `pulumi:"resourceOwnerPassword"`
	// The username associated with the Resource Owner.
	ResourceOwnerUsername *string `pulumi:"resourceOwnerUsername"`
	// Does this Authorization Server support State? If this is set to `true` the client may use the state parameter to raise protocol security.
	SupportState *bool `pulumi:"supportState"`
	// A `tokenBodyParameter` block as defined below.
	TokenBodyParameters []AuthorizationServerTokenBodyParameter `pulumi:"tokenBodyParameters"`
	// The OAUTH Token Endpoint.
	TokenEndpoint *string `pulumi:"tokenEndpoint"`
}

// The set of arguments for constructing a AuthorizationServer resource.
type AuthorizationServerArgs struct {
	// The name of the API Management Service in which this Authorization Server should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The OAUTH Authorization Endpoint.
	AuthorizationEndpoint pulumi.StringInput
	// The HTTP Verbs supported by the Authorization Endpoint. Possible values are `DELETE`, `GET`, `HEAD`, `OPTIONS`, `PATCH`, `POST`, `PUT` and `TRACE`.
	AuthorizationMethods pulumi.StringArrayInput
	// The mechanism by which Access Tokens are passed to the API. Possible values are `authorizationHeader` and `query`.
	BearerTokenSendingMethods pulumi.StringArrayInput
	// The Authentication Methods supported by the Token endpoint of this Authorization Server.. Possible values are `Basic` and `Body`.
	ClientAuthenticationMethods pulumi.StringArrayInput
	// The Client/App ID registered with this Authorization Server.
	ClientId pulumi.StringInput
	// The URI of page where Client/App Registration is performed for this Authorization Server.
	ClientRegistrationEndpoint pulumi.StringInput
	// The Client/App Secret registered with this Authorization Server.
	ClientSecret pulumi.StringPtrInput
	// The Default Scope used when requesting an Access Token, specified as a string containing space-delimited values.
	DefaultScope pulumi.StringPtrInput
	// A description of the Authorization Server, which may contain HTML formatting tags.
	Description pulumi.StringPtrInput
	// The user-friendly name of this Authorization Server.
	DisplayName pulumi.StringInput
	// Form of Authorization Grants required when requesting an Access Token. Possible values are `authorizationCode`, `clientCredentials`, `implicit` and `resourceOwnerPassword`.
	GrantTypes pulumi.StringArrayInput
	// The name of this Authorization Server. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The password associated with the Resource Owner.
	ResourceOwnerPassword pulumi.StringPtrInput
	// The username associated with the Resource Owner.
	ResourceOwnerUsername pulumi.StringPtrInput
	// Does this Authorization Server support State? If this is set to `true` the client may use the state parameter to raise protocol security.
	SupportState pulumi.BoolPtrInput
	// A `tokenBodyParameter` block as defined below.
	TokenBodyParameters AuthorizationServerTokenBodyParameterArrayInput
	// The OAUTH Token Endpoint.
	TokenEndpoint pulumi.StringPtrInput
}

func (AuthorizationServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationServerArgs)(nil)).Elem()
}

type AuthorizationServerInput interface {
	pulumi.Input

	ToAuthorizationServerOutput() AuthorizationServerOutput
	ToAuthorizationServerOutputWithContext(ctx context.Context) AuthorizationServerOutput
}

func (*AuthorizationServer) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationServer)(nil))
}

func (i *AuthorizationServer) ToAuthorizationServerOutput() AuthorizationServerOutput {
	return i.ToAuthorizationServerOutputWithContext(context.Background())
}

func (i *AuthorizationServer) ToAuthorizationServerOutputWithContext(ctx context.Context) AuthorizationServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationServerOutput)
}

func (i *AuthorizationServer) ToAuthorizationServerPtrOutput() AuthorizationServerPtrOutput {
	return i.ToAuthorizationServerPtrOutputWithContext(context.Background())
}

func (i *AuthorizationServer) ToAuthorizationServerPtrOutputWithContext(ctx context.Context) AuthorizationServerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationServerPtrOutput)
}

type AuthorizationServerPtrInput interface {
	pulumi.Input

	ToAuthorizationServerPtrOutput() AuthorizationServerPtrOutput
	ToAuthorizationServerPtrOutputWithContext(ctx context.Context) AuthorizationServerPtrOutput
}

type authorizationServerPtrType AuthorizationServerArgs

func (*authorizationServerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationServer)(nil))
}

func (i *authorizationServerPtrType) ToAuthorizationServerPtrOutput() AuthorizationServerPtrOutput {
	return i.ToAuthorizationServerPtrOutputWithContext(context.Background())
}

func (i *authorizationServerPtrType) ToAuthorizationServerPtrOutputWithContext(ctx context.Context) AuthorizationServerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationServerPtrOutput)
}

// AuthorizationServerArrayInput is an input type that accepts AuthorizationServerArray and AuthorizationServerArrayOutput values.
// You can construct a concrete instance of `AuthorizationServerArrayInput` via:
//
//          AuthorizationServerArray{ AuthorizationServerArgs{...} }
type AuthorizationServerArrayInput interface {
	pulumi.Input

	ToAuthorizationServerArrayOutput() AuthorizationServerArrayOutput
	ToAuthorizationServerArrayOutputWithContext(context.Context) AuthorizationServerArrayOutput
}

type AuthorizationServerArray []AuthorizationServerInput

func (AuthorizationServerArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AuthorizationServer)(nil))
}

func (i AuthorizationServerArray) ToAuthorizationServerArrayOutput() AuthorizationServerArrayOutput {
	return i.ToAuthorizationServerArrayOutputWithContext(context.Background())
}

func (i AuthorizationServerArray) ToAuthorizationServerArrayOutputWithContext(ctx context.Context) AuthorizationServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationServerArrayOutput)
}

// AuthorizationServerMapInput is an input type that accepts AuthorizationServerMap and AuthorizationServerMapOutput values.
// You can construct a concrete instance of `AuthorizationServerMapInput` via:
//
//          AuthorizationServerMap{ "key": AuthorizationServerArgs{...} }
type AuthorizationServerMapInput interface {
	pulumi.Input

	ToAuthorizationServerMapOutput() AuthorizationServerMapOutput
	ToAuthorizationServerMapOutputWithContext(context.Context) AuthorizationServerMapOutput
}

type AuthorizationServerMap map[string]AuthorizationServerInput

func (AuthorizationServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AuthorizationServer)(nil))
}

func (i AuthorizationServerMap) ToAuthorizationServerMapOutput() AuthorizationServerMapOutput {
	return i.ToAuthorizationServerMapOutputWithContext(context.Background())
}

func (i AuthorizationServerMap) ToAuthorizationServerMapOutputWithContext(ctx context.Context) AuthorizationServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationServerMapOutput)
}

type AuthorizationServerOutput struct {
	*pulumi.OutputState
}

func (AuthorizationServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationServer)(nil))
}

func (o AuthorizationServerOutput) ToAuthorizationServerOutput() AuthorizationServerOutput {
	return o
}

func (o AuthorizationServerOutput) ToAuthorizationServerOutputWithContext(ctx context.Context) AuthorizationServerOutput {
	return o
}

func (o AuthorizationServerOutput) ToAuthorizationServerPtrOutput() AuthorizationServerPtrOutput {
	return o.ToAuthorizationServerPtrOutputWithContext(context.Background())
}

func (o AuthorizationServerOutput) ToAuthorizationServerPtrOutputWithContext(ctx context.Context) AuthorizationServerPtrOutput {
	return o.ApplyT(func(v AuthorizationServer) *AuthorizationServer {
		return &v
	}).(AuthorizationServerPtrOutput)
}

type AuthorizationServerPtrOutput struct {
	*pulumi.OutputState
}

func (AuthorizationServerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationServer)(nil))
}

func (o AuthorizationServerPtrOutput) ToAuthorizationServerPtrOutput() AuthorizationServerPtrOutput {
	return o
}

func (o AuthorizationServerPtrOutput) ToAuthorizationServerPtrOutputWithContext(ctx context.Context) AuthorizationServerPtrOutput {
	return o
}

type AuthorizationServerArrayOutput struct{ *pulumi.OutputState }

func (AuthorizationServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthorizationServer)(nil))
}

func (o AuthorizationServerArrayOutput) ToAuthorizationServerArrayOutput() AuthorizationServerArrayOutput {
	return o
}

func (o AuthorizationServerArrayOutput) ToAuthorizationServerArrayOutputWithContext(ctx context.Context) AuthorizationServerArrayOutput {
	return o
}

func (o AuthorizationServerArrayOutput) Index(i pulumi.IntInput) AuthorizationServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AuthorizationServer {
		return vs[0].([]AuthorizationServer)[vs[1].(int)]
	}).(AuthorizationServerOutput)
}

type AuthorizationServerMapOutput struct{ *pulumi.OutputState }

func (AuthorizationServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AuthorizationServer)(nil))
}

func (o AuthorizationServerMapOutput) ToAuthorizationServerMapOutput() AuthorizationServerMapOutput {
	return o
}

func (o AuthorizationServerMapOutput) ToAuthorizationServerMapOutputWithContext(ctx context.Context) AuthorizationServerMapOutput {
	return o
}

func (o AuthorizationServerMapOutput) MapIndex(k pulumi.StringInput) AuthorizationServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AuthorizationServer {
		return vs[0].(map[string]AuthorizationServer)[vs[1].(string)]
	}).(AuthorizationServerOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthorizationServerOutput{})
	pulumi.RegisterOutputType(AuthorizationServerPtrOutput{})
	pulumi.RegisterOutputType(AuthorizationServerArrayOutput{})
	pulumi.RegisterOutputType(AuthorizationServerMapOutput{})
}

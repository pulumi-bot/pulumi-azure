// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an API Management Twitter Identity Provider.
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
// 		_, err = apimanagement.NewIdentityProviderTwitter(ctx, "exampleIdentityProviderTwitter", &apimanagement.IdentityProviderTwitterArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApiManagementName: exampleService.Name,
// 			ApiKey:            pulumi.String("00000000000000000000000000000000"),
// 			ApiSecretKey:      pulumi.String("00000000000000000000000000000000"),
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
// API Management Twitter Identity Provider can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/identityProviderTwitter:IdentityProviderTwitter example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/twitter
// ```
type IdentityProviderTwitter struct {
	pulumi.CustomResourceState

	// App Consumer API key for Twitter.
	ApiKey pulumi.StringOutput `pulumi:"apiKey"`
	// The Name of the API Management Service where this Twitter Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// App Consumer API secret key for Twitter.
	ApiSecretKey pulumi.StringOutput `pulumi:"apiSecretKey"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewIdentityProviderTwitter registers a new resource with the given unique name, arguments, and options.
func NewIdentityProviderTwitter(ctx *pulumi.Context,
	name string, args *IdentityProviderTwitterArgs, opts ...pulumi.ResourceOption) (*IdentityProviderTwitter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiKey == nil {
		return nil, errors.New("invalid value for required argument 'ApiKey'")
	}
	if args.ApiManagementName == nil {
		return nil, errors.New("invalid value for required argument 'ApiManagementName'")
	}
	if args.ApiSecretKey == nil {
		return nil, errors.New("invalid value for required argument 'ApiSecretKey'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource IdentityProviderTwitter
	err := ctx.RegisterResource("azure:apimanagement/identityProviderTwitter:IdentityProviderTwitter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityProviderTwitter gets an existing IdentityProviderTwitter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityProviderTwitter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityProviderTwitterState, opts ...pulumi.ResourceOption) (*IdentityProviderTwitter, error) {
	var resource IdentityProviderTwitter
	err := ctx.ReadResource("azure:apimanagement/identityProviderTwitter:IdentityProviderTwitter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityProviderTwitter resources.
type identityProviderTwitterState struct {
	// App Consumer API key for Twitter.
	ApiKey *string `pulumi:"apiKey"`
	// The Name of the API Management Service where this Twitter Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// App Consumer API secret key for Twitter.
	ApiSecretKey *string `pulumi:"apiSecretKey"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type IdentityProviderTwitterState struct {
	// App Consumer API key for Twitter.
	ApiKey pulumi.StringPtrInput
	// The Name of the API Management Service where this Twitter Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// App Consumer API secret key for Twitter.
	ApiSecretKey pulumi.StringPtrInput
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (IdentityProviderTwitterState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderTwitterState)(nil)).Elem()
}

type identityProviderTwitterArgs struct {
	// App Consumer API key for Twitter.
	ApiKey string `pulumi:"apiKey"`
	// The Name of the API Management Service where this Twitter Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// App Consumer API secret key for Twitter.
	ApiSecretKey string `pulumi:"apiSecretKey"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a IdentityProviderTwitter resource.
type IdentityProviderTwitterArgs struct {
	// App Consumer API key for Twitter.
	ApiKey pulumi.StringInput
	// The Name of the API Management Service where this Twitter Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// App Consumer API secret key for Twitter.
	ApiSecretKey pulumi.StringInput
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (IdentityProviderTwitterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderTwitterArgs)(nil)).Elem()
}

type IdentityProviderTwitterInput interface {
	pulumi.Input

	ToIdentityProviderTwitterOutput() IdentityProviderTwitterOutput
	ToIdentityProviderTwitterOutputWithContext(ctx context.Context) IdentityProviderTwitterOutput
}

func (*IdentityProviderTwitter) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProviderTwitter)(nil))
}

func (i *IdentityProviderTwitter) ToIdentityProviderTwitterOutput() IdentityProviderTwitterOutput {
	return i.ToIdentityProviderTwitterOutputWithContext(context.Background())
}

func (i *IdentityProviderTwitter) ToIdentityProviderTwitterOutputWithContext(ctx context.Context) IdentityProviderTwitterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderTwitterOutput)
}

func (i *IdentityProviderTwitter) ToIdentityProviderTwitterPtrOutput() IdentityProviderTwitterPtrOutput {
	return i.ToIdentityProviderTwitterPtrOutputWithContext(context.Background())
}

func (i *IdentityProviderTwitter) ToIdentityProviderTwitterPtrOutputWithContext(ctx context.Context) IdentityProviderTwitterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderTwitterPtrOutput)
}

type IdentityProviderTwitterPtrInput interface {
	pulumi.Input

	ToIdentityProviderTwitterPtrOutput() IdentityProviderTwitterPtrOutput
	ToIdentityProviderTwitterPtrOutputWithContext(ctx context.Context) IdentityProviderTwitterPtrOutput
}

type identityProviderTwitterPtrType IdentityProviderTwitterArgs

func (*identityProviderTwitterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProviderTwitter)(nil))
}

func (i *identityProviderTwitterPtrType) ToIdentityProviderTwitterPtrOutput() IdentityProviderTwitterPtrOutput {
	return i.ToIdentityProviderTwitterPtrOutputWithContext(context.Background())
}

func (i *identityProviderTwitterPtrType) ToIdentityProviderTwitterPtrOutputWithContext(ctx context.Context) IdentityProviderTwitterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderTwitterOutput).ToIdentityProviderTwitterPtrOutput()
}

// IdentityProviderTwitterArrayInput is an input type that accepts IdentityProviderTwitterArray and IdentityProviderTwitterArrayOutput values.
// You can construct a concrete instance of `IdentityProviderTwitterArrayInput` via:
//
//          IdentityProviderTwitterArray{ IdentityProviderTwitterArgs{...} }
type IdentityProviderTwitterArrayInput interface {
	pulumi.Input

	ToIdentityProviderTwitterArrayOutput() IdentityProviderTwitterArrayOutput
	ToIdentityProviderTwitterArrayOutputWithContext(context.Context) IdentityProviderTwitterArrayOutput
}

type IdentityProviderTwitterArray []IdentityProviderTwitterInput

func (IdentityProviderTwitterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentityProviderTwitter)(nil))
}

func (i IdentityProviderTwitterArray) ToIdentityProviderTwitterArrayOutput() IdentityProviderTwitterArrayOutput {
	return i.ToIdentityProviderTwitterArrayOutputWithContext(context.Background())
}

func (i IdentityProviderTwitterArray) ToIdentityProviderTwitterArrayOutputWithContext(ctx context.Context) IdentityProviderTwitterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderTwitterArrayOutput)
}

// IdentityProviderTwitterMapInput is an input type that accepts IdentityProviderTwitterMap and IdentityProviderTwitterMapOutput values.
// You can construct a concrete instance of `IdentityProviderTwitterMapInput` via:
//
//          IdentityProviderTwitterMap{ "key": IdentityProviderTwitterArgs{...} }
type IdentityProviderTwitterMapInput interface {
	pulumi.Input

	ToIdentityProviderTwitterMapOutput() IdentityProviderTwitterMapOutput
	ToIdentityProviderTwitterMapOutputWithContext(context.Context) IdentityProviderTwitterMapOutput
}

type IdentityProviderTwitterMap map[string]IdentityProviderTwitterInput

func (IdentityProviderTwitterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IdentityProviderTwitter)(nil))
}

func (i IdentityProviderTwitterMap) ToIdentityProviderTwitterMapOutput() IdentityProviderTwitterMapOutput {
	return i.ToIdentityProviderTwitterMapOutputWithContext(context.Background())
}

func (i IdentityProviderTwitterMap) ToIdentityProviderTwitterMapOutputWithContext(ctx context.Context) IdentityProviderTwitterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderTwitterMapOutput)
}

type IdentityProviderTwitterOutput struct {
	*pulumi.OutputState
}

func (IdentityProviderTwitterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProviderTwitter)(nil))
}

func (o IdentityProviderTwitterOutput) ToIdentityProviderTwitterOutput() IdentityProviderTwitterOutput {
	return o
}

func (o IdentityProviderTwitterOutput) ToIdentityProviderTwitterOutputWithContext(ctx context.Context) IdentityProviderTwitterOutput {
	return o
}

func (o IdentityProviderTwitterOutput) ToIdentityProviderTwitterPtrOutput() IdentityProviderTwitterPtrOutput {
	return o.ToIdentityProviderTwitterPtrOutputWithContext(context.Background())
}

func (o IdentityProviderTwitterOutput) ToIdentityProviderTwitterPtrOutputWithContext(ctx context.Context) IdentityProviderTwitterPtrOutput {
	return o.ApplyT(func(v IdentityProviderTwitter) *IdentityProviderTwitter {
		return &v
	}).(IdentityProviderTwitterPtrOutput)
}

type IdentityProviderTwitterPtrOutput struct {
	*pulumi.OutputState
}

func (IdentityProviderTwitterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProviderTwitter)(nil))
}

func (o IdentityProviderTwitterPtrOutput) ToIdentityProviderTwitterPtrOutput() IdentityProviderTwitterPtrOutput {
	return o
}

func (o IdentityProviderTwitterPtrOutput) ToIdentityProviderTwitterPtrOutputWithContext(ctx context.Context) IdentityProviderTwitterPtrOutput {
	return o
}

type IdentityProviderTwitterArrayOutput struct{ *pulumi.OutputState }

func (IdentityProviderTwitterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentityProviderTwitter)(nil))
}

func (o IdentityProviderTwitterArrayOutput) ToIdentityProviderTwitterArrayOutput() IdentityProviderTwitterArrayOutput {
	return o
}

func (o IdentityProviderTwitterArrayOutput) ToIdentityProviderTwitterArrayOutputWithContext(ctx context.Context) IdentityProviderTwitterArrayOutput {
	return o
}

func (o IdentityProviderTwitterArrayOutput) Index(i pulumi.IntInput) IdentityProviderTwitterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IdentityProviderTwitter {
		return vs[0].([]IdentityProviderTwitter)[vs[1].(int)]
	}).(IdentityProviderTwitterOutput)
}

type IdentityProviderTwitterMapOutput struct{ *pulumi.OutputState }

func (IdentityProviderTwitterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IdentityProviderTwitter)(nil))
}

func (o IdentityProviderTwitterMapOutput) ToIdentityProviderTwitterMapOutput() IdentityProviderTwitterMapOutput {
	return o
}

func (o IdentityProviderTwitterMapOutput) ToIdentityProviderTwitterMapOutputWithContext(ctx context.Context) IdentityProviderTwitterMapOutput {
	return o
}

func (o IdentityProviderTwitterMapOutput) MapIndex(k pulumi.StringInput) IdentityProviderTwitterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IdentityProviderTwitter {
		return vs[0].(map[string]IdentityProviderTwitter)[vs[1].(string)]
	}).(IdentityProviderTwitterOutput)
}

func init() {
	pulumi.RegisterOutputType(IdentityProviderTwitterOutput{})
	pulumi.RegisterOutputType(IdentityProviderTwitterPtrOutput{})
	pulumi.RegisterOutputType(IdentityProviderTwitterArrayOutput{})
	pulumi.RegisterOutputType(IdentityProviderTwitterMapOutput{})
}

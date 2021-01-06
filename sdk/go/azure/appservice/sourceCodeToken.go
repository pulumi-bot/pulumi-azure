// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an App Service source control token.
//
// > **NOTE:** Source Control Tokens are configured at the subscription level, not on each App Service - as such this can only be configured Subscription-wide
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/appservice"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := appservice.NewSourceCodeToken(ctx, "example", &appservice.SourceCodeTokenArgs{
// 			Token: pulumi.String("7e57735e77e577e57"),
// 			Type:  pulumi.String("GitHub"),
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
// App Service Source Control Token's can be imported using the `type`, e.g.
//
// ```sh
//  $ pulumi import azure:appservice/sourceCodeToken:SourceCodeToken example GitHub
// ```
type SourceCodeToken struct {
	pulumi.CustomResourceState

	// The OAuth access token.
	Token pulumi.StringOutput `pulumi:"token"`
	// The OAuth access token secret.
	TokenSecret pulumi.StringPtrOutput `pulumi:"tokenSecret"`
	// The source control type. Possible values are `BitBucket`, `Dropbox`, `GitHub` and `OneDrive`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSourceCodeToken registers a new resource with the given unique name, arguments, and options.
func NewSourceCodeToken(ctx *pulumi.Context,
	name string, args *SourceCodeTokenArgs, opts ...pulumi.ResourceOption) (*SourceCodeToken, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource SourceCodeToken
	err := ctx.RegisterResource("azure:appservice/sourceCodeToken:SourceCodeToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceCodeToken gets an existing SourceCodeToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceCodeToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceCodeTokenState, opts ...pulumi.ResourceOption) (*SourceCodeToken, error) {
	var resource SourceCodeToken
	err := ctx.ReadResource("azure:appservice/sourceCodeToken:SourceCodeToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceCodeToken resources.
type sourceCodeTokenState struct {
	// The OAuth access token.
	Token *string `pulumi:"token"`
	// The OAuth access token secret.
	TokenSecret *string `pulumi:"tokenSecret"`
	// The source control type. Possible values are `BitBucket`, `Dropbox`, `GitHub` and `OneDrive`.
	Type *string `pulumi:"type"`
}

type SourceCodeTokenState struct {
	// The OAuth access token.
	Token pulumi.StringPtrInput
	// The OAuth access token secret.
	TokenSecret pulumi.StringPtrInput
	// The source control type. Possible values are `BitBucket`, `Dropbox`, `GitHub` and `OneDrive`.
	Type pulumi.StringPtrInput
}

func (SourceCodeTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceCodeTokenState)(nil)).Elem()
}

type sourceCodeTokenArgs struct {
	// The OAuth access token.
	Token string `pulumi:"token"`
	// The OAuth access token secret.
	TokenSecret *string `pulumi:"tokenSecret"`
	// The source control type. Possible values are `BitBucket`, `Dropbox`, `GitHub` and `OneDrive`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a SourceCodeToken resource.
type SourceCodeTokenArgs struct {
	// The OAuth access token.
	Token pulumi.StringInput
	// The OAuth access token secret.
	TokenSecret pulumi.StringPtrInput
	// The source control type. Possible values are `BitBucket`, `Dropbox`, `GitHub` and `OneDrive`.
	Type pulumi.StringInput
}

func (SourceCodeTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceCodeTokenArgs)(nil)).Elem()
}

type SourceCodeTokenInput interface {
	pulumi.Input

	ToSourceCodeTokenOutput() SourceCodeTokenOutput
	ToSourceCodeTokenOutputWithContext(ctx context.Context) SourceCodeTokenOutput
}

func (*SourceCodeToken) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceCodeToken)(nil))
}

func (i *SourceCodeToken) ToSourceCodeTokenOutput() SourceCodeTokenOutput {
	return i.ToSourceCodeTokenOutputWithContext(context.Background())
}

func (i *SourceCodeToken) ToSourceCodeTokenOutputWithContext(ctx context.Context) SourceCodeTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCodeTokenOutput)
}

func (i *SourceCodeToken) ToSourceCodeTokenPtrOutput() SourceCodeTokenPtrOutput {
	return i.ToSourceCodeTokenPtrOutputWithContext(context.Background())
}

func (i *SourceCodeToken) ToSourceCodeTokenPtrOutputWithContext(ctx context.Context) SourceCodeTokenPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCodeTokenPtrOutput)
}

type SourceCodeTokenPtrInput interface {
	pulumi.Input

	ToSourceCodeTokenPtrOutput() SourceCodeTokenPtrOutput
	ToSourceCodeTokenPtrOutputWithContext(ctx context.Context) SourceCodeTokenPtrOutput
}

type sourceCodeTokenPtrType SourceCodeTokenArgs

func (*sourceCodeTokenPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceCodeToken)(nil))
}

func (i *sourceCodeTokenPtrType) ToSourceCodeTokenPtrOutput() SourceCodeTokenPtrOutput {
	return i.ToSourceCodeTokenPtrOutputWithContext(context.Background())
}

func (i *sourceCodeTokenPtrType) ToSourceCodeTokenPtrOutputWithContext(ctx context.Context) SourceCodeTokenPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCodeTokenOutput).ToSourceCodeTokenPtrOutput()
}

// SourceCodeTokenArrayInput is an input type that accepts SourceCodeTokenArray and SourceCodeTokenArrayOutput values.
// You can construct a concrete instance of `SourceCodeTokenArrayInput` via:
//
//          SourceCodeTokenArray{ SourceCodeTokenArgs{...} }
type SourceCodeTokenArrayInput interface {
	pulumi.Input

	ToSourceCodeTokenArrayOutput() SourceCodeTokenArrayOutput
	ToSourceCodeTokenArrayOutputWithContext(context.Context) SourceCodeTokenArrayOutput
}

type SourceCodeTokenArray []SourceCodeTokenInput

func (SourceCodeTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SourceCodeToken)(nil))
}

func (i SourceCodeTokenArray) ToSourceCodeTokenArrayOutput() SourceCodeTokenArrayOutput {
	return i.ToSourceCodeTokenArrayOutputWithContext(context.Background())
}

func (i SourceCodeTokenArray) ToSourceCodeTokenArrayOutputWithContext(ctx context.Context) SourceCodeTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCodeTokenArrayOutput)
}

// SourceCodeTokenMapInput is an input type that accepts SourceCodeTokenMap and SourceCodeTokenMapOutput values.
// You can construct a concrete instance of `SourceCodeTokenMapInput` via:
//
//          SourceCodeTokenMap{ "key": SourceCodeTokenArgs{...} }
type SourceCodeTokenMapInput interface {
	pulumi.Input

	ToSourceCodeTokenMapOutput() SourceCodeTokenMapOutput
	ToSourceCodeTokenMapOutputWithContext(context.Context) SourceCodeTokenMapOutput
}

type SourceCodeTokenMap map[string]SourceCodeTokenInput

func (SourceCodeTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SourceCodeToken)(nil))
}

func (i SourceCodeTokenMap) ToSourceCodeTokenMapOutput() SourceCodeTokenMapOutput {
	return i.ToSourceCodeTokenMapOutputWithContext(context.Background())
}

func (i SourceCodeTokenMap) ToSourceCodeTokenMapOutputWithContext(ctx context.Context) SourceCodeTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCodeTokenMapOutput)
}

type SourceCodeTokenOutput struct {
	*pulumi.OutputState
}

func (SourceCodeTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceCodeToken)(nil))
}

func (o SourceCodeTokenOutput) ToSourceCodeTokenOutput() SourceCodeTokenOutput {
	return o
}

func (o SourceCodeTokenOutput) ToSourceCodeTokenOutputWithContext(ctx context.Context) SourceCodeTokenOutput {
	return o
}

func (o SourceCodeTokenOutput) ToSourceCodeTokenPtrOutput() SourceCodeTokenPtrOutput {
	return o.ToSourceCodeTokenPtrOutputWithContext(context.Background())
}

func (o SourceCodeTokenOutput) ToSourceCodeTokenPtrOutputWithContext(ctx context.Context) SourceCodeTokenPtrOutput {
	return o.ApplyT(func(v SourceCodeToken) *SourceCodeToken {
		return &v
	}).(SourceCodeTokenPtrOutput)
}

type SourceCodeTokenPtrOutput struct {
	*pulumi.OutputState
}

func (SourceCodeTokenPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceCodeToken)(nil))
}

func (o SourceCodeTokenPtrOutput) ToSourceCodeTokenPtrOutput() SourceCodeTokenPtrOutput {
	return o
}

func (o SourceCodeTokenPtrOutput) ToSourceCodeTokenPtrOutputWithContext(ctx context.Context) SourceCodeTokenPtrOutput {
	return o
}

type SourceCodeTokenArrayOutput struct{ *pulumi.OutputState }

func (SourceCodeTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SourceCodeToken)(nil))
}

func (o SourceCodeTokenArrayOutput) ToSourceCodeTokenArrayOutput() SourceCodeTokenArrayOutput {
	return o
}

func (o SourceCodeTokenArrayOutput) ToSourceCodeTokenArrayOutputWithContext(ctx context.Context) SourceCodeTokenArrayOutput {
	return o
}

func (o SourceCodeTokenArrayOutput) Index(i pulumi.IntInput) SourceCodeTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SourceCodeToken {
		return vs[0].([]SourceCodeToken)[vs[1].(int)]
	}).(SourceCodeTokenOutput)
}

type SourceCodeTokenMapOutput struct{ *pulumi.OutputState }

func (SourceCodeTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SourceCodeToken)(nil))
}

func (o SourceCodeTokenMapOutput) ToSourceCodeTokenMapOutput() SourceCodeTokenMapOutput {
	return o
}

func (o SourceCodeTokenMapOutput) ToSourceCodeTokenMapOutputWithContext(ctx context.Context) SourceCodeTokenMapOutput {
	return o
}

func (o SourceCodeTokenMapOutput) MapIndex(k pulumi.StringInput) SourceCodeTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SourceCodeToken {
		return vs[0].(map[string]SourceCodeToken)[vs[1].(string)]
	}).(SourceCodeTokenOutput)
}

func init() {
	pulumi.RegisterOutputType(SourceCodeTokenOutput{})
	pulumi.RegisterOutputType(SourceCodeTokenPtrOutput{})
	pulumi.RegisterOutputType(SourceCodeTokenArrayOutput{})
	pulumi.RegisterOutputType(SourceCodeTokenMapOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package marketplace

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows accepting the Legal Terms for a Marketplace Image.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/marketplace"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := marketplace.NewAgreement(ctx, "barracuda", &marketplace.AgreementArgs{
// 			Offer:     pulumi.String("waf"),
// 			Plan:      pulumi.String("hourly"),
// 			Publisher: pulumi.String("barracudanetworks"),
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
// Marketplace Agreement can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:marketplace/agreement:Agreement example /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.MarketplaceOrdering/agreements/publisher1/offers/offer1/plans/plan1
// ```
type Agreement struct {
	pulumi.CustomResourceState

	LicenseTextLink pulumi.StringOutput `pulumi:"licenseTextLink"`
	// The Offer of the Marketplace Image. Changing this forces a new resource to be created.
	Offer pulumi.StringOutput `pulumi:"offer"`
	// The Plan of the Marketplace Image. Changing this forces a new resource to be created.
	Plan              pulumi.StringOutput `pulumi:"plan"`
	PrivacyPolicyLink pulumi.StringOutput `pulumi:"privacyPolicyLink"`
	// The Publisher of the Marketplace Image. Changing this forces a new resource to be created.
	Publisher pulumi.StringOutput `pulumi:"publisher"`
}

// NewAgreement registers a new resource with the given unique name, arguments, and options.
func NewAgreement(ctx *pulumi.Context,
	name string, args *AgreementArgs, opts ...pulumi.ResourceOption) (*Agreement, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Offer == nil {
		return nil, errors.New("invalid value for required argument 'Offer'")
	}
	if args.Plan == nil {
		return nil, errors.New("invalid value for required argument 'Plan'")
	}
	if args.Publisher == nil {
		return nil, errors.New("invalid value for required argument 'Publisher'")
	}
	var resource Agreement
	err := ctx.RegisterResource("azure:marketplace/agreement:Agreement", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgreement gets an existing Agreement resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgreement(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgreementState, opts ...pulumi.ResourceOption) (*Agreement, error) {
	var resource Agreement
	err := ctx.ReadResource("azure:marketplace/agreement:Agreement", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Agreement resources.
type agreementState struct {
	LicenseTextLink *string `pulumi:"licenseTextLink"`
	// The Offer of the Marketplace Image. Changing this forces a new resource to be created.
	Offer *string `pulumi:"offer"`
	// The Plan of the Marketplace Image. Changing this forces a new resource to be created.
	Plan              *string `pulumi:"plan"`
	PrivacyPolicyLink *string `pulumi:"privacyPolicyLink"`
	// The Publisher of the Marketplace Image. Changing this forces a new resource to be created.
	Publisher *string `pulumi:"publisher"`
}

type AgreementState struct {
	LicenseTextLink pulumi.StringPtrInput
	// The Offer of the Marketplace Image. Changing this forces a new resource to be created.
	Offer pulumi.StringPtrInput
	// The Plan of the Marketplace Image. Changing this forces a new resource to be created.
	Plan              pulumi.StringPtrInput
	PrivacyPolicyLink pulumi.StringPtrInput
	// The Publisher of the Marketplace Image. Changing this forces a new resource to be created.
	Publisher pulumi.StringPtrInput
}

func (AgreementState) ElementType() reflect.Type {
	return reflect.TypeOf((*agreementState)(nil)).Elem()
}

type agreementArgs struct {
	// The Offer of the Marketplace Image. Changing this forces a new resource to be created.
	Offer string `pulumi:"offer"`
	// The Plan of the Marketplace Image. Changing this forces a new resource to be created.
	Plan string `pulumi:"plan"`
	// The Publisher of the Marketplace Image. Changing this forces a new resource to be created.
	Publisher string `pulumi:"publisher"`
}

// The set of arguments for constructing a Agreement resource.
type AgreementArgs struct {
	// The Offer of the Marketplace Image. Changing this forces a new resource to be created.
	Offer pulumi.StringInput
	// The Plan of the Marketplace Image. Changing this forces a new resource to be created.
	Plan pulumi.StringInput
	// The Publisher of the Marketplace Image. Changing this forces a new resource to be created.
	Publisher pulumi.StringInput
}

func (AgreementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agreementArgs)(nil)).Elem()
}

type AgreementInput interface {
	pulumi.Input

	ToAgreementOutput() AgreementOutput
	ToAgreementOutputWithContext(ctx context.Context) AgreementOutput
}

func (*Agreement) ElementType() reflect.Type {
	return reflect.TypeOf((*Agreement)(nil))
}

func (i *Agreement) ToAgreementOutput() AgreementOutput {
	return i.ToAgreementOutputWithContext(context.Background())
}

func (i *Agreement) ToAgreementOutputWithContext(ctx context.Context) AgreementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgreementOutput)
}

func (i *Agreement) ToAgreementPtrOutput() AgreementPtrOutput {
	return i.ToAgreementPtrOutputWithContext(context.Background())
}

func (i *Agreement) ToAgreementPtrOutputWithContext(ctx context.Context) AgreementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgreementPtrOutput)
}

type AgreementPtrInput interface {
	pulumi.Input

	ToAgreementPtrOutput() AgreementPtrOutput
	ToAgreementPtrOutputWithContext(ctx context.Context) AgreementPtrOutput
}

type agreementPtrType AgreementArgs

func (*agreementPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Agreement)(nil))
}

func (i *agreementPtrType) ToAgreementPtrOutput() AgreementPtrOutput {
	return i.ToAgreementPtrOutputWithContext(context.Background())
}

func (i *agreementPtrType) ToAgreementPtrOutputWithContext(ctx context.Context) AgreementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgreementOutput).ToAgreementPtrOutput()
}

// AgreementArrayInput is an input type that accepts AgreementArray and AgreementArrayOutput values.
// You can construct a concrete instance of `AgreementArrayInput` via:
//
//          AgreementArray{ AgreementArgs{...} }
type AgreementArrayInput interface {
	pulumi.Input

	ToAgreementArrayOutput() AgreementArrayOutput
	ToAgreementArrayOutputWithContext(context.Context) AgreementArrayOutput
}

type AgreementArray []AgreementInput

func (AgreementArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Agreement)(nil))
}

func (i AgreementArray) ToAgreementArrayOutput() AgreementArrayOutput {
	return i.ToAgreementArrayOutputWithContext(context.Background())
}

func (i AgreementArray) ToAgreementArrayOutputWithContext(ctx context.Context) AgreementArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgreementArrayOutput)
}

// AgreementMapInput is an input type that accepts AgreementMap and AgreementMapOutput values.
// You can construct a concrete instance of `AgreementMapInput` via:
//
//          AgreementMap{ "key": AgreementArgs{...} }
type AgreementMapInput interface {
	pulumi.Input

	ToAgreementMapOutput() AgreementMapOutput
	ToAgreementMapOutputWithContext(context.Context) AgreementMapOutput
}

type AgreementMap map[string]AgreementInput

func (AgreementMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Agreement)(nil))
}

func (i AgreementMap) ToAgreementMapOutput() AgreementMapOutput {
	return i.ToAgreementMapOutputWithContext(context.Background())
}

func (i AgreementMap) ToAgreementMapOutputWithContext(ctx context.Context) AgreementMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgreementMapOutput)
}

type AgreementOutput struct {
	*pulumi.OutputState
}

func (AgreementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Agreement)(nil))
}

func (o AgreementOutput) ToAgreementOutput() AgreementOutput {
	return o
}

func (o AgreementOutput) ToAgreementOutputWithContext(ctx context.Context) AgreementOutput {
	return o
}

func (o AgreementOutput) ToAgreementPtrOutput() AgreementPtrOutput {
	return o.ToAgreementPtrOutputWithContext(context.Background())
}

func (o AgreementOutput) ToAgreementPtrOutputWithContext(ctx context.Context) AgreementPtrOutput {
	return o.ApplyT(func(v Agreement) *Agreement {
		return &v
	}).(AgreementPtrOutput)
}

type AgreementPtrOutput struct {
	*pulumi.OutputState
}

func (AgreementPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Agreement)(nil))
}

func (o AgreementPtrOutput) ToAgreementPtrOutput() AgreementPtrOutput {
	return o
}

func (o AgreementPtrOutput) ToAgreementPtrOutputWithContext(ctx context.Context) AgreementPtrOutput {
	return o
}

type AgreementArrayOutput struct{ *pulumi.OutputState }

func (AgreementArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Agreement)(nil))
}

func (o AgreementArrayOutput) ToAgreementArrayOutput() AgreementArrayOutput {
	return o
}

func (o AgreementArrayOutput) ToAgreementArrayOutputWithContext(ctx context.Context) AgreementArrayOutput {
	return o
}

func (o AgreementArrayOutput) Index(i pulumi.IntInput) AgreementOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Agreement {
		return vs[0].([]Agreement)[vs[1].(int)]
	}).(AgreementOutput)
}

type AgreementMapOutput struct{ *pulumi.OutputState }

func (AgreementMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Agreement)(nil))
}

func (o AgreementMapOutput) ToAgreementMapOutput() AgreementMapOutput {
	return o
}

func (o AgreementMapOutput) ToAgreementMapOutputWithContext(ctx context.Context) AgreementMapOutput {
	return o
}

func (o AgreementMapOutput) MapIndex(k pulumi.StringInput) AgreementOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Agreement {
		return vs[0].(map[string]Agreement)[vs[1].(string)]
	}).(AgreementOutput)
}

func init() {
	pulumi.RegisterOutputType(AgreementOutput{})
	pulumi.RegisterOutputType(AgreementPtrOutput{})
	pulumi.RegisterOutputType(AgreementArrayOutput{})
	pulumi.RegisterOutputType(AgreementMapOutput{})
}

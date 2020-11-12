// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an App Service Certificate Order.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/appservice"
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
// 		_, err = appservice.NewCertificateOrder(ctx, "exampleCertificateOrder", &appservice.CertificateOrderArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          pulumi.String("global"),
// 			DistinguishedName: pulumi.String("CN=example.com"),
// 			ProductType:       pulumi.String("Standard"),
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
// App Service Certificate Orders can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:appservice/certificateOrder:CertificateOrder example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.CertificateRegistration/certificateOrders/certificateorder1
// ```
type CertificateOrder struct {
	pulumi.CustomResourceState

	// Reasons why App Service Certificate is not renewable at the current moment.
	AppServiceCertificateNotRenewableReasons pulumi.StringArrayOutput `pulumi:"appServiceCertificateNotRenewableReasons"`
	// true if the certificate should be automatically renewed when it expires; otherwise, false. Defaults to true.
	AutoRenew pulumi.BoolPtrOutput `pulumi:"autoRenew"`
	// State of the Key Vault secret. A `certificates` block as defined below.
	Certificates CertificateOrderCertificateArrayOutput `pulumi:"certificates"`
	// Last CSR that was created for this order.
	Csr pulumi.StringOutput `pulumi:"csr"`
	// The Distinguished Name for the App Service Certificate Order.
	DistinguishedName pulumi.StringOutput `pulumi:"distinguishedName"`
	// Domain verification token.
	DomainVerificationToken pulumi.StringOutput `pulumi:"domainVerificationToken"`
	// Certificate expiration time.
	ExpirationTime pulumi.StringOutput `pulumi:"expirationTime"`
	// Certificate thumbprint intermediate certificate.
	IntermediateThumbprint pulumi.StringOutput `pulumi:"intermediateThumbprint"`
	// Whether the private key is external or not.
	IsPrivateKeyExternal pulumi.BoolOutput `pulumi:"isPrivateKeyExternal"`
	// Certificate key size.  Defaults to 2048.
	KeySize pulumi.IntPtrOutput `pulumi:"keySize"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. Currently the only valid value is `global`.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the certificate. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Certificate product type, such as `Standard` or `WildCard`.
	ProductType pulumi.StringPtrOutput `pulumi:"productType"`
	// The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Certificate thumbprint for root certificate.
	RootThumbprint pulumi.StringOutput `pulumi:"rootThumbprint"`
	// Certificate thumbprint for signed certificate.
	SignedCertificateThumbprint pulumi.StringOutput `pulumi:"signedCertificateThumbprint"`
	// Current order status.
	Status pulumi.StringOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Duration in years (must be between `1` and `3`).  Defaults to `1`.
	ValidityInYears pulumi.IntPtrOutput `pulumi:"validityInYears"`
}

// NewCertificateOrder registers a new resource with the given unique name, arguments, and options.
func NewCertificateOrder(ctx *pulumi.Context,
	name string, args *CertificateOrderArgs, opts ...pulumi.ResourceOption) (*CertificateOrder, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &CertificateOrderArgs{}
	}
	var resource CertificateOrder
	err := ctx.RegisterResource("azure:appservice/certificateOrder:CertificateOrder", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateOrder gets an existing CertificateOrder resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateOrder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateOrderState, opts ...pulumi.ResourceOption) (*CertificateOrder, error) {
	var resource CertificateOrder
	err := ctx.ReadResource("azure:appservice/certificateOrder:CertificateOrder", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateOrder resources.
type certificateOrderState struct {
	// Reasons why App Service Certificate is not renewable at the current moment.
	AppServiceCertificateNotRenewableReasons []string `pulumi:"appServiceCertificateNotRenewableReasons"`
	// true if the certificate should be automatically renewed when it expires; otherwise, false. Defaults to true.
	AutoRenew *bool `pulumi:"autoRenew"`
	// State of the Key Vault secret. A `certificates` block as defined below.
	Certificates []CertificateOrderCertificate `pulumi:"certificates"`
	// Last CSR that was created for this order.
	Csr *string `pulumi:"csr"`
	// The Distinguished Name for the App Service Certificate Order.
	DistinguishedName *string `pulumi:"distinguishedName"`
	// Domain verification token.
	DomainVerificationToken *string `pulumi:"domainVerificationToken"`
	// Certificate expiration time.
	ExpirationTime *string `pulumi:"expirationTime"`
	// Certificate thumbprint intermediate certificate.
	IntermediateThumbprint *string `pulumi:"intermediateThumbprint"`
	// Whether the private key is external or not.
	IsPrivateKeyExternal *bool `pulumi:"isPrivateKeyExternal"`
	// Certificate key size.  Defaults to 2048.
	KeySize *int `pulumi:"keySize"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. Currently the only valid value is `global`.
	Location *string `pulumi:"location"`
	// Specifies the name of the certificate. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Certificate product type, such as `Standard` or `WildCard`.
	ProductType *string `pulumi:"productType"`
	// The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Certificate thumbprint for root certificate.
	RootThumbprint *string `pulumi:"rootThumbprint"`
	// Certificate thumbprint for signed certificate.
	SignedCertificateThumbprint *string `pulumi:"signedCertificateThumbprint"`
	// Current order status.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Duration in years (must be between `1` and `3`).  Defaults to `1`.
	ValidityInYears *int `pulumi:"validityInYears"`
}

type CertificateOrderState struct {
	// Reasons why App Service Certificate is not renewable at the current moment.
	AppServiceCertificateNotRenewableReasons pulumi.StringArrayInput
	// true if the certificate should be automatically renewed when it expires; otherwise, false. Defaults to true.
	AutoRenew pulumi.BoolPtrInput
	// State of the Key Vault secret. A `certificates` block as defined below.
	Certificates CertificateOrderCertificateArrayInput
	// Last CSR that was created for this order.
	Csr pulumi.StringPtrInput
	// The Distinguished Name for the App Service Certificate Order.
	DistinguishedName pulumi.StringPtrInput
	// Domain verification token.
	DomainVerificationToken pulumi.StringPtrInput
	// Certificate expiration time.
	ExpirationTime pulumi.StringPtrInput
	// Certificate thumbprint intermediate certificate.
	IntermediateThumbprint pulumi.StringPtrInput
	// Whether the private key is external or not.
	IsPrivateKeyExternal pulumi.BoolPtrInput
	// Certificate key size.  Defaults to 2048.
	KeySize pulumi.IntPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. Currently the only valid value is `global`.
	Location pulumi.StringPtrInput
	// Specifies the name of the certificate. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Certificate product type, such as `Standard` or `WildCard`.
	ProductType pulumi.StringPtrInput
	// The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Certificate thumbprint for root certificate.
	RootThumbprint pulumi.StringPtrInput
	// Certificate thumbprint for signed certificate.
	SignedCertificateThumbprint pulumi.StringPtrInput
	// Current order status.
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Duration in years (must be between `1` and `3`).  Defaults to `1`.
	ValidityInYears pulumi.IntPtrInput
}

func (CertificateOrderState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateOrderState)(nil)).Elem()
}

type certificateOrderArgs struct {
	// true if the certificate should be automatically renewed when it expires; otherwise, false. Defaults to true.
	AutoRenew *bool `pulumi:"autoRenew"`
	// Last CSR that was created for this order.
	Csr *string `pulumi:"csr"`
	// The Distinguished Name for the App Service Certificate Order.
	DistinguishedName *string `pulumi:"distinguishedName"`
	// Certificate key size.  Defaults to 2048.
	KeySize *int `pulumi:"keySize"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. Currently the only valid value is `global`.
	Location *string `pulumi:"location"`
	// Specifies the name of the certificate. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Certificate product type, such as `Standard` or `WildCard`.
	ProductType *string `pulumi:"productType"`
	// The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Duration in years (must be between `1` and `3`).  Defaults to `1`.
	ValidityInYears *int `pulumi:"validityInYears"`
}

// The set of arguments for constructing a CertificateOrder resource.
type CertificateOrderArgs struct {
	// true if the certificate should be automatically renewed when it expires; otherwise, false. Defaults to true.
	AutoRenew pulumi.BoolPtrInput
	// Last CSR that was created for this order.
	Csr pulumi.StringPtrInput
	// The Distinguished Name for the App Service Certificate Order.
	DistinguishedName pulumi.StringPtrInput
	// Certificate key size.  Defaults to 2048.
	KeySize pulumi.IntPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. Currently the only valid value is `global`.
	Location pulumi.StringPtrInput
	// Specifies the name of the certificate. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Certificate product type, such as `Standard` or `WildCard`.
	ProductType pulumi.StringPtrInput
	// The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Duration in years (must be between `1` and `3`).  Defaults to `1`.
	ValidityInYears pulumi.IntPtrInput
}

func (CertificateOrderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateOrderArgs)(nil)).Elem()
}

type CertificateOrderInput interface {
	pulumi.Input

	ToCertificateOrderOutput() CertificateOrderOutput
	ToCertificateOrderOutputWithContext(ctx context.Context) CertificateOrderOutput
}

func (CertificateOrder) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateOrder)(nil)).Elem()
}

func (i CertificateOrder) ToCertificateOrderOutput() CertificateOrderOutput {
	return i.ToCertificateOrderOutputWithContext(context.Background())
}

func (i CertificateOrder) ToCertificateOrderOutputWithContext(ctx context.Context) CertificateOrderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOrderOutput)
}

type CertificateOrderOutput struct {
	*pulumi.OutputState
}

func (CertificateOrderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateOrderOutput)(nil)).Elem()
}

func (o CertificateOrderOutput) ToCertificateOrderOutput() CertificateOrderOutput {
	return o
}

func (o CertificateOrderOutput) ToCertificateOrderOutputWithContext(ctx context.Context) CertificateOrderOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CertificateOrderOutput{})
}

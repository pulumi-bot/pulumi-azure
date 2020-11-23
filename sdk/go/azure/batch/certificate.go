// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package batch

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a certificate in an Azure Batch account.
//
// ## Import
//
// Batch Certificates can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:batch/certificate:Certificate example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Batch/batchAccounts/batch1/certificates/certificate1
// ```
type Certificate struct {
	pulumi.CustomResourceState

	// Specifies the name of the Batch account. Changing this forces a new resource to be created.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// The base64-encoded contents of the certificate.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// The format of the certificate. Possible values are `Cer` or `Pfx`.
	Format pulumi.StringOutput `pulumi:"format"`
	// The generated name of the certificate.
	Name pulumi.StringOutput `pulumi:"name"`
	// The password to access the certificate's private key. This must and can only be specified when `format` is `Pfx`.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The public key of the certificate.
	PublicData pulumi.StringOutput `pulumi:"publicData"`
	// The name of the resource group in which to create the Batch account. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The thumbprint of the certificate. At this time the only supported value is 'SHA1'.
	Thumbprint          pulumi.StringOutput `pulumi:"thumbprint"`
	ThumbprintAlgorithm pulumi.StringOutput `pulumi:"thumbprintAlgorithm"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Certificate == nil {
		return nil, errors.New("invalid value for required argument 'Certificate'")
	}
	if args.Format == nil {
		return nil, errors.New("invalid value for required argument 'Format'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Thumbprint == nil {
		return nil, errors.New("invalid value for required argument 'Thumbprint'")
	}
	if args.ThumbprintAlgorithm == nil {
		return nil, errors.New("invalid value for required argument 'ThumbprintAlgorithm'")
	}
	var resource Certificate
	err := ctx.RegisterResource("azure:batch/certificate:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("azure:batch/certificate:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
	// Specifies the name of the Batch account. Changing this forces a new resource to be created.
	AccountName *string `pulumi:"accountName"`
	// The base64-encoded contents of the certificate.
	Certificate *string `pulumi:"certificate"`
	// The format of the certificate. Possible values are `Cer` or `Pfx`.
	Format *string `pulumi:"format"`
	// The generated name of the certificate.
	Name *string `pulumi:"name"`
	// The password to access the certificate's private key. This must and can only be specified when `format` is `Pfx`.
	Password *string `pulumi:"password"`
	// The public key of the certificate.
	PublicData *string `pulumi:"publicData"`
	// The name of the resource group in which to create the Batch account. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The thumbprint of the certificate. At this time the only supported value is 'SHA1'.
	Thumbprint          *string `pulumi:"thumbprint"`
	ThumbprintAlgorithm *string `pulumi:"thumbprintAlgorithm"`
}

type CertificateState struct {
	// Specifies the name of the Batch account. Changing this forces a new resource to be created.
	AccountName pulumi.StringPtrInput
	// The base64-encoded contents of the certificate.
	Certificate pulumi.StringPtrInput
	// The format of the certificate. Possible values are `Cer` or `Pfx`.
	Format pulumi.StringPtrInput
	// The generated name of the certificate.
	Name pulumi.StringPtrInput
	// The password to access the certificate's private key. This must and can only be specified when `format` is `Pfx`.
	Password pulumi.StringPtrInput
	// The public key of the certificate.
	PublicData pulumi.StringPtrInput
	// The name of the resource group in which to create the Batch account. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The thumbprint of the certificate. At this time the only supported value is 'SHA1'.
	Thumbprint          pulumi.StringPtrInput
	ThumbprintAlgorithm pulumi.StringPtrInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// Specifies the name of the Batch account. Changing this forces a new resource to be created.
	AccountName string `pulumi:"accountName"`
	// The base64-encoded contents of the certificate.
	Certificate string `pulumi:"certificate"`
	// The format of the certificate. Possible values are `Cer` or `Pfx`.
	Format string `pulumi:"format"`
	// The password to access the certificate's private key. This must and can only be specified when `format` is `Pfx`.
	Password *string `pulumi:"password"`
	// The name of the resource group in which to create the Batch account. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The thumbprint of the certificate. At this time the only supported value is 'SHA1'.
	Thumbprint          string `pulumi:"thumbprint"`
	ThumbprintAlgorithm string `pulumi:"thumbprintAlgorithm"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// Specifies the name of the Batch account. Changing this forces a new resource to be created.
	AccountName pulumi.StringInput
	// The base64-encoded contents of the certificate.
	Certificate pulumi.StringInput
	// The format of the certificate. Possible values are `Cer` or `Pfx`.
	Format pulumi.StringInput
	// The password to access the certificate's private key. This must and can only be specified when `format` is `Pfx`.
	Password pulumi.StringPtrInput
	// The name of the resource group in which to create the Batch account. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The thumbprint of the certificate. At this time the only supported value is 'SHA1'.
	Thumbprint          pulumi.StringInput
	ThumbprintAlgorithm pulumi.StringInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil)).Elem()
}

func (i Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

type CertificateOutput struct {
	*pulumi.OutputState
}

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateOutput)(nil)).Elem()
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CertificateOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keyvault

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Key Vault Certificate.
//
// ## Example Usage
// ### Generating A New Certificate)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/keyvault"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleKeyVault, err := keyvault.NewKeyVault(ctx, "exampleKeyVault", &keyvault.KeyVaultArgs{
// 			Location:                exampleResourceGroup.Location,
// 			ResourceGroupName:       exampleResourceGroup.Name,
// 			TenantId:                pulumi.String(current.TenantId),
// 			SkuName:                 pulumi.String("standard"),
// 			SoftDeleteEnabled:       pulumi.Bool(true),
// 			SoftDeleteRetentionDays: pulumi.Int(7),
// 			AccessPolicies: keyvault.KeyVaultAccessPolicyArray{
// 				&keyvault.KeyVaultAccessPolicyArgs{
// 					TenantId: pulumi.String(current.TenantId),
// 					ObjectId: pulumi.String(current.ObjectId),
// 					CertificatePermissions: pulumi.StringArray{
// 						pulumi.String("create"),
// 						pulumi.String("delete"),
// 						pulumi.String("deleteissuers"),
// 						pulumi.String("get"),
// 						pulumi.String("getissuers"),
// 						pulumi.String("import"),
// 						pulumi.String("list"),
// 						pulumi.String("listissuers"),
// 						pulumi.String("managecontacts"),
// 						pulumi.String("manageissuers"),
// 						pulumi.String("purge"),
// 						pulumi.String("setissuers"),
// 						pulumi.String("update"),
// 					},
// 					KeyPermissions: pulumi.StringArray{
// 						pulumi.String("backup"),
// 						pulumi.String("create"),
// 						pulumi.String("decrypt"),
// 						pulumi.String("delete"),
// 						pulumi.String("encrypt"),
// 						pulumi.String("get"),
// 						pulumi.String("import"),
// 						pulumi.String("list"),
// 						pulumi.String("purge"),
// 						pulumi.String("recover"),
// 						pulumi.String("restore"),
// 						pulumi.String("sign"),
// 						pulumi.String("unwrapKey"),
// 						pulumi.String("update"),
// 						pulumi.String("verify"),
// 						pulumi.String("wrapKey"),
// 					},
// 					SecretPermissions: pulumi.StringArray{
// 						pulumi.String("backup"),
// 						pulumi.String("delete"),
// 						pulumi.String("get"),
// 						pulumi.String("list"),
// 						pulumi.String("purge"),
// 						pulumi.String("recover"),
// 						pulumi.String("restore"),
// 						pulumi.String("set"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = keyvault.NewCertificate(ctx, "exampleCertificate", &keyvault.CertificateArgs{
// 			KeyVaultId: exampleKeyVault.ID(),
// 			CertificatePolicy: &keyvault.CertificateCertificatePolicyArgs{
// 				IssuerParameters: &keyvault.CertificateCertificatePolicyIssuerParametersArgs{
// 					Name: pulumi.String("Self"),
// 				},
// 				KeyProperties: &keyvault.CertificateCertificatePolicyKeyPropertiesArgs{
// 					Exportable: pulumi.Bool(true),
// 					KeySize:    pulumi.Int(2048),
// 					KeyType:    pulumi.String("RSA"),
// 					ReuseKey:   pulumi.Bool(true),
// 				},
// 				LifetimeActions: keyvault.CertificateCertificatePolicyLifetimeActionArray{
// 					&keyvault.CertificateCertificatePolicyLifetimeActionArgs{
// 						Action: &keyvault.CertificateCertificatePolicyLifetimeActionActionArgs{
// 							ActionType: pulumi.String("AutoRenew"),
// 						},
// 						Trigger: &keyvault.CertificateCertificatePolicyLifetimeActionTriggerArgs{
// 							DaysBeforeExpiry: pulumi.Int(30),
// 						},
// 					},
// 				},
// 				SecretProperties: &keyvault.CertificateCertificatePolicySecretPropertiesArgs{
// 					ContentType: pulumi.String("application/x-pkcs12"),
// 				},
// 				X509CertificateProperties: &keyvault.CertificateCertificatePolicyX509CertificatePropertiesArgs{
// 					ExtendedKeyUsages: pulumi.StringArray{
// 						pulumi.String("1.3.6.1.5.5.7.3.1"),
// 					},
// 					KeyUsages: pulumi.StringArray{
// 						pulumi.String("cRLSign"),
// 						pulumi.String("dataEncipherment"),
// 						pulumi.String("digitalSignature"),
// 						pulumi.String("keyAgreement"),
// 						pulumi.String("keyCertSign"),
// 						pulumi.String("keyEncipherment"),
// 					},
// 					SubjectAlternativeNames: &keyvault.CertificateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesArgs{
// 						DnsNames: pulumi.StringArray{
// 							pulumi.String("internal.contoso.com"),
// 							pulumi.String("domain.hello.world"),
// 						},
// 					},
// 					Subject:          pulumi.String("CN=hello-world"),
// 					ValidityInMonths: pulumi.Int(12),
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
// Key Vault Certificates can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:keyvault/certificate:Certificate example "https://example-keyvault.vault.azure.net/certificates/example/fdf067c93bbb4b22bff4d8b7a9a56217"
// ```
type Certificate struct {
	pulumi.CustomResourceState

	// A `certificate` block as defined below, used to Import an existing certificate.
	Certificate CertificateCertificatePtrOutput `pulumi:"certificate"`
	// A `certificateAttribute` block as defined below.
	CertificateAttributes CertificateCertificateAttributeArrayOutput `pulumi:"certificateAttributes"`
	// The raw Key Vault Certificate data represented as a hexadecimal string.
	CertificateData pulumi.StringOutput `pulumi:"certificateData"`
	// A `certificatePolicy` block as defined below.
	CertificatePolicy CertificateCertificatePolicyOutput `pulumi:"certificatePolicy"`
	// The ID of the Key Vault where the Certificate should be created.
	KeyVaultId pulumi.StringOutput `pulumi:"keyVaultId"`
	// Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the associated Key Vault Secret.
	SecretId pulumi.StringOutput `pulumi:"secretId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`
	// The current version of the Key Vault Certificate.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificatePolicy == nil {
		return nil, errors.New("invalid value for required argument 'CertificatePolicy'")
	}
	if args.KeyVaultId == nil {
		return nil, errors.New("invalid value for required argument 'KeyVaultId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure:keyvault/certifiate:Certifiate"),
		},
	})
	opts = append(opts, aliases)
	var resource Certificate
	err := ctx.RegisterResource("azure:keyvault/certificate:Certificate", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure:keyvault/certificate:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
	// A `certificate` block as defined below, used to Import an existing certificate.
	Certificate *CertificateCertificate `pulumi:"certificate"`
	// A `certificateAttribute` block as defined below.
	CertificateAttributes []CertificateCertificateAttribute `pulumi:"certificateAttributes"`
	// The raw Key Vault Certificate data represented as a hexadecimal string.
	CertificateData *string `pulumi:"certificateData"`
	// A `certificatePolicy` block as defined below.
	CertificatePolicy *CertificateCertificatePolicy `pulumi:"certificatePolicy"`
	// The ID of the Key Vault where the Certificate should be created.
	KeyVaultId *string `pulumi:"keyVaultId"`
	// Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the associated Key Vault Secret.
	SecretId *string `pulumi:"secretId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.
	Thumbprint *string `pulumi:"thumbprint"`
	// The current version of the Key Vault Certificate.
	Version *string `pulumi:"version"`
}

type CertificateState struct {
	// A `certificate` block as defined below, used to Import an existing certificate.
	Certificate CertificateCertificatePtrInput
	// A `certificateAttribute` block as defined below.
	CertificateAttributes CertificateCertificateAttributeArrayInput
	// The raw Key Vault Certificate data represented as a hexadecimal string.
	CertificateData pulumi.StringPtrInput
	// A `certificatePolicy` block as defined below.
	CertificatePolicy CertificateCertificatePolicyPtrInput
	// The ID of the Key Vault where the Certificate should be created.
	KeyVaultId pulumi.StringPtrInput
	// Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the associated Key Vault Secret.
	SecretId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.
	Thumbprint pulumi.StringPtrInput
	// The current version of the Key Vault Certificate.
	Version pulumi.StringPtrInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// A `certificate` block as defined below, used to Import an existing certificate.
	Certificate *CertificateCertificate `pulumi:"certificate"`
	// A `certificatePolicy` block as defined below.
	CertificatePolicy CertificateCertificatePolicy `pulumi:"certificatePolicy"`
	// The ID of the Key Vault where the Certificate should be created.
	KeyVaultId string `pulumi:"keyVaultId"`
	// Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// A `certificate` block as defined below, used to Import an existing certificate.
	Certificate CertificateCertificatePtrInput
	// A `certificatePolicy` block as defined below.
	CertificatePolicy CertificateCertificatePolicyInput
	// The ID of the Key Vault where the Certificate should be created.
	KeyVaultId pulumi.StringInput
	// Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil))
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

func (i *Certificate) ToCertificatePtrOutput() CertificatePtrOutput {
	return i.ToCertificatePtrOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificatePtrOutputWithContext(ctx context.Context) CertificatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePtrOutput)
}

type CertificatePtrInput interface {
	pulumi.Input

	ToCertificatePtrOutput() CertificatePtrOutput
	ToCertificatePtrOutputWithContext(ctx context.Context) CertificatePtrOutput
}

type certificatePtrType CertificateArgs

func (*certificatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil))
}

func (i *certificatePtrType) ToCertificatePtrOutput() CertificatePtrOutput {
	return i.ToCertificatePtrOutputWithContext(context.Background())
}

func (i *certificatePtrType) ToCertificatePtrOutputWithContext(ctx context.Context) CertificatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePtrOutput)
}

// CertificateArrayInput is an input type that accepts CertificateArray and CertificateArrayOutput values.
// You can construct a concrete instance of `CertificateArrayInput` via:
//
//          CertificateArray{ CertificateArgs{...} }
type CertificateArrayInput interface {
	pulumi.Input

	ToCertificateArrayOutput() CertificateArrayOutput
	ToCertificateArrayOutputWithContext(context.Context) CertificateArrayOutput
}

type CertificateArray []CertificateInput

func (CertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Certificate)(nil))
}

func (i CertificateArray) ToCertificateArrayOutput() CertificateArrayOutput {
	return i.ToCertificateArrayOutputWithContext(context.Background())
}

func (i CertificateArray) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateArrayOutput)
}

// CertificateMapInput is an input type that accepts CertificateMap and CertificateMapOutput values.
// You can construct a concrete instance of `CertificateMapInput` via:
//
//          CertificateMap{ "key": CertificateArgs{...} }
type CertificateMapInput interface {
	pulumi.Input

	ToCertificateMapOutput() CertificateMapOutput
	ToCertificateMapOutputWithContext(context.Context) CertificateMapOutput
}

type CertificateMap map[string]CertificateInput

func (CertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Certificate)(nil))
}

func (i CertificateMap) ToCertificateMapOutput() CertificateMapOutput {
	return i.ToCertificateMapOutputWithContext(context.Background())
}

func (i CertificateMap) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateMapOutput)
}

type CertificateOutput struct {
	*pulumi.OutputState
}

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil))
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificatePtrOutput() CertificatePtrOutput {
	return o.ToCertificatePtrOutputWithContext(context.Background())
}

func (o CertificateOutput) ToCertificatePtrOutputWithContext(ctx context.Context) CertificatePtrOutput {
	return o.ApplyT(func(v Certificate) *Certificate {
		return &v
	}).(CertificatePtrOutput)
}

type CertificatePtrOutput struct {
	*pulumi.OutputState
}

func (CertificatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil))
}

func (o CertificatePtrOutput) ToCertificatePtrOutput() CertificatePtrOutput {
	return o
}

func (o CertificatePtrOutput) ToCertificatePtrOutputWithContext(ctx context.Context) CertificatePtrOutput {
	return o
}

type CertificateArrayOutput struct{ *pulumi.OutputState }

func (CertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Certificate)(nil))
}

func (o CertificateArrayOutput) ToCertificateArrayOutput() CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) Index(i pulumi.IntInput) CertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Certificate {
		return vs[0].([]Certificate)[vs[1].(int)]
	}).(CertificateOutput)
}

type CertificateMapOutput struct{ *pulumi.OutputState }

func (CertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Certificate)(nil))
}

func (o CertificateMapOutput) ToCertificateMapOutput() CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) MapIndex(k pulumi.StringInput) CertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Certificate {
		return vs[0].(map[string]Certificate)[vs[1].(string)]
	}).(CertificateOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificateOutput{})
	pulumi.RegisterOutputType(CertificatePtrOutput{})
	pulumi.RegisterOutputType(CertificateArrayOutput{})
	pulumi.RegisterOutputType(CertificateMapOutput{})
}

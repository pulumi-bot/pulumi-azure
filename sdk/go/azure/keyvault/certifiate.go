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
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			TenantId:          pulumi.String(current.TenantId),
// 			SkuName:           pulumi.String("standard"),
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
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("Production"),
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
// Deprecated: azure.keyvault.Certifiate has been deprecated in favor of azure.keyvault.Certificate
type Certifiate struct {
	pulumi.CustomResourceState

	// A `certificate` block as defined below, used to Import an existing certificate.
	Certificate CertifiateCertificatePtrOutput `pulumi:"certificate"`
	// A `certificateAttribute` block as defined below.
	CertificateAttributes CertifiateCertificateAttributeArrayOutput `pulumi:"certificateAttributes"`
	// The raw Key Vault Certificate data represented as a hexadecimal string.
	CertificateData pulumi.StringOutput `pulumi:"certificateData"`
	// A `certificatePolicy` block as defined below.
	CertificatePolicy CertifiateCertificatePolicyOutput `pulumi:"certificatePolicy"`
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

// NewCertifiate registers a new resource with the given unique name, arguments, and options.
func NewCertifiate(ctx *pulumi.Context,
	name string, args *CertifiateArgs, opts ...pulumi.ResourceOption) (*Certifiate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificatePolicy == nil {
		return nil, errors.New("invalid value for required argument 'CertificatePolicy'")
	}
	if args.KeyVaultId == nil {
		return nil, errors.New("invalid value for required argument 'KeyVaultId'")
	}
	var resource Certifiate
	err := ctx.RegisterResource("azure:keyvault/certifiate:Certifiate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertifiate gets an existing Certifiate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertifiate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertifiateState, opts ...pulumi.ResourceOption) (*Certifiate, error) {
	var resource Certifiate
	err := ctx.ReadResource("azure:keyvault/certifiate:Certifiate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certifiate resources.
type certifiateState struct {
	// A `certificate` block as defined below, used to Import an existing certificate.
	Certificate *CertifiateCertificate `pulumi:"certificate"`
	// A `certificateAttribute` block as defined below.
	CertificateAttributes []CertifiateCertificateAttribute `pulumi:"certificateAttributes"`
	// The raw Key Vault Certificate data represented as a hexadecimal string.
	CertificateData *string `pulumi:"certificateData"`
	// A `certificatePolicy` block as defined below.
	CertificatePolicy *CertifiateCertificatePolicy `pulumi:"certificatePolicy"`
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

type CertifiateState struct {
	// A `certificate` block as defined below, used to Import an existing certificate.
	Certificate CertifiateCertificatePtrInput
	// A `certificateAttribute` block as defined below.
	CertificateAttributes CertifiateCertificateAttributeArrayInput
	// The raw Key Vault Certificate data represented as a hexadecimal string.
	CertificateData pulumi.StringPtrInput
	// A `certificatePolicy` block as defined below.
	CertificatePolicy CertifiateCertificatePolicyPtrInput
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

func (CertifiateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certifiateState)(nil)).Elem()
}

type certifiateArgs struct {
	// A `certificate` block as defined below, used to Import an existing certificate.
	Certificate *CertifiateCertificate `pulumi:"certificate"`
	// A `certificatePolicy` block as defined below.
	CertificatePolicy CertifiateCertificatePolicy `pulumi:"certificatePolicy"`
	// The ID of the Key Vault where the Certificate should be created.
	KeyVaultId string `pulumi:"keyVaultId"`
	// Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Certifiate resource.
type CertifiateArgs struct {
	// A `certificate` block as defined below, used to Import an existing certificate.
	Certificate CertifiateCertificatePtrInput
	// A `certificatePolicy` block as defined below.
	CertificatePolicy CertifiateCertificatePolicyInput
	// The ID of the Key Vault where the Certificate should be created.
	KeyVaultId pulumi.StringInput
	// Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (CertifiateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certifiateArgs)(nil)).Elem()
}

type CertifiateInput interface {
	pulumi.Input

	ToCertifiateOutput() CertifiateOutput
	ToCertifiateOutputWithContext(ctx context.Context) CertifiateOutput
}

func (Certifiate) ElementType() reflect.Type {
	return reflect.TypeOf((*Certifiate)(nil)).Elem()
}

func (i Certifiate) ToCertifiateOutput() CertifiateOutput {
	return i.ToCertifiateOutputWithContext(context.Background())
}

func (i Certifiate) ToCertifiateOutputWithContext(ctx context.Context) CertifiateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertifiateOutput)
}

type CertifiateOutput struct {
	*pulumi.OutputState
}

func (CertifiateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertifiateOutput)(nil)).Elem()
}

func (o CertifiateOutput) ToCertifiateOutput() CertifiateOutput {
	return o
}

func (o CertifiateOutput) ToCertifiateOutputWithContext(ctx context.Context) CertifiateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CertifiateOutput{})
}

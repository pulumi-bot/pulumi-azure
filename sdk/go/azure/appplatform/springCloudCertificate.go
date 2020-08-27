// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure Spring Cloud Certificate.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/appplatform"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/keyvault"
// 	"github.com/pulumi/pulumi-azuread/sdk/v2/go/azuread"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("Southeast Asia"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt0 := "Azure Spring Cloud Domain-Management"
// 		exampleServicePrincipal, err := azuread.LookupServicePrincipal(ctx, &azuread.LookupServicePrincipalArgs{
// 			DisplayName: &opt0,
// 		}, nil)
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
// 					SecretPermissions: pulumi.StringArray{
// 						pulumi.String("set"),
// 					},
// 					CertificatePermissions: pulumi.StringArray{
// 						pulumi.String("create"),
// 						pulumi.String("delete"),
// 						pulumi.String("get"),
// 						pulumi.String("update"),
// 					},
// 				},
// 				&keyvault.KeyVaultAccessPolicyArgs{
// 					TenantId: pulumi.String(current.TenantId),
// 					ObjectId: pulumi.String(exampleServicePrincipal.ObjectId),
// 					SecretPermissions: pulumi.StringArray{
// 						pulumi.String("get"),
// 						pulumi.String("list"),
// 					},
// 					CertificatePermissions: pulumi.StringArray{
// 						pulumi.String("get"),
// 						pulumi.String("list"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleCertificate, err := keyvault.NewCertificate(ctx, "exampleCertificate", &keyvault.CertificateArgs{
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
// 					KeyUsages: pulumi.StringArray{
// 						pulumi.String("cRLSign"),
// 						pulumi.String("dataEncipherment"),
// 						pulumi.String("digitalSignature"),
// 						pulumi.String("keyAgreement"),
// 						pulumi.String("keyCertSign"),
// 						pulumi.String("keyEncipherment"),
// 					},
// 					Subject:          pulumi.String("CN=contoso.com"),
// 					ValidityInMonths: pulumi.Int(12),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSpringCloudService, err := appplatform.NewSpringCloudService(ctx, "exampleSpringCloudService", &appplatform.SpringCloudServiceArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appplatform.NewSpringCloudCertificate(ctx, "exampleSpringCloudCertificate", &appplatform.SpringCloudCertificateArgs{
// 			ResourceGroupName:     exampleSpringCloudService.ResourceGroupName,
// 			ServiceName:           exampleSpringCloudService.Name,
// 			KeyVaultCertificateId: exampleCertificate.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SpringCloudCertificate struct {
	pulumi.CustomResourceState

	// Specifies the ID of the Key Vault Certificate resource. Changing this forces a new resource to be created.
	KeyVaultCertificateId pulumi.StringOutput `pulumi:"keyVaultCertificateId"`
	// Specifies the name of the Spring Cloud Certificate. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the resource group in which to create the Spring Cloud Certificate. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewSpringCloudCertificate registers a new resource with the given unique name, arguments, and options.
func NewSpringCloudCertificate(ctx *pulumi.Context,
	name string, args *SpringCloudCertificateArgs, opts ...pulumi.ResourceOption) (*SpringCloudCertificate, error) {
	if args == nil || args.KeyVaultCertificateId == nil {
		return nil, errors.New("missing required argument 'KeyVaultCertificateId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &SpringCloudCertificateArgs{}
	}
	var resource SpringCloudCertificate
	err := ctx.RegisterResource("azure:appplatform/springCloudCertificate:SpringCloudCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpringCloudCertificate gets an existing SpringCloudCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpringCloudCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpringCloudCertificateState, opts ...pulumi.ResourceOption) (*SpringCloudCertificate, error) {
	var resource SpringCloudCertificate
	err := ctx.ReadResource("azure:appplatform/springCloudCertificate:SpringCloudCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpringCloudCertificate resources.
type springCloudCertificateState struct {
	// Specifies the ID of the Key Vault Certificate resource. Changing this forces a new resource to be created.
	KeyVaultCertificateId *string `pulumi:"keyVaultCertificateId"`
	// Specifies the name of the Spring Cloud Certificate. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the resource group in which to create the Spring Cloud Certificate. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
	ServiceName *string `pulumi:"serviceName"`
}

type SpringCloudCertificateState struct {
	// Specifies the ID of the Key Vault Certificate resource. Changing this forces a new resource to be created.
	KeyVaultCertificateId pulumi.StringPtrInput
	// Specifies the name of the Spring Cloud Certificate. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the resource group in which to create the Spring Cloud Certificate. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
	ServiceName pulumi.StringPtrInput
}

func (SpringCloudCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*springCloudCertificateState)(nil)).Elem()
}

type springCloudCertificateArgs struct {
	// Specifies the ID of the Key Vault Certificate resource. Changing this forces a new resource to be created.
	KeyVaultCertificateId string `pulumi:"keyVaultCertificateId"`
	// Specifies the name of the Spring Cloud Certificate. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the resource group in which to create the Spring Cloud Certificate. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a SpringCloudCertificate resource.
type SpringCloudCertificateArgs struct {
	// Specifies the ID of the Key Vault Certificate resource. Changing this forces a new resource to be created.
	KeyVaultCertificateId pulumi.StringInput
	// Specifies the name of the Spring Cloud Certificate. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the resource group in which to create the Spring Cloud Certificate. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
	ServiceName pulumi.StringInput
}

func (SpringCloudCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*springCloudCertificateArgs)(nil)).Elem()
}

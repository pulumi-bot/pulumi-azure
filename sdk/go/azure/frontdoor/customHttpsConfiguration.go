// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package frontdoor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages the Custom Https Configuration for an Azure Front Door Frontend Endpoint..
//
// > **NOTE:** Custom https configurations for a Front Door Frontend Endpoint can be defined both within the `frontdoor.Frontdoor` resource via the `customHttpsConfiguration` block and by using a separate resource, as described in the following sections.
//
// > **NOTE:** Defining custom https configurations using a separate `frontdoor.CustomHttpsConfiguration` resource allows for parallel creation/update.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/frontdoor"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/keyvault"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("EastUS2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		vault, err := keyvault.LookupKeyVault(ctx, &keyvault.LookupKeyVaultArgs{
// 			Name:              "example-vault",
// 			ResourceGroupName: "example-vault-rg",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleFrontdoor, err := frontdoor.NewFrontdoor(ctx, "exampleFrontdoor", &frontdoor.FrontdoorArgs{
// 			ResourceGroupName:                       exampleResourceGroup.Name,
// 			EnforceBackendPoolsCertificateNameCheck: pulumi.Bool(false),
// 			RoutingRules: frontdoor.FrontdoorRoutingRuleArray{
// 				&frontdoor.FrontdoorRoutingRuleArgs{
// 					Name: pulumi.String("exampleRoutingRule1"),
// 					AcceptedProtocols: pulumi.StringArray{
// 						pulumi.String("Http"),
// 						pulumi.String("Https"),
// 					},
// 					PatternsToMatches: pulumi.StringArray{
// 						pulumi.String("/*"),
// 					},
// 					FrontendEndpoints: pulumi.StringArray{
// 						pulumi.String("exampleFrontendEndpoint1"),
// 					},
// 					ForwardingConfiguration: &frontdoor.FrontdoorRoutingRuleForwardingConfigurationArgs{
// 						ForwardingProtocol: pulumi.String("MatchRequest"),
// 						BackendPoolName:    pulumi.String("exampleBackendBing"),
// 					},
// 				},
// 			},
// 			BackendPoolLoadBalancings: frontdoor.FrontdoorBackendPoolLoadBalancingArray{
// 				&frontdoor.FrontdoorBackendPoolLoadBalancingArgs{
// 					Name: pulumi.String("exampleLoadBalancingSettings1"),
// 				},
// 			},
// 			BackendPoolHealthProbes: frontdoor.FrontdoorBackendPoolHealthProbeArray{
// 				&frontdoor.FrontdoorBackendPoolHealthProbeArgs{
// 					Name: pulumi.String("exampleHealthProbeSetting1"),
// 				},
// 			},
// 			BackendPools: frontdoor.FrontdoorBackendPoolArray{
// 				&frontdoor.FrontdoorBackendPoolArgs{
// 					Name: pulumi.String("exampleBackendBing"),
// 					Backends: frontdoor.FrontdoorBackendPoolBackendArray{
// 						&frontdoor.FrontdoorBackendPoolBackendArgs{
// 							HostHeader: pulumi.String("www.bing.com"),
// 							Address:    pulumi.String("www.bing.com"),
// 							HttpPort:   pulumi.Int(80),
// 							HttpsPort:  pulumi.Int(443),
// 						},
// 					},
// 					LoadBalancingName: pulumi.String("exampleLoadBalancingSettings1"),
// 					HealthProbeName:   pulumi.String("exampleHealthProbeSetting1"),
// 				},
// 			},
// 			FrontendEndpoints: frontdoor.FrontdoorFrontendEndpointArray{
// 				&frontdoor.FrontdoorFrontendEndpointArgs{
// 					Name:     pulumi.String("exampleFrontendEndpoint1"),
// 					HostName: pulumi.String("example-FrontDoor.azurefd.net"),
// 				},
// 				&frontdoor.FrontdoorFrontendEndpointArgs{
// 					Name:     pulumi.String("exampleFrontendEndpoint2"),
// 					HostName: pulumi.String("examplefd1.examplefd.net"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = frontdoor.NewCustomHttpsConfiguration(ctx, "exampleCustomHttps0", &frontdoor.CustomHttpsConfigurationArgs{
// 			FrontendEndpointId: pulumi.String(exampleFrontdoor.FrontendEndpoints.ApplyT(func(frontendEndpoints []frontdoor.FrontdoorFrontendEndpoint) (string, error) {
// 				return frontendEndpoints[0].Id, nil
// 			}).(pulumi.StringOutput)),
// 			CustomHttpsProvisioningEnabled: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = frontdoor.NewCustomHttpsConfiguration(ctx, "exampleCustomHttps1", &frontdoor.CustomHttpsConfigurationArgs{
// 			FrontendEndpointId: pulumi.String(exampleFrontdoor.FrontendEndpoints.ApplyT(func(frontendEndpoints []frontdoor.FrontdoorFrontendEndpoint) (string, error) {
// 				return frontendEndpoints[1].Id, nil
// 			}).(pulumi.StringOutput)),
// 			CustomHttpsProvisioningEnabled: pulumi.Bool(true),
// 			CustomHttpsConfiguration: &frontdoor.CustomHttpsConfigurationCustomHttpsConfigurationArgs{
// 				CertificateSource:                     pulumi.String("AzureKeyVault"),
// 				AzureKeyVaultCertificateSecretName:    pulumi.String("examplefd1"),
// 				AzureKeyVaultCertificateSecretVersion: pulumi.String("ec8d0737e0df4f4gb52ecea858e97a73"),
// 				AzureKeyVaultCertificateVaultId:       pulumi.String(vault.Id),
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
// Front Door Custom Https Configurations can be imported using the `resource id` of the Frontend Endpoint, e.g.
//
// ```sh
//  $ pulumi import azure:frontdoor/customHttpsConfiguration:CustomHttpsConfiguration example_custom_https_1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/frontdoors/frontdoor1/frontendEndpoints/endpoint1
// ```
type CustomHttpsConfiguration struct {
	pulumi.CustomResourceState

	// A `customHttpsConfiguration` block as defined below.
	CustomHttpsConfiguration CustomHttpsConfigurationCustomHttpsConfigurationPtrOutput `pulumi:"customHttpsConfiguration"`
	// Should the HTTPS protocol be enabled for this custom domain associated with the Front Door?
	CustomHttpsProvisioningEnabled pulumi.BoolOutput `pulumi:"customHttpsProvisioningEnabled"`
	// The ID of the FrontDoor Frontend Endpoint which this configuration refers to.
	FrontendEndpointId pulumi.StringOutput `pulumi:"frontendEndpointId"`
	// Deprecated: This field is no longer used and will be removed in the next major version of the Azure Provider
	ResourceGroupName pulumi.StringPtrOutput `pulumi:"resourceGroupName"`
}

// NewCustomHttpsConfiguration registers a new resource with the given unique name, arguments, and options.
func NewCustomHttpsConfiguration(ctx *pulumi.Context,
	name string, args *CustomHttpsConfigurationArgs, opts ...pulumi.ResourceOption) (*CustomHttpsConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CustomHttpsProvisioningEnabled == nil {
		return nil, errors.New("invalid value for required argument 'CustomHttpsProvisioningEnabled'")
	}
	if args.FrontendEndpointId == nil {
		return nil, errors.New("invalid value for required argument 'FrontendEndpointId'")
	}
	var resource CustomHttpsConfiguration
	err := ctx.RegisterResource("azure:frontdoor/customHttpsConfiguration:CustomHttpsConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomHttpsConfiguration gets an existing CustomHttpsConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomHttpsConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomHttpsConfigurationState, opts ...pulumi.ResourceOption) (*CustomHttpsConfiguration, error) {
	var resource CustomHttpsConfiguration
	err := ctx.ReadResource("azure:frontdoor/customHttpsConfiguration:CustomHttpsConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomHttpsConfiguration resources.
type customHttpsConfigurationState struct {
	// A `customHttpsConfiguration` block as defined below.
	CustomHttpsConfiguration *CustomHttpsConfigurationCustomHttpsConfiguration `pulumi:"customHttpsConfiguration"`
	// Should the HTTPS protocol be enabled for this custom domain associated with the Front Door?
	CustomHttpsProvisioningEnabled *bool `pulumi:"customHttpsProvisioningEnabled"`
	// The ID of the FrontDoor Frontend Endpoint which this configuration refers to.
	FrontendEndpointId *string `pulumi:"frontendEndpointId"`
	// Deprecated: This field is no longer used and will be removed in the next major version of the Azure Provider
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type CustomHttpsConfigurationState struct {
	// A `customHttpsConfiguration` block as defined below.
	CustomHttpsConfiguration CustomHttpsConfigurationCustomHttpsConfigurationPtrInput
	// Should the HTTPS protocol be enabled for this custom domain associated with the Front Door?
	CustomHttpsProvisioningEnabled pulumi.BoolPtrInput
	// The ID of the FrontDoor Frontend Endpoint which this configuration refers to.
	FrontendEndpointId pulumi.StringPtrInput
	// Deprecated: This field is no longer used and will be removed in the next major version of the Azure Provider
	ResourceGroupName pulumi.StringPtrInput
}

func (CustomHttpsConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*customHttpsConfigurationState)(nil)).Elem()
}

type customHttpsConfigurationArgs struct {
	// A `customHttpsConfiguration` block as defined below.
	CustomHttpsConfiguration *CustomHttpsConfigurationCustomHttpsConfiguration `pulumi:"customHttpsConfiguration"`
	// Should the HTTPS protocol be enabled for this custom domain associated with the Front Door?
	CustomHttpsProvisioningEnabled bool `pulumi:"customHttpsProvisioningEnabled"`
	// The ID of the FrontDoor Frontend Endpoint which this configuration refers to.
	FrontendEndpointId string `pulumi:"frontendEndpointId"`
	// Deprecated: This field is no longer used and will be removed in the next major version of the Azure Provider
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a CustomHttpsConfiguration resource.
type CustomHttpsConfigurationArgs struct {
	// A `customHttpsConfiguration` block as defined below.
	CustomHttpsConfiguration CustomHttpsConfigurationCustomHttpsConfigurationPtrInput
	// Should the HTTPS protocol be enabled for this custom domain associated with the Front Door?
	CustomHttpsProvisioningEnabled pulumi.BoolInput
	// The ID of the FrontDoor Frontend Endpoint which this configuration refers to.
	FrontendEndpointId pulumi.StringInput
	// Deprecated: This field is no longer used and will be removed in the next major version of the Azure Provider
	ResourceGroupName pulumi.StringPtrInput
}

func (CustomHttpsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customHttpsConfigurationArgs)(nil)).Elem()
}

type CustomHttpsConfigurationInput interface {
	pulumi.Input

	ToCustomHttpsConfigurationOutput() CustomHttpsConfigurationOutput
	ToCustomHttpsConfigurationOutputWithContext(ctx context.Context) CustomHttpsConfigurationOutput
}

type CustomHttpsConfigurationPtrInput interface {
	pulumi.Input

	ToCustomHttpsConfigurationPtrOutput() CustomHttpsConfigurationPtrOutput
	ToCustomHttpsConfigurationPtrOutputWithContext(ctx context.Context) CustomHttpsConfigurationPtrOutput
}

func (CustomHttpsConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomHttpsConfiguration)(nil)).Elem()
}

func (i CustomHttpsConfiguration) ToCustomHttpsConfigurationOutput() CustomHttpsConfigurationOutput {
	return i.ToCustomHttpsConfigurationOutputWithContext(context.Background())
}

func (i CustomHttpsConfiguration) ToCustomHttpsConfigurationOutputWithContext(ctx context.Context) CustomHttpsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomHttpsConfigurationOutput)
}

func (i CustomHttpsConfiguration) ToCustomHttpsConfigurationPtrOutput() CustomHttpsConfigurationPtrOutput {
	return i.ToCustomHttpsConfigurationPtrOutputWithContext(context.Background())
}

func (i CustomHttpsConfiguration) ToCustomHttpsConfigurationPtrOutputWithContext(ctx context.Context) CustomHttpsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomHttpsConfigurationPtrOutput)
}

type CustomHttpsConfigurationOutput struct {
	*pulumi.OutputState
}

func (CustomHttpsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomHttpsConfigurationOutput)(nil)).Elem()
}

func (o CustomHttpsConfigurationOutput) ToCustomHttpsConfigurationOutput() CustomHttpsConfigurationOutput {
	return o
}

func (o CustomHttpsConfigurationOutput) ToCustomHttpsConfigurationOutputWithContext(ctx context.Context) CustomHttpsConfigurationOutput {
	return o
}

type CustomHttpsConfigurationPtrOutput struct {
	*pulumi.OutputState
}

func (CustomHttpsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomHttpsConfiguration)(nil)).Elem()
}

func (o CustomHttpsConfigurationPtrOutput) ToCustomHttpsConfigurationPtrOutput() CustomHttpsConfigurationPtrOutput {
	return o
}

func (o CustomHttpsConfigurationPtrOutput) ToCustomHttpsConfigurationPtrOutputWithContext(ctx context.Context) CustomHttpsConfigurationPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CustomHttpsConfigurationOutput{})
	pulumi.RegisterOutputType(CustomHttpsConfigurationPtrOutput{})
}

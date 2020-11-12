// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azure

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The provider type for the azurerm package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}
	if args.Environment == nil {
		args.Environment = pulumi.StringPtr(getEnvOrDefault("public", nil, "AZURE_ENVIRONMENT", "ARM_ENVIRONMENT").(string))
	}
	if args.MetadataHost == nil {
		args.MetadataHost = pulumi.StringPtr(getEnvOrDefault("", nil, "ARM_METADATA_HOSTNAME").(string))
	}
	if args.MetadataUrl == nil {
		args.MetadataUrl = pulumi.StringPtr(getEnvOrDefault("", nil, "ARM_METADATA_URL").(string))
	}
	if args.SkipProviderRegistration == nil {
		args.SkipProviderRegistration = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "ARM_SKIP_PROVIDER_REGISTRATION").(bool))
	}
	if args.StorageUseAzuread == nil {
		args.StorageUseAzuread = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "ARM_STORAGE_USE_AZUREAD").(bool))
	}
	if args.SubscriptionId == nil {
		args.SubscriptionId = pulumi.StringPtr(getEnvOrDefault("", nil, "ARM_SUBSCRIPTION_ID").(string))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:azure", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	AuxiliaryTenantIds []string `pulumi:"auxiliaryTenantIds"`
	// The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
	// Certificate
	ClientCertificatePassword *string `pulumi:"clientCertificatePassword"`
	// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
	// Principal using a Client Certificate.
	ClientCertificatePath *string `pulumi:"clientCertificatePath"`
	// The Client ID which should be used.
	ClientId *string `pulumi:"clientId"`
	// The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
	ClientSecret *string `pulumi:"clientSecret"`
	// This will disable the x-ms-correlation-request-id header.
	DisableCorrelationRequestId *bool `pulumi:"disableCorrelationRequestId"`
	// This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
	DisableTerraformPartnerId *bool `pulumi:"disableTerraformPartnerId"`
	// The Cloud Environment which should be used. Possible values are public, usgovernment, german, and china. Defaults to
	// public.
	Environment *string           `pulumi:"environment"`
	Features    *ProviderFeatures `pulumi:"features"`
	// The Hostname which should be used for the Azure Metadata Service.
	MetadataHost *string `pulumi:"metadataHost"`
	// Deprecated - replaced by `metadata_host`.
	//
	// Deprecated: use `metadata_host` instead
	MetadataUrl *string `pulumi:"metadataUrl"`
	// The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected
	// automatically.
	MsiEndpoint *string `pulumi:"msiEndpoint"`
	// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
	PartnerId *string `pulumi:"partnerId"`
	// This will cause the AzureRM Provider to skip verifying the credentials being used are valid.
	SkipCredentialsValidation *bool `pulumi:"skipCredentialsValidation"`
	// Should the AzureRM Provider skip registering all of the Resource Providers that it supports, if they're not already
	// registered?
	SkipProviderRegistration *bool `pulumi:"skipProviderRegistration"`
	// Should the AzureRM Provider use AzureAD to access the Storage Data Plane API's?
	StorageUseAzuread *bool `pulumi:"storageUseAzuread"`
	// The Subscription ID which should be used.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The Tenant ID which should be used.
	TenantId *string `pulumi:"tenantId"`
	// Allowed Managed Service Identity be used for Authentication.
	UseMsi *bool `pulumi:"useMsi"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	AuxiliaryTenantIds pulumi.StringArrayInput
	// The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
	// Certificate
	ClientCertificatePassword pulumi.StringPtrInput
	// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
	// Principal using a Client Certificate.
	ClientCertificatePath pulumi.StringPtrInput
	// The Client ID which should be used.
	ClientId pulumi.StringPtrInput
	// The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
	ClientSecret pulumi.StringPtrInput
	// This will disable the x-ms-correlation-request-id header.
	DisableCorrelationRequestId pulumi.BoolPtrInput
	// This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
	DisableTerraformPartnerId pulumi.BoolPtrInput
	// The Cloud Environment which should be used. Possible values are public, usgovernment, german, and china. Defaults to
	// public.
	Environment pulumi.StringPtrInput
	Features    ProviderFeaturesPtrInput
	// The Hostname which should be used for the Azure Metadata Service.
	MetadataHost pulumi.StringPtrInput
	// Deprecated - replaced by `metadata_host`.
	//
	// Deprecated: use `metadata_host` instead
	MetadataUrl pulumi.StringPtrInput
	// The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected
	// automatically.
	MsiEndpoint pulumi.StringPtrInput
	// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
	PartnerId pulumi.StringPtrInput
	// This will cause the AzureRM Provider to skip verifying the credentials being used are valid.
	SkipCredentialsValidation pulumi.BoolPtrInput
	// Should the AzureRM Provider skip registering all of the Resource Providers that it supports, if they're not already
	// registered?
	SkipProviderRegistration pulumi.BoolPtrInput
	// Should the AzureRM Provider use AzureAD to access the Storage Data Plane API's?
	StorageUseAzuread pulumi.BoolPtrInput
	// The Subscription ID which should be used.
	SubscriptionId pulumi.StringPtrInput
	// The Tenant ID which should be used.
	TenantId pulumi.StringPtrInput
	// Allowed Managed Service Identity be used for Authentication.
	UseMsi pulumi.BoolPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (Provider) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil)).Elem()
}

func (i Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct {
	*pulumi.OutputState
}

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderOutput)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProviderOutput{})
}

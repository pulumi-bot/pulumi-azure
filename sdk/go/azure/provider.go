// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azure

import (
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
	if args.ClientCertificatePassword == nil {
		args.ClientCertificatePassword = pulumi.ToSecret(getEnvOrDefault("", nil, "AZURE_CLIENT_CERTIFICATE_PASSWORD", "ARM_CLIENT_CERTIFICATE_PASSWORD").(string)).(pulumi.StringOutput)
	}
	if args.ClientCertificatePath == nil {
		args.ClientCertificatePath = pulumi.StringPtr(getEnvOrDefault("", nil, "AZURE_CLIENT_CERTIFICATE_PATH", "ARM_CLIENT_CERTIFICATE_PATH").(string))
	}
	if args.ClientId == nil {
		args.ClientId = pulumi.StringPtr(getEnvOrDefault("", nil, "AZURE_CLIENT_ID", "ARM_CLIENT_ID").(string))
	}
	if args.ClientSecret == nil {
		args.ClientSecret = pulumi.ToSecret(getEnvOrDefault("", nil, "AZURE_CLIENT_SECRET", "ARM_CLIENT_SECRET").(string)).(pulumi.StringOutput)
	}
	if args.DisableTerraformPartnerId == nil {
		args.DisableTerraformPartnerId = pulumi.BoolPtr(getEnvOrDefault(true, parseEnvBool, "ARM_DISABLE_TERRAFORM_PARTNER_ID").(bool))
	}
	if args.Environment == nil {
		args.Environment = pulumi.StringPtr(getEnvOrDefault("public", nil, "AZURE_ENVIRONMENT", "ARM_ENVIRONMENT").(string))
	}
	if args.MsiEndpoint == nil {
		args.MsiEndpoint = pulumi.StringPtr(getEnvOrDefault("", nil, "ARM_MSI_ENDPOINT").(string))
	}
	if args.PartnerId == nil {
		args.PartnerId = pulumi.StringPtr(getEnvOrDefault("", nil, "ARM_PARTNER_ID").(string))
	}
	if args.SkipCredentialsValidation == nil {
		args.SkipCredentialsValidation = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "ARM_SKIP_CREDENTIALS_VALIDATION").(bool))
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
	if args.TenantId == nil {
		args.TenantId = pulumi.StringPtr(getEnvOrDefault("", nil, "AZURE_TENANT_ID", "ARM_TENANT_ID").(string))
	}
	if args.UseMsi == nil {
		args.UseMsi = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "ARM_USE_MSI").(bool))
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

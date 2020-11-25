// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.Azure
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("azure");
        public static ImmutableArray<string> AuxiliaryTenantIds { get; set; } = __config.GetObject<ImmutableArray<string>>("auxiliaryTenantIds");

        /// <summary>
        /// The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
        /// Certificate
        /// </summary>
        public static string? ClientCertificatePassword { get; set; } = __config.Get("clientCertificatePassword");

        /// <summary>
        /// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
        /// Principal using a Client Certificate.
        /// </summary>
        public static string? ClientCertificatePath { get; set; } = __config.Get("clientCertificatePath");

        /// <summary>
        /// The Client ID which should be used.
        /// </summary>
        public static string? ClientId { get; set; } = __config.Get("clientId");

        /// <summary>
        /// The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
        /// </summary>
        public static string? ClientSecret { get; set; } = __config.Get("clientSecret");

        /// <summary>
        /// This will disable the x-ms-correlation-request-id header.
        /// </summary>
        public static bool? DisableCorrelationRequestId { get; set; } = __config.GetBoolean("disableCorrelationRequestId");

        /// <summary>
        /// This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
        /// </summary>
        public static bool? DisableTerraformPartnerId { get; set; } = __config.GetBoolean("disableTerraformPartnerId");

        /// <summary>
        /// The Cloud Environment which should be used. Possible values are public, usgovernment, german, and china. Defaults to
        /// public.
        /// </summary>
        public static string? Environment { get; set; } = __config.Get("environment") ?? Utilities.GetEnv("AZURE_ENVIRONMENT", "ARM_ENVIRONMENT") ?? "public";

        public static Pulumi.Azure.Config.Types.Features? Features { get; set; } = __config.GetObject<Pulumi.Azure.Config.Types.Features>("features");

        public static string? Location { get; set; } = __config.Get("location") ?? Utilities.GetEnv("ARM_LOCATION");

        /// <summary>
        /// The Hostname which should be used for the Azure Metadata Service.
        /// </summary>
        public static string? MetadataHost { get; set; } = __config.Get("metadataHost") ?? Utilities.GetEnv("ARM_METADATA_HOSTNAME") ?? "";

        /// <summary>
        /// Deprecated - replaced by `metadata_host`.
        /// </summary>
        public static string? MetadataUrl { get; set; } = __config.Get("metadataUrl") ?? Utilities.GetEnv("ARM_METADATA_URL") ?? "";

        /// <summary>
        /// The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected
        /// automatically.
        /// </summary>
        public static string? MsiEndpoint { get; set; } = __config.Get("msiEndpoint");

        /// <summary>
        /// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
        /// </summary>
        public static string? PartnerId { get; set; } = __config.Get("partnerId");

        /// <summary>
        /// This will cause the AzureRM Provider to skip verifying the credentials being used are valid.
        /// </summary>
        public static bool? SkipCredentialsValidation { get; set; } = __config.GetBoolean("skipCredentialsValidation");

        /// <summary>
        /// Should the AzureRM Provider skip registering all of the Resource Providers that it supports, if they're not already
        /// registered?
        /// </summary>
        public static bool? SkipProviderRegistration { get; set; } = __config.GetBoolean("skipProviderRegistration") ?? Utilities.GetEnvBoolean("ARM_SKIP_PROVIDER_REGISTRATION") ?? false;

        /// <summary>
        /// Should the AzureRM Provider use AzureAD to access the Storage Data Plane API's?
        /// </summary>
        public static bool? StorageUseAzuread { get; set; } = __config.GetBoolean("storageUseAzuread") ?? Utilities.GetEnvBoolean("ARM_STORAGE_USE_AZUREAD") ?? false;

        /// <summary>
        /// The Subscription ID which should be used.
        /// </summary>
        public static string? SubscriptionId { get; set; } = __config.Get("subscriptionId") ?? Utilities.GetEnv("ARM_SUBSCRIPTION_ID") ?? "";

        /// <summary>
        /// The Tenant ID which should be used.
        /// </summary>
        public static string? TenantId { get; set; } = __config.Get("tenantId");

        /// <summary>
        /// Allowed Managed Service Identity be used for Authentication.
        /// </summary>
        public static bool? UseMsi { get; set; } = __config.GetBoolean("useMsi");

        public static class Types
        {

             public class Features
             {
                public Pulumi.Azure.Config.Types.FeaturesKeyVault? KeyVault { get; set; } = null!;
                public Pulumi.Azure.Config.Types.FeaturesNetwork? Network { get; set; } = null!;
                public Pulumi.Azure.Config.Types.FeaturesTemplateDeployment? TemplateDeployment { get; set; } = null!;
                public Pulumi.Azure.Config.Types.FeaturesVirtualMachine? VirtualMachine { get; set; } = null!;
                public Pulumi.Azure.Config.Types.FeaturesVirtualMachineScaleSet? VirtualMachineScaleSet { get; set; } = null!;
            }

             public class FeaturesKeyVault
             {
                public bool? PurgeSoftDeleteOnDestroy { get; set; }
                public bool? RecoverSoftDeletedKeyVaults { get; set; }
            }

             public class FeaturesNetwork
             {
                public bool RelaxedLocking { get; set; }
            }

             public class FeaturesTemplateDeployment
             {
                public bool DeleteNestedItemsDuringDeletion { get; set; }
            }

             public class FeaturesVirtualMachine
             {
                public bool? DeleteOsDiskOnDeletion { get; set; }
                public bool? GracefulShutdown { get; set; }
            }

             public class FeaturesVirtualMachineScaleSet
             {
                public bool RollInstancesWhenRequired { get; set; }
            }
        }
    }
}

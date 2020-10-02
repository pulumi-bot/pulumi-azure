// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

let __config = new pulumi.Config("azure");

export let auxiliaryTenantIds: string[] | undefined = __config.getObject<string[]>("auxiliaryTenantIds");
/**
 * The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
 * Certificate
 */
export let clientCertificatePassword: string | undefined = __config.get("clientCertificatePassword");
/**
 * The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
 * Principal using a Client Certificate.
 */
export let clientCertificatePath: string | undefined = __config.get("clientCertificatePath");
/**
 * The Client ID which should be used.
 */
export let clientId: string | undefined = __config.get("clientId");
/**
 * The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
 */
export let clientSecret: string | undefined = __config.get("clientSecret");
/**
 * This will disable the x-ms-correlation-request-id header.
 */
export let disableCorrelationRequestId: boolean | undefined = __config.getObject<boolean>("disableCorrelationRequestId");
/**
 * This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
 */
export let disableTerraformPartnerId: boolean | undefined = __config.getObject<boolean>("disableTerraformPartnerId");
/**
 * The Cloud Environment which should be used. Possible values are public, usgovernment, german, and china. Defaults to
 * public.
 */
export let environment: string | undefined = __config.get("environment") || (utilities.getEnv("AZURE_ENVIRONMENT", "ARM_ENVIRONMENT") || "public");
export let features: outputs.config.Features | undefined = __config.getObject<outputs.config.Features>("features");
export let location: string | undefined = __config.get("location") || utilities.getEnv("ARM_LOCATION");
/**
 * The Hostname which should be used for the Azure Metadata Service.
 */
export let metadataHost: string | undefined = __config.get("metadataHost") || (utilities.getEnv("ARM_METADATA_HOSTNAME") || "");
/**
 * Deprecated - replaced by `metadata_host`.
 */
export let metadataUrl: string | undefined = __config.get("metadataUrl") || (utilities.getEnv("ARM_METADATA_URL") || "");
/**
 * The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected
 * automatically.
 */
export let msiEndpoint: string | undefined = __config.get("msiEndpoint");
/**
 * A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
 */
export let partnerId: string | undefined = __config.get("partnerId");
/**
 * This will cause the AzureRM Provider to skip verifying the credentials being used are valid.
 */
export let skipCredentialsValidation: boolean | undefined = __config.getObject<boolean>("skipCredentialsValidation");
/**
 * Should the AzureRM Provider skip registering all of the Resource Providers that it supports, if they're not already
 * registered?
 */
export let skipProviderRegistration: boolean | undefined = __config.getObject<boolean>("skipProviderRegistration") || (<any>utilities.getEnvBoolean("ARM_SKIP_PROVIDER_REGISTRATION") || false);
/**
 * Should the AzureRM Provider use AzureAD to access the Storage Data Plane API's?
 */
export let storageUseAzuread: boolean | undefined = __config.getObject<boolean>("storageUseAzuread") || (<any>utilities.getEnvBoolean("ARM_STORAGE_USE_AZUREAD") || false);
/**
 * The Subscription ID which should be used.
 */
export let subscriptionId: string | undefined = __config.get("subscriptionId") || (utilities.getEnv("ARM_SUBSCRIPTION_ID") || "");
/**
 * The Tenant ID which should be used.
 */
export let tenantId: string | undefined = __config.get("tenantId");
/**
 * Allowed Managed Service Identity be used for Authentication.
 */
export let useMsi: boolean | undefined = __config.getObject<boolean>("useMsi");

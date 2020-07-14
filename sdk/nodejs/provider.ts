// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The provider type for the azurerm package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'azure';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }


    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        {
            inputs["auxiliaryTenantIds"] = pulumi.output(args ? args.auxiliaryTenantIds : undefined).apply(JSON.stringify);
            inputs["clientCertificatePassword"] = args ? args.clientCertificatePassword : undefined;
            inputs["clientCertificatePath"] = args ? args.clientCertificatePath : undefined;
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientSecret"] = args ? args.clientSecret : undefined;
            inputs["disableCorrelationRequestId"] = pulumi.output(args ? args.disableCorrelationRequestId : undefined).apply(JSON.stringify);
            inputs["disableTerraformPartnerId"] = pulumi.output(args ? args.disableTerraformPartnerId : undefined).apply(JSON.stringify);
            inputs["environment"] = (args ? args.environment : undefined) || (utilities.getEnv("AZURE_ENVIRONMENT", "ARM_ENVIRONMENT") || "public");
            inputs["features"] = pulumi.output(args ? args.features : undefined).apply(JSON.stringify);
            inputs["metadataUrl"] = (args ? args.metadataUrl : undefined) || (utilities.getEnv("ARM_METADATA_URL") || "");
            inputs["msiEndpoint"] = args ? args.msiEndpoint : undefined;
            inputs["partnerId"] = args ? args.partnerId : undefined;
            inputs["skipCredentialsValidation"] = pulumi.output(args ? args.skipCredentialsValidation : undefined).apply(JSON.stringify);
            inputs["skipProviderRegistration"] = pulumi.output((args ? args.skipProviderRegistration : undefined) || (<any>utilities.getEnvBoolean("ARM_SKIP_PROVIDER_REGISTRATION") || false)).apply(JSON.stringify);
            inputs["storageUseAzuread"] = pulumi.output((args ? args.storageUseAzuread : undefined) || (<any>utilities.getEnvBoolean("ARM_STORAGE_USE_AZUREAD") || false)).apply(JSON.stringify);
            inputs["subscriptionId"] = (args ? args.subscriptionId : undefined) || (utilities.getEnv("ARM_SUBSCRIPTION_ID") || "");
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["useMsi"] = pulumi.output(args ? args.useMsi : undefined).apply(JSON.stringify);
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Provider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    readonly auxiliaryTenantIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
     * Certificate
     */
    readonly clientCertificatePassword?: pulumi.Input<string>;
    /**
     * The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
     * Principal using a Client Certificate.
     */
    readonly clientCertificatePath?: pulumi.Input<string>;
    /**
     * The Client ID which should be used.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
     */
    readonly clientSecret?: pulumi.Input<string>;
    /**
     * This will disable the x-ms-correlation-request-id header.
     */
    readonly disableCorrelationRequestId?: pulumi.Input<boolean>;
    /**
     * This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
     */
    readonly disableTerraformPartnerId?: pulumi.Input<boolean>;
    /**
     * The Cloud Environment which should be used. Possible values are public, usgovernment, german, and china. Defaults to
     * public.
     */
    readonly environment?: pulumi.Input<string>;
    readonly features?: pulumi.Input<inputs.ProviderFeatures>;
    /**
     * The Metadata URL which will be used to obtain the Cloud Environment.
     */
    readonly metadataUrl?: pulumi.Input<string>;
    /**
     * The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected
     * automatically.
     */
    readonly msiEndpoint?: pulumi.Input<string>;
    /**
     * A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
     */
    readonly partnerId?: pulumi.Input<string>;
    /**
     * This will cause the AzureRM Provider to skip verifying the credentials being used are valid.
     */
    readonly skipCredentialsValidation?: pulumi.Input<boolean>;
    /**
     * Should the AzureRM Provider skip registering all of the Resource Providers that it supports, if they're not already
     * registered?
     */
    readonly skipProviderRegistration?: pulumi.Input<boolean>;
    /**
     * Should the AzureRM Provider use AzureAD to access the Storage Data Plane API's?
     */
    readonly storageUseAzuread?: pulumi.Input<boolean>;
    /**
     * The Subscription ID which should be used.
     */
    readonly subscriptionId?: pulumi.Input<string>;
    /**
     * The Tenant ID which should be used.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Allowed Managed Service Identity be used for Authentication.
     */
    readonly useMsi?: pulumi.Input<boolean>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Key Vault Certificate Issuer.
 */
export function getCertificateIssuer(args: GetCertificateIssuerArgs, opts?: pulumi.InvokeOptions): Promise<GetCertificateIssuerResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:keyvault/getCertificateIssuer:getCertificateIssuer", {
        "keyVaultId": args.keyVaultId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getCertificateIssuer.
 */
export interface GetCertificateIssuerArgs {
    /**
     * The ID of the Key Vault in which to locate the Certificate Issuer.
     */
    readonly keyVaultId: string;
    /**
     * The name of the Key Vault Certificate Issuer.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getCertificateIssuer.
 */
export interface GetCertificateIssuerResult {
    /**
     * The account number with the third-party Certificate Issuer.
     */
    readonly accountId: string;
    /**
     * A list of `admin` blocks as defined below.
     */
    readonly admins: outputs.keyvault.GetCertificateIssuerAdmin[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly keyVaultId: string;
    readonly name: string;
    /**
     * The organization ID with the third-party Certificate Issuer.
     */
    readonly orgId: string;
    /**
     * The name of the third-party Certificate Issuer.
     */
    readonly providerName: string;
}

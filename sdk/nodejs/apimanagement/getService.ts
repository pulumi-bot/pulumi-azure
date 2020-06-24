// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing API Management Service.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.apimanagement.getService({
 *     name: "search-api",
 *     resourceGroupName: "search-service",
 * });
 * export const apiManagementId = example.then(example => example.id);
 * ```
 */
export function getService(args: GetServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:apimanagement/getService:getService", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getService.
 */
export interface GetServiceArgs {
    /**
     * The name of the API Management service.
     */
    readonly name: string;
    /**
     * The Name of the Resource Group in which the API Management Service exists.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getService.
 */
export interface GetServiceResult {
    /**
     * One or more `additionalLocation` blocks as defined below
     */
    readonly additionalLocations: outputs.apimanagement.GetServiceAdditionalLocation[];
    /**
     * The URL for the Developer Portal associated with this API Management service.
     */
    readonly developerPortalUrl: string;
    /**
     * Gateway URL of the API Management service in the Region.
     */
    readonly gatewayRegionalUrl: string;
    /**
     * The URL for the API Management Service's Gateway.
     */
    readonly gatewayUrl: string;
    /**
     * A `hostnameConfiguration` block as defined below.
     */
    readonly hostnameConfigurations: outputs.apimanagement.GetServiceHostnameConfiguration[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * (Optional) An `identity` block as defined below.
     */
    readonly identities: outputs.apimanagement.GetServiceIdentity[];
    /**
     * The location name of the additional region among Azure Data center regions.
     */
    readonly location: string;
    /**
     * The URL for the Management API.
     */
    readonly managementApiUrl: string;
    /**
     * Specifies the plan's pricing tier.
     */
    readonly name: string;
    /**
     * The email address from which the notification will be sent.
     */
    readonly notificationSenderEmail: string;
    /**
     * The URL of the Publisher Portal.
     */
    readonly portalUrl: string;
    /**
     * Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU.
     */
    readonly publicIpAddresses: string[];
    /**
     * The email of Publisher/Company of the API Management Service.
     */
    readonly publisherEmail: string;
    /**
     * The name of the Publisher/Company of the API Management Service.
     */
    readonly publisherName: string;
    readonly resourceGroupName: string;
    /**
     * The SCM (Source Code Management) endpoint.
     */
    readonly scmUrl: string;
    readonly skuName: string;
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags: {[key: string]: string};
}

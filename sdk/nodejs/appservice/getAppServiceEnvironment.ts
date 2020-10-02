// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing App Service Environment.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.appservice.getAppServiceEnvironment({
 *     name: "existing-ase",
 *     resourceGroupName: "existing-rg",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getAppServiceEnvironment(args: GetAppServiceEnvironmentArgs, opts?: pulumi.InvokeOptions): Promise<GetAppServiceEnvironmentResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:appservice/getAppServiceEnvironment:getAppServiceEnvironment", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppServiceEnvironment.
 */
export interface GetAppServiceEnvironmentArgs {
    /**
     * The name of this App Service Environment.
     */
    readonly name: string;
    /**
     * The name of the Resource Group where the App Service Environment exists.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getAppServiceEnvironment.
 */
export interface GetAppServiceEnvironmentResult {
    /**
     * The number of app instances per App Service Environment Front End.
     */
    readonly frontEndScaleFactor: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * IP address of internal load balancer of the App Service Environment.
     */
    readonly internalIpAddress: string;
    /**
     * The Azure Region where the App Service Environment exists.
     */
    readonly location: string;
    readonly name: string;
    /**
     * List of outbound IP addresses of the App Service Environment.
     */
    readonly outboundIpAddresses: string[];
    /**
     * The Pricing Tier (Isolated SKU) of the App Service Environment.
     */
    readonly pricingTier: string;
    readonly resourceGroupName: string;
    /**
     * IP address of service endpoint of the App Service Environment.
     */
    readonly serviceIpAddress: string;
    /**
     * A mapping of tags assigned to the App Service Environment.
     */
    readonly tags: {[key: string]: string};
}

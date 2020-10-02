// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Spring Cloud Service.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.appplatform.getSpringCloudService({
 *     name: azurerm_spring_cloud_service.example.name,
 *     resourceGroupName: azurerm_spring_cloud_service.example.resource_group_name,
 * });
 * export const springCloudServiceId = example.then(example => example.id);
 * ```
 */
export function getSpringCloudService(args: GetSpringCloudServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetSpringCloudServiceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:appplatform/getSpringCloudService:getSpringCloudService", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getSpringCloudService.
 */
export interface GetSpringCloudServiceArgs {
    /**
     * Specifies The name of the Spring Cloud Service resource.
     */
    readonly name: string;
    /**
     * Specifies the name of the Resource Group where the Spring Cloud Service exists.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getSpringCloudService.
 */
export interface GetSpringCloudServiceResult {
    /**
     * A `configServerGitSetting` block as defined below.
     */
    readonly configServerGitSettings: outputs.appplatform.GetSpringCloudServiceConfigServerGitSetting[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The location of Spring Cloud Service.
     */
    readonly location: string;
    /**
     * The name to identify on the Git repository.
     */
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * A mapping of tags assigned to Spring Cloud Service.
     */
    readonly tags: {[key: string]: string};
}

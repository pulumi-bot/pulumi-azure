// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access the properties of a Log Profile.
 */
export function getLogProfile(args: GetLogProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetLogProfileResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:monitoring/getLogProfile:getLogProfile", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getLogProfile.
 */
export interface GetLogProfileArgs {
    /**
     * Specifies the Name of the Log Profile.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getLogProfile.
 */
export interface GetLogProfileResult {
    /**
     * List of categories of the logs.
     */
    readonly categories: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of regions for which Activity Log events are stored or streamed.
     */
    readonly locations: string[];
    readonly name: string;
    readonly retentionPolicies: outputs.monitoring.GetLogProfileRetentionPolicy[];
    /**
     * The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to.
     */
    readonly servicebusRuleId: string;
    /**
     * The resource id of the storage account in which the Activity Log is stored.
     */
    readonly storageAccountId: string;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Advisor Recommendations.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.advisor.getRecommendations({
 *     filterByCategories: [
 *         "security",
 *         "cost",
 *     ],
 *     filterByResourceGroups: ["example-resgroups"],
 * });
 * export const recommendations = example.then(example => example.recommendations);
 * ```
 */
export function getRecommendations(args?: GetRecommendationsArgs, opts?: pulumi.InvokeOptions): Promise<GetRecommendationsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:advisor/getRecommendations:getRecommendations", {
        "filterByCategories": args.filterByCategories,
        "filterByResourceGroups": args.filterByResourceGroups,
    }, opts);
}

/**
 * A collection of arguments for invoking getRecommendations.
 */
export interface GetRecommendationsArgs {
    /**
     * Specifies a list of categories in which the Advisor Recommendations will be listed. Possible values are `HighAvailability`, `Security`, `Performance`, `Cost` and `OperationalExcellence`.
     */
    filterByCategories?: string[];
    /**
     * Specifies a list of resource groups about which the Advisor Recommendations will be listed.
     */
    filterByResourceGroups?: string[];
}

/**
 * A collection of values returned by getRecommendations.
 */
export interface GetRecommendationsResult {
    readonly filterByCategories?: string[];
    readonly filterByResourceGroups?: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * One or more `recommendations` blocks as defined below.
     */
    readonly recommendations: outputs.advisor.GetRecommendationsRecommendation[];
}

export function getRecommendationsApply(args?: GetRecommendationsApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRecommendationsResult> {
    return pulumi.output(args).apply(a => getRecommendations(a, opts))
}

/**
 * A collection of arguments for invoking getRecommendations.
 */
export interface GetRecommendationsApplyArgs {
    /**
     * Specifies a list of categories in which the Advisor Recommendations will be listed. Possible values are `HighAvailability`, `Security`, `Performance`, `Cost` and `OperationalExcellence`.
     */
    filterByCategories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies a list of resource groups about which the Advisor Recommendations will be listed.
     */
    filterByResourceGroups?: pulumi.Input<pulumi.Input<string>[]>;
}

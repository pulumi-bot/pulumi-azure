// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Sentinel Alert Rule Template.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.sentinel.getAlertRuleTemplate({
 *     logAnalyticsWorkspaceId: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1",
 *     displayName: "Create incidents based on Azure Security Center for IoT alerts",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getAlertRuleTemplate(args: GetAlertRuleTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetAlertRuleTemplateResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:sentinel/getAlertRuleTemplate:getAlertRuleTemplate", {
        "displayName": args.displayName,
        "logAnalyticsWorkspaceId": args.logAnalyticsWorkspaceId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getAlertRuleTemplate.
 */
export interface GetAlertRuleTemplateArgs {
    /**
     * The display name of this Sentinel Alert Rule Template. Either `displayName` or `name` have to be specified.
     */
    displayName?: string;
    /**
     * The ID of the Log Analytics Workspace.
     */
    logAnalyticsWorkspaceId: string;
    /**
     * The name of this Sentinel Alert Rule Template. Either `displayName` or `name` have to be specified.
     */
    name?: string;
}

/**
 * A collection of values returned by getAlertRuleTemplate.
 */
export interface GetAlertRuleTemplateResult {
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly logAnalyticsWorkspaceId: string;
    readonly name: string;
    /**
     * A `scheduledTemplate` block as defined below. This only applies to Sentinel Scheduled Alert Rule Template.
     */
    readonly scheduledTemplates: outputs.sentinel.GetAlertRuleTemplateScheduledTemplate[];
    /**
     * A `securityIncidentTemplate` block as defined below. This only applies to Sentinel MS Security Incident Alert Rule Template.
     */
    readonly securityIncidentTemplates: outputs.sentinel.GetAlertRuleTemplateSecurityIncidentTemplate[];
}

export function getAlertRuleTemplateOutput(args: GetAlertRuleTemplateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAlertRuleTemplateResult> {
    return pulumi.output(args).apply(a => getAlertRuleTemplate(a, opts))
}

/**
 * A collection of arguments for invoking getAlertRuleTemplate.
 */
export interface GetAlertRuleTemplateOutputArgs {
    /**
     * The display name of this Sentinel Alert Rule Template. Either `displayName` or `name` have to be specified.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The ID of the Log Analytics Workspace.
     */
    logAnalyticsWorkspaceId: pulumi.Input<string>;
    /**
     * The name of this Sentinel Alert Rule Template. Either `displayName` or `name` have to be specified.
     */
    name?: pulumi.Input<string>;
}

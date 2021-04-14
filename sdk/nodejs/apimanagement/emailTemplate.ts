// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a API Management Email Template.
 *
 * ## Import
 *
 * API Management Email Templates can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:apimanagement/emailTemplate:EmailTemplate example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/instance1/templates/template1
 * ```
 */
export class EmailTemplate extends pulumi.CustomResource {
    /**
     * Get an existing EmailTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EmailTemplateState, opts?: pulumi.CustomResourceOptions): EmailTemplate {
        return new EmailTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:apimanagement/emailTemplate:EmailTemplate';

    /**
     * Returns true if the given object is an instance of EmailTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EmailTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EmailTemplate.__pulumiType;
    }

    /**
     * The name of the API Management Service in which the Email Template should exist. Changing this forces a new API Management Email Template to be created.
     */
    public readonly apiManagementName!: pulumi.Output<string>;
    /**
     * The body of the Email. Its format has to be a well-formed HTML document.
     */
    public readonly body!: pulumi.Output<string>;
    /**
     * The description of the Email Template.
     */
    public /*out*/ readonly description!: pulumi.Output<string>;
    /**
     * The name of the Resource Group where the API Management Email Template should exist. Changing this forces a new API Management Email Template to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The subject of the Email.
     */
    public readonly subject!: pulumi.Output<string>;
    /**
     * The name of the Email Template. Possible values are `AccountClosedDeveloper`, `ApplicationApprovedNotificationMessage`, `ConfirmSignUpIdentityDefault`, `EmailChangeIdentityDefault`, `InviteUserNotificationMessage`, `NewCommentNotificationMessage`, `NewDeveloperNotificationMessage`, `NewIssueNotificationMessage`, `PasswordResetByAdminNotificationMessage`, `PasswordResetIdentityDefault`, `PurchaseDeveloperNotificationMessage`, `QuotaLimitApproachingDeveloperNotificationMessage`, `RejectDeveloperNotificationMessage`, `RequestDeveloperNotificationMessage`. Changing this forces a new API Management Email Template to be created.
     */
    public readonly templateName!: pulumi.Output<string>;
    /**
     * The title of the Email Template.
     */
    public /*out*/ readonly title!: pulumi.Output<string>;

    /**
     * Create a EmailTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EmailTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EmailTemplateArgs | EmailTemplateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EmailTemplateState | undefined;
            inputs["apiManagementName"] = state ? state.apiManagementName : undefined;
            inputs["body"] = state ? state.body : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["subject"] = state ? state.subject : undefined;
            inputs["templateName"] = state ? state.templateName : undefined;
            inputs["title"] = state ? state.title : undefined;
        } else {
            const args = argsOrState as EmailTemplateArgs | undefined;
            if ((!args || args.apiManagementName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiManagementName'");
            }
            if ((!args || args.body === undefined) && !opts.urn) {
                throw new Error("Missing required property 'body'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.subject === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subject'");
            }
            if ((!args || args.templateName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateName'");
            }
            inputs["apiManagementName"] = args ? args.apiManagementName : undefined;
            inputs["body"] = args ? args.body : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["subject"] = args ? args.subject : undefined;
            inputs["templateName"] = args ? args.templateName : undefined;
            inputs["description"] = undefined /*out*/;
            inputs["title"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EmailTemplate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EmailTemplate resources.
 */
export interface EmailTemplateState {
    /**
     * The name of the API Management Service in which the Email Template should exist. Changing this forces a new API Management Email Template to be created.
     */
    readonly apiManagementName?: pulumi.Input<string>;
    /**
     * The body of the Email. Its format has to be a well-formed HTML document.
     */
    readonly body?: pulumi.Input<string>;
    /**
     * The description of the Email Template.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the API Management Email Template should exist. Changing this forces a new API Management Email Template to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The subject of the Email.
     */
    readonly subject?: pulumi.Input<string>;
    /**
     * The name of the Email Template. Possible values are `AccountClosedDeveloper`, `ApplicationApprovedNotificationMessage`, `ConfirmSignUpIdentityDefault`, `EmailChangeIdentityDefault`, `InviteUserNotificationMessage`, `NewCommentNotificationMessage`, `NewDeveloperNotificationMessage`, `NewIssueNotificationMessage`, `PasswordResetByAdminNotificationMessage`, `PasswordResetIdentityDefault`, `PurchaseDeveloperNotificationMessage`, `QuotaLimitApproachingDeveloperNotificationMessage`, `RejectDeveloperNotificationMessage`, `RequestDeveloperNotificationMessage`. Changing this forces a new API Management Email Template to be created.
     */
    readonly templateName?: pulumi.Input<string>;
    /**
     * The title of the Email Template.
     */
    readonly title?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EmailTemplate resource.
 */
export interface EmailTemplateArgs {
    /**
     * The name of the API Management Service in which the Email Template should exist. Changing this forces a new API Management Email Template to be created.
     */
    readonly apiManagementName: pulumi.Input<string>;
    /**
     * The body of the Email. Its format has to be a well-formed HTML document.
     */
    readonly body: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the API Management Email Template should exist. Changing this forces a new API Management Email Template to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The subject of the Email.
     */
    readonly subject: pulumi.Input<string>;
    /**
     * The name of the Email Template. Possible values are `AccountClosedDeveloper`, `ApplicationApprovedNotificationMessage`, `ConfirmSignUpIdentityDefault`, `EmailChangeIdentityDefault`, `InviteUserNotificationMessage`, `NewCommentNotificationMessage`, `NewDeveloperNotificationMessage`, `NewIssueNotificationMessage`, `PasswordResetByAdminNotificationMessage`, `PasswordResetIdentityDefault`, `PurchaseDeveloperNotificationMessage`, `QuotaLimitApproachingDeveloperNotificationMessage`, `RejectDeveloperNotificationMessage`, `RequestDeveloperNotificationMessage`. Changing this forces a new API Management Email Template to be created.
     */
    readonly templateName: pulumi.Input<string>;
}
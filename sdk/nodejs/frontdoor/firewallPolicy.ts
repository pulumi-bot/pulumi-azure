// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an Azure Front Door Web Application Firewall Policy instance.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US 2"});
 * const exampleFirewallPolicy = new azure.frontdoor.FirewallPolicy("exampleFirewallPolicy", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     enabled: true,
 *     mode: "Prevention",
 *     redirectUrl: "https://www.contoso.com",
 *     customBlockResponseStatusCode: 403,
 *     customBlockResponseBody: "PGh0bWw+CjxoZWFkZXI+PHRpdGxlPkhlbGxvPC90aXRsZT48L2hlYWRlcj4KPGJvZHk+CkhlbGxvIHdvcmxkCjwvYm9keT4KPC9odG1sPg==",
 *     custom_rule: [
 *         {
 *             name: "Rule1",
 *             enabled: true,
 *             priority: 1,
 *             rateLimitDurationInMinutes: 1,
 *             rateLimitThreshold: 10,
 *             type: "MatchRule",
 *             action: "Block",
 *             match_condition: [{
 *                 matchVariable: "RemoteAddr",
 *                 operator: "IPMatch",
 *                 negationCondition: false,
 *                 matchValues: [
 *                     "192.168.1.0/24",
 *                     "10.0.0.0/24",
 *                 ],
 *             }],
 *         },
 *         {
 *             name: "Rule2",
 *             enabled: true,
 *             priority: 2,
 *             rateLimitDurationInMinutes: 1,
 *             rateLimitThreshold: 10,
 *             type: "MatchRule",
 *             action: "Block",
 *             match_condition: [
 *                 {
 *                     matchVariable: "RemoteAddr",
 *                     operator: "IPMatch",
 *                     negationCondition: false,
 *                     matchValues: ["192.168.1.0/24"],
 *                 },
 *                 {
 *                     matchVariable: "RequestHeader",
 *                     selector: "UserAgent",
 *                     operator: "Contains",
 *                     negationCondition: false,
 *                     matchValues: ["windows"],
 *                     transforms: [
 *                         "Lowercase",
 *                         "Trim",
 *                     ],
 *                 },
 *             ],
 *         },
 *     ],
 *     managed_rule: [
 *         {
 *             type: "DefaultRuleSet",
 *             version: "1.0",
 *             exclusion: [{
 *                 matchVariable: "QueryStringArgNames",
 *                 operator: "Equals",
 *                 selector: "not_suspicious",
 *             }],
 *             override: [
 *                 {
 *                     ruleGroupName: "PHP",
 *                     rule: [{
 *                         ruleId: "933100",
 *                         enabled: false,
 *                         action: "Block",
 *                     }],
 *                 },
 *                 {
 *                     ruleGroupName: "SQLI",
 *                     exclusion: [{
 *                         matchVariable: "QueryStringArgNames",
 *                         operator: "Equals",
 *                         selector: "really_not_suspicious",
 *                     }],
 *                     rule: [{
 *                         ruleId: "942200",
 *                         action: "Block",
 *                         exclusion: [{
 *                             matchVariable: "QueryStringArgNames",
 *                             operator: "Equals",
 *                             selector: "innocent",
 *                         }],
 *                     }],
 *                 },
 *             ],
 *         },
 *         {
 *             type: "Microsoft_BotManagerRuleSet",
 *             version: "1.0",
 *         },
 *     ],
 * });
 * ```
 */
export class FirewallPolicy extends pulumi.CustomResource {
    /**
     * Get an existing FirewallPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallPolicyState, opts?: pulumi.CustomResourceOptions): FirewallPolicy {
        return new FirewallPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:frontdoor/firewallPolicy:FirewallPolicy';

    /**
     * Returns true if the given object is an instance of FirewallPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallPolicy.__pulumiType;
    }

    /**
     * If a `customRule` block's action type is `block`, this is the response body. The body must be specified in base64 encoding.
     */
    public readonly customBlockResponseBody!: pulumi.Output<string | undefined>;
    /**
     * If a `customRule` block's action type is `block`, this is the response status code. Possible values are `200`, `403`, `405`, `406`, or `429`.
     */
    public readonly customBlockResponseStatusCode!: pulumi.Output<number | undefined>;
    /**
     * One or more `customRule` blocks as defined below.
     */
    public readonly customRules!: pulumi.Output<outputs.frontdoor.FirewallPolicyCustomRule[] | undefined>;
    /**
     * Is the policy a enabled state or disabled state. Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * the Frontend Endpoints associated with this Front Door Web Application Firewall policy.
     */
    public /*out*/ readonly frontendEndpointIds!: pulumi.Output<string[]>;
    /**
     * Resource location.
     */
    public /*out*/ readonly location!: pulumi.Output<string>;
    /**
     * One or more `managedRule` blocks as defined below.
     */
    public readonly managedRules!: pulumi.Output<outputs.frontdoor.FirewallPolicyManagedRule[] | undefined>;
    /**
     * The firewall policy mode. Possible values are `Detection`, `Prevention` and defaults to `Prevention`.
     */
    public readonly mode!: pulumi.Output<string | undefined>;
    /**
     * The name of the policy. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * If action type is redirect, this field represents redirect URL for the client.
     */
    public readonly redirectUrl!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource group. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the Web Application Firewall Policy.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a FirewallPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallPolicyArgs | FirewallPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FirewallPolicyState | undefined;
            inputs["customBlockResponseBody"] = state ? state.customBlockResponseBody : undefined;
            inputs["customBlockResponseStatusCode"] = state ? state.customBlockResponseStatusCode : undefined;
            inputs["customRules"] = state ? state.customRules : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["frontendEndpointIds"] = state ? state.frontendEndpointIds : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["managedRules"] = state ? state.managedRules : undefined;
            inputs["mode"] = state ? state.mode : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["redirectUrl"] = state ? state.redirectUrl : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as FirewallPolicyArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["customBlockResponseBody"] = args ? args.customBlockResponseBody : undefined;
            inputs["customBlockResponseStatusCode"] = args ? args.customBlockResponseStatusCode : undefined;
            inputs["customRules"] = args ? args.customRules : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["managedRules"] = args ? args.managedRules : undefined;
            inputs["mode"] = args ? args.mode : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["redirectUrl"] = args ? args.redirectUrl : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["frontendEndpointIds"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(FirewallPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallPolicy resources.
 */
export interface FirewallPolicyState {
    /**
     * If a `customRule` block's action type is `block`, this is the response body. The body must be specified in base64 encoding.
     */
    readonly customBlockResponseBody?: pulumi.Input<string>;
    /**
     * If a `customRule` block's action type is `block`, this is the response status code. Possible values are `200`, `403`, `405`, `406`, or `429`.
     */
    readonly customBlockResponseStatusCode?: pulumi.Input<number>;
    /**
     * One or more `customRule` blocks as defined below.
     */
    readonly customRules?: pulumi.Input<pulumi.Input<inputs.frontdoor.FirewallPolicyCustomRule>[]>;
    /**
     * Is the policy a enabled state or disabled state. Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * the Frontend Endpoints associated with this Front Door Web Application Firewall policy.
     */
    readonly frontendEndpointIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * One or more `managedRule` blocks as defined below.
     */
    readonly managedRules?: pulumi.Input<pulumi.Input<inputs.frontdoor.FirewallPolicyManagedRule>[]>;
    /**
     * The firewall policy mode. Possible values are `Detection`, `Prevention` and defaults to `Prevention`.
     */
    readonly mode?: pulumi.Input<string>;
    /**
     * The name of the policy. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * If action type is redirect, this field represents redirect URL for the client.
     */
    readonly redirectUrl?: pulumi.Input<string>;
    /**
     * The name of the resource group. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the Web Application Firewall Policy.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a FirewallPolicy resource.
 */
export interface FirewallPolicyArgs {
    /**
     * If a `customRule` block's action type is `block`, this is the response body. The body must be specified in base64 encoding.
     */
    readonly customBlockResponseBody?: pulumi.Input<string>;
    /**
     * If a `customRule` block's action type is `block`, this is the response status code. Possible values are `200`, `403`, `405`, `406`, or `429`.
     */
    readonly customBlockResponseStatusCode?: pulumi.Input<number>;
    /**
     * One or more `customRule` blocks as defined below.
     */
    readonly customRules?: pulumi.Input<pulumi.Input<inputs.frontdoor.FirewallPolicyCustomRule>[]>;
    /**
     * Is the policy a enabled state or disabled state. Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * One or more `managedRule` blocks as defined below.
     */
    readonly managedRules?: pulumi.Input<pulumi.Input<inputs.frontdoor.FirewallPolicyManagedRule>[]>;
    /**
     * The firewall policy mode. Possible values are `Detection`, `Prevention` and defaults to `Prevention`.
     */
    readonly mode?: pulumi.Input<string>;
    /**
     * The name of the policy. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * If action type is redirect, this field represents redirect URL for the client.
     */
    readonly redirectUrl?: pulumi.Input<string>;
    /**
     * The name of the resource group. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the Web Application Firewall Policy.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

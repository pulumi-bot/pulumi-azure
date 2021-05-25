// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * FrontDoor Web Application Firewall Policy can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:frontdoor/firewallPolicy:FirewallPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/examplefdwafpolicy
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
     * The Frontend Endpoints associated with this Front Door Web Application Firewall policy.
     */
    public /*out*/ readonly frontendEndpointIds!: pulumi.Output<string[]>;
    /**
     * The Azure Region where this FrontDoor Firewall Policy exists.
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
        opts = opts || {};
        if (opts.id) {
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
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
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
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
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
    customBlockResponseBody?: pulumi.Input<string>;
    /**
     * If a `customRule` block's action type is `block`, this is the response status code. Possible values are `200`, `403`, `405`, `406`, or `429`.
     */
    customBlockResponseStatusCode?: pulumi.Input<number>;
    /**
     * One or more `customRule` blocks as defined below.
     */
    customRules?: pulumi.Input<pulumi.Input<inputs.frontdoor.FirewallPolicyCustomRule>[]>;
    /**
     * Is the policy a enabled state or disabled state. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The Frontend Endpoints associated with this Front Door Web Application Firewall policy.
     */
    frontendEndpointIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Azure Region where this FrontDoor Firewall Policy exists.
     */
    location?: pulumi.Input<string>;
    /**
     * One or more `managedRule` blocks as defined below.
     */
    managedRules?: pulumi.Input<pulumi.Input<inputs.frontdoor.FirewallPolicyManagedRule>[]>;
    /**
     * The firewall policy mode. Possible values are `Detection`, `Prevention` and defaults to `Prevention`.
     */
    mode?: pulumi.Input<string>;
    /**
     * The name of the policy. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * If action type is redirect, this field represents redirect URL for the client.
     */
    redirectUrl?: pulumi.Input<string>;
    /**
     * The name of the resource group. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the Web Application Firewall Policy.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a FirewallPolicy resource.
 */
export interface FirewallPolicyArgs {
    /**
     * If a `customRule` block's action type is `block`, this is the response body. The body must be specified in base64 encoding.
     */
    customBlockResponseBody?: pulumi.Input<string>;
    /**
     * If a `customRule` block's action type is `block`, this is the response status code. Possible values are `200`, `403`, `405`, `406`, or `429`.
     */
    customBlockResponseStatusCode?: pulumi.Input<number>;
    /**
     * One or more `customRule` blocks as defined below.
     */
    customRules?: pulumi.Input<pulumi.Input<inputs.frontdoor.FirewallPolicyCustomRule>[]>;
    /**
     * Is the policy a enabled state or disabled state. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * One or more `managedRule` blocks as defined below.
     */
    managedRules?: pulumi.Input<pulumi.Input<inputs.frontdoor.FirewallPolicyManagedRule>[]>;
    /**
     * The firewall policy mode. Possible values are `Detection`, `Prevention` and defaults to `Prevention`.
     */
    mode?: pulumi.Input<string>;
    /**
     * The name of the policy. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * If action type is redirect, this field represents redirect URL for the client.
     */
    redirectUrl?: pulumi.Input<string>;
    /**
     * The name of the resource group. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the Web Application Firewall Policy.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

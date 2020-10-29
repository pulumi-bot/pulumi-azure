// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows you to Manages a Synapse Firewall Rule.
 */
export class FirewallRule extends pulumi.CustomResource {
    /**
     * Get an existing FirewallRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallRuleState, opts?: pulumi.CustomResourceOptions): FirewallRule {
        return new FirewallRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:synapse/firewallRule:FirewallRule';

    /**
     * Returns true if the given object is an instance of FirewallRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallRule.__pulumiType;
    }

    /**
     * The ending IP address to allow through the firewall for this rule.
     */
    public readonly endIpAddress!: pulumi.Output<string>;
    /**
     * The Name of the firewall rule. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The starting IP address to allow through the firewall for this rule.
     */
    public readonly startIpAddress!: pulumi.Output<string>;
    /**
     * The ID of the Synapse Workspace on which to create the Firewall Rule. Changing this forces a new resource to be created.
     */
    public readonly synapseWorkspaceId!: pulumi.Output<string>;

    /**
     * Create a FirewallRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallRuleArgs | FirewallRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FirewallRuleState | undefined;
            inputs["endIpAddress"] = state ? state.endIpAddress : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["startIpAddress"] = state ? state.startIpAddress : undefined;
            inputs["synapseWorkspaceId"] = state ? state.synapseWorkspaceId : undefined;
        } else {
            const args = argsOrState as FirewallRuleArgs | undefined;
            if (!args || args.endIpAddress === undefined) {
                throw new Error("Missing required property 'endIpAddress'");
            }
            if (!args || args.startIpAddress === undefined) {
                throw new Error("Missing required property 'startIpAddress'");
            }
            if (!args || args.synapseWorkspaceId === undefined) {
                throw new Error("Missing required property 'synapseWorkspaceId'");
            }
            inputs["endIpAddress"] = args ? args.endIpAddress : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["startIpAddress"] = args ? args.startIpAddress : undefined;
            inputs["synapseWorkspaceId"] = args ? args.synapseWorkspaceId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(FirewallRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallRule resources.
 */
export interface FirewallRuleState {
    /**
     * The ending IP address to allow through the firewall for this rule.
     */
    readonly endIpAddress?: pulumi.Input<string>;
    /**
     * The Name of the firewall rule. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The starting IP address to allow through the firewall for this rule.
     */
    readonly startIpAddress?: pulumi.Input<string>;
    /**
     * The ID of the Synapse Workspace on which to create the Firewall Rule. Changing this forces a new resource to be created.
     */
    readonly synapseWorkspaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallRule resource.
 */
export interface FirewallRuleArgs {
    /**
     * The ending IP address to allow through the firewall for this rule.
     */
    readonly endIpAddress: pulumi.Input<string>;
    /**
     * The Name of the firewall rule. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The starting IP address to allow through the firewall for this rule.
     */
    readonly startIpAddress: pulumi.Input<string>;
    /**
     * The ID of the Synapse Workspace on which to create the Firewall Rule. Changing this forces a new resource to be created.
     */
    readonly synapseWorkspaceId: pulumi.Input<string>;
}

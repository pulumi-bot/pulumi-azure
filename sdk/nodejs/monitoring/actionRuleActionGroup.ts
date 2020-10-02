// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Manages an Monitor Action Rule which type is action group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleActionGroup = new azure.monitoring.ActionGroup("exampleActionGroup", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     shortName: "exampleactiongroup",
 * });
 * const exampleActionRuleActionGroup = new azure.monitoring.ActionRuleActionGroup("exampleActionRuleActionGroup", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     actionGroupId: exampleActionGroup.id,
 *     scope: {
 *         type: "ResourceGroup",
 *         resourceIds: [exampleResourceGroup.id],
 *     },
 *     tags: {
 *         foo: "bar",
 *     },
 * });
 * ```
 */
export class ActionRuleActionGroup extends pulumi.CustomResource {
    /**
     * Get an existing ActionRuleActionGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ActionRuleActionGroupState, opts?: pulumi.CustomResourceOptions): ActionRuleActionGroup {
        return new ActionRuleActionGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:monitoring/actionRuleActionGroup:ActionRuleActionGroup';

    /**
     * Returns true if the given object is an instance of ActionRuleActionGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ActionRuleActionGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ActionRuleActionGroup.__pulumiType;
    }

    /**
     * Specifies the resource id of monitor action group.
     */
    public readonly actionGroupId!: pulumi.Output<string>;
    /**
     * A `condition` block as defined below.
     */
    public readonly condition!: pulumi.Output<outputs.monitoring.ActionRuleActionGroupCondition | undefined>;
    /**
     * Specifies a description for the Action Rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Is the Action Rule enabled? Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the name of the Monitor Action Rule. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the name of the resource group in which the Monitor Action Rule should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `scope` block as defined below.
     */
    public readonly scope!: pulumi.Output<outputs.monitoring.ActionRuleActionGroupScope | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a ActionRuleActionGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ActionRuleActionGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ActionRuleActionGroupArgs | ActionRuleActionGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ActionRuleActionGroupState | undefined;
            inputs["actionGroupId"] = state ? state.actionGroupId : undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["scope"] = state ? state.scope : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ActionRuleActionGroupArgs | undefined;
            if (!args || args.actionGroupId === undefined) {
                throw new Error("Missing required property 'actionGroupId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["actionGroupId"] = args ? args.actionGroupId : undefined;
            inputs["condition"] = args ? args.condition : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ActionRuleActionGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ActionRuleActionGroup resources.
 */
export interface ActionRuleActionGroupState {
    /**
     * Specifies the resource id of monitor action group.
     */
    readonly actionGroupId?: pulumi.Input<string>;
    /**
     * A `condition` block as defined below.
     */
    readonly condition?: pulumi.Input<inputs.monitoring.ActionRuleActionGroupCondition>;
    /**
     * Specifies a description for the Action Rule.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Is the Action Rule enabled? Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the Monitor Action Rule. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the name of the resource group in which the Monitor Action Rule should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A `scope` block as defined below.
     */
    readonly scope?: pulumi.Input<inputs.monitoring.ActionRuleActionGroupScope>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a ActionRuleActionGroup resource.
 */
export interface ActionRuleActionGroupArgs {
    /**
     * Specifies the resource id of monitor action group.
     */
    readonly actionGroupId: pulumi.Input<string>;
    /**
     * A `condition` block as defined below.
     */
    readonly condition?: pulumi.Input<inputs.monitoring.ActionRuleActionGroupCondition>;
    /**
     * Specifies a description for the Action Rule.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Is the Action Rule enabled? Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the Monitor Action Rule. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the name of the resource group in which the Monitor Action Rule should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `scope` block as defined below.
     */
    readonly scope?: pulumi.Input<inputs.monitoring.ActionRuleActionGroupScope>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

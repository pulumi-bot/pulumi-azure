// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a integer variable in Azure Automation
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US"});
 * const exampleAccount = new azure.automation.Account("exampleAccount", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: [{
 *         name: "Basic",
 *     }],
 * });
 * const exampleIntVariable = new azure.automation.IntVariable("exampleIntVariable", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     automationAccountName: exampleAccount.name,
 *     value: 1234,
 * });
 * ```
 *
 * ## Import
 *
 * Automation Int Variable can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:automation/intVariable:IntVariable example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tfex-example-rg/providers/Microsoft.Automation/automationAccounts/tfex-example-account/variables/tfex-example-var
 * ```
 */
export class IntVariable extends pulumi.CustomResource {
    /**
     * Get an existing IntVariable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntVariableState, opts?: pulumi.CustomResourceOptions): IntVariable {
        return new IntVariable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:automation/intVariable:IntVariable';

    /**
     * Returns true if the given object is an instance of IntVariable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IntVariable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IntVariable.__pulumiType;
    }

    /**
     * The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
     */
    public readonly automationAccountName!: pulumi.Output<string>;
    /**
     * The description of the Automation Variable.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies if the Automation Variable is encrypted. Defaults to `false`.
     */
    public readonly encrypted!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the Automation Variable. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The value of the Automation Variable as a `integer`.
     */
    public readonly value!: pulumi.Output<number | undefined>;

    /**
     * Create a IntVariable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntVariableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntVariableArgs | IntVariableState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IntVariableState | undefined;
            inputs["automationAccountName"] = state ? state.automationAccountName : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["encrypted"] = state ? state.encrypted : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as IntVariableArgs | undefined;
            if ((!args || args.automationAccountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'automationAccountName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["automationAccountName"] = args ? args.automationAccountName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["encrypted"] = args ? args.encrypted : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["value"] = args ? args.value : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(IntVariable.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IntVariable resources.
 */
export interface IntVariableState {
    /**
     * The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
     */
    readonly automationAccountName?: pulumi.Input<string>;
    /**
     * The description of the Automation Variable.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies if the Automation Variable is encrypted. Defaults to `false`.
     */
    readonly encrypted?: pulumi.Input<boolean>;
    /**
     * The name of the Automation Variable. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The value of the Automation Variable as a `integer`.
     */
    readonly value?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a IntVariable resource.
 */
export interface IntVariableArgs {
    /**
     * The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
     */
    readonly automationAccountName: pulumi.Input<string>;
    /**
     * The description of the Automation Variable.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies if the Automation Variable is encrypted. Defaults to `false`.
     */
    readonly encrypted?: pulumi.Input<boolean>;
    /**
     * The name of the Automation Variable. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The value of the Automation Variable as a `integer`.
     */
    readonly value?: pulumi.Input<number>;
}

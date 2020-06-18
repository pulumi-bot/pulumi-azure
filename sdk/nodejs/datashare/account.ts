// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Data Share Account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAccount = new azure.datashare.Account("exampleAccount", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     identity: {
 *         type: "SystemAssigned",
 *     },
 *     tags: {
 *         foo: "bar",
 *     },
 * });
 * ```
 */
export class Account extends pulumi.CustomResource {
    /**
     * Get an existing Account resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountState, opts?: pulumi.CustomResourceOptions): Account {
        return new Account(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:datashare/account:Account';

    /**
     * Returns true if the given object is an instance of Account.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Account {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Account.__pulumiType;
    }

    /**
     * An `identity` block as defined below.
     */
    public readonly identity!: pulumi.Output<outputs.datashare.AccountIdentity>;
    /**
     * The Azure Region where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name which should be used for this Data Share Account. Changing this forces a new Data Share Account to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Resource Group where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags which should be assigned to the Data Share Account.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Account resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountArgs | AccountState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AccountState | undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as AccountArgs | undefined;
            if (!args || args.identity === undefined) {
                throw new Error("Missing required property 'identity'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Account.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Account resources.
 */
export interface AccountState {
    /**
     * An `identity` block as defined below.
     */
    readonly identity?: pulumi.Input<inputs.datashare.AccountIdentity>;
    /**
     * The Azure Region where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name which should be used for this Data Share Account. Changing this forces a new Data Share Account to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Data Share Account.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Account resource.
 */
export interface AccountArgs {
    /**
     * An `identity` block as defined below.
     */
    readonly identity: pulumi.Input<inputs.datashare.AccountIdentity>;
    /**
     * The Azure Region where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name which should be used for this Data Share Account. Changing this forces a new Data Share Account to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Data Share Account.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

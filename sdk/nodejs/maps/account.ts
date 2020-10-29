// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Azure Maps Account.
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
    public static readonly __pulumiType = 'azure:maps/account:Account';

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
     * The name of the Azure Maps Account. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The primary key used to authenticate and authorize access to the Maps REST APIs.
     */
    public /*out*/ readonly primaryAccessKey!: pulumi.Output<string>;
    /**
     * The name of the Resource Group in which the Azure Maps Account should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The secondary key used to authenticate and authorize access to the Maps REST APIs.
     */
    public /*out*/ readonly secondaryAccessKey!: pulumi.Output<string>;
    /**
     * The sku of the Azure Maps Account. Possible values are `S0` and `S1`.
     */
    public readonly skuName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the Azure Maps Account.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A unique identifier for the Maps Account.
     */
    public /*out*/ readonly xMsClientId!: pulumi.Output<string>;

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
            inputs["name"] = state ? state.name : undefined;
            inputs["primaryAccessKey"] = state ? state.primaryAccessKey : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["secondaryAccessKey"] = state ? state.secondaryAccessKey : undefined;
            inputs["skuName"] = state ? state.skuName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["xMsClientId"] = state ? state.xMsClientId : undefined;
        } else {
            const args = argsOrState as AccountArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.skuName === undefined) {
                throw new Error("Missing required property 'skuName'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["skuName"] = args ? args.skuName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["primaryAccessKey"] = undefined /*out*/;
            inputs["secondaryAccessKey"] = undefined /*out*/;
            inputs["xMsClientId"] = undefined /*out*/;
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
     * The name of the Azure Maps Account. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The primary key used to authenticate and authorize access to the Maps REST APIs.
     */
    readonly primaryAccessKey?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the Azure Maps Account should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The secondary key used to authenticate and authorize access to the Maps REST APIs.
     */
    readonly secondaryAccessKey?: pulumi.Input<string>;
    /**
     * The sku of the Azure Maps Account. Possible values are `S0` and `S1`.
     */
    readonly skuName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the Azure Maps Account.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A unique identifier for the Maps Account.
     */
    readonly xMsClientId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Account resource.
 */
export interface AccountArgs {
    /**
     * The name of the Azure Maps Account. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the Azure Maps Account should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The sku of the Azure Maps Account. Possible values are `S0` and `S1`.
     */
    readonly skuName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the Azure Maps Account.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

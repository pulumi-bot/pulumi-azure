// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Bot Channels Registration.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const current = azure.core.getClientConfig({});
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "northeurope"});
 * const exampleChannelsRegistration = new azure.bot.ChannelsRegistration("exampleChannelsRegistration", {
 *     location: "global",
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "F0",
 *     microsoftAppId: current.then(current => current.clientId),
 * });
 * ```
 *
 * ## Import
 *
 * Bot Channels Registration can be imported using the `resource id`, e.g.
 */
export class ChannelsRegistration extends pulumi.CustomResource {
    /**
     * Get an existing ChannelsRegistration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ChannelsRegistrationState, opts?: pulumi.CustomResourceOptions): ChannelsRegistration {
        return new ChannelsRegistration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:bot/channelsRegistration:ChannelsRegistration';

    /**
     * Returns true if the given object is an instance of ChannelsRegistration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ChannelsRegistration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ChannelsRegistration.__pulumiType;
    }

    /**
     * The Application Insights API Key to associate with the Bot Channels Registration.
     */
    public readonly developerAppInsightsApiKey!: pulumi.Output<string>;
    /**
     * The Application Insights Application ID to associate with the Bot Channels Registration.
     */
    public readonly developerAppInsightsApplicationId!: pulumi.Output<string>;
    /**
     * The Application Insights Key to associate with the Bot Channels Registration.
     */
    public readonly developerAppInsightsKey!: pulumi.Output<string>;
    /**
     * The name of the Bot Channels Registration will be displayed as. This defaults to `name` if not specified.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The Bot Channels Registration endpoint.
     */
    public readonly endpoint!: pulumi.Output<string | undefined>;
    /**
     * The supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
     */
    public readonly microsoftAppId!: pulumi.Output<string>;
    /**
     * Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Bot Channels Registration. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The SKU of the Bot Channels Registration. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
     */
    public readonly sku!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a ChannelsRegistration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ChannelsRegistrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ChannelsRegistrationArgs | ChannelsRegistrationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ChannelsRegistrationState | undefined;
            inputs["developerAppInsightsApiKey"] = state ? state.developerAppInsightsApiKey : undefined;
            inputs["developerAppInsightsApplicationId"] = state ? state.developerAppInsightsApplicationId : undefined;
            inputs["developerAppInsightsKey"] = state ? state.developerAppInsightsKey : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["microsoftAppId"] = state ? state.microsoftAppId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ChannelsRegistrationArgs | undefined;
            if (!args || args.microsoftAppId === undefined) {
                throw new Error("Missing required property 'microsoftAppId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["developerAppInsightsApiKey"] = args ? args.developerAppInsightsApiKey : undefined;
            inputs["developerAppInsightsApplicationId"] = args ? args.developerAppInsightsApplicationId : undefined;
            inputs["developerAppInsightsKey"] = args ? args.developerAppInsightsKey : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["endpoint"] = args ? args.endpoint : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["microsoftAppId"] = args ? args.microsoftAppId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ChannelsRegistration.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ChannelsRegistration resources.
 */
export interface ChannelsRegistrationState {
    /**
     * The Application Insights API Key to associate with the Bot Channels Registration.
     */
    readonly developerAppInsightsApiKey?: pulumi.Input<string>;
    /**
     * The Application Insights Application ID to associate with the Bot Channels Registration.
     */
    readonly developerAppInsightsApplicationId?: pulumi.Input<string>;
    /**
     * The Application Insights Key to associate with the Bot Channels Registration.
     */
    readonly developerAppInsightsKey?: pulumi.Input<string>;
    /**
     * The name of the Bot Channels Registration will be displayed as. This defaults to `name` if not specified.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The Bot Channels Registration endpoint.
     */
    readonly endpoint?: pulumi.Input<string>;
    /**
     * The supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
     */
    readonly microsoftAppId?: pulumi.Input<string>;
    /**
     * Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Bot Channels Registration. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The SKU of the Bot Channels Registration. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
     */
    readonly sku?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a ChannelsRegistration resource.
 */
export interface ChannelsRegistrationArgs {
    /**
     * The Application Insights API Key to associate with the Bot Channels Registration.
     */
    readonly developerAppInsightsApiKey?: pulumi.Input<string>;
    /**
     * The Application Insights Application ID to associate with the Bot Channels Registration.
     */
    readonly developerAppInsightsApplicationId?: pulumi.Input<string>;
    /**
     * The Application Insights Key to associate with the Bot Channels Registration.
     */
    readonly developerAppInsightsKey?: pulumi.Input<string>;
    /**
     * The name of the Bot Channels Registration will be displayed as. This defaults to `name` if not specified.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The Bot Channels Registration endpoint.
     */
    readonly endpoint?: pulumi.Input<string>;
    /**
     * The supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
     */
    readonly microsoftAppId: pulumi.Input<string>;
    /**
     * Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Bot Channels Registration. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The SKU of the Bot Channels Registration. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
     */
    readonly sku: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

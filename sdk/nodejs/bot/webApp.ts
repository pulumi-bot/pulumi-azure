// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Bot Web App.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const current = azure.core.getClientConfig({});
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleWebApp = new azure.bot.WebApp("exampleWebApp", {
 *     location: "global",
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "F0",
 *     microsoftAppId: current.then(current => current.clientId),
 * });
 * ```
 *
 * ## Import
 *
 * Bot Web App's can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:bot/webApp:WebApp example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example
 * ```
 */
export class WebApp extends pulumi.CustomResource {
    /**
     * Get an existing WebApp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebAppState, opts?: pulumi.CustomResourceOptions): WebApp {
        return new WebApp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:bot/webApp:WebApp';

    /**
     * Returns true if the given object is an instance of WebApp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebApp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebApp.__pulumiType;
    }

    /**
     * The Application Insights API Key to associate with the Web App Bot.
     */
    public readonly developerAppInsightsApiKey!: pulumi.Output<string>;
    /**
     * The Application Insights Application ID to associate with the Web App Bot.
     */
    public readonly developerAppInsightsApplicationId!: pulumi.Output<string>;
    /**
     * The Application Insights Key to associate with the Web App Bot.
     */
    public readonly developerAppInsightsKey!: pulumi.Output<string>;
    /**
     * The name of the Web App Bot will be displayed as. This defaults to `name` if not specified.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The Web App Bot endpoint.
     */
    public readonly endpoint!: pulumi.Output<string | undefined>;
    /**
     * The supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * A list of LUIS App IDs to associate with the Web App Bot.
     */
    public readonly luisAppIds!: pulumi.Output<string[] | undefined>;
    /**
     * The LUIS key to associate with the Web App Bot.
     */
    public readonly luisKey!: pulumi.Output<string | undefined>;
    /**
     * The Microsoft Application ID for the Web App Bot. Changing this forces a new resource to be created.
     */
    public readonly microsoftAppId!: pulumi.Output<string>;
    /**
     * Specifies the name of the Web App Bot. Changing this forces a new resource to be created. Must be globally unique.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Web App Bot. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The SKU of the Web App Bot. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
     */
    public readonly sku!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a WebApp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebAppArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebAppArgs | WebAppState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebAppState | undefined;
            inputs["developerAppInsightsApiKey"] = state ? state.developerAppInsightsApiKey : undefined;
            inputs["developerAppInsightsApplicationId"] = state ? state.developerAppInsightsApplicationId : undefined;
            inputs["developerAppInsightsKey"] = state ? state.developerAppInsightsKey : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["luisAppIds"] = state ? state.luisAppIds : undefined;
            inputs["luisKey"] = state ? state.luisKey : undefined;
            inputs["microsoftAppId"] = state ? state.microsoftAppId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as WebAppArgs | undefined;
            if ((!args || args.microsoftAppId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'microsoftAppId'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.sku === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["developerAppInsightsApiKey"] = args ? args.developerAppInsightsApiKey : undefined;
            inputs["developerAppInsightsApplicationId"] = args ? args.developerAppInsightsApplicationId : undefined;
            inputs["developerAppInsightsKey"] = args ? args.developerAppInsightsKey : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["endpoint"] = args ? args.endpoint : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["luisAppIds"] = args ? args.luisAppIds : undefined;
            inputs["luisKey"] = args ? args.luisKey : undefined;
            inputs["microsoftAppId"] = args ? args.microsoftAppId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WebApp.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebApp resources.
 */
export interface WebAppState {
    /**
     * The Application Insights API Key to associate with the Web App Bot.
     */
    developerAppInsightsApiKey?: pulumi.Input<string>;
    /**
     * The Application Insights Application ID to associate with the Web App Bot.
     */
    developerAppInsightsApplicationId?: pulumi.Input<string>;
    /**
     * The Application Insights Key to associate with the Web App Bot.
     */
    developerAppInsightsKey?: pulumi.Input<string>;
    /**
     * The name of the Web App Bot will be displayed as. This defaults to `name` if not specified.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The Web App Bot endpoint.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * The supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * A list of LUIS App IDs to associate with the Web App Bot.
     */
    luisAppIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The LUIS key to associate with the Web App Bot.
     */
    luisKey?: pulumi.Input<string>;
    /**
     * The Microsoft Application ID for the Web App Bot. Changing this forces a new resource to be created.
     */
    microsoftAppId?: pulumi.Input<string>;
    /**
     * Specifies the name of the Web App Bot. Changing this forces a new resource to be created. Must be globally unique.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Web App Bot. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The SKU of the Web App Bot. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
     */
    sku?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a WebApp resource.
 */
export interface WebAppArgs {
    /**
     * The Application Insights API Key to associate with the Web App Bot.
     */
    developerAppInsightsApiKey?: pulumi.Input<string>;
    /**
     * The Application Insights Application ID to associate with the Web App Bot.
     */
    developerAppInsightsApplicationId?: pulumi.Input<string>;
    /**
     * The Application Insights Key to associate with the Web App Bot.
     */
    developerAppInsightsKey?: pulumi.Input<string>;
    /**
     * The name of the Web App Bot will be displayed as. This defaults to `name` if not specified.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The Web App Bot endpoint.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * The supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * A list of LUIS App IDs to associate with the Web App Bot.
     */
    luisAppIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The LUIS key to associate with the Web App Bot.
     */
    luisKey?: pulumi.Input<string>;
    /**
     * The Microsoft Application ID for the Web App Bot. Changing this forces a new resource to be created.
     */
    microsoftAppId: pulumi.Input<string>;
    /**
     * Specifies the name of the Web App Bot. Changing this forces a new resource to be created. Must be globally unique.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Web App Bot. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The SKU of the Web App Bot. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
     */
    sku: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

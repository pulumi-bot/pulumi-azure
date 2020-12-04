// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Logger within an API Management Service.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US"});
 * const exampleInsights = new azure.appinsights.Insights("exampleInsights", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     applicationType: "other",
 * });
 * const exampleService = new azure.apimanagement.Service("exampleService", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     publisherName: "My Company",
 *     publisherEmail: "company@exmaple.com",
 *     skuName: "Developer_1",
 * });
 * const exampleLogger = new azure.apimanagement.Logger("exampleLogger", {
 *     apiManagementName: exampleService.name,
 *     resourceGroupName: exampleResourceGroup.name,
 *     applicationInsights: {
 *         instrumentationKey: exampleInsights.instrumentationKey,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * API Management Loggers can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:apimanagement/logger:Logger example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/example-rg/Microsoft.ApiManagement/service/example-apim/loggers/example-logger
 * ```
 */
export class Logger extends pulumi.CustomResource {
    /**
     * Get an existing Logger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoggerState, opts?: pulumi.CustomResourceOptions): Logger {
        return new Logger(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:apimanagement/logger:Logger';

    /**
     * Returns true if the given object is an instance of Logger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Logger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Logger.__pulumiType;
    }

    /**
     * The name of the API Management Service. Changing this forces a new resource to be created.
     */
    public readonly apiManagementName!: pulumi.Output<string>;
    /**
     * An `applicationInsights` block as documented below.
     */
    public readonly applicationInsights!: pulumi.Output<outputs.apimanagement.LoggerApplicationInsights | undefined>;
    /**
     * Specifies whether records should be buffered in the Logger prior to publishing. Defaults to `true`.
     */
    public readonly buffered!: pulumi.Output<boolean | undefined>;
    /**
     * A description of this Logger.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * An `eventhub` block as documented below.
     */
    public readonly eventhub!: pulumi.Output<outputs.apimanagement.LoggerEventhub | undefined>;
    /**
     * The name of this Logger, which must be unique within the API Management Service. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;

    /**
     * Create a Logger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoggerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoggerArgs | LoggerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LoggerState | undefined;
            inputs["apiManagementName"] = state ? state.apiManagementName : undefined;
            inputs["applicationInsights"] = state ? state.applicationInsights : undefined;
            inputs["buffered"] = state ? state.buffered : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["eventhub"] = state ? state.eventhub : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
        } else {
            const args = argsOrState as LoggerArgs | undefined;
            if ((!args || args.apiManagementName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'apiManagementName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["apiManagementName"] = args ? args.apiManagementName : undefined;
            inputs["applicationInsights"] = args ? args.applicationInsights : undefined;
            inputs["buffered"] = args ? args.buffered : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["eventhub"] = args ? args.eventhub : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Logger.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Logger resources.
 */
export interface LoggerState {
    /**
     * The name of the API Management Service. Changing this forces a new resource to be created.
     */
    readonly apiManagementName?: pulumi.Input<string>;
    /**
     * An `applicationInsights` block as documented below.
     */
    readonly applicationInsights?: pulumi.Input<inputs.apimanagement.LoggerApplicationInsights>;
    /**
     * Specifies whether records should be buffered in the Logger prior to publishing. Defaults to `true`.
     */
    readonly buffered?: pulumi.Input<boolean>;
    /**
     * A description of this Logger.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * An `eventhub` block as documented below.
     */
    readonly eventhub?: pulumi.Input<inputs.apimanagement.LoggerEventhub>;
    /**
     * The name of this Logger, which must be unique within the API Management Service. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Logger resource.
 */
export interface LoggerArgs {
    /**
     * The name of the API Management Service. Changing this forces a new resource to be created.
     */
    readonly apiManagementName: pulumi.Input<string>;
    /**
     * An `applicationInsights` block as documented below.
     */
    readonly applicationInsights?: pulumi.Input<inputs.apimanagement.LoggerApplicationInsights>;
    /**
     * Specifies whether records should be buffered in the Logger prior to publishing. Defaults to `true`.
     */
    readonly buffered?: pulumi.Input<boolean>;
    /**
     * A description of this Logger.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * An `eventhub` block as documented below.
     */
    readonly eventhub?: pulumi.Input<inputs.apimanagement.LoggerEventhub>;
    /**
     * The name of this Logger, which must be unique within the API Management Service. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}

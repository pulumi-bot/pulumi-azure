// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an iot security solution.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleIoTHub = new azure.iot.IoTHub("exampleIoTHub", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     sku: {
 *         name: "S1",
 *         capacity: "1",
 *     },
 * });
 * const exampleSecuritySolution = new azure.iot.SecuritySolution("exampleSecuritySolution", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     displayName: "Iot Security Solution",
 *     iothubIds: [exampleIoTHub.id],
 * });
 * ```
 *
 * ## Import
 *
 * Iot Security Solution can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:iot/securitySolution:SecuritySolution example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Security/IoTSecuritySolutions/solution1
 * ```
 */
export class SecuritySolution extends pulumi.CustomResource {
    /**
     * Get an existing SecuritySolution resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecuritySolutionState, opts?: pulumi.CustomResourceOptions): SecuritySolution {
        return new SecuritySolution(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:iot/securitySolution:SecuritySolution';

    /**
     * Returns true if the given object is an instance of SecuritySolution.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecuritySolution {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecuritySolution.__pulumiType;
    }

    /**
     * Specifies the Display Name for this Iot Security Solution.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Is the Iot Security Solution enabled? Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * A list of data which is to exported to analytic workspace. Valid values include `RawEvents`.
     */
    public readonly eventsToExports!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the IoT Hub resource IDs to which this Iot Security Solution is applied.
     */
    public readonly iothubIds!: pulumi.Output<string[]>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the Log Analytics Workspace ID to which the security data will be sent.
     */
    public readonly logAnalyticsWorkspaceId!: pulumi.Output<string | undefined>;
    /**
     * Should ip addressed be unmasked in the log? Defaults to `false`.
     */
    public readonly logUnmaskedIpsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the name of the Iot Security Solution. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * An Azure Resource Graph query used to set the resources monitored.
     */
    public readonly queryForResources!: pulumi.Output<string>;
    /**
     * A list of subscription Ids on which the user defined resources query should be executed.
     */
    public readonly querySubscriptionIds!: pulumi.Output<string[]>;
    /**
     * A `recommendationsEnabled` block of options to enable or disable as defined below.
     */
    public readonly recommendationsEnabled!: pulumi.Output<outputs.iot.SecuritySolutionRecommendationsEnabled>;
    /**
     * Specifies the name of the resource group in which to create the Iot Security Solution. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a SecuritySolution resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecuritySolutionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecuritySolutionArgs | SecuritySolutionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecuritySolutionState | undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["eventsToExports"] = state ? state.eventsToExports : undefined;
            inputs["iothubIds"] = state ? state.iothubIds : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["logAnalyticsWorkspaceId"] = state ? state.logAnalyticsWorkspaceId : undefined;
            inputs["logUnmaskedIpsEnabled"] = state ? state.logUnmaskedIpsEnabled : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["queryForResources"] = state ? state.queryForResources : undefined;
            inputs["querySubscriptionIds"] = state ? state.querySubscriptionIds : undefined;
            inputs["recommendationsEnabled"] = state ? state.recommendationsEnabled : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as SecuritySolutionArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.iothubIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'iothubIds'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["eventsToExports"] = args ? args.eventsToExports : undefined;
            inputs["iothubIds"] = args ? args.iothubIds : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["logAnalyticsWorkspaceId"] = args ? args.logAnalyticsWorkspaceId : undefined;
            inputs["logUnmaskedIpsEnabled"] = args ? args.logUnmaskedIpsEnabled : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["queryForResources"] = args ? args.queryForResources : undefined;
            inputs["querySubscriptionIds"] = args ? args.querySubscriptionIds : undefined;
            inputs["recommendationsEnabled"] = args ? args.recommendationsEnabled : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SecuritySolution.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecuritySolution resources.
 */
export interface SecuritySolutionState {
    /**
     * Specifies the Display Name for this Iot Security Solution.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Is the Iot Security Solution enabled? Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * A list of data which is to exported to analytic workspace. Valid values include `RawEvents`.
     */
    eventsToExports?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the IoT Hub resource IDs to which this Iot Security Solution is applied.
     */
    iothubIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies the Log Analytics Workspace ID to which the security data will be sent.
     */
    logAnalyticsWorkspaceId?: pulumi.Input<string>;
    /**
     * Should ip addressed be unmasked in the log? Defaults to `false`.
     */
    logUnmaskedIpsEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the Iot Security Solution. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * An Azure Resource Graph query used to set the resources monitored.
     */
    queryForResources?: pulumi.Input<string>;
    /**
     * A list of subscription Ids on which the user defined resources query should be executed.
     */
    querySubscriptionIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A `recommendationsEnabled` block of options to enable or disable as defined below.
     */
    recommendationsEnabled?: pulumi.Input<inputs.iot.SecuritySolutionRecommendationsEnabled>;
    /**
     * Specifies the name of the resource group in which to create the Iot Security Solution. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a SecuritySolution resource.
 */
export interface SecuritySolutionArgs {
    /**
     * Specifies the Display Name for this Iot Security Solution.
     */
    displayName: pulumi.Input<string>;
    /**
     * Is the Iot Security Solution enabled? Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * A list of data which is to exported to analytic workspace. Valid values include `RawEvents`.
     */
    eventsToExports?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the IoT Hub resource IDs to which this Iot Security Solution is applied.
     */
    iothubIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies the Log Analytics Workspace ID to which the security data will be sent.
     */
    logAnalyticsWorkspaceId?: pulumi.Input<string>;
    /**
     * Should ip addressed be unmasked in the log? Defaults to `false`.
     */
    logUnmaskedIpsEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the Iot Security Solution. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * An Azure Resource Graph query used to set the resources monitored.
     */
    queryForResources?: pulumi.Input<string>;
    /**
     * A list of subscription Ids on which the user defined resources query should be executed.
     */
    querySubscriptionIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A `recommendationsEnabled` block of options to enable or disable as defined below.
     */
    recommendationsEnabled?: pulumi.Input<inputs.iot.SecuritySolutionRecommendationsEnabled>;
    /**
     * Specifies the name of the resource group in which to create the Iot Security Solution. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

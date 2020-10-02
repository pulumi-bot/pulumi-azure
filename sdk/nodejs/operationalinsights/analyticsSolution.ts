// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Manages a Log Analytics (formally Operational Insights) Solution.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * import * as random from "@pulumi/random";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "westeurope"});
 * const workspace = new random.RandomId("workspace", {
 *     keepers: {
 *         group_name: exampleResourceGroup.name,
 *     },
 *     byteLength: 8,
 * });
 * const exampleAnalyticsWorkspace = new azure.operationalinsights.AnalyticsWorkspace("exampleAnalyticsWorkspace", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "PerGB2018",
 * });
 * const exampleAnalyticsSolution = new azure.operationalinsights.AnalyticsSolution("exampleAnalyticsSolution", {
 *     solutionName: "ContainerInsights",
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     workspaceResourceId: exampleAnalyticsWorkspace.id,
 *     workspaceName: exampleAnalyticsWorkspace.name,
 *     plan: {
 *         publisher: "Microsoft",
 *         product: "OMSGallery/ContainerInsights",
 *     },
 * });
 * ```
 */
export class AnalyticsSolution extends pulumi.CustomResource {
    /**
     * Get an existing AnalyticsSolution resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AnalyticsSolutionState, opts?: pulumi.CustomResourceOptions): AnalyticsSolution {
        return new AnalyticsSolution(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:operationalinsights/analyticsSolution:AnalyticsSolution';

    /**
     * Returns true if the given object is an instance of AnalyticsSolution.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AnalyticsSolution {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AnalyticsSolution.__pulumiType;
    }

    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * A `plan` block as documented below.
     */
    public readonly plan!: pulumi.Output<outputs.operationalinsights.AnalyticsSolutionPlan>;
    /**
     * The name of the resource group in which the Log Analytics solution is created. Changing this forces a new resource to be created. Note: The solution and its related workspace can only exist in the same resource group.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Specifies the name of the solution to be deployed. See [here for options](https://docs.microsoft.com/en-us/azure/log-analytics/log-analytics-add-solutions).Changing this forces a new resource to be created.
     */
    public readonly solutionName!: pulumi.Output<string>;
    /**
     * The full name of the Log Analytics workspace with which the solution will be linked. Changing this forces a new resource to be created.
     */
    public readonly workspaceName!: pulumi.Output<string>;
    /**
     * The full resource ID of the Log Analytics workspace with which the solution will be linked. Changing this forces a new resource to be created.
     */
    public readonly workspaceResourceId!: pulumi.Output<string>;

    /**
     * Create a AnalyticsSolution resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AnalyticsSolutionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AnalyticsSolutionArgs | AnalyticsSolutionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AnalyticsSolutionState | undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["plan"] = state ? state.plan : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["solutionName"] = state ? state.solutionName : undefined;
            inputs["workspaceName"] = state ? state.workspaceName : undefined;
            inputs["workspaceResourceId"] = state ? state.workspaceResourceId : undefined;
        } else {
            const args = argsOrState as AnalyticsSolutionArgs | undefined;
            if (!args || args.plan === undefined) {
                throw new Error("Missing required property 'plan'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.solutionName === undefined) {
                throw new Error("Missing required property 'solutionName'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            if (!args || args.workspaceResourceId === undefined) {
                throw new Error("Missing required property 'workspaceResourceId'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["plan"] = args ? args.plan : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["solutionName"] = args ? args.solutionName : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["workspaceResourceId"] = args ? args.workspaceResourceId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AnalyticsSolution.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AnalyticsSolution resources.
 */
export interface AnalyticsSolutionState {
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * A `plan` block as documented below.
     */
    readonly plan?: pulumi.Input<inputs.operationalinsights.AnalyticsSolutionPlan>;
    /**
     * The name of the resource group in which the Log Analytics solution is created. Changing this forces a new resource to be created. Note: The solution and its related workspace can only exist in the same resource group.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies the name of the solution to be deployed. See [here for options](https://docs.microsoft.com/en-us/azure/log-analytics/log-analytics-add-solutions).Changing this forces a new resource to be created.
     */
    readonly solutionName?: pulumi.Input<string>;
    /**
     * The full name of the Log Analytics workspace with which the solution will be linked. Changing this forces a new resource to be created.
     */
    readonly workspaceName?: pulumi.Input<string>;
    /**
     * The full resource ID of the Log Analytics workspace with which the solution will be linked. Changing this forces a new resource to be created.
     */
    readonly workspaceResourceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AnalyticsSolution resource.
 */
export interface AnalyticsSolutionArgs {
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * A `plan` block as documented below.
     */
    readonly plan: pulumi.Input<inputs.operationalinsights.AnalyticsSolutionPlan>;
    /**
     * The name of the resource group in which the Log Analytics solution is created. Changing this forces a new resource to be created. Note: The solution and its related workspace can only exist in the same resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the name of the solution to be deployed. See [here for options](https://docs.microsoft.com/en-us/azure/log-analytics/log-analytics-add-solutions).Changing this forces a new resource to be created.
     */
    readonly solutionName: pulumi.Input<string>;
    /**
     * The full name of the Log Analytics workspace with which the solution will be linked. Changing this forces a new resource to be created.
     */
    readonly workspaceName: pulumi.Input<string>;
    /**
     * The full resource ID of the Log Analytics workspace with which the solution will be linked. Changing this forces a new resource to be created.
     */
    readonly workspaceResourceId: pulumi.Input<string>;
}

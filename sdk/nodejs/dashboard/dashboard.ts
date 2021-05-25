// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a shared dashboard in the Azure Portal.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const config = new pulumi.Config();
 * const mdContent = config.get("mdContent") || "# Hello all :)";
 * const videoLink = config.get("videoLink") || "https://www.youtube.com/watch?v=......";
 * const current = azure.core.getSubscription({});
 * const my_group = new azure.core.ResourceGroup("my-group", {location: "West Europe"});
 * const my_board = new azure.dashboard.Dashboard("my-board", {
 *     resourceGroupName: my_group.name,
 *     location: my_group.location,
 *     tags: {
 *         source: "managed",
 *     },
 *     dashboardProperties: current.then(current => `{
 *    "lenses": {
 *         "0": {
 *             "order": 0,
 *             "parts": {
 *                 "0": {
 *                     "position": {
 *                         "x": 0,
 *                         "y": 0,
 *                         "rowSpan": 2,
 *                         "colSpan": 3
 *                     },
 *                     "metadata": {
 *                         "inputs": [],
 *                         "type": "Extension/HubsExtension/PartType/MarkdownPart",
 *                         "settings": {
 *                             "content": {
 *                                 "settings": {
 *                                     "content": "${mdContent}",
 *                                     "subtitle": "",
 *                                     "title": ""
 *                                 }
 *                             }
 *                         }
 *                     }
 *                 },               
 *                 "1": {
 *                     "position": {
 *                         "x": 5,
 *                         "y": 0,
 *                         "rowSpan": 4,
 *                         "colSpan": 6
 *                     },
 *                     "metadata": {
 *                         "inputs": [],
 *                         "type": "Extension/HubsExtension/PartType/VideoPart",
 *                         "settings": {
 *                             "content": {
 *                                 "settings": {
 *                                     "title": "Important Information",
 *                                     "subtitle": "",
 *                                     "src": "${videoLink}",
 *                                     "autoplay": true
 *                                 }
 *                             }
 *                         }
 *                     }
 *                 },
 *                 "2": {
 *                     "position": {
 *                         "x": 0,
 *                         "y": 4,
 *                         "rowSpan": 4,
 *                         "colSpan": 6
 *                     },
 *                     "metadata": {
 *                         "inputs": [
 *                             {
 *                                 "name": "ComponentId",
 *                                 "value": "/subscriptions/${current.subscriptionId}/resourceGroups/myRG/providers/microsoft.insights/components/myWebApp"
 *                             }
 *                         ],
 *                         "type": "Extension/AppInsightsExtension/PartType/AppMapGalPt",
 *                         "settings": {},
 *                         "asset": {
 *                             "idInputName": "ComponentId",
 *                             "type": "ApplicationInsights"
 *                         }
 *                     }
 *                 }              
 *             }
 *         }
 *     },
 *     "metadata": {
 *         "model": {
 *             "timeRange": {
 *                 "value": {
 *                     "relative": {
 *                         "duration": 24,
 *                         "timeUnit": 1
 *                     }
 *                 },
 *                 "type": "MsPortalFx.Composition.Configuration.ValueTypes.TimeRange"
 *             },
 *             "filterLocale": {
 *                 "value": "en-us"
 *             },
 *             "filters": {
 *                 "value": {
 *                     "MsPortalFx_TimeRange": {
 *                         "model": {
 *                             "format": "utc",
 *                             "granularity": "auto",
 *                             "relative": "24h"
 *                         },
 *                         "displayCache": {
 *                             "name": "UTC Time",
 *                             "value": "Past 24 hours"
 *                         },
 *                         "filteredPartIds": [
 *                             "StartboardPart-UnboundPart-ae44fef5-76b8-46b0-86f0-2b3f47bad1c7"
 *                         ]
 *                     }
 *                 }
 *             }
 *         }
 *     }
 * }
 * `),
 * });
 * ```
 *
 * It is recommended to follow the steps outlined
 * [here](https://docs.microsoft.com/en-us/azure/azure-portal/azure-portal-dashboards-create-programmatically#fetch-the-json-representation-of-the-dashboard) to create a Dashboard in the Portal and extract the relevant JSON to use in this resource. From the extracted JSON, the contents of the `properties: {}` object can used. Variables can be injected as needed - see above example.
 *
 * ## Import
 *
 * Dashboards can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:dashboard/dashboard:Dashboard my-board /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Portal/dashboards/00000000-0000-0000-0000-000000000000
 * ```
 *
 *  Note the URI in the above sample can be found using the Resource Explorer tool in the Azure Portal.
 */
export class Dashboard extends pulumi.CustomResource {
    /**
     * Get an existing Dashboard resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DashboardState, opts?: pulumi.CustomResourceOptions): Dashboard {
        return new Dashboard(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:dashboard/dashboard:Dashboard';

    /**
     * Returns true if the given object is an instance of Dashboard.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Dashboard {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Dashboard.__pulumiType;
    }

    /**
     * JSON data representing dashboard body. See above for details on how to obtain this from the Portal.
     */
    public readonly dashboardProperties!: pulumi.Output<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Shared Dashboard. This should be be 64 chars max, only alphanumeric and hyphens (no spaces). For a more friendly display name, add the `hidden-title` tag.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to
     * create the dashboard.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Dashboard resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DashboardArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DashboardArgs | DashboardState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DashboardState | undefined;
            inputs["dashboardProperties"] = state ? state.dashboardProperties : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as DashboardArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["dashboardProperties"] = args ? args.dashboardProperties : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Dashboard.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Dashboard resources.
 */
export interface DashboardState {
    /**
     * JSON data representing dashboard body. See above for details on how to obtain this from the Portal.
     */
    dashboardProperties?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Shared Dashboard. This should be be 64 chars max, only alphanumeric and hyphens (no spaces). For a more friendly display name, add the `hidden-title` tag.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to
     * create the dashboard.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Dashboard resource.
 */
export interface DashboardArgs {
    /**
     * JSON data representing dashboard body. See above for details on how to obtain this from the Portal.
     */
    dashboardProperties?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Shared Dashboard. This should be be 64 chars max, only alphanumeric and hyphens (no spaces). For a more friendly display name, add the `hidden-title` tag.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to
     * create the dashboard.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

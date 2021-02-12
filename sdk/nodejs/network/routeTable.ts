// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Route Table
 *
 * > **NOTE on Route Tables and Routes:** There is both a standalone `route` resource, and allows for Routes to be defined in-line within the `routeTable` resource.
 * At this time you cannot use a Route Table with in-line Routes in conjunction with any Route resources. Doing so will cause a conflict of Route configurations and will overwrite Routes.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US"});
 * const exampleRouteTable = new azure.network.RouteTable("exampleRouteTable", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     disableBgpRoutePropagation: false,
 *     routes: [{
 *         name: "route1",
 *         addressPrefix: "10.1.0.0/16",
 *         nextHopType: "vnetlocal",
 *     }],
 *     tags: {
 *         environment: "Production",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Route Tables can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:network/routeTable:RouteTable example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/routeTables/mytable1
 * ```
 */
export class RouteTable extends pulumi.CustomResource {
    /**
     * Get an existing RouteTable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteTableState, opts?: pulumi.CustomResourceOptions): RouteTable {
        return new RouteTable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/routeTable:RouteTable';

    /**
     * Returns true if the given object is an instance of RouteTable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouteTable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouteTable.__pulumiType;
    }

    /**
     * Boolean flag which controls propagation of routes learned by BGP on that route table. True means disable.
     */
    public readonly disableBgpRoutePropagation!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the route.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the route table. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A list of objects representing routes. Each object accepts the arguments documented below.
     */
    public readonly routes!: pulumi.Output<outputs.network.RouteTableRoute[]>;
    /**
     * The collection of Subnets associated with this route table.
     */
    public /*out*/ readonly subnets!: pulumi.Output<string[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a RouteTable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteTableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteTableArgs | RouteTableState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouteTableState | undefined;
            inputs["disableBgpRoutePropagation"] = state ? state.disableBgpRoutePropagation : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["routes"] = state ? state.routes : undefined;
            inputs["subnets"] = state ? state.subnets : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as RouteTableArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["disableBgpRoutePropagation"] = args ? args.disableBgpRoutePropagation : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["routes"] = args ? args.routes : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["subnets"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RouteTable.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouteTable resources.
 */
export interface RouteTableState {
    /**
     * Boolean flag which controls propagation of routes learned by BGP on that route table. True means disable.
     */
    readonly disableBgpRoutePropagation?: pulumi.Input<boolean>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the route.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the route table. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A list of objects representing routes. Each object accepts the arguments documented below.
     */
    readonly routes?: pulumi.Input<pulumi.Input<inputs.network.RouteTableRoute>[]>;
    /**
     * The collection of Subnets associated with this route table.
     */
    readonly subnets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a RouteTable resource.
 */
export interface RouteTableArgs {
    /**
     * Boolean flag which controls propagation of routes learned by BGP on that route table. True means disable.
     */
    readonly disableBgpRoutePropagation?: pulumi.Input<boolean>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the route.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the route table. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A list of objects representing routes. Each object accepts the arguments documented below.
     */
    readonly routes?: pulumi.Input<pulumi.Input<inputs.network.RouteTableRoute>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

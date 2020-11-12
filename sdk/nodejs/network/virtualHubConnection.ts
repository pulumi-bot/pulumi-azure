// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Connection for a Virtual Hub.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("exampleVirtualNetwork", {
 *     addressSpaces: ["172.0.0.0/16"],
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const test = new azure.network.VirtualWan("test", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 * });
 * const exampleVirtualHub = new azure.network.VirtualHub("exampleVirtualHub", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     virtualWanId: azurerm_virtual_wan.example.id,
 *     addressPrefix: "10.0.1.0/24",
 * });
 * const exampleVirtualHubConnection = new azure.network.VirtualHubConnection("exampleVirtualHubConnection", {
 *     virtualHubId: exampleVirtualHub.id,
 *     remoteVirtualNetworkId: exampleVirtualNetwork.id,
 * });
 * ```
 *
 * ## Import
 *
 * Virtual Hub Connection's can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:network/virtualHubConnection:VirtualHubConnection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/hub1/hubVirtualNetworkConnections/connection1
 * ```
 */
export class VirtualHubConnection extends pulumi.CustomResource {
    /**
     * Get an existing VirtualHubConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualHubConnectionState, opts?: pulumi.CustomResourceOptions): VirtualHubConnection {
        return new VirtualHubConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/virtualHubConnection:VirtualHubConnection';

    /**
     * Returns true if the given object is an instance of VirtualHubConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualHubConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualHubConnection.__pulumiType;
    }

    /**
     * @deprecated Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
     */
    public readonly hubToVitualNetworkTrafficAllowed!: pulumi.Output<boolean | undefined>;
    /**
     * Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created. Defaults to `false`.
     */
    public readonly internetSecurityEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The Name which should be used for this Connection, which must be unique within the Virtual Hub. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the Virtual Network which the Virtual Hub should be connected to. Changing this forces a new resource to be created.
     */
    public readonly remoteVirtualNetworkId!: pulumi.Output<string>;
    /**
     * A `routing` block as defined below.
     */
    public readonly routing!: pulumi.Output<outputs.network.VirtualHubConnectionRouting>;
    /**
     * The ID of the Virtual Hub within which this connection should be created. Changing this forces a new resource to be created.
     */
    public readonly virtualHubId!: pulumi.Output<string>;
    /**
     * @deprecated Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
     */
    public readonly vitualNetworkToHubGatewaysTrafficAllowed!: pulumi.Output<boolean | undefined>;

    /**
     * Create a VirtualHubConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualHubConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualHubConnectionArgs | VirtualHubConnectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VirtualHubConnectionState | undefined;
            inputs["hubToVitualNetworkTrafficAllowed"] = state ? state.hubToVitualNetworkTrafficAllowed : undefined;
            inputs["internetSecurityEnabled"] = state ? state.internetSecurityEnabled : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["remoteVirtualNetworkId"] = state ? state.remoteVirtualNetworkId : undefined;
            inputs["routing"] = state ? state.routing : undefined;
            inputs["virtualHubId"] = state ? state.virtualHubId : undefined;
            inputs["vitualNetworkToHubGatewaysTrafficAllowed"] = state ? state.vitualNetworkToHubGatewaysTrafficAllowed : undefined;
        } else {
            const args = argsOrState as VirtualHubConnectionArgs | undefined;
            if (!args || args.remoteVirtualNetworkId === undefined) {
                throw new Error("Missing required property 'remoteVirtualNetworkId'");
            }
            if (!args || args.virtualHubId === undefined) {
                throw new Error("Missing required property 'virtualHubId'");
            }
            inputs["hubToVitualNetworkTrafficAllowed"] = args ? args.hubToVitualNetworkTrafficAllowed : undefined;
            inputs["internetSecurityEnabled"] = args ? args.internetSecurityEnabled : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["remoteVirtualNetworkId"] = args ? args.remoteVirtualNetworkId : undefined;
            inputs["routing"] = args ? args.routing : undefined;
            inputs["virtualHubId"] = args ? args.virtualHubId : undefined;
            inputs["vitualNetworkToHubGatewaysTrafficAllowed"] = args ? args.vitualNetworkToHubGatewaysTrafficAllowed : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VirtualHubConnection.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualHubConnection resources.
 */
export interface VirtualHubConnectionState {
    /**
     * @deprecated Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
     */
    readonly hubToVitualNetworkTrafficAllowed?: pulumi.Input<boolean>;
    /**
     * Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created. Defaults to `false`.
     */
    readonly internetSecurityEnabled?: pulumi.Input<boolean>;
    /**
     * The Name which should be used for this Connection, which must be unique within the Virtual Hub. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the Virtual Network which the Virtual Hub should be connected to. Changing this forces a new resource to be created.
     */
    readonly remoteVirtualNetworkId?: pulumi.Input<string>;
    /**
     * A `routing` block as defined below.
     */
    readonly routing?: pulumi.Input<inputs.network.VirtualHubConnectionRouting>;
    /**
     * The ID of the Virtual Hub within which this connection should be created. Changing this forces a new resource to be created.
     */
    readonly virtualHubId?: pulumi.Input<string>;
    /**
     * @deprecated Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
     */
    readonly vitualNetworkToHubGatewaysTrafficAllowed?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a VirtualHubConnection resource.
 */
export interface VirtualHubConnectionArgs {
    /**
     * @deprecated Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
     */
    readonly hubToVitualNetworkTrafficAllowed?: pulumi.Input<boolean>;
    /**
     * Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created. Defaults to `false`.
     */
    readonly internetSecurityEnabled?: pulumi.Input<boolean>;
    /**
     * The Name which should be used for this Connection, which must be unique within the Virtual Hub. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the Virtual Network which the Virtual Hub should be connected to. Changing this forces a new resource to be created.
     */
    readonly remoteVirtualNetworkId: pulumi.Input<string>;
    /**
     * A `routing` block as defined below.
     */
    readonly routing?: pulumi.Input<inputs.network.VirtualHubConnectionRouting>;
    /**
     * The ID of the Virtual Hub within which this connection should be created. Changing this forces a new resource to be created.
     */
    readonly virtualHubId: pulumi.Input<string>;
    /**
     * @deprecated Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
     */
    readonly vitualNetworkToHubGatewaysTrafficAllowed?: pulumi.Input<boolean>;
}

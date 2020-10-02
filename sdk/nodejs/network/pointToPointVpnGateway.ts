// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Manages a Point-to-Site VPN Gateway.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = new azure.network.PointToPointVpnGateway("example", {
 *     location: azurerm_resource_group.example.location,
 *     resourceGroupName: azurerm_resource_group.example.resource_group_name,
 *     virtualHubId: azurerm_virtual_hub.example.id,
 *     vpnServerConfigurationId: azurerm_vpn_server_configuration.example.id,
 *     scaleUnit: 1,
 * });
 * ```
 */
export class PointToPointVpnGateway extends pulumi.CustomResource {
    /**
     * Get an existing PointToPointVpnGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PointToPointVpnGatewayState, opts?: pulumi.CustomResourceOptions): PointToPointVpnGateway {
        return new PointToPointVpnGateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/pointToPointVpnGateway:PointToPointVpnGateway';

    /**
     * Returns true if the given object is an instance of PointToPointVpnGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PointToPointVpnGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PointToPointVpnGateway.__pulumiType;
    }

    /**
     * A `connectionConfiguration` block as defined below.
     */
    public readonly connectionConfiguration!: pulumi.Output<outputs.network.PointToPointVpnGatewayConnectionConfiguration>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The Scale Unit for this Point-to-Site VPN Gateway.
     */
    public readonly scaleUnit!: pulumi.Output<number>;
    /**
     * A mapping of tags to assign to the Point-to-Site VPN Gateway.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of the Virtual Hub where this Point-to-Site VPN Gateway should exist. Changing this forces a new resource to be created.
     */
    public readonly virtualHubId!: pulumi.Output<string>;
    /**
     * The ID of the VPN Server Configuration which this Point-to-Site VPN Gateway should use. Changing this forces a new resource to be created.
     */
    public readonly vpnServerConfigurationId!: pulumi.Output<string>;

    /**
     * Create a PointToPointVpnGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PointToPointVpnGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PointToPointVpnGatewayArgs | PointToPointVpnGatewayState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PointToPointVpnGatewayState | undefined;
            inputs["connectionConfiguration"] = state ? state.connectionConfiguration : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["scaleUnit"] = state ? state.scaleUnit : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["virtualHubId"] = state ? state.virtualHubId : undefined;
            inputs["vpnServerConfigurationId"] = state ? state.vpnServerConfigurationId : undefined;
        } else {
            const args = argsOrState as PointToPointVpnGatewayArgs | undefined;
            if (!args || args.connectionConfiguration === undefined) {
                throw new Error("Missing required property 'connectionConfiguration'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.scaleUnit === undefined) {
                throw new Error("Missing required property 'scaleUnit'");
            }
            if (!args || args.virtualHubId === undefined) {
                throw new Error("Missing required property 'virtualHubId'");
            }
            if (!args || args.vpnServerConfigurationId === undefined) {
                throw new Error("Missing required property 'vpnServerConfigurationId'");
            }
            inputs["connectionConfiguration"] = args ? args.connectionConfiguration : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["scaleUnit"] = args ? args.scaleUnit : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualHubId"] = args ? args.virtualHubId : undefined;
            inputs["vpnServerConfigurationId"] = args ? args.vpnServerConfigurationId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(PointToPointVpnGateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PointToPointVpnGateway resources.
 */
export interface PointToPointVpnGatewayState {
    /**
     * A `connectionConfiguration` block as defined below.
     */
    readonly connectionConfiguration?: pulumi.Input<inputs.network.PointToPointVpnGatewayConnectionConfiguration>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The Scale Unit for this Point-to-Site VPN Gateway.
     */
    readonly scaleUnit?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the Point-to-Site VPN Gateway.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the Virtual Hub where this Point-to-Site VPN Gateway should exist. Changing this forces a new resource to be created.
     */
    readonly virtualHubId?: pulumi.Input<string>;
    /**
     * The ID of the VPN Server Configuration which this Point-to-Site VPN Gateway should use. Changing this forces a new resource to be created.
     */
    readonly vpnServerConfigurationId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PointToPointVpnGateway resource.
 */
export interface PointToPointVpnGatewayArgs {
    /**
     * A `connectionConfiguration` block as defined below.
     */
    readonly connectionConfiguration: pulumi.Input<inputs.network.PointToPointVpnGatewayConnectionConfiguration>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The Scale Unit for this Point-to-Site VPN Gateway.
     */
    readonly scaleUnit: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the Point-to-Site VPN Gateway.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the Virtual Hub where this Point-to-Site VPN Gateway should exist. Changing this forces a new resource to be created.
     */
    readonly virtualHubId: pulumi.Input<string>;
    /**
     * The ID of the VPN Server Configuration which this Point-to-Site VPN Gateway should use. Changing this forces a new resource to be created.
     */
    readonly vpnServerConfigurationId: pulumi.Input<string>;
}

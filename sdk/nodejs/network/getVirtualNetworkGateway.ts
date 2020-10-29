// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Virtual Network Gateway.
 */
export function getVirtualNetworkGateway(args: GetVirtualNetworkGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualNetworkGatewayResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:network/getVirtualNetworkGateway:getVirtualNetworkGateway", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getVirtualNetworkGateway.
 */
export interface GetVirtualNetworkGatewayArgs {
    /**
     * Specifies the name of the Virtual Network Gateway.
     */
    readonly name: string;
    /**
     * Specifies the name of the resource group the Virtual Network Gateway is located in.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getVirtualNetworkGateway.
 */
export interface GetVirtualNetworkGatewayResult {
    /**
     * Is this an Active-Active Gateway?
     */
    readonly activeActive: boolean;
    readonly bgpSettings: outputs.network.GetVirtualNetworkGatewayBgpSetting[];
    /**
     * The ID of the local network gateway
     * through which outbound Internet traffic from the virtual network in which the
     * gateway is created will be routed (*forced tunneling*). Refer to the
     * [Azure documentation on forced tunneling](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-forced-tunneling-rm).
     */
    readonly defaultLocalNetworkGatewayId: string;
    /**
     * Will BGP (Border Gateway Protocol) will be enabled
     * for this Virtual Network Gateway.
     */
    readonly enableBgp: boolean;
    /**
     * The Generation of the Virtual Network Gateway.
     */
    readonly generation: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * One or two `ipConfiguration` blocks documented below.
     */
    readonly ipConfigurations: outputs.network.GetVirtualNetworkGatewayIpConfiguration[];
    /**
     * The location/region where the Virtual Network Gateway is located.
     */
    readonly location: string;
    /**
     * The user-defined name of the revoked certificate.
     */
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * Configuration of the size and capacity of the Virtual Network Gateway.
     */
    readonly sku: string;
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags: {[key: string]: string};
    /**
     * The type of the Virtual Network Gateway.
     */
    readonly type: string;
    /**
     * A `vpnClientConfiguration` block which is documented below.
     */
    readonly vpnClientConfigurations: outputs.network.GetVirtualNetworkGatewayVpnClientConfiguration[];
    /**
     * The routing type of the Virtual Network Gateway.
     */
    readonly vpnType: string;
}

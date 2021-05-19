// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Virtual Network Gateway to establish secure, cross-premises connectivity.
 *
 * > **Note:** Please be aware that provisioning a Virtual Network Gateway takes a long time (between 30 minutes and 1 hour)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("exampleVirtualNetwork", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     addressSpaces: ["10.0.0.0/16"],
 * });
 * const exampleSubnet = new azure.network.Subnet("exampleSubnet", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefixes: ["10.0.1.0/24"],
 * });
 * const examplePublicIp = new azure.network.PublicIp("examplePublicIp", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     allocationMethod: "Dynamic",
 * });
 * const exampleVirtualNetworkGateway = new azure.network.VirtualNetworkGateway("exampleVirtualNetworkGateway", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     type: "Vpn",
 *     vpnType: "RouteBased",
 *     activeActive: false,
 *     enableBgp: false,
 *     sku: "Basic",
 *     ipConfigurations: [{
 *         name: "vnetGatewayConfig",
 *         publicIpAddressId: examplePublicIp.id,
 *         privateIpAddressAllocation: "Dynamic",
 *         subnetId: exampleSubnet.id,
 *     }],
 *     vpnClientConfiguration: {
 *         addressSpaces: ["10.2.0.0/24"],
 *         rootCertificates: [{
 *             name: "DigiCert-Federated-ID-Root-CA",
 *             publicCertData: `MIIDuzCCAqOgAwIBAgIQCHTZWCM+IlfFIRXIvyKSrjANBgkqhkiG9w0BAQsFADBn
 * MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
 * d3cuZGlnaWNlcnQuY29tMSYwJAYDVQQDEx1EaWdpQ2VydCBGZWRlcmF0ZWQgSUQg
 * Um9vdCBDQTAeFw0xMzAxMTUxMjAwMDBaFw0zMzAxMTUxMjAwMDBaMGcxCzAJBgNV
 * BAYTAlVTMRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdp
 * Y2VydC5jb20xJjAkBgNVBAMTHURpZ2lDZXJ0IEZlZGVyYXRlZCBJRCBSb290IENB
 * MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvAEB4pcCqnNNOWE6Ur5j
 * QPUH+1y1F9KdHTRSza6k5iDlXq1kGS1qAkuKtw9JsiNRrjltmFnzMZRBbX8Tlfl8
 * zAhBmb6dDduDGED01kBsTkgywYPxXVTKec0WxYEEF0oMn4wSYNl0lt2eJAKHXjNf
 * GTwiibdP8CUR2ghSM2sUTI8Nt1Omfc4SMHhGhYD64uJMbX98THQ/4LMGuYegou+d
 * GTiahfHtjn7AboSEknwAMJHCh5RlYZZ6B1O4QbKJ+34Q0eKgnI3X6Vc9u0zf6DH8
 * Dk+4zQDYRRTqTnVO3VT8jzqDlCRuNtq6YvryOWN74/dq8LQhUnXHvFyrsdMaE1X2
 * DwIDAQABo2MwYTAPBgNVHRMBAf8EBTADAQH/MA4GA1UdDwEB/wQEAwIBhjAdBgNV
 * HQ4EFgQUGRdkFnbGt1EWjKwbUne+5OaZvRYwHwYDVR0jBBgwFoAUGRdkFnbGt1EW
 * jKwbUne+5OaZvRYwDQYJKoZIhvcNAQELBQADggEBAHcqsHkrjpESqfuVTRiptJfP
 * 9JbdtWqRTmOf6uJi2c8YVqI6XlKXsD8C1dUUaaHKLUJzvKiazibVuBwMIT84AyqR
 * QELn3e0BtgEymEygMU569b01ZPxoFSnNXc7qDZBDef8WfqAV/sxkTi8L9BkmFYfL
 * uGLOhRJOFprPdoDIUBB+tmCl3oDcBy3vnUeOEioz8zAkprcb3GHwHAK+vHmmfgcn
 * WsfMLH4JCLa/tRYL+Rw/N3ybCkDp00s0WUZ+AoDywSl0Q/ZEnNY0MsFiw6LyIdbq
 * M/s/1JRtO3bDSzD9TazRVzn2oBqzSa8VgIo5C1nOnoAKJTlsClJKvIhnRlaLQqk=
 * `,
 *         }],
 *         revokedCertificates: [{
 *             name: "Verizon-Global-Root-CA",
 *             thumbprint: "912198EEF23DCAC40939312FEE97DD560BAE49B1",
 *         }],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Virtual Network Gateways can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:network/virtualNetworkGateway:VirtualNetworkGateway exampleGateway /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.Network/virtualNetworkGateways/myGateway1
 * ```
 */
export class VirtualNetworkGateway extends pulumi.CustomResource {
    /**
     * Get an existing VirtualNetworkGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualNetworkGatewayState, opts?: pulumi.CustomResourceOptions): VirtualNetworkGateway {
        return new VirtualNetworkGateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/virtualNetworkGateway:VirtualNetworkGateway';

    /**
     * Returns true if the given object is an instance of VirtualNetworkGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualNetworkGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualNetworkGateway.__pulumiType;
    }

    /**
     * If `true`, an active-active Virtual Network Gateway
     * will be created. An active-active gateway requires a `HighPerformance` or an
     * `UltraPerformance` sku. If `false`, an active-standby gateway will be created.
     * Defaults to `false`.
     */
    public readonly activeActive!: pulumi.Output<boolean>;
    /**
     * A block of `bgpSettings`.
     */
    public readonly bgpSettings!: pulumi.Output<outputs.network.VirtualNetworkGatewayBgpSettings>;
    public readonly customRoute!: pulumi.Output<outputs.network.VirtualNetworkGatewayCustomRoute | undefined>;
    /**
     * The ID of the local network gateway
     * through which outbound Internet traffic from the virtual network in which the
     * gateway is created will be routed (*forced tunnelling*). Refer to the
     * [Azure documentation on forced tunnelling](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-forced-tunneling-rm).
     * If not specified, forced tunnelling is disabled.
     */
    public readonly defaultLocalNetworkGatewayId!: pulumi.Output<string | undefined>;
    /**
     * If `true`, BGP (Border Gateway Protocol) will be enabled
     * for this Virtual Network Gateway. Defaults to `false`.
     */
    public readonly enableBgp!: pulumi.Output<boolean>;
    /**
     * The Generation of the Virtual Network gateway. Possible values include `Generation1`, `Generation2` or `None`.
     */
    public readonly generation!: pulumi.Output<string>;
    /**
     * One or two `ipConfiguration` blocks documented below.
     * An active-standby gateway requires exactly one `ipConfiguration` block whereas
     * an active-active gateway requires exactly two `ipConfiguration` blocks.
     */
    public readonly ipConfigurations!: pulumi.Output<outputs.network.VirtualNetworkGatewayIpConfiguration[]>;
    /**
     * The location/region where the Virtual Network Gateway is
     * located. Changing the location/region forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * A user-defined name of the IP configuration. Defaults to
     * `vnetGatewayConfig`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Should private IP be enabled on this gateway for connections? Changing this forces a new resource to be created.
     */
    public readonly privateIpAddressEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the resource group in which to
     * create the Virtual Network Gateway. Changing the resource group name forces
     * a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Configuration of the size and capacity of the virtual network
     * gateway. Valid options are `Basic`, `Standard`, `HighPerformance`, `UltraPerformance`,
     * `ErGw1AZ`, `ErGw2AZ`, `ErGw3AZ`, `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw4`,`VpnGw5`, `VpnGw1AZ`,
     * `VpnGw2AZ`, `VpnGw3AZ`,`VpnGw4AZ` and `VpnGw5AZ` and depend on the `type`, `vpnType` and
     * `generation` arguments.
     * A `PolicyBased` gateway only supports the `Basic` sku. Further, the `UltraPerformance`
     * sku is only supported by an `ExpressRoute` gateway.
     */
    public readonly sku!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the Virtual Network Gateway. Valid options are
     * `Vpn` or `ExpressRoute`. Changing the type forces a new resource to be created.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * A `vpnClientConfiguration` block which
     * is documented below. In this block the Virtual Network Gateway can be configured
     * to accept IPSec point-to-site connections.
     */
    public readonly vpnClientConfiguration!: pulumi.Output<outputs.network.VirtualNetworkGatewayVpnClientConfiguration | undefined>;
    /**
     * The routing type of the Virtual Network Gateway. Valid
     * options are `RouteBased` or `PolicyBased`. Defaults to `RouteBased`.
     */
    public readonly vpnType!: pulumi.Output<string | undefined>;

    /**
     * Create a VirtualNetworkGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualNetworkGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualNetworkGatewayArgs | VirtualNetworkGatewayState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VirtualNetworkGatewayState | undefined;
            inputs["activeActive"] = state ? state.activeActive : undefined;
            inputs["bgpSettings"] = state ? state.bgpSettings : undefined;
            inputs["customRoute"] = state ? state.customRoute : undefined;
            inputs["defaultLocalNetworkGatewayId"] = state ? state.defaultLocalNetworkGatewayId : undefined;
            inputs["enableBgp"] = state ? state.enableBgp : undefined;
            inputs["generation"] = state ? state.generation : undefined;
            inputs["ipConfigurations"] = state ? state.ipConfigurations : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["privateIpAddressEnabled"] = state ? state.privateIpAddressEnabled : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["vpnClientConfiguration"] = state ? state.vpnClientConfiguration : undefined;
            inputs["vpnType"] = state ? state.vpnType : undefined;
        } else {
            const args = argsOrState as VirtualNetworkGatewayArgs | undefined;
            if ((!args || args.ipConfigurations === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipConfigurations'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.sku === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sku'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            inputs["activeActive"] = args ? args.activeActive : undefined;
            inputs["bgpSettings"] = args ? args.bgpSettings : undefined;
            inputs["customRoute"] = args ? args.customRoute : undefined;
            inputs["defaultLocalNetworkGatewayId"] = args ? args.defaultLocalNetworkGatewayId : undefined;
            inputs["enableBgp"] = args ? args.enableBgp : undefined;
            inputs["generation"] = args ? args.generation : undefined;
            inputs["ipConfigurations"] = args ? args.ipConfigurations : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["privateIpAddressEnabled"] = args ? args.privateIpAddressEnabled : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["vpnClientConfiguration"] = args ? args.vpnClientConfiguration : undefined;
            inputs["vpnType"] = args ? args.vpnType : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(VirtualNetworkGateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualNetworkGateway resources.
 */
export interface VirtualNetworkGatewayState {
    /**
     * If `true`, an active-active Virtual Network Gateway
     * will be created. An active-active gateway requires a `HighPerformance` or an
     * `UltraPerformance` sku. If `false`, an active-standby gateway will be created.
     * Defaults to `false`.
     */
    activeActive?: pulumi.Input<boolean>;
    /**
     * A block of `bgpSettings`.
     */
    bgpSettings?: pulumi.Input<inputs.network.VirtualNetworkGatewayBgpSettings>;
    customRoute?: pulumi.Input<inputs.network.VirtualNetworkGatewayCustomRoute>;
    /**
     * The ID of the local network gateway
     * through which outbound Internet traffic from the virtual network in which the
     * gateway is created will be routed (*forced tunnelling*). Refer to the
     * [Azure documentation on forced tunnelling](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-forced-tunneling-rm).
     * If not specified, forced tunnelling is disabled.
     */
    defaultLocalNetworkGatewayId?: pulumi.Input<string>;
    /**
     * If `true`, BGP (Border Gateway Protocol) will be enabled
     * for this Virtual Network Gateway. Defaults to `false`.
     */
    enableBgp?: pulumi.Input<boolean>;
    /**
     * The Generation of the Virtual Network gateway. Possible values include `Generation1`, `Generation2` or `None`.
     */
    generation?: pulumi.Input<string>;
    /**
     * One or two `ipConfiguration` blocks documented below.
     * An active-standby gateway requires exactly one `ipConfiguration` block whereas
     * an active-active gateway requires exactly two `ipConfiguration` blocks.
     */
    ipConfigurations?: pulumi.Input<pulumi.Input<inputs.network.VirtualNetworkGatewayIpConfiguration>[]>;
    /**
     * The location/region where the Virtual Network Gateway is
     * located. Changing the location/region forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * A user-defined name of the IP configuration. Defaults to
     * `vnetGatewayConfig`.
     */
    name?: pulumi.Input<string>;
    /**
     * Should private IP be enabled on this gateway for connections? Changing this forces a new resource to be created.
     */
    privateIpAddressEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which to
     * create the Virtual Network Gateway. Changing the resource group name forces
     * a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * Configuration of the size and capacity of the virtual network
     * gateway. Valid options are `Basic`, `Standard`, `HighPerformance`, `UltraPerformance`,
     * `ErGw1AZ`, `ErGw2AZ`, `ErGw3AZ`, `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw4`,`VpnGw5`, `VpnGw1AZ`,
     * `VpnGw2AZ`, `VpnGw3AZ`,`VpnGw4AZ` and `VpnGw5AZ` and depend on the `type`, `vpnType` and
     * `generation` arguments.
     * A `PolicyBased` gateway only supports the `Basic` sku. Further, the `UltraPerformance`
     * sku is only supported by an `ExpressRoute` gateway.
     */
    sku?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of the Virtual Network Gateway. Valid options are
     * `Vpn` or `ExpressRoute`. Changing the type forces a new resource to be created.
     */
    type?: pulumi.Input<string>;
    /**
     * A `vpnClientConfiguration` block which
     * is documented below. In this block the Virtual Network Gateway can be configured
     * to accept IPSec point-to-site connections.
     */
    vpnClientConfiguration?: pulumi.Input<inputs.network.VirtualNetworkGatewayVpnClientConfiguration>;
    /**
     * The routing type of the Virtual Network Gateway. Valid
     * options are `RouteBased` or `PolicyBased`. Defaults to `RouteBased`.
     */
    vpnType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VirtualNetworkGateway resource.
 */
export interface VirtualNetworkGatewayArgs {
    /**
     * If `true`, an active-active Virtual Network Gateway
     * will be created. An active-active gateway requires a `HighPerformance` or an
     * `UltraPerformance` sku. If `false`, an active-standby gateway will be created.
     * Defaults to `false`.
     */
    activeActive?: pulumi.Input<boolean>;
    /**
     * A block of `bgpSettings`.
     */
    bgpSettings?: pulumi.Input<inputs.network.VirtualNetworkGatewayBgpSettings>;
    customRoute?: pulumi.Input<inputs.network.VirtualNetworkGatewayCustomRoute>;
    /**
     * The ID of the local network gateway
     * through which outbound Internet traffic from the virtual network in which the
     * gateway is created will be routed (*forced tunnelling*). Refer to the
     * [Azure documentation on forced tunnelling](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-forced-tunneling-rm).
     * If not specified, forced tunnelling is disabled.
     */
    defaultLocalNetworkGatewayId?: pulumi.Input<string>;
    /**
     * If `true`, BGP (Border Gateway Protocol) will be enabled
     * for this Virtual Network Gateway. Defaults to `false`.
     */
    enableBgp?: pulumi.Input<boolean>;
    /**
     * The Generation of the Virtual Network gateway. Possible values include `Generation1`, `Generation2` or `None`.
     */
    generation?: pulumi.Input<string>;
    /**
     * One or two `ipConfiguration` blocks documented below.
     * An active-standby gateway requires exactly one `ipConfiguration` block whereas
     * an active-active gateway requires exactly two `ipConfiguration` blocks.
     */
    ipConfigurations: pulumi.Input<pulumi.Input<inputs.network.VirtualNetworkGatewayIpConfiguration>[]>;
    /**
     * The location/region where the Virtual Network Gateway is
     * located. Changing the location/region forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * A user-defined name of the IP configuration. Defaults to
     * `vnetGatewayConfig`.
     */
    name?: pulumi.Input<string>;
    /**
     * Should private IP be enabled on this gateway for connections? Changing this forces a new resource to be created.
     */
    privateIpAddressEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which to
     * create the Virtual Network Gateway. Changing the resource group name forces
     * a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Configuration of the size and capacity of the virtual network
     * gateway. Valid options are `Basic`, `Standard`, `HighPerformance`, `UltraPerformance`,
     * `ErGw1AZ`, `ErGw2AZ`, `ErGw3AZ`, `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw4`,`VpnGw5`, `VpnGw1AZ`,
     * `VpnGw2AZ`, `VpnGw3AZ`,`VpnGw4AZ` and `VpnGw5AZ` and depend on the `type`, `vpnType` and
     * `generation` arguments.
     * A `PolicyBased` gateway only supports the `Basic` sku. Further, the `UltraPerformance`
     * sku is only supported by an `ExpressRoute` gateway.
     */
    sku: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of the Virtual Network Gateway. Valid options are
     * `Vpn` or `ExpressRoute`. Changing the type forces a new resource to be created.
     */
    type: pulumi.Input<string>;
    /**
     * A `vpnClientConfiguration` block which
     * is documented below. In this block the Virtual Network Gateway can be configured
     * to accept IPSec point-to-site connections.
     */
    vpnClientConfiguration?: pulumi.Input<inputs.network.VirtualNetworkGatewayVpnClientConfiguration>;
    /**
     * The routing type of the Virtual Network Gateway. Valid
     * options are `RouteBased` or `PolicyBased`. Defaults to `RouteBased`.
     */
    vpnType?: pulumi.Input<string>;
}

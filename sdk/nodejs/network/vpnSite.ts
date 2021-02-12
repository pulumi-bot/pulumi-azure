// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a VPN Site.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleVirtualWan = new azure.network.VirtualWan("exampleVirtualWan", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 * });
 * const exampleVpnSite = new azure.network.VpnSite("exampleVpnSite", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     virtualWanId: exampleVirtualWan.id,
 *     links: [{
 *         name: "link1",
 *         ipAddress: "10.0.0.1",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * VPN Sites can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:network/vpnSite:VpnSite example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/vpnSites/site1
 * ```
 */
export class VpnSite extends pulumi.CustomResource {
    /**
     * Get an existing VpnSite resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpnSiteState, opts?: pulumi.CustomResourceOptions): VpnSite {
        return new VpnSite(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/vpnSite:VpnSite';

    /**
     * Returns true if the given object is an instance of VpnSite.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpnSite {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpnSite.__pulumiType;
    }

    /**
     * Specifies a list of IP address CIDRs that are located on your on-premises site. Traffic destined for these address spaces is routed to your local site.
     */
    public readonly addressCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * The model of the VPN device.
     */
    public readonly deviceModel!: pulumi.Output<string | undefined>;
    /**
     * The name of the VPN device vendor.
     */
    public readonly deviceVendor!: pulumi.Output<string | undefined>;
    /**
     * One or more `link` blocks as defined below.
     */
    public readonly links!: pulumi.Output<outputs.network.VpnSiteLink[] | undefined>;
    /**
     * The Azure Region where the VPN Site should exist. Changing this forces a new VPN Site to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name which should be used for this VPN Site. Changing this forces a new VPN Site to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Resource Group where the VPN Site should exist. Changing this forces a new VPN Site to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags which should be assigned to the VPN Site.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of the Virtual Wan where this VPN site resides in. Changing this forces a new VPN Site to be created.
     */
    public readonly virtualWanId!: pulumi.Output<string>;

    /**
     * Create a VpnSite resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpnSiteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpnSiteArgs | VpnSiteState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpnSiteState | undefined;
            inputs["addressCidrs"] = state ? state.addressCidrs : undefined;
            inputs["deviceModel"] = state ? state.deviceModel : undefined;
            inputs["deviceVendor"] = state ? state.deviceVendor : undefined;
            inputs["links"] = state ? state.links : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["virtualWanId"] = state ? state.virtualWanId : undefined;
        } else {
            const args = argsOrState as VpnSiteArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.virtualWanId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'virtualWanId'");
            }
            inputs["addressCidrs"] = args ? args.addressCidrs : undefined;
            inputs["deviceModel"] = args ? args.deviceModel : undefined;
            inputs["deviceVendor"] = args ? args.deviceVendor : undefined;
            inputs["links"] = args ? args.links : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualWanId"] = args ? args.virtualWanId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(VpnSite.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpnSite resources.
 */
export interface VpnSiteState {
    /**
     * Specifies a list of IP address CIDRs that are located on your on-premises site. Traffic destined for these address spaces is routed to your local site.
     */
    readonly addressCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The model of the VPN device.
     */
    readonly deviceModel?: pulumi.Input<string>;
    /**
     * The name of the VPN device vendor.
     */
    readonly deviceVendor?: pulumi.Input<string>;
    /**
     * One or more `link` blocks as defined below.
     */
    readonly links?: pulumi.Input<pulumi.Input<inputs.network.VpnSiteLink>[]>;
    /**
     * The Azure Region where the VPN Site should exist. Changing this forces a new VPN Site to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name which should be used for this VPN Site. Changing this forces a new VPN Site to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the VPN Site should exist. Changing this forces a new VPN Site to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the VPN Site.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the Virtual Wan where this VPN site resides in. Changing this forces a new VPN Site to be created.
     */
    readonly virtualWanId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpnSite resource.
 */
export interface VpnSiteArgs {
    /**
     * Specifies a list of IP address CIDRs that are located on your on-premises site. Traffic destined for these address spaces is routed to your local site.
     */
    readonly addressCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The model of the VPN device.
     */
    readonly deviceModel?: pulumi.Input<string>;
    /**
     * The name of the VPN device vendor.
     */
    readonly deviceVendor?: pulumi.Input<string>;
    /**
     * One or more `link` blocks as defined below.
     */
    readonly links?: pulumi.Input<pulumi.Input<inputs.network.VpnSiteLink>[]>;
    /**
     * The Azure Region where the VPN Site should exist. Changing this forces a new VPN Site to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name which should be used for this VPN Site. Changing this forces a new VPN Site to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the VPN Site should exist. Changing this forces a new VPN Site to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the VPN Site.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the Virtual Wan where this VPN site resides in. Changing this forces a new VPN Site to be created.
     */
    readonly virtualWanId: pulumi.Input<string>;
}

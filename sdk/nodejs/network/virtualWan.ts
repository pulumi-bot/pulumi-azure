// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Virtual WAN.
 *
 * ## Example Usage
 *
 *
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
 * ```
 */
export class VirtualWan extends pulumi.CustomResource {
    /**
     * Get an existing VirtualWan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualWanState, opts?: pulumi.CustomResourceOptions): VirtualWan {
        return new VirtualWan(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/virtualWan:VirtualWan';

    /**
     * Returns true if the given object is an instance of VirtualWan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualWan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualWan.__pulumiType;
    }

    /**
     * Boolean flag to specify whether branch to branch traffic is allowed. Defaults to `true`.
     */
    public readonly allowBranchToBranchTraffic!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean flag to specify whether VNet to VNet traffic is allowed. Defaults to `false`.
     */
    public readonly allowVnetToVnetTraffic!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean flag to specify whether VPN encryption is disabled. Defaults to `false`.
     */
    public readonly disableVpnEncryption!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Virtual WAN. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the Office365 local breakout category. Possible values include: `Optimize`, `OptimizeAndAllow`, `All`, `None`. Defaults to `None`.
     */
    public readonly office365LocalBreakoutCategory!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource group in which to create the Virtual WAN. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the Virtual WAN.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the Virtual WAN type. Possible Values include: `Basic` and `Standard`. Defaults to `Standard`.
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a VirtualWan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualWanArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualWanArgs | VirtualWanState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VirtualWanState | undefined;
            inputs["allowBranchToBranchTraffic"] = state ? state.allowBranchToBranchTraffic : undefined;
            inputs["allowVnetToVnetTraffic"] = state ? state.allowVnetToVnetTraffic : undefined;
            inputs["disableVpnEncryption"] = state ? state.disableVpnEncryption : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["office365LocalBreakoutCategory"] = state ? state.office365LocalBreakoutCategory : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as VirtualWanArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["allowBranchToBranchTraffic"] = args ? args.allowBranchToBranchTraffic : undefined;
            inputs["allowVnetToVnetTraffic"] = args ? args.allowVnetToVnetTraffic : undefined;
            inputs["disableVpnEncryption"] = args ? args.disableVpnEncryption : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["office365LocalBreakoutCategory"] = args ? args.office365LocalBreakoutCategory : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VirtualWan.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualWan resources.
 */
export interface VirtualWanState {
    /**
     * Boolean flag to specify whether branch to branch traffic is allowed. Defaults to `true`.
     */
    readonly allowBranchToBranchTraffic?: pulumi.Input<boolean>;
    /**
     * Boolean flag to specify whether VNet to VNet traffic is allowed. Defaults to `false`.
     */
    readonly allowVnetToVnetTraffic?: pulumi.Input<boolean>;
    /**
     * Boolean flag to specify whether VPN encryption is disabled. Defaults to `false`.
     */
    readonly disableVpnEncryption?: pulumi.Input<boolean>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Virtual WAN. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the Office365 local breakout category. Possible values include: `Optimize`, `OptimizeAndAllow`, `All`, `None`. Defaults to `None`.
     */
    readonly office365LocalBreakoutCategory?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Virtual WAN. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the Virtual WAN.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the Virtual WAN type. Possible Values include: `Basic` and `Standard`. Defaults to `Standard`.
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VirtualWan resource.
 */
export interface VirtualWanArgs {
    /**
     * Boolean flag to specify whether branch to branch traffic is allowed. Defaults to `true`.
     */
    readonly allowBranchToBranchTraffic?: pulumi.Input<boolean>;
    /**
     * Boolean flag to specify whether VNet to VNet traffic is allowed. Defaults to `false`.
     */
    readonly allowVnetToVnetTraffic?: pulumi.Input<boolean>;
    /**
     * Boolean flag to specify whether VPN encryption is disabled. Defaults to `false`.
     */
    readonly disableVpnEncryption?: pulumi.Input<boolean>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Virtual WAN. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the Office365 local breakout category. Possible values include: `Optimize`, `OptimizeAndAllow`, `All`, `None`. Defaults to `None`.
     */
    readonly office365LocalBreakoutCategory?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Virtual WAN. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the Virtual WAN.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the Virtual WAN type. Possible Values include: `Basic` and `Standard`. Defaults to `Standard`.
     */
    readonly type?: pulumi.Input<string>;
}

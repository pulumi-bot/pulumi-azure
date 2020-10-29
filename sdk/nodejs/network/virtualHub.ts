// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Virtual Hub within a Virtual WAN.
 */
export class VirtualHub extends pulumi.CustomResource {
    /**
     * Get an existing VirtualHub resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualHubState, opts?: pulumi.CustomResourceOptions): VirtualHub {
        return new VirtualHub(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/virtualHub:VirtualHub';

    /**
     * Returns true if the given object is an instance of VirtualHub.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualHub {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualHub.__pulumiType;
    }

    /**
     * The Address Prefix which should be used for this Virtual Hub. Changing this forces a new resource to be created.
     */
    public readonly addressPrefix!: pulumi.Output<string>;
    /**
     * Specifies the supported Azure location where the Virtual Hub should exist. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the Virtual Hub. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the name of the Resource Group where the Virtual Hub should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * One or more `route` blocks as defined below.
     */
    public readonly routes!: pulumi.Output<outputs.network.VirtualHubRoute[] | undefined>;
    /**
     * A mapping of tags to assign to the Virtual Hub.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of a Virtual WAN within which the Virtual Hub should be created. Changing this forces a new resource to be created.
     */
    public readonly virtualWanId!: pulumi.Output<string>;

    /**
     * Create a VirtualHub resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualHubArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualHubArgs | VirtualHubState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VirtualHubState | undefined;
            inputs["addressPrefix"] = state ? state.addressPrefix : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["routes"] = state ? state.routes : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["virtualWanId"] = state ? state.virtualWanId : undefined;
        } else {
            const args = argsOrState as VirtualHubArgs | undefined;
            if (!args || args.addressPrefix === undefined) {
                throw new Error("Missing required property 'addressPrefix'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.virtualWanId === undefined) {
                throw new Error("Missing required property 'virtualWanId'");
            }
            inputs["addressPrefix"] = args ? args.addressPrefix : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["routes"] = args ? args.routes : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualWanId"] = args ? args.virtualWanId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VirtualHub.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualHub resources.
 */
export interface VirtualHubState {
    /**
     * The Address Prefix which should be used for this Virtual Hub. Changing this forces a new resource to be created.
     */
    readonly addressPrefix?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the Virtual Hub should exist. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the Virtual Hub. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the name of the Resource Group where the Virtual Hub should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * One or more `route` blocks as defined below.
     */
    readonly routes?: pulumi.Input<pulumi.Input<inputs.network.VirtualHubRoute>[]>;
    /**
     * A mapping of tags to assign to the Virtual Hub.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of a Virtual WAN within which the Virtual Hub should be created. Changing this forces a new resource to be created.
     */
    readonly virtualWanId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VirtualHub resource.
 */
export interface VirtualHubArgs {
    /**
     * The Address Prefix which should be used for this Virtual Hub. Changing this forces a new resource to be created.
     */
    readonly addressPrefix: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the Virtual Hub should exist. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the Virtual Hub. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the name of the Resource Group where the Virtual Hub should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * One or more `route` blocks as defined below.
     */
    readonly routes?: pulumi.Input<pulumi.Input<inputs.network.VirtualHubRoute>[]>;
    /**
     * A mapping of tags to assign to the Virtual Hub.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of a Virtual WAN within which the Virtual Hub should be created. Changing this forces a new resource to be created.
     */
    readonly virtualWanId: pulumi.Input<string>;
}

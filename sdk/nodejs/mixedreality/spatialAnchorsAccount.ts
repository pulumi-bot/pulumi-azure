// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Azure Spatial Anchors Account.
 */
export class SpatialAnchorsAccount extends pulumi.CustomResource {
    /**
     * Get an existing SpatialAnchorsAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SpatialAnchorsAccountState, opts?: pulumi.CustomResourceOptions): SpatialAnchorsAccount {
        return new SpatialAnchorsAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:mixedreality/spatialAnchorsAccount:SpatialAnchorsAccount';

    /**
     * Returns true if the given object is an instance of SpatialAnchorsAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SpatialAnchorsAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SpatialAnchorsAccount.__pulumiType;
    }

    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Spatial Anchors Account. Changing this forces a new resource to be created. Must be globally unique.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Spatial Anchors Account.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a SpatialAnchorsAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SpatialAnchorsAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SpatialAnchorsAccountArgs | SpatialAnchorsAccountState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SpatialAnchorsAccountState | undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as SpatialAnchorsAccountArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SpatialAnchorsAccount.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SpatialAnchorsAccount resources.
 */
export interface SpatialAnchorsAccountState {
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Spatial Anchors Account. Changing this forces a new resource to be created. Must be globally unique.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Spatial Anchors Account.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a SpatialAnchorsAccount resource.
 */
export interface SpatialAnchorsAccountArgs {
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Spatial Anchors Account. Changing this forces a new resource to be created. Must be globally unique.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Spatial Anchors Account.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

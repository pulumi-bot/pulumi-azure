// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Availability Set for Virtual Machines.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAvailabilitySet = new azure.compute.AvailabilitySet("exampleAvailabilitySet", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     tags: {
 *         environment: "Production",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Availability Sets can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:compute/availabilitySet:AvailabilitySet group1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/availabilitySets/webAvailSet
 * ```
 */
export class AvailabilitySet extends pulumi.CustomResource {
    /**
     * Get an existing AvailabilitySet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AvailabilitySetState, opts?: pulumi.CustomResourceOptions): AvailabilitySet {
        return new AvailabilitySet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:compute/availabilitySet:AvailabilitySet';

    /**
     * Returns true if the given object is an instance of AvailabilitySet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AvailabilitySet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AvailabilitySet.__pulumiType;
    }

    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
     */
    public readonly managed!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the name of the availability set. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the number of fault domains that are used. Defaults to `3`. Changing this forces a new resource to be created.
     */
    public readonly platformFaultDomainCount!: pulumi.Output<number | undefined>;
    /**
     * Specifies the number of update domains that are used. Defaults to `5`. Changing this forces a new resource to be created.
     */
    public readonly platformUpdateDomainCount!: pulumi.Output<number | undefined>;
    /**
     * The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created.
     */
    public readonly proximityPlacementGroupId!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a AvailabilitySet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AvailabilitySetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AvailabilitySetArgs | AvailabilitySetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AvailabilitySetState | undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["managed"] = state ? state.managed : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["platformFaultDomainCount"] = state ? state.platformFaultDomainCount : undefined;
            inputs["platformUpdateDomainCount"] = state ? state.platformUpdateDomainCount : undefined;
            inputs["proximityPlacementGroupId"] = state ? state.proximityPlacementGroupId : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as AvailabilitySetArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["managed"] = args ? args.managed : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["platformFaultDomainCount"] = args ? args.platformFaultDomainCount : undefined;
            inputs["platformUpdateDomainCount"] = args ? args.platformUpdateDomainCount : undefined;
            inputs["proximityPlacementGroupId"] = args ? args.proximityPlacementGroupId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AvailabilitySet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AvailabilitySet resources.
 */
export interface AvailabilitySetState {
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
     */
    managed?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the availability set. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the number of fault domains that are used. Defaults to `3`. Changing this forces a new resource to be created.
     */
    platformFaultDomainCount?: pulumi.Input<number>;
    /**
     * Specifies the number of update domains that are used. Defaults to `5`. Changing this forces a new resource to be created.
     */
    platformUpdateDomainCount?: pulumi.Input<number>;
    /**
     * The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created.
     */
    proximityPlacementGroupId?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a AvailabilitySet resource.
 */
export interface AvailabilitySetArgs {
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
     */
    managed?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the availability set. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the number of fault domains that are used. Defaults to `3`. Changing this forces a new resource to be created.
     */
    platformFaultDomainCount?: pulumi.Input<number>;
    /**
     * Specifies the number of update domains that are used. Defaults to `5`. Changing this forces a new resource to be created.
     */
    platformUpdateDomainCount?: pulumi.Input<number>;
    /**
     * The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created.
     */
    proximityPlacementGroupId?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

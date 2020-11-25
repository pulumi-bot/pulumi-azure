// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Virtual Desktop Host Pool.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleHostPool = new azure.desktopvirtualization.HostPool("exampleHostPool", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     friendlyName: "pooleddepthfirst",
 *     validateEnvironment: true,
 *     description: "Acceptance Test: A pooled host pool - pooleddepthfirst",
 *     type: "Pooled",
 *     maximumSessionsAllowed: 50,
 *     loadBalancerType: "DepthFirst",
 * });
 * ```
 *
 * ## Import
 *
 * Virtual Desktop Host Pools can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:desktopvirtualization/hostPool:HostPool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.DesktopVirtualization/hostpools/myhostpool
 * ```
 */
export class HostPool extends pulumi.CustomResource {
    /**
     * Get an existing HostPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HostPoolState, opts?: pulumi.CustomResourceOptions): HostPool {
        return new HostPool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:desktopvirtualization/hostPool:HostPool';

    /**
     * Returns true if the given object is an instance of HostPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HostPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HostPool.__pulumiType;
    }

    /**
     * A description for the Virtual Desktop Host Pool.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A friendly name for the Virtual Desktop Host Pool.
     */
    public readonly friendlyName!: pulumi.Output<string | undefined>;
    /**
     * `BreadthFirst` load balancing distributes new user sessions across all available session hosts in the host pool.
     * `DepthFirst` load balancing distributes new user sessions to an available session host with the highest number of connections but has not reached its maximum session limit threshold.
     * `Persistent` should be used if the host pool type is `Personal`
     */
    public readonly loadBalancerType!: pulumi.Output<string>;
    /**
     * The location/region where the Virtual Desktop Host Pool is
     * located. Changing the location/region forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * A valid integer value from 0 to 999999 for the maximum number of users that have concurrent sessions on a session host.
     * Should only be set if the `type` of your Virtual Desktop Host Pool is `Pooled`.
     */
    public readonly maximumSessionsAllowed!: pulumi.Output<number | undefined>;
    /**
     * The name of the Virtual Desktop Host Pool. Changing the name
     * forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * `Automatic` assignment – The service will select an available host and assign it to an user.
     * `Direct` Assignment – Admin selects a specific host to assign to an user.
     */
    public readonly personalDesktopAssignmentType!: pulumi.Output<string | undefined>;
    /**
     * Option to specify the preferred Application Group type for the Virtual Desktop Host Pool.
     * Valid options are `None`, `Desktop` or `RailApplications`. Default is `None`.
     */
    public readonly preferredAppGroupType!: pulumi.Output<string | undefined>;
    /**
     * A `registrationInfo` block which is documented below. Specifies configuration on the registration information of the Virtual Desktop Host Pool.
     */
    public readonly registrationInfo!: pulumi.Output<outputs.desktopvirtualization.HostPoolRegistrationInfo | undefined>;
    /**
     * The name of the resource group in which to
     * create the Virtual Desktop Host Pool. Changing the resource group name forces
     * a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the Virtual Desktop Host Pool. Valid options are
     * `Personal` or `Pooled`. Changing the type forces a new resource to be created.
     */
    public readonly type!: pulumi.Output<string>;
    public readonly validateEnvironment!: pulumi.Output<boolean | undefined>;

    /**
     * Create a HostPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostPoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HostPoolArgs | HostPoolState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as HostPoolState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["friendlyName"] = state ? state.friendlyName : undefined;
            inputs["loadBalancerType"] = state ? state.loadBalancerType : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["maximumSessionsAllowed"] = state ? state.maximumSessionsAllowed : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["personalDesktopAssignmentType"] = state ? state.personalDesktopAssignmentType : undefined;
            inputs["preferredAppGroupType"] = state ? state.preferredAppGroupType : undefined;
            inputs["registrationInfo"] = state ? state.registrationInfo : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["validateEnvironment"] = state ? state.validateEnvironment : undefined;
        } else {
            const args = argsOrState as HostPoolArgs | undefined;
            if (!args || args.loadBalancerType === undefined) {
                throw new Error("Missing required property 'loadBalancerType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["friendlyName"] = args ? args.friendlyName : undefined;
            inputs["loadBalancerType"] = args ? args.loadBalancerType : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["maximumSessionsAllowed"] = args ? args.maximumSessionsAllowed : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["personalDesktopAssignmentType"] = args ? args.personalDesktopAssignmentType : undefined;
            inputs["preferredAppGroupType"] = args ? args.preferredAppGroupType : undefined;
            inputs["registrationInfo"] = args ? args.registrationInfo : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["validateEnvironment"] = args ? args.validateEnvironment : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(HostPool.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HostPool resources.
 */
export interface HostPoolState {
    /**
     * A description for the Virtual Desktop Host Pool.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A friendly name for the Virtual Desktop Host Pool.
     */
    readonly friendlyName?: pulumi.Input<string>;
    /**
     * `BreadthFirst` load balancing distributes new user sessions across all available session hosts in the host pool.
     * `DepthFirst` load balancing distributes new user sessions to an available session host with the highest number of connections but has not reached its maximum session limit threshold.
     * `Persistent` should be used if the host pool type is `Personal`
     */
    readonly loadBalancerType?: pulumi.Input<string>;
    /**
     * The location/region where the Virtual Desktop Host Pool is
     * located. Changing the location/region forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * A valid integer value from 0 to 999999 for the maximum number of users that have concurrent sessions on a session host.
     * Should only be set if the `type` of your Virtual Desktop Host Pool is `Pooled`.
     */
    readonly maximumSessionsAllowed?: pulumi.Input<number>;
    /**
     * The name of the Virtual Desktop Host Pool. Changing the name
     * forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * `Automatic` assignment – The service will select an available host and assign it to an user.
     * `Direct` Assignment – Admin selects a specific host to assign to an user.
     */
    readonly personalDesktopAssignmentType?: pulumi.Input<string>;
    /**
     * Option to specify the preferred Application Group type for the Virtual Desktop Host Pool.
     * Valid options are `None`, `Desktop` or `RailApplications`. Default is `None`.
     */
    readonly preferredAppGroupType?: pulumi.Input<string>;
    /**
     * A `registrationInfo` block which is documented below. Specifies configuration on the registration information of the Virtual Desktop Host Pool.
     */
    readonly registrationInfo?: pulumi.Input<inputs.desktopvirtualization.HostPoolRegistrationInfo>;
    /**
     * The name of the resource group in which to
     * create the Virtual Desktop Host Pool. Changing the resource group name forces
     * a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of the Virtual Desktop Host Pool. Valid options are
     * `Personal` or `Pooled`. Changing the type forces a new resource to be created.
     */
    readonly type?: pulumi.Input<string>;
    readonly validateEnvironment?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a HostPool resource.
 */
export interface HostPoolArgs {
    /**
     * A description for the Virtual Desktop Host Pool.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A friendly name for the Virtual Desktop Host Pool.
     */
    readonly friendlyName?: pulumi.Input<string>;
    /**
     * `BreadthFirst` load balancing distributes new user sessions across all available session hosts in the host pool.
     * `DepthFirst` load balancing distributes new user sessions to an available session host with the highest number of connections but has not reached its maximum session limit threshold.
     * `Persistent` should be used if the host pool type is `Personal`
     */
    readonly loadBalancerType: pulumi.Input<string>;
    /**
     * The location/region where the Virtual Desktop Host Pool is
     * located. Changing the location/region forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * A valid integer value from 0 to 999999 for the maximum number of users that have concurrent sessions on a session host.
     * Should only be set if the `type` of your Virtual Desktop Host Pool is `Pooled`.
     */
    readonly maximumSessionsAllowed?: pulumi.Input<number>;
    /**
     * The name of the Virtual Desktop Host Pool. Changing the name
     * forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * `Automatic` assignment – The service will select an available host and assign it to an user.
     * `Direct` Assignment – Admin selects a specific host to assign to an user.
     */
    readonly personalDesktopAssignmentType?: pulumi.Input<string>;
    /**
     * Option to specify the preferred Application Group type for the Virtual Desktop Host Pool.
     * Valid options are `None`, `Desktop` or `RailApplications`. Default is `None`.
     */
    readonly preferredAppGroupType?: pulumi.Input<string>;
    /**
     * A `registrationInfo` block which is documented below. Specifies configuration on the registration information of the Virtual Desktop Host Pool.
     */
    readonly registrationInfo?: pulumi.Input<inputs.desktopvirtualization.HostPoolRegistrationInfo>;
    /**
     * The name of the resource group in which to
     * create the Virtual Desktop Host Pool. Changing the resource group name forces
     * a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of the Virtual Desktop Host Pool. Valid options are
     * `Personal` or `Pooled`. Changing the type forces a new resource to be created.
     */
    readonly type: pulumi.Input<string>;
    readonly validateEnvironment?: pulumi.Input<boolean>;
}

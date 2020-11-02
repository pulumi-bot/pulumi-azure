// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a maintenance assignment to Dedicated Host.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleDedicatedHostGroup = new azure.compute.DedicatedHostGroup("exampleDedicatedHostGroup", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     platformFaultDomainCount: 2,
 * });
 * const exampleDedicatedHost = new azure.compute.DedicatedHost("exampleDedicatedHost", {
 *     location: exampleResourceGroup.location,
 *     dedicatedHostGroupId: exampleDedicatedHostGroup.id,
 *     skuName: "DSv3-Type1",
 *     platformFaultDomain: 1,
 * });
 * const exampleConfiguration = new azure.maintenance.Configuration("exampleConfiguration", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     scope: "All",
 * });
 * const exampleAssignmentDedicatedHost = new azure.maintenance.AssignmentDedicatedHost("exampleAssignmentDedicatedHost", {
 *     location: exampleResourceGroup.location,
 *     maintenanceConfigurationId: exampleConfiguration.id,
 *     dedicatedHostId: exampleDedicatedHost.id,
 * });
 * ```
 *
 * ## Import
 *
 * Maintenance Assignment can be imported using the `resource id`, e.g.
 */
export class AssignmentDedicatedHost extends pulumi.CustomResource {
    /**
     * Get an existing AssignmentDedicatedHost resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AssignmentDedicatedHostState, opts?: pulumi.CustomResourceOptions): AssignmentDedicatedHost {
        return new AssignmentDedicatedHost(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:maintenance/assignmentDedicatedHost:AssignmentDedicatedHost';

    /**
     * Returns true if the given object is an instance of AssignmentDedicatedHost.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AssignmentDedicatedHost {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AssignmentDedicatedHost.__pulumiType;
    }

    /**
     * Specifies the Dedicated Host ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
     */
    public readonly dedicatedHostId!: pulumi.Output<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
     */
    public readonly maintenanceConfigurationId!: pulumi.Output<string>;

    /**
     * Create a AssignmentDedicatedHost resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AssignmentDedicatedHostArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AssignmentDedicatedHostArgs | AssignmentDedicatedHostState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AssignmentDedicatedHostState | undefined;
            inputs["dedicatedHostId"] = state ? state.dedicatedHostId : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["maintenanceConfigurationId"] = state ? state.maintenanceConfigurationId : undefined;
        } else {
            const args = argsOrState as AssignmentDedicatedHostArgs | undefined;
            if (!args || args.dedicatedHostId === undefined) {
                throw new Error("Missing required property 'dedicatedHostId'");
            }
            if (!args || args.maintenanceConfigurationId === undefined) {
                throw new Error("Missing required property 'maintenanceConfigurationId'");
            }
            inputs["dedicatedHostId"] = args ? args.dedicatedHostId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["maintenanceConfigurationId"] = args ? args.maintenanceConfigurationId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AssignmentDedicatedHost.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AssignmentDedicatedHost resources.
 */
export interface AssignmentDedicatedHostState {
    /**
     * Specifies the Dedicated Host ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
     */
    readonly dedicatedHostId?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
     */
    readonly maintenanceConfigurationId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AssignmentDedicatedHost resource.
 */
export interface AssignmentDedicatedHostArgs {
    /**
     * Specifies the Dedicated Host ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
     */
    readonly dedicatedHostId: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
     */
    readonly maintenanceConfigurationId: pulumi.Input<string>;
}

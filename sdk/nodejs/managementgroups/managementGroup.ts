// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Management Group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const current = azure.core.getSubscription({});
 * const exampleParent = new azure.management.Group("exampleParent", {
 *     displayName: "ParentGroup",
 *     subscriptionIds: [current.then(current => current.subscriptionId)],
 * });
 * const exampleChild = new azure.management.Group("exampleChild", {
 *     displayName: "ChildGroup",
 *     parentManagementGroupId: exampleParent.id,
 *     subscriptionIds: [current.then(current => current.subscriptionId)],
 * });
 * // other subscription IDs can go here
 * ```
 *
 * ## Import
 *
 * Management Groups can be imported using the `management group resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:managementgroups/managementGroup:ManagementGroup example /providers/Microsoft.Management/managementGroups/group1
 * ```
 *
 * @deprecated azure.managementgroups.ManagementGroup has been deprecated in favor of azure.management.Group
 */
export class ManagementGroup extends pulumi.CustomResource {
    /**
     * Get an existing ManagementGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ManagementGroupState, opts?: pulumi.CustomResourceOptions): ManagementGroup {
        pulumi.log.warn("ManagementGroup is deprecated: azure.managementgroups.ManagementGroup has been deprecated in favor of azure.management.Group")
        return new ManagementGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:managementgroups/managementGroup:ManagementGroup';

    /**
     * Returns true if the given object is an instance of ManagementGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagementGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagementGroup.__pulumiType;
    }

    /**
     * A friendly name for this Management Group. If not specified, this'll be the same as the `name`.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
     *
     * @deprecated Deprecated in favour of `name`
     */
    public readonly groupId!: pulumi.Output<string>;
    /**
     * The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the Parent Management Group. Changing this forces a new resource to be created.
     */
    public readonly parentManagementGroupId!: pulumi.Output<string>;
    /**
     * A list of Subscription GUIDs which should be assigned to the Management Group.
     */
    public readonly subscriptionIds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a ManagementGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated azure.managementgroups.ManagementGroup has been deprecated in favor of azure.management.Group */
    constructor(name: string, args?: ManagementGroupArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azure.managementgroups.ManagementGroup has been deprecated in favor of azure.management.Group */
    constructor(name: string, argsOrState?: ManagementGroupArgs | ManagementGroupState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("ManagementGroup is deprecated: azure.managementgroups.ManagementGroup has been deprecated in favor of azure.management.Group")
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ManagementGroupState | undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["groupId"] = state ? state.groupId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parentManagementGroupId"] = state ? state.parentManagementGroupId : undefined;
            inputs["subscriptionIds"] = state ? state.subscriptionIds : undefined;
        } else {
            const args = argsOrState as ManagementGroupArgs | undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["groupId"] = args ? args.groupId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parentManagementGroupId"] = args ? args.parentManagementGroupId : undefined;
            inputs["subscriptionIds"] = args ? args.subscriptionIds : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ManagementGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ManagementGroup resources.
 */
export interface ManagementGroupState {
    /**
     * A friendly name for this Management Group. If not specified, this'll be the same as the `name`.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
     *
     * @deprecated Deprecated in favour of `name`
     */
    readonly groupId?: pulumi.Input<string>;
    /**
     * The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the Parent Management Group. Changing this forces a new resource to be created.
     */
    readonly parentManagementGroupId?: pulumi.Input<string>;
    /**
     * A list of Subscription GUIDs which should be assigned to the Management Group.
     */
    readonly subscriptionIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a ManagementGroup resource.
 */
export interface ManagementGroupArgs {
    /**
     * A friendly name for this Management Group. If not specified, this'll be the same as the `name`.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
     *
     * @deprecated Deprecated in favour of `name`
     */
    readonly groupId?: pulumi.Input<string>;
    /**
     * The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the Parent Management Group. Changing this forces a new resource to be created.
     */
    readonly parentManagementGroupId?: pulumi.Input<string>;
    /**
     * A list of Subscription GUIDs which should be assigned to the Management Group.
     */
    readonly subscriptionIds?: pulumi.Input<pulumi.Input<string>[]>;
}

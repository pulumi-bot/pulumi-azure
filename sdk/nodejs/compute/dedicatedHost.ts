// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manage a Dedicated Host within a Dedicated Host Group.
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
 * ```
 *
 * ## Import
 *
 * Dedicated Hosts can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:compute/dedicatedHost:DedicatedHost example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/hostGroups/group1/hosts/host1
 * ```
 */
export class DedicatedHost extends pulumi.CustomResource {
    /**
     * Get an existing DedicatedHost resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DedicatedHostState, opts?: pulumi.CustomResourceOptions): DedicatedHost {
        return new DedicatedHost(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:compute/dedicatedHost:DedicatedHost';

    /**
     * Returns true if the given object is an instance of DedicatedHost.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DedicatedHost {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DedicatedHost.__pulumiType;
    }

    /**
     * Should the Dedicated Host automatically be replaced in case of a Hardware Failure? Defaults to `true`.
     */
    public readonly autoReplaceOnFailure!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the ID of the Dedicated Host Group where the Dedicated Host should exist. Changing this forces a new resource to be created.
     */
    public readonly dedicatedHostGroupId!: pulumi.Output<string>;
    /**
     * Specifies the software license type that will be applied to the VMs deployed on the Dedicated Host. Possible values are `None`, `Windows_Server_Hybrid` and `Windows_Server_Perpetual`. Defaults to `None`.
     */
    public readonly licenseType!: pulumi.Output<string | undefined>;
    /**
     * Specify the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of this Dedicated Host. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specify the fault domain of the Dedicated Host Group in which to create the Dedicated Host. Changing this forces a new resource to be created.
     */
    public readonly platformFaultDomain!: pulumi.Output<number>;
    /**
     * Specify the sku name of the Dedicated Host. Possible values are `DSv3-Type1`, `DSv3-Type2`, `DSv4-Type1`, `ESv3-Type1`, `ESv3-Type2`,`FSv2-Type2`. Changing this forces a new resource to be created.
     */
    public readonly skuName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a DedicatedHost resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DedicatedHostArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DedicatedHostArgs | DedicatedHostState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DedicatedHostState | undefined;
            inputs["autoReplaceOnFailure"] = state ? state.autoReplaceOnFailure : undefined;
            inputs["dedicatedHostGroupId"] = state ? state.dedicatedHostGroupId : undefined;
            inputs["licenseType"] = state ? state.licenseType : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["platformFaultDomain"] = state ? state.platformFaultDomain : undefined;
            inputs["skuName"] = state ? state.skuName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as DedicatedHostArgs | undefined;
            if ((!args || args.dedicatedHostGroupId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'dedicatedHostGroupId'");
            }
            if ((!args || args.platformFaultDomain === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'platformFaultDomain'");
            }
            if ((!args || args.skuName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'skuName'");
            }
            inputs["autoReplaceOnFailure"] = args ? args.autoReplaceOnFailure : undefined;
            inputs["dedicatedHostGroupId"] = args ? args.dedicatedHostGroupId : undefined;
            inputs["licenseType"] = args ? args.licenseType : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["platformFaultDomain"] = args ? args.platformFaultDomain : undefined;
            inputs["skuName"] = args ? args.skuName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DedicatedHost.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DedicatedHost resources.
 */
export interface DedicatedHostState {
    /**
     * Should the Dedicated Host automatically be replaced in case of a Hardware Failure? Defaults to `true`.
     */
    readonly autoReplaceOnFailure?: pulumi.Input<boolean>;
    /**
     * Specifies the ID of the Dedicated Host Group where the Dedicated Host should exist. Changing this forces a new resource to be created.
     */
    readonly dedicatedHostGroupId?: pulumi.Input<string>;
    /**
     * Specifies the software license type that will be applied to the VMs deployed on the Dedicated Host. Possible values are `None`, `Windows_Server_Hybrid` and `Windows_Server_Perpetual`. Defaults to `None`.
     */
    readonly licenseType?: pulumi.Input<string>;
    /**
     * Specify the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of this Dedicated Host. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specify the fault domain of the Dedicated Host Group in which to create the Dedicated Host. Changing this forces a new resource to be created.
     */
    readonly platformFaultDomain?: pulumi.Input<number>;
    /**
     * Specify the sku name of the Dedicated Host. Possible values are `DSv3-Type1`, `DSv3-Type2`, `DSv4-Type1`, `ESv3-Type1`, `ESv3-Type2`,`FSv2-Type2`. Changing this forces a new resource to be created.
     */
    readonly skuName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a DedicatedHost resource.
 */
export interface DedicatedHostArgs {
    /**
     * Should the Dedicated Host automatically be replaced in case of a Hardware Failure? Defaults to `true`.
     */
    readonly autoReplaceOnFailure?: pulumi.Input<boolean>;
    /**
     * Specifies the ID of the Dedicated Host Group where the Dedicated Host should exist. Changing this forces a new resource to be created.
     */
    readonly dedicatedHostGroupId: pulumi.Input<string>;
    /**
     * Specifies the software license type that will be applied to the VMs deployed on the Dedicated Host. Possible values are `None`, `Windows_Server_Hybrid` and `Windows_Server_Perpetual`. Defaults to `None`.
     */
    readonly licenseType?: pulumi.Input<string>;
    /**
     * Specify the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of this Dedicated Host. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specify the fault domain of the Dedicated Host Group in which to create the Dedicated Host. Changing this forces a new resource to be created.
     */
    readonly platformFaultDomain: pulumi.Input<number>;
    /**
     * Specify the sku name of the Dedicated Host. Possible values are `DSv3-Type1`, `DSv3-Type2`, `DSv4-Type1`, `ESv3-Type1`, `ESv3-Type2`,`FSv2-Type2`. Changing this forces a new resource to be created.
     */
    readonly skuName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Synapse Spark Pool.
 */
export class SparkPool extends pulumi.CustomResource {
    /**
     * Get an existing SparkPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SparkPoolState, opts?: pulumi.CustomResourceOptions): SparkPool {
        return new SparkPool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:synapse/sparkPool:SparkPool';

    /**
     * Returns true if the given object is an instance of SparkPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SparkPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SparkPool.__pulumiType;
    }

    /**
     * An `autoPause` block as defined below.
     */
    public readonly autoPause!: pulumi.Output<outputs.synapse.SparkPoolAutoPause | undefined>;
    /**
     * An `autoScale` block as defined below. Exactly one of `nodeCount` or `autoScale` must be specified.
     */
    public readonly autoScale!: pulumi.Output<outputs.synapse.SparkPoolAutoScale | undefined>;
    /**
     * A `libraryRequirement` block as defined below.
     */
    public readonly libraryRequirement!: pulumi.Output<outputs.synapse.SparkPoolLibraryRequirement | undefined>;
    /**
     * The name which should be used for this Synapse Spark Pool. Changing this forces a new Synapse Spark Pool to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The number of nodes in the Spark Pool. Exactly one of `nodeCount` or `autoScale` must be specified.
     */
    public readonly nodeCount!: pulumi.Output<number | undefined>;
    /**
     * The level of node in the Spark Pool. Possible value is `Small`, `Medium` and `Large`.
     */
    public readonly nodeSize!: pulumi.Output<string>;
    /**
     * The kind of nodes that the Spark Pool provides. Possible value is `MemoryOptimized`.
     */
    public readonly nodeSizeFamily!: pulumi.Output<string>;
    /**
     * The Spark events folder. Defaults to `/events`.
     */
    public readonly sparkEventsFolder!: pulumi.Output<string | undefined>;
    /**
     * The default folder where Spark logs will be written. Defaults to `/logs`.
     */
    public readonly sparkLogFolder!: pulumi.Output<string | undefined>;
    /**
     * The Apache Spark version. Possible value is `2.4`. Defaults to `2.4`.
     */
    public readonly sparkVersion!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Synapse Workspace where the Synapse Spark Pool should exist. Changing this forces a new Synapse Spark Pool to be created.
     */
    public readonly synapseWorkspaceId!: pulumi.Output<string>;
    /**
     * A mapping of tags which should be assigned to the Synapse Spark Pool.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a SparkPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SparkPoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SparkPoolArgs | SparkPoolState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SparkPoolState | undefined;
            inputs["autoPause"] = state ? state.autoPause : undefined;
            inputs["autoScale"] = state ? state.autoScale : undefined;
            inputs["libraryRequirement"] = state ? state.libraryRequirement : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nodeCount"] = state ? state.nodeCount : undefined;
            inputs["nodeSize"] = state ? state.nodeSize : undefined;
            inputs["nodeSizeFamily"] = state ? state.nodeSizeFamily : undefined;
            inputs["sparkEventsFolder"] = state ? state.sparkEventsFolder : undefined;
            inputs["sparkLogFolder"] = state ? state.sparkLogFolder : undefined;
            inputs["sparkVersion"] = state ? state.sparkVersion : undefined;
            inputs["synapseWorkspaceId"] = state ? state.synapseWorkspaceId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as SparkPoolArgs | undefined;
            if (!args || args.nodeSize === undefined) {
                throw new Error("Missing required property 'nodeSize'");
            }
            if (!args || args.nodeSizeFamily === undefined) {
                throw new Error("Missing required property 'nodeSizeFamily'");
            }
            if (!args || args.synapseWorkspaceId === undefined) {
                throw new Error("Missing required property 'synapseWorkspaceId'");
            }
            inputs["autoPause"] = args ? args.autoPause : undefined;
            inputs["autoScale"] = args ? args.autoScale : undefined;
            inputs["libraryRequirement"] = args ? args.libraryRequirement : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nodeCount"] = args ? args.nodeCount : undefined;
            inputs["nodeSize"] = args ? args.nodeSize : undefined;
            inputs["nodeSizeFamily"] = args ? args.nodeSizeFamily : undefined;
            inputs["sparkEventsFolder"] = args ? args.sparkEventsFolder : undefined;
            inputs["sparkLogFolder"] = args ? args.sparkLogFolder : undefined;
            inputs["sparkVersion"] = args ? args.sparkVersion : undefined;
            inputs["synapseWorkspaceId"] = args ? args.synapseWorkspaceId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SparkPool.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SparkPool resources.
 */
export interface SparkPoolState {
    /**
     * An `autoPause` block as defined below.
     */
    readonly autoPause?: pulumi.Input<inputs.synapse.SparkPoolAutoPause>;
    /**
     * An `autoScale` block as defined below. Exactly one of `nodeCount` or `autoScale` must be specified.
     */
    readonly autoScale?: pulumi.Input<inputs.synapse.SparkPoolAutoScale>;
    /**
     * A `libraryRequirement` block as defined below.
     */
    readonly libraryRequirement?: pulumi.Input<inputs.synapse.SparkPoolLibraryRequirement>;
    /**
     * The name which should be used for this Synapse Spark Pool. Changing this forces a new Synapse Spark Pool to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The number of nodes in the Spark Pool. Exactly one of `nodeCount` or `autoScale` must be specified.
     */
    readonly nodeCount?: pulumi.Input<number>;
    /**
     * The level of node in the Spark Pool. Possible value is `Small`, `Medium` and `Large`.
     */
    readonly nodeSize?: pulumi.Input<string>;
    /**
     * The kind of nodes that the Spark Pool provides. Possible value is `MemoryOptimized`.
     */
    readonly nodeSizeFamily?: pulumi.Input<string>;
    /**
     * The Spark events folder. Defaults to `/events`.
     */
    readonly sparkEventsFolder?: pulumi.Input<string>;
    /**
     * The default folder where Spark logs will be written. Defaults to `/logs`.
     */
    readonly sparkLogFolder?: pulumi.Input<string>;
    /**
     * The Apache Spark version. Possible value is `2.4`. Defaults to `2.4`.
     */
    readonly sparkVersion?: pulumi.Input<string>;
    /**
     * The ID of the Synapse Workspace where the Synapse Spark Pool should exist. Changing this forces a new Synapse Spark Pool to be created.
     */
    readonly synapseWorkspaceId?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Synapse Spark Pool.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a SparkPool resource.
 */
export interface SparkPoolArgs {
    /**
     * An `autoPause` block as defined below.
     */
    readonly autoPause?: pulumi.Input<inputs.synapse.SparkPoolAutoPause>;
    /**
     * An `autoScale` block as defined below. Exactly one of `nodeCount` or `autoScale` must be specified.
     */
    readonly autoScale?: pulumi.Input<inputs.synapse.SparkPoolAutoScale>;
    /**
     * A `libraryRequirement` block as defined below.
     */
    readonly libraryRequirement?: pulumi.Input<inputs.synapse.SparkPoolLibraryRequirement>;
    /**
     * The name which should be used for this Synapse Spark Pool. Changing this forces a new Synapse Spark Pool to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The number of nodes in the Spark Pool. Exactly one of `nodeCount` or `autoScale` must be specified.
     */
    readonly nodeCount?: pulumi.Input<number>;
    /**
     * The level of node in the Spark Pool. Possible value is `Small`, `Medium` and `Large`.
     */
    readonly nodeSize: pulumi.Input<string>;
    /**
     * The kind of nodes that the Spark Pool provides. Possible value is `MemoryOptimized`.
     */
    readonly nodeSizeFamily: pulumi.Input<string>;
    /**
     * The Spark events folder. Defaults to `/events`.
     */
    readonly sparkEventsFolder?: pulumi.Input<string>;
    /**
     * The default folder where Spark logs will be written. Defaults to `/logs`.
     */
    readonly sparkLogFolder?: pulumi.Input<string>;
    /**
     * The Apache Spark version. Possible value is `2.4`. Defaults to `2.4`.
     */
    readonly sparkVersion?: pulumi.Input<string>;
    /**
     * The ID of the Synapse Workspace where the Synapse Spark Pool should exist. Changing this forces a new Synapse Spark Pool to be created.
     */
    readonly synapseWorkspaceId: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Synapse Spark Pool.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Synapse Spark Pool.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 *     accountKind: "StorageV2",
 *     isHnsEnabled: "true",
 * });
 * const exampleDataLakeGen2Filesystem = new azure.storage.DataLakeGen2Filesystem("exampleDataLakeGen2Filesystem", {storageAccountId: exampleAccount.id});
 * const exampleWorkspace = new azure.synapse.Workspace("exampleWorkspace", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     storageDataLakeGen2FilesystemId: exampleDataLakeGen2Filesystem.id,
 *     sqlAdministratorLogin: "sqladminuser",
 *     sqlAdministratorLoginPassword: "H@Sh1CoR3!",
 * });
 * const exampleSparkPool = new azure.synapse.SparkPool("exampleSparkPool", {
 *     synapseWorkspaceId: exampleWorkspace.id,
 *     nodeSizeFamily: "MemoryOptimized",
 *     nodeSize: "Small",
 *     autoScale: {
 *         maxNodeCount: 50,
 *         minNodeCount: 3,
 *     },
 *     autoPause: {
 *         delayInMinutes: 15,
 *     },
 *     tags: {
 *         ENV: "Production",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Synapse Spark Pool can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:synapse/sparkPool:SparkPool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/bigDataPools/sparkPool1
 * ```
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
        opts = opts || {};
        if (opts.id) {
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
            if ((!args || args.nodeSize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeSize'");
            }
            if ((!args || args.nodeSizeFamily === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeSizeFamily'");
            }
            if ((!args || args.synapseWorkspaceId === undefined) && !opts.urn) {
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
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
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
    autoPause?: pulumi.Input<inputs.synapse.SparkPoolAutoPause>;
    /**
     * An `autoScale` block as defined below. Exactly one of `nodeCount` or `autoScale` must be specified.
     */
    autoScale?: pulumi.Input<inputs.synapse.SparkPoolAutoScale>;
    /**
     * A `libraryRequirement` block as defined below.
     */
    libraryRequirement?: pulumi.Input<inputs.synapse.SparkPoolLibraryRequirement>;
    /**
     * The name which should be used for this Synapse Spark Pool. Changing this forces a new Synapse Spark Pool to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The number of nodes in the Spark Pool. Exactly one of `nodeCount` or `autoScale` must be specified.
     */
    nodeCount?: pulumi.Input<number>;
    /**
     * The level of node in the Spark Pool. Possible value is `Small`, `Medium` and `Large`.
     */
    nodeSize?: pulumi.Input<string>;
    /**
     * The kind of nodes that the Spark Pool provides. Possible value is `MemoryOptimized`.
     */
    nodeSizeFamily?: pulumi.Input<string>;
    /**
     * The Spark events folder. Defaults to `/events`.
     */
    sparkEventsFolder?: pulumi.Input<string>;
    /**
     * The default folder where Spark logs will be written. Defaults to `/logs`.
     */
    sparkLogFolder?: pulumi.Input<string>;
    /**
     * The Apache Spark version. Possible value is `2.4`. Defaults to `2.4`.
     */
    sparkVersion?: pulumi.Input<string>;
    /**
     * The ID of the Synapse Workspace where the Synapse Spark Pool should exist. Changing this forces a new Synapse Spark Pool to be created.
     */
    synapseWorkspaceId?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Synapse Spark Pool.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a SparkPool resource.
 */
export interface SparkPoolArgs {
    /**
     * An `autoPause` block as defined below.
     */
    autoPause?: pulumi.Input<inputs.synapse.SparkPoolAutoPause>;
    /**
     * An `autoScale` block as defined below. Exactly one of `nodeCount` or `autoScale` must be specified.
     */
    autoScale?: pulumi.Input<inputs.synapse.SparkPoolAutoScale>;
    /**
     * A `libraryRequirement` block as defined below.
     */
    libraryRequirement?: pulumi.Input<inputs.synapse.SparkPoolLibraryRequirement>;
    /**
     * The name which should be used for this Synapse Spark Pool. Changing this forces a new Synapse Spark Pool to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The number of nodes in the Spark Pool. Exactly one of `nodeCount` or `autoScale` must be specified.
     */
    nodeCount?: pulumi.Input<number>;
    /**
     * The level of node in the Spark Pool. Possible value is `Small`, `Medium` and `Large`.
     */
    nodeSize: pulumi.Input<string>;
    /**
     * The kind of nodes that the Spark Pool provides. Possible value is `MemoryOptimized`.
     */
    nodeSizeFamily: pulumi.Input<string>;
    /**
     * The Spark events folder. Defaults to `/events`.
     */
    sparkEventsFolder?: pulumi.Input<string>;
    /**
     * The default folder where Spark logs will be written. Defaults to `/logs`.
     */
    sparkLogFolder?: pulumi.Input<string>;
    /**
     * The Apache Spark version. Possible value is `2.4`. Defaults to `2.4`.
     */
    sparkVersion?: pulumi.Input<string>;
    /**
     * The ID of the Synapse Workspace where the Synapse Spark Pool should exist. Changing this forces a new Synapse Spark Pool to be created.
     */
    synapseWorkspaceId: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Synapse Spark Pool.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

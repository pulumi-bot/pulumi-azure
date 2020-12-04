// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Queue within an Azure Storage Account.
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
 * });
 * const exampleQueue = new azure.storage.Queue("exampleQueue", {storageAccountName: exampleAccount.name});
 * ```
 *
 * ## Import
 *
 * Storage Queue's can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:storage/queue:Queue queue1 https://example.queue.core.windows.net/queue1
 * ```
 */
export class Queue extends pulumi.CustomResource {
    /**
     * Get an existing Queue resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QueueState, opts?: pulumi.CustomResourceOptions): Queue {
        return new Queue(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:storage/queue:Queue';

    /**
     * Returns true if the given object is an instance of Queue.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Queue {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Queue.__pulumiType;
    }

    /**
     * A mapping of MetaData which should be assigned to this Storage Queue.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the Queue which should be created within the Storage Account. Must be unique within the storage account the queue is located.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the Storage Account in which the Storage Queue should exist. Changing this forces a new resource to be created.
     */
    public readonly storageAccountName!: pulumi.Output<string>;

    /**
     * Create a Queue resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: QueueArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QueueArgs | QueueState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as QueueState | undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["storageAccountName"] = state ? state.storageAccountName : undefined;
        } else {
            const args = argsOrState as QueueArgs | undefined;
            if ((!args || args.storageAccountName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'storageAccountName'");
            }
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["storageAccountName"] = args ? args.storageAccountName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Queue.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Queue resources.
 */
export interface QueueState {
    /**
     * A mapping of MetaData which should be assigned to this Storage Queue.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the Queue which should be created within the Storage Account. Must be unique within the storage account the queue is located.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the Storage Account in which the Storage Queue should exist. Changing this forces a new resource to be created.
     */
    readonly storageAccountName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Queue resource.
 */
export interface QueueArgs {
    /**
     * A mapping of MetaData which should be assigned to this Storage Queue.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the Queue which should be created within the Storage Account. Must be unique within the storage account the queue is located.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the Storage Account in which the Storage Queue should exist. Changing this forces a new resource to be created.
     */
    readonly storageAccountName: pulumi.Input<string>;
}

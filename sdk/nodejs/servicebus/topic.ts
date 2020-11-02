// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a ServiceBus Topic.
 *
 * **Note** Topics can only be created in Namespaces with an SKU of `standard` or higher.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleNamespace = new azure.servicebus.Namespace("exampleNamespace", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "Standard",
 *     tags: {
 *         source: "example",
 *     },
 * });
 * const exampleTopic = new azure.servicebus.Topic("exampleTopic", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     namespaceName: exampleNamespace.name,
 *     enablePartitioning: true,
 * });
 * ```
 *
 * ## Import
 *
 * Service Bus Topics can be imported using the `resource id`, e.g.
 */
export class Topic extends pulumi.CustomResource {
    /**
     * Get an existing Topic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TopicState, opts?: pulumi.CustomResourceOptions): Topic {
        return new Topic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:servicebus/topic:Topic';

    /**
     * Returns true if the given object is an instance of Topic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Topic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Topic.__pulumiType;
    }

    /**
     * The ISO 8601 timespan duration of the idle interval after which the
     * Topic is automatically deleted, minimum of 5 minutes.
     */
    public readonly autoDeleteOnIdle!: pulumi.Output<string>;
    /**
     * The ISO 8601 timespan duration of TTL of messages sent to this topic if no
     * TTL value is set on the message itself.
     */
    public readonly defaultMessageTtl!: pulumi.Output<string>;
    /**
     * The ISO 8601 timespan duration during which
     * duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
     */
    public readonly duplicateDetectionHistoryTimeWindow!: pulumi.Output<string>;
    /**
     * Boolean flag which controls if server-side
     * batched operations are enabled. Defaults to false.
     */
    public readonly enableBatchedOperations!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean flag which controls whether Express Entities
     * are enabled. An express topic holds a message in memory temporarily before writing
     * it to persistent storage. Defaults to false.
     */
    public readonly enableExpress!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean flag which controls whether to enable
     * the topic to be partitioned across multiple message brokers. Defaults to false.
     * Changing this forces a new resource to be created.
     */
    public readonly enablePartitioning!: pulumi.Output<boolean | undefined>;
    /**
     * Integer value which controls the size of
     * memory allocated for the topic. For supported values see the "Queue/topic size"
     * section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
     */
    public readonly maxSizeInMegabytes!: pulumi.Output<number>;
    /**
     * Specifies the name of the ServiceBus Topic resource. Changing this forces a
     * new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the ServiceBus Namespace to create
     * this topic in. Changing this forces a new resource to be created.
     */
    public readonly namespaceName!: pulumi.Output<string>;
    /**
     * Boolean flag which controls whether
     * the Topic requires duplicate detection. Defaults to false. Changing this forces
     * a new resource to be created.
     */
    public readonly requiresDuplicateDetection!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the resource group in which to
     * create the namespace. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * Boolean flag which controls whether the Topic
     * supports ordering. Defaults to false.
     */
    public readonly supportOrdering!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Topic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TopicArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TopicArgs | TopicState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TopicState | undefined;
            inputs["autoDeleteOnIdle"] = state ? state.autoDeleteOnIdle : undefined;
            inputs["defaultMessageTtl"] = state ? state.defaultMessageTtl : undefined;
            inputs["duplicateDetectionHistoryTimeWindow"] = state ? state.duplicateDetectionHistoryTimeWindow : undefined;
            inputs["enableBatchedOperations"] = state ? state.enableBatchedOperations : undefined;
            inputs["enableExpress"] = state ? state.enableExpress : undefined;
            inputs["enablePartitioning"] = state ? state.enablePartitioning : undefined;
            inputs["maxSizeInMegabytes"] = state ? state.maxSizeInMegabytes : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namespaceName"] = state ? state.namespaceName : undefined;
            inputs["requiresDuplicateDetection"] = state ? state.requiresDuplicateDetection : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["supportOrdering"] = state ? state.supportOrdering : undefined;
        } else {
            const args = argsOrState as TopicArgs | undefined;
            if (!args || args.namespaceName === undefined) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["autoDeleteOnIdle"] = args ? args.autoDeleteOnIdle : undefined;
            inputs["defaultMessageTtl"] = args ? args.defaultMessageTtl : undefined;
            inputs["duplicateDetectionHistoryTimeWindow"] = args ? args.duplicateDetectionHistoryTimeWindow : undefined;
            inputs["enableBatchedOperations"] = args ? args.enableBatchedOperations : undefined;
            inputs["enableExpress"] = args ? args.enableExpress : undefined;
            inputs["enablePartitioning"] = args ? args.enablePartitioning : undefined;
            inputs["maxSizeInMegabytes"] = args ? args.maxSizeInMegabytes : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["requiresDuplicateDetection"] = args ? args.requiresDuplicateDetection : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["supportOrdering"] = args ? args.supportOrdering : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure:eventhub/topic:Topic" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Topic.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Topic resources.
 */
export interface TopicState {
    /**
     * The ISO 8601 timespan duration of the idle interval after which the
     * Topic is automatically deleted, minimum of 5 minutes.
     */
    readonly autoDeleteOnIdle?: pulumi.Input<string>;
    /**
     * The ISO 8601 timespan duration of TTL of messages sent to this topic if no
     * TTL value is set on the message itself.
     */
    readonly defaultMessageTtl?: pulumi.Input<string>;
    /**
     * The ISO 8601 timespan duration during which
     * duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
     */
    readonly duplicateDetectionHistoryTimeWindow?: pulumi.Input<string>;
    /**
     * Boolean flag which controls if server-side
     * batched operations are enabled. Defaults to false.
     */
    readonly enableBatchedOperations?: pulumi.Input<boolean>;
    /**
     * Boolean flag which controls whether Express Entities
     * are enabled. An express topic holds a message in memory temporarily before writing
     * it to persistent storage. Defaults to false.
     */
    readonly enableExpress?: pulumi.Input<boolean>;
    /**
     * Boolean flag which controls whether to enable
     * the topic to be partitioned across multiple message brokers. Defaults to false.
     * Changing this forces a new resource to be created.
     */
    readonly enablePartitioning?: pulumi.Input<boolean>;
    /**
     * Integer value which controls the size of
     * memory allocated for the topic. For supported values see the "Queue/topic size"
     * section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
     */
    readonly maxSizeInMegabytes?: pulumi.Input<number>;
    /**
     * Specifies the name of the ServiceBus Topic resource. Changing this forces a
     * new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the ServiceBus Namespace to create
     * this topic in. Changing this forces a new resource to be created.
     */
    readonly namespaceName?: pulumi.Input<string>;
    /**
     * Boolean flag which controls whether
     * the Topic requires duplicate detection. Defaults to false. Changing this forces
     * a new resource to be created.
     */
    readonly requiresDuplicateDetection?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which to
     * create the namespace. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Boolean flag which controls whether the Topic
     * supports ordering. Defaults to false.
     */
    readonly supportOrdering?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Topic resource.
 */
export interface TopicArgs {
    /**
     * The ISO 8601 timespan duration of the idle interval after which the
     * Topic is automatically deleted, minimum of 5 minutes.
     */
    readonly autoDeleteOnIdle?: pulumi.Input<string>;
    /**
     * The ISO 8601 timespan duration of TTL of messages sent to this topic if no
     * TTL value is set on the message itself.
     */
    readonly defaultMessageTtl?: pulumi.Input<string>;
    /**
     * The ISO 8601 timespan duration during which
     * duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
     */
    readonly duplicateDetectionHistoryTimeWindow?: pulumi.Input<string>;
    /**
     * Boolean flag which controls if server-side
     * batched operations are enabled. Defaults to false.
     */
    readonly enableBatchedOperations?: pulumi.Input<boolean>;
    /**
     * Boolean flag which controls whether Express Entities
     * are enabled. An express topic holds a message in memory temporarily before writing
     * it to persistent storage. Defaults to false.
     */
    readonly enableExpress?: pulumi.Input<boolean>;
    /**
     * Boolean flag which controls whether to enable
     * the topic to be partitioned across multiple message brokers. Defaults to false.
     * Changing this forces a new resource to be created.
     */
    readonly enablePartitioning?: pulumi.Input<boolean>;
    /**
     * Integer value which controls the size of
     * memory allocated for the topic. For supported values see the "Queue/topic size"
     * section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
     */
    readonly maxSizeInMegabytes?: pulumi.Input<number>;
    /**
     * Specifies the name of the ServiceBus Topic resource. Changing this forces a
     * new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the ServiceBus Namespace to create
     * this topic in. Changing this forces a new resource to be created.
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * Boolean flag which controls whether
     * the Topic requires duplicate detection. Defaults to false. Changing this forces
     * a new resource to be created.
     */
    readonly requiresDuplicateDetection?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which to
     * create the namespace. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Boolean flag which controls whether the Topic
     * supports ordering. Defaults to false.
     */
    readonly supportOrdering?: pulumi.Input<boolean>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Redis Linked Server (ie Geo Location)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example_primaryResourceGroup = new azure.core.ResourceGroup("example-primaryResourceGroup", {location: "East US"});
 * const example_primaryCache = new azure.redis.Cache("example-primaryCache", {
 *     location: example_primaryResourceGroup.location,
 *     resourceGroupName: example_primaryResourceGroup.name,
 *     capacity: 1,
 *     family: "P",
 *     skuName: "Premium",
 *     enableNonSslPort: false,
 *     redisConfiguration: {
 *         maxmemoryReserved: 2,
 *         maxmemoryDelta: 2,
 *         maxmemoryPolicy: "allkeys-lru",
 *     },
 * });
 * const example_secondaryResourceGroup = new azure.core.ResourceGroup("example-secondaryResourceGroup", {location: "West US"});
 * const example_secondaryCache = new azure.redis.Cache("example-secondaryCache", {
 *     location: example_secondaryResourceGroup.location,
 *     resourceGroupName: example_secondaryResourceGroup.name,
 *     capacity: 1,
 *     family: "P",
 *     skuName: "Premium",
 *     enableNonSslPort: false,
 *     redisConfiguration: {
 *         maxmemoryReserved: 2,
 *         maxmemoryDelta: 2,
 *         maxmemoryPolicy: "allkeys-lru",
 *     },
 * });
 * const example_link = new azure.redis.LinkedServer("example-link", {
 *     targetRedisCacheName: example_primaryCache.name,
 *     resourceGroupName: example_primaryCache.resourceGroupName,
 *     linkedRedisCacheId: example_secondaryCache.id,
 *     linkedRedisCacheLocation: example_secondaryCache.location,
 *     serverRole: "Secondary",
 * });
 * ```
 *
 * ## Import
 *
 * Rediss can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:redis/linkedServer:LinkedServer example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/Redis/cache1/linkedServers/cache2
 * ```
 */
export class LinkedServer extends pulumi.CustomResource {
    /**
     * Get an existing LinkedServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LinkedServerState, opts?: pulumi.CustomResourceOptions): LinkedServer {
        return new LinkedServer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:redis/linkedServer:LinkedServer';

    /**
     * Returns true if the given object is an instance of LinkedServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LinkedServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LinkedServer.__pulumiType;
    }

    /**
     * The ID of the linked Redis cache. Changing this forces a new Redis to be created.
     */
    public readonly linkedRedisCacheId!: pulumi.Output<string>;
    /**
     * The location of the linked Redis cache. Changing this forces a new Redis to be created.
     */
    public readonly linkedRedisCacheLocation!: pulumi.Output<string>;
    /**
     * The name of the linked server.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The name of the Resource Group where the Redis caches exists. Changing this forces a new Redis to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The role of the linked Redis cache (eg "Secondary"). Changing this forces a new Redis to be created.
     */
    public readonly serverRole!: pulumi.Output<string>;
    /**
     * The name of Redis cache to link with. Changing this forces a new Redis to be created. (eg The primary role)
     */
    public readonly targetRedisCacheName!: pulumi.Output<string>;

    /**
     * Create a LinkedServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LinkedServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LinkedServerArgs | LinkedServerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LinkedServerState | undefined;
            inputs["linkedRedisCacheId"] = state ? state.linkedRedisCacheId : undefined;
            inputs["linkedRedisCacheLocation"] = state ? state.linkedRedisCacheLocation : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["serverRole"] = state ? state.serverRole : undefined;
            inputs["targetRedisCacheName"] = state ? state.targetRedisCacheName : undefined;
        } else {
            const args = argsOrState as LinkedServerArgs | undefined;
            if (!args || args.linkedRedisCacheId === undefined) {
                throw new Error("Missing required property 'linkedRedisCacheId'");
            }
            if (!args || args.linkedRedisCacheLocation === undefined) {
                throw new Error("Missing required property 'linkedRedisCacheLocation'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serverRole === undefined) {
                throw new Error("Missing required property 'serverRole'");
            }
            if (!args || args.targetRedisCacheName === undefined) {
                throw new Error("Missing required property 'targetRedisCacheName'");
            }
            inputs["linkedRedisCacheId"] = args ? args.linkedRedisCacheId : undefined;
            inputs["linkedRedisCacheLocation"] = args ? args.linkedRedisCacheLocation : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serverRole"] = args ? args.serverRole : undefined;
            inputs["targetRedisCacheName"] = args ? args.targetRedisCacheName : undefined;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LinkedServer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LinkedServer resources.
 */
export interface LinkedServerState {
    /**
     * The ID of the linked Redis cache. Changing this forces a new Redis to be created.
     */
    readonly linkedRedisCacheId?: pulumi.Input<string>;
    /**
     * The location of the linked Redis cache. Changing this forces a new Redis to be created.
     */
    readonly linkedRedisCacheLocation?: pulumi.Input<string>;
    /**
     * The name of the linked server.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Redis caches exists. Changing this forces a new Redis to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The role of the linked Redis cache (eg "Secondary"). Changing this forces a new Redis to be created.
     */
    readonly serverRole?: pulumi.Input<string>;
    /**
     * The name of Redis cache to link with. Changing this forces a new Redis to be created. (eg The primary role)
     */
    readonly targetRedisCacheName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LinkedServer resource.
 */
export interface LinkedServerArgs {
    /**
     * The ID of the linked Redis cache. Changing this forces a new Redis to be created.
     */
    readonly linkedRedisCacheId: pulumi.Input<string>;
    /**
     * The location of the linked Redis cache. Changing this forces a new Redis to be created.
     */
    readonly linkedRedisCacheLocation: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Redis caches exists. Changing this forces a new Redis to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The role of the linked Redis cache (eg "Secondary"). Changing this forces a new Redis to be created.
     */
    readonly serverRole: pulumi.Input<string>;
    /**
     * The name of Redis cache to link with. Changing this forces a new Redis to be created. (eg The primary role)
     */
    readonly targetRedisCacheName: pulumi.Input<string>;
}

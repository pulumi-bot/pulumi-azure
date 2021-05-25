// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Storage Sync Cloud Endpoint.
 *
 * > **NOTE:** Please ensure Azure File Sync has access to the storage account in your subscription, which indicates that `Microsoft.StorageSync` is assigned role `Reader and Data Access` ( refer to details [here](https://docs.microsoft.com/en-us/azure/storage/files/storage-sync-files-troubleshoot?tabs=portal1%2Cazure-portal#common-troubleshooting-steps)).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleSync = new azure.storage.Sync("exampleSync", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 * });
 * const exampleSyncGroup = new azure.storage.SyncGroup("exampleSyncGroup", {storageSyncId: exampleSync.id});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const exampleShare = new azure.storage.Share("exampleShare", {
 *     storageAccountName: exampleAccount.name,
 *     acls: [{
 *         id: "GhostedRecall",
 *         accessPolicies: [{
 *             permissions: "r",
 *         }],
 *     }],
 * });
 * const exampleSyncCloudEndpoint = new azure.storage.SyncCloudEndpoint("exampleSyncCloudEndpoint", {
 *     storageSyncGroupId: exampleSyncGroup.id,
 *     fileShareName: exampleShare.name,
 *     storageAccountId: exampleAccount.id,
 * });
 * ```
 *
 * ## Import
 *
 * Storage Sync Cloud Endpoints can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:storage/syncCloudEndpoint:SyncCloudEndpoint example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StorageSync/storageSyncServices/sync1/syncGroups/syncgroup1/cloudEndpoints/cloudEndpoint1
 * ```
 */
export class SyncCloudEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing SyncCloudEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SyncCloudEndpointState, opts?: pulumi.CustomResourceOptions): SyncCloudEndpoint {
        return new SyncCloudEndpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:storage/syncCloudEndpoint:SyncCloudEndpoint';

    /**
     * Returns true if the given object is an instance of SyncCloudEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SyncCloudEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SyncCloudEndpoint.__pulumiType;
    }

    /**
     * The Storage Share name to be synchronized in this Storage Sync Cloud Endpoint. Changing this forces a new Storage Sync Cloud Endpoint to be created.
     */
    public readonly fileShareName!: pulumi.Output<string>;
    /**
     * The name which should be used for this Storage Sync Cloud Endpoint. Changing this forces a new Storage Sync Cloud Endpoint to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the Storage Account where the Storage Share exists. Changing this forces a new Storage Sync Cloud Endpoint to be created.
     */
    public readonly storageAccountId!: pulumi.Output<string>;
    /**
     * The Tenant ID of the Storage Account where the Storage Share exists. Changing this forces a new Storage Sync Cloud Endpoint to be created. Defaults to the current tenant id.
     */
    public readonly storageAccountTenantId!: pulumi.Output<string>;
    /**
     * The ID of the Storage Sync Group where this Cloud Endpoint should be created. Changing this forces a new Storage Sync Cloud Endpoint to be created.
     */
    public readonly storageSyncGroupId!: pulumi.Output<string>;

    /**
     * Create a SyncCloudEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SyncCloudEndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SyncCloudEndpointArgs | SyncCloudEndpointState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SyncCloudEndpointState | undefined;
            inputs["fileShareName"] = state ? state.fileShareName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["storageAccountId"] = state ? state.storageAccountId : undefined;
            inputs["storageAccountTenantId"] = state ? state.storageAccountTenantId : undefined;
            inputs["storageSyncGroupId"] = state ? state.storageSyncGroupId : undefined;
        } else {
            const args = argsOrState as SyncCloudEndpointArgs | undefined;
            if ((!args || args.fileShareName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileShareName'");
            }
            if ((!args || args.storageAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageAccountId'");
            }
            if ((!args || args.storageSyncGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageSyncGroupId'");
            }
            inputs["fileShareName"] = args ? args.fileShareName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["storageAccountId"] = args ? args.storageAccountId : undefined;
            inputs["storageAccountTenantId"] = args ? args.storageAccountTenantId : undefined;
            inputs["storageSyncGroupId"] = args ? args.storageSyncGroupId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SyncCloudEndpoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SyncCloudEndpoint resources.
 */
export interface SyncCloudEndpointState {
    /**
     * The Storage Share name to be synchronized in this Storage Sync Cloud Endpoint. Changing this forces a new Storage Sync Cloud Endpoint to be created.
     */
    fileShareName?: pulumi.Input<string>;
    /**
     * The name which should be used for this Storage Sync Cloud Endpoint. Changing this forces a new Storage Sync Cloud Endpoint to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the Storage Account where the Storage Share exists. Changing this forces a new Storage Sync Cloud Endpoint to be created.
     */
    storageAccountId?: pulumi.Input<string>;
    /**
     * The Tenant ID of the Storage Account where the Storage Share exists. Changing this forces a new Storage Sync Cloud Endpoint to be created. Defaults to the current tenant id.
     */
    storageAccountTenantId?: pulumi.Input<string>;
    /**
     * The ID of the Storage Sync Group where this Cloud Endpoint should be created. Changing this forces a new Storage Sync Cloud Endpoint to be created.
     */
    storageSyncGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SyncCloudEndpoint resource.
 */
export interface SyncCloudEndpointArgs {
    /**
     * The Storage Share name to be synchronized in this Storage Sync Cloud Endpoint. Changing this forces a new Storage Sync Cloud Endpoint to be created.
     */
    fileShareName: pulumi.Input<string>;
    /**
     * The name which should be used for this Storage Sync Cloud Endpoint. Changing this forces a new Storage Sync Cloud Endpoint to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the Storage Account where the Storage Share exists. Changing this forces a new Storage Sync Cloud Endpoint to be created.
     */
    storageAccountId: pulumi.Input<string>;
    /**
     * The Tenant ID of the Storage Account where the Storage Share exists. Changing this forces a new Storage Sync Cloud Endpoint to be created. Defaults to the current tenant id.
     */
    storageAccountTenantId?: pulumi.Input<string>;
    /**
     * The ID of the Storage Sync Group where this Cloud Endpoint should be created. Changing this forces a new Storage Sync Cloud Endpoint to be created.
     */
    storageSyncGroupId: pulumi.Input<string>;
}

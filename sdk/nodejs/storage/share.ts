// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a File Share within Azure Storage.
 *
 * > **Note:** The storage share supports two storage tiers: premium and standard. Standard file shares are created in general purpose (GPv1 or GPv2) storage accounts and premium file shares are created in FileStorage storage accounts. For further information, refer to the section "What storage tiers are supported in Azure Files?" of [documentation](https://docs.microsoft.com/en-us/azure/storage/files/storage-files-faq#general).
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
 * const exampleShare = new azure.storage.Share("exampleShare", {
 *     storageAccountName: exampleAccount.name,
 *     quota: 50,
 *     acls: [{
 *         id: "MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI",
 *         accessPolicies: [{
 *             permissions: "rwdl",
 *             start: "2019-07-02T09:38:21.0000000Z",
 *             expiry: "2019-07-02T10:38:21.0000000Z",
 *         }],
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Storage Shares can be imported using the `resource id`, e.g.
 */
export class Share extends pulumi.CustomResource {
    /**
     * Get an existing Share resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ShareState, opts?: pulumi.CustomResourceOptions): Share {
        return new Share(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:storage/share:Share';

    /**
     * Returns true if the given object is an instance of Share.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Share {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Share.__pulumiType;
    }

    /**
     * One or more `acl` blocks as defined below.
     */
    public readonly acls!: pulumi.Output<outputs.storage.ShareAcl[] | undefined>;
    /**
     * A mapping of MetaData for this File Share.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string}>;
    /**
     * The name of the share. Must be unique within the storage account where the share is located.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The maximum size of the share, in gigabytes. For Standard storage accounts, this must be greater than 0 and less than 5120 GB (5 TB). For Premium FileStorage storage accounts, this must be greater than 100 GB and less than 102400 GB (100 TB). Default is 5120.
     */
    public readonly quota!: pulumi.Output<number | undefined>;
    /**
     * The Resource Manager ID of this File Share.
     */
    public /*out*/ readonly resourceManagerId!: pulumi.Output<string>;
    /**
     * Specifies the storage account in which to create the share.
     * Changing this forces a new resource to be created.
     */
    public readonly storageAccountName!: pulumi.Output<string>;
    /**
     * The URL of the File Share
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a Share resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ShareArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ShareArgs | ShareState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ShareState | undefined;
            inputs["acls"] = state ? state.acls : undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["quota"] = state ? state.quota : undefined;
            inputs["resourceManagerId"] = state ? state.resourceManagerId : undefined;
            inputs["storageAccountName"] = state ? state.storageAccountName : undefined;
            inputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as ShareArgs | undefined;
            if (!args || args.storageAccountName === undefined) {
                throw new Error("Missing required property 'storageAccountName'");
            }
            inputs["acls"] = args ? args.acls : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["quota"] = args ? args.quota : undefined;
            inputs["storageAccountName"] = args ? args.storageAccountName : undefined;
            inputs["resourceManagerId"] = undefined /*out*/;
            inputs["url"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Share.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Share resources.
 */
export interface ShareState {
    /**
     * One or more `acl` blocks as defined below.
     */
    readonly acls?: pulumi.Input<pulumi.Input<inputs.storage.ShareAcl>[]>;
    /**
     * A mapping of MetaData for this File Share.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the share. Must be unique within the storage account where the share is located.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The maximum size of the share, in gigabytes. For Standard storage accounts, this must be greater than 0 and less than 5120 GB (5 TB). For Premium FileStorage storage accounts, this must be greater than 100 GB and less than 102400 GB (100 TB). Default is 5120.
     */
    readonly quota?: pulumi.Input<number>;
    /**
     * The Resource Manager ID of this File Share.
     */
    readonly resourceManagerId?: pulumi.Input<string>;
    /**
     * Specifies the storage account in which to create the share.
     * Changing this forces a new resource to be created.
     */
    readonly storageAccountName?: pulumi.Input<string>;
    /**
     * The URL of the File Share
     */
    readonly url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Share resource.
 */
export interface ShareArgs {
    /**
     * One or more `acl` blocks as defined below.
     */
    readonly acls?: pulumi.Input<pulumi.Input<inputs.storage.ShareAcl>[]>;
    /**
     * A mapping of MetaData for this File Share.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the share. Must be unique within the storage account where the share is located.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The maximum size of the share, in gigabytes. For Standard storage accounts, this must be greater than 0 and less than 5120 GB (5 TB). For Premium FileStorage storage accounts, this must be greater than 100 GB and less than 102400 GB (100 TB). Default is 5120.
     */
    readonly quota?: pulumi.Input<number>;
    /**
     * Specifies the storage account in which to create the share.
     * Changing this forces a new resource to be created.
     */
    readonly storageAccountName: pulumi.Input<string>;
}

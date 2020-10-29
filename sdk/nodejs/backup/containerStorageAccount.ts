// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages registration of a storage account with Azure Backup. Storage accounts must be registered with an Azure Recovery Vault in order to backup file shares within the storage account. Registering a storage account with a vault creates what is known as a protection container within Azure Recovery Services. Once the container is created, Azure file shares within the storage account can be backed up using the `azure.backup.ProtectedFileShare` resource.
 *
 * > **NOTE:** Azure Backup for Azure File Shares is currently in public preview. During the preview, the service is subject to additional limitations and unsupported backup scenarios. [Read More](https://docs.microsoft.com/en-us/azure/backup/backup-azure-files#limitations-for-azure-file-share-backup-during-preview)
 */
export class ContainerStorageAccount extends pulumi.CustomResource {
    /**
     * Get an existing ContainerStorageAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContainerStorageAccountState, opts?: pulumi.CustomResourceOptions): ContainerStorageAccount {
        return new ContainerStorageAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:backup/containerStorageAccount:ContainerStorageAccount';

    /**
     * Returns true if the given object is an instance of ContainerStorageAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContainerStorageAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContainerStorageAccount.__pulumiType;
    }

    /**
     * The name of the vault where the storage account will be registered.
     */
    public readonly recoveryVaultName!: pulumi.Output<string>;
    /**
     * Name of the resource group where the vault is located.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The ID of the Storage Account to be registered
     */
    public readonly storageAccountId!: pulumi.Output<string>;

    /**
     * Create a ContainerStorageAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContainerStorageAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContainerStorageAccountArgs | ContainerStorageAccountState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ContainerStorageAccountState | undefined;
            inputs["recoveryVaultName"] = state ? state.recoveryVaultName : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["storageAccountId"] = state ? state.storageAccountId : undefined;
        } else {
            const args = argsOrState as ContainerStorageAccountArgs | undefined;
            if (!args || args.recoveryVaultName === undefined) {
                throw new Error("Missing required property 'recoveryVaultName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.storageAccountId === undefined) {
                throw new Error("Missing required property 'storageAccountId'");
            }
            inputs["recoveryVaultName"] = args ? args.recoveryVaultName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["storageAccountId"] = args ? args.storageAccountId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ContainerStorageAccount.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContainerStorageAccount resources.
 */
export interface ContainerStorageAccountState {
    /**
     * The name of the vault where the storage account will be registered.
     */
    readonly recoveryVaultName?: pulumi.Input<string>;
    /**
     * Name of the resource group where the vault is located.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The ID of the Storage Account to be registered
     */
    readonly storageAccountId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ContainerStorageAccount resource.
 */
export interface ContainerStorageAccountArgs {
    /**
     * The name of the vault where the storage account will be registered.
     */
    readonly recoveryVaultName: pulumi.Input<string>;
    /**
     * Name of the resource group where the vault is located.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The ID of the Storage Account to be registered
     */
    readonly storageAccountId: pulumi.Input<string>;
}

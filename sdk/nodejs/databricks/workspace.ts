// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Databricks Workspace
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleWorkspace = new azure.databricks.Workspace("exampleWorkspace", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     sku: "standard",
 *     tags: {
 *         Environment: "Production",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Databrick Workspaces can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:databricks/workspace:Workspace workspace1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Databricks/workspaces/workspace1
 * ```
 */
export class Workspace extends pulumi.CustomResource {
    /**
     * Get an existing Workspace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkspaceState, opts?: pulumi.CustomResourceOptions): Workspace {
        return new Workspace(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:databricks/workspace:Workspace';

    /**
     * Returns true if the given object is an instance of Workspace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Workspace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Workspace.__pulumiType;
    }

    /**
     * A `customParameters` block as documented below.
     */
    public readonly customParameters!: pulumi.Output<outputs.databricks.WorkspaceCustomParameters>;
    /**
     * Is the workspace enabled for customer managed key encryption? If `true` this enables the Managed Identity for the managed storage account. Possible values are `true` or `false`. Defaults to `false`. This field is only valid if the Databricks Workspace `sku` is set to `premium`. Changing this forces a new resource to be created.
     */
    public readonly customerManagedKeyEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Is the Databricks File System root file system enabled with a secondary layer of encryption with platform managed keys? Possible values are `true` or `false`. Defaults to `false`. This field is only valid if the Databricks Workspace `sku` is set to `premium`. Changing this forces a new resource to be created.
     */
    public readonly infrastructureEncryptionEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The ID of the Managed Resource Group created by the Databricks Workspace.
     */
    public /*out*/ readonly managedResourceGroupId!: pulumi.Output<string>;
    /**
     * The name of the resource group where Azure should place the managed Databricks resources. Changing this forces a new resource to be created.
     */
    public readonly managedResourceGroupName!: pulumi.Output<string>;
    /**
     * Specifies the name of the Databricks Workspace resource. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Resource Group in which the Databricks Workspace should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The `sku` to use for the Databricks Workspace. Possible values are `standard`, `premium`, or `trial`. Changing this can force a new resource to be created in some circumstances.
     */
    public readonly sku!: pulumi.Output<string>;
    /**
     * A `storageAccountIdentity` block as documented below.
     */
    public /*out*/ readonly storageAccountIdentities!: pulumi.Output<outputs.databricks.WorkspaceStorageAccountIdentity[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The unique identifier of the databricks workspace in Databricks control plane.
     */
    public /*out*/ readonly workspaceId!: pulumi.Output<string>;
    /**
     * The workspace URL which is of the format 'adb-{workspaceId}.{random}.azuredatabricks.net'
     */
    public /*out*/ readonly workspaceUrl!: pulumi.Output<string>;

    /**
     * Create a Workspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkspaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkspaceArgs | WorkspaceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WorkspaceState | undefined;
            inputs["customParameters"] = state ? state.customParameters : undefined;
            inputs["customerManagedKeyEnabled"] = state ? state.customerManagedKeyEnabled : undefined;
            inputs["infrastructureEncryptionEnabled"] = state ? state.infrastructureEncryptionEnabled : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["managedResourceGroupId"] = state ? state.managedResourceGroupId : undefined;
            inputs["managedResourceGroupName"] = state ? state.managedResourceGroupName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["storageAccountIdentities"] = state ? state.storageAccountIdentities : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["workspaceId"] = state ? state.workspaceId : undefined;
            inputs["workspaceUrl"] = state ? state.workspaceUrl : undefined;
        } else {
            const args = argsOrState as WorkspaceArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.sku === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["customParameters"] = args ? args.customParameters : undefined;
            inputs["customerManagedKeyEnabled"] = args ? args.customerManagedKeyEnabled : undefined;
            inputs["infrastructureEncryptionEnabled"] = args ? args.infrastructureEncryptionEnabled : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["managedResourceGroupName"] = args ? args.managedResourceGroupName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["managedResourceGroupId"] = undefined /*out*/;
            inputs["storageAccountIdentities"] = undefined /*out*/;
            inputs["workspaceId"] = undefined /*out*/;
            inputs["workspaceUrl"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Workspace.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Workspace resources.
 */
export interface WorkspaceState {
    /**
     * A `customParameters` block as documented below.
     */
    customParameters?: pulumi.Input<inputs.databricks.WorkspaceCustomParameters>;
    /**
     * Is the workspace enabled for customer managed key encryption? If `true` this enables the Managed Identity for the managed storage account. Possible values are `true` or `false`. Defaults to `false`. This field is only valid if the Databricks Workspace `sku` is set to `premium`. Changing this forces a new resource to be created.
     */
    customerManagedKeyEnabled?: pulumi.Input<boolean>;
    /**
     * Is the Databricks File System root file system enabled with a secondary layer of encryption with platform managed keys? Possible values are `true` or `false`. Defaults to `false`. This field is only valid if the Databricks Workspace `sku` is set to `premium`. Changing this forces a new resource to be created.
     */
    infrastructureEncryptionEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The ID of the Managed Resource Group created by the Databricks Workspace.
     */
    managedResourceGroupId?: pulumi.Input<string>;
    /**
     * The name of the resource group where Azure should place the managed Databricks resources. Changing this forces a new resource to be created.
     */
    managedResourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies the name of the Databricks Workspace resource. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the Databricks Workspace should exist. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The `sku` to use for the Databricks Workspace. Possible values are `standard`, `premium`, or `trial`. Changing this can force a new resource to be created in some circumstances.
     */
    sku?: pulumi.Input<string>;
    /**
     * A `storageAccountIdentity` block as documented below.
     */
    storageAccountIdentities?: pulumi.Input<pulumi.Input<inputs.databricks.WorkspaceStorageAccountIdentity>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The unique identifier of the databricks workspace in Databricks control plane.
     */
    workspaceId?: pulumi.Input<string>;
    /**
     * The workspace URL which is of the format 'adb-{workspaceId}.{random}.azuredatabricks.net'
     */
    workspaceUrl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Workspace resource.
 */
export interface WorkspaceArgs {
    /**
     * A `customParameters` block as documented below.
     */
    customParameters?: pulumi.Input<inputs.databricks.WorkspaceCustomParameters>;
    /**
     * Is the workspace enabled for customer managed key encryption? If `true` this enables the Managed Identity for the managed storage account. Possible values are `true` or `false`. Defaults to `false`. This field is only valid if the Databricks Workspace `sku` is set to `premium`. Changing this forces a new resource to be created.
     */
    customerManagedKeyEnabled?: pulumi.Input<boolean>;
    /**
     * Is the Databricks File System root file system enabled with a secondary layer of encryption with platform managed keys? Possible values are `true` or `false`. Defaults to `false`. This field is only valid if the Databricks Workspace `sku` is set to `premium`. Changing this forces a new resource to be created.
     */
    infrastructureEncryptionEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the resource group where Azure should place the managed Databricks resources. Changing this forces a new resource to be created.
     */
    managedResourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies the name of the Databricks Workspace resource. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the Databricks Workspace should exist. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The `sku` to use for the Databricks Workspace. Possible values are `standard`, `premium`, or `trial`. Changing this can force a new resource to be created in some circumstances.
     */
    sku: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

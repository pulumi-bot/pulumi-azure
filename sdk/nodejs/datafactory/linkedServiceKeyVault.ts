// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Linked Service (connection) between Key Vault and Azure Data Factory.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const current = azure.core.getClientConfig({});
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleKeyVault = new azure.keyvault.KeyVault("exampleKeyVault", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     tenantId: current.then(current => current.tenantId),
 *     skuName: "standard",
 * });
 * const exampleFactory = new azure.datafactory.Factory("exampleFactory", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleLinkedServiceKeyVault = new azure.datafactory.LinkedServiceKeyVault("exampleLinkedServiceKeyVault", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     dataFactoryName: exampleFactory.name,
 *     keyVaultId: exampleKeyVault.id,
 * });
 * ```
 *
 * ## Import
 *
 * Data Factory Key Vault Linked Service's can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:datafactory/linkedServiceKeyVault:LinkedServiceKeyVault example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
 * ```
 */
export class LinkedServiceKeyVault extends pulumi.CustomResource {
    /**
     * Get an existing LinkedServiceKeyVault resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LinkedServiceKeyVaultState, opts?: pulumi.CustomResourceOptions): LinkedServiceKeyVault {
        return new LinkedServiceKeyVault(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:datafactory/linkedServiceKeyVault:LinkedServiceKeyVault';

    /**
     * Returns true if the given object is an instance of LinkedServiceKeyVault.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LinkedServiceKeyVault {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LinkedServiceKeyVault.__pulumiType;
    }

    /**
     * A map of additional properties to associate with the Data Factory Linked Service Key Vault.
     */
    public readonly additionalProperties!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * List of tags that can be used for describing the Data Factory Linked Service Key Vault.
     */
    public readonly annotations!: pulumi.Output<string[] | undefined>;
    /**
     * The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
     */
    public readonly dataFactoryName!: pulumi.Output<string>;
    /**
     * The description for the Data Factory Linked Service Key Vault.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The integration runtime reference to associate with the Data Factory Linked Service Key Vault.
     */
    public readonly integrationRuntimeName!: pulumi.Output<string | undefined>;
    /**
     * The ID the Azure Key Vault resource.
     */
    public readonly keyVaultId!: pulumi.Output<string>;
    /**
     * Specifies the name of the Data Factory Linked Service Key Vault. Changing this forces a new resource to be created. Must be unique within a data
     * factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A map of parameters to associate with the Data Factory Linked Service Key Vault.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the resource group in which to create the Data Factory Linked Service Key Vault. Changing this forces a new resource
     */
    public readonly resourceGroupName!: pulumi.Output<string>;

    /**
     * Create a LinkedServiceKeyVault resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LinkedServiceKeyVaultArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LinkedServiceKeyVaultArgs | LinkedServiceKeyVaultState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LinkedServiceKeyVaultState | undefined;
            inputs["additionalProperties"] = state ? state.additionalProperties : undefined;
            inputs["annotations"] = state ? state.annotations : undefined;
            inputs["dataFactoryName"] = state ? state.dataFactoryName : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["integrationRuntimeName"] = state ? state.integrationRuntimeName : undefined;
            inputs["keyVaultId"] = state ? state.keyVaultId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
        } else {
            const args = argsOrState as LinkedServiceKeyVaultArgs | undefined;
            if ((!args || args.dataFactoryName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataFactoryName'");
            }
            if ((!args || args.keyVaultId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyVaultId'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["additionalProperties"] = args ? args.additionalProperties : undefined;
            inputs["annotations"] = args ? args.annotations : undefined;
            inputs["dataFactoryName"] = args ? args.dataFactoryName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["integrationRuntimeName"] = args ? args.integrationRuntimeName : undefined;
            inputs["keyVaultId"] = args ? args.keyVaultId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(LinkedServiceKeyVault.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LinkedServiceKeyVault resources.
 */
export interface LinkedServiceKeyVaultState {
    /**
     * A map of additional properties to associate with the Data Factory Linked Service Key Vault.
     */
    additionalProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of tags that can be used for describing the Data Factory Linked Service Key Vault.
     */
    annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
     */
    dataFactoryName?: pulumi.Input<string>;
    /**
     * The description for the Data Factory Linked Service Key Vault.
     */
    description?: pulumi.Input<string>;
    /**
     * The integration runtime reference to associate with the Data Factory Linked Service Key Vault.
     */
    integrationRuntimeName?: pulumi.Input<string>;
    /**
     * The ID the Azure Key Vault resource.
     */
    keyVaultId?: pulumi.Input<string>;
    /**
     * Specifies the name of the Data Factory Linked Service Key Vault. Changing this forces a new resource to be created. Must be unique within a data
     * factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    name?: pulumi.Input<string>;
    /**
     * A map of parameters to associate with the Data Factory Linked Service Key Vault.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the resource group in which to create the Data Factory Linked Service Key Vault. Changing this forces a new resource
     */
    resourceGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LinkedServiceKeyVault resource.
 */
export interface LinkedServiceKeyVaultArgs {
    /**
     * A map of additional properties to associate with the Data Factory Linked Service Key Vault.
     */
    additionalProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of tags that can be used for describing the Data Factory Linked Service Key Vault.
     */
    annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
     */
    dataFactoryName: pulumi.Input<string>;
    /**
     * The description for the Data Factory Linked Service Key Vault.
     */
    description?: pulumi.Input<string>;
    /**
     * The integration runtime reference to associate with the Data Factory Linked Service Key Vault.
     */
    integrationRuntimeName?: pulumi.Input<string>;
    /**
     * The ID the Azure Key Vault resource.
     */
    keyVaultId: pulumi.Input<string>;
    /**
     * Specifies the name of the Data Factory Linked Service Key Vault. Changing this forces a new resource to be created. Must be unique within a data
     * factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    name?: pulumi.Input<string>;
    /**
     * A map of parameters to associate with the Data Factory Linked Service Key Vault.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the resource group in which to create the Data Factory Linked Service Key Vault. Changing this forces a new resource
     */
    resourceGroupName: pulumi.Input<string>;
}

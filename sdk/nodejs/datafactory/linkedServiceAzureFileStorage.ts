// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Linked Service (connection) between a SFTP Server and Azure Data Factory.
 *
 * > **Note:** All arguments including the client secret will be stored in the raw state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "northeurope"});
 * const exampleAccount = exampleResourceGroup.name.apply(name => azure.storage.getAccount({
 *     name: "storageaccountname",
 *     resourceGroupName: name,
 * }));
 * const exampleFactory = new azure.datafactory.Factory("exampleFactory", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleLinkedServiceAzureFileStorage = new azure.datafactory.LinkedServiceAzureFileStorage("exampleLinkedServiceAzureFileStorage", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     dataFactoryName: exampleFactory.name,
 *     connectionString: exampleAccount.primaryConnectionString,
 * });
 * ```
 *
 * ## Import
 *
 * Data Factory Linked Service's can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:datafactory/linkedServiceAzureFileStorage:LinkedServiceAzureFileStorage example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
 * ```
 */
export class LinkedServiceAzureFileStorage extends pulumi.CustomResource {
    /**
     * Get an existing LinkedServiceAzureFileStorage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LinkedServiceAzureFileStorageState, opts?: pulumi.CustomResourceOptions): LinkedServiceAzureFileStorage {
        return new LinkedServiceAzureFileStorage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:datafactory/linkedServiceAzureFileStorage:LinkedServiceAzureFileStorage';

    /**
     * Returns true if the given object is an instance of LinkedServiceAzureFileStorage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LinkedServiceAzureFileStorage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LinkedServiceAzureFileStorage.__pulumiType;
    }

    /**
     * A map of additional properties to associate with the Data Factory Linked Service.
     */
    public readonly additionalProperties!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * List of tags that can be used for describing the Data Factory Linked Service.
     */
    public readonly annotations!: pulumi.Output<string[] | undefined>;
    /**
     * The connection string.
     */
    public readonly connectionString!: pulumi.Output<string>;
    /**
     * The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
     */
    public readonly dataFactoryName!: pulumi.Output<string>;
    /**
     * The description for the Data Factory Linked Service.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly host!: pulumi.Output<string | undefined>;
    /**
     * The integration runtime reference to associate with the Data Factory Linked Service.
     */
    public readonly integrationRuntimeName!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A map of parameters to associate with the Data Factory Linked Service.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    public readonly userId!: pulumi.Output<string | undefined>;

    /**
     * Create a LinkedServiceAzureFileStorage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LinkedServiceAzureFileStorageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LinkedServiceAzureFileStorageArgs | LinkedServiceAzureFileStorageState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LinkedServiceAzureFileStorageState | undefined;
            inputs["additionalProperties"] = state ? state.additionalProperties : undefined;
            inputs["annotations"] = state ? state.annotations : undefined;
            inputs["connectionString"] = state ? state.connectionString : undefined;
            inputs["dataFactoryName"] = state ? state.dataFactoryName : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["host"] = state ? state.host : undefined;
            inputs["integrationRuntimeName"] = state ? state.integrationRuntimeName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as LinkedServiceAzureFileStorageArgs | undefined;
            if (!args || args.connectionString === undefined) {
                throw new Error("Missing required property 'connectionString'");
            }
            if (!args || args.dataFactoryName === undefined) {
                throw new Error("Missing required property 'dataFactoryName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["additionalProperties"] = args ? args.additionalProperties : undefined;
            inputs["annotations"] = args ? args.annotations : undefined;
            inputs["connectionString"] = args ? args.connectionString : undefined;
            inputs["dataFactoryName"] = args ? args.dataFactoryName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["host"] = args ? args.host : undefined;
            inputs["integrationRuntimeName"] = args ? args.integrationRuntimeName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["userId"] = args ? args.userId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LinkedServiceAzureFileStorage.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LinkedServiceAzureFileStorage resources.
 */
export interface LinkedServiceAzureFileStorageState {
    /**
     * A map of additional properties to associate with the Data Factory Linked Service.
     */
    readonly additionalProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of tags that can be used for describing the Data Factory Linked Service.
     */
    readonly annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The connection string.
     */
    readonly connectionString?: pulumi.Input<string>;
    /**
     * The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
     */
    readonly dataFactoryName?: pulumi.Input<string>;
    /**
     * The description for the Data Factory Linked Service.
     */
    readonly description?: pulumi.Input<string>;
    readonly host?: pulumi.Input<string>;
    /**
     * The integration runtime reference to associate with the Data Factory Linked Service.
     */
    readonly integrationRuntimeName?: pulumi.Input<string>;
    /**
     * Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A map of parameters to associate with the Data Factory Linked Service.
     */
    readonly parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly password?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    readonly userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LinkedServiceAzureFileStorage resource.
 */
export interface LinkedServiceAzureFileStorageArgs {
    /**
     * A map of additional properties to associate with the Data Factory Linked Service.
     */
    readonly additionalProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of tags that can be used for describing the Data Factory Linked Service.
     */
    readonly annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The connection string.
     */
    readonly connectionString: pulumi.Input<string>;
    /**
     * The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
     */
    readonly dataFactoryName: pulumi.Input<string>;
    /**
     * The description for the Data Factory Linked Service.
     */
    readonly description?: pulumi.Input<string>;
    readonly host?: pulumi.Input<string>;
    /**
     * The integration runtime reference to associate with the Data Factory Linked Service.
     */
    readonly integrationRuntimeName?: pulumi.Input<string>;
    /**
     * Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A map of parameters to associate with the Data Factory Linked Service.
     */
    readonly parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly password?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
     */
    readonly resourceGroupName: pulumi.Input<string>;
    readonly userId?: pulumi.Input<string>;
}

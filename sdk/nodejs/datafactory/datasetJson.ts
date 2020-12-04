// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an Azure JSON Dataset inside an Azure Data Factory.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "northeurope"});
 * const exampleFactory = new azure.datafactory.Factory("exampleFactory", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleLinkedServiceWeb = new azure.datafactory.LinkedServiceWeb("exampleLinkedServiceWeb", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     dataFactoryName: exampleFactory.name,
 *     authenticationType: "Anonymous",
 *     url: "https://www.bing.com",
 * });
 * const exampleDatasetJson = new azure.datafactory.DatasetJson("exampleDatasetJson", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     dataFactoryName: exampleFactory.name,
 *     linkedServiceName: exampleLinkedServiceWeb.name,
 *     httpServerLocation: {
 *         relativeUrl: "/fizz/buzz/",
 *         path: "foo/bar/",
 *         filename: "foo.txt",
 *     },
 *     encoding: "UTF-8",
 * });
 * ```
 *
 * ## Import
 *
 * Data Factory Datasets can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:datafactory/datasetJson:DatasetJson example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
 * ```
 */
export class DatasetJson extends pulumi.CustomResource {
    /**
     * Get an existing DatasetJson resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatasetJsonState, opts?: pulumi.CustomResourceOptions): DatasetJson {
        return new DatasetJson(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:datafactory/datasetJson:DatasetJson';

    /**
     * Returns true if the given object is an instance of DatasetJson.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatasetJson {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatasetJson.__pulumiType;
    }

    /**
     * A map of additional properties to associate with the Data Factory Dataset.
     */
    public readonly additionalProperties!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * List of tags that can be used for describing the Data Factory Dataset.
     */
    public readonly annotations!: pulumi.Output<string[] | undefined>;
    /**
     * A `azureBlobStorageLocation` block as defined below.
     */
    public readonly azureBlobStorageLocation!: pulumi.Output<outputs.datafactory.DatasetJsonAzureBlobStorageLocation | undefined>;
    /**
     * The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
     */
    public readonly dataFactoryName!: pulumi.Output<string>;
    /**
     * The description for the Data Factory Dataset.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The encoding format for the file.
     */
    public readonly encoding!: pulumi.Output<string | undefined>;
    /**
     * The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
     */
    public readonly folder!: pulumi.Output<string | undefined>;
    /**
     * A `httpServerLocation` block as defined below.
     */
    public readonly httpServerLocation!: pulumi.Output<outputs.datafactory.DatasetJsonHttpServerLocation | undefined>;
    /**
     * The Data Factory Linked Service name in which to associate the Dataset with.
     */
    public readonly linkedServiceName!: pulumi.Output<string>;
    /**
     * Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A map of parameters to associate with the Data Factory Dataset.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `schemaColumn` block as defined below.
     */
    public readonly schemaColumns!: pulumi.Output<outputs.datafactory.DatasetJsonSchemaColumn[] | undefined>;

    /**
     * Create a DatasetJson resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatasetJsonArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatasetJsonArgs | DatasetJsonState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DatasetJsonState | undefined;
            inputs["additionalProperties"] = state ? state.additionalProperties : undefined;
            inputs["annotations"] = state ? state.annotations : undefined;
            inputs["azureBlobStorageLocation"] = state ? state.azureBlobStorageLocation : undefined;
            inputs["dataFactoryName"] = state ? state.dataFactoryName : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["encoding"] = state ? state.encoding : undefined;
            inputs["folder"] = state ? state.folder : undefined;
            inputs["httpServerLocation"] = state ? state.httpServerLocation : undefined;
            inputs["linkedServiceName"] = state ? state.linkedServiceName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["schemaColumns"] = state ? state.schemaColumns : undefined;
        } else {
            const args = argsOrState as DatasetJsonArgs | undefined;
            if ((!args || args.dataFactoryName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'dataFactoryName'");
            }
            if ((!args || args.linkedServiceName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'linkedServiceName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["additionalProperties"] = args ? args.additionalProperties : undefined;
            inputs["annotations"] = args ? args.annotations : undefined;
            inputs["azureBlobStorageLocation"] = args ? args.azureBlobStorageLocation : undefined;
            inputs["dataFactoryName"] = args ? args.dataFactoryName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["encoding"] = args ? args.encoding : undefined;
            inputs["folder"] = args ? args.folder : undefined;
            inputs["httpServerLocation"] = args ? args.httpServerLocation : undefined;
            inputs["linkedServiceName"] = args ? args.linkedServiceName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["schemaColumns"] = args ? args.schemaColumns : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DatasetJson.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatasetJson resources.
 */
export interface DatasetJsonState {
    /**
     * A map of additional properties to associate with the Data Factory Dataset.
     */
    readonly additionalProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of tags that can be used for describing the Data Factory Dataset.
     */
    readonly annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A `azureBlobStorageLocation` block as defined below.
     */
    readonly azureBlobStorageLocation?: pulumi.Input<inputs.datafactory.DatasetJsonAzureBlobStorageLocation>;
    /**
     * The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
     */
    readonly dataFactoryName?: pulumi.Input<string>;
    /**
     * The description for the Data Factory Dataset.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The encoding format for the file.
     */
    readonly encoding?: pulumi.Input<string>;
    /**
     * The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * A `httpServerLocation` block as defined below.
     */
    readonly httpServerLocation?: pulumi.Input<inputs.datafactory.DatasetJsonHttpServerLocation>;
    /**
     * The Data Factory Linked Service name in which to associate the Dataset with.
     */
    readonly linkedServiceName?: pulumi.Input<string>;
    /**
     * Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A map of parameters to associate with the Data Factory Dataset.
     */
    readonly parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A `schemaColumn` block as defined below.
     */
    readonly schemaColumns?: pulumi.Input<pulumi.Input<inputs.datafactory.DatasetJsonSchemaColumn>[]>;
}

/**
 * The set of arguments for constructing a DatasetJson resource.
 */
export interface DatasetJsonArgs {
    /**
     * A map of additional properties to associate with the Data Factory Dataset.
     */
    readonly additionalProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of tags that can be used for describing the Data Factory Dataset.
     */
    readonly annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A `azureBlobStorageLocation` block as defined below.
     */
    readonly azureBlobStorageLocation?: pulumi.Input<inputs.datafactory.DatasetJsonAzureBlobStorageLocation>;
    /**
     * The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
     */
    readonly dataFactoryName: pulumi.Input<string>;
    /**
     * The description for the Data Factory Dataset.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The encoding format for the file.
     */
    readonly encoding?: pulumi.Input<string>;
    /**
     * The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * A `httpServerLocation` block as defined below.
     */
    readonly httpServerLocation?: pulumi.Input<inputs.datafactory.DatasetJsonHttpServerLocation>;
    /**
     * The Data Factory Linked Service name in which to associate the Dataset with.
     */
    readonly linkedServiceName: pulumi.Input<string>;
    /**
     * Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A map of parameters to associate with the Data Factory Dataset.
     */
    readonly parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `schemaColumn` block as defined below.
     */
    readonly schemaColumns?: pulumi.Input<pulumi.Input<inputs.datafactory.DatasetJsonSchemaColumn>[]>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Linked Service (connection) between a Web Server and Azure Data Factory.
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
 * const exampleFactory = new azure.datafactory.Factory("exampleFactory", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleLinkedServiceWeb = new azure.datafactory.LinkedServiceWeb("exampleLinkedServiceWeb", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     dataFactoryName: exampleFactory.name,
 *     authenticationType: "Anonymous",
 *     url: "http://www.bing.com",
 * });
 * ```
 *
 * ## Import
 *
 * Data Factory Linked Service's can be imported using the `resource id`, e.g.
 */
export class LinkedServiceWeb extends pulumi.CustomResource {
    /**
     * Get an existing LinkedServiceWeb resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LinkedServiceWebState, opts?: pulumi.CustomResourceOptions): LinkedServiceWeb {
        return new LinkedServiceWeb(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:datafactory/linkedServiceWeb:LinkedServiceWeb';

    /**
     * Returns true if the given object is an instance of LinkedServiceWeb.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LinkedServiceWeb {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LinkedServiceWeb.__pulumiType;
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
     * The type of authentication used to connect to the web table source. Valid options are `Anonymous`, `Basic` and `ClientCertificate`.
     */
    public readonly authenticationType!: pulumi.Output<string>;
    /**
     * The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
     */
    public readonly dataFactoryName!: pulumi.Output<string>;
    /**
     * The description for the Data Factory Linked Service.
     */
    public readonly description!: pulumi.Output<string | undefined>;
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
    /**
     * The URL of the web service endpoint (e.g. http://www.microsoft.com).
     */
    public readonly url!: pulumi.Output<string>;
    public readonly username!: pulumi.Output<string | undefined>;

    /**
     * Create a LinkedServiceWeb resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LinkedServiceWebArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LinkedServiceWebArgs | LinkedServiceWebState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LinkedServiceWebState | undefined;
            inputs["additionalProperties"] = state ? state.additionalProperties : undefined;
            inputs["annotations"] = state ? state.annotations : undefined;
            inputs["authenticationType"] = state ? state.authenticationType : undefined;
            inputs["dataFactoryName"] = state ? state.dataFactoryName : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["integrationRuntimeName"] = state ? state.integrationRuntimeName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as LinkedServiceWebArgs | undefined;
            if (!args || args.authenticationType === undefined) {
                throw new Error("Missing required property 'authenticationType'");
            }
            if (!args || args.dataFactoryName === undefined) {
                throw new Error("Missing required property 'dataFactoryName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.url === undefined) {
                throw new Error("Missing required property 'url'");
            }
            inputs["additionalProperties"] = args ? args.additionalProperties : undefined;
            inputs["annotations"] = args ? args.annotations : undefined;
            inputs["authenticationType"] = args ? args.authenticationType : undefined;
            inputs["dataFactoryName"] = args ? args.dataFactoryName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["integrationRuntimeName"] = args ? args.integrationRuntimeName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["url"] = args ? args.url : undefined;
            inputs["username"] = args ? args.username : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LinkedServiceWeb.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LinkedServiceWeb resources.
 */
export interface LinkedServiceWebState {
    /**
     * A map of additional properties to associate with the Data Factory Linked Service.
     */
    readonly additionalProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of tags that can be used for describing the Data Factory Linked Service.
     */
    readonly annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of authentication used to connect to the web table source. Valid options are `Anonymous`, `Basic` and `ClientCertificate`.
     */
    readonly authenticationType?: pulumi.Input<string>;
    /**
     * The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
     */
    readonly dataFactoryName?: pulumi.Input<string>;
    /**
     * The description for the Data Factory Linked Service.
     */
    readonly description?: pulumi.Input<string>;
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
    /**
     * The URL of the web service endpoint (e.g. http://www.microsoft.com).
     */
    readonly url?: pulumi.Input<string>;
    readonly username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LinkedServiceWeb resource.
 */
export interface LinkedServiceWebArgs {
    /**
     * A map of additional properties to associate with the Data Factory Linked Service.
     */
    readonly additionalProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of tags that can be used for describing the Data Factory Linked Service.
     */
    readonly annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of authentication used to connect to the web table source. Valid options are `Anonymous`, `Basic` and `ClientCertificate`.
     */
    readonly authenticationType: pulumi.Input<string>;
    /**
     * The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
     */
    readonly dataFactoryName: pulumi.Input<string>;
    /**
     * The description for the Data Factory Linked Service.
     */
    readonly description?: pulumi.Input<string>;
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
    /**
     * The URL of the web service endpoint (e.g. http://www.microsoft.com).
     */
    readonly url: pulumi.Input<string>;
    readonly username?: pulumi.Input<string>;
}

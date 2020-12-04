// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an IoT Central Application
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleApplication = new azure.iotcentral.Application("exampleApplication", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     subDomain: "example-iotcentral-app-subdomain",
 *     displayName: "example-iotcentral-app-display-name",
 *     sku: "S1",
 *     template: "iotc-default@1.0.0",
 *     tags: {
 *         Foo: "Bar",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * The IoT Central Application can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:iotcentral/application:Application example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.IoTCentral/IoTApps/app1
 * ```
 */
export class Application extends pulumi.CustomResource {
    /**
     * Get an existing Application resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationState, opts?: pulumi.CustomResourceOptions): Application {
        return new Application(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:iotcentral/application:Application';

    /**
     * Returns true if the given object is an instance of Application.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Application {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Application.__pulumiType;
    }

    /**
     * A `displayName` name. Custom display name for the IoT Central application. Default is resource name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Specifies the supported Azure location where the resource has to be create. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `sku` name. Possible values is `ST1`, `ST2`, Default value is `ST1`
     */
    public readonly sku!: pulumi.Output<string | undefined>;
    /**
     * A `subDomain` name. Subdomain for the IoT Central URL. Each application must have a unique subdomain.
     */
    public readonly subDomain!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A `template` name. IoT Central application template name. Default is a custom application.
     */
    public readonly template!: pulumi.Output<string>;

    /**
     * Create a Application resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationArgs | ApplicationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ApplicationState | undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["subDomain"] = state ? state.subDomain : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["template"] = state ? state.template : undefined;
        } else {
            const args = argsOrState as ApplicationArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.subDomain === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'subDomain'");
            }
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["subDomain"] = args ? args.subDomain : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["template"] = args ? args.template : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Application.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Application resources.
 */
export interface ApplicationState {
    /**
     * A `displayName` name. Custom display name for the IoT Central application. Default is resource name.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource has to be create. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A `sku` name. Possible values is `ST1`, `ST2`, Default value is `ST1`
     */
    readonly sku?: pulumi.Input<string>;
    /**
     * A `subDomain` name. Subdomain for the IoT Central URL. Each application must have a unique subdomain.
     */
    readonly subDomain?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A `template` name. IoT Central application template name. Default is a custom application.
     */
    readonly template?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Application resource.
 */
export interface ApplicationArgs {
    /**
     * A `displayName` name. Custom display name for the IoT Central application. Default is resource name.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource has to be create. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `sku` name. Possible values is `ST1`, `ST2`, Default value is `ST1`
     */
    readonly sku?: pulumi.Input<string>;
    /**
     * A `subDomain` name. Subdomain for the IoT Central URL. Each application must have a unique subdomain.
     */
    readonly subDomain: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A `template` name. IoT Central application template name. Default is a custom application.
     */
    readonly template?: pulumi.Input<string>;
}

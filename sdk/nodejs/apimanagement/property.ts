// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an API Management Property.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US"});
 * const exampleService = new azure.apimanagement.Service("exampleService", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     publisherName: "pub1",
 *     publisherEmail: "pub1@email.com",
 *     skuName: "Developer_1",
 * });
 * const exampleProperty = new azure.apimanagement.Property("exampleProperty", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     apiManagementName: exampleService.name,
 *     displayName: "ExampleProperty",
 *     value: "Example Value",
 * });
 * ```
 *
 * ## Import
 *
 * API Management Properties can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:apimanagement/property:Property example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.ApiManagement/service/example-apim/namedValues/example-apimp
 * ```
 */
export class Property extends pulumi.CustomResource {
    /**
     * Get an existing Property resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PropertyState, opts?: pulumi.CustomResourceOptions): Property {
        return new Property(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:apimanagement/property:Property';

    /**
     * Returns true if the given object is an instance of Property.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Property {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Property.__pulumiType;
    }

    /**
     * The name of the API Management Service in which the API Management Property should exist. Changing this forces a new resource to be created.
     */
    public readonly apiManagementName!: pulumi.Output<string>;
    /**
     * The display name of this API Management Property.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The name of the API Management Property. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Resource Group in which the API Management Property should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Specifies whether the API Management Property is secret. Valid values are `true` or `false`. The default value is `false`.
     */
    public readonly secret!: pulumi.Output<boolean | undefined>;
    /**
     * A list of tags to be applied to the API Management Property.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The value of this API Management Property.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a Property resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PropertyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PropertyArgs | PropertyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PropertyState | undefined;
            inputs["apiManagementName"] = state ? state.apiManagementName : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["secret"] = state ? state.secret : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as PropertyArgs | undefined;
            if ((!args || args.apiManagementName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'apiManagementName'");
            }
            if ((!args || args.displayName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.value === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'value'");
            }
            inputs["apiManagementName"] = args ? args.apiManagementName : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["secret"] = args ? args.secret : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["value"] = args ? args.value : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Property.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Property resources.
 */
export interface PropertyState {
    /**
     * The name of the API Management Service in which the API Management Property should exist. Changing this forces a new resource to be created.
     */
    readonly apiManagementName?: pulumi.Input<string>;
    /**
     * The display name of this API Management Property.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The name of the API Management Property. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the API Management Property should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies whether the API Management Property is secret. Valid values are `true` or `false`. The default value is `false`.
     */
    readonly secret?: pulumi.Input<boolean>;
    /**
     * A list of tags to be applied to the API Management Property.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The value of this API Management Property.
     */
    readonly value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Property resource.
 */
export interface PropertyArgs {
    /**
     * The name of the API Management Service in which the API Management Property should exist. Changing this forces a new resource to be created.
     */
    readonly apiManagementName: pulumi.Input<string>;
    /**
     * The display name of this API Management Property.
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * The name of the API Management Property. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the API Management Property should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies whether the API Management Property is secret. Valid values are `true` or `false`. The default value is `false`.
     */
    readonly secret?: pulumi.Input<boolean>;
    /**
     * A list of tags to be applied to the API Management Property.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The value of this API Management Property.
     */
    readonly value: pulumi.Input<string>;
}

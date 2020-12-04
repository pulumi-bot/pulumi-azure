// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an API Management Microsoft Identity Provider.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleService = new azure.apimanagement.Service("exampleService", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     publisherName: "My Company",
 *     publisherEmail: "company@mycompany.io",
 *     skuName: "Developer_1",
 * });
 * const exampleIdentityProviderMicrosoft = new azure.apimanagement.IdentityProviderMicrosoft("exampleIdentityProviderMicrosoft", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     apiManagementName: exampleService.name,
 *     clientId: "00000000-0000-0000-0000-000000000000",
 *     clientSecret: "00000000000000000000000000000000",
 * });
 * ```
 *
 * ## Import
 *
 * API Management Microsoft Identity Provider can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:apimanagement/identityProviderMicrosoft:IdentityProviderMicrosoft example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/microsoft
 * ```
 */
export class IdentityProviderMicrosoft extends pulumi.CustomResource {
    /**
     * Get an existing IdentityProviderMicrosoft resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IdentityProviderMicrosoftState, opts?: pulumi.CustomResourceOptions): IdentityProviderMicrosoft {
        return new IdentityProviderMicrosoft(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:apimanagement/identityProviderMicrosoft:IdentityProviderMicrosoft';

    /**
     * Returns true if the given object is an instance of IdentityProviderMicrosoft.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IdentityProviderMicrosoft {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IdentityProviderMicrosoft.__pulumiType;
    }

    /**
     * The Name of the API Management Service where this Microsoft Identity Provider should be created. Changing this forces a new resource to be created.
     */
    public readonly apiManagementName!: pulumi.Output<string>;
    /**
     * Client Id of the Azure AD Application.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * Client secret of the Azure AD Application.
     */
    public readonly clientSecret!: pulumi.Output<string>;
    /**
     * The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;

    /**
     * Create a IdentityProviderMicrosoft resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IdentityProviderMicrosoftArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IdentityProviderMicrosoftArgs | IdentityProviderMicrosoftState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IdentityProviderMicrosoftState | undefined;
            inputs["apiManagementName"] = state ? state.apiManagementName : undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientSecret"] = state ? state.clientSecret : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
        } else {
            const args = argsOrState as IdentityProviderMicrosoftArgs | undefined;
            if ((!args || args.apiManagementName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'apiManagementName'");
            }
            if ((!args || args.clientId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.clientSecret === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'clientSecret'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["apiManagementName"] = args ? args.apiManagementName : undefined;
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientSecret"] = args ? args.clientSecret : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IdentityProviderMicrosoft.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IdentityProviderMicrosoft resources.
 */
export interface IdentityProviderMicrosoftState {
    /**
     * The Name of the API Management Service where this Microsoft Identity Provider should be created. Changing this forces a new resource to be created.
     */
    readonly apiManagementName?: pulumi.Input<string>;
    /**
     * Client Id of the Azure AD Application.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * Client secret of the Azure AD Application.
     */
    readonly clientSecret?: pulumi.Input<string>;
    /**
     * The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IdentityProviderMicrosoft resource.
 */
export interface IdentityProviderMicrosoftArgs {
    /**
     * The Name of the API Management Service where this Microsoft Identity Provider should be created. Changing this forces a new resource to be created.
     */
    readonly apiManagementName: pulumi.Input<string>;
    /**
     * Client Id of the Azure AD Application.
     */
    readonly clientId: pulumi.Input<string>;
    /**
     * Client secret of the Azure AD Application.
     */
    readonly clientSecret: pulumi.Input<string>;
    /**
     * The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}

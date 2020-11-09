// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an API Operation within an API Management Service.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleApi = azure.apimanagement.getApi({
 *     name: "search-api",
 *     apiManagementName: "search-api-management",
 *     resourceGroupName: "search-service",
 *     revision: "2",
 * });
 * const exampleApiOperation = new azure.apimanagement.ApiOperation("exampleApiOperation", {
 *     operationId: "user-delete",
 *     apiName: exampleApi.then(exampleApi => exampleApi.name),
 *     apiManagementName: exampleApi.then(exampleApi => exampleApi.apiManagementName),
 *     resourceGroupName: exampleApi.then(exampleApi => exampleApi.resourceGroupName),
 *     displayName: "Delete User Operation",
 *     method: "DELETE",
 *     urlTemplate: "/users/{id}/delete",
 *     description: "This can only be done by the logged in user.",
 *     responses: [{
 *         statusCode: 200,
 *     }],
 * });
 * ```
 */
export class ApiOperation extends pulumi.CustomResource {
    /**
     * Get an existing ApiOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiOperationState, opts?: pulumi.CustomResourceOptions): ApiOperation {
        return new ApiOperation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:apimanagement/apiOperation:ApiOperation';

    /**
     * Returns true if the given object is an instance of ApiOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiOperation.__pulumiType;
    }

    /**
     * The Name of the API Management Service where the API exists. Changing this forces a new resource to be created.
     */
    public readonly apiManagementName!: pulumi.Output<string>;
    /**
     * The name of the API within the API Management Service where this API Operation should be created. Changing this forces a new resource to be created.
     */
    public readonly apiName!: pulumi.Output<string>;
    /**
     * A description for this API Operation, which may include HTML formatting tags.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Display Name for this API Management Operation.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The HTTP Method used for this API Management Operation, like `GET`, `DELETE`, `PUT` or `POST` - but not limited to these values.
     */
    public readonly method!: pulumi.Output<string>;
    /**
     * A unique identifier for this API Operation. Changing this forces a new resource to be created.
     */
    public readonly operationId!: pulumi.Output<string>;
    /**
     * A `request` block as defined below.
     */
    public readonly request!: pulumi.Output<outputs.apimanagement.ApiOperationRequest>;
    /**
     * The Name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * One or more `response` blocks as defined below.
     */
    public readonly responses!: pulumi.Output<outputs.apimanagement.ApiOperationResponse[] | undefined>;
    /**
     * One or more `templateParameter` blocks as defined below.
     */
    public readonly templateParameters!: pulumi.Output<outputs.apimanagement.ApiOperationTemplateParameter[] | undefined>;
    /**
     * The relative URL Template identifying the target resource for this operation, which may include parameters.
     */
    public readonly urlTemplate!: pulumi.Output<string>;

    /**
     * Create a ApiOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiOperationArgs | ApiOperationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ApiOperationState | undefined;
            inputs["apiManagementName"] = state ? state.apiManagementName : undefined;
            inputs["apiName"] = state ? state.apiName : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["method"] = state ? state.method : undefined;
            inputs["operationId"] = state ? state.operationId : undefined;
            inputs["request"] = state ? state.request : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["responses"] = state ? state.responses : undefined;
            inputs["templateParameters"] = state ? state.templateParameters : undefined;
            inputs["urlTemplate"] = state ? state.urlTemplate : undefined;
        } else {
            const args = argsOrState as ApiOperationArgs | undefined;
            if (!args || args.apiManagementName === undefined) {
                throw new Error("Missing required property 'apiManagementName'");
            }
            if (!args || args.apiName === undefined) {
                throw new Error("Missing required property 'apiName'");
            }
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            if (!args || args.method === undefined) {
                throw new Error("Missing required property 'method'");
            }
            if (!args || args.operationId === undefined) {
                throw new Error("Missing required property 'operationId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.urlTemplate === undefined) {
                throw new Error("Missing required property 'urlTemplate'");
            }
            inputs["apiManagementName"] = args ? args.apiManagementName : undefined;
            inputs["apiName"] = args ? args.apiName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["method"] = args ? args.method : undefined;
            inputs["operationId"] = args ? args.operationId : undefined;
            inputs["request"] = args ? args.request : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["responses"] = args ? args.responses : undefined;
            inputs["templateParameters"] = args ? args.templateParameters : undefined;
            inputs["urlTemplate"] = args ? args.urlTemplate : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ApiOperation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApiOperation resources.
 */
export interface ApiOperationState {
    /**
     * The Name of the API Management Service where the API exists. Changing this forces a new resource to be created.
     */
    readonly apiManagementName?: pulumi.Input<string>;
    /**
     * The name of the API within the API Management Service where this API Operation should be created. Changing this forces a new resource to be created.
     */
    readonly apiName?: pulumi.Input<string>;
    /**
     * A description for this API Operation, which may include HTML formatting tags.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Display Name for this API Management Operation.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The HTTP Method used for this API Management Operation, like `GET`, `DELETE`, `PUT` or `POST` - but not limited to these values.
     */
    readonly method?: pulumi.Input<string>;
    /**
     * A unique identifier for this API Operation. Changing this forces a new resource to be created.
     */
    readonly operationId?: pulumi.Input<string>;
    /**
     * A `request` block as defined below.
     */
    readonly request?: pulumi.Input<inputs.apimanagement.ApiOperationRequest>;
    /**
     * The Name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * One or more `response` blocks as defined below.
     */
    readonly responses?: pulumi.Input<pulumi.Input<inputs.apimanagement.ApiOperationResponse>[]>;
    /**
     * One or more `templateParameter` blocks as defined below.
     */
    readonly templateParameters?: pulumi.Input<pulumi.Input<inputs.apimanagement.ApiOperationTemplateParameter>[]>;
    /**
     * The relative URL Template identifying the target resource for this operation, which may include parameters.
     */
    readonly urlTemplate?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApiOperation resource.
 */
export interface ApiOperationArgs {
    /**
     * The Name of the API Management Service where the API exists. Changing this forces a new resource to be created.
     */
    readonly apiManagementName: pulumi.Input<string>;
    /**
     * The name of the API within the API Management Service where this API Operation should be created. Changing this forces a new resource to be created.
     */
    readonly apiName: pulumi.Input<string>;
    /**
     * A description for this API Operation, which may include HTML formatting tags.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Display Name for this API Management Operation.
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * The HTTP Method used for this API Management Operation, like `GET`, `DELETE`, `PUT` or `POST` - but not limited to these values.
     */
    readonly method: pulumi.Input<string>;
    /**
     * A unique identifier for this API Operation. Changing this forces a new resource to be created.
     */
    readonly operationId: pulumi.Input<string>;
    /**
     * A `request` block as defined below.
     */
    readonly request?: pulumi.Input<inputs.apimanagement.ApiOperationRequest>;
    /**
     * The Name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * One or more `response` blocks as defined below.
     */
    readonly responses?: pulumi.Input<pulumi.Input<inputs.apimanagement.ApiOperationResponse>[]>;
    /**
     * One or more `templateParameter` blocks as defined below.
     */
    readonly templateParameters?: pulumi.Input<pulumi.Input<inputs.apimanagement.ApiOperationTemplateParameter>[]>;
    /**
     * The relative URL Template identifying the target resource for this operation, which may include parameters.
     */
    readonly urlTemplate: pulumi.Input<string>;
}

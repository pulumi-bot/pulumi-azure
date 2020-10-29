// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a HTTP Request Trigger within a Logic App Workflow
 */
export class TriggerHttpRequest extends pulumi.CustomResource {
    /**
     * Get an existing TriggerHttpRequest resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TriggerHttpRequestState, opts?: pulumi.CustomResourceOptions): TriggerHttpRequest {
        return new TriggerHttpRequest(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:logicapps/triggerHttpRequest:TriggerHttpRequest';

    /**
     * Returns true if the given object is an instance of TriggerHttpRequest.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TriggerHttpRequest {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TriggerHttpRequest.__pulumiType;
    }

    /**
     * Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
     */
    public readonly logicAppId!: pulumi.Output<string>;
    /**
     * Specifies the HTTP Method which the request be using. Possible values include `DELETE`, `GET`, `PATCH`, `POST` or `PUT`.
     */
    public readonly method!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the HTTP Request Trigger to be created within the Logic App Workflow. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the Relative Path used for this Request.
     */
    public readonly relativePath!: pulumi.Output<string | undefined>;
    /**
     * A JSON Blob defining the Schema of the incoming request. This needs to be valid JSON.
     */
    public readonly schema!: pulumi.Output<string>;

    /**
     * Create a TriggerHttpRequest resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TriggerHttpRequestArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TriggerHttpRequestArgs | TriggerHttpRequestState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TriggerHttpRequestState | undefined;
            inputs["logicAppId"] = state ? state.logicAppId : undefined;
            inputs["method"] = state ? state.method : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["relativePath"] = state ? state.relativePath : undefined;
            inputs["schema"] = state ? state.schema : undefined;
        } else {
            const args = argsOrState as TriggerHttpRequestArgs | undefined;
            if (!args || args.logicAppId === undefined) {
                throw new Error("Missing required property 'logicAppId'");
            }
            if (!args || args.schema === undefined) {
                throw new Error("Missing required property 'schema'");
            }
            inputs["logicAppId"] = args ? args.logicAppId : undefined;
            inputs["method"] = args ? args.method : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["relativePath"] = args ? args.relativePath : undefined;
            inputs["schema"] = args ? args.schema : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TriggerHttpRequest.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TriggerHttpRequest resources.
 */
export interface TriggerHttpRequestState {
    /**
     * Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
     */
    readonly logicAppId?: pulumi.Input<string>;
    /**
     * Specifies the HTTP Method which the request be using. Possible values include `DELETE`, `GET`, `PATCH`, `POST` or `PUT`.
     */
    readonly method?: pulumi.Input<string>;
    /**
     * Specifies the name of the HTTP Request Trigger to be created within the Logic App Workflow. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the Relative Path used for this Request.
     */
    readonly relativePath?: pulumi.Input<string>;
    /**
     * A JSON Blob defining the Schema of the incoming request. This needs to be valid JSON.
     */
    readonly schema?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TriggerHttpRequest resource.
 */
export interface TriggerHttpRequestArgs {
    /**
     * Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
     */
    readonly logicAppId: pulumi.Input<string>;
    /**
     * Specifies the HTTP Method which the request be using. Possible values include `DELETE`, `GET`, `PATCH`, `POST` or `PUT`.
     */
    readonly method?: pulumi.Input<string>;
    /**
     * Specifies the name of the HTTP Request Trigger to be created within the Logic App Workflow. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the Relative Path used for this Request.
     */
    readonly relativePath?: pulumi.Input<string>;
    /**
     * A JSON Blob defining the Schema of the incoming request. This needs to be valid JSON.
     */
    readonly schema: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a backend within an API Management Service.
 *
 * ## Import
 *
 * API Management backends can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:apimanagement/backend:Backend example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/backends/backend1
 * ```
 */
export class Backend extends pulumi.CustomResource {
    /**
     * Get an existing Backend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackendState, opts?: pulumi.CustomResourceOptions): Backend {
        return new Backend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:apimanagement/backend:Backend';

    /**
     * Returns true if the given object is an instance of Backend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Backend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Backend.__pulumiType;
    }

    /**
     * The Name of the API Management Service where this backend should be created. Changing this forces a new resource to be created.
     */
    public readonly apiManagementName!: pulumi.Output<string>;
    /**
     * A `credentials` block as documented below.
     */
    public readonly credentials!: pulumi.Output<outputs.apimanagement.BackendCredentials | undefined>;
    /**
     * The description of the backend.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the API Management backend. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The protocol used by the backend host. Possible values are `http` or `soap`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * A `proxy` block as documented below.
     */
    public readonly proxy!: pulumi.Output<outputs.apimanagement.BackendProxy | undefined>;
    /**
     * The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The management URI of the backend host in an external system. This URI can be the ARM Resource ID of Logic Apps, Function Apps or API Apps, or the management endpoint of a Service Fabric cluster.
     */
    public readonly resourceId!: pulumi.Output<string | undefined>;
    /**
     * A `serviceFabricCluster` block as documented below.
     */
    public readonly serviceFabricCluster!: pulumi.Output<outputs.apimanagement.BackendServiceFabricCluster | undefined>;
    /**
     * The title of the backend.
     */
    public readonly title!: pulumi.Output<string | undefined>;
    /**
     * A `tls` block as documented below.
     */
    public readonly tls!: pulumi.Output<outputs.apimanagement.BackendTls | undefined>;
    /**
     * The URL of the backend host.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a Backend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackendArgs | BackendState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BackendState | undefined;
            inputs["apiManagementName"] = state ? state.apiManagementName : undefined;
            inputs["credentials"] = state ? state.credentials : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["proxy"] = state ? state.proxy : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["resourceId"] = state ? state.resourceId : undefined;
            inputs["serviceFabricCluster"] = state ? state.serviceFabricCluster : undefined;
            inputs["title"] = state ? state.title : undefined;
            inputs["tls"] = state ? state.tls : undefined;
            inputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as BackendArgs | undefined;
            if (!args || args.apiManagementName === undefined) {
                throw new Error("Missing required property 'apiManagementName'");
            }
            if (!args || args.protocol === undefined) {
                throw new Error("Missing required property 'protocol'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.url === undefined) {
                throw new Error("Missing required property 'url'");
            }
            inputs["apiManagementName"] = args ? args.apiManagementName : undefined;
            inputs["credentials"] = args ? args.credentials : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["proxy"] = args ? args.proxy : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["serviceFabricCluster"] = args ? args.serviceFabricCluster : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["tls"] = args ? args.tls : undefined;
            inputs["url"] = args ? args.url : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Backend.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Backend resources.
 */
export interface BackendState {
    /**
     * The Name of the API Management Service where this backend should be created. Changing this forces a new resource to be created.
     */
    readonly apiManagementName?: pulumi.Input<string>;
    /**
     * A `credentials` block as documented below.
     */
    readonly credentials?: pulumi.Input<inputs.apimanagement.BackendCredentials>;
    /**
     * The description of the backend.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the API Management backend. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The protocol used by the backend host. Possible values are `http` or `soap`.
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * A `proxy` block as documented below.
     */
    readonly proxy?: pulumi.Input<inputs.apimanagement.BackendProxy>;
    /**
     * The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The management URI of the backend host in an external system. This URI can be the ARM Resource ID of Logic Apps, Function Apps or API Apps, or the management endpoint of a Service Fabric cluster.
     */
    readonly resourceId?: pulumi.Input<string>;
    /**
     * A `serviceFabricCluster` block as documented below.
     */
    readonly serviceFabricCluster?: pulumi.Input<inputs.apimanagement.BackendServiceFabricCluster>;
    /**
     * The title of the backend.
     */
    readonly title?: pulumi.Input<string>;
    /**
     * A `tls` block as documented below.
     */
    readonly tls?: pulumi.Input<inputs.apimanagement.BackendTls>;
    /**
     * The URL of the backend host.
     */
    readonly url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Backend resource.
 */
export interface BackendArgs {
    /**
     * The Name of the API Management Service where this backend should be created. Changing this forces a new resource to be created.
     */
    readonly apiManagementName: pulumi.Input<string>;
    /**
     * A `credentials` block as documented below.
     */
    readonly credentials?: pulumi.Input<inputs.apimanagement.BackendCredentials>;
    /**
     * The description of the backend.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the API Management backend. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The protocol used by the backend host. Possible values are `http` or `soap`.
     */
    readonly protocol: pulumi.Input<string>;
    /**
     * A `proxy` block as documented below.
     */
    readonly proxy?: pulumi.Input<inputs.apimanagement.BackendProxy>;
    /**
     * The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The management URI of the backend host in an external system. This URI can be the ARM Resource ID of Logic Apps, Function Apps or API Apps, or the management endpoint of a Service Fabric cluster.
     */
    readonly resourceId?: pulumi.Input<string>;
    /**
     * A `serviceFabricCluster` block as documented below.
     */
    readonly serviceFabricCluster?: pulumi.Input<inputs.apimanagement.BackendServiceFabricCluster>;
    /**
     * The title of the backend.
     */
    readonly title?: pulumi.Input<string>;
    /**
     * A `tls` block as documented below.
     */
    readonly tls?: pulumi.Input<inputs.apimanagement.BackendTls>;
    /**
     * The URL of the backend host.
     */
    readonly url: pulumi.Input<string>;
}

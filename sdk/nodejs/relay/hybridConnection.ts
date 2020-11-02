// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Azure Relay Hybrid Connection.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleNamespace = new azure.relay.Namespace("exampleNamespace", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     skuName: "Standard",
 *     tags: {
 *         source: "managed",
 *     },
 * });
 * const exampleHybridConnection = new azure.relay.HybridConnection("exampleHybridConnection", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     relayNamespaceName: exampleNamespace.name,
 *     requiresClientAuthorization: false,
 *     userMetadata: "testmetadata",
 * });
 * ```
 *
 * ## Import
 *
 * Relay Hybrid Connection's can be imported using the `resource id`, e.g.
 */
export class HybridConnection extends pulumi.CustomResource {
    /**
     * Get an existing HybridConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HybridConnectionState, opts?: pulumi.CustomResourceOptions): HybridConnection {
        return new HybridConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:relay/hybridConnection:HybridConnection';

    /**
     * Returns true if the given object is an instance of HybridConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HybridConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HybridConnection.__pulumiType;
    }

    /**
     * Specifies the name of the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Azure Relay in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
     */
    public readonly relayNamespaceName!: pulumi.Output<string>;
    /**
     * Specify if client authorization is needed for this hybrid connection. True by default. Changing this forces a new resource to be created.
     */
    public readonly requiresClientAuthorization!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the resource group in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
     */
    public readonly userMetadata!: pulumi.Output<string | undefined>;

    /**
     * Create a HybridConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HybridConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HybridConnectionArgs | HybridConnectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as HybridConnectionState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["relayNamespaceName"] = state ? state.relayNamespaceName : undefined;
            inputs["requiresClientAuthorization"] = state ? state.requiresClientAuthorization : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["userMetadata"] = state ? state.userMetadata : undefined;
        } else {
            const args = argsOrState as HybridConnectionArgs | undefined;
            if (!args || args.relayNamespaceName === undefined) {
                throw new Error("Missing required property 'relayNamespaceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["relayNamespaceName"] = args ? args.relayNamespaceName : undefined;
            inputs["requiresClientAuthorization"] = args ? args.requiresClientAuthorization : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["userMetadata"] = args ? args.userMetadata : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(HybridConnection.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HybridConnection resources.
 */
export interface HybridConnectionState {
    /**
     * Specifies the name of the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Azure Relay in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
     */
    readonly relayNamespaceName?: pulumi.Input<string>;
    /**
     * Specify if client authorization is needed for this hybrid connection. True by default. Changing this forces a new resource to be created.
     */
    readonly requiresClientAuthorization?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
     */
    readonly userMetadata?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HybridConnection resource.
 */
export interface HybridConnectionArgs {
    /**
     * Specifies the name of the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Azure Relay in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
     */
    readonly relayNamespaceName: pulumi.Input<string>;
    /**
     * Specify if client authorization is needed for this hybrid connection. True by default. Changing this forces a new resource to be created.
     */
    readonly requiresClientAuthorization?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
     */
    readonly userMetadata?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an App Service Hybrid Connection for an existing App Service, Relay and Service Bus.
 *
 * ## Example Usage
 *
 * This example provisions an App Service, a Relay Hybrid Connection, and a Service Bus using their outputs to create the App Service Hybrid Connection.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const examplePlan = new azure.appservice.Plan("examplePlan", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: {
 *         tier: "Standard",
 *         size: "S1",
 *     },
 * });
 * const exampleAppService = new azure.appservice.AppService("exampleAppService", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     appServicePlanId: examplePlan.id,
 * });
 * const exampleNamespace = new azure.relay.Namespace("exampleNamespace", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     skuName: "Standard",
 * });
 * const exampleHybridConnection = new azure.relay.HybridConnection("exampleHybridConnection", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     relayNamespaceName: exampleNamespace.name,
 *     userMetadata: "examplemetadata",
 * });
 * const exampleAppservice_hybridConnectionHybridConnection = new azure.appservice.HybridConnection("exampleAppservice/hybridConnectionHybridConnection", {
 *     appServiceName: exampleAppService.name,
 *     resourceGroupName: exampleResourceGroup.name,
 *     relayId: exampleHybridConnection.id,
 *     hostname: "testhostname.example",
 *     port: 8080,
 *     sendKeyName: "exampleSharedAccessKey",
 * });
 * ```
 *
 * ## Import
 *
 * App Service Hybrid Connections can be imported using the `resource id`, e.g.
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
    public static readonly __pulumiType = 'azure:appservice/hybridConnection:HybridConnection';

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
     * Specifies the name of the App Service. Changing this forces a new resource to be created.
     */
    public readonly appServiceName!: pulumi.Output<string>;
    /**
     * The hostname of the endpoint.
     */
    public readonly hostname!: pulumi.Output<string>;
    /**
     * The name of the Relay Namespace.
     */
    public /*out*/ readonly namespaceName!: pulumi.Output<string>;
    /**
     * The port of the endpoint.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * The ID of the Service Bus Relay. Changing this forces a new resource to be created.
     */
    public readonly relayId!: pulumi.Output<string>;
    public /*out*/ readonly relayName!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the App Service.  Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The name of the Service Bus key which has Send permissions. Defaults to `RootManageSharedAccessKey`.
     */
    public readonly sendKeyName!: pulumi.Output<string | undefined>;
    /**
     * The value of the Service Bus Primary Access key.
     */
    public /*out*/ readonly sendKeyValue!: pulumi.Output<string>;
    /**
     * The name of the Service Bus namespace.
     */
    public /*out*/ readonly serviceBusNamespace!: pulumi.Output<string>;
    /**
     * The suffix for the service bus endpoint.
     */
    public /*out*/ readonly serviceBusSuffix!: pulumi.Output<string>;

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
            inputs["appServiceName"] = state ? state.appServiceName : undefined;
            inputs["hostname"] = state ? state.hostname : undefined;
            inputs["namespaceName"] = state ? state.namespaceName : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["relayId"] = state ? state.relayId : undefined;
            inputs["relayName"] = state ? state.relayName : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sendKeyName"] = state ? state.sendKeyName : undefined;
            inputs["sendKeyValue"] = state ? state.sendKeyValue : undefined;
            inputs["serviceBusNamespace"] = state ? state.serviceBusNamespace : undefined;
            inputs["serviceBusSuffix"] = state ? state.serviceBusSuffix : undefined;
        } else {
            const args = argsOrState as HybridConnectionArgs | undefined;
            if (!args || args.appServiceName === undefined) {
                throw new Error("Missing required property 'appServiceName'");
            }
            if (!args || args.hostname === undefined) {
                throw new Error("Missing required property 'hostname'");
            }
            if (!args || args.port === undefined) {
                throw new Error("Missing required property 'port'");
            }
            if (!args || args.relayId === undefined) {
                throw new Error("Missing required property 'relayId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["appServiceName"] = args ? args.appServiceName : undefined;
            inputs["hostname"] = args ? args.hostname : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["relayId"] = args ? args.relayId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sendKeyName"] = args ? args.sendKeyName : undefined;
            inputs["namespaceName"] = undefined /*out*/;
            inputs["relayName"] = undefined /*out*/;
            inputs["sendKeyValue"] = undefined /*out*/;
            inputs["serviceBusNamespace"] = undefined /*out*/;
            inputs["serviceBusSuffix"] = undefined /*out*/;
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
     * Specifies the name of the App Service. Changing this forces a new resource to be created.
     */
    readonly appServiceName?: pulumi.Input<string>;
    /**
     * The hostname of the endpoint.
     */
    readonly hostname?: pulumi.Input<string>;
    /**
     * The name of the Relay Namespace.
     */
    readonly namespaceName?: pulumi.Input<string>;
    /**
     * The port of the endpoint.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * The ID of the Service Bus Relay. Changing this forces a new resource to be created.
     */
    readonly relayId?: pulumi.Input<string>;
    readonly relayName?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the App Service.  Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The name of the Service Bus key which has Send permissions. Defaults to `RootManageSharedAccessKey`.
     */
    readonly sendKeyName?: pulumi.Input<string>;
    /**
     * The value of the Service Bus Primary Access key.
     */
    readonly sendKeyValue?: pulumi.Input<string>;
    /**
     * The name of the Service Bus namespace.
     */
    readonly serviceBusNamespace?: pulumi.Input<string>;
    /**
     * The suffix for the service bus endpoint.
     */
    readonly serviceBusSuffix?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HybridConnection resource.
 */
export interface HybridConnectionArgs {
    /**
     * Specifies the name of the App Service. Changing this forces a new resource to be created.
     */
    readonly appServiceName: pulumi.Input<string>;
    /**
     * The hostname of the endpoint.
     */
    readonly hostname: pulumi.Input<string>;
    /**
     * The port of the endpoint.
     */
    readonly port: pulumi.Input<number>;
    /**
     * The ID of the Service Bus Relay. Changing this forces a new resource to be created.
     */
    readonly relayId: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the App Service.  Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Service Bus key which has Send permissions. Defaults to `RootManageSharedAccessKey`.
     */
    readonly sendKeyName?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an App Service Virtual Network Association (this is for the [Regional VNet Integration](https://docs.microsoft.com/en-us/azure/app-service/web-sites-integrate-with-vnet#regional-vnet-integration) which is still in preview).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "uksouth"});
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("exampleVirtualNetwork", {
 *     addressSpaces: ["10.0.0.0/16"],
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleSubnet = new azure.network.Subnet("exampleSubnet", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefixes: ["10.0.1.0/24"],
 *     delegations: [{
 *         name: "example-delegation",
 *         serviceDelegation: {
 *             name: "Microsoft.Web/serverFarms",
 *             actions: ["Microsoft.Network/virtualNetworks/subnets/action"],
 *         },
 *     }],
 * });
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
 * const exampleVirtualNetworkSwiftConnection = new azure.appservice.VirtualNetworkSwiftConnection("exampleVirtualNetworkSwiftConnection", {
 *     appServiceId: exampleAppService.id,
 *     subnetId: exampleSubnet.id,
 * });
 * ```
 *
 * ## Import
 *
 * App Service Virtual Network Associations can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:appservice/virtualNetworkSwiftConnection:VirtualNetworkSwiftConnection myassociation /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/networkconfig/virtualNetwork
 * ```
 */
export class VirtualNetworkSwiftConnection extends pulumi.CustomResource {
    /**
     * Get an existing VirtualNetworkSwiftConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualNetworkSwiftConnectionState, opts?: pulumi.CustomResourceOptions): VirtualNetworkSwiftConnection {
        return new VirtualNetworkSwiftConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:appservice/virtualNetworkSwiftConnection:VirtualNetworkSwiftConnection';

    /**
     * Returns true if the given object is an instance of VirtualNetworkSwiftConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualNetworkSwiftConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualNetworkSwiftConnection.__pulumiType;
    }

    /**
     * The ID of the App Service to associate to the VNet. Changing this forces a new resource to be created.
     */
    public readonly appServiceId!: pulumi.Output<string>;
    /**
     * The ID of the subnet the app service will be associated to (the subnet must have a `serviceDelegation` configured for `Microsoft.Web/serverFarms`).
     */
    public readonly subnetId!: pulumi.Output<string>;

    /**
     * Create a VirtualNetworkSwiftConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualNetworkSwiftConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualNetworkSwiftConnectionArgs | VirtualNetworkSwiftConnectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VirtualNetworkSwiftConnectionState | undefined;
            inputs["appServiceId"] = state ? state.appServiceId : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
        } else {
            const args = argsOrState as VirtualNetworkSwiftConnectionArgs | undefined;
            if ((!args || args.appServiceId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'appServiceId'");
            }
            if ((!args || args.subnetId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'subnetId'");
            }
            inputs["appServiceId"] = args ? args.appServiceId : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VirtualNetworkSwiftConnection.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualNetworkSwiftConnection resources.
 */
export interface VirtualNetworkSwiftConnectionState {
    /**
     * The ID of the App Service to associate to the VNet. Changing this forces a new resource to be created.
     */
    readonly appServiceId?: pulumi.Input<string>;
    /**
     * The ID of the subnet the app service will be associated to (the subnet must have a `serviceDelegation` configured for `Microsoft.Web/serverFarms`).
     */
    readonly subnetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VirtualNetworkSwiftConnection resource.
 */
export interface VirtualNetworkSwiftConnectionArgs {
    /**
     * The ID of the App Service to associate to the VNet. Changing this forces a new resource to be created.
     */
    readonly appServiceId: pulumi.Input<string>;
    /**
     * The ID of the subnet the app service will be associated to (the subnet must have a `serviceDelegation` configured for `Microsoft.Web/serverFarms`).
     */
    readonly subnetId: pulumi.Input<string>;
}

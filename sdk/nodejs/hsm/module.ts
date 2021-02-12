// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Dedicated Hardware Security Module.
 *
 * > **Note**: Before using this resource, it's required to submit the request of registering the providers and features with Azure CLI `az provider register --namespace Microsoft.HardwareSecurityModules && az feature register --namespace Microsoft.HardwareSecurityModules --name AzureDedicatedHSM && az provider register --namespace Microsoft.Network && az feature register --namespace Microsoft.Network --name AllowBaremetalServers` and ask service team (hsmrequest@microsoft.com) to approve. See more details from https://docs.microsoft.com/en-us/azure/dedicated-hsm/tutorial-deploy-hsm-cli#prerequisites.
 *
 * > **Note**: If the quota is not enough in some region, please submit the quota request to service team.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("exampleVirtualNetwork", {
 *     addressSpaces: ["10.2.0.0/16"],
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleSubnet = new azure.network.Subnet("exampleSubnet", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefixes: ["10.2.0.0/24"],
 * });
 * const example2 = new azure.network.Subnet("example2", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefixes: ["10.2.1.0/24"],
 *     delegations: [{
 *         name: "first",
 *         serviceDelegation: {
 *             name: "Microsoft.HardwareSecurityModules/dedicatedHSMs",
 *             actions: [
 *                 "Microsoft.Network/networkinterfaces/*",
 *                 "Microsoft.Network/virtualNetworks/subnets/join/action",
 *             ],
 *         },
 *     }],
 * });
 * const example3 = new azure.network.Subnet("example3", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefixes: ["10.2.255.0/26"],
 * });
 * const examplePublicIp = new azure.network.PublicIp("examplePublicIp", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     allocationMethod: "Dynamic",
 * });
 * const exampleVirtualNetworkGateway = new azure.network.VirtualNetworkGateway("exampleVirtualNetworkGateway", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     type: "ExpressRoute",
 *     vpnType: "PolicyBased",
 *     sku: "Standard",
 *     ipConfigurations: [{
 *         publicIpAddressId: examplePublicIp.id,
 *         privateIpAddressAllocation: "Dynamic",
 *         subnetId: example3.id,
 *     }],
 * });
 * const exampleModule = new azure.hsm.Module("exampleModule", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     skuName: "SafeNet Luna Network HSM A790",
 *     networkProfile: {
 *         networkInterfacePrivateIpAddresses: ["10.2.1.8"],
 *         subnetId: example2.id,
 *     },
 *     stampId: "stamp2",
 *     tags: {
 *         env: "Test",
 *     },
 * }, {
 *     dependsOn: [exampleVirtualNetworkGateway],
 * });
 * ```
 *
 * ## Import
 *
 * Dedicated Hardware Security Module can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:hsm/module:Module example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/hsm1
 * ```
 */
export class Module extends pulumi.CustomResource {
    /**
     * Get an existing Module resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ModuleState, opts?: pulumi.CustomResourceOptions): Module {
        return new Module(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:hsm/module:Module';

    /**
     * Returns true if the given object is an instance of Module.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Module {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Module.__pulumiType;
    }

    /**
     * The Azure Region where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name which should be used for this Dedicated Hardware Security Module. Changing this forces a new Dedicated Hardware Security Module to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A `networkProfile` block as defined below.
     */
    public readonly networkProfile!: pulumi.Output<outputs.hsm.ModuleNetworkProfile>;
    /**
     * The name of the Resource Group where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The sku name of the dedicated hardware security module. Changing this forces a new Dedicated Hardware Security Module to be created.
     */
    public readonly skuName!: pulumi.Output<string>;
    /**
     * The ID of the stamp. Possible values are `stamp1` or `stamp2`. Changing this forces a new Dedicated Hardware Security Module to be created.
     */
    public readonly stampId!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags which should be assigned to the Dedicated Hardware Security Module.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The Dedicated Hardware Security Module zones. Changing this forces a new Dedicated Hardware Security Module to be created.
     */
    public readonly zones!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Module resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ModuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ModuleArgs | ModuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ModuleState | undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkProfile"] = state ? state.networkProfile : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["skuName"] = state ? state.skuName : undefined;
            inputs["stampId"] = state ? state.stampId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["zones"] = state ? state.zones : undefined;
        } else {
            const args = argsOrState as ModuleArgs | undefined;
            if ((!args || args.networkProfile === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkProfile'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.skuName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'skuName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkProfile"] = args ? args.networkProfile : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["skuName"] = args ? args.skuName : undefined;
            inputs["stampId"] = args ? args.stampId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["zones"] = args ? args.zones : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Module.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Module resources.
 */
export interface ModuleState {
    /**
     * The Azure Region where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name which should be used for this Dedicated Hardware Security Module. Changing this forces a new Dedicated Hardware Security Module to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `networkProfile` block as defined below.
     */
    readonly networkProfile?: pulumi.Input<inputs.hsm.ModuleNetworkProfile>;
    /**
     * The name of the Resource Group where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The sku name of the dedicated hardware security module. Changing this forces a new Dedicated Hardware Security Module to be created.
     */
    readonly skuName?: pulumi.Input<string>;
    /**
     * The ID of the stamp. Possible values are `stamp1` or `stamp2`. Changing this forces a new Dedicated Hardware Security Module to be created.
     */
    readonly stampId?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Dedicated Hardware Security Module.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Dedicated Hardware Security Module zones. Changing this forces a new Dedicated Hardware Security Module to be created.
     */
    readonly zones?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Module resource.
 */
export interface ModuleArgs {
    /**
     * The Azure Region where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name which should be used for this Dedicated Hardware Security Module. Changing this forces a new Dedicated Hardware Security Module to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `networkProfile` block as defined below.
     */
    readonly networkProfile: pulumi.Input<inputs.hsm.ModuleNetworkProfile>;
    /**
     * The name of the Resource Group where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The sku name of the dedicated hardware security module. Changing this forces a new Dedicated Hardware Security Module to be created.
     */
    readonly skuName: pulumi.Input<string>;
    /**
     * The ID of the stamp. Possible values are `stamp1` or `stamp2`. Changing this forces a new Dedicated Hardware Security Module to be created.
     */
    readonly stampId?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Dedicated Hardware Security Module.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Dedicated Hardware Security Module zones. Changing this forces a new Dedicated Hardware Security Module to be created.
     */
    readonly zones?: pulumi.Input<pulumi.Input<string>[]>;
}

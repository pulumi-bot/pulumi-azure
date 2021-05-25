// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a subnet. Subnets represent network segments within the IP space defined by the virtual network.
 *
 * > **NOTE on Virtual Networks and Subnet's:** This provider currently
 * provides both a standalone Subnet resource, and allows for Subnets to be defined in-line within the Virtual Network resource.
 * At this time you cannot use a Virtual Network with in-line Subnets in conjunction with any Subnet resources. Doing so will cause a conflict of Subnet configurations and will overwrite Subnet's.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
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
 *         name: "delegation",
 *         serviceDelegation: {
 *             name: "Microsoft.ContainerInstance/containerGroups",
 *             actions: [
 *                 "Microsoft.Network/virtualNetworks/subnets/join/action",
 *                 "Microsoft.Network/virtualNetworks/subnets/prepareNetworkPolicies/action",
 *             ],
 *         },
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Subnets can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:network/subnet:Subnet exampleSubnet /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1/subnets/mysubnet1
 * ```
 */
export class Subnet extends pulumi.CustomResource {
    /**
     * Get an existing Subnet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubnetState, opts?: pulumi.CustomResourceOptions): Subnet {
        return new Subnet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/subnet:Subnet';

    /**
     * Returns true if the given object is an instance of Subnet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Subnet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Subnet.__pulumiType;
    }

    /**
     * The address prefix to use for the subnet.
     *
     * @deprecated Use the `address_prefixes` property instead.
     */
    public readonly addressPrefix!: pulumi.Output<string>;
    /**
     * The address prefixes to use for the subnet.
     */
    public readonly addressPrefixes!: pulumi.Output<string[]>;
    /**
     * One or more `delegation` blocks as defined below.
     */
    public readonly delegations!: pulumi.Output<outputs.network.SubnetDelegation[] | undefined>;
    /**
     * Enable or Disable network policies for the private link endpoint on the subnet. Default value is `false`. Conflicts with enforce_private_link_service_network_policies.
     */
    public readonly enforcePrivateLinkEndpointNetworkPolicies!: pulumi.Output<boolean | undefined>;
    /**
     * Enable or Disable network policies for the private link service on the subnet. Default valule is `false`. Conflicts with `enforcePrivateLinkEndpointNetworkPolicies`.
     */
    public readonly enforcePrivateLinkServiceNetworkPolicies!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the subnet. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the subnet. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The list of IDs of Service Endpoint Policies to associate with the subnet.
     */
    public readonly serviceEndpointPolicyIds!: pulumi.Output<string[] | undefined>;
    /**
     * The list of Service endpoints to associate with the subnet. Possible values include: `Microsoft.AzureActiveDirectory`, `Microsoft.AzureCosmosDB`, `Microsoft.ContainerRegistry`, `Microsoft.EventHub`, `Microsoft.KeyVault`, `Microsoft.ServiceBus`, `Microsoft.Sql`, `Microsoft.Storage` and `Microsoft.Web`.
     */
    public readonly serviceEndpoints!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the virtual network to which to attach the subnet. Changing this forces a new resource to be created.
     */
    public readonly virtualNetworkName!: pulumi.Output<string>;

    /**
     * Create a Subnet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubnetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubnetArgs | SubnetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SubnetState | undefined;
            inputs["addressPrefix"] = state ? state.addressPrefix : undefined;
            inputs["addressPrefixes"] = state ? state.addressPrefixes : undefined;
            inputs["delegations"] = state ? state.delegations : undefined;
            inputs["enforcePrivateLinkEndpointNetworkPolicies"] = state ? state.enforcePrivateLinkEndpointNetworkPolicies : undefined;
            inputs["enforcePrivateLinkServiceNetworkPolicies"] = state ? state.enforcePrivateLinkServiceNetworkPolicies : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["serviceEndpointPolicyIds"] = state ? state.serviceEndpointPolicyIds : undefined;
            inputs["serviceEndpoints"] = state ? state.serviceEndpoints : undefined;
            inputs["virtualNetworkName"] = state ? state.virtualNetworkName : undefined;
        } else {
            const args = argsOrState as SubnetArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.virtualNetworkName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'virtualNetworkName'");
            }
            inputs["addressPrefix"] = args ? args.addressPrefix : undefined;
            inputs["addressPrefixes"] = args ? args.addressPrefixes : undefined;
            inputs["delegations"] = args ? args.delegations : undefined;
            inputs["enforcePrivateLinkEndpointNetworkPolicies"] = args ? args.enforcePrivateLinkEndpointNetworkPolicies : undefined;
            inputs["enforcePrivateLinkServiceNetworkPolicies"] = args ? args.enforcePrivateLinkServiceNetworkPolicies : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceEndpointPolicyIds"] = args ? args.serviceEndpointPolicyIds : undefined;
            inputs["serviceEndpoints"] = args ? args.serviceEndpoints : undefined;
            inputs["virtualNetworkName"] = args ? args.virtualNetworkName : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Subnet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Subnet resources.
 */
export interface SubnetState {
    /**
     * The address prefix to use for the subnet.
     *
     * @deprecated Use the `address_prefixes` property instead.
     */
    addressPrefix?: pulumi.Input<string>;
    /**
     * The address prefixes to use for the subnet.
     */
    addressPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * One or more `delegation` blocks as defined below.
     */
    delegations?: pulumi.Input<pulumi.Input<inputs.network.SubnetDelegation>[]>;
    /**
     * Enable or Disable network policies for the private link endpoint on the subnet. Default value is `false`. Conflicts with enforce_private_link_service_network_policies.
     */
    enforcePrivateLinkEndpointNetworkPolicies?: pulumi.Input<boolean>;
    /**
     * Enable or Disable network policies for the private link service on the subnet. Default valule is `false`. Conflicts with `enforcePrivateLinkEndpointNetworkPolicies`.
     */
    enforcePrivateLinkServiceNetworkPolicies?: pulumi.Input<boolean>;
    /**
     * The name of the subnet. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the subnet. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The list of IDs of Service Endpoint Policies to associate with the subnet.
     */
    serviceEndpointPolicyIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of Service endpoints to associate with the subnet. Possible values include: `Microsoft.AzureActiveDirectory`, `Microsoft.AzureCosmosDB`, `Microsoft.ContainerRegistry`, `Microsoft.EventHub`, `Microsoft.KeyVault`, `Microsoft.ServiceBus`, `Microsoft.Sql`, `Microsoft.Storage` and `Microsoft.Web`.
     */
    serviceEndpoints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the virtual network to which to attach the subnet. Changing this forces a new resource to be created.
     */
    virtualNetworkName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Subnet resource.
 */
export interface SubnetArgs {
    /**
     * The address prefix to use for the subnet.
     *
     * @deprecated Use the `address_prefixes` property instead.
     */
    addressPrefix?: pulumi.Input<string>;
    /**
     * The address prefixes to use for the subnet.
     */
    addressPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * One or more `delegation` blocks as defined below.
     */
    delegations?: pulumi.Input<pulumi.Input<inputs.network.SubnetDelegation>[]>;
    /**
     * Enable or Disable network policies for the private link endpoint on the subnet. Default value is `false`. Conflicts with enforce_private_link_service_network_policies.
     */
    enforcePrivateLinkEndpointNetworkPolicies?: pulumi.Input<boolean>;
    /**
     * Enable or Disable network policies for the private link service on the subnet. Default valule is `false`. Conflicts with `enforcePrivateLinkEndpointNetworkPolicies`.
     */
    enforcePrivateLinkServiceNetworkPolicies?: pulumi.Input<boolean>;
    /**
     * The name of the subnet. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the subnet. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The list of IDs of Service Endpoint Policies to associate with the subnet.
     */
    serviceEndpointPolicyIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of Service endpoints to associate with the subnet. Possible values include: `Microsoft.AzureActiveDirectory`, `Microsoft.AzureCosmosDB`, `Microsoft.ContainerRegistry`, `Microsoft.EventHub`, `Microsoft.KeyVault`, `Microsoft.ServiceBus`, `Microsoft.Sql`, `Microsoft.Storage` and `Microsoft.Web`.
     */
    serviceEndpoints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the virtual network to which to attach the subnet. Changing this forces a new resource to be created.
     */
    virtualNetworkName: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages the association between a Network Interface and a Application Security Group.
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
 * });
 * const exampleApplicationSecurityGroup = new azure.network.ApplicationSecurityGroup("exampleApplicationSecurityGroup", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleNetworkInterface = new azure.network.NetworkInterface("exampleNetworkInterface", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     ipConfigurations: [{
 *         name: "testconfiguration1",
 *         subnetId: exampleSubnet.id,
 *         privateIpAddressAllocation: "Dynamic",
 *     }],
 * });
 * const exampleNetworkInterfaceApplicationSecurityGroupAssociation = new azure.network.NetworkInterfaceApplicationSecurityGroupAssociation("exampleNetworkInterfaceApplicationSecurityGroupAssociation", {
 *     networkInterfaceId: exampleNetworkInterface.id,
 *     applicationSecurityGroupId: exampleApplicationSecurityGroup.id,
 * });
 * ```
 *
 * ## Import
 *
 * Associations between Network Interfaces and Application Security Groups can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:network/networkInterfaceApplicationSecurityGroupAssociation:NetworkInterfaceApplicationSecurityGroupAssociation association1 "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.network/networkInterfaces/nic1|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/applicationSecurityGroups/securityGroup1"
 * ```
 */
export class NetworkInterfaceApplicationSecurityGroupAssociation extends pulumi.CustomResource {
    /**
     * Get an existing NetworkInterfaceApplicationSecurityGroupAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkInterfaceApplicationSecurityGroupAssociationState, opts?: pulumi.CustomResourceOptions): NetworkInterfaceApplicationSecurityGroupAssociation {
        return new NetworkInterfaceApplicationSecurityGroupAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/networkInterfaceApplicationSecurityGroupAssociation:NetworkInterfaceApplicationSecurityGroupAssociation';

    /**
     * Returns true if the given object is an instance of NetworkInterfaceApplicationSecurityGroupAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkInterfaceApplicationSecurityGroupAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkInterfaceApplicationSecurityGroupAssociation.__pulumiType;
    }

    /**
     * The ID of the Application Security Group which this Network Interface which should be connected to. Changing this forces a new resource to be created.
     */
    public readonly applicationSecurityGroupId!: pulumi.Output<string>;
    /**
     * The ID of the Network Interface. Changing this forces a new resource to be created.
     */
    public readonly networkInterfaceId!: pulumi.Output<string>;

    /**
     * Create a NetworkInterfaceApplicationSecurityGroupAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkInterfaceApplicationSecurityGroupAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkInterfaceApplicationSecurityGroupAssociationArgs | NetworkInterfaceApplicationSecurityGroupAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkInterfaceApplicationSecurityGroupAssociationState | undefined;
            inputs["applicationSecurityGroupId"] = state ? state.applicationSecurityGroupId : undefined;
            inputs["networkInterfaceId"] = state ? state.networkInterfaceId : undefined;
        } else {
            const args = argsOrState as NetworkInterfaceApplicationSecurityGroupAssociationArgs | undefined;
            if ((!args || args.applicationSecurityGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationSecurityGroupId'");
            }
            if ((!args || args.networkInterfaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkInterfaceId'");
            }
            inputs["applicationSecurityGroupId"] = args ? args.applicationSecurityGroupId : undefined;
            inputs["networkInterfaceId"] = args ? args.networkInterfaceId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NetworkInterfaceApplicationSecurityGroupAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkInterfaceApplicationSecurityGroupAssociation resources.
 */
export interface NetworkInterfaceApplicationSecurityGroupAssociationState {
    /**
     * The ID of the Application Security Group which this Network Interface which should be connected to. Changing this forces a new resource to be created.
     */
    applicationSecurityGroupId?: pulumi.Input<string>;
    /**
     * The ID of the Network Interface. Changing this forces a new resource to be created.
     */
    networkInterfaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkInterfaceApplicationSecurityGroupAssociation resource.
 */
export interface NetworkInterfaceApplicationSecurityGroupAssociationArgs {
    /**
     * The ID of the Application Security Group which this Network Interface which should be connected to. Changing this forces a new resource to be created.
     */
    applicationSecurityGroupId: pulumi.Input<string>;
    /**
     * The ID of the Network Interface. Changing this forces a new resource to be created.
     */
    networkInterfaceId: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Associates a NAT Gateway with a Subnet within a Virtual Network.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "East US 2"});
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("exampleVirtualNetwork", {
 *     addressSpaces: ["10.0.0.0/16"],
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleSubnet = new azure.network.Subnet("exampleSubnet", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefixes: ["10.0.2.0/24"],
 * });
 * const exampleNatGateway = new azure.network.NatGateway("exampleNatGateway", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleSubnetNatGatewayAssociation = new azure.network.SubnetNatGatewayAssociation("exampleSubnetNatGatewayAssociation", {
 *     subnetId: exampleSubnet.id,
 *     natGatewayId: exampleNatGateway.id,
 * });
 * ```
 *
 * ## Import
 *
 * Subnet NAT Gateway Associations can be imported using the `resource id` of the Subnet, e.g.
 */
export class SubnetNatGatewayAssociation extends pulumi.CustomResource {
    /**
     * Get an existing SubnetNatGatewayAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubnetNatGatewayAssociationState, opts?: pulumi.CustomResourceOptions): SubnetNatGatewayAssociation {
        return new SubnetNatGatewayAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/subnetNatGatewayAssociation:SubnetNatGatewayAssociation';

    /**
     * Returns true if the given object is an instance of SubnetNatGatewayAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SubnetNatGatewayAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SubnetNatGatewayAssociation.__pulumiType;
    }

    /**
     * The ID of the NAT Gateway which should be associated with the Subnet. Changing this forces a new resource to be created.
     */
    public readonly natGatewayId!: pulumi.Output<string>;
    /**
     * The ID of the Subnet. Changing this forces a new resource to be created.
     */
    public readonly subnetId!: pulumi.Output<string>;

    /**
     * Create a SubnetNatGatewayAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubnetNatGatewayAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubnetNatGatewayAssociationArgs | SubnetNatGatewayAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SubnetNatGatewayAssociationState | undefined;
            inputs["natGatewayId"] = state ? state.natGatewayId : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
        } else {
            const args = argsOrState as SubnetNatGatewayAssociationArgs | undefined;
            if (!args || args.natGatewayId === undefined) {
                throw new Error("Missing required property 'natGatewayId'");
            }
            if (!args || args.subnetId === undefined) {
                throw new Error("Missing required property 'subnetId'");
            }
            inputs["natGatewayId"] = args ? args.natGatewayId : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SubnetNatGatewayAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SubnetNatGatewayAssociation resources.
 */
export interface SubnetNatGatewayAssociationState {
    /**
     * The ID of the NAT Gateway which should be associated with the Subnet. Changing this forces a new resource to be created.
     */
    readonly natGatewayId?: pulumi.Input<string>;
    /**
     * The ID of the Subnet. Changing this forces a new resource to be created.
     */
    readonly subnetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SubnetNatGatewayAssociation resource.
 */
export interface SubnetNatGatewayAssociationArgs {
    /**
     * The ID of the NAT Gateway which should be associated with the Subnet. Changing this forces a new resource to be created.
     */
    readonly natGatewayId: pulumi.Input<string>;
    /**
     * The ID of the Subnet. Changing this forces a new resource to be created.
     */
    readonly subnetId: pulumi.Input<string>;
}

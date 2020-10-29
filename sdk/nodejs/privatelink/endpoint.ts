// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Private Endpoint.
 *
 * > **NOTE** Private Endpoint is currently in Public Preview.
 *
 * Azure Private Endpoint is a network interface that connects you privately and securely to a service powered by Azure Private Link. Private Endpoint uses a private IP address from your VNet, effectively bringing the service into your VNet. The service could be an Azure service such as Azure Storage, SQL, etc. or your own Private Link Service.
 */
export class Endpoint extends pulumi.CustomResource {
    /**
     * Get an existing Endpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointState, opts?: pulumi.CustomResourceOptions): Endpoint {
        return new Endpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:privatelink/endpoint:Endpoint';

    /**
     * Returns true if the given object is an instance of Endpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Endpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Endpoint.__pulumiType;
    }

    public /*out*/ readonly customDnsConfigs!: pulumi.Output<outputs.privatelink.EndpointCustomDnsConfig[]>;
    /**
     * The supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the Name of the Private Endpoint. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly privateDnsZoneConfigs!: pulumi.Output<outputs.privatelink.EndpointPrivateDnsZoneConfig[]>;
    /**
     * A `privateDnsZoneGroup` block as defined below.
     */
    public readonly privateDnsZoneGroup!: pulumi.Output<outputs.privatelink.EndpointPrivateDnsZoneGroup | undefined>;
    /**
     * A `privateServiceConnection` block as defined below.
     */
    public readonly privateServiceConnection!: pulumi.Output<outputs.privatelink.EndpointPrivateServiceConnection>;
    /**
     * Specifies the Name of the Resource Group within which the Private Endpoint should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The ID of the Subnet from which Private IP Addresses will be allocated for this Private Endpoint. Changing this forces a new resource to be created.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Endpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndpointArgs | EndpointState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EndpointState | undefined;
            inputs["customDnsConfigs"] = state ? state.customDnsConfigs : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["privateDnsZoneConfigs"] = state ? state.privateDnsZoneConfigs : undefined;
            inputs["privateDnsZoneGroup"] = state ? state.privateDnsZoneGroup : undefined;
            inputs["privateServiceConnection"] = state ? state.privateServiceConnection : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as EndpointArgs | undefined;
            if (!args || args.privateServiceConnection === undefined) {
                throw new Error("Missing required property 'privateServiceConnection'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.subnetId === undefined) {
                throw new Error("Missing required property 'subnetId'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["privateDnsZoneGroup"] = args ? args.privateDnsZoneGroup : undefined;
            inputs["privateServiceConnection"] = args ? args.privateServiceConnection : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["customDnsConfigs"] = undefined /*out*/;
            inputs["privateDnsZoneConfigs"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Endpoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Endpoint resources.
 */
export interface EndpointState {
    readonly customDnsConfigs?: pulumi.Input<pulumi.Input<inputs.privatelink.EndpointCustomDnsConfig>[]>;
    /**
     * The supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the Name of the Private Endpoint. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    readonly privateDnsZoneConfigs?: pulumi.Input<pulumi.Input<inputs.privatelink.EndpointPrivateDnsZoneConfig>[]>;
    /**
     * A `privateDnsZoneGroup` block as defined below.
     */
    readonly privateDnsZoneGroup?: pulumi.Input<inputs.privatelink.EndpointPrivateDnsZoneGroup>;
    /**
     * A `privateServiceConnection` block as defined below.
     */
    readonly privateServiceConnection?: pulumi.Input<inputs.privatelink.EndpointPrivateServiceConnection>;
    /**
     * Specifies the Name of the Resource Group within which the Private Endpoint should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The ID of the Subnet from which Private IP Addresses will be allocated for this Private Endpoint. Changing this forces a new resource to be created.
     */
    readonly subnetId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Endpoint resource.
 */
export interface EndpointArgs {
    /**
     * The supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the Name of the Private Endpoint. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `privateDnsZoneGroup` block as defined below.
     */
    readonly privateDnsZoneGroup?: pulumi.Input<inputs.privatelink.EndpointPrivateDnsZoneGroup>;
    /**
     * A `privateServiceConnection` block as defined below.
     */
    readonly privateServiceConnection: pulumi.Input<inputs.privatelink.EndpointPrivateServiceConnection>;
    /**
     * Specifies the Name of the Resource Group within which the Private Endpoint should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The ID of the Subnet from which Private IP Addresses will be allocated for this Private Endpoint. Changing this forces a new resource to be created.
     */
    readonly subnetId: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Private Link Service.
 *
 * > **NOTE** Private Link is now in [GA](https://docs.microsoft.com/en-gb/azure/private-link/).
 */
export class LinkService extends pulumi.CustomResource {
    /**
     * Get an existing LinkService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LinkServiceState, opts?: pulumi.CustomResourceOptions): LinkService {
        return new LinkService(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:privatedns/linkService:LinkService';

    /**
     * Returns true if the given object is an instance of LinkService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LinkService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LinkService.__pulumiType;
    }

    /**
     * A globally unique DNS Name for your Private Link Service. You can use this alias to request a connection to your Private Link Service.
     */
    public /*out*/ readonly alias!: pulumi.Output<string>;
    /**
     * A list of Subscription UUID/GUID's that will be automatically be able to use this Private Link Service.
     */
    public readonly autoApprovalSubscriptionIds!: pulumi.Output<string[] | undefined>;
    /**
     * Should the Private Link Service support the Proxy Protocol? Defaults to `false`.
     */
    public readonly enableProxyProtocol!: pulumi.Output<boolean | undefined>;
    /**
     * A list of Frontend IP Configuration ID's from a Standard Load Balancer, where traffic from the Private Link Service should be routed. You can use Load Balancer Rules to direct this traffic to appropriate backend pools where your applications are running.
     */
    public readonly loadBalancerFrontendIpConfigurationIds!: pulumi.Output<string[]>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of this Private Link Service. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * One or more (up to 8) `natIpConfiguration` block as defined below.
     */
    public readonly natIpConfigurations!: pulumi.Output<outputs.privatedns.LinkServiceNatIpConfiguration[]>;
    /**
     * The name of the Resource Group where the Private Link Service should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A list of Subscription UUID/GUID's that will be able to see this Private Link Service.
     */
    public readonly visibilitySubscriptionIds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a LinkService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LinkServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LinkServiceArgs | LinkServiceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LinkServiceState | undefined;
            inputs["alias"] = state ? state.alias : undefined;
            inputs["autoApprovalSubscriptionIds"] = state ? state.autoApprovalSubscriptionIds : undefined;
            inputs["enableProxyProtocol"] = state ? state.enableProxyProtocol : undefined;
            inputs["loadBalancerFrontendIpConfigurationIds"] = state ? state.loadBalancerFrontendIpConfigurationIds : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["natIpConfigurations"] = state ? state.natIpConfigurations : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["visibilitySubscriptionIds"] = state ? state.visibilitySubscriptionIds : undefined;
        } else {
            const args = argsOrState as LinkServiceArgs | undefined;
            if (!args || args.loadBalancerFrontendIpConfigurationIds === undefined) {
                throw new Error("Missing required property 'loadBalancerFrontendIpConfigurationIds'");
            }
            if (!args || args.natIpConfigurations === undefined) {
                throw new Error("Missing required property 'natIpConfigurations'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["autoApprovalSubscriptionIds"] = args ? args.autoApprovalSubscriptionIds : undefined;
            inputs["enableProxyProtocol"] = args ? args.enableProxyProtocol : undefined;
            inputs["loadBalancerFrontendIpConfigurationIds"] = args ? args.loadBalancerFrontendIpConfigurationIds : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["natIpConfigurations"] = args ? args.natIpConfigurations : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["visibilitySubscriptionIds"] = args ? args.visibilitySubscriptionIds : undefined;
            inputs["alias"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LinkService.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LinkService resources.
 */
export interface LinkServiceState {
    /**
     * A globally unique DNS Name for your Private Link Service. You can use this alias to request a connection to your Private Link Service.
     */
    readonly alias?: pulumi.Input<string>;
    /**
     * A list of Subscription UUID/GUID's that will be automatically be able to use this Private Link Service.
     */
    readonly autoApprovalSubscriptionIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Should the Private Link Service support the Proxy Protocol? Defaults to `false`.
     */
    readonly enableProxyProtocol?: pulumi.Input<boolean>;
    /**
     * A list of Frontend IP Configuration ID's from a Standard Load Balancer, where traffic from the Private Link Service should be routed. You can use Load Balancer Rules to direct this traffic to appropriate backend pools where your applications are running.
     */
    readonly loadBalancerFrontendIpConfigurationIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of this Private Link Service. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * One or more (up to 8) `natIpConfiguration` block as defined below.
     */
    readonly natIpConfigurations?: pulumi.Input<pulumi.Input<inputs.privatedns.LinkServiceNatIpConfiguration>[]>;
    /**
     * The name of the Resource Group where the Private Link Service should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of Subscription UUID/GUID's that will be able to see this Private Link Service.
     */
    readonly visibilitySubscriptionIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a LinkService resource.
 */
export interface LinkServiceArgs {
    /**
     * A list of Subscription UUID/GUID's that will be automatically be able to use this Private Link Service.
     */
    readonly autoApprovalSubscriptionIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Should the Private Link Service support the Proxy Protocol? Defaults to `false`.
     */
    readonly enableProxyProtocol?: pulumi.Input<boolean>;
    /**
     * A list of Frontend IP Configuration ID's from a Standard Load Balancer, where traffic from the Private Link Service should be routed. You can use Load Balancer Rules to direct this traffic to appropriate backend pools where your applications are running.
     */
    readonly loadBalancerFrontendIpConfigurationIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of this Private Link Service. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * One or more (up to 8) `natIpConfiguration` block as defined below.
     */
    readonly natIpConfigurations: pulumi.Input<pulumi.Input<inputs.privatedns.LinkServiceNatIpConfiguration>[]>;
    /**
     * The name of the Resource Group where the Private Link Service should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of Subscription UUID/GUID's that will be able to see this Private Link Service.
     */
    readonly visibilitySubscriptionIds?: pulumi.Input<pulumi.Input<string>[]>;
}

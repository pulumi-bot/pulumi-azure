// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an App Service Environment.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "westeurope"});
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("exampleVirtualNetwork", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     addressSpaces: ["10.0.0.0/16"],
 * });
 * const ase = new azure.network.Subnet("ase", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefix: "10.0.1.0/24",
 * });
 * const gateway = new azure.network.Subnet("gateway", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefix: "10.0.2.0/24",
 * });
 * const exampleEnvironment = new azure.appservice.Environment("exampleEnvironment", {
 *     subnetId: ase.id,
 *     pricingTier: "I2",
 *     frontEndScaleFactor: 10,
 * });
 * ```
 */
export class Environment extends pulumi.CustomResource {
    /**
     * Get an existing Environment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EnvironmentState, opts?: pulumi.CustomResourceOptions): Environment {
        return new Environment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:appservice/environment:Environment';

    /**
     * Returns true if the given object is an instance of Environment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Environment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Environment.__pulumiType;
    }

    /**
     * Scale factor for front end instances. Possible values are between `5` and `15`. Defaults to `15`.
     */
    public readonly frontEndScaleFactor!: pulumi.Output<number | undefined>;
    /**
     * Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment. Possible values are `None`, `Web` and `Publishing`. Defaults to `None`.
     */
    public readonly internalLoadBalancingMode!: pulumi.Output<string | undefined>;
    /**
     * The location where the App Service Environment exists.
     */
    public /*out*/ readonly location!: pulumi.Output<string>;
    /**
     * The name of the App Service Environment. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Pricing tier for the front end instances. Possible values are `I1`, `I2` and `I3`. Defaults to `I1`.
     */
    public readonly pricingTier!: pulumi.Output<string | undefined>;
    /**
     * The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnetId`).
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Environment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnvironmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EnvironmentArgs | EnvironmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EnvironmentState | undefined;
            inputs["frontEndScaleFactor"] = state ? state.frontEndScaleFactor : undefined;
            inputs["internalLoadBalancingMode"] = state ? state.internalLoadBalancingMode : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["pricingTier"] = state ? state.pricingTier : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as EnvironmentArgs | undefined;
            if (!args || args.subnetId === undefined) {
                throw new Error("Missing required property 'subnetId'");
            }
            inputs["frontEndScaleFactor"] = args ? args.frontEndScaleFactor : undefined;
            inputs["internalLoadBalancingMode"] = args ? args.internalLoadBalancingMode : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["pricingTier"] = args ? args.pricingTier : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["location"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Environment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Environment resources.
 */
export interface EnvironmentState {
    /**
     * Scale factor for front end instances. Possible values are between `5` and `15`. Defaults to `15`.
     */
    readonly frontEndScaleFactor?: pulumi.Input<number>;
    /**
     * Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment. Possible values are `None`, `Web` and `Publishing`. Defaults to `None`.
     */
    readonly internalLoadBalancingMode?: pulumi.Input<string>;
    /**
     * The location where the App Service Environment exists.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the App Service Environment. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Pricing tier for the front end instances. Possible values are `I1`, `I2` and `I3`. Defaults to `I1`.
     */
    readonly pricingTier?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnetId`).
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
     */
    readonly subnetId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Environment resource.
 */
export interface EnvironmentArgs {
    /**
     * Scale factor for front end instances. Possible values are between `5` and `15`. Defaults to `15`.
     */
    readonly frontEndScaleFactor?: pulumi.Input<number>;
    /**
     * Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment. Possible values are `None`, `Web` and `Publishing`. Defaults to `None`.
     */
    readonly internalLoadBalancingMode?: pulumi.Input<string>;
    /**
     * The name of the App Service Environment. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Pricing tier for the front end instances. Possible values are `I1`, `I2` and `I3`. Defaults to `I1`.
     */
    readonly pricingTier?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnetId`).
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
     */
    readonly subnetId: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

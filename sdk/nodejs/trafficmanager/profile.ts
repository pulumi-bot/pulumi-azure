// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Traffic Manager Profile to which multiple endpoints can be attached.
 *
 * @deprecated azure.trafficmanager.Profile has been deprecated in favor of azure.network.TrafficManagerProfile
 */
export class Profile extends pulumi.CustomResource {
    /**
     * Get an existing Profile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProfileState, opts?: pulumi.CustomResourceOptions): Profile {
        pulumi.log.warn("Profile is deprecated: azure.trafficmanager.Profile has been deprecated in favor of azure.network.TrafficManagerProfile")
        return new Profile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:trafficmanager/profile:Profile';

    /**
     * Returns true if the given object is an instance of Profile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Profile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Profile.__pulumiType;
    }

    /**
     * This block specifies the DNS configuration of the Profile, it supports the fields documented below.
     */
    public readonly dnsConfig!: pulumi.Output<outputs.trafficmanager.ProfileDnsConfig>;
    /**
     * The FQDN of the created Profile.
     */
    public /*out*/ readonly fqdn!: pulumi.Output<string>;
    /**
     * This block specifies the Endpoint monitoring configuration for the Profile, it supports the fields documented below.
     */
    public readonly monitorConfig!: pulumi.Output<outputs.trafficmanager.ProfileMonitorConfig>;
    /**
     * The name of the Traffic Manager profile. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The status of the profile, can be set to either `Enabled` or `Disabled`. Defaults to `Enabled`.
     */
    public readonly profileStatus!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Traffic Manager profile.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the algorithm used to route traffic, possible values are:
     */
    public readonly trafficRoutingMethod!: pulumi.Output<string>;

    /**
     * Create a Profile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated azure.trafficmanager.Profile has been deprecated in favor of azure.network.TrafficManagerProfile */
    constructor(name: string, args: ProfileArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azure.trafficmanager.Profile has been deprecated in favor of azure.network.TrafficManagerProfile */
    constructor(name: string, argsOrState?: ProfileArgs | ProfileState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Profile is deprecated: azure.trafficmanager.Profile has been deprecated in favor of azure.network.TrafficManagerProfile")
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProfileState | undefined;
            inputs["dnsConfig"] = state ? state.dnsConfig : undefined;
            inputs["fqdn"] = state ? state.fqdn : undefined;
            inputs["monitorConfig"] = state ? state.monitorConfig : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["profileStatus"] = state ? state.profileStatus : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["trafficRoutingMethod"] = state ? state.trafficRoutingMethod : undefined;
        } else {
            const args = argsOrState as ProfileArgs | undefined;
            if (!args || args.dnsConfig === undefined) {
                throw new Error("Missing required property 'dnsConfig'");
            }
            if (!args || args.monitorConfig === undefined) {
                throw new Error("Missing required property 'monitorConfig'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.trafficRoutingMethod === undefined) {
                throw new Error("Missing required property 'trafficRoutingMethod'");
            }
            inputs["dnsConfig"] = args ? args.dnsConfig : undefined;
            inputs["monitorConfig"] = args ? args.monitorConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["profileStatus"] = args ? args.profileStatus : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["trafficRoutingMethod"] = args ? args.trafficRoutingMethod : undefined;
            inputs["fqdn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Profile.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Profile resources.
 */
export interface ProfileState {
    /**
     * This block specifies the DNS configuration of the Profile, it supports the fields documented below.
     */
    readonly dnsConfig?: pulumi.Input<inputs.trafficmanager.ProfileDnsConfig>;
    /**
     * The FQDN of the created Profile.
     */
    readonly fqdn?: pulumi.Input<string>;
    /**
     * This block specifies the Endpoint monitoring configuration for the Profile, it supports the fields documented below.
     */
    readonly monitorConfig?: pulumi.Input<inputs.trafficmanager.ProfileMonitorConfig>;
    /**
     * The name of the Traffic Manager profile. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The status of the profile, can be set to either `Enabled` or `Disabled`. Defaults to `Enabled`.
     */
    readonly profileStatus?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Traffic Manager profile.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the algorithm used to route traffic, possible values are:
     */
    readonly trafficRoutingMethod?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Profile resource.
 */
export interface ProfileArgs {
    /**
     * This block specifies the DNS configuration of the Profile, it supports the fields documented below.
     */
    readonly dnsConfig: pulumi.Input<inputs.trafficmanager.ProfileDnsConfig>;
    /**
     * This block specifies the Endpoint monitoring configuration for the Profile, it supports the fields documented below.
     */
    readonly monitorConfig: pulumi.Input<inputs.trafficmanager.ProfileMonitorConfig>;
    /**
     * The name of the Traffic Manager profile. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The status of the profile, can be set to either `Enabled` or `Disabled`. Defaults to `Enabled`.
     */
    readonly profileStatus?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Traffic Manager profile.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the algorithm used to route traffic, possible values are:
     */
    readonly trafficRoutingMethod: pulumi.Input<string>;
}

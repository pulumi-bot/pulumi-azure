// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a VPN Server Configuration.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/vpn_server_configuration.html.markdown.
 */
export class VpnServerConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing VpnServerConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpnServerConfigurationState, opts?: pulumi.CustomResourceOptions): VpnServerConfiguration {
        return new VpnServerConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/vpnServerConfiguration:VpnServerConfiguration';

    /**
     * Returns true if the given object is an instance of VpnServerConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpnServerConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpnServerConfiguration.__pulumiType;
    }

    /**
     * A `azureActiveDirectoryAuthentication` block as defined below.
     */
    public readonly azureActiveDirectoryAuthentications!: pulumi.Output<outputs.network.VpnServerConfigurationAzureActiveDirectoryAuthentication[] | undefined>;
    /**
     * One or more `clientRevokedCertificate` blocks as defined below.
     */
    public readonly clientRevokedCertificates!: pulumi.Output<outputs.network.VpnServerConfigurationClientRevokedCertificate[] | undefined>;
    /**
     * One or more `clientRootCertificate` blocks as defined below.
     */
    public readonly clientRootCertificates!: pulumi.Output<outputs.network.VpnServerConfigurationClientRootCertificate[] | undefined>;
    /**
     * A `ipsecPolicy` block as defined below.
     */
    public readonly ipsecPolicy!: pulumi.Output<outputs.network.VpnServerConfigurationIpsecPolicy | undefined>;
    /**
     * The Azure location where this VPN Server Configuration should be created. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The Name which should be used for this VPN Server Configuration. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A `radiusServer` block as defined below.
     */
    public readonly radiusServer!: pulumi.Output<outputs.network.VpnServerConfigurationRadiusServer | undefined>;
    /**
     * The Name of the Resource Group in which this VPN Server Configuration should be created. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * A list of one of more Authentication Types applicable for this VPN Server Configuration. Possible values are `AAD` (Azure Active Directory), `Certificate` and `Radius`.
     */
    public readonly vpnAuthenticationTypes!: pulumi.Output<string>;
    /**
     * A list of VPN Protocols to use for this Server Configuration. Possible values are `IkeV2` and `OpenVPN`.
     */
    public readonly vpnProtocols!: pulumi.Output<string[]>;

    /**
     * Create a VpnServerConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpnServerConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpnServerConfigurationArgs | VpnServerConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VpnServerConfigurationState | undefined;
            inputs["azureActiveDirectoryAuthentications"] = state ? state.azureActiveDirectoryAuthentications : undefined;
            inputs["clientRevokedCertificates"] = state ? state.clientRevokedCertificates : undefined;
            inputs["clientRootCertificates"] = state ? state.clientRootCertificates : undefined;
            inputs["ipsecPolicy"] = state ? state.ipsecPolicy : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["radiusServer"] = state ? state.radiusServer : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vpnAuthenticationTypes"] = state ? state.vpnAuthenticationTypes : undefined;
            inputs["vpnProtocols"] = state ? state.vpnProtocols : undefined;
        } else {
            const args = argsOrState as VpnServerConfigurationArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.vpnAuthenticationTypes === undefined) {
                throw new Error("Missing required property 'vpnAuthenticationTypes'");
            }
            inputs["azureActiveDirectoryAuthentications"] = args ? args.azureActiveDirectoryAuthentications : undefined;
            inputs["clientRevokedCertificates"] = args ? args.clientRevokedCertificates : undefined;
            inputs["clientRootCertificates"] = args ? args.clientRootCertificates : undefined;
            inputs["ipsecPolicy"] = args ? args.ipsecPolicy : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["radiusServer"] = args ? args.radiusServer : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpnAuthenticationTypes"] = args ? args.vpnAuthenticationTypes : undefined;
            inputs["vpnProtocols"] = args ? args.vpnProtocols : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VpnServerConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpnServerConfiguration resources.
 */
export interface VpnServerConfigurationState {
    /**
     * A `azureActiveDirectoryAuthentication` block as defined below.
     */
    readonly azureActiveDirectoryAuthentications?: pulumi.Input<pulumi.Input<inputs.network.VpnServerConfigurationAzureActiveDirectoryAuthentication>[]>;
    /**
     * One or more `clientRevokedCertificate` blocks as defined below.
     */
    readonly clientRevokedCertificates?: pulumi.Input<pulumi.Input<inputs.network.VpnServerConfigurationClientRevokedCertificate>[]>;
    /**
     * One or more `clientRootCertificate` blocks as defined below.
     */
    readonly clientRootCertificates?: pulumi.Input<pulumi.Input<inputs.network.VpnServerConfigurationClientRootCertificate>[]>;
    /**
     * A `ipsecPolicy` block as defined below.
     */
    readonly ipsecPolicy?: pulumi.Input<inputs.network.VpnServerConfigurationIpsecPolicy>;
    /**
     * The Azure location where this VPN Server Configuration should be created. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The Name which should be used for this VPN Server Configuration. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `radiusServer` block as defined below.
     */
    readonly radiusServer?: pulumi.Input<inputs.network.VpnServerConfigurationRadiusServer>;
    /**
     * The Name of the Resource Group in which this VPN Server Configuration should be created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of one of more Authentication Types applicable for this VPN Server Configuration. Possible values are `AAD` (Azure Active Directory), `Certificate` and `Radius`.
     */
    readonly vpnAuthenticationTypes?: pulumi.Input<string>;
    /**
     * A list of VPN Protocols to use for this Server Configuration. Possible values are `IkeV2` and `OpenVPN`.
     */
    readonly vpnProtocols?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a VpnServerConfiguration resource.
 */
export interface VpnServerConfigurationArgs {
    /**
     * A `azureActiveDirectoryAuthentication` block as defined below.
     */
    readonly azureActiveDirectoryAuthentications?: pulumi.Input<pulumi.Input<inputs.network.VpnServerConfigurationAzureActiveDirectoryAuthentication>[]>;
    /**
     * One or more `clientRevokedCertificate` blocks as defined below.
     */
    readonly clientRevokedCertificates?: pulumi.Input<pulumi.Input<inputs.network.VpnServerConfigurationClientRevokedCertificate>[]>;
    /**
     * One or more `clientRootCertificate` blocks as defined below.
     */
    readonly clientRootCertificates?: pulumi.Input<pulumi.Input<inputs.network.VpnServerConfigurationClientRootCertificate>[]>;
    /**
     * A `ipsecPolicy` block as defined below.
     */
    readonly ipsecPolicy?: pulumi.Input<inputs.network.VpnServerConfigurationIpsecPolicy>;
    /**
     * The Azure location where this VPN Server Configuration should be created. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The Name which should be used for this VPN Server Configuration. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `radiusServer` block as defined below.
     */
    readonly radiusServer?: pulumi.Input<inputs.network.VpnServerConfigurationRadiusServer>;
    /**
     * The Name of the Resource Group in which this VPN Server Configuration should be created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of one of more Authentication Types applicable for this VPN Server Configuration. Possible values are `AAD` (Azure Active Directory), `Certificate` and `Radius`.
     */
    readonly vpnAuthenticationTypes: pulumi.Input<string>;
    /**
     * A list of VPN Protocols to use for this Server Configuration. Possible values are `IkeV2` and `OpenVPN`.
     */
    readonly vpnProtocols?: pulumi.Input<pulumi.Input<string>[]>;
}
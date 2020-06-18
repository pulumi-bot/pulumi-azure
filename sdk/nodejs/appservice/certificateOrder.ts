// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an App Service Certificate Order.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleCertificateOrder = new azure.appservice.CertificateOrder("exampleCertificateOrder", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: "global",
 *     distinguishedName: "CN=example.com",
 *     productType: "Standard",
 * });
 * ```
 */
export class CertificateOrder extends pulumi.CustomResource {
    /**
     * Get an existing CertificateOrder resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CertificateOrderState, opts?: pulumi.CustomResourceOptions): CertificateOrder {
        return new CertificateOrder(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:appservice/certificateOrder:CertificateOrder';

    /**
     * Returns true if the given object is an instance of CertificateOrder.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CertificateOrder {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CertificateOrder.__pulumiType;
    }

    /**
     * Reasons why App Service Certificate is not renewable at the current moment.
     */
    public /*out*/ readonly appServiceCertificateNotRenewableReasons!: pulumi.Output<string[]>;
    /**
     * true if the certificate should be automatically renewed when it expires; otherwise, false. Defaults to true.
     */
    public readonly autoRenew!: pulumi.Output<boolean | undefined>;
    /**
     * State of the Key Vault secret. A `certificates` block as defined below.
     */
    public /*out*/ readonly certificates!: pulumi.Output<outputs.appservice.CertificateOrderCertificate[]>;
    /**
     * Last CSR that was created for this order.
     */
    public readonly csr!: pulumi.Output<string>;
    /**
     * The Distinguished Name for the App Service Certificate Order.
     */
    public readonly distinguishedName!: pulumi.Output<string>;
    /**
     * Domain verification token.
     */
    public /*out*/ readonly domainVerificationToken!: pulumi.Output<string>;
    /**
     * Certificate expiration time.
     */
    public /*out*/ readonly expirationTime!: pulumi.Output<string>;
    /**
     * Certificate thumbprint intermediate certificate.
     */
    public /*out*/ readonly intermediateThumbprint!: pulumi.Output<string>;
    /**
     * Whether the private key is external or not.
     */
    public /*out*/ readonly isPrivateKeyExternal!: pulumi.Output<boolean>;
    /**
     * Certificate key size.  Defaults to 2048.
     */
    public readonly keySize!: pulumi.Output<number | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. Currently the only valid value is `global`.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the certificate. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Certificate product type, such as `Standard` or `WildCard`.
     */
    public readonly productType!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Certificate thumbprint for root certificate.
     */
    public /*out*/ readonly rootThumbprint!: pulumi.Output<string>;
    /**
     * Certificate thumbprint for signed certificate.
     */
    public /*out*/ readonly signedCertificateThumbprint!: pulumi.Output<string>;
    /**
     * Current order status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Duration in years (must be between `1` and `3`).  Defaults to `1`.
     */
    public readonly validityInYears!: pulumi.Output<number | undefined>;

    /**
     * Create a CertificateOrder resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateOrderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertificateOrderArgs | CertificateOrderState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CertificateOrderState | undefined;
            inputs["appServiceCertificateNotRenewableReasons"] = state ? state.appServiceCertificateNotRenewableReasons : undefined;
            inputs["autoRenew"] = state ? state.autoRenew : undefined;
            inputs["certificates"] = state ? state.certificates : undefined;
            inputs["csr"] = state ? state.csr : undefined;
            inputs["distinguishedName"] = state ? state.distinguishedName : undefined;
            inputs["domainVerificationToken"] = state ? state.domainVerificationToken : undefined;
            inputs["expirationTime"] = state ? state.expirationTime : undefined;
            inputs["intermediateThumbprint"] = state ? state.intermediateThumbprint : undefined;
            inputs["isPrivateKeyExternal"] = state ? state.isPrivateKeyExternal : undefined;
            inputs["keySize"] = state ? state.keySize : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["productType"] = state ? state.productType : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["rootThumbprint"] = state ? state.rootThumbprint : undefined;
            inputs["signedCertificateThumbprint"] = state ? state.signedCertificateThumbprint : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["validityInYears"] = state ? state.validityInYears : undefined;
        } else {
            const args = argsOrState as CertificateOrderArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["autoRenew"] = args ? args.autoRenew : undefined;
            inputs["csr"] = args ? args.csr : undefined;
            inputs["distinguishedName"] = args ? args.distinguishedName : undefined;
            inputs["keySize"] = args ? args.keySize : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["productType"] = args ? args.productType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["validityInYears"] = args ? args.validityInYears : undefined;
            inputs["appServiceCertificateNotRenewableReasons"] = undefined /*out*/;
            inputs["certificates"] = undefined /*out*/;
            inputs["domainVerificationToken"] = undefined /*out*/;
            inputs["expirationTime"] = undefined /*out*/;
            inputs["intermediateThumbprint"] = undefined /*out*/;
            inputs["isPrivateKeyExternal"] = undefined /*out*/;
            inputs["rootThumbprint"] = undefined /*out*/;
            inputs["signedCertificateThumbprint"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CertificateOrder.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CertificateOrder resources.
 */
export interface CertificateOrderState {
    /**
     * Reasons why App Service Certificate is not renewable at the current moment.
     */
    readonly appServiceCertificateNotRenewableReasons?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * true if the certificate should be automatically renewed when it expires; otherwise, false. Defaults to true.
     */
    readonly autoRenew?: pulumi.Input<boolean>;
    /**
     * State of the Key Vault secret. A `certificates` block as defined below.
     */
    readonly certificates?: pulumi.Input<pulumi.Input<inputs.appservice.CertificateOrderCertificate>[]>;
    /**
     * Last CSR that was created for this order.
     */
    readonly csr?: pulumi.Input<string>;
    /**
     * The Distinguished Name for the App Service Certificate Order.
     */
    readonly distinguishedName?: pulumi.Input<string>;
    /**
     * Domain verification token.
     */
    readonly domainVerificationToken?: pulumi.Input<string>;
    /**
     * Certificate expiration time.
     */
    readonly expirationTime?: pulumi.Input<string>;
    /**
     * Certificate thumbprint intermediate certificate.
     */
    readonly intermediateThumbprint?: pulumi.Input<string>;
    /**
     * Whether the private key is external or not.
     */
    readonly isPrivateKeyExternal?: pulumi.Input<boolean>;
    /**
     * Certificate key size.  Defaults to 2048.
     */
    readonly keySize?: pulumi.Input<number>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. Currently the only valid value is `global`.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the certificate. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Certificate product type, such as `Standard` or `WildCard`.
     */
    readonly productType?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Certificate thumbprint for root certificate.
     */
    readonly rootThumbprint?: pulumi.Input<string>;
    /**
     * Certificate thumbprint for signed certificate.
     */
    readonly signedCertificateThumbprint?: pulumi.Input<string>;
    /**
     * Current order status.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Duration in years (must be between `1` and `3`).  Defaults to `1`.
     */
    readonly validityInYears?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a CertificateOrder resource.
 */
export interface CertificateOrderArgs {
    /**
     * true if the certificate should be automatically renewed when it expires; otherwise, false. Defaults to true.
     */
    readonly autoRenew?: pulumi.Input<boolean>;
    /**
     * Last CSR that was created for this order.
     */
    readonly csr?: pulumi.Input<string>;
    /**
     * The Distinguished Name for the App Service Certificate Order.
     */
    readonly distinguishedName?: pulumi.Input<string>;
    /**
     * Certificate key size.  Defaults to 2048.
     */
    readonly keySize?: pulumi.Input<number>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. Currently the only valid value is `global`.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the certificate. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Certificate product type, such as `Standard` or `WildCard`.
     */
    readonly productType?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Duration in years (must be between `1` and `3`).  Defaults to `1`.
     */
    readonly validityInYears?: pulumi.Input<number>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an API Management Product.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleService = new azure.apimanagement.Service("exampleService", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     publisherName: "My Company",
 *     publisherEmail: "company@exmaple.com",
 *     skuName: "Developer_1",
 * });
 * const exampleProduct = new azure.apimanagement.Product("exampleProduct", {
 *     productId: "test-product",
 *     apiManagementName: exampleService.name,
 *     resourceGroupName: exampleResourceGroup.name,
 *     displayName: "Test Product",
 *     subscriptionRequired: true,
 *     approvalRequired: true,
 *     published: true,
 * });
 * ```
 *
 * ## Import
 *
 * API Management Products can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:apimanagement/product:Product example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/products/myproduct
 * ```
 */
export class Product extends pulumi.CustomResource {
    /**
     * Get an existing Product resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProductState, opts?: pulumi.CustomResourceOptions): Product {
        return new Product(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:apimanagement/product:Product';

    /**
     * Returns true if the given object is an instance of Product.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Product {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Product.__pulumiType;
    }

    /**
     * The name of the API Management Service. Changing this forces a new resource to be created.
     */
    public readonly apiManagementName!: pulumi.Output<string>;
    /**
     * Do subscribers need to be approved prior to being able to use the Product?
     */
    public readonly approvalRequired!: pulumi.Output<boolean | undefined>;
    /**
     * A description of this Product, which may include HTML formatting tags.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Display Name for this API Management Product.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The Identifier for this Product, which must be unique within the API Management Service. Changing this forces a new resource to be created.
     */
    public readonly productId!: pulumi.Output<string>;
    /**
     * Is this Product Published?
     */
    public readonly published!: pulumi.Output<boolean>;
    /**
     * The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Is a Subscription required to access API's included in this Product?
     */
    public readonly subscriptionRequired!: pulumi.Output<boolean>;
    /**
     * The number of subscriptions a user can have to this Product at the same time.
     */
    public readonly subscriptionsLimit!: pulumi.Output<number | undefined>;
    /**
     * The Terms and Conditions for this Product, which must be accepted by Developers before they can begin the Subscription process.
     */
    public readonly terms!: pulumi.Output<string | undefined>;

    /**
     * Create a Product resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProductArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProductArgs | ProductState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProductState | undefined;
            inputs["apiManagementName"] = state ? state.apiManagementName : undefined;
            inputs["approvalRequired"] = state ? state.approvalRequired : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["productId"] = state ? state.productId : undefined;
            inputs["published"] = state ? state.published : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["subscriptionRequired"] = state ? state.subscriptionRequired : undefined;
            inputs["subscriptionsLimit"] = state ? state.subscriptionsLimit : undefined;
            inputs["terms"] = state ? state.terms : undefined;
        } else {
            const args = argsOrState as ProductArgs | undefined;
            if (!args || args.apiManagementName === undefined) {
                throw new Error("Missing required property 'apiManagementName'");
            }
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            if (!args || args.productId === undefined) {
                throw new Error("Missing required property 'productId'");
            }
            if (!args || args.published === undefined) {
                throw new Error("Missing required property 'published'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.subscriptionRequired === undefined) {
                throw new Error("Missing required property 'subscriptionRequired'");
            }
            inputs["apiManagementName"] = args ? args.apiManagementName : undefined;
            inputs["approvalRequired"] = args ? args.approvalRequired : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["productId"] = args ? args.productId : undefined;
            inputs["published"] = args ? args.published : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["subscriptionRequired"] = args ? args.subscriptionRequired : undefined;
            inputs["subscriptionsLimit"] = args ? args.subscriptionsLimit : undefined;
            inputs["terms"] = args ? args.terms : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Product.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Product resources.
 */
export interface ProductState {
    /**
     * The name of the API Management Service. Changing this forces a new resource to be created.
     */
    readonly apiManagementName?: pulumi.Input<string>;
    /**
     * Do subscribers need to be approved prior to being able to use the Product?
     */
    readonly approvalRequired?: pulumi.Input<boolean>;
    /**
     * A description of this Product, which may include HTML formatting tags.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Display Name for this API Management Product.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The Identifier for this Product, which must be unique within the API Management Service. Changing this forces a new resource to be created.
     */
    readonly productId?: pulumi.Input<string>;
    /**
     * Is this Product Published?
     */
    readonly published?: pulumi.Input<boolean>;
    /**
     * The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Is a Subscription required to access API's included in this Product?
     */
    readonly subscriptionRequired?: pulumi.Input<boolean>;
    /**
     * The number of subscriptions a user can have to this Product at the same time.
     */
    readonly subscriptionsLimit?: pulumi.Input<number>;
    /**
     * The Terms and Conditions for this Product, which must be accepted by Developers before they can begin the Subscription process.
     */
    readonly terms?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Product resource.
 */
export interface ProductArgs {
    /**
     * The name of the API Management Service. Changing this forces a new resource to be created.
     */
    readonly apiManagementName: pulumi.Input<string>;
    /**
     * Do subscribers need to be approved prior to being able to use the Product?
     */
    readonly approvalRequired?: pulumi.Input<boolean>;
    /**
     * A description of this Product, which may include HTML formatting tags.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Display Name for this API Management Product.
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * The Identifier for this Product, which must be unique within the API Management Service. Changing this forces a new resource to be created.
     */
    readonly productId: pulumi.Input<string>;
    /**
     * Is this Product Published?
     */
    readonly published: pulumi.Input<boolean>;
    /**
     * The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Is a Subscription required to access API's included in this Product?
     */
    readonly subscriptionRequired: pulumi.Input<boolean>;
    /**
     * The number of subscriptions a user can have to this Product at the same time.
     */
    readonly subscriptionsLimit?: pulumi.Input<number>;
    /**
     * The Terms and Conditions for this Product, which must be accepted by Developers before they can begin the Subscription process.
     */
    readonly terms?: pulumi.Input<string>;
}

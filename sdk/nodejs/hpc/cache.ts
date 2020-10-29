// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a HPC Cache.
 *
 * > **Note**: During the first several months of the GA release, a request must be made to the Azure HPC Cache team to add your subscription to the access list before it can be used to create a cache instance. Fill out [this form](https://aka.ms/onboard-hpc-cache) to request access.
 *
 * > **NOTE:**: By request of the service team the provider no longer automatically registering the `Microsoft.StorageCache` Resource Provider for this resource. To register it you can run `az provider register --namespace 'Microsoft.StorageCache'`.
 */
export class Cache extends pulumi.CustomResource {
    /**
     * Get an existing Cache resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CacheState, opts?: pulumi.CustomResourceOptions): Cache {
        return new Cache(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:hpc/cache:Cache';

    /**
     * Returns true if the given object is an instance of Cache.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cache {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cache.__pulumiType;
    }

    /**
     * The size of the HPC Cache, in GB. Possible values are `3072`, `6144`, `12288`, `24576`, and `49152`. Changing this forces a new resource to be created.
     */
    public readonly cacheSizeInGb!: pulumi.Output<number>;
    /**
     * Specifies the supported Azure Region where the HPC Cache should be created. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * A list of IP Addresses where the HPC Cache can be mounted.
     */
    public /*out*/ readonly mountAddresses!: pulumi.Output<string[]>;
    /**
     * The IPv4 maximum transmission unit configured for the subnet of the HPC Cache. Possible values range from 576 - 1500. Defaults to 1500.
     */
    public readonly mtu!: pulumi.Output<number | undefined>;
    /**
     * The name of the HPC Cache. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Resource Group in which to create the HPC Cache. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Whether root squash property is enabled for this HPC Cache.
     */
    public readonly rootSquashEnabled!: pulumi.Output<boolean>;
    /**
     * The SKU of HPC Cache to use. Possible values are `Standard_2G`, `Standard_4G` and `Standard_8G`. Changing this forces a new resource to be created.
     */
    public readonly skuName!: pulumi.Output<string>;
    /**
     * The ID of the Subnet for the HPC Cache. Changing this forces a new resource to be created.
     */
    public readonly subnetId!: pulumi.Output<string>;

    /**
     * Create a Cache resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CacheArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CacheArgs | CacheState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CacheState | undefined;
            inputs["cacheSizeInGb"] = state ? state.cacheSizeInGb : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["mountAddresses"] = state ? state.mountAddresses : undefined;
            inputs["mtu"] = state ? state.mtu : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["rootSquashEnabled"] = state ? state.rootSquashEnabled : undefined;
            inputs["skuName"] = state ? state.skuName : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
        } else {
            const args = argsOrState as CacheArgs | undefined;
            if (!args || args.cacheSizeInGb === undefined) {
                throw new Error("Missing required property 'cacheSizeInGb'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.skuName === undefined) {
                throw new Error("Missing required property 'skuName'");
            }
            if (!args || args.subnetId === undefined) {
                throw new Error("Missing required property 'subnetId'");
            }
            inputs["cacheSizeInGb"] = args ? args.cacheSizeInGb : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["mtu"] = args ? args.mtu : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["rootSquashEnabled"] = args ? args.rootSquashEnabled : undefined;
            inputs["skuName"] = args ? args.skuName : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["mountAddresses"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Cache.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cache resources.
 */
export interface CacheState {
    /**
     * The size of the HPC Cache, in GB. Possible values are `3072`, `6144`, `12288`, `24576`, and `49152`. Changing this forces a new resource to be created.
     */
    readonly cacheSizeInGb?: pulumi.Input<number>;
    /**
     * Specifies the supported Azure Region where the HPC Cache should be created. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * A list of IP Addresses where the HPC Cache can be mounted.
     */
    readonly mountAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The IPv4 maximum transmission unit configured for the subnet of the HPC Cache. Possible values range from 576 - 1500. Defaults to 1500.
     */
    readonly mtu?: pulumi.Input<number>;
    /**
     * The name of the HPC Cache. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which to create the HPC Cache. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Whether root squash property is enabled for this HPC Cache.
     */
    readonly rootSquashEnabled?: pulumi.Input<boolean>;
    /**
     * The SKU of HPC Cache to use. Possible values are `Standard_2G`, `Standard_4G` and `Standard_8G`. Changing this forces a new resource to be created.
     */
    readonly skuName?: pulumi.Input<string>;
    /**
     * The ID of the Subnet for the HPC Cache. Changing this forces a new resource to be created.
     */
    readonly subnetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cache resource.
 */
export interface CacheArgs {
    /**
     * The size of the HPC Cache, in GB. Possible values are `3072`, `6144`, `12288`, `24576`, and `49152`. Changing this forces a new resource to be created.
     */
    readonly cacheSizeInGb: pulumi.Input<number>;
    /**
     * Specifies the supported Azure Region where the HPC Cache should be created. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The IPv4 maximum transmission unit configured for the subnet of the HPC Cache. Possible values range from 576 - 1500. Defaults to 1500.
     */
    readonly mtu?: pulumi.Input<number>;
    /**
     * The name of the HPC Cache. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which to create the HPC Cache. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Whether root squash property is enabled for this HPC Cache.
     */
    readonly rootSquashEnabled?: pulumi.Input<boolean>;
    /**
     * The SKU of HPC Cache to use. Possible values are `Standard_2G`, `Standard_4G` and `Standard_8G`. Changing this forces a new resource to be created.
     */
    readonly skuName: pulumi.Input<string>;
    /**
     * The ID of the Subnet for the HPC Cache. Changing this forces a new resource to be created.
     */
    readonly subnetId: pulumi.Input<string>;
}

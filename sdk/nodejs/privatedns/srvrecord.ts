// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Enables you to manage DNS SRV Records within Azure Private DNS.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = new azure.core.ResourceGroup("example", {location: "West US"});
 * const testZone = new azure.privatedns.Zone("testZone", {resourceGroupName: azurerm_resource_group.test.name});
 * const testSRVRecord = new azure.privatedns.SRVRecord("testSRVRecord", {
 *     resourceGroupName: azurerm_resource_group.test.name,
 *     zoneName: testZone.name,
 *     ttl: 300,
 *     records: [
 *         {
 *             priority: 1,
 *             weight: 5,
 *             port: 8080,
 *             target: "target1.contoso.com",
 *         },
 *         {
 *             priority: 10,
 *             weight: 10,
 *             port: 8080,
 *             target: "target2.contoso.com",
 *         },
 *     ],
 *     tags: {
 *         Environment: "Production",
 *     },
 * });
 * ```
 */
export class SRVRecord extends pulumi.CustomResource {
    /**
     * Get an existing SRVRecord resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SRVRecordState, opts?: pulumi.CustomResourceOptions): SRVRecord {
        return new SRVRecord(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:privatedns/sRVRecord:SRVRecord';

    /**
     * Returns true if the given object is an instance of SRVRecord.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SRVRecord {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SRVRecord.__pulumiType;
    }

    /**
     * The FQDN of the DNS SRV Record.
     */
    public /*out*/ readonly fqdn!: pulumi.Output<string>;
    /**
     * The name of the DNS SRV Record. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * One or more `record` blocks as defined below.
     */
    public readonly records!: pulumi.Output<outputs.privatedns.SRVRecordRecord[]>;
    /**
     * Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly ttl!: pulumi.Output<number>;
    /**
     * Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly zoneName!: pulumi.Output<string>;

    /**
     * Create a SRVRecord resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SRVRecordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SRVRecordArgs | SRVRecordState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SRVRecordState | undefined;
            inputs["fqdn"] = state ? state.fqdn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["records"] = state ? state.records : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
            inputs["zoneName"] = state ? state.zoneName : undefined;
        } else {
            const args = argsOrState as SRVRecordArgs | undefined;
            if (!args || args.records === undefined) {
                throw new Error("Missing required property 'records'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.ttl === undefined) {
                throw new Error("Missing required property 'ttl'");
            }
            if (!args || args.zoneName === undefined) {
                throw new Error("Missing required property 'zoneName'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["records"] = args ? args.records : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["zoneName"] = args ? args.zoneName : undefined;
            inputs["fqdn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SRVRecord.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SRVRecord resources.
 */
export interface SRVRecordState {
    /**
     * The FQDN of the DNS SRV Record.
     */
    readonly fqdn?: pulumi.Input<string>;
    /**
     * The name of the DNS SRV Record. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * One or more `record` blocks as defined below.
     */
    readonly records?: pulumi.Input<pulumi.Input<inputs.privatedns.SRVRecordRecord>[]>;
    /**
     * Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly ttl?: pulumi.Input<number>;
    /**
     * Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
     */
    readonly zoneName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SRVRecord resource.
 */
export interface SRVRecordArgs {
    /**
     * The name of the DNS SRV Record. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * One or more `record` blocks as defined below.
     */
    readonly records: pulumi.Input<pulumi.Input<inputs.privatedns.SRVRecordRecord>[]>;
    /**
     * Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly ttl: pulumi.Input<number>;
    /**
     * Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
     */
    readonly zoneName: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Enables you to manage DNS CAA Records within Azure DNS.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US"});
 * const exampleZone = new azure.dns.Zone("exampleZone", {resourceGroupName: exampleResourceGroup.name});
 * const exampleCaaRecord = new azure.dns.CaaRecord("exampleCaaRecord", {
 *     zoneName: exampleZone.name,
 *     resourceGroupName: exampleResourceGroup.name,
 *     ttl: 300,
 *     records: [
 *         {
 *             flags: 0,
 *             tag: "issue",
 *             value: "example.com",
 *         },
 *         {
 *             flags: 0,
 *             tag: "issue",
 *             value: "example.net",
 *         },
 *         {
 *             flags: 0,
 *             tag: "issuewild",
 *             value: ";",
 *         },
 *         {
 *             flags: 0,
 *             tag: "iodef",
 *             value: "mailto:user@nonexisting.tld",
 *         },
 *     ],
 *     tags: {
 *         Environment: "Production",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * CAA records can be imported using the `resource id`, e.g.
 */
export class CaaRecord extends pulumi.CustomResource {
    /**
     * Get an existing CaaRecord resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CaaRecordState, opts?: pulumi.CustomResourceOptions): CaaRecord {
        return new CaaRecord(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:dns/caaRecord:CaaRecord';

    /**
     * Returns true if the given object is an instance of CaaRecord.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CaaRecord {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CaaRecord.__pulumiType;
    }

    /**
     * The FQDN of the DNS CAA Record.
     */
    public /*out*/ readonly fqdn!: pulumi.Output<string>;
    /**
     * The name of the DNS CAA Record. If you are creating the record in the apex of the zone use `"@"` as the name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of values that make up the CAA record. Each `record` block supports fields documented below.
     */
    public readonly records!: pulumi.Output<outputs.dns.CaaRecordRecord[]>;
    /**
     * Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The Time To Live (TTL) of the DNS record in seconds.
     */
    public readonly ttl!: pulumi.Output<number>;
    /**
     * Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly zoneName!: pulumi.Output<string>;

    /**
     * Create a CaaRecord resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CaaRecordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CaaRecordArgs | CaaRecordState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CaaRecordState | undefined;
            inputs["fqdn"] = state ? state.fqdn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["records"] = state ? state.records : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
            inputs["zoneName"] = state ? state.zoneName : undefined;
        } else {
            const args = argsOrState as CaaRecordArgs | undefined;
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
        super(CaaRecord.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CaaRecord resources.
 */
export interface CaaRecordState {
    /**
     * The FQDN of the DNS CAA Record.
     */
    readonly fqdn?: pulumi.Input<string>;
    /**
     * The name of the DNS CAA Record. If you are creating the record in the apex of the zone use `"@"` as the name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of values that make up the CAA record. Each `record` block supports fields documented below.
     */
    readonly records?: pulumi.Input<pulumi.Input<inputs.dns.CaaRecordRecord>[]>;
    /**
     * Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Time To Live (TTL) of the DNS record in seconds.
     */
    readonly ttl?: pulumi.Input<number>;
    /**
     * Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
     */
    readonly zoneName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CaaRecord resource.
 */
export interface CaaRecordArgs {
    /**
     * The name of the DNS CAA Record. If you are creating the record in the apex of the zone use `"@"` as the name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of values that make up the CAA record. Each `record` block supports fields documented below.
     */
    readonly records: pulumi.Input<pulumi.Input<inputs.dns.CaaRecordRecord>[]>;
    /**
     * Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Time To Live (TTL) of the DNS record in seconds.
     */
    readonly ttl: pulumi.Input<number>;
    /**
     * Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
     */
    readonly zoneName: pulumi.Input<string>;
}

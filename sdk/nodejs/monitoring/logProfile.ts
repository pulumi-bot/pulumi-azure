// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a [Log Profile](https://docs.microsoft.com/en-us/azure/monitoring-and-diagnostics/monitoring-overview-activity-logs#export-the-activity-log-with-a-log-profile). A Log Profile configures how Activity Logs are exported.
 *
 * > **NOTE:** It's only possible to configure one Log Profile per Subscription. If you are trying to create more than one Log Profile, an error with `StatusCode=409` will occur.
 */
export class LogProfile extends pulumi.CustomResource {
    /**
     * Get an existing LogProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogProfileState, opts?: pulumi.CustomResourceOptions): LogProfile {
        return new LogProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:monitoring/logProfile:LogProfile';

    /**
     * Returns true if the given object is an instance of LogProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogProfile.__pulumiType;
    }

    /**
     * List of categories of the logs.
     */
    public readonly categories!: pulumi.Output<string[]>;
    /**
     * List of regions for which Activity Log events are stored or streamed.
     */
    public readonly locations!: pulumi.Output<string[]>;
    /**
     * The name of the Log Profile. Changing this forces a
     * new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A `retentionPolicy` block as documented below. A retention policy for how long Activity Logs are retained in the storage account.
     */
    public readonly retentionPolicy!: pulumi.Output<outputs.monitoring.LogProfileRetentionPolicy>;
    /**
     * The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to. At least one of `storageAccountId` or `servicebusRuleId` must be set.
     */
    public readonly servicebusRuleId!: pulumi.Output<string | undefined>;
    /**
     * The resource ID of the storage account in which the Activity Log is stored. At least one of `storageAccountId` or `servicebusRuleId` must be set.
     */
    public readonly storageAccountId!: pulumi.Output<string | undefined>;

    /**
     * Create a LogProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LogProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogProfileArgs | LogProfileState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LogProfileState | undefined;
            inputs["categories"] = state ? state.categories : undefined;
            inputs["locations"] = state ? state.locations : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["retentionPolicy"] = state ? state.retentionPolicy : undefined;
            inputs["servicebusRuleId"] = state ? state.servicebusRuleId : undefined;
            inputs["storageAccountId"] = state ? state.storageAccountId : undefined;
        } else {
            const args = argsOrState as LogProfileArgs | undefined;
            if (!args || args.categories === undefined) {
                throw new Error("Missing required property 'categories'");
            }
            if (!args || args.locations === undefined) {
                throw new Error("Missing required property 'locations'");
            }
            if (!args || args.retentionPolicy === undefined) {
                throw new Error("Missing required property 'retentionPolicy'");
            }
            inputs["categories"] = args ? args.categories : undefined;
            inputs["locations"] = args ? args.locations : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["retentionPolicy"] = args ? args.retentionPolicy : undefined;
            inputs["servicebusRuleId"] = args ? args.servicebusRuleId : undefined;
            inputs["storageAccountId"] = args ? args.storageAccountId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LogProfile.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogProfile resources.
 */
export interface LogProfileState {
    /**
     * List of categories of the logs.
     */
    readonly categories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of regions for which Activity Log events are stored or streamed.
     */
    readonly locations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the Log Profile. Changing this forces a
     * new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `retentionPolicy` block as documented below. A retention policy for how long Activity Logs are retained in the storage account.
     */
    readonly retentionPolicy?: pulumi.Input<inputs.monitoring.LogProfileRetentionPolicy>;
    /**
     * The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to. At least one of `storageAccountId` or `servicebusRuleId` must be set.
     */
    readonly servicebusRuleId?: pulumi.Input<string>;
    /**
     * The resource ID of the storage account in which the Activity Log is stored. At least one of `storageAccountId` or `servicebusRuleId` must be set.
     */
    readonly storageAccountId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogProfile resource.
 */
export interface LogProfileArgs {
    /**
     * List of categories of the logs.
     */
    readonly categories: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of regions for which Activity Log events are stored or streamed.
     */
    readonly locations: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the Log Profile. Changing this forces a
     * new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `retentionPolicy` block as documented below. A retention policy for how long Activity Logs are retained in the storage account.
     */
    readonly retentionPolicy: pulumi.Input<inputs.monitoring.LogProfileRetentionPolicy>;
    /**
     * The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to. At least one of `storageAccountId` or `servicebusRuleId` must be set.
     */
    readonly servicebusRuleId?: pulumi.Input<string>;
    /**
     * The resource ID of the storage account in which the Activity Log is stored. At least one of `storageAccountId` or `servicebusRuleId` must be set.
     */
    readonly storageAccountId?: pulumi.Input<string>;
}

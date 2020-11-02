// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a ServiceBus Subscription Rule.
 *
 * ## Example Usage
 * ### SQL Filter)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleNamespace = new azure.servicebus.Namespace("exampleNamespace", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "Standard",
 *     tags: {
 *         source: "example",
 *     },
 * });
 * const exampleTopic = new azure.servicebus.Topic("exampleTopic", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     namespaceName: exampleNamespace.name,
 *     enablePartitioning: true,
 * });
 * const exampleSubscription = new azure.servicebus.Subscription("exampleSubscription", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     namespaceName: exampleNamespace.name,
 *     topicName: exampleTopic.name,
 *     maxDeliveryCount: 1,
 * });
 * const exampleSubscriptionRule = new azure.servicebus.SubscriptionRule("exampleSubscriptionRule", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     namespaceName: exampleNamespace.name,
 *     topicName: exampleTopic.name,
 *     subscriptionName: exampleSubscription.name,
 *     filterType: "SqlFilter",
 *     sqlFilter: "colour = 'red'",
 * });
 * ```
 * ### Correlation Filter)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleNamespace = new azure.servicebus.Namespace("exampleNamespace", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "Standard",
 *     tags: {
 *         source: "example",
 *     },
 * });
 * const exampleTopic = new azure.servicebus.Topic("exampleTopic", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     namespaceName: exampleNamespace.name,
 *     enablePartitioning: true,
 * });
 * const exampleSubscription = new azure.servicebus.Subscription("exampleSubscription", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     namespaceName: exampleNamespace.name,
 *     topicName: exampleTopic.name,
 *     maxDeliveryCount: 1,
 * });
 * const exampleSubscriptionRule = new azure.servicebus.SubscriptionRule("exampleSubscriptionRule", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     namespaceName: exampleNamespace.name,
 *     topicName: exampleTopic.name,
 *     subscriptionName: exampleSubscription.name,
 *     filterType: "CorrelationFilter",
 *     correlationFilter: {
 *         correlationId: "high",
 *         label: "red",
 *         properties: {
 *             customProperty: "value",
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Service Bus Subscription Rule can be imported using the `resource id`, e.g.
 */
export class SubscriptionRule extends pulumi.CustomResource {
    /**
     * Get an existing SubscriptionRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubscriptionRuleState, opts?: pulumi.CustomResourceOptions): SubscriptionRule {
        return new SubscriptionRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:servicebus/subscriptionRule:SubscriptionRule';

    /**
     * Returns true if the given object is an instance of SubscriptionRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SubscriptionRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SubscriptionRule.__pulumiType;
    }

    /**
     * Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
     */
    public readonly action!: pulumi.Output<string | undefined>;
    /**
     * A `correlationFilter` block as documented below to be evaluated against a BrokeredMessage. Required when `filterType` is set to `CorrelationFilter`.
     */
    public readonly correlationFilter!: pulumi.Output<outputs.servicebus.SubscriptionRuleCorrelationFilter | undefined>;
    /**
     * Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
     */
    public readonly filterType!: pulumi.Output<string>;
    /**
     * Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the ServiceBus Namespace in which the ServiceBus Topic exists. Changing this forces a new resource to be created.
     */
    public readonly namespaceName!: pulumi.Output<string>;
    /**
     * The name of the resource group in the ServiceBus Namespace exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filterType` is set to `SqlFilter`.
     */
    public readonly sqlFilter!: pulumi.Output<string | undefined>;
    /**
     * The name of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
     */
    public readonly subscriptionName!: pulumi.Output<string>;
    /**
     * The name of the ServiceBus Topic in which the ServiceBus Subscription exists. Changing this forces a new resource to be created.
     */
    public readonly topicName!: pulumi.Output<string>;

    /**
     * Create a SubscriptionRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubscriptionRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubscriptionRuleArgs | SubscriptionRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SubscriptionRuleState | undefined;
            inputs["action"] = state ? state.action : undefined;
            inputs["correlationFilter"] = state ? state.correlationFilter : undefined;
            inputs["filterType"] = state ? state.filterType : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namespaceName"] = state ? state.namespaceName : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sqlFilter"] = state ? state.sqlFilter : undefined;
            inputs["subscriptionName"] = state ? state.subscriptionName : undefined;
            inputs["topicName"] = state ? state.topicName : undefined;
        } else {
            const args = argsOrState as SubscriptionRuleArgs | undefined;
            if (!args || args.filterType === undefined) {
                throw new Error("Missing required property 'filterType'");
            }
            if (!args || args.namespaceName === undefined) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.subscriptionName === undefined) {
                throw new Error("Missing required property 'subscriptionName'");
            }
            if (!args || args.topicName === undefined) {
                throw new Error("Missing required property 'topicName'");
            }
            inputs["action"] = args ? args.action : undefined;
            inputs["correlationFilter"] = args ? args.correlationFilter : undefined;
            inputs["filterType"] = args ? args.filterType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sqlFilter"] = args ? args.sqlFilter : undefined;
            inputs["subscriptionName"] = args ? args.subscriptionName : undefined;
            inputs["topicName"] = args ? args.topicName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure:eventhub/subscriptionRule:SubscriptionRule" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(SubscriptionRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SubscriptionRule resources.
 */
export interface SubscriptionRuleState {
    /**
     * Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
     */
    readonly action?: pulumi.Input<string>;
    /**
     * A `correlationFilter` block as documented below to be evaluated against a BrokeredMessage. Required when `filterType` is set to `CorrelationFilter`.
     */
    readonly correlationFilter?: pulumi.Input<inputs.servicebus.SubscriptionRuleCorrelationFilter>;
    /**
     * Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
     */
    readonly filterType?: pulumi.Input<string>;
    /**
     * Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the ServiceBus Namespace in which the ServiceBus Topic exists. Changing this forces a new resource to be created.
     */
    readonly namespaceName?: pulumi.Input<string>;
    /**
     * The name of the resource group in the ServiceBus Namespace exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filterType` is set to `SqlFilter`.
     */
    readonly sqlFilter?: pulumi.Input<string>;
    /**
     * The name of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
     */
    readonly subscriptionName?: pulumi.Input<string>;
    /**
     * The name of the ServiceBus Topic in which the ServiceBus Subscription exists. Changing this forces a new resource to be created.
     */
    readonly topicName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SubscriptionRule resource.
 */
export interface SubscriptionRuleArgs {
    /**
     * Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
     */
    readonly action?: pulumi.Input<string>;
    /**
     * A `correlationFilter` block as documented below to be evaluated against a BrokeredMessage. Required when `filterType` is set to `CorrelationFilter`.
     */
    readonly correlationFilter?: pulumi.Input<inputs.servicebus.SubscriptionRuleCorrelationFilter>;
    /**
     * Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
     */
    readonly filterType: pulumi.Input<string>;
    /**
     * Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the ServiceBus Namespace in which the ServiceBus Topic exists. Changing this forces a new resource to be created.
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * The name of the resource group in the ServiceBus Namespace exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filterType` is set to `SqlFilter`.
     */
    readonly sqlFilter?: pulumi.Input<string>;
    /**
     * The name of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
     */
    readonly subscriptionName: pulumi.Input<string>;
    /**
     * The name of the ServiceBus Topic in which the ServiceBus Subscription exists. Changing this forces a new resource to be created.
     */
    readonly topicName: pulumi.Input<string>;
}

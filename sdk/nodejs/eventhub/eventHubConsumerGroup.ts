// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Event Hubs Consumer Group as a nested resource within an Event Hub.
 *
 * @deprecated azure.eventhub.EventHubConsumerGroup has been deprecated in favor of azure.eventhub.ConsumerGroup
 */
export class EventHubConsumerGroup extends pulumi.CustomResource {
    /**
     * Get an existing EventHubConsumerGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventHubConsumerGroupState, opts?: pulumi.CustomResourceOptions): EventHubConsumerGroup {
        pulumi.log.warn("EventHubConsumerGroup is deprecated: azure.eventhub.EventHubConsumerGroup has been deprecated in favor of azure.eventhub.ConsumerGroup")
        return new EventHubConsumerGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:eventhub/eventHubConsumerGroup:EventHubConsumerGroup';

    /**
     * Returns true if the given object is an instance of EventHubConsumerGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventHubConsumerGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventHubConsumerGroup.__pulumiType;
    }

    /**
     * Specifies the name of the EventHub. Changing this forces a new resource to be created.
     */
    public readonly eventhubName!: pulumi.Output<string>;
    /**
     * Specifies the name of the EventHub Consumer Group resource. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
     */
    public readonly namespaceName!: pulumi.Output<string>;
    /**
     * The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Specifies the user metadata.
     */
    public readonly userMetadata!: pulumi.Output<string | undefined>;

    /**
     * Create a EventHubConsumerGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated azure.eventhub.EventHubConsumerGroup has been deprecated in favor of azure.eventhub.ConsumerGroup */
    constructor(name: string, args: EventHubConsumerGroupArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azure.eventhub.EventHubConsumerGroup has been deprecated in favor of azure.eventhub.ConsumerGroup */
    constructor(name: string, argsOrState?: EventHubConsumerGroupArgs | EventHubConsumerGroupState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("EventHubConsumerGroup is deprecated: azure.eventhub.EventHubConsumerGroup has been deprecated in favor of azure.eventhub.ConsumerGroup")
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EventHubConsumerGroupState | undefined;
            inputs["eventhubName"] = state ? state.eventhubName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namespaceName"] = state ? state.namespaceName : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["userMetadata"] = state ? state.userMetadata : undefined;
        } else {
            const args = argsOrState as EventHubConsumerGroupArgs | undefined;
            if (!args || args.eventhubName === undefined) {
                throw new Error("Missing required property 'eventhubName'");
            }
            if (!args || args.namespaceName === undefined) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["eventhubName"] = args ? args.eventhubName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["userMetadata"] = args ? args.userMetadata : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EventHubConsumerGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventHubConsumerGroup resources.
 */
export interface EventHubConsumerGroupState {
    /**
     * Specifies the name of the EventHub. Changing this forces a new resource to be created.
     */
    readonly eventhubName?: pulumi.Input<string>;
    /**
     * Specifies the name of the EventHub Consumer Group resource. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
     */
    readonly namespaceName?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies the user metadata.
     */
    readonly userMetadata?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EventHubConsumerGroup resource.
 */
export interface EventHubConsumerGroupArgs {
    /**
     * Specifies the name of the EventHub. Changing this forces a new resource to be created.
     */
    readonly eventhubName: pulumi.Input<string>;
    /**
     * Specifies the name of the EventHub Consumer Group resource. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the user metadata.
     */
    readonly userMetadata?: pulumi.Input<string>;
}

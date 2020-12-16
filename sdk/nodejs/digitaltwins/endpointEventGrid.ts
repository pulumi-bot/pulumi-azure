// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Digital Twins Event Grid Endpoint.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleInstance = new azure.digitaltwins.Instance("exampleInstance", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 * });
 * const exampleTopic = new azure.eventgrid.Topic("exampleTopic", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleEndpointEventGrid = new azure.digitaltwins.EndpointEventGrid("exampleEndpointEventGrid", {
 *     digitalTwinsId: exampleInstance.id,
 *     eventgridTopicEndpoint: exampleTopic.endpoint,
 *     eventgridTopicPrimaryAccessKey: exampleTopic.primaryAccessKey,
 *     eventgridTopicSecondaryAccessKey: exampleTopic.secondaryAccessKey,
 * });
 * ```
 *
 * ## Import
 *
 * Digital Twins Eventgrid Endpoints can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:digitaltwins/endpointEventGrid:EndpointEventGrid example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DigitalTwins/digitalTwinsInstances/dt1/endpoints/ep1
 * ```
 */
export class EndpointEventGrid extends pulumi.CustomResource {
    /**
     * Get an existing EndpointEventGrid resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointEventGridState, opts?: pulumi.CustomResourceOptions): EndpointEventGrid {
        return new EndpointEventGrid(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:digitaltwins/endpointEventGrid:EndpointEventGrid';

    /**
     * Returns true if the given object is an instance of EndpointEventGrid.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EndpointEventGrid {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EndpointEventGrid.__pulumiType;
    }

    /**
     * The storage secret of the dead-lettering, whose format is `https://<storageAccountname>.blob.core.windows.net/<containerName>?<SASToken>`. When an endpoint can't deliver an event within a certain time period or after trying to deliver the event a certain number of times, it can send the undelivered event to a storage account.
     */
    public readonly deadLetterStorageSecret!: pulumi.Output<string | undefined>;
    /**
     * The resource ID of the Digital Twins Instance. Changing this forces a new Digital Twins Eventgrid Endpoint to be created.
     */
    public readonly digitalTwinsId!: pulumi.Output<string>;
    /**
     * The endpoint of the Event Grid Topic.
     */
    public readonly eventgridTopicEndpoint!: pulumi.Output<string>;
    /**
     * The primary access key of the Event Grid Topic.
     */
    public readonly eventgridTopicPrimaryAccessKey!: pulumi.Output<string>;
    /**
     * The secondary access key of the Event Grid Topic.
     */
    public readonly eventgridTopicSecondaryAccessKey!: pulumi.Output<string>;
    /**
     * The name which should be used for this Digital Twins Eventgrid Endpoint. Changing this forces a new Digital Twins Eventgrid Endpoint to be created.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a EndpointEventGrid resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndpointEventGridArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndpointEventGridArgs | EndpointEventGridState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EndpointEventGridState | undefined;
            inputs["deadLetterStorageSecret"] = state ? state.deadLetterStorageSecret : undefined;
            inputs["digitalTwinsId"] = state ? state.digitalTwinsId : undefined;
            inputs["eventgridTopicEndpoint"] = state ? state.eventgridTopicEndpoint : undefined;
            inputs["eventgridTopicPrimaryAccessKey"] = state ? state.eventgridTopicPrimaryAccessKey : undefined;
            inputs["eventgridTopicSecondaryAccessKey"] = state ? state.eventgridTopicSecondaryAccessKey : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as EndpointEventGridArgs | undefined;
            if ((!args || args.digitalTwinsId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'digitalTwinsId'");
            }
            if ((!args || args.eventgridTopicEndpoint === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'eventgridTopicEndpoint'");
            }
            if ((!args || args.eventgridTopicPrimaryAccessKey === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'eventgridTopicPrimaryAccessKey'");
            }
            if ((!args || args.eventgridTopicSecondaryAccessKey === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'eventgridTopicSecondaryAccessKey'");
            }
            inputs["deadLetterStorageSecret"] = args ? args.deadLetterStorageSecret : undefined;
            inputs["digitalTwinsId"] = args ? args.digitalTwinsId : undefined;
            inputs["eventgridTopicEndpoint"] = args ? args.eventgridTopicEndpoint : undefined;
            inputs["eventgridTopicPrimaryAccessKey"] = args ? args.eventgridTopicPrimaryAccessKey : undefined;
            inputs["eventgridTopicSecondaryAccessKey"] = args ? args.eventgridTopicSecondaryAccessKey : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EndpointEventGrid.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EndpointEventGrid resources.
 */
export interface EndpointEventGridState {
    /**
     * The storage secret of the dead-lettering, whose format is `https://<storageAccountname>.blob.core.windows.net/<containerName>?<SASToken>`. When an endpoint can't deliver an event within a certain time period or after trying to deliver the event a certain number of times, it can send the undelivered event to a storage account.
     */
    readonly deadLetterStorageSecret?: pulumi.Input<string>;
    /**
     * The resource ID of the Digital Twins Instance. Changing this forces a new Digital Twins Eventgrid Endpoint to be created.
     */
    readonly digitalTwinsId?: pulumi.Input<string>;
    /**
     * The endpoint of the Event Grid Topic.
     */
    readonly eventgridTopicEndpoint?: pulumi.Input<string>;
    /**
     * The primary access key of the Event Grid Topic.
     */
    readonly eventgridTopicPrimaryAccessKey?: pulumi.Input<string>;
    /**
     * The secondary access key of the Event Grid Topic.
     */
    readonly eventgridTopicSecondaryAccessKey?: pulumi.Input<string>;
    /**
     * The name which should be used for this Digital Twins Eventgrid Endpoint. Changing this forces a new Digital Twins Eventgrid Endpoint to be created.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EndpointEventGrid resource.
 */
export interface EndpointEventGridArgs {
    /**
     * The storage secret of the dead-lettering, whose format is `https://<storageAccountname>.blob.core.windows.net/<containerName>?<SASToken>`. When an endpoint can't deliver an event within a certain time period or after trying to deliver the event a certain number of times, it can send the undelivered event to a storage account.
     */
    readonly deadLetterStorageSecret?: pulumi.Input<string>;
    /**
     * The resource ID of the Digital Twins Instance. Changing this forces a new Digital Twins Eventgrid Endpoint to be created.
     */
    readonly digitalTwinsId: pulumi.Input<string>;
    /**
     * The endpoint of the Event Grid Topic.
     */
    readonly eventgridTopicEndpoint: pulumi.Input<string>;
    /**
     * The primary access key of the Event Grid Topic.
     */
    readonly eventgridTopicPrimaryAccessKey: pulumi.Input<string>;
    /**
     * The secondary access key of the Event Grid Topic.
     */
    readonly eventgridTopicSecondaryAccessKey: pulumi.Input<string>;
    /**
     * The name which should be used for this Digital Twins Eventgrid Endpoint. Changing this forces a new Digital Twins Eventgrid Endpoint to be created.
     */
    readonly name?: pulumi.Input<string>;
}
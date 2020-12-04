// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Application Insights WebTest.
 *
 * ## Import
 *
 * Application Insights Web Tests can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:appinsights/webTest:WebTest my_test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.insights/webtests/my_test
 * ```
 */
export class WebTest extends pulumi.CustomResource {
    /**
     * Get an existing WebTest resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebTestState, opts?: pulumi.CustomResourceOptions): WebTest {
        return new WebTest(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:appinsights/webTest:WebTest';

    /**
     * Returns true if the given object is an instance of WebTest.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebTest {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebTest.__pulumiType;
    }

    /**
     * The ID of the Application Insights component on which the WebTest operates. Changing this forces a new resource to be created.
     */
    public readonly applicationInsightsId!: pulumi.Output<string>;
    /**
     * An XML configuration specification for a WebTest.
     */
    public readonly configuration!: pulumi.Output<string>;
    /**
     * Purpose/user defined descriptive test for this WebTest.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Is the test actively being monitored.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Interval in seconds between test runs for this WebTest. Default is `300`.
     */
    public readonly frequency!: pulumi.Output<number | undefined>;
    /**
     * A list of where to physically run the tests from to give global coverage for accessibility of your application.
     */
    public readonly geoLocations!: pulumi.Output<string[]>;
    /**
     * = (Required) The kind of web test that this web test watches. Choices are `ping` and `multistep`.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * The location of the resource group.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Application Insights WebTest. Changing this forces a
     * new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Allow for retries should this WebTest fail.
     */
    public readonly retryEnabled!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly syntheticMonitorId!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Seconds until this WebTest will timeout and fail. Default is `30`.
     */
    public readonly timeout!: pulumi.Output<number | undefined>;

    /**
     * Create a WebTest resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebTestArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebTestArgs | WebTestState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as WebTestState | undefined;
            inputs["applicationInsightsId"] = state ? state.applicationInsightsId : undefined;
            inputs["configuration"] = state ? state.configuration : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["frequency"] = state ? state.frequency : undefined;
            inputs["geoLocations"] = state ? state.geoLocations : undefined;
            inputs["kind"] = state ? state.kind : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["retryEnabled"] = state ? state.retryEnabled : undefined;
            inputs["syntheticMonitorId"] = state ? state.syntheticMonitorId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["timeout"] = state ? state.timeout : undefined;
        } else {
            const args = argsOrState as WebTestArgs | undefined;
            if ((!args || args.applicationInsightsId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'applicationInsightsId'");
            }
            if ((!args || args.configuration === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'configuration'");
            }
            if ((!args || args.geoLocations === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'geoLocations'");
            }
            if ((!args || args.kind === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'kind'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["applicationInsightsId"] = args ? args.applicationInsightsId : undefined;
            inputs["configuration"] = args ? args.configuration : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["frequency"] = args ? args.frequency : undefined;
            inputs["geoLocations"] = args ? args.geoLocations : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["retryEnabled"] = args ? args.retryEnabled : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["timeout"] = args ? args.timeout : undefined;
            inputs["syntheticMonitorId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(WebTest.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebTest resources.
 */
export interface WebTestState {
    /**
     * The ID of the Application Insights component on which the WebTest operates. Changing this forces a new resource to be created.
     */
    readonly applicationInsightsId?: pulumi.Input<string>;
    /**
     * An XML configuration specification for a WebTest.
     */
    readonly configuration?: pulumi.Input<string>;
    /**
     * Purpose/user defined descriptive test for this WebTest.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Is the test actively being monitored.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Interval in seconds between test runs for this WebTest. Default is `300`.
     */
    readonly frequency?: pulumi.Input<number>;
    /**
     * A list of where to physically run the tests from to give global coverage for accessibility of your application.
     */
    readonly geoLocations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * = (Required) The kind of web test that this web test watches. Choices are `ping` and `multistep`.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * The location of the resource group.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Application Insights WebTest. Changing this forces a
     * new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Allow for retries should this WebTest fail.
     */
    readonly retryEnabled?: pulumi.Input<boolean>;
    readonly syntheticMonitorId?: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Seconds until this WebTest will timeout and fail. Default is `30`.
     */
    readonly timeout?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a WebTest resource.
 */
export interface WebTestArgs {
    /**
     * The ID of the Application Insights component on which the WebTest operates. Changing this forces a new resource to be created.
     */
    readonly applicationInsightsId: pulumi.Input<string>;
    /**
     * An XML configuration specification for a WebTest.
     */
    readonly configuration: pulumi.Input<string>;
    /**
     * Purpose/user defined descriptive test for this WebTest.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Is the test actively being monitored.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Interval in seconds between test runs for this WebTest. Default is `300`.
     */
    readonly frequency?: pulumi.Input<number>;
    /**
     * A list of where to physically run the tests from to give global coverage for accessibility of your application.
     */
    readonly geoLocations: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * = (Required) The kind of web test that this web test watches. Choices are `ping` and `multistep`.
     */
    readonly kind: pulumi.Input<string>;
    /**
     * The location of the resource group.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Application Insights WebTest. Changing this forces a
     * new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Allow for retries should this WebTest fail.
     */
    readonly retryEnabled?: pulumi.Input<boolean>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Seconds until this WebTest will timeout and fail. Default is `30`.
     */
    readonly timeout?: pulumi.Input<number>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Azure Spring Cloud Deployment with a Java runtime.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleSpringCloudService = new azure.appplatform.SpringCloudService("exampleSpringCloudService", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 * });
 * const exampleSpringCloudApp = new azure.appplatform.SpringCloudApp("exampleSpringCloudApp", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     serviceName: exampleSpringCloudService.name,
 *     identity: {
 *         type: "SystemAssigned",
 *     },
 * });
 * const exampleSpringCloudJavaDeployment = new azure.appplatform.SpringCloudJavaDeployment("exampleSpringCloudJavaDeployment", {
 *     springCloudAppId: exampleSpringCloudApp.id,
 *     cpu: 2,
 *     instanceCount: 2,
 *     jvmOptions: "-XX:+PrintGC",
 *     memoryInGb: 4,
 *     runtimeVersion: "Java_11",
 *     environmentVariables: {
 *         Foo: "Bar",
 *         Env: "Staging",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Spring Cloud Deployment can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:appplatform/springCloudJavaDeployment:SpringCloudJavaDeployment example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/Spring/service1/apps/app1/deployments/deploy1
 * ```
 */
export class SpringCloudJavaDeployment extends pulumi.CustomResource {
    /**
     * Get an existing SpringCloudJavaDeployment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SpringCloudJavaDeploymentState, opts?: pulumi.CustomResourceOptions): SpringCloudJavaDeployment {
        return new SpringCloudJavaDeployment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:appplatform/springCloudJavaDeployment:SpringCloudJavaDeployment';

    /**
     * Returns true if the given object is an instance of SpringCloudJavaDeployment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SpringCloudJavaDeployment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SpringCloudJavaDeployment.__pulumiType;
    }

    /**
     * Specifies the required cpu of the Spring Cloud Deployment. Possible Values are between `1` and `4`. Defaults to `1` if not specified.
     */
    public readonly cpu!: pulumi.Output<number | undefined>;
    /**
     * Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
     */
    public readonly environmentVariables!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between `1` and `500`. Defaults to `1` if not specified.
     */
    public readonly instanceCount!: pulumi.Output<number | undefined>;
    /**
     * Specifies the jvm option of the Spring Cloud Deployment.
     */
    public readonly jvmOptions!: pulumi.Output<string | undefined>;
    /**
     * Specifies the required memory size of the Spring Cloud Deployment. Possible Values are between `1` and `8`. Defaults to `1` if not specified.
     */
    public readonly memoryInGb!: pulumi.Output<number | undefined>;
    /**
     * Specifies the name of the Spring Cloud Deployment. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the runtime version of the Spring Cloud Deployment. Possible Values are `Java_8` and `Java_11`. Defaults to `Java_8`.
     */
    public readonly runtimeVersion!: pulumi.Output<string | undefined>;
    /**
     * Specifies the id of the Spring Cloud Application in which to create the Deployment. Changing this forces a new resource to be created.
     */
    public readonly springCloudAppId!: pulumi.Output<string>;

    /**
     * Create a SpringCloudJavaDeployment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SpringCloudJavaDeploymentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SpringCloudJavaDeploymentArgs | SpringCloudJavaDeploymentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SpringCloudJavaDeploymentState | undefined;
            inputs["cpu"] = state ? state.cpu : undefined;
            inputs["environmentVariables"] = state ? state.environmentVariables : undefined;
            inputs["instanceCount"] = state ? state.instanceCount : undefined;
            inputs["jvmOptions"] = state ? state.jvmOptions : undefined;
            inputs["memoryInGb"] = state ? state.memoryInGb : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["runtimeVersion"] = state ? state.runtimeVersion : undefined;
            inputs["springCloudAppId"] = state ? state.springCloudAppId : undefined;
        } else {
            const args = argsOrState as SpringCloudJavaDeploymentArgs | undefined;
            if ((!args || args.springCloudAppId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'springCloudAppId'");
            }
            inputs["cpu"] = args ? args.cpu : undefined;
            inputs["environmentVariables"] = args ? args.environmentVariables : undefined;
            inputs["instanceCount"] = args ? args.instanceCount : undefined;
            inputs["jvmOptions"] = args ? args.jvmOptions : undefined;
            inputs["memoryInGb"] = args ? args.memoryInGb : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["runtimeVersion"] = args ? args.runtimeVersion : undefined;
            inputs["springCloudAppId"] = args ? args.springCloudAppId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SpringCloudJavaDeployment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SpringCloudJavaDeployment resources.
 */
export interface SpringCloudJavaDeploymentState {
    /**
     * Specifies the required cpu of the Spring Cloud Deployment. Possible Values are between `1` and `4`. Defaults to `1` if not specified.
     */
    readonly cpu?: pulumi.Input<number>;
    /**
     * Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
     */
    readonly environmentVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between `1` and `500`. Defaults to `1` if not specified.
     */
    readonly instanceCount?: pulumi.Input<number>;
    /**
     * Specifies the jvm option of the Spring Cloud Deployment.
     */
    readonly jvmOptions?: pulumi.Input<string>;
    /**
     * Specifies the required memory size of the Spring Cloud Deployment. Possible Values are between `1` and `8`. Defaults to `1` if not specified.
     */
    readonly memoryInGb?: pulumi.Input<number>;
    /**
     * Specifies the name of the Spring Cloud Deployment. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the runtime version of the Spring Cloud Deployment. Possible Values are `Java_8` and `Java_11`. Defaults to `Java_8`.
     */
    readonly runtimeVersion?: pulumi.Input<string>;
    /**
     * Specifies the id of the Spring Cloud Application in which to create the Deployment. Changing this forces a new resource to be created.
     */
    readonly springCloudAppId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SpringCloudJavaDeployment resource.
 */
export interface SpringCloudJavaDeploymentArgs {
    /**
     * Specifies the required cpu of the Spring Cloud Deployment. Possible Values are between `1` and `4`. Defaults to `1` if not specified.
     */
    readonly cpu?: pulumi.Input<number>;
    /**
     * Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
     */
    readonly environmentVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between `1` and `500`. Defaults to `1` if not specified.
     */
    readonly instanceCount?: pulumi.Input<number>;
    /**
     * Specifies the jvm option of the Spring Cloud Deployment.
     */
    readonly jvmOptions?: pulumi.Input<string>;
    /**
     * Specifies the required memory size of the Spring Cloud Deployment. Possible Values are between `1` and `8`. Defaults to `1` if not specified.
     */
    readonly memoryInGb?: pulumi.Input<number>;
    /**
     * Specifies the name of the Spring Cloud Deployment. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the runtime version of the Spring Cloud Deployment. Possible Values are `Java_8` and `Java_11`. Defaults to `Java_8`.
     */
    readonly runtimeVersion?: pulumi.Input<string>;
    /**
     * Specifies the id of the Spring Cloud Application in which to create the Deployment. Changing this forces a new resource to be created.
     */
    readonly springCloudAppId: pulumi.Input<string>;
}

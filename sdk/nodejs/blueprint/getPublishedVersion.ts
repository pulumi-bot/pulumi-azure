// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Blueprint Published Version
 *
 * > **NOTE:** Azure Blueprints are in Preview and potentially subject to breaking change without notice.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const current = azure.core.getSubscription({});
 * const test = current.then(current => azure.blueprint.getPublishedVersion({
 *     scopeId: current.id,
 *     blueprintName: "exampleBluePrint",
 *     version: "dev_v2.3",
 * }));
 * ```
 */
export function getPublishedVersion(args: GetPublishedVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetPublishedVersionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:blueprint/getPublishedVersion:getPublishedVersion", {
        "blueprintName": args.blueprintName,
        "scopeId": args.scopeId,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getPublishedVersion.
 */
export interface GetPublishedVersionArgs {
    /**
     * The name of the Blueprint Definition
     */
    blueprintName: string;
    /**
     * The ID of the Management Group / Subscription where this Blueprint Definition is stored.
     */
    scopeId: string;
    /**
     * The Version name of the Published Version of the Blueprint Definition
     */
    version: string;
}

/**
 * A collection of values returned by getPublishedVersion.
 */
export interface GetPublishedVersionResult {
    readonly blueprintName: string;
    /**
     * The description of the Blueprint Published Version
     */
    readonly description: string;
    /**
     * The display name of the Blueprint Published Version
     */
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lastModified: string;
    readonly scopeId: string;
    /**
     * The target scope
     */
    readonly targetScope: string;
    readonly timeCreated: string;
    /**
     * The type of the Blueprint
     */
    readonly type: string;
    readonly version: string;
}

export function getPublishedVersionOutput(args: GetPublishedVersionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPublishedVersionResult> {
    return pulumi.output(args).apply(a => getPublishedVersion(a, opts))
}

/**
 * A collection of arguments for invoking getPublishedVersion.
 */
export interface GetPublishedVersionOutputArgs {
    /**
     * The name of the Blueprint Definition
     */
    blueprintName: pulumi.Input<string>;
    /**
     * The ID of the Management Group / Subscription where this Blueprint Definition is stored.
     */
    scopeId: pulumi.Input<string>;
    /**
     * The Version name of the Published Version of the Blueprint Definition
     */
    version: pulumi.Input<string>;
}

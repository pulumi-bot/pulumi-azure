// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Azure Blueprint Definition
 *
 * > **NOTE:** Azure Blueprints are in Preview and potentially subject to breaking change without notice.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const current = azure.core.getClientConfig({});
 * const root = current.then(current => azure.management.getGroup({
 *     name: current.tenantId,
 * }));
 * const example = root.then(root => azure.blueprint.getDefinition({
 *     name: "exampleManagementGroupBP",
 *     scopeId: root.id,
 * }));
 * ```
 */
export function getDefinition(args: GetDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetDefinitionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:blueprint/getDefinition:getDefinition", {
        "name": args.name,
        "scopeId": args.scopeId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDefinition.
 */
export interface GetDefinitionArgs {
    /**
     * The name of the Blueprint.
     */
    readonly name: string;
    /**
     * The ID of the Subscription or Management Group, as the scope at which the blueprint definition is stored.
     */
    readonly scopeId: string;
}

/**
 * A collection of values returned by getDefinition.
 */
export interface GetDefinitionResult {
    /**
     * The description of the Blueprint Definition.
     */
    readonly description: string;
    /**
     * The display name of the Blueprint Definition.
     */
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The timestamp of when this last modification was saved to the Blueprint Definition.
     */
    readonly lastModified: string;
    readonly name: string;
    readonly scopeId: string;
    /**
     * The target scope.
     */
    readonly targetScope: string;
    /**
     * The timestamp of when this Blueprint Definition was created.
     */
    readonly timeCreated: string;
    /**
     * A list of versions published for this Blueprint Definition.
     */
    readonly versions: string[];
}

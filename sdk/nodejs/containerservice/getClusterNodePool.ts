// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Kubernetes Cluster Node Pool.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.containerservice.getClusterNodePool({
 *     name: "existing",
 *     kubernetesClusterName: "existing-cluster",
 *     resourceGroupName: "existing-resource-group",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getClusterNodePool(args: GetClusterNodePoolArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterNodePoolResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:containerservice/getClusterNodePool:getClusterNodePool", {
        "kubernetesClusterName": args.kubernetesClusterName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getClusterNodePool.
 */
export interface GetClusterNodePoolArgs {
    /**
     * The Name of the Kubernetes Cluster where this Node Pool is located.
     */
    readonly kubernetesClusterName: string;
    /**
     * The name of this Kubernetes Cluster Node Pool.
     */
    readonly name: string;
    /**
     * The name of the Resource Group where the Kubernetes Cluster exists.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getClusterNodePool.
 */
export interface GetClusterNodePoolResult {
    /**
     * A list of Availability Zones in which the Nodes in this Node Pool exists.
     */
    readonly availabilityZones: string[];
    /**
     * Does this Node Pool have Auto-Scaling enabled?
     */
    readonly enableAutoScaling: boolean;
    /**
     * Do nodes in this Node Pool have a Public IP Address?
     */
    readonly enableNodePublicIp: boolean;
    /**
     * The eviction policy used for Virtual Machines in the Virtual Machine Scale Set, when `priority` is set to `Spot`.
     */
    readonly evictionPolicy: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly kubernetesClusterName: string;
    /**
     * The maximum number of Nodes allowed when auto-scaling is enabled.
     */
    readonly maxCount: number;
    /**
     * The maximum number of Pods allowed on each Node in this Node Pool.
     */
    readonly maxPods: number;
    /**
     * The minimum number of Nodes allowed when auto-scaling is enabled.
     */
    readonly minCount: number;
    /**
     * The Mode for this Node Pool, specifying how these Nodes should be used (for either System or User resources).
     */
    readonly mode: string;
    readonly name: string;
    /**
     * The current number of Nodes in the Node Pool.
     */
    readonly nodeCount: number;
    /**
     * A map of Kubernetes Labels applied to each Node in this Node Pool.
     */
    readonly nodeLabels: {[key: string]: string};
    /**
     * A map of Kubernetes Taints applied to each Node in this Node Pool.
     */
    readonly nodeTaints: string[];
    /**
     * The version of Kubernetes configured on each Node in this Node Pool.
     */
    readonly orchestratorVersion: string;
    /**
     * The size of the OS Disk on each Node in this Node Pool.
     */
    readonly osDiskSizeGb: number;
    /**
     * The operating system used on each Node in this Node Pool.
     */
    readonly osType: string;
    /**
     * The priority of the Virtual Machines in the Virtual Machine Scale Set backing this Node Pool.
     */
    readonly priority: string;
    readonly resourceGroupName: string;
    /**
     * The maximum price being paid for Virtual Machines in this Scale Set. `-1` means the current on-demand price for a Virtual Machine.
     */
    readonly spotMaxPrice: number;
    /**
     * A mapping of tags assigned to the Kubernetes Cluster Node Pool.
     */
    readonly tags: {[key: string]: string};
    /**
     * The size of the Virtual Machines used in the Virtual Machine Scale Set backing this Node Pool.
     */
    readonly vmSize: string;
    /**
     * The ID of the Subnet in which this Node Pool exists.
     */
    readonly vnetSubnetId: string;
}

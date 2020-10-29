// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Network Interface.
 */
export function getNetworkInterface(args: GetNetworkInterfaceArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkInterfaceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:network/getNetworkInterface:getNetworkInterface", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetworkInterface.
 */
export interface GetNetworkInterfaceArgs {
    /**
     * Specifies the name of the Network Interface.
     */
    readonly name: string;
    /**
     * Specifies the name of the resource group the Network Interface is located in.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getNetworkInterface.
 */
export interface GetNetworkInterfaceResult {
    /**
     * List of DNS servers applied to the specified Network Interface.
     */
    readonly appliedDnsServers: string[];
    /**
     * The list of DNS servers used by the specified Network Interface.
     */
    readonly dnsServers: string[];
    /**
     * Indicates if accelerated networking is set on the specified Network Interface.
     */
    readonly enableAcceleratedNetworking: boolean;
    /**
     * Indicate if IP forwarding is set on the specified Network Interface.
     */
    readonly enableIpForwarding: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The internal dns name label of the specified Network Interface.
     */
    readonly internalDnsNameLabel: string;
    /**
     * One or more `ipConfiguration` blocks as defined below.
     */
    readonly ipConfigurations: outputs.network.GetNetworkInterfaceIpConfiguration[];
    /**
     * The location of the specified Network Interface.
     */
    readonly location: string;
    /**
     * The MAC address used by the specified Network Interface.
     */
    readonly macAddress: string;
    /**
     * The name of the IP Configuration.
     */
    readonly name: string;
    /**
     * The ID of the network security group associated to the specified Network Interface.
     */
    readonly networkSecurityGroupId: string;
    /**
     * The Private IP Address assigned to this Network Interface.
     */
    readonly privateIpAddress: string;
    /**
     * The list of private ip addresses associates to the specified Network Interface.
     */
    readonly privateIpAddresses: string[];
    readonly resourceGroupName: string;
    /**
     * List the tags associated to the specified Network Interface.
     */
    readonly tags: {[key: string]: string};
    /**
     * The ID of the virtual machine that the specified Network Interface is attached to.
     */
    readonly virtualMachineId: string;
}

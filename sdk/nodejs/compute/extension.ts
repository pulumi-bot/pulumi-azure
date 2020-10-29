// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Virtual Machine Extension to provide post deployment configuration
 * and run automated tasks.
 *
 * > **NOTE:** Custom Script Extensions for Linux & Windows require that the `commandToExecute` returns a `0` exit code to be classified as successfully deployed. You can achieve this by appending `exit 0` to the end of your `commandToExecute`.
 *
 * > **NOTE:** Custom Script Extensions require that the Azure Virtual Machine Guest Agent is running on the Virtual Machine.
 */
export class Extension extends pulumi.CustomResource {
    /**
     * Get an existing Extension resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExtensionState, opts?: pulumi.CustomResourceOptions): Extension {
        return new Extension(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:compute/extension:Extension';

    /**
     * Returns true if the given object is an instance of Extension.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Extension {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Extension.__pulumiType;
    }

    /**
     * Specifies if the platform deploys
     * the latest minor version update to the `typeHandlerVersion` specified.
     */
    public readonly autoUpgradeMinorVersion!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the virtual machine extension peering. Changing
     * this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The protectedSettings passed to the
     * extension, like settings, these are specified as a JSON object in a string.
     */
    public readonly protectedSettings!: pulumi.Output<string | undefined>;
    /**
     * The publisher of the extension, available publishers
     * can be found by using the Azure CLI.
     */
    public readonly publisher!: pulumi.Output<string>;
    /**
     * The settings passed to the extension, these are
     * specified as a JSON object in a string.
     */
    public readonly settings!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of extension, available types for a publisher can
     * be found using the Azure CLI.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Specifies the version of the extension to
     * use, available versions can be found using the Azure CLI.
     */
    public readonly typeHandlerVersion!: pulumi.Output<string>;
    /**
     * The ID of the Virtual Machine. Changing this forces a new resource to be created
     */
    public readonly virtualMachineId!: pulumi.Output<string>;

    /**
     * Create a Extension resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExtensionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExtensionArgs | ExtensionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ExtensionState | undefined;
            inputs["autoUpgradeMinorVersion"] = state ? state.autoUpgradeMinorVersion : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["protectedSettings"] = state ? state.protectedSettings : undefined;
            inputs["publisher"] = state ? state.publisher : undefined;
            inputs["settings"] = state ? state.settings : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["typeHandlerVersion"] = state ? state.typeHandlerVersion : undefined;
            inputs["virtualMachineId"] = state ? state.virtualMachineId : undefined;
        } else {
            const args = argsOrState as ExtensionArgs | undefined;
            if (!args || args.publisher === undefined) {
                throw new Error("Missing required property 'publisher'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            if (!args || args.typeHandlerVersion === undefined) {
                throw new Error("Missing required property 'typeHandlerVersion'");
            }
            if (!args || args.virtualMachineId === undefined) {
                throw new Error("Missing required property 'virtualMachineId'");
            }
            inputs["autoUpgradeMinorVersion"] = args ? args.autoUpgradeMinorVersion : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["protectedSettings"] = args ? args.protectedSettings : undefined;
            inputs["publisher"] = args ? args.publisher : undefined;
            inputs["settings"] = args ? args.settings : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["typeHandlerVersion"] = args ? args.typeHandlerVersion : undefined;
            inputs["virtualMachineId"] = args ? args.virtualMachineId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Extension.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Extension resources.
 */
export interface ExtensionState {
    /**
     * Specifies if the platform deploys
     * the latest minor version update to the `typeHandlerVersion` specified.
     */
    readonly autoUpgradeMinorVersion?: pulumi.Input<boolean>;
    /**
     * The name of the virtual machine extension peering. Changing
     * this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The protectedSettings passed to the
     * extension, like settings, these are specified as a JSON object in a string.
     */
    readonly protectedSettings?: pulumi.Input<string>;
    /**
     * The publisher of the extension, available publishers
     * can be found by using the Azure CLI.
     */
    readonly publisher?: pulumi.Input<string>;
    /**
     * The settings passed to the extension, these are
     * specified as a JSON object in a string.
     */
    readonly settings?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of extension, available types for a publisher can
     * be found using the Azure CLI.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * Specifies the version of the extension to
     * use, available versions can be found using the Azure CLI.
     */
    readonly typeHandlerVersion?: pulumi.Input<string>;
    /**
     * The ID of the Virtual Machine. Changing this forces a new resource to be created
     */
    readonly virtualMachineId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Extension resource.
 */
export interface ExtensionArgs {
    /**
     * Specifies if the platform deploys
     * the latest minor version update to the `typeHandlerVersion` specified.
     */
    readonly autoUpgradeMinorVersion?: pulumi.Input<boolean>;
    /**
     * The name of the virtual machine extension peering. Changing
     * this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The protectedSettings passed to the
     * extension, like settings, these are specified as a JSON object in a string.
     */
    readonly protectedSettings?: pulumi.Input<string>;
    /**
     * The publisher of the extension, available publishers
     * can be found by using the Azure CLI.
     */
    readonly publisher: pulumi.Input<string>;
    /**
     * The settings passed to the extension, these are
     * specified as a JSON object in a string.
     */
    readonly settings?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of extension, available types for a publisher can
     * be found using the Azure CLI.
     */
    readonly type: pulumi.Input<string>;
    /**
     * Specifies the version of the extension to
     * use, available versions can be found using the Azure CLI.
     */
    readonly typeHandlerVersion: pulumi.Input<string>;
    /**
     * The ID of the Virtual Machine. Changing this forces a new resource to be created
     */
    readonly virtualMachineId: pulumi.Input<string>;
}

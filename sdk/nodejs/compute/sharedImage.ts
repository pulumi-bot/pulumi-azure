// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Shared Image within a Shared Image Gallery.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleSharedImageGallery = new azure.compute.SharedImageGallery("exampleSharedImageGallery", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     description: "Shared images and things.",
 *     tags: {
 *         Hello: "There",
 *         World: "Example",
 *     },
 * });
 * const exampleSharedImage = new azure.compute.SharedImage("exampleSharedImage", {
 *     galleryName: exampleSharedImageGallery.name,
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     osType: "Linux",
 *     identifier: {
 *         publisher: "PublisherName",
 *         offer: "OfferName",
 *         sku: "ExampleSku",
 *     },
 * });
 * ```
 */
export class SharedImage extends pulumi.CustomResource {
    /**
     * Get an existing SharedImage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SharedImageState, opts?: pulumi.CustomResourceOptions): SharedImage {
        return new SharedImage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:compute/sharedImage:SharedImage';

    /**
     * Returns true if the given object is an instance of SharedImage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SharedImage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SharedImage.__pulumiType;
    }

    /**
     * A description of this Shared Image.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The End User Licence Agreement for the Shared Image.
     */
    public readonly eula!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
     */
    public readonly galleryName!: pulumi.Output<string>;
    /**
     * The generation of HyperV that the Virtual Machine used to create the Shared Image is based on. Possible values are `V1` and `V2`. Defaults to `V1`. Changing this forces a new resource to be created.
     */
    public readonly hyperVGeneration!: pulumi.Output<string | undefined>;
    /**
     * An `identifier` block as defined below.
     */
    public readonly identifier!: pulumi.Output<outputs.compute.SharedImageIdentifier>;
    /**
     * Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Shared Image. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`.
     */
    public readonly osType!: pulumi.Output<string>;
    /**
     * The URI containing the Privacy Statement associated with this Shared Image.
     */
    public readonly privacyStatementUri!: pulumi.Output<string | undefined>;
    /**
     * The URI containing the Release Notes associated with this Shared Image.
     */
    public readonly releaseNoteUri!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the Shared Image.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a SharedImage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SharedImageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SharedImageArgs | SharedImageState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SharedImageState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["eula"] = state ? state.eula : undefined;
            inputs["galleryName"] = state ? state.galleryName : undefined;
            inputs["hyperVGeneration"] = state ? state.hyperVGeneration : undefined;
            inputs["identifier"] = state ? state.identifier : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["osType"] = state ? state.osType : undefined;
            inputs["privacyStatementUri"] = state ? state.privacyStatementUri : undefined;
            inputs["releaseNoteUri"] = state ? state.releaseNoteUri : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as SharedImageArgs | undefined;
            if (!args || args.galleryName === undefined) {
                throw new Error("Missing required property 'galleryName'");
            }
            if (!args || args.identifier === undefined) {
                throw new Error("Missing required property 'identifier'");
            }
            if (!args || args.osType === undefined) {
                throw new Error("Missing required property 'osType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["eula"] = args ? args.eula : undefined;
            inputs["galleryName"] = args ? args.galleryName : undefined;
            inputs["hyperVGeneration"] = args ? args.hyperVGeneration : undefined;
            inputs["identifier"] = args ? args.identifier : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["osType"] = args ? args.osType : undefined;
            inputs["privacyStatementUri"] = args ? args.privacyStatementUri : undefined;
            inputs["releaseNoteUri"] = args ? args.releaseNoteUri : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SharedImage.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SharedImage resources.
 */
export interface SharedImageState {
    /**
     * A description of this Shared Image.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The End User Licence Agreement for the Shared Image.
     */
    readonly eula?: pulumi.Input<string>;
    /**
     * Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
     */
    readonly galleryName?: pulumi.Input<string>;
    /**
     * The generation of HyperV that the Virtual Machine used to create the Shared Image is based on. Possible values are `V1` and `V2`. Defaults to `V1`. Changing this forces a new resource to be created.
     */
    readonly hyperVGeneration?: pulumi.Input<string>;
    /**
     * An `identifier` block as defined below.
     */
    readonly identifier?: pulumi.Input<inputs.compute.SharedImageIdentifier>;
    /**
     * Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Shared Image. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`.
     */
    readonly osType?: pulumi.Input<string>;
    /**
     * The URI containing the Privacy Statement associated with this Shared Image.
     */
    readonly privacyStatementUri?: pulumi.Input<string>;
    /**
     * The URI containing the Release Notes associated with this Shared Image.
     */
    readonly releaseNoteUri?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the Shared Image.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a SharedImage resource.
 */
export interface SharedImageArgs {
    /**
     * A description of this Shared Image.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The End User Licence Agreement for the Shared Image.
     */
    readonly eula?: pulumi.Input<string>;
    /**
     * Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
     */
    readonly galleryName: pulumi.Input<string>;
    /**
     * The generation of HyperV that the Virtual Machine used to create the Shared Image is based on. Possible values are `V1` and `V2`. Defaults to `V1`. Changing this forces a new resource to be created.
     */
    readonly hyperVGeneration?: pulumi.Input<string>;
    /**
     * An `identifier` block as defined below.
     */
    readonly identifier: pulumi.Input<inputs.compute.SharedImageIdentifier>;
    /**
     * Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Shared Image. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`.
     */
    readonly osType: pulumi.Input<string>;
    /**
     * The URI containing the Privacy Statement associated with this Shared Image.
     */
    readonly privacyStatementUri?: pulumi.Input<string>;
    /**
     * The URI containing the Release Notes associated with this Shared Image.
     */
    readonly releaseNoteUri?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the Shared Image.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

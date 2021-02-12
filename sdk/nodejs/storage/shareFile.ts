// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a File within an Azure Storage File Share.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const exampleShare = new azure.storage.Share("exampleShare", {
 *     storageAccountName: exampleAccount.name,
 *     quota: 50,
 * });
 * const exampleShareFile = new azure.storage.ShareFile("exampleShareFile", {
 *     storageShareId: exampleShare.id,
 *     source: "some-local-file.zip",
 * });
 * ```
 *
 * ## Import
 *
 * Directories within an Azure Storage File Share can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:storage/shareFile:ShareFile example https://account1.file.core.windows.net/share1/file1
 * ```
 */
export class ShareFile extends pulumi.CustomResource {
    /**
     * Get an existing ShareFile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ShareFileState, opts?: pulumi.CustomResourceOptions): ShareFile {
        return new ShareFile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:storage/shareFile:ShareFile';

    /**
     * Returns true if the given object is an instance of ShareFile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ShareFile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ShareFile.__pulumiType;
    }

    /**
     * Sets the file’s Content-Disposition header.
     */
    public readonly contentDisposition!: pulumi.Output<string | undefined>;
    /**
     * Specifies which content encodings have been applied to the file.
     */
    public readonly contentEncoding!: pulumi.Output<string | undefined>;
    /**
     * The MD5 sum of the file contents. Changing this forces a new resource to be created.
     */
    public readonly contentMd5!: pulumi.Output<string | undefined>;
    /**
     * The content type of the share file. Defaults to `application/octet-stream`.
     */
    public readonly contentType!: pulumi.Output<string | undefined>;
    /**
     * A mapping of metadata to assign to this file.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name (or path) of the File that should be created within this File Share. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The storage share directory that you would like the file placed into. Changing this forces a new resource to be created.
     */
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * An absolute path to a file on the local system.
     */
    public readonly source!: pulumi.Output<string | undefined>;
    /**
     * The Storage Share ID in which this file will be placed into. Changing this forces a new resource to be created.
     */
    public readonly storageShareId!: pulumi.Output<string>;

    /**
     * Create a ShareFile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ShareFileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ShareFileArgs | ShareFileState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ShareFileState | undefined;
            inputs["contentDisposition"] = state ? state.contentDisposition : undefined;
            inputs["contentEncoding"] = state ? state.contentEncoding : undefined;
            inputs["contentMd5"] = state ? state.contentMd5 : undefined;
            inputs["contentType"] = state ? state.contentType : undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["source"] = state ? state.source : undefined;
            inputs["storageShareId"] = state ? state.storageShareId : undefined;
        } else {
            const args = argsOrState as ShareFileArgs | undefined;
            if ((!args || args.storageShareId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageShareId'");
            }
            inputs["contentDisposition"] = args ? args.contentDisposition : undefined;
            inputs["contentEncoding"] = args ? args.contentEncoding : undefined;
            inputs["contentMd5"] = args ? args.contentMd5 : undefined;
            inputs["contentType"] = args ? args.contentType : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["storageShareId"] = args ? args.storageShareId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ShareFile.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ShareFile resources.
 */
export interface ShareFileState {
    /**
     * Sets the file’s Content-Disposition header.
     */
    readonly contentDisposition?: pulumi.Input<string>;
    /**
     * Specifies which content encodings have been applied to the file.
     */
    readonly contentEncoding?: pulumi.Input<string>;
    /**
     * The MD5 sum of the file contents. Changing this forces a new resource to be created.
     */
    readonly contentMd5?: pulumi.Input<string>;
    /**
     * The content type of the share file. Defaults to `application/octet-stream`.
     */
    readonly contentType?: pulumi.Input<string>;
    /**
     * A mapping of metadata to assign to this file.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name (or path) of the File that should be created within this File Share. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The storage share directory that you would like the file placed into. Changing this forces a new resource to be created.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * An absolute path to a file on the local system.
     */
    readonly source?: pulumi.Input<string>;
    /**
     * The Storage Share ID in which this file will be placed into. Changing this forces a new resource to be created.
     */
    readonly storageShareId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ShareFile resource.
 */
export interface ShareFileArgs {
    /**
     * Sets the file’s Content-Disposition header.
     */
    readonly contentDisposition?: pulumi.Input<string>;
    /**
     * Specifies which content encodings have been applied to the file.
     */
    readonly contentEncoding?: pulumi.Input<string>;
    /**
     * The MD5 sum of the file contents. Changing this forces a new resource to be created.
     */
    readonly contentMd5?: pulumi.Input<string>;
    /**
     * The content type of the share file. Defaults to `application/octet-stream`.
     */
    readonly contentType?: pulumi.Input<string>;
    /**
     * A mapping of metadata to assign to this file.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name (or path) of the File that should be created within this File Share. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The storage share directory that you would like the file placed into. Changing this forces a new resource to be created.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * An absolute path to a file on the local system.
     */
    readonly source?: pulumi.Input<string>;
    /**
     * The Storage Share ID in which this file will be placed into. Changing this forces a new resource to be created.
     */
    readonly storageShareId: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows you to set a user or group as the AD administrator for an PostgreSQL server in Azure
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const current = azure.core.getClientConfig({});
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleServer = new azure.postgresql.Server("exampleServer", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     version: "9.6",
 *     administratorLogin: "4dm1n157r470r",
 *     administratorLoginPassword: "4-v3ry-53cr37-p455w0rd",
 * });
 * const exampleActiveDirectoryAdministrator = new azure.postgresql.ActiveDirectoryAdministrator("exampleActiveDirectoryAdministrator", {
 *     serverName: exampleServer.name,
 *     resourceGroupName: exampleResourceGroup.name,
 *     login: "sqladmin",
 *     tenantId: current.then(current => current.tenantId),
 *     objectId: current.then(current => current.objectId),
 * });
 * ```
 *
 * ## Import
 *
 * A PostgreSQL Active Directory Administrator can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:postgresql/activeDirectoryAdministrator:ActiveDirectoryAdministrator administrator /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.DBforPostgreSQL/servers/myserver/administrators/activeDirectory
 * ```
 */
export class ActiveDirectoryAdministrator extends pulumi.CustomResource {
    /**
     * Get an existing ActiveDirectoryAdministrator resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ActiveDirectoryAdministratorState, opts?: pulumi.CustomResourceOptions): ActiveDirectoryAdministrator {
        return new ActiveDirectoryAdministrator(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:postgresql/activeDirectoryAdministrator:ActiveDirectoryAdministrator';

    /**
     * Returns true if the given object is an instance of ActiveDirectoryAdministrator.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ActiveDirectoryAdministrator {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ActiveDirectoryAdministrator.__pulumiType;
    }

    /**
     * The login name of the principal to set as the server administrator
     */
    public readonly login!: pulumi.Output<string>;
    /**
     * The ID of the principal to set as the server administrator
     */
    public readonly objectId!: pulumi.Output<string>;
    /**
     * The name of the resource group for the PostgreSQL server. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The name of the PostgreSQL Server on which to set the administrator. Changing this forces a new resource to be created.
     */
    public readonly serverName!: pulumi.Output<string>;
    /**
     * The Azure Tenant ID
     */
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a ActiveDirectoryAdministrator resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ActiveDirectoryAdministratorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ActiveDirectoryAdministratorArgs | ActiveDirectoryAdministratorState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ActiveDirectoryAdministratorState | undefined;
            inputs["login"] = state ? state.login : undefined;
            inputs["objectId"] = state ? state.objectId : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["serverName"] = state ? state.serverName : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as ActiveDirectoryAdministratorArgs | undefined;
            if ((!args || args.login === undefined) && !opts.urn) {
                throw new Error("Missing required property 'login'");
            }
            if ((!args || args.objectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'objectId'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.serverName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverName'");
            }
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            inputs["login"] = args ? args.login : undefined;
            inputs["objectId"] = args ? args.objectId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ActiveDirectoryAdministrator.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ActiveDirectoryAdministrator resources.
 */
export interface ActiveDirectoryAdministratorState {
    /**
     * The login name of the principal to set as the server administrator
     */
    login?: pulumi.Input<string>;
    /**
     * The ID of the principal to set as the server administrator
     */
    objectId?: pulumi.Input<string>;
    /**
     * The name of the resource group for the PostgreSQL server. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The name of the PostgreSQL Server on which to set the administrator. Changing this forces a new resource to be created.
     */
    serverName?: pulumi.Input<string>;
    /**
     * The Azure Tenant ID
     */
    tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ActiveDirectoryAdministrator resource.
 */
export interface ActiveDirectoryAdministratorArgs {
    /**
     * The login name of the principal to set as the server administrator
     */
    login: pulumi.Input<string>;
    /**
     * The ID of the principal to set as the server administrator
     */
    objectId: pulumi.Input<string>;
    /**
     * The name of the resource group for the PostgreSQL server. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the PostgreSQL Server on which to set the administrator. Changing this forces a new resource to be created.
     */
    serverName: pulumi.Input<string>;
    /**
     * The Azure Tenant ID
     */
    tenantId: pulumi.Input<string>;
}

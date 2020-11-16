// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Automation Runbook.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAccount = new azure.automation.Account("exampleAccount", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     skuName: "Basic",
 * });
 * const exampleRunBook = new azure.automation.RunBook("exampleRunBook", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     automationAccountName: exampleAccount.name,
 *     logVerbose: "true",
 *     logProgress: "true",
 *     description: "This is an example runbook",
 *     runbookType: "PowerShellWorkflow",
 *     publishContentLink: {
 *         uri: "https://raw.githubusercontent.com/Azure/azure-quickstart-templates/c4935ffb69246a6058eb24f54640f53f69d3ac9f/101-automation-runbook-getvms/Runbooks/Get-AzureVMTutorial.ps1",
 *     },
 * });
 * ```
 */
export class RunBook extends pulumi.CustomResource {
    /**
     * Get an existing RunBook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RunBookState, opts?: pulumi.CustomResourceOptions): RunBook {
        return new RunBook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:automation/runBook:RunBook';

    /**
     * Returns true if the given object is an instance of RunBook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RunBook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RunBook.__pulumiType;
    }

    /**
     * The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
     */
    public readonly automationAccountName!: pulumi.Output<string>;
    /**
     * The desired content of the runbook.
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * A description for this credential.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly jobSchedules!: pulumi.Output<outputs.automation.RunBookJobSchedule[]>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Progress log option.
     */
    public readonly logProgress!: pulumi.Output<boolean>;
    /**
     * Verbose log option.
     */
    public readonly logVerbose!: pulumi.Output<boolean>;
    /**
     * Specifies the name of the Runbook. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The published runbook content link.
     */
    public readonly publishContentLink!: pulumi.Output<outputs.automation.RunBookPublishContentLink | undefined>;
    /**
     * The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The type of the runbook - can be either `Graph`, `GraphPowerShell`, `GraphPowerShellWorkflow`, `PowerShellWorkflow`, `PowerShell` or `Script`.
     */
    public readonly runbookType!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a RunBook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RunBookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RunBookArgs | RunBookState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RunBookState | undefined;
            inputs["automationAccountName"] = state ? state.automationAccountName : undefined;
            inputs["content"] = state ? state.content : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["jobSchedules"] = state ? state.jobSchedules : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["logProgress"] = state ? state.logProgress : undefined;
            inputs["logVerbose"] = state ? state.logVerbose : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["publishContentLink"] = state ? state.publishContentLink : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["runbookType"] = state ? state.runbookType : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as RunBookArgs | undefined;
            if (!args || args.automationAccountName === undefined) {
                throw new Error("Missing required property 'automationAccountName'");
            }
            if (!args || args.logProgress === undefined) {
                throw new Error("Missing required property 'logProgress'");
            }
            if (!args || args.logVerbose === undefined) {
                throw new Error("Missing required property 'logVerbose'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.runbookType === undefined) {
                throw new Error("Missing required property 'runbookType'");
            }
            inputs["automationAccountName"] = args ? args.automationAccountName : undefined;
            inputs["content"] = args ? args.content : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["jobSchedules"] = args ? args.jobSchedules : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["logProgress"] = args ? args.logProgress : undefined;
            inputs["logVerbose"] = args ? args.logVerbose : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["publishContentLink"] = args ? args.publishContentLink : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["runbookType"] = args ? args.runbookType : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RunBook.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RunBook resources.
 */
export interface RunBookState {
    /**
     * The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
     */
    readonly automationAccountName?: pulumi.Input<string>;
    /**
     * The desired content of the runbook.
     */
    readonly content?: pulumi.Input<string>;
    /**
     * A description for this credential.
     */
    readonly description?: pulumi.Input<string>;
    readonly jobSchedules?: pulumi.Input<pulumi.Input<inputs.automation.RunBookJobSchedule>[]>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Progress log option.
     */
    readonly logProgress?: pulumi.Input<boolean>;
    /**
     * Verbose log option.
     */
    readonly logVerbose?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the Runbook. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The published runbook content link.
     */
    readonly publishContentLink?: pulumi.Input<inputs.automation.RunBookPublishContentLink>;
    /**
     * The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The type of the runbook - can be either `Graph`, `GraphPowerShell`, `GraphPowerShellWorkflow`, `PowerShellWorkflow`, `PowerShell` or `Script`.
     */
    readonly runbookType?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a RunBook resource.
 */
export interface RunBookArgs {
    /**
     * The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
     */
    readonly automationAccountName: pulumi.Input<string>;
    /**
     * The desired content of the runbook.
     */
    readonly content?: pulumi.Input<string>;
    /**
     * A description for this credential.
     */
    readonly description?: pulumi.Input<string>;
    readonly jobSchedules?: pulumi.Input<pulumi.Input<inputs.automation.RunBookJobSchedule>[]>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Progress log option.
     */
    readonly logProgress: pulumi.Input<boolean>;
    /**
     * Verbose log option.
     */
    readonly logVerbose: pulumi.Input<boolean>;
    /**
     * Specifies the name of the Runbook. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The published runbook content link.
     */
    readonly publishContentLink?: pulumi.Input<inputs.automation.RunBookPublishContentLink>;
    /**
     * The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The type of the runbook - can be either `Graph`, `GraphPowerShell`, `GraphPowerShellWorkflow`, `PowerShellWorkflow`, `PowerShell` or `Script`.
     */
    readonly runbookType: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

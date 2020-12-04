// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./account";
export * from "./boolVariable";
export * from "./certificate";
export * from "./connection";
export * from "./connectionCertificate";
export * from "./connectionClassicCertificate";
export * from "./connectionServicePrincipal";
export * from "./credential";
export * from "./dateTimeVariable";
export * from "./dscConfiguration";
export * from "./dscNodeConfiguration";
export * from "./getAccount";
export * from "./getBoolVariable";
export * from "./getDateTimeVariable";
export * from "./getIntVariable";
export * from "./getStringVariable";
export * from "./intVariable";
export * from "./jobSchedule";
export * from "./module";
export * from "./runBook";
export * from "./schedule";
export * from "./stringVariable";

// Import resources to register:
import { Account } from "./account";
import { BoolVariable } from "./boolVariable";
import { Certificate } from "./certificate";
import { Connection } from "./connection";
import { ConnectionCertificate } from "./connectionCertificate";
import { ConnectionClassicCertificate } from "./connectionClassicCertificate";
import { ConnectionServicePrincipal } from "./connectionServicePrincipal";
import { Credential } from "./credential";
import { DateTimeVariable } from "./dateTimeVariable";
import { DscConfiguration } from "./dscConfiguration";
import { DscNodeConfiguration } from "./dscNodeConfiguration";
import { IntVariable } from "./intVariable";
import { JobSchedule } from "./jobSchedule";
import { Module } from "./module";
import { RunBook } from "./runBook";
import { Schedule } from "./schedule";
import { StringVariable } from "./stringVariable";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure:automation/account:Account":
                return new Account(name, <any>undefined, { urn })
            case "azure:automation/boolVariable:BoolVariable":
                return new BoolVariable(name, <any>undefined, { urn })
            case "azure:automation/certificate:Certificate":
                return new Certificate(name, <any>undefined, { urn })
            case "azure:automation/connection:Connection":
                return new Connection(name, <any>undefined, { urn })
            case "azure:automation/connectionCertificate:ConnectionCertificate":
                return new ConnectionCertificate(name, <any>undefined, { urn })
            case "azure:automation/connectionClassicCertificate:ConnectionClassicCertificate":
                return new ConnectionClassicCertificate(name, <any>undefined, { urn })
            case "azure:automation/connectionServicePrincipal:ConnectionServicePrincipal":
                return new ConnectionServicePrincipal(name, <any>undefined, { urn })
            case "azure:automation/credential:Credential":
                return new Credential(name, <any>undefined, { urn })
            case "azure:automation/dateTimeVariable:DateTimeVariable":
                return new DateTimeVariable(name, <any>undefined, { urn })
            case "azure:automation/dscConfiguration:DscConfiguration":
                return new DscConfiguration(name, <any>undefined, { urn })
            case "azure:automation/dscNodeConfiguration:DscNodeConfiguration":
                return new DscNodeConfiguration(name, <any>undefined, { urn })
            case "azure:automation/intVariable:IntVariable":
                return new IntVariable(name, <any>undefined, { urn })
            case "azure:automation/jobSchedule:JobSchedule":
                return new JobSchedule(name, <any>undefined, { urn })
            case "azure:automation/module:Module":
                return new Module(name, <any>undefined, { urn })
            case "azure:automation/runBook:RunBook":
                return new RunBook(name, <any>undefined, { urn })
            case "azure:automation/schedule:Schedule":
                return new Schedule(name, <any>undefined, { urn })
            case "azure:automation/stringVariable:StringVariable":
                return new StringVariable(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure", "automation/account", _module)
pulumi.runtime.registerResourceModule("azure", "automation/boolVariable", _module)
pulumi.runtime.registerResourceModule("azure", "automation/certificate", _module)
pulumi.runtime.registerResourceModule("azure", "automation/connection", _module)
pulumi.runtime.registerResourceModule("azure", "automation/connectionCertificate", _module)
pulumi.runtime.registerResourceModule("azure", "automation/connectionClassicCertificate", _module)
pulumi.runtime.registerResourceModule("azure", "automation/connectionServicePrincipal", _module)
pulumi.runtime.registerResourceModule("azure", "automation/credential", _module)
pulumi.runtime.registerResourceModule("azure", "automation/dateTimeVariable", _module)
pulumi.runtime.registerResourceModule("azure", "automation/dscConfiguration", _module)
pulumi.runtime.registerResourceModule("azure", "automation/dscNodeConfiguration", _module)
pulumi.runtime.registerResourceModule("azure", "automation/intVariable", _module)
pulumi.runtime.registerResourceModule("azure", "automation/jobSchedule", _module)
pulumi.runtime.registerResourceModule("azure", "automation/module", _module)
pulumi.runtime.registerResourceModule("azure", "automation/runBook", _module)
pulumi.runtime.registerResourceModule("azure", "automation/schedule", _module)
pulumi.runtime.registerResourceModule("azure", "automation/stringVariable", _module)

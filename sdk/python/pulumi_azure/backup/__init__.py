# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .container_storage_account import *
from .get_policy_vm import *
from .policy_file_share import *
from .policy_vm import *
from .protected_file_share import *
from .protected_vm import *
from ._inputs import *
from . import outputs

import pulumi

class Module(pulumi.runtime.ResourceModule):
    def version(self):
        return None

    def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
        if typ == "azure:backup/containerStorageAccount:ContainerStorageAccount":
            return ContainerStorageAccount(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:backup/policyFileShare:PolicyFileShare":
            return PolicyFileShare(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:backup/policyVM:PolicyVM":
            return PolicyVM(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:backup/protectedFileShare:ProtectedFileShare":
            return ProtectedFileShare(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:backup/protectedVM:ProtectedVM":
            return ProtectedVM(name, pulumi.ResourceOptions(urn=urn))
        else:
            raise Exception(f"unknown resource type {typ}")


_module_instance = Module()
pulumi.runtime.register_resource_module("azure", "backup/containerStorageAccount", _module_instance)
pulumi.runtime.register_resource_module("azure", "backup/policyFileShare", _module_instance)
pulumi.runtime.register_resource_module("azure", "backup/policyVM", _module_instance)
pulumi.runtime.register_resource_module("azure", "backup/protectedFileShare", _module_instance)
pulumi.runtime.register_resource_module("azure", "backup/protectedVM", _module_instance)

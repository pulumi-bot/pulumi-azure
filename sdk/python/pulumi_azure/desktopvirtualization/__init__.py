# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .application_group import *
from .host_pool import *
from .workspace import *
from .workspace_application_group_association import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi

    class Module(pulumi.runtime.ResourceModule):
        def version(self):
            return None

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "azure:desktopvirtualization/applicationGroup:ApplicationGroup":
                return ApplicationGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:desktopvirtualization/hostPool:HostPool":
                return HostPool(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:desktopvirtualization/workspace:Workspace":
                return Workspace(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:desktopvirtualization/workspaceApplicationGroupAssociation:WorkspaceApplicationGroupAssociation":
                return WorkspaceApplicationGroupAssociation(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure", "desktopvirtualization/applicationGroup", _module_instance)
    pulumi.runtime.register_resource_module("azure", "desktopvirtualization/hostPool", _module_instance)
    pulumi.runtime.register_resource_module("azure", "desktopvirtualization/workspace", _module_instance)
    pulumi.runtime.register_resource_module("azure", "desktopvirtualization/workspaceApplicationGroupAssociation", _module_instance)

_register_module()

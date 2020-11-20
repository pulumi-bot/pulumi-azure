# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .assignment import *
from .definition import *
from .get_role_definition import *
from ._inputs import *
from . import outputs

import pulumi

class Module(pulumi.runtime.ResourceModule):
    def version(self):
        return None

    def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
        if typ == "azure:role/assignment:Assignment":
            return Assignment(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:role/definition:Definition":
            return Definition(name, pulumi.ResourceOptions(urn=urn))
        else:
            raise Exception(f"unknown resource type {typ}")


_module_instance = Module()
pulumi.runtime.register_resource_module("azure", "role/assignment", _module_instance)
pulumi.runtime.register_resource_module("azure", "role/definition", _module_instance)

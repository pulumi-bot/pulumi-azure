# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .account import *
from .get_account import *
from .get_pool import *
from .get_snapshot import *
from .get_volume import *
from .pool import *
from .snapshot import *
from .volume import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi

    class Module(pulumi.runtime.ResourceModule):
        def version(self):
            return None

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "azure:netapp/account:Account":
                return Account(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:netapp/pool:Pool":
                return Pool(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:netapp/snapshot:Snapshot":
                return Snapshot(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:netapp/volume:Volume":
                return Volume(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure", "netapp/account", _module_instance)
    pulumi.runtime.register_resource_module("azure", "netapp/pool", _module_instance)
    pulumi.runtime.register_resource_module("azure", "netapp/snapshot", _module_instance)
    pulumi.runtime.register_resource_module("azure", "netapp/volume", _module_instance)

_register_module()

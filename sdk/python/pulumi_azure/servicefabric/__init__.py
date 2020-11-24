# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .cluster import *
from .mesh_application import *
from .mesh_local_network import *
from .mesh_secret import *
from .mesh_secret_value import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi

    class Module(pulumi.runtime.ResourceModule):
        def version(self):
            return None

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "azure:servicefabric/cluster:Cluster":
                return Cluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:servicefabric/meshApplication:MeshApplication":
                return MeshApplication(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:servicefabric/meshLocalNetwork:MeshLocalNetwork":
                return MeshLocalNetwork(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:servicefabric/meshSecret:MeshSecret":
                return MeshSecret(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:servicefabric/meshSecretValue:MeshSecretValue":
                return MeshSecretValue(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure", "servicefabric/cluster", _module_instance)
    pulumi.runtime.register_resource_module("azure", "servicefabric/meshApplication", _module_instance)
    pulumi.runtime.register_resource_module("azure", "servicefabric/meshLocalNetwork", _module_instance)
    pulumi.runtime.register_resource_module("azure", "servicefabric/meshSecret", _module_instance)
    pulumi.runtime.register_resource_module("azure", "servicefabric/meshSecretValue", _module_instance)

_register_module()

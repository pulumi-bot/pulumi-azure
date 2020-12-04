# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .backend_address_pool import *
from .get_backend_address_pool import *
from .get_lb import *
from .get_lb_rule import *
from .load_balancer import *
from .nat_pool import *
from .nat_rule import *
from .outbound_rule import *
from .probe import *
from .rule import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "azure:lb/backendAddressPool:BackendAddressPool":
                return BackendAddressPool(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:lb/loadBalancer:LoadBalancer":
                return LoadBalancer(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:lb/natPool:NatPool":
                return NatPool(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:lb/natRule:NatRule":
                return NatRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:lb/outboundRule:OutboundRule":
                return OutboundRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:lb/probe:Probe":
                return Probe(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:lb/rule:Rule":
                return Rule(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure", "lb/backendAddressPool", _module_instance)
    pulumi.runtime.register_resource_module("azure", "lb/loadBalancer", _module_instance)
    pulumi.runtime.register_resource_module("azure", "lb/natPool", _module_instance)
    pulumi.runtime.register_resource_module("azure", "lb/natRule", _module_instance)
    pulumi.runtime.register_resource_module("azure", "lb/outboundRule", _module_instance)
    pulumi.runtime.register_resource_module("azure", "lb/probe", _module_instance)
    pulumi.runtime.register_resource_module("azure", "lb/rule", _module_instance)

_register_module()

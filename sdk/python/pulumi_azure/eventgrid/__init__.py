# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .domain import *
from .domain_topic import *
from .event_subscription import *
from .get_system_topic import *
from .get_topic import *
from .topic import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi

    class Module(pulumi.runtime.ResourceModule):
        def version(self):
            return None

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "azure:eventgrid/domain:Domain":
                return Domain(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventgrid/domainTopic:DomainTopic":
                return DomainTopic(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventgrid/eventSubscription:EventSubscription":
                return EventSubscription(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventgrid/getSystemTopic:getSystemTopic":
                return GetSystemTopic(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:eventgrid/topic:Topic":
                return Topic(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure", "eventgrid/domain", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventgrid/domainTopic", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventgrid/eventSubscription", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventgrid/getSystemTopic", _module_instance)
    pulumi.runtime.register_resource_module("azure", "eventgrid/topic", _module_instance)

_register_module()

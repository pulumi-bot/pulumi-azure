# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .analytics_account import *
from .analytics_firewall_rule import *
from .get_store import *
from .store import *
from .store_file import *
from .store_firewall_rule import *

import pulumi

class Module(pulumi.runtime.ResourceModule):
    def version(self):
        return None

    def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
        if typ == "azure:datalake/analyticsAccount:AnalyticsAccount":
            return AnalyticsAccount(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:datalake/analyticsFirewallRule:AnalyticsFirewallRule":
            return AnalyticsFirewallRule(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:datalake/store:Store":
            return Store(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:datalake/storeFile:StoreFile":
            return StoreFile(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:datalake/storeFirewallRule:StoreFirewallRule":
            return StoreFirewallRule(name, pulumi.ResourceOptions(urn=urn))
        else:
            raise Exception(f"unknown resource type {typ}")


_module_instance = Module()
pulumi.runtime.register_resource_module("azure", "datalake/analyticsAccount", _module_instance)
pulumi.runtime.register_resource_module("azure", "datalake/analyticsFirewallRule", _module_instance)
pulumi.runtime.register_resource_module("azure", "datalake/store", _module_instance)
pulumi.runtime.register_resource_module("azure", "datalake/storeFile", _module_instance)
pulumi.runtime.register_resource_module("azure", "datalake/storeFirewallRule", _module_instance)

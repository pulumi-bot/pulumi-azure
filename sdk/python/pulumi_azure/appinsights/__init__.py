# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .analytics_item import *
from .api_key import *
from .get_insights import *
from .insights import *
from .web_test import *

import pulumi

class Module(pulumi.runtime.ResourceModule):
    def version(self):
        return None

    def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
        if typ == "azure:appinsights/analyticsItem:AnalyticsItem":
            return AnalyticsItem(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:appinsights/apiKey:ApiKey":
            return ApiKey(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:appinsights/insights:Insights":
            return Insights(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:appinsights/webTest:WebTest":
            return WebTest(name, pulumi.ResourceOptions(urn=urn))
        else:
            raise Exception(f"unknown resource type {typ}")


_module_instance = Module()
pulumi.runtime.register_resource_module("azure", "appinsights/analyticsItem", _module_instance)
pulumi.runtime.register_resource_module("azure", "appinsights/apiKey", _module_instance)
pulumi.runtime.register_resource_module("azure", "appinsights/insights", _module_instance)
pulumi.runtime.register_resource_module("azure", "appinsights/webTest", _module_instance)

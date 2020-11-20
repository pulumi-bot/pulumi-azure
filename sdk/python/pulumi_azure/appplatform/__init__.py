# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .get_spring_cloud_service import *
from .spring_cloud_app import *
from .spring_cloud_certificate import *
from .spring_cloud_service import *
from ._inputs import *
from . import outputs

import pulumi

class Module(pulumi.runtime.ResourceModule):
    def version(self):
        return None

    def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
        if typ == "azure:appplatform/springCloudApp:SpringCloudApp":
            return SpringCloudApp(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:appplatform/springCloudCertificate:SpringCloudCertificate":
            return SpringCloudCertificate(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "azure:appplatform/springCloudService:SpringCloudService":
            return SpringCloudService(name, pulumi.ResourceOptions(urn=urn))
        else:
            raise Exception(f"unknown resource type {typ}")


_module_instance = Module()
pulumi.runtime.register_resource_module("azure", "appplatform/springCloudApp", _module_instance)
pulumi.runtime.register_resource_module("azure", "appplatform/springCloudCertificate", _module_instance)
pulumi.runtime.register_resource_module("azure", "appplatform/springCloudService", _module_instance)

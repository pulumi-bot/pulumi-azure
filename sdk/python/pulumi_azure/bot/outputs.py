# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

@pulumi.output_type
class ChannelDirectLineSite(dict):
    enabled: Optional[bool] = pulumi.output_property("enabled")
    enhanced_authentication_enabled: Optional[bool] = pulumi.output_property("enhancedAuthenticationEnabled")
    id: Optional[str] = pulumi.output_property("id")
    key: Optional[str] = pulumi.output_property("key")
    key2: Optional[str] = pulumi.output_property("key2")
    name: str = pulumi.output_property("name")
    trusted_origins: Optional[List[str]] = pulumi.output_property("trustedOrigins")
    v1_allowed: Optional[bool] = pulumi.output_property("v1Allowed")
    v3_allowed: Optional[bool] = pulumi.output_property("v3Allowed")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


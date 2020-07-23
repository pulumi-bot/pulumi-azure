# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

@pulumi.output_type
class ApplicationPlan(dict):
    name: str = pulumi.output_property("name")
    """
    Specifies the name of the plan from the marketplace.
    """
    product: str = pulumi.output_property("product")
    """
    Specifies the product of the plan from the marketplace.
    """
    promotion_code: Optional[str] = pulumi.output_property("promotionCode")
    """
    Specifies the promotion code to use with the plan.
    """
    publisher: str = pulumi.output_property("publisher")
    """
    Specifies the publisher of the plan.
    """
    version: str = pulumi.output_property("version")
    """
    Specifies the version of the plan from the marketplace.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DefinitionAuthorization(dict):
    role_definition_id: str = pulumi.output_property("roleDefinitionId")
    """
    Specifies a role definition identifier for the provider. This role will define all the permissions that the provider must have on the managed application's container resource group. This role definition cannot have permission to delete the resource group.
    """
    service_principal_id: str = pulumi.output_property("servicePrincipalId")
    """
    Specifies a service principal identifier for the provider. This is the identity that the provider will use to call ARM to manage the managed application resources.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


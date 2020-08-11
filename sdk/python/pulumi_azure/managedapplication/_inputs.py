# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'ApplicationPlanArgs',
    'DefinitionAuthorizationArgs',
]

@pulumi.input_type
class ApplicationPlanArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 product: pulumi.Input[str],
                 publisher: pulumi.Input[str],
                 version: pulumi.Input[str],
                 promotion_code: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Specifies the name of the plan from the marketplace.
        :param pulumi.Input[str] product: Specifies the product of the plan from the marketplace.
        :param pulumi.Input[str] publisher: Specifies the publisher of the plan.
        :param pulumi.Input[str] version: Specifies the version of the plan from the marketplace.
        :param pulumi.Input[str] promotion_code: Specifies the promotion code to use with the plan.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "product", product)
        pulumi.set(__self__, "publisher", publisher)
        pulumi.set(__self__, "version", version)
        pulumi.set(__self__, "promotionCode", promotion_code)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the plan from the marketplace.
        """
        ...

    @name.setter
    def name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def product(self) -> pulumi.Input[str]:
        """
        Specifies the product of the plan from the marketplace.
        """
        ...

    @product.setter
    def product(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def publisher(self) -> pulumi.Input[str]:
        """
        Specifies the publisher of the plan.
        """
        ...

    @publisher.setter
    def publisher(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[str]:
        """
        Specifies the version of the plan from the marketplace.
        """
        ...

    @version.setter
    def version(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="promotionCode")
    def promotion_code(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the promotion code to use with the plan.
        """
        ...

    @promotion_code.setter
    def promotion_code(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class DefinitionAuthorizationArgs:
    def __init__(__self__, *,
                 role_definition_id: pulumi.Input[str],
                 service_principal_id: pulumi.Input[str]):
        """
        :param pulumi.Input[str] role_definition_id: Specifies a role definition identifier for the provider. This role will define all the permissions that the provider must have on the managed application's container resource group. This role definition cannot have permission to delete the resource group.
        :param pulumi.Input[str] service_principal_id: Specifies a service principal identifier for the provider. This is the identity that the provider will use to call ARM to manage the managed application resources.
        """
        pulumi.set(__self__, "roleDefinitionId", role_definition_id)
        pulumi.set(__self__, "servicePrincipalId", service_principal_id)

    @property
    @pulumi.getter(name="roleDefinitionId")
    def role_definition_id(self) -> pulumi.Input[str]:
        """
        Specifies a role definition identifier for the provider. This role will define all the permissions that the provider must have on the managed application's container resource group. This role definition cannot have permission to delete the resource group.
        """
        ...

    @role_definition_id.setter
    def role_definition_id(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> pulumi.Input[str]:
        """
        Specifies a service principal identifier for the provider. This is the identity that the provider will use to call ARM to manage the managed application resources.
        """
        ...

    @service_principal_id.setter
    def service_principal_id(self, value: pulumi.Input[str]):
        ...



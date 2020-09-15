# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'AnalyticsSolutionPlanArgs',
]

@pulumi.input_type
class AnalyticsSolutionPlanArgs:
    def __init__(__self__, *,
                 product: pulumi.Input[str],
                 publisher: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 promotion_code: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] product: The product name of the solution. For example `OMSGallery/Containers`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] publisher: The publisher of the solution. For example `Microsoft`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] promotion_code: A promotion code to be used with the solution.
        """
        pulumi.set(__self__, "product", product)
        pulumi.set(__self__, "publisher", publisher)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if promotion_code is not None:
            pulumi.set(__self__, "promotion_code", promotion_code)

    @property
    @pulumi.getter
    def product(self) -> pulumi.Input[str]:
        """
        The product name of the solution. For example `OMSGallery/Containers`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "product")

    @product.setter
    def product(self, value: pulumi.Input[str]):
        pulumi.set(self, "product", value)

    @property
    @pulumi.getter
    def publisher(self) -> pulumi.Input[str]:
        """
        The publisher of the solution. For example `Microsoft`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "publisher")

    @publisher.setter
    def publisher(self, value: pulumi.Input[str]):
        pulumi.set(self, "publisher", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="promotionCode")
    def promotion_code(self) -> Optional[pulumi.Input[str]]:
        """
        A promotion code to be used with the solution.
        """
        return pulumi.get(self, "promotion_code")

    @promotion_code.setter
    def promotion_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "promotion_code", value)



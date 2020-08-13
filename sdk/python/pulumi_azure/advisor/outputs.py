# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetRecommendationsRecommendationResult',
]

@pulumi.output_type
class GetRecommendationsRecommendationResult(dict):
    @property
    @pulumi.getter
    def category(self) -> str:
        """
        The category of the recommendation.
        """
        ...

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the issue or the opportunity identified by the recommendation.
        """
        ...

    @property
    @pulumi.getter
    def impact(self) -> str:
        """
        The business impact of the recommendation.
        """
        ...

    @property
    @pulumi.getter(name="recommendationName")
    def recommendation_name(self) -> str:
        """
        The name of the Advisor Recommendation.
        """
        ...

    @property
    @pulumi.getter(name="recommendationTypeId")
    def recommendation_type_id(self) -> str:
        """
        The recommendation type id of the Advisor Recommendation.
        """
        ...

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> str:
        """
        The name of the identified resource of the Advisor Recommendation.
        """
        ...

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> str:
        """
        The type of the identified resource of the Advisor Recommendation.
        """
        ...

    @property
    @pulumi.getter(name="suppressionNames")
    def suppression_names(self) -> List[str]:
        """
        A list of Advisor Suppression names of the Advisor Recommendation.
        """
        ...

    @property
    @pulumi.getter(name="updatedTime")
    def updated_time(self) -> str:
        """
        The most recent time that Advisor checked the validity of the recommendation..
        """
        ...



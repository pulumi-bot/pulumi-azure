# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetRecommendationsResult:
    """
    A collection of values returned by getRecommendations.
    """
    def __init__(__self__, filter_by_categories=None, filter_by_resource_groups=None, id=None, recommendations=None):
        if filter_by_categories and not isinstance(filter_by_categories, list):
            raise TypeError("Expected argument 'filter_by_categories' to be a list")
        __self__.filter_by_categories = filter_by_categories
        if filter_by_resource_groups and not isinstance(filter_by_resource_groups, list):
            raise TypeError("Expected argument 'filter_by_resource_groups' to be a list")
        __self__.filter_by_resource_groups = filter_by_resource_groups
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if recommendations and not isinstance(recommendations, list):
            raise TypeError("Expected argument 'recommendations' to be a list")
        __self__.recommendations = recommendations
        """
        One or more `recommendations` blocks as defined below.
        """
class AwaitableGetRecommendationsResult(GetRecommendationsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRecommendationsResult(
            filter_by_categories=self.filter_by_categories,
            filter_by_resource_groups=self.filter_by_resource_groups,
            id=self.id,
            recommendations=self.recommendations)

def get_recommendations(filter_by_categories=None,filter_by_resource_groups=None,opts=None):
    """
    Use this data source to access information about an existing Advisor Recommendations.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.advisor.get_recommendations(filter_by_categories=[
            "security",
            "cost",
        ],
        filter_by_resource_groups=["example-resgroups"])
    pulumi.export("recommendations", example.recommendations)
    ```


    :param list filter_by_categories: Specifies a list of categories in which the Advisor Recommendations will be listed. Possible values are `HighAvailability`, `Security`, `Performance`, `Cost` and `OperationalExcellence`.
    :param list filter_by_resource_groups: Specifies a list of resource groups about which the Advisor Recommendations will be listed.
    """
    __args__ = dict()


    __args__['filterByCategories'] = filter_by_categories
    __args__['filterByResourceGroups'] = filter_by_resource_groups
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:advisor/getRecommendations:getRecommendations', __args__, opts=opts).value

    return AwaitableGetRecommendationsResult(
        filter_by_categories=__ret__.get('filterByCategories'),
        filter_by_resource_groups=__ret__.get('filterByResourceGroups'),
        id=__ret__.get('id'),
        recommendations=__ret__.get('recommendations'))

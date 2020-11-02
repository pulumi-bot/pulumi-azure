# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['TimeSeriesInsightsAccessPolicy']


class TimeSeriesInsightsAccessPolicy(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 principal_object_id: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 time_series_insights_environment_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Azure IoT Time Series Insights Access Policy.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="northeurope")
        example_time_series_insights_standard_environment = azure.iot.TimeSeriesInsightsStandardEnvironment("exampleTimeSeriesInsightsStandardEnvironment",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku_name="S1_1",
            data_retention_time="P30D")
        example_time_series_insights_access_policy = azure.iot.TimeSeriesInsightsAccessPolicy("exampleTimeSeriesInsightsAccessPolicy",
            time_series_insights_environment_id=example_time_series_insights_standard_environment.name,
            principal_object_id="aGUID",
            roles=["Reader"])
        ```

        ## Import

        Azure IoT Time Series Insights Access Policy can be imported using the `resource id`, e.g.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the Azure IoT Time Series Insights Access Policy.
        :param pulumi.Input[str] name: Specifies the name of the Azure IoT Time Series Insights Access Policy. Changing this forces a new resource to be created. Must be globally unique.
        :param pulumi.Input[str] principal_object_id: The id of the principal in Azure Active Directory.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: A list of roles to apply to the Access Policy. Valid values include `Contributor` and `Reader`.
        :param pulumi.Input[str] time_series_insights_environment_id: The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            __props__['name'] = name
            if principal_object_id is None:
                raise TypeError("Missing required property 'principal_object_id'")
            __props__['principal_object_id'] = principal_object_id
            if roles is None:
                raise TypeError("Missing required property 'roles'")
            __props__['roles'] = roles
            if time_series_insights_environment_id is None:
                raise TypeError("Missing required property 'time_series_insights_environment_id'")
            __props__['time_series_insights_environment_id'] = time_series_insights_environment_id
        super(TimeSeriesInsightsAccessPolicy, __self__).__init__(
            'azure:iot/timeSeriesInsightsAccessPolicy:TimeSeriesInsightsAccessPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            principal_object_id: Optional[pulumi.Input[str]] = None,
            roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            time_series_insights_environment_id: Optional[pulumi.Input[str]] = None) -> 'TimeSeriesInsightsAccessPolicy':
        """
        Get an existing TimeSeriesInsightsAccessPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the Azure IoT Time Series Insights Access Policy.
        :param pulumi.Input[str] name: Specifies the name of the Azure IoT Time Series Insights Access Policy. Changing this forces a new resource to be created. Must be globally unique.
        :param pulumi.Input[str] principal_object_id: The id of the principal in Azure Active Directory.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: A list of roles to apply to the Access Policy. Valid values include `Contributor` and `Reader`.
        :param pulumi.Input[str] time_series_insights_environment_id: The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["name"] = name
        __props__["principal_object_id"] = principal_object_id
        __props__["roles"] = roles
        __props__["time_series_insights_environment_id"] = time_series_insights_environment_id
        return TimeSeriesInsightsAccessPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the Azure IoT Time Series Insights Access Policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Azure IoT Time Series Insights Access Policy. Changing this forces a new resource to be created. Must be globally unique.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="principalObjectId")
    def principal_object_id(self) -> pulumi.Output[str]:
        """
        The id of the principal in Azure Active Directory.
        """
        return pulumi.get(self, "principal_object_id")

    @property
    @pulumi.getter
    def roles(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of roles to apply to the Access Policy. Valid values include `Contributor` and `Reader`.
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter(name="timeSeriesInsightsEnvironmentId")
    def time_series_insights_environment_id(self) -> pulumi.Output[str]:
        """
        The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "time_series_insights_environment_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


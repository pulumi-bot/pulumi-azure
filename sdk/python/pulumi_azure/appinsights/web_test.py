# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class WebTest(pulumi.CustomResource):
    application_insights_id: pulumi.Output[str]
    """
    The ID of the Application Insights component on which the WebTest operates. Changing this forces a new resource to be created.
    """
    configuration: pulumi.Output[str]
    """
    An XML configuration specification for a WebTest.
    """
    description: pulumi.Output[str]
    """
    Purpose/user defined descriptive test for this WebTest.
    """
    enabled: pulumi.Output[bool]
    """
    Is the test actively being monitored.
    """
    frequency: pulumi.Output[float]
    """
    Interval in seconds between test runs for this WebTest. Default is `300`.
    """
    geo_locations: pulumi.Output[list]
    """
    A list of where to physically run the tests from to give global coverage for accessibility of your application.
    """
    kind: pulumi.Output[str]
    """
    = (Required) The kind of web test that this web test watches. Choices are `ping` and `multistep`.
    """
    location: pulumi.Output[str]
    """
    The location of the resource group.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Application Insights WebTest. Changing this forces a
    new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    retry_enabled: pulumi.Output[bool]
    """
    Allow for retries should this WebTest fail.
    """
    synthetic_monitor_id: pulumi.Output[str]
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    timeout: pulumi.Output[float]
    """
    Seconds until this WebTest will timeout and fail. Default is `30`.
    """
    def __init__(__self__, resource_name, opts=None, application_insights_id=None, configuration=None, description=None, enabled=None, frequency=None, geo_locations=None, kind=None, location=None, name=None, resource_group_name=None, retry_enabled=None, tags=None, timeout=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Application Insights WebTest.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_insights_id: The ID of the Application Insights component on which the WebTest operates. Changing this forces a new resource to be created.
        :param pulumi.Input[str] configuration: An XML configuration specification for a WebTest.
        :param pulumi.Input[str] description: Purpose/user defined descriptive test for this WebTest.
        :param pulumi.Input[bool] enabled: Is the test actively being monitored.
        :param pulumi.Input[float] frequency: Interval in seconds between test runs for this WebTest. Default is `300`.
        :param pulumi.Input[list] geo_locations: A list of where to physically run the tests from to give global coverage for accessibility of your application.
        :param pulumi.Input[str] kind: = (Required) The kind of web test that this web test watches. Choices are `ping` and `multistep`.
        :param pulumi.Input[str] location: The location of the resource group.
        :param pulumi.Input[str] name: Specifies the name of the Application Insights WebTest. Changing this forces a
               new resource to be created.
        :param pulumi.Input[bool] retry_enabled: Allow for retries should this WebTest fail.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[float] timeout: Seconds until this WebTest will timeout and fail. Default is `30`.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if application_insights_id is None:
                raise TypeError("Missing required property 'application_insights_id'")
            __props__['application_insights_id'] = application_insights_id
            if configuration is None:
                raise TypeError("Missing required property 'configuration'")
            __props__['configuration'] = configuration
            __props__['description'] = description
            __props__['enabled'] = enabled
            __props__['frequency'] = frequency
            if geo_locations is None:
                raise TypeError("Missing required property 'geo_locations'")
            __props__['geo_locations'] = geo_locations
            if kind is None:
                raise TypeError("Missing required property 'kind'")
            __props__['kind'] = kind
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['retry_enabled'] = retry_enabled
            __props__['tags'] = tags
            __props__['timeout'] = timeout
            __props__['synthetic_monitor_id'] = None
        super(WebTest, __self__).__init__(
            'azure:appinsights/webTest:WebTest',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, application_insights_id=None, configuration=None, description=None, enabled=None, frequency=None, geo_locations=None, kind=None, location=None, name=None, resource_group_name=None, retry_enabled=None, synthetic_monitor_id=None, tags=None, timeout=None):
        """
        Get an existing WebTest resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_insights_id: The ID of the Application Insights component on which the WebTest operates. Changing this forces a new resource to be created.
        :param pulumi.Input[str] configuration: An XML configuration specification for a WebTest.
        :param pulumi.Input[str] description: Purpose/user defined descriptive test for this WebTest.
        :param pulumi.Input[bool] enabled: Is the test actively being monitored.
        :param pulumi.Input[float] frequency: Interval in seconds between test runs for this WebTest. Default is `300`.
        :param pulumi.Input[list] geo_locations: A list of where to physically run the tests from to give global coverage for accessibility of your application.
        :param pulumi.Input[str] kind: = (Required) The kind of web test that this web test watches. Choices are `ping` and `multistep`.
        :param pulumi.Input[str] location: The location of the resource group.
        :param pulumi.Input[str] name: Specifies the name of the Application Insights WebTest. Changing this forces a
               new resource to be created.
        :param pulumi.Input[bool] retry_enabled: Allow for retries should this WebTest fail.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[float] timeout: Seconds until this WebTest will timeout and fail. Default is `30`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["application_insights_id"] = application_insights_id
        __props__["configuration"] = configuration
        __props__["description"] = description
        __props__["enabled"] = enabled
        __props__["frequency"] = frequency
        __props__["geo_locations"] = geo_locations
        __props__["kind"] = kind
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["retry_enabled"] = retry_enabled
        __props__["synthetic_monitor_id"] = synthetic_monitor_id
        __props__["tags"] = tags
        __props__["timeout"] = timeout
        return WebTest(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

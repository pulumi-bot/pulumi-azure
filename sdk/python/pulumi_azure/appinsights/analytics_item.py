# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['AnalyticsItem']


class AnalyticsItem(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_insights_id: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 function_alias: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Application Insights Analytics Item component.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_insights = azure.appinsights.Insights("exampleInsights",
            location="West Europe",
            resource_group_name=example_resource_group.name,
            application_type="web")
        example_analytics_item = azure.appinsights.AnalyticsItem("exampleAnalyticsItem",
            application_insights_id=example_insights.id,
            content="requests //simple example query",
            scope="shared",
            type="query")
        ```

        ## Import

        Application Insights Analytics Items can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:appinsights/analyticsItem:AnalyticsItem example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.insights/components/analyticsItems/11111111-1111-1111-1111-111111111111
        ```

         To find the Analytics Item ID you can query the REST API using the [`az rest` CLI command](https://docs.microsoft.com/en-us/cli/azure/reference-index?view=azure-cli-latest#az-rest), e.g. az rest --method GET --uri "https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.insights/components/appinsightstest/analyticsItems?api-version=2015-05-01"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_insights_id: The ID of the Application Insights component on which the Analytics Item exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] content: The content for the Analytics Item, for example the query text if `type` is `query`.
        :param pulumi.Input[str] function_alias: The alias to use for the function. Required when `type` is `function`.
        :param pulumi.Input[str] name: Specifies the name of the Application Insights Analytics Item. Changing this forces a new resource to be created.
        :param pulumi.Input[str] scope: The scope for the Analytics Item. Can be `shared` or `user`. Changing this forces a new resource to be created. Must be `shared` for functions.
        :param pulumi.Input[str] type: The type of Analytics Item to create. Can be one of `query`, `function`, `folder`, `recent`. Changing this forces a new resource to be created.
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

            if application_insights_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_insights_id'")
            __props__['application_insights_id'] = application_insights_id
            if content is None and not opts.urn:
                raise TypeError("Missing required property 'content'")
            __props__['content'] = content
            __props__['function_alias'] = function_alias
            __props__['name'] = name
            if scope is None and not opts.urn:
                raise TypeError("Missing required property 'scope'")
            __props__['scope'] = scope
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['time_created'] = None
            __props__['time_modified'] = None
            __props__['version'] = None
        super(AnalyticsItem, __self__).__init__(
            'azure:appinsights/analyticsItem:AnalyticsItem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_insights_id: Optional[pulumi.Input[str]] = None,
            content: Optional[pulumi.Input[str]] = None,
            function_alias: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            scope: Optional[pulumi.Input[str]] = None,
            time_created: Optional[pulumi.Input[str]] = None,
            time_modified: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'AnalyticsItem':
        """
        Get an existing AnalyticsItem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_insights_id: The ID of the Application Insights component on which the Analytics Item exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] content: The content for the Analytics Item, for example the query text if `type` is `query`.
        :param pulumi.Input[str] function_alias: The alias to use for the function. Required when `type` is `function`.
        :param pulumi.Input[str] name: Specifies the name of the Application Insights Analytics Item. Changing this forces a new resource to be created.
        :param pulumi.Input[str] scope: The scope for the Analytics Item. Can be `shared` or `user`. Changing this forces a new resource to be created. Must be `shared` for functions.
        :param pulumi.Input[str] time_created: A string containing the time the Analytics Item was created.
        :param pulumi.Input[str] time_modified: A string containing the time the Analytics Item was last modified.
        :param pulumi.Input[str] type: The type of Analytics Item to create. Can be one of `query`, `function`, `folder`, `recent`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] version: A string indicating the version of the query format
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["application_insights_id"] = application_insights_id
        __props__["content"] = content
        __props__["function_alias"] = function_alias
        __props__["name"] = name
        __props__["scope"] = scope
        __props__["time_created"] = time_created
        __props__["time_modified"] = time_modified
        __props__["type"] = type
        __props__["version"] = version
        return AnalyticsItem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationInsightsId")
    def application_insights_id(self) -> pulumi.Output[str]:
        """
        The ID of the Application Insights component on which the Analytics Item exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "application_insights_id")

    @property
    @pulumi.getter
    def content(self) -> pulumi.Output[str]:
        """
        The content for the Analytics Item, for example the query text if `type` is `query`.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="functionAlias")
    def function_alias(self) -> pulumi.Output[Optional[str]]:
        """
        The alias to use for the function. Required when `type` is `function`.
        """
        return pulumi.get(self, "function_alias")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Application Insights Analytics Item. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[str]:
        """
        The scope for the Analytics Item. Can be `shared` or `user`. Changing this forces a new resource to be created. Must be `shared` for functions.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> pulumi.Output[str]:
        """
        A string containing the time the Analytics Item was created.
        """
        return pulumi.get(self, "time_created")

    @property
    @pulumi.getter(name="timeModified")
    def time_modified(self) -> pulumi.Output[str]:
        """
        A string containing the time the Analytics Item was last modified.
        """
        return pulumi.get(self, "time_modified")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of Analytics Item to create. Can be one of `query`, `function`, `folder`, `recent`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        A string indicating the version of the query format
        """
        return pulumi.get(self, "version")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


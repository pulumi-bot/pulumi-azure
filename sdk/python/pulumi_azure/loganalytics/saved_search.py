# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['SavedSearch']


class SavedSearch(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 function_alias: Optional[pulumi.Input[str]] = None,
                 function_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 log_analytics_workspace_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Log Analytics (formally Operational Insights) Saved Search.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="East US")
        example_analytics_workspace = azure.operationalinsights.AnalyticsWorkspace("exampleAnalyticsWorkspace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="PerGB2018",
            retention_in_days=30)
        example_saved_search = azure.loganalytics.SavedSearch("exampleSavedSearch",
            log_analytics_workspace_id=azurerm_log_analytics_workspace["test"]["id"],
            category="exampleCategory",
            display_name="exampleDisplayName",
            query="exampleQuery")
        ```

        ## Import

        Log Analytics Saved Searches can be imported using the `resource id`, e.g.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] category: The category that the Saved Search will be listed under. Changing this forces a new resource to be created.
        :param pulumi.Input[str] display_name: The name that Saved Search will be displayed as. Changing this forces a new resource to be created.
        :param pulumi.Input[str] function_alias: The function alias if the query serves as a function. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] function_parameters: The function parameters if the query serves as a function. Changing this forces a new resource to be created.
        :param pulumi.Input[str] log_analytics_workspace_id: Specifies the ID of the Log Analytics Workspace that the Saved Search will be associated with. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Log Analytics Saved Search. Changing this forces a new resource to be created.
        :param pulumi.Input[str] query: The query expression for the saved search. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Logs Analytics Saved Search.
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

            if category is None:
                raise TypeError("Missing required property 'category'")
            __props__['category'] = category
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            __props__['function_alias'] = function_alias
            __props__['function_parameters'] = function_parameters
            if log_analytics_workspace_id is None:
                raise TypeError("Missing required property 'log_analytics_workspace_id'")
            __props__['log_analytics_workspace_id'] = log_analytics_workspace_id
            __props__['name'] = name
            if query is None:
                raise TypeError("Missing required property 'query'")
            __props__['query'] = query
            __props__['tags'] = tags
        super(SavedSearch, __self__).__init__(
            'azure:loganalytics/savedSearch:SavedSearch',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            category: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            function_alias: Optional[pulumi.Input[str]] = None,
            function_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            log_analytics_workspace_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            query: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'SavedSearch':
        """
        Get an existing SavedSearch resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] category: The category that the Saved Search will be listed under. Changing this forces a new resource to be created.
        :param pulumi.Input[str] display_name: The name that Saved Search will be displayed as. Changing this forces a new resource to be created.
        :param pulumi.Input[str] function_alias: The function alias if the query serves as a function. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] function_parameters: The function parameters if the query serves as a function. Changing this forces a new resource to be created.
        :param pulumi.Input[str] log_analytics_workspace_id: Specifies the ID of the Log Analytics Workspace that the Saved Search will be associated with. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Log Analytics Saved Search. Changing this forces a new resource to be created.
        :param pulumi.Input[str] query: The query expression for the saved search. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Logs Analytics Saved Search.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["category"] = category
        __props__["display_name"] = display_name
        __props__["function_alias"] = function_alias
        __props__["function_parameters"] = function_parameters
        __props__["log_analytics_workspace_id"] = log_analytics_workspace_id
        __props__["name"] = name
        __props__["query"] = query
        __props__["tags"] = tags
        return SavedSearch(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def category(self) -> pulumi.Output[str]:
        """
        The category that the Saved Search will be listed under. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The name that Saved Search will be displayed as. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="functionAlias")
    def function_alias(self) -> pulumi.Output[Optional[str]]:
        """
        The function alias if the query serves as a function. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "function_alias")

    @property
    @pulumi.getter(name="functionParameters")
    def function_parameters(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The function parameters if the query serves as a function. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "function_parameters")

    @property
    @pulumi.getter(name="logAnalyticsWorkspaceId")
    def log_analytics_workspace_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of the Log Analytics Workspace that the Saved Search will be associated with. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "log_analytics_workspace_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Log Analytics Saved Search. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def query(self) -> pulumi.Output[str]:
        """
        The query expression for the saved search. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "query")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags which should be assigned to the Logs Analytics Saved Search.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class RouteFilter(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The Azure Region where the Route Filter should exist. Changing this forces a new Route Filter to be created.
    """
    name: pulumi.Output[str]
    """
    The Name which should be used for this Route Filter.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group where the Route Filter should exist. Changing this forces a new Route Filter to be created.
    """
    rule: pulumi.Output[dict]
    """
    A `rules` block as defined below.

      * `access` (`str`) - The access type of the rule. The only possible value is `Allow`.
      * `communities` (`list`) - The collection for bgp community values to filter on. e.g. ['12076:5010','12076:5020'].
      * `name` (`str`) - The name of the route filter rule.
      * `ruleType` (`str`) - The rule type of the rule. The only possible value is `Community`.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags which should be assigned to the Route Filter.
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, resource_group_name=None, rule=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Route Filter.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example = azure.network.RouteFilter("example",
            location="East US",
            resource_group_name="example",
            rule={
                "access": "Allow",
                "communities": ["12076:52004"],
                "name": "rule",
                "ruleType": "Community",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The Azure Region where the Route Filter should exist. Changing this forces a new Route Filter to be created.
        :param pulumi.Input[str] name: The Name which should be used for this Route Filter.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Route Filter should exist. Changing this forces a new Route Filter to be created.
        :param pulumi.Input[dict] rule: A `rules` block as defined below.
        :param pulumi.Input[dict] tags: A mapping of tags which should be assigned to the Route Filter.

        The **rule** object supports the following:

          * `access` (`pulumi.Input[str]`) - The access type of the rule. The only possible value is `Allow`.
          * `communities` (`pulumi.Input[list]`) - The collection for bgp community values to filter on. e.g. ['12076:5010','12076:5020'].
          * `name` (`pulumi.Input[str]`) - The name of the route filter rule.
          * `ruleType` (`pulumi.Input[str]`) - The rule type of the rule. The only possible value is `Community`.
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

            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['rule'] = rule
            __props__['tags'] = tags
        super(RouteFilter, __self__).__init__(
            'azure:network/routeFilter:RouteFilter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, location=None, name=None, resource_group_name=None, rule=None, tags=None):
        """
        Get an existing RouteFilter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The Azure Region where the Route Filter should exist. Changing this forces a new Route Filter to be created.
        :param pulumi.Input[str] name: The Name which should be used for this Route Filter.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Route Filter should exist. Changing this forces a new Route Filter to be created.
        :param pulumi.Input[dict] rule: A `rules` block as defined below.
        :param pulumi.Input[dict] tags: A mapping of tags which should be assigned to the Route Filter.

        The **rule** object supports the following:

          * `access` (`pulumi.Input[str]`) - The access type of the rule. The only possible value is `Allow`.
          * `communities` (`pulumi.Input[list]`) - The collection for bgp community values to filter on. e.g. ['12076:5010','12076:5020'].
          * `name` (`pulumi.Input[str]`) - The name of the route filter rule.
          * `ruleType` (`pulumi.Input[str]`) - The rule type of the rule. The only possible value is `Community`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["rule"] = rule
        __props__["tags"] = tags
        return RouteFilter(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

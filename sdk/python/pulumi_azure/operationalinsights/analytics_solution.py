# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class AnalyticsSolution(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    plan: pulumi.Output[dict]
    """
    A `plan` block as documented below.

      * `name` (`str`)
      * `product` (`str`) - The product name of the solution. For example `OMSGallery/Containers`. Changing this forces a new resource to be created.
      * `promotionCode` (`str`) - A promotion code to be used with the solution.
      * `publisher` (`str`) - The publisher of the solution. For example `Microsoft`. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which the Log Analytics solution is created. Changing this forces a new resource to be created. Note: The solution and its related workspace can only exist in the same resource group.
    """
    solution_name: pulumi.Output[str]
    """
    Specifies the name of the solution to be deployed. See [here for options](https://docs.microsoft.com/en-us/azure/log-analytics/log-analytics-add-solutions).Changing this forces a new resource to be created.
    """
    workspace_name: pulumi.Output[str]
    """
    The full name of the Log Analytics workspace with which the solution will be linked. Changing this forces a new resource to be created.
    """
    workspace_resource_id: pulumi.Output[str]
    """
    The full resource ID of the Log Analytics workspace with which the solution will be linked. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, location=None, plan=None, resource_group_name=None, solution_name=None, workspace_name=None, workspace_resource_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Log Analytics (formally Operational Insights) Solution.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure
        import pulumi_random as random

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="westeurope")
        workspace = random.RandomId("workspace",
            keepers={
                "group_name": example_resource_group.name,
            },
            byte_length=8)
        example_analytics_workspace = azure.operationalinsights.AnalyticsWorkspace("exampleAnalyticsWorkspace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="PerGB2018")
        example_analytics_solution = azure.operationalinsights.AnalyticsSolution("exampleAnalyticsSolution",
            solution_name="ContainerInsights",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            workspace_resource_id=example_analytics_workspace.id,
            workspace_name=example_analytics_workspace.name,
            plan={
                "publisher": "Microsoft",
                "product": "OMSGallery/ContainerInsights",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] plan: A `plan` block as documented below.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Log Analytics solution is created. Changing this forces a new resource to be created. Note: The solution and its related workspace can only exist in the same resource group.
        :param pulumi.Input[str] solution_name: Specifies the name of the solution to be deployed. See [here for options](https://docs.microsoft.com/en-us/azure/log-analytics/log-analytics-add-solutions).Changing this forces a new resource to be created.
        :param pulumi.Input[str] workspace_name: The full name of the Log Analytics workspace with which the solution will be linked. Changing this forces a new resource to be created.
        :param pulumi.Input[str] workspace_resource_id: The full resource ID of the Log Analytics workspace with which the solution will be linked. Changing this forces a new resource to be created.

        The **plan** object supports the following:

          * `name` (`pulumi.Input[str]`)
          * `product` (`pulumi.Input[str]`) - The product name of the solution. For example `OMSGallery/Containers`. Changing this forces a new resource to be created.
          * `promotionCode` (`pulumi.Input[str]`) - A promotion code to be used with the solution.
          * `publisher` (`pulumi.Input[str]`) - The publisher of the solution. For example `Microsoft`. Changing this forces a new resource to be created.
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
            if plan is None:
                raise TypeError("Missing required property 'plan'")
            __props__['plan'] = plan
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if solution_name is None:
                raise TypeError("Missing required property 'solution_name'")
            __props__['solution_name'] = solution_name
            if workspace_name is None:
                raise TypeError("Missing required property 'workspace_name'")
            __props__['workspace_name'] = workspace_name
            if workspace_resource_id is None:
                raise TypeError("Missing required property 'workspace_resource_id'")
            __props__['workspace_resource_id'] = workspace_resource_id
        super(AnalyticsSolution, __self__).__init__(
            'azure:operationalinsights/analyticsSolution:AnalyticsSolution',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, location=None, plan=None, resource_group_name=None, solution_name=None, workspace_name=None, workspace_resource_id=None):
        """
        Get an existing AnalyticsSolution resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] plan: A `plan` block as documented below.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Log Analytics solution is created. Changing this forces a new resource to be created. Note: The solution and its related workspace can only exist in the same resource group.
        :param pulumi.Input[str] solution_name: Specifies the name of the solution to be deployed. See [here for options](https://docs.microsoft.com/en-us/azure/log-analytics/log-analytics-add-solutions).Changing this forces a new resource to be created.
        :param pulumi.Input[str] workspace_name: The full name of the Log Analytics workspace with which the solution will be linked. Changing this forces a new resource to be created.
        :param pulumi.Input[str] workspace_resource_id: The full resource ID of the Log Analytics workspace with which the solution will be linked. Changing this forces a new resource to be created.

        The **plan** object supports the following:

          * `name` (`pulumi.Input[str]`)
          * `product` (`pulumi.Input[str]`) - The product name of the solution. For example `OMSGallery/Containers`. Changing this forces a new resource to be created.
          * `promotionCode` (`pulumi.Input[str]`) - A promotion code to be used with the solution.
          * `publisher` (`pulumi.Input[str]`) - The publisher of the solution. For example `Microsoft`. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["location"] = location
        __props__["plan"] = plan
        __props__["resource_group_name"] = resource_group_name
        __props__["solution_name"] = solution_name
        __props__["workspace_name"] = workspace_name
        __props__["workspace_resource_id"] = workspace_resource_id
        return AnalyticsSolution(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

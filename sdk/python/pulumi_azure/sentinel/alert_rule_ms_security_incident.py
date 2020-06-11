# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class AlertRuleMsSecurityIncident(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    The description of this Sentinel MS Security Incident Alert Rule.
    """
    display_name: pulumi.Output[str]
    """
    The friendly name of this Sentinel MS Security Incident Alert Rule.
    """
    enabled: pulumi.Output[bool]
    """
    Should this Sentinel MS Security Incident Alert Rule be enabled? Defaults to `true`.
    """
    log_analytics_workspace_id: pulumi.Output[str]
    """
    The ID of the Log Analytics Workspace this Sentinel MS Security Incident Alert Rule belongs to. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
    """
    name: pulumi.Output[str]
    """
    The name which should be used for this Sentinel MS Security Incident Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
    """
    product_filter: pulumi.Output[str]
    """
    The Microsoft Security Service from where the alert will be generated. Possible values are `Azure Active Directory Identity Protection`, `Azure Advanced Threat Protection`, `Azure Security Center`, `Azure Security Center for IoT` and `Microsoft Cloud App Security`.
    """
    severity_filters: pulumi.Output[list]
    """
    Only create incidents from alerts when alert severity level is contained in this list. Possible values are `High`, `Medium`, `Low` and `Informational`.
    """
    text_whitelists: pulumi.Output[list]
    """
    Only create incidents from alerts when alert name contain text in this list. No filter will happen if this field is absent.
    """
    def __init__(__self__, resource_name, opts=None, description=None, display_name=None, enabled=None, log_analytics_workspace_id=None, name=None, product_filter=None, severity_filters=None, text_whitelists=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Sentinel MS Security Incident Alert Rule.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_analytics_workspace = azure.operationalinsights.AnalyticsWorkspace("exampleAnalyticsWorkspace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="pergb2018")
        example_alert_rule_ms_security_incident = azure.sentinel.AlertRuleMsSecurityIncident("exampleAlertRuleMsSecurityIncident",
            log_analytics_workspace_id=example_analytics_workspace.id,
            product_filter="Microsoft Cloud App Security",
            display_name="example rule",
            severity_filters=["High"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of this Sentinel MS Security Incident Alert Rule.
        :param pulumi.Input[str] display_name: The friendly name of this Sentinel MS Security Incident Alert Rule.
        :param pulumi.Input[bool] enabled: Should this Sentinel MS Security Incident Alert Rule be enabled? Defaults to `true`.
        :param pulumi.Input[str] log_analytics_workspace_id: The ID of the Log Analytics Workspace this Sentinel MS Security Incident Alert Rule belongs to. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
        :param pulumi.Input[str] name: The name which should be used for this Sentinel MS Security Incident Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
        :param pulumi.Input[str] product_filter: The Microsoft Security Service from where the alert will be generated. Possible values are `Azure Active Directory Identity Protection`, `Azure Advanced Threat Protection`, `Azure Security Center`, `Azure Security Center for IoT` and `Microsoft Cloud App Security`.
        :param pulumi.Input[list] severity_filters: Only create incidents from alerts when alert severity level is contained in this list. Possible values are `High`, `Medium`, `Low` and `Informational`.
        :param pulumi.Input[list] text_whitelists: Only create incidents from alerts when alert name contain text in this list. No filter will happen if this field is absent.
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

            __props__['description'] = description
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            __props__['enabled'] = enabled
            if log_analytics_workspace_id is None:
                raise TypeError("Missing required property 'log_analytics_workspace_id'")
            __props__['log_analytics_workspace_id'] = log_analytics_workspace_id
            __props__['name'] = name
            if product_filter is None:
                raise TypeError("Missing required property 'product_filter'")
            __props__['product_filter'] = product_filter
            if severity_filters is None:
                raise TypeError("Missing required property 'severity_filters'")
            __props__['severity_filters'] = severity_filters
            __props__['text_whitelists'] = text_whitelists
        super(AlertRuleMsSecurityIncident, __self__).__init__(
            'azure:sentinel/alertRuleMsSecurityIncident:AlertRuleMsSecurityIncident',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, display_name=None, enabled=None, log_analytics_workspace_id=None, name=None, product_filter=None, severity_filters=None, text_whitelists=None):
        """
        Get an existing AlertRuleMsSecurityIncident resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of this Sentinel MS Security Incident Alert Rule.
        :param pulumi.Input[str] display_name: The friendly name of this Sentinel MS Security Incident Alert Rule.
        :param pulumi.Input[bool] enabled: Should this Sentinel MS Security Incident Alert Rule be enabled? Defaults to `true`.
        :param pulumi.Input[str] log_analytics_workspace_id: The ID of the Log Analytics Workspace this Sentinel MS Security Incident Alert Rule belongs to. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
        :param pulumi.Input[str] name: The name which should be used for this Sentinel MS Security Incident Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
        :param pulumi.Input[str] product_filter: The Microsoft Security Service from where the alert will be generated. Possible values are `Azure Active Directory Identity Protection`, `Azure Advanced Threat Protection`, `Azure Security Center`, `Azure Security Center for IoT` and `Microsoft Cloud App Security`.
        :param pulumi.Input[list] severity_filters: Only create incidents from alerts when alert severity level is contained in this list. Possible values are `High`, `Medium`, `Low` and `Informational`.
        :param pulumi.Input[list] text_whitelists: Only create incidents from alerts when alert name contain text in this list. No filter will happen if this field is absent.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["display_name"] = display_name
        __props__["enabled"] = enabled
        __props__["log_analytics_workspace_id"] = log_analytics_workspace_id
        __props__["name"] = name
        __props__["product_filter"] = product_filter
        __props__["severity_filters"] = severity_filters
        __props__["text_whitelists"] = text_whitelists
        return AlertRuleMsSecurityIncident(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


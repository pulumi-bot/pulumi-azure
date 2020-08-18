# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['BoolVariable']


class BoolVariable(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 automation_account_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encrypted: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a boolean variable in Azure Automation

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] description: The description of the Automation Variable.
        :param pulumi.Input[bool] encrypted: Specifies if the Automation Variable is encrypted. Defaults to `false`.
        :param pulumi.Input[str] name: The name of the Automation Variable. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] value: The value of the Automation Variable as a `boolean`.
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

            if automation_account_name is None:
                raise TypeError("Missing required property 'automation_account_name'")
            __props__['automation_account_name'] = automation_account_name
            __props__['description'] = description
            __props__['encrypted'] = encrypted
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['value'] = value
        super(BoolVariable, __self__).__init__(
            'azure:automation/boolVariable:BoolVariable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            automation_account_name: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            encrypted: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[bool]] = None) -> 'BoolVariable':
        """
        Get an existing BoolVariable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] description: The description of the Automation Variable.
        :param pulumi.Input[bool] encrypted: Specifies if the Automation Variable is encrypted. Defaults to `false`.
        :param pulumi.Input[str] name: The name of the Automation Variable. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] value: The value of the Automation Variable as a `boolean`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["automation_account_name"] = automation_account_name
        __props__["description"] = description
        __props__["encrypted"] = encrypted
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["value"] = value
        return BoolVariable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="automationAccountName")
    def automation_account_name(self) -> str:
        """
        The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "automation_account_name")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the Automation Variable.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def encrypted(self) -> Optional[bool]:
        """
        Specifies if the Automation Variable is encrypted. Defaults to `false`.
        """
        return pulumi.get(self, "encrypted")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the Automation Variable. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def value(self) -> Optional[bool]:
        """
        The value of the Automation Variable as a `boolean`.
        """
        return pulumi.get(self, "value")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


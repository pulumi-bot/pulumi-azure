# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['DscNodeConfiguration']


class DscNodeConfiguration(pulumi.CustomResource):
    automation_account_name: pulumi.Output[str] = pulumi.output_property("automationAccountName")
    """
    The name of the automation account in which the DSC Node Configuration is created. Changing this forces a new resource to be created.
    """
    configuration_name: pulumi.Output[str] = pulumi.output_property("configurationName")
    content_embedded: pulumi.Output[str] = pulumi.output_property("contentEmbedded")
    """
    The PowerShell DSC Node Configuration (mof content).
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    Specifies the name of the DSC Node Configuration. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str] = pulumi.output_property("resourceGroupName")
    """
    The name of the resource group in which the DSC Node Configuration is created. Changing this forces a new resource to be created.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, automation_account_name: Optional[pulumi.Input[str]] = None, content_embedded: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manages a Automation DSC Node Configuration.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.automation.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku=[{
                "name": "Basic",
            }])
        example_dsc_configuration = azure.automation.DscConfiguration("exampleDscConfiguration",
            resource_group_name=example_resource_group.name,
            automation_account_name=example_account.name,
            location=example_resource_group.location,
            content_embedded="configuration test {}")
        example_dsc_node_configuration = azure.automation.DscNodeConfiguration("exampleDscNodeConfiguration",
            resource_group_name=example_resource_group.name,
            automation_account_name=example_account.name,
            content_embedded=\"\"\"instance of MSFT_FileDirectoryConfiguration as $MSFT_FileDirectoryConfiguration1ref
        {
          ResourceID = "[File]bla";
          Ensure = "Present";
          Contents = "bogus Content";
          DestinationPath = "c:\\bogus.txt";
          ModuleName = "PSDesiredStateConfiguration";
          SourceInfo = "::3::9::file";
          ModuleVersion = "1.0";
          ConfigurationName = "bla";
        };
        instance of OMI_ConfigurationDocument
        {
          Version="2.0.0";
          MinimumCompatibleVersion = "1.0.0";
          CompatibleVersionAdditionalProperties= {"Omi_BaseResource:ConfigurationName"};
          Author="bogusAuthor";
          GenerationDate="06/15/2018 14:06:24";
          GenerationHost="bogusComputer";
          Name="test";
        };
        \"\"\",
            opts=ResourceOptions(depends_on=[example_dsc_configuration]))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account in which the DSC Node Configuration is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] content_embedded: The PowerShell DSC Node Configuration (mof content).
        :param pulumi.Input[str] name: Specifies the name of the DSC Node Configuration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the DSC Node Configuration is created. Changing this forces a new resource to be created.
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
            if content_embedded is None:
                raise TypeError("Missing required property 'content_embedded'")
            __props__['content_embedded'] = content_embedded
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['configuration_name'] = None
        super(DscNodeConfiguration, __self__).__init__(
            'azure:automation/dscNodeConfiguration:DscNodeConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, automation_account_name: Optional[pulumi.Input[str]] = None, configuration_name: Optional[pulumi.Input[str]] = None, content_embedded: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None) -> 'DscNodeConfiguration':
        """
        Get an existing DscNodeConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account in which the DSC Node Configuration is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] content_embedded: The PowerShell DSC Node Configuration (mof content).
        :param pulumi.Input[str] name: Specifies the name of the DSC Node Configuration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the DSC Node Configuration is created. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["automation_account_name"] = automation_account_name
        __props__["configuration_name"] = configuration_name
        __props__["content_embedded"] = content_embedded
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        return DscNodeConfiguration(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


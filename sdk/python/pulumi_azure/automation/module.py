# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Module(pulumi.CustomResource):
    automation_account_name: pulumi.Output[str]
    """
    The name of the automation account in which the Module is created. Changing this forces a new resource to be created.
    """
    module_link: pulumi.Output[dict]
    """
    The published Module link.

      * `hash` (`dict`)
        * `algorithm` (`str`)
        * `value` (`str`)

      * `uri` (`str`) - The uri of the module content (zip or nupkg).
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Module. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which the Module is created. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, automation_account_name=None, module_link=None, name=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Automation Module.

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
        example_module = azure.automation.Module("exampleModule",
            resource_group_name=example_resource_group.name,
            automation_account_name=example_account.name,
            module_link={
                "uri": "https://devopsgallerystorage.blob.core.windows.net/packages/xactivedirectory.2.19.0.nupkg",
            })
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account in which the Module is created. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] module_link: The published Module link.
        :param pulumi.Input[str] name: Specifies the name of the Module. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Module is created. Changing this forces a new resource to be created.

        The **module_link** object supports the following:

          * `hash` (`pulumi.Input[dict]`)
            * `algorithm` (`pulumi.Input[str]`)
            * `value` (`pulumi.Input[str]`)

          * `uri` (`pulumi.Input[str]`) - The uri of the module content (zip or nupkg).
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

            if automation_account_name is None:
                raise TypeError("Missing required property 'automation_account_name'")
            __props__['automation_account_name'] = automation_account_name
            if module_link is None:
                raise TypeError("Missing required property 'module_link'")
            __props__['module_link'] = module_link
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
        super(Module, __self__).__init__(
            'azure:automation/module:Module',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, automation_account_name=None, module_link=None, name=None, resource_group_name=None):
        """
        Get an existing Module resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account in which the Module is created. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] module_link: The published Module link.
        :param pulumi.Input[str] name: Specifies the name of the Module. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Module is created. Changing this forces a new resource to be created.

        The **module_link** object supports the following:

          * `hash` (`pulumi.Input[dict]`)
            * `algorithm` (`pulumi.Input[str]`)
            * `value` (`pulumi.Input[str]`)

          * `uri` (`pulumi.Input[str]`) - The uri of the module content (zip or nupkg).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["automation_account_name"] = automation_account_name
        __props__["module_link"] = module_link
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        return Module(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


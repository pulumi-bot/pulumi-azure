# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Application(pulumi.CustomResource):
    display_name: pulumi.Output[str]
    """
    A `display_name` name. Custom display name for the IoT Central application. Default is resource name.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource has to be create. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
    """
    sku: pulumi.Output[str]
    """
    A `sku` name. Possible values is `ST1`, `ST2`, Default value is `ST1`
    """
    sub_domain: pulumi.Output[str]
    """
    A `sub_domain` name. Subdomain for the IoT Central URL. Each application must have a unique subdomain.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    template: pulumi.Output[str]
    """
    A `template` name. IoT Central application template name. Default is a custom application.
    """
    def __init__(__self__, resource_name, opts=None, display_name=None, location=None, name=None, resource_group_name=None, sku=None, sub_domain=None, tags=None, template=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an IoT Central Application

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_application = azure.iotcentral.Application("exampleApplication",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            sub_domain="example-iotcentral-app-subdomain",
            display_name="example-iotcentral-app-display-name",
            sku="S1",
            template="iotc-default@1.0.0",
            tags={
                "Foo": "Bar",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: A `display_name` name. Custom display name for the IoT Central application. Default is resource name.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource has to be create. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku: A `sku` name. Possible values is `ST1`, `ST2`, Default value is `ST1`
        :param pulumi.Input[str] sub_domain: A `sub_domain` name. Subdomain for the IoT Central URL. Each application must have a unique subdomain.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] template: A `template` name. IoT Central application template name. Default is a custom application.
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

            __props__['display_name'] = display_name
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            if sub_domain is None:
                raise TypeError("Missing required property 'sub_domain'")
            __props__['sub_domain'] = sub_domain
            __props__['tags'] = tags
            __props__['template'] = template
        super(Application, __self__).__init__(
            'azure:iotcentral/application:Application',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, display_name=None, location=None, name=None, resource_group_name=None, sku=None, sub_domain=None, tags=None, template=None):
        """
        Get an existing Application resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: A `display_name` name. Custom display name for the IoT Central application. Default is resource name.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource has to be create. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku: A `sku` name. Possible values is `ST1`, `ST2`, Default value is `ST1`
        :param pulumi.Input[str] sub_domain: A `sub_domain` name. Subdomain for the IoT Central URL. Each application must have a unique subdomain.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] template: A `template` name. IoT Central application template name. Default is a custom application.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["display_name"] = display_name
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["sku"] = sku
        __props__["sub_domain"] = sub_domain
        __props__["tags"] = tags
        __props__["template"] = template
        return Application(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


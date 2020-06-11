# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class ConfigurationStore(pulumi.CustomResource):
    endpoint: pulumi.Output[str]
    """
    The URL of the App Configuration.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the App Configuration. Changing this forces a new resource to be created.
    """
    primary_read_keys: pulumi.Output[list]
    """
    A `primary_read_key` block as defined below containing the primary read access key.

      * `connection_string` (`str`) - The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.
      * `id` (`str`) - The ID of the Access Key.
      * `secret` (`str`) - The Secret of the Access Key.
    """
    primary_write_keys: pulumi.Output[list]
    """
    A `primary_write_key` block as defined below containing the primary write access key.

      * `connection_string` (`str`) - The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.
      * `id` (`str`) - The ID of the Access Key.
      * `secret` (`str`) - The Secret of the Access Key.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the App Configuration. Changing this forces a new resource to be created.
    """
    secondary_read_keys: pulumi.Output[list]
    """
    A `secondary_read_key` block as defined below containing the secondary read access key.

      * `connection_string` (`str`) - The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.
      * `id` (`str`) - The ID of the Access Key.
      * `secret` (`str`) - The Secret of the Access Key.
    """
    secondary_write_keys: pulumi.Output[list]
    """
    A `secondary_write_key` block as defined below containing the secondary write access key.

      * `connection_string` (`str`) - The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.
      * `id` (`str`) - The ID of the Access Key.
      * `secret` (`str`) - The Secret of the Access Key.
    """
    sku: pulumi.Output[str]
    """
    The SKU name of the the App Configuration. Possible values are `free` and `standard`.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, resource_group_name=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Azure App Configuration.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        rg = azure.core.ResourceGroup("rg", location="West Europe")
        appconf = azure.appconfiguration.ConfigurationStore("appconf",
            resource_group_name=rg.name,
            location=rg.location)
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the App Configuration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the App Configuration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku: The SKU name of the the App Configuration. Possible values are `free` and `standard`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
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

            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['endpoint'] = None
            __props__['primary_read_keys'] = None
            __props__['primary_write_keys'] = None
            __props__['secondary_read_keys'] = None
            __props__['secondary_write_keys'] = None
        super(ConfigurationStore, __self__).__init__(
            'azure:appconfiguration/configurationStore:ConfigurationStore',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, endpoint=None, location=None, name=None, primary_read_keys=None, primary_write_keys=None, resource_group_name=None, secondary_read_keys=None, secondary_write_keys=None, sku=None, tags=None):
        """
        Get an existing ConfigurationStore resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint: The URL of the App Configuration.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the App Configuration. Changing this forces a new resource to be created.
        :param pulumi.Input[list] primary_read_keys: A `primary_read_key` block as defined below containing the primary read access key.
        :param pulumi.Input[list] primary_write_keys: A `primary_write_key` block as defined below containing the primary write access key.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the App Configuration. Changing this forces a new resource to be created.
        :param pulumi.Input[list] secondary_read_keys: A `secondary_read_key` block as defined below containing the secondary read access key.
        :param pulumi.Input[list] secondary_write_keys: A `secondary_write_key` block as defined below containing the secondary write access key.
        :param pulumi.Input[str] sku: The SKU name of the the App Configuration. Possible values are `free` and `standard`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **primary_read_keys** object supports the following:

          * `connection_string` (`pulumi.Input[str]`) - The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.
          * `id` (`pulumi.Input[str]`) - The ID of the Access Key.
          * `secret` (`pulumi.Input[str]`) - The Secret of the Access Key.

        The **primary_write_keys** object supports the following:

          * `connection_string` (`pulumi.Input[str]`) - The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.
          * `id` (`pulumi.Input[str]`) - The ID of the Access Key.
          * `secret` (`pulumi.Input[str]`) - The Secret of the Access Key.

        The **secondary_read_keys** object supports the following:

          * `connection_string` (`pulumi.Input[str]`) - The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.
          * `id` (`pulumi.Input[str]`) - The ID of the Access Key.
          * `secret` (`pulumi.Input[str]`) - The Secret of the Access Key.

        The **secondary_write_keys** object supports the following:

          * `connection_string` (`pulumi.Input[str]`) - The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.
          * `id` (`pulumi.Input[str]`) - The ID of the Access Key.
          * `secret` (`pulumi.Input[str]`) - The Secret of the Access Key.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["endpoint"] = endpoint
        __props__["location"] = location
        __props__["name"] = name
        __props__["primary_read_keys"] = primary_read_keys
        __props__["primary_write_keys"] = primary_write_keys
        __props__["resource_group_name"] = resource_group_name
        __props__["secondary_read_keys"] = secondary_read_keys
        __props__["secondary_write_keys"] = secondary_write_keys
        __props__["sku"] = sku
        __props__["tags"] = tags
        return ConfigurationStore(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

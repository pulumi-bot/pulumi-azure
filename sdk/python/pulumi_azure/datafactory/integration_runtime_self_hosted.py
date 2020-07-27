# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class IntegrationRuntimeSelfHosted(pulumi.CustomResource):
    auth_key1: pulumi.Output[str]
    """
    The primary integration runtime authentication key.
    """
    auth_key2: pulumi.Output[str]
    """
    The secondary integration runtime authentication key.
    """
    data_factory_name: pulumi.Output[str]
    """
    Changing this forces a new Data Factory to be created.
    """
    description: pulumi.Output[str]
    """
    Integration runtime description.
    """
    name: pulumi.Output[str]
    """
    The name which should be used for this Data Factory. Changing this forces a new Data Factory to be created.
    """
    rbac_authorizations: pulumi.Output[list]
    """
    A `rbac_authorization` block as defined below.

      * `resource_id` (`str`) - The resource identifier of the integration runtime to be shared. Changing this forces a new Data Factory to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group where the Data Factory should exist. Changing this forces a new Data Factory to be created.
    """
    def __init__(__self__, resource_name, opts=None, data_factory_name=None, description=None, name=None, rbac_authorizations=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Data Factory Self-host Integration Runtime.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="eastus")
        example_factory = azure.datafactory.Factory("exampleFactory",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_integration_runtime_self_hosted = azure.datafactory.IntegrationRuntimeSelfHosted("exampleIntegrationRuntimeSelfHosted",
            resource_group_name="example",
            data_factory_name="example")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_factory_name: Changing this forces a new Data Factory to be created.
        :param pulumi.Input[str] description: Integration runtime description.
        :param pulumi.Input[str] name: The name which should be used for this Data Factory. Changing this forces a new Data Factory to be created.
        :param pulumi.Input[list] rbac_authorizations: A `rbac_authorization` block as defined below.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Data Factory should exist. Changing this forces a new Data Factory to be created.

        The **rbac_authorizations** object supports the following:

          * `resource_id` (`pulumi.Input[str]`) - The resource identifier of the integration runtime to be shared. Changing this forces a new Data Factory to be created.
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

            if data_factory_name is None:
                raise TypeError("Missing required property 'data_factory_name'")
            __props__['data_factory_name'] = data_factory_name
            __props__['description'] = description
            __props__['name'] = name
            __props__['rbac_authorizations'] = rbac_authorizations
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['auth_key1'] = None
            __props__['auth_key2'] = None
        super(IntegrationRuntimeSelfHosted, __self__).__init__(
            'azure:datafactory/integrationRuntimeSelfHosted:IntegrationRuntimeSelfHosted',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, auth_key1=None, auth_key2=None, data_factory_name=None, description=None, name=None, rbac_authorizations=None, resource_group_name=None):
        """
        Get an existing IntegrationRuntimeSelfHosted resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_key1: The primary integration runtime authentication key.
        :param pulumi.Input[str] auth_key2: The secondary integration runtime authentication key.
        :param pulumi.Input[str] data_factory_name: Changing this forces a new Data Factory to be created.
        :param pulumi.Input[str] description: Integration runtime description.
        :param pulumi.Input[str] name: The name which should be used for this Data Factory. Changing this forces a new Data Factory to be created.
        :param pulumi.Input[list] rbac_authorizations: A `rbac_authorization` block as defined below.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Data Factory should exist. Changing this forces a new Data Factory to be created.

        The **rbac_authorizations** object supports the following:

          * `resource_id` (`pulumi.Input[str]`) - The resource identifier of the integration runtime to be shared. Changing this forces a new Data Factory to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auth_key1"] = auth_key1
        __props__["auth_key2"] = auth_key2
        __props__["data_factory_name"] = data_factory_name
        __props__["description"] = description
        __props__["name"] = name
        __props__["rbac_authorizations"] = rbac_authorizations
        __props__["resource_group_name"] = resource_group_name
        return IntegrationRuntimeSelfHosted(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

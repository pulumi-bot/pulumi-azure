# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class ConnectionServicePrincipal(pulumi.CustomResource):
    application_id: pulumi.Output[str]
    """
    The (Client) ID of the Service Principal.
    """
    automation_account_name: pulumi.Output[str]
    """
    The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
    """
    certificate_thumbprint: pulumi.Output[str]
    """
    The thumbprint of the Service Principal Certificate.
    """
    description: pulumi.Output[str]
    """
    A description for this Connection.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Connection. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
    """
    subscription_id: pulumi.Output[str]
    """
    The subscription GUID.
    """
    tenant_id: pulumi.Output[str]
    """
    The ID of the Tenant the Service Principal is assigned in.
    """
    def __init__(__self__, resource_name, opts=None, application_id=None, automation_account_name=None, certificate_thumbprint=None, description=None, name=None, resource_group_name=None, subscription_id=None, tenant_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Automation Connection with type `AzureServicePrincipal`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_client_config = azure.core.get_client_config()
        example_account = azure.automation.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku=[{
                "name": "Basic",
            }])
        example_connection_service_principal = azure.automation.ConnectionServicePrincipal("exampleConnectionServicePrincipal",
            resource_group_name=example_resource_group.name,
            automation_account_name=example_account.name,
            application_id="00000000-0000-0000-0000-000000000000",
            tenant_id=example_client_config.tenant_id,
            subscription_id=example_client_config.subscription_id,
            certificate_thumbprint=(lambda path: open(path).read())("automation_certificate_test.thumb"))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The (Client) ID of the Service Principal.
        :param pulumi.Input[str] automation_account_name: The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] certificate_thumbprint: The thumbprint of the Service Principal Certificate.
        :param pulumi.Input[str] description: A description for this Connection.
        :param pulumi.Input[str] name: Specifies the name of the Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subscription_id: The subscription GUID.
        :param pulumi.Input[str] tenant_id: The ID of the Tenant the Service Principal is assigned in.
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

            if application_id is None:
                raise TypeError("Missing required property 'application_id'")
            __props__['application_id'] = application_id
            if automation_account_name is None:
                raise TypeError("Missing required property 'automation_account_name'")
            __props__['automation_account_name'] = automation_account_name
            if certificate_thumbprint is None:
                raise TypeError("Missing required property 'certificate_thumbprint'")
            __props__['certificate_thumbprint'] = certificate_thumbprint
            __props__['description'] = description
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if subscription_id is None:
                raise TypeError("Missing required property 'subscription_id'")
            __props__['subscription_id'] = subscription_id
            if tenant_id is None:
                raise TypeError("Missing required property 'tenant_id'")
            __props__['tenant_id'] = tenant_id
        super(ConnectionServicePrincipal, __self__).__init__(
            'azure:automation/connectionServicePrincipal:ConnectionServicePrincipal',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, application_id=None, automation_account_name=None, certificate_thumbprint=None, description=None, name=None, resource_group_name=None, subscription_id=None, tenant_id=None):
        """
        Get an existing ConnectionServicePrincipal resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The (Client) ID of the Service Principal.
        :param pulumi.Input[str] automation_account_name: The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] certificate_thumbprint: The thumbprint of the Service Principal Certificate.
        :param pulumi.Input[str] description: A description for this Connection.
        :param pulumi.Input[str] name: Specifies the name of the Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subscription_id: The subscription GUID.
        :param pulumi.Input[str] tenant_id: The ID of the Tenant the Service Principal is assigned in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["application_id"] = application_id
        __props__["automation_account_name"] = automation_account_name
        __props__["certificate_thumbprint"] = certificate_thumbprint
        __props__["description"] = description
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["subscription_id"] = subscription_id
        __props__["tenant_id"] = tenant_id
        return ConnectionServicePrincipal(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

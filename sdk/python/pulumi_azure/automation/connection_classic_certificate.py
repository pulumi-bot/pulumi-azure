# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['ConnectionClassicCertificate']


class ConnectionClassicCertificate(pulumi.CustomResource):
    automation_account_name: pulumi.Output[str] = pulumi.output_property("automationAccountName")
    """
    The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
    """
    certificate_asset_name: pulumi.Output[str] = pulumi.output_property("certificateAssetName")
    """
    The name of the certificate asset.
    """
    description: pulumi.Output[Optional[str]] = pulumi.output_property("description")
    """
    A description for this Connection.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    Specifies the name of the Connection. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str] = pulumi.output_property("resourceGroupName")
    """
    The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
    """
    subscription_id: pulumi.Output[str] = pulumi.output_property("subscriptionId")
    """
    The id of subscription.
    """
    subscription_name: pulumi.Output[str] = pulumi.output_property("subscriptionName")
    """
    The name of subscription.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, automation_account_name: Optional[pulumi.Input[str]] = None, certificate_asset_name: Optional[pulumi.Input[str]] = None, description: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, subscription_id: Optional[pulumi.Input[str]] = None, subscription_name: Optional[pulumi.Input[str]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manages an Automation Connection with type `AzureClassicCertificate`.

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
        example_connection_classic_certificate = azure.automation.ConnectionClassicCertificate("exampleConnectionClassicCertificate",
            resource_group_name=example_resource_group.name,
            automation_account_name=example_account.name,
            certificate_asset_name="cert1",
            subscription_name="subs1",
            subscription_id=example_client_config.subscription_id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] certificate_asset_name: The name of the certificate asset.
        :param pulumi.Input[str] description: A description for this Connection.
        :param pulumi.Input[str] name: Specifies the name of the Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subscription_id: The id of subscription.
        :param pulumi.Input[str] subscription_name: The name of subscription.
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
            if certificate_asset_name is None:
                raise TypeError("Missing required property 'certificate_asset_name'")
            __props__['certificate_asset_name'] = certificate_asset_name
            __props__['description'] = description
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if subscription_id is None:
                raise TypeError("Missing required property 'subscription_id'")
            __props__['subscription_id'] = subscription_id
            if subscription_name is None:
                raise TypeError("Missing required property 'subscription_name'")
            __props__['subscription_name'] = subscription_name
        super(ConnectionClassicCertificate, __self__).__init__(
            'azure:automation/connectionClassicCertificate:ConnectionClassicCertificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, automation_account_name: Optional[pulumi.Input[str]] = None, certificate_asset_name: Optional[pulumi.Input[str]] = None, description: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, resource_group_name: Optional[pulumi.Input[str]] = None, subscription_id: Optional[pulumi.Input[str]] = None, subscription_name: Optional[pulumi.Input[str]] = None) -> 'ConnectionClassicCertificate':
        """
        Get an existing ConnectionClassicCertificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] certificate_asset_name: The name of the certificate asset.
        :param pulumi.Input[str] description: A description for this Connection.
        :param pulumi.Input[str] name: Specifies the name of the Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subscription_id: The id of subscription.
        :param pulumi.Input[str] subscription_name: The name of subscription.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["automation_account_name"] = automation_account_name
        __props__["certificate_asset_name"] = certificate_asset_name
        __props__["description"] = description
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["subscription_id"] = subscription_id
        __props__["subscription_name"] = subscription_name
        return ConnectionClassicCertificate(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


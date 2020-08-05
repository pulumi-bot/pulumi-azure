# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class IotHubCertificate(pulumi.CustomResource):
    certificate_content: pulumi.Output[str]
    """
    The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
    """
    iot_dps_name: pulumi.Output[str]
    """
    The name of the IoT Device Provisioning Service that this certificate will be attached to. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Iot Device Provisioning Service Certificate resource. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group under which the Iot Device Provisioning Service Certificate resource has to be created. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, certificate_content=None, iot_dps_name=None, name=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an IotHub Device Provisioning Service Certificate.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_content: The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
        :param pulumi.Input[str] iot_dps_name: The name of the IoT Device Provisioning Service that this certificate will be attached to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Iot Device Provisioning Service Certificate resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group under which the Iot Device Provisioning Service Certificate resource has to be created. Changing this forces a new resource to be created.
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

            if certificate_content is None:
                raise TypeError("Missing required property 'certificate_content'")
            __props__['certificate_content'] = certificate_content
            if iot_dps_name is None:
                raise TypeError("Missing required property 'iot_dps_name'")
            __props__['iot_dps_name'] = iot_dps_name
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
        super(IotHubCertificate, __self__).__init__(
            'azure:iot/iotHubCertificate:IotHubCertificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, certificate_content=None, iot_dps_name=None, name=None, resource_group_name=None):
        """
        Get an existing IotHubCertificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_content: The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
        :param pulumi.Input[str] iot_dps_name: The name of the IoT Device Provisioning Service that this certificate will be attached to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Iot Device Provisioning Service Certificate resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group under which the Iot Device Provisioning Service Certificate resource has to be created. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["certificate_content"] = certificate_content
        __props__["iot_dps_name"] = iot_dps_name
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        return IotHubCertificate(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

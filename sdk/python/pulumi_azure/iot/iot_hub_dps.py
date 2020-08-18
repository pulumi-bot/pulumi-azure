# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['IotHubDps']


class IotHubDps(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 linked_hubs: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['IotHubDpsLinkedHubArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['IotHubDpsSkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an IotHub Device Provisioning Service.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_iot_hub_dps = azure.iot.IotHubDps("exampleIotHubDps",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            sku={
                "name": "S1",
                "capacity": 1,
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['IotHubDpsLinkedHubArgs']]]] linked_hubs: A `linked_hub` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['IotHubDpsSkuArgs']] sku: A `sku` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
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

            __props__['linked_hubs'] = linked_hubs
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku is None:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['allocation_policy'] = None
            __props__['device_provisioning_host_name'] = None
            __props__['id_scope'] = None
            __props__['service_operations_host_name'] = None
        super(IotHubDps, __self__).__init__(
            'azure:iot/iotHubDps:IotHubDps',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            allocation_policy: Optional[pulumi.Input[str]] = None,
            device_provisioning_host_name: Optional[pulumi.Input[str]] = None,
            id_scope: Optional[pulumi.Input[str]] = None,
            linked_hubs: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['IotHubDpsLinkedHubArgs']]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            service_operations_host_name: Optional[pulumi.Input[str]] = None,
            sku: Optional[pulumi.Input[pulumi.InputType['IotHubDpsSkuArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'IotHubDps':
        """
        Get an existing IotHubDps resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allocation_policy: The allocation policy of the IoT Device Provisioning Service.
        :param pulumi.Input[str] device_provisioning_host_name: The device endpoint of the IoT Device Provisioning Service.
        :param pulumi.Input[str] id_scope: The unique identifier of the IoT Device Provisioning Service.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['IotHubDpsLinkedHubArgs']]]] linked_hubs: A `linked_hub` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_operations_host_name: The service endpoint of the IoT Device Provisioning Service.
        :param pulumi.Input[pulumi.InputType['IotHubDpsSkuArgs']] sku: A `sku` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allocation_policy"] = allocation_policy
        __props__["device_provisioning_host_name"] = device_provisioning_host_name
        __props__["id_scope"] = id_scope
        __props__["linked_hubs"] = linked_hubs
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["service_operations_host_name"] = service_operations_host_name
        __props__["sku"] = sku
        __props__["tags"] = tags
        return IotHubDps(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allocationPolicy")
    def allocation_policy(self) -> str:
        """
        The allocation policy of the IoT Device Provisioning Service.
        """
        return pulumi.get(self, "allocation_policy")

    @property
    @pulumi.getter(name="deviceProvisioningHostName")
    def device_provisioning_host_name(self) -> str:
        """
        The device endpoint of the IoT Device Provisioning Service.
        """
        return pulumi.get(self, "device_provisioning_host_name")

    @property
    @pulumi.getter(name="idScope")
    def id_scope(self) -> str:
        """
        The unique identifier of the IoT Device Provisioning Service.
        """
        return pulumi.get(self, "id_scope")

    @property
    @pulumi.getter(name="linkedHubs")
    def linked_hubs(self) -> Optional[List['outputs.IotHubDpsLinkedHub']]:
        """
        A `linked_hub` block as defined below.
        """
        return pulumi.get(self, "linked_hubs")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="serviceOperationsHostName")
    def service_operations_host_name(self) -> str:
        """
        The service endpoint of the IoT Device Provisioning Service.
        """
        return pulumi.get(self, "service_operations_host_name")

    @property
    @pulumi.getter
    def sku(self) -> 'outputs.IotHubDpsSku':
        """
        A `sku` block as defined below.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


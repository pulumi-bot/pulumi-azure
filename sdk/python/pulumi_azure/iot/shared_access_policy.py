# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['SharedAccessPolicy']


class SharedAccessPolicy(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_connect: Optional[pulumi.Input[bool]] = None,
                 iothub_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 registry_read: Optional[pulumi.Input[bool]] = None,
                 registry_write: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_connect: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an IotHub Shared Access Policy

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_io_t_hub = azure.iot.IoTHub("exampleIoTHub",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            sku=azure.iot.IoTHubSkuArgs(
                name="S1",
                capacity=1,
            ))
        example_shared_access_policy = azure.iot.SharedAccessPolicy("exampleSharedAccessPolicy",
            resource_group_name=example_resource_group.name,
            iothub_name=example_io_t_hub.name,
            registry_read=True,
            registry_write=True)
        ```

        ## Import

        IoTHub Shared Access Policies can be imported using the `resource id`, e.g.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] device_connect: Adds `DeviceConnect` permission to this Shared Access Account. It allows sending and receiving on the device-side endpoints.
        :param pulumi.Input[str] iothub_name: The name of the IoTHub to which this Shared Access Policy belongs. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the IotHub Shared Access Policy resource. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] registry_read: Adds `RegistryRead` permission to this Shared Access Account. It allows read access to the identity registry.
        :param pulumi.Input[bool] registry_write: Adds `RegistryWrite` permission to this Shared Access Account. It allows write access to the identity registry.
        :param pulumi.Input[str] resource_group_name: The name of the resource group under which the IotHub Shared Access Policy resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] service_connect: Adds `ServiceConnect` permission to this Shared Access Account. It allows sending and receiving on the cloud-side endpoints.
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

            __props__['device_connect'] = device_connect
            if iothub_name is None:
                raise TypeError("Missing required property 'iothub_name'")
            __props__['iothub_name'] = iothub_name
            __props__['name'] = name
            __props__['registry_read'] = registry_read
            __props__['registry_write'] = registry_write
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['service_connect'] = service_connect
            __props__['primary_connection_string'] = None
            __props__['primary_key'] = None
            __props__['secondary_connection_string'] = None
            __props__['secondary_key'] = None
        super(SharedAccessPolicy, __self__).__init__(
            'azure:iot/sharedAccessPolicy:SharedAccessPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            device_connect: Optional[pulumi.Input[bool]] = None,
            iothub_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            primary_connection_string: Optional[pulumi.Input[str]] = None,
            primary_key: Optional[pulumi.Input[str]] = None,
            registry_read: Optional[pulumi.Input[bool]] = None,
            registry_write: Optional[pulumi.Input[bool]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            secondary_connection_string: Optional[pulumi.Input[str]] = None,
            secondary_key: Optional[pulumi.Input[str]] = None,
            service_connect: Optional[pulumi.Input[bool]] = None) -> 'SharedAccessPolicy':
        """
        Get an existing SharedAccessPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] device_connect: Adds `DeviceConnect` permission to this Shared Access Account. It allows sending and receiving on the device-side endpoints.
        :param pulumi.Input[str] iothub_name: The name of the IoTHub to which this Shared Access Policy belongs. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the IotHub Shared Access Policy resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] primary_connection_string: The primary connection string of the Shared Access Policy.
        :param pulumi.Input[str] primary_key: The primary key used to create the authentication token.
        :param pulumi.Input[bool] registry_read: Adds `RegistryRead` permission to this Shared Access Account. It allows read access to the identity registry.
        :param pulumi.Input[bool] registry_write: Adds `RegistryWrite` permission to this Shared Access Account. It allows write access to the identity registry.
        :param pulumi.Input[str] resource_group_name: The name of the resource group under which the IotHub Shared Access Policy resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secondary_connection_string: The secondary connection string of the Shared Access Policy.
        :param pulumi.Input[str] secondary_key: The secondary key used to create the authentication token.
        :param pulumi.Input[bool] service_connect: Adds `ServiceConnect` permission to this Shared Access Account. It allows sending and receiving on the cloud-side endpoints.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["device_connect"] = device_connect
        __props__["iothub_name"] = iothub_name
        __props__["name"] = name
        __props__["primary_connection_string"] = primary_connection_string
        __props__["primary_key"] = primary_key
        __props__["registry_read"] = registry_read
        __props__["registry_write"] = registry_write
        __props__["resource_group_name"] = resource_group_name
        __props__["secondary_connection_string"] = secondary_connection_string
        __props__["secondary_key"] = secondary_key
        __props__["service_connect"] = service_connect
        return SharedAccessPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deviceConnect")
    def device_connect(self) -> pulumi.Output[Optional[bool]]:
        """
        Adds `DeviceConnect` permission to this Shared Access Account. It allows sending and receiving on the device-side endpoints.
        """
        return pulumi.get(self, "device_connect")

    @property
    @pulumi.getter(name="iothubName")
    def iothub_name(self) -> pulumi.Output[str]:
        """
        The name of the IoTHub to which this Shared Access Policy belongs. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "iothub_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the IotHub Shared Access Policy resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="primaryConnectionString")
    def primary_connection_string(self) -> pulumi.Output[str]:
        """
        The primary connection string of the Shared Access Policy.
        """
        return pulumi.get(self, "primary_connection_string")

    @property
    @pulumi.getter(name="primaryKey")
    def primary_key(self) -> pulumi.Output[str]:
        """
        The primary key used to create the authentication token.
        """
        return pulumi.get(self, "primary_key")

    @property
    @pulumi.getter(name="registryRead")
    def registry_read(self) -> pulumi.Output[Optional[bool]]:
        """
        Adds `RegistryRead` permission to this Shared Access Account. It allows read access to the identity registry.
        """
        return pulumi.get(self, "registry_read")

    @property
    @pulumi.getter(name="registryWrite")
    def registry_write(self) -> pulumi.Output[Optional[bool]]:
        """
        Adds `RegistryWrite` permission to this Shared Access Account. It allows write access to the identity registry.
        """
        return pulumi.get(self, "registry_write")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group under which the IotHub Shared Access Policy resource has to be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="secondaryConnectionString")
    def secondary_connection_string(self) -> pulumi.Output[str]:
        """
        The secondary connection string of the Shared Access Policy.
        """
        return pulumi.get(self, "secondary_connection_string")

    @property
    @pulumi.getter(name="secondaryKey")
    def secondary_key(self) -> pulumi.Output[str]:
        """
        The secondary key used to create the authentication token.
        """
        return pulumi.get(self, "secondary_key")

    @property
    @pulumi.getter(name="serviceConnect")
    def service_connect(self) -> pulumi.Output[Optional[bool]]:
        """
        Adds `ServiceConnect` permission to this Shared Access Account. It allows sending and receiving on the cloud-side endpoints.
        """
        return pulumi.get(self, "service_connect")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


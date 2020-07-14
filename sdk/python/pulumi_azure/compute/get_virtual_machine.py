# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class GetVirtualMachineResult:
    """
    A collection of values returned by getVirtualMachine.
    """
    def __init__(__self__, id=None, identities=None, location=None, name=None, resource_group_name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if identities and not isinstance(identities, list):
            raise TypeError("Expected argument 'identities' to be a list")
        __self__.identities = identities
        """
        A `identity` block as defined below.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name


class AwaitableGetVirtualMachineResult(GetVirtualMachineResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVirtualMachineResult(
            id=self.id,
            identities=self.identities,
            location=self.location,
            name=self.name,
            resource_group_name=self.resource_group_name)


def get_virtual_machine(name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing Virtual Machine.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.compute.get_virtual_machine(name="production",
        resource_group_name="networking")
    pulumi.export("virtualMachineId", example.id)
    ```


    :param str name: Specifies the name of the Virtual Machine.
    :param str resource_group_name: Specifies the name of the resource group the Virtual Machine is located in.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:compute/getVirtualMachine:getVirtualMachine', __args__, opts=opts).value

    return AwaitableGetVirtualMachineResult(
        id=__ret__.get('id'),
        identities=__ret__.get('identities'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        resource_group_name=__ret__.get('resourceGroupName'))

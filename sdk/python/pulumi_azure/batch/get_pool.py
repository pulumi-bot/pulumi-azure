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

__all__ = [
    'GetPoolResult',
    'AwaitableGetPoolResult',
    'get_pool',
]


class GetPoolResult:
    """
    A collection of values returned by getPool.
    """
    def __init__(__self__, account_name=None, auto_scales=None, certificates=None, container_configurations=None, display_name=None, fixed_scales=None, id=None, max_tasks_per_node=None, metadata=None, name=None, network_configuration=None, node_agent_sku_id=None, resource_group_name=None, start_task=None, storage_image_references=None, vm_size=None):
        if account_name and not isinstance(account_name, str):
            raise TypeError("Expected argument 'account_name' to be a str")
        __self__.account_name = account_name
        """
        The name of the Batch account.
        """
        if auto_scales and not isinstance(auto_scales, list):
            raise TypeError("Expected argument 'auto_scales' to be a list")
        __self__.auto_scales = auto_scales
        """
        A `auto_scale` block that describes the scale settings when using auto scale.
        """
        if certificates and not isinstance(certificates, list):
            raise TypeError("Expected argument 'certificates' to be a list")
        __self__.certificates = certificates
        """
        One or more `certificate` blocks that describe the certificates installed on each compute node in the pool.
        """
        if container_configurations and not isinstance(container_configurations, list):
            raise TypeError("Expected argument 'container_configurations' to be a list")
        __self__.container_configurations = container_configurations
        """
        The container configuration used in the pool's VMs.
        """
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        __self__.display_name = display_name
        if fixed_scales and not isinstance(fixed_scales, list):
            raise TypeError("Expected argument 'fixed_scales' to be a list")
        __self__.fixed_scales = fixed_scales
        """
        A `fixed_scale` block that describes the scale settings when using fixed scale.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if max_tasks_per_node and not isinstance(max_tasks_per_node, float):
            raise TypeError("Expected argument 'max_tasks_per_node' to be a float")
        __self__.max_tasks_per_node = max_tasks_per_node
        """
        The maximum number of tasks that can run concurrently on a single compute node in the pool.
        """
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        __self__.metadata = metadata
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the endpoint.
        """
        if network_configuration and not isinstance(network_configuration, dict):
            raise TypeError("Expected argument 'network_configuration' to be a dict")
        __self__.network_configuration = network_configuration
        if node_agent_sku_id and not isinstance(node_agent_sku_id, str):
            raise TypeError("Expected argument 'node_agent_sku_id' to be a str")
        __self__.node_agent_sku_id = node_agent_sku_id
        """
        The Sku of the node agents in the Batch pool.
        """
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if start_task and not isinstance(start_task, dict):
            raise TypeError("Expected argument 'start_task' to be a dict")
        __self__.start_task = start_task
        """
        A `start_task` block that describes the start task settings for the Batch pool.
        """
        if storage_image_references and not isinstance(storage_image_references, list):
            raise TypeError("Expected argument 'storage_image_references' to be a list")
        __self__.storage_image_references = storage_image_references
        """
        The reference of the storage image used by the nodes in the Batch pool.
        """
        if vm_size and not isinstance(vm_size, str):
            raise TypeError("Expected argument 'vm_size' to be a str")
        __self__.vm_size = vm_size
        """
        The size of the VM created in the Batch pool.
        """


class AwaitableGetPoolResult(GetPoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPoolResult(
            account_name=self.account_name,
            auto_scales=self.auto_scales,
            certificates=self.certificates,
            container_configurations=self.container_configurations,
            display_name=self.display_name,
            fixed_scales=self.fixed_scales,
            id=self.id,
            max_tasks_per_node=self.max_tasks_per_node,
            metadata=self.metadata,
            name=self.name,
            network_configuration=self.network_configuration,
            node_agent_sku_id=self.node_agent_sku_id,
            resource_group_name=self.resource_group_name,
            start_task=self.start_task,
            storage_image_references=self.storage_image_references,
            vm_size=self.vm_size)


def get_pool(account_name: Optional[str] = None,
             certificates: Optional[List[pulumi.InputType['GetPoolCertificateArgs']]] = None,
             name: Optional[str] = None,
             network_configuration: Optional[pulumi.InputType['GetPoolNetworkConfigurationArgs']] = None,
             resource_group_name: Optional[str] = None,
             start_task: Optional[pulumi.InputType['GetPoolStartTaskArgs']] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPoolResult:
    """
    Use this data source to access information about an existing Batch pool

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.batch.get_pool(account_name="testbatchaccount",
        name="testbatchpool",
        resource_group_name="test")
    ```


    :param str account_name: The name of the Batch account.
    :param List[pulumi.InputType['GetPoolCertificateArgs']] certificates: One or more `certificate` blocks that describe the certificates installed on each compute node in the pool.
    :param str name: The name of the endpoint.
    :param pulumi.InputType['GetPoolStartTaskArgs'] start_task: A `start_task` block that describes the start task settings for the Batch pool.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['certificates'] = certificates
    __args__['name'] = name
    __args__['networkConfiguration'] = network_configuration
    __args__['resourceGroupName'] = resource_group_name
    __args__['startTask'] = start_task
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:batch/getPool:getPool', __args__, opts=opts).value

    return AwaitableGetPoolResult(
        account_name=__ret__.get('accountName'),
        auto_scales=__ret__.get('autoScales'),
        certificates=__ret__.get('certificates'),
        container_configurations=__ret__.get('containerConfigurations'),
        display_name=__ret__.get('displayName'),
        fixed_scales=__ret__.get('fixedScales'),
        id=__ret__.get('id'),
        max_tasks_per_node=__ret__.get('maxTasksPerNode'),
        metadata=__ret__.get('metadata'),
        name=__ret__.get('name'),
        network_configuration=__ret__.get('networkConfiguration'),
        node_agent_sku_id=__ret__.get('nodeAgentSkuId'),
        resource_group_name=__ret__.get('resourceGroupName'),
        start_task=__ret__.get('startTask'),
        storage_image_references=__ret__.get('storageImageReferences'),
        vm_size=__ret__.get('vmSize'))

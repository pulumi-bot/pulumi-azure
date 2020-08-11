# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'ClusterIdentityArgs',
    'ClusterOptimizedAutoScaleArgs',
    'ClusterSkuArgs',
    'ClusterVirtualNetworkConfigurationArgs',
]

@pulumi.input_type
class ClusterIdentityArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 identity_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 principal_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: Specifies the type of Managed Service Identity that is configured on this Kusto Cluster. Possible values are: `SystemAssigned` (where Azure will generate a Service Principal for you).
        :param pulumi.Input[List[pulumi.Input[str]]] identity_ids: The list of user identities associated with the Kusto cluster.
        :param pulumi.Input[str] principal_id: Specifies the Principal ID of the System Assigned Managed Service Identity that is configured on this Kusto Cluster.
        :param pulumi.Input[str] tenant_id: Specifies the Tenant ID of the System Assigned Managed Service Identity that is configured on this Kusto Cluster.
        """
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "identityIds", identity_ids)
        pulumi.set(__self__, "principalId", principal_id)
        pulumi.set(__self__, "tenantId", tenant_id)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Specifies the type of Managed Service Identity that is configured on this Kusto Cluster. Possible values are: `SystemAssigned` (where Azure will generate a Service Principal for you).
        """
        ...

    @type.setter
    def type(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="identityIds")
    def identity_ids(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        The list of user identities associated with the Kusto cluster.
        """
        ...

    @identity_ids.setter
    def identity_ids(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        ...

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the Principal ID of the System Assigned Managed Service Identity that is configured on this Kusto Cluster.
        """
        ...

    @principal_id.setter
    def principal_id(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the Tenant ID of the System Assigned Managed Service Identity that is configured on this Kusto Cluster.
        """
        ...

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class ClusterOptimizedAutoScaleArgs:
    def __init__(__self__, *,
                 maximum_instances: pulumi.Input[float],
                 minimum_instances: pulumi.Input[float]):
        """
        :param pulumi.Input[float] maximum_instances: The maximum number of allowed instances. Must between `0` and `1000`.
        :param pulumi.Input[float] minimum_instances: The minimum number of allowed instances. Must between `0` and `1000`.
        """
        pulumi.set(__self__, "maximumInstances", maximum_instances)
        pulumi.set(__self__, "minimumInstances", minimum_instances)

    @property
    @pulumi.getter(name="maximumInstances")
    def maximum_instances(self) -> pulumi.Input[float]:
        """
        The maximum number of allowed instances. Must between `0` and `1000`.
        """
        ...

    @maximum_instances.setter
    def maximum_instances(self, value: pulumi.Input[float]):
        ...

    @property
    @pulumi.getter(name="minimumInstances")
    def minimum_instances(self) -> pulumi.Input[float]:
        """
        The minimum number of allowed instances. Must between `0` and `1000`.
        """
        ...

    @minimum_instances.setter
    def minimum_instances(self, value: pulumi.Input[float]):
        ...


@pulumi.input_type
class ClusterSkuArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 capacity: Optional[pulumi.Input[float]] = None):
        """
        :param pulumi.Input[str] name: The name of the SKU. Valid values are: `Dev(No SLA)_Standard_D11_v2`, `Dev(No SLA)_Standard_E2a_v4`, `Standard_D11_v2`, `Standard_D12_v2`, `Standard_D13_v2`, `Standard_D14_v2`, `Standard_DS13_v2+1TB_PS`, `Standard_DS13_v2+2TB_PS`, `Standard_DS14_v2+3TB_PS`, `Standard_DS14_v2+4TB_PS`, `Standard_E16as_v4+3TB_PS`, `Standard_E16as_v4+4TB_PS`, `Standard_E16a_v4`, `Standard_E2a_v4`, `Standard_E4a_v4`, `Standard_E8as_v4+1TB_PS`, `Standard_E8as_v4+2TB_PS`, `Standard_E8a_v4`, `Standard_L16s`, `Standard_L4s` and `Standard_L8s`
        :param pulumi.Input[float] capacity: Specifies the node count for the cluster. Boundaries depend on the sku name.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "capacity", capacity)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the SKU. Valid values are: `Dev(No SLA)_Standard_D11_v2`, `Dev(No SLA)_Standard_E2a_v4`, `Standard_D11_v2`, `Standard_D12_v2`, `Standard_D13_v2`, `Standard_D14_v2`, `Standard_DS13_v2+1TB_PS`, `Standard_DS13_v2+2TB_PS`, `Standard_DS14_v2+3TB_PS`, `Standard_DS14_v2+4TB_PS`, `Standard_E16as_v4+3TB_PS`, `Standard_E16as_v4+4TB_PS`, `Standard_E16a_v4`, `Standard_E2a_v4`, `Standard_E4a_v4`, `Standard_E8as_v4+1TB_PS`, `Standard_E8as_v4+2TB_PS`, `Standard_E8a_v4`, `Standard_L16s`, `Standard_L4s` and `Standard_L8s`
        """
        ...

    @name.setter
    def name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def capacity(self) -> Optional[pulumi.Input[float]]:
        """
        Specifies the node count for the cluster. Boundaries depend on the sku name.
        """
        ...

    @capacity.setter
    def capacity(self, value: Optional[pulumi.Input[float]]):
        ...


@pulumi.input_type
class ClusterVirtualNetworkConfigurationArgs:
    def __init__(__self__, *,
                 data_management_public_ip_id: pulumi.Input[str],
                 engine_public_ip_id: pulumi.Input[str],
                 subnet_id: pulumi.Input[str]):
        """
        :param pulumi.Input[str] data_management_public_ip_id: Data management's service public IP address resource id.
        :param pulumi.Input[str] engine_public_ip_id: Engine service's public IP address resource id.
        :param pulumi.Input[str] subnet_id: The subnet resource id.
        """
        pulumi.set(__self__, "dataManagementPublicIpId", data_management_public_ip_id)
        pulumi.set(__self__, "enginePublicIpId", engine_public_ip_id)
        pulumi.set(__self__, "subnetId", subnet_id)

    @property
    @pulumi.getter(name="dataManagementPublicIpId")
    def data_management_public_ip_id(self) -> pulumi.Input[str]:
        """
        Data management's service public IP address resource id.
        """
        ...

    @data_management_public_ip_id.setter
    def data_management_public_ip_id(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="enginePublicIpId")
    def engine_public_ip_id(self) -> pulumi.Input[str]:
        """
        Engine service's public IP address resource id.
        """
        ...

    @engine_public_ip_id.setter
    def engine_public_ip_id(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        The subnet resource id.
        """
        ...

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        ...



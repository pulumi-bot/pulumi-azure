# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'ClusterIdentity',
    'ClusterOptimizedAutoScale',
    'ClusterSku',
    'ClusterVirtualNetworkConfiguration',
]

@pulumi.output_type
class ClusterIdentity(dict):
    def __init__(__self__, *,
                 type: str,
                 identity_ids: Optional[Sequence[str]] = None,
                 principal_id: Optional[str] = None,
                 tenant_id: Optional[str] = None):
        """
        :param str type: Specifies the type of Managed Service Identity that is configured on this Kusto Cluster. Possible values are: `SystemAssigned` (where Azure will generate a Service Principal for you).
        :param Sequence[str] identity_ids: The list of user identities associated with the Kusto cluster.
        :param str principal_id: Specifies the Principal ID of the System Assigned Managed Service Identity that is configured on this Kusto Cluster.
        :param str tenant_id: Specifies the Tenant ID of the System Assigned Managed Service Identity that is configured on this Kusto Cluster.
        """
        pulumi.set(__self__, "type", type)
        if identity_ids is not None:
            pulumi.set(__self__, "identity_ids", identity_ids)
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Specifies the type of Managed Service Identity that is configured on this Kusto Cluster. Possible values are: `SystemAssigned` (where Azure will generate a Service Principal for you).
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="identityIds")
    def identity_ids(self) -> Optional[Sequence[str]]:
        """
        The list of user identities associated with the Kusto cluster.
        """
        return pulumi.get(self, "identity_ids")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[str]:
        """
        Specifies the Principal ID of the System Assigned Managed Service Identity that is configured on this Kusto Cluster.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        """
        Specifies the Tenant ID of the System Assigned Managed Service Identity that is configured on this Kusto Cluster.
        """
        return pulumi.get(self, "tenant_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterOptimizedAutoScale(dict):
    def __init__(__self__, *,
                 maximum_instances: float,
                 minimum_instances: float):
        """
        :param float maximum_instances: The maximum number of allowed instances. Must between `0` and `1000`.
        :param float minimum_instances: The minimum number of allowed instances. Must between `0` and `1000`.
        """
        pulumi.set(__self__, "maximum_instances", maximum_instances)
        pulumi.set(__self__, "minimum_instances", minimum_instances)

    @property
    @pulumi.getter(name="maximumInstances")
    def maximum_instances(self) -> float:
        """
        The maximum number of allowed instances. Must between `0` and `1000`.
        """
        return pulumi.get(self, "maximum_instances")

    @property
    @pulumi.getter(name="minimumInstances")
    def minimum_instances(self) -> float:
        """
        The minimum number of allowed instances. Must between `0` and `1000`.
        """
        return pulumi.get(self, "minimum_instances")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterSku(dict):
    def __init__(__self__, *,
                 name: str,
                 capacity: Optional[float] = None):
        """
        :param str name: The name of the SKU. Valid values are: `Dev(No SLA)_Standard_D11_v2`, `Dev(No SLA)_Standard_E2a_v4`, `Standard_D11_v2`, `Standard_D12_v2`, `Standard_D13_v2`, `Standard_D14_v2`, `Standard_DS13_v2+1TB_PS`, `Standard_DS13_v2+2TB_PS`, `Standard_DS14_v2+3TB_PS`, `Standard_DS14_v2+4TB_PS`, `Standard_E16as_v4+3TB_PS`, `Standard_E16as_v4+4TB_PS`, `Standard_E16a_v4`, `Standard_E2a_v4`, `Standard_E4a_v4`, `Standard_E8as_v4+1TB_PS`, `Standard_E8as_v4+2TB_PS`, `Standard_E8a_v4`, `Standard_L16s`, `Standard_L4s` and `Standard_L8s`
        :param float capacity: Specifies the node count for the cluster. Boundaries depend on the sku name.
        """
        pulumi.set(__self__, "name", name)
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the SKU. Valid values are: `Dev(No SLA)_Standard_D11_v2`, `Dev(No SLA)_Standard_E2a_v4`, `Standard_D11_v2`, `Standard_D12_v2`, `Standard_D13_v2`, `Standard_D14_v2`, `Standard_DS13_v2+1TB_PS`, `Standard_DS13_v2+2TB_PS`, `Standard_DS14_v2+3TB_PS`, `Standard_DS14_v2+4TB_PS`, `Standard_E16as_v4+3TB_PS`, `Standard_E16as_v4+4TB_PS`, `Standard_E16a_v4`, `Standard_E2a_v4`, `Standard_E4a_v4`, `Standard_E8as_v4+1TB_PS`, `Standard_E8as_v4+2TB_PS`, `Standard_E8a_v4`, `Standard_L16s`, `Standard_L4s` and `Standard_L8s`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def capacity(self) -> Optional[float]:
        """
        Specifies the node count for the cluster. Boundaries depend on the sku name.
        """
        return pulumi.get(self, "capacity")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterVirtualNetworkConfiguration(dict):
    def __init__(__self__, *,
                 data_management_public_ip_id: str,
                 engine_public_ip_id: str,
                 subnet_id: str):
        """
        :param str data_management_public_ip_id: Data management's service public IP address resource id.
        :param str engine_public_ip_id: Engine service's public IP address resource id.
        :param str subnet_id: The subnet resource id.
        """
        pulumi.set(__self__, "data_management_public_ip_id", data_management_public_ip_id)
        pulumi.set(__self__, "engine_public_ip_id", engine_public_ip_id)
        pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="dataManagementPublicIpId")
    def data_management_public_ip_id(self) -> str:
        """
        Data management's service public IP address resource id.
        """
        return pulumi.get(self, "data_management_public_ip_id")

    @property
    @pulumi.getter(name="enginePublicIpId")
    def engine_public_ip_id(self) -> str:
        """
        Engine service's public IP address resource id.
        """
        return pulumi.get(self, "engine_public_ip_id")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        The subnet resource id.
        """
        return pulumi.get(self, "subnet_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop



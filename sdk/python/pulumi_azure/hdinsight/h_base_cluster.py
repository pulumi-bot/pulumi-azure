# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['HBaseClusterArgs', 'HBaseCluster']

@pulumi.input_type
class HBaseClusterArgs:
    def __init__(__self__, *,
                 cluster_version: pulumi.Input[str],
                 component_version: pulumi.Input['HBaseClusterComponentVersionArgs'],
                 gateway: pulumi.Input['HBaseClusterGatewayArgs'],
                 resource_group_name: pulumi.Input[str],
                 roles: pulumi.Input['HBaseClusterRolesArgs'],
                 tier: pulumi.Input[str],
                 location: Optional[pulumi.Input[str]] = None,
                 metastores: Optional[pulumi.Input['HBaseClusterMetastoresArgs']] = None,
                 monitor: Optional[pulumi.Input['HBaseClusterMonitorArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 storage_account_gen2: Optional[pulumi.Input['HBaseClusterStorageAccountGen2Args']] = None,
                 storage_accounts: Optional[pulumi.Input[Sequence[pulumi.Input['HBaseClusterStorageAccountArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tls_min_version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a HBaseCluster resource.
        :param pulumi.Input[str] cluster_version: Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
        :param pulumi.Input['HBaseClusterComponentVersionArgs'] component_version: A `component_version` block as defined below.
        :param pulumi.Input['HBaseClusterGatewayArgs'] gateway: A `gateway` block as defined below.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group in which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
        :param pulumi.Input['HBaseClusterRolesArgs'] roles: A `roles` block as defined below.
        :param pulumi.Input[str] tier: Specifies the Tier which should be used for this HDInsight HBase Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the Azure Region which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
        :param pulumi.Input['HBaseClusterMetastoresArgs'] metastores: A `metastores` block as defined below.
        :param pulumi.Input['HBaseClusterMonitorArgs'] monitor: A `monitor` block as defined below.
        :param pulumi.Input[str] name: Specifies the name for this HDInsight HBase Cluster. Changing this forces a new resource to be created.
        :param pulumi.Input['HBaseClusterStorageAccountGen2Args'] storage_account_gen2: A `storage_account_gen2` block as defined below.
        :param pulumi.Input[Sequence[pulumi.Input['HBaseClusterStorageAccountArgs']]] storage_accounts: One or more `storage_account` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of Tags which should be assigned to this HDInsight HBase Cluster.
        """
        pulumi.set(__self__, "cluster_version", cluster_version)
        pulumi.set(__self__, "component_version", component_version)
        pulumi.set(__self__, "gateway", gateway)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "roles", roles)
        pulumi.set(__self__, "tier", tier)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if metastores is not None:
            pulumi.set(__self__, "metastores", metastores)
        if monitor is not None:
            pulumi.set(__self__, "monitor", monitor)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if storage_account_gen2 is not None:
            pulumi.set(__self__, "storage_account_gen2", storage_account_gen2)
        if storage_accounts is not None:
            pulumi.set(__self__, "storage_accounts", storage_accounts)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tls_min_version is not None:
            pulumi.set(__self__, "tls_min_version", tls_min_version)

    @property
    @pulumi.getter(name="clusterVersion")
    def cluster_version(self) -> pulumi.Input[str]:
        """
        Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "cluster_version")

    @cluster_version.setter
    def cluster_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_version", value)

    @property
    @pulumi.getter(name="componentVersion")
    def component_version(self) -> pulumi.Input['HBaseClusterComponentVersionArgs']:
        """
        A `component_version` block as defined below.
        """
        return pulumi.get(self, "component_version")

    @component_version.setter
    def component_version(self, value: pulumi.Input['HBaseClusterComponentVersionArgs']):
        pulumi.set(self, "component_version", value)

    @property
    @pulumi.getter
    def gateway(self) -> pulumi.Input['HBaseClusterGatewayArgs']:
        """
        A `gateway` block as defined below.
        """
        return pulumi.get(self, "gateway")

    @gateway.setter
    def gateway(self, value: pulumi.Input['HBaseClusterGatewayArgs']):
        pulumi.set(self, "gateway", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the Resource Group in which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def roles(self) -> pulumi.Input['HBaseClusterRolesArgs']:
        """
        A `roles` block as defined below.
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: pulumi.Input['HBaseClusterRolesArgs']):
        pulumi.set(self, "roles", value)

    @property
    @pulumi.getter
    def tier(self) -> pulumi.Input[str]:
        """
        Specifies the Tier which should be used for this HDInsight HBase Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "tier")

    @tier.setter
    def tier(self, value: pulumi.Input[str]):
        pulumi.set(self, "tier", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the Azure Region which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def metastores(self) -> Optional[pulumi.Input['HBaseClusterMetastoresArgs']]:
        """
        A `metastores` block as defined below.
        """
        return pulumi.get(self, "metastores")

    @metastores.setter
    def metastores(self, value: Optional[pulumi.Input['HBaseClusterMetastoresArgs']]):
        pulumi.set(self, "metastores", value)

    @property
    @pulumi.getter
    def monitor(self) -> Optional[pulumi.Input['HBaseClusterMonitorArgs']]:
        """
        A `monitor` block as defined below.
        """
        return pulumi.get(self, "monitor")

    @monitor.setter
    def monitor(self, value: Optional[pulumi.Input['HBaseClusterMonitorArgs']]):
        pulumi.set(self, "monitor", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name for this HDInsight HBase Cluster. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="storageAccountGen2")
    def storage_account_gen2(self) -> Optional[pulumi.Input['HBaseClusterStorageAccountGen2Args']]:
        """
        A `storage_account_gen2` block as defined below.
        """
        return pulumi.get(self, "storage_account_gen2")

    @storage_account_gen2.setter
    def storage_account_gen2(self, value: Optional[pulumi.Input['HBaseClusterStorageAccountGen2Args']]):
        pulumi.set(self, "storage_account_gen2", value)

    @property
    @pulumi.getter(name="storageAccounts")
    def storage_accounts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['HBaseClusterStorageAccountArgs']]]]:
        """
        One or more `storage_account` block as defined below.
        """
        return pulumi.get(self, "storage_accounts")

    @storage_accounts.setter
    def storage_accounts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['HBaseClusterStorageAccountArgs']]]]):
        pulumi.set(self, "storage_accounts", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of Tags which should be assigned to this HDInsight HBase Cluster.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tlsMinVersion")
    def tls_min_version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tls_min_version")

    @tls_min_version.setter
    def tls_min_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_min_version", value)


class HBaseCluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HBaseClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a HDInsight HBase Cluster.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS")
        example_container = azure.storage.Container("exampleContainer",
            storage_account_name=example_account.name,
            container_access_type="private")
        example_h_base_cluster = azure.hdinsight.HBaseCluster("exampleHBaseCluster",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            cluster_version="3.6",
            tier="Standard",
            component_version=azure.hdinsight.HBaseClusterComponentVersionArgs(
                hbase="1.1",
            ),
            gateway=azure.hdinsight.HBaseClusterGatewayArgs(
                enabled=True,
                username="acctestusrgw",
                password="Password123!",
            ),
            storage_accounts=[azure.hdinsight.HBaseClusterStorageAccountArgs(
                storage_container_id=example_container.id,
                storage_account_key=example_account.primary_access_key,
                is_default=True,
            )],
            roles=azure.hdinsight.HBaseClusterRolesArgs(
                head_node=azure.hdinsight.HBaseClusterRolesHeadNodeArgs(
                    vm_size="Standard_D3_V2",
                    username="acctestusrvm",
                    password="AccTestvdSC4daf986!",
                ),
                worker_node=azure.hdinsight.HBaseClusterRolesWorkerNodeArgs(
                    vm_size="Standard_D3_V2",
                    username="acctestusrvm",
                    password="AccTestvdSC4daf986!",
                    target_instance_count=3,
                ),
                zookeeper_node=azure.hdinsight.HBaseClusterRolesZookeeperNodeArgs(
                    vm_size="Standard_D3_V2",
                    username="acctestusrvm",
                    password="AccTestvdSC4daf986!",
                ),
            ))
        ```

        ## Import

        HDInsight HBase Clusters can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:hdinsight/hBaseCluster:HBaseCluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.HDInsight/clusters/cluster1
        ```

        :param str resource_name: The name of the resource.
        :param HBaseClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_version: Optional[pulumi.Input[str]] = None,
                 component_version: Optional[pulumi.Input[pulumi.InputType['HBaseClusterComponentVersionArgs']]] = None,
                 gateway: Optional[pulumi.Input[pulumi.InputType['HBaseClusterGatewayArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 metastores: Optional[pulumi.Input[pulumi.InputType['HBaseClusterMetastoresArgs']]] = None,
                 monitor: Optional[pulumi.Input[pulumi.InputType['HBaseClusterMonitorArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[pulumi.InputType['HBaseClusterRolesArgs']]] = None,
                 storage_account_gen2: Optional[pulumi.Input[pulumi.InputType['HBaseClusterStorageAccountGen2Args']]] = None,
                 storage_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['HBaseClusterStorageAccountArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tier: Optional[pulumi.Input[str]] = None,
                 tls_min_version: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a HDInsight HBase Cluster.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS")
        example_container = azure.storage.Container("exampleContainer",
            storage_account_name=example_account.name,
            container_access_type="private")
        example_h_base_cluster = azure.hdinsight.HBaseCluster("exampleHBaseCluster",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            cluster_version="3.6",
            tier="Standard",
            component_version=azure.hdinsight.HBaseClusterComponentVersionArgs(
                hbase="1.1",
            ),
            gateway=azure.hdinsight.HBaseClusterGatewayArgs(
                enabled=True,
                username="acctestusrgw",
                password="Password123!",
            ),
            storage_accounts=[azure.hdinsight.HBaseClusterStorageAccountArgs(
                storage_container_id=example_container.id,
                storage_account_key=example_account.primary_access_key,
                is_default=True,
            )],
            roles=azure.hdinsight.HBaseClusterRolesArgs(
                head_node=azure.hdinsight.HBaseClusterRolesHeadNodeArgs(
                    vm_size="Standard_D3_V2",
                    username="acctestusrvm",
                    password="AccTestvdSC4daf986!",
                ),
                worker_node=azure.hdinsight.HBaseClusterRolesWorkerNodeArgs(
                    vm_size="Standard_D3_V2",
                    username="acctestusrvm",
                    password="AccTestvdSC4daf986!",
                    target_instance_count=3,
                ),
                zookeeper_node=azure.hdinsight.HBaseClusterRolesZookeeperNodeArgs(
                    vm_size="Standard_D3_V2",
                    username="acctestusrvm",
                    password="AccTestvdSC4daf986!",
                ),
            ))
        ```

        ## Import

        HDInsight HBase Clusters can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:hdinsight/hBaseCluster:HBaseCluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.HDInsight/clusters/cluster1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_version: Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['HBaseClusterComponentVersionArgs']] component_version: A `component_version` block as defined below.
        :param pulumi.Input[pulumi.InputType['HBaseClusterGatewayArgs']] gateway: A `gateway` block as defined below.
        :param pulumi.Input[str] location: Specifies the Azure Region which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['HBaseClusterMetastoresArgs']] metastores: A `metastores` block as defined below.
        :param pulumi.Input[pulumi.InputType['HBaseClusterMonitorArgs']] monitor: A `monitor` block as defined below.
        :param pulumi.Input[str] name: Specifies the name for this HDInsight HBase Cluster. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group in which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['HBaseClusterRolesArgs']] roles: A `roles` block as defined below.
        :param pulumi.Input[pulumi.InputType['HBaseClusterStorageAccountGen2Args']] storage_account_gen2: A `storage_account_gen2` block as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['HBaseClusterStorageAccountArgs']]]] storage_accounts: One or more `storage_account` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of Tags which should be assigned to this HDInsight HBase Cluster.
        :param pulumi.Input[str] tier: Specifies the Tier which should be used for this HDInsight HBase Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HBaseClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_version: Optional[pulumi.Input[str]] = None,
                 component_version: Optional[pulumi.Input[pulumi.InputType['HBaseClusterComponentVersionArgs']]] = None,
                 gateway: Optional[pulumi.Input[pulumi.InputType['HBaseClusterGatewayArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 metastores: Optional[pulumi.Input[pulumi.InputType['HBaseClusterMetastoresArgs']]] = None,
                 monitor: Optional[pulumi.Input[pulumi.InputType['HBaseClusterMonitorArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[pulumi.InputType['HBaseClusterRolesArgs']]] = None,
                 storage_account_gen2: Optional[pulumi.Input[pulumi.InputType['HBaseClusterStorageAccountGen2Args']]] = None,
                 storage_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['HBaseClusterStorageAccountArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tier: Optional[pulumi.Input[str]] = None,
                 tls_min_version: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            if cluster_version is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_version'")
            __props__['cluster_version'] = cluster_version
            if component_version is None and not opts.urn:
                raise TypeError("Missing required property 'component_version'")
            __props__['component_version'] = component_version
            if gateway is None and not opts.urn:
                raise TypeError("Missing required property 'gateway'")
            __props__['gateway'] = gateway
            __props__['location'] = location
            __props__['metastores'] = metastores
            __props__['monitor'] = monitor
            __props__['name'] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if roles is None and not opts.urn:
                raise TypeError("Missing required property 'roles'")
            __props__['roles'] = roles
            __props__['storage_account_gen2'] = storage_account_gen2
            __props__['storage_accounts'] = storage_accounts
            __props__['tags'] = tags
            if tier is None and not opts.urn:
                raise TypeError("Missing required property 'tier'")
            __props__['tier'] = tier
            __props__['tls_min_version'] = tls_min_version
            __props__['https_endpoint'] = None
            __props__['ssh_endpoint'] = None
        super(HBaseCluster, __self__).__init__(
            'azure:hdinsight/hBaseCluster:HBaseCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_version: Optional[pulumi.Input[str]] = None,
            component_version: Optional[pulumi.Input[pulumi.InputType['HBaseClusterComponentVersionArgs']]] = None,
            gateway: Optional[pulumi.Input[pulumi.InputType['HBaseClusterGatewayArgs']]] = None,
            https_endpoint: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            metastores: Optional[pulumi.Input[pulumi.InputType['HBaseClusterMetastoresArgs']]] = None,
            monitor: Optional[pulumi.Input[pulumi.InputType['HBaseClusterMonitorArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            roles: Optional[pulumi.Input[pulumi.InputType['HBaseClusterRolesArgs']]] = None,
            ssh_endpoint: Optional[pulumi.Input[str]] = None,
            storage_account_gen2: Optional[pulumi.Input[pulumi.InputType['HBaseClusterStorageAccountGen2Args']]] = None,
            storage_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['HBaseClusterStorageAccountArgs']]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tier: Optional[pulumi.Input[str]] = None,
            tls_min_version: Optional[pulumi.Input[str]] = None) -> 'HBaseCluster':
        """
        Get an existing HBaseCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_version: Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['HBaseClusterComponentVersionArgs']] component_version: A `component_version` block as defined below.
        :param pulumi.Input[pulumi.InputType['HBaseClusterGatewayArgs']] gateway: A `gateway` block as defined below.
        :param pulumi.Input[str] https_endpoint: The HTTPS Connectivity Endpoint for this HDInsight HBase Cluster.
        :param pulumi.Input[str] location: Specifies the Azure Region which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['HBaseClusterMetastoresArgs']] metastores: A `metastores` block as defined below.
        :param pulumi.Input[pulumi.InputType['HBaseClusterMonitorArgs']] monitor: A `monitor` block as defined below.
        :param pulumi.Input[str] name: Specifies the name for this HDInsight HBase Cluster. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group in which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['HBaseClusterRolesArgs']] roles: A `roles` block as defined below.
        :param pulumi.Input[str] ssh_endpoint: The SSH Connectivity Endpoint for this HDInsight HBase Cluster.
        :param pulumi.Input[pulumi.InputType['HBaseClusterStorageAccountGen2Args']] storage_account_gen2: A `storage_account_gen2` block as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['HBaseClusterStorageAccountArgs']]]] storage_accounts: One or more `storage_account` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of Tags which should be assigned to this HDInsight HBase Cluster.
        :param pulumi.Input[str] tier: Specifies the Tier which should be used for this HDInsight HBase Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cluster_version"] = cluster_version
        __props__["component_version"] = component_version
        __props__["gateway"] = gateway
        __props__["https_endpoint"] = https_endpoint
        __props__["location"] = location
        __props__["metastores"] = metastores
        __props__["monitor"] = monitor
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["roles"] = roles
        __props__["ssh_endpoint"] = ssh_endpoint
        __props__["storage_account_gen2"] = storage_account_gen2
        __props__["storage_accounts"] = storage_accounts
        __props__["tags"] = tags
        __props__["tier"] = tier
        __props__["tls_min_version"] = tls_min_version
        return HBaseCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterVersion")
    def cluster_version(self) -> pulumi.Output[str]:
        """
        Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "cluster_version")

    @property
    @pulumi.getter(name="componentVersion")
    def component_version(self) -> pulumi.Output['outputs.HBaseClusterComponentVersion']:
        """
        A `component_version` block as defined below.
        """
        return pulumi.get(self, "component_version")

    @property
    @pulumi.getter
    def gateway(self) -> pulumi.Output['outputs.HBaseClusterGateway']:
        """
        A `gateway` block as defined below.
        """
        return pulumi.get(self, "gateway")

    @property
    @pulumi.getter(name="httpsEndpoint")
    def https_endpoint(self) -> pulumi.Output[str]:
        """
        The HTTPS Connectivity Endpoint for this HDInsight HBase Cluster.
        """
        return pulumi.get(self, "https_endpoint")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the Azure Region which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def metastores(self) -> pulumi.Output[Optional['outputs.HBaseClusterMetastores']]:
        """
        A `metastores` block as defined below.
        """
        return pulumi.get(self, "metastores")

    @property
    @pulumi.getter
    def monitor(self) -> pulumi.Output[Optional['outputs.HBaseClusterMonitor']]:
        """
        A `monitor` block as defined below.
        """
        return pulumi.get(self, "monitor")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name for this HDInsight HBase Cluster. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Resource Group in which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def roles(self) -> pulumi.Output['outputs.HBaseClusterRoles']:
        """
        A `roles` block as defined below.
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter(name="sshEndpoint")
    def ssh_endpoint(self) -> pulumi.Output[str]:
        """
        The SSH Connectivity Endpoint for this HDInsight HBase Cluster.
        """
        return pulumi.get(self, "ssh_endpoint")

    @property
    @pulumi.getter(name="storageAccountGen2")
    def storage_account_gen2(self) -> pulumi.Output[Optional['outputs.HBaseClusterStorageAccountGen2']]:
        """
        A `storage_account_gen2` block as defined below.
        """
        return pulumi.get(self, "storage_account_gen2")

    @property
    @pulumi.getter(name="storageAccounts")
    def storage_accounts(self) -> pulumi.Output[Optional[Sequence['outputs.HBaseClusterStorageAccount']]]:
        """
        One or more `storage_account` block as defined below.
        """
        return pulumi.get(self, "storage_accounts")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of Tags which should be assigned to this HDInsight HBase Cluster.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def tier(self) -> pulumi.Output[str]:
        """
        Specifies the Tier which should be used for this HDInsight HBase Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "tier")

    @property
    @pulumi.getter(name="tlsMinVersion")
    def tls_min_version(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "tls_min_version")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EventhubDataConnectionArgs', 'EventhubDataConnection']

@pulumi.input_type
class EventhubDataConnectionArgs:
    def __init__(__self__, *,
                 cluster_name: pulumi.Input[str],
                 consumer_group: pulumi.Input[str],
                 database_name: pulumi.Input[str],
                 eventhub_id: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 compression: Optional[pulumi.Input[str]] = None,
                 data_format: Optional[pulumi.Input[str]] = None,
                 event_system_properties: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 mapping_rule_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 table_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EventhubDataConnection resource.
        :param pulumi.Input[str] cluster_name: Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] consumer_group: Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
        :param pulumi.Input[str] database_name: Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] eventhub_id: Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] compression: Specifies compression type for the connection. Allowed values: `GZip` and `None`. Defaults to `None`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] data_format: Specifies the data format of the EventHub messages. Allowed values: `APACHEAVRO`, `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_system_properties: Specifies a list of system properties for the Event Hub.
        :param pulumi.Input[str] location: The location where the Kusto Database should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] mapping_rule_name: Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
        :param pulumi.Input[str] name: The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
        :param pulumi.Input[str] table_name: Specifies the target table name used for the message ingestion. Table must exist before resource is created.
        """
        pulumi.set(__self__, "cluster_name", cluster_name)
        pulumi.set(__self__, "consumer_group", consumer_group)
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "eventhub_id", eventhub_id)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if compression is not None:
            pulumi.set(__self__, "compression", compression)
        if data_format is not None:
            pulumi.set(__self__, "data_format", data_format)
        if event_system_properties is not None:
            pulumi.set(__self__, "event_system_properties", event_system_properties)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if mapping_rule_name is not None:
            pulumi.set(__self__, "mapping_rule_name", mapping_rule_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if table_name is not None:
            pulumi.set(__self__, "table_name", table_name)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="consumerGroup")
    def consumer_group(self) -> pulumi.Input[str]:
        """
        Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "consumer_group")

    @consumer_group.setter
    def consumer_group(self, value: pulumi.Input[str]):
        pulumi.set(self, "consumer_group", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="eventhubId")
    def eventhub_id(self) -> pulumi.Input[str]:
        """
        Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "eventhub_id")

    @eventhub_id.setter
    def eventhub_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "eventhub_id", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def compression(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies compression type for the connection. Allowed values: `GZip` and `None`. Defaults to `None`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "compression")

    @compression.setter
    def compression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compression", value)

    @property
    @pulumi.getter(name="dataFormat")
    def data_format(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the data format of the EventHub messages. Allowed values: `APACHEAVRO`, `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
        """
        return pulumi.get(self, "data_format")

    @data_format.setter
    def data_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_format", value)

    @property
    @pulumi.getter(name="eventSystemProperties")
    def event_system_properties(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies a list of system properties for the Event Hub.
        """
        return pulumi.get(self, "event_system_properties")

    @event_system_properties.setter
    def event_system_properties(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "event_system_properties", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location where the Kusto Database should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="mappingRuleName")
    def mapping_rule_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
        """
        return pulumi.get(self, "mapping_rule_name")

    @mapping_rule_name.setter
    def mapping_rule_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mapping_rule_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the target table name used for the message ingestion. Table must exist before resource is created.
        """
        return pulumi.get(self, "table_name")

    @table_name.setter
    def table_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "table_name", value)


@pulumi.input_type
class _EventhubDataConnectionState:
    def __init__(__self__, *,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 compression: Optional[pulumi.Input[str]] = None,
                 consumer_group: Optional[pulumi.Input[str]] = None,
                 data_format: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 event_system_properties: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 eventhub_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 mapping_rule_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 table_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EventhubDataConnection resources.
        :param pulumi.Input[str] cluster_name: Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] compression: Specifies compression type for the connection. Allowed values: `GZip` and `None`. Defaults to `None`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] consumer_group: Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
        :param pulumi.Input[str] data_format: Specifies the data format of the EventHub messages. Allowed values: `APACHEAVRO`, `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
        :param pulumi.Input[str] database_name: Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_system_properties: Specifies a list of system properties for the Event Hub.
        :param pulumi.Input[str] eventhub_id: Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: The location where the Kusto Database should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] mapping_rule_name: Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
        :param pulumi.Input[str] name: The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] table_name: Specifies the target table name used for the message ingestion. Table must exist before resource is created.
        """
        if cluster_name is not None:
            pulumi.set(__self__, "cluster_name", cluster_name)
        if compression is not None:
            pulumi.set(__self__, "compression", compression)
        if consumer_group is not None:
            pulumi.set(__self__, "consumer_group", consumer_group)
        if data_format is not None:
            pulumi.set(__self__, "data_format", data_format)
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if event_system_properties is not None:
            pulumi.set(__self__, "event_system_properties", event_system_properties)
        if eventhub_id is not None:
            pulumi.set(__self__, "eventhub_id", eventhub_id)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if mapping_rule_name is not None:
            pulumi.set(__self__, "mapping_rule_name", mapping_rule_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if table_name is not None:
            pulumi.set(__self__, "table_name", table_name)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter
    def compression(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies compression type for the connection. Allowed values: `GZip` and `None`. Defaults to `None`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "compression")

    @compression.setter
    def compression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compression", value)

    @property
    @pulumi.getter(name="consumerGroup")
    def consumer_group(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "consumer_group")

    @consumer_group.setter
    def consumer_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "consumer_group", value)

    @property
    @pulumi.getter(name="dataFormat")
    def data_format(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the data format of the EventHub messages. Allowed values: `APACHEAVRO`, `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
        """
        return pulumi.get(self, "data_format")

    @data_format.setter
    def data_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_format", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="eventSystemProperties")
    def event_system_properties(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies a list of system properties for the Event Hub.
        """
        return pulumi.get(self, "event_system_properties")

    @event_system_properties.setter
    def event_system_properties(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "event_system_properties", value)

    @property
    @pulumi.getter(name="eventhubId")
    def eventhub_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "eventhub_id")

    @eventhub_id.setter
    def eventhub_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eventhub_id", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location where the Kusto Database should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="mappingRuleName")
    def mapping_rule_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
        """
        return pulumi.get(self, "mapping_rule_name")

    @mapping_rule_name.setter
    def mapping_rule_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mapping_rule_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the target table name used for the message ingestion. Table must exist before resource is created.
        """
        return pulumi.get(self, "table_name")

    @table_name.setter
    def table_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "table_name", value)


class EventhubDataConnection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 compression: Optional[pulumi.Input[str]] = None,
                 consumer_group: Optional[pulumi.Input[str]] = None,
                 data_format: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 event_system_properties: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 eventhub_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 mapping_rule_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 table_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Kusto (also known as Azure Data Explorer) EventHub Data Connection

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        rg = azure.core.ResourceGroup("rg", location="West Europe")
        cluster = azure.kusto.Cluster("cluster",
            location=rg.location,
            resource_group_name=rg.name,
            sku=azure.kusto.ClusterSkuArgs(
                name="Standard_D13_v2",
                capacity=2,
            ))
        database = azure.kusto.Database("database",
            resource_group_name=rg.name,
            location=rg.location,
            cluster_name=cluster.name,
            hot_cache_period="P7D",
            soft_delete_period="P31D")
        eventhub_ns = azure.eventhub.EventHubNamespace("eventhubNs",
            location=rg.location,
            resource_group_name=rg.name,
            sku="Standard")
        eventhub = azure.eventhub.EventHub("eventhub",
            namespace_name=eventhub_ns.name,
            resource_group_name=rg.name,
            partition_count=1,
            message_retention=1)
        consumer_group = azure.eventhub.ConsumerGroup("consumerGroup",
            namespace_name=eventhub_ns.name,
            eventhub_name=eventhub.name,
            resource_group_name=rg.name)
        eventhub_connection = azure.kusto.EventhubDataConnection("eventhubConnection",
            resource_group_name=rg.name,
            location=rg.location,
            cluster_name=cluster.name,
            database_name=database.name,
            eventhub_id=eventhub.id,
            consumer_group=consumer_group.name,
            table_name="my-table",
            mapping_rule_name="my-table-mapping",
            data_format="JSON")
        #(Optional)
        ```

        ## Import

        Kusto EventHub Data Connections can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:kusto/eventhubDataConnection:EventhubDataConnection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/Databases/database1/DataConnections/eventHubConnection1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_name: Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] compression: Specifies compression type for the connection. Allowed values: `GZip` and `None`. Defaults to `None`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] consumer_group: Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
        :param pulumi.Input[str] data_format: Specifies the data format of the EventHub messages. Allowed values: `APACHEAVRO`, `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
        :param pulumi.Input[str] database_name: Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_system_properties: Specifies a list of system properties for the Event Hub.
        :param pulumi.Input[str] eventhub_id: Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: The location where the Kusto Database should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] mapping_rule_name: Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
        :param pulumi.Input[str] name: The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] table_name: Specifies the target table name used for the message ingestion. Table must exist before resource is created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EventhubDataConnectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Kusto (also known as Azure Data Explorer) EventHub Data Connection

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        rg = azure.core.ResourceGroup("rg", location="West Europe")
        cluster = azure.kusto.Cluster("cluster",
            location=rg.location,
            resource_group_name=rg.name,
            sku=azure.kusto.ClusterSkuArgs(
                name="Standard_D13_v2",
                capacity=2,
            ))
        database = azure.kusto.Database("database",
            resource_group_name=rg.name,
            location=rg.location,
            cluster_name=cluster.name,
            hot_cache_period="P7D",
            soft_delete_period="P31D")
        eventhub_ns = azure.eventhub.EventHubNamespace("eventhubNs",
            location=rg.location,
            resource_group_name=rg.name,
            sku="Standard")
        eventhub = azure.eventhub.EventHub("eventhub",
            namespace_name=eventhub_ns.name,
            resource_group_name=rg.name,
            partition_count=1,
            message_retention=1)
        consumer_group = azure.eventhub.ConsumerGroup("consumerGroup",
            namespace_name=eventhub_ns.name,
            eventhub_name=eventhub.name,
            resource_group_name=rg.name)
        eventhub_connection = azure.kusto.EventhubDataConnection("eventhubConnection",
            resource_group_name=rg.name,
            location=rg.location,
            cluster_name=cluster.name,
            database_name=database.name,
            eventhub_id=eventhub.id,
            consumer_group=consumer_group.name,
            table_name="my-table",
            mapping_rule_name="my-table-mapping",
            data_format="JSON")
        #(Optional)
        ```

        ## Import

        Kusto EventHub Data Connections can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:kusto/eventhubDataConnection:EventhubDataConnection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/Databases/database1/DataConnections/eventHubConnection1
        ```

        :param str resource_name: The name of the resource.
        :param EventhubDataConnectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EventhubDataConnectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 compression: Optional[pulumi.Input[str]] = None,
                 consumer_group: Optional[pulumi.Input[str]] = None,
                 data_format: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 event_system_properties: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 eventhub_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 mapping_rule_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 table_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EventhubDataConnectionArgs.__new__(EventhubDataConnectionArgs)

            if cluster_name is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_name'")
            __props__.__dict__["cluster_name"] = cluster_name
            __props__.__dict__["compression"] = compression
            if consumer_group is None and not opts.urn:
                raise TypeError("Missing required property 'consumer_group'")
            __props__.__dict__["consumer_group"] = consumer_group
            __props__.__dict__["data_format"] = data_format
            if database_name is None and not opts.urn:
                raise TypeError("Missing required property 'database_name'")
            __props__.__dict__["database_name"] = database_name
            __props__.__dict__["event_system_properties"] = event_system_properties
            if eventhub_id is None and not opts.urn:
                raise TypeError("Missing required property 'eventhub_id'")
            __props__.__dict__["eventhub_id"] = eventhub_id
            __props__.__dict__["location"] = location
            __props__.__dict__["mapping_rule_name"] = mapping_rule_name
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["table_name"] = table_name
        super(EventhubDataConnection, __self__).__init__(
            'azure:kusto/eventhubDataConnection:EventhubDataConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_name: Optional[pulumi.Input[str]] = None,
            compression: Optional[pulumi.Input[str]] = None,
            consumer_group: Optional[pulumi.Input[str]] = None,
            data_format: Optional[pulumi.Input[str]] = None,
            database_name: Optional[pulumi.Input[str]] = None,
            event_system_properties: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            eventhub_id: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            mapping_rule_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            table_name: Optional[pulumi.Input[str]] = None) -> 'EventhubDataConnection':
        """
        Get an existing EventhubDataConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_name: Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] compression: Specifies compression type for the connection. Allowed values: `GZip` and `None`. Defaults to `None`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] consumer_group: Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
        :param pulumi.Input[str] data_format: Specifies the data format of the EventHub messages. Allowed values: `APACHEAVRO`, `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
        :param pulumi.Input[str] database_name: Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_system_properties: Specifies a list of system properties for the Event Hub.
        :param pulumi.Input[str] eventhub_id: Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: The location where the Kusto Database should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] mapping_rule_name: Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
        :param pulumi.Input[str] name: The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] table_name: Specifies the target table name used for the message ingestion. Table must exist before resource is created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EventhubDataConnectionState.__new__(_EventhubDataConnectionState)

        __props__.__dict__["cluster_name"] = cluster_name
        __props__.__dict__["compression"] = compression
        __props__.__dict__["consumer_group"] = consumer_group
        __props__.__dict__["data_format"] = data_format
        __props__.__dict__["database_name"] = database_name
        __props__.__dict__["event_system_properties"] = event_system_properties
        __props__.__dict__["eventhub_id"] = eventhub_id
        __props__.__dict__["location"] = location
        __props__.__dict__["mapping_rule_name"] = mapping_rule_name
        __props__.__dict__["name"] = name
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["table_name"] = table_name
        return EventhubDataConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter
    def compression(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies compression type for the connection. Allowed values: `GZip` and `None`. Defaults to `None`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "compression")

    @property
    @pulumi.getter(name="consumerGroup")
    def consumer_group(self) -> pulumi.Output[str]:
        """
        Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "consumer_group")

    @property
    @pulumi.getter(name="dataFormat")
    def data_format(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the data format of the EventHub messages. Allowed values: `APACHEAVRO`, `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
        """
        return pulumi.get(self, "data_format")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="eventSystemProperties")
    def event_system_properties(self) -> pulumi.Output[Sequence[str]]:
        """
        Specifies a list of system properties for the Event Hub.
        """
        return pulumi.get(self, "event_system_properties")

    @property
    @pulumi.getter(name="eventhubId")
    def eventhub_id(self) -> pulumi.Output[str]:
        """
        Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "eventhub_id")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location where the Kusto Database should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="mappingRuleName")
    def mapping_rule_name(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
        """
        return pulumi.get(self, "mapping_rule_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the target table name used for the message ingestion. Table must exist before resource is created.
        """
        return pulumi.get(self, "table_name")


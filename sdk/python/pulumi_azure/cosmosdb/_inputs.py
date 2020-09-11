# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'AccountCapabilityArgs',
    'AccountConsistencyPolicyArgs',
    'AccountGeoLocationArgs',
    'AccountVirtualNetworkRuleArgs',
    'GremlinGraphConflictResolutionPolicyArgs',
    'GremlinGraphIndexPolicyArgs',
    'GremlinGraphUniqueKeyArgs',
    'MongoCollectionIndexArgs',
    'MongoCollectionSystemIndexArgs',
    'SqlContainerUniqueKeyArgs',
]

@pulumi.input_type
class AccountCapabilityArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str]):
        """
        :param pulumi.Input[str] name: Specifies the name of the CosmosDB Account. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the CosmosDB Account. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class AccountConsistencyPolicyArgs:
    def __init__(__self__, *,
                 consistency_level: pulumi.Input[str],
                 max_interval_in_seconds: Optional[pulumi.Input[int]] = None,
                 max_staleness_prefix: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] consistency_level: The Consistency Level to use for this CosmosDB Account - can be either `BoundedStaleness`, `Eventual`, `Session`, `Strong` or `ConsistentPrefix`.
        :param pulumi.Input[int] max_interval_in_seconds: When used with the Bounded Staleness consistency level, this value represents the time amount of staleness (in seconds) tolerated. Accepted range for this value is `5` - `86400` (1 day). Defaults to `5`. Required when `consistency_level` is set to `BoundedStaleness`.
        :param pulumi.Input[int] max_staleness_prefix: When used with the Bounded Staleness consistency level, this value represents the number of stale requests tolerated. Accepted range for this value is `10` – `2147483647`. Defaults to `100`. Required when `consistency_level` is set to `BoundedStaleness`.
        """
        pulumi.set(__self__, "consistency_level", consistency_level)
        if max_interval_in_seconds is not None:
            pulumi.set(__self__, "max_interval_in_seconds", max_interval_in_seconds)
        if max_staleness_prefix is not None:
            pulumi.set(__self__, "max_staleness_prefix", max_staleness_prefix)

    @property
    @pulumi.getter(name="consistencyLevel")
    def consistency_level(self) -> pulumi.Input[str]:
        """
        The Consistency Level to use for this CosmosDB Account - can be either `BoundedStaleness`, `Eventual`, `Session`, `Strong` or `ConsistentPrefix`.
        """
        return pulumi.get(self, "consistency_level")

    @consistency_level.setter
    def consistency_level(self, value: pulumi.Input[str]):
        pulumi.set(self, "consistency_level", value)

    @property
    @pulumi.getter(name="maxIntervalInSeconds")
    def max_interval_in_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        When used with the Bounded Staleness consistency level, this value represents the time amount of staleness (in seconds) tolerated. Accepted range for this value is `5` - `86400` (1 day). Defaults to `5`. Required when `consistency_level` is set to `BoundedStaleness`.
        """
        return pulumi.get(self, "max_interval_in_seconds")

    @max_interval_in_seconds.setter
    def max_interval_in_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_interval_in_seconds", value)

    @property
    @pulumi.getter(name="maxStalenessPrefix")
    def max_staleness_prefix(self) -> Optional[pulumi.Input[int]]:
        """
        When used with the Bounded Staleness consistency level, this value represents the number of stale requests tolerated. Accepted range for this value is `10` – `2147483647`. Defaults to `100`. Required when `consistency_level` is set to `BoundedStaleness`.
        """
        return pulumi.get(self, "max_staleness_prefix")

    @max_staleness_prefix.setter
    def max_staleness_prefix(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_staleness_prefix", value)


@pulumi.input_type
class AccountGeoLocationArgs:
    def __init__(__self__, *,
                 failover_priority: pulumi.Input[int],
                 location: pulumi.Input[str],
                 id: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] failover_priority: The failover priority of the region. A failover priority of `0` indicates a write region. The maximum value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the regions in which the database account exists. Changing this causes the location to be re-provisioned and cannot be changed for the location with failover priority `0`.
        :param pulumi.Input[str] location: The name of the Azure region to host replicated data.
        :param pulumi.Input[str] id: The ID of the virtual network subnet.
        :param pulumi.Input[str] prefix: The string used to generate the document endpoints for this region. If not specified it defaults to `${cosmosdb_account.name}-${location}`. Changing this causes the location to be deleted and re-provisioned and cannot be changed for the location with failover priority `0`.
        """
        pulumi.set(__self__, "failover_priority", failover_priority)
        pulumi.set(__self__, "location", location)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if prefix is not None:
            warnings.warn("This is deprecated because the service no longer accepts this as an input since Apr 25, 2019", DeprecationWarning)
            pulumi.log.warn("prefix is deprecated: This is deprecated because the service no longer accepts this as an input since Apr 25, 2019")
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)

    @property
    @pulumi.getter(name="failoverPriority")
    def failover_priority(self) -> pulumi.Input[int]:
        """
        The failover priority of the region. A failover priority of `0` indicates a write region. The maximum value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the regions in which the database account exists. Changing this causes the location to be re-provisioned and cannot be changed for the location with failover priority `0`.
        """
        return pulumi.get(self, "failover_priority")

    @failover_priority.setter
    def failover_priority(self, value: pulumi.Input[int]):
        pulumi.set(self, "failover_priority", value)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        """
        The name of the Azure region to host replicated data.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the virtual network subnet.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The string used to generate the document endpoints for this region. If not specified it defaults to `${cosmosdb_account.name}-${location}`. Changing this causes the location to be deleted and re-provisioned and cannot be changed for the location with failover priority `0`.
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)


@pulumi.input_type
class AccountVirtualNetworkRuleArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 ignore_missing_vnet_service_endpoint: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[str] id: The ID of the virtual network subnet.
        :param pulumi.Input[bool] ignore_missing_vnet_service_endpoint: If set to true, the specified subnet will be added as a virtual network rule even if its CosmosDB service endpoint is not active. Defaults to `false`.
        """
        pulumi.set(__self__, "id", id)
        if ignore_missing_vnet_service_endpoint is not None:
            pulumi.set(__self__, "ignore_missing_vnet_service_endpoint", ignore_missing_vnet_service_endpoint)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        The ID of the virtual network subnet.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="ignoreMissingVnetServiceEndpoint")
    def ignore_missing_vnet_service_endpoint(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to true, the specified subnet will be added as a virtual network rule even if its CosmosDB service endpoint is not active. Defaults to `false`.
        """
        return pulumi.get(self, "ignore_missing_vnet_service_endpoint")

    @ignore_missing_vnet_service_endpoint.setter
    def ignore_missing_vnet_service_endpoint(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_missing_vnet_service_endpoint", value)


@pulumi.input_type
class GremlinGraphConflictResolutionPolicyArgs:
    def __init__(__self__, *,
                 mode: pulumi.Input[str],
                 conflict_resolution_path: Optional[pulumi.Input[str]] = None,
                 conflict_resolution_procedure: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] mode: Indicates the conflict resolution mode. Possible values include: `LastWriterWins`, `Custom`.
        :param pulumi.Input[str] conflict_resolution_path: The conflict resolution path in the case of LastWriterWins mode.
        :param pulumi.Input[str] conflict_resolution_procedure: The procedure to resolve conflicts in the case of custom mode.
        """
        pulumi.set(__self__, "mode", mode)
        if conflict_resolution_path is not None:
            pulumi.set(__self__, "conflict_resolution_path", conflict_resolution_path)
        if conflict_resolution_procedure is not None:
            pulumi.set(__self__, "conflict_resolution_procedure", conflict_resolution_procedure)

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Input[str]:
        """
        Indicates the conflict resolution mode. Possible values include: `LastWriterWins`, `Custom`.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter(name="conflictResolutionPath")
    def conflict_resolution_path(self) -> Optional[pulumi.Input[str]]:
        """
        The conflict resolution path in the case of LastWriterWins mode.
        """
        return pulumi.get(self, "conflict_resolution_path")

    @conflict_resolution_path.setter
    def conflict_resolution_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "conflict_resolution_path", value)

    @property
    @pulumi.getter(name="conflictResolutionProcedure")
    def conflict_resolution_procedure(self) -> Optional[pulumi.Input[str]]:
        """
        The procedure to resolve conflicts in the case of custom mode.
        """
        return pulumi.get(self, "conflict_resolution_procedure")

    @conflict_resolution_procedure.setter
    def conflict_resolution_procedure(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "conflict_resolution_procedure", value)


@pulumi.input_type
class GremlinGraphIndexPolicyArgs:
    def __init__(__self__, *,
                 indexing_mode: pulumi.Input[str],
                 automatic: Optional[pulumi.Input[bool]] = None,
                 excluded_paths: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 included_paths: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] indexing_mode: Indicates the indexing mode. Possible values include: `Consistent`, `Lazy`, `None`.
        :param pulumi.Input[bool] automatic: Indicates if the indexing policy is automatic. Defaults to `true`.
        :param pulumi.Input[List[pulumi.Input[str]]] excluded_paths: List of paths to exclude from indexing. Required if `indexing_mode` is `Consistent` or `Lazy`.
        :param pulumi.Input[List[pulumi.Input[str]]] included_paths: List of paths to include in the indexing. Required if `indexing_mode` is `Consistent` or `Lazy`.
        """
        pulumi.set(__self__, "indexing_mode", indexing_mode)
        if automatic is not None:
            pulumi.set(__self__, "automatic", automatic)
        if excluded_paths is not None:
            pulumi.set(__self__, "excluded_paths", excluded_paths)
        if included_paths is not None:
            pulumi.set(__self__, "included_paths", included_paths)

    @property
    @pulumi.getter(name="indexingMode")
    def indexing_mode(self) -> pulumi.Input[str]:
        """
        Indicates the indexing mode. Possible values include: `Consistent`, `Lazy`, `None`.
        """
        return pulumi.get(self, "indexing_mode")

    @indexing_mode.setter
    def indexing_mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "indexing_mode", value)

    @property
    @pulumi.getter
    def automatic(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if the indexing policy is automatic. Defaults to `true`.
        """
        return pulumi.get(self, "automatic")

    @automatic.setter
    def automatic(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "automatic", value)

    @property
    @pulumi.getter(name="excludedPaths")
    def excluded_paths(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        List of paths to exclude from indexing. Required if `indexing_mode` is `Consistent` or `Lazy`.
        """
        return pulumi.get(self, "excluded_paths")

    @excluded_paths.setter
    def excluded_paths(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "excluded_paths", value)

    @property
    @pulumi.getter(name="includedPaths")
    def included_paths(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        List of paths to include in the indexing. Required if `indexing_mode` is `Consistent` or `Lazy`.
        """
        return pulumi.get(self, "included_paths")

    @included_paths.setter
    def included_paths(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "included_paths", value)


@pulumi.input_type
class GremlinGraphUniqueKeyArgs:
    def __init__(__self__, *,
                 paths: pulumi.Input[List[pulumi.Input[str]]]):
        """
        :param pulumi.Input[List[pulumi.Input[str]]] paths: A list of paths to use for this unique key.
        """
        pulumi.set(__self__, "paths", paths)

    @property
    @pulumi.getter
    def paths(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        A list of paths to use for this unique key.
        """
        return pulumi.get(self, "paths")

    @paths.setter
    def paths(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "paths", value)


@pulumi.input_type
class MongoCollectionIndexArgs:
    def __init__(__self__, *,
                 keys: pulumi.Input[List[pulumi.Input[str]]],
                 unique: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[List[pulumi.Input[str]]] keys: Specifies the list of user settable keys for each Cosmos DB Mongo Collection.
        :param pulumi.Input[bool] unique: Is the index unique or not? Defaults to `false`.
        """
        pulumi.set(__self__, "keys", keys)
        if unique is not None:
            pulumi.set(__self__, "unique", unique)

    @property
    @pulumi.getter
    def keys(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        Specifies the list of user settable keys for each Cosmos DB Mongo Collection.
        """
        return pulumi.get(self, "keys")

    @keys.setter
    def keys(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "keys", value)

    @property
    @pulumi.getter
    def unique(self) -> Optional[pulumi.Input[bool]]:
        """
        Is the index unique or not? Defaults to `false`.
        """
        return pulumi.get(self, "unique")

    @unique.setter
    def unique(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "unique", value)


@pulumi.input_type
class MongoCollectionSystemIndexArgs:
    def __init__(__self__, *,
                 keys: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 unique: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[List[pulumi.Input[str]]] keys: Specifies the list of user settable keys for each Cosmos DB Mongo Collection.
        :param pulumi.Input[bool] unique: Is the index unique or not? Defaults to `false`.
        """
        if keys is not None:
            pulumi.set(__self__, "keys", keys)
        if unique is not None:
            pulumi.set(__self__, "unique", unique)

    @property
    @pulumi.getter
    def keys(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        Specifies the list of user settable keys for each Cosmos DB Mongo Collection.
        """
        return pulumi.get(self, "keys")

    @keys.setter
    def keys(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "keys", value)

    @property
    @pulumi.getter
    def unique(self) -> Optional[pulumi.Input[bool]]:
        """
        Is the index unique or not? Defaults to `false`.
        """
        return pulumi.get(self, "unique")

    @unique.setter
    def unique(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "unique", value)


@pulumi.input_type
class SqlContainerUniqueKeyArgs:
    def __init__(__self__, *,
                 paths: pulumi.Input[List[pulumi.Input[str]]]):
        """
        :param pulumi.Input[List[pulumi.Input[str]]] paths: A list of paths to use for this unique key.
        """
        pulumi.set(__self__, "paths", paths)

    @property
    @pulumi.getter
    def paths(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        A list of paths to use for this unique key.
        """
        return pulumi.get(self, "paths")

    @paths.setter
    def paths(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "paths", value)



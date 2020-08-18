# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GetClusterResult:
    """
    A collection of values returned by getCluster.
    """
    def __init__(__self__, data_ingestion_uri=None, id=None, location=None, name=None, resource_group_name=None, tags=None, uri=None):
        if data_ingestion_uri and not isinstance(data_ingestion_uri, str):
            raise TypeError("Expected argument 'data_ingestion_uri' to be a str")
        __self__.data_ingestion_uri = data_ingestion_uri
        """
        The Kusto Cluster URI to be used for data ingestion.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
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
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        if uri and not isinstance(uri, str):
            raise TypeError("Expected argument 'uri' to be a str")
        __self__.uri = uri
        """
        The FQDN of the Azure Kusto Cluster.
        """


class AwaitableGetClusterResult(GetClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterResult(
            data_ingestion_uri=self.data_ingestion_uri,
            id=self.id,
            location=self.location,
            name=self.name,
            resource_group_name=self.resource_group_name,
            tags=self.tags,
            uri=self.uri)


def get_cluster(name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing Kusto (also known as Azure Data Explorer) Cluster

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.kusto.get_cluster(name="kustocluster",
        resource_group_name="test_resource_group")
    ```


    :param str name: Specifies the name of the Kusto Cluster.
    :param str resource_group_name: The name of the Resource Group where the Kusto Cluster exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:kusto/getCluster:getCluster', __args__, opts=opts).value

    return AwaitableGetClusterResult(
        data_ingestion_uri=__ret__.get('dataIngestionUri'),
        id=__ret__.get('id'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        resource_group_name=__ret__.get('resourceGroupName'),
        tags=__ret__.get('tags'),
        uri=__ret__.get('uri'))

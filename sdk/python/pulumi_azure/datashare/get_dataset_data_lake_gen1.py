# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetDatasetDataLakeGen1Result',
    'AwaitableGetDatasetDataLakeGen1Result',
    'get_dataset_data_lake_gen1',
]


@pulumi.output_type
class _GetDatasetDataLakeGen1Result(dict):
    data_lake_store_id: str = pulumi.property("dataLakeStoreId")
    data_share_id: str = pulumi.property("dataShareId")
    display_name: str = pulumi.property("displayName")
    file_name: str = pulumi.property("fileName")
    folder_path: str = pulumi.property("folderPath")
    id: str = pulumi.property("id")
    name: str = pulumi.property("name")


class GetDatasetDataLakeGen1Result:
    """
    A collection of values returned by getDatasetDataLakeGen1.
    """
    def __init__(__self__, data_lake_store_id=None, data_share_id=None, display_name=None, file_name=None, folder_path=None, id=None, name=None):
        if data_lake_store_id and not isinstance(data_lake_store_id, str):
            raise TypeError("Expected argument 'data_lake_store_id' to be a str")
        __self__.data_lake_store_id = data_lake_store_id
        """
        The resource ID of the Data Lake Store to be shared with the receiver.
        """
        if data_share_id and not isinstance(data_share_id, str):
            raise TypeError("Expected argument 'data_share_id' to be a str")
        __self__.data_share_id = data_share_id
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        __self__.display_name = display_name
        """
        The displayed name of the Data Share Dataset.
        """
        if file_name and not isinstance(file_name, str):
            raise TypeError("Expected argument 'file_name' to be a str")
        __self__.file_name = file_name
        """
        The file name of the data lake store to be shared with the receiver.
        """
        if folder_path and not isinstance(folder_path, str):
            raise TypeError("Expected argument 'folder_path' to be a str")
        __self__.folder_path = folder_path
        """
        The folder path of the data lake store to be shared with the receiver.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name


class AwaitableGetDatasetDataLakeGen1Result(GetDatasetDataLakeGen1Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatasetDataLakeGen1Result(
            data_lake_store_id=self.data_lake_store_id,
            data_share_id=self.data_share_id,
            display_name=self.display_name,
            file_name=self.file_name,
            folder_path=self.folder_path,
            id=self.id,
            name=self.name)


def get_dataset_data_lake_gen1(data_share_id: Optional[str] = None,
                               name: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatasetDataLakeGen1Result:
    """
    Use this data source to access information about an existing DataShareDataLakeGen1Dataset.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.datashare.get_dataset_data_lake_gen1(name="example-dsdsdlg1",
        data_share_id="example-share-id")
    pulumi.export("id", example.id)
    ```


    :param str data_share_id: The resource ID of the Data Share where this Data Share Data Lake Gen1 Dataset should be created.
    :param str name: The name of the Data Share Data Lake Gen1 Dataset.
    """
    __args__ = dict()
    __args__['dataShareId'] = data_share_id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:datashare/getDatasetDataLakeGen1:getDatasetDataLakeGen1', __args__, opts=opts, typ=_GetDatasetDataLakeGen1Result).value

    return AwaitableGetDatasetDataLakeGen1Result(
        data_lake_store_id=_utilities.get_dict_value(__ret__, 'dataLakeStoreId'),
        data_share_id=_utilities.get_dict_value(__ret__, 'dataShareId'),
        display_name=_utilities.get_dict_value(__ret__, 'displayName'),
        file_name=_utilities.get_dict_value(__ret__, 'fileName'),
        folder_path=_utilities.get_dict_value(__ret__, 'folderPath'),
        id=_utilities.get_dict_value(__ret__, 'id'),
        name=_utilities.get_dict_value(__ret__, 'name'))

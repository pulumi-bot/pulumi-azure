# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['StoreFile']


class StoreFile(pulumi.CustomResource):
    account_name: pulumi.Output[str] = pulumi.property("accountName")
    """
    Specifies the name of the Data Lake Store for which the File should created.
    """

    local_file_path: pulumi.Output[str] = pulumi.property("localFilePath")
    """
    The path to the local file to be added to the Data Lake Store.
    """

    remote_file_path: pulumi.Output[str] = pulumi.property("remoteFilePath")
    """
    The path created for the file on the Data Lake Store.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 local_file_path: Optional[pulumi.Input[str]] = None,
                 remote_file_path: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Azure Data Lake Store File.

        > **Note:** If you want to change the data in the remote file without changing the `local_file_path`, then
        taint the resource so the `datalake.StoreFile` gets recreated with the new data.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: Specifies the name of the Data Lake Store for which the File should created.
        :param pulumi.Input[str] local_file_path: The path to the local file to be added to the Data Lake Store.
        :param pulumi.Input[str] remote_file_path: The path created for the file on the Data Lake Store.
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

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            if local_file_path is None:
                raise TypeError("Missing required property 'local_file_path'")
            __props__['local_file_path'] = local_file_path
            if remote_file_path is None:
                raise TypeError("Missing required property 'remote_file_path'")
            __props__['remote_file_path'] = remote_file_path
        super(StoreFile, __self__).__init__(
            'azure:datalake/storeFile:StoreFile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            account_name: Optional[pulumi.Input[str]] = None,
            local_file_path: Optional[pulumi.Input[str]] = None,
            remote_file_path: Optional[pulumi.Input[str]] = None) -> 'StoreFile':
        """
        Get an existing StoreFile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: Specifies the name of the Data Lake Store for which the File should created.
        :param pulumi.Input[str] local_file_path: The path to the local file to be added to the Data Lake Store.
        :param pulumi.Input[str] remote_file_path: The path created for the file on the Data Lake Store.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_name"] = account_name
        __props__["local_file_path"] = local_file_path
        __props__["remote_file_path"] = remote_file_path
        return StoreFile(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


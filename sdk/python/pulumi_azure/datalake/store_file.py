# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['StoreFileArgs', 'StoreFile']

@pulumi.input_type
class StoreFileArgs:
    def __init__(__self__, *,
                 account_name: pulumi.Input[str],
                 local_file_path: pulumi.Input[str],
                 remote_file_path: pulumi.Input[str]):
        """
        The set of arguments for constructing a StoreFile resource.
        :param pulumi.Input[str] account_name: Specifies the name of the Data Lake Store for which the File should created.
        :param pulumi.Input[str] local_file_path: The path to the local file to be added to the Data Lake Store.
        :param pulumi.Input[str] remote_file_path: The path created for the file on the Data Lake Store.
        """
        pulumi.set(__self__, "account_name", account_name)
        pulumi.set(__self__, "local_file_path", local_file_path)
        pulumi.set(__self__, "remote_file_path", remote_file_path)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the Data Lake Store for which the File should created.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="localFilePath")
    def local_file_path(self) -> pulumi.Input[str]:
        """
        The path to the local file to be added to the Data Lake Store.
        """
        return pulumi.get(self, "local_file_path")

    @local_file_path.setter
    def local_file_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "local_file_path", value)

    @property
    @pulumi.getter(name="remoteFilePath")
    def remote_file_path(self) -> pulumi.Input[str]:
        """
        The path created for the file on the Data Lake Store.
        """
        return pulumi.get(self, "remote_file_path")

    @remote_file_path.setter
    def remote_file_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "remote_file_path", value)


class StoreFile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
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

        ## Import

        Data Lake Store File's can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:datalake/storeFile:StoreFile txt
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: Specifies the name of the Data Lake Store for which the File should created.
        :param pulumi.Input[str] local_file_path: The path to the local file to be added to the Data Lake Store.
        :param pulumi.Input[str] remote_file_path: The path created for the file on the Data Lake Store.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StoreFileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Azure Data Lake Store File.

        > **Note:** If you want to change the data in the remote file without changing the `local_file_path`, then
        taint the resource so the `datalake.StoreFile` gets recreated with the new data.

        ## Import

        Data Lake Store File's can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:datalake/storeFile:StoreFile txt
        ```

        :param str resource_name: The name of the resource.
        :param StoreFileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StoreFileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 local_file_path: Optional[pulumi.Input[str]] = None,
                 remote_file_path: Optional[pulumi.Input[str]] = None,
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

            if account_name is None and not opts.urn:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            if local_file_path is None and not opts.urn:
                raise TypeError("Missing required property 'local_file_path'")
            __props__['local_file_path'] = local_file_path
            if remote_file_path is None and not opts.urn:
                raise TypeError("Missing required property 'remote_file_path'")
            __props__['remote_file_path'] = remote_file_path
        super(StoreFile, __self__).__init__(
            'azure:datalake/storeFile:StoreFile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_name: Optional[pulumi.Input[str]] = None,
            local_file_path: Optional[pulumi.Input[str]] = None,
            remote_file_path: Optional[pulumi.Input[str]] = None) -> 'StoreFile':
        """
        Get an existing StoreFile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
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

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Data Lake Store for which the File should created.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="localFilePath")
    def local_file_path(self) -> pulumi.Output[str]:
        """
        The path to the local file to be added to the Data Lake Store.
        """
        return pulumi.get(self, "local_file_path")

    @property
    @pulumi.getter(name="remoteFilePath")
    def remote_file_path(self) -> pulumi.Output[str]:
        """
        The path created for the file on the Data Lake Store.
        """
        return pulumi.get(self, "remote_file_path")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


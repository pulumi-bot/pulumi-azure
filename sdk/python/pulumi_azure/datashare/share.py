# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Share']


class Share(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 snapshot_schedule: Optional[pulumi.Input[pulumi.InputType['ShareSnapshotScheduleArgs']]] = None,
                 terms: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Data Share.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.datashare.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tags={
                "foo": "bar",
            })
        example_share = azure.datashare.Share("exampleShare",
            account_id=example_account.id,
            kind="CopyBased",
            description="example desc",
            terms="example terms",
            snapshot_schedule=azure.datashare.ShareSnapshotScheduleArgs(
                name="example-ss",
                recurrence="Day",
                start_time="2020-04-17T04:47:52.9614956Z",
            ))
        ```

        ## Import

        Data Shares can be imported using the `resource id`, e.g.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The ID of the Data Share account in which the Data Share is created. Changing this forces a new Data Share to be created.
        :param pulumi.Input[str] description: The Data Share's description.
        :param pulumi.Input[str] kind: The kind of the Data Share. Possible values are `CopyBased` and `InPlace`. Changing this forces a new Data Share to be created.
        :param pulumi.Input[str] name: The name which should be used for this Data Share. Changing this forces a new Data Share to be created.
        :param pulumi.Input[pulumi.InputType['ShareSnapshotScheduleArgs']] snapshot_schedule: A `snapshot_schedule` block as defined below.
        :param pulumi.Input[str] terms: The terms of the Data Share.
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

            if account_id is None:
                raise TypeError("Missing required property 'account_id'")
            __props__['account_id'] = account_id
            __props__['description'] = description
            if kind is None:
                raise TypeError("Missing required property 'kind'")
            __props__['kind'] = kind
            __props__['name'] = name
            __props__['snapshot_schedule'] = snapshot_schedule
            __props__['terms'] = terms
        super(Share, __self__).__init__(
            'azure:datashare/share:Share',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            kind: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            snapshot_schedule: Optional[pulumi.Input[pulumi.InputType['ShareSnapshotScheduleArgs']]] = None,
            terms: Optional[pulumi.Input[str]] = None) -> 'Share':
        """
        Get an existing Share resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The ID of the Data Share account in which the Data Share is created. Changing this forces a new Data Share to be created.
        :param pulumi.Input[str] description: The Data Share's description.
        :param pulumi.Input[str] kind: The kind of the Data Share. Possible values are `CopyBased` and `InPlace`. Changing this forces a new Data Share to be created.
        :param pulumi.Input[str] name: The name which should be used for this Data Share. Changing this forces a new Data Share to be created.
        :param pulumi.Input[pulumi.InputType['ShareSnapshotScheduleArgs']] snapshot_schedule: A `snapshot_schedule` block as defined below.
        :param pulumi.Input[str] terms: The terms of the Data Share.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_id"] = account_id
        __props__["description"] = description
        __props__["kind"] = kind
        __props__["name"] = name
        __props__["snapshot_schedule"] = snapshot_schedule
        __props__["terms"] = terms
        return Share(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[str]:
        """
        The ID of the Data Share account in which the Data Share is created. Changing this forces a new Data Share to be created.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The Data Share's description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        The kind of the Data Share. Possible values are `CopyBased` and `InPlace`. Changing this forces a new Data Share to be created.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this Data Share. Changing this forces a new Data Share to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="snapshotSchedule")
    def snapshot_schedule(self) -> pulumi.Output[Optional['outputs.ShareSnapshotSchedule']]:
        """
        A `snapshot_schedule` block as defined below.
        """
        return pulumi.get(self, "snapshot_schedule")

    @property
    @pulumi.getter
    def terms(self) -> pulumi.Output[Optional[str]]:
        """
        The terms of the Data Share.
        """
        return pulumi.get(self, "terms")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


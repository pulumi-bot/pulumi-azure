# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['LogProfileArgs', 'LogProfile']

@pulumi.input_type
class LogProfileArgs:
    def __init__(__self__, *,
                 categories: pulumi.Input[Sequence[pulumi.Input[str]]],
                 locations: pulumi.Input[Sequence[pulumi.Input[str]]],
                 retention_policy: pulumi.Input['LogProfileRetentionPolicyArgs'],
                 name: Optional[pulumi.Input[str]] = None,
                 servicebus_rule_id: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LogProfile resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] categories: List of categories of the logs.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] locations: List of regions for which Activity Log events are stored or streamed.
        :param pulumi.Input['LogProfileRetentionPolicyArgs'] retention_policy: A `retention_policy` block as documented below. A retention policy for how long Activity Logs are retained in the storage account.
        :param pulumi.Input[str] name: The name of the Log Profile. Changing this forces a
               new resource to be created.
        :param pulumi.Input[str] servicebus_rule_id: The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to. At least one of `storage_account_id` or `servicebus_rule_id` must be set.
        :param pulumi.Input[str] storage_account_id: The resource ID of the storage account in which the Activity Log is stored. At least one of `storage_account_id` or `servicebus_rule_id` must be set.
        """
        pulumi.set(__self__, "categories", categories)
        pulumi.set(__self__, "locations", locations)
        pulumi.set(__self__, "retention_policy", retention_policy)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if servicebus_rule_id is not None:
            pulumi.set(__self__, "servicebus_rule_id", servicebus_rule_id)
        if storage_account_id is not None:
            pulumi.set(__self__, "storage_account_id", storage_account_id)

    @property
    @pulumi.getter
    def categories(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of categories of the logs.
        """
        return pulumi.get(self, "categories")

    @categories.setter
    def categories(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "categories", value)

    @property
    @pulumi.getter
    def locations(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of regions for which Activity Log events are stored or streamed.
        """
        return pulumi.get(self, "locations")

    @locations.setter
    def locations(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "locations", value)

    @property
    @pulumi.getter(name="retentionPolicy")
    def retention_policy(self) -> pulumi.Input['LogProfileRetentionPolicyArgs']:
        """
        A `retention_policy` block as documented below. A retention policy for how long Activity Logs are retained in the storage account.
        """
        return pulumi.get(self, "retention_policy")

    @retention_policy.setter
    def retention_policy(self, value: pulumi.Input['LogProfileRetentionPolicyArgs']):
        pulumi.set(self, "retention_policy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Log Profile. Changing this forces a
        new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="servicebusRuleId")
    def servicebus_rule_id(self) -> Optional[pulumi.Input[str]]:
        """
        The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to. At least one of `storage_account_id` or `servicebus_rule_id` must be set.
        """
        return pulumi.get(self, "servicebus_rule_id")

    @servicebus_rule_id.setter
    def servicebus_rule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "servicebus_rule_id", value)

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource ID of the storage account in which the Activity Log is stored. At least one of `storage_account_id` or `servicebus_rule_id` must be set.
        """
        return pulumi.get(self, "storage_account_id")

    @storage_account_id.setter
    def storage_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_id", value)


@pulumi.input_type
class _LogProfileState:
    def __init__(__self__, *,
                 categories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retention_policy: Optional[pulumi.Input['LogProfileRetentionPolicyArgs']] = None,
                 servicebus_rule_id: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LogProfile resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] categories: List of categories of the logs.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] locations: List of regions for which Activity Log events are stored or streamed.
        :param pulumi.Input[str] name: The name of the Log Profile. Changing this forces a
               new resource to be created.
        :param pulumi.Input['LogProfileRetentionPolicyArgs'] retention_policy: A `retention_policy` block as documented below. A retention policy for how long Activity Logs are retained in the storage account.
        :param pulumi.Input[str] servicebus_rule_id: The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to. At least one of `storage_account_id` or `servicebus_rule_id` must be set.
        :param pulumi.Input[str] storage_account_id: The resource ID of the storage account in which the Activity Log is stored. At least one of `storage_account_id` or `servicebus_rule_id` must be set.
        """
        if categories is not None:
            pulumi.set(__self__, "categories", categories)
        if locations is not None:
            pulumi.set(__self__, "locations", locations)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if retention_policy is not None:
            pulumi.set(__self__, "retention_policy", retention_policy)
        if servicebus_rule_id is not None:
            pulumi.set(__self__, "servicebus_rule_id", servicebus_rule_id)
        if storage_account_id is not None:
            pulumi.set(__self__, "storage_account_id", storage_account_id)

    @property
    @pulumi.getter
    def categories(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of categories of the logs.
        """
        return pulumi.get(self, "categories")

    @categories.setter
    def categories(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "categories", value)

    @property
    @pulumi.getter
    def locations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of regions for which Activity Log events are stored or streamed.
        """
        return pulumi.get(self, "locations")

    @locations.setter
    def locations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "locations", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Log Profile. Changing this forces a
        new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="retentionPolicy")
    def retention_policy(self) -> Optional[pulumi.Input['LogProfileRetentionPolicyArgs']]:
        """
        A `retention_policy` block as documented below. A retention policy for how long Activity Logs are retained in the storage account.
        """
        return pulumi.get(self, "retention_policy")

    @retention_policy.setter
    def retention_policy(self, value: Optional[pulumi.Input['LogProfileRetentionPolicyArgs']]):
        pulumi.set(self, "retention_policy", value)

    @property
    @pulumi.getter(name="servicebusRuleId")
    def servicebus_rule_id(self) -> Optional[pulumi.Input[str]]:
        """
        The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to. At least one of `storage_account_id` or `servicebus_rule_id` must be set.
        """
        return pulumi.get(self, "servicebus_rule_id")

    @servicebus_rule_id.setter
    def servicebus_rule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "servicebus_rule_id", value)

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource ID of the storage account in which the Activity Log is stored. At least one of `storage_account_id` or `servicebus_rule_id` must be set.
        """
        return pulumi.get(self, "storage_account_id")

    @storage_account_id.setter
    def storage_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_id", value)


class LogProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 categories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retention_policy: Optional[pulumi.Input[pulumi.InputType['LogProfileRetentionPolicyArgs']]] = None,
                 servicebus_rule_id: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a [Log Profile](https://docs.microsoft.com/en-us/azure/monitoring-and-diagnostics/monitoring-overview-activity-logs#export-the-activity-log-with-a-log-profile). A Log Profile configures how Activity Logs are exported.

        > **NOTE:** It's only possible to configure one Log Profile per Subscription. If you are trying to create more than one Log Profile, an error with `StatusCode=409` will occur.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="GRS")
        example_event_hub_namespace = azure.eventhub.EventHubNamespace("exampleEventHubNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard",
            capacity=2)
        example_log_profile = azure.monitoring.LogProfile("exampleLogProfile",
            categories=[
                "Action",
                "Delete",
                "Write",
            ],
            locations=[
                "westus",
                "global",
            ],
            servicebus_rule_id=example_event_hub_namespace.id.apply(lambda id: f"{id}/authorizationrules/RootManageSharedAccessKey"),
            storage_account_id=example_account.id,
            retention_policy=azure.monitoring.LogProfileRetentionPolicyArgs(
                enabled=True,
                days=7,
            ))
        ```

        ## Import

        A Log Profile can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:monitoring/logProfile:LogProfile example /subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.insights/logprofiles/test
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] categories: List of categories of the logs.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] locations: List of regions for which Activity Log events are stored or streamed.
        :param pulumi.Input[str] name: The name of the Log Profile. Changing this forces a
               new resource to be created.
        :param pulumi.Input[pulumi.InputType['LogProfileRetentionPolicyArgs']] retention_policy: A `retention_policy` block as documented below. A retention policy for how long Activity Logs are retained in the storage account.
        :param pulumi.Input[str] servicebus_rule_id: The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to. At least one of `storage_account_id` or `servicebus_rule_id` must be set.
        :param pulumi.Input[str] storage_account_id: The resource ID of the storage account in which the Activity Log is stored. At least one of `storage_account_id` or `servicebus_rule_id` must be set.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LogProfileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a [Log Profile](https://docs.microsoft.com/en-us/azure/monitoring-and-diagnostics/monitoring-overview-activity-logs#export-the-activity-log-with-a-log-profile). A Log Profile configures how Activity Logs are exported.

        > **NOTE:** It's only possible to configure one Log Profile per Subscription. If you are trying to create more than one Log Profile, an error with `StatusCode=409` will occur.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="GRS")
        example_event_hub_namespace = azure.eventhub.EventHubNamespace("exampleEventHubNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard",
            capacity=2)
        example_log_profile = azure.monitoring.LogProfile("exampleLogProfile",
            categories=[
                "Action",
                "Delete",
                "Write",
            ],
            locations=[
                "westus",
                "global",
            ],
            servicebus_rule_id=example_event_hub_namespace.id.apply(lambda id: f"{id}/authorizationrules/RootManageSharedAccessKey"),
            storage_account_id=example_account.id,
            retention_policy=azure.monitoring.LogProfileRetentionPolicyArgs(
                enabled=True,
                days=7,
            ))
        ```

        ## Import

        A Log Profile can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:monitoring/logProfile:LogProfile example /subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.insights/logprofiles/test
        ```

        :param str resource_name: The name of the resource.
        :param LogProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LogProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 categories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retention_policy: Optional[pulumi.Input[pulumi.InputType['LogProfileRetentionPolicyArgs']]] = None,
                 servicebus_rule_id: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = LogProfileArgs.__new__(LogProfileArgs)

            if categories is None and not opts.urn:
                raise TypeError("Missing required property 'categories'")
            __props__.__dict__["categories"] = categories
            if locations is None and not opts.urn:
                raise TypeError("Missing required property 'locations'")
            __props__.__dict__["locations"] = locations
            __props__.__dict__["name"] = name
            if retention_policy is None and not opts.urn:
                raise TypeError("Missing required property 'retention_policy'")
            __props__.__dict__["retention_policy"] = retention_policy
            __props__.__dict__["servicebus_rule_id"] = servicebus_rule_id
            __props__.__dict__["storage_account_id"] = storage_account_id
        super(LogProfile, __self__).__init__(
            'azure:monitoring/logProfile:LogProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            categories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            locations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            retention_policy: Optional[pulumi.Input[pulumi.InputType['LogProfileRetentionPolicyArgs']]] = None,
            servicebus_rule_id: Optional[pulumi.Input[str]] = None,
            storage_account_id: Optional[pulumi.Input[str]] = None) -> 'LogProfile':
        """
        Get an existing LogProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] categories: List of categories of the logs.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] locations: List of regions for which Activity Log events are stored or streamed.
        :param pulumi.Input[str] name: The name of the Log Profile. Changing this forces a
               new resource to be created.
        :param pulumi.Input[pulumi.InputType['LogProfileRetentionPolicyArgs']] retention_policy: A `retention_policy` block as documented below. A retention policy for how long Activity Logs are retained in the storage account.
        :param pulumi.Input[str] servicebus_rule_id: The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to. At least one of `storage_account_id` or `servicebus_rule_id` must be set.
        :param pulumi.Input[str] storage_account_id: The resource ID of the storage account in which the Activity Log is stored. At least one of `storage_account_id` or `servicebus_rule_id` must be set.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LogProfileState.__new__(_LogProfileState)

        __props__.__dict__["categories"] = categories
        __props__.__dict__["locations"] = locations
        __props__.__dict__["name"] = name
        __props__.__dict__["retention_policy"] = retention_policy
        __props__.__dict__["servicebus_rule_id"] = servicebus_rule_id
        __props__.__dict__["storage_account_id"] = storage_account_id
        return LogProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def categories(self) -> pulumi.Output[Sequence[str]]:
        """
        List of categories of the logs.
        """
        return pulumi.get(self, "categories")

    @property
    @pulumi.getter
    def locations(self) -> pulumi.Output[Sequence[str]]:
        """
        List of regions for which Activity Log events are stored or streamed.
        """
        return pulumi.get(self, "locations")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Log Profile. Changing this forces a
        new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="retentionPolicy")
    def retention_policy(self) -> pulumi.Output['outputs.LogProfileRetentionPolicy']:
        """
        A `retention_policy` block as documented below. A retention policy for how long Activity Logs are retained in the storage account.
        """
        return pulumi.get(self, "retention_policy")

    @property
    @pulumi.getter(name="servicebusRuleId")
    def servicebus_rule_id(self) -> pulumi.Output[Optional[str]]:
        """
        The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to. At least one of `storage_account_id` or `servicebus_rule_id` must be set.
        """
        return pulumi.get(self, "servicebus_rule_id")

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> pulumi.Output[Optional[str]]:
        """
        The resource ID of the storage account in which the Activity Log is stored. At least one of `storage_account_id` or `servicebus_rule_id` must be set.
        """
        return pulumi.get(self, "storage_account_id")


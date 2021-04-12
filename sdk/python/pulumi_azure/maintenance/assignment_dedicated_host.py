# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['AssignmentDedicatedHostArgs', 'AssignmentDedicatedHost']

@pulumi.input_type
class AssignmentDedicatedHostArgs:
    def __init__(__self__, *,
                 dedicated_host_id: pulumi.Input[str],
                 maintenance_configuration_id: pulumi.Input[str],
                 location: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AssignmentDedicatedHost resource.
        :param pulumi.Input[str] dedicated_host_id: Specifies the Dedicated Host ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
        :param pulumi.Input[str] maintenance_configuration_id: Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "dedicated_host_id", dedicated_host_id)
        pulumi.set(__self__, "maintenance_configuration_id", maintenance_configuration_id)
        if location is not None:
            pulumi.set(__self__, "location", location)

    @property
    @pulumi.getter(name="dedicatedHostId")
    def dedicated_host_id(self) -> pulumi.Input[str]:
        """
        Specifies the Dedicated Host ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "dedicated_host_id")

    @dedicated_host_id.setter
    def dedicated_host_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "dedicated_host_id", value)

    @property
    @pulumi.getter(name="maintenanceConfigurationId")
    def maintenance_configuration_id(self) -> pulumi.Input[str]:
        """
        Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "maintenance_configuration_id")

    @maintenance_configuration_id.setter
    def maintenance_configuration_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "maintenance_configuration_id", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)


class AssignmentDedicatedHost(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dedicated_host_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 maintenance_configuration_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a maintenance assignment to Dedicated Host.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_dedicated_host_group = azure.compute.DedicatedHostGroup("exampleDedicatedHostGroup",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            platform_fault_domain_count=2)
        example_dedicated_host = azure.compute.DedicatedHost("exampleDedicatedHost",
            location=example_resource_group.location,
            dedicated_host_group_id=example_dedicated_host_group.id,
            sku_name="DSv3-Type1",
            platform_fault_domain=1)
        example_configuration = azure.maintenance.Configuration("exampleConfiguration",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            scope="All")
        example_assignment_dedicated_host = azure.maintenance.AssignmentDedicatedHost("exampleAssignmentDedicatedHost",
            location=example_resource_group.location,
            maintenance_configuration_id=example_configuration.id,
            dedicated_host_id=example_dedicated_host.id)
        ```

        ## Import

        Maintenance Assignment can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:maintenance/assignmentDedicatedHost:AssignmentDedicatedHost example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resGroup1/providers/microsoft.compute/hostGroups/group1/hosts/host1/providers/Microsoft.Maintenance/configurationAssignments/assign1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dedicated_host_id: Specifies the Dedicated Host ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] maintenance_configuration_id: Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AssignmentDedicatedHostArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a maintenance assignment to Dedicated Host.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_dedicated_host_group = azure.compute.DedicatedHostGroup("exampleDedicatedHostGroup",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            platform_fault_domain_count=2)
        example_dedicated_host = azure.compute.DedicatedHost("exampleDedicatedHost",
            location=example_resource_group.location,
            dedicated_host_group_id=example_dedicated_host_group.id,
            sku_name="DSv3-Type1",
            platform_fault_domain=1)
        example_configuration = azure.maintenance.Configuration("exampleConfiguration",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            scope="All")
        example_assignment_dedicated_host = azure.maintenance.AssignmentDedicatedHost("exampleAssignmentDedicatedHost",
            location=example_resource_group.location,
            maintenance_configuration_id=example_configuration.id,
            dedicated_host_id=example_dedicated_host.id)
        ```

        ## Import

        Maintenance Assignment can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:maintenance/assignmentDedicatedHost:AssignmentDedicatedHost example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resGroup1/providers/microsoft.compute/hostGroups/group1/hosts/host1/providers/Microsoft.Maintenance/configurationAssignments/assign1
        ```

        :param str resource_name: The name of the resource.
        :param AssignmentDedicatedHostArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AssignmentDedicatedHostArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dedicated_host_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 maintenance_configuration_id: Optional[pulumi.Input[str]] = None,
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

            if dedicated_host_id is None and not opts.urn:
                raise TypeError("Missing required property 'dedicated_host_id'")
            __props__['dedicated_host_id'] = dedicated_host_id
            __props__['location'] = location
            if maintenance_configuration_id is None and not opts.urn:
                raise TypeError("Missing required property 'maintenance_configuration_id'")
            __props__['maintenance_configuration_id'] = maintenance_configuration_id
        super(AssignmentDedicatedHost, __self__).__init__(
            'azure:maintenance/assignmentDedicatedHost:AssignmentDedicatedHost',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dedicated_host_id: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            maintenance_configuration_id: Optional[pulumi.Input[str]] = None) -> 'AssignmentDedicatedHost':
        """
        Get an existing AssignmentDedicatedHost resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dedicated_host_id: Specifies the Dedicated Host ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] maintenance_configuration_id: Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["dedicated_host_id"] = dedicated_host_id
        __props__["location"] = location
        __props__["maintenance_configuration_id"] = maintenance_configuration_id
        return AssignmentDedicatedHost(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dedicatedHostId")
    def dedicated_host_id(self) -> pulumi.Output[str]:
        """
        Specifies the Dedicated Host ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "dedicated_host_id")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maintenanceConfigurationId")
    def maintenance_configuration_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "maintenance_configuration_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


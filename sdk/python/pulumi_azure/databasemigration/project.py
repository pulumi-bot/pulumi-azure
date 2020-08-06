# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Project(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specify the name of the database migration project. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    Name of the resource group in which to create the database migration project. Changing this forces a new resource to be created.
    """
    service_name: pulumi.Output[str]
    """
    Name of the database migration service where resource belongs to. Changing this forces a new resource to be created.
    """
    source_platform: pulumi.Output[str]
    """
    The platform type of the migration source. Currently only support: `SQL`(on-premises SQL Server). Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assigned to the resource.
    """
    target_platform: pulumi.Output[str]
    """
    The platform type of the migration target. Currently only support: `SQLDB`(Azure SQL Database). Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, resource_group_name=None, service_name=None, source_platform=None, tags=None, target_platform=None, __props__=None, __name__=None, __opts__=None):
        """
        Manage a Azure Database Migration Project.

        > **NOTE:** Destroying a Database Migration Project will leave any outstanding tasks untouched. This is to avoid unexpectedly deleting any tasks managed outside of this provider.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specify the name of the database migration project. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Name of the resource group in which to create the database migration project. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_name: Name of the database migration service where resource belongs to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] source_platform: The platform type of the migration source. Currently only support: `SQL`(on-premises SQL Server). Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assigned to the resource.
        :param pulumi.Input[str] target_platform: The platform type of the migration target. Currently only support: `SQLDB`(Azure SQL Database). Changing this forces a new resource to be created.
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

            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if service_name is None:
                raise TypeError("Missing required property 'service_name'")
            __props__['service_name'] = service_name
            if source_platform is None:
                raise TypeError("Missing required property 'source_platform'")
            __props__['source_platform'] = source_platform
            __props__['tags'] = tags
            if target_platform is None:
                raise TypeError("Missing required property 'target_platform'")
            __props__['target_platform'] = target_platform
        super(Project, __self__).__init__(
            'azure:databasemigration/project:Project',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, location=None, name=None, resource_group_name=None, service_name=None, source_platform=None, tags=None, target_platform=None):
        """
        Get an existing Project resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specify the name of the database migration project. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Name of the resource group in which to create the database migration project. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_name: Name of the database migration service where resource belongs to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] source_platform: The platform type of the migration source. Currently only support: `SQL`(on-premises SQL Server). Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assigned to the resource.
        :param pulumi.Input[str] target_platform: The platform type of the migration target. Currently only support: `SQLDB`(Azure SQL Database). Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["service_name"] = service_name
        __props__["source_platform"] = source_platform
        __props__["tags"] = tags
        __props__["target_platform"] = target_platform
        return Project(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

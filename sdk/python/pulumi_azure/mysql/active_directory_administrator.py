# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['ActiveDirectoryAdministratorArgs', 'ActiveDirectoryAdministrator']

@pulumi.input_type
class ActiveDirectoryAdministratorArgs:
    def __init__(__self__, *,
                 login: pulumi.Input[str],
                 object_id: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 server_name: pulumi.Input[str],
                 tenant_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ActiveDirectoryAdministrator resource.
        :param pulumi.Input[str] login: The login name of the principal to set as the server administrator
        :param pulumi.Input[str] object_id: The ID of the principal to set as the server administrator
        :param pulumi.Input[str] resource_group_name: The name of the resource group for the MySQL server. Changing this forces a new resource to be created.
        :param pulumi.Input[str] server_name: The name of the MySQL Server on which to set the administrator. Changing this forces a new resource to be created.
        :param pulumi.Input[str] tenant_id: The Azure Tenant ID
        """
        pulumi.set(__self__, "login", login)
        pulumi.set(__self__, "object_id", object_id)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "server_name", server_name)
        pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def login(self) -> pulumi.Input[str]:
        """
        The login name of the principal to set as the server administrator
        """
        return pulumi.get(self, "login")

    @login.setter
    def login(self, value: pulumi.Input[str]):
        pulumi.set(self, "login", value)

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> pulumi.Input[str]:
        """
        The ID of the principal to set as the server administrator
        """
        return pulumi.get(self, "object_id")

    @object_id.setter
    def object_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "object_id", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group for the MySQL server. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> pulumi.Input[str]:
        """
        The name of the MySQL Server on which to set the administrator. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "server_name")

    @server_name.setter
    def server_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_name", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Input[str]:
        """
        The Azure Tenant ID
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "tenant_id", value)


class ActiveDirectoryAdministrator(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 login: Optional[pulumi.Input[str]] = None,
                 object_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Allows you to set a user or group as the AD administrator for an MySQL server in Azure

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_server = azure.mysql.Server("exampleServer",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            administrator_login="mysqladminun",
            administrator_login_password="H@Sh1CoR3!",
            sku_name="B_Gen5_2",
            storage_mb=5120,
            version="5.7")
        example_active_directory_administrator = azure.mysql.ActiveDirectoryAdministrator("exampleActiveDirectoryAdministrator",
            server_name=example_server.name,
            resource_group_name=example_resource_group.name,
            login="sqladmin",
            tenant_id=current.tenant_id,
            object_id=current.object_id)
        ```

        ## Import

        A MySQL Active Directory Administrator can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:mysql/activeDirectoryAdministrator:ActiveDirectoryAdministrator administrator /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.DBforMySQL/servers/myserver/administrators/activeDirectory
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] login: The login name of the principal to set as the server administrator
        :param pulumi.Input[str] object_id: The ID of the principal to set as the server administrator
        :param pulumi.Input[str] resource_group_name: The name of the resource group for the MySQL server. Changing this forces a new resource to be created.
        :param pulumi.Input[str] server_name: The name of the MySQL Server on which to set the administrator. Changing this forces a new resource to be created.
        :param pulumi.Input[str] tenant_id: The Azure Tenant ID
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ActiveDirectoryAdministratorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows you to set a user or group as the AD administrator for an MySQL server in Azure

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_server = azure.mysql.Server("exampleServer",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            administrator_login="mysqladminun",
            administrator_login_password="H@Sh1CoR3!",
            sku_name="B_Gen5_2",
            storage_mb=5120,
            version="5.7")
        example_active_directory_administrator = azure.mysql.ActiveDirectoryAdministrator("exampleActiveDirectoryAdministrator",
            server_name=example_server.name,
            resource_group_name=example_resource_group.name,
            login="sqladmin",
            tenant_id=current.tenant_id,
            object_id=current.object_id)
        ```

        ## Import

        A MySQL Active Directory Administrator can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:mysql/activeDirectoryAdministrator:ActiveDirectoryAdministrator administrator /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.DBforMySQL/servers/myserver/administrators/activeDirectory
        ```

        :param str resource_name: The name of the resource.
        :param ActiveDirectoryAdministratorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ActiveDirectoryAdministratorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 login: Optional[pulumi.Input[str]] = None,
                 object_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
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

            if login is None and not opts.urn:
                raise TypeError("Missing required property 'login'")
            __props__['login'] = login
            if object_id is None and not opts.urn:
                raise TypeError("Missing required property 'object_id'")
            __props__['object_id'] = object_id
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if server_name is None and not opts.urn:
                raise TypeError("Missing required property 'server_name'")
            __props__['server_name'] = server_name
            if tenant_id is None and not opts.urn:
                raise TypeError("Missing required property 'tenant_id'")
            __props__['tenant_id'] = tenant_id
        super(ActiveDirectoryAdministrator, __self__).__init__(
            'azure:mysql/activeDirectoryAdministrator:ActiveDirectoryAdministrator',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            login: Optional[pulumi.Input[str]] = None,
            object_id: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            server_name: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None) -> 'ActiveDirectoryAdministrator':
        """
        Get an existing ActiveDirectoryAdministrator resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] login: The login name of the principal to set as the server administrator
        :param pulumi.Input[str] object_id: The ID of the principal to set as the server administrator
        :param pulumi.Input[str] resource_group_name: The name of the resource group for the MySQL server. Changing this forces a new resource to be created.
        :param pulumi.Input[str] server_name: The name of the MySQL Server on which to set the administrator. Changing this forces a new resource to be created.
        :param pulumi.Input[str] tenant_id: The Azure Tenant ID
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["login"] = login
        __props__["object_id"] = object_id
        __props__["resource_group_name"] = resource_group_name
        __props__["server_name"] = server_name
        __props__["tenant_id"] = tenant_id
        return ActiveDirectoryAdministrator(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def login(self) -> pulumi.Output[str]:
        """
        The login name of the principal to set as the server administrator
        """
        return pulumi.get(self, "login")

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> pulumi.Output[str]:
        """
        The ID of the principal to set as the server administrator
        """
        return pulumi.get(self, "object_id")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group for the MySQL server. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> pulumi.Output[str]:
        """
        The name of the MySQL Server on which to set the administrator. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "server_name")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The Azure Tenant ID
        """
        return pulumi.get(self, "tenant_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'DatasetAzureBlobSchemaColumnArgs',
    'DatasetCosmosDBApiSchemaColumnArgs',
    'DatasetDelimitedTextAzureBlobStorageLocationArgs',
    'DatasetDelimitedTextHttpServerLocationArgs',
    'DatasetDelimitedTextSchemaColumnArgs',
    'DatasetHttpSchemaColumnArgs',
    'DatasetJsonAzureBlobStorageLocationArgs',
    'DatasetJsonHttpServerLocationArgs',
    'DatasetJsonSchemaColumnArgs',
    'DatasetMysqlSchemaColumnArgs',
    'DatasetPostgresqlSchemaColumnArgs',
    'DatasetSqlServerTableSchemaColumnArgs',
    'FactoryGithubConfigurationArgs',
    'FactoryIdentityArgs',
    'FactoryVstsConfigurationArgs',
    'IntegrationRuntimeManagedCatalogInfoArgs',
    'IntegrationRuntimeManagedCustomSetupScriptArgs',
    'IntegrationRuntimeManagedVnetIntegrationArgs',
    'IntegrationRuntimeSelfHostedRbacAuthorizationArgs',
]

@pulumi.input_type
class DatasetAzureBlobSchemaColumnArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: The name of the column.
        :param pulumi.Input[str] description: The description of the column.
        :param pulumi.Input[str] type: Type of the column. Valid values are `Byte`, `Byte[]`, `Boolean`, `Date`, `DateTime`,`DateTimeOffset`, `Decimal`, `Double`, `Guid`, `Int16`, `Int32`, `Int64`, `Single`, `String`, `TimeSpan`. Please note these values are case sensitive.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the column.
        """
        ...

    @name.setter
    def name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the column.
        """
        ...

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the column. Valid values are `Byte`, `Byte[]`, `Boolean`, `Date`, `DateTime`,`DateTimeOffset`, `Decimal`, `Double`, `Guid`, `Int16`, `Int32`, `Int64`, `Single`, `String`, `TimeSpan`. Please note these values are case sensitive.
        """
        ...

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class DatasetCosmosDBApiSchemaColumnArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: The name of the column.
        :param pulumi.Input[str] description: The description of the column.
        :param pulumi.Input[str] type: Type of the column. Valid values are `Byte`, `Byte[]`, `Boolean`, `Date`, `DateTime`,`DateTimeOffset`, `Decimal`, `Double`, `Guid`, `Int16`, `Int32`, `Int64`, `Single`, `String`, `TimeSpan`. Please note these values are case sensitive.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the column.
        """
        ...

    @name.setter
    def name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the column.
        """
        ...

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the column. Valid values are `Byte`, `Byte[]`, `Boolean`, `Date`, `DateTime`,`DateTimeOffset`, `Decimal`, `Double`, `Guid`, `Int16`, `Int32`, `Int64`, `Single`, `String`, `TimeSpan`. Please note these values are case sensitive.
        """
        ...

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class DatasetDelimitedTextAzureBlobStorageLocationArgs:
    def __init__(__self__, *,
                 container: pulumi.Input[str],
                 filename: pulumi.Input[str],
                 path: pulumi.Input[str]):
        """
        :param pulumi.Input[str] container: The container on the Azure Blob Storage Account hosting the file.
        :param pulumi.Input[str] filename: The filename of the file on the web server.
        :param pulumi.Input[str] path: The folder path to the file on the web server.
        """
        pulumi.set(__self__, "container", container)
        pulumi.set(__self__, "filename", filename)
        pulumi.set(__self__, "path", path)

    @property
    @pulumi.getter
    def container(self) -> pulumi.Input[str]:
        """
        The container on the Azure Blob Storage Account hosting the file.
        """
        ...

    @container.setter
    def container(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def filename(self) -> pulumi.Input[str]:
        """
        The filename of the file on the web server.
        """
        ...

    @filename.setter
    def filename(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        """
        The folder path to the file on the web server.
        """
        ...

    @path.setter
    def path(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class DatasetDelimitedTextHttpServerLocationArgs:
    def __init__(__self__, *,
                 filename: pulumi.Input[str],
                 path: pulumi.Input[str],
                 relative_url: pulumi.Input[str]):
        """
        :param pulumi.Input[str] filename: The filename of the file on the web server.
        :param pulumi.Input[str] path: The folder path to the file on the web server.
        :param pulumi.Input[str] relative_url: The base URL to the web server hosting the file.
        """
        pulumi.set(__self__, "filename", filename)
        pulumi.set(__self__, "path", path)
        pulumi.set(__self__, "relativeUrl", relative_url)

    @property
    @pulumi.getter
    def filename(self) -> pulumi.Input[str]:
        """
        The filename of the file on the web server.
        """
        ...

    @filename.setter
    def filename(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        """
        The folder path to the file on the web server.
        """
        ...

    @path.setter
    def path(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="relativeUrl")
    def relative_url(self) -> pulumi.Input[str]:
        """
        The base URL to the web server hosting the file.
        """
        ...

    @relative_url.setter
    def relative_url(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class DatasetDelimitedTextSchemaColumnArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: The name of the column.
        :param pulumi.Input[str] description: The description of the column.
        :param pulumi.Input[str] type: Type of the column. Valid values are `Byte`, `Byte[]`, `Boolean`, `Date`, `DateTime`,`DateTimeOffset`, `Decimal`, `Double`, `Guid`, `Int16`, `Int32`, `Int64`, `Single`, `String`, `TimeSpan`. Please note these values are case sensitive.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the column.
        """
        ...

    @name.setter
    def name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the column.
        """
        ...

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the column. Valid values are `Byte`, `Byte[]`, `Boolean`, `Date`, `DateTime`,`DateTimeOffset`, `Decimal`, `Double`, `Guid`, `Int16`, `Int32`, `Int64`, `Single`, `String`, `TimeSpan`. Please note these values are case sensitive.
        """
        ...

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class DatasetHttpSchemaColumnArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: The name of the column.
        :param pulumi.Input[str] description: The description of the column.
        :param pulumi.Input[str] type: Type of the column. Valid values are `Byte`, `Byte[]`, `Boolean`, `Date`, `DateTime`,`DateTimeOffset`, `Decimal`, `Double`, `Guid`, `Int16`, `Int32`, `Int64`, `Single`, `String`, `TimeSpan`. Please note these values are case sensitive.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the column.
        """
        ...

    @name.setter
    def name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the column.
        """
        ...

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the column. Valid values are `Byte`, `Byte[]`, `Boolean`, `Date`, `DateTime`,`DateTimeOffset`, `Decimal`, `Double`, `Guid`, `Int16`, `Int32`, `Int64`, `Single`, `String`, `TimeSpan`. Please note these values are case sensitive.
        """
        ...

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class DatasetJsonAzureBlobStorageLocationArgs:
    def __init__(__self__, *,
                 container: pulumi.Input[str],
                 filename: pulumi.Input[str],
                 path: pulumi.Input[str]):
        """
        :param pulumi.Input[str] filename: The filename of the file on the web server.
        :param pulumi.Input[str] path: The folder path to the file on the web server.
        """
        pulumi.set(__self__, "container", container)
        pulumi.set(__self__, "filename", filename)
        pulumi.set(__self__, "path", path)

    @property
    @pulumi.getter
    def container(self) -> pulumi.Input[str]:
        ...

    @container.setter
    def container(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def filename(self) -> pulumi.Input[str]:
        """
        The filename of the file on the web server.
        """
        ...

    @filename.setter
    def filename(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        """
        The folder path to the file on the web server.
        """
        ...

    @path.setter
    def path(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class DatasetJsonHttpServerLocationArgs:
    def __init__(__self__, *,
                 filename: pulumi.Input[str],
                 path: pulumi.Input[str],
                 relative_url: pulumi.Input[str]):
        """
        :param pulumi.Input[str] filename: The filename of the file on the web server.
        :param pulumi.Input[str] path: The folder path to the file on the web server.
        :param pulumi.Input[str] relative_url: The base URL to the web server hosting the file.
        """
        pulumi.set(__self__, "filename", filename)
        pulumi.set(__self__, "path", path)
        pulumi.set(__self__, "relativeUrl", relative_url)

    @property
    @pulumi.getter
    def filename(self) -> pulumi.Input[str]:
        """
        The filename of the file on the web server.
        """
        ...

    @filename.setter
    def filename(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        """
        The folder path to the file on the web server.
        """
        ...

    @path.setter
    def path(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="relativeUrl")
    def relative_url(self) -> pulumi.Input[str]:
        """
        The base URL to the web server hosting the file.
        """
        ...

    @relative_url.setter
    def relative_url(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class DatasetJsonSchemaColumnArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: The name of the column.
        :param pulumi.Input[str] description: The description of the column.
        :param pulumi.Input[str] type: Type of the column. Valid values are `Byte`, `Byte[]`, `Boolean`, `Date`, `DateTime`,`DateTimeOffset`, `Decimal`, `Double`, `Guid`, `Int16`, `Int32`, `Int64`, `Single`, `String`, `TimeSpan`. Please note these values are case sensitive.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the column.
        """
        ...

    @name.setter
    def name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the column.
        """
        ...

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the column. Valid values are `Byte`, `Byte[]`, `Boolean`, `Date`, `DateTime`,`DateTimeOffset`, `Decimal`, `Double`, `Guid`, `Int16`, `Int32`, `Int64`, `Single`, `String`, `TimeSpan`. Please note these values are case sensitive.
        """
        ...

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class DatasetMysqlSchemaColumnArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: The name of the column.
        :param pulumi.Input[str] description: The description of the column.
        :param pulumi.Input[str] type: Type of the column. Valid values are `Byte`, `Byte[]`, `Boolean`, `Date`, `DateTime`,`DateTimeOffset`, `Decimal`, `Double`, `Guid`, `Int16`, `Int32`, `Int64`, `Single`, `String`, `TimeSpan`. Please note these values are case sensitive.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the column.
        """
        ...

    @name.setter
    def name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the column.
        """
        ...

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the column. Valid values are `Byte`, `Byte[]`, `Boolean`, `Date`, `DateTime`,`DateTimeOffset`, `Decimal`, `Double`, `Guid`, `Int16`, `Int32`, `Int64`, `Single`, `String`, `TimeSpan`. Please note these values are case sensitive.
        """
        ...

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class DatasetPostgresqlSchemaColumnArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: The name of the column.
        :param pulumi.Input[str] description: The description of the column.
        :param pulumi.Input[str] type: Type of the column. Valid values are `Byte`, `Byte[]`, `Boolean`, `Date`, `DateTime`,`DateTimeOffset`, `Decimal`, `Double`, `Guid`, `Int16`, `Int32`, `Int64`, `Single`, `String`, `TimeSpan`. Please note these values are case sensitive.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the column.
        """
        ...

    @name.setter
    def name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the column.
        """
        ...

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the column. Valid values are `Byte`, `Byte[]`, `Boolean`, `Date`, `DateTime`,`DateTimeOffset`, `Decimal`, `Double`, `Guid`, `Int16`, `Int32`, `Int64`, `Single`, `String`, `TimeSpan`. Please note these values are case sensitive.
        """
        ...

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class DatasetSqlServerTableSchemaColumnArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: The name of the column.
        :param pulumi.Input[str] description: The description of the column.
        :param pulumi.Input[str] type: Type of the column. Valid values are `Byte`, `Byte[]`, `Boolean`, `Date`, `DateTime`,`DateTimeOffset`, `Decimal`, `Double`, `Guid`, `Int16`, `Int32`, `Int64`, `Single`, `String`, `TimeSpan`. Please note these values are case sensitive.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the column.
        """
        ...

    @name.setter
    def name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the column.
        """
        ...

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the column. Valid values are `Byte`, `Byte[]`, `Boolean`, `Date`, `DateTime`,`DateTimeOffset`, `Decimal`, `Double`, `Guid`, `Int16`, `Int32`, `Int64`, `Single`, `String`, `TimeSpan`. Please note these values are case sensitive.
        """
        ...

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class FactoryGithubConfigurationArgs:
    def __init__(__self__, *,
                 account_name: pulumi.Input[str],
                 branch_name: pulumi.Input[str],
                 git_url: pulumi.Input[str],
                 repository_name: pulumi.Input[str],
                 root_folder: pulumi.Input[str]):
        """
        :param pulumi.Input[str] account_name: Specifies the GitHub account name.
        :param pulumi.Input[str] branch_name: Specifies the branch of the repository to get code from.
        :param pulumi.Input[str] git_url: Specifies the GitHub Enterprise host name. For example: https://github.mydomain.com. Use https://github.com for open source repositories.
        :param pulumi.Input[str] repository_name: Specifies the name of the git repository.
        :param pulumi.Input[str] root_folder: Specifies the root folder within the repository. Set to `/` for the top level.
        """
        pulumi.set(__self__, "accountName", account_name)
        pulumi.set(__self__, "branchName", branch_name)
        pulumi.set(__self__, "gitUrl", git_url)
        pulumi.set(__self__, "repositoryName", repository_name)
        pulumi.set(__self__, "rootFolder", root_folder)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Input[str]:
        """
        Specifies the GitHub account name.
        """
        ...

    @account_name.setter
    def account_name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="branchName")
    def branch_name(self) -> pulumi.Input[str]:
        """
        Specifies the branch of the repository to get code from.
        """
        ...

    @branch_name.setter
    def branch_name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="gitUrl")
    def git_url(self) -> pulumi.Input[str]:
        """
        Specifies the GitHub Enterprise host name. For example: https://github.mydomain.com. Use https://github.com for open source repositories.
        """
        ...

    @git_url.setter
    def git_url(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="repositoryName")
    def repository_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the git repository.
        """
        ...

    @repository_name.setter
    def repository_name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="rootFolder")
    def root_folder(self) -> pulumi.Input[str]:
        """
        Specifies the root folder within the repository. Set to `/` for the top level.
        """
        ...

    @root_folder.setter
    def root_folder(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class FactoryIdentityArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 principal_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: Specifies the identity type of the Data Factory. At this time the only allowed value is `SystemAssigned`.
        :param pulumi.Input[str] principal_id: The ID of the Principal (Client) in Azure Active Directory
        :param pulumi.Input[str] tenant_id: Specifies the Tenant ID associated with the VSTS account.
        """
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "principalId", principal_id)
        pulumi.set(__self__, "tenantId", tenant_id)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Specifies the identity type of the Data Factory. At this time the only allowed value is `SystemAssigned`.
        """
        ...

    @type.setter
    def type(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Principal (Client) in Azure Active Directory
        """
        ...

    @principal_id.setter
    def principal_id(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the Tenant ID associated with the VSTS account.
        """
        ...

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class FactoryVstsConfigurationArgs:
    def __init__(__self__, *,
                 account_name: pulumi.Input[str],
                 branch_name: pulumi.Input[str],
                 project_name: pulumi.Input[str],
                 repository_name: pulumi.Input[str],
                 root_folder: pulumi.Input[str],
                 tenant_id: pulumi.Input[str]):
        """
        :param pulumi.Input[str] account_name: Specifies the VSTS account name.
        :param pulumi.Input[str] branch_name: Specifies the branch of the repository to get code from.
        :param pulumi.Input[str] project_name: Specifies the name of the VSTS project.
        :param pulumi.Input[str] repository_name: Specifies the name of the git repository.
        :param pulumi.Input[str] root_folder: Specifies the root folder within the repository. Set to `/` for the top level.
        :param pulumi.Input[str] tenant_id: Specifies the Tenant ID associated with the VSTS account.
        """
        pulumi.set(__self__, "accountName", account_name)
        pulumi.set(__self__, "branchName", branch_name)
        pulumi.set(__self__, "projectName", project_name)
        pulumi.set(__self__, "repositoryName", repository_name)
        pulumi.set(__self__, "rootFolder", root_folder)
        pulumi.set(__self__, "tenantId", tenant_id)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Input[str]:
        """
        Specifies the VSTS account name.
        """
        ...

    @account_name.setter
    def account_name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="branchName")
    def branch_name(self) -> pulumi.Input[str]:
        """
        Specifies the branch of the repository to get code from.
        """
        ...

    @branch_name.setter
    def branch_name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the VSTS project.
        """
        ...

    @project_name.setter
    def project_name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="repositoryName")
    def repository_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the git repository.
        """
        ...

    @repository_name.setter
    def repository_name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="rootFolder")
    def root_folder(self) -> pulumi.Input[str]:
        """
        Specifies the root folder within the repository. Set to `/` for the top level.
        """
        ...

    @root_folder.setter
    def root_folder(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Input[str]:
        """
        Specifies the Tenant ID associated with the VSTS account.
        """
        ...

    @tenant_id.setter
    def tenant_id(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class IntegrationRuntimeManagedCatalogInfoArgs:
    def __init__(__self__, *,
                 administrator_login: pulumi.Input[str],
                 administrator_password: pulumi.Input[str],
                 server_endpoint: pulumi.Input[str],
                 pricing_tier: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] administrator_login: Administrator login name for the SQL Server.
        :param pulumi.Input[str] administrator_password: Administrator login password for the SQL Server.
        :param pulumi.Input[str] server_endpoint: The endpoint of an Azure SQL Server that will be used to host the SSIS catalog.
        :param pulumi.Input[str] pricing_tier: Pricing tier for the database that will be created for the SSIS catalog. Valid values are: `Basic`, `Standard`, `Premium` and `PremiumRS`.
        """
        pulumi.set(__self__, "administratorLogin", administrator_login)
        pulumi.set(__self__, "administratorPassword", administrator_password)
        pulumi.set(__self__, "serverEndpoint", server_endpoint)
        pulumi.set(__self__, "pricingTier", pricing_tier)

    @property
    @pulumi.getter(name="administratorLogin")
    def administrator_login(self) -> pulumi.Input[str]:
        """
        Administrator login name for the SQL Server.
        """
        ...

    @administrator_login.setter
    def administrator_login(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="administratorPassword")
    def administrator_password(self) -> pulumi.Input[str]:
        """
        Administrator login password for the SQL Server.
        """
        ...

    @administrator_password.setter
    def administrator_password(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="serverEndpoint")
    def server_endpoint(self) -> pulumi.Input[str]:
        """
        The endpoint of an Azure SQL Server that will be used to host the SSIS catalog.
        """
        ...

    @server_endpoint.setter
    def server_endpoint(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="pricingTier")
    def pricing_tier(self) -> Optional[pulumi.Input[str]]:
        """
        Pricing tier for the database that will be created for the SSIS catalog. Valid values are: `Basic`, `Standard`, `Premium` and `PremiumRS`.
        """
        ...

    @pricing_tier.setter
    def pricing_tier(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class IntegrationRuntimeManagedCustomSetupScriptArgs:
    def __init__(__self__, *,
                 blob_container_uri: pulumi.Input[str],
                 sas_token: pulumi.Input[str]):
        """
        :param pulumi.Input[str] blob_container_uri: The blob endpoint for the container which contains a custom setup script that will be run on every node on startup. See [https://docs.microsoft.com/en-us/azure/data-factory/how-to-configure-azure-ssis-ir-custom-setup](https://docs.microsoft.com/en-us/azure/data-factory/how-to-configure-azure-ssis-ir-custom-setup) for more information.
        :param pulumi.Input[str] sas_token: A container SAS token that gives access to the files. See [https://docs.microsoft.com/en-us/azure/data-factory/how-to-configure-azure-ssis-ir-custom-setup](https://docs.microsoft.com/en-us/azure/data-factory/how-to-configure-azure-ssis-ir-custom-setup) for more information.
        """
        pulumi.set(__self__, "blobContainerUri", blob_container_uri)
        pulumi.set(__self__, "sasToken", sas_token)

    @property
    @pulumi.getter(name="blobContainerUri")
    def blob_container_uri(self) -> pulumi.Input[str]:
        """
        The blob endpoint for the container which contains a custom setup script that will be run on every node on startup. See [https://docs.microsoft.com/en-us/azure/data-factory/how-to-configure-azure-ssis-ir-custom-setup](https://docs.microsoft.com/en-us/azure/data-factory/how-to-configure-azure-ssis-ir-custom-setup) for more information.
        """
        ...

    @blob_container_uri.setter
    def blob_container_uri(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="sasToken")
    def sas_token(self) -> pulumi.Input[str]:
        """
        A container SAS token that gives access to the files. See [https://docs.microsoft.com/en-us/azure/data-factory/how-to-configure-azure-ssis-ir-custom-setup](https://docs.microsoft.com/en-us/azure/data-factory/how-to-configure-azure-ssis-ir-custom-setup) for more information.
        """
        ...

    @sas_token.setter
    def sas_token(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class IntegrationRuntimeManagedVnetIntegrationArgs:
    def __init__(__self__, *,
                 subnet_name: pulumi.Input[str],
                 vnet_id: pulumi.Input[str]):
        """
        :param pulumi.Input[str] subnet_name: Name of the subnet to which the nodes of the Managed Integration Runtime will be added.
        :param pulumi.Input[str] vnet_id: ID of the virtual network to which the nodes of the Managed Integration Runtime will be added.
        """
        pulumi.set(__self__, "subnetName", subnet_name)
        pulumi.set(__self__, "vnetId", vnet_id)

    @property
    @pulumi.getter(name="subnetName")
    def subnet_name(self) -> pulumi.Input[str]:
        """
        Name of the subnet to which the nodes of the Managed Integration Runtime will be added.
        """
        ...

    @subnet_name.setter
    def subnet_name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="vnetId")
    def vnet_id(self) -> pulumi.Input[str]:
        """
        ID of the virtual network to which the nodes of the Managed Integration Runtime will be added.
        """
        ...

    @vnet_id.setter
    def vnet_id(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class IntegrationRuntimeSelfHostedRbacAuthorizationArgs:
    def __init__(__self__, *,
                 resource_id: pulumi.Input[str]):
        """
        :param pulumi.Input[str] resource_id: The resource identifier of the integration runtime to be shared. Changing this forces a new Data Factory to be created.
        """
        pulumi.set(__self__, "resourceId", resource_id)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        The resource identifier of the integration runtime to be shared. Changing this forces a new Data Factory to be created.
        """
        ...

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        ...



# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'AccountActiveDirectory',
    'VolumeExportPolicyRule',
]

@pulumi.output_type
class AccountActiveDirectory(dict):
    def __init__(__self__, *,
                 dns_servers: List[str],
                 domain: str,
                 password: str,
                 smb_server_name: str,
                 username: str,
                 organizational_unit: Optional[str] = None):
        """
        :param List[str] dns_servers: A list of DNS server IP addresses for the Active Directory domain. Only allows `IPv4` address.
        :param str domain: The name of the Active Directory domain.
        :param str password: The password associated with the `username`.
        :param str smb_server_name: The NetBIOS name which should be used for the NetApp SMB Server, which will be registered as a computer account in the AD and used to mount volumes.
        :param str username: The Username of Active Directory Domain Administrator.
        :param str organizational_unit: The Organizational Unit (OU) within the Active Directory Domain.
        """
        pulumi.set(__self__, "dns_servers", dns_servers)
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "smb_server_name", smb_server_name)
        pulumi.set(__self__, "username", username)
        if organizational_unit is not None:
            pulumi.set(__self__, "organizational_unit", organizational_unit)

    @property
    @pulumi.getter(name="dnsServers")
    def dns_servers(self) -> List[str]:
        """
        A list of DNS server IP addresses for the Active Directory domain. Only allows `IPv4` address.
        """
        return pulumi.get(self, "dns_servers")

    @property
    @pulumi.getter
    def domain(self) -> str:
        """
        The name of the Active Directory domain.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def password(self) -> str:
        """
        The password associated with the `username`.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="smbServerName")
    def smb_server_name(self) -> str:
        """
        The NetBIOS name which should be used for the NetApp SMB Server, which will be registered as a computer account in the AD and used to mount volumes.
        """
        return pulumi.get(self, "smb_server_name")

    @property
    @pulumi.getter
    def username(self) -> str:
        """
        The Username of Active Directory Domain Administrator.
        """
        return pulumi.get(self, "username")

    @property
    @pulumi.getter(name="organizationalUnit")
    def organizational_unit(self) -> Optional[str]:
        """
        The Organizational Unit (OU) within the Active Directory Domain.
        """
        return pulumi.get(self, "organizational_unit")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class VolumeExportPolicyRule(dict):
    def __init__(__self__, *,
                 allowed_clients: List[str],
                 rule_index: int,
                 cifs_enabled: Optional[bool] = None,
                 nfsv3_enabled: Optional[bool] = None,
                 nfsv4_enabled: Optional[bool] = None,
                 protocols_enabled: Optional[str] = None,
                 unix_read_only: Optional[bool] = None,
                 unix_read_write: Optional[bool] = None):
        """
        :param List[str] allowed_clients: A list of allowed clients IPv4 addresses.
        :param int rule_index: The index number of the rule.
        :param bool cifs_enabled: Is the CIFS protocol allowed?
        :param bool nfsv3_enabled: Is the NFSv3 protocol allowed?
        :param bool nfsv4_enabled: Is the NFSv4 protocol allowed?
        :param str protocols_enabled: A list of allowed protocols. Valid values include `CIFS`, `NFSv3`, or `NFSv4.1`. Only one value is supported at this time. This replaces the previous arguments: `cifs_enabled`, `nfsv3_enabled` and `nfsv4_enabled`.
        :param bool unix_read_only: Is the file system on unix read only?
        :param bool unix_read_write: Is the file system on unix read and write?
        """
        pulumi.set(__self__, "allowed_clients", allowed_clients)
        pulumi.set(__self__, "rule_index", rule_index)
        if cifs_enabled is not None:
            pulumi.set(__self__, "cifs_enabled", cifs_enabled)
        if nfsv3_enabled is not None:
            pulumi.set(__self__, "nfsv3_enabled", nfsv3_enabled)
        if nfsv4_enabled is not None:
            pulumi.set(__self__, "nfsv4_enabled", nfsv4_enabled)
        if protocols_enabled is not None:
            pulumi.set(__self__, "protocols_enabled", protocols_enabled)
        if unix_read_only is not None:
            pulumi.set(__self__, "unix_read_only", unix_read_only)
        if unix_read_write is not None:
            pulumi.set(__self__, "unix_read_write", unix_read_write)

    @property
    @pulumi.getter(name="allowedClients")
    def allowed_clients(self) -> List[str]:
        """
        A list of allowed clients IPv4 addresses.
        """
        return pulumi.get(self, "allowed_clients")

    @property
    @pulumi.getter(name="ruleIndex")
    def rule_index(self) -> int:
        """
        The index number of the rule.
        """
        return pulumi.get(self, "rule_index")

    @property
    @pulumi.getter(name="cifsEnabled")
    def cifs_enabled(self) -> Optional[bool]:
        """
        Is the CIFS protocol allowed?
        """
        return pulumi.get(self, "cifs_enabled")

    @property
    @pulumi.getter(name="nfsv3Enabled")
    def nfsv3_enabled(self) -> Optional[bool]:
        """
        Is the NFSv3 protocol allowed?
        """
        return pulumi.get(self, "nfsv3_enabled")

    @property
    @pulumi.getter(name="nfsv4Enabled")
    def nfsv4_enabled(self) -> Optional[bool]:
        """
        Is the NFSv4 protocol allowed?
        """
        return pulumi.get(self, "nfsv4_enabled")

    @property
    @pulumi.getter(name="protocolsEnabled")
    def protocols_enabled(self) -> Optional[str]:
        """
        A list of allowed protocols. Valid values include `CIFS`, `NFSv3`, or `NFSv4.1`. Only one value is supported at this time. This replaces the previous arguments: `cifs_enabled`, `nfsv3_enabled` and `nfsv4_enabled`.
        """
        return pulumi.get(self, "protocols_enabled")

    @property
    @pulumi.getter(name="unixReadOnly")
    def unix_read_only(self) -> Optional[bool]:
        """
        Is the file system on unix read only?
        """
        return pulumi.get(self, "unix_read_only")

    @property
    @pulumi.getter(name="unixReadWrite")
    def unix_read_write(self) -> Optional[bool]:
        """
        Is the file system on unix read and write?
        """
        return pulumi.get(self, "unix_read_write")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop



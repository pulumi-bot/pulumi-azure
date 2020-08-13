# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'LinkServiceNatIpConfiguration',
    'MxRecordRecord',
    'SRVRecordRecord',
    'TxtRecordRecord',
]

@pulumi.output_type
class LinkServiceNatIpConfiguration(dict):
    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name which should be used for the NAT IP Configuration. Changing this forces a new resource to be created.
        """
        ...

    @property
    @pulumi.getter
    def primary(self) -> bool:
        """
        Is this is the Primary IP Configuration? Changing this forces a new resource to be created.
        """
        ...

    @property
    @pulumi.getter(name="privateIpAddress")
    def private_ip_address(self) -> Optional[str]:
        """
        Specifies a Private Static IP Address for this IP Configuration.
        """
        ...

    @property
    @pulumi.getter(name="privateIpAddressVersion")
    def private_ip_address_version(self) -> Optional[str]:
        """
        The version of the IP Protocol which should be used. At this time the only supported value is `IPv4`. Defaults to `IPv4`.
        """
        ...

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        Specifies the ID of the Subnet which should be used for the Private Link Service.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MxRecordRecord(dict):
    @property
    @pulumi.getter
    def exchange(self) -> str:
        """
        The FQDN of the exchange to MX record points to.
        """
        ...

    @property
    @pulumi.getter
    def preference(self) -> float:
        """
        The preference of the MX record.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SRVRecordRecord(dict):
    @property
    @pulumi.getter
    def port(self) -> float:
        """
        The Port the service is listening on.
        """
        ...

    @property
    @pulumi.getter
    def priority(self) -> float:
        """
        The priority of the SRV record.
        """
        ...

    @property
    @pulumi.getter
    def target(self) -> str:
        """
        The FQDN of the service.
        """
        ...

    @property
    @pulumi.getter
    def weight(self) -> float:
        """
        The Weight of the SRV record.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TxtRecordRecord(dict):
    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value of the TXT record. Max length: 1024 characters
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop



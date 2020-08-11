# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'ClusterAzureActiveDirectoryArgs',
    'ClusterCertificateArgs',
    'ClusterCertificateCommonNamesArgs',
    'ClusterCertificateCommonNamesCommonNameArgs',
    'ClusterClientCertificateCommonNameArgs',
    'ClusterClientCertificateThumbprintArgs',
    'ClusterDiagnosticsConfigArgs',
    'ClusterFabricSettingArgs',
    'ClusterNodeTypeArgs',
    'ClusterNodeTypeApplicationPortsArgs',
    'ClusterNodeTypeEphemeralPortsArgs',
    'ClusterReverseProxyCertificateArgs',
]

@pulumi.input_type
class ClusterAzureActiveDirectoryArgs:
    def __init__(__self__, *,
                 client_application_id: pulumi.Input[str],
                 cluster_application_id: pulumi.Input[str],
                 tenant_id: pulumi.Input[str]):
        """
        :param pulumi.Input[str] client_application_id: The Azure Active Directory Client ID which should be used for the Client Application.
        :param pulumi.Input[str] cluster_application_id: The Azure Active Directory Cluster Application ID.
        :param pulumi.Input[str] tenant_id: The Azure Active Directory Tenant ID.
        """
        pulumi.set(__self__, "clientApplicationId", client_application_id)
        pulumi.set(__self__, "clusterApplicationId", cluster_application_id)
        pulumi.set(__self__, "tenantId", tenant_id)

    @property
    @pulumi.getter(name="clientApplicationId")
    def client_application_id(self) -> pulumi.Input[str]:
        """
        The Azure Active Directory Client ID which should be used for the Client Application.
        """
        ...

    @client_application_id.setter
    def client_application_id(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="clusterApplicationId")
    def cluster_application_id(self) -> pulumi.Input[str]:
        """
        The Azure Active Directory Cluster Application ID.
        """
        ...

    @cluster_application_id.setter
    def cluster_application_id(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Input[str]:
        """
        The Azure Active Directory Tenant ID.
        """
        ...

    @tenant_id.setter
    def tenant_id(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class ClusterCertificateArgs:
    def __init__(__self__, *,
                 thumbprint: pulumi.Input[str],
                 x509_store_name: pulumi.Input[str],
                 thumbprint_secondary: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] thumbprint: The Thumbprint of the Certificate.
        :param pulumi.Input[str] x509_store_name: The X509 Store where the Certificate Exists, such as `My`.
        :param pulumi.Input[str] thumbprint_secondary: The Secondary Thumbprint of the Certificate.
        """
        pulumi.set(__self__, "thumbprint", thumbprint)
        pulumi.set(__self__, "x509StoreName", x509_store_name)
        pulumi.set(__self__, "thumbprintSecondary", thumbprint_secondary)

    @property
    @pulumi.getter
    def thumbprint(self) -> pulumi.Input[str]:
        """
        The Thumbprint of the Certificate.
        """
        ...

    @thumbprint.setter
    def thumbprint(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="x509StoreName")
    def x509_store_name(self) -> pulumi.Input[str]:
        """
        The X509 Store where the Certificate Exists, such as `My`.
        """
        ...

    @x509_store_name.setter
    def x509_store_name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="thumbprintSecondary")
    def thumbprint_secondary(self) -> Optional[pulumi.Input[str]]:
        """
        The Secondary Thumbprint of the Certificate.
        """
        ...

    @thumbprint_secondary.setter
    def thumbprint_secondary(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class ClusterCertificateCommonNamesArgs:
    def __init__(__self__, *,
                 common_names: pulumi.Input[List[pulumi.Input['ClusterCertificateCommonNamesCommonNameArgs']]],
                 x509_store_name: pulumi.Input[str]):
        """
        :param pulumi.Input[List[pulumi.Input['ClusterCertificateCommonNamesCommonNameArgs']]] common_names: A `common_names` block as defined below.
        :param pulumi.Input[str] x509_store_name: The X509 Store where the Certificate Exists, such as `My`.
        """
        pulumi.set(__self__, "commonNames", common_names)
        pulumi.set(__self__, "x509StoreName", x509_store_name)

    @property
    @pulumi.getter(name="commonNames")
    def common_names(self) -> pulumi.Input[List[pulumi.Input['ClusterCertificateCommonNamesCommonNameArgs']]]:
        """
        A `common_names` block as defined below.
        """
        ...

    @common_names.setter
    def common_names(self, value: pulumi.Input[List[pulumi.Input['ClusterCertificateCommonNamesCommonNameArgs']]]):
        ...

    @property
    @pulumi.getter(name="x509StoreName")
    def x509_store_name(self) -> pulumi.Input[str]:
        """
        The X509 Store where the Certificate Exists, such as `My`.
        """
        ...

    @x509_store_name.setter
    def x509_store_name(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class ClusterCertificateCommonNamesCommonNameArgs:
    def __init__(__self__, *,
                 certificate_common_name: pulumi.Input[str],
                 certificate_issuer_thumbprint: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] certificate_common_name: The common or subject name of the certificate.
        :param pulumi.Input[str] certificate_issuer_thumbprint: The Issuer Thumbprint of the Certificate.
        """
        pulumi.set(__self__, "certificateCommonName", certificate_common_name)
        pulumi.set(__self__, "certificateIssuerThumbprint", certificate_issuer_thumbprint)

    @property
    @pulumi.getter(name="certificateCommonName")
    def certificate_common_name(self) -> pulumi.Input[str]:
        """
        The common or subject name of the certificate.
        """
        ...

    @certificate_common_name.setter
    def certificate_common_name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="certificateIssuerThumbprint")
    def certificate_issuer_thumbprint(self) -> Optional[pulumi.Input[str]]:
        """
        The Issuer Thumbprint of the Certificate.
        """
        ...

    @certificate_issuer_thumbprint.setter
    def certificate_issuer_thumbprint(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class ClusterClientCertificateCommonNameArgs:
    def __init__(__self__, *,
                 common_name: pulumi.Input[str],
                 is_admin: pulumi.Input[bool],
                 issuer_thumbprint: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] is_admin: Does the Client Certificate have Admin Access to the cluster? Non-admin clients can only perform read only operations on the cluster.
        """
        pulumi.set(__self__, "commonName", common_name)
        pulumi.set(__self__, "isAdmin", is_admin)
        pulumi.set(__self__, "issuerThumbprint", issuer_thumbprint)

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> pulumi.Input[str]:
        ...

    @common_name.setter
    def common_name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="isAdmin")
    def is_admin(self) -> pulumi.Input[bool]:
        """
        Does the Client Certificate have Admin Access to the cluster? Non-admin clients can only perform read only operations on the cluster.
        """
        ...

    @is_admin.setter
    def is_admin(self, value: pulumi.Input[bool]):
        ...

    @property
    @pulumi.getter(name="issuerThumbprint")
    def issuer_thumbprint(self) -> Optional[pulumi.Input[str]]:
        ...

    @issuer_thumbprint.setter
    def issuer_thumbprint(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class ClusterClientCertificateThumbprintArgs:
    def __init__(__self__, *,
                 is_admin: pulumi.Input[bool],
                 thumbprint: pulumi.Input[str]):
        """
        :param pulumi.Input[bool] is_admin: Does the Client Certificate have Admin Access to the cluster? Non-admin clients can only perform read only operations on the cluster.
        :param pulumi.Input[str] thumbprint: The Thumbprint associated with the Client Certificate.
        """
        pulumi.set(__self__, "isAdmin", is_admin)
        pulumi.set(__self__, "thumbprint", thumbprint)

    @property
    @pulumi.getter(name="isAdmin")
    def is_admin(self) -> pulumi.Input[bool]:
        """
        Does the Client Certificate have Admin Access to the cluster? Non-admin clients can only perform read only operations on the cluster.
        """
        ...

    @is_admin.setter
    def is_admin(self, value: pulumi.Input[bool]):
        ...

    @property
    @pulumi.getter
    def thumbprint(self) -> pulumi.Input[str]:
        """
        The Thumbprint associated with the Client Certificate.
        """
        ...

    @thumbprint.setter
    def thumbprint(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class ClusterDiagnosticsConfigArgs:
    def __init__(__self__, *,
                 blob_endpoint: pulumi.Input[str],
                 protected_account_key_name: pulumi.Input[str],
                 queue_endpoint: pulumi.Input[str],
                 storage_account_name: pulumi.Input[str],
                 table_endpoint: pulumi.Input[str]):
        """
        :param pulumi.Input[str] blob_endpoint: The Blob Endpoint of the Storage Account.
        :param pulumi.Input[str] protected_account_key_name: The protected diagnostics storage key name, such as `StorageAccountKey1`.
        :param pulumi.Input[str] queue_endpoint: The Queue Endpoint of the Storage Account.
        :param pulumi.Input[str] storage_account_name: The name of the Storage Account where the Diagnostics should be sent to.
        :param pulumi.Input[str] table_endpoint: The Table Endpoint of the Storage Account.
        """
        pulumi.set(__self__, "blobEndpoint", blob_endpoint)
        pulumi.set(__self__, "protectedAccountKeyName", protected_account_key_name)
        pulumi.set(__self__, "queueEndpoint", queue_endpoint)
        pulumi.set(__self__, "storageAccountName", storage_account_name)
        pulumi.set(__self__, "tableEndpoint", table_endpoint)

    @property
    @pulumi.getter(name="blobEndpoint")
    def blob_endpoint(self) -> pulumi.Input[str]:
        """
        The Blob Endpoint of the Storage Account.
        """
        ...

    @blob_endpoint.setter
    def blob_endpoint(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="protectedAccountKeyName")
    def protected_account_key_name(self) -> pulumi.Input[str]:
        """
        The protected diagnostics storage key name, such as `StorageAccountKey1`.
        """
        ...

    @protected_account_key_name.setter
    def protected_account_key_name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="queueEndpoint")
    def queue_endpoint(self) -> pulumi.Input[str]:
        """
        The Queue Endpoint of the Storage Account.
        """
        ...

    @queue_endpoint.setter
    def queue_endpoint(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="storageAccountName")
    def storage_account_name(self) -> pulumi.Input[str]:
        """
        The name of the Storage Account where the Diagnostics should be sent to.
        """
        ...

    @storage_account_name.setter
    def storage_account_name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="tableEndpoint")
    def table_endpoint(self) -> pulumi.Input[str]:
        """
        The Table Endpoint of the Storage Account.
        """
        ...

    @table_endpoint.setter
    def table_endpoint(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class ClusterFabricSettingArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] name: The name of the Fabric Setting, such as `Security` or `Federation`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map containing settings for the specified Fabric Setting.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "parameters", parameters)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the Fabric Setting, such as `Security` or `Federation`.
        """
        ...

    @name.setter
    def name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map containing settings for the specified Fabric Setting.
        """
        ...

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        ...


@pulumi.input_type
class ClusterNodeTypeArgs:
    def __init__(__self__, *,
                 client_endpoint_port: pulumi.Input[float],
                 http_endpoint_port: pulumi.Input[float],
                 instance_count: pulumi.Input[float],
                 is_primary: pulumi.Input[bool],
                 name: pulumi.Input[str],
                 application_ports: Optional[pulumi.Input['ClusterNodeTypeApplicationPortsArgs']] = None,
                 capacities: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 durability_level: Optional[pulumi.Input[str]] = None,
                 ephemeral_ports: Optional[pulumi.Input['ClusterNodeTypeEphemeralPortsArgs']] = None,
                 placement_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 reverse_proxy_endpoint_port: Optional[pulumi.Input[float]] = None):
        """
        :param pulumi.Input[float] client_endpoint_port: The Port used for the Client Endpoint for this Node Type. Changing this forces a new resource to be created.
        :param pulumi.Input[float] http_endpoint_port: The Port used for the HTTP Endpoint for this Node Type. Changing this forces a new resource to be created.
        :param pulumi.Input[float] instance_count: The number of nodes for this Node Type.
        :param pulumi.Input[bool] is_primary: Is this the Primary Node Type? Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Node Type. Changing this forces a new resource to be created.
        :param pulumi.Input['ClusterNodeTypeApplicationPortsArgs'] application_ports: A `application_ports` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] capacities: The capacity tags applied to the nodes in the node type, the cluster resource manager uses these tags to understand how much resource a node has.
        :param pulumi.Input[str] durability_level: The Durability Level for this Node Type. Possible values include `Bronze`, `Gold` and `Silver`. Defaults to `Bronze`. Changing this forces a new resource to be created.
        :param pulumi.Input['ClusterNodeTypeEphemeralPortsArgs'] ephemeral_ports: A `ephemeral_ports` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] placement_properties: The placement tags applied to nodes in the node type, which can be used to indicate where certain services (workload) should run.
        :param pulumi.Input[float] reverse_proxy_endpoint_port: The Port used for the Reverse Proxy Endpoint  for this Node Type. Changing this will upgrade the cluster.
        """
        pulumi.set(__self__, "clientEndpointPort", client_endpoint_port)
        pulumi.set(__self__, "httpEndpointPort", http_endpoint_port)
        pulumi.set(__self__, "instanceCount", instance_count)
        pulumi.set(__self__, "isPrimary", is_primary)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "applicationPorts", application_ports)
        pulumi.set(__self__, "capacities", capacities)
        pulumi.set(__self__, "durabilityLevel", durability_level)
        pulumi.set(__self__, "ephemeralPorts", ephemeral_ports)
        pulumi.set(__self__, "placementProperties", placement_properties)
        pulumi.set(__self__, "reverseProxyEndpointPort", reverse_proxy_endpoint_port)

    @property
    @pulumi.getter(name="clientEndpointPort")
    def client_endpoint_port(self) -> pulumi.Input[float]:
        """
        The Port used for the Client Endpoint for this Node Type. Changing this forces a new resource to be created.
        """
        ...

    @client_endpoint_port.setter
    def client_endpoint_port(self, value: pulumi.Input[float]):
        ...

    @property
    @pulumi.getter(name="httpEndpointPort")
    def http_endpoint_port(self) -> pulumi.Input[float]:
        """
        The Port used for the HTTP Endpoint for this Node Type. Changing this forces a new resource to be created.
        """
        ...

    @http_endpoint_port.setter
    def http_endpoint_port(self, value: pulumi.Input[float]):
        ...

    @property
    @pulumi.getter(name="instanceCount")
    def instance_count(self) -> pulumi.Input[float]:
        """
        The number of nodes for this Node Type.
        """
        ...

    @instance_count.setter
    def instance_count(self, value: pulumi.Input[float]):
        ...

    @property
    @pulumi.getter(name="isPrimary")
    def is_primary(self) -> pulumi.Input[bool]:
        """
        Is this the Primary Node Type? Changing this forces a new resource to be created.
        """
        ...

    @is_primary.setter
    def is_primary(self, value: pulumi.Input[bool]):
        ...

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the Node Type. Changing this forces a new resource to be created.
        """
        ...

    @name.setter
    def name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="applicationPorts")
    def application_ports(self) -> Optional[pulumi.Input['ClusterNodeTypeApplicationPortsArgs']]:
        """
        A `application_ports` block as defined below.
        """
        ...

    @application_ports.setter
    def application_ports(self, value: Optional[pulumi.Input['ClusterNodeTypeApplicationPortsArgs']]):
        ...

    @property
    @pulumi.getter
    def capacities(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The capacity tags applied to the nodes in the node type, the cluster resource manager uses these tags to understand how much resource a node has.
        """
        ...

    @capacities.setter
    def capacities(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        ...

    @property
    @pulumi.getter(name="durabilityLevel")
    def durability_level(self) -> Optional[pulumi.Input[str]]:
        """
        The Durability Level for this Node Type. Possible values include `Bronze`, `Gold` and `Silver`. Defaults to `Bronze`. Changing this forces a new resource to be created.
        """
        ...

    @durability_level.setter
    def durability_level(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter(name="ephemeralPorts")
    def ephemeral_ports(self) -> Optional[pulumi.Input['ClusterNodeTypeEphemeralPortsArgs']]:
        """
        A `ephemeral_ports` block as defined below.
        """
        ...

    @ephemeral_ports.setter
    def ephemeral_ports(self, value: Optional[pulumi.Input['ClusterNodeTypeEphemeralPortsArgs']]):
        ...

    @property
    @pulumi.getter(name="placementProperties")
    def placement_properties(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The placement tags applied to nodes in the node type, which can be used to indicate where certain services (workload) should run.
        """
        ...

    @placement_properties.setter
    def placement_properties(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        ...

    @property
    @pulumi.getter(name="reverseProxyEndpointPort")
    def reverse_proxy_endpoint_port(self) -> Optional[pulumi.Input[float]]:
        """
        The Port used for the Reverse Proxy Endpoint  for this Node Type. Changing this will upgrade the cluster.
        """
        ...

    @reverse_proxy_endpoint_port.setter
    def reverse_proxy_endpoint_port(self, value: Optional[pulumi.Input[float]]):
        ...


@pulumi.input_type
class ClusterNodeTypeApplicationPortsArgs:
    def __init__(__self__, *,
                 end_port: pulumi.Input[float],
                 start_port: pulumi.Input[float]):
        """
        :param pulumi.Input[float] end_port: The end of the Application Port Range on this Node Type.
        :param pulumi.Input[float] start_port: The start of the Application Port Range on this Node Type.
        """
        pulumi.set(__self__, "endPort", end_port)
        pulumi.set(__self__, "startPort", start_port)

    @property
    @pulumi.getter(name="endPort")
    def end_port(self) -> pulumi.Input[float]:
        """
        The end of the Application Port Range on this Node Type.
        """
        ...

    @end_port.setter
    def end_port(self, value: pulumi.Input[float]):
        ...

    @property
    @pulumi.getter(name="startPort")
    def start_port(self) -> pulumi.Input[float]:
        """
        The start of the Application Port Range on this Node Type.
        """
        ...

    @start_port.setter
    def start_port(self, value: pulumi.Input[float]):
        ...


@pulumi.input_type
class ClusterNodeTypeEphemeralPortsArgs:
    def __init__(__self__, *,
                 end_port: pulumi.Input[float],
                 start_port: pulumi.Input[float]):
        """
        :param pulumi.Input[float] end_port: The end of the Ephemeral Port Range on this Node Type.
        :param pulumi.Input[float] start_port: The start of the Ephemeral Port Range on this Node Type.
        """
        pulumi.set(__self__, "endPort", end_port)
        pulumi.set(__self__, "startPort", start_port)

    @property
    @pulumi.getter(name="endPort")
    def end_port(self) -> pulumi.Input[float]:
        """
        The end of the Ephemeral Port Range on this Node Type.
        """
        ...

    @end_port.setter
    def end_port(self, value: pulumi.Input[float]):
        ...

    @property
    @pulumi.getter(name="startPort")
    def start_port(self) -> pulumi.Input[float]:
        """
        The start of the Ephemeral Port Range on this Node Type.
        """
        ...

    @start_port.setter
    def start_port(self, value: pulumi.Input[float]):
        ...


@pulumi.input_type
class ClusterReverseProxyCertificateArgs:
    def __init__(__self__, *,
                 thumbprint: pulumi.Input[str],
                 x509_store_name: pulumi.Input[str],
                 thumbprint_secondary: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] thumbprint: The Thumbprint of the Certificate.
        :param pulumi.Input[str] x509_store_name: The X509 Store where the Certificate Exists, such as `My`.
        :param pulumi.Input[str] thumbprint_secondary: The Secondary Thumbprint of the Certificate.
        """
        pulumi.set(__self__, "thumbprint", thumbprint)
        pulumi.set(__self__, "x509StoreName", x509_store_name)
        pulumi.set(__self__, "thumbprintSecondary", thumbprint_secondary)

    @property
    @pulumi.getter
    def thumbprint(self) -> pulumi.Input[str]:
        """
        The Thumbprint of the Certificate.
        """
        ...

    @thumbprint.setter
    def thumbprint(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="x509StoreName")
    def x509_store_name(self) -> pulumi.Input[str]:
        """
        The X509 Store where the Certificate Exists, such as `My`.
        """
        ...

    @x509_store_name.setter
    def x509_store_name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="thumbprintSecondary")
    def thumbprint_secondary(self) -> Optional[pulumi.Input[str]]:
        """
        The Secondary Thumbprint of the Certificate.
        """
        ...

    @thumbprint_secondary.setter
    def thumbprint_secondary(self, value: Optional[pulumi.Input[str]]):
        ...



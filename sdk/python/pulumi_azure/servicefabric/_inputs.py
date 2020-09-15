# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
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
        pulumi.set(__self__, "client_application_id", client_application_id)
        pulumi.set(__self__, "cluster_application_id", cluster_application_id)
        pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="clientApplicationId")
    def client_application_id(self) -> pulumi.Input[str]:
        """
        The Azure Active Directory Client ID which should be used for the Client Application.
        """
        return pulumi.get(self, "client_application_id")

    @client_application_id.setter
    def client_application_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_application_id", value)

    @property
    @pulumi.getter(name="clusterApplicationId")
    def cluster_application_id(self) -> pulumi.Input[str]:
        """
        The Azure Active Directory Cluster Application ID.
        """
        return pulumi.get(self, "cluster_application_id")

    @cluster_application_id.setter
    def cluster_application_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_application_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Input[str]:
        """
        The Azure Active Directory Tenant ID.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "tenant_id", value)


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
        pulumi.set(__self__, "x509_store_name", x509_store_name)
        if thumbprint_secondary is not None:
            pulumi.set(__self__, "thumbprint_secondary", thumbprint_secondary)

    @property
    @pulumi.getter
    def thumbprint(self) -> pulumi.Input[str]:
        """
        The Thumbprint of the Certificate.
        """
        return pulumi.get(self, "thumbprint")

    @thumbprint.setter
    def thumbprint(self, value: pulumi.Input[str]):
        pulumi.set(self, "thumbprint", value)

    @property
    @pulumi.getter(name="x509StoreName")
    def x509_store_name(self) -> pulumi.Input[str]:
        """
        The X509 Store where the Certificate Exists, such as `My`.
        """
        return pulumi.get(self, "x509_store_name")

    @x509_store_name.setter
    def x509_store_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "x509_store_name", value)

    @property
    @pulumi.getter(name="thumbprintSecondary")
    def thumbprint_secondary(self) -> Optional[pulumi.Input[str]]:
        """
        The Secondary Thumbprint of the Certificate.
        """
        return pulumi.get(self, "thumbprint_secondary")

    @thumbprint_secondary.setter
    def thumbprint_secondary(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "thumbprint_secondary", value)


@pulumi.input_type
class ClusterCertificateCommonNamesArgs:
    def __init__(__self__, *,
                 common_names: pulumi.Input[Sequence[pulumi.Input['ClusterCertificateCommonNamesCommonNameArgs']]],
                 x509_store_name: pulumi.Input[str]):
        """
        :param pulumi.Input[Sequence[pulumi.Input['ClusterCertificateCommonNamesCommonNameArgs']]] common_names: A `common_names` block as defined below.
        :param pulumi.Input[str] x509_store_name: The X509 Store where the Certificate Exists, such as `My`.
        """
        pulumi.set(__self__, "common_names", common_names)
        pulumi.set(__self__, "x509_store_name", x509_store_name)

    @property
    @pulumi.getter(name="commonNames")
    def common_names(self) -> pulumi.Input[Sequence[pulumi.Input['ClusterCertificateCommonNamesCommonNameArgs']]]:
        """
        A `common_names` block as defined below.
        """
        return pulumi.get(self, "common_names")

    @common_names.setter
    def common_names(self, value: pulumi.Input[Sequence[pulumi.Input['ClusterCertificateCommonNamesCommonNameArgs']]]):
        pulumi.set(self, "common_names", value)

    @property
    @pulumi.getter(name="x509StoreName")
    def x509_store_name(self) -> pulumi.Input[str]:
        """
        The X509 Store where the Certificate Exists, such as `My`.
        """
        return pulumi.get(self, "x509_store_name")

    @x509_store_name.setter
    def x509_store_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "x509_store_name", value)


@pulumi.input_type
class ClusterCertificateCommonNamesCommonNameArgs:
    def __init__(__self__, *,
                 certificate_common_name: pulumi.Input[str],
                 certificate_issuer_thumbprint: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] certificate_common_name: The common or subject name of the certificate.
        :param pulumi.Input[str] certificate_issuer_thumbprint: The Issuer Thumbprint of the Certificate.
        """
        pulumi.set(__self__, "certificate_common_name", certificate_common_name)
        if certificate_issuer_thumbprint is not None:
            pulumi.set(__self__, "certificate_issuer_thumbprint", certificate_issuer_thumbprint)

    @property
    @pulumi.getter(name="certificateCommonName")
    def certificate_common_name(self) -> pulumi.Input[str]:
        """
        The common or subject name of the certificate.
        """
        return pulumi.get(self, "certificate_common_name")

    @certificate_common_name.setter
    def certificate_common_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "certificate_common_name", value)

    @property
    @pulumi.getter(name="certificateIssuerThumbprint")
    def certificate_issuer_thumbprint(self) -> Optional[pulumi.Input[str]]:
        """
        The Issuer Thumbprint of the Certificate.
        """
        return pulumi.get(self, "certificate_issuer_thumbprint")

    @certificate_issuer_thumbprint.setter
    def certificate_issuer_thumbprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_issuer_thumbprint", value)


@pulumi.input_type
class ClusterClientCertificateCommonNameArgs:
    def __init__(__self__, *,
                 common_name: pulumi.Input[str],
                 is_admin: pulumi.Input[bool],
                 issuer_thumbprint: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] common_name: The common or subject name of the certificate.
        :param pulumi.Input[bool] is_admin: Does the Client Certificate have Admin Access to the cluster? Non-admin clients can only perform read only operations on the cluster.
        """
        pulumi.set(__self__, "common_name", common_name)
        pulumi.set(__self__, "is_admin", is_admin)
        if issuer_thumbprint is not None:
            pulumi.set(__self__, "issuer_thumbprint", issuer_thumbprint)

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> pulumi.Input[str]:
        """
        The common or subject name of the certificate.
        """
        return pulumi.get(self, "common_name")

    @common_name.setter
    def common_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "common_name", value)

    @property
    @pulumi.getter(name="isAdmin")
    def is_admin(self) -> pulumi.Input[bool]:
        """
        Does the Client Certificate have Admin Access to the cluster? Non-admin clients can only perform read only operations on the cluster.
        """
        return pulumi.get(self, "is_admin")

    @is_admin.setter
    def is_admin(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_admin", value)

    @property
    @pulumi.getter(name="issuerThumbprint")
    def issuer_thumbprint(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "issuer_thumbprint")

    @issuer_thumbprint.setter
    def issuer_thumbprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuer_thumbprint", value)


@pulumi.input_type
class ClusterClientCertificateThumbprintArgs:
    def __init__(__self__, *,
                 is_admin: pulumi.Input[bool],
                 thumbprint: pulumi.Input[str]):
        """
        :param pulumi.Input[bool] is_admin: Does the Client Certificate have Admin Access to the cluster? Non-admin clients can only perform read only operations on the cluster.
        :param pulumi.Input[str] thumbprint: The Thumbprint associated with the Client Certificate.
        """
        pulumi.set(__self__, "is_admin", is_admin)
        pulumi.set(__self__, "thumbprint", thumbprint)

    @property
    @pulumi.getter(name="isAdmin")
    def is_admin(self) -> pulumi.Input[bool]:
        """
        Does the Client Certificate have Admin Access to the cluster? Non-admin clients can only perform read only operations on the cluster.
        """
        return pulumi.get(self, "is_admin")

    @is_admin.setter
    def is_admin(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_admin", value)

    @property
    @pulumi.getter
    def thumbprint(self) -> pulumi.Input[str]:
        """
        The Thumbprint associated with the Client Certificate.
        """
        return pulumi.get(self, "thumbprint")

    @thumbprint.setter
    def thumbprint(self, value: pulumi.Input[str]):
        pulumi.set(self, "thumbprint", value)


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
        pulumi.set(__self__, "blob_endpoint", blob_endpoint)
        pulumi.set(__self__, "protected_account_key_name", protected_account_key_name)
        pulumi.set(__self__, "queue_endpoint", queue_endpoint)
        pulumi.set(__self__, "storage_account_name", storage_account_name)
        pulumi.set(__self__, "table_endpoint", table_endpoint)

    @property
    @pulumi.getter(name="blobEndpoint")
    def blob_endpoint(self) -> pulumi.Input[str]:
        """
        The Blob Endpoint of the Storage Account.
        """
        return pulumi.get(self, "blob_endpoint")

    @blob_endpoint.setter
    def blob_endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "blob_endpoint", value)

    @property
    @pulumi.getter(name="protectedAccountKeyName")
    def protected_account_key_name(self) -> pulumi.Input[str]:
        """
        The protected diagnostics storage key name, such as `StorageAccountKey1`.
        """
        return pulumi.get(self, "protected_account_key_name")

    @protected_account_key_name.setter
    def protected_account_key_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "protected_account_key_name", value)

    @property
    @pulumi.getter(name="queueEndpoint")
    def queue_endpoint(self) -> pulumi.Input[str]:
        """
        The Queue Endpoint of the Storage Account.
        """
        return pulumi.get(self, "queue_endpoint")

    @queue_endpoint.setter
    def queue_endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "queue_endpoint", value)

    @property
    @pulumi.getter(name="storageAccountName")
    def storage_account_name(self) -> pulumi.Input[str]:
        """
        The name of the Storage Account where the Diagnostics should be sent to.
        """
        return pulumi.get(self, "storage_account_name")

    @storage_account_name.setter
    def storage_account_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_account_name", value)

    @property
    @pulumi.getter(name="tableEndpoint")
    def table_endpoint(self) -> pulumi.Input[str]:
        """
        The Table Endpoint of the Storage Account.
        """
        return pulumi.get(self, "table_endpoint")

    @table_endpoint.setter
    def table_endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "table_endpoint", value)


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
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the Fabric Setting, such as `Security` or `Federation`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map containing settings for the specified Fabric Setting.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "parameters", value)


@pulumi.input_type
class ClusterNodeTypeArgs:
    def __init__(__self__, *,
                 client_endpoint_port: pulumi.Input[int],
                 http_endpoint_port: pulumi.Input[int],
                 instance_count: pulumi.Input[int],
                 is_primary: pulumi.Input[bool],
                 name: pulumi.Input[str],
                 application_ports: Optional[pulumi.Input['ClusterNodeTypeApplicationPortsArgs']] = None,
                 capacities: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 durability_level: Optional[pulumi.Input[str]] = None,
                 ephemeral_ports: Optional[pulumi.Input['ClusterNodeTypeEphemeralPortsArgs']] = None,
                 placement_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 reverse_proxy_endpoint_port: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] client_endpoint_port: The Port used for the Client Endpoint for this Node Type. Changing this forces a new resource to be created.
        :param pulumi.Input[int] http_endpoint_port: The Port used for the HTTP Endpoint for this Node Type. Changing this forces a new resource to be created.
        :param pulumi.Input[int] instance_count: The number of nodes for this Node Type.
        :param pulumi.Input[bool] is_primary: Is this the Primary Node Type? Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Node Type. Changing this forces a new resource to be created.
        :param pulumi.Input['ClusterNodeTypeApplicationPortsArgs'] application_ports: A `application_ports` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] capacities: The capacity tags applied to the nodes in the node type, the cluster resource manager uses these tags to understand how much resource a node has.
        :param pulumi.Input[str] durability_level: The Durability Level for this Node Type. Possible values include `Bronze`, `Gold` and `Silver`. Defaults to `Bronze`. Changing this forces a new resource to be created.
        :param pulumi.Input['ClusterNodeTypeEphemeralPortsArgs'] ephemeral_ports: A `ephemeral_ports` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] placement_properties: The placement tags applied to nodes in the node type, which can be used to indicate where certain services (workload) should run.
        :param pulumi.Input[int] reverse_proxy_endpoint_port: The Port used for the Reverse Proxy Endpoint  for this Node Type. Changing this will upgrade the cluster.
        """
        pulumi.set(__self__, "client_endpoint_port", client_endpoint_port)
        pulumi.set(__self__, "http_endpoint_port", http_endpoint_port)
        pulumi.set(__self__, "instance_count", instance_count)
        pulumi.set(__self__, "is_primary", is_primary)
        pulumi.set(__self__, "name", name)
        if application_ports is not None:
            pulumi.set(__self__, "application_ports", application_ports)
        if capacities is not None:
            pulumi.set(__self__, "capacities", capacities)
        if durability_level is not None:
            pulumi.set(__self__, "durability_level", durability_level)
        if ephemeral_ports is not None:
            pulumi.set(__self__, "ephemeral_ports", ephemeral_ports)
        if placement_properties is not None:
            pulumi.set(__self__, "placement_properties", placement_properties)
        if reverse_proxy_endpoint_port is not None:
            pulumi.set(__self__, "reverse_proxy_endpoint_port", reverse_proxy_endpoint_port)

    @property
    @pulumi.getter(name="clientEndpointPort")
    def client_endpoint_port(self) -> pulumi.Input[int]:
        """
        The Port used for the Client Endpoint for this Node Type. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "client_endpoint_port")

    @client_endpoint_port.setter
    def client_endpoint_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "client_endpoint_port", value)

    @property
    @pulumi.getter(name="httpEndpointPort")
    def http_endpoint_port(self) -> pulumi.Input[int]:
        """
        The Port used for the HTTP Endpoint for this Node Type. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "http_endpoint_port")

    @http_endpoint_port.setter
    def http_endpoint_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "http_endpoint_port", value)

    @property
    @pulumi.getter(name="instanceCount")
    def instance_count(self) -> pulumi.Input[int]:
        """
        The number of nodes for this Node Type.
        """
        return pulumi.get(self, "instance_count")

    @instance_count.setter
    def instance_count(self, value: pulumi.Input[int]):
        pulumi.set(self, "instance_count", value)

    @property
    @pulumi.getter(name="isPrimary")
    def is_primary(self) -> pulumi.Input[bool]:
        """
        Is this the Primary Node Type? Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "is_primary")

    @is_primary.setter
    def is_primary(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_primary", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the Node Type. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="applicationPorts")
    def application_ports(self) -> Optional[pulumi.Input['ClusterNodeTypeApplicationPortsArgs']]:
        """
        A `application_ports` block as defined below.
        """
        return pulumi.get(self, "application_ports")

    @application_ports.setter
    def application_ports(self, value: Optional[pulumi.Input['ClusterNodeTypeApplicationPortsArgs']]):
        pulumi.set(self, "application_ports", value)

    @property
    @pulumi.getter
    def capacities(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The capacity tags applied to the nodes in the node type, the cluster resource manager uses these tags to understand how much resource a node has.
        """
        return pulumi.get(self, "capacities")

    @capacities.setter
    def capacities(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "capacities", value)

    @property
    @pulumi.getter(name="durabilityLevel")
    def durability_level(self) -> Optional[pulumi.Input[str]]:
        """
        The Durability Level for this Node Type. Possible values include `Bronze`, `Gold` and `Silver`. Defaults to `Bronze`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "durability_level")

    @durability_level.setter
    def durability_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "durability_level", value)

    @property
    @pulumi.getter(name="ephemeralPorts")
    def ephemeral_ports(self) -> Optional[pulumi.Input['ClusterNodeTypeEphemeralPortsArgs']]:
        """
        A `ephemeral_ports` block as defined below.
        """
        return pulumi.get(self, "ephemeral_ports")

    @ephemeral_ports.setter
    def ephemeral_ports(self, value: Optional[pulumi.Input['ClusterNodeTypeEphemeralPortsArgs']]):
        pulumi.set(self, "ephemeral_ports", value)

    @property
    @pulumi.getter(name="placementProperties")
    def placement_properties(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The placement tags applied to nodes in the node type, which can be used to indicate where certain services (workload) should run.
        """
        return pulumi.get(self, "placement_properties")

    @placement_properties.setter
    def placement_properties(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "placement_properties", value)

    @property
    @pulumi.getter(name="reverseProxyEndpointPort")
    def reverse_proxy_endpoint_port(self) -> Optional[pulumi.Input[int]]:
        """
        The Port used for the Reverse Proxy Endpoint  for this Node Type. Changing this will upgrade the cluster.
        """
        return pulumi.get(self, "reverse_proxy_endpoint_port")

    @reverse_proxy_endpoint_port.setter
    def reverse_proxy_endpoint_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "reverse_proxy_endpoint_port", value)


@pulumi.input_type
class ClusterNodeTypeApplicationPortsArgs:
    def __init__(__self__, *,
                 end_port: pulumi.Input[int],
                 start_port: pulumi.Input[int]):
        """
        :param pulumi.Input[int] end_port: The end of the Application Port Range on this Node Type.
        :param pulumi.Input[int] start_port: The start of the Application Port Range on this Node Type.
        """
        pulumi.set(__self__, "end_port", end_port)
        pulumi.set(__self__, "start_port", start_port)

    @property
    @pulumi.getter(name="endPort")
    def end_port(self) -> pulumi.Input[int]:
        """
        The end of the Application Port Range on this Node Type.
        """
        return pulumi.get(self, "end_port")

    @end_port.setter
    def end_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "end_port", value)

    @property
    @pulumi.getter(name="startPort")
    def start_port(self) -> pulumi.Input[int]:
        """
        The start of the Application Port Range on this Node Type.
        """
        return pulumi.get(self, "start_port")

    @start_port.setter
    def start_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "start_port", value)


@pulumi.input_type
class ClusterNodeTypeEphemeralPortsArgs:
    def __init__(__self__, *,
                 end_port: pulumi.Input[int],
                 start_port: pulumi.Input[int]):
        """
        :param pulumi.Input[int] end_port: The end of the Ephemeral Port Range on this Node Type.
        :param pulumi.Input[int] start_port: The start of the Ephemeral Port Range on this Node Type.
        """
        pulumi.set(__self__, "end_port", end_port)
        pulumi.set(__self__, "start_port", start_port)

    @property
    @pulumi.getter(name="endPort")
    def end_port(self) -> pulumi.Input[int]:
        """
        The end of the Ephemeral Port Range on this Node Type.
        """
        return pulumi.get(self, "end_port")

    @end_port.setter
    def end_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "end_port", value)

    @property
    @pulumi.getter(name="startPort")
    def start_port(self) -> pulumi.Input[int]:
        """
        The start of the Ephemeral Port Range on this Node Type.
        """
        return pulumi.get(self, "start_port")

    @start_port.setter
    def start_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "start_port", value)


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
        pulumi.set(__self__, "x509_store_name", x509_store_name)
        if thumbprint_secondary is not None:
            pulumi.set(__self__, "thumbprint_secondary", thumbprint_secondary)

    @property
    @pulumi.getter
    def thumbprint(self) -> pulumi.Input[str]:
        """
        The Thumbprint of the Certificate.
        """
        return pulumi.get(self, "thumbprint")

    @thumbprint.setter
    def thumbprint(self, value: pulumi.Input[str]):
        pulumi.set(self, "thumbprint", value)

    @property
    @pulumi.getter(name="x509StoreName")
    def x509_store_name(self) -> pulumi.Input[str]:
        """
        The X509 Store where the Certificate Exists, such as `My`.
        """
        return pulumi.get(self, "x509_store_name")

    @x509_store_name.setter
    def x509_store_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "x509_store_name", value)

    @property
    @pulumi.getter(name="thumbprintSecondary")
    def thumbprint_secondary(self) -> Optional[pulumi.Input[str]]:
        """
        The Secondary Thumbprint of the Certificate.
        """
        return pulumi.get(self, "thumbprint_secondary")

    @thumbprint_secondary.setter
    def thumbprint_secondary(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "thumbprint_secondary", value)



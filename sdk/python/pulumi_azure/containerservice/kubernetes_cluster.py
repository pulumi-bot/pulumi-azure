# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['KubernetesCluster']


class KubernetesCluster(pulumi.CustomResource):
    addon_profile: pulumi.Output['outputs.KubernetesClusterAddonProfile'] = pulumi.property("addonProfile")
    """
    A `addon_profile` block as defined below.
    """

    api_server_authorized_ip_ranges: pulumi.Output[Optional[List[str]]] = pulumi.property("apiServerAuthorizedIpRanges")
    """
    The IP ranges to whitelist for incoming traffic to the masters.
    """

    auto_scaler_profile: pulumi.Output['outputs.KubernetesClusterAutoScalerProfile'] = pulumi.property("autoScalerProfile")
    """
    A `auto_scaler_profile` block as defined below.
    """

    default_node_pool: pulumi.Output['outputs.KubernetesClusterDefaultNodePool'] = pulumi.property("defaultNodePool")
    """
    A `default_node_pool` block as defined below.
    """

    disk_encryption_set_id: pulumi.Output[Optional[str]] = pulumi.property("diskEncryptionSetId")
    """
    The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
    """

    dns_prefix: pulumi.Output[str] = pulumi.property("dnsPrefix")
    """
    DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
    """

    enable_pod_security_policy: pulumi.Output[Optional[bool]] = pulumi.property("enablePodSecurityPolicy")
    """
    Whether Pod Security Policies are enabled. Note that this also requires role based access control to be enabled.
    """

    fqdn: pulumi.Output[str] = pulumi.property("fqdn")
    """
    The FQDN of the Azure Kubernetes Managed Cluster.
    """

    identity: pulumi.Output[Optional['outputs.KubernetesClusterIdentity']] = pulumi.property("identity")
    """
    A `identity` block as defined below. Changing this forces a new resource to be created.
    """

    kube_admin_config_raw: pulumi.Output[str] = pulumi.property("kubeAdminConfigRaw")
    """
    Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled.
    """

    kube_admin_configs: pulumi.Output[List['outputs.KubernetesClusterKubeAdminConfig']] = pulumi.property("kubeAdminConfigs")
    """
    A `kube_admin_config` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled.
    """

    kube_config_raw: pulumi.Output[str] = pulumi.property("kubeConfigRaw")
    """
    Raw Kubernetes config to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools
    """

    kube_configs: pulumi.Output[List['outputs.KubernetesClusterKubeConfig']] = pulumi.property("kubeConfigs")
    """
    A `kube_config` block as defined below.
    """

    kubelet_identities: pulumi.Output[List['outputs.KubernetesClusterKubeletIdentity']] = pulumi.property("kubeletIdentities")
    """
    A `kubelet_identity` block as defined below.
    """

    kubernetes_version: pulumi.Output[str] = pulumi.property("kubernetesVersion")
    """
    Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
    """

    linux_profile: pulumi.Output[Optional['outputs.KubernetesClusterLinuxProfile']] = pulumi.property("linuxProfile")
    """
    A `linux_profile` block as defined below.
    """

    location: pulumi.Output[str] = pulumi.property("location")
    """
    The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
    """

    network_profile: pulumi.Output['outputs.KubernetesClusterNetworkProfile'] = pulumi.property("networkProfile")
    """
    A `network_profile` block as defined below.
    """

    node_resource_group: pulumi.Output[str] = pulumi.property("nodeResourceGroup")
    """
    The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
    """

    private_cluster_enabled: pulumi.Output[bool] = pulumi.property("privateClusterEnabled")
    """
    Should this Kubernetes Cluster have it's API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
    """

    private_fqdn: pulumi.Output[str] = pulumi.property("privateFqdn")
    """
    The FQDN for the Kubernetes Cluster when private link has been enabled, which is only resolvable inside the Virtual Network used by the Kubernetes Cluster.
    """

    private_link_enabled: pulumi.Output[bool] = pulumi.property("privateLinkEnabled")

    resource_group_name: pulumi.Output[str] = pulumi.property("resourceGroupName")
    """
    Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
    """

    role_based_access_control: pulumi.Output['outputs.KubernetesClusterRoleBasedAccessControl'] = pulumi.property("roleBasedAccessControl")
    """
    A `role_based_access_control` block. Changing this forces a new resource to be created.
    """

    service_principal: pulumi.Output[Optional['outputs.KubernetesClusterServicePrincipal']] = pulumi.property("servicePrincipal")
    """
    A `service_principal` block as documented below.
    """

    sku_tier: pulumi.Output[Optional[str]] = pulumi.property("skuTier")
    """
    The SKU Tier that should be used for this Kubernetes Cluster. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
    """

    tags: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("tags")
    """
    A mapping of tags to assign to the resource.
    """

    windows_profile: pulumi.Output['outputs.KubernetesClusterWindowsProfile'] = pulumi.property("windowsProfile")
    """
    A `windows_profile` block as defined below.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 addon_profile: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterAddonProfileArgs']]] = None,
                 api_server_authorized_ip_ranges: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 auto_scaler_profile: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterAutoScalerProfileArgs']]] = None,
                 default_node_pool: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterDefaultNodePoolArgs']]] = None,
                 disk_encryption_set_id: Optional[pulumi.Input[str]] = None,
                 dns_prefix: Optional[pulumi.Input[str]] = None,
                 enable_pod_security_policy: Optional[pulumi.Input[bool]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterIdentityArgs']]] = None,
                 kubernetes_version: Optional[pulumi.Input[str]] = None,
                 linux_profile: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterLinuxProfileArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_profile: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterNetworkProfileArgs']]] = None,
                 node_resource_group: Optional[pulumi.Input[str]] = None,
                 private_cluster_enabled: Optional[pulumi.Input[bool]] = None,
                 private_link_enabled: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 role_based_access_control: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterRoleBasedAccessControlArgs']]] = None,
                 service_principal: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterServicePrincipalArgs']]] = None,
                 sku_tier: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 windows_profile: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterWindowsProfileArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Managed Kubernetes Cluster (also known as AKS / Azure Kubernetes Service)

        ## Example Usage

        This example provisions a basic Managed Kubernetes Cluster.

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_kubernetes_cluster = azure.containerservice.KubernetesCluster("exampleKubernetesCluster",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            dns_prefix="exampleaks1",
            default_node_pool={
                "name": "default",
                "node_count": 1,
                "vm_size": "Standard_D2_v2",
            },
            identity={
                "type": "SystemAssigned",
            },
            tags={
                "Environment": "Production",
            })
        pulumi.export("clientCertificate", example_kubernetes_cluster.kube_configs[0]["clientCertificate"])
        pulumi.export("kubeConfig", example_kubernetes_cluster.kube_config_raw)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterAddonProfileArgs']] addon_profile: A `addon_profile` block as defined below.
        :param pulumi.Input[List[pulumi.Input[str]]] api_server_authorized_ip_ranges: The IP ranges to whitelist for incoming traffic to the masters.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterAutoScalerProfileArgs']] auto_scaler_profile: A `auto_scaler_profile` block as defined below.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterDefaultNodePoolArgs']] default_node_pool: A `default_node_pool` block as defined below.
        :param pulumi.Input[str] disk_encryption_set_id: The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
        :param pulumi.Input[str] dns_prefix: DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] enable_pod_security_policy: Whether Pod Security Policies are enabled. Note that this also requires role based access control to be enabled.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterIdentityArgs']] identity: A `identity` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[str] kubernetes_version: Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
        :param pulumi.Input[pulumi.InputType['KubernetesClusterLinuxProfileArgs']] linux_profile: A `linux_profile` block as defined below.
        :param pulumi.Input[str] location: The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterNetworkProfileArgs']] network_profile: A `network_profile` block as defined below.
        :param pulumi.Input[str] node_resource_group: The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] private_cluster_enabled: Should this Kubernetes Cluster have it's API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterRoleBasedAccessControlArgs']] role_based_access_control: A `role_based_access_control` block. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterServicePrincipalArgs']] service_principal: A `service_principal` block as documented below.
        :param pulumi.Input[str] sku_tier: The SKU Tier that should be used for this Kubernetes Cluster. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterWindowsProfileArgs']] windows_profile: A `windows_profile` block as defined below.
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

            __props__['addon_profile'] = addon_profile
            __props__['api_server_authorized_ip_ranges'] = api_server_authorized_ip_ranges
            __props__['auto_scaler_profile'] = auto_scaler_profile
            if default_node_pool is None:
                raise TypeError("Missing required property 'default_node_pool'")
            __props__['default_node_pool'] = default_node_pool
            __props__['disk_encryption_set_id'] = disk_encryption_set_id
            if dns_prefix is None:
                raise TypeError("Missing required property 'dns_prefix'")
            __props__['dns_prefix'] = dns_prefix
            __props__['enable_pod_security_policy'] = enable_pod_security_policy
            __props__['identity'] = identity
            __props__['kubernetes_version'] = kubernetes_version
            __props__['linux_profile'] = linux_profile
            __props__['location'] = location
            __props__['name'] = name
            __props__['network_profile'] = network_profile
            __props__['node_resource_group'] = node_resource_group
            __props__['private_cluster_enabled'] = private_cluster_enabled
            if private_link_enabled is not None:
                warnings.warn("Deprecated in favor of `private_cluster_enabled`", DeprecationWarning)
                pulumi.log.warn("private_link_enabled is deprecated: Deprecated in favor of `private_cluster_enabled`")
            __props__['private_link_enabled'] = private_link_enabled
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['role_based_access_control'] = role_based_access_control
            __props__['service_principal'] = service_principal
            __props__['sku_tier'] = sku_tier
            __props__['tags'] = tags
            __props__['windows_profile'] = windows_profile
            __props__['fqdn'] = None
            __props__['kube_admin_config_raw'] = None
            __props__['kube_admin_configs'] = None
            __props__['kube_config_raw'] = None
            __props__['kube_configs'] = None
            __props__['kubelet_identities'] = None
            __props__['private_fqdn'] = None
        super(KubernetesCluster, __self__).__init__(
            'azure:containerservice/kubernetesCluster:KubernetesCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            addon_profile: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterAddonProfileArgs']]] = None,
            api_server_authorized_ip_ranges: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            auto_scaler_profile: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterAutoScalerProfileArgs']]] = None,
            default_node_pool: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterDefaultNodePoolArgs']]] = None,
            disk_encryption_set_id: Optional[pulumi.Input[str]] = None,
            dns_prefix: Optional[pulumi.Input[str]] = None,
            enable_pod_security_policy: Optional[pulumi.Input[bool]] = None,
            fqdn: Optional[pulumi.Input[str]] = None,
            identity: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterIdentityArgs']]] = None,
            kube_admin_config_raw: Optional[pulumi.Input[str]] = None,
            kube_admin_configs: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['KubernetesClusterKubeAdminConfigArgs']]]]] = None,
            kube_config_raw: Optional[pulumi.Input[str]] = None,
            kube_configs: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['KubernetesClusterKubeConfigArgs']]]]] = None,
            kubelet_identities: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['KubernetesClusterKubeletIdentityArgs']]]]] = None,
            kubernetes_version: Optional[pulumi.Input[str]] = None,
            linux_profile: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterLinuxProfileArgs']]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_profile: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterNetworkProfileArgs']]] = None,
            node_resource_group: Optional[pulumi.Input[str]] = None,
            private_cluster_enabled: Optional[pulumi.Input[bool]] = None,
            private_fqdn: Optional[pulumi.Input[str]] = None,
            private_link_enabled: Optional[pulumi.Input[bool]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            role_based_access_control: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterRoleBasedAccessControlArgs']]] = None,
            service_principal: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterServicePrincipalArgs']]] = None,
            sku_tier: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            windows_profile: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterWindowsProfileArgs']]] = None) -> 'KubernetesCluster':
        """
        Get an existing KubernetesCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterAddonProfileArgs']] addon_profile: A `addon_profile` block as defined below.
        :param pulumi.Input[List[pulumi.Input[str]]] api_server_authorized_ip_ranges: The IP ranges to whitelist for incoming traffic to the masters.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterAutoScalerProfileArgs']] auto_scaler_profile: A `auto_scaler_profile` block as defined below.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterDefaultNodePoolArgs']] default_node_pool: A `default_node_pool` block as defined below.
        :param pulumi.Input[str] disk_encryption_set_id: The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
        :param pulumi.Input[str] dns_prefix: DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] enable_pod_security_policy: Whether Pod Security Policies are enabled. Note that this also requires role based access control to be enabled.
        :param pulumi.Input[str] fqdn: The FQDN of the Azure Kubernetes Managed Cluster.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterIdentityArgs']] identity: A `identity` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[str] kube_admin_config_raw: Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['KubernetesClusterKubeAdminConfigArgs']]]] kube_admin_configs: A `kube_admin_config` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled.
        :param pulumi.Input[str] kube_config_raw: Raw Kubernetes config to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['KubernetesClusterKubeConfigArgs']]]] kube_configs: A `kube_config` block as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['KubernetesClusterKubeletIdentityArgs']]]] kubelet_identities: A `kubelet_identity` block as defined below.
        :param pulumi.Input[str] kubernetes_version: Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
        :param pulumi.Input[pulumi.InputType['KubernetesClusterLinuxProfileArgs']] linux_profile: A `linux_profile` block as defined below.
        :param pulumi.Input[str] location: The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterNetworkProfileArgs']] network_profile: A `network_profile` block as defined below.
        :param pulumi.Input[str] node_resource_group: The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] private_cluster_enabled: Should this Kubernetes Cluster have it's API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] private_fqdn: The FQDN for the Kubernetes Cluster when private link has been enabled, which is only resolvable inside the Virtual Network used by the Kubernetes Cluster.
        :param pulumi.Input[str] resource_group_name: Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterRoleBasedAccessControlArgs']] role_based_access_control: A `role_based_access_control` block. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterServicePrincipalArgs']] service_principal: A `service_principal` block as documented below.
        :param pulumi.Input[str] sku_tier: The SKU Tier that should be used for this Kubernetes Cluster. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterWindowsProfileArgs']] windows_profile: A `windows_profile` block as defined below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["addon_profile"] = addon_profile
        __props__["api_server_authorized_ip_ranges"] = api_server_authorized_ip_ranges
        __props__["auto_scaler_profile"] = auto_scaler_profile
        __props__["default_node_pool"] = default_node_pool
        __props__["disk_encryption_set_id"] = disk_encryption_set_id
        __props__["dns_prefix"] = dns_prefix
        __props__["enable_pod_security_policy"] = enable_pod_security_policy
        __props__["fqdn"] = fqdn
        __props__["identity"] = identity
        __props__["kube_admin_config_raw"] = kube_admin_config_raw
        __props__["kube_admin_configs"] = kube_admin_configs
        __props__["kube_config_raw"] = kube_config_raw
        __props__["kube_configs"] = kube_configs
        __props__["kubelet_identities"] = kubelet_identities
        __props__["kubernetes_version"] = kubernetes_version
        __props__["linux_profile"] = linux_profile
        __props__["location"] = location
        __props__["name"] = name
        __props__["network_profile"] = network_profile
        __props__["node_resource_group"] = node_resource_group
        __props__["private_cluster_enabled"] = private_cluster_enabled
        __props__["private_fqdn"] = private_fqdn
        __props__["private_link_enabled"] = private_link_enabled
        __props__["resource_group_name"] = resource_group_name
        __props__["role_based_access_control"] = role_based_access_control
        __props__["service_principal"] = service_principal
        __props__["sku_tier"] = sku_tier
        __props__["tags"] = tags
        __props__["windows_profile"] = windows_profile
        return KubernetesCluster(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

